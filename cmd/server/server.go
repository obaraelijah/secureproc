package main

import "github.com/obaraelijah/secureproc/pkg/command"

func main() {
	err := command.RunJobmanagerServer(
		"tcp",
		":24482",
		"certs/ca.cert.pem",
		"certs/server.cert.pem",
		"certs/server.key.pem",
	)
	if err != nil {
		panic(err)
	}
}
