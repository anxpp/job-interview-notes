package main

import (
	"github.com/google/gops/agent"
	"log"
	"time"
)

func main() {
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Hour)
}
