package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/mercuryoio/tonlib-go"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

var transactionsCmd = &cobra.Command{
	Use:   "transactions",
	Short: "Get transaction info",
	Long: `Get transaction info command. It contains 4 attributes:
- path2configfile. see tonlib.config.json.example
- address
- publicKey
- secret
- password
- lt (optional)
- hash (optional)
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 5 && len(args) != 7 {
			return fmt.Errorf("you have to use 5 or 7 args for this command \n")
		}
		_, err := os.Stat(args[0])
		if err != nil {
			errors.New("please choose config path")
		}
		return nil
	},
	Run: transactions,
}

func transactions(cmd *cobra.Command, args []string) {
	//panic("currently doesn't work, because of primary ton changes - working to get it back")
	confPath := args[0]
	address := args[1]
	publicKey := args[2]
	secret := args[3]
	password := args[4]
	var lt tonlib.JSONInt64
	var hash string
	var err error
	addr := tonlib.NewAccountAddress(address)

	// init client
	err = initClient(confPath)
	if err != nil {
		fmt.Println("init connection error: ", err)
		os.Exit(0)
	}

	if len(args) == 7 {
		// parse lt
		ltInt, err := strconv.ParseInt(args[5], 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse lt as integer number: `%s`. %s", args[2], err)
		}
		lt = tonlib.JSONInt64(ltInt)
		hash = args[6]
	} else {
		accState, err := tonClient.GetAccountState(*addr)
		if err != nil {
			log.Fatalf("Failed to get account state: %#v", err)
		}

		lt = accState.LastTransactionId.Lt
		hash = accState.LastTransactionId.Hash
	}

	// prepare input key
	pKey := tonlib.TONPrivateKey{PublicKey: publicKey, Secret: secret}
	inputKey := tonlib.InputKey{
		Type:          "inputKeyRegular",
		LocalPassword: base64.StdEncoding.EncodeToString(tonlib.SecureBytes(password)),
		Key:           pKey,
	}

	txs, err := tonClient.RawGetTransactions(*addr, *tonlib.NewInternalTransactionId(hash, lt), inputKey)
	if err != nil {
		fmt.Println("get wallet address error: ", err)
		os.Exit(0)
	}
	for i := 0; i < len(txs.Transactions); i++ {
		tx := txs.Transactions[i]
		fmt.Printf("Got a result: data: %v; type: %v; transaction id: %v; fee: %v; inMsg: %v;\n", tx.Data, tx.Type, tx.TransactionId, tx.Fee, tx.InMsg)
		for j := 0; j < len(tx.OutMsgs); j++ {
			msg := tx.OutMsgs[j]
			fmt.Printf("Got a out msg: message: %v; \n", msg.Message)
		}
	}

	fmt.Printf("Got a result: transactions count: %v. Errors: %v. \n", len(txs.Transactions), err)
}
