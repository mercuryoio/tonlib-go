package v2

import (
	"encoding/json"
	"fmt"
)

// Init
// @param options
func (client *Client) Init(options Options) (*OptionsInfo, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type    string  `json:"@type"`
			Extra   string  `json:"@extra"`
			Options Options `json:"options"`
		}{
			Type:    "init",
			Extra:   requestID,
			Options: options,
		},
		requestID,
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
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type  string `json:"@type"`
			Extra string `json:"@extra"`
		}{
			Type:  "close",
			Extra: requestID,
		},
		requestID,
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
func (client *Client) OptionsSetConfig(config Config) (*OptionsConfigInfo, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type   string `json:"@type"`
			Extra  string `json:"@extra"`
			Config Config `json:"config"`
		}{
			Type:   "options.setConfig",
			Extra:  requestID,
			Config: config,
		},
		requestID,
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
func (client *Client) OptionsValidateConfig(config Config) (*OptionsConfigInfo, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type   string `json:"@type"`
			Extra  string `json:"@extra"`
			Config Config `json:"config"`
		}{
			Type:   "options.validateConfig",
			Extra:  requestID,
			Config: config,
		},
		requestID,
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
func (client *Client) CreateNewKey(localPassword SecureBytes, mnemonicPassword SecureBytes, randomExtraSeed SecureBytes) (*Key, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type             string      `json:"@type"`
			Extra            string      `json:"@extra"`
			LocalPassword    SecureBytes `json:"local_password"`
			MnemonicPassword SecureBytes `json:"mnemonic_password"`
			RandomExtraSeed  SecureBytes `json:"random_extra_seed"`
		}{
			Type:             "createNewKey",
			Extra:            requestID,
			LocalPassword:    localPassword,
			MnemonicPassword: mnemonicPassword,
			RandomExtraSeed:  randomExtraSeed,
		},
		requestID,
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
func (client *Client) DeleteKey(key Key) (*Ok, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type  string `json:"@type"`
			Extra string `json:"@extra"`
			Key   Key    `json:"key"`
		}{
			Type:  "deleteKey",
			Extra: requestID,
			Key:   key,
		},
		requestID,
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
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type  string `json:"@type"`
			Extra string `json:"@extra"`
		}{
			Type:  "deleteAllKeys",
			Extra: requestID,
		},
		requestID,
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
func (client *Client) ExportKey(inputKey InputKey) (*ExportedKey, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type     string   `json:"@type"`
			Extra    string   `json:"@extra"`
			InputKey InputKey `json:"input_key"`
		}{
			Type:     "exportKey",
			Extra:    requestID,
			InputKey: inputKey,
		},
		requestID,
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
func (client *Client) ExportPemKey(inputKey InputKey, keyPassword SecureBytes) (*ExportedPemKey, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type        string      `json:"@type"`
			Extra       string      `json:"@extra"`
			InputKey    InputKey    `json:"input_key"`
			KeyPassword SecureBytes `json:"key_password"`
		}{
			Type:        "exportPemKey",
			Extra:       requestID,
			InputKey:    inputKey,
			KeyPassword: keyPassword,
		},
		requestID,
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
func (client *Client) ExportEncryptedKey(inputKey InputKey, keyPassword SecureBytes) (*ExportedEncryptedKey, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type        string      `json:"@type"`
			Extra       string      `json:"@extra"`
			InputKey    InputKey    `json:"input_key"`
			KeyPassword SecureBytes `json:"key_password"`
		}{
			Type:        "exportEncryptedKey",
			Extra:       requestID,
			InputKey:    inputKey,
			KeyPassword: keyPassword,
		},
		requestID,
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

