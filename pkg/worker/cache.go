package worker

import (
	"math/rand"
	"time"

	"github.com/jacekk/dead-simple-proxy-server/pkg/storage"
	"github.com/pkg/errors"
)

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
		return errors.Wrapf(err, "failed refreshing '%s' item", randomItem.URL)
	}

	return nil
}

func refreshConfigItem(loggr *loggerImpl, item storage.Item) error {
	loggr.Info("Refreshing '%s'", item.ID)

	return nil
}
