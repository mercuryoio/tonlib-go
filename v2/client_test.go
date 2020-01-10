package v2

import (
	"encoding/base64"
	"fmt"
	"testing"
)

const (
	TestAccountAddress = "EQDfYZhDfNJ0EePoT5ibfI9oG9bWIU6g872oX5h9rL5PHY9a"
	TestTxLt           = 289040000001
	TestTxHash         = "V6R8l0hTjpGb/HHHtDwrMk1KxTDLpfz5h7PINr1crp4="
	TestAmount         = "100000000"
	TestPassword       = "test_password"
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
	loc := SecureBytes("dsfsdfdsf")
	mem := SecureBytes("1w1w1w1w1")
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
	loc := SecureBytes("dsfsdfdsf")
	mem := SecureBytes("1w1w1w1w1")
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
	locPassword := "odododoee22221"
	loc := SecureBytes(locPassword)
	mem := SecureBytes("sdlkdslk11")
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
	},)
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
	locPassword := "wefasfsafw"
	loc := SecureBytes(locPassword)
	mem := SecureBytes("sdlkdslk11")
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