// ExportUnencryptedKey
// @param inputKey
func (client *Client) ExportUnencryptedKey(inputKey InputKey) (*ExportedUnencryptedKey, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type     string   `json:"@type"`
			Extra    string   `json:"@extra"`
			InputKey InputKey `json:"input_key"`
		}{
			Type:     "exportUnencryptedKey",
			Extra:    requestID,
			InputKey: inputKey,
		},
		requestID,
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var exportedUnencryptedKey ExportedUnencryptedKey
	err = json.Unmarshal(result.Raw, &exportedUnencryptedKey)
	return &exportedUnencryptedKey, err

}

// ImportKey
// @param exportedKey
// @param localPassword
// @param mnemonicPassword
func (client *Client) ImportKey(exportedKey ExportedKey, localPassword SecureBytes, mnemonicPassword SecureBytes) (*Key, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type             string      `json:"@type"`
			Extra            string      `json:"@extra"`
			ExportedKey      ExportedKey `json:"exported_key"`
			LocalPassword    SecureBytes `json:"local_password"`
			MnemonicPassword SecureBytes `json:"mnemonic_password"`
		}{
			Type:             "importKey",
			Extra:            requestID,
			ExportedKey:      exportedKey,
			LocalPassword:    localPassword,
			MnemonicPassword: mnemonicPassword,
		},
		requestID,
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
func (client *Client) ImportPemKey(exportedKey ExportedPemKey, keyPassword SecureBytes, localPassword SecureBytes) (*Key, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type          string         `json:"@type"`
			Extra         string         `json:"@extra"`
			ExportedKey   ExportedPemKey `json:"exported_key"`
			KeyPassword   SecureBytes    `json:"key_password"`
			LocalPassword SecureBytes    `json:"local_password"`
		}{
			Type:          "importPemKey",
			Extra:         requestID,
			ExportedKey:   exportedKey,
			KeyPassword:   keyPassword,
			LocalPassword: localPassword,
		},
		requestID,
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
func (client *Client) ImportEncryptedKey(exportedEncryptedKey ExportedEncryptedKey, keyPassword SecureBytes, localPassword SecureBytes) (*Key, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type                 string               `json:"@type"`
			Extra                string               `json:"@extra"`
			ExportedEncryptedKey ExportedEncryptedKey `json:"exported_encrypted_key"`
			KeyPassword          SecureBytes          `json:"key_password"`
			LocalPassword        SecureBytes          `json:"local_password"`
		}{
			Type:                 "importEncryptedKey",
			Extra:                requestID,
			ExportedEncryptedKey: exportedEncryptedKey,
			KeyPassword:          keyPassword,
			LocalPassword:        localPassword,
		},
		requestID,
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

// ImportUnencryptedKey
// @param exportedUnencryptedKey
// @param localPassword
func (client *Client) ImportUnencryptedKey(exportedUnencryptedKey ExportedUnencryptedKey, localPassword SecureBytes) (*Key, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type                   string                 `json:"@type"`
			Extra                  string                 `json:"@extra"`
			ExportedUnencryptedKey ExportedUnencryptedKey `json:" exported_unencrypted_key"`
			LocalPassword          SecureBytes            `json:"local_password"`
		}{
			Type:                   "importUnencryptedKey",
			Extra:                  requestID,
			ExportedUnencryptedKey: exportedUnencryptedKey,
			LocalPassword:          localPassword,
		},
		requestID,
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

