package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/spf13/viper"
)

const (
	PORT string = "18411"
)

func main() {
	c := make(chan os.Signal, 1)

	w := Watcher{
		pipeName: "dalsk",
		filePath: "/tmp/dalsk.com",
		writer:   NewEndpointWriter(nil),
	}

	w.Watch()

	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL)

	<-c

	log.Println("Exiting App")

}
