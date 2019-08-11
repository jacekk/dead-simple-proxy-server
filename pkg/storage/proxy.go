package storage

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

// GetProxyConfigBySlug - returns URL based on given slug
func GetProxyConfigBySlug(slug string) (ConfigItem, error) {
	var item ConfigItem
	config, err := readConfig()
	if err != nil {
		log.Printf("%+v", errors.Wrap(err, "failed to read and parse proxy config"))
		return item, err
	}

	item, isFound := config[slug]
	if !isFound {
		errMsg := fmt.Sprintf("failed to find proxy config for %q slug", slug)
		log.Print(errMsg)
		return item, errors.New(errMsg)
	}

	return item, nil
}
