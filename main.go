package main

import (
	"log"
)

func main() {
	parseFlags()
	setupLog()

	log.Print("Starting Kodi Chrome-TV Pilot")

	checkFlags()
	prepare()

	go startIRCHandler()

	if err := startBrowser(); err != nil {
		log.Println("FATAL: ", err)
	}

	tearDown()
}
