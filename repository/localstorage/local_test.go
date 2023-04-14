package localstorage_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	domain "github.com/xoltawn/simple-file-storage-file-service/domain"
	"github.com/xoltawn/simple-file-storage-file-service/repository/localstorage"
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
			sut := localstorage.NewLocalStorage(path)
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
			sut := localstorage.NewLocalStorage(path)
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
			sut := localstorage.NewLocalStorage(path)
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
			sut := localstorage.NewLocalStorage(path)
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
			fileInto := &domain.File{
				LocalName:     uuid.New().String(),
				FileExtension: "png",
			}

			//act
			sut := localstorage.NewLocalStorage(path)
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
			fileInto := &domain.File{
				LocalName:     uuid.New().String(),
				FileExtension: "png",
			}

			//act
			sut := localstorage.NewLocalStorage(path)
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

func TestRemoveFiles(t *testing.T) {
	//arrange
	sut := localstorage.NewLocalStorage(path)
	filesToDelete := []*domain.File{}
	filesWithBytes := []*domain.FileWithBytes{
		{
			File: &domain.File{
				OriginalURL:   "OriginalUrl1",
				LocalName:     "LocalName1",
				FileExtension: "png",
				FileSize:      1,
				CreatedAt:     "CreatedAt1",
			},
			Data: []byte{},
		},
		{
			File: &domain.File{
				OriginalURL:   "OriginalUrl1",
				LocalName:     "LocalName1",
				FileExtension: "gif",
				FileSize:      1,
				CreatedAt:     "CreatedAt1",
			},
			Data: []byte{},
		},
	}
	for _, f := range filesWithBytes {
		err := sut.SaveFile(context.TODO(), f.Data, (*f).File, path)
		if err != nil {
			log.Fatal(err)
		}

		filesToDelete = append(filesToDelete, f.File)
	}
	//act
	err := sut.RemoveFiles(context.TODO(), filesToDelete)

	//assert
	assert.NoError(t, err)
	for _, fd := range filesToDelete {
		assert.NoFileExists(t, fmt.Sprint(path, "/", fd.LocalName, ".", fd.FileExtension))
	}

	//tearup
	err = os.RemoveAll(path)
	if err != nil {
		log.Fatal(err)
	}

}
