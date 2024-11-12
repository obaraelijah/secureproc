//go:build integration
// +build integration

package pidnamespace_test

import (
	"strings"
	"testing"

	"github.com/obaraelijah/secureproc/pkg/jobmanager"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_jobPidIsOne(t *testing.T) {
	job := jobmanager.NewJob("theOwner", "my-test", nil,
		"/bin/bash",
		"-c",
		"echo $$",
	)
	defer job.Stop()

	err := job.Start()
	require.Nil(t, err)

	output := <-job.StdoutStream().Stream()
	require.NotNil(t, output)

	outputStr := strings.TrimSpace(string(output))
	assert.Equal(t, "1", outputStr)
}
