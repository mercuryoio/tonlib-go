package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/mercuryoio/tonlib-go"
	"github.com/spf13/cobra"
)

var tonClient *tonlib.Client

func init() {
	rootCmd.AddCommand(sendMessageCmd, sendFileCmd, createPKCmd, rawAccountStateCmd,
		walletAddressCmd, walletStateCmd, sendGrammCmd, deletePKCmd, exportPKCmd, transactionsCmd)
}

func initClient(configPath string) error {
	cnf, err := tonlib.ParseConfigFile(configPath)
	if err != nil {
		return err
	}
	tonClient, err = tonlib.NewClient(cnf, tonlib.Config{})
	if err != nil {
		fmt.Errorf("Init client error: %v. ", err)
	}
	return nil
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
