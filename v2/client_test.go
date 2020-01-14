package v2

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"
)

const (
	TestAccountAddress = "EQDfYZhDfNJ0EePoT5ibfI9oG9bWIU6g872oX5h9rL5PHY9a"
	TestTxLt           = 289040000001
	TestTxHash         = "V6R8l0hTjpGb/HHHtDwrMk1KxTDLpfz5h7PINr1crp4="
	TestAmount         = 100000000
	TestPassword       = "test_password"
	TestAddress        = "EQDfYZhDfNJ0EePoT5ibfI9oG9bWIU6g872oX5h9rL5PHY9a"
)

func TestClient_InitWallet(t *testing.T) {
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
	cln, err := NewClient(&req, Config{})
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
	cln, err := NewClient(&req, Config{})
	if err != nil {
		t.Fatal("Init client error. ", err)
	}
	defer cln.Destroy()

	// prepare data
	loc := SecureBytes(TestPassword)
	mem := SecureBytes(TestPassword)
	seed := SecureBytes("")

	// create new key
	pKey, err := cln.Createnewkey(&loc, &mem, &seed)
	if err != nil {
		t.Fatal("Ton create key for init wallet error", err)
	}
	fmt.Println(fmt.Sprintf("pKey: %#v", pKey))
}

func TestClient_Deletekey(t *testing.T) {
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
	cln, err := NewClient(&req, Config{})
	if err != nil {
		t.Fatal("Init client error. ", err)
	}
	defer cln.Destroy()

	// prepare data
	loc := SecureBytes(TestPassword)
	mem := SecureBytes(TestPassword)
	seed := SecureBytes("")

	// create new key
	pKey, err := cln.Createnewkey(&loc, &mem, &seed)
	if err != nil {
		t.Fatal("Ton create key for init wallet error", err)
	}
	fmt.Println(fmt.Sprintf("pKey: %#v", pKey))

	// delete new key
	ok, err := cln.Deletekey(pKey)
	if err != nil {
		t.Fatal("failed to delete new key", err)
	}
	fmt.Println(fmt.Sprintf("Ok: %#v, err: %v. ", ok, err))
}

func TestClient_Exportkey(t *testing.T) {
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
	cln, err := NewClient(&req, Config{})
	if err != nil {
		t.Fatal("Init client error. ", err)
	}
	defer cln.Destroy()

	// prepare data
	loc := SecureBytes(TestPassword)
	mem := SecureBytes(TestPassword)
	seed := SecureBytes("")

	// create new key
	pKey, err := cln.Createnewkey(&loc, &mem, &seed)
	if err != nil {
		t.Fatal("Ton create key for init wallet error", err)
	}
	fmt.Println(fmt.Sprintf("pKey: %#v", pKey))

	// export key
	exportedKey, err := cln.Exportkey(&InputKey{
		"inputKeyRegular",
		base64.StdEncoding.EncodeToString(loc),
		TONPrivateKey{
			pKey.PublicKey,
			base64.StdEncoding.EncodeToString((*pKey.Secret)[:]),
		},
	}, )
	if err != nil {
		t.Fatal("Ton export key error", err)
	}
	fmt.Println(fmt.Sprintf("exported key: %#v", exportedKey))
}

func TestClient_Exportpemkey(t *testing.T) {
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
	cln, err := NewClient(&req, Config{})
	if err != nil {
		t.Fatal("Init client error. ", err)
	}
	defer cln.Destroy()

	// prepare data
	loc := SecureBytes(TestPassword)
	mem := SecureBytes(TestPassword)
	seed := SecureBytes("")

	// create new key
	pKey, err := cln.Createnewkey(&loc, &mem, &seed)
	if err != nil {
		t.Fatal("Ton create key for init wallet error", err)
	}
	fmt.Println(fmt.Sprintf("pKey: %#v.", pKey))

	// export key
	exportedKey, err := cln.Exportpemkey(&InputKey{
		"inputKeyRegular",
		base64.StdEncoding.EncodeToString(loc),
		TONPrivateKey{
			pKey.PublicKey,
			base64.StdEncoding.EncodeToString((*pKey.Secret)[:]),
		},
	}, &loc)
	if err != nil {
		t.Fatal("Ton export key error", err)
	}
	fmt.Println(fmt.Sprintf("exported key: %#v", exportedKey))
}

