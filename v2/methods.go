package v2

import (
	"encoding/json"
	"fmt"
)

// Init
// @param options
func (client *Client) Init(options *Options) (*OptionsInfo, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type    string   `json:"@type"`
			Options *Options `json:"options"`
		}{
			Type:    "init",
			Options: options,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var optionsInfo OptionsInfo
	err = json.Unmarshal(result.Raw, &optionsInfo)
	return &optionsInfo, err

}

// Close
func (client *Client) Close() (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
		}{
			Type: "close",
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// OptionsSetconfig
// @param config
func (client *Client) OptionsSetconfig(config *Config) (*OptionsConfigInfo, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type   string  `json:"@type"`
			Config *Config `json:"config"`
		}{
			Type:   "options.setConfig",
			Config: config,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var optionsConfigInfo OptionsConfigInfo
	err = json.Unmarshal(result.Raw, &optionsConfigInfo)
	return &optionsConfigInfo, err

}

// OptionsValidateconfig
// @param config
func (client *Client) OptionsValidateconfig(config *Config) (*OptionsConfigInfo, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type   string  `json:"@type"`
			Config *Config `json:"config"`
		}{
			Type:   "options.validateConfig",
			Config: config,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var optionsConfigInfo OptionsConfigInfo
	err = json.Unmarshal(result.Raw, &optionsConfigInfo)
	return &optionsConfigInfo, err

}

// Createnewkey
// @param mnemonicPassword
// @param randomExtraSeed
// @param localPassword
func (client *Client) Createnewkey(mnemonicPassword *SecureBytes, randomExtraSeed *SecureBytes, localPassword *SecureBytes) (*Key, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type             string       `json:"@type"`
			MnemonicPassword *SecureBytes `json:"mnemonic_password"`
			RandomExtraSeed  *SecureBytes `json:"random_extra_seed"`
			LocalPassword    *SecureBytes `json:"local_password"`
		}{
			Type:             "createNewKey",
			MnemonicPassword: mnemonicPassword,
			RandomExtraSeed:  randomExtraSeed,
			LocalPassword:    localPassword,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var key Key
	err = json.Unmarshal(result.Raw, &key)
	return &key, err

}

// Deletekey
// @param key
func (client *Client) Deletekey(key *Key) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
			Key  *Key   `json:"key"`
		}{
			Type: "deleteKey",
			Key:  key,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// Deleteallkeys
func (client *Client) Deleteallkeys() (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
		}{
			Type: "deleteAllKeys",
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// Exportkey
// @param inputKey
func (client *Client) Exportkey(inputKey *InputKey) (*ExportedKey, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type     string    `json:"@type"`
			InputKey *InputKey `json:"input_key"`
		}{
			Type:     "exportKey",
			InputKey: inputKey,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var exportedKey ExportedKey
	err = json.Unmarshal(result.Raw, &exportedKey)
	return &exportedKey, err

}

// Exportpemkey
// @param inputKey
// @param keyPassword
func (client *Client) Exportpemkey(inputKey *InputKey, keyPassword *SecureBytes) (*ExportedPemKey, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type        string       `json:"@type"`
			InputKey    *InputKey    `json:"input_key"`
			KeyPassword *SecureBytes `json:"key_password"`
		}{
			Type:        "exportPemKey",
			InputKey:    inputKey,
			KeyPassword: keyPassword,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var exportedPemKey ExportedPemKey
	err = json.Unmarshal(result.Raw, &exportedPemKey)
	return &exportedPemKey, err

}

// Exportencryptedkey
// @param inputKey
// @param keyPassword
func (client *Client) Exportencryptedkey(inputKey *InputKey, keyPassword *SecureBytes) (*ExportedEncryptedKey, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type        string       `json:"@type"`
			InputKey    *InputKey    `json:"input_key"`
			KeyPassword *SecureBytes `json:"key_password"`
		}{
			Type:        "exportEncryptedKey",
			InputKey:    inputKey,
			KeyPassword: keyPassword,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var exportedEncryptedKey ExportedEncryptedKey
	err = json.Unmarshal(result.Raw, &exportedEncryptedKey)
	return &exportedEncryptedKey, err

}

// Importkey
// @param localPassword
// @param mnemonicPassword
// @param exportedKey
func (client *Client) Importkey(localPassword *SecureBytes, mnemonicPassword *SecureBytes, exportedKey *ExportedKey) (*Key, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type             string       `json:"@type"`
			LocalPassword    *SecureBytes `json:"local_password"`
			MnemonicPassword *SecureBytes `json:"mnemonic_password"`
			ExportedKey      *ExportedKey `json:"exported_key"`
		}{
			Type:             "importKey",
			LocalPassword:    localPassword,
			MnemonicPassword: mnemonicPassword,
			ExportedKey:      exportedKey,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var key Key
	err = json.Unmarshal(result.Raw, &key)
	return &key, err

}

// Importpemkey
// @param localPassword
// @param keyPassword
// @param exportedKey
func (client *Client) Importpemkey(localPassword *SecureBytes, keyPassword *SecureBytes, exportedKey *ExportedPemKey) (*Key, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type          string          `json:"@type"`
			LocalPassword *SecureBytes    `json:"local_password"`
			KeyPassword   *SecureBytes    `json:"key_password"`
			ExportedKey   *ExportedPemKey `json:"exported_key"`
		}{
			Type:          "importPemKey",
			LocalPassword: localPassword,
			KeyPassword:   keyPassword,
			ExportedKey:   exportedKey,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var keyDummy Key
	err = json.Unmarshal(result.Raw, &keyDummy)
	return &keyDummy, err

}

// Importencryptedkey
// @param localPassword
// @param keyPassword
// @param exportedEncryptedKey
func (client *Client) Importencryptedkey(localPassword *SecureBytes, keyPassword *SecureBytes, exportedEncryptedKey *ExportedEncryptedKey) (*Key, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type                 string                `json:"@type"`
			LocalPassword        *SecureBytes          `json:"local_password"`
			KeyPassword          *SecureBytes          `json:"key_password"`
			ExportedEncryptedKey *ExportedEncryptedKey `json:"exported_encrypted_key"`
		}{
			Type:                 "importEncryptedKey",
			LocalPassword:        localPassword,
			KeyPassword:          keyPassword,
			ExportedEncryptedKey: exportedEncryptedKey,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var keyDummy Key
	err = json.Unmarshal(result.Raw, &keyDummy)
	return &keyDummy, err

}

// Changelocalpassword
// @param inputKey
// @param newLocalPassword
func (client *Client) Changelocalpassword(inputKey *InputKey, newLocalPassword *SecureBytes) (*Key, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type             string       `json:"@type"`
			InputKey         *InputKey    `json:"input_key"`
			NewLocalPassword *SecureBytes `json:"new_local_password"`
		}{
			Type:             "changeLocalPassword",
			InputKey:         inputKey,
			NewLocalPassword: newLocalPassword,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var key Key
	err = json.Unmarshal(result.Raw, &key)
	return &key, err

}

// Encrypt
// @param decryptedData
// @param secret
func (client *Client) Encrypt(decryptedData *SecureBytes, secret *SecureBytes) (*Data, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type          string       `json:"@type"`
			DecryptedData *SecureBytes `json:"decrypted_data"`
			Secret        *SecureBytes `json:"secret"`
		}{
			Type:          "encrypt",
			DecryptedData: decryptedData,
			Secret:        secret,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var data Data
	err = json.Unmarshal(result.Raw, &data)
	return &data, err

}

// Decrypt
// @param encryptedData
// @param secret
func (client *Client) Decrypt(encryptedData *SecureBytes, secret *SecureBytes) (*Data, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type          string       `json:"@type"`
			EncryptedData *SecureBytes `json:"encrypted_data"`
			Secret        *SecureBytes `json:"secret"`
		}{
			Type:          "decrypt",
			EncryptedData: encryptedData,
			Secret:        secret,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var data Data
	err = json.Unmarshal(result.Raw, &data)
	return &data, err

}

// Kdf
// @param iterations
// @param password
// @param salt
func (client *Client) Kdf(iterations int32, password *SecureBytes, salt *SecureBytes) (*Data, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type       string       `json:"@type"`
			Iterations int32        `json:"iterations"`
			Password   *SecureBytes `json:"password"`
			Salt       *SecureBytes `json:"salt"`
		}{
			Type:       "kdf",
			Iterations: iterations,
			Password:   password,
			Salt:       salt,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var data Data
	err = json.Unmarshal(result.Raw, &data)
	return &data, err

}

// Unpackaccountaddress
// @param accountAddress
func (client *Client) Unpackaccountaddress(accountAddress string) (*UnpackedAccountAddress, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type           string `json:"@type"`
			AccountAddress string `json:"account_address"`
		}{
			Type:           "unpackAccountAddress",
			AccountAddress: accountAddress,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var unpackedAccountAddress UnpackedAccountAddress
	err = json.Unmarshal(result.Raw, &unpackedAccountAddress)
	return &unpackedAccountAddress, err

}

// Packaccountaddress
// @param accountAddress
func (client *Client) Packaccountaddress(accountAddress *UnpackedAccountAddress) (*AccountAddress, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type           string                  `json:"@type"`
			AccountAddress *UnpackedAccountAddress `json:"account_address"`
		}{
			Type:           "packAccountAddress",
			AccountAddress: accountAddress,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var accountAddressDummy AccountAddress
	err = json.Unmarshal(result.Raw, &accountAddressDummy)
	return &accountAddressDummy, err

}

// Getbip39hints
// @param prefix
func (client *Client) Getbip39hints(prefix string) (*Bip39Hints, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type   string `json:"@type"`
			Prefix string `json:"prefix"`
		}{
			Type:   "getBip39Hints",
			Prefix: prefix,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var bip39Hints Bip39Hints
	err = json.Unmarshal(result.Raw, &bip39Hints)
	return &bip39Hints, err

}

// RawGetaccountaddress
// @param inititalAccountState
func (client *Client) RawGetaccountaddress(inititalAccountState *RawInitialAccountState) (*AccountAddress, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type                 string                  `json:"@type"`
			InititalAccountState *RawInitialAccountState `json:"initital_account_state"`
		}{
			Type:                 "raw.getAccountAddress",
			InititalAccountState: inititalAccountState,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var accountAddress AccountAddress
	err = json.Unmarshal(result.Raw, &accountAddress)
	return &accountAddress, err

}

// RawGetaccountstate
// @param accountAddress
func (client *Client) RawGetaccountstate(accountAddress *AccountAddress) (*RawAccountState, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type           string          `json:"@type"`
			AccountAddress *AccountAddress `json:"account_address"`
		}{
			Type:           "raw.getAccountState",
			AccountAddress: accountAddress,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var rawAccountState RawAccountState
	err = json.Unmarshal(result.Raw, &rawAccountState)
	return &rawAccountState, err

}

// RawGettransactions
// @param accountAddress
// @param fromTransactionId
func (client *Client) RawGettransactions(accountAddress *AccountAddress, fromTransactionId *InternalTransactionId) (*RawTransactions, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type              string                 `json:"@type"`
			AccountAddress    *AccountAddress        `json:"account_address"`
			FromTransactionId *InternalTransactionId `json:"from_transaction_id"`
		}{
			Type:              "raw.getTransactions",
			AccountAddress:    accountAddress,
			FromTransactionId: fromTransactionId,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var rawTransactions RawTransactions
	err = json.Unmarshal(result.Raw, &rawTransactions)
	return &rawTransactions, err

}

// RawSendmessage
// @param body
func (client *Client) RawSendmessage(body []byte) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
			Body []byte `json:"body"`
		}{
			Type: "raw.sendMessage",
			Body: body,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var Ok Ok
	err = json.Unmarshal(result.Raw, &Ok)
	return &Ok, err

}

// RawCreateandsendmessage
// @param destination
// @param initialAccountState
// @param data
func (client *Client) RawCreateandsendmessage(destination *AccountAddress, initialAccountState []byte, data []byte) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type                string          `json:"@type"`
			Destination         *AccountAddress `json:"destination"`
			InitialAccountState []byte          `json:"initial_account_state"`
			Data                []byte          `json:"data"`
		}{
			Type:                "raw.createAndSendMessage",
			Destination:         destination,
			InitialAccountState: initialAccountState,
			Data:                data,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RawCreatequery
// @param destination
// @param initCode
// @param initData
// @param body
func (client *Client) RawCreatequery(destination *AccountAddress, initCode []byte, initData []byte, body []byte) (*QueryInfo, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type        string          `json:"@type"`
			Destination *AccountAddress `json:"destination"`
			InitCode    []byte          `json:"init_code"`
			InitData    []byte          `json:"init_data"`
			Body        []byte          `json:"body"`
		}{
			Type:        "raw.createQuery",
			Destination: destination,
			InitCode:    initCode,
			InitData:    initData,
			Body:        body,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var queryInfo QueryInfo
	err = json.Unmarshal(result.Raw, &queryInfo)
	return &queryInfo, err

}

// TestwalletInit
// @param privateKey
func (client *Client) TestwalletInit(privateKey *InputKey) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type       string    `json:"@type"`
			PrivateKey *InputKey `json:"private_key"`
		}{
			Type:       "testWallet.init",
			PrivateKey: privateKey,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// TestwalletGetaccountaddress
// @param inititalAccountState
func (client *Client) TestwalletGetaccountaddress(inititalAccountState *TestWalletInitialAccountState) (*AccountAddress, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type                 string                         `json:"@type"`
			InititalAccountState *TestWalletInitialAccountState `json:"initital_account_state"`
		}{
			Type:                 "testWallet.getAccountAddress",
			InititalAccountState: inititalAccountState,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var accountAddress AccountAddress
	err = json.Unmarshal(result.Raw, &accountAddress)
	return &accountAddress, err

}

// TestwalletGetaccountstate
// @param accountAddress
func (client *Client) TestwalletGetaccountstate(accountAddress *AccountAddress) (*TestWalletAccountState, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type           string          `json:"@type"`
			AccountAddress *AccountAddress `json:"account_address"`
		}{
			Type:           "testWallet.getAccountState",
			AccountAddress: accountAddress,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var testWalletAccountState TestWalletAccountState
	err = json.Unmarshal(result.Raw, &testWalletAccountState)
	return &testWalletAccountState, err

}

// TestwalletSendgrams
// @param amount
// @param message
// @param privateKey
// @param destination
// @param seqno
func (client *Client) TestwalletSendgrams(amount JSONInt64, message []byte, privateKey *InputKey, destination *AccountAddress, seqno int32) (*SendGramsResult, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type        string          `json:"@type"`
			Amount      JSONInt64       `json:"amount"`
			Message     []byte          `json:"message"`
			PrivateKey  *InputKey       `json:"private_key"`
			Destination *AccountAddress `json:"destination"`
			Seqno       int32           `json:"seqno"`
		}{
			Type:        "testWallet.sendGrams",
			Amount:      amount,
			Message:     message,
			PrivateKey:  privateKey,
			Destination: destination,
			Seqno:       seqno,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var sendGramsResult SendGramsResult
	err = json.Unmarshal(result.Raw, &sendGramsResult)
	return &sendGramsResult, err

}

// WalletInit
// @param privateKey
func (client *Client) WalletInit(privateKey *InputKey) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type       string    `json:"@type"`
			PrivateKey *InputKey `json:"private_key"`
		}{
			Type:       "wallet.init",
			PrivateKey: privateKey,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// WalletGetaccountaddress
// @param inititalAccountState
func (client *Client) WalletGetaccountaddress(inititalAccountState *WalletInitialAccountState) (*AccountAddress, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type                 string                     `json:"@type"`
			InititalAccountState *WalletInitialAccountState `json:"initital_account_state"`
		}{
			Type:                 "wallet.getAccountAddress",
			InititalAccountState: inititalAccountState,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var accountAddress AccountAddress
	err = json.Unmarshal(result.Raw, &accountAddress)
	return &accountAddress, err

}

// WalletGetaccountstate
// @param accountAddress
func (client *Client) WalletGetaccountstate(accountAddress *AccountAddress) (*WalletAccountState, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type           string          `json:"@type"`
			AccountAddress *AccountAddress `json:"account_address"`
		}{
			Type:           "wallet.getAccountState",
			AccountAddress: accountAddress,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var walletAccountState WalletAccountState
	err = json.Unmarshal(result.Raw, &walletAccountState)
	return &walletAccountState, err

}

// WalletSendgrams
// @param validUntil
// @param amount
// @param message
// @param privateKey
// @param destination
// @param seqno
func (client *Client) WalletSendgrams(validUntil int64, amount JSONInt64, message []byte, privateKey *InputKey, destination *AccountAddress, seqno int32) (*SendGramsResult, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type        string          `json:"@type"`
			ValidUntil  int64           `json:"valid_until"`
			Amount      JSONInt64       `json:"amount"`
			Message     []byte          `json:"message"`
			PrivateKey  *InputKey       `json:"private_key"`
			Destination *AccountAddress `json:"destination"`
			Seqno       int32           `json:"seqno"`
		}{
			Type:        "wallet.sendGrams",
			ValidUntil:  validUntil,
			Amount:      amount,
			Message:     message,
			PrivateKey:  privateKey,
			Destination: destination,
			Seqno:       seqno,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var sendGramsResult SendGramsResult
	err = json.Unmarshal(result.Raw, &sendGramsResult)
	return &sendGramsResult, err

}

// WalletV3Getaccountaddress
// @param inititalAccountState
func (client *Client) WalletV3Getaccountaddress(inititalAccountState *WalletV3InitialAccountState) (*AccountAddress, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type                 string                       `json:"@type"`
			InititalAccountState *WalletV3InitialAccountState `json:"initital_account_state"`
		}{
			Type:                 "wallet.v3.getAccountAddress",
			InititalAccountState: inititalAccountState,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var accountAddress AccountAddress
	err = json.Unmarshal(result.Raw, &accountAddress)
	return &accountAddress, err

}

// TestgiverGetaccountstate
func (client *Client) TestgiverGetaccountstate() (*TestGiverAccountState, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
		}{
			Type: "testGiver.getAccountState",
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var testGiverAccountState TestGiverAccountState
	err = json.Unmarshal(result.Raw, &testGiverAccountState)
	return &testGiverAccountState, err

}

// TestgiverGetaccountaddress
func (client *Client) TestgiverGetaccountaddress() (*AccountAddress, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
		}{
			Type: "testGiver.getAccountAddress",
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var accountAddress AccountAddress
	err = json.Unmarshal(result.Raw, &accountAddress)
	return &accountAddress, err

}

// TestgiverSendgrams
// @param destination
// @param seqno
// @param amount
// @param message
func (client *Client) TestgiverSendgrams(destination *AccountAddress, seqno int32, amount JSONInt64, message []byte) (*SendGramsResult, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type        string          `json:"@type"`
			Destination *AccountAddress `json:"destination"`
			Seqno       int32           `json:"seqno"`
			Amount      JSONInt64       `json:"amount"`
			Message     []byte          `json:"message"`
		}{
			Type:        "testGiver.sendGrams",
			Destination: destination,
			Seqno:       seqno,
			Amount:      amount,
			Message:     message,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var sendGramsResult SendGramsResult
	err = json.Unmarshal(result.Raw, &sendGramsResult)
	return &sendGramsResult, err

}

// Sync
func (client *Client) Sync() (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
		}{
			Type: "sync",
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GenericGetaccountstate
// @param accountAddress
func (client *Client) GenericGetaccountstate(accountAddress *AccountAddress) (*GenericAccountState, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type           string          `json:"@type"`
			AccountAddress *AccountAddress `json:"account_address"`
		}{
			Type:           "generic.getAccountState",
			AccountAddress: accountAddress,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var genericAccountState GenericAccountState
	err = json.Unmarshal(result.Raw, &genericAccountState)
	return &genericAccountState, err

}

// GenericSendgrams
// @param privateKey
// @param source
// @param destination
// @param amount
// @param timeout
// @param allowSendToUninited
// @param message
func (client *Client) GenericSendgrams(privateKey *InputKey, source *AccountAddress, destination *AccountAddress, amount JSONInt64, timeout int32, allowSendToUninited bool, message []byte) (*SendGramsResult, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type                string          `json:"@type"`
			PrivateKey          *InputKey       `json:"private_key"`
			Source              *AccountAddress `json:"source"`
			Destination         *AccountAddress `json:"destination"`
			Amount              JSONInt64       `json:"amount"`
			Timeout             int32           `json:"timeout"`
			AllowSendToUninited bool            `json:"allow_send_to_uninited"`
			Message             []byte          `json:"message"`
		}{
			Type:                "generic.sendGrams",
			PrivateKey:          privateKey,
			Source:              source,
			Destination:         destination,
			Amount:              amount,
			Timeout:             timeout,
			AllowSendToUninited: allowSendToUninited,
			Message:             message,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var sendGramsResult SendGramsResult
	err = json.Unmarshal(result.Raw, &sendGramsResult)
	return &sendGramsResult, err

}

// GenericCreatesendgramsquery
// @param message
// @param privateKey
// @param source
// @param destination
// @param amount
// @param timeout
// @param allowSendToUninited
func (client *Client) GenericCreatesendgramsquery(message []byte, privateKey *InputKey, source *AccountAddress, destination *AccountAddress, amount JSONInt64, timeout int32, allowSendToUninited bool) (*QueryInfo, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type                string          `json:"@type"`
			Message             []byte          `json:"message"`
			PrivateKey          *InputKey       `json:"private_key"`
			Source              *AccountAddress `json:"source"`
			Destination         *AccountAddress `json:"destination"`
			Amount              JSONInt64       `json:"amount"`
			Timeout             int32           `json:"timeout"`
			AllowSendToUninited bool            `json:"allow_send_to_uninited"`
		}{
			Type:                "generic.createSendGramsQuery",
			Message:             message,
			PrivateKey:          privateKey,
			Source:              source,
			Destination:         destination,
			Amount:              amount,
			Timeout:             timeout,
			AllowSendToUninited: allowSendToUninited,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var queryInfo QueryInfo
	err = json.Unmarshal(result.Raw, &queryInfo)
	return &queryInfo, err

}

// QuerySend
// @param id
func (client *Client) QuerySend(id int64) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
			Id   int64  `json:"id"`
		}{
			Type: "query.send",
			Id:   id,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// QueryForget
// @param id
func (client *Client) QueryForget(id int64) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
			Id   int64  `json:"id"`
		}{
			Type: "query.forget",
			Id:   id,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// QueryEstimatefees
// @param id
// @param ignoreChksig
func (client *Client) QueryEstimatefees(id int64, ignoreChksig bool) (*QueryFees, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type         string `json:"@type"`
			Id           int64  `json:"id"`
			IgnoreChksig bool   `json:"ignore_chksig"`
		}{
			Type:         "query.estimateFees",
			Id:           id,
			IgnoreChksig: ignoreChksig,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var queryFees QueryFees
	err = json.Unmarshal(result.Raw, &queryFees)
	return &queryFees, err

}

// QueryGetinfo
// @param id
func (client *Client) QueryGetinfo(id int64) (*QueryInfo, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
			Id   int64  `json:"id"`
		}{
			Type: "query.getInfo",
			Id:   id,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var queryInfo QueryInfo
	err = json.Unmarshal(result.Raw, &queryInfo)
	return &queryInfo, err

}

// SmcLoad
// @param accountAddress
func (client *Client) SmcLoad(accountAddress *AccountAddress) (*SmcInfo, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type           string          `json:"@type"`
			AccountAddress *AccountAddress `json:"account_address"`
		}{
			Type:           "smc.load",
			AccountAddress: accountAddress,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var smcInfo SmcInfo
	err = json.Unmarshal(result.Raw, &smcInfo)
	return &smcInfo, err

}

// SmcGetcode
// @param id
func (client *Client) SmcGetcode(id int64) (*TvmCell, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
			Id   int64  `json:"id"`
		}{
			Type: "smc.getCode",
			Id:   id,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var tvmCell TvmCell
	err = json.Unmarshal(result.Raw, &tvmCell)
	return &tvmCell, err

}

// SmcGetdata
// @param id
func (client *Client) SmcGetdata(id int64) (*TvmCell, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
			Id   int64  `json:"id"`
		}{
			Type: "smc.getData",
			Id:   id,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var tvmCell TvmCell
	err = json.Unmarshal(result.Raw, &tvmCell)
	return &tvmCell, err

}

// SmcGetstate
// @param id
func (client *Client) SmcGetstate(id int64) (*TvmCell, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
			Id   int64  `json:"id"`
		}{
			Type: "smc.getState",
			Id:   id,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var tvmCell TvmCell
	err = json.Unmarshal(result.Raw, &tvmCell)
	return &tvmCell, err

}

// SmcRungetmethod
// @param method
// @param stack
// @param id
func (client *Client) SmcRungetmethod(method *SmcMethodId, stack []TvmStackEntry, id int64) (*SmcRunResult, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type   string          `json:"@type"`
			Method *SmcMethodId    `json:"method"`
			Stack  []TvmStackEntry `json:"stack"`
			Id     int64           `json:"id"`
		}{
			Type:   "smc.runGetMethod",
			Method: method,
			Stack:  stack,
			Id:     id,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var smcRunResult SmcRunResult
	err = json.Unmarshal(result.Raw, &smcRunResult)
	return &smcRunResult, err

}

// Onliteserverqueryresult
// @param id
// @param bytes
func (client *Client) Onliteserverqueryresult(id JSONInt64, bytes []byte) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type  string    `json:"@type"`
			Id    JSONInt64 `json:"id"`
			Bytes []byte    `json:"bytes"`
		}{
			Type:  "onLiteServerQueryResult",
			Id:    id,
			Bytes: bytes,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// Onliteserverqueryerror
// @param id
// @param error
func (client *Client) Onliteserverqueryerror(id JSONInt64, error *Error) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type  string    `json:"@type"`
			Id    JSONInt64 `json:"id"`
			Error *Error    `json:"error"`
		}{
			Type:  "onLiteServerQueryError",
			Id:    id,
			Error: error,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// Runtests
// @param dir
func (client *Client) Runtests(dir string) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
			Dir  string `json:"dir"`
		}{
			Type: "runTests",
			Dir:  dir,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// LiteserverGetinfo
func (client *Client) LiteserverGetinfo() (*LiteServerInfo, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
		}{
			Type: "liteServer.getInfo",
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var liteServerInfo LiteServerInfo
	err = json.Unmarshal(result.Raw, &liteServerInfo)
	return &liteServerInfo, err

}

// Setlogstream Sets new log stream for internal logging of tonlib. This is an offline method. Can be called before authorization. Can be called synchronously
// @param logStream New log stream
func (client *Client) Setlogstream(logStream LogStream) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type      string    `json:"@type"`
			LogStream LogStream `json:"log_stream"`
		}{
			Type:      "setLogStream",
			LogStream: logStream,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// Getlogstream Returns information about currently used log stream for internal logging of tonlib. This is an offline method. Can be called before authorization. Can be called synchronously
func (client *Client) Getlogstream() (LogStream, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
		}{
			Type: "getLogStream",
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch LogStreamEnum(result.Data["@type"].(string)) {

	case LogStreamDefaultType:
		var logStream LogStreamDefault
		err = json.Unmarshal(result.Raw, &logStream)
		return &logStream, err

	case LogStreamFileType:
		var logStream LogStreamFile
		err = json.Unmarshal(result.Raw, &logStream)
		return &logStream, err

	case LogStreamEmptyType:
		var logStream LogStreamEmpty
		err = json.Unmarshal(result.Raw, &logStream)
		return &logStream, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// Setlogverbositylevel Sets the verbosity level of the internal logging of tonlib. This is an offline method. Can be called before authorization. Can be called synchronously
// @param newVerbosityLevel New value of the verbosity level for logging. Value 0 corresponds to fatal errors, value 1 corresponds to errors, value 2 corresponds to warnings and debug warnings, value 3 corresponds to informational, value 4 corresponds to debug, value 5 corresponds to verbose debug, value greater than 5 and up to 1023 can be used to enable even more logging
func (client *Client) Setlogverbositylevel(newVerbosityLevel int32) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type              string `json:"@type"`
			NewVerbosityLevel int32  `json:"new_verbosity_level"`
		}{
			Type:              "setLogVerbosityLevel",
			NewVerbosityLevel: newVerbosityLevel,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// Getlogverbositylevel Returns current verbosity level of the internal logging of tonlib. This is an offline method. Can be called before authorization. Can be called synchronously
func (client *Client) Getlogverbositylevel() (*LogVerbosityLevel, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
		}{
			Type: "getLogVerbosityLevel",
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var logVerbosityLevel LogVerbosityLevel
	err = json.Unmarshal(result.Raw, &logVerbosityLevel)
	return &logVerbosityLevel, err

}

// Getlogtags Returns list of available tonlib internal log tags, for example, ["actor", "binlog", "connections", "notifications", "proxy"]. This is an offline method. Can be called before authorization. Can be called synchronously
func (client *Client) Getlogtags() (*LogTags, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
		}{
			Type: "getLogTags",
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var logTags LogTags
	err = json.Unmarshal(result.Raw, &logTags)
	return &logTags, err

}

// Setlogtagverbositylevel Sets the verbosity level for a specified tonlib internal log tag. This is an offline method. Can be called before authorization. Can be called synchronously
// @param tag Logging tag to change verbosity level
// @param newVerbosityLevel New verbosity level; 1-1024
func (client *Client) Setlogtagverbositylevel(tag string, newVerbosityLevel int32) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type              string `json:"@type"`
			Tag               string `json:"tag"`
			NewVerbosityLevel int32  `json:"new_verbosity_level"`
		}{
			Type:              "setLogTagVerbosityLevel",
			Tag:               tag,
			NewVerbosityLevel: newVerbosityLevel,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// Getlogtagverbositylevel Returns current verbosity level for a specified tonlib internal log tag. This is an offline method. Can be called before authorization. Can be called synchronously
// @param tag Logging tag to change verbosity level
func (client *Client) Getlogtagverbositylevel(tag string) (*LogVerbosityLevel, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
			Tag  string `json:"tag"`
		}{
			Type: "getLogTagVerbosityLevel",
			Tag:  tag,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var logVerbosityLevel LogVerbosityLevel
	err = json.Unmarshal(result.Raw, &logVerbosityLevel)
	return &logVerbosityLevel, err

}

// Addlogmessage Adds a message to tonlib internal log. This is an offline method. Can be called before authorization. Can be called synchronously
// @param verbosityLevel Minimum verbosity level needed for the message to be logged, 0-1023
// @param text Text of a message to log
func (client *Client) Addlogmessage(verbosityLevel int32, text string) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type           string `json:"@type"`
			VerbosityLevel int32  `json:"verbosity_level"`
			Text           string `json:"text"`
		}{
			Type:           "addLogMessage",
			VerbosityLevel: verbosityLevel,
			Text:           text,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}
