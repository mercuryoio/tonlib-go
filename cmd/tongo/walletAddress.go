package main

import (
	"errors"
	"fmt"
	"github.com/mercuryoio/tonlib-go"
	"github.com/spf13/cobra"
	"os"
)

var walletAddressCmd = &cobra.Command{
	Use:   "walletAddress",
	Short: "Get wallet addresses",
	Long: `Get wallet addresses command. It contains 4 attributes:
- path2configfile. see tonlib.config.json.example
- public key
- secret
- password
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
	Run: walletAddress,
}

func walletAddress(cmd *cobra.Command, args []string) {
	err := initClient(args[0])
	if err != nil {
		fmt.Println("init connection error: ", err)
		os.Exit(0)
	}

	pKey := &tonlib.TONPrivateKey{PublicKey: args[1], Secret: args[2]}
	err = tonClient.InitWallet(pKey, []byte(args[3]))
	if err != nil {
		fmt.Println("init wallet error: ", err)
		os.Exit(0)
	}

	addr, err := tonClient.WalletGetAddress(pKey.PublicKey)
	if err != nil {
		fmt.Println("get wallet address error: ", err)
		os.Exit(0)
	}
	addrUP, err := tonClient.UnpackAccountAddress(addr.AccountAddress)
	if err != nil {
		fmt.Println("unpack wallet address error: ", err)
		os.Exit(0)
	}
	addrUP.Bounceable = false
	address, err := tonClient.PackAccountAddress(addrUP, addr.AccountAddress)
	if err != nil {
		fmt.Println("pack wallet address error: ", err)
		os.Exit(0)
	}

	fmt.Printf("Got a result! Origin address :%v; Unpack address: %v. Bounceable address(for first incoming): %v. \n", addr.AccountAddress, addrUP.Addr, address)
}