// ChangeLocalPassword
// @param inputKey
// @param newLocalPassword
func (client *Client) ChangeLocalPassword(inputKey InputKey, newLocalPassword SecureBytes) (*Key, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type             string      `json:"@type"`
			Extra            string      `json:"@extra"`
			InputKey         InputKey    `json:"input_key"`
			NewLocalPassword SecureBytes `json:"new_local_password"`
		}{
			Type:             "changeLocalPassword",
			Extra:            requestID,
			InputKey:         inputKey,
			NewLocalPassword: newLocalPassword,
		},
		requestID,
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
func (client *Client) Encrypt(decryptedData SecureBytes, secret SecureBytes) (*Data, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type          string      `json:"@type"`
			Extra         string      `json:"@extra"`
			DecryptedData SecureBytes `json:"decrypted_data"`
			Secret        SecureBytes `json:"secret"`
		}{
			Type:          "encrypt",
			Extra:         requestID,
			DecryptedData: decryptedData,
			Secret:        secret,
		},
		requestID,
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
func (client *Client) Decrypt(encryptedData SecureBytes, secret SecureBytes) (*Data, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type          string      `json:"@type"`
			Extra         string      `json:"@extra"`
			EncryptedData SecureBytes `json:"encrypted_data"`
			Secret        SecureBytes `json:"secret"`
		}{
			Type:          "decrypt",
			Extra:         requestID,
			EncryptedData: encryptedData,
			Secret:        secret,
		},
		requestID,
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
func (client *Client) Kdf(iterations int32, password SecureBytes, salt SecureBytes) (*Data, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type       string      `json:"@type"`
			Extra      string      `json:"@extra"`
			Iterations int32       `json:"iterations"`
			Password   SecureBytes `json:"password"`
			Salt       SecureBytes `json:"salt"`
		}{
			Type:       "kdf",
			Extra:      requestID,
			Iterations: iterations,
			Password:   password,
			Salt:       salt,
		},
		requestID,
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
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type           string `json:"@type"`
			Extra          string `json:"@extra"`
			AccountAddress string `json:"account_address"`
		}{
			Type:           "unpackAccountAddress",
			Extra:          requestID,
			AccountAddress: accountAddress,
		},
		requestID,
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
func (client *Client) PackAccountAddress(accountAddress UnpackedAccountAddress) (*AccountAddress, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type           string                 `json:"@type"`
			Extra          string                 `json:"@extra"`
			AccountAddress UnpackedAccountAddress `json:"account_address"`
		}{
			Type:           "packAccountAddress",
			Extra:          requestID,
			AccountAddress: accountAddress,
		},
		requestID,
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
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type   string `json:"@type"`
			Extra  string `json:"@extra"`
			Prefix string `json:"prefix"`
		}{
			Type:   "getBip39Hints",
			Extra:  requestID,
			Prefix: prefix,
		},
		requestID,
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

// RawGetAccountState
// @param accountAddress
func (client *Client) RawGetAccountState(accountAddress AccountAddress) (*RawFullAccountState, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type           string         `json:"@type"`
			Extra          string         `json:"@extra"`
			AccountAddress AccountAddress `json:"account_address"`
		}{
			Type:           "raw.getAccountState",
			Extra:          requestID,
			AccountAddress: accountAddress,
		},
		requestID,
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var rawFullAccountState RawFullAccountState
	err = json.Unmarshal(result.Raw, &rawFullAccountState)
	return &rawFullAccountState, err

}

// RawGetTransactions
// @param accountAddress
// @param fromTransactionId
func (client *Client) RawGetTransactions(accountAddress AccountAddress, fromTransactionId InternalTransactionId) (*RawTransactions, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type              string                `json:"@type"`
			Extra             string                `json:"@extra"`
			AccountAddress    AccountAddress        `json:"account_address"`
			FromTransactionId InternalTransactionId `json:"from_transaction_id"`
		}{
			Type:              "raw.getTransactions",
			Extra:             requestID,
			AccountAddress:    accountAddress,
			FromTransactionId: fromTransactionId,
		},
		requestID,
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
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type  string `json:"@type"`
			Extra string `json:"@extra"`
			Body  []byte `json:"body"`
		}{
			Type:  "raw.sendMessage",
			Extra: requestID,
			Body:  body,
		},
		requestID,
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
func (client *Client) RawCreateAndSendMessage(data []byte, destination AccountAddress, initialAccountState []byte) (*Ok, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type                string         `json:"@type"`
			Extra               string         `json:"@extra"`
			Data                []byte         `json:"data"`
			Destination         AccountAddress `json:"destination"`
			InitialAccountState []byte         `json:"initial_account_state"`
		}{
			Type:                "raw.createAndSendMessage",
			Extra:               requestID,
			Data:                data,
			Destination:         destination,
			InitialAccountState: initialAccountState,
		},
		requestID,
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
func (client *Client) RawCreateQuery(body []byte, destination AccountAddress, initCode []byte, initData []byte) (*QueryInfo, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type        string         `json:"@type"`
			Extra       string         `json:"@extra"`
			Body        []byte         `json:"body"`
			Destination AccountAddress `json:"destination"`
			InitCode    []byte         `json:"init_code"`
			InitData    []byte         `json:"init_data"`
		}{
			Type:        "raw.createQuery",
			Extra:       requestID,
			Body:        body,
			Destination: destination,
			InitCode:    initCode,
			InitData:    initData,
		},
		requestID,
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

