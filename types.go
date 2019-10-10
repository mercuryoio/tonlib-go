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

// TONResult is used to unmarshal received json strings into
type TONResult struct {
	Data TONResponse
	Raw  []byte
}

// TONAccountAddress
type TONAccountAddress struct {
	AccountAddress string `json:"account_address"`
}

// GetHEXAddress
func (a TONAccountAddress) GetHEXAddress() string {
	data, err := base64url.Decode(a.AccountAddress)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", data)
}

// LocalKey
type LocalKey struct {
	PublicKey string `json:"public_key"`
	Secret    string `json:"secret"`
}

// InputKey
type InputKey struct {
	LocalPassword string   `json:"local_password"`
	Key           LocalKey `json:"key"`
}

// InternalTransactionId
type InternalTransactionId struct {
	Lt   string `json:"lt"`
	Hash string `json:"hash"`
}

// ValidatorConfig
type ValidatorConfig struct {
	Type      string    `json:"@type"`
	ZeroState ZeroState `json:"zero_state"`
}

// ZeroState
type ZeroState struct {
	Workchain int    `json:"workchain"`
	Shard     int64  `json:"shard"`
	Seqno     int    `json:"seqno"`
	RootHash  string `json:"root_hash"`
	FileHash  string `json:"file_hash"`
}

// TONConfigOption
type TONConfigOption struct {
	Type         string          `json:"@type"`
	Config       TONConfig       `json:"config"`
	KeystoreType TONKeystoreType `json:"keystore_type"`
}

// TONKeystoreType
type TONKeystoreType struct {
	Type      string `json:"@type"`
	Directory string `json:"directory"`
}

// TONConfig
type TONConfig struct {
	Config                 string `json:"config"`
	BlockchainName         string `json:"blockchain_name"`
	UseCallbacksForNetwork bool   `json:"use_callbacks_for_network"`
	IgnoreCache            bool   `json:"ignore_cache"`
}

// TONConfigServer
type TONConfigServer struct {
	Liteservers []TONLiteservierConfig `json:"liteservers"`
	Validator   ValidatorConfig        `json:"validator"`
}

// TONLiteservierConfig
type TONLiteservierConfig struct {
	Type string            `json:"@type"`
	Ip   int64             `json:"ip"`
	Port string            `json:"port"`
	ID   map[string]string `json:"id"`
}

// TONInitRequest
type TONInitRequest struct {
	Type    string          `json:"@type"`
	Options TONConfigOption `json:"options"`
}

// TONMsg
type TONMsg struct {
	Type        string          `json:"@type"`
	Sourse      string          `json:"sourse"`
	Destination string          `json:"destination"`
	Value       decimal.Decimal `json:"value"`
	Message     string          `json:"message"`
	FwdFee      decimal.Decimal `json:"fwd_fee"`
	IhrFee      decimal.Decimal `json:"ihr_fee"`
	CreatedLT   string          `json:"created_lt"`
	BodyHash    string          `json:"body_hash"`
}

// TONTransaction
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

// TONTransactionsResponse
type TONTransactionsResponse struct {
	Type         string           `json:"@type"`
	Transactions []TONTransaction `json:"transactions"`
}

// TONTransactionID
type TONTransactionID struct {
	Type string `json:"@type"`
	Lt   string `json:"lt"`
	Hash string `json:"hash"`
}

// TONAccountState
type TONAccountState struct {
	Type              string           `json:"@type"`
	Code              string           `json:"code"`
	Message           string           `json:"message"`
	Balance           decimal.Decimal  `json:"balance"`
	LastTransactionID TONTransactionID `json:"last_transaction_id"`
	FrozenHash        string           `json:"frozen_hash"`
	SyncUTime         uint
}

// TONUnpackedAddress
type TONUnpackedAddress struct {
	WorkchainID int    `json:"workchain_id"`
	Bounceable  bool   `json:"bounceable"`
	Testnet     bool   `json:"testnet"`
	Addr        string `json:"addr"`
	Type        string `json:"@type"`
}

// GetHEXAddress
func (a TONUnpackedAddress) GetHEXAddress() string {
	data, err := base64url.Decode(a.Addr)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", data)
}

// TONPrivateKey
type TONPrivateKey struct {
	PublicKey string `json:"public_key"`
	Secret    string `json:"secret"`
}

// TONPrivateKeyResponse
type TONPrivateKeyResponse struct {
	Type string `json:"@type"`
	TONPrivateKey
}

func (k TONPrivateKey) getInputKey(password []byte) InputKey {
	return InputKey{
		Key: LocalKey{
			PublicKey: k.PublicKey,
			Secret:    k.Secret,
		},
		LocalPassword: base64.StdEncoding.EncodeToString(password),
	}
}

// TONFileConfig
type TONFileConfig struct {
	Config struct {
		Config                 TONConfigServer `json:"config"`
		BlockchainName         string          `json:"blockchain_name"`
		UseCallbacksForNetwork bool            `json:"use_callbacks_for_network"`
		IgnoreCache            bool            `json:"ignore_cache"`
	} `json:"config"`
	KeystoreType TONKeystoreType `json:"keystore_type"`
}

// GetConfig
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
