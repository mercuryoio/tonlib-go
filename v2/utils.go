package v2

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type TONInitConfig struct {
	Config   Config       `json:"config"`
	Keystore KeyStoreType `json:"keystore_type"`
}

// ParseConfigFile parse JSON config file to
func ParseConfigFile(path string) (*TONInitConfig, error) {
	jsonFile, err := os.Open(path)
	defer jsonFile.Close()
	if err != nil {
		return nil, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	conf := new(TONInitConfig)
	err = json.Unmarshal(byteValue, &conf)
	if err != nil {
		return nil, err
	}

	conf.Config.Type = "config"
	return conf, nil
}
