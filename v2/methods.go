package v2

import (
	"encoding/json"
	"fmt"
)

// Init
// @param options
func (client *Client) Init(options Options) (*OptionsInfo, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type    string  `json:"@type"`
			Options Options `json:"options"`
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
func (client *Client) OptionsSetConfig(config Config) (*OptionsConfigInfo, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type   string `json:"@type"`
			Config Config `json:"config"`
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
func (client *Client) OptionsValidateConfig(config Config) (*OptionsConfigInfo, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type   string `json:"@type"`
			Config Config `json:"config"`
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
func (client *Client) CreateNewKey(localPassword SecureBytes, mnemonicPassword SecureBytes, randomExtraSeed SecureBytes) (*Key, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type             string      `json:"@type"`
			LocalPassword    SecureBytes `json:"local_password"`
			MnemonicPassword SecureBytes `json:"mnemonic_password"`
			RandomExtraSeed  SecureBytes `json:"random_extra_seed"`
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
func (client *Client) DeleteKey(key Key) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
			Key  Key    `json:"key"`
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
func (client *Client) ExportKey(inputKey InputKey) (*ExportedKey, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type     string   `json:"@type"`
			InputKey InputKey `json:"input_key"`
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
func (client *Client) ExportPemKey(inputKey InputKey, keyPassword SecureBytes) (*ExportedPemKey, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type        string      `json:"@type"`
			InputKey    InputKey    `json:"input_key"`
			KeyPassword SecureBytes `json:"key_password"`
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
func (client *Client) ExportEncryptedKey(inputKey InputKey, keyPassword SecureBytes) (*ExportedEncryptedKey, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type        string      `json:"@type"`
			InputKey    InputKey    `json:"input_key"`
			KeyPassword SecureBytes `json:"key_password"`
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

// ExportUnencryptedKey
// @param inputKey
func (client *Client) ExportUnencryptedKey(inputKey InputKey) (*ExportedUnencryptedKey, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type     string   `json:"@type"`
			InputKey InputKey `json:"input_key"`
		}{
			Type:     "exportUnencryptedKey",
			InputKey: inputKey,
		},
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
	result, err := client.executeAsynchronously(
		struct {
			Type             string      `json:"@type"`
			ExportedKey      ExportedKey `json:"exported_key"`
			LocalPassword    SecureBytes `json:"local_password"`
			MnemonicPassword SecureBytes `json:"mnemonic_password"`
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
func (client *Client) ImportPemKey(exportedKey ExportedPemKey, keyPassword SecureBytes, localPassword SecureBytes) (*Key, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type          string         `json:"@type"`
			ExportedKey   ExportedPemKey `json:"exported_key"`
			KeyPassword   SecureBytes    `json:"key_password"`
			LocalPassword SecureBytes    `json:"local_password"`
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
func (client *Client) ImportEncryptedKey(exportedEncryptedKey ExportedEncryptedKey, keyPassword SecureBytes, localPassword SecureBytes) (*Key, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type                 string               `json:"@type"`
			ExportedEncryptedKey ExportedEncryptedKey `json:"exported_encrypted_key"`
			KeyPassword          SecureBytes          `json:"key_password"`
			LocalPassword        SecureBytes          `json:"local_password"`
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

// ImportUnencryptedKey
// @param exportedUnencryptedKey
// @param localPassword
func (client *Client) ImportUnencryptedKey(exportedUnencryptedKey ExportedUnencryptedKey, localPassword SecureBytes) (*Key, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type                   string                 `json:"@type"`
			ExportedUnencryptedKey ExportedUnencryptedKey `json:" exported_unencrypted_key"`
			LocalPassword          SecureBytes            `json:"local_password"`
		}{
			Type:                   "importUnencryptedKey",
			ExportedUnencryptedKey: exportedUnencryptedKey,
			LocalPassword:          localPassword,
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

// ChangeLocalPassword
// @param inputKey
// @param newLocalPassword
func (client *Client) ChangeLocalPassword(inputKey InputKey, newLocalPassword SecureBytes) (*Key, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type             string      `json:"@type"`
			InputKey         InputKey    `json:"input_key"`
			NewLocalPassword SecureBytes `json:"new_local_password"`
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
func (client *Client) Encrypt(decryptedData SecureBytes, secret SecureBytes) (*Data, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type          string      `json:"@type"`
			DecryptedData SecureBytes `json:"decrypted_data"`
			Secret        SecureBytes `json:"secret"`
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
func (client *Client) Decrypt(encryptedData SecureBytes, secret SecureBytes) (*Data, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type          string      `json:"@type"`
			EncryptedData SecureBytes `json:"encrypted_data"`
			Secret        SecureBytes `json:"secret"`
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
func (client *Client) Kdf(iterations int32, password SecureBytes, salt SecureBytes) (*Data, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type       string      `json:"@type"`
			Iterations int32       `json:"iterations"`
			Password   SecureBytes `json:"password"`
			Salt       SecureBytes `json:"salt"`
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
func (client *Client) PackAccountAddress(accountAddress UnpackedAccountAddress) (*AccountAddress, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type           string                 `json:"@type"`
			AccountAddress UnpackedAccountAddress `json:"account_address"`
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

// RawGetAccountState
// @param accountAddress
func (client *Client) RawGetAccountState(accountAddress AccountAddress) (*RawFullAccountState, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type           string         `json:"@type"`
			AccountAddress AccountAddress `json:"account_address"`
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

	var rawFullAccountState RawFullAccountState
	err = json.Unmarshal(result.Raw, &rawFullAccountState)
	return &rawFullAccountState, err

}

// RawGetTransactions
// @param accountAddress
// @param fromTransactionId
// @param privateKey
func (client *Client) RawGetTransactions(accountAddress AccountAddress, fromTransactionId InternalTransactionId, privateKey InputKey) (*RawTransactions, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type              string                `json:"@type"`
			AccountAddress    AccountAddress        `json:"account_address"`
			FromTransactionId InternalTransactionId `json:"from_transaction_id"`
			PrivateKey        InputKey              `json:"private_key"`
		}{
			Type:              "raw.getTransactions",
			AccountAddress:    accountAddress,
			FromTransactionId: fromTransactionId,
			PrivateKey:        privateKey,
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
func (client *Client) RawCreateAndSendMessage(data []byte, destination AccountAddress, initialAccountState []byte) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type                string         `json:"@type"`
			Data                []byte         `json:"data"`
			Destination         AccountAddress `json:"destination"`
			InitialAccountState []byte         `json:"initial_account_state"`
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
func (client *Client) RawCreateQuery(body []byte, destination AccountAddress, initCode []byte, initData []byte) (*QueryInfo, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type        string         `json:"@type"`
			Body        []byte         `json:"body"`
			Destination AccountAddress `json:"destination"`
			InitCode    []byte         `json:"init_code"`
			InitData    []byte         `json:"init_data"`
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

// GetAccountAddress
// @param initialAccountState
// @param revision
// @param workchainId
func (client *Client) GetAccountAddress(initialAccountState InitialAccountState, revision int32, workchainId int32) (*AccountAddress, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type                string              `json:"@type"`
			InitialAccountState InitialAccountState `json:"initial_account_state"`
			Revision            int32               `json:"revision"`
			WorkchainId         int32               `json:"workchain_id"`
		}{
			Type:                "getAccountAddress",
			InitialAccountState: initialAccountState,
			Revision:            revision,
			WorkchainId:         workchainId,
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

// GuessAccountRevision
// @param initialAccountState
// @param workchainId
func (client *Client) GuessAccountRevision(initialAccountState InitialAccountState, workchainId int32) (*AccountRevisionList, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type                string              `json:"@type"`
			InitialAccountState InitialAccountState `json:"initial_account_state"`
			WorkchainId         int32               `json:"workchain_id"`
		}{
			Type:                "guessAccountRevision",
			InitialAccountState: initialAccountState,
			WorkchainId:         workchainId,
		},
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

// GuessAccount
// @param publicKey
// @param rwalletInitPublicKey
func (client *Client) GuessAccount(publicKey string, rwalletInitPublicKey string) (*AccountRevisionList, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type                 string `json:"@type"`
			PublicKey            string `json:"public_key"`
			RwalletInitPublicKey string `json:"rwallet_init_public_key"`
		}{
			Type:                 "guessAccount",
			PublicKey:            publicKey,
			RwalletInitPublicKey: rwalletInitPublicKey,
		},
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
	result, err := client.executeAsynchronously(
		struct {
			Type           string         `json:"@type"`
			AccountAddress AccountAddress `json:"account_address"`
		}{
			Type:           "getAccountState",
			AccountAddress: accountAddress,
		},
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
// @param initialAccountState
// @param privateKey
// @param timeout
func (client *Client) CreateQuery(action Action, address AccountAddress, initialAccountState InitialAccountState, privateKey InputKey, timeout int32) (*QueryInfo, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type                string              `json:"@type"`
			Action              Action              `json:"action"`
			Address             AccountAddress      `json:"address"`
			InitialAccountState InitialAccountState `json:"initial_account_state"`
			PrivateKey          InputKey            `json:"private_key"`
			Timeout             int32               `json:"timeout"`
		}{
			Type:                "createQuery",
			Action:              action,
			Address:             address,
			InitialAccountState: initialAccountState,
			PrivateKey:          privateKey,
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

// MsgDecrypt
// @param data
// @param inputKey
func (client *Client) MsgDecrypt(data MsgDataEncryptedArray, inputKey InputKey) (*MsgDataDecryptedArray, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type     string                `json:"@type"`
			Data     MsgDataEncryptedArray `json:"data"`
			InputKey InputKey              `json:"input_key"`
		}{
			Type:     "msg.decrypt",
			Data:     data,
			InputKey: inputKey,
		},
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
	result, err := client.executeAsynchronously(
		struct {
			Type  string           `json:"@type"`
			Data  MsgDataEncrypted `json:"data"`
			Proof []byte           `json:"proof"`
		}{
			Type:  "msg.decryptWithProof",
			Data:  data,
			Proof: proof,
		},
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
func (client *Client) SmcLoad(accountAddress AccountAddress) (*SmcInfo, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type           string         `json:"@type"`
			AccountAddress AccountAddress `json:"account_address"`
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
func (client *Client) SmcRunGetMethod(id int64, method SmcMethodId, stack []TvmStackEntry) (*SmcRunResult, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type   string          `json:"@type"`
			Id     int64           `json:"id"`
			Method SmcMethodId     `json:"method"`
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

// DnsResolve
// @param accountAddress
// @param category
// @param name
// @param ttl
func (client *Client) DnsResolve(accountAddress AccountAddress, category int32, name string, ttl int32) (*DnsResolved, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type           string         `json:"@type"`
			AccountAddress AccountAddress `json:"account_address"`
			Category       int32          `json:"category"`
			Name           string         `json:"name"`
			Ttl            int32          `json:"ttl"`
		}{
			Type:           "dns.resolve",
			AccountAddress: accountAddress,
			Category:       category,
			Name:           name,
			Ttl:            ttl,
		},
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

// PchanSignPromise
// @param inputKey
// @param promise
func (client *Client) PchanSignPromise(inputKey InputKey, promise PchanPromise) (*PchanPromise, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type     string       `json:"@type"`
			InputKey InputKey     `json:"input_key"`
			Promise  PchanPromise `json:"promise"`
		}{
			Type:     "pchan.signPromise",
			InputKey: inputKey,
			Promise:  promise,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var pchanPromise PchanPromise
	err = json.Unmarshal(result.Raw, &pchanPromise)
	return &pchanPromise, err

}

// PchanValidatePromise
// @param promise
// @param publicKey
func (client *Client) PchanValidatePromise(promise PchanPromise, publicKey []byte) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type      string       `json:"@type"`
			Promise   PchanPromise `json:"promise"`
			PublicKey []byte       `json:"public_key"`
		}{
			Type:      "pchan.validatePromise",
			Promise:   promise,
			PublicKey: publicKey,
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

// PchanPackPromise
// @param promise
func (client *Client) PchanPackPromise(promise PchanPromise) (*Data, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type    string       `json:"@type"`
			Promise PchanPromise `json:"promise"`
		}{
			Type:    "pchan.packPromise",
			Promise: promise,
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

// PchanUnpackPromise
// @param data
func (client *Client) PchanUnpackPromise(data SecureBytes) (*PchanPromise, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string      `json:"@type"`
			Data SecureBytes `json:"data"`
		}{
			Type: "pchan.unpackPromise",
			Data: data,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var pchanPromise PchanPromise
	err = json.Unmarshal(result.Raw, &pchanPromise)
	return &pchanPromise, err

}

// BlocksGetMasterchainInfo
func (client *Client) BlocksGetMasterchainInfo() (*BlocksMasterchainInfo, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string `json:"@type"`
		}{
			Type: "blocks.getMasterchainInfo",
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var blocksMasterchainInfo BlocksMasterchainInfo
	err = json.Unmarshal(result.Raw, &blocksMasterchainInfo)
	return &blocksMasterchainInfo, err

}

// BlocksGetShards
// @param id
func (client *Client) BlocksGetShards(id TonBlockIdExt) (*BlocksShards, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string        `json:"@type"`
			Id   TonBlockIdExt `json:"id"`
		}{
			Type: "blocks.getShards",
			Id:   id,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var blocksShards BlocksShards
	err = json.Unmarshal(result.Raw, &blocksShards)
	return &blocksShards, err

}

// BlocksLookupBlock
// @param id
// @param lt
// @param mode
// @param utime
func (client *Client) BlocksLookupBlock(id TonBlockId, lt JSONInt64, mode int32, utime int32) (*TonBlockIdExt, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type  string     `json:"@type"`
			Id    TonBlockId `json:"id"`
			Lt    JSONInt64  `json:"lt"`
			Mode  int32      `json:"mode"`
			Utime int32      `json:"utime"`
		}{
			Type:  "blocks.lookupBlock",
			Id:    id,
			Lt:    lt,
			Mode:  mode,
			Utime: utime,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var tonBlockIdExt TonBlockIdExt
	err = json.Unmarshal(result.Raw, &tonBlockIdExt)
	return &tonBlockIdExt, err

}

// BlocksGetTransactions
// @param after
// @param count
// @param id
// @param mode
func (client *Client) BlocksGetTransactions(after BlocksAccountTransactionId, count int32, id TonBlockIdExt, mode int32) (*BlocksTransactions, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type  string                     `json:"@type"`
			After BlocksAccountTransactionId `json:"after"`
			Count int32                      `json:"count"`
			Id    TonBlockIdExt              `json:"id"`
			Mode  int32                      `json:"mode"`
		}{
			Type:  "blocks.getTransactions",
			After: after,
			Count: count,
			Id:    id,
			Mode:  mode,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var blocksTransactions BlocksTransactions
	err = json.Unmarshal(result.Raw, &blocksTransactions)
	return &blocksTransactions, err

}

// BlocksGetBlockHeader
// @param id
func (client *Client) BlocksGetBlockHeader(id TonBlockIdExt) (*BlocksHeader, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type string        `json:"@type"`
			Id   TonBlockIdExt `json:"id"`
		}{
			Type: "blocks.getBlockHeader",
			Id:   id,
		},
	)

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var blocksHeader BlocksHeader
	err = json.Unmarshal(result.Raw, &blocksHeader)
	return &blocksHeader, err

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
func (client *Client) OnLiteServerQueryError(error Error, id JSONInt64) (*Ok, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type  string    `json:"@type"`
			Error Error     `json:"error"`
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

// WithBlock
// @param function
// @param id
func (client *Client) WithBlock(function Function, id TonBlockIdExt) (*Object, error) {
	result, err := client.executeAsynchronously(
		struct {
			Type     string        `json:"@type"`
			Function Function      `json:"function"`
			Id       TonBlockIdExt `json:"id"`
		}{
			Type:     "withBlock",
			Function: function,
			Id:       id,
		},
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
