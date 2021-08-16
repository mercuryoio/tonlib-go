package v2

//#cgo linux CFLAGS: -I./lib/linux
//#cgo darwin CFLAGS: -I./lib/darwin
//#cgo linux LDFLAGS: -L./lib/linux -ltonlibjson -ltonlibjson_private -ltonlibjson_static -ltonlib
//#cgo darwin LDFLAGS: -L./lib/darwin -ltonlibjson -ltonlibjson_private -ltonlibjson_static -ltonlib
//#include <stdlib.h>
//#include <./lib/tonlib_client_json.h>
import "C"
import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	DEFAULT_TIMEOUT     = 4.5
	DefaultRetries      = 600 // 600 * 100ms = 1min
	DefaultRetryTimeout = 100 * time.Millisecond
	SinkLimitTries      = 1000
)

type InputKey struct {
	Type          string        `json:"@type"`
	LocalPassword string        `json:"local_password"`
	Key           TONPrivateKey `json:"key"`
}
type TONPrivateKey struct {
	PublicKey string `json:"public_key"`
	Secret    string `json:"secret"`
}

type SyncState struct {
	Type         string `json:"@type"`
	FromSeqno    int    `json:"from_seqno"`
	ToSeqno      int    `json:"to_seqno"`
	CurrentSeqno int    `json:"current_seqno"`
}

// KeyStoreType directory
type KeyStoreType struct {
	Type      string `json:"@type"`
	Directory string `json:"directory"`
}

// TONResponse alias for use in TONResult
type TONResponse map[string]interface{}

func (o TONResponse) getType() (string, bool) {
	t, ok := o["@type"].(string)
	if !ok {
		t = "undef"
	}

	return t, ok
}

func (o TONResponse) getExtra() (string, bool) {
	extra, ok := o["@extra"].(string)
	if !ok {
		extra = "undef"
	}

	return extra, ok
}

// TONResult is used to unmarshal received json strings into
type TONResult struct {
	Data TONResponse
	Raw  []byte
}

// Client is the Telegram TdLib client
type Client struct {
	extraToResponse map[string]*TONResult

	respMapMu *sync.RWMutex
	receiveMu *sync.Mutex
	serviceMu *sync.RWMutex

	uniqClientID string

	uniqExtra *uint64
	stopChan  chan struct{}

	client        unsafe.Pointer
	config        Config
	timeout       int64
	clientLogging bool
	tonLogging    int32
	options       Options
}

type TonInitRequest struct {
	Type    string  `json:"@type"`
	Options Options `json:"options"`
}

func NewClient(tonCnf *TonInitRequest, config Config, timeout int64,
	clientLogging bool, tonLogging int32) (*Client, error) {
	uniqClientID := strconv.FormatUint(rand.Uint64(), 10)

	return NewClientWithUniqID(tonCnf, config, timeout, clientLogging, tonLogging, uniqClientID)
}

// NewClientWithUniqID Creates a new instance of TONLib.
// @uniqClientID - is not important for executing. Using just for logging for identify client
func NewClientWithUniqID(tonCnf *TonInitRequest, config Config, timeout int64,
	clientLogging bool, tonLogging int32, uniqClientID string) (*Client, error) {
	rand.Seed(time.Now().UnixNano())

	cli := C.tonlib_client_json_create()

	extra := uint64(0)

	client := Client{
		extraToResponse: map[string]*TONResult{},
		respMapMu:       &sync.RWMutex{},
		receiveMu:       &sync.Mutex{},
		serviceMu:       &sync.RWMutex{},

		uniqClientID: uniqClientID,

		uniqExtra: &extra,
		stopChan:  make(chan struct{}),

		client:        cli,
		config:        config,
		timeout:       timeout,
		clientLogging: clientLogging,
		tonLogging:    tonLogging,
		options:       tonCnf.Options,
	}

	// disable ton logs if needed
	err := client.executeSetLogLevel(tonLogging)
	if err != nil {
		return nil, errors.New(
			client.formatLogString("", fmt.Sprintf("error set log level. Error:%v", err)))
	}

	go client.receiveWorker()

	optionsInfo, err := client.Init(tonCnf.Options)
	if err != nil {
		return nil, errors.New(
			client.formatLogString(optionsInfo.tonCommon.Extra, fmt.Sprintf("error init client. Error:%v", err)))
	}
	if optionsInfo.tonCommon.Type == "options.info" {
		return &client, nil
	}

	client.stopChan <- struct{}{}
	if optionsInfo.tonCommon.Type == "error" {
		return nil, fmt.Errorf(client.formatLogString(optionsInfo.tonCommon.Extra, "error client init"))
	}

	return nil, fmt.Errorf(client.formatLogString(optionsInfo.tonCommon.Extra, "error NewClient"))
}

