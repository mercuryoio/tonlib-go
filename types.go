package tonlib

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dvsekhvalnov/jose2go/base64url"
	"github.com/shopspring/decimal"
)

// TONResponse alias for use in TONResult
type TONResponse map[string]interface{}

// TONResult is used to unmarshal recieved json strings into
type TONResult struct {
	Data TONResponse
	Raw  []byte
}

type TONAccountAddress struct {
	AccountAddress string `json:"account_address"`
}

func (a TONAccountAddress) GetHEXAddress() string {
	data, err := base64url.Decode(a.AccountAddress)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", data)
}

type InputKey struct {
	LocalPassword string        `json:"local_password"`
	Key           TONPrivateKey `json:"key"`
}

type InternalTransactionId struct {
	Lt   string `json:"lt"`
	Hash string `json:"hash"`
}

type ValidatorConfig struct {
	Type      string    `json:"@type"`
	ZeroState ZeroState `json:"zero_state"`
}
type ZeroState struct {
	Workchain int    `json:"workchain"`
	Shard     int64  `json:"shard"`
	Seqno     int    `json:"seqno"`
	RootHash  string `json:"root_hash"`
	FileHash  string `json:"file_hash"`
}

type TONConfigOption struct {
	Type         string          `json:"@type"`
	Config       TONConfig       `json:"config"`
	KeystoreType TONKeystoreType `json:"keystore_type"`
}

type TONKeystoreType struct {
	Type      string `json:"@type"`
	Directory string `json:"directory"`
}

type TONConfig struct {
	Config                 string `json:"config"`
	BlockchainName         string `json:"blockchain_name"`
	UseCallbacksForNetwork bool   `json:"use_callbacks_for_network"`
	IgnoreCache            bool   `json:"ignore_cache"`
}

type TONConfigServer struct {
	Liteservers []TONLiteservierConfig `json:"liteservers"`
	Validator   ValidatorConfig        `json:"validator"`
}

type TONLiteservierConfig struct {
	Type string            `json:"@type"`
	Ip   int64             `json:"ip"`
	Port string            `json:"port"`
	ID   map[string]string `json:"id"`
}

type TONInitRequest struct {
	Type    string          `json:"@type"`
	Options TONConfigOption `json:"options"`
}

type TONMsg struct {
	Type        string          `json:"@type"`
	Source      string          `json:"source"`
	Destination string          `json:"destination"`
	Value       decimal.Decimal `json:"value"`
	Message     string          `json:"message"`
	FwdFee      decimal.Decimal `json:"fwd_fee"`
	IhrFee      decimal.Decimal `json:"ihr_fee"`
	CreatedLT   string          `json:"created_lt"`
	BodyHash    string          `json:"body_hash"`
}

type TONTransaction struct {
	Type                  string           `json:"@type"`
	Utime                 uint             `json:"utime"`
	Data                  string           `json:"data"`
	TransactionID         TONTransactionID `json:"transaction_id"`
	PreviousTransactionID TONTransactionID `json:"previous_transaction_id"`
	StorageFee            decimal.Decimal  `json:"storage_fee"`
	OtherFee              decimal.Decimal  `json:"other_fee"`
	Fee                   decimal.Decimal  `json:"fee"`
	InMsg                 TONMsg           `json:"in_msg"`
	OutMsgs               []TONMsg         `json:"out_msgs"`
}

type TONTransactionsResponse struct {
	Type         string           `json:"@type"`
	Transactions []TONTransaction `json:"transactions"`
}

type TONTransactionID struct {
	Type string `json:"@type"`
	Lt   string `json:"lt"`
	Hash string `json:"hash"`
}

type TONAccountState struct {
	Type              string           `json:"@type"`
	Code              string           `json:"code"`
	Message           string           `json:"message"`
	Balance           decimal.Decimal  `json:"balance"`
	LastTransactionID TONTransactionID `json:"last_transaction_id"`
	FrozenHash        string           `json:"frozen_hash"`
	SyncUTime         uint
}

type TONUnpackedAddress struct {
	WorkchainID int    `json:"workchain_id"`
	Bounceable  bool   `json:"bounceable"`
	Testnet     bool   `json:"testnet"`
	Addr        string `json:"addr"`
	Type        string `json:"@type"`
}

func (a TONUnpackedAddress) GetHEXAddress() string {
	data, err := base64url.Decode(a.Addr)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", data)
}

type TONPrivateKey struct {
	PublicKey string `json:"public_key"`
	Secret    string `json:"secret"`
}

type TONPrivateKeyResponse struct {
	Type string `json:"@type"`
	TONPrivateKey
}

type TONSyncState struct {
	FromSeqno    int `json:"from_seqno"`
	ToSeqno      int `json:"to_seqno"`
	CurrentSeqno int `json:"current_seqno"`
}

func (k TONPrivateKey) getInputKey(password []byte) InputKey {
	return InputKey{
		Key: TONPrivateKey{
			PublicKey: k.PublicKey,
			Secret:    k.Secret,
		},
		LocalPassword: base64.StdEncoding.EncodeToString(password),
	}
}

type TONFileConfig struct {
	Config struct {
		Config                 TONConfigServer `json:"config"`
		BlockchainName         string          `json:"blockchain_name"`
		UseCallbacksForNetwork bool            `json:"use_callbacks_for_network"`
		IgnoreCache            bool            `json:"ignore_cache"`
	} `json:"config"`
	KeystoreType TONKeystoreType `json:"keystore_type"`
}

func (c TONFileConfig) GetConfig() *TONInitRequest {
	confStr, _ := json.Marshal(c.Config.Config)
	data := &TONInitRequest{
		Type: "init",
		Options: TONConfigOption{
			Type: "options",
			Config: TONConfig{
				Config:                 string(confStr),
				BlockchainName:         c.Config.BlockchainName,
				IgnoreCache:            c.Config.IgnoreCache,
				UseCallbacksForNetwork: c.Config.UseCallbacksForNetwork,
			},
			KeystoreType: c.KeystoreType,
		},
	}
	return data
}
