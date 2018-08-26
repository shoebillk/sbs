// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blob.proto

package blob

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PushStatusCode int32

const (
	PushStatusCode_Unknown PushStatusCode = 0
	PushStatusCode_Ok      PushStatusCode = 1
	PushStatusCode_Failed  PushStatusCode = 2
)

var PushStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}
var PushStatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x PushStatusCode) String() string {
	return proto.EnumName(PushStatusCode_name, int32(x))
}
func (PushStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_blob_15818a742587a8ac, []int{0}
}

type Chunk struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Content              []byte   `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_blob_15818a742587a8ac, []int{0}
}
func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (dst *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(dst, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Chunk) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type PushStatus struct {
	Message              string         `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code                 PushStatusCode `protobuf:"varint,2,opt,name=Code,proto3,enum=PushStatusCode" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PushStatus) Reset()         { *m = PushStatus{} }
func (m *PushStatus) String() string { return proto.CompactTextString(m) }
func (*PushStatus) ProtoMessage()    {}
func (*PushStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_blob_15818a742587a8ac, []int{1}
}
func (m *PushStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushStatus.Unmarshal(m, b)
}
func (m *PushStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushStatus.Marshal(b, m, deterministic)
}
func (dst *PushStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushStatus.Merge(dst, src)
}
func (m *PushStatus) XXX_Size() int {
	return xxx_messageInfo_PushStatus.Size(m)
}
func (m *PushStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_PushStatus.DiscardUnknown(m)
}

var xxx_messageInfo_PushStatus proto.InternalMessageInfo

func (m *PushStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PushStatus) GetCode() PushStatusCode {
	if m != nil {
		return m.Code
	}
	return PushStatusCode_Unknown
}

type GetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_blob_15818a742587a8ac, []int{2}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Chunk)(nil), "Chunk")
	proto.RegisterType((*PushStatus)(nil), "PushStatus")
	proto.RegisterType((*GetRequest)(nil), "GetRequest")
	proto.RegisterEnum("PushStatusCode", PushStatusCode_name, PushStatusCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlobServiceClient is the client API for BlobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlobServiceClient interface {
	Push(ctx context.Context, opts ...grpc.CallOption) (BlobService_PushClient, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (BlobService_GetClient, error)
}

type blobServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlobServiceClient(cc *grpc.ClientConn) BlobServiceClient {
	return &blobServiceClient{cc}
}

func (c *blobServiceClient) Push(ctx context.Context, opts ...grpc.CallOption) (BlobService_PushClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BlobService_serviceDesc.Streams[0], "/BlobService/Push", opts...)
	if err != nil {
		return nil, err
	}
	x := &blobServicePushClient{stream}
	return x, nil
}

type BlobService_PushClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*PushStatus, error)
	grpc.ClientStream
}

type blobServicePushClient struct {
	grpc.ClientStream
}

func (x *blobServicePushClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *blobServicePushClient) CloseAndRecv() (*PushStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PushStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blobServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (BlobService_GetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BlobService_serviceDesc.Streams[1], "/BlobService/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &blobServiceGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlobService_GetClient interface {
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type blobServiceGetClient struct {
	grpc.ClientStream
}

func (x *blobServiceGetClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlobServiceServer is the server API for BlobService service.
type BlobServiceServer interface {
	Push(BlobService_PushServer) error
	Get(*GetRequest, BlobService_GetServer) error
}

func RegisterBlobServiceServer(s *grpc.Server, srv BlobServiceServer) {
	s.RegisterService(&_BlobService_serviceDesc, srv)
}

func _BlobService_Push_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BlobServiceServer).Push(&blobServicePushServer{stream})
}

type BlobService_PushServer interface {
	SendAndClose(*PushStatus) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type blobServicePushServer struct {
	grpc.ServerStream
}

func (x *blobServicePushServer) SendAndClose(m *PushStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *blobServicePushServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BlobService_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlobServiceServer).Get(m, &blobServiceGetServer{stream})
}

type BlobService_GetServer interface {
	Send(*Chunk) error
	grpc.ServerStream
}

type blobServiceGetServer struct {
	grpc.ServerStream
}

func (x *blobServiceGetServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

var _BlobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BlobService",
	HandlerType: (*BlobServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Push",
			Handler:       _BlobService_Push_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Get",
			Handler:       _BlobService_Get_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blob.proto",
}

func init() { proto.RegisterFile("blob.proto", fileDescriptor_blob_15818a742587a8ac) }

var fileDescriptor_blob_15818a742587a8ac = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xdd, 0x4a, 0xc3, 0x40,
	0x10, 0x85, 0xb3, 0x6b, 0x4d, 0xf1, 0x44, 0x62, 0x98, 0xab, 0x50, 0x44, 0xcb, 0x7a, 0x13, 0xbc,
	0x08, 0xb6, 0xbe, 0x81, 0x01, 0x4b, 0x11, 0x7f, 0x48, 0xf1, 0x01, 0x12, 0x33, 0xd8, 0x92, 0xb0,
	0xab, 0xdd, 0x8d, 0xbe, 0xbe, 0x64, 0x69, 0xa9, 0x7a, 0x79, 0x76, 0xcf, 0x37, 0x7c, 0x33, 0x40,
	0xdd, 0x99, 0x3a, 0xff, 0xd8, 0x1a, 0x67, 0xd4, 0x0c, 0xc7, 0xc5, 0xba, 0xd7, 0x2d, 0xc5, 0x90,
	0xcb, 0x26, 0x15, 0x53, 0x91, 0x9d, 0x94, 0x72, 0xd9, 0x50, 0x8a, 0x71, 0x61, 0xb4, 0x63, 0xed,
	0x52, 0x39, 0x15, 0xd9, 0x69, 0xb9, 0x8f, 0xea, 0x01, 0x78, 0xe9, 0xed, 0x7a, 0xe5, 0x2a, 0xd7,
	0xdb, 0xa1, 0xf7, 0xc8, 0xd6, 0x56, 0xef, 0xbc, 0x83, 0xf7, 0x91, 0xae, 0x30, 0x2a, 0x4c, 0xc3,
	0x1e, 0x8f, 0xe7, 0x67, 0xf9, 0x01, 0x1a, 0x9e, 0x4b, 0xff, 0xa9, 0xce, 0x81, 0x05, 0xbb, 0x92,
	0x3f, 0x7b, 0xb6, 0xee, 0xbf, 0xc4, 0xf5, 0x0c, 0xf1, 0x5f, 0x8a, 0x22, 0x8c, 0x5f, 0x75, 0xab,
	0xcd, 0xb7, 0x4e, 0x02, 0x0a, 0x21, 0x9f, 0xdb, 0x44, 0x10, 0x10, 0xde, 0x57, 0x9b, 0x8e, 0x9b,
	0x44, 0xce, 0x9f, 0x10, 0xdd, 0x75, 0xa6, 0x5e, 0xf1, 0xf6, 0x6b, 0xf3, 0xc6, 0x74, 0x89, 0xd1,
	0x30, 0x81, 0xc2, 0xdc, 0xaf, 0x39, 0x89, 0x7e, 0x69, 0xa8, 0x20, 0x13, 0x74, 0x81, 0xa3, 0x05,
	0x3b, 0x8a, 0xf2, 0x83, 0xc6, 0x64, 0x57, 0x56, 0xc1, 0x8d, 0xa8, 0x43, 0x7f, 0xa7, 0xdb, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x23, 0x07, 0xdf, 0x7e, 0x35, 0x01, 0x00, 0x00,
}