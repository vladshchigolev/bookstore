// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: api/proto/bookstore.proto

package api

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

// BookStorageClient is the client API for BookStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookStorageClient interface {
	GetBooks(ctx context.Context, in *Author, opts ...grpc.CallOption) (*BooksSet, error)
	GetAuthors(ctx context.Context, in *Title, opts ...grpc.CallOption) (*AuthorsSet, error)
}

type bookStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewBookStorageClient(cc grpc.ClientConnInterface) BookStorageClient {
	return &bookStorageClient{cc}
}

func (c *bookStorageClient) GetBooks(ctx context.Context, in *Author, opts ...grpc.CallOption) (*BooksSet, error) {
	out := new(BooksSet)
	err := c.cc.Invoke(ctx, "/api.BookStorage/getBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookStorageClient) GetAuthors(ctx context.Context, in *Title, opts ...grpc.CallOption) (*AuthorsSet, error) {
	out := new(AuthorsSet)
	err := c.cc.Invoke(ctx, "/api.BookStorage/getAuthors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookStorageServer is the server API for BookStorage service.
// All implementations must embed UnimplementedBookStorageServer
// for forward compatibility
type BookStorageServer interface {
	GetBooks(context.Context, *Author) (*BooksSet, error)
	GetAuthors(context.Context, *Title) (*AuthorsSet, error)
	//mustEmbedUnimplementedBookStorageServer()
}

// UnimplementedBookStorageServer must be embedded to have forward compatible implementations.
type UnimplementedBookStorageServer struct {
}

func (UnimplementedBookStorageServer) GetBooks(context.Context, *Author) (*BooksSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooks not implemented")
}
func (UnimplementedBookStorageServer) GetAuthors(context.Context, *Title) (*AuthorsSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthors not implemented")
}
func (UnimplementedBookStorageServer) mustEmbedUnimplementedBookStorageServer() {}

// UnsafeBookStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookStorageServer will
// result in compilation errors.
type UnsafeBookStorageServer interface {
	mustEmbedUnimplementedBookStorageServer()
}

func RegisterBookStorageServer(s grpc.ServiceRegistrar, srv BookStorageServer) {
	s.RegisterService(&BookStorage_ServiceDesc, srv)
}

func _BookStorage_GetBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Author)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookStorageServer).GetBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BookStorage/getBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookStorageServer).GetBooks(ctx, req.(*Author))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookStorage_GetAuthors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Title)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookStorageServer).GetAuthors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BookStorage/getAuthors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookStorageServer).GetAuthors(ctx, req.(*Title))
	}
	return interceptor(ctx, in, info, handler)
}

// BookStorage_ServiceDesc is the grpc.ServiceDesc for BookStorage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookStorage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.BookStorage",
	HandlerType: (*BookStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getBooks",
			Handler:    _BookStorage_GetBooks_Handler,
		},
		{
			MethodName: "getAuthors",
			Handler:    _BookStorage_GetAuthors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/bookstore.proto",
}