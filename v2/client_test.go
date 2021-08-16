package v2

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/mercuryoio/tonlib-go"
)

const (
	TestAccountAddress  = "UQB-8ullbN9cxf1ILhLBLwjScjCEFJnimZ9WG3xiR6O5wr0e"
	TestAccountPassword = "testmm"
	TestAccountPublic   = "PuauM0qjqVGAr8t1YNfbqeY8pTtCniGk7lwGFll0oZZt60wh"
	TestAccountSecret   = "AWO+xDBnFHnB5SJ0270Xd4YvPb86iUPSQr1DLgdmbW8="
	TestTxLt            = 289040000001
	TestTxHash          = "V6R8l0hTjpGb/HHHtDwrMk1KxTDLpfz5h7PINr1crp4="
	TestAmount          = 100000000
	TestPassword        = "test_password"
	DefaultTestTimeout  = 10
)

func TestClient_NewClient(t *testing.T) {
	// parse config
	options, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("parse config error. ", err)
	}

	// make req
	req := TonInitRequest{
		"init",
		*options,
	}

	// create client
	cln, err := NewClient(&req, Config{}, DefaultTestTimeout, true, 100)
	if err != nil {
		t.Fatal("Init client error. ", err)
	}
	defer cln.Destroy()
}

func TestClient_CreateNewKey(t *testing.T) {
	// parse config
	options, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("parse config error. ", err)
	}

	// make req
	req := TonInitRequest{
		"init",
		*options,
	}

	// create client
	cln, err := NewClient(&req, Config{}, DefaultTestTimeout, true, 100)
	if err != nil {
		t.Fatal("Init client error. ", err)
	}
	defer cln.Destroy()

	// prepare data
	loc := SecureBytes(TestPassword)
	mem := SecureBytes(TestPassword)
	seed := SecureBytes("")

	// create new key
	pKey, err := cln.CreateNewKey(loc, mem, seed)
	if err != nil {
		t.Fatal("Ton create key for init wallet error", err)
	}
	fmt.Println(fmt.Sprintf("pKey: %#v", pKey))
}

func TestClient_DeleteKey(t *testing.T) {
	// parse config
	options, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("parse config error. ", err)
	}

	// make req
	req := TonInitRequest{
		"init",
		*options,
	}

	// create client
	cln, err := NewClient(&req, Config{}, DefaultTestTimeout, true, 100)
	if err != nil {
		t.Fatal("Init client error. ", err)
	}
	defer cln.Destroy()

	// prepare data
	loc := SecureBytes(TestPassword)
	mem := SecureBytes(TestPassword)
	seed := SecureBytes("")

	// create new key
	pKey, err := cln.CreateNewKey(loc, mem, seed)
	if err != nil {
		t.Fatal("Ton create key for init wallet error", err)
	}
	fmt.Println(fmt.Sprintf("pKey: %s, secret: %s.", pKey.PublicKey, string(pKey.Secret)))

	// delete new key
	ok, err := cln.DeleteKey(*pKey)
	if err != nil {
		t.Fatal("failed to delete new key", err)
	}
	fmt.Println(fmt.Sprintf("Ok: %#v, err: %v. ", ok, err))
}

func TestClient_ExportKey(t *testing.T) {
	// parse config
	options, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("parse config error. ", err)
	}

	// make req
	req := TonInitRequest{
		"init",
		*options,
	}

	// create client
	cln, err := NewClient(&req, Config{}, DefaultTestTimeout, true, 100)
	if err != nil {
		t.Fatal("Init client error. ", err)
	}
	defer cln.Destroy()

	// prepare data
	loc := SecureBytes(TestPassword)
	mem := SecureBytes(TestPassword)
	seed := SecureBytes("")

	// create new key
	pKey, err := cln.CreateNewKey(loc, mem, seed)
	if err != nil {
		t.Fatal("Ton create key for init wallet error", err)
	}
	fmt.Println(fmt.Sprintf("pKey: %#v", pKey))

	// export key
	exportedKey, err := cln.ExportKey(InputKey{
		"inputKeyRegular",
		base64.StdEncoding.EncodeToString(loc),
		TONPrivateKey{
			pKey.PublicKey,
			pKey.Secret,
		},
	})
	if err != nil {
		t.Fatal("Ton export key error", err)
	}
	fmt.Println(fmt.Sprintf("exported key: %#v", exportedKey))
}

