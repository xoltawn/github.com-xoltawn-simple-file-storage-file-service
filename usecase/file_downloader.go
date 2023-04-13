package usecase

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/xoltawn/simple-file-storage-file-service/domain"
)

type fileDownloader struct {
}

// NewFileDownloader returns a new file downloader
func NewFileDownloader() *fileDownloader {
	return &fileDownloader{}
}

func (d *fileDownloader) Download(file *domain.FileWithBytes) (err error) {
	resp, err := http.Get(file.OriginalURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bytes := make([]byte, resp.ContentLength)
	_, err = resp.Body.Read(bytes)
	if err != nil {
		return err
	}

	(*file).Data = bytes
	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return err
	}

	(*file).FileSize = int64(size)
	_, extension, ok := strings.Cut(http.DetectContentType(bytes), "/")
	if !ok {
		return domain.ErrFileExtensionNotSupported
	}

	(*file).FileExtension = extension
	return
}
