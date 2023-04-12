package localstorage_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/xoltawn/simple-file-storage-file-service/repository/localstorage"
	_domain "github.com/xoltawn/simple-file-storage-sharedparts/domain"
)

const (
	path          = "../../tmp"
	recursivePath = "../../tmp/tmp2/tmp3"
)

func TestCreatePathIfNotExist(t *testing.T) {
	t.Run("if dir exists it simply exits without error", func(t *testing.T) {
		t.Run("recursive path", func(t *testing.T) {
			//arrange
			err := os.MkdirAll(recursivePath, os.ModePerm)
			if err != nil {
				log.Fatalln(err)
			}

			//act
			sut := localstorage.NewLocalStorage()
			err = sut.CreatePathIfNotExist(recursivePath)

			//assert
			assert.NoError(t, err)

			//tearup
			err = os.RemoveAll(recursivePath)
			if err != nil {
				log.Fatal(err)
			}

		})

		t.Run("non-recursive path", func(t *testing.T) {
			//arrange
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				log.Fatalln(err)
			}

			//act
			sut := localstorage.NewLocalStorage()
			err = sut.CreatePathIfNotExist(path)

			//assert
			assert.NoError(t, err)

			//tearup
			err = os.RemoveAll(path)
			if err != nil {
				log.Fatal(err)
			}
		})
	})

	t.Run("if dir not exists it simply then we create it", func(t *testing.T) {
		t.Run("non-recursive", func(t *testing.T) {
			//arrange
			err := os.RemoveAll(path)
			if err != nil {
				log.Fatal(err)
			}

			//act
			sut := localstorage.NewLocalStorage()
			err = sut.CreatePathIfNotExist(path)

			//assert
			assert.NoError(t, err)

			//tearup
			err = os.RemoveAll(path)
			if err != nil {
				log.Fatal(err)
			}

		})

		t.Run("recursive", func(t *testing.T) {
			//arrange
			err := os.RemoveAll(recursivePath)
			if err != nil {
				log.Fatal(err)
			}

			//act
			sut := localstorage.NewLocalStorage()
			err = sut.CreatePathIfNotExist(recursivePath)

			//assert
			assert.NoError(t, err)

			//tearup
			err = os.RemoveAll(path)
			if err != nil {
				log.Fatal(err)
			}
		})
	})
}

func TestSaveFile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		t.Run("non-recursive", func(t *testing.T) {
			//arrange
			fileBytes := []byte{}
			fileInto := &_domain.File{
				LocalName:     uuid.New().String(),
				FileExtension: "png",
			}

			//act
			sut := localstorage.NewLocalStorage()
			err := sut.SaveFile(context.TODO(), fileBytes, fileInto, path)

			//assert
			assert.NoError(t, err)

			//tearup
			err = os.RemoveAll(path)
			if err != nil {
				log.Fatal(err)
			}
		})

		t.Run("recursive", func(t *testing.T) {
			//arrange
			fileBytes := []byte{}
			fileInto := &_domain.File{
				LocalName:     uuid.New().String(),
				FileExtension: "png",
			}

			//act
			sut := localstorage.NewLocalStorage()
			err := sut.SaveFile(context.TODO(), fileBytes, fileInto, recursivePath)

			//assert
			assert.NoError(t, err)

			//tearup
			err = os.RemoveAll(path)
			if err != nil {
				log.Fatal(err)
			}
		})
	})
}
