package worker

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const loggerPrefix = "[WORKER]"
const sleepVarName = "WORKER_SLEEP_SECONDS"

var random *rand.Rand

func init() {
	s1 := rand.NewSource(time.Now().UnixNano())
	random = rand.New(s1)
}

func info(baseFormat string, args ...interface{}) {
	format := fmt.Sprintf("%s %s \n", loggerPrefix, baseFormat)
	fmt.Fprintf(os.Stdout, format, args...)
}

// InitWorker -
func InitWorker() error {
	info("Starting ...")
	workerErr := make(chan error)
	defer close(workerErr)

	sleepTime, err := strconv.Atoi(os.Getenv(sleepVarName))
	if err != nil {
		return errors.Wrapf(err, "failed to parsed '%s' ", sleepVarName)
	}

	info("Will run every %d seconds", sleepTime)

	go func() {
		for {
			time.Sleep(time.Second * time.Duration(sleepTime))
			randomNumber := random.Intn(100)
			if randomNumber > 90 {
				workerErr <- fmt.Errorf("random number is '%d' so above 90", randomNumber)
				continue
			}
			info("tick %d", randomNumber)
		}
	}()

	return <-workerErr
}
