package serverv1

import (
	"context"
	"errors"

	"github.com/obaraelijah/secureproc/pkg/io"
	"github.com/obaraelijah/secureproc/pkg/jobmanager"
	v1 "github.com/obaraelijah/secureproc/service/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserIDContext struct{}

// jobmanagerServer implements the gRPC handler for the jobmanager service.
type jobmanagerServer struct {
	v1.UnimplementedJobManagerServer
	jm *jobmanager.Manager
}

// NewJobmanagerServer creates and returns a jobmanagerServer.
func NewJobmanagerServer() *jobmanagerServer {
	return &jobmanagerServer{
		jm: jobmanager.NewManager(),
	}
}

func (s *jobmanagerServer) Start(ctx context.Context, jcr *v1.JobCreationRequest) (*v1.Job, error) {
	userID, err := s.userIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	job, err := s.jm.Start(userID, jcr.GetName(), jcr.GetProgramPath(), jcr.GetArguments())
	if err != nil {
		return nil, status.Errorf(errorToGRPCErrorCode(err), err.Error())
	}

	jobResponse := &v1.Job{
		Id:   &v1.JobID{Id: job.ID().String()},
		Name: job.Name(),
	}

	return jobResponse, nil
}

func (s *jobmanagerServer) Stop(ctx context.Context, requestJobID *v1.JobID) (*v1.NilMessage, error) {
	userID, err := s.userIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	err = s.jm.Stop(userID, requestJobID.Id)
	if err != nil {
		return nil, status.Errorf(errorToGRPCErrorCode(err), err.Error())
	}

	return &v1.NilMessage{}, nil
}

func internalToExternalStatusV1(internalStatus *jobmanager.JobStatus) *v1.JobStatus {
	errMsg := ""

	if internalStatus.RunError != nil {
		errMsg = internalStatus.RunError.Error()
	}

	return &v1.JobStatus{
		Job: &v1.Job{
			Id: &v1.JobID{
				Id: internalStatus.ID,
			},
			Name: internalStatus.Name,
		},
		IsRunning:    internalStatus.Running,
		ExitCode:     int32(internalStatus.ExitCode),
		SignalNumber: int32(internalStatus.SignalNum),
		ErrorMessage: errMsg,
	}
}

func (s *jobmanagerServer) Query(ctx context.Context, requestJobID *v1.JobID) (*v1.JobStatus, error) {
	userID, err := s.userIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	jobStatus, err := s.jm.Status(userID, requestJobID.Id)
	if err != nil {
		return nil, status.Errorf(errorToGRPCErrorCode(err), err.Error())
	}

	return internalToExternalStatusV1(jobStatus), nil
}

func (s *jobmanagerServer) List(ctx context.Context, _ *v1.NilMessage) (*v1.JobStatusList, error) {
	userID, err := s.userIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	statusList := s.jm.List(userID)

	responseStatusList := &v1.JobStatusList{
		JobStatusList: make([]*v1.JobStatus, 0, len(statusList)),
	}

	for _, status := range statusList {
		responseStatusList.JobStatusList =
			append(responseStatusList.JobStatusList, internalToExternalStatusV1(status))
	}

	return responseStatusList, nil
}

func (s *jobmanagerServer) StreamOutput(
	request *v1.StreamOutputRequest,
	response v1.JobManager_StreamOutputServer,
) error {

	userID, err := s.userIDFromContext(response.Context())
	if err != nil {
		return err
	}

	var byteStream *io.ByteStream

	switch streamType := request.GetOutputStream(); streamType {
	case v1.OutputStream_STDOUT:
		byteStream, err = s.jm.StdoutStream(userID, request.JobID.Id)

	case v1.OutputStream_STDERR:
		byteStream, err = s.jm.StderrStream(userID, request.JobID.Id)

	default:
		return status.Errorf(codes.InvalidArgument, "unsupported stream: %v", streamType)
	}

	if err != nil {
		return status.Errorf(errorToGRPCErrorCode(err), err.Error())
	}

	for data := range byteStream.Stream() {
		response.Send(&v1.JobOutput{Output: data})
	}

	return nil
}

// userIDFromContext extracts and returns the userID from the given context.
// If the userID doesn't exist, returns an error ready to be returned by a
// gRPC API.
func (s *jobmanagerServer) userIDFromContext(ctx context.Context) (string, error) {
	if userID, ok := ctx.Value(&UserIDContext{}).(string); ok && userID != "" {
		return userID, nil
	}

	return "", status.Errorf(codes.Unauthenticated, "jobmanager: user unauthenticated")
}

// errorToGRPCErrorCode maps the given error to a suitable gRPC error code.
// If no mapping is found, it will return codes.Internal.
func errorToGRPCErrorCode(err error) codes.Code {
	code := codes.Internal

	if errors.Is(err, &jobmanager.JobExistsError{}) {
		code = codes.AlreadyExists
	} else if errors.Is(err, &jobmanager.JobNotFoundError{}) {
		code = codes.NotFound
	} else if errors.Is(err, &jobmanager.InvalidJobID{}) {
		code = codes.InvalidArgument
	}

	return code
}
