package worker

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/jacekk/dead-simple-proxy-server/pkg/storage"
	"github.com/pkg/errors"
)

const isBodyCached = true

var random *rand.Rand

func init() {
	src := rand.NewSource(time.Now().UnixNano())
	random = rand.New(src)
}

func refreshRandomCache(loggr *loggerImpl) error {
	cached, err := storage.GetCached()
	if err != nil {
		return errors.Wrap(err, "failed reading cached items")
	}
	if len(cached) == 0 {
		return nil
	}

	randomIndex := random.Intn(len(cached))
	randomItem := cached[randomIndex]
	err = refreshConfigItem(loggr, randomItem)
	if err != nil {
		return errors.Wrapf(err, "failed refreshing '%s' ", randomItem.ID)
	}

	return nil
}

func refreshConfigItem(loggr *loggerImpl, item storage.Item) error {
	loggr.Info("Refreshing '%s' ...", item.ID)
	resp, err := http.Get(item.URL)
	if err != nil {
		return errors.Wrap(err, "failed making get request")
	}
	defer resp.Body.Close()

	bodyPath := storage.SlugCachePath(item.ID, "body")
	headersPath := storage.SlugCachePath(item.ID, "headers")

	// cache body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed reading body")
	}
	for from, to := range item.BodyRewrite {
		body = bytes.Replace(body, []byte(from), []byte(to), -1)
	}
	err = ioutil.WriteFile(bodyPath, body, 0644)
	if err != nil {
		return errors.Wrap(err, "failed writing body to cache")
	}

	// cache headers
	resp.Header.Set("X-proxy-cached-at", time.Now().Format(http.TimeFormat))
	headers, err := json.Marshal(resp.Header)
	if err != nil {
		return errors.Wrap(err, "failed to serialize headers")
	}
	err = ioutil.WriteFile(headersPath, headers, 0644)
	if err != nil {
		return errors.Wrap(err, "failed writing headers to cache")
	}

	loggr.Info("Refreshed '%s'", item.ID)

	return nil
}
