package command

import (
	"context"
	"log"
	"net"

	"github.com/obaraelijah/secureproc/server/jobmanager/serverv1"
	"github.com/obaraelijah/secureproc/service/jobmanager/jobmanagerv1"
	"github.com/obaraelijah/secureproc/util/grpcutil"
	"google.golang.org/grpc"
)

// RunJobmanagerServer runs a JobmanagerServer on the given network, listening
// on the given address, with the given CA certificate and server certificate
// and key.
func RunJobmanagerServer(
	ctx context.Context,
	listener net.Listener,
	caCert, serverCert, serverKey []byte,
) error {
	tc, err := grpcutil.NewServerTransportCredentials(caCert, serverCert, serverKey)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(tc),
		grpc.UnaryInterceptor(grpcutil.UnaryGetUserIDFromContextInterceptor),
		grpc.StreamInterceptor(grpcutil.StreamGetUserIDFromContextInterceptor),
	)

	jobmanagerv1.RegisterJobManagerServer(grpcServer, serverv1.NewJobmanagerServer())

	errChan := make(chan error)

	go func() {
		log.Println("Server ready.")
		if err := grpcServer.Serve(listener); err != nil {
			errChan <- err
			return
		}
	}()

	select {
	case <-ctx.Done():
	case err = <-errChan:
	}

	grpcServer.GracefulStop()

	return err
}
