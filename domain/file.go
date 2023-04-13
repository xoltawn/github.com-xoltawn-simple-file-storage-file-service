package domain

import (
	"context"
	"errors"

	"github.com/xoltawn/simple-file-storage-sharedparts/domain"
)

var (
	ErrInsertingRecord = errors.New("error inserting record")
)

// File is the domain objects for stored files
type File struct {
	//OriginalUrl indicates the url from which the file was downloaded(used when file is downloaded from a link)
	OriginalURL string `json:"original_url"`
	//LocalName is the name given on storing
	LocalName string `json:"local_name"`
	//FileExtension ...
	FileExtension string `json:"file_extension"`
	//FileSize ...
	FileSize int64 `json:"file_size"`
	//CreatedAt ...
	CreatedAt string `json:"created_at"`
}

// FileStorage provices interface for any type of file storages
type FileStorage interface {
	SaveFile(ctx context.Context, fileBytes []byte, fileInfo *File, path string) (err error)
}

// FileRepository ...
type FileRepository interface {
	SaveFile(ctx context.Context, fileInfo *File) (err error)
	SaveMutltipleFiles(ctx context.Context, files []*File) (err error)
	FetchFiles(ctx context.Context, limit, offset int) (files []domain.File, err error)
}

// FileUsecase ...
type FileUsecase interface {
	SaveFile(ctx context.Context, fileBytes []byte, fileInfo *File) (err error)
}
