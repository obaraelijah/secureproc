package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/obaraelijah/secureproc/pkg/command"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	err := command.RunJobmanagerServer(
		"tcp",
		":24482",
		"certs/ca.cert.pem",
		"certs/server.cert.pem",
		"certs/server.key.pem",
		stop,
	)
	if err != nil {
		panic(err)
	}
}
