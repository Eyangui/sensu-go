// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hook.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import encoding_binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A HookRequest represents a request to execute a hook
type HookRequest struct {
	// Config is the specification of a hook.
	Config *HookConfig `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
}

func (m *HookRequest) Reset()                    { *m = HookRequest{} }
func (m *HookRequest) String() string            { return proto.CompactTextString(m) }
func (*HookRequest) ProtoMessage()               {}
func (*HookRequest) Descriptor() ([]byte, []int) { return fileDescriptorHook, []int{0} }

func (m *HookRequest) GetConfig() *HookConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

// HookConfig is the specification of a hook
type HookConfig struct {
	// Name is the unique identifier for a hook
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Command is the command to be executed
	Command string `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	// Timeout is the timeout, in seconds, at which the hook has to run
	Timeout uint32 `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Stdin indicates if hook requests have stdin enabled
	Stdin bool `protobuf:"varint,4,opt,name=stdin,proto3" json:"stdin,omitempty"`
	// Environment indicates to which env a hook belongs to
	Environment string `protobuf:"bytes,5,opt,name=environment,proto3" json:"environment,omitempty"`
	// Organization indicates to which org a hook belongs to
	Organization string `protobuf:"bytes,6,opt,name=organization,proto3" json:"organization,omitempty"`
}

func (m *HookConfig) Reset()                    { *m = HookConfig{} }
func (m *HookConfig) String() string            { return proto.CompactTextString(m) }
func (*HookConfig) ProtoMessage()               {}
func (*HookConfig) Descriptor() ([]byte, []int) { return fileDescriptorHook, []int{1} }

func (m *HookConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HookConfig) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *HookConfig) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *HookConfig) GetStdin() bool {
	if m != nil {
		return m.Stdin
	}
	return false
}

func (m *HookConfig) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *HookConfig) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

// A Hook is a hook specification and optionally the results of the hook's
// execution.
type Hook struct {
	// Config is the specification of a hook
	Config *HookConfig `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	// Duration of execution
	Duration float64 `protobuf:"fixed64,2,opt,name=duration,proto3" json:"duration,omitempty"`
	// Executed describes the time in which the hook request was executed
	Executed int64 `protobuf:"varint,3,opt,name=executed,proto3" json:"executed,omitempty"`
	// Issued describes the time in which the hook request was issued
	Issued int64 `protobuf:"varint,4,opt,name=issued,proto3" json:"issued,omitempty"`
	// Output from the execution of Command
	Output string `protobuf:"bytes,5,opt,name=output,proto3" json:"output,omitempty"`
	// Status is the exit status code produced by the hook
	Status int32 `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *Hook) Reset()                    { *m = Hook{} }
func (m *Hook) String() string            { return proto.CompactTextString(m) }
func (*Hook) ProtoMessage()               {}
func (*Hook) Descriptor() ([]byte, []int) { return fileDescriptorHook, []int{2} }

func (m *Hook) GetConfig() *HookConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Hook) GetDuration() float64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Hook) GetExecuted() int64 {
	if m != nil {
		return m.Executed
	}
	return 0
}

func (m *Hook) GetIssued() int64 {
	if m != nil {
		return m.Issued
	}
	return 0
}

func (m *Hook) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *Hook) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*HookRequest)(nil), "sensu.types.HookRequest")
	proto.RegisterType((*HookConfig)(nil), "sensu.types.HookConfig")
	proto.RegisterType((*Hook)(nil), "sensu.types.Hook")
}
func (this *HookRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*HookRequest)
	if !ok {
		that2, ok := that.(HookRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Config.Equal(that1.Config) {
		return false
	}
	return true
}
func (this *HookConfig) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*HookConfig)
	if !ok {
		that2, ok := that.(HookConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Command != that1.Command {
		return false
	}
	if this.Timeout != that1.Timeout {
		return false
	}
	if this.Stdin != that1.Stdin {
		return false
	}
	if this.Environment != that1.Environment {
		return false
	}
	if this.Organization != that1.Organization {
		return false
	}
	return true
}
func (this *Hook) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Hook)
	if !ok {
		that2, ok := that.(Hook)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Config.Equal(that1.Config) {
		return false
	}
	if this.Duration != that1.Duration {
		return false
	}
	if this.Executed != that1.Executed {
		return false
	}
	if this.Issued != that1.Issued {
		return false
	}
	if this.Output != that1.Output {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	return true
}
func (m *HookRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HookRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Config != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHook(dAtA, i, uint64(m.Config.Size()))
		n1, err := m.Config.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *HookConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HookConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHook(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Command) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHook(dAtA, i, uint64(len(m.Command)))
		i += copy(dAtA[i:], m.Command)
	}
	if m.Timeout != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintHook(dAtA, i, uint64(m.Timeout))
	}
	if m.Stdin {
		dAtA[i] = 0x20
		i++
		if m.Stdin {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Environment) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintHook(dAtA, i, uint64(len(m.Environment)))
		i += copy(dAtA[i:], m.Environment)
	}
	if len(m.Organization) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintHook(dAtA, i, uint64(len(m.Organization)))
		i += copy(dAtA[i:], m.Organization)
	}
	return i, nil
}

func (m *Hook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Hook) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Config != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHook(dAtA, i, uint64(m.Config.Size()))
		n2, err := m.Config.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Duration != 0 {
		dAtA[i] = 0x11
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Duration))))
		i += 8
	}
	if m.Executed != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintHook(dAtA, i, uint64(m.Executed))
	}
	if m.Issued != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintHook(dAtA, i, uint64(m.Issued))
	}
	if len(m.Output) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintHook(dAtA, i, uint64(len(m.Output)))
		i += copy(dAtA[i:], m.Output)
	}
	if m.Status != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintHook(dAtA, i, uint64(m.Status))
	}
	return i, nil
}

func encodeVarintHook(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedHookRequest(r randyHook, easy bool) *HookRequest {
	this := &HookRequest{}
	if r.Intn(10) != 0 {
		this.Config = NewPopulatedHookConfig(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedHookConfig(r randyHook, easy bool) *HookConfig {
	this := &HookConfig{}
	this.Name = string(randStringHook(r))
	this.Command = string(randStringHook(r))
	this.Timeout = uint32(r.Uint32())
	this.Stdin = bool(bool(r.Intn(2) == 0))
	this.Environment = string(randStringHook(r))
	this.Organization = string(randStringHook(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedHook(r randyHook, easy bool) *Hook {
	this := &Hook{}
	if r.Intn(10) != 0 {
		this.Config = NewPopulatedHookConfig(r, easy)
	}
	this.Duration = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.Duration *= -1
	}
	this.Executed = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Executed *= -1
	}
	this.Issued = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Issued *= -1
	}
	this.Output = string(randStringHook(r))
	this.Status = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Status *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyHook interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneHook(r randyHook) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringHook(r randyHook) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneHook(r)
	}
	return string(tmps)
}
func randUnrecognizedHook(r randyHook, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldHook(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldHook(dAtA []byte, r randyHook, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateHook(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateHook(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateHook(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateHook(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateHook(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateHook(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateHook(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *HookRequest) Size() (n int) {
	var l int
	_ = l
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovHook(uint64(l))
	}
	return n
}

func (m *HookConfig) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovHook(uint64(l))
	}
	l = len(m.Command)
	if l > 0 {
		n += 1 + l + sovHook(uint64(l))
	}
	if m.Timeout != 0 {
		n += 1 + sovHook(uint64(m.Timeout))
	}
	if m.Stdin {
		n += 2
	}
	l = len(m.Environment)
	if l > 0 {
		n += 1 + l + sovHook(uint64(l))
	}
	l = len(m.Organization)
	if l > 0 {
		n += 1 + l + sovHook(uint64(l))
	}
	return n
}

func (m *Hook) Size() (n int) {
	var l int
	_ = l
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovHook(uint64(l))
	}
	if m.Duration != 0 {
		n += 9
	}
	if m.Executed != 0 {
		n += 1 + sovHook(uint64(m.Executed))
	}
	if m.Issued != 0 {
		n += 1 + sovHook(uint64(m.Issued))
	}
	l = len(m.Output)
	if l > 0 {
		n += 1 + l + sovHook(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovHook(uint64(m.Status))
	}
	return n
}

func sovHook(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHook(x uint64) (n int) {
	return sovHook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HookRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHook
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
			return fmt.Errorf("proto: HookRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HookRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
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
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &HookConfig{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHook
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
func (m *HookConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHook
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
			return fmt.Errorf("proto: HookConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HookConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
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
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Command", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
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
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Command = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stdin", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
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
			m.Stdin = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Environment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
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
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Environment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Organization", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
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
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Organization = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHook
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
func (m *Hook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHook
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
			return fmt.Errorf("proto: Hook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Hook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
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
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &HookConfig{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Duration = float64(math.Float64frombits(v))
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Executed", wireType)
			}
			m.Executed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Executed |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issued", wireType)
			}
			m.Issued = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Issued |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Output", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
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
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Output = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHook
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
func skipHook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHook
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
					return 0, ErrIntOverflowHook
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
					return 0, ErrIntOverflowHook
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
				return 0, ErrInvalidLengthHook
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHook
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
				next, err := skipHook(dAtA[start:])
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
	ErrInvalidLengthHook = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHook   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("hook.proto", fileDescriptorHook) }

var fileDescriptorHook = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0x7f, 0xff, 0x4d, 0x42, 0x71, 0x60, 0xb1, 0x10, 0x44, 0x1d, 0x42, 0x14, 0x96, 0x2e,
	0xa4, 0x12, 0x88, 0x17, 0x28, 0x0c, 0xcc, 0x1e, 0xd9, 0xd2, 0xc4, 0x4d, 0xad, 0x2a, 0xbe, 0x25,
	0xb6, 0x11, 0xf0, 0x24, 0x3c, 0x02, 0x23, 0x23, 0x0b, 0x7b, 0x47, 0x9e, 0x00, 0x41, 0x78, 0x09,
	0x46, 0x94, 0x9b, 0xb4, 0x94, 0x95, 0xed, 0x7e, 0xe7, 0x9c, 0xdc, 0xdc, 0x23, 0x53, 0x3a, 0x03,
	0x98, 0x27, 0x8b, 0x0a, 0x0c, 0x30, 0x5f, 0x0b, 0xa5, 0x6d, 0x62, 0xee, 0x16, 0x42, 0x0f, 0x8e,
	0x0b, 0x69, 0x66, 0x76, 0x92, 0x64, 0x50, 0x8e, 0x0a, 0x28, 0x60, 0x84, 0x99, 0x89, 0x9d, 0x22,
	0x21, 0xe0, 0xd4, 0x7e, 0x3b, 0xf0, 0x53, 0xad, 0x85, 0x69, 0x21, 0xbe, 0xa0, 0xfe, 0x25, 0xc0,
	0x9c, 0x8b, 0x6b, 0x2b, 0xb4, 0x61, 0x67, 0xd4, 0xcb, 0x40, 0x4d, 0x65, 0x11, 0x90, 0x88, 0x0c,
	0xfd, 0x93, 0x83, 0x64, 0xe3, 0x47, 0x49, 0x93, 0x3c, 0x47, 0x7b, 0xec, 0x2c, 0xdf, 0x0e, 0x09,
	0xef, 0xc2, 0xf1, 0x13, 0xa1, 0xf4, 0xc7, 0x64, 0x8c, 0x3a, 0x2a, 0x2d, 0x05, 0xee, 0xd8, 0xe6,
	0x38, 0xb3, 0x80, 0x6e, 0x65, 0x50, 0x96, 0xa9, 0xca, 0x83, 0xff, 0x28, 0xaf, 0xb0, 0x71, 0x8c,
	0x2c, 0x05, 0x58, 0x13, 0xf4, 0x22, 0x32, 0xdc, 0xe5, 0x2b, 0x64, 0x7b, 0xd4, 0xd5, 0x26, 0x97,
	0x2a, 0x70, 0x22, 0x32, 0xec, 0xf3, 0x16, 0x58, 0x44, 0x7d, 0xa1, 0x6e, 0x64, 0x05, 0xaa, 0x14,
	0xca, 0x04, 0x2e, 0x6e, 0xdb, 0x94, 0x58, 0x4c, 0x77, 0xa0, 0x2a, 0x52, 0x25, 0xef, 0x53, 0x23,
	0x41, 0x05, 0x1e, 0x46, 0x7e, 0x69, 0xf1, 0x0b, 0xa1, 0x4e, 0x73, 0xf2, 0x1f, 0x2b, 0xb3, 0x01,
	0xed, 0xe7, 0xb6, 0x6a, 0xf7, 0x37, 0x85, 0x08, 0x5f, 0x73, 0xe3, 0x89, 0x5b, 0x91, 0x59, 0x23,
	0x72, 0xac, 0xd4, 0xe3, 0x6b, 0x66, 0xfb, 0xd4, 0x93, 0x5a, 0x5b, 0x91, 0x63, 0xa9, 0x1e, 0xef,
	0xa8, 0xd1, 0xc1, 0x9a, 0x85, 0x5d, 0x15, 0xea, 0xa8, 0xd1, 0xb5, 0x49, 0x8d, 0xd5, 0xd8, 0xc2,
	0xe5, 0x1d, 0x8d, 0x8f, 0xbe, 0x3e, 0x42, 0xf2, 0x58, 0x87, 0xe4, 0xb9, 0x0e, 0xc9, 0xb2, 0x0e,
	0xc9, 0x6b, 0x1d, 0x92, 0xf7, 0x3a, 0x24, 0x0f, 0x9f, 0xe1, 0xbf, 0x2b, 0x17, 0xaf, 0x9f, 0x78,
	0xf8, 0xc8, 0xa7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x84, 0xec, 0xa3, 0xf6, 0x3b, 0x02, 0x00,
	0x00,
}
