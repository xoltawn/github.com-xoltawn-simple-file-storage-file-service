package localstorage

import (
	"context"
	"fmt"
	"os"

	"github.com/xoltawn/simple-file-storage-file-service/domain"
)

type localstorage struct {
	imagesPath string
}

// NewLocalstorage is the builder function for localstorage
func NewLocalStorage(imagesPath string) *localstorage {
	return &localstorage{imagesPath: imagesPath}
}

func (s *localstorage) SaveFile(ctx context.Context, fileBytes []byte, fileInfo *domain.File, path string) (err error) {
	err = s.CreatePathIfNotExist(path)
	if err != nil {
		return
	}

	path = fmt.Sprint(path, "/", fileInfo.LocalName, ".", fileInfo.FileExtension)
	err = os.WriteFile(path, fileBytes, 0644)
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

func (s *localstorage) RemoveFiles(ctx context.Context, files []*domain.File) (err error) {
	for _, file := range files {
		err = os.Remove(fmt.Sprint(s.imagesPath, "/", file.LocalName, ".", file.FileExtension))
		if err != nil {
			return
		}
	}
	return
}
