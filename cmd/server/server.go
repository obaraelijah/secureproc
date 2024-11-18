package main

import (
	"context"
	"net"
	"os/signal"
	"syscall"

	"github.com/obaraelijah/secureproc/certs"
	"github.com/obaraelijah/secureproc/pkg/command"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	listener, err := net.Listen("tcp", ":24482")
	if err != nil {
		panic(err)
	}

	err = command.RunJobmanagerServer(
		ctx,
		listener,
		certs.CACert,
		certs.ServerCert,
		certs.ServerKey,
	)

	if err != nil {
		panic(err)
	}
}
