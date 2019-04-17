package client

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

//RunAndHold hold an client
func RunAndHold() {
	forever := make(chan bool)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		log.Infof("Waiting SIGTERM...")
		<-c
		log.Infof("Do clean jobs...")
		forever <- true
	}()
	<-forever
}
