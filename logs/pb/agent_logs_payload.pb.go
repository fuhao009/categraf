// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: agent_logs_payload.proto

/*
	Package pb is a generated protocol buffer package.

	It is generated from these files:
		agent_logs_payload.proto

	It has these top-level messages:
		Log
*/
package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Log struct {
	Message   string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// from host
	Hostname string `protobuf:"bytes,4,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// from config
	Service string `protobuf:"bytes,5,opt,name=service,proto3" json:"service,omitempty"`
	Source  string `protobuf:"bytes,6,opt,name=source,proto3" json:"source,omitempty"`
	// from config, container tags, ...
	Tags []string `protobuf:"bytes,7,rep,name=tags" json:"tags,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptorAgentLogsPayload, []int{0} }

func (m *Log) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Log) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Log) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Log) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Log) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Log) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Log) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*Log)(nil), "pb.Log")
}
func (m *Log) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Log) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAgentLogsPayload(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if len(m.Status) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAgentLogsPayload(dAtA, i, uint64(len(m.Status)))
		i += copy(dAtA[i:], m.Status)
	}
	if m.Timestamp != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintAgentLogsPayload(dAtA, i, uint64(m.Timestamp))
	}
	if len(m.Hostname) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintAgentLogsPayload(dAtA, i, uint64(len(m.Hostname)))
		i += copy(dAtA[i:], m.Hostname)
	}
	if len(m.Service) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintAgentLogsPayload(dAtA, i, uint64(len(m.Service)))
		i += copy(dAtA[i:], m.Service)
	}
	if len(m.Source) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintAgentLogsPayload(dAtA, i, uint64(len(m.Source)))
		i += copy(dAtA[i:], m.Source)
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			dAtA[i] = 0x3a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintAgentLogsPayload(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Log) Size() (n int) {
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovAgentLogsPayload(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovAgentLogsPayload(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovAgentLogsPayload(uint64(m.Timestamp))
	}
	l = len(m.Hostname)
	if l > 0 {
		n += 1 + l + sovAgentLogsPayload(uint64(l))
	}
	l = len(m.Service)
	if l > 0 {
		n += 1 + l + sovAgentLogsPayload(uint64(l))
	}
	l = len(m.Source)
	if l > 0 {
		n += 1 + l + sovAgentLogsPayload(uint64(l))
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovAgentLogsPayload(uint64(l))
		}
	}
	return n
}

func sovAgentLogsPayload(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAgentLogsPayload(x uint64) (n int) {
	return sovAgentLogsPayload(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Log) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAgentLogsPayload
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
			return fmt.Errorf("proto: Log: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Log: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgentLogsPayload
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
				return ErrInvalidLengthAgentLogsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgentLogsPayload
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
				return ErrInvalidLengthAgentLogsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgentLogsPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hostname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgentLogsPayload
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
				return ErrInvalidLengthAgentLogsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hostname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgentLogsPayload
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
				return ErrInvalidLengthAgentLogsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Service = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgentLogsPayload
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
				return ErrInvalidLengthAgentLogsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Source = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgentLogsPayload
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
				return ErrInvalidLengthAgentLogsPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAgentLogsPayload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAgentLogsPayload
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
func skipAgentLogsPayload(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAgentLogsPayload
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
					return 0, ErrIntOverflowAgentLogsPayload
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
					return 0, ErrIntOverflowAgentLogsPayload
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
				return 0, ErrInvalidLengthAgentLogsPayload
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAgentLogsPayload
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
				next, err := skipAgentLogsPayload(dAtA[start:])
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
	ErrInvalidLengthAgentLogsPayload = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAgentLogsPayload   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("agent_logs_payload.proto", fileDescriptorAgentLogsPayload) }

var fileDescriptorAgentLogsPayload = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x99, 0xa6, 0xa6, 0x66, 0x10, 0x84, 0x59, 0xc8, 0x20, 0x12, 0x82, 0x6e, 0xb2, 0x31,
	0x59, 0xf8, 0x04, 0x96, 0x2e, 0x5d, 0x94, 0x2c, 0xdd, 0x94, 0x3b, 0xc9, 0x30, 0x0d, 0x74, 0x7a,
	0x87, 0xdc, 0x1b, 0xc1, 0x27, 0xf3, 0x15, 0x5c, 0xfa, 0x08, 0x92, 0x27, 0x91, 0x4e, 0xa3, 0xee,
	0xce, 0x37, 0x3f, 0x7c, 0x9c, 0x23, 0x35, 0x38, 0x7b, 0xe4, 0xdd, 0x01, 0x1d, 0xed, 0x02, 0xbc,
	0x1f, 0x10, 0xba, 0x2a, 0x0c, 0xc8, 0xa8, 0x16, 0xc1, 0xdc, 0x7f, 0x08, 0x99, 0xbc, 0xa0, 0x53,
	0x5a, 0xae, 0xbc, 0x25, 0x02, 0x67, 0xb5, 0x28, 0x44, 0x99, 0x35, 0xbf, 0xa8, 0x6e, 0x64, 0x4a,
	0x0c, 0x3c, 0x92, 0x5e, 0xc4, 0x8b, 0x99, 0xd4, 0x9d, 0xcc, 0xb8, 0xf7, 0x96, 0x18, 0x7c, 0xd0,
	0x49, 0x21, 0xca, 0xa4, 0xf9, 0x3f, 0x50, 0xb7, 0xf2, 0x72, 0x8f, 0xc4, 0x47, 0xf0, 0x56, 0x2f,
	0xe3, 0xbf, 0x3f, 0x3e, 0xb9, 0xc8, 0x0e, 0x6f, 0x7d, 0x6b, 0xf5, 0xc5, 0xd9, 0x35, 0x63, 0x74,
	0xe1, 0x38, 0xb4, 0x56, 0xa7, 0xb3, 0x2b, 0x92, 0x52, 0x72, 0xc9, 0xe0, 0x48, 0xaf, 0x8a, 0xa4,
	0xcc, 0x9a, 0x98, 0xd7, 0xcd, 0xe7, 0x94, 0x8b, 0xaf, 0x29, 0x17, 0xdf, 0x53, 0x2e, 0xe4, 0x75,
	0x8b, 0xbe, 0xea, 0xba, 0x2a, 0x96, 0xad, 0x82, 0x59, 0x5f, 0x3d, 0x9f, 0xd2, 0xf6, 0x5c, 0x78,
	0x2b, 0x5e, 0x1f, 0x5c, 0xcf, 0xfb, 0xd1, 0x54, 0x2d, 0xfa, 0x7a, 0x03, 0x0c, 0x1b, 0x74, 0x75,
	0x7c, 0xfc, 0x38, 0x8f, 0x52, 0x07, 0x63, 0xd2, 0x38, 0xcc, 0xd3, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x11, 0x4f, 0x88, 0xb8, 0x34, 0x01, 0x00, 0x00,
}