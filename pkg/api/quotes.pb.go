// Code generated by protoc-gen-go. DO NOT EDIT.
// source: quotes.proto

package api

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

type QuoteResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Genre                string   `protobuf:"bytes,2,opt,name=genre,proto3" json:"genre,omitempty"`
	Body                 string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Author               string   `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuoteResponse) Reset()         { *m = QuoteResponse{} }
func (m *QuoteResponse) String() string { return proto.CompactTextString(m) }
func (*QuoteResponse) ProtoMessage()    {}
func (*QuoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa8c397fe3f70398, []int{0}
}

func (m *QuoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuoteResponse.Unmarshal(m, b)
}
func (m *QuoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuoteResponse.Marshal(b, m, deterministic)
}
func (m *QuoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuoteResponse.Merge(m, src)
}
func (m *QuoteResponse) XXX_Size() int {
	return xxx_messageInfo_QuoteResponse.Size(m)
}
func (m *QuoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuoteResponse proto.InternalMessageInfo

func (m *QuoteResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *QuoteResponse) GetGenre() string {
	if m != nil {
		return m.Genre
	}
	return ""
}

func (m *QuoteResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *QuoteResponse) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

type QuoteRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Genre                string   `protobuf:"bytes,2,opt,name=genre,proto3" json:"genre,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuoteRequest) Reset()         { *m = QuoteRequest{} }
func (m *QuoteRequest) String() string { return proto.CompactTextString(m) }
func (*QuoteRequest) ProtoMessage()    {}
func (*QuoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa8c397fe3f70398, []int{1}
}

func (m *QuoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuoteRequest.Unmarshal(m, b)
}
func (m *QuoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuoteRequest.Marshal(b, m, deterministic)
}
func (m *QuoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuoteRequest.Merge(m, src)
}
func (m *QuoteRequest) XXX_Size() int {
	return xxx_messageInfo_QuoteRequest.Size(m)
}
func (m *QuoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuoteRequest proto.InternalMessageInfo

func (m *QuoteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *QuoteRequest) GetGenre() string {
	if m != nil {
		return m.Genre
	}
	return ""
}

func init() {
	proto.RegisterType((*QuoteResponse)(nil), "api.QuoteResponse")
	proto.RegisterType((*QuoteRequest)(nil), "api.QuoteRequest")
}

func init() { proto.RegisterFile("quotes.proto", fileDescriptor_aa8c397fe3f70398) }

var fileDescriptor_aa8c397fe3f70398 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2c, 0xcd, 0x2f,
	0x49, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x4a, 0xe4,
	0xe2, 0x0d, 0x04, 0x09, 0x06, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0xf1, 0x71, 0x31,
	0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x31, 0x65, 0xa6, 0x08, 0x89, 0x70, 0xb1,
	0xa6, 0xa7, 0xe6, 0x15, 0xa5, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0x42,
	0x5c, 0x2c, 0x49, 0xf9, 0x29, 0x95, 0x12, 0xcc, 0x60, 0x41, 0x30, 0x5b, 0x48, 0x8c, 0x8b, 0x2d,
	0xb1, 0xb4, 0x24, 0x23, 0xbf, 0x48, 0x82, 0x05, 0x2c, 0x0a, 0xe5, 0x29, 0x99, 0x70, 0xf1, 0x40,
	0xad, 0x28, 0x2c, 0x4d, 0x2d, 0x2e, 0x21, 0xce, 0x06, 0x23, 0x33, 0x2e, 0x36, 0xb0, 0xae, 0x62,
	0x21, 0x1d, 0x2e, 0x66, 0xf7, 0xd4, 0x12, 0x21, 0x41, 0xbd, 0xc4, 0x82, 0x4c, 0x3d, 0x64, 0x93,
	0xa4, 0x84, 0x90, 0x85, 0x20, 0xee, 0x4f, 0x62, 0x03, 0x7b, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xd0, 0x10, 0x0c, 0x33, 0xec, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QuotesClient is the client API for Quotes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuotesClient interface {
	Get(ctx context.Context, in *QuoteRequest, opts ...grpc.CallOption) (*QuoteResponse, error)
}

type quotesClient struct {
	cc *grpc.ClientConn
}

func NewQuotesClient(cc *grpc.ClientConn) QuotesClient {
	return &quotesClient{cc}
}

func (c *quotesClient) Get(ctx context.Context, in *QuoteRequest, opts ...grpc.CallOption) (*QuoteResponse, error) {
	out := new(QuoteResponse)
	err := c.cc.Invoke(ctx, "/api.Quotes/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuotesServer is the server API for Quotes service.
type QuotesServer interface {
	Get(context.Context, *QuoteRequest) (*QuoteResponse, error)
}

// UnimplementedQuotesServer can be embedded to have forward compatible implementations.
type UnimplementedQuotesServer struct {
}

func (*UnimplementedQuotesServer) Get(ctx context.Context, req *QuoteRequest) (*QuoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterQuotesServer(s *grpc.Server, srv QuotesServer) {
	s.RegisterService(&_Quotes_serviceDesc, srv)
}

func _Quotes_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Quotes/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotesServer).Get(ctx, req.(*QuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Quotes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Quotes",
	HandlerType: (*QuotesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Quotes_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quotes.proto",
}
