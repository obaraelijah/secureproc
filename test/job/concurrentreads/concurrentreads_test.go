//go:build integration
// +build integration

package concurrentreads_test

import (
	"bytes"
	"io"
	"sync"
	"testing"

	"github.com/obaraelijah/secureproc/pkg/jobmanager"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_concurrentReads(t *testing.T) {
	numValues := 100 * 1000

	job := jobmanager.NewJob("theOwner", "my-test", nil,
		"/bin/bash",
		"-c",
		"for ((i = 0; i < 100; ++i)); do for((j = 0; j < 1000; ++j)); do echo $RANDOM; done; sleep 0.25; done",
	)
	defer job.Stop()

	err := job.Start()
	require.Nil(t, err)

	const numGoroutines = 100
	var buckets [numGoroutines]bytes.Buffer

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(goroutineNum int) {
			for output := range job.StdoutStream().Stream() {
				buckets[goroutineNum].Write(output)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	for i := 0; i < numValues; i++ {
		expectedValue, err := buckets[0].ReadString('\n')
		assert.Nil(t, err)

		for j := 1; j < len(buckets); j++ {
			value, err := buckets[j].ReadString('\n')
			assert.Nil(t, err)

			assert.Equal(t, expectedValue, value)
		}
	}

	// There should be no more values; all readers should be at EOF
	for i := 0; i < len(buckets); i++ {
		_, err := buckets[i].ReadString('\n')
		assert.Equal(t, err, io.EOF)
	}
}
