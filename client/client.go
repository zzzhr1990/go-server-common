package client

import (
	"os"
	"os/signal"
	"syscall"
	// log "github.com/sirupsen/logrus"
)

// RunAndHold hold an client
func RunAndHold(f func(text string)) {
	forever := make(chan bool)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		if f != nil {
			f("hold program & waiting SIGTERM...")
		}
		<-c
		if f != nil {
			f("starting shutown...")
		}
		forever <- true
	}()
	<-forever
}
