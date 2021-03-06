// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/orgsvc/v1/orgsvc.proto

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

// list response
type ListResponse struct {
	Data                 []*OrgResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a94a9b62e908a3f, []int{0}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetData() []*OrgResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

// list request
type ListRequest struct {
	// id of admin user
	AdminUserId          string   `protobuf:"bytes,1,opt,name=adminUserId,proto3" json:"adminUserId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a94a9b62e908a3f, []int{1}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetAdminUserId() string {
	if m != nil {
		return m.AdminUserId
	}
	return ""
}

// delete user response
type DeleteResonse struct {
	Errors               []*ErrorType `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DeleteResonse) Reset()         { *m = DeleteResonse{} }
func (m *DeleteResonse) String() string { return proto.CompactTextString(m) }
func (*DeleteResonse) ProtoMessage()    {}
func (*DeleteResonse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a94a9b62e908a3f, []int{2}
}

func (m *DeleteResonse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResonse.Unmarshal(m, b)
}
func (m *DeleteResonse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResonse.Marshal(b, m, deterministic)
}
func (m *DeleteResonse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResonse.Merge(m, src)
}
func (m *DeleteResonse) XXX_Size() int {
	return xxx_messageInfo_DeleteResonse.Size(m)
}
func (m *DeleteResonse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResonse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResonse proto.InternalMessageInfo

func (m *DeleteResonse) GetErrors() []*ErrorType {
	if m != nil {
		return m.Errors
	}
	return nil
}

// delete user request
type DeleteRequest struct {
	// id of user
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a94a9b62e908a3f, []int{3}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// org response
type OrgResponse struct {
	// unique id of org
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AdminUserId          string   `protobuf:"bytes,3,opt,name=adminUserId,proto3" json:"adminUserId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrgResponse) Reset()         { *m = OrgResponse{} }
func (m *OrgResponse) String() string { return proto.CompactTextString(m) }
func (*OrgResponse) ProtoMessage()    {}
func (*OrgResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a94a9b62e908a3f, []int{4}
}

func (m *OrgResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrgResponse.Unmarshal(m, b)
}
func (m *OrgResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrgResponse.Marshal(b, m, deterministic)
}
func (m *OrgResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrgResponse.Merge(m, src)
}
func (m *OrgResponse) XXX_Size() int {
	return xxx_messageInfo_OrgResponse.Size(m)
}
func (m *OrgResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrgResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrgResponse proto.InternalMessageInfo

func (m *OrgResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OrgResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OrgResponse) GetAdminUserId() string {
	if m != nil {
		return m.AdminUserId
	}
	return ""
}

// create org request
type CreateRequest struct {
	// name or org
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// user id of admin
	AdminUserId          string   `protobuf:"bytes,2,opt,name=adminUserId,proto3" json:"adminUserId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a94a9b62e908a3f, []int{5}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetAdminUserId() string {
	if m != nil {
		return m.AdminUserId
	}
	return ""
}

// error type
type ErrorType struct {
	// http error code
	StatusCode int32 `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	// error detail
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorType) Reset()         { *m = ErrorType{} }
func (m *ErrorType) String() string { return proto.CompactTextString(m) }
func (*ErrorType) ProtoMessage()    {}
func (*ErrorType) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a94a9b62e908a3f, []int{6}
}

func (m *ErrorType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorType.Unmarshal(m, b)
}
func (m *ErrorType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorType.Marshal(b, m, deterministic)
}
func (m *ErrorType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorType.Merge(m, src)
}
func (m *ErrorType) XXX_Size() int {
	return xxx_messageInfo_ErrorType.Size(m)
}
func (m *ErrorType) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorType.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorType proto.InternalMessageInfo

func (m *ErrorType) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *ErrorType) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ListResponse)(nil), "pb.orgsvc.v1.ListResponse")
	proto.RegisterType((*ListRequest)(nil), "pb.orgsvc.v1.ListRequest")
	proto.RegisterType((*DeleteResonse)(nil), "pb.orgsvc.v1.DeleteResonse")
	proto.RegisterType((*DeleteRequest)(nil), "pb.orgsvc.v1.DeleteRequest")
	proto.RegisterType((*OrgResponse)(nil), "pb.orgsvc.v1.OrgResponse")
	proto.RegisterType((*CreateRequest)(nil), "pb.orgsvc.v1.CreateRequest")
	proto.RegisterType((*ErrorType)(nil), "pb.orgsvc.v1.ErrorType")
}

func init() { proto.RegisterFile("pb/orgsvc/v1/orgsvc.proto", fileDescriptor_0a94a9b62e908a3f) }

var fileDescriptor_0a94a9b62e908a3f = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x4f, 0x83, 0x40,
	0x10, 0x15, 0x5a, 0x31, 0x1d, 0xda, 0x1e, 0xf6, 0x22, 0xc5, 0x44, 0x1b, 0x4e, 0xbd, 0x08, 0x69,
	0x3d, 0x1b, 0x4d, 0x3f, 0x0e, 0x26, 0x26, 0x4d, 0xa8, 0x5e, 0xbc, 0x6d, 0xcb, 0x84, 0x90, 0xd8,
	0x82, 0x3b, 0x5b, 0x12, 0x7f, 0xa7, 0x7f, 0xc8, 0x14, 0x58, 0xbb, 0x60, 0x7b, 0x1b, 0xe6, 0x3d,
	0xde, 0x7b, 0x0c, 0x0f, 0x06, 0xd9, 0x3a, 0x48, 0x45, 0x4c, 0xf9, 0x26, 0xc8, 0xc7, 0xd5, 0xe4,
	0x67, 0x22, 0x95, 0x29, 0xeb, 0x66, 0x6b, 0xbf, 0x5a, 0xe4, 0x63, 0xef, 0x11, 0xba, 0xaf, 0x09,
	0xc9, 0x10, 0x29, 0x4b, 0x77, 0x84, 0xec, 0x1e, 0xda, 0x11, 0x97, 0xdc, 0x31, 0x86, 0xad, 0x91,
	0x3d, 0x19, 0xf8, 0x3a, 0xd9, 0x5f, 0x8a, 0x58, 0x11, 0xc3, 0x82, 0xe6, 0x05, 0x60, 0x97, 0xaf,
	0x7f, 0xed, 0x91, 0x24, 0x1b, 0x82, 0xcd, 0xa3, 0x6d, 0xb2, 0x7b, 0x27, 0x14, 0x2f, 0x91, 0x63,
	0x0c, 0x8d, 0x51, 0x27, 0xd4, 0x57, 0xde, 0x33, 0xf4, 0xe6, 0xf8, 0x89, 0x12, 0x43, 0xa4, 0xc2,
	0x30, 0x00, 0x0b, 0x85, 0x48, 0x05, 0x55, 0x96, 0xd7, 0x75, 0xcb, 0xc5, 0x01, 0x7b, 0xfb, 0xce,
	0x30, 0xac, 0x68, 0xde, 0xdd, 0x51, 0xa1, 0x34, 0xed, 0x83, 0x99, 0x28, 0x2f, 0x33, 0x89, 0xbc,
	0x15, 0xd8, 0x5a, 0xd0, 0x26, 0xcc, 0x18, 0xb4, 0x77, 0x7c, 0x8b, 0x8e, 0x59, 0x6c, 0x8a, 0xb9,
	0x99, 0xbb, 0xf5, 0x3f, 0xf7, 0x02, 0x7a, 0x33, 0x81, 0xfc, 0xe8, 0xaa, 0x64, 0x8c, 0xf3, 0x32,
	0xe6, 0x29, 0x99, 0xce, 0xdf, 0x17, 0xb1, 0x5b, 0x00, 0x92, 0x5c, 0xee, 0x69, 0x96, 0x46, 0xa5,
	0xd0, 0x65, 0xa8, 0x6d, 0x98, 0x03, 0x57, 0x5b, 0x24, 0xe2, 0xb1, 0x0a, 0xab, 0x1e, 0x27, 0x3f,
	0x06, 0x58, 0x4b, 0x11, 0xaf, 0xf2, 0x0d, 0x9b, 0x82, 0x55, 0x06, 0x63, 0x37, 0xf5, 0xcb, 0xd5,
	0xe2, 0xba, 0xe7, 0xff, 0xa4, 0x77, 0xc1, 0xe6, 0x60, 0x95, 0x27, 0x6d, 0x6a, 0xd4, 0x0e, 0xed,
	0x9e, 0x01, 0xa9, 0x52, 0x79, 0x82, 0xf6, 0xa1, 0x0b, 0xac, 0x61, 0xa5, 0xf5, 0xc3, 0x75, 0x4f,
	0x41, 0x2a, 0xc6, 0xb4, 0xff, 0xd1, 0xd5, 0x6b, 0xbb, 0xb6, 0x8a, 0xc2, 0x3e, 0xfc, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x43, 0xd2, 0x3f, 0x77, 0xcd, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrgSvcClient is the client API for OrgSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrgSvcClient interface {
	// create new org
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*OrgResponse, error)
	// delete org
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResonse, error)
	// get user org
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type orgSvcClient struct {
	cc *grpc.ClientConn
}

func NewOrgSvcClient(cc *grpc.ClientConn) OrgSvcClient {
	return &orgSvcClient{cc}
}

func (c *orgSvcClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*OrgResponse, error) {
	out := new(OrgResponse)
	err := c.cc.Invoke(ctx, "/pb.orgsvc.v1.OrgSvc/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgSvcClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResonse, error) {
	out := new(DeleteResonse)
	err := c.cc.Invoke(ctx, "/pb.orgsvc.v1.OrgSvc/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgSvcClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/pb.orgsvc.v1.OrgSvc/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrgSvcServer is the server API for OrgSvc service.
type OrgSvcServer interface {
	// create new org
	Create(context.Context, *CreateRequest) (*OrgResponse, error)
	// delete org
	Delete(context.Context, *DeleteRequest) (*DeleteResonse, error)
	// get user org
	List(context.Context, *ListRequest) (*ListResponse, error)
}

// UnimplementedOrgSvcServer can be embedded to have forward compatible implementations.
type UnimplementedOrgSvcServer struct {
}

func (*UnimplementedOrgSvcServer) Create(ctx context.Context, req *CreateRequest) (*OrgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedOrgSvcServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResonse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedOrgSvcServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterOrgSvcServer(s *grpc.Server, srv OrgSvcServer) {
	s.RegisterService(&_OrgSvc_serviceDesc, srv)
}

func _OrgSvc_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgSvcServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.orgsvc.v1.OrgSvc/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgSvcServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgSvc_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgSvcServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.orgsvc.v1.OrgSvc/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgSvcServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgSvc_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgSvcServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.orgsvc.v1.OrgSvc/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgSvcServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrgSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.orgsvc.v1.OrgSvc",
	HandlerType: (*OrgSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OrgSvc_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _OrgSvc_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _OrgSvc_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/orgsvc/v1/orgsvc.proto",
}
