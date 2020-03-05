package v2

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type tonCommon struct {
	Type  string `json:"@type"`
	Extra string `json:"@extra"`
}

type SecureBytes []byte
type SecureString string
type Bytes []byte
type TvmStackEntry interface{}
type SmcMethodId interface {}
type TvmNumber interface {}
type GenericAccountState string

// JSONInt64 alias for int64, in order to deal with json big number problem
type JSONInt64 int64

// MarshalJSON marshals to json
func (jsonInt *JSONInt64) MarshalJSON() ([]byte, error) {
	intStr := strconv.FormatInt(int64(*jsonInt), 10)
	return []byte(intStr), nil
}

// UnmarshalJSON unmarshals from json
func (jsonInt *JSONInt64) UnmarshalJSON(b []byte) error {
	intStr := string(b)
	intStr = strings.Replace(intStr, "\"", "", 2)
	jsonBigInt, err := strconv.ParseInt(intStr, 10, 64)
	if err != nil {
		return err
	}
	*jsonInt = JSONInt64(jsonBigInt)
	return nil
}

// TonMessage is the interface for all messages send and received to/from tonlib
type TonMessage interface {
	MessageType() string
}

// LogStreamEnum Alias for abstract LogStream 'Sub-Classes', used as constant-enum here
type LogStreamEnum string

// LogStream enums
const (
	LogStreamDefaultType LogStreamEnum = "logStreamDefault"
	LogStreamFileType    LogStreamEnum = "logStreamFile"
	LogStreamEmptyType   LogStreamEnum = "logStreamEmpty"
) // LogStream Describes a stream to which tonlib internal log is written
type LogStream interface {
	GetLogStreamEnum() LogStreamEnum
}

// Double
type Double struct {
	tonCommon
}

// MessageType return the string telegram-type of Double
func (double *Double) MessageType() string {
	return "double"
}

// NewDouble creates a new Double
//
func NewDouble() *Double {
	doubleTemp := Double{
		tonCommon: tonCommon{Type: "double"},
	}

	return &doubleTemp
}

// String
type String struct {
	tonCommon
}

// MessageType return the string telegram-type of String
func (string *String) MessageType() string {
	return "string"
}

// NewString creates a new String
//
func NewString() *String {
	stringTemp := String{
		tonCommon: tonCommon{Type: "string"},
	}

	return &stringTemp
}

// Int32
type Int32 struct {
	tonCommon
}

// MessageType return the string telegram-type of Int32
func (int32 *Int32) MessageType() string {
	return "int32"
}

// NewInt32 creates a new Int32
//
func NewInt32() *Int32 {
	int32Temp := Int32{
		tonCommon: tonCommon{Type: "int32"},
	}

	return &int32Temp
}

// Int53
type Int53 struct {
	tonCommon
}

// MessageType return the string telegram-type of Int53
func (int53 *Int53) MessageType() string {
	return "int53"
}

// NewInt53 creates a new Int53
//
func NewInt53() *Int53 {
	int53Temp := Int53{
		tonCommon: tonCommon{Type: "int53"},
	}

	return &int53Temp
}

// Int64
type Int64 struct {
	tonCommon
}

// MessageType return the string telegram-type of Int64
func (int64 *Int64) MessageType() string {
	return "int64"
}

// NewInt64 creates a new Int64
//
func NewInt64() *Int64 {
	int64Temp := Int64{
		tonCommon: tonCommon{Type: "int64"},
	}

	return &int64Temp
}

// Object
type Object struct {
	tonCommon
}

// MessageType return the string telegram-type of Object
func (object *Object) MessageType() string {
	return "object"
}

// NewObject creates a new Object
//
func NewObject() *Object {
	objectTemp := Object{
		tonCommon: tonCommon{Type: "object"},
	}

	return &objectTemp
}

// Function
type Function struct {
	tonCommon
}

// MessageType return the string telegram-type of Function
func (function *Function) MessageType() string {
	return "function"
}

// NewFunction creates a new Function
//
func NewFunction() *Function {
	functionTemp := Function{
		tonCommon: tonCommon{Type: "function"},
	}

	return &functionTemp
}

// BoolFalse
type BoolFalse struct {
	tonCommon
}

// MessageType return the string telegram-type of BoolFalse
func (boolFalse *BoolFalse) MessageType() string {
	return "boolFalse"
}

// NewBoolFalse creates a new BoolFalse
//
func NewBoolFalse() *BoolFalse {
	boolFalseTemp := BoolFalse{
		tonCommon: tonCommon{Type: "boolFalse"},
	}

	return &boolFalseTemp
}

// BoolTrue
type BoolTrue struct {
	tonCommon
}

// MessageType return the string telegram-type of BoolTrue
func (boolTrue *BoolTrue) MessageType() string {
	return "boolTrue"
}

// NewBoolTrue creates a new BoolTrue
//
func NewBoolTrue() *BoolTrue {
	boolTrueTemp := BoolTrue{
		tonCommon: tonCommon{Type: "boolTrue"},
	}

	return &boolTrueTemp
}

// Error
type Error struct {
	tonCommon
	Code    int32  `json:"code"`    //
	Message string `json:"message"` //
}

// MessageType return the string telegram-type of Error
func (error *Error) MessageType() string {
	return "error"
}

// NewError creates a new Error
//
// @param code
// @param message
func NewError(code int32, message string) *Error {
	errorTemp := Error{
		tonCommon: tonCommon{Type: "error"},
		Code:      code,
		Message:   message,
	}

	return &errorTemp
}

// Ok
type Ok struct {
	tonCommon
}

// MessageType return the string telegram-type of Ok
func (ok *Ok) MessageType() string {
	return "ok"
}

// NewOk creates a new Ok
//
func NewOk() *Ok {
	okTemp := Ok{
		tonCommon: tonCommon{Type: "ok"},
	}

	return &okTemp
}

// KeyStoreTypeDirectory
type KeyStoreTypeDirectory struct {
	tonCommon
	Directory string `json:"directory"` //
}

// MessageType return the string telegram-type of KeyStoreTypeDirectory
func (keyStoreTypeDirectory *KeyStoreTypeDirectory) MessageType() string {
	return "keyStoreTypeDirectory"
}

// NewKeyStoreTypeDirectory creates a new KeyStoreTypeDirectory
//
// @param directory
func NewKeyStoreTypeDirectory(directory string) *KeyStoreTypeDirectory {
	keyStoreTypeDirectoryTemp := KeyStoreTypeDirectory{
		tonCommon: tonCommon{Type: "keyStoreTypeDirectory"},
		Directory: directory,
	}

	return &keyStoreTypeDirectoryTemp
}

// KeyStoreTypeInMemory
type KeyStoreTypeInMemory struct {
	tonCommon
}

// MessageType return the string telegram-type of KeyStoreTypeInMemory
func (keyStoreTypeInMemory *KeyStoreTypeInMemory) MessageType() string {
	return "keyStoreTypeInMemory"
}

// NewKeyStoreTypeInMemory creates a new KeyStoreTypeInMemory
//
func NewKeyStoreTypeInMemory() *KeyStoreTypeInMemory {
	keyStoreTypeInMemoryTemp := KeyStoreTypeInMemory{
		tonCommon: tonCommon{Type: "keyStoreTypeInMemory"},
	}

	return &keyStoreTypeInMemoryTemp
}

// Config
type Config struct {
	tonCommon
	BlockchainName         string `json:"blockchain_name"`           //
	Config                 string `json:"config"`                    //
	IgnoreCache            bool   `json:"ignore_cache"`              //
	UseCallbacksForNetwork bool   `json:"use_callbacks_for_network"` //
}

// MessageType return the string telegram-type of Config
func (config *Config) MessageType() string {
	return "config"
}

// NewConfig creates a new Config
//
// @param blockchainName
// @param config
// @param ignoreCache
// @param useCallbacksForNetwork
func NewConfig(blockchainName string, config string, ignoreCache bool, useCallbacksForNetwork bool) *Config {
	configTemp := Config{
		tonCommon:              tonCommon{Type: "config"},
		BlockchainName:         blockchainName,
		Config:                 config,
		IgnoreCache:            ignoreCache,
		UseCallbacksForNetwork: useCallbacksForNetwork,
	}

	return &configTemp
}

// Options
type Options struct {
	tonCommon
	Config       *Config       `json:"config"`        //
	KeystoreType *KeyStoreType `json:"keystore_type"` //
}

// MessageType return the string telegram-type of Options
func (options *Options) MessageType() string {
	return "options"
}

// NewOptions creates a new Options
//
// @param config
// @param keystoreType
func NewOptions(config *Config, keystoreType *KeyStoreType) *Options {
	optionsTemp := Options{
		tonCommon:    tonCommon{Type: "options"},
		Config:       config,
		KeystoreType: keystoreType,
	}

	return &optionsTemp
}

// OptionsConfigInfo
type OptionsConfigInfo struct {
	tonCommon
	DefaultWalletId JSONInt64 `json:"default_wallet_id"` //
}

// MessageType return the string telegram-type of OptionsConfigInfo
func (optionsConfigInfo *OptionsConfigInfo) MessageType() string {
	return "options.configInfo"
}

