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

// OptionsSetConfig
// @param config
func (client *Client) OptionsSetConfig(config *Config) (*OptionsConfigInfo, error) {
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

// OptionsValidateConfig
// @param config
func (client *Client) OptionsValidateConfig(config *Config) (*OptionsConfigInfo, error) {
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

// CreateNewKey
// @param localPassword
// @param mnemonicPassword
// @param randomExtraSeed
func (client *Client) CreateNewKey(localPassword *SecureBytes, mnemonicPassword *SecureBytes, randomExtraSeed *SecureBytes) (*Key, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type             string       `json:"@type"`
			LocalPassword    *SecureBytes `json:"local_password"`
			MnemonicPassword *SecureBytes `json:"mnemonic_password"`
			RandomExtraSeed  *SecureBytes `json:"random_extra_seed"`
		}{
			Type:             "createNewKey",
			LocalPassword:    localPassword,
			MnemonicPassword: mnemonicPassword,
			RandomExtraSeed:  randomExtraSeed,
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

// DeleteKey
// @param key
func (client *Client) DeleteKey(key *Key) (*Ok, error) {
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

// DeleteAllKeys
func (client *Client) DeleteAllKeys() (*Ok, error) {
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

// ExportKey
// @param inputKey
func (client *Client) ExportKey(inputKey *InputKey) (*ExportedKey, error) {
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

// ExportPemKey
// @param inputKey
// @param keyPassword
func (client *Client) ExportPemKey(inputKey *InputKey, keyPassword *SecureBytes) (*ExportedPemKey, error) {
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

// ExportEncryptedKey
// @param inputKey
// @param keyPassword
func (client *Client) ExportEncryptedKey(inputKey *InputKey, keyPassword *SecureBytes) (*ExportedEncryptedKey, error) {
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

// ImportKey
// @param exportedKey
// @param localPassword
// @param mnemonicPassword
func (client *Client) ImportKey(exportedKey *ExportedKey, localPassword *SecureBytes, mnemonicPassword *SecureBytes) (*Key, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type             string       `json:"@type"`
			ExportedKey      *ExportedKey `json:"exported_key"`
			LocalPassword    *SecureBytes `json:"local_password"`
			MnemonicPassword *SecureBytes `json:"mnemonic_password"`
		}{
			Type:             "importKey",
			ExportedKey:      exportedKey,
			LocalPassword:    localPassword,
			MnemonicPassword: mnemonicPassword,
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

// ImportPemKey
// @param exportedKey
// @param keyPassword
// @param localPassword
func (client *Client) ImportPemKey(exportedKey *ExportedPemKey, keyPassword *SecureBytes, localPassword *SecureBytes) (*Key, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type          string          `json:"@type"`
			ExportedKey   *ExportedPemKey `json:"exported_key"`
			KeyPassword   *SecureBytes    `json:"key_password"`
			LocalPassword *SecureBytes    `json:"local_password"`
		}{
			Type:          "importPemKey",
			ExportedKey:   exportedKey,
			KeyPassword:   keyPassword,
			LocalPassword: localPassword,
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

// ImportEncryptedKey
// @param exportedEncryptedKey
// @param keyPassword
// @param localPassword
func (client *Client) ImportEncryptedKey(exportedEncryptedKey *ExportedEncryptedKey, keyPassword *SecureBytes, localPassword *SecureBytes) (*Key, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type                 string                `json:"@type"`
			ExportedEncryptedKey *ExportedEncryptedKey `json:"exported_encrypted_key"`
			KeyPassword          *SecureBytes          `json:"key_password"`
			LocalPassword        *SecureBytes          `json:"local_password"`
		}{
			Type:                 "importEncryptedKey",
			ExportedEncryptedKey: exportedEncryptedKey,
			KeyPassword:          keyPassword,
			LocalPassword:        localPassword,
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

// ChangeLocalPassword
// @param inputKey
// @param newLocalPassword
func (client *Client) ChangeLocalPassword(inputKey *InputKey, newLocalPassword *SecureBytes) (*Key, error) {
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

// UnpackAccountAddress
// @param accountAddress
func (client *Client) UnpackAccountAddress(accountAddress string) (*UnpackedAccountAddress, error) {
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

// PackAccountAddress
// @param accountAddress
func (client *Client) PackAccountAddress(accountAddress *UnpackedAccountAddress) (*AccountAddress, error) {
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

// GetBip39Hints
// @param prefix
func (client *Client) GetBip39Hints(prefix string) (*Bip39Hints, error) {
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

// RawGetAccountAddress
// @param inititalAccountState
func (client *Client) RawGetAccountAddress(inititalAccountState *RawInitialAccountState) (*AccountAddress, error) {
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

// RawGetAccountState
// @param accountAddress
func (client *Client) RawGetAccountState(accountAddress *AccountAddress) (*RawAccountState, error) {
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

// RawGetTransactions
// @param accountAddress
// @param fromTransactionId
func (client *Client) RawGetTransactions(accountAddress *AccountAddress, fromTransactionId *InternalTransactionId) (*RawTransactions, error) {
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

// RawSendMessage
// @param body
func (client *Client) RawSendMessage(body []byte) (*Ok, error) {
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

// RawCreateAndSendMessage
// @param data
// @param destination
// @param initialAccountState
func (client *Client) RawCreateAndSendMessage(data []byte, destination *AccountAddress, initialAccountState []byte) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type                string          `json:"@type"`
			Data                []byte          `json:"data"`
			Destination         *AccountAddress `json:"destination"`
			InitialAccountState []byte          `json:"initial_account_state"`
		}{
			Type:                "raw.createAndSendMessage",
			Data:                data,
			Destination:         destination,
			InitialAccountState: initialAccountState,
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

// RawCreateQuery
// @param body
// @param destination
// @param initCode
// @param initData
func (client *Client) RawCreateQuery(body []byte, destination *AccountAddress, initCode []byte, initData []byte) (*QueryInfo, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type        string          `json:"@type"`
			Body        []byte          `json:"body"`
			Destination *AccountAddress `json:"destination"`
			InitCode    []byte          `json:"init_code"`
			InitData    []byte          `json:"init_data"`
		}{
			Type:        "raw.createQuery",
			Body:        body,
			Destination: destination,
			InitCode:    initCode,
			InitData:    initData,
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

// TestWalletInit
// @param privateKey
func (client *Client) TestWalletInit(privateKey *InputKey) (*Ok, error) {
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

// TestWalletGetAccountAddress
// @param inititalAccountState
func (client *Client) TestWalletGetAccountAddress(inititalAccountState *TestWalletInitialAccountState) (*AccountAddress, error) {
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

// TestWalletGetAccountState
// @param accountAddress
func (client *Client) TestWalletGetAccountState(accountAddress *AccountAddress) (*TestWalletAccountState, error) {
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

// TestWalletSendGrams
// @param amount
// @param destination
// @param message
// @param privateKey
// @param seqno
func (client *Client) TestWalletSendGrams(amount JSONInt64, destination *AccountAddress, message []byte, privateKey *InputKey, seqno int32) (*SendGramsResult, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type        string          `json:"@type"`
			Amount      JSONInt64       `json:"amount"`
			Destination *AccountAddress `json:"destination"`
			Message     []byte          `json:"message"`
			PrivateKey  *InputKey       `json:"private_key"`
			Seqno       int32           `json:"seqno"`
		}{
			Type:        "testWallet.sendGrams",
			Amount:      amount,
			Destination: destination,
			Message:     message,
			PrivateKey:  privateKey,
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

// WalletGetAccountAddress
// @param inititalAccountState
func (client *Client) WalletGetAccountAddress(inititalAccountState *WalletInitialAccountState) (*AccountAddress, error) {
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

// WalletGetAccountState
// @param accountAddress
func (client *Client) WalletGetAccountState(accountAddress *AccountAddress) (*WalletAccountState, error) {
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

// WalletSendGrams
// @param amount
// @param destination
// @param message
// @param privateKey
// @param seqno
// @param validUntil
func (client *Client) WalletSendGrams(amount JSONInt64, destination *AccountAddress, message []byte, privateKey *InputKey, seqno int32, validUntil int64) (*SendGramsResult, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type        string          `json:"@type"`
			Amount      JSONInt64       `json:"amount"`
			Destination *AccountAddress `json:"destination"`
			Message     []byte          `json:"message"`
			PrivateKey  *InputKey       `json:"private_key"`
			Seqno       int32           `json:"seqno"`
			ValidUntil  int64           `json:"valid_until"`
		}{
			Type:        "wallet.sendGrams",
			Amount:      amount,
			Destination: destination,
			Message:     message,
			PrivateKey:  privateKey,
			Seqno:       seqno,
			ValidUntil:  validUntil,
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

// WalletV3GetAccountAddress
// @param inititalAccountState
func (client *Client) WalletV3GetAccountAddress(inititalAccountState *WalletV3InitialAccountState) (*AccountAddress, error) {
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

// TestGiverGetAccountState
func (client *Client) TestGiverGetAccountState() (*TestGiverAccountState, error) {
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

// TestGiverGetAccountAddress
func (client *Client) TestGiverGetAccountAddress() (*AccountAddress, error) {
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

// TestGiverSendGrams
// @param amount
// @param destination
// @param message
// @param seqno
func (client *Client) TestGiverSendGrams(amount JSONInt64, destination *AccountAddress, message []byte, seqno int32) (*SendGramsResult, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type        string          `json:"@type"`
			Amount      JSONInt64       `json:"amount"`
			Destination *AccountAddress `json:"destination"`
			Message     []byte          `json:"message"`
			Seqno       int32           `json:"seqno"`
		}{
			Type:        "testGiver.sendGrams",
			Amount:      amount,
			Destination: destination,
			Message:     message,
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

// GenericGetAccountState
// @param accountAddress
func (client *Client) GenericGetAccountState(accountAddress *AccountAddress) (*GenericAccountState, error) {
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

// GenericSendGrams
// @param allowSendToUninited
// @param amount
// @param destination
// @param message
// @param privateKey
// @param source
// @param timeout
func (client *Client) GenericSendGrams(allowSendToUninited bool, amount JSONInt64, destination *AccountAddress, message []byte, privateKey *InputKey, source *AccountAddress, timeout int32) (*SendGramsResult, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type                string          `json:"@type"`
			AllowSendToUninited bool            `json:"allow_send_to_uninited"`
			Amount              JSONInt64       `json:"amount"`
			Destination         *AccountAddress `json:"destination"`
			Message             []byte          `json:"message"`
			PrivateKey          *InputKey       `json:"private_key"`
			Source              *AccountAddress `json:"source"`
			Timeout             int32           `json:"timeout"`
		}{
			Type:                "generic.sendGrams",
			AllowSendToUninited: allowSendToUninited,
			Amount:              amount,
			Destination:         destination,
			Message:             message,
			PrivateKey:          privateKey,
			Source:              source,
			Timeout:             timeout,
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

// GenericCreateSendGramsQuery
// @param allowSendToUninited
// @param amount
// @param destination
// @param message
// @param privateKey
// @param source
// @param timeout
func (client *Client) GenericCreateSendGramsQuery(allowSendToUninited bool, amount JSONInt64, destination *AccountAddress, message []byte, privateKey *InputKey, source *AccountAddress, timeout int32) (*QueryInfo, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type                string          `json:"@type"`
			AllowSendToUninited bool            `json:"allow_send_to_uninited"`
			Amount              JSONInt64       `json:"amount"`
			Destination         *AccountAddress `json:"destination"`
			Message             []byte          `json:"message"`
			PrivateKey          *InputKey       `json:"private_key"`
			Source              *AccountAddress `json:"source"`
			Timeout             int32           `json:"timeout"`
		}{
			Type:                "generic.createSendGramsQuery",
			AllowSendToUninited: allowSendToUninited,
			Amount:              amount,
			Destination:         destination,
			Message:             message,
			PrivateKey:          privateKey,
			Source:              source,
			Timeout:             timeout,
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

// QueryEstimateFees
// @param id
// @param ignoreChksig
func (client *Client) QueryEstimateFees(id int64, ignoreChksig bool) (*QueryFees, error) {
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

// QueryGetInfo
// @param id
func (client *Client) QueryGetInfo(id int64) (*QueryInfo, error) {
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

// SmcGetCode
// @param id
func (client *Client) SmcGetCode(id int64) (*TvmCell, error) {
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

// SmcGetData
// @param id
func (client *Client) SmcGetData(id int64) (*TvmCell, error) {
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

// SmcGetState
// @param id
func (client *Client) SmcGetState(id int64) (*TvmCell, error) {
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

// SmcRunGetMethod
// @param id
// @param method
// @param stack
func (client *Client) SmcRunGetMethod(id int64, method *SmcMethodId, stack []TvmStackEntry) (*SmcRunResult, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type   string          `json:"@type"`
			Id     int64           `json:"id"`
			Method *SmcMethodId    `json:"method"`
			Stack  []TvmStackEntry `json:"stack"`
		}{
			Type:   "smc.runGetMethod",
			Id:     id,
			Method: method,
			Stack:  stack,
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

// OnLiteServerQueryResult
// @param bytes
// @param id
func (client *Client) OnLiteServerQueryResult(bytes []byte, id JSONInt64) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type  string    `json:"@type"`
			Bytes []byte    `json:"bytes"`
			Id    JSONInt64 `json:"id"`
		}{
			Type:  "onLiteServerQueryResult",
			Bytes: bytes,
			Id:    id,
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

// OnLiteServerQueryError
// @param error
// @param id
func (client *Client) OnLiteServerQueryError(error *Error, id JSONInt64) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type  string    `json:"@type"`
			Error *Error    `json:"error"`
			Id    JSONInt64 `json:"id"`
		}{
			Type:  "onLiteServerQueryError",
			Error: error,
			Id:    id,
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

// RunTests
// @param dir
func (client *Client) RunTests(dir string) (*Ok, error) {
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

// LiteServerGetInfo
func (client *Client) LiteServerGetInfo() (*LiteServerInfo, error) {
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

// SetLogStream Sets new log stream for internal logging of tonlib. This is an offline method. Can be called before authorization. Can be called synchronously
// @param logStream New log stream
func (client *Client) SetLogStream(logStream LogStream) (*Ok, error) {
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

// GetLogStream Returns information about currently used log stream for internal logging of tonlib. This is an offline method. Can be called before authorization. Can be called synchronously
func (client *Client) GetLogStream() (LogStream, error) {
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

// SetLogVerbosityLevel Sets the verbosity level of the internal logging of tonlib. This is an offline method. Can be called before authorization. Can be called synchronously
// @param newVerbosityLevel New value of the verbosity level for logging. Value 0 corresponds to fatal errors, value 1 corresponds to errors, value 2 corresponds to warnings and debug warnings, value 3 corresponds to informational, value 4 corresponds to debug, value 5 corresponds to verbose debug, value greater than 5 and up to 1023 can be used to enable even more logging
func (client *Client) SetLogVerbosityLevel(newVerbosityLevel int32) (*Ok, error) {
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

// GetLogVerbosityLevel Returns current verbosity level of the internal logging of tonlib. This is an offline method. Can be called before authorization. Can be called synchronously
func (client *Client) GetLogVerbosityLevel() (*LogVerbosityLevel, error) {
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

// GetLogTags Returns list of available tonlib internal log tags, for example, ["actor", "binlog", "connections", "notifications", "proxy"]. This is an offline method. Can be called before authorization. Can be called synchronously
func (client *Client) GetLogTags() (*LogTags, error) {
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

// SetLogTagVerbosityLevel Sets the verbosity level for a specified tonlib internal log tag. This is an offline method. Can be called before authorization. Can be called synchronously
// @param newVerbosityLevel New verbosity level; 1-1024
// @param tag Logging tag to change verbosity level
func (client *Client) SetLogTagVerbosityLevel(newVerbosityLevel int32, tag string) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type              string `json:"@type"`
			NewVerbosityLevel int32  `json:"new_verbosity_level"`
			Tag               string `json:"tag"`
		}{
			Type:              "setLogTagVerbosityLevel",
			NewVerbosityLevel: newVerbosityLevel,
			Tag:               tag,
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

// GetLogTagVerbosityLevel Returns current verbosity level for a specified tonlib internal log tag. This is an offline method. Can be called before authorization. Can be called synchronously
// @param tag Logging tag to change verbosity level
func (client *Client) GetLogTagVerbosityLevel(tag string) (*LogVerbosityLevel, error) {
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

// AddLogMessage Adds a message to tonlib internal log. This is an offline method. Can be called before authorization. Can be called synchronously
// @param text Text of a message to log
// @param verbosityLevel Minimum verbosity level needed for the message to be logged, 0-1023
func (client *Client) AddLogMessage(text string, verbosityLevel int32) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type           string `json:"@type"`
			Text           string `json:"text"`
			VerbosityLevel int32  `json:"verbosity_level"`
		}{
			Type:           "addLogMessage",
			Text:           text,
			VerbosityLevel: verbosityLevel,
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
