// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: src/vizier/services/metadata/storepb/store.proto

package storepb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	proto1 "pixielabs.ai/pixielabs/src/common/base/proto"
	logical "pixielabs.ai/pixielabs/src/stirling/dynamic_tracing/ir/logical"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ProbeInfo struct {
	ProbeID string           `protobuf:"bytes,1,opt,name=probe_id,json=probeId,proto3" json:"probe_id,omitempty"`
	Program *logical.Program `protobuf:"bytes,2,opt,name=program,proto3" json:"program,omitempty"`
}

func (m *ProbeInfo) Reset()      { *m = ProbeInfo{} }
func (*ProbeInfo) ProtoMessage() {}
func (*ProbeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_27ea71ea705227d1, []int{0}
}
func (m *ProbeInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProbeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProbeInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProbeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeInfo.Merge(m, src)
}
func (m *ProbeInfo) XXX_Size() int {
	return m.Size()
}
func (m *ProbeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeInfo proto.InternalMessageInfo

func (m *ProbeInfo) GetProbeID() string {
	if m != nil {
		return m.ProbeID
	}
	return ""
}

func (m *ProbeInfo) GetProgram() *logical.Program {
	if m != nil {
		return m.Program
	}
	return nil
}

type AgentProbeStatus struct {
	State   proto1.LifeCycleState `protobuf:"varint,1,opt,name=state,proto3,enum=pl.statuspb.LifeCycleState" json:"state,omitempty"`
	Status  *proto1.Status        `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	ProbeID string                `protobuf:"bytes,3,opt,name=probe_id,json=probeId,proto3" json:"probe_id,omitempty"`
	AgentID string                `protobuf:"bytes,4,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
}

func (m *AgentProbeStatus) Reset()      { *m = AgentProbeStatus{} }
func (*AgentProbeStatus) ProtoMessage() {}
func (*AgentProbeStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_27ea71ea705227d1, []int{1}
}
func (m *AgentProbeStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AgentProbeStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AgentProbeStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AgentProbeStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentProbeStatus.Merge(m, src)
}
func (m *AgentProbeStatus) XXX_Size() int {
	return m.Size()
}
func (m *AgentProbeStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentProbeStatus.DiscardUnknown(m)
}

var xxx_messageInfo_AgentProbeStatus proto.InternalMessageInfo

func (m *AgentProbeStatus) GetState() proto1.LifeCycleState {
	if m != nil {
		return m.State
	}
	return proto1.UNKNOWN_STATE
}

func (m *AgentProbeStatus) GetStatus() *proto1.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AgentProbeStatus) GetProbeID() string {
	if m != nil {
		return m.ProbeID
	}
	return ""
}

func (m *AgentProbeStatus) GetAgentID() string {
	if m != nil {
		return m.AgentID
	}
	return ""
}

func init() {
	proto.RegisterType((*ProbeInfo)(nil), "pl.vizier.services.metadata.ProbeInfo")
	proto.RegisterType((*AgentProbeStatus)(nil), "pl.vizier.services.metadata.AgentProbeStatus")
}

func init() {
	proto.RegisterFile("src/vizier/services/metadata/storepb/store.proto", fileDescriptor_27ea71ea705227d1)
}

