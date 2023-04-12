package domain

import (
	"context"

	_domain "github.com/xoltawn/simple-file-storage-sharedparts/domain"
)

// FileStorage provices interface for any type of file storages
type FileStorage interface {
	SaveFile(ctx context.Context, fileBytes []byte, fileInfo *_domain.File, path string) (err error)
}

// FileRepository ...
type FileRepository interface {
	SaveFile(ctx context.Context, fileInfo *_domain.File) (err error)
}

// FileUsecase ...
type FileUsecase interface {
	SaveFile(ctx context.Context, fileBytes []byte, fileInfo *_domain.File) (err error)
}
