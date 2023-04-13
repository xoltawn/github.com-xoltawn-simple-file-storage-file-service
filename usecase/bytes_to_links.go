package usecase

import (
	"bufio"

	"github.com/xoltawn/simple-file-storage-file-service/domain"
)

type bytesToLinksConvertor struct {
	bytesToReaderConvertor domain.BytesToReaderConvertor
	linkValidator          domain.LinkValidator
}

// NewBytesToLinksConvertor is the builder function for BytesToLinksConvertor
func NewBytesToLinksConvertor(bytesToReaderConvertor domain.BytesToReaderConvertor, linkValidator domain.LinkValidator) *bytesToLinksConvertor {
	return &bytesToLinksConvertor{
		bytesToReaderConvertor: bytesToReaderConvertor,
		linkValidator:          linkValidator,
	}
}

func (c *bytesToLinksConvertor) Convert(data []byte) (links []string, err error) {
	fileReader, err := c.bytesToReaderConvertor.Convert(data)
	if err != nil {
		return
	}

	linksScanner := bufio.NewScanner(fileReader)

	for linksScanner.Scan() {
		if ok := c.linkValidator.IsLink(linksScanner.Text()); ok {
			links = append(links, linksScanner.Text())
		}
	}
	return
}
