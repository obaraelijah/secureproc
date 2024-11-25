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
	RunE: func(cmd *cobra.Command, args []string) error {
		return runServer(cmd.Context())
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

func runServer(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	listener, err := net.Listen("tcp", argAddress)
	if err != nil {
		return err
	}
	defer listener.Close()

	err = command.RunJobmanagerServer(
		ctx,
		listener,
		certs.CACert,
		certs.ServerCert,
		certs.ServerKey,
	)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	rootCmd.Execute()
}
