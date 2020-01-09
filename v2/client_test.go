package v2

import (
	"fmt"
	"testing"
)

func TestClientInitWallet(t *testing.T) {
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


func TestClientCreateNewKey(t *testing.T) {
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
