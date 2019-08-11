package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/jacekk/dead-simple-proxy-server/pkg/helpers"
	"github.com/pkg/errors"
)

const configPathEnv = "PROXY_CONFIG_PATH"

// ConfigItem -
type ConfigItem struct {
	URL         string            `json:"url"`
	BodyRewrite map[string]string `json:"bodyRewrite"`
}

type proxyConfig map[string]ConfigItem

// @todo replace with (sqlite) database.
func readConfig() (proxyConfig, error) {
	var config proxyConfig

	configPath := os.Getenv(configPathEnv)
	if configPath == "" {
		errMsg := fmt.Sprintf("Proxy config path is not set via '%s' system variable.", configPathEnv)
		return config, errors.New(errMsg)
	}

	fullPath := path.Join(helpers.GetProjectDir(), configPath)
	jsonFile, err := os.Open(fullPath)
	if err != nil {
		return config, err
	}

	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
