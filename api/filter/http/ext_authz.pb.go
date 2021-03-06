// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/filter/http/ext_authz.proto

package envoy_api_v2_filter_http

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v22 "github.com/envoyproxy/go-control-plane/api"
import _ "github.com/lyft/protoc-gen-validate/validate"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// [#not-implemented-hide:]
// External Authorization filter calls out to an external service over the
// gRPC Authorization API defined by :ref:`external_auth <envoy_api_msg_auth.CheckRequest>`.
// A failed check will cause this filter to return 403 Forbidden.
type ExtAuthz struct {
	// The external authorization gRPC service configuration.
	GrpcService *envoy_api_v22.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService" json:"grpc_service,omitempty"`
	// The filter's behaviour in case the external authorization service does
	// not respond back. If set to true then in case of failure to get a
	// response back from the authorization service allow the traffic.
	// Defaults to false.
	// If set to true and the response from the authorization service is NOT
	// Denied then the traffic will be permitted.
	FailureModeAllow bool `protobuf:"varint,2,opt,name=failure_mode_allow,json=failureModeAllow,proto3" json:"failure_mode_allow,omitempty"`
}

func (m *ExtAuthz) Reset()                    { *m = ExtAuthz{} }
func (m *ExtAuthz) String() string            { return proto.CompactTextString(m) }
func (*ExtAuthz) ProtoMessage()               {}
func (*ExtAuthz) Descriptor() ([]byte, []int) { return fileDescriptorExtAuthz, []int{0} }

func (m *ExtAuthz) GetGrpcService() *envoy_api_v22.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func (m *ExtAuthz) GetFailureModeAllow() bool {
	if m != nil {
		return m.FailureModeAllow
	}
	return false
}

func init() {
	proto.RegisterType((*ExtAuthz)(nil), "envoy.api.v2.filter.http.ExtAuthz")
}
func (m *ExtAuthz) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtAuthz) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.GrpcService != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintExtAuthz(dAtA, i, uint64(m.GrpcService.Size()))
		n1, err := m.GrpcService.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.FailureModeAllow {
		dAtA[i] = 0x10
		i++
		if m.FailureModeAllow {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintExtAuthz(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ExtAuthz) Size() (n int) {
	var l int
	_ = l
	if m.GrpcService != nil {
		l = m.GrpcService.Size()
		n += 1 + l + sovExtAuthz(uint64(l))
	}
	if m.FailureModeAllow {
		n += 2
	}
	return n
}

func sovExtAuthz(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozExtAuthz(x uint64) (n int) {
	return sovExtAuthz(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExtAuthz) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExtAuthz
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExtAuthz: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtAuthz: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrpcService", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtAuthz
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthExtAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GrpcService == nil {
				m.GrpcService = &envoy_api_v22.GrpcService{}
			}
			if err := m.GrpcService.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailureModeAllow", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtAuthz
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FailureModeAllow = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipExtAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExtAuthz
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
func skipExtAuthz(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExtAuthz
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowExtAuthz
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowExtAuthz
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthExtAuthz
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowExtAuthz
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
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
				next, err := skipExtAuthz(dAtA[start:])
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
	ErrInvalidLengthExtAuthz = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExtAuthz   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api/filter/http/ext_authz.proto", fileDescriptorExtAuthz) }

var fileDescriptorExtAuthz = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2c, 0xc8, 0xd4,
	0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2, 0xcf, 0x28, 0x29, 0x29, 0xd0, 0x4f, 0xad, 0x28, 0x89,
	0x4f, 0x2c, 0x2d, 0xc9, 0xa8, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x48, 0xcd, 0x2b,
	0xcb, 0xaf, 0xd4, 0x4b, 0x2c, 0xc8, 0xd4, 0x2b, 0x33, 0xd2, 0x83, 0xa8, 0xd4, 0x03, 0xa9, 0x94,
	0x12, 0x03, 0x69, 0x4d, 0x2f, 0x2a, 0x48, 0x8e, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x85,
	0xe8, 0x90, 0x12, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20, 0x12,
	0x4a, 0x65, 0x5c, 0x1c, 0xae, 0x15, 0x25, 0x8e, 0x20, 0xc3, 0x85, 0x6c, 0xb8, 0x78, 0x90, 0xb5,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x49, 0xea, 0xa1, 0xd8, 0xe6, 0x5e, 0x54, 0x90, 0x1c,
	0x0c, 0x51, 0x10, 0xc4, 0x9d, 0x8e, 0xe0, 0x08, 0xe9, 0x70, 0x09, 0xa5, 0x25, 0x66, 0xe6, 0x94,
	0x16, 0xa5, 0xc6, 0xe7, 0xe6, 0xa7, 0xa4, 0xc6, 0x27, 0xe6, 0xe4, 0xe4, 0x97, 0x4b, 0x30, 0x29,
	0x30, 0x6a, 0x70, 0x04, 0x09, 0x40, 0x65, 0x7c, 0xf3, 0x53, 0x52, 0x1d, 0x41, 0xe2, 0x4e, 0x3c,
	0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0x63, 0x12, 0x1b, 0xd8,
	0x31, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xd6, 0x31, 0xcd, 0xfa, 0x00, 0x00, 0x00,
}
