package tonlib

import (
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
		t.Errorf("Config file not found: %v. ", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Errorf("Init client error: %v. ", err)
	}
	defer cln.Destroy()
	_, err = cln.GetAccountState(TEST_ADDRESS)
	if err != nil {
		t.Errorf("Ton get account state error: %v. ", err)
	}
}

func TestClient_GetAccountTransactions(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Errorf("Config file not found: %v. ", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Errorf("Init client error: %v. ", err)
	}
	defer cln.Destroy()
	_, err = cln.GetAccountTransactions(TEST_ADDRESS, TEST_TX_LT, TEST_TX_HASH)
	if err != nil {
		t.Errorf("Ton get account txs error: %v. ", err)
	}
}

func TestClient_CreatePrivateKey(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Errorf("Config file not found: %v. ", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Errorf("Init client error: %v. ", err)
	}
	defer cln.Destroy()
	_, err = cln.CreatePrivateKey([]byte(TEST_PASSWORD))
	if err != nil {
		t.Errorf("Ton create key error: %v. ", err)
	}
}

//func TestClient_ChangeLocalPassword(t *testing.T) {
//	cln, err := NewClient(cnf, Config{})
//	if err != nil {
//		t.Errorf("Init client error: %v. ", err)
//	}
//	defer cln.Destroy()
//	pKey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD))
//	if err != nil {
//		t.Errorf("Ton create key for change password error: %v. ", err)
//	}
//	_, err = cln.ChangeLocalPassword(pKey, []byte(TEST_PASSWORD), []byte(TEST_CHANGE_PASSWORD))
//	if err != nil {
//		t.Errorf("Ton change key password error: %v. ", err)
//	}
//}

func TestClient_InitWallet(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Errorf("Config file not found: %v. ", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Errorf("Init client error: %v. ", err)
	}
	defer cln.Destroy()
	pKey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD))
	if err != nil {
		t.Errorf("Ton create key for init wallet error: %v. ", err)
	}
	err = cln.InitWallet(pKey, []byte(TEST_PASSWORD))
	if err != nil {
		t.Errorf("Ton init wallet error: %v. ", err)
	}
}

func TestClient_WalletGetAddress(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Errorf("Config file not found: %v. ", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Errorf("Init client error: %v. ", err)
	}
	defer cln.Destroy()
	pKey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD))
	if err != nil {
		t.Errorf("Ton create key for get wallet address error: %v. ", err)
	}
	err = cln.InitWallet(pKey, []byte(TEST_PASSWORD))
	if err != nil {
		t.Errorf("Ton init wallet for get wallet address error: %v. ", err)
	}
	_, err = cln.WalletGetAddress(pKey.PublicKey)
	if err != nil {
		t.Errorf("Ton get wallet address error: %v. ", err)
	}
}

func TestClient_WalletSendGRAMM2Address(t *testing.T) {
	cnf, err := ParseConfigFile("./tonlib.config.json.example")
	if err != nil {
		t.Errorf("Config file not found: %v. ", err)
	}
	cln, err := NewClient(cnf, Config{})
	if err != nil {
		t.Errorf("Init client error: %v. ", err)
	}
	defer cln.Destroy()
	pKey, err := cln.CreatePrivateKey([]byte(TEST_PASSWORD))
	if err != nil {
		t.Errorf("Ton create key for init wallet error: %v. ", err)
	}
	err = cln.InitWallet(pKey, []byte(TEST_PASSWORD))
	if err != nil {
		t.Errorf("Ton init wallet for send gramms error: %v. ", err)
	}
	address, err := cln.WalletGetAddress(pKey.PublicKey)
	if err != nil {
		t.Errorf("Ton get address for send grams error: %v. ", err)
	}
	err = cln.WalletSendGRAMM2Address(pKey, []byte(TEST_PASSWORD), address.AccountAddress, TEST_ADDRESS, TEST_AMOUNT)
	if err != nil {
		t.Errorf("Ton send gramms error: %v. ", err)
	}
}
