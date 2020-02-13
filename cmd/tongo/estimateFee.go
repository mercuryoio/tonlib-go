package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/mercuryoio/tonlib-go"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

var estimateFeeCmd = &cobra.Command{
	Use:   "estimateFee",
	Short: "estimateFee from local account to destination command",
	Long: `estimateFee command. It contains four attributes:
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
	Run: estimateFee,
}

func estimateFee(cmd *cobra.Command, args []string) {
	confPath := args[0]
	publicKey := args[1]
	secret := args[2]
	password := args[3]
	destinationAddr := args[4]
	// parse amount
	amount, err := strconv.ParseInt(args[5], 10, 64)
	if err != nil {
		log.Fatalf("failed to parse amount argument: %s as int. err: %s. ", args[5], err)
	}

	message := ""
	if len(args) > 6 {
		message = args[6]
	}
	err = initClient(confPath)
	if err != nil {
		fmt.Println("init connection error: ", err)
		os.Exit(0)
	}
	pKey := tonlib.TONPrivateKey{PublicKey: publicKey, Secret: secret}

	// prepare input key
	inputKey := tonlib.InputKey{
		Type: "inputKeyRegular",
		LocalPassword: base64.StdEncoding.EncodeToString(tonlib.SecureBytes(password)),
		Key: pKey,
	}
	//_, err = tonClient.WalletInit(&inputKey)
	//if err != nil {
	//	fmt.Println("init wallet error: ", err)
	//	os.Exit(0)
	//}

	// get wallet adress info
	addr, err := tonClient.GetAccountAddress(tonlib.NewWalletInitialAccountState(pKey.PublicKey), 0)
	if err != nil {
		fmt.Println("get wallet address error: ", err)
		os.Exit(0)
	}

	// get query info
	//true,
	//	tonlib.JSONInt64(amount),
	//	tonlib.NewAccountAddress(destinationAddr),
	//	[]byte(message),
	//	&inputKey,
	//	addr,
	//	300, // ti

	msgAction := tonlib.NewActionMsg(
		true,
		[]tonlib.MsgMessage{*tonlib.NewMsgMessage(
			tonlib.JSONInt64(amount),
			tonlib.NewMsgDataText(message),
			tonlib.NewAccountAddress(destinationAddr),
			)},
		)
	queryInfo, err := tonClient.CreateQuery(
		msgAction,
		*addr,
		inputKey,
		300, // time out of sending money not executing request
	)
	fmt.Println(fmt.Sprintf("queryInfo: %#v. err: %#v. ", queryInfo, err))
	if err != nil{
		fmt.Printf("Failed to create query with  error: %v \n", err)
		os.Exit(1)
	}

	// get fee
	fees, err := tonClient.QueryEstimateFees(queryInfo.Id, false)
	fmt.Println(fmt.Sprintf("fees: %#v. err: %#v. ", fees, err))

}
