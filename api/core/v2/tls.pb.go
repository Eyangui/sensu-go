// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tls.proto

package v2

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	io "io"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// TLSOptions holds TLS options that are used across the varying Sensu
// components
type TLSOptions struct {
	CertFile             string   `protobuf:"bytes,1,opt,name=cert_file,json=certFile,proto3" json:"cert_file,omitempty"`
	KeyFile              string   `protobuf:"bytes,2,opt,name=key_file,json=keyFile,proto3" json:"key_file,omitempty"`
	TrustedCAFile        string   `protobuf:"bytes,3,opt,name=trusted_ca_file,json=trustedCaFile,proto3" json:"trusted_ca_file,omitempty"`
	InsecureSkipVerify   bool     `protobuf:"varint,4,opt,name=insecure_skip_verify,json=insecureSkipVerify,proto3" json:"insecure_skip_verify"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TLSOptions) Reset()         { *m = TLSOptions{} }
func (m *TLSOptions) String() string { return proto.CompactTextString(m) }
func (*TLSOptions) ProtoMessage()    {}
func (*TLSOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_adf82c87377d3c77, []int{0}
}
func (m *TLSOptions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TLSOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TLSOptions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TLSOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLSOptions.Merge(m, src)
}
func (m *TLSOptions) XXX_Size() int {
	return m.Size()
}
func (m *TLSOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_TLSOptions.DiscardUnknown(m)
}

var xxx_messageInfo_TLSOptions proto.InternalMessageInfo

func (m *TLSOptions) GetCertFile() string {
	if m != nil {
		return m.CertFile
	}
	return ""
}

func (m *TLSOptions) GetKeyFile() string {
	if m != nil {
		return m.KeyFile
	}
	return ""
}

func (m *TLSOptions) GetTrustedCAFile() string {
	if m != nil {
		return m.TrustedCAFile
	}
	return ""
}

func (m *TLSOptions) GetInsecureSkipVerify() bool {
	if m != nil {
		return m.InsecureSkipVerify
	}
	return false
}

func init() {
	proto.RegisterType((*TLSOptions)(nil), "sensu.core.v2.TLSOptions")
}

func init() { proto.RegisterFile("tls.proto", fileDescriptor_adf82c87377d3c77) }

var fileDescriptor_adf82c87377d3c77 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0xc9, 0x29, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2d, 0x4e, 0xcd, 0x2b, 0x2e, 0xd5, 0x4b, 0xce, 0x2f,
	0x4a, 0xd5, 0x2b, 0x33, 0x92, 0x32, 0x49, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x4a, 0x2a, 0x4d, 0x73, 0x28, 0x33, 0xd4, 0x33,
	0xd2, 0x33, 0x04, 0x0b, 0x82, 0xc5, 0xc0, 0x2c, 0x88, 0x21, 0x4a, 0xa7, 0x19, 0xb9, 0xb8, 0x42,
	0x7c, 0x82, 0xfd, 0x0b, 0x4a, 0x32, 0xf3, 0xf3, 0x8a, 0x85, 0xa4, 0xb9, 0x38, 0x93, 0x53, 0x8b,
	0x4a, 0xe2, 0xd3, 0x32, 0x73, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x38, 0x40, 0x02,
	0x6e, 0x99, 0x39, 0xa9, 0x42, 0x92, 0x5c, 0x1c, 0xd9, 0xa9, 0x95, 0x10, 0x39, 0x26, 0xb0, 0x1c,
	0x7b, 0x76, 0x6a, 0x25, 0x58, 0xca, 0x92, 0x8b, 0xbf, 0xa4, 0xa8, 0xb4, 0xb8, 0x24, 0x35, 0x25,
	0x3e, 0x39, 0x11, 0xa2, 0x82, 0x19, 0xa4, 0xc2, 0x49, 0xf0, 0xd1, 0x3d, 0x79, 0xde, 0x10, 0x88,
	0x94, 0xb3, 0x23, 0x48, 0x6d, 0x10, 0x2f, 0x54, 0xa5, 0x73, 0x22, 0x58, 0xab, 0x17, 0x97, 0x48,
	0x66, 0x5e, 0x71, 0x6a, 0x72, 0x69, 0x51, 0x6a, 0x7c, 0x71, 0x76, 0x66, 0x41, 0x7c, 0x59, 0x6a,
	0x51, 0x66, 0x5a, 0xa5, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0x87, 0x93, 0xc4, 0xab, 0x7b, 0xf2, 0x58,
	0xe5, 0x83, 0x84, 0x60, 0xa2, 0xc1, 0xd9, 0x99, 0x05, 0x61, 0x60, 0x31, 0x27, 0x85, 0x1f, 0x0f,
	0xe5, 0x18, 0x57, 0x3c, 0x92, 0x63, 0xdc, 0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x8a, 0xa9, 0xcc, 0x28,
	0x89, 0x0d, 0xec, 0x6d, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xef, 0x29, 0x21, 0x48,
	0x01, 0x00, 0x00,
}

func (this *TLSOptions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TLSOptions)
	if !ok {
		that2, ok := that.(TLSOptions)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CertFile != that1.CertFile {
		return false
	}
	if this.KeyFile != that1.KeyFile {
		return false
	}
	if this.TrustedCAFile != that1.TrustedCAFile {
		return false
	}
	if this.InsecureSkipVerify != that1.InsecureSkipVerify {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *TLSOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TLSOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.CertFile) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTls(dAtA, i, uint64(len(m.CertFile)))
		i += copy(dAtA[i:], m.CertFile)
	}
	if len(m.KeyFile) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTls(dAtA, i, uint64(len(m.KeyFile)))
		i += copy(dAtA[i:], m.KeyFile)
	}
	if len(m.TrustedCAFile) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTls(dAtA, i, uint64(len(m.TrustedCAFile)))
		i += copy(dAtA[i:], m.TrustedCAFile)
	}
	if m.InsecureSkipVerify {
		dAtA[i] = 0x20
		i++
		if m.InsecureSkipVerify {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintTls(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedTLSOptions(r randyTls, easy bool) *TLSOptions {
	this := &TLSOptions{}
	this.CertFile = string(randStringTls(r))
	this.KeyFile = string(randStringTls(r))
	this.TrustedCAFile = string(randStringTls(r))
	this.InsecureSkipVerify = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTls(r, 5)
	}
	return this
}

type randyTls interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTls(r randyTls) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTls(r randyTls) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneTls(r)
	}
	return string(tmps)
}
func randUnrecognizedTls(r randyTls, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTls(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTls(dAtA []byte, r randyTls, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTls(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateTls(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateTls(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTls(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTls(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTls(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTls(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *TLSOptions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CertFile)
	if l > 0 {
		n += 1 + l + sovTls(uint64(l))
	}
	l = len(m.KeyFile)
	if l > 0 {
		n += 1 + l + sovTls(uint64(l))
	}
	l = len(m.TrustedCAFile)
	if l > 0 {
		n += 1 + l + sovTls(uint64(l))
	}
	if m.InsecureSkipVerify {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTls(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTls(x uint64) (n int) {
	return sovTls(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TLSOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTls
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TLSOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TLSOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertFile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertFile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyFile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyFile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustedCAFile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTls
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTls
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrustedCAFile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InsecureSkipVerify", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTls
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InsecureSkipVerify = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTls(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTls
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTls
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTls(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTls
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
					return 0, ErrIntOverflowTls
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
					return 0, ErrIntOverflowTls
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
			if length < 0 {
				return 0, ErrInvalidLengthTls
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTls
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTls
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
				next, err := skipTls(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTls
				}
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
	ErrInvalidLengthTls = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTls   = fmt.Errorf("proto: integer overflow")
)
