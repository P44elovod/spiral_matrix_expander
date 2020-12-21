package util

import (
	"log"
)

//FailOnError is a wraper for logging errors
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
