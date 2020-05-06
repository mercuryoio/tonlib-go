package v2

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"os"
)

// TonlibConfigServer stores config for tonlibjson.
type TonlibConfigServer struct {
	Liteservers []TonlibLiteServerConfig `json:"liteservers"`
	Validator   ValidatorConfig          `json:"validator"`
}

// TonlibLiteServerConfig stores lite-server credentials.
type TonlibLiteServerConfig struct {
	Type string            `json:"@type"`
	Ip   int64             `json:"ip"`
	Port string            `json:"port"`
	ID   map[string]string `json:"id"`
}

// ValidatorConfig stores validator config.
type ValidatorConfig struct {
	Type      string    `json:"@type"`
	ZeroState ZeroState `json:"zero_state"`
}

// ZeroState stores zero_state params from config.
type ZeroState struct {
	Workchain int    `json:"workchain"`
	Shard     int64  `json:"shard"`
	Seqno     int    `json:"seqno"`
	RootHash  string `json:"root_hash"`
	FileHash  string `json:"file_hash"`
}

// TonlibConfigFileConfig struct for config file network params.
type TonlibConfigFileConfig struct {
	Config                 TonlibConfigServer `json:"config"`
	BlockchainName         string             `json:"blockchain_name"`
	UseCallbacksForNetwork bool               `json:"use_callbacks_for_network"`
	IgnoreCache            bool               `json:"ignore_cache"`
}

// TonlibConfigFile handles key store type.
type TonlibConfigFile struct {
	Config   TonlibConfigFileConfig `json:"config"`
	Keystore KeyStoreType           `json:"keystore_type"`
}

// ParseConfigFile parses JSON config file.
func ParseConfigFile(path string) (*Options, error) {
	jsonFile, err := os.Open(path)
	defer jsonFile.Close()
	if err != nil {
		return nil, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	conf := new(TonlibConfigFile)
	err = json.Unmarshal(byteValue, &conf)
	if err != nil {
		return nil, err
	}

	// marshal back Internal config
	internalConfig, err := json.Marshal(&conf.Config.Config)
	if err != nil {
		return nil, err
	}

	return NewOptions(
		NewConfig(
			conf.Config.BlockchainName,
			string(internalConfig),
			conf.Config.IgnoreCache,
			conf.Config.UseCallbacksForNetwork,
		),
		&conf.Keystore,
	), nil
}

func hex2int(hexStr string) *big.Int {
	i := new(big.Int)
	i.SetString(hexStr, 16)
	return i
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !info.IsDir()
}
