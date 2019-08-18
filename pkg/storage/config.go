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

// Item -
type Item struct {
	BodyRewrite    map[string]string `json:"bodyRewrite"`
	ID             string            `json:"id"`
	IsCacheEnabled bool              `json:"isCacheEnabled"`
	URL            string            `json:"url"`
}

// Items -
type Items []Item

func readConfig() (Items, error) {
	var config Items

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
