package main

import (
	"flag"
	"log"
	"thinblue/internal/core"
)

func main() {

	c := Config{}
	err := c.InitialConfig()
	if err != nil {
		log.Fatalf("Error initializing configuration: %v", err)
	}

	flag.Parse()

	core.SetupLogger(c.debug_mode)

}
