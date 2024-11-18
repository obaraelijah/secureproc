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
	network, address string,
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

	go func() {
		l, err := net.Listen(network, address)
		if err != nil {
			panic(err)
		}

		log.Println("Server ready.")
		if err := grpcServer.Serve(l); err != nil {
			panic(err)
		}
	}()

	<-ctx.Done()
	grpcServer.GracefulStop()

	return nil
}
