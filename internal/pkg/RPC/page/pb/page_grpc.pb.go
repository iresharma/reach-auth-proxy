// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: page.proto

package pb

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

// PagePackageClient is the client API for PagePackage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PagePackageClient interface {
	GetPage(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Page, error)
	CreatePage(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*Page, error)
	CreateTemplate(ctx context.Context, in *TemplateRequest, opts ...grpc.CallOption) (*Template, error)
	UpdateTemplate(ctx context.Context, in *IdResponse, opts ...grpc.CallOption) (*VoidResponse, error)
	CreateLink(ctx context.Context, in *CreateLinkRequest, opts ...grpc.CallOption) (*PageLinks, error)
	UpdateLink(ctx context.Context, in *PageLinks, opts ...grpc.CallOption) (*VoidResponse, error)
	CreateMetaLink(ctx context.Context, in *Meta, opts ...grpc.CallOption) (*Meta, error)
	UpdateMetaLink(ctx context.Context, in *Meta, opts ...grpc.CallOption) (*VoidResponse, error)
}

type pagePackageClient struct {
	cc grpc.ClientConnInterface
}

func NewPagePackageClient(cc grpc.ClientConnInterface) PagePackageClient {
	return &pagePackageClient{cc}
}

func (c *pagePackageClient) GetPage(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Page, error) {
	out := new(Page)
	err := c.cc.Invoke(ctx, "/page_package.PagePackage/GetPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pagePackageClient) CreatePage(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*Page, error) {
	out := new(Page)
	err := c.cc.Invoke(ctx, "/page_package.PagePackage/CreatePage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pagePackageClient) CreateTemplate(ctx context.Context, in *TemplateRequest, opts ...grpc.CallOption) (*Template, error) {
	out := new(Template)
	err := c.cc.Invoke(ctx, "/page_package.PagePackage/CreateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pagePackageClient) UpdateTemplate(ctx context.Context, in *IdResponse, opts ...grpc.CallOption) (*VoidResponse, error) {
	out := new(VoidResponse)
	err := c.cc.Invoke(ctx, "/page_package.PagePackage/UpdateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pagePackageClient) CreateLink(ctx context.Context, in *CreateLinkRequest, opts ...grpc.CallOption) (*PageLinks, error) {
	out := new(PageLinks)
	err := c.cc.Invoke(ctx, "/page_package.PagePackage/CreateLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pagePackageClient) UpdateLink(ctx context.Context, in *PageLinks, opts ...grpc.CallOption) (*VoidResponse, error) {
	out := new(VoidResponse)
	err := c.cc.Invoke(ctx, "/page_package.PagePackage/UpdateLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pagePackageClient) CreateMetaLink(ctx context.Context, in *Meta, opts ...grpc.CallOption) (*Meta, error) {
	out := new(Meta)
	err := c.cc.Invoke(ctx, "/page_package.PagePackage/CreateMetaLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pagePackageClient) UpdateMetaLink(ctx context.Context, in *Meta, opts ...grpc.CallOption) (*VoidResponse, error) {
	out := new(VoidResponse)
	err := c.cc.Invoke(ctx, "/page_package.PagePackage/UpdateMetaLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PagePackageServer is the server API for PagePackage service.
// All implementations must embed UnimplementedPagePackageServer
// for forward compatibility
type PagePackageServer interface {
	GetPage(context.Context, *IdRequest) (*Page, error)
	CreatePage(context.Context, *PageRequest) (*Page, error)
	CreateTemplate(context.Context, *TemplateRequest) (*Template, error)
	UpdateTemplate(context.Context, *IdResponse) (*VoidResponse, error)
	CreateLink(context.Context, *CreateLinkRequest) (*PageLinks, error)
	UpdateLink(context.Context, *PageLinks) (*VoidResponse, error)
	CreateMetaLink(context.Context, *Meta) (*Meta, error)
	UpdateMetaLink(context.Context, *Meta) (*VoidResponse, error)
	mustEmbedUnimplementedPagePackageServer()
}

// UnimplementedPagePackageServer must be embedded to have forward compatible implementations.
type UnimplementedPagePackageServer struct {
}

func (UnimplementedPagePackageServer) GetPage(context.Context, *IdRequest) (*Page, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPage not implemented")
}
func (UnimplementedPagePackageServer) CreatePage(context.Context, *PageRequest) (*Page, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePage not implemented")
}
func (UnimplementedPagePackageServer) CreateTemplate(context.Context, *TemplateRequest) (*Template, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemplate not implemented")
}
func (UnimplementedPagePackageServer) UpdateTemplate(context.Context, *IdResponse) (*VoidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTemplate not implemented")
}
func (UnimplementedPagePackageServer) CreateLink(context.Context, *CreateLinkRequest) (*PageLinks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLink not implemented")
}
func (UnimplementedPagePackageServer) UpdateLink(context.Context, *PageLinks) (*VoidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLink not implemented")
}
func (UnimplementedPagePackageServer) CreateMetaLink(context.Context, *Meta) (*Meta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMetaLink not implemented")
}
func (UnimplementedPagePackageServer) UpdateMetaLink(context.Context, *Meta) (*VoidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMetaLink not implemented")
}
func (UnimplementedPagePackageServer) mustEmbedUnimplementedPagePackageServer() {}

// UnsafePagePackageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PagePackageServer will
// result in compilation errors.
type UnsafePagePackageServer interface {
	mustEmbedUnimplementedPagePackageServer()
}

func RegisterPagePackageServer(s grpc.ServiceRegistrar, srv PagePackageServer) {
	s.RegisterService(&PagePackage_ServiceDesc, srv)
}

func _PagePackage_GetPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PagePackageServer).GetPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/page_package.PagePackage/GetPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PagePackageServer).GetPage(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PagePackage_CreatePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PagePackageServer).CreatePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/page_package.PagePackage/CreatePage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PagePackageServer).CreatePage(ctx, req.(*PageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PagePackage_CreateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PagePackageServer).CreateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/page_package.PagePackage/CreateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PagePackageServer).CreateTemplate(ctx, req.(*TemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PagePackage_UpdateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PagePackageServer).UpdateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/page_package.PagePackage/UpdateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PagePackageServer).UpdateTemplate(ctx, req.(*IdResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _PagePackage_CreateLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PagePackageServer).CreateLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/page_package.PagePackage/CreateLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PagePackageServer).CreateLink(ctx, req.(*CreateLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PagePackage_UpdateLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageLinks)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PagePackageServer).UpdateLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/page_package.PagePackage/UpdateLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PagePackageServer).UpdateLink(ctx, req.(*PageLinks))
	}
	return interceptor(ctx, in, info, handler)
}

func _PagePackage_CreateMetaLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Meta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PagePackageServer).CreateMetaLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/page_package.PagePackage/CreateMetaLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PagePackageServer).CreateMetaLink(ctx, req.(*Meta))
	}
	return interceptor(ctx, in, info, handler)
}

