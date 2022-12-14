// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: calculator.proto

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

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	CalculateSum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	DeterminePrimeFactors(ctx context.Context, in *PrimeFactorsRequest, opts ...grpc.CallOption) (CalculatorService_DeterminePrimeFactorsClient, error)
	CalculateAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_CalculateAverageClient, error)
	CalculateMax(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_CalculateMaxClient, error)
	// rpc call to understand error handling
	CalculateSqRt(ctx context.Context, in *SqRtRequest, opts ...grpc.CallOption) (*SqRtResponse, error)
	// rpc call to understand contexts with deadline and with values
	CalculateSumWithDeadLine(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) CalculateSum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/calculateSum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) DeterminePrimeFactors(ctx context.Context, in *PrimeFactorsRequest, opts ...grpc.CallOption) (CalculatorService_DeterminePrimeFactorsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[0], "/calculator.CalculatorService/determinePrimeFactors", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceDeterminePrimeFactorsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_DeterminePrimeFactorsClient interface {
	Recv() (*PrimeFactorsResponse, error)
	grpc.ClientStream
}

type calculatorServiceDeterminePrimeFactorsClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceDeterminePrimeFactorsClient) Recv() (*PrimeFactorsResponse, error) {
	m := new(PrimeFactorsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) CalculateAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_CalculateAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[1], "/calculator.CalculatorService/calculateAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceCalculateAverageClient{stream}
	return x, nil
}

type CalculatorService_CalculateAverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceCalculateAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceCalculateAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceCalculateAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) CalculateMax(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_CalculateMaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[2], "/calculator.CalculatorService/calculateMax", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceCalculateMaxClient{stream}
	return x, nil
}

type CalculatorService_CalculateMaxClient interface {
	Send(*MaxRequest) error
	Recv() (*MaxResponse, error)
	grpc.ClientStream
}

type calculatorServiceCalculateMaxClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceCalculateMaxClient) Send(m *MaxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceCalculateMaxClient) Recv() (*MaxResponse, error) {
	m := new(MaxResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) CalculateSqRt(ctx context.Context, in *SqRtRequest, opts ...grpc.CallOption) (*SqRtResponse, error) {
	out := new(SqRtResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/calculateSqRt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) CalculateSumWithDeadLine(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/calculateSumWithDeadLine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations must embed UnimplementedCalculatorServiceServer
// for forward compatibility
type CalculatorServiceServer interface {
	CalculateSum(context.Context, *SumRequest) (*SumResponse, error)
	DeterminePrimeFactors(*PrimeFactorsRequest, CalculatorService_DeterminePrimeFactorsServer) error
	CalculateAverage(CalculatorService_CalculateAverageServer) error
	CalculateMax(CalculatorService_CalculateMaxServer) error
	// rpc call to understand error handling
	CalculateSqRt(context.Context, *SqRtRequest) (*SqRtResponse, error)
	// rpc call to understand contexts with deadline and with values
	CalculateSumWithDeadLine(context.Context, *SumRequest) (*SumResponse, error)
	mustEmbedUnimplementedCalculatorServiceServer()
}

// UnimplementedCalculatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (UnimplementedCalculatorServiceServer) CalculateSum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateSum not implemented")
}
func (UnimplementedCalculatorServiceServer) DeterminePrimeFactors(*PrimeFactorsRequest, CalculatorService_DeterminePrimeFactorsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeterminePrimeFactors not implemented")
}
func (UnimplementedCalculatorServiceServer) CalculateAverage(CalculatorService_CalculateAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method CalculateAverage not implemented")
}
func (UnimplementedCalculatorServiceServer) CalculateMax(CalculatorService_CalculateMaxServer) error {
	return status.Errorf(codes.Unimplemented, "method CalculateMax not implemented")
}
func (UnimplementedCalculatorServiceServer) CalculateSqRt(context.Context, *SqRtRequest) (*SqRtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateSqRt not implemented")
}
func (UnimplementedCalculatorServiceServer) CalculateSumWithDeadLine(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateSumWithDeadLine not implemented")
}
func (UnimplementedCalculatorServiceServer) mustEmbedUnimplementedCalculatorServiceServer() {}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_CalculateSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).CalculateSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/calculateSum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).CalculateSum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_DeterminePrimeFactors_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeFactorsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).DeterminePrimeFactors(m, &calculatorServiceDeterminePrimeFactorsServer{stream})
}

type CalculatorService_DeterminePrimeFactorsServer interface {
	Send(*PrimeFactorsResponse) error
	grpc.ServerStream
}

type calculatorServiceDeterminePrimeFactorsServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceDeterminePrimeFactorsServer) Send(m *PrimeFactorsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_CalculateAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).CalculateAverage(&calculatorServiceCalculateAverageServer{stream})
}

type CalculatorService_CalculateAverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceCalculateAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceCalculateAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceCalculateAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_CalculateMax_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).CalculateMax(&calculatorServiceCalculateMaxServer{stream})
}

type CalculatorService_CalculateMaxServer interface {
	Send(*MaxResponse) error
	Recv() (*MaxRequest, error)
	grpc.ServerStream
}

type calculatorServiceCalculateMaxServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceCalculateMaxServer) Send(m *MaxResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceCalculateMaxServer) Recv() (*MaxRequest, error) {
	m := new(MaxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_CalculateSqRt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqRtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).CalculateSqRt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/calculateSqRt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).CalculateSqRt(ctx, req.(*SqRtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_CalculateSumWithDeadLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).CalculateSumWithDeadLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/calculateSumWithDeadLine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).CalculateSumWithDeadLine(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "calculateSum",
			Handler:    _CalculatorService_CalculateSum_Handler,
		},
		{
			MethodName: "calculateSqRt",
			Handler:    _CalculatorService_CalculateSqRt_Handler,
		},
		{
			MethodName: "calculateSumWithDeadLine",
			Handler:    _CalculatorService_CalculateSumWithDeadLine_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "determinePrimeFactors",
			Handler:       _CalculatorService_DeterminePrimeFactors_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "calculateAverage",
			Handler:       _CalculatorService_CalculateAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "calculateMax",
			Handler:       _CalculatorService_CalculateMax_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
