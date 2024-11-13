package serverv1

import (
	"context"

	v1 "github.com/obaraelijah/secureproc/service/v1"
)

// jobmanagerServer implements the gRPC handler for the jobmanager service.
type jobmanagerServer struct {
	v1.UnimplementedJobManagerServer
}

func NewJobmanagerServer() *jobmanagerServer {
	return &jobmanagerServer{}
}

func (s *jobmanagerServer) Start(ctx context.Context, jcr *v1.JobCreationRequest) (*v1.Job, error) {
	return s.UnimplementedJobManagerServer.Start(ctx, jcr)
}

func (s *jobmanagerServer) Stop(ctx context.Context, requestJobID *v1.JobID) (*v1.NilMessage, error) {
	return s.UnimplementedJobManagerServer.Stop(ctx, requestJobID)
}

func (s *jobmanagerServer) Query(ctx context.Context, requestJobID *v1.JobID) (*v1.JobStatus, error) {
	return s.UnimplementedJobManagerServer.Query(ctx, requestJobID)
}

func (s *jobmanagerServer) List(ctx context.Context, nm *v1.NilMessage) (*v1.JobStatusList, error) {
	return s.UnimplementedJobManagerServer.List(ctx, nm)
}

func (s *jobmanagerServer) StreamOutput(
	request *v1.StreamOutputRequest,
	response v1.JobManager_StreamOutputServer,
) error {
	return s.UnimplementedJobManagerServer.StreamOutput(request, response)
}
