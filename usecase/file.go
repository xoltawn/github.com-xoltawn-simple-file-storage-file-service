package usecase

import (
	"context"

	"github.com/xoltawn/simple-file-storage-file-service/domain"
)

type fileUsecase struct {
	fileStorage domain.FileStorage
	fileRepo    domain.FileRepository
	imagesPath  string
}

func NewFileUsecase(fileStorage domain.FileStorage, fileRepo domain.FileRepository, imagesPath string) *fileUsecase {
	return &fileUsecase{fileStorage: fileStorage, fileRepo: fileRepo, imagesPath: imagesPath}
}

func (f *fileUsecase) SaveFile(ctx context.Context, fileBytes []byte, fileInfo *domain.File) (err error) {
	err = f.fileStorage.SaveFile(ctx, fileBytes, fileInfo, f.imagesPath)
	return
}

func (f *fileUsecase) FetchFiles(ctx context.Context, limit, offset int) (files []domain.File, err error) {
	return
}
