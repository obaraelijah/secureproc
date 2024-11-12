package main

import (
	"fmt"
	"strings"

	"github.com/obaraelijah/secureproc/pkg/jobmanager"
)

func runTest() {

	job := jobmanager.NewJob("theOwner", "my-test", nil,
		"/bin/bash",
		"-c",
		"echo $$",
	)

	if err := job.Start(); err != nil {
		panic(err)
	}

	output := <-job.StdoutStream().Stream()
	if output == nil {
		panic("Received nil response")
	}

	outputStr := strings.TrimSpace(string(output))

	if outputStr != "1" {
		panic("The pid of the process should have been 1, not " + outputStr)
	}

	fmt.Println(outputStr)
}

// Sample run:
//     Determining the job's PID in its namespace
//     1

func main() {
	fmt.Println("Determining the job's PID in its namespace")
	runTest()
}
