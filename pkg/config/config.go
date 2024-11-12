package config

import (
	"os"
	"path"
)

// Note: Generally I would avoid having a "config.go" as a place for a bunch of
//
//	unrelated constants.  However, here these constants represent the
//	hard-coded values for the exercise, so I thought this might make it
//	easier to adjust the values for experimentation if I put them all in
//	one place.
var (
	CgexecPath string
)

// init sets CgexecPath based on the position of the current executable
func init() {
	const binaryName = "/cgexec"

	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	CgexecPath = path.Dir(exe) + binaryName

	if _, err := os.Stat(CgexecPath); err != nil {
		// Fallback to a well-known location; this would normally not be in /tmp
		CgexecPath = "/tmp" + binaryName
		if _, err = os.Stat(CgexecPath); err != nil {
			panic(err)
		}
	}
}

const (
	CgroupDefaultCpuLimit        = 0.5
	CgroupDefaultMemoryLimit     = "2M"
	CgroupDefaultBlkioDevice     = "8:16"
	CgroupDefaultBlkioWriteLimit = CgroupDefaultBlkioDevice + " 20971520"
	CgroupDefaultBlkioReadLimit  = CgroupDefaultBlkioDevice + " 41943040"
)