func (client *Client) formatLogString(extra string, logStr string) string {
	return fmt.Sprintf("clientID:%s, extra:%s, log:%s", client.uniqClientID, extra, logStr)
}

func (client *Client) getNewExtra() string {
	return strconv.FormatUint(atomic.AddUint64(client.uniqExtra, 1), 10)
}

func (client *Client) sendAsync(extra string, data interface{}) error {
	req, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if client.clientLogging {
		fmt.Println(client.formatLogString(extra, fmt.Sprintf("call: %s", string(req))))
	}

	cs := C.CString(string(req))
	defer C.free(unsafe.Pointer(cs))

	// send may be async
	C.tonlib_client_json_send(client.client, cs)

	return nil
}

func (client *Client) execReceive() []byte {
	client.receiveMu.Lock()
	defer client.receiveMu.Unlock()

	// receive must be sync
	result := C.tonlib_client_json_receive(client.client, DEFAULT_TIMEOUT)

	if result == nil {
		return nil
	}

	res := C.GoString(result)
	resB := []byte(res)

	return resB
}

func (client *Client) receiveOne() error {
	// there is possible situation that one thread on sync and other thread exec Receive.
	// This situation handled in sync method via check seq_to=0
	// It is more safe to use Lock/defer Unlock but it is slow down app for 2-5 times
	client.serviceMu.RLock()
	client.serviceMu.RUnlock()

	resB := client.execReceive()
	if resB == nil {
		return nil
	}

	var updateData TONResponse
	err := json.Unmarshal(resB, &updateData)
	if err != nil {
		return err
	}

	if client.clientLogging {
		extra, _ := updateData.getExtra()
		fmt.Println(client.formatLogString(extra, fmt.Sprintf("fetch data: %s", string(resB))))
	}

	data := &TONResult{Data: updateData, Raw: resB}

	client.serviceMu.Lock()
	defer client.serviceMu.Unlock()

	err = client.putResponseToStore(data)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) putResponseToStore(data *TONResult) error {
	client.checkNeedService(data)

	extra, ok := data.Data.getExtra()
	if !ok {
		return errors.New(
			client.formatLogString(extra,
				fmt.Sprintf("error put receive to store. Extra field not found in responce. Skip. Resp: %s",
					string(data.Raw))))
	}

	client.respMapMu.Lock()
	defer client.respMapMu.Unlock()

	if _, ok = client.extraToResponse[extra]; ok {
		return errors.New(
			client.formatLogString(extra,
				fmt.Sprintf("error put receive to store. Received duplicate extra in responce. Skip. Resp: %s",
					string(data.Raw))))
	}

	client.extraToResponse[extra] = data

	return nil
}

func (client *Client) checkNeedService(data *TONResult) {
	extra, _ := data.Data.getExtra()
	respType, ok := data.Data.getType()
	if !ok {
		if client.clientLogging {
			fmt.Println(client.formatLogString(extra, "error. checkNeedService. Resp does not contains type field"))
		}

		return
	}

	switch respType {
	case "updateSendLiteServerQuery":
		go func() {
			_, err := client.onLiteServerQueryResult(data.Data["data"].([]byte), data.Data["id"].(JSONInt64))
			if err != nil {
				fmt.Println(
					client.formatLogString(extra, fmt.Sprintf("error OnLiteServerQueryResult. Error: %v", err)))
				return
			}
		}()
		break

	case "updateSyncState":
		syncResp := struct {
			Type      string    `json:"@type"`
			SyncState SyncState `json:"sync_state"`
		}{}
		err := json.Unmarshal(data.Raw, &syncResp)
		if err != nil {
			if err != nil {
				if client.clientLogging {
					fmt.Println(
						client.formatLogString(extra, fmt.Sprintf("error unmarshal SyncState. Error: %v", err)))
				}
				return
			}
		}

		// seq_to=0 node set for new sync request. It is leaked sync msg out of sync method
		if syncResp.SyncState.ToSeqno != 0 {
			if client.clientLogging {
				fmt.Println(
					client.formatLogString(extra,
						fmt.Sprintf("error sync receive seq_to=0. "+
							"Possible sinc msg received out of sync method. Skip. Resp:%s", string(data.Raw))))
			}
			return
		}

		err = client.sync(syncResp.SyncState)
		if err != nil {
			if client.clientLogging {
				fmt.Println(client.formatLogString(extra, fmt.Sprintf("error sync. Error: %v", err)))
			}
			return
		}
		break

	default:
		return
	}
}

