package cgroupv1_test

import (
	"fmt"
	"testing"

	"github.com/obaraelijah/secureproc/pkg/adaptation/os"
	"github.com/obaraelijah/secureproc/pkg/adaptation/os/ostest"
	"github.com/obaraelijah/secureproc/pkg/cgroup/cgroupv1"
	"github.com/stretchr/testify/assert"
)

func Test_memory_Apply(t *testing.T) {
	path := "/sys/fs/cgroup/jobs/889f7cc2-9935-4773-aaa1-b94478abc923"
	writeRecorder := ostest.WriteFileMock{}
	adapter := &os.Adapter{
		WriteFileFn: writeRecorder.WriteFile,
	}

	limit := "500M"
	mem := &cgroupv1.MemoryController{OsAdapter: adapter, Limit: limit}
	mem.Apply(path)

	assert.Equal(t, 1, len(writeRecorder.Events))
	assert.Equal(t, fmt.Sprintf("%s/%s", path, cgroupv1.MemoryLimitInBytesFilename), writeRecorder.Events[0].Name)
	assert.Equal(t, []byte(limit), writeRecorder.Events[0].Data)
}
