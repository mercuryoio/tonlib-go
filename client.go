package tonlib

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
	_ "github.com/mercuryoio/tonlib-go/lib"
	"math/rand"
	"strconv"
	"time"
	"unsafe"
)

// Client is the Telegram TdLib client
type Client struct {
	client unsafe.Pointer
	config Config
	wallet *TonWallet
}

// Config holds tonlibParameters
type Config struct {
	Timeout float32
}

// NewClient Creates a new instance of TONLib.
func NewClient(tonCnf *TONInitRequest, config Config) (*Client, error) {
	rand.Seed(time.Now().UnixNano())

	client := Client{client: C.tonlib_client_json_create(), config: combineConfig(config)}
	resp, err := client.executeAsynchronously(tonCnf)
	if err != nil {
		return &client, err
	}
	if st, ok := resp.Data["@type"]; ok && st == "ok" {
		return &client, nil
	}
	if st, ok := resp.Data["@type"]; ok && st == "error" {
		return &client, fmt.Errorf("Error ton client init. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	}
	fmt.Println("Init ton client result: ", string(resp.Raw), err)
	return &client, fmt.Errorf("Error ton client init. ")
}

// InitWallet wallet.init and set it as default wallet for client
func (client *Client) InitWallet(key *TONPrivateKey, password []byte) (err error) {
	st := struct {
		Type       string   `json:"@type"`
		PrivateKey InputKey `json:"private_key"`
	}{
		Type:       "wallet.init",
		PrivateKey: key.getInputKey(password),
	}
	resp, err := client.executeAsynchronously(st)
	if err != nil {
		return err
	}
	if st, ok := resp.Data["@type"]; ok && st == "error" {
		return fmt.Errorf("Error ton client init. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	}
	if st, ok := resp.Data["@type"]; ok && st == "ok" {
		client.wallet = new(TonWallet)
		client.wallet.client = client
		return nil
	}
	return fmt.Errorf("Error ton TonWallet init. ")
}

// WalletGetAddress get wallet address method
func (client *Client) WalletGetAddress(pubKey string) (*TONAccountAddress, error) {
	if client.wallet == nil {
		return nil, fmt.Errorf("You must init wallet before. ")
	}
	return client.wallet.getAddress(pubKey)
}

// WalletState get wallet state
func (client *Client) WalletState(address string) (*TONAccountState, error) {
	if client.wallet == nil {
		return nil, fmt.Errorf("You must init wallet before. ")
	}
	return client.wallet.getState(address)
}

// WalletSendGRAMM2Address send GRAM to address
func (client *Client) WalletSendGRAMM2Address(key *TONPrivateKey, password []byte, fromAddress, toAddress string, amount string) (*TONResult, error) {
	if client.wallet == nil {
		return nil, fmt.Errorf("You must init wallet before. ")
	}
	return client.wallet.sendGRAMM2Address(key, password, fromAddress, toAddress, amount)
}

// UnpackAccountAddress get full address in HEX
func (client *Client) UnpackAccountAddress(address string) (*TONUnpackedAddress, error) {
	st := struct {
		Type           string `json:"@type"`
		AccountAddress string `json:"account_address"`
	}{
		Type:           "unpackAccountAddress",
		AccountAddress: address,
	}
	resp, err := client.executeAsynchronously(st)
	if err != nil {
		return nil, err
	}
	if st, ok := resp.Data["@type"]; ok && st == "error" {
		return nil, fmt.Errorf("Error ton client init. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	}

	addressSt := TONUnpackedAddress{}
	err = json.Unmarshal(resp.Raw, &addressSt)
	if err != nil {
		return nil, err
	}
	return &addressSt, nil
}

// PackAccountAddress get short address from full address HEX
func (client *Client) PackAccountAddress(packedAddr *TONUnpackedAddress, address string) (string, error) {
	fmt.Println("addr: ", packedAddr.Addr, address)
	st := struct {
		Type           string             `json:"@type"`
		AccountAddress TONUnpackedAddress `json:"account_address"`
	}{
		Type:           "packAccountAddress",
		AccountAddress: *packedAddr,
	}
	resp, err := client.executeAsynchronously(st)
	if err != nil {
		return "", err
	}
	if st, ok := resp.Data["@type"]; ok && st == "error" {
		return "", fmt.Errorf("Error ton client init. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	}

	addressSt := struct {
		TONAccountAddress
		Type string `json:"@type"`
	}{}
	err = json.Unmarshal(resp.Raw, &addressSt)
	if err != nil {
		return "", err
	}
	return addressSt.TONAccountAddress.AccountAddress, nil
}

// GetAccountState raw.getAccountState take account state
func (client *Client) GetAccountState(address string) (state *TONAccountState, err error) {
	st := struct {
		Type           string            `json:"@type"`
		AccountAddress TONAccountAddress `json:"account_address"`
	}{
		Type: "raw.getAccountState",
		AccountAddress: TONAccountAddress{
			AccountAddress: address,
		},
	}
	resp, err := client.executeAsynchronously(st)
	if err != nil {
		return state, err
	}
	if st, ok := resp.Data["@type"]; ok && st == "error" {
		return state, fmt.Errorf("Error ton get account sate. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	}

	state = new(TONAccountState)
	err = json.Unmarshal(resp.Raw, state)
	if err != nil {
		return state, err
	}
	return state, nil
}

// SendGrams2Address generic.sendGrams sends GRAM to address and returns transaction`s hash
func (client *Client) SendGrams2Address(key *TONPrivateKey, password []byte, fromAddress, toAddress, amount, message string) (string, error) {
	st := struct {
		Type        string            `json:"@type"`
		Seqno       int64             `json:"seqno"`
		Amount      string            `json:"amount"`
		PrivateKey  InputKey          `json:"private_key"`
		Destination TONAccountAddress `json:"destination"`
		ValidUntil  uint              `json:"valid_until"`
		Source      TONAccountAddress `json:"source"`
		Message     []byte            `json:"message"`
	}{
		Type:       "generic.sendGrams",
		PrivateKey: key.getInputKey(password),
		Amount:     amount,
		Destination: TONAccountAddress{
			AccountAddress: toAddress,
		},
		Seqno: 2,
		Source: TONAccountAddress{
			AccountAddress: fromAddress,
		},
		Message: []byte(message),
	}
	resp, err := client.executeAsynchronously(st)
	if err != nil {
		return "", err
	}
	if st, ok := resp.Data["@type"]; ok && st == "error" {
		return "", fmt.Errorf("Error ton send grams. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	}

	r := struct {
		SentUntil int    `json:"sent_until"`
		BodyHash  string `json:"body_hash"`
	}{}
	err = json.Unmarshal(resp.Raw, &r)
	if err != nil {
		return "", err
	}
	return r.BodyHash, nil
}

// CreateQuery4SendGrams2Address generic.createSendGramsQuery sends GRAM to address and returns transaction`s hash
// timeout must be between 0 and 300
func (client *Client) CreateQuery4SendGrams2Address(key *TONPrivateKey, password []byte, fromAddress, toAddress, amount, message string, timeout uint, allow_send_to_uninited bool) (string, error) {
	st := struct {
		Type                string            `json:"@type"`
		Seqno               int64             `json:"seqno"`
		Amount              string            `json:"amount"`
		PrivateKey          InputKey          `json:"private_key"`
		Destination         TONAccountAddress `json:"destination"`
		ValidUntil          uint              `json:"valid_until"`
		Source              TONAccountAddress `json:"source"`
		Message             []byte            `json:"message"`
		Timeout             uint              `json:"timeout"`
		AllowSendToUninited bool              `json:"allow_send_to_uninited"`
	}{
		Type:       "generic.createSendGramsQuery",
		PrivateKey: key.getInputKey(password),
		Amount:     amount,
		Destination: TONAccountAddress{
			AccountAddress: toAddress,
		},
		Seqno: 2,
		Source: TONAccountAddress{
			AccountAddress: fromAddress,
		},
		Message:             []byte(message),
		AllowSendToUninited: allow_send_to_uninited,
		Timeout:             timeout,
	}
	resp, err := client.executeAsynchronously(st)
	if err != nil {
		return "", err
	}
	if st, ok := resp.Data["@type"]; ok && st == "error" {
		return "", fmt.Errorf("Error ton create query for sending grams. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	}

	r := struct {
		SentUntil int    `json:"sent_until"`
		BodyHash  string `json:"body_hash"`
	}{}
	err = json.Unmarshal(resp.Raw, &r)
	if err != nil {
		return "", err
	}
	return r.BodyHash, nil
}

// SendMessage raw.sendMessage to address
func (client *Client) SendMessage(destinationAddress string, initialAccountState, data []byte) (res *TONResult, err error) {
	st := struct {
		Type                string            `json:"@type"`
		Destination         TONAccountAddress `json:"destination"`
		InitialAccountState []byte            `json:"initial_account_state"`
		Data                []byte            `json:"data"`
	}{
		Type: "raw.sendMessage",
		Data: data,
		Destination: TONAccountAddress{
			AccountAddress: destinationAddress,
		},
		InitialAccountState: initialAccountState,
	}
	return client.executeAsynchronously(st)
}

// CreateAndSendMessage raw.sendMessage to address
func (client *Client) CreateAndSendMessage(destinationAddress string, initialAccountState, data []byte) (res *TONResult, err error) {
	st := struct {
		Type                string            `json:"@type"`
		Destination         TONAccountAddress `json:"destination"`
		InitialAccountState []byte            `json:"initial_account_state"`
		Data                []byte            `json:"data"`
	}{
		Type: "raw.createAndSendMessage",
		Data: data,
		Destination: TONAccountAddress{
			AccountAddress: destinationAddress,
		},
		InitialAccountState: initialAccountState,
	}
	return client.executeAsynchronously(st)
}

// GetAccountTransactions raw.getTransactions fetch address`s transactions
func (client *Client) GetAccountTransactions(address string, lt string, hash string) (txs *TONTransactionsResponse, err error) {
	st := struct {
		Type              string                 `json:"@type"`
		AccountAddress    TONAccountAddress      `json:"account_address"`
		FromTransactionId *InternalTransactionId `json:"from_transaction_id"`
	}{
		Type: "raw.getTransactions",
		AccountAddress: TONAccountAddress{
			AccountAddress: address,
		},
		FromTransactionId: &InternalTransactionId{
			Lt:   lt,
			Hash: hash,
		},
	}
	resp, err := client.executeAsynchronously(st)
	if err != nil {
		return txs, err
	}
	if st, ok := resp.Data["@type"]; ok && st == "error" {
		return txs, fmt.Errorf("Error ton get account transactions. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	}

	txs = new(TONTransactionsResponse)
	err = json.Unmarshal(resp.Raw, txs)
	return txs, err
}

//sync node`s blocks to current
func (client *Client) Sync(fromBlock, toBlock, currentBlock int) error {
	data := struct {
		Type      string       `json:"@type"`
		SyncState TONSyncState `json:"sync_state"`
	}{
		Type: "sync",
		SyncState: TONSyncState{
			FromSeqno:    fromBlock,
			CurrentSeqno: currentBlock,
			ToSeqno:      toBlock,
		},
	}
	req, err := json.Marshal(data)
	if err != nil {
		return err
	}
	cs := C.CString(string(req))
	defer C.free(unsafe.Pointer(cs))

	C.tonlib_client_json_send(client.client, cs)
	for {
		result := C.tonlib_client_json_receive(client.client, DEFAULT_TIMEOUT)

		for result == nil {
			fmt.Println("empty response. next attempt")
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
		fmt.Println("sync result", string(resB))
		if err != nil {
			return err
		}
		if syncResp.Type == "ok" || syncResp.SyncState.Type == "syncStateDone" {
			return nil
		}
	}
}

func (client *Client) Destroy() {
	C.tonlib_client_json_destroy(client.client)
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
		updateReq := map[string]string{}
		err = json.Unmarshal(resB, &updateReq)
		if err == nil {
			client.updateSendLiteServerQuery(updateReq["id"], updateReq["data"])
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
		err = client.Sync(syncResp.SyncState.FromSeqno, syncResp.SyncState.ToSeqno, syncResp.SyncState.CurrentSeqno)
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

//
func (client *Client) updateSendLiteServerError(id, data string) (res *TONResult, err error) {
	queryID, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	st := struct {
		Type  string `json:"@type"`
		Id    int32  `json:"id"`
		Bytes []byte `json:"bytes"`
	}{
		Type:  "onLiteServerQueryError",
		Bytes: []byte(data),
		Id:    int32(queryID),
	}
	return client.executeAsynchronously(st)
}

func (client *Client) getLogStream(id, data string) (res *TONResult, err error) {
	queryID, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	st := struct {
		Type  string `json:"@type"`
		Id    int32  `json:"id"`
		Bytes []byte `json:"bytes"`
	}{
		Type:  "getLogStream",
		Bytes: []byte(data),
		Id:    int32(queryID),
	}
	return client.executeAsynchronously(st)
}

func (client *Client) updateSendLiteServerQuery(id, data string) (res *TONResult, err error) {
	queryID, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	st := struct {
		Type  string `json:"@type"`
		Id    int32  `json:"id"`
		Bytes []byte `json:"bytes"`
	}{
		Type:  "onLiteServerQueryResult",
		Bytes: []byte(data),
		Id:    int32(queryID),
	}
	return client.executeAsynchronously(st)
}
