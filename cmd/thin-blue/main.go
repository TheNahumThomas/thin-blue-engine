package main

import (
	"flag"
	"fmt"
	"log"
	"thinblue/internal/core"
)

func main() {

	c := Config{}
	err := c.InitialConfig()
	if err != nil {
		fmt.Println("Error initializing configuration:", err)
	}

	flag.Parse()

	err = core.SetupLogger(c.debug_mode)
	if err != nil {
		log.Panicln("Failed to setup logger: \n", err)
	}

	log.Println("Engine logging started with debug value:", c.debug_mode)

}
