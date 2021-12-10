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
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	// password of user
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// confirm password
	ConfirmPassword      string   `protobuf:"bytes,5,opt,name=confirmPassword,proto3" json:"confirmPassword,omitempty"`
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

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

// register response
type RegisterResponse struct {
	// unique id of user
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Errors               []*Error `protobuf:"bytes,5,rep,name=errors,proto3" json:"errors,omitempty"`
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

func (m *RegisterResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
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

func (m *RegisterResponse) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	// http error code
	StatusCode int32 `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	// error detail
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ad5f245eaa35e9d, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "pb.registrationsvc.v1.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "pb.registrationsvc.v1.RegisterResponse")
	proto.RegisterType((*Error)(nil), "pb.registrationsvc.v1.Error")
}

func init() {
	proto.RegisterFile("pb/registrationsvc/v1/registrationsvc.proto", fileDescriptor_5ad5f245eaa35e9d)
}

var fileDescriptor_5ad5f245eaa35e9d = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xdf, 0x4a, 0xfb, 0x30,
	0x14, 0xc7, 0x7f, 0x6d, 0xd7, 0xfd, 0xb6, 0xa3, 0x38, 0x09, 0x0e, 0xcb, 0x18, 0x32, 0x7a, 0xa1,
	0x05, 0xa1, 0x63, 0xd3, 0x17, 0x50, 0xf1, 0x56, 0x24, 0x97, 0x82, 0x17, 0xe9, 0x9a, 0x8d, 0xc0,
	0xda, 0xd4, 0x9c, 0xac, 0x3e, 0x92, 0x3e, 0xa6, 0x2c, 0x6d, 0x5c, 0x2d, 0x9b, 0x5e, 0x7e, 0xff,
	0x1c, 0xfa, 0xe9, 0xc9, 0x81, 0xeb, 0x22, 0x99, 0x2a, 0xbe, 0x12, 0xa8, 0x15, 0xd3, 0x42, 0xe6,
	0x58, 0x2e, 0xa6, 0xe5, 0xac, 0x6d, 0xc5, 0x85, 0x92, 0x5a, 0x92, 0x61, 0x91, 0xc4, 0xed, 0xa4,
	0x9c, 0x85, 0x1f, 0x0e, 0x0c, 0xa8, 0xb1, 0xb9, 0xa2, 0xfc, 0x6d, 0xc3, 0x51, 0x93, 0x31, 0xf4,
	0x97, 0x42, 0xa1, 0x7e, 0x62, 0x19, 0x0f, 0x9c, 0x89, 0x13, 0xf5, 0xe9, 0xce, 0x20, 0x23, 0xe8,
	0xad, 0x59, 0x1d, 0xba, 0x26, 0xfc, 0xd6, 0xe4, 0x0c, 0x7c, 0x9e, 0x31, 0xb1, 0x0e, 0x3c, 0x13,
	0x54, 0x62, 0x3b, 0x51, 0x30, 0xc4, 0x77, 0xa9, 0xd2, 0xa0, 0x53, 0x4d, 0x58, 0x4d, 0x22, 0x18,
	0x2c, 0x64, 0xbe, 0x14, 0x2a, 0x7b, 0xb6, 0x15, 0xdf, 0x54, 0xda, 0x76, 0xf8, 0xe9, 0xc0, 0xe9,
	0x8e, 0x14, 0x0b, 0x99, 0x23, 0x27, 0x27, 0xe0, 0x8a, 0xb4, 0x66, 0x74, 0x45, 0xfa, 0x13, 0xdd,
	0xfd, 0x0d, 0xdd, 0x3b, 0x84, 0xde, 0x69, 0xa2, 0xdf, 0x42, 0x97, 0x2b, 0x25, 0x15, 0x06, 0xfe,
	0xc4, 0x8b, 0x8e, 0xe6, 0xe3, 0x78, 0xef, 0x1a, 0xe3, 0xc7, 0x6d, 0x89, 0xd6, 0xdd, 0xf0, 0x0e,
	0x7c, 0x63, 0x90, 0x0b, 0x00, 0xd4, 0x4c, 0x6f, 0xf0, 0x41, 0xa6, 0xd5, 0x2a, 0x7d, 0xda, 0x70,
	0x48, 0x00, 0xff, 0x33, 0x8e, 0xc8, 0x56, 0x16, 0xd6, 0xca, 0x79, 0x06, 0xc7, 0xb4, 0xf1, 0x19,
	0xf2, 0x0a, 0x3d, 0xfb, 0xf3, 0xe4, 0xf2, 0x00, 0x44, 0xeb, 0x1d, 0x47, 0x57, 0x7f, 0xf6, 0xaa,
	0x2d, 0x86, 0xff, 0xee, 0xcf, 0x5f, 0x86, 0x7b, 0x8f, 0x29, 0xe9, 0x9a, 0xeb, 0xb9, 0xf9, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x02, 0x36, 0x34, 0x16, 0x6c, 0x02, 0x00, 0x00,
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
