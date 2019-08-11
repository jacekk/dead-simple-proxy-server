package storage

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
)

// GetBodyRewriteSource - returns key of bodyRewrite based on given entry value
func GetBodyRewriteSource(id string) (string, error) {
	config, err := readConfig()
	if err != nil {
		log.Printf("%+v", errors.Wrap(err, "failed to read and parse proxy config"))
		return "", err
	}

	for _, configItem := range config {
		for key, value := range configItem.BodyRewrite {
			if strings.Contains(value, id) {
				return key, nil
			}
		}
	}

	errMsg := fmt.Sprintf("failed to find proxy rewrite for %q id", id)
	log.Print(errMsg)

	return "", errors.New(errMsg)
}
