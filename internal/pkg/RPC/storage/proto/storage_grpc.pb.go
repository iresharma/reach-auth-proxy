// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: storage.proto

package proto

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

// FileServerPackageClient is the client API for FileServerPackage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServerPackageClient interface {
	InitializeFileServer(ctx context.Context, in *InitServerRequest, opts ...grpc.CallOption) (*InitServerResponse, error)
	GetServerInfo(ctx context.Context, in *GetServerRequest, opts ...grpc.CallOption) (*GetServerResponse, error)
	ListFile(ctx context.Context, in *ListFileRequest, opts ...grpc.CallOption) (*ListFileResponse, error)
	PreSignedGet(ctx context.Context, in *FileOperationRequest, opts ...grpc.CallOption) (*GetFileResponse, error)
	PreSignedDelete(ctx context.Context, in *FileOperationRequest, opts ...grpc.CallOption) (*OkResponse, error)
	PreSignedPut(ctx context.Context, in *FileOperationRequest, opts ...grpc.CallOption) (*GetFileResponse, error)
}

type fileServerPackageClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServerPackageClient(cc grpc.ClientConnInterface) FileServerPackageClient {
	return &fileServerPackageClient{cc}
}

func (c *fileServerPackageClient) InitializeFileServer(ctx context.Context, in *InitServerRequest, opts ...grpc.CallOption) (*InitServerResponse, error) {
	out := new(InitServerResponse)
	err := c.cc.Invoke(ctx, "/file_server_package.FileServerPackage/InitializeFileServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServerPackageClient) GetServerInfo(ctx context.Context, in *GetServerRequest, opts ...grpc.CallOption) (*GetServerResponse, error) {
	out := new(GetServerResponse)
	err := c.cc.Invoke(ctx, "/file_server_package.FileServerPackage/GetServerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServerPackageClient) ListFile(ctx context.Context, in *ListFileRequest, opts ...grpc.CallOption) (*ListFileResponse, error) {
	out := new(ListFileResponse)
	err := c.cc.Invoke(ctx, "/file_server_package.FileServerPackage/ListFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServerPackageClient) PreSignedGet(ctx context.Context, in *FileOperationRequest, opts ...grpc.CallOption) (*GetFileResponse, error) {
	out := new(GetFileResponse)
	err := c.cc.Invoke(ctx, "/file_server_package.FileServerPackage/PreSignedGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServerPackageClient) PreSignedDelete(ctx context.Context, in *FileOperationRequest, opts ...grpc.CallOption) (*OkResponse, error) {
	out := new(OkResponse)
	err := c.cc.Invoke(ctx, "/file_server_package.FileServerPackage/PreSignedDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServerPackageClient) PreSignedPut(ctx context.Context, in *FileOperationRequest, opts ...grpc.CallOption) (*GetFileResponse, error) {
	out := new(GetFileResponse)
	err := c.cc.Invoke(ctx, "/file_server_package.FileServerPackage/PreSignedPut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServerPackageServer is the server API for FileServerPackage service.
// All implementations must embed UnimplementedFileServerPackageServer
// for forward compatibility
type FileServerPackageServer interface {
	InitializeFileServer(context.Context, *InitServerRequest) (*InitServerResponse, error)
	GetServerInfo(context.Context, *GetServerRequest) (*GetServerResponse, error)
	ListFile(context.Context, *ListFileRequest) (*ListFileResponse, error)
	PreSignedGet(context.Context, *FileOperationRequest) (*GetFileResponse, error)
	PreSignedDelete(context.Context, *FileOperationRequest) (*OkResponse, error)
	PreSignedPut(context.Context, *FileOperationRequest) (*GetFileResponse, error)
	mustEmbedUnimplementedFileServerPackageServer()
}

// UnimplementedFileServerPackageServer must be embedded to have forward compatible implementations.
type UnimplementedFileServerPackageServer struct {
}

func (UnimplementedFileServerPackageServer) InitializeFileServer(context.Context, *InitServerRequest) (*InitServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitializeFileServer not implemented")
}
func (UnimplementedFileServerPackageServer) GetServerInfo(context.Context, *GetServerRequest) (*GetServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerInfo not implemented")
}
func (UnimplementedFileServerPackageServer) ListFile(context.Context, *ListFileRequest) (*ListFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFile not implemented")
}
func (UnimplementedFileServerPackageServer) PreSignedGet(context.Context, *FileOperationRequest) (*GetFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreSignedGet not implemented")
}
func (UnimplementedFileServerPackageServer) PreSignedDelete(context.Context, *FileOperationRequest) (*OkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreSignedDelete not implemented")
}
func (UnimplementedFileServerPackageServer) PreSignedPut(context.Context, *FileOperationRequest) (*GetFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreSignedPut not implemented")
}
func (UnimplementedFileServerPackageServer) mustEmbedUnimplementedFileServerPackageServer() {}

// UnsafeFileServerPackageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServerPackageServer will
// result in compilation errors.
type UnsafeFileServerPackageServer interface {
	mustEmbedUnimplementedFileServerPackageServer()
}

func RegisterFileServerPackageServer(s grpc.ServiceRegistrar, srv FileServerPackageServer) {
	s.RegisterService(&FileServerPackage_ServiceDesc, srv)
}

func _FileServerPackage_InitializeFileServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServerPackageServer).InitializeFileServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file_server_package.FileServerPackage/InitializeFileServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServerPackageServer).InitializeFileServer(ctx, req.(*InitServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileServerPackage_GetServerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServerPackageServer).GetServerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file_server_package.FileServerPackage/GetServerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServerPackageServer).GetServerInfo(ctx, req.(*GetServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileServerPackage_ListFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServerPackageServer).ListFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file_server_package.FileServerPackage/ListFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServerPackageServer).ListFile(ctx, req.(*ListFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileServerPackage_PreSignedGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServerPackageServer).PreSignedGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file_server_package.FileServerPackage/PreSignedGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServerPackageServer).PreSignedGet(ctx, req.(*FileOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileServerPackage_PreSignedDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServerPackageServer).PreSignedDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file_server_package.FileServerPackage/PreSignedDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServerPackageServer).PreSignedDelete(ctx, req.(*FileOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileServerPackage_PreSignedPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServerPackageServer).PreSignedPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file_server_package.FileServerPackage/PreSignedPut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServerPackageServer).PreSignedPut(ctx, req.(*FileOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileServerPackage_ServiceDesc is the grpc.ServiceDesc for FileServerPackage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileServerPackage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "file_server_package.FileServerPackage",
	HandlerType: (*FileServerPackageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitializeFileServer",
			Handler:    _FileServerPackage_InitializeFileServer_Handler,
		},
		{
			MethodName: "GetServerInfo",
			Handler:    _FileServerPackage_GetServerInfo_Handler,
		},
		{
			MethodName: "ListFile",
			Handler:    _FileServerPackage_ListFile_Handler,
		},
		{
			MethodName: "PreSignedGet",
			Handler:    _FileServerPackage_PreSignedGet_Handler,
		},
		{
			MethodName: "PreSignedDelete",
			Handler:    _FileServerPackage_PreSignedDelete_Handler,
		},
		{
			MethodName: "PreSignedPut",
			Handler:    _FileServerPackage_PreSignedPut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage.proto",
}
