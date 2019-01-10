package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

type chromeExts []string

func (i *chromeExts) String() string {
	return strings.Join(*i, " ")
}

func (i *chromeExts) Set(value string) error {
	*i = append(*i, value)

	return nil
}

func (i *chromeExts) Value() []string {
	return *i
}

var BrowserUrl string
var ChromePath string
var ChromeExts chromeExts
var LogFile string

var ChromeArgs []string

func parseFlags() {
	f := flag.NewFlagSet("default", flag.ExitOnError)

	f.Var(&ChromeExts, "ext-path", "path to unpacked chrome extension directory")
	f.StringVar(&LogFile, "log-path", "", "path to log directory")
	f.StringVar(&ChromePath, "chrome-path", "", "path to chrome executable")
	f.StringVar(&BrowserUrl, "url", "", "url to load")

	mainArgs, cArgs := getSplitArgs()
	ChromeArgs = cArgs

	f.Parse(mainArgs)
}

func getSplitArgs() (mainArgs []string, chromeArgs []string) {
	index := -1
	for i, arg := range os.Args {
		if arg == "--chrome-args" {
			index = i
		}
	}

	if index == -1 {
		return os.Args[1:], chromeArgs
	} else {
		return os.Args[1:index], os.Args[index+1:]
	}
}

func checkFlags() {
	if ChromePath == "" {
		log.Fatal("Chrome path must be set!")
	}

	if BrowserUrl == "" {
		log.Fatal("URL must be set!")
	}
}
