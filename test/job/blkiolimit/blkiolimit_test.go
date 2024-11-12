//go:build integration
// +build integration

package blkiolimit_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/obaraelijah/secureproc/pkg/cgroup/cgroupv1"
	"github.com/obaraelijah/secureproc/pkg/jobmanager"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	// Todo: writing to the root filesystem isn't ideal.  In a real scenario
	//       this would be configurable.
	tmpFileDirectory = "/"

	deviceString = "8:16"
)

func Test_blkiolimit(t *testing.T) {
	noLimit := runTest(t)

	// The device portion must be a device, not a partition
	deviceString := fmt.Sprintf("%s %d", deviceString, 1024*1024*20)
	withLimit := runTest(t, &cgroupv1.BlockIOController{
		ReadBpsDevice:  deviceString,
		WriteBpsDevice: deviceString,
	})

	// Give it a little wiggle room.  This might need some additional experimentation
	// to dial in on a suitable value.
	const limitThreshold = 2.0

	assert.Less(t, withLimit, noLimit)
	assert.Less(t, withLimit, 20.0+limitThreshold)
}

func runTest(t *testing.T, controllers ...cgroupv1.Controller) float64 {

	file, err := ioutil.TempFile(tmpFileDirectory, "blkiolimit-test")
	require.Nil(t, err)
	defer os.Remove(file.Name())

	job := jobmanager.NewJob("theOwner", "my-test", controllers,
		"/bin/bash",
		"-c",
		"/bin/dd if=/dev/zero of="+file.Name()+" bs=4096 count=100000 oflag=direct 2>&1 |"+
			"grep copied | sed -e 's-^.*, --' -e 's/ .*$//'",
	)

	require.Nil(t, job.Start())

	allOutput := bytes.Buffer{}
	for output := range job.StdoutStream().Stream() {
		allOutput.Write(output)
	}

	output, err := allOutput.ReadString('\n')
	assert.Nil(t, err)

	value, err := strconv.ParseFloat(strings.TrimSpace(strings.TrimSpace(output)), 64)
	assert.Nil(t, err)

	return value
}
