package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
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
	lt := args[2]
	hash := args[3]
	err := initClient(confPath)
	if err != nil {
		fmt.Println("init connection error: ", err)
		os.Exit(0)
	}

	txs, err := tonClient.GetAccountTransactions(address, lt, hash)
	if err != nil {
		fmt.Println("get wallet address error: ", err)
		os.Exit(0)
	}
	for i := 0; i < len(txs.Transactions); i++ {
		tx := txs.Transactions[i]
		fmt.Printf("Got a result: data: %v; type: %v; transaction id: %v; fee: %v; inMsg: %v; outMsg: %v \n", tx.Data, tx.Type, tx.TransactionID, tx.Fee, tx.InMsg, tx.OutMsgs)
	}

	fmt.Printf("Got a result: transactions count: %v. Errors: %v. \n", len(txs.Transactions), err)
}
