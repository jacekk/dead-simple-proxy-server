package storage

import (
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
func SlugCachePath(slug string) string {
	baseDir := helpers.GetProjectDir()

	return path.Join(baseDir, "cache", slug+".txt")
}