func TestClient_RawGetaccountstate(t *testing.T) {
	// parse config
	options, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("TestClient_RawGetaccountstate failed parse config error. ", err)
	}

	// make req
	req := TonInitRequest{
		"init",
		*options,
	}

	// create client
	cln, err := NewClient(&req, Config{})
	if err != nil {
		t.Fatal("TestClient_RawGetaccountstate Init client error. ", err)
	}
	defer cln.Destroy()

	ok, err := cln.RawGetaccountstate(NewAccountAddress(TestAccountAddress))
	if err != nil {
		t.Fatal("TestClient_RawGetaccountstate failed to RawGetaccountstate(): ", err)
	}

	fmt.Printf("TestClient_RawGetaccountstate: ok: %#v, err: %v. ", ok, err)
}

func TestClient_WalletInit(t *testing.T) {
	// parse config
	options, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("TestClient_WalletInit failed parse config error. ", err)
	}

	// make req
	req := TonInitRequest{
		"init",
		*options,
	}

	// create client
	cln, err := NewClient(&req, Config{})
	if err != nil {
		t.Fatal("TestClient_WalletInit Init client error. ", err)
	}
	defer cln.Destroy()

	// prepare data
	loc := SecureBytes(TestPassword)
	mem := SecureBytes(TestPassword)
	seed := SecureBytes("")

	// create new key
	pKey, err := cln.Createnewkey(&loc, &mem, &seed)
	if err != nil {
		t.Fatal("TestClient_WalletInit create key for init wallet error", err)
	}
	fmt.Println(fmt.Sprintf("TestClient_WalletInit pKey: %#v", pKey))

	// init wallet
	ok, err := cln.WalletInit(
		&InputKey{
			"inputKeyRegular",
			base64.StdEncoding.EncodeToString(loc),
			TONPrivateKey{
				pKey.PublicKey,
				base64.StdEncoding.EncodeToString((*pKey.Secret)[:]),
			},
		},
	)
	if err != nil {
		t.Fatal("TestClient_WalletInit failed to WalletInit(): ", err)
	}

	fmt.Printf("TestClient_WalletInit: ok: %#v, err: %v. ", ok, err)
}

func TestClient_WalletGetaccountaddress(t *testing.T) {
	// parse config
	options, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("TestClient_WalletGetaccountaddress failed parse config error. ", err)
	}

	// make req
	req := TonInitRequest{
		"init",
		*options,
	}

	// create client
	cln, err := NewClient(&req, Config{})
	if err != nil {
		t.Fatal("TestClient_WalletGetaccountaddress Init client error. ", err)
	}
	defer cln.Destroy()

	// prepare data
	loc := SecureBytes(TestPassword)
	mem := SecureBytes(TestPassword)
	seed := SecureBytes("")

	// create new key
	pKey, err := cln.Createnewkey(&loc, &mem, &seed)
	if err != nil {
		t.Fatal("TestClient_WalletGetaccountaddress create key for init wallet error", err)
	}
	fmt.Println(fmt.Sprintf("TestClient_WalletGetaccountaddress pKey: %#v", pKey))

	// init wallet
	ok, err := cln.WalletInit(
		&InputKey{
			"inputKeyRegular",
			base64.StdEncoding.EncodeToString(loc),
			TONPrivateKey{
				pKey.PublicKey,
				base64.StdEncoding.EncodeToString((*pKey.Secret)[:]),
			},
		},
	)
	if err != nil {
		t.Fatal("TestClient_WalletGetaccountaddress failed to WalletInit(): ", err)
	}

	fmt.Printf("TestClient_WalletGetaccountaddress: init wallet ok: %#v, err: %v. ", ok, err)

	// get wallet adress info
	addrr, err := cln.WalletGetaccountaddress(NewWalletInitialAccountState(pKey.PublicKey))
	if err != nil {
		t.Fatal("TestClient_WalletGetaccountaddress failed to WalletGetaccountaddress(): ", err)
	}

	fmt.Printf("TestClient_WalletGetaccountaddress: get account adress addr: %#v, err: %v. ", addrr, err)
}

