package jobmanagertest

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/obaraelijah/secureproc/pkg/cgroup/v1"
	"github.com/obaraelijah/secureproc/pkg/io"
	"github.com/obaraelijah/secureproc/pkg/jobmanager"
)

type mockJob struct {
	name    string
	id      uuid.UUID
	running bool
	stdout  io.OutputBuffer
	stderr  io.OutputBuffer
}

func NewMockJob(
	jobName string, controllers []cgroup.Controller,
	programPath string,
	arguments ...string,
) jobmanager.Job {
	return &mockJob{
		name:   jobName,
		id:     uuid.New(),
		stdout: io.NewMemoryBuffer(),
		stderr: io.NewMemoryBuffer(),
	}
}

func (m *mockJob) Start() error {

	if m.running {
		return fmt.Errorf("job %s (%v) has already been started", m.name, m.id)
	}

	m.running = true
	_, _ = m.stdout.Write([]byte("this is standard output"))
	_, _ = m.stderr.Write([]byte("this is standard error"))

	return nil
}

func (m *mockJob) Stop() error {
	m.running = false
	m.stdout.Close()
	m.stderr.Close()
	return nil
}

func (m *mockJob) StdoutStream() *io.ByteStream {
	return io.NewByteStream(m.stdout)
}

func (m *mockJob) StderrStream() *io.ByteStream {
	return io.NewByteStream(m.stderr)
}

func (m *mockJob) Status() *jobmanager.JobStatus {
	exitCode := -1

	if m.running {
		exitCode = 0
	}

	return &jobmanager.JobStatus{
		Name:     m.name,
		Id:       m.id.String(),
		Running:  m.running,
		Pid:      1234,
		ExitCode: exitCode,
	}
}

func (m *mockJob) Id() uuid.UUID {
	return m.id
}