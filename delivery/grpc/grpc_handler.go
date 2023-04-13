package grpc

import (
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
		return
	}

	filesWithByte := []*domain.FileWithBytes{}

	for _, l := range links {
		filesWithByte = append(filesWithByte, &domain.FileWithBytes{
			File: domain.File{
				OriginalURL: l,
			},
		})
	}

	err = h.fileUsecase.SaveMutltipleFiles(ctx, filesWithByte)
	return
}
func (h *fileGRPCHandler) FetchFiles(context.Context, *_filepb.FetchFilesRequest) (res *_filepb.FetchFilesResponse, err error) {
	return
}
func (h *fileGRPCHandler) UploadFile(context.Context, *_filepb.UploadFileRequest) (res *_filepb.UploadFileResponse, err error) {
	return
}
