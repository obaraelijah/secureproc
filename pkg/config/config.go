package config

import (
	"fmt"
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
	if dir, ok := os.LookupEnv("CGEXEC_PATH"); ok {
		fmt.Println(dir)
		CgexecPath = dir
		return
	}

	const binaryName = "/cgexec"

	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	CgexecPath = path.Dir(exe) + binaryName
}

const (
	CgroupDefaultCpuLimit        = 0.5
	CgroupDefaultMemoryLimit     = "2M"
	CgroupDefaultBlkioDevice     = "8:16"
	CgroupDefaultBlkioWriteLimit = CgroupDefaultBlkioDevice + " 20971520"
	CgroupDefaultBlkioReadLimit  = CgroupDefaultBlkioDevice + " 41943040"
)
