//go:build integration
// +build integration

package memorylimit_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/obaraelijah/secureproc/pkg/cgroup/cgroupv1"
	"github.com/obaraelijah/secureproc/pkg/jobmanager"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_memorylimit(t *testing.T) {
	noLimitCount := runTest(t)
	require.Equal(t, 0, noLimitCount)

	limitCount := runTest(t, &cgroupv1.MemoryController{Limit: "1M"})
	require.Greater(t, limitCount, 0)
}

func runTest(t *testing.T, controllers ...cgroupv1.Controller) int {

	cmd := fmt.Sprintf("/usr/bin/stress-ng --vm 1 --vm-bytes %d --timeout 10 --oomable -v 2>&1 | grep 'OOM killer'", 1024*1024*1024)

	job := jobmanager.NewJob("theOwner", "my-test", controllers,
		"/bin/bash",
		"-c",
		cmd)

	require.Nil(t, job.Start())

	outputBuffer := bytes.Buffer{}
	for output := range job.StdoutStream().Stream() {
		outputBuffer.Write(output)
	}

	lineCount := 0

	var err error

	for {
		_, err = outputBuffer.ReadString('\n')
		if err != nil {
			assert.Equal(t, err, io.EOF)
			break
		}

		lineCount++
	}

	return lineCount
}
