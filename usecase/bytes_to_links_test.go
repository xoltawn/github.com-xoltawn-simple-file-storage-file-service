package usecase_test

import (
	"bytes"
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
	data := []byte(
		`http://placehold.co/600x400
https://placehold.co/600x400
+`,
	)
	t.Run("if bytesToReader returns error, return err", func(t *testing.T) {
		//assert
		bytesToReaderConvertor := _mocks.NewMockBytesToReaderConvertor(ctrl)
		bytesToReaderConvertor.EXPECT().Convert(gomock.Any()).Return(nil, sampleErr)
		sut := usecase.NewBytesToLinksConvertor(bytesToReaderConvertor, nil)

		//act
		_, err := sut.Convert(data)

		//assert
		assert.Equal(t, sampleErr, err)
	})

	t.Run("if each link is valid , it will be added to the result", func(t *testing.T) {
		//assert
		reader := bytes.NewReader(data)

		bytesToReaderConvertor := _mocks.NewMockBytesToReaderConvertor(ctrl)
		bytesToReaderConvertor.EXPECT().Convert(gomock.Any()).Return(reader, nil)
		linkValidator := usecase.NewLinkValidator()
		sut := usecase.NewBytesToLinksConvertor(bytesToReaderConvertor, linkValidator)

		//act
		links, err := sut.Convert(data)

		//assert
		assert.NoError(t, err)
		assert.Equal(t, 2, len(links))
	})
}
