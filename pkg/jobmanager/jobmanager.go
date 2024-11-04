package jobmanager

import (
	"os/exec"
	"sync"

	"github.com/google/uuid"
	"github.com/obaraelijah/secureproc/pkg/io"
)

type job struct {
	mutex        sync.Mutex
	id           uuid.UUID
	name         string
	cmd          exec.Cmd
	stdoutBuffer io.OutputBuffer
	stderrBuffer io.OutputBuffer
}

func NewJob(name string, command string, args ...string) *job {
	return NewJobDetailed(name, io.NewMemoryBuffer(), io.NewMemoryBuffer(), command, args...)
}

func NewJobDetailed(
	name string,
	stdoutBuffer io.OutputBuffer,
	stderrBuffer io.OutputBuffer,
	command string,
	args ...string,
) *job {
	return &job{
		id:           uuid.New(),
		name:         name,
		stdoutBuffer: stdoutBuffer,
		stderrBuffer: stderrBuffer,
	}
}