// NewOptionsConfigInfo creates a new OptionsConfigInfo
//
// @param defaultWalletId
func NewOptionsConfigInfo(defaultWalletId JSONInt64) *OptionsConfigInfo {
	optionsConfigInfoTemp := OptionsConfigInfo{
		tonCommon:       tonCommon{Type: "options.configInfo"},
		DefaultWalletId: defaultWalletId,
	}

	return &optionsConfigInfoTemp
}

// OptionsInfo
type OptionsInfo struct {
	tonCommon
	ConfigInfo *OptionsConfigInfo `json:"config_info"` //
}

// MessageType return the string telegram-type of OptionsInfo
func (optionsInfo *OptionsInfo) MessageType() string {
	return "options.info"
}

// NewOptionsInfo creates a new OptionsInfo
//
// @param configInfo
func NewOptionsInfo(configInfo *OptionsConfigInfo) *OptionsInfo {
	optionsInfoTemp := OptionsInfo{
		tonCommon:  tonCommon{Type: "options.info"},
		ConfigInfo: configInfo,
	}

	return &optionsInfoTemp
}

// InputKeyRegular
type InputKeyRegular struct {
	tonCommon
	Key           *Key         `json:"key"`            //
	LocalPassword *SecureBytes `json:"local_password"` //
}

// MessageType return the string telegram-type of InputKeyRegular
func (inputKeyRegular *InputKeyRegular) MessageType() string {
	return "inputKeyRegular"
}

// NewInputKeyRegular creates a new InputKeyRegular
//
// @param key
// @param localPassword
func NewInputKeyRegular(key *Key, localPassword *SecureBytes) *InputKeyRegular {
	inputKeyRegularTemp := InputKeyRegular{
		tonCommon:     tonCommon{Type: "inputKeyRegular"},
		Key:           key,
		LocalPassword: localPassword,
	}

	return &inputKeyRegularTemp
}

// InputKeyFake
type InputKeyFake struct {
	tonCommon
}

// MessageType return the string telegram-type of InputKeyFake
func (inputKeyFake *InputKeyFake) MessageType() string {
	return "inputKeyFake"
}

// NewInputKeyFake creates a new InputKeyFake
//
func NewInputKeyFake() *InputKeyFake {
	inputKeyFakeTemp := InputKeyFake{
		tonCommon: tonCommon{Type: "inputKeyFake"},
	}

	return &inputKeyFakeTemp
}

// ExportedKey
type ExportedKey struct {
	tonCommon
	WordList []SecureString `json:"word_list"` //
}

// MessageType return the string telegram-type of ExportedKey
func (exportedKey *ExportedKey) MessageType() string {
	return "exportedKey"
}

// NewExportedKey creates a new ExportedKey
//
// @param wordList
func NewExportedKey(wordList []SecureString) *ExportedKey {
	exportedKeyTemp := ExportedKey{
		tonCommon: tonCommon{Type: "exportedKey"},
		WordList:  wordList,
	}

	return &exportedKeyTemp
}

// ExportedPemKey
type ExportedPemKey struct {
	tonCommon
	Pem *SecureString `json:"pem"` //
}

// MessageType return the string telegram-type of ExportedPemKey
func (exportedPemKey *ExportedPemKey) MessageType() string {
	return "exportedPemKey"
}

// NewExportedPemKey creates a new ExportedPemKey
//
// @param pem
func NewExportedPemKey(pem *SecureString) *ExportedPemKey {
	exportedPemKeyTemp := ExportedPemKey{
		tonCommon: tonCommon{Type: "exportedPemKey"},
		Pem:       pem,
	}

	return &exportedPemKeyTemp
}

// ExportedEncryptedKey
type ExportedEncryptedKey struct {
	tonCommon
	Data *SecureBytes `json:"data"` //
}

// MessageType return the string telegram-type of ExportedEncryptedKey
func (exportedEncryptedKey *ExportedEncryptedKey) MessageType() string {
	return "exportedEncryptedKey"
}

// NewExportedEncryptedKey creates a new ExportedEncryptedKey
//
// @param data
func NewExportedEncryptedKey(data *SecureBytes) *ExportedEncryptedKey {
	exportedEncryptedKeyTemp := ExportedEncryptedKey{
		tonCommon: tonCommon{Type: "exportedEncryptedKey"},
		Data:      data,
	}

	return &exportedEncryptedKeyTemp
}

// ExportedUnencryptedKey
type ExportedUnencryptedKey struct {
	tonCommon
	Data *SecureBytes `json:"data"` //
}

// MessageType return the string telegram-type of ExportedUnencryptedKey
func (exportedUnencryptedKey *ExportedUnencryptedKey) MessageType() string {
	return "exportedUnencryptedKey"
}

// NewExportedUnencryptedKey creates a new ExportedUnencryptedKey
//
// @param data
func NewExportedUnencryptedKey(data *SecureBytes) *ExportedUnencryptedKey {
	exportedUnencryptedKeyTemp := ExportedUnencryptedKey{
		tonCommon: tonCommon{Type: "exportedUnencryptedKey"},
		Data:      data,
	}

	return &exportedUnencryptedKeyTemp
}

// Bip39Hints
type Bip39Hints struct {
	tonCommon
	Words []string `json:"words"` //
}

// MessageType return the string telegram-type of Bip39Hints
func (bip39Hints *Bip39Hints) MessageType() string {
	return "bip39Hints"
}

// NewBip39Hints creates a new Bip39Hints
//
// @param words
func NewBip39Hints(words []string) *Bip39Hints {
	bip39HintsTemp := Bip39Hints{
		tonCommon: tonCommon{Type: "bip39Hints"},
		Words:     words,
	}

	return &bip39HintsTemp
}

// AdnlAddress
type AdnlAddress struct {
	tonCommon
	AdnlAddress string `json:"adnl_address"` //
}

// MessageType return the string telegram-type of AdnlAddress
func (adnlAddress *AdnlAddress) MessageType() string {
	return "adnlAddress"
}

// NewAdnlAddress creates a new AdnlAddress
//
// @param adnlAddress
func NewAdnlAddress(adnlAddress string) *AdnlAddress {
	adnlAddressTemp := AdnlAddress{
		tonCommon:   tonCommon{Type: "adnlAddress"},
		AdnlAddress: adnlAddress,
	}

	return &adnlAddressTemp
}

// AccountAddress
type AccountAddress struct {
	tonCommon
	AccountAddress string `json:"account_address"` //
}

// MessageType return the string telegram-type of AccountAddress
func (accountAddress *AccountAddress) MessageType() string {
	return "accountAddress"
}

// NewAccountAddress creates a new AccountAddress
//
// @param accountAddress
func NewAccountAddress(accountAddress string) *AccountAddress {
	accountAddressTemp := AccountAddress{
		tonCommon:      tonCommon{Type: "accountAddress"},
		AccountAddress: accountAddress,
	}

	return &accountAddressTemp
}

// AccountRevisionList
type AccountRevisionList struct {
	tonCommon
	Revisions []int32 `json:"revisions"` //
}

// MessageType return the string telegram-type of AccountRevisionList
func (accountRevisionList *AccountRevisionList) MessageType() string {
	return "accountRevisionList"
}

// NewAccountRevisionList creates a new AccountRevisionList
//
// @param revisions
func NewAccountRevisionList(revisions []int32) *AccountRevisionList {
	accountRevisionListTemp := AccountRevisionList{
		tonCommon: tonCommon{Type: "accountRevisionList"},
		Revisions: revisions,
	}

	return &accountRevisionListTemp
}

// UnpackedAccountAddress
type UnpackedAccountAddress struct {
	tonCommon
	Addr        string `json:"addr"`         //
	Bounceable  bool   `json:"bounceable"`   //
	Testnet     bool   `json:"testnet"`      //
	WorkchainId int32  `json:"workchain_id"` //
}

// MessageType return the string telegram-type of UnpackedAccountAddress
func (unpackedAccountAddress *UnpackedAccountAddress) MessageType() string {
	return "unpackedAccountAddress"
}

// NewUnpackedAccountAddress creates a new UnpackedAccountAddress
//
// @param addr
// @param bounceable
// @param testnet
// @param workchainId
func NewUnpackedAccountAddress(addr string, bounceable bool, testnet bool, workchainId int32) *UnpackedAccountAddress {
	unpackedAccountAddressTemp := UnpackedAccountAddress{
		tonCommon:   tonCommon{Type: "unpackedAccountAddress"},
		Addr:        addr,
		Bounceable:  bounceable,
		Testnet:     testnet,
		WorkchainId: workchainId,
	}

	return &unpackedAccountAddressTemp
}

// InternalTransactionId
type InternalTransactionId struct {
	tonCommon
	Hash string    `json:"hash"` //
	Lt   JSONInt64 `json:"lt"`   //
}

// MessageType return the string telegram-type of InternalTransactionId
func (internalTransactionId *InternalTransactionId) MessageType() string {
	return "internal.transactionId"
}

// NewInternalTransactionId creates a new InternalTransactionId
//
// @param hash
// @param lt
func NewInternalTransactionId(hash string, lt JSONInt64) *InternalTransactionId {
	internalTransactionIdTemp := InternalTransactionId{
		tonCommon: tonCommon{Type: "internal.transactionId"},
		Hash:      hash,
		Lt:        lt,
	}

	return &internalTransactionIdTemp
}

