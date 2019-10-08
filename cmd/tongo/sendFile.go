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
- destinationAddress
- path2boc path to boc binary file
- path2initialAccountState path to boc file with initial account state (optional) 
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return fmt.Errorf("you have to use minimum 3 args for this commaond \n")
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
	bocFile, err := ioutil.ReadFile(args[2])
	if err != nil {
		fmt.Println("boc file dosn't exist", err)
		os.Exit(0)
	}
	bocInitialStateFile := []byte{}
	if len(args) > 3 {
		bocInitialStateFile, err = ioutil.ReadFile(args[3])
		if err != nil {
			fmt.Println("boc file dosn't exist", err)
			os.Exit(0)
		}
	}

	res, err := tonClient.SendMessage(args[1], bocInitialStateFile, bocFile)
	fmt.Printf("Got a result: %v. Errors: %v", res, err)
}
