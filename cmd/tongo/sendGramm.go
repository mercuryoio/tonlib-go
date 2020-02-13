package main

import (
	"errors"
	"fmt"
	"github.com/mercuryoio/tonlib-go"
	"github.com/spf13/cobra"
	"os"
)

var sendGrammCmd = &cobra.Command{
	Use:   "sendGramm",
	Short: "Send gramm from local account to destination command",
	Long: `Send gramm command. It contains four attributes:
- path2configfile. see tonlib.config.json.example
- public key
- secret
- password
- addressDestination
- amount
- message for destination. not required
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 6 {
			return fmt.Errorf("you have to use minimum six args for this commaond \n")
		}
		_, err := os.Stat(args[0])
		if err != nil {
			errors.New("please choose config path")
		}
		return nil
	},
	Run: sendGramm,
}

func sendGramm(cmd *cobra.Command, args []string) {
	confPath := args[0]
	publicKey := args[1]
	secret := args[2]
	//password := args[3]
	//destinationAddr := args[4]
	//// parse amount
	//amount, err := strconv.ParseInt(args[5], 10, 64)
	//if err != nil {
	//	log.Fatalf("failed to parse amount argument: %s as int. err: %s. ", args[5], err)
	//}

	//message := ""
	//if len(args) > 6 {
	//	message = args[6]
	//}

	err := initClient(confPath)
	if err != nil {
		fmt.Println("init connection error: ", err)
		os.Exit(0)
	}
	pKey := tonlib.TONPrivateKey{PublicKey: publicKey, Secret: secret}

	// prepare input key
	//inputKey := tonlib.InputKey{
	//	Type: "inputKeyRegular",
	//	LocalPassword: base64.StdEncoding.EncodeToString(tonlib.SecureBytes(password)),
	//	Key: pKey,
	//}
	//_, err = tonClient.WalletInit(&inputKey)
	//if err != nil {
	//	fmt.Println("init wallet error: ", err)
	//	os.Exit(0)
	//}

	// get wallet adress info
	addr, err := tonClient.GetAccountAddress(tonlib.NewWalletInitialAccountState(pKey.PublicKey), 0)
	if err != nil {
		fmt.Println("get wallet address error: ", err, addr)
		os.Exit(0)
	}

	// send grams
	//sendResult, err := tonClient.GenericSendGrams(
	//	true,
	//	tonlib.JSONInt64(amount),
	//	tonlib.NewAccountAddress(destinationAddr),
	//	[]byte(message),
	//	&inputKey,
	//	addr,
	//	5,
	//)
	//fmt.Printf("Got a result: hash %v. Errors: %v \n", sendResult, err)
}
