package jobmanager

import (
	"os/exec"
	"sync"

	"github.com/google/uuid"
)

type job struct {
	mutex sync.Mutex
	id    uuid.UUID
	name  string
	cmd   exec.Cmd
}
