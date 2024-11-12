package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"

	"github.com/obaraelijah/secureproc/pkg/jobmanager"
)

func runTest() {

	numValues := 100 * 1000

	job := jobmanager.NewJob("theOwner", "my-test", nil,
		"/bin/bash",
		"-c",
		"for ((i = 0; i < 100; ++i)); do for((j = 0; j < 1000; ++j)); do echo $RANDOM; done; sleep 0.25; done",
	)

	if err := job.Start(); err != nil {
		panic(err)
	}

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
		if err != nil {
			panic(fmt.Sprintf("Unexpected error at value number %d, goroutine 0: %v", i, err))
		}

		for j := 1; j < len(buckets); j++ {
			value, err := buckets[j].ReadString('\n')
			if err != nil {
				panic(fmt.Sprintf("Unexpected error at value number %d, goroutine %d: %v", i, j, err))
			}

			if expectedValue != value {
				panic(fmt.Sprintf("value mismatch at %d: %s/%s\n", i, expectedValue, value))
			}
		}
	}

	// There should be no more values; all readers should be at EOF
	for i := 0; i < len(buckets); i++ {
		_, err := buckets[i].ReadString('\n')
		if err != io.EOF {
			panic(fmt.Sprintf("Unexpected additional data from goroutine %d", i))
		}
	}

	fmt.Printf("Matched %d matched generated values across %d goroutines\n", numValues, numGoroutines)
}

// The job generates 100000 random numbers and prints them to standard output
// The program starts 100 goroutines to consume that output.  Each goroutine
// counts and prints the number of bytes that it receives.

func main() {
	runTest()
}
