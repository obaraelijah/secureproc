package testserverv1

import (
	"context"

	"github.com/obaraelijah/secureproc/service/jobmanager/jobmanagerv1"
	"google.golang.org/grpc/metadata"
)

// MockJobManagerStreamServer mocks the APIs used by a JobManager server.
type MockJobmanagerStreamServer struct {
	LastJobOutput *jobmanagerv1.JobOutput
	SendCount     int
	SendError     error
	NextContext   context.Context
}

func (m *MockJobmanagerStreamServer) Send(output *jobmanagerv1.JobOutput) error {
	m.SendCount++
	m.LastJobOutput = output
	return m.SendError
}

func (m *MockJobmanagerStreamServer) Context() context.Context {
	return m.NextContext
}

// SetHeader is not yet implemented; it will panic.
func (m *MockJobmanagerStreamServer) SetHeader(metadata.MD) error {
	panic("unimplemented")
}

// SendHeader is not yet implemented; it will panic.
func (m *MockJobmanagerStreamServer) SendHeader(metadata.MD) error {
	panic("unimplemented")
}

// SetTrailer is not yet implemented; it will panic.
func (m *MockJobmanagerStreamServer) SetTrailer(metadata.MD) {
	panic("unimplemented")
}

// SendMsg is not yet implemented; it will panic.
func (m *MockJobmanagerStreamServer) SendMsg(interface{}) error {
	panic("unimplemented")
}

// RecvMsg is not yet implemented; it will panic.
func (m *MockJobmanagerStreamServer) RecvMsg(interface{}) error {
	panic("unimplemented")
}
