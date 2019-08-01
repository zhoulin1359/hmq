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
	Qos                  int32    `protobuf:"varint,2,opt,name=Qos,proto3" json:"Qos,omitempty"`
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

func (m *QuerySubscribeRequest) GetQos() int32 {
	if m != nil {
		return m.Qos
	}
	return 0
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
	Qos                  int32    `protobuf:"varint,3,opt,name=Qos,proto3" json:"Qos,omitempty"`
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

func (m *DeliverMessageRequest) GetQos() int32 {
	if m != nil {
		return m.Qos
	}
	return 0
}

type QueryShareSubscribeRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=Topic,proto3" json:"Topic,omitempty"`
	Qos                  int32    `protobuf:"varint,2,opt,name=Qos,proto3" json:"Qos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryShareSubscribeRequest) Reset()         { *m = QueryShareSubscribeRequest{} }
func (m *QueryShareSubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryShareSubscribeRequest) ProtoMessage()    {}
func (*QueryShareSubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_935f0990d4a84183, []int{3}
}

func (m *QueryShareSubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryShareSubscribeRequest.Unmarshal(m, b)
}
func (m *QueryShareSubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryShareSubscribeRequest.Marshal(b, m, deterministic)
}
func (m *QueryShareSubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryShareSubscribeRequest.Merge(m, src)
}
func (m *QueryShareSubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_QueryShareSubscribeRequest.Size(m)
}
func (m *QueryShareSubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryShareSubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryShareSubscribeRequest proto.InternalMessageInfo

func (m *QueryShareSubscribeRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *QueryShareSubscribeRequest) GetQos() int32 {
	if m != nil {
		return m.Qos
	}
	return 0
}

type Response struct {
	RetCode              int32    `protobuf:"varint,1,opt,name=RetCode,proto3" json:"RetCode,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_935f0990d4a84183, []int{4}
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

type ShareSubscribeResponse struct {
	RetCode              int32    `protobuf:"varint,1,opt,name=RetCode,proto3" json:"RetCode,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	ShareSubCount        int32    `protobuf:"varint,3,opt,name=ShareSubCount,proto3" json:"ShareSubCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShareSubscribeResponse) Reset()         { *m = ShareSubscribeResponse{} }
func (m *ShareSubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*ShareSubscribeResponse) ProtoMessage()    {}
func (*ShareSubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_935f0990d4a84183, []int{5}
}

func (m *ShareSubscribeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareSubscribeResponse.Unmarshal(m, b)
}
func (m *ShareSubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareSubscribeResponse.Marshal(b, m, deterministic)
}
func (m *ShareSubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareSubscribeResponse.Merge(m, src)
}
func (m *ShareSubscribeResponse) XXX_Size() int {
	return xxx_messageInfo_ShareSubscribeResponse.Size(m)
}
func (m *ShareSubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareSubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShareSubscribeResponse proto.InternalMessageInfo

func (m *ShareSubscribeResponse) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *ShareSubscribeResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ShareSubscribeResponse) GetShareSubCount() int32 {
	if m != nil {
		return m.ShareSubCount
	}
	return 0
}

func init() {
	proto.RegisterType((*QuerySubscribeRequest)(nil), "QuerySubscribeRequest")
	proto.RegisterType((*QueryConnectRequest)(nil), "QueryConnectRequest")
	proto.RegisterType((*DeliverMessageRequest)(nil), "DeliverMessageRequest")
	proto.RegisterType((*QueryShareSubscribeRequest)(nil), "QueryShareSubscribeRequest")
	proto.RegisterType((*Response)(nil), "Response")
	proto.RegisterType((*ShareSubscribeResponse)(nil), "ShareSubscribeResponse")
}

func init() { proto.RegisterFile("hmq.proto", fileDescriptor_935f0990d4a84183) }

var fileDescriptor_935f0990d4a84183 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4f, 0xc2, 0x30,
	0x14, 0xc6, 0x99, 0x04, 0x65, 0x2f, 0x48, 0x4c, 0x05, 0x24, 0xf3, 0x42, 0x1a, 0x0f, 0x9c, 0x9a,
	0xa0, 0xe1, 0xaa, 0x87, 0x71, 0xd0, 0x03, 0x89, 0x2b, 0x5e, 0x3c, 0x6e, 0xe3, 0x45, 0x96, 0xcc,
	0x75, 0xb4, 0x1d, 0x09, 0xff, 0xbb, 0x07, 0xb3, 0xb1, 0xaa, 0x35, 0x33, 0x26, 0xdc, 0xf6, 0xad,
	0xf9, 0xbe, 0xf7, 0xbd, 0xfe, 0x0a, 0xee, 0xe6, 0x7d, 0xcb, 0x72, 0x29, 0xb4, 0xa0, 0x0f, 0x30,
	0x0c, 0x0a, 0x94, 0xfb, 0x55, 0x11, 0xa9, 0x58, 0x26, 0x11, 0x72, 0xdc, 0x16, 0xa8, 0x34, 0x19,
	0x40, 0xe7, 0x45, 0xe4, 0x49, 0x3c, 0x76, 0x26, 0xce, 0xd4, 0xe5, 0x07, 0x41, 0x2e, 0xa0, 0x1d,
	0x08, 0x35, 0x3e, 0x99, 0x38, 0xd3, 0x0e, 0x2f, 0x3f, 0xe9, 0x0c, 0x2e, 0xab, 0x00, 0x5f, 0x64,
	0x19, 0xc6, 0xda, 0xd8, 0x3d, 0xe8, 0xfa, 0x69, 0x82, 0x99, 0x7e, 0x5a, 0xd4, 0x09, 0x5f, 0x9a,
	0xbe, 0xc2, 0x70, 0x81, 0x69, 0xb2, 0x43, 0xb9, 0x44, 0xa5, 0xc2, 0xb7, 0x7f, 0x66, 0x8e, 0xe1,
	0xec, 0x39, 0xdc, 0xa7, 0x22, 0x5c, 0x57, 0x73, 0x7b, 0xdc, 0x48, 0xd3, 0xa6, 0xfd, 0xdd, 0x66,
	0x01, 0xde, 0x61, 0x9d, 0x4d, 0x28, 0xf1, 0xe8, 0x9d, 0xee, 0xa1, 0xcb, 0x51, 0xe5, 0x22, 0x53,
	0x58, 0x4e, 0xe7, 0xa8, 0x7d, 0xb1, 0xc6, 0xca, 0xd5, 0xe1, 0x46, 0x96, 0x27, 0x75, 0xff, 0xca,
	0xeb, 0x72, 0x23, 0xa9, 0x84, 0xd1, 0xef, 0x02, 0xc7, 0xa7, 0x91, 0x1b, 0x38, 0x37, 0x69, 0xbe,
	0x28, 0x32, 0x5d, 0xef, 0x6b, 0xff, 0xbc, 0xfd, 0x70, 0x00, 0x1e, 0x97, 0xc1, 0x0a, 0xe5, 0x2e,
	0x89, 0x91, 0xcc, 0xa1, 0x6f, 0x73, 0x25, 0x23, 0xd6, 0x08, 0xda, 0x73, 0x99, 0x69, 0x47, 0x5b,
	0x64, 0x06, 0xbd, 0x9f, 0x34, 0xc9, 0x80, 0x35, 0xc0, 0xb5, 0x2d, 0x73, 0xe8, 0xdb, 0x34, 0xc9,
	0x88, 0x35, 0xe2, 0xb5, 0x6d, 0xcb, 0xfa, 0xdd, 0xd8, 0x17, 0x45, 0xae, 0xd9, 0xdf, 0xfc, 0xbc,
	0x2b, 0xd6, 0x7c, 0xad, 0xb4, 0x15, 0x9d, 0x56, 0xcf, 0xf9, 0xee, 0x33, 0x00, 0x00, 0xff, 0xff,
	0x82, 0x90, 0x64, 0x25, 0xdb, 0x02, 0x00, 0x00,
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
	QueryShareSubscribe(ctx context.Context, in *QueryShareSubscribeRequest, opts ...grpc.CallOption) (*ShareSubscribeResponse, error)
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

func (c *hMQServiceClient) QueryShareSubscribe(ctx context.Context, in *QueryShareSubscribeRequest, opts ...grpc.CallOption) (*ShareSubscribeResponse, error) {
	out := new(ShareSubscribeResponse)
	err := c.cc.Invoke(ctx, "/HMQService/QueryShareSubscribe", in, out, opts...)
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
	QueryShareSubscribe(context.Context, *QueryShareSubscribeRequest) (*ShareSubscribeResponse, error)
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
func (*UnimplementedHMQServiceServer) QueryShareSubscribe(ctx context.Context, req *QueryShareSubscribeRequest) (*ShareSubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryShareSubscribe not implemented")
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

func _HMQService_QueryShareSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryShareSubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HMQServiceServer).QueryShareSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HMQService/QueryShareSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HMQServiceServer).QueryShareSubscribe(ctx, req.(*QueryShareSubscribeRequest))
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
		{
			MethodName: "QueryShareSubscribe",
			Handler:    _HMQService_QueryShareSubscribe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hmq.proto",
}