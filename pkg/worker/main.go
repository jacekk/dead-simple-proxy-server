package worker

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const prefix = "[WORKER]"

func info(format string, args ...interface{}) {
	prefixedFormat := fmt.Sprintf("%s %s \n", prefix, format)
	fmt.Fprintf(os.Stdout, prefixedFormat, args...)
}

// Init -
func Init() error {
	info("Starting ...")

	go func() {
		for {
			time.Sleep(time.Second * 10)
			info("tick")
		}
	}()

	info("Press 'Enter' to exit ...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	return nil
}
