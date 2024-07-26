package main

import (
	"os"
	"os/signal"
	"syscall"
)

func handleInterupt() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	<-sc
	// This channel will signal the writter GoR to write data from the in-mem cache to persistent storage
	triggerChannel <- true
}
