package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/mercuryoio/tonlib-go/v2"
)

func main() {
	// parse config
	options, err := v2.ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		log.Fatal("failed parse config error. ", err)
	}

	// make req
	req := v2.TonInitRequest{
		"init",
		*options,
	}
	cln, err := v2.NewClient(&req, v2.Config{}, 10, true, 9)
	if err != nil {
		log.Fatalln("Init client error", err)
	}
	defer cln.Destroy()

	// create private key
	// prepare data
	loc := v2.SecureBytes(TEST_PASSWORD)
	mem := v2.SecureBytes(TEST_PASSWORD)
	seed := v2.SecureBytes("")

	// create new key
	pKey, err := cln.CreateNewKey(loc, mem, seed)
	if err != nil {
		log.Fatalln("Ton create key for send grams error", err)
	}

	// prepare input key
	inputKey := v2.InputKey{
		"inputKeyRegular",
		base64.StdEncoding.EncodeToString(loc),
		v2.TONPrivateKey{
			pKey.PublicKey,
			base64.StdEncoding.EncodeToString([]byte(pKey.Secret)),
		},
	}

	_, err = cln.WalletInit(&inputKey)
	if err != nil {
		log.Fatalln("Ton init wallet for send gramms error", err)
	}
	address, err := cln.WalletGetAccountAddress(v2.NewWalletInitialAccountState(pKey.PublicKey))
	if err != nil {
		log.Fatalln("Ton get address for send grams error", err)
	}

	// send grams
	sendResult, err := cln.GenericSendGrams(
		true,
		TEST_AMOUNT,
		v2.NewAccountAddress(TEST_ADDRESS),
		[]byte(""),
		&inputKey,
		address,
		5,
	)
	if err != nil {
		log.Fatalln("Ton send grams error ", err)
	}
	fmt.Println(fmt.Sprintf("Send grams with resp: %#v. ", sendResult))
}
