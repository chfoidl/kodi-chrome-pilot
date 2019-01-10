package main

import (
	"log"
	"os/exec"
	"time"
)

func prepare() (err error) {
	stopKodi := exec.Command("sh", "-c", "systemctl is-active kodi && systemctl stop kodi")
	stopPA := exec.Command("sh", "-c", "systemctl is-active pulseaudio && systemctl stop pulseaudio")

	log.Print("Stopping kodi")

	err = stopKodi.Run()
	if err != nil {
		return err
	}

	log.Print("Stopping pulse audio")

	err = stopPA.Run()
	if err != nil {
		return err
	}

	return nil
}

func tearDown() (err error) {
	log.Print("Restoring system...")

	startKodi := exec.Command("systemctl", "restart", "kodi")
	startPA := exec.Command("systemctl", "start", "pulseaudio")

	err = startKodi.Start()
	if err != nil {
		return err
	}

	time.Sleep(5 * time.Second)

	err = startPA.Start()
	if err != nil {
		return err
	}

	killBrowser()

	return nil
}
