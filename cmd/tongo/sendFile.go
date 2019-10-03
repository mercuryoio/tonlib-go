package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var sendFileCmd = &cobra.Command{
	Use:   "sendFile",
	Short: "Send boc file command",
	Long: `Send message command. It contains four attributes:
- path2configfile. see tonlib.config.json.example
- initialAddress
- destinationAddress
- path2boc path to boc binary file 
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
	Run: sendFile,
}

func sendFile(cmd *cobra.Command, args []string) {
	err := initClient(args[0])
	if err != nil {
		fmt.Println("init connection error: ", err)
		os.Exit(0)
	}
	bocFile, err := ioutil.ReadFile(args[3])
	if err != nil {
		fmt.Println("boc file dosn't exist", err)
		os.Exit(0)
	}

	res, err := tonClient.SendMessage(args[1], args[2], string(bocFile))
	fmt.Printf("Got a result: %v. Errors: %v", res, err)
}
