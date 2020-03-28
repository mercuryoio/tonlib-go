package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/varche1/tonlib-go"
	"github.com/spf13/cobra"
	"os"
)

var exportPKCmd = &cobra.Command{
	Use:   "exportPK",
	Short: "Export exists local private key",
	Long: `"Export exists local private key command. It contains two attributes:
- path2configfile. see tonlib.config.json.example
- public key
- secret
- password

And returns words list of exporting mnemonics between [] and error.
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return fmt.Errorf("you have to use minimum three args for this commaond \n")
		}
		_, err := os.Stat(args[0])
		if err != nil {
			errors.New("please choose config path")
		}
		return nil
	},
	Run: exportPK,
}

func exportPK(cmd *cobra.Command, args []string) {
	confPath := args[0]
	publicKey := args[1]
	secret := args[2]
	password := args[3]
	err := initClient(confPath)
	locPass := tonlib.SecureBytes(password)

	if err != nil {
		fmt.Println("init connection error: ", err)
		os.Exit(0)
	}
	pKey := tonlib.TONPrivateKey{PublicKey: publicKey, Secret: secret}

	result, err := tonClient.ExportPemKey(&tonlib.InputKey{
		"inputKeyRegular",
		base64.StdEncoding.EncodeToString(locPass),
		pKey,
	}, &locPass)

	fmt.Printf("Got a result: %#v. Errors: %v. \n", result, err)
}
