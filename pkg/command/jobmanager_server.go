package command

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/obaraelijah/secureproc/server/serverv1"
	v1 "github.com/obaraelijah/secureproc/service/v1"
	"github.com/obaraelijah/secureproc/util/grpcutil"

	"google.golang.org/grpc"
)

// RunJobmanagerServer runs a JobmanagerServer on the given network, listening
// on the given address, with the given CA certificate and server certificate
// and key.
func RunJobmanagerServer(network, address, caCert, serverCert, serverKey string) error {
	tc, err := grpcutil.NewServerTransportCredentials(caCert, serverCert, serverKey)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(tc),
		grpc.UnaryInterceptor(grpcutil.UnaryGetUserIDFromContextInterceptor),
		grpc.StreamInterceptor(grpcutil.StreamGetUserIDFromContextInterceptor),
	)

	v1.RegisterJobManagerServer(grpcServer, serverv1.NewJobmanagerServer())

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

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	grpcServer.GracefulStop()

	return nil
}
