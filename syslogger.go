package main

import (
	"log"
	"log/syslog"
)

func main() {
	s, err := syslog.Dial("tcp", ":5000", syslog.LOG_INFO, "test")
	if err != nil {
		log.Fatal(err)
	}

	s.Info("pants")
}
