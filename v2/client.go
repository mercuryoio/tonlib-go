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
	DEFAULT_TIMEOUT = 4.5
	DefaultRetries  = 600 // 600 * 100ms = 1min
	DefaultRetryTimeout = 100 * time.Millisecond
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

// TONResult is used to unmarshal received json strings into
type TONResult struct {
	Data TONResponse
	Raw  []byte
}

// Client is the Telegram TdLib client
type Client struct {
	extraToResponse map[string]*TONResult
	respMapMu       *sync.RWMutex

	wgService *sync.WaitGroup

	receiveMu *sync.Mutex

	serviceMu *sync.RWMutex

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

// NewClient Creates a new instance of TONLib.
func NewClient(tonCnf *TonInitRequest, config Config, timeout int64, clientLogging bool, tonLogging int32) (*Client, error) {
	rand.Seed(time.Now().UnixNano())

	cli := C.tonlib_client_json_create()

	extra := uint64(0)

	client := Client{
		extraToResponse: map[string]*TONResult{},
		respMapMu:       &sync.RWMutex{},
		receiveMu:       &sync.Mutex{},
		serviceMu: &sync.RWMutex{},

		wgService: &sync.WaitGroup{},

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
		return &client, err
	}

	go client.receiveWorker()
	optionsInfo, err := client.Init(tonCnf.Options)
	if err != nil {
		return &client, err
	}
	if optionsInfo.tonCommon.Type == "options.info" {
		return &client, nil
	}

	client.stopChan <- struct{}{}
	if optionsInfo.tonCommon.Type == "error" {
		return &client, fmt.Errorf("Error ton client init. Message: %s. ", optionsInfo.tonCommon.Extra)
	}
	return &client, fmt.Errorf("Error ton client init. ")
}

func (client *Client) getNewExtra() string {
	return strconv.FormatUint(atomic.AddUint64(client.uniqExtra, 1), 10)
}

func (client *Client) sendAsync(data interface{}) error {
	req, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if client.clientLogging {
		fmt.Println(fmt.Sprintf("call: %s", string(req)))
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
		fmt.Println("fetch data: ", string(resB))
	}

	data := &TONResult{Data: updateData, Raw: resB}

	client.checkNeedService(data)

	extraRaw, ok := updateData["@extra"]
	if !ok {
		return errors.New(fmt.Sprintf("extra field not found in responce. Skip. Resp: %v", updateData))
	}

	extra, ok := extraRaw.(string)
	if !ok || extra == "" {
		return errors.New(fmt.Sprintf("extra field is emptry in responce. Skip. Resp: %v", updateData))
	}

	client.respMapMu.Lock()
	defer client.respMapMu.Unlock()

	if _, ok = client.extraToResponse[extra]; ok {
		return errors.New(fmt.Sprintf("received duplicate extra in responce. Skip. Extra: %s, Resp: %v", extra, updateData))
	}

	client.extraToResponse[extra] = data

	return nil
}

func (client *Client) checkNeedService(data *TONResult) {
	respType, ok := data.Data["@type"]
	if !ok {
		if client.clientLogging {
			fmt.Println("error. checkNeedService. Resp does not contains type field")
		}

		return
	}

	//client.wgService.Wait()

	client.serviceMu.Lock()
	defer client.serviceMu.Unlock()
	//fmt.Println("wg.Add")
	//client.wgService.Add(1)
	//defer func() {
	//	fmt.Println("wg.Done")
	//	client.wgService.Done()
	//}()

	switch respType.(string) {
	case "updateSendLiteServerQuery":
		_, err := client.onLiteServerQueryResult(data.Data["data"].([]byte), data.Data["id"].(JSONInt64))
		if err != nil {
			fmt.Println(fmt.Sprintf("error OnLiteServerQueryResult. Error: %v", err))
			return
		}
		break

	case "updateSyncState":
		syncResp := struct {
			Type      string    `json:"@type"`
			SyncState SyncState `json:"sync_state"`
		}{}
		err := json.Unmarshal(data.Raw, &syncResp)
		if err != nil {
			if err != nil {
				fmt.Println(fmt.Sprintf("error unmarshal SyncState. Error: %v", err))
				return
			}
		}

		if client.clientLogging {
			fmt.Println("run sync", string(data.Raw))
		}

		err = client.syncNew(syncResp.SyncState)
		if err != nil {
			fmt.Println(fmt.Sprintf("error Sync. Error: %v", err))
			return
		}
		break

	default:
		return
	}

	//client.respMapMu.Lock()
	//defer client.respMapMu.Unlock()
	//
	//delete(client.extraToResponse, extra)
}

func (client *Client) receiveWorker() {
	for {
		//client.wgService.Wait()

		select {
		case <-client.stopChan:
			if client.clientLogging {
				fmt.Println("receiveWorker. Stop")
			}
			return

		default:
			if err := client.receiveOne(); err != nil {
				fmt.Println(err)
			}
			client.respMapMu.RLock()
			fmt.Println(fmt.Sprintf("receiveOne. len:%d", len(client.extraToResponse)))
			client.respMapMu.RUnlock()
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
		//client.wgService.Wait()
		client.serviceMu.RLock()
		client.serviceMu.RUnlock()

		val := client.getStoredResponse(extra)
		if val != nil {
			client.respMapMu.Lock()
			defer client.respMapMu.Unlock()

			delete(client.extraToResponse, extra)

			return val, nil
		}

		//if client.clientLogging {
		//	fmt.Println(fmt.Sprintf("empty response. next attempt for extra: %s. AttemptId: %d", extra, i))
		//}

		time.Sleep(DefaultRetryTimeout)
	}

	return nil, errors.New(fmt.Sprintf("timeout for wait resp for extra: %s", extra))
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
		fmt.Println("call execute setLogVerbosityLevel: ", string(req))
	}
	C.tonlib_client_json_execute(client.client, cs)
	return nil
}

func (client *Client) executeAsynchronously(extra string, data interface{}) (*TONResult, error) {
	//client.wgService.Wait()
	client.serviceMu.RLock()
	client.serviceMu.RUnlock()

	err := client.sendAsync(data)
	if err != nil {
		return nil, err
	}

	resp, err := client.waitResponse(extra, DefaultRetries)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (client *Client) onLiteServerQueryResult(bytes []byte, id JSONInt64) (*Ok, error) {
	extra := client.getNewExtra()
	err := client.sendAsync(
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
execute ton-lib asynchronously
*/
func (client *Client) executeAsynchronouslyOld(data interface{}) (*TONResult, error) {
	req, err := json.Marshal(data)
	if err != nil {
		return &TONResult{}, err
	}
	cs := C.CString(string(req))
	defer C.free(unsafe.Pointer(cs))

	if client.clientLogging {
		fmt.Println(fmt.Sprintf("call: %s", string(req)))
	}
	// send may be async
	C.tonlib_client_json_send(client.client, cs)
	// receive must be sync
	result := C.tonlib_client_json_receive(client.client, DEFAULT_TIMEOUT)

	num := 0
	for result == nil {
		if num >= DefaultRetries {
			return &TONResult{}, fmt.Errorf("Client.executeAsynchronously: exided limit of retries to get json response from TON C`s lib. ")
		}
		time.Sleep(1 * time.Second)
		result = C.tonlib_client_json_receive(client.client, DEFAULT_TIMEOUT)
		num += 1
	}

	var updateData TONResponse
	res := C.GoString(result)
	resB := []byte(res)

	err = json.Unmarshal(resB, &updateData)
	if err != nil {
		return nil, err
	}
	if client.clientLogging {
		fmt.Println("fetch data: ", string(resB))
	}
	if st, ok := updateData["@type"]; ok && st == "updateSendLiteServerQuery" {
		err = json.Unmarshal(resB, &updateData)
		if err == nil {
			_, err = client.OnLiteServerQueryResult(updateData["data"].([]byte), updateData["id"].(JSONInt64))
		}
	}
	if st, ok := updateData["@type"]; ok && st == "updateSyncState" {
		syncResp := struct {
			Type      string    `json:"@type"`
			SyncState SyncState `json:"sync_state"`
		}{}
		err = json.Unmarshal(resB, &syncResp)
		if err != nil {
			return &TONResult{}, err
		}
		if client.clientLogging {
			fmt.Println("run sync", updateData)
		}
		res, err = client.Sync(syncResp.SyncState)
		if err != nil {
			return &TONResult{}, err
		}
		if res != "" {
			updateData = TONResponse{}
			resB := []byte(res)
			err = json.Unmarshal(resB, &updateData)
			if err != nil {
				return &TONResult{}, err
			}
			return &TONResult{Data: updateData, Raw: resB}, err
		}
	}
	return &TONResult{Data: updateData, Raw: resB}, err
}

/**
execute ton-lib synchronously
*/
func (client *Client) executeSynchronously(data interface{}) (*TONResult, error) {
	req, err := json.Marshal(data)
	if err!= nil {
		return nil, err
	}

	cs := C.CString(string(req))
	defer C.free(unsafe.Pointer(cs))

	if client.clientLogging {
		fmt.Println(fmt.Sprintf("call sycroniously: %v", data))
	}

	result := C.tonlib_client_json_execute(client.client, cs)

	var updateData TONResponse
	res := C.GoString(result)
	resB := []byte(res)

	if client.clientLogging {
		fmt.Println(fmt.Sprintf("fetch sycroniously: %s", string(resB)))
	}

	err = json.Unmarshal(resB, &updateData)

	return &TONResult{Data: updateData, Raw: resB}, err
}

func (client *Client) Destroy() {
	if client.clientLogging {
		fmt.Println("destroy client")
	}
	client.stopChan <- struct{}{}
	C.tonlib_client_json_destroy(client.client)
}

func (client *Client) syncNew(syncState SyncState) error {
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

	err := client.sendAsync(data)
	if err != nil {
		return err
	}

	for {
		time.Sleep(1 * time.Second)
		//result, err := client.waitResponse(extra, DefaultRetries*2)
		result := client.execReceive()
		//if err != nil {
		//	return err
		//}

		if result == nil {
			continue
		}

		if client.clientLogging {
			fmt.Println(fmt.Sprintf("sync fetch: %s", string(result)))
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
			return fmt.Errorf("Got an error response from ton: `%s` ", string(result))
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
			// continue updating
			continue
		}

		return nil
	}
}

//sync node`s blocks to current
func (client *Client) Sync(syncState SyncState) (string, error) {
	data := struct {
		Type      string    `json:"@type"`
		SyncState SyncState `json:"sync_state"`
	}{
		Type:      "sync",
		SyncState: syncState,
	}
	req, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	cs := C.CString(string(req))
	defer C.free(unsafe.Pointer(cs))
	if client.clientLogging {
		fmt.Println("call (sync)", string(req))
	}
	C.tonlib_client_json_send(client.client, cs)
	for {
		time.Sleep(1 * time.Second)
		result := C.tonlib_client_json_receive(client.client, DEFAULT_TIMEOUT)
		for result == nil {
			if client.clientLogging {
				fmt.Println("empty response. next attempt")
			}
			time.Sleep(1 * time.Second)
			result = C.tonlib_client_json_receive(client.client, DEFAULT_TIMEOUT)
		}
		syncResp := struct {
			Type      string    `json:"@type"`
			SyncState SyncState `json:"sync_state"`
		}{}
		res := C.GoString(result)
		resB := []byte(res)
		err = json.Unmarshal(resB, &syncResp)
		if client.clientLogging {
			fmt.Println("sync result #1: ", res)
		}
		if err != nil {
			return "", err
		}
		if syncResp.Type == "error" {
			return "", fmt.Errorf("Got an error response from ton: `%s` ", res)
		}
		if syncResp.SyncState.Type == "syncStateDone" {
			result := C.tonlib_client_json_receive(client.client, DEFAULT_TIMEOUT)
			syncResp = struct {
				Type      string    `json:"@type"`
				SyncState SyncState `json:"sync_state"`
			}{}
			res = C.GoString(result)
			resB := []byte(res)
			err = json.Unmarshal(resB, &syncResp)
			if client.clientLogging {
				fmt.Println("sync result #2: ", string(resB))
			}
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
			// continue updating
			continue
		}

		return res, nil
	}
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
		return fmt.Errorf("Error ton client init. Message: %s. ", optionsInfo.tonCommon.Extra)
	}
	return fmt.Errorf("Unexpected client init response. %#v", optionsInfo)
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
