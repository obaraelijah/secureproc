package cgroupv1_test

import (
	"fmt"
	"testing"

	"github.com/obaraelijah/secureproc/pkg/adaptation/os"
	"github.com/obaraelijah/secureproc/pkg/adaptation/os/ostest"
	"github.com/obaraelijah/secureproc/pkg/cgroup/cgroupv1"

	"github.com/stretchr/testify/assert"
)

func Test_cpu_Apply(t *testing.T) {
	path := "/sys/fs/cgroup/jobs/889f7cc2-9935-4773-aaa1-b94478abc923"
	writeRecorder := ostest.WriteFileMock{}
	adapter := &os.Adapter{
		WriteFileFn: writeRecorder.WriteFile,
	}

	cpu := cgroupv1.CpuController{OsAdapter: adapter, Cpus: 2.0}
	cpu.Apply(path)

	assert.Equal(t, 2, len(writeRecorder.Events))
	assert.Equal(t, fmt.Sprintf("%s/%s", path, cgroupv1.CpuPeriodFilename), writeRecorder.Events[0].Name)
	assert.Equal(t, []byte("100000"), writeRecorder.Events[0].Data)

	assert.Equal(t, fmt.Sprintf("%s/%s", path, cgroupv1.CpuQuotaFilename), writeRecorder.Events[1].Name)
	assert.Equal(t, []byte("200000"), writeRecorder.Events[1].Data)
}