// TonBlockId
type TonBlockId struct {
	tonCommon
	Seqno     int32     `json:"seqno"`     //
	Shard     JSONInt64 `json:"shard"`     //
	Workchain int32     `json:"workchain"` //
}

// MessageType return the string telegram-type of TonBlockId
func (tonBlockId *TonBlockId) MessageType() string {
	return "ton.blockId"
}

// NewTonBlockId creates a new TonBlockId
//
// @param seqno
// @param shard
// @param workchain
func NewTonBlockId(seqno int32, shard JSONInt64, workchain int32) *TonBlockId {
	tonBlockIdTemp := TonBlockId{
		tonCommon: tonCommon{Type: "ton.blockId"},
		Seqno:     seqno,
		Shard:     shard,
		Workchain: workchain,
	}

	return &tonBlockIdTemp
}

// TonBlockIdExt
type TonBlockIdExt struct {
	tonCommon
	FileHash  string    `json:"file_hash"` //
	RootHash  string    `json:"root_hash"` //
	Seqno     int32     `json:"seqno"`     //
	Shard     JSONInt64 `json:"shard"`     //
	Workchain int32     `json:"workchain"` //
}

// MessageType return the string telegram-type of TonBlockIdExt
func (tonBlockIdExt *TonBlockIdExt) MessageType() string {
	return "ton.blockIdExt"
}

// NewTonBlockIdExt creates a new TonBlockIdExt
//
// @param fileHash
// @param rootHash
// @param seqno
// @param shard
// @param workchain
func NewTonBlockIdExt(fileHash string, rootHash string, seqno int32, shard JSONInt64, workchain int32) *TonBlockIdExt {
	tonBlockIdExtTemp := TonBlockIdExt{
		tonCommon: tonCommon{Type: "ton.blockIdExt"},
		FileHash:  fileHash,
		RootHash:  rootHash,
		Seqno:     seqno,
		Shard:     shard,
		Workchain: workchain,
	}

	return &tonBlockIdExtTemp
}

// RawFullAccountState
type RawFullAccountState struct {
	tonCommon
	Balance           JSONInt64              `json:"balance"`             //
	BlockId           *TonBlockIdExt         `json:"block_id"`            //
	Code              string                 `json:"code"`                //
	Data              string                 `json:"data"`                //
	FrozenHash        string                 `json:"frozen_hash"`         //
	LastTransactionId *InternalTransactionId `json:"last_transaction_id"` //
	SyncUtime         int64                  `json:"sync_utime"`          //
}

// MessageType return the string telegram-type of RawFullAccountState
func (rawFullAccountState *RawFullAccountState) MessageType() string {
	return "raw.fullAccountState"
}

// NewRawFullAccountState creates a new RawFullAccountState
//
// @param balance
// @param blockId
// @param code
// @param data
// @param frozenHash
// @param lastTransactionId
// @param syncUtime
func NewRawFullAccountState(balance JSONInt64, blockId *TonBlockIdExt, code string, data string, frozenHash string, lastTransactionId *InternalTransactionId, syncUtime int64) *RawFullAccountState {
	rawFullAccountStateTemp := RawFullAccountState{
		tonCommon:         tonCommon{Type: "raw.fullAccountState"},
		Balance:           balance,
		BlockId:           blockId,
		Code:              code,
		Data:              data,
		FrozenHash:        frozenHash,
		LastTransactionId: lastTransactionId,
		SyncUtime:         syncUtime,
	}

	return &rawFullAccountStateTemp
}

// RawMessage
type RawMessage struct {
	tonCommon
	BodyHash           string    `json:"body_hash"`            //
	CreatedLt          JSONInt64 `json:"created_lt"`           //
	Destination        string    `json:"destination"`          //
	FwdFee             JSONInt64 `json:"fwd_fee"`              //
	IhrFee             JSONInt64 `json:"ihr_fee"`              //
	IsMessageEncrypted bool      `json:"is_message_encrypted"` //
	Message            string    `json:"message"`              //
	Source             string    `json:"source"`               //
	Value              JSONInt64 `json:"value"`                //
}

// MessageType return the string telegram-type of RawMessage
func (rawMessage *RawMessage) MessageType() string {
	return "raw.message"
}

// NewRawMessage creates a new RawMessage
//
// @param bodyHash
// @param createdLt
// @param destination
// @param fwdFee
// @param ihrFee
// @param isMessageEncrypted
// @param message
// @param source
// @param value
func NewRawMessage(bodyHash string, createdLt JSONInt64, destination string, fwdFee JSONInt64, ihrFee JSONInt64, isMessageEncrypted bool, message string, source string, value JSONInt64) *RawMessage {
	rawMessageTemp := RawMessage{
		tonCommon:          tonCommon{Type: "raw.message"},
		BodyHash:           bodyHash,
		CreatedLt:          createdLt,
		Destination:        destination,
		FwdFee:             fwdFee,
		IhrFee:             ihrFee,
		IsMessageEncrypted: isMessageEncrypted,
		Message:            message,
		Source:             source,
		Value:              value,
	}

	return &rawMessageTemp
}

// RawTransaction
type RawTransaction struct {
	tonCommon
	Data          string                 `json:"data"`           //
	Fee           JSONInt64              `json:"fee"`            //
	InMsg         *RawMessage            `json:"in_msg"`         //
	OtherFee      JSONInt64              `json:"other_fee"`      //
	OutMsgs       []RawMessage           `json:"out_msgs"`       //
	StorageFee    JSONInt64              `json:"storage_fee"`    //
	TransactionId *InternalTransactionId `json:"transaction_id"` //
	Utime         int64                  `json:"utime"`          //
}

// MessageType return the string telegram-type of RawTransaction
func (rawTransaction *RawTransaction) MessageType() string {
	return "raw.transaction"
}

// NewRawTransaction creates a new RawTransaction
//
// @param data
// @param fee
// @param inMsg
// @param otherFee
// @param outMsgs
// @param storageFee
// @param transactionId
// @param utime
func NewRawTransaction(data string, fee JSONInt64, inMsg *RawMessage, otherFee JSONInt64, outMsgs []RawMessage, storageFee JSONInt64, transactionId *InternalTransactionId, utime int64) *RawTransaction {
	rawTransactionTemp := RawTransaction{
		tonCommon:     tonCommon{Type: "raw.transaction"},
		Data:          data,
		Fee:           fee,
		InMsg:         inMsg,
		OtherFee:      otherFee,
		OutMsgs:       outMsgs,
		StorageFee:    storageFee,
		TransactionId: transactionId,
		Utime:         utime,
	}

	return &rawTransactionTemp
}

// RawTransactions
type RawTransactions struct {
	tonCommon
	PreviousTransactionId *InternalTransactionId `json:"previous_transaction_id"` //
	Transactions          []RawTransaction       `json:"transactions"`            //
}

// MessageType return the string telegram-type of RawTransactions
func (rawTransactions *RawTransactions) MessageType() string {
	return "raw.transactions"
}

// NewRawTransactions creates a new RawTransactions
//
// @param previousTransactionId
// @param transactions
func NewRawTransactions(previousTransactionId *InternalTransactionId, transactions []RawTransaction) *RawTransactions {
	rawTransactionsTemp := RawTransactions{
		tonCommon:             tonCommon{Type: "raw.transactions"},
		PreviousTransactionId: previousTransactionId,
		Transactions:          transactions,
	}

	return &rawTransactionsTemp
}

// RawInitialAccountState
type RawInitialAccountState struct {
	tonCommon
	Code string `json:"code"` //
	Data string `json:"data"` //
}

// MessageType return the string telegram-type of RawInitialAccountState
func (rawInitialAccountState *RawInitialAccountState) MessageType() string {
	return "raw.initialAccountState"
}

// NewRawInitialAccountState creates a new RawInitialAccountState
//
// @param code
// @param data
func NewRawInitialAccountState(code string, data string) *RawInitialAccountState {
	rawInitialAccountStateTemp := RawInitialAccountState{
		tonCommon: tonCommon{Type: "raw.initialAccountState"},
		Code:      code,
		Data:      data,
	}

	return &rawInitialAccountStateTemp
}

// TestGiverInitialAccountState
type TestGiverInitialAccountState struct {
	tonCommon
}

// MessageType return the string telegram-type of TestGiverInitialAccountState
func (testGiverInitialAccountState *TestGiverInitialAccountState) MessageType() string {
	return "testGiver.initialAccountState"
}

// NewTestGiverInitialAccountState creates a new TestGiverInitialAccountState
//
func NewTestGiverInitialAccountState() *TestGiverInitialAccountState {
	testGiverInitialAccountStateTemp := TestGiverInitialAccountState{
		tonCommon: tonCommon{Type: "testGiver.initialAccountState"},
	}

	return &testGiverInitialAccountStateTemp
}

// TestWalletInitialAccountState
type TestWalletInitialAccountState struct {
	tonCommon
	PublicKey string `json:"public_key"` //
}

// MessageType return the string telegram-type of TestWalletInitialAccountState
func (testWalletInitialAccountState *TestWalletInitialAccountState) MessageType() string {
	return "testWallet.initialAccountState"
}

