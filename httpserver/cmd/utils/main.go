package utils

import (
	"log"
	"os"
)

func Fatal(err error) {
	log.Fatal(err)
	os.Exit(1)
}
