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
	"regexp"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	NodeTimeoutSeconds    = 4.5
	WaitTimeout            = 60 * time.Second
	ForceUnlockSyncTimeout = 10 * time.Second
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
		t = ""
	}

	return t, ok
}

func (o TONResponse) getExtra() (string, bool) {
	extra, ok := o["@extra"].(string)
	if !ok {
		extra = ""
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
	syncInProgress bool

	extraToRespChan map[string]chan *TONResult

	respMapMu *sync.RWMutex
	receiveMu *sync.Mutex
	syncMu    *sync.RWMutex
	syncUtilsMu    *sync.Mutex

	uniqClientID string

	uniqExtra *uint64

	filterLogRe *regexp.Regexp

	stopChan chan struct{}
	syncChan chan *TONResult

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
		syncInProgress:  false,
		extraToRespChan: make(map[string]chan *TONResult),
		respMapMu:       &sync.RWMutex{},
		receiveMu:       &sync.Mutex{},
		syncMu:          &sync.RWMutex{},
		syncUtilsMu: &sync.Mutex{},

		uniqClientID: uniqClientID,

		uniqExtra: &extra,

		filterLogRe: regexp.MustCompile(`"(local_password|mnemonic_password|secret|word_list)":("(\\"|[^"])*"|\[("(\\"|[^"])*"(,"(\\"|[^"])*")*)?\])`),

		stopChan: make(chan struct{}),
		syncChan: make(chan *TONResult),

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
			client.formatLog("", fmt.Sprintf("error set log level. Error:%v", err)))
	}

	go client.receiveWorker()
	go client.syncWorker()

	optionsInfo, err := client.Init(tonCnf.Options)
	if err != nil {
		return nil, errors.New(
			client.formatLog(optionsInfo.tonCommon.Extra, fmt.Sprintf("error init client. Error:%v", err)))
	}
	if optionsInfo.tonCommon.Type == "options.info" {
		return &client, nil
	}

	client.stopChan <- struct{}{}
	close(client.stopChan)
	close(client.syncChan)
	if optionsInfo.tonCommon.Type == "error" {
		return nil, fmt.Errorf(client.formatLog(optionsInfo.tonCommon.Extra, "error client init"))
	}

	return nil, fmt.Errorf(client.formatLog(optionsInfo.tonCommon.Extra, "error NewClient"))
}

func (client *Client) formatLog(extra, logStr string) string {
	logStr = client.filterLogRe.ReplaceAllString(logStr, "\"$1\":\"hided\"")

	return fmt.Sprintf("clientID:%s, extra:%s, log:%s", client.uniqClientID, extra, logStr)
}

func (client *Client) printLog(extra, logStr string) {
	if client.clientLogging {
		fmt.Println(client.formatLog(extra, logStr))
	}
}

func (client *Client) getNewExtra() string {
	return strconv.FormatUint(atomic.AddUint64(client.uniqExtra, 1), 10)
}

func (client *Client) sendAsync(extra string, data interface{}) error {
	req, err := json.Marshal(data)
	if err != nil {
		return err
	}

	client.printLog(extra, fmt.Sprintf("call: %s", string(req)))

	cs := C.CString(string(req))
	defer C.free(unsafe.Pointer(cs))

	// send may be async
	C.tonlib_client_json_send(client.client, cs)

	return nil
}

func (client *Client) execReceive() []byte {
	// receive must be sync
	result := C.tonlib_client_json_receive(client.client, NodeTimeoutSeconds)

	if result == nil {
		return nil
	}

	res := C.GoString(result)
	resB := []byte(res)

	return resB
}

func (client *Client) receiveOne() error {
	// receive and sync must be syncrinious
	client.receiveMu.Lock()
	defer client.receiveMu.Unlock()

	resB := client.execReceive()
	if resB == nil {
		return nil
	}

	var updateData TONResponse
	err := json.Unmarshal(resB, &updateData)
	if err != nil {
		return err
	}

	data := &TONResult{Data: updateData, Raw: resB}

	extra, extraOk := data.Data.getExtra()

	client.printLog(extra, fmt.Sprintf("fetch data: %s", string(resB)))

	respType, ok := data.Data.getType()
	if !ok {
		return errors.New(
			client.formatLog(extra,
				fmt.Sprintf("error. Resp does not contains type field. Resp: %s", string(data.Raw))))
	}

	needService := client.checkNeedService(extra, respType, data)
	if needService {
		return nil
	}

	if !extraOk {
		return errors.New(
			client.formatLog(extra,
				fmt.Sprintf("error. Resp does not contains extra field. Resp: %s", string(data.Raw))))
	}

	ch, err := client.getFromStore(extra)
	if err != nil {
		return err
	}

	ch <- data

	return nil
}

func (client *Client) checkNeedService(extra, respType string, data *TONResult) bool {
	switch respType {
	case "updateSendLiteServerQuery":
		go func() {
			_, err := client.onLiteServerQueryResult(data.Data["data"].([]byte), data.Data["id"].(JSONInt64))
			if err != nil {
				client.printLog(extra, fmt.Sprintf("error OnLiteServerQueryResult. Error: %v", err))
			}
		}()
		return true

	case "updateSyncState":
		go func() {
			client.syncChan <- data
		}()
		return true

	case "ton.blockIdExt":
		return true

	default:
		return false
	}
}

func (client *Client) receiveWorker() {
	for {
		select {
		case <-client.stopChan:
			client.printLog("", "receiveWorker. Stop")
			return

		default:
			if err := client.receiveOne(); err != nil {
				client.printLog("", fmt.Sprintf("error receiveOne. Error: %v", err))
			}

			break
		}
	}
}

func (client *Client) timerSyncProgress(t *time.Timer) {
	<-t.C

	client.syncUnlock()
}

