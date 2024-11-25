//go:build integration
// +build integration

package server_test

import (
	"context"
	"net"
	"sync"
	"testing"

	"github.com/obaraelijah/secureproc/certs"
	"github.com/obaraelijah/secureproc/pkg/client/jobmanager"
	"github.com/obaraelijah/secureproc/pkg/command"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_clientServer_clientCertNotSignedByTrustedCA(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	hostPort, err := runServer(ctx, &wg, t, certs.CACert, certs.ServerCert, certs.ServerKey)
	require.Nil(t, err)

	client, err := jobmanager.NewClient("badclient", hostPort)
	require.Nil(t, err)
	defer client.Close()

	_, err = client.List(context.Background())

	assert.Error(t, err)
	s, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.Unavailable, s.Code())

	cancel()
	wg.Wait()
}

func Test_clientServer_serverCertNotSignedByTrustedCA(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	hostPort, err := runServer(ctx, &wg, t, certs.CACert, certs.BadServerCert, certs.BadServerKey)
	require.Nil(t, err)

	client, err := jobmanager.NewClient("client1", hostPort)
	require.Nil(t, err)
	defer client.Close()

	_, err = client.List(context.Background())

	assert.Error(t, err)
	s, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.Unavailable, s.Code())

	cancel()
	wg.Wait()
}

func Test_clientServer_TooWeakServerCert(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	hostPort, err := runServer(ctx, &wg, t, certs.CACert, certs.WeakServerCert, certs.WeakServerKey)
	require.Nil(t, err)

	client, err := jobmanager.NewClient("client1", hostPort)
	require.Nil(t, err)
	defer client.Close()

	_, err = client.List(context.Background())

	assert.Error(t, err)

	s, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.Unavailable, s.Code())

	cancel()
	wg.Wait()
}

func Test_clientServer_TooWeakClientCert(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	hostPort, err := runServer(ctx, &wg, t, certs.CACert, certs.ServerCert, certs.ServerKey)
	require.Nil(t, err)

	client, err := jobmanager.NewClient("weakclient", hostPort)
	require.Nil(t, err)
	defer client.Close()

	_, err = client.List(context.Background())

	assert.Error(t, err)
	s, ok := status.FromError(err)
	assert.True(t, ok)
	assert.Equal(t, codes.Unavailable, s.Code())

	cancel()
	wg.Wait()
}

func Test_clientServer_Success(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	hostPort, err := runServer(ctx, &wg, t, certs.CACert, certs.ServerCert, certs.ServerKey)
	require.Nil(t, err)

	client, err := jobmanager.NewClient("client1", hostPort)
	require.Nil(t, err)
	defer client.Close()

	_, err = client.List(context.Background())

	assert.Nil(t, err)

	cancel()
	wg.Wait()
}

func Test_clientServer_Multitenant(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	hostPort, err := runServer(ctx, &wg, t, certs.CACert, certs.ServerCert, certs.ServerKey)
	require.Nil(t, err)

	user1Client, err := jobmanager.NewClient("client1", hostPort)
	require.Nil(t, err)
	defer user1Client.Close()

	_, err = user1Client.Start(context.Background(), "myjob", "/bin/true")
	assert.Nil(t, err)

	user2Client, err := jobmanager.NewClient("client2", hostPort)
	require.Nil(t, err)
	defer user2Client.Close()

	jobList, err := user2Client.List(context.Background())

	assert.Nil(t, err)
	assert.Equal(t, 0, len(jobList))

	cancel()
	wg.Wait()
}

func Test_clientServer_AdministratorCanSeeAllJobs(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	hostPort, err := runServer(ctx, &wg, t, certs.CACert, certs.ServerCert, certs.ServerKey)
	require.Nil(t, err)

	user1Client, err := jobmanager.NewClient("client1", hostPort)
	require.Nil(t, err)
	defer user1Client.Close()

	_, err = user1Client.Start(context.Background(), "myjob", "/bin/true")
	assert.Nil(t, err)

	adminClient, err := jobmanager.NewClient(jobmanager.Superuser, hostPort)
	require.Nil(t, err)
	defer user1Client.Close()

	jobList, err := adminClient.List(context.Background())

	assert.Nil(t, err)
	assert.Equal(t, 1, len(jobList))

	cancel()
	wg.Wait()
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
		if ruErr != nil {
			t.Error()
		}
		wg.Done()
	}()

	return listener.Addr().String(), nil
}
