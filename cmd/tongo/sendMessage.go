package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var sendMessageCmd = &cobra.Command{
	Use:   "sendMessage",
	Short: "Send short message command",
	Long: `Send message command. It contains three attributes:
- path2configfile. see tonlib.config.json.example
- destinationAddress
- data in boc format
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 3 {
			return fmt.Errorf("you have to use three args for this commaond \n")
		}
		_, err := os.Stat(args[0])
		if err != nil {
			errors.New("please choose config path")
		}
		return nil
	},
	Run: sendMessage,
}

func sendMessage(cmd *cobra.Command, args []string) {
	err := initClient(args[0])
	if err != nil {
		fmt.Println("init connection error: ", err)
		os.Exit(0)
	}
	res, err := tonClient.SendMessage(args[1], []byte{}, []byte(args[2]))
	fmt.Printf("Got a result: %v. Errors: %v", res, err)
}
