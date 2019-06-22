package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/jacekk/dead-simple-proxy-server/pkg/helpers"
	"github.com/pkg/errors"
)

const configPath = "storage/proxy-config.json"

type proxyConfig map[string]string

// @todo replace with (sqlite) database.
func readConfig() (proxyConfig, error) {
	var config proxyConfig

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

// GetProxyBySlug - returns URL based on given slug
func GetProxyBySlug(slug string) string {
	config, err := readConfig()
	if err != nil {
		log.Printf("%+v", errors.Wrap(err, "failed to read and parse proxy config"))

		return ""
	}

	return config[slug]
}
