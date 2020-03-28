package main

import (
	"errors"
	"fmt"
	tonlib "github.com/varche1/tonlib-go/v2"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var runSmcMethodCmd = &cobra.Command{
	Use:   "runSmcMethod",
	Short: "Run smc method",
	Long: `Run smc method. It contains thre or more attributes:
- path2configfile. see tonlib.config.json.example
- smc address % - in the beginig - required!
- smc method name - required
- extra decimal args - as many as you want
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return fmt.Errorf("you have to use more than two args for this commaond \n")
		}
		_, err := os.Stat(args[0])
		if err != nil {
			errors.New("please choose config path")
		}

		// check address
		if len(args[1]) < 2 {
			fmt.Printf("failed to parse adress: %s/n", args[1])
			os.Exit(1)
		}
		if args[1][0:1] != "%" {
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

	// parse address
	// cut  % in the begining of address- escape symbol
	address := args[1][1:]
	methodName := args[2]

	// parse extra params
	params := []tonlib.TvmStackEntry{}
	for _, arg := range (args[3:]) {
		params = append(params, tonlib.NewTvmStackEntryNumber(tonlib.NewTvmNumberDecimal(arg)))
	}

	// load adresss
	smcInfo, err := tonClient.SmcLoad(*tonlib.NewAccountAddress(address))
	if err != nil {
		log.Fatal("Failed to SmcLoad: ", err)
	}
	fmt.Printf("smcInfo: %#v \n", smcInfo)

	// run smc
	res, err := tonClient.SmcRunGetMethod(
		smcInfo.Id,
		tonlib.NewSmcMethodIdName(methodName),
		params,
	)
	if err != nil {
		log.Fatal("Failed to run smc method: ", err)
	}
	fmt.Printf("Got a result: SmcRunResult :%#v;. Errors: %v. \n", res, err)
}
