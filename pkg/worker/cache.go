package worker

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httputil"
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

	cachePath := storage.SlugCachePath(item.ID)
	body, err := httputil.DumpResponse(resp, isBodyCached)
	if err != nil {
		return errors.Wrap(err, "failed dumping reponse")
	}
	// @todo rewrite body
	err = ioutil.WriteFile(cachePath, body, 0644)
	if err != nil {
		return errors.Wrap(err, "failed writing to cache")
	}

	loggr.Info("Refreshed '%s'", item.ID)

	return nil
}
