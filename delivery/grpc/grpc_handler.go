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

	_, err = h.bytesToLinksConvertor.Convert(req.Links)
	if err != nil {
		return
	}

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