func (client *Client) receiveWorker() {
	for {
		select {
		case <-client.stopChan:
			if client.clientLogging {
				fmt.Println(client.formatLogString("", "receiveWorker. Stop"))
			}
			return

		default:
			if err := client.receiveOne(); err != nil {
				if client.clientLogging {
					fmt.Println(client.formatLogString("", fmt.Sprintf("error receiveOne. Error: %v", err)))
				}
			}

			break
		}
	}
}

func (client *Client) getStoredResponse(extra string) *TONResult {
	client.respMapMu.RLock()
	defer client.respMapMu.RUnlock()

	val, ok := client.extraToResponse[extra]
	if !ok {
		return nil
	}

	return val
}

func (client *Client) waitResponse(extra string, retryCount int) (*TONResult, error) {
	for i := 0; i < retryCount; i++ {
		client.serviceMu.RLock()
		client.serviceMu.RUnlock()

		val := client.getStoredResponse(extra)
		if val != nil {
			client.respMapMu.Lock()
			defer client.respMapMu.Unlock()

			delete(client.extraToResponse, extra)

			return val, nil
		}

		time.Sleep(DefaultRetryTimeout)
	}

	return nil, errors.New("timeout for wait resp for extra")
}

// disable ton client C lib`s logs
func (client *Client) executeSetLogLevel(logLevel int32) error {
	data := struct {
		Type              string `json:"@type"`
		NewVerbosityLevel int32  `json:"new_verbosity_level"`
	}{
		Type:              "setLogVerbosityLevel",
		NewVerbosityLevel: logLevel,
	}
	req, err := json.Marshal(data)
	if err != nil {
		return err
	}
	cs := C.CString(string(req))
	defer C.free(unsafe.Pointer(cs))

	if client.clientLogging {
		fmt.Println(
			client.formatLogString("", fmt.Sprintf("call execute setLogVerbosityLevel: %s", string(req))))
	}

	C.tonlib_client_json_execute(client.client, cs)

	return nil
}

func (client *Client) executeAsynchronously(extra string, data interface{}) (*TONResult, error) {
	err := client.sendAsync(extra, data)
	if err != nil {
		return nil, errors.New(client.formatLogString(extra, fmt.Sprintf("error sendAsync. Error: %v", err)))
	}

	resp, err := client.waitResponse(extra, DefaultRetries)
	if err != nil {
		return nil, errors.New(client.formatLogString(extra, fmt.Sprintf("error waitResponse. Error: %v", err)))
	}

	return resp, nil
}

/**
execute ton-lib asynchronously
*/
func (client *Client) onLiteServerQueryResult(bytes []byte, id JSONInt64) (*Ok, error) {
	extra := client.getNewExtra()
	err := client.sendAsync(extra,
		struct {
			Type  string    `json:"@type"`
			Extra string    `json:"@extra"`
			Bytes []byte    `json:"bytes"`
			Id    JSONInt64 `json:"id"`
		}{
			Type:  "onLiteServerQueryResult",
			Bytes: bytes,
			Extra: extra,
			Id:    id,
		},
	)
	if err != nil {
		return nil, err
	}

	result, err := client.waitResponse(extra, DefaultRetries)
	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)

	return &ok, err
}

/**
execute ton-lib synchronously
*/
func (client *Client) executeSynchronously(data interface{}) (*TONResult, error) {
	req, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	cs := C.CString(string(req))
	defer C.free(unsafe.Pointer(cs))

	if client.clientLogging {
		fmt.Println(client.formatLogString("", fmt.Sprintf("call sycroniously: %v", data)))
	}

	result := C.tonlib_client_json_execute(client.client, cs)

	var updateData TONResponse
	res := C.GoString(result)
	resB := []byte(res)

	if client.clientLogging {
		fmt.Println(client.formatLogString("", fmt.Sprintf("fetch sycroniously: %s", string(resB))))
	}

	err = json.Unmarshal(resB, &updateData)

	return &TONResult{Data: updateData, Raw: resB}, err
}

func (client *Client) Destroy() {
	if client.clientLogging {
		fmt.Println(client.formatLogString("", "destroy client"))
	}
	client.stopChan <- struct{}{}
	C.tonlib_client_json_destroy(client.client)
}

