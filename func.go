package tonlib

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net"
	"os"
)

// add default values to custom config
func combineConfig(config Config) Config {
	if config.Timeout == 0 {
		config.Timeout = DEFAULT_TIMEOUT
	}
	return config
}

// return inet atom from IP
func InetAton(ip string) int64 {
	ip4 := net.ParseIP(ip)
	ipv4Int := big.NewInt(0)
	ipv4Int.SetBytes(ip4.To4())
	return ipv4Int.Int64()
}

// parse JSON config file to
func ParseConfigFile(path string) (*TONInitRequest, error) {
	jsonFile, err := os.Open(path)
	defer jsonFile.Close()
	if err != nil {
		return nil, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	conf := new(TONFileConfig)
	err = json.Unmarshal(byteValue, &conf)
	if err != nil {
		return nil, err
	}
	return conf.GetConfig(), nil
}