// NewTestWalletInitialAccountState creates a new TestWalletInitialAccountState
//
// @param publicKey
func NewTestWalletInitialAccountState(publicKey string) *TestWalletInitialAccountState {
	testWalletInitialAccountStateTemp := TestWalletInitialAccountState{
		tonCommon: tonCommon{Type: "testWallet.initialAccountState"},
		PublicKey: publicKey,
	}

	return &testWalletInitialAccountStateTemp
}

// WalletInitialAccountState
type WalletInitialAccountState struct {
	tonCommon
	PublicKey string `json:"public_key"` //
}

// MessageType return the string telegram-type of WalletInitialAccountState
func (walletInitialAccountState *WalletInitialAccountState) MessageType() string {
	return "wallet.initialAccountState"
}

// NewWalletInitialAccountState creates a new WalletInitialAccountState
//
// @param publicKey
func NewWalletInitialAccountState(publicKey string) *WalletInitialAccountState {
	walletInitialAccountStateTemp := WalletInitialAccountState{
		tonCommon: tonCommon{Type: "wallet.initialAccountState"},
		PublicKey: publicKey,
	}

	return &walletInitialAccountStateTemp
}

// WalletV3InitialAccountState
type WalletV3InitialAccountState struct {
	tonCommon
	PublicKey string    `json:"public_key"` //
	WalletId  JSONInt64 `json:"wallet_id"`  //
}

// MessageType return the string telegram-type of WalletV3InitialAccountState
func (walletV3InitialAccountState *WalletV3InitialAccountState) MessageType() string {
	return "wallet.v3.initialAccountState"
}

// NewWalletV3InitialAccountState creates a new WalletV3InitialAccountState
//
// @param publicKey
// @param walletId
func NewWalletV3InitialAccountState(publicKey string, walletId JSONInt64) *WalletV3InitialAccountState {
	walletV3InitialAccountStateTemp := WalletV3InitialAccountState{
		tonCommon: tonCommon{Type: "wallet.v3.initialAccountState"},
		PublicKey: publicKey,
		WalletId:  walletId,
	}

	return &walletV3InitialAccountStateTemp
}

// WalletHighloadV1InitialAccountState
type WalletHighloadV1InitialAccountState struct {
	tonCommon
	PublicKey string    `json:"public_key"` //
	WalletId  JSONInt64 `json:"wallet_id"`  //
}

// MessageType return the string telegram-type of WalletHighloadV1InitialAccountState
func (walletHighloadV1InitialAccountState *WalletHighloadV1InitialAccountState) MessageType() string {
	return "wallet.highload.v1.initialAccountState"
}

// NewWalletHighloadV1InitialAccountState creates a new WalletHighloadV1InitialAccountState
//
// @param publicKey
// @param walletId
func NewWalletHighloadV1InitialAccountState(publicKey string, walletId JSONInt64) *WalletHighloadV1InitialAccountState {
	walletHighloadV1InitialAccountStateTemp := WalletHighloadV1InitialAccountState{
		tonCommon: tonCommon{Type: "wallet.highload.v1.initialAccountState"},
		PublicKey: publicKey,
		WalletId:  walletId,
	}

	return &walletHighloadV1InitialAccountStateTemp
}

// WalletHighloadV2InitialAccountState
type WalletHighloadV2InitialAccountState struct {
	tonCommon
	PublicKey string    `json:"public_key"` //
	WalletId  JSONInt64 `json:"wallet_id"`  //
}

// MessageType return the string telegram-type of WalletHighloadV2InitialAccountState
func (walletHighloadV2InitialAccountState *WalletHighloadV2InitialAccountState) MessageType() string {
	return "wallet.highload.v2.initialAccountState"
}

// NewWalletHighloadV2InitialAccountState creates a new WalletHighloadV2InitialAccountState
//
// @param publicKey
// @param walletId
func NewWalletHighloadV2InitialAccountState(publicKey string, walletId JSONInt64) *WalletHighloadV2InitialAccountState {
	walletHighloadV2InitialAccountStateTemp := WalletHighloadV2InitialAccountState{
		tonCommon: tonCommon{Type: "wallet.highload.v2.initialAccountState"},
		PublicKey: publicKey,
		WalletId:  walletId,
	}

	return &walletHighloadV2InitialAccountStateTemp
}

// DnsInitialAccountState
type DnsInitialAccountState struct {
	tonCommon
	PublicKey string    `json:"public_key"` //
	WalletId  JSONInt64 `json:"wallet_id"`  //
}

// MessageType return the string telegram-type of DnsInitialAccountState
func (dnsInitialAccountState *DnsInitialAccountState) MessageType() string {
	return "dns.initialAccountState"
}

// NewDnsInitialAccountState creates a new DnsInitialAccountState
//
// @param publicKey
// @param walletId
func NewDnsInitialAccountState(publicKey string, walletId JSONInt64) *DnsInitialAccountState {
	dnsInitialAccountStateTemp := DnsInitialAccountState{
		tonCommon: tonCommon{Type: "dns.initialAccountState"},
		PublicKey: publicKey,
		WalletId:  walletId,
	}

	return &dnsInitialAccountStateTemp
}

// RawAccountState
type RawAccountState struct {
	tonCommon
	Code       string `json:"code"`        //
	Data       string `json:"data"`        //
	FrozenHash string `json:"frozen_hash"` //
}

// MessageType return the string telegram-type of RawAccountState
func (rawAccountState *RawAccountState) MessageType() string {
	return "raw.accountState"
}

// NewRawAccountState creates a new RawAccountState
//
// @param code
// @param data
// @param frozenHash
func NewRawAccountState(code string, data string, frozenHash string) *RawAccountState {
	rawAccountStateTemp := RawAccountState{
		tonCommon:  tonCommon{Type: "raw.accountState"},
		Code:       code,
		Data:       data,
		FrozenHash: frozenHash,
	}

	return &rawAccountStateTemp
}

// TestWalletAccountState
type TestWalletAccountState struct {
	tonCommon
	Seqno int32 `json:"seqno"` //
}

// MessageType return the string telegram-type of TestWalletAccountState
func (testWalletAccountState *TestWalletAccountState) MessageType() string {
	return "testWallet.accountState"
}

// NewTestWalletAccountState creates a new TestWalletAccountState
//
// @param seqno
func NewTestWalletAccountState(seqno int32) *TestWalletAccountState {
	testWalletAccountStateTemp := TestWalletAccountState{
		tonCommon: tonCommon{Type: "testWallet.accountState"},
		Seqno:     seqno,
	}

	return &testWalletAccountStateTemp
}

// WalletAccountState
type WalletAccountState struct {
	tonCommon
	Seqno int32 `json:"seqno"` //
}

// MessageType return the string telegram-type of WalletAccountState
func (walletAccountState *WalletAccountState) MessageType() string {
	return "wallet.accountState"
}

// NewWalletAccountState creates a new WalletAccountState
//
// @param seqno
func NewWalletAccountState(seqno int32) *WalletAccountState {
	walletAccountStateTemp := WalletAccountState{
		tonCommon: tonCommon{Type: "wallet.accountState"},
		Seqno:     seqno,
	}

	return &walletAccountStateTemp
}

// WalletV3AccountState
type WalletV3AccountState struct {
	tonCommon
	Seqno    int32     `json:"seqno"`     //
	WalletId JSONInt64 `json:"wallet_id"` //
}

// MessageType return the string telegram-type of WalletV3AccountState
func (walletV3AccountState *WalletV3AccountState) MessageType() string {
	return "wallet.v3.accountState"
}

// NewWalletV3AccountState creates a new WalletV3AccountState
//
// @param seqno
// @param walletId
func NewWalletV3AccountState(seqno int32, walletId JSONInt64) *WalletV3AccountState {
	walletV3AccountStateTemp := WalletV3AccountState{
		tonCommon: tonCommon{Type: "wallet.v3.accountState"},
		Seqno:     seqno,
		WalletId:  walletId,
	}

	return &walletV3AccountStateTemp
}

// WalletHighloadV1AccountState
type WalletHighloadV1AccountState struct {
	tonCommon
	Seqno    int32     `json:"seqno"`     //
	WalletId JSONInt64 `json:"wallet_id"` //
}

// MessageType return the string telegram-type of WalletHighloadV1AccountState
func (walletHighloadV1AccountState *WalletHighloadV1AccountState) MessageType() string {
	return "wallet.highload.v1.accountState"
}

// NewWalletHighloadV1AccountState creates a new WalletHighloadV1AccountState
//
// @param seqno
// @param walletId
func NewWalletHighloadV1AccountState(seqno int32, walletId JSONInt64) *WalletHighloadV1AccountState {
	walletHighloadV1AccountStateTemp := WalletHighloadV1AccountState{
		tonCommon: tonCommon{Type: "wallet.highload.v1.accountState"},
		Seqno:     seqno,
		WalletId:  walletId,
	}

	return &walletHighloadV1AccountStateTemp
}

// WalletHighloadV2AccountState
type WalletHighloadV2AccountState struct {
	tonCommon
	WalletId JSONInt64 `json:"wallet_id"` //
}

