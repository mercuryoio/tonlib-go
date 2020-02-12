package main

import (
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
- lt
- hash
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 4 {
			return fmt.Errorf("you have to use four args for this commaond \n")
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
	confPath := args[0]
	address := args[1]
	// parse lt
	lt, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil{
		log.Fatalf("Failed to parse lt as integer number: `%s`. %s", args[2], err)
	}
	hash := args[3]
	err = initClient(confPath)
	if err != nil {
		fmt.Println("init connection error: ", err)
		os.Exit(0)
	}

	txs, err := tonClient.RawGetTransactions(*tonlib.NewAccountAddress(address), *tonlib.NewInternalTransactionId(hash, tonlib.JSONInt64(lt)), tonlib.InputKey{})
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
