package tonlib

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestClient_GetAccountState(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("Config file not found", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Fatal("Init client error", err)
	}
	defer cln.Destroy()
	_, err = cln.GetAccountState(TEST_ADDRESS)
	if err != nil {
		t.Fatal("Ton get account state error", err)
	}
}

func TestClient_GetAccountTransactions(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("Config file not found", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Fatal("Init client error", err)
	}
	defer cln.Destroy()
	state, err := cln.GetAccountState(TEST_ADDRESS)
	if err != nil {
		t.Fatal("Get state error error", err)
	}
	_, err = cln.GetAccountTransactions(TEST_ADDRESS, state.LastTransactionID.Lt, state.LastTransactionID.Hash)
	if err != nil {
		t.Fatal("Ton get account txs error", err)
	}
}

func TestClient_InitWallet(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("Config file not found", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Fatal("Init client error", err)
	}
	defer cln.Destroy()
	pKey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton create key for init wallet error", err)
	}
	err = cln.InitWallet(pKey, []byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton init wallet error", err)
	}
}

func TestClient_WalletGetAddress(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("Config file not found", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Fatal("Init client error", err)
	}
	defer cln.Destroy()
	pKey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton create key for get wallet address error", err)
	}
	err = cln.InitWallet(pKey, []byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton init wallet for get wallet address error", err)
	}
	_, err = cln.WalletGetAddress(pKey.PublicKey)
	if err != nil {
		t.Fatal("Ton get wallet address error", err)
	}
}

func TestClient_WalletSendGRAMM2Address(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("Config file not found", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Fatal("Init client error", err)
	}
	defer cln.Destroy()
	pKey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton create key for wallet send gramms error", err)
	}
	err = cln.InitWallet(pKey, []byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton init wallet for wallet send gramms error", err)
	}
	address, err := cln.WalletGetAddress(pKey.PublicKey)
	if err != nil {
		t.Fatal("Ton get address for wallet send grams error", err)
	}
	_, err = cln.WalletSendGRAMM2Address(pKey, []byte(TEST_PASSWORD), address.AccountAddress, TEST_ADDRESS, TEST_AMOUNT)
	if err != nil {
		t.Fatal("Ton wallet send gramms error", err)
	}
}

func TestClient_SendGRAMM2Address(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("Config file not found", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Fatal("Init client error", err)
	}
	defer cln.Destroy()
	pKey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton create key for send grams error", err)
	}
	err = cln.InitWallet(pKey, []byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton init wallet for send gramms error", err)
	}
	address, err := cln.WalletGetAddress(pKey.PublicKey)
	if err != nil {
		t.Fatal("Ton get address for send grams error", err)
	}
	_, err = cln.SendGrams2Address(pKey, []byte(TEST_PASSWORD), address.AccountAddress, TEST_ADDRESS, TEST_AMOUNT, "")
	if err != nil && err.Error() != "Error ton send gramms. Code 500. Message NOT_ENOUGH_FUNDS. " {
		t.Fatal("Ton send grams error", err)
	}
}

//func TestClient_CreateQuery4SendGRAMM2Address(t *testing.T) {
//	cnf, err := ParseConfigFile("./tonlib.config.json.example")
//	if err != nil {
//		t.Fatal("Config file not found", err)
//	}
//	cln, err := NewClient(cnf, Config{})
//	if err != nil {
//		t.Fatal("Init client error", err)
//	}
//	defer cln.Destroy()
//	pKey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
//	if err != nil {
//		t.Fatal("Ton create key for send grams error", err)
//	}
//	err = cln.InitWallet(pKey, []byte(TEST_PASSWORD))
//	if err != nil {
//		t.Fatal("Ton init wallet for send gramms error", err)
//	}
//	address, err := cln.WalletGetAddress(pKey.PublicKey)
//	if err != nil {
//		t.Fatal("Ton get address for send grams error", err)
//	}
//	_, err = cln.CreateQuery4SendGrams2Address(pKey, []byte(TEST_PASSWORD), address.AccountAddress, TEST_ADDRESS, TEST_AMOUNT, "", 0, true)
//	if err != nil && err.Error() != "Error ton create query for sending grams. Code 500. Message NOT_ENOUGH_FUNDS. " {
//		t.Fatal("Ton send grams error", err)
//	}
//	t.Fail()
//}

func TestClient_CreateAndSendMessage(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("Config file not found", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Fatal("Init client error", err)
	}
	defer cln.Destroy()
	pKey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton create key for send grams error", err)
	}
	err = cln.InitWallet(pKey, []byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton init wallet for send gramms error", err)
	}
	address, err := cln.WalletGetAddress(pKey.PublicKey)
	if err != nil {
		t.Fatal("Ton get address for send grams error", err)
	}
	bocFile, err := ioutil.ReadFile("./testgiver-query.boc")
	if err != nil {
		fmt.Println("boc file dosn't exist", err)
		os.Exit(0)
	}

	_, err = cln.CreateAndSendMessage(address.AccountAddress, []byte{}, bocFile)
	if err != nil {
		t.Fatal("Ton send gramms error", err)
	}
}