// GetAccountAddress
// @param initialAccountState
// @param revision
func (client *Client) GetAccountAddress(initialAccountState InitialAccountState, revision int32) (*AccountAddress, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type                string              `json:"@type"`
			Extra               string              `json:"@extra"`
			InitialAccountState InitialAccountState `json:"initial_account_state"`
			Revision            int32               `json:"revision"`
		}{
			Type:                "getAccountAddress",
			Extra:               requestID,
			InitialAccountState: initialAccountState,
			Revision:            revision,
		},
		requestID,
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

// GuessAccountRevision
// @param initialAccountState
func (client *Client) GuessAccountRevision(initialAccountState InitialAccountState) (*AccountRevisionList, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type                string              `json:"@type"`
			Extra               string              `json:"@extra"`
			InitialAccountState InitialAccountState `json:"initial_account_state"`
		}{
			Type:                "guessAccountRevision",
			Extra:               requestID,
			InitialAccountState: initialAccountState,
		},
		requestID,
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var accountRevisionList AccountRevisionList
	err = json.Unmarshal(result.Raw, &accountRevisionList)
	return &accountRevisionList, err

}

// GetAccountState
// @param accountAddress
func (client *Client) GetAccountState(accountAddress AccountAddress) (*FullAccountState, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type           string         `json:"@type"`
			Extra          string         `json:"@extra"`
			AccountAddress AccountAddress `json:"account_address"`
		}{
			Type:           "getAccountState",
			Extra:          requestID,
			AccountAddress: accountAddress,
		},
		requestID,
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var fullAccountState FullAccountState
	err = json.Unmarshal(result.Raw, &fullAccountState)
	return &fullAccountState, err

}

// CreateQuery
// @param action
// @param address
// @param privateKey
// @param timeout
func (client *Client) CreateQuery(action Action, address AccountAddress, privateKey InputKey, timeout int32) (*QueryInfo, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronouslyExactlyOnce(
		struct {
			Type       string         `json:"@type"`
			Extra      string         `json:"@extra"`
			Action     Action         `json:"action"`
			Address    AccountAddress `json:"address"`
			PrivateKey InputKey       `json:"private_key"`
			Timeout    int32          `json:"timeout"`
		}{
			Type:       "createQuery",
			Extra:      requestID,
			Action:     action,
			Address:    address,
			PrivateKey: privateKey,
			Timeout:    timeout,
		},
		requestID,
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

// MsgDecrypt
// @param data
// @param inputKey
func (client *Client) MsgDecrypt(data MsgDataEncryptedArray, inputKey InputKey) (*MsgDataDecryptedArray, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type     string                `json:"@type"`
			Extra    string                `json:"@extra"`
			Data     MsgDataEncryptedArray `json:"data"`
			InputKey InputKey              `json:"input_key"`
		}{
			Type:     "msg.decrypt",
			Extra:    requestID,
			Data:     data,
			InputKey: inputKey,
		},
		requestID,
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var msgDataDecryptedArray MsgDataDecryptedArray
	err = json.Unmarshal(result.Raw, &msgDataDecryptedArray)
	return &msgDataDecryptedArray, err

}

// MsgDecryptWithProof
// @param data
// @param proof
func (client *Client) MsgDecryptWithProof(data MsgDataEncrypted, proof []byte) (*MsgData, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type  string           `json:"@type"`
			Extra string           `json:"@extra"`
			Data  MsgDataEncrypted `json:"data"`
			Proof []byte           `json:"proof"`
		}{
			Type:  "msg.decryptWithProof",
			Extra: requestID,
			Data:  data,
			Proof: proof,
		},
		requestID,
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var msgData MsgData
	err = json.Unmarshal(result.Raw, &msgData)
	return &msgData, err

}

