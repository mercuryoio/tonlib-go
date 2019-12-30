package v2

//#cgo linux CFLAGS: -I../lib/linux
//#cgo darwin CFLAGS: -I../lib/darwin
//#cgo linux LDFLAGS: -L../lib/linux -ltonlibjson -ltonlibjson_private -ltonlibjson_static -ltonlib
//#cgo darwin LDFLAGS: -L../lib/darwin -ltonlibjson -ltonlibjson_private -ltonlibjson_static -ltonlib
//#include <stdlib.h>
//#include <../lib/tonlib_client_json.h>
import "C"
import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

const (
	DEFAULT_TIMEOUT = 4.5
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
	client unsafe.Pointer
	config Config
	// wallet *TonWallet
}

type TonInitRequest struct {
	Type    string  `json:"@type"`
	Options Options `json:"options"`
}

// NewClient Creates a new instance of TONLib.
func NewClient(tonCnf *TonInitRequest, config Config) (*Client, error) {
	rand.Seed(time.Now().UnixNano())

	client := Client{client: C.tonlib_client_json_create(), config: config}
	optionsInfo, err := client.Init(&tonCnf.Options)
	//resp, err := client.executeAsynchronously(tonCnf)
	if err != nil {
		return &client, err
	}
	//if st, ok := resp.Data["@type"]; ok && st == "ok" {
	//	return &client, nil
	//}
	//if st, ok := resp.Data["@type"]; ok && st == "error" {
	//	return &client, fmt.Errorf("Error ton client init. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	//}
	fmt.Println("Init ton client result: ", optionsInfo, err)
	return &client, fmt.Errorf("Error ton client init. ")
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

	fmt.Println("call", string(req))
	C.tonlib_client_json_send(client.client, cs)
	result := C.tonlib_client_json_receive(client.client, DEFAULT_TIMEOUT)

	for result == nil {
		time.Sleep(1 * time.Second)
		result = C.tonlib_client_json_receive(client.client, DEFAULT_TIMEOUT)
	}

	var updateData TONResponse
	res := C.GoString(result)
	resB := []byte(res)
	err = json.Unmarshal(resB, &updateData)
	fmt.Println("fetch data: ", string(resB))
	if st, ok := updateData["@type"]; ok && st == "updateSendLiteServerQuery" {
		err = json.Unmarshal(resB, &updateData)
		if err == nil {
			_, err = client.Onliteserverqueryresult(updateData["data"].([]byte), updateData["id"].(JSONInt64))
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
		fmt.Println("run sync", updateData)
		_, err = client.Sync()
		if err != nil {
			return &TONResult{}, err
		}
		return client.executeAsynchronously(data)
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
