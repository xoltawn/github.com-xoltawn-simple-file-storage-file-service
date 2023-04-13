package usecase

type bytesToLinksConvertor struct {
}

// NewBytesToLinksConvertor is the builder function for BytesToLinksConvertor
func NewBytesToLinksConvertor() *bytesToLinksConvertor {
	return &bytesToLinksConvertor{}
}

func (c *bytesToLinksConvertor) Convert([]byte) (links []string, err error) {
	return
}
