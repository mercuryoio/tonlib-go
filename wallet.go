package tonlib

import "C"
import (
	"encoding/json"
	"fmt"
)

type TonWallet struct {
	client *Client
}

// get TonWallet address
func (wallet *TonWallet) getAddress(pubKey string) (*TONAccountAddress, error) {
	//decodedKey, _ := base64.StdEncoding.DecodeString(pubKey)
	st := struct {
		Type                 string `json:"@type"`
		InititalAccountState struct {
			PublicKey string `json:"public_key"`
		} `json:"initital_account_state"`
	}{
		Type: "wallet.getAccountAddress",
		InititalAccountState: struct {
			PublicKey string `json:"public_key"`
		}{
			PublicKey: pubKey,
		},
	}
	resp, err := wallet.client.executeAsynchronously(st)
	if err != nil {
		return nil, err
	}
	if st, ok := resp.Data["@type"]; ok && st == "error" {
		return nil, fmt.Errorf("Error ton client init. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	}

	addressSt := struct {
		TONAccountAddress
		Type string `json:"@type"`
	}{}
	err = json.Unmarshal(resp.Raw, &addressSt)
	if err != nil {
		return nil, err
	}
	return &addressSt.TONAccountAddress, nil
}

// send gramm to address
func (wallet *TonWallet) sendGRAMM2Address(key *TONPrivateKey, password []byte, fromAddress, toAddress string, amount string) error {
	st := struct {
		Type        string            `json:"@type"`
		Seqno       int64             `json:"seqno"`
		Amount      string            `json:"amount"`
		PrivateKey  InputKey          `json:"private_key"`
		Destination TONAccountAddress `json:"destination"`
		ValidUntil  uint              `json:"valid_until"`
		Source      TONAccountAddress `json:"source"`
	}{
		Type:       "wallet.sendGrams",
		PrivateKey: key.getInputKey(password),
		Amount:     amount,
		Destination: TONAccountAddress{
			AccountAddress: toAddress,
		},
		Seqno: 2,
		Source: TONAccountAddress{
			AccountAddress: fromAddress,
		},
	}
	resp, err := wallet.client.executeAsynchronously(st)
	if err != nil {
		return err
	}
	if st, ok := resp.Data["@type"]; ok && st == "error" {
		return fmt.Errorf("Error ton send gramms. Code %v. Message %s. ", resp.Data["code"], resp.Data["message"])
	}

	var updateData TONResponse
	return json.Unmarshal(resp.Raw, &updateData)
}