// QuerySend
// @param id
func (client *Client) QuerySend(id int64) (*Ok, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type  string `json:"@type"`
			Extra string `json:"@extra"`
			Id    int64  `json:"id"`
		}{
			Type:  "query.send",
			Extra: requestID,
			Id:    id,
		},
		requestID,
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
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type  string `json:"@type"`
			Extra string `json:"@extra"`
			Id    int64  `json:"id"`
		}{
			Type:  "query.forget",
			Extra: requestID,
			Id:    id,
		},
		requestID,
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

// QueryGetInfo
// @param id
func (client *Client) QueryGetInfo(id int64) (*QueryInfo, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type  string `json:"@type"`
			Extra string `json:"@extra"`
			Id    int64  `json:"id"`
		}{
			Type:  "query.getInfo",
			Extra: requestID,
			Id:    id,
		},
		requestID,
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
func (client *Client) SmcLoad(accountAddress AccountAddress) (*SmcInfo, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type           string         `json:"@type"`
			Extra          string         `json:"@extra"`
			AccountAddress AccountAddress `json:"account_address"`
		}{
			Type:           "smc.load",
			Extra:          requestID,
			AccountAddress: accountAddress,
		},
		requestID,
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
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type  string `json:"@type"`
			Extra string `json:"@extra"`
			Id    int64  `json:"id"`
		}{
			Type:  "smc.getCode",
			Extra: requestID,
			Id:    id,
		},
		requestID,
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
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type  string `json:"@type"`
			Extra string `json:"@extra"`
			Id    int64  `json:"id"`
		}{
			Type:  "smc.getData",
			Extra: requestID,
			Id:    id,
		},
		requestID,
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
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type  string `json:"@type"`
			Extra string `json:"@extra"`
			Id    int64  `json:"id"`
		}{
			Type:  "smc.getState",
			Extra: requestID,
			Id:    id,
		},
		requestID,
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
func (client *Client) SmcRunGetMethod(id int64, method SmcMethodId, stack []TvmStackEntry) (*SmcRunResult, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type   string          `json:"@type"`
			Extra  string          `json:"@extra"`
			Id     int64           `json:"id"`
			Method SmcMethodId     `json:"method"`
			Stack  []TvmStackEntry `json:"stack"`
		}{
			Type:   "smc.runGetMethod",
			Extra:  requestID,
			Id:     id,
			Method: method,
			Stack:  stack,
		},
		requestID,
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

// DnsResolve
// @param accountAddress
// @param category
// @param name
// @param ttl
func (client *Client) DnsResolve(accountAddress AccountAddress, category int32, name string, ttl int32) (*DnsResolved, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type           string         `json:"@type"`
			Extra          string         `json:"@extra"`
			AccountAddress AccountAddress `json:"account_address"`
			Category       int32          `json:"category"`
			Name           string         `json:"name"`
			Ttl            int32          `json:"ttl"`
		}{
			Type:           "dns.resolve",
			Extra:          requestID,
			AccountAddress: accountAddress,
			Category:       category,
			Name:           name,
			Ttl:            ttl,
		},
		requestID,
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var dnsResolved DnsResolved
	err = json.Unmarshal(result.Raw, &dnsResolved)
	return &dnsResolved, err

}

