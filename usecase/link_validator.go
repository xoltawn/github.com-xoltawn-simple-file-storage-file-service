package usecase

import "strings"

type linkValidator struct {
}

// NewlinkValidator is the builder function for linkValidator
func NewLinkValidator() *linkValidator {
	return &linkValidator{}
}

func (c *linkValidator) IsLink(txt string) (result bool) {
	txt = strings.Trim(txt, " ")
	if strings.HasPrefix(txt, "http://") || strings.HasPrefix(txt, "https://") {
		return true
	}
	return
}
