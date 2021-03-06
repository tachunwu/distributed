// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package calvinv1

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

// SequencerServiceClient is the client API for SequencerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SequencerServiceClient interface {
	ReceiveNewTxn(ctx context.Context, in *ReceiveNewTxnRequest, opts ...grpc.CallOption) (*ReceiveNewTxnResponse, error)
}

type sequencerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSequencerServiceClient(cc grpc.ClientConnInterface) SequencerServiceClient {
	return &sequencerServiceClient{cc}
}

func (c *sequencerServiceClient) ReceiveNewTxn(ctx context.Context, in *ReceiveNewTxnRequest, opts ...grpc.CallOption) (*ReceiveNewTxnResponse, error) {
	out := new(ReceiveNewTxnResponse)
	err := c.cc.Invoke(ctx, "/calvin.v1.SequencerService/ReceiveNewTxn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SequencerServiceServer is the server API for SequencerService service.
// All implementations must embed UnimplementedSequencerServiceServer
// for forward compatibility
type SequencerServiceServer interface {
	ReceiveNewTxn(context.Context, *ReceiveNewTxnRequest) (*ReceiveNewTxnResponse, error)
	mustEmbedUnimplementedSequencerServiceServer()
}

// UnimplementedSequencerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSequencerServiceServer struct {
}

func (UnimplementedSequencerServiceServer) ReceiveNewTxn(context.Context, *ReceiveNewTxnRequest) (*ReceiveNewTxnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveNewTxn not implemented")
}
func (UnimplementedSequencerServiceServer) mustEmbedUnimplementedSequencerServiceServer() {}

// UnsafeSequencerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SequencerServiceServer will
// result in compilation errors.
type UnsafeSequencerServiceServer interface {
	mustEmbedUnimplementedSequencerServiceServer()
}

func RegisterSequencerServiceServer(s grpc.ServiceRegistrar, srv SequencerServiceServer) {
	s.RegisterService(&SequencerService_ServiceDesc, srv)
}

func _SequencerService_ReceiveNewTxn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveNewTxnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SequencerServiceServer).ReceiveNewTxn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calvin.v1.SequencerService/ReceiveNewTxn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SequencerServiceServer).ReceiveNewTxn(ctx, req.(*ReceiveNewTxnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SequencerService_ServiceDesc is the grpc.ServiceDesc for SequencerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SequencerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calvin.v1.SequencerService",
	HandlerType: (*SequencerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveNewTxn",
			Handler:    _SequencerService_ReceiveNewTxn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calvin/v1/sequencer.proto",
}
