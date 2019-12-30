package v2

import (
	"fmt"
	"testing"
)

func TestClient_InitWallet(t *testing.T) {
	cnf := TonInitRequest{
		"init",
		*NewOptions(
			NewConfig(`{"liteservers":[{"@type":"liteserver.desc","ip":1137658550,"port":"4924","id":{"@type":"pub.ed25519","key":"peJTw/arlRfssgTuf9BMypJzqOi7SXEqSPSWiEw2U1M="}}],"validator":{"@type":"validator.config.global","zero_state":{"workchain":-1,"shard":-9223372036854775808,"seqno":0,"root_hash":"VCSXxDHhTALFxReyTZRd8E4Ya3ySOmpOWAS4rBX9XBY=","file_hash":"eh9yveSz1qMdJ7mOsO+I+H77jkLr9NpAuEkoJuseXBo="}}}`,
				"my-chain", false, true, ),
			&KeyStoreType{
				"keyStoreTypeDirectory",
				"./test.keys",
			}, ),
	}

	cln, err := NewClient(&cnf, Config{})
	if err != nil {
		t.Fatal("Init client error. ", err)
	}
	defer cln.Destroy()
	loc := SecureBytes("dsfsdfdsf")
	mem := SecureBytes("1w1w1w1w1")
	seed := SecureBytes("")
	pKey, err := cln.Createnewkey(&loc, &mem, &seed)
	if err != nil {
		t.Fatal("Ton create key for init wallet error", err)
	}
	fmt.Println(fmt.Sprintf("pKey: %#v", pKey))
}
