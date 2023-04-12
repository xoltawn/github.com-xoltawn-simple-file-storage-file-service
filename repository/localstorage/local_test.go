package localstorage_test

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xoltawn/simple-file-storage-file-service/repository/localstorage"
)

const (
	path          = "../../tmp"
	recursivePath = "../../tmp/tmp1/tmp2"
)

func TestCreatePathIfNotExist(t *testing.T) {
	t.Run("if dir exists it simply exits without error", func(t *testing.T) {
		t.Run("non-recursive path", func(t *testing.T) {
			//arrange
			err := os.Mkdir(path, os.ModePerm)
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
			err = os.RemoveAll(recursivePath)
			if err != nil {
				log.Fatal(err)
			}
		})
	})
}
