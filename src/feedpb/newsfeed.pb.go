// Code generated by protoc-gen-gogo.
// source: newsfeed.proto
// DO NOT EDIT!

package feedpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for NewsFeed service

type NewsFeedClient interface {
	Get(ctx context.Context, in *FeedGetRequest, opts ...grpc.CallOption) (*FeedGetResponse, error)
}

type newsFeedClient struct {
	cc *grpc.ClientConn
}

func NewNewsFeedClient(cc *grpc.ClientConn) NewsFeedClient {
	return &newsFeedClient{cc}
}

func (c *newsFeedClient) Get(ctx context.Context, in *FeedGetRequest, opts ...grpc.CallOption) (*FeedGetResponse, error) {
	out := new(FeedGetResponse)
	err := grpc.Invoke(ctx, "/feed.v1.NewsFeed/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NewsFeed service

type NewsFeedServer interface {
	Get(context.Context, *FeedGetRequest) (*FeedGetResponse, error)
}

func RegisterNewsFeedServer(s *grpc.Server, srv NewsFeedServer) {
	s.RegisterService(&_NewsFeed_serviceDesc, srv)
}

func _NewsFeed_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsFeedServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feed.v1.NewsFeed/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsFeedServer).Get(ctx, req.(*FeedGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NewsFeed_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feed.v1.NewsFeed",
	HandlerType: (*NewsFeedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _NewsFeed_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptorNewsfeed,
}

func init() { proto.RegisterFile("newsfeed.proto", fileDescriptorNewsfeed) }

var fileDescriptorNewsfeed = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4b, 0x2d, 0x2f,
	0x4e, 0x4b, 0x4d, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x07, 0xb3, 0xcb, 0x0c,
	0xa5, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2,
	0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xca, 0xa4, 0x78, 0x92, 0xf3, 0x73, 0x73,
	0xf3, 0xf3, 0x20, 0x3c, 0xa3, 0x28, 0x2e, 0x0e, 0xbf, 0xd4, 0xf2, 0x62, 0xb7, 0xd4, 0xd4, 0x14,
	0x21, 0x3f, 0x2e, 0x66, 0xf7, 0xd4, 0x12, 0x21, 0x71, 0x3d, 0xa8, 0x41, 0x7a, 0x20, 0x51, 0xf7,
	0xd4, 0x92, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x29, 0x09, 0x4c, 0x89, 0xe2, 0x82, 0xfc,
	0xbc, 0xe2, 0x54, 0x25, 0x91, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0xf1, 0x09, 0xf1, 0xe8, 0x83, 0x54,
	0xe8, 0x97, 0x19, 0xea, 0xa7, 0xa7, 0x96, 0x38, 0x49, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x44, 0xb1, 0x81, 0xa4, 0x0b, 0x92,
	0x92, 0xd8, 0xc0, 0x96, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x42, 0xc6, 0x80, 0xcd, 0xc3,
	0x00, 0x00, 0x00,
}