// OnLiteServerQueryResult
// @param bytes
// @param id
func (client *Client) OnLiteServerQueryResult(bytes []byte, id JSONInt64) (*Ok, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type  string    `json:"@type"`
			Extra string    `json:"@extra"`
			Bytes []byte    `json:"bytes"`
			Id    JSONInt64 `json:"id"`
		}{
			Type:  "onLiteServerQueryResult",
			Extra: requestID,
			Bytes: bytes,
			Id:    id,
		},
		requestID,
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
func (client *Client) OnLiteServerQueryError(error Error, id JSONInt64) (*Ok, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type  string    `json:"@type"`
			Extra string    `json:"@extra"`
			Error Error     `json:"error"`
			Id    JSONInt64 `json:"id"`
		}{
			Type:  "onLiteServerQueryError",
			Extra: requestID,
			Error: error,
			Id:    id,
		},
		requestID,
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

// WithBlock
// @param function
// @param id
func (client *Client) WithBlock(function Function, id TonBlockIdExt) (*Object, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type     string        `json:"@type"`
			Extra    string        `json:"@extra"`
			Function Function      `json:"function"`
			Id       TonBlockIdExt `json:"id"`
		}{
			Type:     "withBlock",
			Extra:    requestID,
			Function: function,
			Id:       id,
		},
		requestID,
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var object Object
	err = json.Unmarshal(result.Raw, &object)
	return &object, err

}

// RunTests
// @param dir
func (client *Client) RunTests(dir string) (*Ok, error) {
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type  string `json:"@type"`
			Extra string `json:"@extra"`
			Dir   string `json:"dir"`
		}{
			Type:  "runTests",
			Extra: requestID,
			Dir:   dir,
		},
		requestID,
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
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type  string `json:"@type"`
			Extra string `json:"@extra"`
		}{
			Type:  "liteServer.getInfo",
			Extra: requestID,
		},
		requestID,
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
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type      string    `json:"@type"`
			Extra     string    `json:"@extra"`
			LogStream LogStream `json:"log_stream"`
		}{
			Type:      "setLogStream",
			Extra:     requestID,
			LogStream: logStream,
		},
		requestID,
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
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type  string `json:"@type"`
			Extra string `json:"@extra"`
		}{
			Type:  "getLogStream",
			Extra: requestID,
		},
		requestID,
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
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type              string `json:"@type"`
			Extra             string `json:"@extra"`
			NewVerbosityLevel int32  `json:"new_verbosity_level"`
		}{
			Type:              "setLogVerbosityLevel",
			Extra:             requestID,
			NewVerbosityLevel: newVerbosityLevel,
		},
		requestID,
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
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type  string `json:"@type"`
			Extra string `json:"@extra"`
		}{
			Type:  "getLogVerbosityLevel",
			Extra: requestID,
		},
		requestID,
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
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type  string `json:"@type"`
			Extra string `json:"@extra"`
		}{
			Type:  "getLogTags",
			Extra: requestID,
		},
		requestID,
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
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type              string `json:"@type"`
			Extra             string `json:"@extra"`
			NewVerbosityLevel int32  `json:"new_verbosity_level"`
			Tag               string `json:"tag"`
		}{
			Type:              "setLogTagVerbosityLevel",
			Extra:             requestID,
			NewVerbosityLevel: newVerbosityLevel,
			Tag:               tag,
		},
		requestID,
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
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type  string `json:"@type"`
			Extra string `json:"@extra"`
			Tag   string `json:"tag"`
		}{
			Type:  "getLogTagVerbosityLevel",
			Extra: requestID,
			Tag:   tag,
		},
		requestID,
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
	requestID := client.GenerateRequestID()
	result, err := client.executeAsynchronously(
		struct {
			Type           string `json:"@type"`
			Extra          string `json:"@extra"`
			Text           string `json:"text"`
			VerbosityLevel int32  `json:"verbosity_level"`
		}{
			Type:           "addLogMessage",
			Extra:          requestID,
			Text:           text,
			VerbosityLevel: verbosityLevel,
		},
		requestID,
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
