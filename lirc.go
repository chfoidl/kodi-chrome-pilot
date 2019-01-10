package main

import (
	"log"
	"time"

	"github.com/chrisxf/lirc"
	"github.com/micmonay/keybd_event"
)

var kb keybd_event.KeyBonding

func pressKey(key int) {
	kb.Clear()
	kb.SetKeys(key)

	err := kb.Launching()
	if err != nil {
		log.Fatal("Could not send keyboard key ", key)
	}
}

func keyDown(event lirc.Event) {
	log.Println("Incoming DOWN -> Simulate DOWN keystroke")
	pressKey(keybd_event.VK_DOWN)
}

func keyUp(event lirc.Event) {
	log.Println("Incoming UP -> Simulate UP keystroke")
	pressKey(keybd_event.VK_UP)
}

func keyLeft(event lirc.Event) {
	log.Println("Incoming LEFT -> Simulate LEFT keystroke")
	pressKey(keybd_event.VK_LEFT)
}

func keyRight(event lirc.Event) {
	log.Println("Incoming RIGHT -> Simulate RIGHT keystroke")
	pressKey(keybd_event.VK_RIGHT)
}

func keyOK(event lirc.Event) {
	log.Println("Incoming OK -> Simulate ENTER keystroke")
	pressKey(keybd_event.VK_ENTER)
}

func keyExit(event lirc.Event) {
	log.Println("Incoming EXIT -> Simulate ESC keystroke")
	pressKey(keybd_event.VK_ESC)
}

func keyRed(event lirc.Event) {
	log.Println("Incoming RED -> Simulate HOME keystroke")
	pressKey(keybd_event.VK_HOME)
}

func keyGreen(event lirc.Event) {
	log.Println("Incoming GREEN -> Simulate END keystroke")
	pressKey(keybd_event.VK_END)
}

func keyNumeric0(event lirc.Event) {
	log.Println("Incoming NUMERIC_0 -> Killing chrome")

	ProcChrome.Process.Kill()
}

func handleIRC() (err error) {
	kb, err = keybd_event.NewKeyBonding()
	if err != nil {
		return err
	}

	time.Sleep(2 * time.Second)

	ir, err := lirc.Init("/var/run/lirc/lircd")
	if err != nil {
		return err
	}

	log.Println("Ready to accept IR...")

	ir.Handle("", "KEY_UP", keyUp)
	ir.Handle("", "KEY_DOWN", keyDown)
	ir.Handle("", "KEY_LEFT", keyLeft)
	ir.Handle("", "KEY_RIGHT", keyRight)
	ir.Handle("", "KEY_OK", keyOK)
	ir.Handle("", "KEY_EXIT", keyExit)
	ir.Handle("", "KEY_RED", keyRed)
	ir.Handle("", "KEY_GREEN", keyGreen)
	ir.Handle("", "KEY_NUMERIC_0", keyNumeric0)

	ir.Run()

	return nil
}

func startIRCHandler() {
	err := handleIRC()
	if err != nil {
		log.Println(err)
	}
}