var fileDescriptor_27ea71ea705227d1 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0xae, 0xd3, 0x30,
	0x14, 0xc6, 0x63, 0xfe, 0xdc, 0x3f, 0xbe, 0x12, 0x42, 0x61, 0xb9, 0xba, 0x95, 0x4c, 0xd5, 0xa1,
	0xaa, 0x84, 0xb0, 0xa1, 0x8c, 0x4c, 0xb4, 0x2c, 0x91, 0x18, 0xaa, 0xb0, 0xb1, 0x54, 0xb6, 0xe3,
	0x06, 0x4b, 0x49, 0x1c, 0xd9, 0x6e, 0x45, 0x19, 0x10, 0x8f, 0xc0, 0x63, 0xf0, 0x1c, 0x4c, 0x8c,
	0x1d, 0x3b, 0x21, 0xea, 0x2e, 0x8c, 0x7d, 0x04, 0x64, 0xbb, 0x01, 0x95, 0x01, 0xdd, 0xc9, 0x3e,
	0x3a, 0xbf, 0xef, 0x3b, 0xdf, 0x49, 0x0c, 0x9f, 0x19, 0xcd, 0xc9, 0x4a, 0x7e, 0x94, 0x42, 0x13,
	0x23, 0xf4, 0x4a, 0x72, 0x61, 0x48, 0x2d, 0x2c, 0x2d, 0xa8, 0xa5, 0xc4, 0x58, 0xa5, 0x45, 0xcb,
	0xe2, 0x89, 0x5b, 0xad, 0xac, 0x4a, 0x7b, 0x6d, 0x85, 0xa3, 0x00, 0x77, 0x02, 0xdc, 0x09, 0x6e,
	0x06, 0xde, 0x8e, 0xab, 0xba, 0x56, 0x0d, 0x61, 0xd4, 0x08, 0x12, 0x34, 0xc4, 0x58, 0x6a, 0x97,
	0x26, 0x1a, 0xdc, 0x3c, 0x2d, 0xa5, 0x7d, 0xbf, 0x64, 0x98, 0xab, 0x9a, 0x94, 0xaa, 0x54, 0x91,
	0x61, 0xcb, 0x45, 0xa8, 0xa2, 0xc0, 0xdf, 0x3a, 0xdc, 0x5b, 0x1a, 0x2b, 0x75, 0x25, 0x9b, 0x92,
	0x14, 0xeb, 0x86, 0xd6, 0x92, 0xcf, 0xad, 0xa6, 0xdc, 0xd7, 0x52, 0x93, 0x4a, 0x95, 0x92, 0xd3,
	0x2a, 0xe2, 0x83, 0x4f, 0xf0, 0x72, 0xa6, 0x15, 0x13, 0x59, 0xb3, 0x50, 0xe9, 0x10, 0x5e, 0xb4,
	0xbe, 0x98, 0xcb, 0xe2, 0x1a, 0xf4, 0xc1, 0xe8, 0x72, 0x72, 0xe5, 0x7e, 0x3c, 0x3e, 0x8f, 0xc0,
	0xeb, 0xfc, 0x3c, 0x34, 0xb3, 0x22, 0xcd, 0xa0, 0xbf, 0x96, 0x9a, 0xd6, 0xd7, 0x77, 0xfa, 0x60,
	0x74, 0x35, 0x26, 0xb8, 0xad, 0x70, 0x37, 0x14, 0xff, 0x33, 0x14, 0x4b, 0x8d, 0xbb, 0xa1, 0xb3,
	0x28, 0xcb, 0x3b, 0xfd, 0xe0, 0x1b, 0x80, 0x0f, 0x5f, 0x95, 0xa2, 0xb1, 0x61, 0xc8, 0xdb, 0xb0,
	0x78, 0xfa, 0x1c, 0xde, 0xf7, 0x9f, 0x40, 0x84, 0x10, 0x0f, 0xc6, 0xbd, 0xe8, 0xee, 0x5b, 0x2d,
	0xc3, 0x6f, 0xe4, 0x42, 0x4c, 0xd7, 0xbc, 0x0a, 0xb0, 0xc8, 0x23, 0x99, 0x3e, 0x81, 0x67, 0x91,
	0x38, 0x26, 0x7a, 0x74, 0xa2, 0x89, 0xbe, 0xf9, 0x11, 0x39, 0xd9, 0xf3, 0xee, 0x7f, 0xf6, 0x1c,
	0xc2, 0x0b, 0xea, 0xb3, 0x79, 0xee, 0xde, 0x5f, 0x2e, 0xe4, 0xf5, 0x5c, 0x68, 0x66, 0xc5, 0x64,
	0xbd, 0xd9, 0xa1, 0x64, 0xbb, 0x43, 0xc9, 0x61, 0x87, 0xc0, 0x67, 0x87, 0xc0, 0x57, 0x87, 0xc0,
	0x77, 0x87, 0xc0, 0xc6, 0x21, 0xf0, 0xd3, 0x21, 0xf0, 0xcb, 0xa1, 0xe4, 0xe0, 0x10, 0xf8, 0xb2,
	0x47, 0xc9, 0x66, 0x8f, 0x92, 0xed, 0x1e, 0x25, 0xef, 0xa6, 0xad, 0xfc, 0x20, 0x45, 0x45, 0x99,
	0xc1, 0x54, 0x92, 0x3f, 0x05, 0xb9, 0xcd, 0x13, 0x7b, 0x79, 0x3c, 0xd9, 0x59, 0xf8, 0x8d, 0x2f,
	0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x2b, 0xd8, 0x75, 0x99, 0x02, 0x00, 0x00,
}

