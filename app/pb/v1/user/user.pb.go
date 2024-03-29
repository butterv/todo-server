// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1/user/user.proto

package user

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateUserRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8fd1c8d22ccfa2e, []int{0}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateUserResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8fd1c8d22ccfa2e, []int{1}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateUserRequest)(nil), "user.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "user.CreateUserResponse")
}

func init() { proto.RegisterFile("proto/v1/user/user.proto", fileDescriptor_a8fd1c8d22ccfa2e) }

var fileDescriptor_a8fd1c8d22ccfa2e = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0x95, 0xfe, 0x7d, 0xad, 0xfb, 0x09, 0x15, 0xab, 0x52, 0xad, 0x88, 0x01, 0x65, 0x01,
	0x81, 0x1a, 0x0b, 0xd8, 0x98, 0x50, 0xb9, 0x83, 0xa0, 0x2e, 0x4c, 0x38, 0xcd, 0x51, 0x6a, 0x29,
	0xcd, 0x31, 0x3e, 0x4e, 0xd8, 0xb9, 0x05, 0xee, 0x86, 0x89, 0x7b, 0xe0, 0x16, 0x58, 0xb8, 0x09,
	0x84, 0xe2, 0x28, 0x80, 0xd4, 0xc5, 0x7a, 0xed, 0xe7, 0xf1, 0xdf, 0xcb, 0x84, 0xb1, 0xe8, 0x50,
	0xd6, 0x17, 0xb2, 0x22, 0xb0, 0x7e, 0x88, 0xfd, 0x12, 0x1f, 0x34, 0x39, 0x3c, 0xca, 0x11, 0xf3,
	0x02, 0xa4, 0x32, 0x5a, 0xaa, 0xb2, 0x44, 0xa7, 0x9c, 0xc6, 0x92, 0x5a, 0x27, 0x5c, 0xd4, 0xaa,
	0xd0, 0x99, 0x72, 0x20, 0xbb, 0xd0, 0x82, 0xe8, 0x81, 0x1d, 0xde, 0x5a, 0x50, 0x0e, 0xd6, 0x04,
	0x36, 0x81, 0xc7, 0x0a, 0xc8, 0xf1, 0x88, 0x0d, 0x61, 0xa7, 0x74, 0x21, 0x82, 0xe3, 0xe0, 0x74,
	0xb2, 0xfa, 0xff, 0xfa, 0xf9, 0xd6, 0xff, 0x67, 0x87, 0xb3, 0xbe, 0xf8, 0x0a, 0x92, 0x16, 0xf1,
	0x13, 0x36, 0x36, 0x8a, 0xe8, 0x09, 0x6d, 0x26, 0x7a, 0x5e, 0x9b, 0x36, 0xda, 0xc8, 0x0e, 0x66,
	0x63, 0x71, 0x93, 0xfc, 0xc0, 0x68, 0xce, 0xf8, 0xdf, 0x1b, 0xc8, 0x60, 0x49, 0x70, 0x99, 0xb1,
	0x69, 0x33, 0xbf, 0x03, 0x5b, 0xeb, 0x0d, 0xf0, 0x35, 0x63, 0xbf, 0x12, 0x5f, 0xc4, 0xfe, 0x7b,
	0x7b, 0x0f, 0x0b, 0xc5, 0x3e, 0x68, 0xcf, 0x8b, 0xe6, 0xcf, 0xef, 0x1f, 0x2f, 0xbd, 0x83, 0x68,
	0xd2, 0x35, 0x44, 0xd7, 0xc1, 0xd9, 0x6a, 0x79, 0x7f, 0x9e, 0x6b, 0xb7, 0xad, 0xd2, 0x78, 0x83,
	0x3b, 0xa9, 0xc9, 0xd1, 0x56, 0x3a, 0xcc, 0x70, 0x49, 0x60, 0x6b, 0xb0, 0x52, 0x19, 0x23, 0x4d,
	0xda, 0x6d, 0x49, 0x47, 0xbe, 0x93, 0xab, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x23, 0x39, 0x04,
	0x72, 0x6c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/user/user.proto",
}
