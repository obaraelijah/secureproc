//go:build integration
// +build integration

package server_test

import (
	"context"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/obaraelijah/secureproc/certs"
	"github.com/obaraelijah/secureproc/pkg/command"
	"github.com/obaraelijah/secureproc/service/jobmanager/jobmanagerv1"
	"github.com/obaraelijah/secureproc/util/grpcutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

// Note: There is some fragility with these tests.
//       * This expects port 12345 to be available (so we wouldn't want to
//         run multiple instances of this test on the same host in parallel)
//       * There's a Sleep() between starting the server and trying to connect
//         to it.  That's open to race conditions.  I'd like to find a way to
//         know that the server is up before we try to connect to it.

func Test_clientServer_clientCertNotSignedByTrustedCA(t *testing.T) {
	stop := make(chan os.Signal, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		_ = command.RunJobmanagerServer(
			"tcp",
			":12345",
			certs.CACert,
			certs.ServerCert,
			certs.ServerKey,
			stop)
		wg.Done()
	}()

	waitForServer()

	tc, err := grpcutil.NewClientTransportCredentials(
		certs.CACert,
		certs.BadClientCert,
		certs.BadClientKey,
	)
	require.Nil(t, err)

	conn, err := grpc.Dial("localhost:12345", grpc.WithTransportCredentials(tc))
	require.Nil(t, err)
	defer conn.Close()

	client := jobmanagerv1.NewJobManagerClient(conn)
	_, err = client.List(context.Background(), &jobmanagerv1.NilMessage{})

	assert.Error(t, err)

	stop <- os.Kill
	wg.Wait()
}

func Test_clientServer_serverCertNotSignedByTrustedCA(t *testing.T) {
	stop := make(chan os.Signal, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		_ = command.RunJobmanagerServer(
			"tcp",
			":12345",
			certs.CACert,
			certs.BadServerCert,
			certs.BadServerKey,
			stop)
		wg.Done()
	}()

	waitForServer()

	tc, err := grpcutil.NewClientTransportCredentials(
		certs.CACert,
		certs.Client1Cert,
		certs.Client1Key,
	)
	require.Nil(t, err)

	conn, err := grpc.Dial("localhost:12345", grpc.WithTransportCredentials(tc))
	require.Nil(t, err)
	defer conn.Close()

	client := jobmanagerv1.NewJobManagerClient(conn)

	ctx := grpcutil.AttachUserIDToContext(context.Background(), "user1")

	_, err = client.List(ctx, &jobmanagerv1.NilMessage{})

	fmt.Println(err)
	assert.Error(t, err)

	stop <- os.Kill
	wg.Wait()
}

func Test_clientServer_TooWeakServerCert(t *testing.T) {
	stop := make(chan os.Signal, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		_ = command.RunJobmanagerServer(
			"tcp",
			":12345",
			certs.CACert,
			certs.WeakServerCert,
			certs.WeakServerKey,
			stop)
		wg.Done()
	}()

	waitForServer()

	tc, err := grpcutil.NewClientTransportCredentials(
		certs.CACert,
		certs.Client1Cert,
		certs.Client1Key,
	)
	require.Nil(t, err)

	conn, err := grpc.Dial("localhost:12345", grpc.WithTransportCredentials(tc))
	require.Nil(t, err)
	defer conn.Close()

	client := jobmanagerv1.NewJobManagerClient(conn)

	ctx := grpcutil.AttachUserIDToContext(context.Background(), "user1")

	_, err = client.List(ctx, &jobmanagerv1.NilMessage{})

	fmt.Println(err)
	assert.Error(t, err)

	stop <- os.Kill
	wg.Wait()
}

func Test_clientServer_TooWeakClientCert(t *testing.T) {
	stop := make(chan os.Signal, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		_ = command.RunJobmanagerServer(
			"tcp",
			":12345",
			certs.CACert,
			certs.ServerCert,
			certs.ServerKey,
			stop)
		wg.Done()
	}()

	waitForServer()

	tc, err := grpcutil.NewClientTransportCredentials(
		certs.CACert,
		certs.WeakClientCert,
		certs.WeakClientKey,
	)
	require.Nil(t, err)

	conn, err := grpc.Dial("localhost:12345", grpc.WithTransportCredentials(tc))
	require.Nil(t, err)
	defer conn.Close()

	client := jobmanagerv1.NewJobManagerClient(conn)

	ctx := grpcutil.AttachUserIDToContext(context.Background(), "weakclient")

	_, err = client.List(ctx, &jobmanagerv1.NilMessage{})

	fmt.Println(err)
	assert.Error(t, err)

	stop <- os.Kill
	wg.Wait()
}

func Test_clientServer_Success(t *testing.T) {
	stop := make(chan os.Signal, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		_ = command.RunJobmanagerServer(
			"tcp",
			":12345",
			certs.CACert,
			certs.ServerCert,
			certs.ServerKey,
			stop)
		wg.Done()
	}()

	waitForServer()

	tc, err := grpcutil.NewClientTransportCredentials(
		certs.CACert,
		certs.Client1Cert,
		certs.Client1Key,
	)
	require.Nil(t, err)

	conn, err := grpc.Dial("localhost:12345", grpc.WithTransportCredentials(tc))
	require.Nil(t, err)
	defer conn.Close()

	client := jobmanagerv1.NewJobManagerClient(conn)

	ctx := grpcutil.AttachUserIDToContext(context.Background(), "user1")

	_, err = client.List(ctx, &jobmanagerv1.NilMessage{})

	assert.Nil(t, err)

	stop <- os.Kill
	wg.Wait()
}

// waitForServer waits for the server to come up before attempting a network
// connection to it.
//
// This is gross, but I didn't find a better way to check.
func waitForServer() {
	time.Sleep(1 * time.Second)
}
