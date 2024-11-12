package main

import (
	"encoding/json"
	"fmt"

	"github.com/obaraelijah/secureproc/pkg/jobmanager"
)

func runTest() {

	job := jobmanager.NewJob("theOwner", "my-test", nil,
		"/bin/ip",
		"-j",
		"link",
	)

	if err := job.Start(); err != nil {
		panic(err)
	}

	var outputBuffer []byte

	for output := range job.StdoutStream().Stream() {
		outputBuffer = append(outputBuffer, output...)
	}

	type iface struct {
		Ifindex *int    `json:"ifindex,omitempty"`
		Ifname  *string `json:"ifname,omitempty"`
	}
	var ifaceList []iface

	if err := json.Unmarshal(outputBuffer, &ifaceList); err != nil {
		panic(err)
	}

	if len(ifaceList) != 2 {
		panic(fmt.Sprintf("Expected 2, found: %d", len(ifaceList)))
	}

	fmt.Println("Found expected number of network interface in new network namespace (2)")
}

// Sample run:
//     Running test to list all network interfaces avaialble to a job
//     Found expected number of network interface in new network namespace (2)

func main() {
	fmt.Println("Running test to list all network interfaces avaialble to a job")
	runTest()
}
