package domain

import (
	"context"
	"errors"
	"io"
)

var (
	ErrInsertingRecord           = errors.New("error inserting record")
	ErrFileExtensionNotSupported = errors.New("File Extension Not Supported")
)

// File is the domain objects for stored files
type File struct {
	ID uint `gorm:"primarykey"`
	//OriginalUrl indicates the url from which the file was downloaded(used when file is downloaded from a link)
	OriginalURL string `json:"original_url"`
	//LocalName is the name given on storing
	LocalName string `json:"local_name"`
	//FileExtension ...
	FileExtension string `json:"file_extension"`
	//FileSize ...
	FileSize int64 `json:"file_size"`
	//CreatedAt ...
	CreatedAt string `json:"created_at"`
}

// FileWithBytes contains File struct and also []byte of the file that contains the file content
type FileWithBytes struct {
	File
	Data []byte
}

// FileStorage provices interface for any type of file storages
//
//go:generate mockgen --destination=mocks/file_storage.go . FileStorage
type FileStorage interface {
	SaveFile(ctx context.Context, fileBytes []byte, fileInfo *File, path string) (err error)
	RemoveFiles(ctx context.Context, files []*File) (err error)
}

// FileRepository ...
//
//go:generate mockgen --destination=mocks/file_repository.go . FileRepository
type FileRepository interface {
	SaveFile(ctx context.Context, fileInfo *File) (err error)
	SaveMutltipleFiles(ctx context.Context, files []*File) (err error)
	FetchFiles(ctx context.Context, limit, offset int) (files []File, err error)
}

// FileUsecase ...
//
//go:generate mockgen --destination=mocks/file_usecase.go . FileUsecase
type FileUsecase interface {
	SaveFile(ctx context.Context, fileBytes []byte, fileInfo *File) (err error)
	SaveMutltipleFiles(ctx context.Context, files []*FileWithBytes) (err error)
	FetchFiles(ctx context.Context, limit, offset int) (files []File, err error)
}

// BytesToLinksConvertor is reponsible for converting the text file in the format of bytes to array of links to download
//
//go:generate mockgen --destination=mocks/bytes_to_links_convertor.go . BytesToLinksConvertor
type BytesToLinksConvertor interface {
	Convert([]byte) ([]string, error)
}

// LinkValidator is reponsible for checking whether the given text file is a valid link
//
//go:generate mockgen --destination=mocks/link_validator.go . LinkValidator
type LinkValidator interface {
	IsLink(txt string) (result bool)
}

// BytesToReaderConvertor is reponsible for converting []byte to io.Reader
//
//go:generate mockgen --destination=mocks/bytes_to_reader_convertor.go . BytesToReaderConvertor
type BytesToReaderConvertor interface {
	Convert([]byte) (reader io.Reader, err error)
}

// FileDownloader is reponsible for downloading the file from the given URL
// and update the needed info in the FileWithBytes struct (Data []byte, FileSize, FileExtension)
//
//go:generate mockgen --destination=mocks/file_downloader.go . FileDownloader
type FileDownloader interface {
	Download(*FileWithBytes) (err error)
}
