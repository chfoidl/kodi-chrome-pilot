package main

import (
	"log"
	"os"
)

func setupLog() {
	if LogFile == "" {
		return
	}

	f, err := os.OpenFile(LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}

	defer f.Close()

	log.SetOutput(f)
}