func TestClient_WalletGetaccountstate(t *testing.T) {
	// parse config
	options, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("TestClient_WalletGetaccountstate failed parse config error. ", err)
	}

	// make req
	req := TonInitRequest{
		"init",
		*options,
	}

	// create client
	cln, err := NewClient(&req, Config{})
	if err != nil {
		t.Fatal("TestClient_WalletGetaccountstate Init client error. ", err)
	}
	defer cln.Destroy()

	// prepare data
	loc := SecureBytes(TestPassword)
	mem := SecureBytes(TestPassword)
	seed := SecureBytes("")

	// create new key
	pKey, err := cln.Createnewkey(&loc, &mem, &seed)
	if err != nil {
		t.Fatal("TestClient_WalletGetaccountstate create key for init wallet error", err)
	}
	fmt.Println(fmt.Sprintf("TestClient_WalletGetaccountstate pKey: %#v", pKey))

	// init wallet
	ok, err := cln.WalletInit(
		&InputKey{
			"inputKeyRegular",
			base64.StdEncoding.EncodeToString(loc),
			TONPrivateKey{
				pKey.PublicKey,
				base64.StdEncoding.EncodeToString((*pKey.Secret)[:]),
			},
		},
	)
	if err != nil {
		t.Fatal("TestClient_WalletGetaccountstate failed to WalletInit(): ", err)
	}
	fmt.Printf("TestClient_WalletGetaccountstate: init wallet ok: %#v, err: %v. ", ok, err)

	// get wallet adress info
	addrr, err := cln.WalletGetaccountaddress(NewWalletInitialAccountState(pKey.PublicKey))
	if err != nil {
		t.Fatal("TestClient_WalletGetaccountstate failed to WalletGetaccountaddress(): ", err)
	}

	// get wallet account state info
	state, err := cln.WalletGetaccountstate(addrr)
	if err != nil {
		t.Fatal("TestClient_WalletGetaccountstate failed to WalletGetaccountstate(): ", err)
	}

	fmt.Printf("TestClient_WalletGetaccountstate: get account stater: %#v, err: %v. ", state, err)
}

func TestClient_WalletSendgrams(t *testing.T) {
	// parse config
	options, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("TestClient_WalletSendgrams failed parse config error. ", err)
	}

	// make req
	req := TonInitRequest{
		"init",
		*options,
	}

	// create client
	cln, err := NewClient(&req, Config{})
	if err != nil {
		t.Fatal("TestClient_WalletSendgrams Init client error. ", err)
	}
	defer cln.Destroy()

	// prepare data
	loc := SecureBytes(TestPassword)
	mem := SecureBytes(TestPassword)
	seed := SecureBytes("")

	// create new key
	pKey, err := cln.Createnewkey(&loc, &mem, &seed)
	if err != nil {
		t.Fatal("TestClient_WalletSendgrams create key for init wallet error", err)
	}
	fmt.Println(fmt.Sprintf("TestClient_WalletSendgrams pKey: %#v", pKey))

	// prepare input key
	inputKey := InputKey{
		"inputKeyRegular",
		base64.StdEncoding.EncodeToString(loc),
		TONPrivateKey{
			pKey.PublicKey,
			base64.StdEncoding.EncodeToString((*pKey.Secret)[:]),
		},
	}

	// init wallet
	ok, err := cln.WalletInit(
		&inputKey,
	)
	if err != nil {
		t.Fatal("TestClient_WalletSendgrams failed to WalletInit(): ", err)
	}
	fmt.Printf("TestClient_WalletSendgrams: init wallet ok: %#v, err: %v. \n", ok, err)

	// send grams
	fmt.Println("starts wallet send gramm")
	sendResult, err := cln.WalletSendgrams(
		2,
		0,
		TestAmount,
		[]byte("test send grams"),
		&inputKey,
		NewAccountAddress(TestAddress),
	)
	if err != nil {
		t.Fatal("TestClient_WalletSendgrams failed to WalletSendgrams(): ", err)
	}
	fmt.Printf("TestClient_WalletSendgrams: send grams: %#v, err: %v. ", sendResult, err)
}

func TestClient_RawGettransactions(t *testing.T) {
	// parse config
	options, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("TestClient_RawGettransactions failed parse config error. ", err)
	}

	// make req
	req := TonInitRequest{
		"init",
		*options,
	}

	// create client
	cln, err := NewClient(&req, Config{})
	if err != nil {
		t.Fatal("TestClient_RawGettransactions Init client error. ", err)
	}
	defer cln.Destroy()

	// prepare data
	addr := NewAccountAddress(TestAddress)

	// get account state
	state, err := cln.RawGetaccountstate(addr)
	if err != nil {
		t.Fatal("Get state error error", err)
	}

	_, err = cln.RawGettransactions(
		state.LastTransactionId,
		addr,
	)
	if err != nil {
		t.Fatal("Ton get account txs error", err)
	}
}

