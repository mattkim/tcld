// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/namespaceservice/v1/service.proto

package namespaceservice

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("api/namespaceservice/v1/service.proto", fileDescriptor_d746e5fd89aff5eb)
}

var fileDescriptor_d746e5fd89aff5eb = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xaf, 0x4f, 0x03, 0x31,
	0x14, 0xc7, 0x5b, 0x01, 0xa2, 0x21, 0xfc, 0xa8, 0x21, 0x19, 0xc9, 0x13, 0x24, 0x38, 0xe8, 0xb1,
	0x81, 0x1b, 0x06, 0x10, 0x18, 0x82, 0xd8, 0x82, 0xc1, 0x90, 0xdb, 0xed, 0x05, 0x4a, 0x76, 0xbb,
	0xd2, 0xf6, 0x4e, 0x21, 0xb0, 0x38, 0x34, 0x7f, 0x01, 0x7f, 0x0a, 0x72, 0x72, 0x92, 0x75, 0x06,
	0xb9, 0x3f, 0x81, 0xc0, 0xd6, 0x1b, 0x8c, 0x94, 0x0c, 0xdc, 0xe5, 0xde, 0xe7, 0xfb, 0x3e, 0xdf,
	0xa4, 0x8f, 0x6d, 0xc5, 0x4a, 0x46, 0xdd, 0x38, 0x45, 0xa3, 0xe2, 0x04, 0x0d, 0xea, 0x42, 0x26,
	0x18, 0x15, 0xd5, 0x68, 0xf2, 0x29, 0x94, 0xce, 0x6c, 0xc6, 0xd7, 0x63, 0x25, 0xc5, 0x2c, 0x26,
	0x8a, 0x6a, 0x45, 0x84, 0xf2, 0x1a, 0x6f, 0x73, 0x34, 0xf6, 0x52, 0xa3, 0x51, 0x59, 0xd7, 0x4c,
	0x16, 0xd5, 0x1e, 0x16, 0xd8, 0xea, 0x99, 0xc7, 0x9b, 0x63, 0x9c, 0xdf, 0xb1, 0xb5, 0x06, 0x5e,
	0x49, 0x63, 0x51, 0x97, 0x33, 0x5e, 0x15, 0x01, 0xa7, 0xf8, 0xc1, 0x36, 0xc6, 0xae, 0x4a, 0xed,
	0x2f, 0x91, 0x71, 0xad, 0x4d, 0xc2, 0x0d, 0x5b, 0x3e, 0x95, 0xc6, 0x96, 0x23, 0xc3, 0x45, 0x70,
	0xcf, 0x77, 0xd0, 0x7b, 0xa3, 0xb9, 0xf9, 0x52, 0x9a, 0xb2, 0xa5, 0x13, 0x9c, 0x8e, 0xf8, 0x76,
	0x70, 0xc5, 0x57, 0xcc, 0x0b, 0x77, 0xe6, 0xa4, 0x4b, 0x5d, 0xc1, 0x56, 0xce, 0x55, 0x3b, 0xb6,
	0x38, 0x35, 0x86, 0x4b, 0xcf, 0x90, 0x5e, 0xba, 0x3b, 0x7f, 0xa0, 0xf4, 0x3e, 0x51, 0xb6, 0xd1,
	0xc0, 0x8f, 0xcc, 0x71, 0x6e, 0x6c, 0x96, 0x36, 0x31, 0xd6, 0xc9, 0xf5, 0xa1, 0xb5, 0x5a, 0xb6,
	0x72, 0x8b, 0xbc, 0xfe, 0xcb, 0x8b, 0x05, 0x53, 0xbe, 0xd0, 0xc1, 0xff, 0xc2, 0xbe, 0xdc, 0xd1,
	0x4d, 0x6f, 0x00, 0xa4, 0x3f, 0x00, 0x32, 0x1a, 0x00, 0xbd, 0x77, 0x40, 0x9f, 0x1d, 0xd0, 0x17,
	0x07, 0xb4, 0xe7, 0x80, 0xbe, 0x3a, 0xa0, 0x6f, 0x0e, 0xc8, 0xc8, 0x01, 0x7d, 0x1c, 0x02, 0xe9,
	0x0d, 0x81, 0xf4, 0x87, 0x40, 0x2e, 0xf6, 0x6d, 0xaa, 0x74, 0x47, 0x24, 0x9d, 0x2c, 0x6f, 0x47,
	0x81, 0xeb, 0xaf, 0xcf, 0xfe, 0x6b, 0x2d, 0x7e, 0x9e, 0xff, 0xde, 0x7b, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc2, 0x8d, 0xce, 0x50, 0x70, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NamespaceServiceClient is the client API for NamespaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NamespaceServiceClient interface {
	// RegisterNamespace creates a new namespace on temporal cloud.
	RegisterNamespace(ctx context.Context, in *RegisterNamespaceRequest, opts ...grpc.CallOption) (*RegisterNamespaceResponse, error)
	// ListNamespaces lists the names of all known namespaces on temporal cloud.
	ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error)
	// GetNamespace describes the namespace in detail.
	GetNamespace(ctx context.Context, in *GetNamespaceRequest, opts ...grpc.CallOption) (*GetNamespaceResponse, error)
	// UpdateNamespace updates an existing namespace on temporal cloud.
	UpdateNamespace(ctx context.Context, in *UpdateNamespaceRequest, opts ...grpc.CallOption) (*UpdateNamespaceResponse, error)
	// RenameCustomSearchAttribute renames an existing custom search attribute for a given namespace on temporal cloud.
	RenameCustomSearchAttribute(ctx context.Context, in *RenameCustomSearchAttributeRequest, opts ...grpc.CallOption) (*RenameCustomSearchAttributeResponse, error)
}

