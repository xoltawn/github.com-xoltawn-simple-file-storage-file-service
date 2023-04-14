package usecase

import (
	"context"
	"log"
	"net/http"
	"strings"
	"sync"
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
	(*fileInfo).LocalName = uuid.NewString()
	_, extension, ok := strings.Cut(http.DetectContentType(fileBytes), "/")
	if !ok {
		return domain.ErrFileExtensionNotSupported
	}
	(*fileInfo).FileExtension = extension
	(*fileInfo).FileSize = int64(len(fileBytes))

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
	var wg sync.WaitGroup
	var once sync.Once
	for _, fileToDownload := range filesWithByte {
		wg.Add(1)
		go func(fileToDownload *domain.FileWithBytes) {
			defer wg.Done()
			downloadErr := f.fileDownloader.Download(fileToDownload)
			if downloadErr != nil {
				once.Do(func() { err = downloadErr })
			}
			// wg.Done()

		}(fileToDownload)
	}
	wg.Wait()
	if err != nil {
		return
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
