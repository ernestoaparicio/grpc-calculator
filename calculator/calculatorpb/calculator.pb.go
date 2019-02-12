// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type InputNumbers struct {
	Number_1             int32    `protobuf:"varint,1,opt,name=number_1,json=number1,proto3" json:"number_1,omitempty"`
	Number_2             int32    `protobuf:"varint,2,opt,name=number_2,json=number2,proto3" json:"number_2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InputNumbers) Reset()         { *m = InputNumbers{} }
func (m *InputNumbers) String() string { return proto.CompactTextString(m) }
func (*InputNumbers) ProtoMessage()    {}
func (*InputNumbers) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{0}
}

func (m *InputNumbers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputNumbers.Unmarshal(m, b)
}
func (m *InputNumbers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputNumbers.Marshal(b, m, deterministic)
}
func (m *InputNumbers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputNumbers.Merge(m, src)
}
func (m *InputNumbers) XXX_Size() int {
	return xxx_messageInfo_InputNumbers.Size(m)
}
func (m *InputNumbers) XXX_DiscardUnknown() {
	xxx_messageInfo_InputNumbers.DiscardUnknown(m)
}

var xxx_messageInfo_InputNumbers proto.InternalMessageInfo

func (m *InputNumbers) GetNumber_1() int32 {
	if m != nil {
		return m.Number_1
	}
	return 0
}

func (m *InputNumbers) GetNumber_2() int32 {
	if m != nil {
		return m.Number_2
	}
	return 0
}

type CalculatorRequest struct {
	InputNumbers         *InputNumbers `protobuf:"bytes,1,opt,name=inputNumbers,proto3" json:"inputNumbers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CalculatorRequest) Reset()         { *m = CalculatorRequest{} }
func (m *CalculatorRequest) String() string { return proto.CompactTextString(m) }
func (*CalculatorRequest) ProtoMessage()    {}
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{1}
}

func (m *CalculatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorRequest.Unmarshal(m, b)
}
func (m *CalculatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorRequest.Marshal(b, m, deterministic)
}
func (m *CalculatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorRequest.Merge(m, src)
}
func (m *CalculatorRequest) XXX_Size() int {
	return xxx_messageInfo_CalculatorRequest.Size(m)
}
func (m *CalculatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorRequest proto.InternalMessageInfo

func (m *CalculatorRequest) GetInputNumbers() *InputNumbers {
	if m != nil {
		return m.InputNumbers
	}
	return nil
}

type CalculatorResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorResponse) Reset()         { *m = CalculatorResponse{} }
func (m *CalculatorResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatorResponse) ProtoMessage()    {}
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{2}
}

func (m *CalculatorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorResponse.Unmarshal(m, b)
}
func (m *CalculatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorResponse.Marshal(b, m, deterministic)
}
func (m *CalculatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorResponse.Merge(m, src)
}
func (m *CalculatorResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatorResponse.Size(m)
}
func (m *CalculatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorResponse proto.InternalMessageInfo

func (m *CalculatorResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*InputNumbers)(nil), "calculator.InputNumbers")
	proto.RegisterType((*CalculatorRequest)(nil), "calculator.CalculatorRequest")
	proto.RegisterType((*CalculatorResponse)(nil), "calculator.CalculatorResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_7f42938f8c8365cf)
}

var fileDescriptor_7f42938f8c8365cf = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0x4e, 0xcc, 0x49,
	0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x47, 0x30, 0x0b, 0x92, 0x90, 0x38, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0x5c, 0x08, 0x11, 0x25, 0x17, 0x2e, 0x1e, 0xcf, 0xbc, 0x82, 0xd2, 0x12,
	0xbf, 0xd2, 0xdc, 0xa4, 0xd4, 0xa2, 0x62, 0x21, 0x49, 0x2e, 0x8e, 0x3c, 0x30, 0x33, 0xde, 0x50,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x35, 0x88, 0x1d, 0xc2, 0x37, 0x44, 0x92, 0x32, 0x92, 0x60, 0x42,
	0x96, 0x32, 0x52, 0x0a, 0xe4, 0x12, 0x74, 0x86, 0x9b, 0x19, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c,
	0x22, 0x64, 0xc3, 0xc5, 0x93, 0x89, 0x64, 0x34, 0xd8, 0x38, 0x6e, 0x23, 0x09, 0x3d, 0x24, 0xf7,
	0x20, 0x5b, 0x1d, 0x84, 0xa2, 0x5a, 0x49, 0x87, 0x4b, 0x08, 0xd9, 0xc8, 0xe2, 0x82, 0xfc, 0xbc,
	0xe2, 0x54, 0x21, 0x31, 0x2e, 0xb6, 0xa2, 0xd4, 0xe2, 0xd2, 0x9c, 0x12, 0xa8, 0xe3, 0xa0, 0x3c,
	0xa3, 0x6e, 0x26, 0x64, 0x17, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0xf9, 0x72, 0x71,
	0x21, 0x04, 0x85, 0x64, 0x91, 0x6d, 0xc6, 0x70, 0xae, 0x94, 0x1c, 0x2e, 0x69, 0x88, 0xd5, 0x4a,
	0x0c, 0x42, 0x01, 0x5c, 0xbc, 0x08, 0x71, 0xc7, 0x94, 0x14, 0xca, 0x4d, 0x0c, 0x45, 0xf6, 0x64,
	0x70, 0x69, 0x52, 0x49, 0x51, 0x62, 0x72, 0x09, 0xc5, 0xc6, 0x3a, 0xf1, 0x45, 0xf1, 0x20, 0xa7,
	0x80, 0x24, 0x36, 0x70, 0xbc, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x43, 0xf0, 0xcd, 0x6e,
	0x23, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	//Unary
	Calculator(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
	CalculatorAdd(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
	CalculatorSubtract(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Calculator(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Calculator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) CalculatorAdd(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/CalculatorAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) CalculatorSubtract(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/CalculatorSubtract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	//Unary
	Calculator(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
	CalculatorAdd(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
	CalculatorSubtract(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Calculator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Calculator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Calculator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Calculator(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_CalculatorAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).CalculatorAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/CalculatorAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).CalculatorAdd(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_CalculatorSubtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).CalculatorSubtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/CalculatorSubtract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).CalculatorSubtract(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculator",
			Handler:    _CalculatorService_Calculator_Handler,
		},
		{
			MethodName: "CalculatorAdd",
			Handler:    _CalculatorService_CalculatorAdd_Handler,
		},
		{
			MethodName: "CalculatorSubtract",
			Handler:    _CalculatorService_CalculatorSubtract_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
