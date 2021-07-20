package main

import (
	"fmt"
	tonlib "github.com/mercuryoio/tonlib-go/v2"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

const walletID = 698983191

var tonClient *tonlib.Client

func init() {
	rootCmd.AddCommand(sendMessageCmd, sendFileCmd, createPKCmd, rawAccountStateCmd, walletAddressCmd, walletStateCmd,
		sendGrammCmd, deletePKCmd, exportPKCmd, transactionsCmd, estimateFeeCmd, runSmcMethodCmd)
}

func initClient(configPath string) error {
	options, err := tonlib.ParseConfigFile(configPath)
	if err != nil {
		return err
	}

	// make req
	req := tonlib.TonInitRequest{
		"init",
		*options,
	}

	tonClient, err = tonlib.NewClient(&req, tonlib.Config{}, 10, true, 9)
	if err != nil {
		err = fmt.Errorf("Init client error: %v. ", err)
	}
	return err
}

var rootCmd = &cobra.Command{
	Use:   "help",
	Short: `Ton console tool used tonlib`,
	Long:  ``,
}

// Execute CLI application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		<-c
		os.Exit(0)
	}()

	Execute()
}