// MessageType return the string telegram-type of WalletHighloadV2AccountState
func (walletHighloadV2AccountState *WalletHighloadV2AccountState) MessageType() string {
	return "wallet.highload.v2.accountState"
}

// NewWalletHighloadV2AccountState creates a new WalletHighloadV2AccountState
//
// @param walletId
func NewWalletHighloadV2AccountState(walletId JSONInt64) *WalletHighloadV2AccountState {
	walletHighloadV2AccountStateTemp := WalletHighloadV2AccountState{
		tonCommon: tonCommon{Type: "wallet.highload.v2.accountState"},
		WalletId:  walletId,
	}

	return &walletHighloadV2AccountStateTemp
}

// TestGiverAccountState
type TestGiverAccountState struct {
	tonCommon
	Seqno int32 `json:"seqno"` //
}

// MessageType return the string telegram-type of TestGiverAccountState
func (testGiverAccountState *TestGiverAccountState) MessageType() string {
	return "testGiver.accountState"
}

// NewTestGiverAccountState creates a new TestGiverAccountState
//
// @param seqno
func NewTestGiverAccountState(seqno int32) *TestGiverAccountState {
	testGiverAccountStateTemp := TestGiverAccountState{
		tonCommon: tonCommon{Type: "testGiver.accountState"},
		Seqno:     seqno,
	}

	return &testGiverAccountStateTemp
}

// DnsAccountState
type DnsAccountState struct {
	tonCommon
	WalletId JSONInt64 `json:"wallet_id"` //
}

// MessageType return the string telegram-type of DnsAccountState
func (dnsAccountState *DnsAccountState) MessageType() string {
	return "dns.accountState"
}

// NewDnsAccountState creates a new DnsAccountState
//
// @param walletId
func NewDnsAccountState(walletId JSONInt64) *DnsAccountState {
	dnsAccountStateTemp := DnsAccountState{
		tonCommon: tonCommon{Type: "dns.accountState"},
		WalletId:  walletId,
	}

	return &dnsAccountStateTemp
}

// UninitedAccountState
type UninitedAccountState struct {
	tonCommon
	FrozenHash string `json:"frozen_hash"` //
}

// MessageType return the string telegram-type of UninitedAccountState
func (uninitedAccountState *UninitedAccountState) MessageType() string {
	return "uninited.accountState"
}

// NewUninitedAccountState creates a new UninitedAccountState
//
// @param frozenHash
func NewUninitedAccountState(frozenHash string) *UninitedAccountState {
	uninitedAccountStateTemp := UninitedAccountState{
		tonCommon:  tonCommon{Type: "uninited.accountState"},
		FrozenHash: frozenHash,
	}

	return &uninitedAccountStateTemp
}

// FullAccountState
type FullAccountState struct {
	tonCommon
	AccountState      *AccountState          `json:"account_state"`       //
	Balance           JSONInt64              `json:"balance"`             //
	BlockId           *TonBlockIdExt         `json:"block_id"`            //
	LastTransactionId *InternalTransactionId `json:"last_transaction_id"` //
	SyncUtime         int64                  `json:"sync_utime"`          //
}

// MessageType return the string telegram-type of FullAccountState
func (fullAccountState *FullAccountState) MessageType() string {
	return "fullAccountState"
}

// NewFullAccountState creates a new FullAccountState
//
// @param accountState
// @param balance
// @param blockId
// @param lastTransactionId
// @param syncUtime
func NewFullAccountState(accountState *AccountState, balance JSONInt64, blockId *TonBlockIdExt, lastTransactionId *InternalTransactionId, syncUtime int64) *FullAccountState {
	fullAccountStateTemp := FullAccountState{
		tonCommon:         tonCommon{Type: "fullAccountState"},
		AccountState:      accountState,
		Balance:           balance,
		BlockId:           blockId,
		LastTransactionId: lastTransactionId,
		SyncUtime:         syncUtime,
	}

	return &fullAccountStateTemp
}

// SyncStateDone
type SyncStateDone struct {
	tonCommon
}

// MessageType return the string telegram-type of SyncStateDone
func (syncStateDone *SyncStateDone) MessageType() string {
	return "syncStateDone"
}

// NewSyncStateDone creates a new SyncStateDone
//
func NewSyncStateDone() *SyncStateDone {
	syncStateDoneTemp := SyncStateDone{
		tonCommon: tonCommon{Type: "syncStateDone"},
	}

	return &syncStateDoneTemp
}

// SyncStateInProgress
type SyncStateInProgress struct {
	tonCommon
	CurrentSeqno int32 `json:"current_seqno"` //
	FromSeqno    int32 `json:"from_seqno"`    //
	ToSeqno      int32 `json:"to_seqno"`      //
}

// MessageType return the string telegram-type of SyncStateInProgress
func (syncStateInProgress *SyncStateInProgress) MessageType() string {
	return "syncStateInProgress"
}

// NewSyncStateInProgress creates a new SyncStateInProgress
//
// @param currentSeqno
// @param fromSeqno
// @param toSeqno
func NewSyncStateInProgress(currentSeqno int32, fromSeqno int32, toSeqno int32) *SyncStateInProgress {
	syncStateInProgressTemp := SyncStateInProgress{
		tonCommon:    tonCommon{Type: "syncStateInProgress"},
		CurrentSeqno: currentSeqno,
		FromSeqno:    fromSeqno,
		ToSeqno:      toSeqno,
	}

	return &syncStateInProgressTemp
}

// MsgDataText
type MsgDataText struct {
	tonCommon
	Text string `json:"text"` //
}

// MessageType return the string telegram-type of MsgDataText
func (msgDataText *MsgDataText) MessageType() string {
	return "msg.dataText"
}

// NewMsgDataText creates a new MsgDataText
//
// @param text
func NewMsgDataText(text string) *MsgDataText {
	msgDataTextTemp := MsgDataText{
		tonCommon: tonCommon{Type: "msg.dataText"},
		Text:      text,
	}

	return &msgDataTextTemp
}

// MsgDataEncryptedText
type MsgDataEncryptedText struct {
	tonCommon
	Text string `json:"text"` //
}

// MessageType return the string telegram-type of MsgDataEncryptedText
func (msgDataEncryptedText *MsgDataEncryptedText) MessageType() string {
	return "msg.dataEncryptedText"
}

// NewMsgDataEncryptedText creates a new MsgDataEncryptedText
//
// @param text
func NewMsgDataEncryptedText(text string) *MsgDataEncryptedText {
	msgDataEncryptedTextTemp := MsgDataEncryptedText{
		tonCommon: tonCommon{Type: "msg.dataEncryptedText"},
		Text:      text,
	}

	return &msgDataEncryptedTextTemp
}

// MsgMessage
type MsgMessage struct {
	tonCommon
	Amount      JSONInt64       `json:"amount"`      //
	Data        MsgData         `json:"data"`        //
	Destination *AccountAddress `json:"destination"` //
}

// MessageType return the string telegram-type of MsgMessage
func (msgMessage *MsgMessage) MessageType() string {
	return "msg.message"
}

// NewMsgMessage creates a new MsgMessage
//
// @param amount
// @param data
// @param destination
func NewMsgMessage(amount JSONInt64, data MsgData, destination *AccountAddress) *MsgMessage {
	msgMessageTemp := MsgMessage{
		tonCommon:   tonCommon{Type: "msg.message"},
		Amount:      amount,
		Data:        data,
		Destination: destination,
	}

	return &msgMessageTemp
}

// DnsEntryDataUnknown
type DnsEntryDataUnknown struct {
	tonCommon
	Bytes string `json:"bytes"` //
}

// MessageType return the string telegram-type of DnsEntryDataUnknown
func (dnsEntryDataUnknown *DnsEntryDataUnknown) MessageType() string {
	return "dns.entryDataUnknown"
}

// NewDnsEntryDataUnknown creates a new DnsEntryDataUnknown
//
// @param bytes
func NewDnsEntryDataUnknown(bytes string) *DnsEntryDataUnknown {
	dnsEntryDataUnknownTemp := DnsEntryDataUnknown{
		tonCommon: tonCommon{Type: "dns.entryDataUnknown"},
		Bytes:     bytes,
	}

	return &dnsEntryDataUnknownTemp
}

// DnsEntryDataText
type DnsEntryDataText struct {
	tonCommon
	Text string `json:"text"` //
}

// MessageType return the string telegram-type of DnsEntryDataText
func (dnsEntryDataText *DnsEntryDataText) MessageType() string {
	return "dns.entryDataText"
}

// NewDnsEntryDataText creates a new DnsEntryDataText
//
// @param text
func NewDnsEntryDataText(text string) *DnsEntryDataText {
	dnsEntryDataTextTemp := DnsEntryDataText{
		tonCommon: tonCommon{Type: "dns.entryDataText"},
		Text:      text,
	}

	return &dnsEntryDataTextTemp
}

// DnsEntryDataNextResolver
type DnsEntryDataNextResolver struct {
	tonCommon
	Resolver *AccountAddress `json:"resolver"` //
}

// MessageType return the string telegram-type of DnsEntryDataNextResolver
func (dnsEntryDataNextResolver *DnsEntryDataNextResolver) MessageType() string {
	return "dns.entryDataNextResolver"
}

