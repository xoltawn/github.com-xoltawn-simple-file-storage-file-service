package grpc_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	_grpc "github.com/xoltawn/simple-file-storage-file-service/delivery/grpc"
	_filepb "github.com/xoltawn/simple-file-storage-file-service/delivery/grpc/filepb"
	_mocks "github.com/xoltawn/simple-file-storage-file-service/domain/mocks"
)

var (
	sampleErr = errors.New("sample error")
)

func TestDownloadFromTextFile(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("if bytes to links convertor returns err, return err from handler", func(t *testing.T) {
		//arrange
		req := &_filepb.DownloadFromTextFileRequest{}

		bytesToLinksConvertor := _mocks.NewMockBytesToLinksConvertor(ctrl)
		bytesToLinksConvertor.EXPECT().Convert(gomock.Any()).Return(nil, sampleErr)

		//act
		sut := _grpc.NewFileGRPCHandler(nil, bytesToLinksConvertor)
		_, err := sut.DownloadFromTextFile(context.TODO(), req)

		//assert
		assert.Error(t, err)
		assert.Equal(t, sampleErr, err)
		// assert.Equal(t, expRes, res)
	})
	// t.Run("if usecase inside grpc handler returns err, return err from handler", func(t *testing.T) {
	// 	//arrange
	// 	req := &_filepb.DownloadFromTextFileRequest{}

	// 	fileUsecase := _mocks.NewMockFileUsecase(ctrl)
	// 	fileUsecase.EXPECT().SaveMutltipleFiles(context.TODO(), gomock.Any()).Return(sampleErr)

	// 	//act
	// 	sut := _grpc.NewFileGRPCHandler(fileUsecase)
	// 	_, err := sut.DownloadFromTextFile(context.TODO(), req)

	// 	//assert
	// 	assert.Error(t, err)
	// 	assert.Equal(t, sampleErr, err)
	// 	// assert.Equal(t, expRes, res)
	// })

}