func TestClient_ExportPemKey(t *testing.T) {
	// parse config
	options, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("parse config error. ", err)
	}

	// make req
	req := TonInitRequest{
		"init",
		*options,
	}

	// create client
	cln, err := NewClient(&req, Config{}, DefaultTestTimeout, true, 100)
	if err != nil {
		t.Fatal("Init client error. ", err)
	}
	defer cln.Destroy()

	// prepare data
	loc := SecureBytes(TestPassword)
	mem := SecureBytes(TestPassword)
	seed := SecureBytes("")

	// create new key
	pKey, err := cln.CreateNewKey(loc, mem, seed)
	if err != nil {
		t.Fatal("Ton create key for init wallet error", err)
	}
	fmt.Println(fmt.Sprintf("pKey: %#v.", pKey))

	// export key
	exportedKey, err := cln.ExportPemKey(InputKey{
		"inputKeyRegular",
		base64.StdEncoding.EncodeToString(loc),
		TONPrivateKey{
			pKey.PublicKey,
			pKey.Secret,
		},
	}, loc)
	if err != nil {
		t.Fatal("Ton export key error", err)
	}
	fmt.Println(fmt.Sprintf("exported key: %#v", exportedKey))
}

func TestClient_RawGetAccountState(t *testing.T) {
	// parse config
	options, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("TestClient_RawGetAccountState failed parse config error. ", err)
	}

	// make req
	req := TonInitRequest{
		"init",
		*options,
	}

	// create client
	cln, err := NewClient(&req, Config{}, DefaultTestTimeout, true, 100)
	if err != nil {
		t.Fatal("TestClient_RawGetAccountState Init client error. ", err)
	}
	defer cln.Destroy()

	ok, err := cln.RawGetAccountState(*NewAccountAddress(TestAccountAddress))
	if err != nil {
		t.Fatal("TestClient_RawGetAccountState failed to RawGetAccountState(): ", err)
	}

	fmt.Printf("TestClient_RawGetAccountState: ok: %#v, err: %v. ", ok, err)
}

func TestClient_WalletGetAccountAddress(t *testing.T) {
	// parse config
	options, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("TestClient_WalletGetAccountAddress failed parse config error. ", err)
	}

	// make req
	req := TonInitRequest{
		"init",
		*options,
	}

	// create client
	cln, err := NewClient(&req, Config{}, DefaultTestTimeout, true, 100)
	if err != nil {
		t.Fatal("TestClient_WalletGetAccountAddress Init client error. ", err)
	}
	defer cln.Destroy()

	// prepare data
	loc := SecureBytes(TestPassword)
	mem := SecureBytes(TestPassword)
	seed := SecureBytes("")

	// create new key
	pKey, err := cln.CreateNewKey(loc, mem, seed)
	if err != nil {
		t.Fatal("TestClient_WalletGetAccountAddress create key for init wallet error", err)
	}
	fmt.Println(fmt.Sprintf("TestClient_WalletGetAccountAddress pKey: %#v", pKey))

	// get wallet adress info
	addrr, err := cln.GetAccountAddress(tonlib.NewWalletInitialAccountState(pKey.PublicKey), 0, 0)
	if err != nil {
		t.Fatal("TestClient_WalletGetAccountAddress failed to WalletGetAccountAddress(): ", err)
	}

	fmt.Printf("TestClient_WalletGetAccountAddress: get account adress addr: %#v, err: %v. ", addrr, err)
}

func TestClient_WalletGetAccountState(t *testing.T) {
	// parse config
	options, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("TestClient_WalletGetAccountState failed parse config error. ", err)
	}

	// make req
	req := TonInitRequest{
		"init",
		*options,
	}

	// create client
	cln, err := NewClient(&req, Config{}, DefaultTestTimeout, true, 100)
	if err != nil {
		t.Fatal("TestClient_WalletGetAccountState Init client error. ", err)
	}
	defer cln.Destroy()

	// prepare data
	loc := SecureBytes(TestPassword)
	mem := SecureBytes(TestPassword)
	seed := SecureBytes("")

	// create new key
	pKey, err := cln.CreateNewKey(loc, mem, seed)
	if err != nil {
		t.Fatal("TestClient_WalletGetAccountState create key for init wallet error", err)
	}
	fmt.Println(fmt.Sprintf("TestClient_WalletGetAccountState pKey: %#v", pKey))

	// get wallet adress info
	addrr, err := cln.GetAccountAddress(tonlib.NewWalletInitialAccountState(pKey.PublicKey), 0, 0)
	if err != nil {
		t.Fatal("TestClient_WalletGetAccountState failed to WalletGetAccountAddress(): ", err)
	}

	// get wallet account state info
	state, err := cln.GetAccountState(*addrr)
	if err != nil {
		t.Fatal("TestClient_WalletGetAccountState failed to WalletGetAccountState(): ", err)
	}

	fmt.Printf("TestClient_WalletGetAccountState: get account stater: %#v, err: %v. ", state, err)
}