// NewDnsEntryDataNextResolver creates a new DnsEntryDataNextResolver
//
// @param resolver
func NewDnsEntryDataNextResolver(resolver *AccountAddress) *DnsEntryDataNextResolver {
	dnsEntryDataNextResolverTemp := DnsEntryDataNextResolver{
		tonCommon: tonCommon{Type: "dns.entryDataNextResolver"},
		Resolver:  resolver,
	}

	return &dnsEntryDataNextResolverTemp
}

// DnsEntryDataSmcAddress
type DnsEntryDataSmcAddress struct {
	tonCommon
	SmcAddress *AccountAddress `json:"smc_address"` //
}

// MessageType return the string telegram-type of DnsEntryDataSmcAddress
func (dnsEntryDataSmcAddress *DnsEntryDataSmcAddress) MessageType() string {
	return "dns.entryDataSmcAddress"
}

// NewDnsEntryDataSmcAddress creates a new DnsEntryDataSmcAddress
//
// @param smcAddress
func NewDnsEntryDataSmcAddress(smcAddress *AccountAddress) *DnsEntryDataSmcAddress {
	dnsEntryDataSmcAddressTemp := DnsEntryDataSmcAddress{
		tonCommon:  tonCommon{Type: "dns.entryDataSmcAddress"},
		SmcAddress: smcAddress,
	}

	return &dnsEntryDataSmcAddressTemp
}

// DnsEntryDataAdnlAddress
type DnsEntryDataAdnlAddress struct {
	tonCommon
	AdnlAddress *AdnlAddress `json:"adnl_address"` //
}

// MessageType return the string telegram-type of DnsEntryDataAdnlAddress
func (dnsEntryDataAdnlAddress *DnsEntryDataAdnlAddress) MessageType() string {
	return "dns.entryDataAdnlAddress"
}

// NewDnsEntryDataAdnlAddress creates a new DnsEntryDataAdnlAddress
//
// @param adnlAddress
func NewDnsEntryDataAdnlAddress(adnlAddress *AdnlAddress) *DnsEntryDataAdnlAddress {
	dnsEntryDataAdnlAddressTemp := DnsEntryDataAdnlAddress{
		tonCommon:   tonCommon{Type: "dns.entryDataAdnlAddress"},
		AdnlAddress: adnlAddress,
	}

	return &dnsEntryDataAdnlAddressTemp
}

// DnsEntry
type DnsEntry struct {
	tonCommon
	Category int32         `json:"category"` //
	Entry    *DnsEntryData `json:"entry"`    //
	Name     string        `json:"name"`     //
}

// MessageType return the string telegram-type of DnsEntry
func (dnsEntry *DnsEntry) MessageType() string {
	return "dns.entry"
}

// NewDnsEntry creates a new DnsEntry
//
// @param category
// @param entry
// @param name
func NewDnsEntry(category int32, entry *DnsEntryData, name string) *DnsEntry {
	dnsEntryTemp := DnsEntry{
		tonCommon: tonCommon{Type: "dns.entry"},
		Category:  category,
		Entry:     entry,
		Name:      name,
	}

	return &dnsEntryTemp
}

// DnsActionDeleteAll
type DnsActionDeleteAll struct {
	tonCommon
}

// MessageType return the string telegram-type of DnsActionDeleteAll
func (dnsActionDeleteAll *DnsActionDeleteAll) MessageType() string {
	return "dns.actionDeleteAll"
}

// NewDnsActionDeleteAll creates a new DnsActionDeleteAll
//
func NewDnsActionDeleteAll() *DnsActionDeleteAll {
	dnsActionDeleteAllTemp := DnsActionDeleteAll{
		tonCommon: tonCommon{Type: "dns.actionDeleteAll"},
	}

	return &dnsActionDeleteAllTemp
}

// DnsActionDelete
type DnsActionDelete struct {
	tonCommon
	Category int32  `json:"category"` //
	Name     string `json:"name"`     //
}

// MessageType return the string telegram-type of DnsActionDelete
func (dnsActionDelete *DnsActionDelete) MessageType() string {
	return "dns.actionDelete"
}

// NewDnsActionDelete creates a new DnsActionDelete
//
// @param category
// @param name
func NewDnsActionDelete(category int32, name string) *DnsActionDelete {
	dnsActionDeleteTemp := DnsActionDelete{
		tonCommon: tonCommon{Type: "dns.actionDelete"},
		Category:  category,
		Name:      name,
	}

	return &dnsActionDeleteTemp
}

// DnsActionSet
type DnsActionSet struct {
	tonCommon
	Entry *DnsEntry `json:"entry"` //
}

// MessageType return the string telegram-type of DnsActionSet
func (dnsActionSet *DnsActionSet) MessageType() string {
	return "dns.actionSet"
}

// NewDnsActionSet creates a new DnsActionSet
//
// @param entry
func NewDnsActionSet(entry *DnsEntry) *DnsActionSet {
	dnsActionSetTemp := DnsActionSet{
		tonCommon: tonCommon{Type: "dns.actionSet"},
		Entry:     entry,
	}

	return &dnsActionSetTemp
}

// DnsResolved
type DnsResolved struct {
	tonCommon
	Entries []DnsEntry `json:"entries"` //
}

// MessageType return the string telegram-type of DnsResolved
func (dnsResolved *DnsResolved) MessageType() string {
	return "dns.resolved"
}

// NewDnsResolved creates a new DnsResolved
//
// @param entries
func NewDnsResolved(entries []DnsEntry) *DnsResolved {
	dnsResolvedTemp := DnsResolved{
		tonCommon: tonCommon{Type: "dns.resolved"},
		Entries:   entries,
	}

	return &dnsResolvedTemp
}

// ActionNoop
type ActionNoop struct {
	tonCommon
}

// MessageType return the string telegram-type of ActionNoop
func (actionNoop *ActionNoop) MessageType() string {
	return "actionNoop"
}

// NewActionNoop creates a new ActionNoop
//
func NewActionNoop() *ActionNoop {
	actionNoopTemp := ActionNoop{
		tonCommon: tonCommon{Type: "actionNoop"},
	}

	return &actionNoopTemp
}

// ActionMsg
type ActionMsg struct {
	tonCommon
	AllowSendToUninited bool         `json:"allow_send_to_uninited"` //
	Messages            []MsgMessage `json:"messages"`               //
}

// MessageType return the string telegram-type of ActionMsg
func (actionMsg *ActionMsg) MessageType() string {
	return "actionMsg"
}

// NewActionMsg creates a new ActionMsg
//
// @param allowSendToUninited
// @param messages
func NewActionMsg(allowSendToUninited bool, messages []MsgMessage) *ActionMsg {
	actionMsgTemp := ActionMsg{
		tonCommon:           tonCommon{Type: "actionMsg"},
		AllowSendToUninited: allowSendToUninited,
		Messages:            messages,
	}

	return &actionMsgTemp
}

// ActionDns
type ActionDns struct {
	tonCommon
	Actions []DnsAction `json:"actions"` //
}

// MessageType return the string telegram-type of ActionDns
func (actionDns *ActionDns) MessageType() string {
	return "actionDns"
}

// NewActionDns creates a new ActionDns
//
// @param actions
func NewActionDns(actions []DnsAction) *ActionDns {
	actionDnsTemp := ActionDns{
		tonCommon: tonCommon{Type: "actionDns"},
		Actions:   actions,
	}

	return &actionDnsTemp
}

// Fees
type Fees struct {
	tonCommon
	FwdFee     int64 `json:"fwd_fee"`     //
	GasFee     int64 `json:"gas_fee"`     //
	InFwdFee   int64 `json:"in_fwd_fee"`  //
	StorageFee int64 `json:"storage_fee"` //
}

// MessageType return the string telegram-type of Fees
func (fees *Fees) MessageType() string {
	return "fees"
}

// NewFees creates a new Fees
//
// @param fwdFee
// @param gasFee
// @param inFwdFee
// @param storageFee
func NewFees(fwdFee int64, gasFee int64, inFwdFee int64, storageFee int64) *Fees {
	feesTemp := Fees{
		tonCommon:  tonCommon{Type: "fees"},
		FwdFee:     fwdFee,
		GasFee:     gasFee,
		InFwdFee:   inFwdFee,
		StorageFee: storageFee,
	}

	return &feesTemp
}

// QueryFees
type QueryFees struct {
	tonCommon
	DestinationFees []Fees `json:"destination_fees"` //
	SourceFees      *Fees  `json:"source_fees"`      //
}

// MessageType return the string telegram-type of QueryFees
func (queryFees *QueryFees) MessageType() string {
	return "query.fees"
}

// NewQueryFees creates a new QueryFees
//
// @param destinationFees
// @param sourceFees
func NewQueryFees(destinationFees []Fees, sourceFees *Fees) *QueryFees {
	queryFeesTemp := QueryFees{
		tonCommon:       tonCommon{Type: "query.fees"},
		DestinationFees: destinationFees,
		SourceFees:      sourceFees,
	}

	return &queryFeesTemp
}

// QueryInfo
type QueryInfo struct {
	tonCommon
	BodyHash   string `json:"body_hash"`   //
	Id         int64  `json:"id"`          //
	ValidUntil int64  `json:"valid_until"` //
}

