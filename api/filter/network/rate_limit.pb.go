// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/filter/network/rate_limit.proto

package network

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v28 "github.com/envoyproxy/go-control-plane/api"
import google_protobuf2 "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RateLimit struct {
	// The prefix to use when emitting :ref:`statistics <config_network_filters_rate_limit_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The rate limit domain to use in the rate limit service request.
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// The rate limit descriptor list to use in the rate limit service request.
	Descriptors []*envoy_api_v28.RateLimitDescriptor `protobuf:"bytes,3,rep,name=descriptors" json:"descriptors,omitempty"`
	// The timeout in milliseconds for the rate limit service RPC. If not
	// set, this defaults to 20ms.
	Timeout *google_protobuf2.Duration `protobuf:"bytes,4,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *RateLimit) Reset()                    { *m = RateLimit{} }
func (m *RateLimit) String() string            { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()               {}
func (*RateLimit) Descriptor() ([]byte, []int) { return fileDescriptorRateLimit, []int{0} }

func (m *RateLimit) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetDescriptors() []*envoy_api_v28.RateLimitDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *RateLimit) GetTimeout() *google_protobuf2.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.api.v2.filter.network.RateLimit")
}
func (m *RateLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RateLimit) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StatPrefix) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRateLimit(dAtA, i, uint64(len(m.StatPrefix)))
		i += copy(dAtA[i:], m.StatPrefix)
	}
	if len(m.Domain) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRateLimit(dAtA, i, uint64(len(m.Domain)))
		i += copy(dAtA[i:], m.Domain)
	}
	if len(m.Descriptors) > 0 {
		for _, msg := range m.Descriptors {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintRateLimit(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Timeout != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRateLimit(dAtA, i, uint64(m.Timeout.Size()))
		n1, err := m.Timeout.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintRateLimit(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RateLimit) Size() (n int) {
	var l int
	_ = l
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sovRateLimit(uint64(l))
	}
	l = len(m.Domain)
	if l > 0 {
		n += 1 + l + sovRateLimit(uint64(l))
	}
	if len(m.Descriptors) > 0 {
		for _, e := range m.Descriptors {
			l = e.Size()
			n += 1 + l + sovRateLimit(uint64(l))
		}
	}
	if m.Timeout != nil {
		l = m.Timeout.Size()
		n += 1 + l + sovRateLimit(uint64(l))
	}
	return n
}

func sovRateLimit(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRateLimit(x uint64) (n int) {
	return sovRateLimit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RateLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRateLimit
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
			return fmt.Errorf("proto: RateLimit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RateLimit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Descriptors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
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
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Descriptors = append(m.Descriptors, &envoy_api_v28.RateLimitDescriptor{})
			if err := m.Descriptors[len(m.Descriptors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRateLimit
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
				return ErrInvalidLengthRateLimit
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timeout == nil {
				m.Timeout = &google_protobuf2.Duration{}
			}
			if err := m.Timeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRateLimit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRateLimit
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
func skipRateLimit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRateLimit
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
					return 0, ErrIntOverflowRateLimit
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
					return 0, ErrIntOverflowRateLimit
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
				return 0, ErrInvalidLengthRateLimit
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRateLimit
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
				next, err := skipRateLimit(dAtA[start:])
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
	ErrInvalidLengthRateLimit = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRateLimit   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api/filter/network/rate_limit.proto", fileDescriptorRateLimit) }

var fileDescriptorRateLimit = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x4f, 0x4e, 0x83, 0x40,
	0x14, 0x87, 0x33, 0x50, 0xab, 0x0c, 0xba, 0x61, 0x23, 0xd6, 0x84, 0x50, 0xdd, 0x10, 0x17, 0x33,
	0x09, 0xbd, 0x01, 0xe9, 0xd2, 0x44, 0xc3, 0x05, 0x9a, 0xa9, 0x0c, 0xcd, 0x8b, 0xc0, 0x90, 0xe1,
	0x81, 0x7a, 0x0d, 0x8f, 0xe3, 0xca, 0xa5, 0x4b, 0x8f, 0xd0, 0xb0, 0xf3, 0x16, 0x86, 0x7f, 0x8d,
	0x76, 0x37, 0x93, 0xef, 0xcb, 0xef, 0xe5, 0xa3, 0xb7, 0xa2, 0x04, 0x9e, 0x42, 0x86, 0x52, 0xf3,
	0x42, 0xe2, 0x8b, 0xd2, 0xcf, 0x5c, 0x0b, 0x94, 0x9b, 0x0c, 0x72, 0x40, 0x56, 0x6a, 0x85, 0xca,
	0xb9, 0x96, 0x45, 0xa3, 0xde, 0x98, 0x28, 0x81, 0x35, 0x21, 0x1b, 0x6c, 0x36, 0xda, 0x8b, 0x8b,
	0x6e, 0x41, 0x67, 0xd5, 0xe0, 0x2e, 0xbc, 0x9d, 0x52, 0xbb, 0x4c, 0xf2, 0xfe, 0xb7, 0xad, 0x53,
	0x9e, 0xd4, 0x5a, 0x20, 0xa8, 0x62, 0xe4, 0x97, 0x8d, 0xc8, 0x20, 0x11, 0x28, 0xf9, 0xf4, 0x18,
	0xc0, 0xcd, 0x9e, 0x50, 0x2b, 0x16, 0x28, 0xef, 0xbb, 0xc3, 0xce, 0x1d, 0xb5, 0x2b, 0x14, 0xb8,
	0x29, 0xb5, 0x4c, 0xe1, 0xd5, 0x25, 0x3e, 0x09, 0xac, 0xc8, 0xfa, 0xf8, 0xf9, 0x34, 0x67, 0xda,
	0xf0, 0x49, 0x4c, 0x3b, 0xfa, 0xd8, 0x43, 0x67, 0x49, 0xe7, 0x89, 0xca, 0x05, 0x14, 0xae, 0x71,
	0xac, 0x8d, 0xc0, 0x79, 0xa0, 0x76, 0x22, 0xab, 0x27, 0x0d, 0x25, 0x2a, 0x5d, 0xb9, 0xa6, 0x6f,
	0x06, 0x76, 0xb8, 0x64, 0xff, 0xba, 0x0e, 0xc7, 0xd7, 0x07, 0x33, 0xa2, 0xdd, 0xd4, 0xc9, 0x3b,
	0x31, 0xce, 0x48, 0xfc, 0x77, 0xc1, 0x59, 0xd1, 0x53, 0x84, 0x5c, 0xaa, 0x1a, 0xdd, 0x99, 0x4f,
	0x02, 0x3b, 0xbc, 0x62, 0x43, 0x38, 0x9b, 0xc2, 0xd9, 0x7a, 0x0c, 0x8f, 0x27, 0x33, 0x3a, 0xff,
	0x6a, 0x3d, 0xf2, 0xdd, 0x7a, 0x64, 0xdf, 0x7a, 0x64, 0x3b, 0xef, 0xcd, 0xd5, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x41, 0xa9, 0x2d, 0x62, 0x83, 0x01, 0x00, 0x00,
}
