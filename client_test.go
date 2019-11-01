package tonlib

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const (
	TEST_PASSWORD        = "test_password"
	TEST_CHANGE_PASSWORD = "new_password"
	TEST_ADDRESS         = "EQDfYZhDfNJ0EePoT5ibfI9oG9bWIU6g872oX5h9rL5PHY9a"
	TEST_TX_LT           = 289040000001
	TEST_TX_HASH         = "V6R8l0hTjpGb/HHHtDwrMk1KxTDLpfz5h7PINr1crp4="
	TEST_AMOUNT          = "100000000"
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

func TestClient_CreatePrivateKey(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("Config file not found", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Fatal("Init client error", err)
	}
	defer cln.Destroy()
	_, err = cln.CreatePrivateKey([]byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton create key error", err)
	}
}

func TestClient_DeletePrivateKey(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("Config file not found", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Fatal("Init client error", err)
	}
	defer cln.Destroy()
	pkey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton create key error", err)
	}
	if err = cln.DeletePrivateKey(pkey, []byte(TEST_PASSWORD)); err != nil {
		t.Fatal("Ton delete key error", err)
	}
}

func TestClient_ExportPrivateKey(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("Config file not found", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Fatal("Init client error", err)
	}
	defer cln.Destroy()
	pkey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton create key error", err)
	}
	if _, err = cln.ExportPrivateKey(pkey, []byte(TEST_PASSWORD)); err != nil {
		t.Fatal("Ton export private key error", err)
	}
}

func TestClient_ExportPemKey(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("Config file not found", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Fatal("Init client error", err)
	}
	defer cln.Destroy()
	pkey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton create key error", err)
	}
	if pem, err := cln.ExportPemKey(pkey, []byte(TEST_PASSWORD), []byte(TEST_PASSWORD)); err != nil || len(pem) == 0 {
		t.Fatal("Ton export pem key error", err)
	}
}

func TestClient_ExportEncryptedKey(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("Config file not found", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Fatal("Init client error", err)
	}
	defer cln.Destroy()
	pkey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton create key error", err)
	}
	if _, err = cln.ExportEncryptedKey(pkey, []byte(TEST_PASSWORD), []byte(TEST_PASSWORD)); err != nil {
		t.Fatal("Ton export pem key error", err)
	}
}

func TestClient_ChangeLocalPassword(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Fatal("Init client error", err)
	}
	defer cln.Destroy()
	pKey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton create key for change password error", err)
	}
	_, err = cln.ChangeLocalPassword(pKey, []byte(TEST_PASSWORD), []byte(TEST_CHANGE_PASSWORD))
	if err != nil {
		t.Fatal("Ton change key password error", err)
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
	pKey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD))
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
	pKey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD))
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
	pKey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD))
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
	pKey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD))
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
	_, err = cln.SendGRAMM2Address(pKey, []byte(TEST_PASSWORD), address.AccountAddress, TEST_ADDRESS, TEST_AMOUNT, "")
	if err != nil && err.Error() != "Error ton send gramms. Code 500. Message NOT_ENOUGH_FUNDS. " {
		t.Fatal("Ton send gramms error", err)
	}
}

func TestClient_ImportPemKey(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("Config file not found", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Fatal("Init client error", err)
	}
	defer cln.Destroy()
	pkey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton create key error", err)
	}
	key, err := cln.ExportPemKey(pkey, []byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton export pem key error", err)
	}

	_, err = cln.ImportPemKey(key, []byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton import pem key error", err)
	}
}

func TestClient_ImportEncryptedKey(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Fatal("Config file not found", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Fatal("Init client error", err)
	}
	defer cln.Destroy()
	pkey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton create key error", err)
	}
	key, err := cln.ExportEncryptedKey(pkey, []byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton export encrypted key error", err)
	}

	_, err = cln.ImportEncryptedKey(key, []byte(TEST_PASSWORD), []byte(TEST_PASSWORD))
	if err != nil {
		t.Fatal("Ton import encrypted key error", err)
	}
}

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
	pKey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD))
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
