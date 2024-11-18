package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/obaraelijah/secureproc/certs"
	"github.com/obaraelijah/secureproc/pkg/command"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	err := command.RunJobmanagerServer(
		"tcp",
		":24482",
		certs.CACert,
		certs.ServerCert,
		certs.ServerKey,
		stop,
	)
	if err != nil {
		panic(err)
	}
}
