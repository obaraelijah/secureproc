package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/obaraelijah/secureproc/pkg/cgroup/cgroupv1"
	"github.com/obaraelijah/secureproc/pkg/jobmanager"
)

func runTest(controllers ...cgroupv1.Controller) float64 {

	job := jobmanager.NewJob("theOwner", "my-test", controllers,
		"/bin/bash",
		"-c",
		"/usr/bin/stress-ng --cpu 1 --timeout 10 --times 2>&1 | "+
			"grep 'user time' | sed -e s'/.*( *//' -e 's/%.$//'",
	)

	if err := job.Start(); err != nil {
		panic(err)
	}

	allOutput := bytes.Buffer{}

	for output := range job.StdoutStream().Stream() {
		allOutput.Write(output)
	}

	output, err := allOutput.ReadString('\n')
	if err != nil {
		panic(err)
	}

	value, err := strconv.ParseFloat(strings.TrimSpace(output), 64)
	if err != nil {
		panic(err)
	}

	fmt.Printf("user time: %3.2f\n", value)

	return value
}

// Sample run:
//     $ sudo go run test/job/cpulimit/cpulimit.go
//     Running CPU test with no cgroup constraints
//     user time: 8.33
//     Running CPU test with cgroup constraints at 0.5 CPU
//     user time: 4.18

func main() {
	fmt.Println("Running CPU test with no cgroup constraints")
	oneCpuResult := runTest()

	fmt.Println("Running CPU test with cgroup constraints at 0.5 CPU")
	halfCpuResult := runTest(&cgroupv1.CpuController{Cpus: 0.5})

	if !aboutHalf(oneCpuResult, halfCpuResult) {
		panic(fmt.Sprintf("%3.2f is not about half of %3.2f", halfCpuResult, oneCpuResult))
	}
}

func aboutHalf(firstResult, secondResult float64) bool {
	const closenessThreshold float64 = 0.5

	return math.Abs((firstResult/2.0)-secondResult) <= closenessThreshold
}
