package storage

import (
	"fmt"
	"path"

	"github.com/jacekk/dead-simple-proxy-server/pkg/helpers"
	"github.com/pkg/errors"
)

// GetCached -
func GetCached() (Items, error) {
	cached := make(Items, 0)
	config, err := readConfig()
	if err != nil {
		return cached, errors.Wrap(err, "failed to read and parse proxy config")
	}

	for _, item := range config {
		if item.IsCacheEnabled {
			cached = append(cached, item)
		}
	}

	return cached, nil
}

// SlugCachePath -
func SlugCachePath(slug, fileType string) string {
	baseDir := helpers.GetProjectDir()
	fileName := fmt.Sprintf("%s.%s.txt", slug, fileType)

	return path.Join(baseDir, "cache", fileName)
}
