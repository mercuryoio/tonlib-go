package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rawAccountStateCmd = &cobra.Command{
	Use:   "rawAccountState",
	Short: "Get raw account state",
	Long: `Get raw account state command. It contains two attributes:
- path2configfile. see tonlib.config.json.example
- account address
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("you have to use two args for this commaond \n")
		}
		_, err := os.Stat(args[0])
		if err != nil {
			errors.New("please choose config path")
		}
		return nil
	},
	Run: rawAccountState,
}

func rawAccountState(cmd *cobra.Command, args []string) {
	err := initClient(args[0])
	if err != nil {
		fmt.Println("init connection error: ", err)
		os.Exit(0)
	}
	res, err := tonClient.GetAccountState(args[1])
	fmt.Printf("Got a result: balance :%v; last transaction id: %v. Errors: %v. \n", res.Balance, res.LastTransactionID, err)
}
