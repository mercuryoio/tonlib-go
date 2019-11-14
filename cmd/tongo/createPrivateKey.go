package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var createPKCmd = &cobra.Command{
	Use:   "createPK",
	Short: "Create new private key",
	Long: `Create new private key command. It contains three attributes:
- path2configfile. see tonlib.config.json.example
- localPassword local password for key
- mnemonicPassword password for mnemonics when you're exporting key. it's not required
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("you have to use minimum two args for this commaond \n")
		}
		_, err := os.Stat(args[0])
		if err != nil {
			errors.New("please choose config path")
		}
		return nil
	},
	Run: createPK,
}

func createPK(cmd *cobra.Command, args []string) {
	err := initClient(args[0])
	if err != nil {
		fmt.Println("init connection error: ", err)
		os.Exit(0)
	}
	mnPass := ""
	if len(args) > 2 {
		mnPass = args[2]
	}
	pKey, err := tonClient.CreatePrivateKey([]byte(args[1]), []byte(mnPass))
	fmt.Printf("Got a result: publicKey :%v; secret: %v. Errors: %v. \n", pKey.PublicKey, pKey.Secret, err)
}
