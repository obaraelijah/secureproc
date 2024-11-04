package cgroup_test

import (
	"fmt"
	"testing"

	"github.com/obaraelijah/secureproc/pkg/adaptation/os"
	"github.com/obaraelijah/secureproc/pkg/adaptation/os/ostest"
	"github.com/obaraelijah/secureproc/pkg/cgroup/v1"
	"github.com/stretchr/testify/assert"
)

func Test_blkio_Apply(t *testing.T) {
	path := "/sys/fs/cgroup/jobs/889f7cc2-9935-4773-aaa1-b94478abc923"
	writeRecorder := ostest.WriteFileRecorder{}
	adapter := &os.Adapter{
		WriteFileFn: writeRecorder.WriteFile,
	}

	readBps := "1:2 1G"
	writeBps := "1:3 900M"

	blkio := cgroup.NewBlockIoControllerDetailed(adapter).
		SetReadBpsDevice(readBps).
		SetWriteBpsDevice(writeBps)

	blkio.Apply(path)

	assert.Equal(t, 2, len(writeRecorder.Events))
	assert.Equal(t, fmt.Sprintf("%s/%s", path, cgroup.BlkioThrottleReadBpsDevice), writeRecorder.Events[0].Name)
	assert.Equal(t, []byte(readBps), writeRecorder.Events[0].Data)

	assert.Equal(t, fmt.Sprintf("%s/%s", path, cgroup.BlkioThrottleWriteBpsDevice), writeRecorder.Events[1].Name)
	assert.Equal(t, []byte(writeBps), writeRecorder.Events[1].Data)
}