package main

import (
	"github.com/mercuryoio/tonlib-go"
	"log"
	"time"
)

func main() {
	cnf, err := tonlib.ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		log.Fatalln("Config file not found", err)
	}
	cln, err := tonlib.NewClient(cnf, tonlib.Config{})
	if err != nil {
		log.Fatalln("Init client error", err)
	}
	defer cln.Destroy()
	pKey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
	if err != nil {
		log.Fatalln("Ton create key for send grams error", err)
	}
	err = cln.InitWallet(pKey, []byte(TEST_PASSWORD))
	if err != nil {
		log.Fatalln("Ton init wallet for send gramms error", err)
	}
	address, err := cln.WalletGetAddress(pKey.PublicKey)
	if err != nil {
		log.Fatalln("Ton get address for send grams error", err)
	}
	// todo create query doesn't work now
	_, err = cln.CreateQuery4SendGrams2Address(pKey, []byte(TEST_PASSWORD), address.AccountAddress, TEST_ADDRESS, TEST_AMOUNT, "", 0, true)
	if err != nil {
		log.Fatalln("Ton send grams error", err)
	}
	time.Sleep(5 * time.Second)
}
