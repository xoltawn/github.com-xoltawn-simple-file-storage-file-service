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

	t.Run("if no error occurres, retursn nil for error and the files requested", func(t *testing.T) {
		//arrange
		expFiles := []domain.File{
			{
				OriginalURL:   "OriginalUrl1",
				LocalName:     "LocalName1",
				FileExtension: "FileExtension1",
				FileSize:      1,
				CreatedAt:     "CreatedAt1",
			},
		}
		fileRepo := _mocks.NewMockFileRepository(ctrl)
		fileRepo.EXPECT().FetchFiles(context.TODO(), gomock.Any(), gomock.Any()).Return(expFiles, nil)

		//act
		sut := usecase.NewFileUsecase(nil, fileRepo, imagesPath)
		resFiles, err := sut.FetchFiles(context.TODO(), limit, offset)

		//assert
		assert.NoError(t, err)
		assert.Equal(t, expFiles, resFiles)
	})

}

func TestSaveMutltipleFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	filesWithBytes := []*domain.FileWithBytes{
		{
			File: domain.File{
				OriginalURL:   "OriginalUrl1",
				LocalName:     "LocalName1",
				FileExtension: "FileExtension1",
				FileSize:      1,
				CreatedAt:     "CreatedAt1",
			},
			Data: []byte{},
		},
		{
			File: domain.File{
				OriginalURL:   "OriginalUrl1",
				LocalName:     "LocalName1",
				FileExtension: "FileExtension1",
				FileSize:      1,
				CreatedAt:     "CreatedAt1",
			},
			Data: []byte{},
		},
	}

	// files := []*domain.File{
	// 	{
	// 		OriginalURL:   "OriginalUrl1",
	// 		LocalName:     "LocalName1",
	// 		FileExtension: "FileExtension1",
	// 		FileSize:      1,
	// 		CreatedAt:     "CreatedAt1",
	// 	},
	// 	{
	// 		OriginalURL:   "OriginalUrl2",
	// 		LocalName:     "LocalName2",
	// 		FileExtension: "FileExtension2",
	// 		FileSize:      2,
	// 		CreatedAt:     "CreatedAt2",
	// 	},
	// }
	t.Run("if saving file in file storage returns error, then return error", func(t *testing.T) {
		//arrange
		fileStorage := _mocks.NewMockFileStorage(ctrl)
		fileStorage.EXPECT().SaveFile(context.TODO(), gomock.Any(), gomock.Any(), gomock.Any()).Return(expErr)

		//act
		sut := usecase.NewFileUsecase(fileStorage, nil, imagesPath)
		err := sut.SaveMutltipleFiles(context.TODO(), filesWithBytes)

		//assert
		assert.Error(t, expErr, err)
	})

	t.Run("if saving file to storage is successful, SaveFile in storage must be called that times", func(t *testing.T) {
		t.Run("if SaveMutltipleFiles in repo has error then return err", func(t *testing.T) {
			t.Run("remove created files", func(t *testing.T) {
				//arrange
				fileStorage := _mocks.NewMockFileStorage(ctrl)
				fileStorage.EXPECT().SaveFile(context.TODO(), gomock.Any(), gomock.Any(), gomock.Any()).Times(len(filesWithBytes)).Return(nil)
				fileStorage.EXPECT().RemoveFiles(context.TODO(), gomock.Any()).Times(1).Return(nil)

				fileRepo := _mocks.NewMockFileRepository(ctrl)
				fileRepo.EXPECT().SaveMutltipleFiles(context.TODO(), gomock.Any()).Return(expErr)

				//act
				sut := usecase.NewFileUsecase(fileStorage, fileRepo, imagesPath)
				err := sut.SaveMutltipleFiles(context.TODO(), filesWithBytes)

				//assert
				assert.Error(t, err)
			})
		})

		t.Run("if SaveMutltipleFiles in repo has no error then return nil", func(t *testing.T) {
			//arrange
			fileStorage := _mocks.NewMockFileStorage(ctrl)
			fileStorage.EXPECT().SaveFile(context.TODO(), gomock.Any(), gomock.Any(), gomock.Any()).Times(len(filesWithBytes)).Return(nil)
			fileRepo := _mocks.NewMockFileRepository(ctrl)
			fileRepo.EXPECT().SaveMutltipleFiles(context.TODO(), gomock.Any()).Return(nil)

			//act
			sut := usecase.NewFileUsecase(fileStorage, fileRepo, imagesPath)
			err := sut.SaveMutltipleFiles(context.TODO(), filesWithBytes)

			//assert
			assert.NoError(t, err)
		})
	})

}
