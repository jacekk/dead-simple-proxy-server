package worker

import (
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const loggerPrefix = "[WORKER]"
const sleepVarName = "WORKER_SLEEP_SECONDS"

// InitWorker -
func InitWorker() error {
	loggr := new(loggerImpl)
	loggr.Info("Starting ...")
	workerErr := make(chan error)
	defer close(workerErr)

	sleepTime, err := strconv.Atoi(os.Getenv(sleepVarName))
	if err != nil {
		return errors.Wrapf(err, "failed to parsed '%s' ", sleepVarName)
	}

	loggr.Info("Will run every %d seconds", sleepTime)

	go func() {
		for {
			time.Sleep(time.Second * time.Duration(sleepTime))
			err := refreshRandomCache(loggr)
			if err != nil {
				workerErr <- err
			}
		}
	}()

	return <-workerErr
}
