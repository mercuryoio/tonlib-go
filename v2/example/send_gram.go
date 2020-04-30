package main

import (
	"encoding/base64"
	"fmt"
	tonlib "github.com/varche1/tonlib-go/v2"
	"log"
)

func main() {
	// parse config
	options, err := tonlib.ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		log.Fatal("failed parse config error. ", err)
	}

	// make req
	req := tonlib.TonInitRequest{
		"init",
		*options,
	}
	cln, _, err := tonlib.NewClient(&req, tonlib.Config{}, 10, true, 9)
	if err != nil {
		log.Fatalln("Init client error", err)
	}
	defer cln.Destroy()

	// create private key
	// prepare data
	loc := tonlib.SecureBytes(TEST_PASSWORD)
	mem := tonlib.SecureBytes(TEST_PASSWORD)
	seed := tonlib.SecureBytes("")

	// create new key
	pKey, err := cln.CreateNewKey(&loc, &mem, &seed)
	if err != nil {
		log.Fatalln("Ton create key for send grams error", err)
	}

	// prepare input key
	inputKey := tonlib.InputKey{
		"inputKeyRegular",
		base64.StdEncoding.EncodeToString(loc),
		tonlib.TONPrivateKey{
			pKey.PublicKey,
			base64.StdEncoding.EncodeToString((*pKey.Secret)[:]),
		},
	}

	_, err = cln.WalletInit(&inputKey)
	if err != nil {
		log.Fatalln("Ton init wallet for send gramms error", err)
	}
	address, err := cln.WalletGetAccountAddress(tonlib.NewWalletInitialAccountState(pKey.PublicKey))
	if err != nil {
		log.Fatalln("Ton get address for send grams error", err)
	}

	// send grams
	sendResult, err := cln.GenericSendGrams(
		true,
		TEST_AMOUNT,
		tonlib.NewAccountAddress(TEST_ADDRESS),
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