func TestClient_GenericSendgrams(t *testing.T) {
	// parse config
	options, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("TestClient_GenericSendgrams failed parse config error. ", err)
	}

	// make req
	req := TonInitRequest{
		"init",
		*options,
	}

	// create client
	cln, err := NewClient(&req, Config{})
	if err != nil {
		t.Fatal("TestClient_GenericSendgrams Init client error. ", err)
	}
	defer cln.Destroy()

	// prepare data
	loc := SecureBytes(TestPassword)
	mem := SecureBytes(TestPassword)
	seed := SecureBytes("")

	// create new key
	pKey, err := cln.Createnewkey(&loc, &mem, &seed)
	if err != nil {
		t.Fatal("TestClient_GenericSendgrams create key for init wallet error", err)
	}
	fmt.Println(fmt.Sprintf("TestClient_GenericSendgrams pKey: %#v", pKey))

	// prepare input key
	inputKey := InputKey{
		"inputKeyRegular",
		base64.StdEncoding.EncodeToString(loc),
		TONPrivateKey{
			pKey.PublicKey,
			base64.StdEncoding.EncodeToString((*pKey.Secret)[:]),
		},
	}

	// init wallet
	ok, err := cln.WalletInit(
		&inputKey,
	)
	if err != nil {
		t.Fatal("TestClient_GenericSendgrams failed to WalletInit(): ", err)
	}
	fmt.Printf("TestClient_GenericSendgrams: init wallet ok: %#v, err: %v. \n", ok, err)

	// get wallet adress info
	addrr, err := cln.WalletGetaccountaddress(NewWalletInitialAccountState(pKey.PublicKey))
	if err != nil {
		t.Fatal("TestClient_GenericSendgrams failed to WalletGetaccountaddress(): ", err)
	}
	fmt.Printf("TestClient_GenericSendgrams: get account adress addr: %#v, err: %v. ", addrr, err)

	// send grams
	sendResult, err := cln.GenericSendgrams(
		true,
		[]byte(""),
		&inputKey,
		addrr,
		NewAccountAddress(TestAddress),
		TestAmount,
		5,
	)
	if err != nil {
		t.Fatal("TestClient_GenericSendgrams failed to GenericSendgrams(): ", err)
	}
	fmt.Printf("TestClient_GenericSendgrams: sent grams: %#v, err: %v. ", sendResult, err)
}

func TestClient_RawCreateandsendmessage(t *testing.T) {
	// parse config
	options, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("TestClient_RawCreateandsendmessage failed parse config error. ", err)
	}

	// make req
	req := TonInitRequest{
		"init",
		*options,
	}

	// create client
	cln, err := NewClient(&req, Config{})
	if err != nil {
		t.Fatal("TestClient_RawCreateandsendmessage Init client error. ", err)
	}
	defer cln.Destroy()

	// prepare data
	loc := SecureBytes(TestPassword)
	mem := SecureBytes(TestPassword)
	seed := SecureBytes("")

	// create new key
	pKey, err := cln.Createnewkey(&loc, &mem, &seed)
	if err != nil {
		t.Fatal("TestClient_RawCreateandsendmessage create key for init wallet error", err)
	}
	fmt.Println(fmt.Sprintf("TestClient_RawCreateandsendmessage pKey: %#v", pKey))

	// prepare input key
	inputKey := InputKey{
		"inputKeyRegular",
		base64.StdEncoding.EncodeToString(loc),
		TONPrivateKey{
			pKey.PublicKey,
			base64.StdEncoding.EncodeToString((*pKey.Secret)[:]),
		},
	}

	// init wallet
	ok, err := cln.WalletInit(
		&inputKey,
	)
	if err != nil {
		t.Fatal("TestClient_RawCreateandsendmessage failed to WalletInit(): ", err)
	}
	fmt.Printf("TestClient_RawCreateandsendmessage: init wallet ok: %#v, err: %v. \n", ok, err)

	// get wallet address info
	addrr, err := cln.WalletGetaccountaddress(NewWalletInitialAccountState(pKey.PublicKey))
	if err != nil {
		t.Fatal("TestClient_RawCreateandsendmessage failed to WalletGetaccountaddress(): ", err)
	}
	fmt.Printf("TestClient_RawCreateandsendmessage: get account adress addr: %#v, err: %v. ", addrr, err)

	// read test message from file
	bocFile, err := ioutil.ReadFile("./testgiver-query.boc")
	if err != nil {
		t.Fatal("TestClient_RawCreateandsendmessage: boc file dosn't exist", err)
	}

	// send msg
	msgSentOk, err := cln.RawCreateandsendmessage(
		addrr,
		[]byte{},
		bocFile,
	)
	if err != nil {
		t.Fatal("TestClient_RawCreateandsendmessage failed to RawCreateandsendmessage(): ", err)
	}
	fmt.Printf("TestClient_RawCreateandsendmessage: create and send msg msgSentOk: %#v, err: %v. ", msgSentOk, err)
}
