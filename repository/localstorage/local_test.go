package localstorage_test

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xoltawn/simple-file-storage-file-service/repository/localstorage"
)

const (
	path = "../../tmp"
)

func TestCreatePathIfNotExist(t *testing.T) {
	t.Run("if dir exists it simply exits without error", func(t *testing.T) {
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

	t.Run("if dir not exists it simply then we create it", func(t *testing.T) {
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
}
