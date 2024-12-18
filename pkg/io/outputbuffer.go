package io

import (
	goio "io"
)

// OutputBuffer is an abstraction over buffers to which the job manager
// can write output.
type OutputBuffer interface {
	goio.Writer
	goio.ReaderAt
	goio.Closer

	waitForChange(nextByte int64) (bufferSize int64, closed bool)
}
