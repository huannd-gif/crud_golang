package main

import (
	"api_crud/port"
	"flag"
)

var (
	runType = flag.String("runtype", "http", "server type")
)

func main() {
	flag.Parse()

	switch *runType {
	case "http":
		port.HttpApp()
	case "call_create":
		port.RabbitMQApp()
	}

}
