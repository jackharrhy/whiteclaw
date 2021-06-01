package main

import (
	"io"
	"log"
	"time"

	"github.com/gliderlabs/ssh"
)

var sayings = []string{
	"its the weekend baby",
	"FUCK",
	"its friday baby",
	"FUCK",
	"drinking a whiteclaw baby",
	"FUCK",
	"i deserve a truffle pizza",
	"NO I DONT",
	"BECAUSE ITS FUCKING SUSHI NIGHT",
	"FUCK",
	"ITS THE WEEKEND BABY",
	"FUCK",
}

func main() {
	ssh.Handle(func(s ssh.Session) {
		for _, saying := range sayings {
			time.Sleep(800 * time.Millisecond)
			io.WriteString(s, saying+"\n")
		}
	})

	log.Println("starting ssh server on port 2222...")
	log.Fatal(ssh.ListenAndServe(":2222", nil))
}
