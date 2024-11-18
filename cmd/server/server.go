package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/obaraelijah/secureproc/certs"
	"github.com/obaraelijah/secureproc/pkg/command"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	err := command.RunJobmanagerServer(
		ctx,
		"tcp",
		":24482",
		certs.CACert,
		certs.ServerCert,
		certs.ServerKey,
	)
	if err != nil {
		panic(err)
	}
}
