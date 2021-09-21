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
	defer cln.Close()

	// prepare data
	// generate new 24 words mnemo on tonwallet.me or https://github.com/cryptoboyio/ton-mnemonic
	// mnemonicPass should be empty
	var mnemonic []string
	var mnemonicPass string

	walletID := int64(698983191)
	revision := 1
	workchainID := 0

	wordList := make([]v2.SecureString, 0, len(mnemonic))
	for _, word := range mnemonic {
		wordList = append(wordList, v2.SecureString(word))
	}

	exportedKey := v2.NewExportedKey(wordList)
	key, err := cln.ImportKey(*exportedKey, v2.SecureBytes(mnemonicPass), v2.SecureBytes(mnemonicPass))
	if err != nil {
		log.Fatalln("Import key error", err)
	}

	sourceAccState := v2.NewWalletV3InitialAccountState(key.PublicKey, v2.JSONInt64(walletID))
	accountAddr, err := cln.GetAccountAddress(sourceAccState, int32(revision), int32(workchainID))
	if err != nil {
		log.Fatalln("Get account address error", err)
	}

	msgAction := v2.NewActionMsg(
		true,
		[]v2.MsgMessage{*v2.NewMsgMessage(
			v2.JSONInt64(TEST_AMOUNT),
			v2.NewMsgDataText(""),
			v2.NewAccountAddress(TEST_ADDRESS),
			"",
			-1,
		)},
	)

	inputKey := v2.InputKey{
		Type:          "inputKeyRegular",
		LocalPassword: base64.StdEncoding.EncodeToString(v2.SecureBytes(mnemonicPass)),
		Key:           v2.TONPrivateKey{Type: "key", PublicKey: key.PublicKey, Secret: key.Secret},
	}

	queryInfo, err := cln.CreateQuery(
		msgAction,
		*accountAddr,
		sourceAccState,
		inputKey,
		90, // If this timeout will be exceeded - all request are go as usual but grams wil not be sent
	)
	if err != nil {
		log.Fatalln("Create query error", err)
	}

	ok, err := cln.QuerySend(queryInfo.Id)
	if err != nil {
		log.Fatalln("Send query error", err)
	}

	fmt.Println(fmt.Sprintf("Send grams with queryInfo: %#v, queryResult: %#v", queryInfo, ok))
}
