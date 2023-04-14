package grpc

import (
	"os"

	_filepb "github.com/xoltawn/simple-file-storage-file-service/delivery/grpc/filepb"
	"github.com/xoltawn/simple-file-storage-file-service/domain"

	"context"
)

type fileGRPCHandler struct {
	fileUsecase           domain.FileUsecase
	bytesToLinksConvertor domain.BytesToLinksConvertor
}

//go:generate mockgen --source=filepb/file_grpc.pb.go --destination=mocks/file_service_server.go . FileServiceServer
func NewFileGRPCHandler(fileUsecase domain.FileUsecase, bytesToLinksConvertor domain.BytesToLinksConvertor) *fileGRPCHandler {
	return &fileGRPCHandler{
		fileUsecase:           fileUsecase,
		bytesToLinksConvertor: bytesToLinksConvertor,
	}
}

func (h *fileGRPCHandler) DownloadFromTextFile(ctx context.Context, req *_filepb.DownloadFromTextFileRequest) (res *_filepb.DownloadFromTextFileResponse, err error) {

	links, err := h.bytesToLinksConvertor.Convert(req.Links)
	if err != nil {
		return &_filepb.DownloadFromTextFileResponse{}, err
	}

	filesWithByte := []*domain.FileWithBytes{}

	for _, l := range links {
		filesWithByte = append(filesWithByte, &domain.FileWithBytes{
			File: &domain.File{
				OriginalURL: l,
			},
		})
	}

	err = h.fileUsecase.SaveMutltipleFiles(ctx, filesWithByte)
	return &_filepb.DownloadFromTextFileResponse{}, err
}
func (h *fileGRPCHandler) FetchFiles(ctx context.Context, req *_filepb.FetchFilesRequest) (res *_filepb.FetchFilesResponse, err error) {
	files, err := h.fileUsecase.FetchFiles(ctx, int(req.Limit), int(req.Offset))
	if err != nil {
		return &_filepb.FetchFilesResponse{}, err
	}

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
	return &_filepb.FetchFilesResponse{
		Files: expFiles,
	}, err

}
func (h *fileGRPCHandler) UploadFile(ctx context.Context, req *_filepb.UploadFileRequest) (res *_filepb.UploadFileResponse, err error) {
	fileInfo := &domain.File{}
	err = h.fileUsecase.SaveFile(ctx, req.GetFile(), fileInfo, os.Getenv("USER_CONTENT_PATH"))
	if err != nil {
		return &_filepb.UploadFileResponse{}, err
	}
	return &_filepb.UploadFileResponse{
		File: &_filepb.File{
			OriginalUrl:   fileInfo.OriginalURL,
			LocalName:     fileInfo.LocalName,
			FileExtension: fileInfo.FileExtension,
			FileSize:      fileInfo.FileSize,
			CreatedAt:     fileInfo.CreatedAt,
		},
	}, err
}