func _PagePackage_UpdateMetaLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Meta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PagePackageServer).UpdateMetaLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/page_package.PagePackage/UpdateMetaLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PagePackageServer).UpdateMetaLink(ctx, req.(*Meta))
	}
	return interceptor(ctx, in, info, handler)
}

// PagePackage_ServiceDesc is the grpc.ServiceDesc for PagePackage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PagePackage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "page_package.PagePackage",
	HandlerType: (*PagePackageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPage",
			Handler:    _PagePackage_GetPage_Handler,
		},
		{
			MethodName: "CreatePage",
			Handler:    _PagePackage_CreatePage_Handler,
		},
		{
			MethodName: "CreateTemplate",
			Handler:    _PagePackage_CreateTemplate_Handler,
		},
		{
			MethodName: "UpdateTemplate",
			Handler:    _PagePackage_UpdateTemplate_Handler,
		},
		{
			MethodName: "CreateLink",
			Handler:    _PagePackage_CreateLink_Handler,
		},
		{
			MethodName: "UpdateLink",
			Handler:    _PagePackage_UpdateLink_Handler,
		},
		{
			MethodName: "CreateMetaLink",
			Handler:    _PagePackage_CreateMetaLink_Handler,
		},
		{
			MethodName: "UpdateMetaLink",
			Handler:    _PagePackage_UpdateMetaLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "page.proto",
}
