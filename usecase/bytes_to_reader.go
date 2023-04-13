package usecase

import (
	"bytes"
	"io"
)

type bytesToReaderConvertor struct {
}

// NewBytesToReaderConvertor is the builder function for BytesToReaderConvertor
func NewBytesToReaderConvertor() *bytesToReaderConvertor {
	return &bytesToReaderConvertor{}
}

func (c *bytesToReaderConvertor) Convert(data []byte) (reader io.Reader, err error) {
	reader = bytes.NewReader(data)
	return
}
