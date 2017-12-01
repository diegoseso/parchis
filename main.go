package main

import (
	"github.com/diegoseso/parchis/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	server := server.NewServer()

	c := make(chan os.Signal, 1)

	signal.Notify(
		c,
		os.Kill,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	//serve until signal
	go func() {
		<-c
		server.Stop()
	}()

	server.Run()
}
