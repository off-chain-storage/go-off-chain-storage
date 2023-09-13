package utils

import (
	"embed"
	"time"

	log "github.com/sirupsen/logrus"
)

var abiJSON embed.FS

func Log_init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: time.StampMilli,
	})
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ReadABIFile(path string) []byte {
	data, err := abiJSON.ReadFile(path)
	CheckErr(err)

	return data
}
