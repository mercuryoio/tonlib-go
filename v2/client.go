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
	"sort"
	"strconv"
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
func NewClient(tonCnf *TonInitRequest, config Config, timeout int64, clientLogging bool, tonLogging int32) (*Client, error) {
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
		return &client, err
	}

	optionsInfo, err := client.Init(tonCnf.Options)
	if err != nil {
		return &client, err
	}
	if optionsInfo.tonCommon.Type == "options.info" {
		return &client, nil
	}
	if optionsInfo.tonCommon.Type == "error" {
		return &client, fmt.Errorf("Error ton client init. Message: %s. ", optionsInfo.tonCommon.Extra)
	}
	return &client, fmt.Errorf("Error ton client init. ")
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

/**
execute ton-lib asynchronously
*/
func (client *Client) executeAsynchronously(data interface{}) (*TONResult, error) {
	req, err := json.Marshal(data)
	if err != nil {
		return &TONResult{}, err
	}
	cs := C.CString(string(req))
	defer C.free(unsafe.Pointer(cs))

	if client.clientLogging {
		fmt.Println("call", string(req))
	}
	C.tonlib_client_json_send(client.client, cs)
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

// QueryEstimateFees
// sometimes it`s respond with "@type: ok" instead of "query.fees"
// @param id
// @param ignoreChksig
func (client *Client) QueryEstimateFees(id int64, ignoreChksig bool) (*QueryFees, error) {
	callData := struct {
		Type         string `json:"@type"`
		Id           int64  `json:"id"`
		IgnoreChksig bool   `json:"ignore_chksig"`
	}{
		Type:         "query.estimateFees",
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
			result, err := client.executeAsynchronously(callData)
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
func (client *Client) UpdateTonConnection() (error) {
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

type parsedParticipant struct {
	Stake int64
	MaxFactor int64
	Id string
}

func (client *Client) GetElectionStackes(electorAddress string, minStake, maxStake, maxStakeFactor, minValidators int64) (minEffectStake, maxEffectStake int64, err error) {
	participants, err := client.GetParticipantListExtended(electorAddress)
	if err != nil{
		return 0,0 , fmt.Errorf("failed to get participants")
	}

	// parse data
	parsedPartisipants := []parsedParticipant{}
	for _, p := range *participants{
		stake, err := strconv.ParseInt(p.Stake, 10, 64)
		if err != nil{
			return 0,0 , fmt.Errorf("failed to parse participant %v stake. %v", p, err)
		}
		maxFacor, err := strconv.ParseInt(p.Stake, 10, 64)
		if err != nil{
			return 0,0 , fmt.Errorf("failed to parse participant %v maxFactor. %v", p, err)
		}
		parsedPartisipants = append(parsedPartisipants, parsedParticipant{stake, min(maxStakeFactor, maxFacor), p.Id})
	}

	// sort array
	sort.Slice(parsedPartisipants, func(i, j int) bool {
		return parsedPartisipants[i].Stake > parsedPartisipants[j].Stake
	})

	// think that lis already sorted in desc
	if int64(len(parsedPartisipants) + 1) <= minValidators{
		minEffectStake = parsedPartisipants[len(parsedPartisipants)-1].Stake
		minEffectStake = min(minEffectStake, minEffectStake*parsedPartisipants[len(parsedPartisipants)-1].MaxFactor >> 16)

		maxEffectStake = parsedPartisipants[0].Stake
		maxEffectStake = min(maxEffectStake, parsedPartisipants[0].MaxFactor*minEffectStake >> 16)
		return minStake, maxEffectStake, nil
	}

	var bestStake, m, i int64 = 0, 0, 0
	for _, p := range parsedPartisipants{
		i += 1
		if (p.Stake >= minStake) {
			totalStake, err := computeTotalStake(&parsedPartisipants, i, p.Stake)
			if err != nil{
				return 0, 0, fmt.Errorf("failed to  computeTotalStake(participants, i (%d), stake (%d)): %v", i, p.Stake, err)
			}
			if (totalStake > bestStake){
				bestStake = totalStake
				m = i
			}
		}
	}

	// evaulate  minEffectStake
	minEffectStake = parsedPartisipants[m].Stake
	minEffectStake = min(minEffectStake, (parsedPartisipants[m].MaxFactor*minEffectStake) >> 16);

	// evaluate maxEffectStake
	newMinStake := minEffectStake
	if m >1 {
		newMinStake = parsedPartisipants[m-1].Stake
		newMinStake = min(newMinStake, (parsedPartisipants[m-1].MaxFactor*newMinStake) >> 16);
	}
	maxEffectStake = parsedPartisipants[0].Stake

	// real stake may differ cause m number may be extended
	// calc real effective second max effective stake
	maxEffectStake = min(maxEffectStake, (parsedPartisipants[0].MaxFactor*newMinStake) >> 16);
	// calc new max stake
	maxEffectStake = min(maxEffectStake + 1, (maxStakeFactor*newMinStake) >> 16);
	return minEffectStake, maxEffectStake, nil
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
