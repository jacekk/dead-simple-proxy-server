package worker

import (
	"bufio"
	"log"
	"os"
	"time"
)

const prefix = "[WORKER]"

// Init -
func Init() error {
	log.Printf("%s Starting ...", prefix)

	go func() {
		for {
			time.Sleep(time.Second * 10)
			log.Printf("%s tick ", prefix)
		}
	}()

	log.Printf("%s Press 'Enter' to exit ...", prefix)
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	return nil
}
