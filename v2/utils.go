package v2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
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

func max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func min(x, y int64) int64 {
	if x > y {
		return y
	}
	return x
}

func computeTotalStake(l *[]ElectionParticipant, n, m_stake int64) (int64, error) {
	var totStake int64
	var i int64 = 0
	if n > int64(len(*l)) {
		return 0, fmt.Errorf("list has not enought length")
	}
	for i=0;  i < n; i++ {
		p := (*l)[i]
		stake, err := strconv.ParseInt(p.Stake, 10, 64)
		if err != nil{
			return 0, err
		}
		maxF, err := strconv.ParseInt(p.MaxFactor, 10, 64)
		if err != nil{
			return 0, err
		}
		stake = min(stake, (maxF* m_stake) >> 16);
		totStake += stake;
	}
	return totStake, nil
}