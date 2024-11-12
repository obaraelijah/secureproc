//go:build integration
// +build integration

package networknamespace_test

import (
	"encoding/json"
	"testing"

	"github.com/obaraelijah/secureproc/pkg/jobmanager"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_networknamespace(t *testing.T) {
	job := jobmanager.NewJob("theOwner", "my-test", nil,
		"/bin/ip",
		"-j",
		"link",
	)

	require.Nil(t, job.Start())
	defer job.Stop()

	var outputBuffer []byte

	for output := range job.StdoutStream().Stream() {
		outputBuffer = append(outputBuffer, output...)
	}

	type iface struct {
		Ifname *string `json:"ifname,omitempty"`
	}
	var ifaceList []iface

	err := json.Unmarshal(outputBuffer, &ifaceList)
	assert.Nil(t, err)

	require.Equal(t, 2, len(ifaceList))
	require.NotNil(t, ifaceList[0].Ifname)
	require.NotNil(t, ifaceList[1].Ifname)
	assert.Equal(t, "lo", *ifaceList[0].Ifname)
	assert.Equal(t, "sit0", *ifaceList[1].Ifname)
}
