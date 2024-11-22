package main

import (
	"context"
	"net"
	"os/signal"
	"syscall"

	"github.com/obaraelijah/secureproc/certs"
	"github.com/obaraelijah/secureproc/pkg/command"
	"github.com/spf13/cobra"
)

var (
	argAddress string
)

var rootCmd = &cobra.Command{
	Use:   "jobmanager",
	Short: "Run the job manager server",
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&argAddress,
		"address",
		"a",
		":24482",
		"The <address>:<port> on which this server should listen for incoming requests")
}

func runServer() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	listener, err := net.Listen("tcp", argAddress)
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

func main() {
	rootCmd.Execute()
}
