package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var ProcChrome *exec.Cmd

func startBrowser() (err error) {
	log.Print("Starting chrome...")

	args := ChromeArgs

	for _, ext := range ChromeExts {
		if _, err := os.Stat(ext); err == nil {
			log.Print("Adding chrome extension ", ext)

			args = append(args, "--load-extension="+ext)
		} else {
			log.Print("Could not add chrome extension ", ext, ": ", err)
		}
	}

	args = append(args, "--kiosk")
	args = append(args, BrowserUrl)

	log.Println(args)

	ProcChrome = exec.Command(ChromePath, args...)

	err = ProcChrome.Run()
	if err != nil {
		fmt.Println("Error running chrome", err)
		return err
	}

	return nil
}

func killBrowser() {
	cmd := exec.Command("sh", "-c", "pgrep chrome | xargs kill")
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(out)
}
