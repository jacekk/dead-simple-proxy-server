package storage

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

// GetProxyConfigBySlug - returns URL based on given slug
func GetProxyConfigBySlug(slug string) (Item, error) {
	config, err := readConfig()
	if err != nil {
		log.Printf("%+v", errors.Wrap(err, "failed to read and parse proxy config"))
		return Item{}, err
	}

	for _, item := range config {
		if item.ID == slug {
			return item, nil
		}
	}

	errMsg := fmt.Sprintf("failed to find proxy config for %q slug", slug)
	log.Print(errMsg)

	return Item{}, errors.New(errMsg)
}