func TestClient_RawGetTransactions(t *testing.T) {
	// parse config
	options, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("TestClient_RawGetTransactions failed parse config error. ", err)
	}

	// make req
	req := TonInitRequest{
		"init",
		*options,
	}

	// create client
	cln, err := NewClient(&req, Config{}, DefaultTestTimeout, true, 100)
	if err != nil {
		t.Fatal("TestClient_RawGetTransactions Init client error. ", err)
	}
	defer cln.Destroy()

	// prepare data
	addr := NewAccountAddress(TestAccountAddress)
	inputKey := InputKey{
		Type:          "inputKeyRegular",
		LocalPassword: base64.StdEncoding.EncodeToString(SecureBytes(TestAccountPassword)),
		Key:           TONPrivateKey{PublicKey: TestAccountPublic, Secret: TestAccountSecret},
	}

	// get account state
	state, err := cln.RawGetAccountState(*addr)
	if err != nil {
		t.Fatal("Get state error error", err)
	}

	_, err = cln.RawGetTransactions(
		*addr,
		*state.LastTransactionId,
		inputKey,
	)
	if err != nil {
		t.Fatal("Ton get account txs error", err)
	}
}

func TestClient_RawCreateAndSendMessage(t *testing.T) {
	// parse config
	options, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("TestClient_RawCreateAndSendMessage failed parse config error. ", err)
	}

	// make req
	req := TonInitRequest{
		"init",
		*options,
	}

	// create client
	cln, err := NewClient(&req, Config{}, DefaultTestTimeout, true, 100)
	if err != nil {
		t.Fatal("TestClient_RawCreateAndSendMessage Init client error. ", err)
	}
	defer cln.Destroy()

	// prepare data
	loc := SecureBytes(TestPassword)
	mem := SecureBytes(TestPassword)
	seed := SecureBytes("")

	// create new key
	pKey, err := cln.CreateNewKey(loc, mem, seed)
	if err != nil {
		t.Fatal("TestClient_RawCreateAndSendMessage create key for init wallet error", err)
	}
	fmt.Println(fmt.Sprintf("TestClient_RawCreateAndSendMessage pKey: %#v", pKey))

	// get wallet address info
	addrr, err := cln.GetAccountAddress(tonlib.NewWalletInitialAccountState(pKey.PublicKey), 0, 0)
	if err != nil {
		t.Fatal("TestClient_RawCreateAndSendMessage failed to WalletGetAccountAddress(): ", err)
	}
	fmt.Printf("TestClient_RawCreateAndSendMessage: get account adress addr: %#v, err: %v. ", addrr, err)

	// read test message from file
	bocFile, err := ioutil.ReadFile("./testgiver-query.boc")
	if err != nil {
		t.Fatal("TestClient_RawCreateAndSendMessage: boc file dosn't exist", err)
	}

	// send msg
	msgSentOk, err := cln.RawCreateAndSendMessage(
		bocFile,
		*addrr,
		[]byte{},
	)
	if err != nil {
		t.Fatal("TestClient_RawCreateAndSendMessage failed to RawCreateAndSendMessage(): ", err)
	}
	fmt.Printf("TestClient_RawCreateAndSendMessage: create and send msg msgSentOk: %#v, err: %v. ", msgSentOk, err)
}

func TestGetType(t *testing.T) {
	wantValue := "query.estimateFees"

	data := struct {
		Type         string `json:"@type"`
		Id           int64  `json:"id"`
		IgnoreChksig bool   `json:"ignore_chksig"`
	}{
		Type:         wantValue,
		Id:           1,
		IgnoreChksig: true,
	}

	actualValue, err := getType(interface{}(data))
	if err != nil {
		t.Errorf("Error getType. Data: %v. Error: %v", data, err)
	}

	if wantValue != actualValue {
		t.Errorf("Error getType. WantValue: %s, ActualValue: %s", wantValue, actualValue)
	}
}