// sync node`s blocks to current
func (client *Client) sync(syncState SyncState) (err error) {
	extra := client.getNewExtra()
	data := struct {
		Type      string    `json:"@type"`
		Extra     string    `json:"extra"`
		SyncState SyncState `json:"sync_state"`
	}{
		Type:      "sync",
		Extra:     extra,
		SyncState: syncState,
	}

	req, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if client.clientLogging {
		fmt.Println(client.formatLogString(extra,
			fmt.Sprintf("run sync. call: %s", string(req))))
	}

	cs := C.CString(string(req))
	defer C.free(unsafe.Pointer(cs))

	// send may be async
	C.tonlib_client_json_send(client.client, cs)

	for i := 0; i < SinkLimitTries; i++ {
		result := client.execReceive()
		if err != nil {
			return err
		}

		if result == nil {
			continue
		}

		if client.clientLogging {
			fmt.Println(client.formatLogString(extra,
				fmt.Sprintf("sync iter:%d, fetch: %s", i, string(result))))
		}

		syncResp := struct {
			Type      string    `json:"@type"`
			SyncState SyncState `json:"sync_state"`
		}{}
		err = json.Unmarshal(result, &syncResp)
		if err != nil {
			return err
		}

		if syncResp.Type == "error" {
			return fmt.Errorf(
				client.formatLogString(extra,
					fmt.Sprintf("sync. Got an error response from ton: `%s`", string(result))))
		}

		if syncResp.SyncState.Type == "syncStateDone" {
			return nil
		}
		if syncResp.Type == "ok" {
			// continue updating
			continue
		}
		if syncResp.Type == "ton.blockIdExt" {
			// continue updating
			continue
		}
		if syncResp.Type == "updateSyncState" {
			time.Sleep(100 * time.Millisecond)
			// continue updating
			continue
		}

		i--

		if client.clientLogging {
			fmt.Println(
				client.formatLogString(extra,
					fmt.Sprintf("sync. WARNING received unexpected type. Resp: %s", string(result))))
		}

		var tonResp TONResponse
		err = json.Unmarshal(result, &tonResp)
		if err != nil {
			return err
		}
		err = client.putResponseToStore(&TONResult{Data: tonResp, Raw: result})
		if err != nil {
			return err
		}

		continue
	}

	return errors.New(client.formatLogString(extra, "sync reach limit of tries. Interrupt sync"))
}

// for now - a few requests may works wrong, cause it some times get respose form previos reqest for a few times
func (client *Client) UpdateTonConnection() error {
	_, err := client.Close()
	if err != nil {
		return err
	}
	// destroy old c.ient
	client.Destroy()

	// create new C client
	client.client = C.tonlib_client_json_create()
	// set log level
	err = client.executeSetLogLevel(client.tonLogging)
	if err != nil {
		return err
	}

	// init client
	optionsInfo, err := client.Init(client.options)
	if err != nil {
		return err
	}
	if optionsInfo.tonCommon.Type == "options.info" {
		return nil
	}
	if optionsInfo.tonCommon.Type == "error" {
		return fmt.Errorf(
			client.formatLogString(optionsInfo.Extra,
				fmt.Sprintf("error update ton connection. Ton client init. Message: %s. ",
					optionsInfo.tonCommon.Extra)))
	}

	return fmt.Errorf(
		client.formatLogString(optionsInfo.Extra,
			fmt.Sprintf("error update ton connection. Unexpected client init response. %#v", optionsInfo)))
}

// key struct cause it strings values no bytes
// Key
type Key struct {
	tonCommon
	PublicKey string `json:"public_key"` //
	Secret    string `json:"secret"`     //
}

// MessageType return the string telegram-type of Key
func (key *Key) MessageType() string {
	return "key"
}

// NewKey creates a new Key
//
// @param publicKey
// @param secret
func NewKey(publicKey string, secret string) *Key {
	keyTemp := Key{
		tonCommon: tonCommon{Type: "key"},
		PublicKey: publicKey,
		Secret:    secret,
	}

	return &keyTemp
}

// because of different subclasses in common class InitialAccountState and  AccountState
// InitialAccountState
type InitialAccountState interface{ MessageType() string }

type AccountState RawAccountState

type MsgData interface{}
type DnsEntryData string

type Action interface{ MessageType() string }
type DnsAction Action

type PchanState interface{ MessageType() string }
type PchanAction interface{ MessageType() string }
