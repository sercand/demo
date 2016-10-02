// Code generated by protoc-gen-gogo.
// source: feedprovider.proto
// DO NOT EDIT!

/*
	Package feedpb is a generated protocol buffer package.

	It is generated from these files:
		feedprovider.proto
		newsfeed.proto

	It has these top-level messages:
		ProviderGetRequest
		FeedItem
		FeedGetRequest
		FeedGetResponse
*/
package feedpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProviderGetRequest struct {
	Request *FeedGetRequest `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
}

func (m *ProviderGetRequest) Reset()                    { *m = ProviderGetRequest{} }
func (m *ProviderGetRequest) String() string            { return proto.CompactTextString(m) }
func (*ProviderGetRequest) ProtoMessage()               {}
func (*ProviderGetRequest) Descriptor() ([]byte, []int) { return fileDescriptorFeedprovider, []int{0} }

func (m *ProviderGetRequest) GetRequest() *FeedGetRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func init() {
	proto.RegisterType((*ProviderGetRequest)(nil), "feed.v1.ProviderGetRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for FeedProvider service

type FeedProviderClient interface {
	Get(ctx context.Context, in *ProviderGetRequest, opts ...grpc.CallOption) (*FeedGetResponse, error)
}

type feedProviderClient struct {
	cc *grpc.ClientConn
}

func NewFeedProviderClient(cc *grpc.ClientConn) FeedProviderClient {
	return &feedProviderClient{cc}
}

func (c *feedProviderClient) Get(ctx context.Context, in *ProviderGetRequest, opts ...grpc.CallOption) (*FeedGetResponse, error) {
	out := new(FeedGetResponse)
	err := grpc.Invoke(ctx, "/feed.v1.FeedProvider/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FeedProvider service

type FeedProviderServer interface {
	Get(context.Context, *ProviderGetRequest) (*FeedGetResponse, error)
}

func RegisterFeedProviderServer(s *grpc.Server, srv FeedProviderServer) {
	s.RegisterService(&_FeedProvider_serviceDesc, srv)
}

func _FeedProvider_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedProviderServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feed.v1.FeedProvider/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedProviderServer).Get(ctx, req.(*ProviderGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feed.v1.FeedProvider",
	HandlerType: (*FeedProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _FeedProvider_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptorFeedprovider,
}

func (m *ProviderGetRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ProviderGetRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Request != nil {
		data[i] = 0xa
		i++
		i = encodeVarintFeedprovider(data, i, uint64(m.Request.Size()))
		n1, err := m.Request.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeFixed64Feedprovider(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Feedprovider(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintFeedprovider(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *ProviderGetRequest) Size() (n int) {
	var l int
	_ = l
	if m.Request != nil {
		l = m.Request.Size()
		n += 1 + l + sovFeedprovider(uint64(l))
	}
	return n
}

func sovFeedprovider(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFeedprovider(x uint64) (n int) {
	return sovFeedprovider(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProviderGetRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeedprovider
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProviderGetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProviderGetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeedprovider
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeedprovider
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Request == nil {
				m.Request = &FeedGetRequest{}
			}
			if err := m.Request.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeedprovider(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFeedprovider
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFeedprovider(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFeedprovider
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFeedprovider
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFeedprovider
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthFeedprovider
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFeedprovider
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipFeedprovider(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthFeedprovider = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFeedprovider   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("feedprovider.proto", fileDescriptorFeedprovider) }

var fileDescriptorFeedprovider = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x4b, 0x4d, 0x4d,
	0x29, 0x28, 0xca, 0x2f, 0xcb, 0x4c, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x07, 0x89, 0xe9, 0x95, 0x19, 0x4a, 0xf1, 0xe5, 0xa5, 0x96, 0x17, 0x83, 0x39, 0x60, 0x09, 0x25,
	0x77, 0x2e, 0xa1, 0x00, 0xa8, 0x52, 0xf7, 0xd4, 0x92, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12,
	0x21, 0x43, 0x2e, 0xf6, 0x22, 0x08, 0x53, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x5c, 0x0f,
	0x6a, 0x80, 0x9e, 0x5b, 0x6a, 0x6a, 0x0a, 0x42, 0x65, 0x10, 0x4c, 0x9d, 0x91, 0x0f, 0x17, 0x0f,
	0x48, 0x0a, 0x66, 0x98, 0x90, 0x0d, 0x17, 0xb3, 0x7b, 0x6a, 0x89, 0x90, 0x34, 0x5c, 0x23, 0xa6,
	0x35, 0x52, 0x12, 0x98, 0xa6, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x3a, 0x49, 0x9c, 0x78, 0x24,
	0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x44, 0xb1,
	0x81, 0x7d, 0x95, 0x94, 0xc4, 0x06, 0x76, 0xb7, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x3e,
	0x17, 0xe9, 0xe6, 0x00, 0x00, 0x00,
}