// MessageType return the string telegram-type of QueryInfo
func (queryInfo *QueryInfo) MessageType() string {
	return "query.info"
}

// NewQueryInfo creates a new QueryInfo
//
// @param bodyHash
// @param id
// @param validUntil
func NewQueryInfo(bodyHash string, id int64, validUntil int64) *QueryInfo {
	queryInfoTemp := QueryInfo{
		tonCommon:  tonCommon{Type: "query.info"},
		BodyHash:   bodyHash,
		Id:         id,
		ValidUntil: validUntil,
	}

	return &queryInfoTemp
}

// TvmSlice
type TvmSlice struct {
	tonCommon
	Bytes string `json:"bytes"` //
}

// MessageType return the string telegram-type of TvmSlice
func (tvmSlice *TvmSlice) MessageType() string {
	return "tvm.slice"
}

// NewTvmSlice creates a new TvmSlice
//
// @param bytes
func NewTvmSlice(bytes string) *TvmSlice {
	tvmSliceTemp := TvmSlice{
		tonCommon: tonCommon{Type: "tvm.slice"},
		Bytes:     bytes,
	}

	return &tvmSliceTemp
}

// TvmCell
type TvmCell struct {
	tonCommon
	Bytes string `json:"bytes"` //
}

// MessageType return the string telegram-type of TvmCell
func (tvmCell *TvmCell) MessageType() string {
	return "tvm.cell"
}

// NewTvmCell creates a new TvmCell
//
// @param bytes
func NewTvmCell(bytes string) *TvmCell {
	tvmCellTemp := TvmCell{
		tonCommon: tonCommon{Type: "tvm.cell"},
		Bytes:     bytes,
	}

	return &tvmCellTemp
}

// TvmNumberDecimal
type TvmNumberDecimal struct {
	tonCommon
	Number string `json:"number"` //
}

// MessageType return the string telegram-type of TvmNumberDecimal
func (tvmNumberDecimal *TvmNumberDecimal) MessageType() string {
	return "tvm.numberDecimal"
}

// NewTvmNumberDecimal creates a new TvmNumberDecimal
//
// @param number
func NewTvmNumberDecimal(number string) *TvmNumberDecimal {
	tvmNumberDecimalTemp := TvmNumberDecimal{
		tonCommon: tonCommon{Type: "tvm.numberDecimal"},
		Number:    number,
	}

	return &tvmNumberDecimalTemp
}

// TvmTuple
type TvmTuple struct {
	tonCommon
	Elements []TvmStackEntry `json:"elements"` //
}

// MessageType return the string telegram-type of TvmTuple
func (tvmTuple *TvmTuple) MessageType() string {
	return "tvm.tuple"
}

// NewTvmTuple creates a new TvmTuple
//
// @param elements
func NewTvmTuple(elements []TvmStackEntry) *TvmTuple {
	tvmTupleTemp := TvmTuple{
		tonCommon: tonCommon{Type: "tvm.tuple"},
		Elements:  elements,
	}

	return &tvmTupleTemp
}

// TvmList
type TvmList struct {
	tonCommon
	Elements []TvmStackEntry `json:"elements"` //
}

// MessageType return the string telegram-type of TvmList
func (tvmList *TvmList) MessageType() string {
	return "tvm.list"
}

// NewTvmList creates a new TvmList
//
// @param elements
func NewTvmList(elements []TvmStackEntry) *TvmList {
	tvmListTemp := TvmList{
		tonCommon: tonCommon{Type: "tvm.list"},
		Elements:  elements,
	}

	return &tvmListTemp
}

// TvmStackEntrySlice
type TvmStackEntrySlice struct {
	tonCommon
	Slice *TvmSlice `json:"slice"` //
}

// MessageType return the string telegram-type of TvmStackEntrySlice
func (tvmStackEntrySlice *TvmStackEntrySlice) MessageType() string {
	return "tvm.stackEntrySlice"
}

// NewTvmStackEntrySlice creates a new TvmStackEntrySlice
//
// @param slice
func NewTvmStackEntrySlice(slice *TvmSlice) *TvmStackEntrySlice {
	tvmStackEntrySliceTemp := TvmStackEntrySlice{
		tonCommon: tonCommon{Type: "tvm.stackEntrySlice"},
		Slice:     slice,
	}

	return &tvmStackEntrySliceTemp
}

// TvmStackEntryCell
type TvmStackEntryCell struct {
	tonCommon
	Cell *TvmCell `json:"cell"` //
}

// MessageType return the string telegram-type of TvmStackEntryCell
func (tvmStackEntryCell *TvmStackEntryCell) MessageType() string {
	return "tvm.stackEntryCell"
}

// NewTvmStackEntryCell creates a new TvmStackEntryCell
//
// @param cell
func NewTvmStackEntryCell(cell *TvmCell) *TvmStackEntryCell {
	tvmStackEntryCellTemp := TvmStackEntryCell{
		tonCommon: tonCommon{Type: "tvm.stackEntryCell"},
		Cell:      cell,
	}

	return &tvmStackEntryCellTemp
}

// TvmStackEntryNumber
type TvmStackEntryNumber struct {
	tonCommon
	Number TvmNumber `json:"number"` //
}

// MessageType return the string telegram-type of TvmStackEntryNumber
func (tvmStackEntryNumber *TvmStackEntryNumber) MessageType() string {
	return "tvm.stackEntryNumber"
}

// NewTvmStackEntryNumber creates a new TvmStackEntryNumber
//
// @param number
func NewTvmStackEntryNumber(number TvmNumber) *TvmStackEntryNumber {
	tvmStackEntryNumberTemp := TvmStackEntryNumber{
		tonCommon: tonCommon{Type: "tvm.stackEntryNumber"},
		Number:    number,
	}

	return &tvmStackEntryNumberTemp
}

// TvmStackEntryTuple
type TvmStackEntryTuple struct {
	tonCommon
	Tuple *TvmTuple `json:"tuple"` //
}

// MessageType return the string telegram-type of TvmStackEntryTuple
func (tvmStackEntryTuple *TvmStackEntryTuple) MessageType() string {
	return "tvm.stackEntryTuple"
}

// NewTvmStackEntryTuple creates a new TvmStackEntryTuple
//
// @param tuple
func NewTvmStackEntryTuple(tuple *TvmTuple) *TvmStackEntryTuple {
	tvmStackEntryTupleTemp := TvmStackEntryTuple{
		tonCommon: tonCommon{Type: "tvm.stackEntryTuple"},
		Tuple:     tuple,
	}

	return &tvmStackEntryTupleTemp
}

// TvmStackEntryList
type TvmStackEntryList struct {
	tonCommon
	List *TvmList `json:"list"` //
}

// MessageType return the string telegram-type of TvmStackEntryList
func (tvmStackEntryList *TvmStackEntryList) MessageType() string {
	return "tvm.stackEntryList"
}

// NewTvmStackEntryList creates a new TvmStackEntryList
//
// @param list
func NewTvmStackEntryList(list *TvmList) *TvmStackEntryList {
	tvmStackEntryListTemp := TvmStackEntryList{
		tonCommon: tonCommon{Type: "tvm.stackEntryList"},
		List:      list,
	}

	return &tvmStackEntryListTemp
}

// TvmStackEntryUnsupported
type TvmStackEntryUnsupported struct {
	tonCommon
}

// MessageType return the string telegram-type of TvmStackEntryUnsupported
func (tvmStackEntryUnsupported *TvmStackEntryUnsupported) MessageType() string {
	return "tvm.stackEntryUnsupported"
}

// NewTvmStackEntryUnsupported creates a new TvmStackEntryUnsupported
//
func NewTvmStackEntryUnsupported() *TvmStackEntryUnsupported {
	tvmStackEntryUnsupportedTemp := TvmStackEntryUnsupported{
		tonCommon: tonCommon{Type: "tvm.stackEntryUnsupported"},
	}

	return &tvmStackEntryUnsupportedTemp
}

// SmcInfo
type SmcInfo struct {
	tonCommon
	Id int64 `json:"id"` //
}

// MessageType return the string telegram-type of SmcInfo
func (smcInfo *SmcInfo) MessageType() string {
	return "smc.info"
}

// NewSmcInfo creates a new SmcInfo
//
// @param id
func NewSmcInfo(id int64) *SmcInfo {
	smcInfoTemp := SmcInfo{
		tonCommon: tonCommon{Type: "smc.info"},
		Id:        id,
	}

	return &smcInfoTemp
}

// SmcMethodIdNumber
type SmcMethodIdNumber struct {
	tonCommon
	Number int32 `json:"number"` //
}

// MessageType return the string telegram-type of SmcMethodIdNumber
func (smcMethodIdNumber *SmcMethodIdNumber) MessageType() string {
	return "smc.methodIdNumber"
}

// NewSmcMethodIdNumber creates a new SmcMethodIdNumber
//
// @param number
func NewSmcMethodIdNumber(number int32) *SmcMethodIdNumber {
	smcMethodIdNumberTemp := SmcMethodIdNumber{
		tonCommon: tonCommon{Type: "smc.methodIdNumber"},
		Number:    number,
	}

	return &smcMethodIdNumberTemp
}

