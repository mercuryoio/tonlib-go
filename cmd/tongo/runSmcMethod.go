package main

import (
	"errors"
	"fmt"
	"github.com/mercuryoio/tonlib-go"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var runSmcMethodCmd = &cobra.Command{
	Use:   "runSmcMethod",
	Short: "Get raw account state",
	Long: `Get raw account state command. It contains two attributes:
- path2configfile. see tonlib.config.json.example
- smc address % - in the beginig - required!
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("you have to use more than two args for this commaond \n")
		}
		_, err := os.Stat(args[0])
		if err != nil {
			errors.New("please choose config path")
		}

		// check address
		if len(args[1]) < 2{
			fmt.Printf("failed to parse adress: %s/n", args[1])
			os.Exit(1)
		}
		if args[1][0:1] != "%"{
			fmt.Printf("adress have to start with `%` but got: `%s` /n", args[1])
			os.Exit(1)
		}
		return nil
	},
	Run: runSmcMethod,
}

func runSmcMethod(cmd *cobra.Command, args []string) {
	err := initClient(args[0])
	if err != nil {
		fmt.Println("init connection error: ", err)
		os.Exit(0)
	}

	// cut  % in the begining of address- escape symbol
	address := args[1][1:]

	// parse address
	// load adresss
	smcInfo, err := tonClient.SmcLoad(*tonlib.NewAccountAddress(address))
	fmt.Println("smcInfo: %#v, err: %v",smcInfo ,err)
	if err != nil{
		log.Fatal("Failed to SmcLoad: ", err)
	}

	// run smc
	res, err := tonClient.SmcRunGetMethod(
		smcInfo.Id,
		tonlib.NewSmcMethodIdName("active_election_id"),
		[]tonlib.TvmStackEntry{},
		)
	fmt.Println("run smc: %#v, err: %v",res ,err)
	if err != nil{
		log.Fatal("Failed to run smc method: ", err)
	}
	fmt.Printf("Got a result: SmcRunResult :%#v;. Errors: %v. \n", res, err)
}
