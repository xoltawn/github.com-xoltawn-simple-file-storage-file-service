package localstorage

import (
	"context"
	"os"

	"github.com/xoltawn/simple-file-storage-sharedparts/domain"
)

type localstorage struct {
}

// NewLocalstorage is the builder function for localstorage
func NewLocalStorage() *localstorage {
	return &localstorage{}
}

func (s *localstorage) SaveFile(ctx context.Context, fileBytes []byte, fileInfo *domain.File, path string) (err error) {
	return
}

func (s *localstorage) CreatePathIfNotExist(path string) (err error) {
	if _, err = os.Stat(path); err == nil {
		//path exists so exit the function
		return
	}

	//the path does not exist so we create one
	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		return
	}

	return
}