// SmcMethodIdName
type SmcMethodIdName struct {
	tonCommon
	Name string `json:"name"` //
}

// MessageType return the string telegram-type of SmcMethodIdName
func (smcMethodIdName *SmcMethodIdName) MessageType() string {
	return "smc.methodIdName"
}

// NewSmcMethodIdName creates a new SmcMethodIdName
//
// @param name
func NewSmcMethodIdName(name string) *SmcMethodIdName {
	smcMethodIdNameTemp := SmcMethodIdName{
		tonCommon: tonCommon{Type: "smc.methodIdName"},
		Name:      name,
	}

	return &smcMethodIdNameTemp
}

// SmcRunResult
type SmcRunResult struct {
	tonCommon
	ExitCode int32           `json:"exit_code"` //
	GasUsed  int64           `json:"gas_used"`  //
	Stack    []TvmStackEntry `json:"stack"`     //
}

// MessageType return the string telegram-type of SmcRunResult
func (smcRunResult *SmcRunResult) MessageType() string {
	return "smc.runResult"
}

// NewSmcRunResult creates a new SmcRunResult
//
// @param exitCode
// @param gasUsed
// @param stack
func NewSmcRunResult(exitCode int32, gasUsed int64, stack []TvmStackEntry) *SmcRunResult {
	smcRunResultTemp := SmcRunResult{
		tonCommon: tonCommon{Type: "smc.runResult"},
		ExitCode:  exitCode,
		GasUsed:   gasUsed,
		Stack:     stack,
	}

	return &smcRunResultTemp
}

// UpdateSendLiteServerQuery
type UpdateSendLiteServerQuery struct {
	tonCommon
	Data string    `json:"data"` //
	Id   JSONInt64 `json:"id"`   //
}

// MessageType return the string telegram-type of UpdateSendLiteServerQuery
func (updateSendLiteServerQuery *UpdateSendLiteServerQuery) MessageType() string {
	return "updateSendLiteServerQuery"
}

// NewUpdateSendLiteServerQuery creates a new UpdateSendLiteServerQuery
//
// @param data
// @param id
func NewUpdateSendLiteServerQuery(data string, id JSONInt64) *UpdateSendLiteServerQuery {
	updateSendLiteServerQueryTemp := UpdateSendLiteServerQuery{
		tonCommon: tonCommon{Type: "updateSendLiteServerQuery"},
		Data:      data,
		Id:        id,
	}

	return &updateSendLiteServerQueryTemp
}

// UpdateSyncState
type UpdateSyncState struct {
	tonCommon
	SyncState *SyncState `json:"sync_state"` //
}

// MessageType return the string telegram-type of UpdateSyncState
func (updateSyncState *UpdateSyncState) MessageType() string {
	return "updateSyncState"
}

// NewUpdateSyncState creates a new UpdateSyncState
//
// @param syncState
func NewUpdateSyncState(syncState *SyncState) *UpdateSyncState {
	updateSyncStateTemp := UpdateSyncState{
		tonCommon: tonCommon{Type: "updateSyncState"},
		SyncState: syncState,
	}

	return &updateSyncStateTemp
}

// LogStreamDefault The log is written to stderr or an OS specific log
type LogStreamDefault struct {
	tonCommon
}

// MessageType return the string telegram-type of LogStreamDefault
func (logStreamDefault *LogStreamDefault) MessageType() string {
	return "logStreamDefault"
}

// NewLogStreamDefault creates a new LogStreamDefault
//
func NewLogStreamDefault() *LogStreamDefault {
	logStreamDefaultTemp := LogStreamDefault{
		tonCommon: tonCommon{Type: "logStreamDefault"},
	}

	return &logStreamDefaultTemp
}

// GetLogStreamEnum return the enum type of this object
func (logStreamDefault *LogStreamDefault) GetLogStreamEnum() LogStreamEnum {
	return LogStreamDefaultType
}

// LogStreamFile The log is written to a file
type LogStreamFile struct {
	tonCommon
	MaxFileSize int64  `json:"max_file_size"` // Maximum size of the file to where the internal tonlib log is written before the file will be auto-rotated
	Path        string `json:"path"`          // Path to the file to where the internal tonlib log will be written
}

// MessageType return the string telegram-type of LogStreamFile
func (logStreamFile *LogStreamFile) MessageType() string {
	return "logStreamFile"
}

// NewLogStreamFile creates a new LogStreamFile
//
// @param maxFileSize Maximum size of the file to where the internal tonlib log is written before the file will be auto-rotated
// @param path Path to the file to where the internal tonlib log will be written
func NewLogStreamFile(maxFileSize int64, path string) *LogStreamFile {
	logStreamFileTemp := LogStreamFile{
		tonCommon:   tonCommon{Type: "logStreamFile"},
		MaxFileSize: maxFileSize,
		Path:        path,
	}

	return &logStreamFileTemp
}

// GetLogStreamEnum return the enum type of this object
func (logStreamFile *LogStreamFile) GetLogStreamEnum() LogStreamEnum {
	return LogStreamFileType
}

// LogStreamEmpty The log is written nowhere
type LogStreamEmpty struct {
	tonCommon
}

// MessageType return the string telegram-type of LogStreamEmpty
func (logStreamEmpty *LogStreamEmpty) MessageType() string {
	return "logStreamEmpty"
}

// NewLogStreamEmpty creates a new LogStreamEmpty
//
func NewLogStreamEmpty() *LogStreamEmpty {
	logStreamEmptyTemp := LogStreamEmpty{
		tonCommon: tonCommon{Type: "logStreamEmpty"},
	}

	return &logStreamEmptyTemp
}

// GetLogStreamEnum return the enum type of this object
func (logStreamEmpty *LogStreamEmpty) GetLogStreamEnum() LogStreamEnum {
	return LogStreamEmptyType
}

// LogVerbosityLevel Contains a tonlib internal log verbosity level
type LogVerbosityLevel struct {
	tonCommon
	VerbosityLevel int32 `json:"verbosity_level"` // Log verbosity level
}

// MessageType return the string telegram-type of LogVerbosityLevel
func (logVerbosityLevel *LogVerbosityLevel) MessageType() string {
	return "logVerbosityLevel"
}

// NewLogVerbosityLevel creates a new LogVerbosityLevel
//
// @param verbosityLevel Log verbosity level
func NewLogVerbosityLevel(verbosityLevel int32) *LogVerbosityLevel {
	logVerbosityLevelTemp := LogVerbosityLevel{
		tonCommon:      tonCommon{Type: "logVerbosityLevel"},
		VerbosityLevel: verbosityLevel,
	}

	return &logVerbosityLevelTemp
}

// LogTags Contains a list of available tonlib internal log tags
type LogTags struct {
	tonCommon
	Tags []string `json:"tags"` // List of log tags
}

// MessageType return the string telegram-type of LogTags
func (logTags *LogTags) MessageType() string {
	return "logTags"
}

// NewLogTags creates a new LogTags
//
// @param tags List of log tags
func NewLogTags(tags []string) *LogTags {
	logTagsTemp := LogTags{
		tonCommon: tonCommon{Type: "logTags"},
		Tags:      tags,
	}

	return &logTagsTemp
}

// Data
type Data struct {
	tonCommon
	Bytes *SecureBytes `json:"bytes"` //
}

// MessageType return the string telegram-type of Data
func (data *Data) MessageType() string {
	return "data"
}

// NewData creates a new Data
//
// @param bytes
func NewData(bytes *SecureBytes) *Data {
	dataTemp := Data{
		tonCommon: tonCommon{Type: "data"},
		Bytes:     bytes,
	}

	return &dataTemp
}

// LiteServerInfo
type LiteServerInfo struct {
	tonCommon
	Capabilities JSONInt64 `json:"capabilities"` //
	Now          int64     `json:"now"`          //
	Version      int32     `json:"version"`      //
}

// MessageType return the string telegram-type of LiteServerInfo
func (liteServerInfo *LiteServerInfo) MessageType() string {
	return "liteServer.info"
}

// NewLiteServerInfo creates a new LiteServerInfo
//
// @param capabilities
// @param now
// @param version
func NewLiteServerInfo(capabilities JSONInt64, now int64, version int32) *LiteServerInfo {
	liteServerInfoTemp := LiteServerInfo{
		tonCommon:    tonCommon{Type: "liteServer.info"},
		Capabilities: capabilities,
		Now:          now,
		Version:      version,
	}

	return &liteServerInfoTemp
}

func unmarshalLogStream(rawMsg *json.RawMessage) (LogStream, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch LogStreamEnum(objMap["@type"].(string)) {
	case LogStreamDefaultType:
		var logStreamDefault LogStreamDefault
		err := json.Unmarshal(*rawMsg, &logStreamDefault)
		return &logStreamDefault, err

	case LogStreamFileType:
		var logStreamFile LogStreamFile
		err := json.Unmarshal(*rawMsg, &logStreamFile)
		return &logStreamFile, err

	case LogStreamEmptyType:
		var logStreamEmpty LogStreamEmpty
		err := json.Unmarshal(*rawMsg, &logStreamEmpty)
		return &logStreamEmpty, err

	default:
		return nil, fmt.Errorf("Error unmarshaling, unknown type:" + objMap["@type"].(string))
	}
}