func (client *Client) syncLock() {
	client.syncUtilsMu.Lock()
	defer client.syncUtilsMu.Unlock()

	if !client.syncInProgress {
		client.syncInProgress = true
		client.syncMu.Lock()

		client.printLog("", "sync lock")
	}
}

func (client *Client) syncUnlock() {
	client.syncUtilsMu.Lock()
	defer client.syncUtilsMu.Unlock()

	if client.syncInProgress {
		client.syncInProgress = false
		client.syncMu.Unlock()

		client.printLog("", "sync unlock")
	}
}

func (client *Client) syncWorker() {
	for {
		data, ok := <-client.syncChan
		if !ok {
			client.printLog("", "sync. stop sync worker")
			return
		}

		respType, _ := data.Data.getType()
		if respType != "updateSyncState" {
			client.printLog("", fmt.Sprintf("sync. received wrong type. Skip. Resp:%s", string(data.Raw)))
			continue
		}

		syncResp := struct {
			Type      string    `json:"@type"`
			SyncState SyncState `json:"sync_state"`
		}{}

		err := json.Unmarshal(data.Raw, &syncResp)
		if err != nil {
			client.syncUnlock()

			client.printLog("",
				fmt.Sprintf("sync. Error unmarshal response. Skip. Resp:%s, Error:%v", string(data.Raw), err))

			continue
		}

		if syncResp.SyncState.Type == "syncStateDone" {
			client.syncUnlock()

			client.printLog("", "sync. Received syncStateDone. SyncInProgress=false")

			continue
		}

		if client.syncInProgress {
			client.printLog("", fmt.Sprintf("sync. Sync in progress. Resp: %s", string(data.Raw)))

			continue
		}

		// for sync we must stop send and receive
		extra := client.getNewExtra()
		startSyncData := struct {
			Type      string    `json:"@type"`
			Extra     string    `json:"extra"`
			SyncState SyncState `json:"sync_state"`
		}{
			Type:      "sync",
			Extra:     extra,
			SyncState: syncResp.SyncState,
		}

		client.syncLock()
		t := time.NewTimer(ForceUnlockSyncTimeout)
		go client.timerSyncProgress(t)

		client.printLog(extra, fmt.Sprintf("sync. Start sync. SyncInProgress=true. Req:%+v", startSyncData))

		// resend start sync msg is ok. We will just receive one more ton.blockIdExt message
		if err = client.sendAsync(extra, startSyncData); err != nil {
			client.printLog(extra,
				fmt.Sprintf("sync. Error send start sync. Req:%+v, Error:%v", startSyncData, err))
		}
	}
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

	client.printLog("", fmt.Sprintf("call execute setLogVerbosityLevel: %s", string(req)))

	C.tonlib_client_json_execute(client.client, cs)

	return nil
}

func (client *Client) putToStore(extra string, ch chan *TONResult) {
	client.respMapMu.Lock()
	defer client.respMapMu.Unlock()

	client.extraToRespChan[extra] = ch
}

func (client *Client) getFromStore(extra string) (chan *TONResult, error) {
	client.respMapMu.Lock()
	defer client.respMapMu.Unlock()

	ch, ok := client.extraToRespChan[extra]
	if !ok {
		return nil, errors.New("extra not found in store")
	}

	return ch, nil
}

func (client *Client) deleteFromStore(extra string) {
	client.respMapMu.Lock()
	defer client.respMapMu.Unlock()

	delete(client.extraToRespChan, extra)
}

func (client *Client) executeAsynchronously(extra string, data interface{}) (*TONResult, error) {
	ch := make(chan *TONResult)
	defer close(ch)

	client.putToStore(extra, ch)
	defer client.deleteFromStore(extra)

	client.syncMu.RLock()
	client.syncMu.RUnlock()

	err := client.sendAsync(extra, data)
	if err != nil {
		return nil, errors.New(client.formatLog(extra, fmt.Sprintf("error sendAsync. Error: %v", err)))
	}

	timeout := time.NewTimer(WaitTimeout)
	defer timeout.Stop()

	select {
	case <-timeout.C:
		return nil, errors.New(client.formatLog(extra, "timeout for wait response"))
	case resp := <-ch:
		extraResp, _ := resp.Data.getExtra()
		if extraResp != extra {
			return nil, errors.New(client.formatLog(extra,
				fmt.Sprintf("req extra != resp extra. ReqExtra: %s, RespExtra: %s, Req:%+v, Resp:%+v",
					extra, extraResp, data, resp.Data)))
		}

		return resp, nil
	}
}

/**
execute ton-lib asynchronously
*/
func (client *Client) onLiteServerQueryResult(bytes []byte, id JSONInt64) (*Ok, error) {
	extra := client.getNewExtra()
	result, err := client.executeAsynchronously(extra,
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

	client.printLog("", fmt.Sprintf("call sycroniously: %v", data))

	result := C.tonlib_client_json_execute(client.client, cs)

	var updateData TONResponse
	res := C.GoString(result)
	resB := []byte(res)

	client.printLog("", fmt.Sprintf("fetch sycroniously: %s", string(resB)))

	err = json.Unmarshal(resB, &updateData)

	return &TONResult{Data: updateData, Raw: resB}, err
}

func (client *Client) Destroy() {
	client.printLog("", "destroy client")
	client.stopChan <- struct{}{}
	close(client.stopChan)
	close(client.syncChan)
	C.tonlib_client_json_destroy(client.client)
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
			client.formatLog(optionsInfo.Extra,
				fmt.Sprintf("error update ton connection. Ton client init. Message: %+v", optionsInfo)))
	}

	return fmt.Errorf(
		client.formatLog(optionsInfo.Extra,
			fmt.Sprintf("error update ton connection. Unexpected client init response. %+v", optionsInfo)))
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
