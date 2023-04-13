// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package filepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	DownloadFromTextFile(ctx context.Context, in *DownloadFromTextFileRequest, opts ...grpc.CallOption) (*DownloadFromTextFileResponse, error)
	FetchFiles(ctx context.Context, in *FetchFilesRequest, opts ...grpc.CallOption) (*FetchFilesResponse, error)
	UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileResponse, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) DownloadFromTextFile(ctx context.Context, in *DownloadFromTextFileRequest, opts ...grpc.CallOption) (*DownloadFromTextFileResponse, error) {
	out := new(DownloadFromTextFileResponse)
	err := c.cc.Invoke(ctx, "/file.FileService/DownloadFromTextFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) FetchFiles(ctx context.Context, in *FetchFilesRequest, opts ...grpc.CallOption) (*FetchFilesResponse, error) {
	out := new(FetchFilesResponse)
	err := c.cc.Invoke(ctx, "/file.FileService/FetchFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileResponse, error) {
	out := new(UploadFileResponse)
	err := c.cc.Invoke(ctx, "/file.FileService/UploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations should embed UnimplementedFileServiceServer
// for forward compatibility
type FileServiceServer interface {
	DownloadFromTextFile(context.Context, *DownloadFromTextFileRequest) (*DownloadFromTextFileResponse, error)
	FetchFiles(context.Context, *FetchFilesRequest) (*FetchFilesResponse, error)
	UploadFile(context.Context, *UploadFileRequest) (*UploadFileResponse, error)
}

// UnimplementedFileServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (UnimplementedFileServiceServer) DownloadFromTextFile(context.Context, *DownloadFromTextFileRequest) (*DownloadFromTextFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadFromTextFile not implemented")
}
func (UnimplementedFileServiceServer) FetchFiles(context.Context, *FetchFilesRequest) (*FetchFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchFiles not implemented")
}
func (UnimplementedFileServiceServer) UploadFile(context.Context, *UploadFileRequest) (*UploadFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_DownloadFromTextFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadFromTextFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).DownloadFromTextFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.FileService/DownloadFromTextFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).DownloadFromTextFile(ctx, req.(*DownloadFromTextFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_FetchFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).FetchFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.FileService/FetchFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).FetchFiles(ctx, req.(*FetchFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.FileService/UploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).UploadFile(ctx, req.(*UploadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "file.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DownloadFromTextFile",
			Handler:    _FileService_DownloadFromTextFile_Handler,
		},
		{
			MethodName: "FetchFiles",
			Handler:    _FileService_FetchFiles_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _FileService_UploadFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "delivery/grpc/file.proto",
}
