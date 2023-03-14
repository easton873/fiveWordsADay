package log

import (
	"io"
	"log"
)

func LogError(msg string, err error) {
	log.Printf("%s: %v\n", msg, err)
}

func CloseAndLog(closer io.Closer, msg string) {
	err := closer.Close()
	if err != nil {
		LogError(msg, err)
	}
}
