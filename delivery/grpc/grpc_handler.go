package grpc

import (
	_filepb "github.com/xoltawn/simple-file-storage-file-service/delivery/grpc/filepb"
	"github.com/xoltawn/simple-file-storage-file-service/domain"

	"context"
)

type fileGRPCHandler struct {
	fileUsecase domain.FileUsecase
}

//go:generate mockgen --source=filepb/file_grpc.pb.go --destination=mocks/file_service_server.go . FileServiceServer
func NewFileGRPCHandler(fileUsecase domain.FileUsecase) *fileGRPCHandler {
	return &fileGRPCHandler{fileUsecase: fileUsecase}
}

func (h *fileGRPCHandler) DownloadFromTextFile(ctx context.Context, req *_filepb.DownloadFromTextFileRequest) (res *_filepb.DownloadFromTextFileResponse, err error) {
	filesWithByte := []*domain.FileWithBytes{}

	err = h.fileUsecase.SaveMutltipleFiles(ctx, filesWithByte)
	return
}
func (h *fileGRPCHandler) FetchFiles(context.Context, *_filepb.FetchFilesRequest) (res *_filepb.FetchFilesResponse, err error) {
	return
}
func (h *fileGRPCHandler) UploadFile(context.Context, *_filepb.UploadFileRequest) (res *_filepb.UploadFileResponse, err error) {
	return
}
