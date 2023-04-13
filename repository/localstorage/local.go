package localstorage

import (
	"context"
	"fmt"
	"os"

	"github.com/xoltawn/simple-file-storage-file-service/domain"
)

type localstorage struct {
}

// NewLocalstorage is the builder function for localstorage
func NewLocalStorage() *localstorage {
	return &localstorage{}
}

func (s *localstorage) SaveFile(ctx context.Context, fileBytes []byte, fileInfo *domain.File, path string) (err error) {
	err = s.CreatePathIfNotExist(path)
	if err != nil {
		return
	}

	err = os.WriteFile(fmt.Sprint(path, "/", fileInfo.LocalName, ".", fileInfo.FileExtension), fileBytes, 0644)
	if err != nil {
		return
	}

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
