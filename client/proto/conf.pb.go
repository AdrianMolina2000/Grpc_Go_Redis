// Code generated by protoc-gen-go. DO NOT EDIT.
// source: conf.proto

package dataproto

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

type RequestData struct {
	Team1                string   `protobuf:"bytes,1,opt,name=team1,proto3" json:"team1,omitempty"`
	Team2                string   `protobuf:"bytes,2,opt,name=team2,proto3" json:"team2,omitempty"`
	Score                string   `protobuf:"bytes,3,opt,name=score,proto3" json:"score,omitempty"`
	Phase                string   `protobuf:"bytes,4,opt,name=phase,proto3" json:"phase,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestData) Reset()         { *m = RequestData{} }
func (m *RequestData) String() string { return proto.CompactTextString(m) }
func (*RequestData) ProtoMessage()    {}
func (*RequestData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b6ecbfc68e85c65, []int{0}
}

func (m *RequestData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestData.Unmarshal(m, b)
}
func (m *RequestData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestData.Marshal(b, m, deterministic)
}
func (m *RequestData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestData.Merge(m, src)
}
func (m *RequestData) XXX_Size() int {
	return xxx_messageInfo_RequestData.Size(m)
}
func (m *RequestData) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestData.DiscardUnknown(m)
}

var xxx_messageInfo_RequestData proto.InternalMessageInfo

func (m *RequestData) GetTeam1() string {
	if m != nil {
		return m.Team1
	}
	return ""
}

func (m *RequestData) GetTeam2() string {
	if m != nil {
		return m.Team2
	}
	return ""
}

func (m *RequestData) GetScore() string {
	if m != nil {
		return m.Score
	}
	return ""
}

func (m *RequestData) GetPhase() string {
	if m != nil {
		return m.Phase
	}
	return ""
}

type ResponseData struct {
	Respuesta            string   `protobuf:"bytes,1,opt,name=respuesta,proto3" json:"respuesta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseData) Reset()         { *m = ResponseData{} }
func (m *ResponseData) String() string { return proto.CompactTextString(m) }
func (*ResponseData) ProtoMessage()    {}
func (*ResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b6ecbfc68e85c65, []int{1}
}

func (m *ResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseData.Unmarshal(m, b)
}
func (m *ResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseData.Marshal(b, m, deterministic)
}
func (m *ResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseData.Merge(m, src)
}
func (m *ResponseData) XXX_Size() int {
	return xxx_messageInfo_ResponseData.Size(m)
}
func (m *ResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseData proto.InternalMessageInfo

func (m *ResponseData) GetRespuesta() string {
	if m != nil {
		return m.Respuesta
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestData)(nil), "dataproto.requestData")
	proto.RegisterType((*ResponseData)(nil), "dataproto.responseData")
}

func init() { proto.RegisterFile("conf.proto", fileDescriptor_0b6ecbfc68e85c65) }

var fileDescriptor_0b6ecbfc68e85c65 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x3f, 0xcb, 0xc2, 0x30,
	0x10, 0x87, 0xdf, 0xbe, 0xfe, 0xa3, 0x57, 0xa7, 0x22, 0x1a, 0xc4, 0x41, 0x3a, 0x39, 0x48, 0xc1,
	0xfa, 0x05, 0x54, 0xdc, 0xdc, 0x1c, 0xdd, 0xce, 0xf6, 0xac, 0x8b, 0x4d, 0xbd, 0xc4, 0xef, 0x2f,
	0xb9, 0xa8, 0xd1, 0x29, 0x79, 0x1e, 0x02, 0xcf, 0x2f, 0x00, 0xa5, 0x6e, 0x2e, 0x79, 0xcb, 0xda,
	0xea, 0x34, 0xae, 0xd0, 0xa2, 0x5c, 0x33, 0x82, 0x84, 0xe9, 0xfe, 0x20, 0x63, 0xf7, 0x68, 0x31,
	0x1d, 0x41, 0xcf, 0x12, 0xde, 0x56, 0x2a, 0x9a, 0x47, 0x8b, 0xf8, 0xe8, 0xe1, 0x6d, 0x0b, 0xf5,
	0x1f, 0x6c, 0xe1, 0xac, 0x29, 0x35, 0x93, 0xea, 0x78, 0x2b, 0xe0, 0x6c, 0x7b, 0x45, 0x43, 0xaa,
	0xeb, 0xad, 0x40, 0xb6, 0x84, 0x21, 0x93, 0x69, 0x75, 0x63, 0x48, 0x3a, 0x33, 0x88, 0x1d, 0xbb,
	0x2e, 0xbe, 0x5a, 0x41, 0x14, 0x07, 0x18, 0x60, 0x55, 0xc9, 0xc3, 0x0d, 0x24, 0xdb, 0x9a, 0xa9,
	0x46, 0x16, 0x1c, 0xe7, 0x9f, 0xe9, 0xf9, 0xd7, 0xee, 0xe9, 0xe4, 0xc7, 0x87, 0x50, 0xf6, 0xb7,
	0x4b, 0x4e, 0xe1, 0xbb, 0xe7, 0xbe, 0x1c, 0xeb, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0x53,
	0xaf, 0x52, 0x0e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AddDataClient is the client API for AddData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddDataClient interface {
	AgregarData(ctx context.Context, in *RequestData, opts ...grpc.CallOption) (*ResponseData, error)
}

type addDataClient struct {
	cc *grpc.ClientConn
}

func NewAddDataClient(cc *grpc.ClientConn) AddDataClient {
	return &addDataClient{cc}
}

func (c *addDataClient) AgregarData(ctx context.Context, in *RequestData, opts ...grpc.CallOption) (*ResponseData, error) {
	out := new(ResponseData)
	err := c.cc.Invoke(ctx, "/dataproto.addData/AgregarData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddDataServer is the server API for AddData service.
type AddDataServer interface {
	AgregarData(context.Context, *RequestData) (*ResponseData, error)
}

// UnimplementedAddDataServer can be embedded to have forward compatible implementations.
type UnimplementedAddDataServer struct {
}

func (*UnimplementedAddDataServer) AgregarData(ctx context.Context, req *RequestData) (*ResponseData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AgregarData not implemented")
}

func RegisterAddDataServer(s *grpc.Server, srv AddDataServer) {
	s.RegisterService(&_AddData_serviceDesc, srv)
}

func _AddData_AgregarData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddDataServer).AgregarData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataproto.addData/AgregarData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddDataServer).AgregarData(ctx, req.(*RequestData))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddData_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dataproto.addData",
	HandlerType: (*AddDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AgregarData",
			Handler:    _AddData_AgregarData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conf.proto",
}
