// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hmq.proto

package hmq

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type QuerySubscribeRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=Topic,proto3" json:"Topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuerySubscribeRequest) Reset()         { *m = QuerySubscribeRequest{} }
func (m *QuerySubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySubscribeRequest) ProtoMessage()    {}
func (*QuerySubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_935f0990d4a84183, []int{0}
}

func (m *QuerySubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuerySubscribeRequest.Unmarshal(m, b)
}
func (m *QuerySubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuerySubscribeRequest.Marshal(b, m, deterministic)
}
func (m *QuerySubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySubscribeRequest.Merge(m, src)
}
func (m *QuerySubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_QuerySubscribeRequest.Size(m)
}
func (m *QuerySubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySubscribeRequest proto.InternalMessageInfo

func (m *QuerySubscribeRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

type QueryConnectRequest struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryConnectRequest) Reset()         { *m = QueryConnectRequest{} }
func (m *QueryConnectRequest) String() string { return proto.CompactTextString(m) }
func (*QueryConnectRequest) ProtoMessage()    {}
func (*QueryConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_935f0990d4a84183, []int{1}
}

func (m *QueryConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryConnectRequest.Unmarshal(m, b)
}
func (m *QueryConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryConnectRequest.Marshal(b, m, deterministic)
}
func (m *QueryConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryConnectRequest.Merge(m, src)
}
func (m *QueryConnectRequest) XXX_Size() int {
	return xxx_messageInfo_QueryConnectRequest.Size(m)
}
func (m *QueryConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryConnectRequest proto.InternalMessageInfo

func (m *QueryConnectRequest) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

type DeliverMessageRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=Topic,proto3" json:"Topic,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=Payload,proto3" json:"Payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeliverMessageRequest) Reset()         { *m = DeliverMessageRequest{} }
func (m *DeliverMessageRequest) String() string { return proto.CompactTextString(m) }
func (*DeliverMessageRequest) ProtoMessage()    {}
func (*DeliverMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_935f0990d4a84183, []int{2}
}

func (m *DeliverMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliverMessageRequest.Unmarshal(m, b)
}
func (m *DeliverMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliverMessageRequest.Marshal(b, m, deterministic)
}
func (m *DeliverMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliverMessageRequest.Merge(m, src)
}
func (m *DeliverMessageRequest) XXX_Size() int {
	return xxx_messageInfo_DeliverMessageRequest.Size(m)
}
func (m *DeliverMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliverMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeliverMessageRequest proto.InternalMessageInfo

func (m *DeliverMessageRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *DeliverMessageRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Response struct {
	RetCode              int32    `protobuf:"varint,2,opt,name=RetCode,proto3" json:"RetCode,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_935f0990d4a84183, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*QuerySubscribeRequest)(nil), "QuerySubscribeRequest")
	proto.RegisterType((*QueryConnectRequest)(nil), "QueryConnectRequest")
	proto.RegisterType((*DeliverMessageRequest)(nil), "DeliverMessageRequest")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("hmq.proto", fileDescriptor_935f0990d4a84183) }

var fileDescriptor_935f0990d4a84183 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x40, 0x5d, 0xa5, 0xda, 0x1d, 0x4a, 0x0f, 0xb1, 0x2d, 0xa1, 0xa7, 0x92, 0x53, 0x2f, 0x06,
	0xaa, 0xf4, 0xea, 0x65, 0x0b, 0xea, 0xa1, 0x60, 0x53, 0x7f, 0x60, 0x37, 0x1d, 0x34, 0xb0, 0x26,
	0xdb, 0x24, 0x5b, 0xe8, 0x47, 0xf9, 0x8f, 0xd2, 0xd8, 0x88, 0x81, 0xc5, 0xe3, 0x23, 0x6f, 0x32,
	0xe4, 0x05, 0xf2, 0x8f, 0xcf, 0x3d, 0x6f, 0xac, 0xf1, 0x86, 0xdd, 0xc1, 0x78, 0xd3, 0xa2, 0x3d,
	0x6e, 0xdb, 0xca, 0x49, 0xab, 0x2a, 0x14, 0xb8, 0x6f, 0xd1, 0x79, 0x32, 0x82, 0xde, 0x9b, 0x69,
	0x94, 0xa4, 0xd9, 0x2c, 0x9b, 0xe7, 0xe2, 0x07, 0xd8, 0x02, 0x6e, 0x83, 0x5e, 0x18, 0xad, 0x51,
	0xfa, 0x28, 0x4f, 0xa1, 0x5f, 0xd4, 0x0a, 0xb5, 0x7f, 0x59, 0x9d, 0xfd, 0x5f, 0x66, 0x4f, 0x30,
	0x5e, 0x61, 0xad, 0x0e, 0x68, 0xd7, 0xe8, 0x5c, 0xf9, 0xfe, 0xff, 0x06, 0x42, 0xe1, 0xe6, 0xb5,
	0x3c, 0xd6, 0xa6, 0xdc, 0xd1, 0xcb, 0x59, 0x36, 0x1f, 0x88, 0x88, 0xec, 0x11, 0xfa, 0x02, 0x5d,
	0x63, 0xb4, 0xc3, 0x93, 0x25, 0xd0, 0x17, 0x66, 0x87, 0xc1, 0xea, 0x89, 0x88, 0xa7, 0x93, 0xf3,
	0x1e, 0x7a, 0x15, 0xee, 0x8d, 0x78, 0xff, 0x95, 0x01, 0x3c, 0xaf, 0x37, 0x5b, 0xb4, 0x07, 0x25,
	0x91, 0x2c, 0x61, 0x98, 0xbe, 0x9c, 0x4c, 0x78, 0x67, 0x8a, 0x69, 0xce, 0xe3, 0x5e, 0x76, 0x41,
	0x16, 0x30, 0xf8, 0x5b, 0x80, 0x8c, 0x78, 0x47, 0x90, 0x74, 0x64, 0x09, 0xc3, 0xb4, 0x00, 0x99,
	0xf0, 0xce, 0x24, 0xc9, 0x58, 0x75, 0x1d, 0x7e, 0xe8, 0xe1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x0f,
	0x76, 0x58, 0x73, 0xae, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HMQServiceClient is the client API for HMQService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HMQServiceClient interface {
	QuerySubscribe(ctx context.Context, in *QuerySubscribeRequest, opts ...grpc.CallOption) (*Response, error)
	QueryConnect(ctx context.Context, in *QueryConnectRequest, opts ...grpc.CallOption) (*Response, error)
	DeliverMessage(ctx context.Context, in *DeliverMessageRequest, opts ...grpc.CallOption) (*Response, error)
}

type hMQServiceClient struct {
	cc *grpc.ClientConn
}

func NewHMQServiceClient(cc *grpc.ClientConn) HMQServiceClient {
	return &hMQServiceClient{cc}
}

func (c *hMQServiceClient) QuerySubscribe(ctx context.Context, in *QuerySubscribeRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/HMQService/QuerySubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hMQServiceClient) QueryConnect(ctx context.Context, in *QueryConnectRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/HMQService/QueryConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hMQServiceClient) DeliverMessage(ctx context.Context, in *DeliverMessageRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/HMQService/DeliverMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HMQServiceServer is the server API for HMQService service.
type HMQServiceServer interface {
	QuerySubscribe(context.Context, *QuerySubscribeRequest) (*Response, error)
	QueryConnect(context.Context, *QueryConnectRequest) (*Response, error)
	DeliverMessage(context.Context, *DeliverMessageRequest) (*Response, error)
}

// UnimplementedHMQServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHMQServiceServer struct {
}

func (*UnimplementedHMQServiceServer) QuerySubscribe(ctx context.Context, req *QuerySubscribeRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySubscribe not implemented")
}
func (*UnimplementedHMQServiceServer) QueryConnect(ctx context.Context, req *QueryConnectRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryConnect not implemented")
}
func (*UnimplementedHMQServiceServer) DeliverMessage(ctx context.Context, req *DeliverMessageRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeliverMessage not implemented")
}

func RegisterHMQServiceServer(s *grpc.Server, srv HMQServiceServer) {
	s.RegisterService(&_HMQService_serviceDesc, srv)
}

func _HMQService_QuerySubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HMQServiceServer).QuerySubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HMQService/QuerySubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HMQServiceServer).QuerySubscribe(ctx, req.(*QuerySubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HMQService_QueryConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HMQServiceServer).QueryConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HMQService/QueryConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HMQServiceServer).QueryConnect(ctx, req.(*QueryConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HMQService_DeliverMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliverMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HMQServiceServer).DeliverMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HMQService/DeliverMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HMQServiceServer).DeliverMessage(ctx, req.(*DeliverMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HMQService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "HMQService",
	HandlerType: (*HMQServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QuerySubscribe",
			Handler:    _HMQService_QuerySubscribe_Handler,
		},
		{
			MethodName: "QueryConnect",
			Handler:    _HMQService_QueryConnect_Handler,
		},
		{
			MethodName: "DeliverMessage",
			Handler:    _HMQService_DeliverMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hmq.proto",
}
