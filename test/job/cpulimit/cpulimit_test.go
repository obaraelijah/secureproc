//go:build integration
// +build integration

package cpulimit_test

import (
	"bytes"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/obaraelijah/secureproc/pkg/cgroup/cgroupv1"
	"github.com/obaraelijah/secureproc/pkg/jobmanager"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_cpulimit(t *testing.T) {
	oneCpuResult := runTest(t)
	halfCpuResult := runTest(t, &cgroupv1.CpuController{Cpus: 0.5})

	assert.True(t, aboutHalf(oneCpuResult, halfCpuResult))
}

func runTest(t *testing.T, controllers ...cgroupv1.Controller) float64 {

	job := jobmanager.NewJob("theOwner", "my-test", controllers,
		"/bin/bash",
		"-c",
		"/usr/bin/stress-ng --cpu 1 --timeout 10 --times 2>&1 | "+
			"grep 'user time' | sed -e s'/.*( *//' -e 's/%.$//'",
	)

	require.Nil(t, job.Start())

	allOutput := bytes.Buffer{}

	for output := range job.StdoutStream().Stream() {
		allOutput.Write(output)
	}

	output, err := allOutput.ReadString('\n')
	assert.Nil(t, err)

	value, err := strconv.ParseFloat(strings.TrimSpace(output), 64)
	assert.Nil(t, err)

	return value
}

func aboutHalf(firstResult, secondResult float64) bool {
	const closenessThreshold float64 = 0.5

	return math.Abs((firstResult/2.0)-secondResult) <= closenessThreshold
}
