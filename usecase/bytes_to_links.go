package usecase

import "strings"

type bytesToLinksConvertor struct {
}

// NewBytesToLinksConvertor is the builder function for BytesToLinksConvertor
func NewBytesToLinksConvertor() *bytesToLinksConvertor {
	return &bytesToLinksConvertor{}
}

func (c *bytesToLinksConvertor) Parse([]byte) (links []string, err error) {
	return
}

func (c *bytesToLinksConvertor) IsLink(txt string) (result bool) {
	if strings.HasPrefix(txt, "http://") || strings.HasPrefix(txt, "https://") {
		return true
	}
	return
}
