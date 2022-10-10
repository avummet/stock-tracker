package common

import (
	log "github.com/sirupsen/logrus"
)

func FatalIfErr(err error) {
	if err != nil {
		log.WithError(err).Fatal("Something failed")
	}
}
