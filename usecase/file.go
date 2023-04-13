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
	if err != nil {
		return
	}

	err = f.fileRepo.SaveFile(ctx, fileInfo)
	if err != nil {
		return
	}
	return
}

func (f *fileUsecase) FetchFiles(ctx context.Context, limit, offset int) (files []domain.File, err error) {
	files, err = f.fileRepo.FetchFiles(ctx, limit, offset)
	return
}

func (f *fileUsecase) SaveMutltipleFiles(ctx context.Context, files []*domain.FileWithBytes) (err error) {
	for _, fileInfo := range files {
		err = f.fileStorage.SaveFile(ctx, fileInfo.Data, &fileInfo.File, f.imagesPath)
		if err != nil {
			return
		}
	}
	return
}
