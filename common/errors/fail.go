package errors

import "log"

func FailOnError(err error, info string) {
	if err != nil {
		log.Fatalf("%s: %v\n", info, err)
	}
}
