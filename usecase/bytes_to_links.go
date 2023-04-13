package usecase

import "github.com/xoltawn/simple-file-storage-file-service/domain"

type bytesToLinksConvertor struct {
	bytesToReaderConvertor domain.BytesToReaderConvertor
}

// NewBytesToLinksConvertor is the builder function for BytesToLinksConvertor
func NewBytesToLinksConvertor(bytesToReaderConvertor domain.BytesToReaderConvertor) *bytesToLinksConvertor {
	return &bytesToLinksConvertor{bytesToReaderConvertor: bytesToReaderConvertor}
}

func (c *bytesToLinksConvertor) Convert(data []byte) (links []string, err error) {
	_, err = c.bytesToReaderConvertor.Convert(data)
	if err != nil {
		return
	}
	return
}