type namespaceServiceClient struct {
	cc *grpc.ClientConn
}

func NewNamespaceServiceClient(cc *grpc.ClientConn) NamespaceServiceClient {
	return &namespaceServiceClient{cc}
}

func (c *namespaceServiceClient) RegisterNamespace(ctx context.Context, in *RegisterNamespaceRequest, opts ...grpc.CallOption) (*RegisterNamespaceResponse, error) {
	out := new(RegisterNamespaceResponse)
	err := c.cc.Invoke(ctx, "/api.namespaceservice.v1.NamespaceService/RegisterNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error) {
	out := new(ListNamespacesResponse)
	err := c.cc.Invoke(ctx, "/api.namespaceservice.v1.NamespaceService/ListNamespaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) GetNamespace(ctx context.Context, in *GetNamespaceRequest, opts ...grpc.CallOption) (*GetNamespaceResponse, error) {
	out := new(GetNamespaceResponse)
	err := c.cc.Invoke(ctx, "/api.namespaceservice.v1.NamespaceService/GetNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) UpdateNamespace(ctx context.Context, in *UpdateNamespaceRequest, opts ...grpc.CallOption) (*UpdateNamespaceResponse, error) {
	out := new(UpdateNamespaceResponse)
	err := c.cc.Invoke(ctx, "/api.namespaceservice.v1.NamespaceService/UpdateNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) RenameCustomSearchAttribute(ctx context.Context, in *RenameCustomSearchAttributeRequest, opts ...grpc.CallOption) (*RenameCustomSearchAttributeResponse, error) {
	out := new(RenameCustomSearchAttributeResponse)
	err := c.cc.Invoke(ctx, "/api.namespaceservice.v1.NamespaceService/RenameCustomSearchAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NamespaceServiceServer is the server API for NamespaceService service.
type NamespaceServiceServer interface {
	// RegisterNamespace creates a new namespace on temporal cloud.
	RegisterNamespace(context.Context, *RegisterNamespaceRequest) (*RegisterNamespaceResponse, error)
	// ListNamespaces lists the names of all known namespaces on temporal cloud.
	ListNamespaces(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error)
	// GetNamespace describes the namespace in detail.
	GetNamespace(context.Context, *GetNamespaceRequest) (*GetNamespaceResponse, error)
	// UpdateNamespace updates an existing namespace on temporal cloud.
	UpdateNamespace(context.Context, *UpdateNamespaceRequest) (*UpdateNamespaceResponse, error)
	// RenameCustomSearchAttribute renames an existing custom search attribute for a given namespace on temporal cloud.
	RenameCustomSearchAttribute(context.Context, *RenameCustomSearchAttributeRequest) (*RenameCustomSearchAttributeResponse, error)
}

// UnimplementedNamespaceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNamespaceServiceServer struct {
}

func (*UnimplementedNamespaceServiceServer) RegisterNamespace(ctx context.Context, req *RegisterNamespaceRequest) (*RegisterNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNamespace not implemented")
}
func (*UnimplementedNamespaceServiceServer) ListNamespaces(ctx context.Context, req *ListNamespacesRequest) (*ListNamespacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNamespaces not implemented")
}
func (*UnimplementedNamespaceServiceServer) GetNamespace(ctx context.Context, req *GetNamespaceRequest) (*GetNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNamespace not implemented")
}
func (*UnimplementedNamespaceServiceServer) UpdateNamespace(ctx context.Context, req *UpdateNamespaceRequest) (*UpdateNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNamespace not implemented")
}
func (*UnimplementedNamespaceServiceServer) RenameCustomSearchAttribute(ctx context.Context, req *RenameCustomSearchAttributeRequest) (*RenameCustomSearchAttributeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameCustomSearchAttribute not implemented")
}

func RegisterNamespaceServiceServer(s *grpc.Server, srv NamespaceServiceServer) {
	s.RegisterService(&_NamespaceService_serviceDesc, srv)
}

func _NamespaceService_RegisterNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).RegisterNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.namespaceservice.v1.NamespaceService/RegisterNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).RegisterNamespace(ctx, req.(*RegisterNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_ListNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNamespacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).ListNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.namespaceservice.v1.NamespaceService/ListNamespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).ListNamespaces(ctx, req.(*ListNamespacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_GetNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).GetNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.namespaceservice.v1.NamespaceService/GetNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).GetNamespace(ctx, req.(*GetNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_UpdateNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).UpdateNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.namespaceservice.v1.NamespaceService/UpdateNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).UpdateNamespace(ctx, req.(*UpdateNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_RenameCustomSearchAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameCustomSearchAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).RenameCustomSearchAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.namespaceservice.v1.NamespaceService/RenameCustomSearchAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).RenameCustomSearchAttribute(ctx, req.(*RenameCustomSearchAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NamespaceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.namespaceservice.v1.NamespaceService",
	HandlerType: (*NamespaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNamespace",
			Handler:    _NamespaceService_RegisterNamespace_Handler,
		},
		{
			MethodName: "ListNamespaces",
			Handler:    _NamespaceService_ListNamespaces_Handler,
		},
		{
			MethodName: "GetNamespace",
			Handler:    _NamespaceService_GetNamespace_Handler,
		},
		{
			MethodName: "UpdateNamespace",
			Handler:    _NamespaceService_UpdateNamespace_Handler,
		},
		{
			MethodName: "RenameCustomSearchAttribute",
			Handler:    _NamespaceService_RenameCustomSearchAttribute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/namespaceservice/v1/service.proto",
}
