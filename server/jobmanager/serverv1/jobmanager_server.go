package serverv1

import (
	"context"

	"github.com/obaraelijah/secureproc/pkg/io"
	"github.com/obaraelijah/secureproc/pkg/jobmanager"
	"github.com/obaraelijah/secureproc/service/jobmanager/jobmanagerv1"
)

// jobmanagerServer implements the gRPC handler for the jobmanager service.
type jobmanagerServer struct {
	jobmanagerv1.UnimplementedJobManagerServer
	jm *jobmanager.Manager
}

// NewJobmanagerServer creates and returns a jobmanagerServer.
func NewJobmanagerServer() *jobmanagerServer {
	return NewJobManagerServerDetailed(jobmanager.NewManager())
}

// NewJobManagerServerDetailed creates a new jobmanagerServer with a custom
// underlying jobmanager.Manager.
func NewJobManagerServerDetailed(manager *jobmanager.Manager) *jobmanagerServer {
	return &jobmanagerServer{
		jm: manager,
	}
}

func (s *jobmanagerServer) Start(
	ctx context.Context,
	jcr *jobmanagerv1.JobCreationRequest,
) (*jobmanagerv1.Job, error) {

	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	job, err := s.jm.Start(userID, jcr.GetName(), jcr.GetProgramPath(), jcr.GetArguments())
	if err != nil {
		return nil, err
	}

	jobResponse := &jobmanagerv1.Job{
		Id:   &jobmanagerv1.JobID{Id: job.ID().String()},
		Name: job.Name(),
	}

	return jobResponse, nil
}

func (s *jobmanagerServer) Stop(
	ctx context.Context,
	requestJobID *jobmanagerv1.JobID,
) (*jobmanagerv1.NilMessage, error) {

	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	err = s.jm.Stop(userID, requestJobID.Id)
	if err != nil {
		return nil, err
	}

	return &jobmanagerv1.NilMessage{}, nil
}

func internalToExternalStatusV1(internalStatus *jobmanager.JobStatus) *jobmanagerv1.JobStatus {
	errMsg := ""

	if internalStatus.RunError != nil {
		errMsg = internalStatus.RunError.Error()
	}

	return &jobmanagerv1.JobStatus{
		Job: &jobmanagerv1.Job{
			Id: &jobmanagerv1.JobID{
				Id: internalStatus.ID,
			},
			Name: internalStatus.Name,
		},
		Owner:        internalStatus.Owner,
		IsRunning:    internalStatus.Running,
		Pid:          int32(internalStatus.Pid),
		ExitCode:     int32(internalStatus.ExitCode),
		SignalNumber: int32(internalStatus.SignalNum),
		ErrorMessage: errMsg,
	}
}

func (s *jobmanagerServer) Query(
	ctx context.Context,
	requestJobID *jobmanagerv1.JobID,
) (*jobmanagerv1.JobStatus, error) {

	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	jobStatus, err := s.jm.Status(userID, requestJobID.Id)
	if err != nil {
		return nil, err
	}

	return internalToExternalStatusV1(jobStatus), nil
}

func (s *jobmanagerServer) List(
	ctx context.Context,
	_ *jobmanagerv1.NilMessage,
) (*jobmanagerv1.JobStatusList, error) {

	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	statusList := s.jm.List(userID)

	responseStatusList := &jobmanagerv1.JobStatusList{
		JobStatusList: make([]*jobmanagerv1.JobStatus, 0, len(statusList)),
	}

	for _, status := range statusList {
		responseStatusList.JobStatusList =
			append(responseStatusList.JobStatusList, internalToExternalStatusV1(status))
	}

	return responseStatusList, nil
}

func (s *jobmanagerServer) StreamOutput(
	request *jobmanagerv1.StreamOutputRequest,
	response jobmanagerv1.JobManager_StreamOutputServer,
) error {

	userID, err := GetUserIDFromContext(response.Context())
	if err != nil {
		return err
	}

	var byteStream *io.ByteStream

	switch streamType := request.GetOutputStream(); streamType {
	case jobmanagerv1.OutputStream_STDOUT:
		byteStream, err = s.jm.StdoutStream(userID, request.JobID.Id)

	case jobmanagerv1.OutputStream_STDERR:
		byteStream, err = s.jm.StderrStream(userID, request.JobID.Id)

	default:
		return jobmanager.InvalidArgument
	}

	if err != nil {
		return err
	}
	defer byteStream.Close()

	for {
		select {
		// We don't have a specific requirement to support a deadline for the
		// stream operation from our client. But, since this is an API, someone
		// might develop a client for it that does want to set a deadline.
		// This will support that.
		case <-response.Context().Done():
			return context.DeadlineExceeded

		case data, ok := <-byteStream.Stream():
			if !ok {
				return nil
			}
			response.Send(&jobmanagerv1.JobOutput{Output: data})
		}
	}
}
