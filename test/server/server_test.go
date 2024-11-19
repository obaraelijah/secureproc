//go:build integration
// +build integration

package server_test

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"testing"

	"github.com/obaraelijah/secureproc/certs"
	"github.com/obaraelijah/secureproc/pkg/command"
	"github.com/obaraelijah/secureproc/server/jobmanager/serverv1"
	"github.com/obaraelijah/secureproc/service/jobmanager/jobmanagerv1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func Test_clientServer_clientCertNotSignedByTrustedCA(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	port, err := runServer(ctx, &wg, t, certs.CACert, certs.ServerCert, certs.ServerKey)
	require.Nil(t, err)

	tc, err := certs.NewClientTransportCredentials(
		certs.CACert,
		certs.BadClientCert,
		certs.BadClientKey,
	)
	require.Nil(t, err)

	conn, err := grpc.Dial("localhost:"+port, grpc.WithTransportCredentials(tc))
	require.Nil(t, err)
	defer conn.Close()

	client := jobmanagerv1.NewJobManagerClient(conn)
	_, err = client.List(context.Background(), &jobmanagerv1.NilMessage{})

	assert.Error(t, err)

	cancel()
	wg.Wait()
}

func Test_clientServer_serverCertNotSignedByTrustedCA(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	port, err := runServer(ctx, &wg, t, certs.CACert, certs.BadServerCert, certs.BadServerKey)
	require.Nil(t, err)

	tc, err := certs.NewClientTransportCredentials(
		certs.CACert,
		certs.Client1Cert,
		certs.Client1Key,
	)
	require.Nil(t, err)

	conn, err := grpc.Dial("localhost:"+port, grpc.WithTransportCredentials(tc))
	require.Nil(t, err)
	defer conn.Close()

	client := jobmanagerv1.NewJobManagerClient(conn)

	opCtx := serverv1.AttachUserIDToContext(context.Background(), "user1")

	_, err = client.List(opCtx, &jobmanagerv1.NilMessage{})

	fmt.Println(err)
	assert.Error(t, err)

	cancel()
	wg.Wait()
}

func Test_clientServer_TooWeakServerCert(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	port, err := runServer(ctx, &wg, t, certs.CACert, certs.WeakServerCert, certs.WeakServerKey)
	require.Nil(t, err)

	tc, err := certs.NewClientTransportCredentials(
		certs.CACert,
		certs.Client1Cert,
		certs.Client1Key,
	)
	require.Nil(t, err)

	conn, err := grpc.Dial("localhost:"+port, grpc.WithTransportCredentials(tc))
	require.Nil(t, err)
	defer conn.Close()

	client := jobmanagerv1.NewJobManagerClient(conn)

	opCtx := serverv1.AttachUserIDToContext(context.Background(), "user1")

	_, err = client.List(opCtx, &jobmanagerv1.NilMessage{})

	fmt.Println(err)
	assert.Error(t, err)

	cancel()
	wg.Wait()
}

func Test_clientServer_TooWeakClientCert(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	port, err := runServer(ctx, &wg, t, certs.CACert, certs.ServerCert, certs.ServerKey)
	require.Nil(t, err)

	tc, err := certs.NewClientTransportCredentials(
		certs.CACert,
		certs.WeakClientCert,
		certs.WeakClientKey,
	)
	require.Nil(t, err)

	conn, err := grpc.Dial("localhost:"+port, grpc.WithTransportCredentials(tc))
	require.Nil(t, err)
	defer conn.Close()

	client := jobmanagerv1.NewJobManagerClient(conn)

	opCtx := serverv1.AttachUserIDToContext(context.Background(), "weakclient")

	_, err = client.List(opCtx, &jobmanagerv1.NilMessage{})

	fmt.Println(err)
	assert.Error(t, err)

	cancel()
	wg.Wait()
}

func Test_clientServer_Success(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	port, err := runServer(ctx, &wg, t, certs.CACert, certs.ServerCert, certs.ServerKey)
	require.Nil(t, err)

	tc, err := certs.NewClientTransportCredentials(
		certs.CACert,
		certs.Client1Cert,
		certs.Client1Key,
	)
	require.Nil(t, err)

	conn, err := grpc.Dial("localhost:"+port, grpc.WithTransportCredentials(tc))
	require.Nil(t, err)
	defer conn.Close()

	client := jobmanagerv1.NewJobManagerClient(conn)

	opCtx := serverv1.AttachUserIDToContext(context.Background(), "user1")

	_, err = client.List(opCtx, &jobmanagerv1.NilMessage{})

	assert.Nil(t, err)

	cancel()
	wg.Wait()
}

func getPort(address string) (string, error) {
	tokens := strings.Split(address, ":")
	if len(tokens) == 0 {
		return "", fmt.Errorf("failed to find port in address '%s'", address)
	}

	return tokens[len(tokens)-1], nil
}

func runServer(
	ctx context.Context,
	wg *sync.WaitGroup,
	t *testing.T,
	caCert, serverCert, serverKey []byte,
) (port string, err error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return "", err
	}

	go func() {
		runErr := command.RunJobmanagerServer(
			ctx, listener, caCert, serverCert, serverKey)
		assert.Nil(t, runErr)
		wg.Done()
	}()

	return getPort(listener.Addr().String())
}
