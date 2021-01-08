package testhelper

import "errors"

type MockErrorIoReader struct {
}

func (m MockErrorIoReader) Read(x []byte) (int, error) {
	return 0, errors.New("read error")
}
