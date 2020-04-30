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
	"fmt"
	"math/rand"
	"sync"
	"time"
	"unsafe"
)

const (
	DEFAULT_TIMEOUT = 4.5
	DefaultRetries  = 10
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
func NewClient(tonCnf *TonInitRequest, config Config, timeout int64, clientLogging bool,
	tonLogging int32,
) (*Client, int64, error) {
	rand.Seed(time.Now().UnixNano())

	client := Client{
		client:        C.tonlib_client_json_create(),
		config:        config,
		timeout:       timeout,
		clientLogging: clientLogging,
		tonLogging:    tonLogging,
		options:       tonCnf.Options,
	}

	// disable ton logs if needed
	err := client.executeSetLogLevel(tonLogging)
	if err != nil {
		return &client, 0, err
	}

	optionsInfo, err := client.Init(tonCnf.Options)
	if err != nil {
		return &client, 0, err
	}
	if optionsInfo.tonCommon.Type == "options.info" {
		return &client, int64(optionsInfo.ConfigInfo.DefaultWalletId), nil
	}
	if optionsInfo.tonCommon.Type == "error" {
		return &client, 0, fmt.Errorf("Error ton client init. Message: %s. ", optionsInfo.tonCommon.Extra)
	}
	return &client, 0, fmt.Errorf("Error ton client init. ")
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

func (client *Client) GenerateRequestID() string {
	return RandSeq(20)
}

func (client *Client) receive(maxRetries int, requestID *string) (*TONResult, error) {
	attemptWaiting := 1 * time.Second

	if maxRetries <= 0 {
		return nil, fmt.Errorf("Client.executeAsynchronously: exceed limit of retries to get json response from TON C`s lib. ")
	}

	received := C.tonlib_client_json_receive(client.client, DEFAULT_TIMEOUT)
	maxRetries -= 1

	for received == nil {
		fmt.Printf("fetch nothing. Wait for: %s and try again\n", attemptWaiting)
		if maxRetries <= 0 {
			return nil, fmt.Errorf("Client.executeAsynchronously: exceed limit of retries to get json response from TON C`s lib. ")
		}
		time.Sleep(attemptWaiting)

		received = C.tonlib_client_json_receive(client.client, DEFAULT_TIMEOUT)
		maxRetries -= 1
	}

	res := C.GoString(received)
	resB := []byte(res)

	if client.clientLogging {
		fmt.Println("fetch data: ", string(resB))
	}

	var response TONResponse
	err := json.Unmarshal(resB, &response)
	if err != nil {
		return nil, err
	}

	result := TONResult{Data: response, Raw: resB}

	if requestID != nil && !client.isExtraMatch(result, *requestID) {
		fmt.Println("fetched data is not fit by @extra. Try to fetch again")
		return client.receive(maxRetries, requestID)
	}

	return &result, nil
}

func (client *Client) receiveAny(maxRetries int) (*TONResult, error) {
	return client.receive(maxRetries, nil)
}

func (client *Client) receiveExactly(maxRetries int, requestID string) (*TONResult, error) {
	return client.receive(maxRetries, &requestID)
}

func (client *Client) send(data interface{}) error {
	req, err := json.Marshal(data)
	if err != nil {
		return err
	}
	cs := C.CString(string(req))
	defer C.free(unsafe.Pointer(cs))

	if client.clientLogging {
		fmt.Println("call", string(req))
	}
	C.tonlib_client_json_send(client.client, cs)

	return nil
}

func (client *Client) isExtraMatch(result TONResult, requestID string) bool {
	extra, ok := result.Data["@extra"]
	if !ok {
		return false
	}

	if extra != requestID {
		return false
	}

	return true
}

func (client *Client) executeAsynchronously(data interface{}, requestID string) (*TONResult, error) {
	return client.executeAsynchronouslyCommon(data, requestID, false)
}

func (client *Client) executeAsynchronouslyExactlyOnce(data interface{}, requestID string) (*TONResult, error) {
	return client.executeAsynchronouslyCommon(data, requestID, true)
}

/**
execute ton-lib asynchronously
*/
func (client *Client) executeAsynchronouslyCommon(data interface{}, requestID string, exactlyOnce bool) (*TONResult, error) {
	if err := client.send(data); err != nil {
		return nil, err
	}

	result, err := client.receiveAny(DefaultRetries)
	if err != nil {
		return nil, err
	}

	resultType, ok := result.Data["@type"]
	if !ok {
		return nil, fmt.Errorf("got response invalid struct: field `@type` is undefined")
	}

	if resultType == "updateSendLiteServerQuery" {
		_, err = client.OnLiteServerQueryResult(result.Data["data"].([]byte), result.Data["id"].(JSONInt64))
		if err != nil {
			return nil, err
		}
		if exactlyOnce {
			return nil, fmt.Errorf("can't send to lib more than once [by getting updateSendLiteServerQuery]")
		}
		return client.executeAsynchronouslyCommon(data, requestID, exactlyOnce)
	}

	if resultType == "updateSyncState" {
		syncResp := struct {
			Type      string    `json:"@type"`
			SyncState SyncState `json:"sync_state"`
		}{}
		err = json.Unmarshal(result.Raw, &syncResp)
		if err != nil {
			return nil, err
		}
		if client.clientLogging {
			fmt.Println("run sync", syncResp)
		}
		resultWhileSync, err := client.Sync(syncResp.SyncState, &requestID)
		if err != nil {
			return nil, err
		}

		if resultWhileSync != nil && client.isExtraMatch(*resultWhileSync, requestID) {
			fmt.Printf("got result while sync: %s\n", resultWhileSync.Raw)
			return resultWhileSync, nil
		}

		if exactlyOnce {
			return nil, fmt.Errorf("can't send to lib more than once [by getting updateSyncState]")
		}
		return client.executeAsynchronouslyCommon(data, requestID, exactlyOnce)
	}

	if !client.isExtraMatch(*result, requestID) {
		fmt.Printf("@extra is not the same: real:%s expected:%s. Skip result: %s\n", result.Data["@extra"], requestID, result.Raw)
		return client.receiveExactly(DefaultRetries, requestID)
	}

	return result, nil
}

/**
execute ton-lib synchronously
*/
func (client *Client) executeSynchronously(data interface{}) (*TONResult, error) {
	req, _ := json.Marshal(data)
	cs := C.CString(string(req))
	defer C.free(unsafe.Pointer(cs))
	result := C.tonlib_client_json_execute(client.client, cs)

	var updateData TONResponse
	res := C.GoString(result)
	resB := []byte(res)
	err := json.Unmarshal(resB, &updateData)
	return &TONResult{Data: updateData, Raw: resB}, err
}

func (client *Client) Destroy() {
	C.tonlib_client_json_destroy(client.client)
}

//sync node`s blocks to current
func (client *Client) Sync(syncState SyncState, requestID *string) (*TONResult, error) {
	data := struct {
		Type      string    `json:"@type"`
		SyncState SyncState `json:"sync_state"`
	}{
		Type:      "sync",
		SyncState: syncState,
	}
	if err := client.send(data); err != nil {
		return nil, err
	}

	var meaningfulResult *TONResult

	syncAttempt := 0
	for {
		syncAttempt += 1

		result, err := client.receiveAny(DefaultRetries)
		if err != nil {
			return nil, err
		}

		syncResp := struct {
			Type      string    `json:"@type"`
			SyncState SyncState `json:"sync_state"`
		}{}
		err = json.Unmarshal(result.Raw, &syncResp)
		if err != nil {
			return nil, err
		}

		if client.clientLogging {
			fmt.Printf("sync result #%d: %s \n", syncAttempt, result.Raw)
		}

		if requestID != nil && client.isExtraMatch(*result, *requestID) {
			if meaningfulResult != nil {
				return nil, fmt.Errorf("Got several fit result while sync. It's unexpected behavior: `%s` ", result.Raw)
			}
			meaningfulResult = result
			fmt.Printf("catch meaningful result while sync: %s \n", result.Raw)
			continue
		}

		if syncResp.Type == "error" {
			return nil, fmt.Errorf("Got an error response from ton: `%s` ", result.Raw)
		}

		if syncResp.SyncState.Type == "syncStateDone" {
			break
		}

		if syncResp.Type == "ton.blockIdExt" {
			break
		}

		if syncResp.Type == "updateSyncState" {
			// continue updating
			continue
		}
	}

	return meaningfulResult, nil
}

// QueryEstimateFees
// sometimes it`s respond with "@type: ok" instead of "query.fees"
// @param id
// @param ignoreChksig
func (client *Client) QueryEstimateFees(id int64, ignoreChksig bool) (*QueryFees, error) {
	requestID := client.GenerateRequestID()
	callData := struct {
		Type         string `json:"@type"`
		Extra        string `json:"@extra"`
		Id           int64  `json:"id"`
		IgnoreChksig bool   `json:"ignore_chksig"`
	}{
		Type:         "query.estimateFees",
		Extra:        requestID,
		Id:           id,
		IgnoreChksig: ignoreChksig,
	}

	type Exit struct {
		Exit bool
		sync.Mutex
	}

	var queryFees QueryFees

	type Resp struct {
		Fee   *QueryFees
		Error error
	}

	resultChan := make(chan Resp, 1)
	ticker := time.NewTimer(time.Duration(client.timeout) * time.Second)
	exit := &Exit{false, sync.Mutex{}}

	go func() {
		for true {
			result, err := client.executeAsynchronously(callData, requestID)
			// if timeout reached - close chan and exit
			exit.Lock()
			if exit.Exit {
				exit.Unlock()
				close(resultChan)
				return
			}
			exit.Unlock()

			if err != nil {
				resultChan <- Resp{nil, err}
				return
			}

			if result.Data["@type"].(string) == "error" {
				resultChan <- Resp{nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])}
				return
			}

			if result.Data["@type"].(string) == "query.fees" {
				err = json.Unmarshal(result.Raw, &queryFees)
				resultChan <- Resp{&queryFees, err}
				return
			}
		}
	}()

	select {
	case _ = <-ticker.C:
		// notify gorutine that performing requests to TON
		exit.Lock()
		exit.Exit = true
		exit.Unlock()

		ticker.Stop()
		return nil, fmt.Errorf("timeout")
	case result := <-resultChan:
		ticker.Stop()
		return result.Fee, result.Error
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

type AccountState CommonAccountState

type MsgData interface{ MessageType() string }
type DnsEntryData string

type Action interface{ MessageType() string }
type DnsAction Action

type CommonAccountState struct {
	tonCommon
	Code       string    `json:"code"`
	Data       string    `json:"data"`
	FrozenHash []byte    `json:"frozen_hash"`
	PublicKey  string    `json:"public_key"`
	WalletId   JSONInt64 `json:"wallet_id"`
	Seqno      JSONInt64 `json:"seqno"`
}
