package usecase

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/xoltawn/simple-file-storage-file-service/domain"
)

type fileUsecase struct {
	fileStorage    domain.FileStorage
	fileRepo       domain.FileRepository
	fileDownloader domain.FileDownloader
	imagesPath     string
}

func NewFileUsecase(
	fileStorage domain.FileStorage,
	fileRepo domain.FileRepository,
	fileDownloader domain.FileDownloader,
	imagesPath string,
) *fileUsecase {
	return &fileUsecase{
		fileStorage:    fileStorage,
		fileRepo:       fileRepo,
		imagesPath:     imagesPath,
		fileDownloader: fileDownloader,
	}
}

func (f *fileUsecase) SaveFile(ctx context.Context, fileBytes []byte, fileInfo *domain.File, path string) (err error) {

	err = f.fileStorage.SaveFile(ctx, fileBytes, fileInfo, path)
	if err != nil {
		return
	}

	(*fileInfo).CreatedAt = time.Now().UTC().String()
	(*fileInfo).LocalName = uuid.NewString()

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

func (f *fileUsecase) SaveMutltipleFiles(ctx context.Context, filesWithByte []*domain.FileWithBytes) (err error) {
	//TODO: concurrent implementation
	for _, fileToDownload := range filesWithByte {
		downloadErr := f.fileDownloader.Download(fileToDownload)
		if downloadErr != nil {
			return
		}
	}

	files := []*domain.File{}
	for _, fileInfo := range filesWithByte {
		(*fileInfo).LocalName = uuid.NewString()
		err = f.fileStorage.SaveFile(ctx, fileInfo.Data, fileInfo.File, f.imagesPath)
		if err != nil {
			return
		}

		(*fileInfo).CreatedAt = time.Now().UTC().String()
		f := ((*fileInfo).File)
		if f != nil {
			files = append(files, f)
		}
	}

	err = f.fileRepo.SaveMutltipleFiles(ctx, files)
	if err != nil {
		removeErr := f.fileStorage.RemoveFiles(ctx, files)
		if removeErr != nil {
			log.Println(err)
		}
	}
	return
}
