package usecase_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/xoltawn/simple-file-storage-file-service/domain"
	_mocks "github.com/xoltawn/simple-file-storage-file-service/domain/mocks"
	"github.com/xoltawn/simple-file-storage-file-service/usecase"
	"golang.org/x/net/context"
)

var (
	expErr     = errors.New("error")
	imagesPath = "images"
)

func TestSaveFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	fileBytes := []byte{}
	file := domain.File{}
	t.Run("if saving file in file storage returns error, then return error", func(t *testing.T) {
		//arrange
		fileStorage := _mocks.NewMockFileStorage(ctrl)
		fileStorage.EXPECT().SaveFile(context.TODO(), gomock.Any(), gomock.Any(), gomock.Any()).Return(expErr)

		//act
		sut := usecase.NewFileUsecase(fileStorage, nil, imagesPath)
		err := sut.SaveFile(context.TODO(), fileBytes, &file)

		//assert
		assert.Error(t, expErr, err)
	})

	t.Run("if saving file inf in file repository returns error, then return error", func(t *testing.T) {
		//arrange
		fileStorage := _mocks.NewMockFileStorage(ctrl)
		fileStorage.EXPECT().SaveFile(context.TODO(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		fileRepo := _mocks.NewMockFileRepository(ctrl)
		fileRepo.EXPECT().SaveFile(context.TODO(), gomock.Any()).Return(expErr)

		//act
		sut := usecase.NewFileUsecase(fileStorage, fileRepo, imagesPath)
		err := sut.SaveFile(context.TODO(), fileBytes, &file)

		//assert
		assert.Error(t, expErr, err)
	})

	t.Run("if no error occurres, then return nil", func(t *testing.T) {
		//arrange
		fileStorage := _mocks.NewMockFileStorage(ctrl)
		fileStorage.EXPECT().SaveFile(context.TODO(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		fileRepo := _mocks.NewMockFileRepository(ctrl)
		fileRepo.EXPECT().SaveFile(context.TODO(), gomock.Any()).Return(nil)

		//act
		sut := usecase.NewFileUsecase(fileStorage, fileRepo, imagesPath)
		err := sut.SaveFile(context.TODO(), fileBytes, &file)

		//assert
		assert.NoError(t, err)
	})
}

func TestFetchFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	limit := 10
	offset := 0
	t.Run("if fetching file in file repo returns error, then return error with no files", func(t *testing.T) {
		//arrange
		fileRepo := _mocks.NewMockFileRepository(ctrl)
		fileRepo.EXPECT().FetchFiles(context.TODO(), gomock.Any(), gomock.Any()).Return([]domain.File{}, expErr)

		//act
		sut := usecase.NewFileUsecase(nil, fileRepo, imagesPath)
		resFiles, err := sut.FetchFiles(context.TODO(), limit, offset)

		//assert
		assert.Error(t, expErr, err)
		assert.Empty(t, resFiles)
	})

}
