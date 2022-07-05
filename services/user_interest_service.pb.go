// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/user_interest_service.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	resources "github.com/kritzware/google-ads-go/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request message for [UserInterestService.GetUserInterest][google.ads.googleads.v9.services.UserInterestService.GetUserInterest].
type GetUserInterestRequest struct {
	// Resource name of the UserInterest to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInterestRequest) Reset()         { *m = GetUserInterestRequest{} }
func (m *GetUserInterestRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserInterestRequest) ProtoMessage()    {}
func (*GetUserInterestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a3b67691e846704, []int{0}
}

func (m *GetUserInterestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInterestRequest.Unmarshal(m, b)
}
func (m *GetUserInterestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInterestRequest.Marshal(b, m, deterministic)
}
func (m *GetUserInterestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInterestRequest.Merge(m, src)
}
func (m *GetUserInterestRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserInterestRequest.Size(m)
}
func (m *GetUserInterestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInterestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInterestRequest proto.InternalMessageInfo

func (m *GetUserInterestRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserInterestRequest)(nil), "google.ads.googleads.v9.services.GetUserInterestRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/user_interest_service.proto", fileDescriptor_6a3b67691e846704)
}

var fileDescriptor_6a3b67691e846704 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc6, 0x49, 0x04, 0xc1, 0xa0, 0x08, 0x11, 0xa4, 0x14, 0x87, 0x52, 0x3b, 0x48, 0x87, 0xbb,
	0xa8, 0x88, 0x72, 0xda, 0x21, 0x5d, 0xaa, 0x8b, 0x94, 0x8a, 0x1d, 0x24, 0x50, 0x62, 0xf3, 0x12,
	0x02, 0xcd, 0x5d, 0xbd, 0xf7, 0xd2, 0x45, 0x5c, 0xfc, 0x0a, 0x7e, 0x03, 0x47, 0x77, 0xbf, 0x84,
	0xe0, 0xe4, 0x57, 0x70, 0xf2, 0x4b, 0x28, 0xe9, 0xe5, 0x42, 0xd5, 0x96, 0x6e, 0x0f, 0x6f, 0x9e,
	0xdf, 0xf3, 0xfe, 0xc9, 0x39, 0x67, 0xb1, 0x10, 0xf1, 0x08, 0x68, 0x18, 0x21, 0xd5, 0x32, 0x57,
	0x13, 0x8f, 0x22, 0xc8, 0x49, 0x32, 0x04, 0xa4, 0x19, 0x82, 0x1c, 0x24, 0x5c, 0x81, 0x04, 0x54,
	0x83, 0xa2, 0x4c, 0xc6, 0x52, 0x28, 0xe1, 0xd6, 0x34, 0x42, 0xc2, 0x08, 0x49, 0x49, 0x93, 0x89,
	0x47, 0x0c, 0x5d, 0x3d, 0x5a, 0x94, 0x2f, 0x01, 0x45, 0x26, 0xff, 0x35, 0xd0, 0xc1, 0xd5, 0x1d,
	0x83, 0x8d, 0x13, 0x1a, 0x72, 0x2e, 0x54, 0xa8, 0x12, 0xc1, 0x51, 0x7f, 0xad, 0xb7, 0x9c, 0xed,
	0x0e, 0xa8, 0x6b, 0x04, 0x79, 0x51, 0x60, 0x3d, 0xb8, 0xcb, 0x00, 0x95, 0xbb, 0xeb, 0x6c, 0x98,
	0xe0, 0x01, 0x0f, 0x53, 0xa8, 0x58, 0x35, 0x6b, 0x6f, 0xad, 0xb7, 0x6e, 0x8a, 0x97, 0x61, 0x0a,
	0x07, 0xef, 0x96, 0xb3, 0x35, 0x0b, 0x5f, 0xe9, 0x61, 0xdd, 0x57, 0xcb, 0xd9, 0xfc, 0x93, 0xeb,
	0x9e, 0x90, 0x65, 0x2b, 0x92, 0xf9, 0xa3, 0x54, 0xe9, 0x42, 0xb2, 0x5c, 0x9d, 0xcc, 0x72, 0xf5,
	0xe3, 0xc7, 0x8f, 0xcf, 0x27, 0x7b, 0xdf, 0xa5, 0xf9, 0x79, 0xee, 0x7f, 0xad, 0xd1, 0x1a, 0x66,
	0xa8, 0x44, 0x0a, 0x12, 0x69, 0x73, 0x7a, 0x2f, 0x03, 0x21, 0x6d, 0x3e, 0xb4, 0xbf, 0x2d, 0xa7,
	0x31, 0x14, 0xe9, 0xd2, 0x49, 0xdb, 0x95, 0x39, 0x5b, 0x77, 0xf3, 0x8b, 0x76, 0xad, 0x9b, 0xf3,
	0x82, 0x8e, 0xc5, 0x28, 0xe4, 0x31, 0x11, 0x32, 0xa6, 0x31, 0xf0, 0xe9, 0xbd, 0xcd, 0x8f, 0x1b,
	0x27, 0xb8, 0xf8, 0x9d, 0x9c, 0x1a, 0xf1, 0x6c, 0xaf, 0x74, 0x7c, 0xff, 0xc5, 0xae, 0x75, 0x74,
	0xa0, 0x1f, 0x21, 0xd1, 0x32, 0x57, 0x7d, 0x8f, 0x14, 0x8d, 0xf1, 0xcd, 0x58, 0x02, 0x3f, 0xc2,
	0xa0, 0xb4, 0x04, 0x7d, 0x2f, 0x30, 0x96, 0x2f, 0xbb, 0xa1, 0xeb, 0x8c, 0xf9, 0x11, 0x32, 0x56,
	0x9a, 0x18, 0xeb, 0x7b, 0x8c, 0x19, 0xdb, 0xed, 0xea, 0x74, 0xce, 0xc3, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xa8, 0xbe, 0x8f, 0x2a, 0xce, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserInterestServiceClient is the client API for UserInterestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserInterestServiceClient interface {
	// Returns the requested user interest in full detail
	GetUserInterest(ctx context.Context, in *GetUserInterestRequest, opts ...grpc.CallOption) (*resources.UserInterest, error)
}

type userInterestServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserInterestServiceClient(cc *grpc.ClientConn) UserInterestServiceClient {
	return &userInterestServiceClient{cc}
}

func (c *userInterestServiceClient) GetUserInterest(ctx context.Context, in *GetUserInterestRequest, opts ...grpc.CallOption) (*resources.UserInterest, error) {
	out := new(resources.UserInterest)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.UserInterestService/GetUserInterest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInterestServiceServer is the server API for UserInterestService service.
type UserInterestServiceServer interface {
	// Returns the requested user interest in full detail
	GetUserInterest(context.Context, *GetUserInterestRequest) (*resources.UserInterest, error)
}

func RegisterUserInterestServiceServer(s *grpc.Server, srv UserInterestServiceServer) {
	s.RegisterService(&_UserInterestService_serviceDesc, srv)
}

func _UserInterestService_GetUserInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInterestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInterestServiceServer).GetUserInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.UserInterestService/GetUserInterest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInterestServiceServer).GetUserInterest(ctx, req.(*GetUserInterestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserInterestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.UserInterestService",
	HandlerType: (*UserInterestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInterest",
			Handler:    _UserInterestService_GetUserInterest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/user_interest_service.proto",
}
