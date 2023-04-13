package usecase_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	_mocks "github.com/xoltawn/simple-file-storage-file-service/domain/mocks"
	"github.com/xoltawn/simple-file-storage-file-service/usecase"
)

var (
	sampleErr = errors.New("sample error")
)

func TestConvert(t *testing.T) {
	ctrl := gomock.NewController(t)
	data := []byte{}
	t.Run("if bytesToReader returns error, return err", func(t *testing.T) {
		//assert
		bytesToReaderConvertor := _mocks.NewMockBytesToReaderConvertor(ctrl)
		bytesToReaderConvertor.EXPECT().Convert(gomock.Any()).Return(nil, sampleErr)
		sut := usecase.NewBytesToLinksConvertor(bytesToReaderConvertor)

		//act
		_, err := sut.Convert(data)

		//assert
		assert.Equal(t, sampleErr, err)

	})
}
