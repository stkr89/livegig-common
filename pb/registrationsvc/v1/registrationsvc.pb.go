// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/registrationsvc/v1/registrationsvc.proto

package v1

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

// register request
type RegisterRequest struct {
	// first name of user
	FirstName string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	// last name of user
	LastName string `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	// email address of user
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ad5f245eaa35e9d, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *RegisterRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

// register response
type RegisterResponse struct {
	// unique id of user
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ad5f245eaa35e9d, []int{1}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RegisterResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *RegisterResponse) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *RegisterResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "pb.registrationsvc.v1.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "pb.registrationsvc.v1.RegisterResponse")
}

func init() {
	proto.RegisterFile("pb/registrationsvc/v1/registrationsvc.proto", fileDescriptor_5ad5f245eaa35e9d)
}

var fileDescriptor_5ad5f245eaa35e9d = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2e, 0x48, 0xd2, 0x2f,
	0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x29, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x2e, 0x4b, 0xd6, 0x2f,
	0x33, 0x44, 0x17, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2d, 0x48, 0xd2, 0x43, 0x97,
	0x29, 0x33, 0x54, 0x4a, 0xe4, 0xe2, 0x0f, 0x02, 0x8b, 0xa6, 0x16, 0x05, 0xa5, 0x16, 0x96, 0xa6,
	0x16, 0x97, 0x08, 0xc9, 0x70, 0x71, 0xa6, 0x65, 0x16, 0x15, 0x97, 0xf8, 0x25, 0xe6, 0xa6, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x21, 0x04, 0x84, 0xa4, 0xb8, 0x38, 0x72, 0x12, 0xa1, 0x92,
	0x4c, 0x60, 0x49, 0x38, 0x5f, 0x48, 0x84, 0x8b, 0x35, 0x35, 0x37, 0x31, 0x33, 0x47, 0x82, 0x19,
	0x2c, 0x01, 0xe1, 0x28, 0x15, 0x71, 0x09, 0x20, 0xac, 0x28, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15,
	0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x01, 0x1b, 0xce, 0x1a, 0xc4, 0x94, 0x99, 0x82, 0x6a, 0x27, 0x13,
	0x3e, 0x3b, 0x99, 0x71, 0xd9, 0xc9, 0x82, 0x64, 0xa7, 0x51, 0x2e, 0x17, 0x4f, 0x10, 0x92, 0x67,
	0x85, 0x62, 0xb9, 0x38, 0x60, 0x6e, 0x10, 0x52, 0xd3, 0xc3, 0x1a, 0x14, 0x7a, 0x68, 0xe1, 0x20,
	0xa5, 0x4e, 0x50, 0x1d, 0xc4, 0x33, 0x4a, 0x0c, 0x4e, 0xe2, 0x51, 0xa2, 0x58, 0xe3, 0x22, 0x89,
	0x0d, 0x1c, 0xf8, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x34, 0xbe, 0x3f, 0xab, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegistrationClient is the client API for Registration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistrationClient interface {
	// register new user
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type registrationClient struct {
	cc *grpc.ClientConn
}

func NewRegistrationClient(cc *grpc.ClientConn) RegistrationClient {
	return &registrationClient{cc}
}

func (c *registrationClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/pb.registrationsvc.v1.Registration/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistrationServer is the server API for Registration service.
type RegistrationServer interface {
	// register new user
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
}

// UnimplementedRegistrationServer can be embedded to have forward compatible implementations.
type UnimplementedRegistrationServer struct {
}

func (*UnimplementedRegistrationServer) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterRegistrationServer(s *grpc.Server, srv RegistrationServer) {
	s.RegisterService(&_Registration_serviceDesc, srv)
}

func _Registration_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.registrationsvc.v1.Registration/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Registration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.registrationsvc.v1.Registration",
	HandlerType: (*RegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Registration_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/registrationsvc/v1/registrationsvc.proto",
}
