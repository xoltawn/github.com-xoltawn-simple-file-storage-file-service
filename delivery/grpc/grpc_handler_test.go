package grpc_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	_grpc "github.com/xoltawn/simple-file-storage-file-service/delivery/grpc"
	_filepb "github.com/xoltawn/simple-file-storage-file-service/delivery/grpc/filepb"
	"github.com/xoltawn/simple-file-storage-file-service/domain"
	_mocks "github.com/xoltawn/simple-file-storage-file-service/domain/mocks"
)

var (
	sampleErr = errors.New("sample error")
)

func TestDownloadFromTextFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	expRes := &_filepb.DownloadFromTextFileResponse{}

	t.Run("if bytes to links convertor returns err, return err from handler", func(t *testing.T) {
		//arrange
		req := &_filepb.DownloadFromTextFileRequest{}

		bytesToLinksConvertor := _mocks.NewMockBytesToLinksConvertor(ctrl)
		bytesToLinksConvertor.EXPECT().Convert(gomock.Any()).Return(nil, sampleErr)

		//act
		sut := _grpc.NewFileGRPCHandler(nil, bytesToLinksConvertor)
		res, err := sut.DownloadFromTextFile(context.TODO(), req)

		//assert
		assert.Error(t, err)
		assert.Equal(t, sampleErr, err)
		assert.Equal(t, expRes, res)
	})
	t.Run("if usecase inside grpc handler returns err, return err from handler", func(t *testing.T) {
		//arrange
		req := &_filepb.DownloadFromTextFileRequest{}

		bytesToLinksConvertor := _mocks.NewMockBytesToLinksConvertor(ctrl)
		bytesToLinksConvertor.EXPECT().Convert(gomock.Any()).Return(nil, nil)
		fileUsecase := _mocks.NewMockFileUsecase(ctrl)
		fileUsecase.EXPECT().SaveMutltipleFiles(context.TODO(), gomock.Any()).Return(sampleErr)

		//act
		sut := _grpc.NewFileGRPCHandler(fileUsecase, bytesToLinksConvertor)
		res, err := sut.DownloadFromTextFile(context.TODO(), req)

		//assert
		assert.Error(t, err)
		assert.Equal(t, sampleErr, err)
		assert.Equal(t, expRes, res)
	})

	t.Run("if no err occures, return empty response with nil err", func(t *testing.T) {
		//arrange
		req := &_filepb.DownloadFromTextFileRequest{}

		bytesToLinksConvertor := _mocks.NewMockBytesToLinksConvertor(ctrl)
		bytesToLinksConvertor.EXPECT().Convert(gomock.Any()).Return(nil, nil)
		fileUsecase := _mocks.NewMockFileUsecase(ctrl)
		fileUsecase.EXPECT().SaveMutltipleFiles(context.TODO(), gomock.Any()).Return(nil)

		//act
		sut := _grpc.NewFileGRPCHandler(fileUsecase, bytesToLinksConvertor)
		res, err := sut.DownloadFromTextFile(context.TODO(), req)

		//assert
		assert.NoError(t, err)
		assert.Equal(t, expRes, res)
	})

}

func TestFetchFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("if FetchFiles in usecase returns err, return err from handler with empty files in response", func(t *testing.T) {
		//arrange
		req := &_filepb.FetchFilesRequest{}
		expRes := &_filepb.FetchFilesResponse{}
		emptySliceOfFiles := make([]domain.File, 0)

		fileUsecase := _mocks.NewMockFileUsecase(ctrl)
		fileUsecase.EXPECT().FetchFiles(context.TODO(), gomock.Any(), gomock.Any()).Return(emptySliceOfFiles, sampleErr)

		//act
		sut := _grpc.NewFileGRPCHandler(fileUsecase, nil)
		res, err := sut.FetchFiles(context.TODO(), req)

		//assert
		assert.Error(t, err)
		assert.Equal(t, sampleErr, err)
		assert.Equal(t, expRes, res)
	})

	t.Run("if FetchFiles in usecase returns no err, return requested files in response with nil err", func(t *testing.T) {
		//arrange
		files := []domain.File{
			{
				OriginalURL:   "OriginalUrl1",
				LocalName:     "LocalName1",
				FileExtension: "gif",
				FileSize:      1,
				CreatedAt:     "CreatedAt1",
			},
			{
				OriginalURL:   "OriginalUrl2",
				LocalName:     "LocalName2",
				FileExtension: "png",
				FileSize:      2,
				CreatedAt:     "CreatedAt2",
			},
		}

		req := &_filepb.FetchFilesRequest{}
		expFiles := []*_filepb.File{}
		for _, file := range files {
			pfFile := &_filepb.File{
				OriginalUrl:   file.OriginalURL,
				LocalName:     file.LocalName,
				FileExtension: file.FileExtension,
				FileSize:      file.FileSize,
				CreatedAt:     file.CreatedAt,
			}
			expFiles = append(expFiles, pfFile)
		}
		expRes := &_filepb.FetchFilesResponse{
			Files: expFiles,
		}

		fileUsecase := _mocks.NewMockFileUsecase(ctrl)
		fileUsecase.EXPECT().FetchFiles(context.TODO(), gomock.Any(), gomock.Any()).Return(files, nil)

		//act
		sut := _grpc.NewFileGRPCHandler(fileUsecase, nil)
		res, err := sut.FetchFiles(context.TODO(), req)

		//assert
		assert.NoError(t, err)
		assert.Equal(t, expRes, res)
	})
}

func TestUploadFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("if SaveFile of usecase returns err, return empty res and err", func(t *testing.T) {
		//arrange
		req := &_filepb.UploadFileRequest{}
		expRes := &_filepb.UploadFileResponse{}

		fileUsecase := _mocks.NewMockFileUsecase(ctrl)
		fileUsecase.EXPECT().SaveFile(context.TODO(), gomock.Any(), gomock.Any(), gomock.Any()).Return(sampleErr)

		//act
		sut := _grpc.NewFileGRPCHandler(fileUsecase, nil)
		res, err := sut.UploadFile(context.TODO(), req)

		//assert
		assert.Error(t, err)
		assert.Equal(t, sampleErr, err)
		assert.Equal(t, expRes, res)
	})

	t.Run("if SaveFile of usecase returns no err, return file info in res and nil err", func(t *testing.T) {
		//arrange
		pfFile := &_filepb.File{}
		req := &_filepb.UploadFileRequest{}
		expRes := &_filepb.UploadFileResponse{
			File: pfFile,
		}

		fileUsecase := _mocks.NewMockFileUsecase(ctrl)
		fileUsecase.EXPECT().SaveFile(context.TODO(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		//act
		sut := _grpc.NewFileGRPCHandler(fileUsecase, nil)
		res, err := sut.UploadFile(context.TODO(), req)

		//assert
		assert.NoError(t, err)
		assert.Equal(t, expRes, res)
	})

}
