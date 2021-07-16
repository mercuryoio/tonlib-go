package v2

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"os"
)

type TonlibConfigServer struct {
	Liteservers []TonlibListenserverConfig `json:"liteservers"`
	Validator   ValidatorConfig            `json:"validator"`
}

type TonlibListenserverConfig struct {
	Type string            `json:"@type"`
	Ip   int64             `json:"ip"`
	Port string            `json:"port"`
	ID   map[string]string `json:"id"`
}
type ValidatorConfig struct {
	Type      string      `json:"@type"`
	ZeroState InitBlock   `json:"zero_state"`
	InitBlock InitBlock   `json:"init_block,omitempty"`
	Hardforks []InitBlock `json:"hardforks,omitempty"`
}

type InitBlock struct {
	Workchain int    `json:"workchain"`
	Shard     int64  `json:"shard"`
	Seqno     int    `json:"seqno"`
	RootHash  string `json:"root_hash"`
	FileHash  string `json:"file_hash"`
}

type TonlibConfigFileConfig struct {
	Config                 TonlibConfigServer `json:"config"`
	BlockchainName         string             `json:"blockchain_name"`
	UseCallbacksForNetwork bool               `json:"use_callbacks_for_network"`
	IgnoreCache            bool               `json:"ignore_cache"`
}

type TonlibConfigFile struct {
	Config   TonlibConfigFileConfig `json:"config"`
	Keystore KeyStoreType           `json:"keystore_type"`
}

// ParseConfigFile parse JSON config file to
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
