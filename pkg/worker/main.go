package worker

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const loggerPrefix = "[WORKER]"

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

	go func() {
		for {
			time.Sleep(time.Second * 5)
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
