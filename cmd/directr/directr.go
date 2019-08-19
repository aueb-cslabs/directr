package main

import (
	"flag"
	"log"
	"os"

	"github.com/enderian/directrd/pkg/types"
	"github.com/kardianos/service"
)

var svcConfig = &service.Config{
	Name:        "directr",
	DisplayName: "directr",
	Description: "The agent process that communicates with directrd.",
}

var hostname string
var config *types.Configuration
var logger service.Logger

var (
	configFlag = flag.String("config", "config.yml", "Configuration file to be used")
)

func main() {
	flag.Parse()

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatalf("%s", err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatalf("%s", err)
	}

	if conf, err := types.LoadConfiguration(*configFlag); err == nil {
		config = conf
	} else {
		log.Fatalf("%s", err)
	}

	if len(os.Args) < 2 {
		log.Println("starting directr in interactive mode")
		err = s.Run()
	} else {
		err = service.Control(s, os.Args[1])
	}
	if err != nil {
		log.Fatalf("%s", err)
	}
}
