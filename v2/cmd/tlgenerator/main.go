package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"os/signal"
)

var syncWithTlCmd = &cobra.Command{
	Use:   "tlgenerator",
	Short: "tlgenerator /path/to/tl/file.tl",
	Long:  `tlgenerator`,
	Run:   syncWtithTl,
}

func syncWtithTl(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		log.Fatal("require only 1 argument - path to tl file")
		return
	}
	tlFilePath := args[0]

	fmt.Println("Open file: ", tlFilePath)
	tlFile, err := os.OpenFile(tlFilePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer tlFile.Close()

	fmt.Println("Read file")
	tlReader := bufio.NewReader(tlFile)

	// eide file string by strign and parse it
	err, entities, interfaces, enums := parseTlFile(tlReader)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Parsed entities: %d, interfaces: %d, enums: %d. ", len(*entities), len(*interfaces), len(*enums))

	// generate gp`s structs based on parsed entities
	structsContent, methodsContetnt := generateStructsFromTnEntities("tonlib", entities, interfaces, enums)
	structsFilePath := "./structs.go"
	methodsFilePath := "./methods.go"

	// delete and create new files
	// structures file
	_ = os.Remove(structsFilePath)
	structsFile, err := os.OpenFile(structsFilePath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer structsFile.Close()
	if err != nil {
		log.Fatalf("error openning file %v", err)
	}
	wgo := bufio.NewWriter(structsFile)
	_, err = wgo.Write([]byte(*structsContent))
	if err != nil {
		log.Fatal(err)
	}

	// methods file
	_ = os.Remove(methodsFilePath)
	methodsFile, err := os.OpenFile(methodsFilePath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer methodsFile.Close()

	if err != nil {
		log.Fatalf("error openning file %v", err)
	}
	wgo = bufio.NewWriter(methodsFile)
	_, err = wgo.Write([]byte(*methodsContetnt))
	if err != nil {
		log.Fatal(err)
	}

	// format files
	cmd1 := exec.Command("gofmt", "-w", methodsFilePath)
	_ = cmd1.Run()
	cmd2 := exec.Command("gofmt", "-w", structsFilePath)
	_ = cmd2.Run()

}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		<-c
		os.Exit(0)
	}()

	err := syncWithTlCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
