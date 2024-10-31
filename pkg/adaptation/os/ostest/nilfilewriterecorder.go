package ostest

import "github.com/obaraelijah/secureproc/pkg/adaptation/os"

type NilFileWriter struct{}

func (w *NilFileWriter) WriteFile(string, []byte, os.FileMode) error {
	return nil
}
