package main

import (
	"flag"
	"fmt"
	"log"
	"thinblue/api"
	"thinblue/internal/core"
	"thinblue/internal/ingest"
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

	go api.StartSysLogUDPServer(c.port)
	parser := ingest.NewSyslogParser()
	i := []byte(`<165>1 2025-06-22T20:03:13.123456+01:00 myhost.example.com myapp AUTH-AUDIT 12345 [audit@18060 op="login" result="success" user="john.doe" ip="192.168.1.100" sessionID="ABCDEF123456" clientType="web-browser"][rule_engine@32473 ruleId="LoginSuccess" confidence="0.9" tags="authentication,behavior"] \xEF\xBB\xBFUser 'john.doe' from 192.168.1.100 successfully logged in to MyApp via web browser. Session ID: ABCDEF123456.`)
	_, err = parser.BuildLogObject(i)
	if err != nil {
		log.Panicln("Failed to parse syslog message: \n", err)
	}

}