func (this *ProbeInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProbeInfo)
	if !ok {
		that2, ok := that.(ProbeInfo)
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
	if this.ProbeID != that1.ProbeID {
		return false
	}
	if !this.Program.Equal(that1.Program) {
		return false
	}
	return true
}
func (this *AgentProbeStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AgentProbeStatus)
	if !ok {
		that2, ok := that.(AgentProbeStatus)
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
	if this.State != that1.State {
		return false
	}
	if !this.Status.Equal(that1.Status) {
		return false
	}
	if this.ProbeID != that1.ProbeID {
		return false
	}
	if this.AgentID != that1.AgentID {
		return false
	}
	return true
}
func (this *ProbeInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&storepb.ProbeInfo{")
	s = append(s, "ProbeID: "+fmt.Sprintf("%#v", this.ProbeID)+",\n")
	if this.Program != nil {
		s = append(s, "Program: "+fmt.Sprintf("%#v", this.Program)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *AgentProbeStatus) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&storepb.AgentProbeStatus{")
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	if this.Status != nil {
		s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	}
	s = append(s, "ProbeID: "+fmt.Sprintf("%#v", this.ProbeID)+",\n")
	s = append(s, "AgentID: "+fmt.Sprintf("%#v", this.AgentID)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringStore(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ProbeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProbeInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProbeInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Program != nil {
		{
			size, err := m.Program.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStore(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ProbeID) > 0 {
		i -= len(m.ProbeID)
		copy(dAtA[i:], m.ProbeID)
		i = encodeVarintStore(dAtA, i, uint64(len(m.ProbeID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AgentProbeStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AgentProbeStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AgentProbeStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AgentID) > 0 {
		i -= len(m.AgentID)
		copy(dAtA[i:], m.AgentID)
		i = encodeVarintStore(dAtA, i, uint64(len(m.AgentID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ProbeID) > 0 {
		i -= len(m.ProbeID)
		copy(dAtA[i:], m.ProbeID)
		i = encodeVarintStore(dAtA, i, uint64(len(m.ProbeID)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Status != nil {
		{
			size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStore(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.State != 0 {
		i = encodeVarintStore(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStore(dAtA []byte, offset int, v uint64) int {
	offset -= sovStore(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProbeInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ProbeID)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	if m.Program != nil {
		l = m.Program.Size()
		n += 1 + l + sovStore(uint64(l))
	}
	return n
}

func (m *AgentProbeStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.State != 0 {
		n += 1 + sovStore(uint64(m.State))
	}
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovStore(uint64(l))
	}
	l = len(m.ProbeID)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	l = len(m.AgentID)
	if l > 0 {
		n += 1 + l + sovStore(uint64(l))
	}
	return n
}

func sovStore(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStore(x uint64) (n int) {
	return sovStore(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ProbeInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProbeInfo{`,
		`ProbeID:` + fmt.Sprintf("%v", this.ProbeID) + `,`,
		`Program:` + strings.Replace(fmt.Sprintf("%v", this.Program), "Program", "logical.Program", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AgentProbeStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AgentProbeStatus{`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`Status:` + strings.Replace(fmt.Sprintf("%v", this.Status), "Status", "proto1.Status", 1) + `,`,
		`ProbeID:` + fmt.Sprintf("%v", this.ProbeID) + `,`,
		`AgentID:` + fmt.Sprintf("%v", this.AgentID) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringStore(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ProbeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: ProbeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProbeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProbeID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProbeID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Program", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Program == nil {
				m.Program = &logical.Program{}
			}
			if err := m.Program.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStore
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStore
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
func (m *AgentProbeStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
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
			return fmt.Errorf("proto: AgentProbeStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AgentProbeStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= proto1.LifeCycleState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &proto1.Status{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProbeID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProbeID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgentID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
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
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStore
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AgentID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStore(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStore
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStore
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
func skipStore(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStore
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
					return 0, ErrIntOverflowStore
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStore
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
				return 0, ErrInvalidLengthStore
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStore
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStore
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStore        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStore          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStore = fmt.Errorf("proto: unexpected end of group")
)
