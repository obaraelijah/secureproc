package command

import (
	"github.com/obaraelijah/secureproc/pkg/adaptation/os"
	"github.com/obaraelijah/secureproc/pkg/adaptation/os/syscall"
)

// Cgexec adds the current process to 0 or more specified cgroups, then
// execs the specfied command.  The format of args is:
//
//	args[0:n]   - cgroups files
//	args[n:n+1] - --
//	args[n+2:]  - command to exec and its arguments
//
// It returns an error if it failed to add itself to the requested cgroups
// or if it fails to exec the command.
func Cgexec(args []string) error {
	return CgexecDetailed(args, nil, nil)
}

// CgexecDetailed is wrapped by Cgexec and performs the same operation,
// optionally with concrete os and syscall adapters.
func CgexecDetailed(args []string, osa *os.Adapter, sa *syscall.Adapter) error {
	//TODO: dfs
}
