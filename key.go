package tonlib

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// CreatePrivateKey createNewKey: create privateKey
func (client *Client) CreatePrivateKey(localPass, mnemonicPass []byte) (key *TONPrivateKey, err error) {
	st := struct {
		Type             string `json:"@type"`
		LocalPassword    string `json:"local_password"`
		MnemonicPassword string `json:"mnemonic_password"`
		RandomExtraSeed  string `json:"random_extra_seed"`
	}{
		Type:             "createNewKey",
		LocalPassword:    base64.StdEncoding.EncodeToString(localPass),
		MnemonicPassword: base64.StdEncoding.EncodeToString(mnemonicPass),
	}
	resp, err := client.executeAsynchronously(st)
	if err != nil {
		return key, err
	}
	if st, ok := resp.Data["@type"]; ok && st == "error" {
		return key, fmt.Errorf("Error ton create private key. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	}

	key = new(TONPrivateKey)
	err = json.Unmarshal(resp.Raw, key)
	return key, err
}

// DeletePrivateKey deleteKey: delete private key
func (client *Client) DeletePrivateKey(key *TONPrivateKey, password []byte) (err error) {
	k := key.getInputKey(password)
	st := struct {
		Type string        `json:"@type"`
		Key  TONPrivateKey `json:"key"`
	}{
		Type: "deleteKey",
		Key:  k.Key,
	}
	resp, err := client.executeAsynchronously(st)
	if err != nil {
		return err
	}
	if st, ok := resp.Data["@type"]; ok && st == "error" {
		return fmt.Errorf("Error ton create private key. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	}

	return nil
}

// ExportPrivateKey exportKey: export private key
func (client *Client) ExportPrivateKey(key *TONPrivateKey, password []byte) (wordList []string, err error) {
	st := struct {
		Type     string   `json:"@type"`
		InputKey InputKey `json:"input_key"`
	}{
		Type:     "exportKey",
		InputKey: key.getInputKey(password),
	}
	resp, err := client.executeAsynchronously(st)
	if err != nil {
		return []string{}, err
	}
	if st, ok := resp.Data["@type"]; ok && st == "error" {
		return []string{}, fmt.Errorf("Error ton export private key. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	}

	mm := struct {
		WordList []string `json:"word_list"`
	}{}
	err = json.Unmarshal(resp.Raw, &mm)
	if err != nil {
		return []string{}, err
	}
	return mm.WordList, nil
}

// ExportPemKey exportPemKey: export pem
func (client *Client) ExportPemKey(key *TONPrivateKey, password, pemPassword []byte) (pem string, err error) {
	st := struct {
		Type        string   `json:"@type"`
		InputKey    InputKey `json:"input_key"`
		KeyPassword string   `json:"key_password"`
	}{
		Type:        "exportPemKey",
		InputKey:    key.getInputKey(password),
		KeyPassword: base64.StdEncoding.EncodeToString(pemPassword),
	}
	resp, err := client.executeAsynchronously(st)
	if err != nil {
		return "", err
	}
	if st, ok := resp.Data["@type"]; ok && st == "error" {
		return "", fmt.Errorf("Error ton export pem key. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	}

	p := struct {
		Pem string `json:"pem"`
	}{}
	err = json.Unmarshal(resp.Raw, &p)
	if err != nil {
		return "", err
	}
	return p.Pem, nil
}

// ExportEncryptedKey exportEncryptedKey: export encrypted key
func (client *Client) ExportEncryptedKey(key *TONPrivateKey, password, keyPassword []byte) (expKey *TONEncryptedKey, err error) {
	expKey = new(TONEncryptedKey)
	st := struct {
		Type        string   `json:"@type"`
		InputKey    InputKey `json:"input_key"`
		KeyPassword string   `json:"key_password"`
	}{
		Type:        "exportEncryptedKey",
		InputKey:    key.getInputKey(password),
		KeyPassword: base64.StdEncoding.EncodeToString(keyPassword),
	}
	resp, err := client.executeAsynchronously(st)
	if err != nil {
		return expKey, err
	}
	if st, ok := resp.Data["@type"]; ok && st == "error" {
		return expKey, fmt.Errorf("Error ton export encrypted key. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	}

	err = json.Unmarshal(resp.Raw, expKey)
	if err != nil {
		return expKey, err
	}
	return expKey, nil
}

// ImportPemKey importPemKey: import exported pem key
func (client *Client) ImportPemKey(pem string, pemPassword, localPass []byte) (key *TONPrivateKey, err error) {
	st := struct {
		Type          string `json:"@type"`
		KeyPassword   []byte `json:"key_password"`
		LocalPassword []byte `json:"local_password"`
		ExportedKey   struct {
			Pem string `json:"pem"`
		} `json:"exported_key"`
	}{
		Type:          "importPemKey",
		KeyPassword:   pemPassword,
		LocalPassword: localPass,
		ExportedKey: struct {
			Pem string `json:"pem"`
		}{
			Pem: pem,
		},
	}
	resp, err := client.executeAsynchronously(st)
	if err != nil {
		return key, err
	}
	if st, ok := resp.Data["@type"]; ok && st == "error" {
		return key, fmt.Errorf("Error ton import pem key. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	}

	key = new(TONPrivateKey)
	err = json.Unmarshal(resp.Raw, key)
	return key, err
}

// ImportEncryptedKey importEncryptedKey: import exported encrypted key
func (client *Client) ImportEncryptedKey(expKey *TONEncryptedKey, keyPassword, localPass []byte) (key *TONPrivateKey, err error) {
	st := struct {
		Type                 string           `json:"@type"`
		KeyPassword          []byte           `json:"key_password"`
		LocalPassword        []byte           `json:"local_password"`
		ExportedEncryptedKey *TONEncryptedKey `json:"exported_encrypted_key"`
	}{
		Type:                 "importEncryptedKey",
		KeyPassword:          keyPassword,
		LocalPassword:        localPass,
		ExportedEncryptedKey: expKey,
	}
	fmt.Println(expKey.Data)
	resp, err := client.executeAsynchronously(st)
	if err != nil {
		return key, err
	}
	if st, ok := resp.Data["@type"]; ok && st == "error" {
		return key, fmt.Errorf("Error ton import pem key. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	}

	key = new(TONPrivateKey)
	err = json.Unmarshal(resp.Raw, key)
	return key, err
}

// ImportKey importKey: import exported key
func (client *Client) ImportKey(wordList []string, mnemonicPass, localPass []byte) (key *TONPrivateKey, err error) {
	st := struct {
		Type             string `json:"@type"`
		MnemonicPassword []byte `json:"mnemonic_password"`
		LocalPassword    []byte `json:"local_password"`
		ExportedKey      struct {
			WordList []string `json:"word_list"`
		} `json:"exported_key"`
	}{
		Type:             "importKey",
		MnemonicPassword: mnemonicPass,
		LocalPassword:    localPass,
		ExportedKey: struct {
			WordList []string `json:"word_list"`
		}{
			WordList: wordList,
		},
	}
	resp, err := client.executeAsynchronously(st)
	if err != nil {
		return key, err
	}
	if st, ok := resp.Data["@type"]; ok && st == "error" {
		return key, fmt.Errorf("Error ton create private key. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	}

	key = new(TONPrivateKey)
	err = json.Unmarshal(resp.Raw, key)
	return key, err
}

// ChangeLocalPassword changeLocalPassword: change localPassword
func (client *Client) ChangeLocalPassword(key *TONPrivateKey, password, newPassword []byte) (*TONPrivateKey, error) {
	st := struct {
		Type             string   `json:"@type"`
		NewLocalPassword string   `json:"new_local_password"`
		InputKey         InputKey `json:"input_key"`
	}{
		Type:             "changeLocalPassword",
		NewLocalPassword: base64.StdEncoding.EncodeToString(password),
		InputKey:         key.getInputKey(password),
	}
	resp, err := client.executeAsynchronously(st)
	if err != nil {
		return key, err
	}
	if st, ok := resp.Data["@type"]; ok && st == "error" {
		return key, fmt.Errorf("Error ton change key passowrd. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	}
	key = new(TONPrivateKey)
	err = json.Unmarshal(resp.Raw, key)
	return key, err
}
