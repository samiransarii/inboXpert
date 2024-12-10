// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: email_categorization_service.proto

package emailcategorization

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	EmailCategorizationService_CategorizeEmail_FullMethodName       = "/inboxpert.services.categorization.v1.EmailCategorizationService/CategorizeEmail"
	EmailCategorizationService_BatchCategorizeEmails_FullMethodName = "/inboxpert.services.categorization.v1.EmailCategorizationService/BatchCategorizeEmails"
)

// EmailCategorizationServiceClient is the client API for EmailCategorizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailCategorizationServiceClient interface {
	CategorizeEmail(ctx context.Context, in *CategorizeRequest, opts ...grpc.CallOption) (*CategorizeResponse, error)
	BatchCategorizeEmails(ctx context.Context, in *BatchCategorizeRequest, opts ...grpc.CallOption) (*BatchCategorizeResponse, error)
}

type emailCategorizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailCategorizationServiceClient(cc grpc.ClientConnInterface) EmailCategorizationServiceClient {
	return &emailCategorizationServiceClient{cc}
}

func (c *emailCategorizationServiceClient) CategorizeEmail(ctx context.Context, in *CategorizeRequest, opts ...grpc.CallOption) (*CategorizeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategorizeResponse)
	err := c.cc.Invoke(ctx, EmailCategorizationService_CategorizeEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailCategorizationServiceClient) BatchCategorizeEmails(ctx context.Context, in *BatchCategorizeRequest, opts ...grpc.CallOption) (*BatchCategorizeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchCategorizeResponse)
	err := c.cc.Invoke(ctx, EmailCategorizationService_BatchCategorizeEmails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailCategorizationServiceServer is the server API for EmailCategorizationService service.
// All implementations must embed UnimplementedEmailCategorizationServiceServer
// for forward compatibility.
type EmailCategorizationServiceServer interface {
	CategorizeEmail(context.Context, *CategorizeRequest) (*CategorizeResponse, error)
	BatchCategorizeEmails(context.Context, *BatchCategorizeRequest) (*BatchCategorizeResponse, error)
	mustEmbedUnimplementedEmailCategorizationServiceServer()
}

// UnimplementedEmailCategorizationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEmailCategorizationServiceServer struct{}

func (UnimplementedEmailCategorizationServiceServer) CategorizeEmail(context.Context, *CategorizeRequest) (*CategorizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategorizeEmail not implemented")
}
func (UnimplementedEmailCategorizationServiceServer) BatchCategorizeEmails(context.Context, *BatchCategorizeRequest) (*BatchCategorizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCategorizeEmails not implemented")
}
func (UnimplementedEmailCategorizationServiceServer) mustEmbedUnimplementedEmailCategorizationServiceServer() {
}
func (UnimplementedEmailCategorizationServiceServer) testEmbeddedByValue() {}

// UnsafeEmailCategorizationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailCategorizationServiceServer will
// result in compilation errors.
type UnsafeEmailCategorizationServiceServer interface {
	mustEmbedUnimplementedEmailCategorizationServiceServer()
}

func RegisterEmailCategorizationServiceServer(s grpc.ServiceRegistrar, srv EmailCategorizationServiceServer) {
	// If the following call pancis, it indicates UnimplementedEmailCategorizationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EmailCategorizationService_ServiceDesc, srv)
}

func _EmailCategorizationService_CategorizeEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailCategorizationServiceServer).CategorizeEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailCategorizationService_CategorizeEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailCategorizationServiceServer).CategorizeEmail(ctx, req.(*CategorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailCategorizationService_BatchCategorizeEmails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchCategorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailCategorizationServiceServer).BatchCategorizeEmails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailCategorizationService_BatchCategorizeEmails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailCategorizationServiceServer).BatchCategorizeEmails(ctx, req.(*BatchCategorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmailCategorizationService_ServiceDesc is the grpc.ServiceDesc for EmailCategorizationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailCategorizationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inboxpert.services.categorization.v1.EmailCategorizationService",
	HandlerType: (*EmailCategorizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CategorizeEmail",
			Handler:    _EmailCategorizationService_CategorizeEmail_Handler,
		},
		{
			MethodName: "BatchCategorizeEmails",
			Handler:    _EmailCategorizationService_BatchCategorizeEmails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email_categorization_service.proto",
}