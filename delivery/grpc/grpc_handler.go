package grpc

type fileGRPCHandler struct {
}

//go:generate mockgen --source=filepb/file_grpc.pb.go --destination=mocks/file_service_server.go . FileServiceServer
func NewFileGRPCHandler() *fileGRPCHandler {
	return &fileGRPCHandler{}
}
