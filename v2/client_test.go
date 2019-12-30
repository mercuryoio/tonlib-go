package v2

import (
	"testing"
)

func TestClient_InitWallet(t *testing.T) {
	cnf := TonInitRequest{
		"init",
		*NewOptions(
			NewConfig("", "", false, true, ),
			&KeyStoreType{
				"keyStoreTypeDirectory",
				"./test.keys",
			}, ),
	}

	cln, err := NewClient(&cnf, Config{})
	if err != nil {
		t.Fatal("Init client error", err)
	}
	defer cln.Destroy()
	loc := SecureBytes("dsfsdfdsf")
	mem := SecureBytes("1w1w1w1w1")
	seed := SecureBytes("")
	pKey, err := cln.Createnewkey(&loc, &mem, &seed)
	if err != nil {
		t.Fatal("Ton create key for init wallet error", err)
	}
	println("pKey", pKey)
}
