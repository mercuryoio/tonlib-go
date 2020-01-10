package v2

import (
	"fmt"
	"testing"
)

const (
	TestAccountAddress = "EQDfYZhDfNJ0EePoT5ibfI9oG9bWIU6g872oX5h9rL5PHY9a"
	TestTxLt           = 289040000001
	TestTxHash         = "V6R8l0hTjpGb/HHHtDwrMk1KxTDLpfz5h7PINr1crp4="
	TestAmount         = "100000000"
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
		locPassword,
		TONPrivateKey{
			pKey.PublicKey,
			fmt.Sprintf("%x", (*pKey.Secret)[:]),
		},
	})
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
		locPassword,
		TONPrivateKey{
			pKey.PublicKey,
			fmt.Sprintf("%x", (*pKey.Secret)[:]),
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
