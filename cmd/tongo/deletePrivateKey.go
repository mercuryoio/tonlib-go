package main

import (
	"errors"
	"fmt"
	"github.com/mercuryoio/tonlib-go"
	"github.com/spf13/cobra"
	"os"
)

var deletePKCmd = &cobra.Command{
	Use:   "deletePK",
	Short: "Delete exists local private key",
	Long: `"Delete exists local private key command. It contains two attributes:
- path2configfile. see tonlib.config.json.example
- public key
- secret
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 3 {
			return fmt.Errorf("you have to use four args for this commaond \n")
		}
		_, err := os.Stat(args[0])
		if err != nil {
			return errors.New("please choose config path")
		}
		return nil
	},
	Run: deletePK,
}

func deletePK(cmd *cobra.Command, args []string) {
	confPath := args[0]
	publicKey := args[1]
	secret := args[2]
	err := initClient(confPath)
	if err != nil {
		fmt.Println("init connection error: ", err)
		os.Exit(0)
	}
	pKey := tonlib.NewKey(publicKey, secret)
	ok, err := tonClient.DeleteKey(*pKey)
	fmt.Printf("Got a result: %#v,  error: %v. \n", ok, err)
}
