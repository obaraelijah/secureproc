package cgroup

import "github.com/obaraelijah/secureproc/pkg/adaptation/os"

const (
	MemoryLimitInBytesFilename = "memory.limit_in_bytes"
)

type memory struct {
	base
	limit *string
}

func NewMemory() *memory {
	return NewMemoryDetailed(nil)
}

func NewMemoryDetailed(osAdapter *os.Adapter) *memory {
	return &memory{
		base: newBase("memory", osAdapter),
	}
}

func (m *memory) SetLimit(value string) *memory {
	m.limit = &value

	return m
}

func (m *memory) Apply(path string) error {
	if m.limit != nil {
		if err := m.write([]byte(*m.limit), "%s/%s", path, MemoryLimitInBytesFilename); err != nil {
			return err
		}
	}

	return nil
}
