package testutils

import "errors"

var (
	FakeReadError  = errors.New("fake read error")
	FakeWriteError = errors.New("fake write error")
)

type CustomReader struct {
}

func NewCustomReader() *CustomReader {
	return &CustomReader{}
}

func (r *CustomReader) Read(p []byte) (n int, err error) {
	return 0, FakeReadError
}

type CustomWriter struct {
}

func NewCustomWriter() *CustomWriter {
	return &CustomWriter{}
}

func (w *CustomWriter) Write(p []byte) (n int, err error) {
	return 0, FakeWriteError
}
