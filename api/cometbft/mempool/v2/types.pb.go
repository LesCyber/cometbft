// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cometbft/mempool/v2/types.proto

package v2

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// Txs contains a list of transaction from the mempool.
type Txs struct {
	Txs [][]byte `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (m *Txs) Reset()         { *m = Txs{} }
func (m *Txs) String() string { return proto.CompactTextString(m) }
func (*Txs) ProtoMessage()    {}
func (*Txs) Descriptor() ([]byte, []int) {
	return fileDescriptor_f354aa43d1c2a8af, []int{0}
}
func (m *Txs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Txs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Txs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Txs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Txs.Merge(m, src)
}
func (m *Txs) XXX_Size() int {
	return m.Size()
}
func (m *Txs) XXX_DiscardUnknown() {
	xxx_messageInfo_Txs.DiscardUnknown(m)
}

var xxx_messageInfo_Txs proto.InternalMessageInfo

func (m *Txs) GetTxs() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

// HaveTx is sent by the DOG protocol to signal a peer that the sender already
// has a transaction.
type HaveTx struct {
	TxKey []byte `protobuf:"bytes,1,opt,name=tx_key,json=txKey,proto3" json:"tx_key,omitempty"`
}

func (m *HaveTx) Reset()         { *m = HaveTx{} }
func (m *HaveTx) String() string { return proto.CompactTextString(m) }
func (*HaveTx) ProtoMessage()    {}
func (*HaveTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_f354aa43d1c2a8af, []int{1}
}
func (m *HaveTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HaveTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HaveTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HaveTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HaveTx.Merge(m, src)
}
func (m *HaveTx) XXX_Size() int {
	return m.Size()
}
func (m *HaveTx) XXX_DiscardUnknown() {
	xxx_messageInfo_HaveTx.DiscardUnknown(m)
}

var xxx_messageInfo_HaveTx proto.InternalMessageInfo

func (m *HaveTx) GetTxKey() []byte {
	if m != nil {
		return m.TxKey
	}
	return nil
}

// ResetRoute is sent by the DOG protocol to signal a peer to reset a (random)
// route to the sender.
type ResetRoute struct {
}

func (m *ResetRoute) Reset()         { *m = ResetRoute{} }
func (m *ResetRoute) String() string { return proto.CompactTextString(m) }
func (*ResetRoute) ProtoMessage()    {}
func (*ResetRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_f354aa43d1c2a8af, []int{2}
}
func (m *ResetRoute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResetRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResetRoute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResetRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetRoute.Merge(m, src)
}
func (m *ResetRoute) XXX_Size() int {
	return m.Size()
}
func (m *ResetRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetRoute.DiscardUnknown(m)
}

var xxx_messageInfo_ResetRoute proto.InternalMessageInfo

// Message is an abstract mempool message.
type Message struct {
	// Sum of all possible messages.
	//
	// Types that are valid to be assigned to Sum:
	//
	//	*Message_Txs
	//	*Message_HaveTx
	//	*Message_ResetRoute
	Sum isMessage_Sum `protobuf_oneof:"sum"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_f354aa43d1c2a8af, []int{3}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

type isMessage_Sum interface {
	isMessage_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Message_Txs struct {
	Txs *Txs `protobuf:"bytes,1,opt,name=txs,proto3,oneof" json:"txs,omitempty"`
}
type Message_HaveTx struct {
	HaveTx *HaveTx `protobuf:"bytes,2,opt,name=have_tx,json=haveTx,proto3,oneof" json:"have_tx,omitempty"`
}
type Message_ResetRoute struct {
	ResetRoute *ResetRoute `protobuf:"bytes,3,opt,name=reset_route,json=resetRoute,proto3,oneof" json:"reset_route,omitempty"`
}

func (*Message_Txs) isMessage_Sum()        {}
func (*Message_HaveTx) isMessage_Sum()     {}
func (*Message_ResetRoute) isMessage_Sum() {}

func (m *Message) GetSum() isMessage_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Message) GetTxs() *Txs {
	if x, ok := m.GetSum().(*Message_Txs); ok {
		return x.Txs
	}
	return nil
}

func (m *Message) GetHaveTx() *HaveTx {
	if x, ok := m.GetSum().(*Message_HaveTx); ok {
		return x.HaveTx
	}
	return nil
}

func (m *Message) GetResetRoute() *ResetRoute {
	if x, ok := m.GetSum().(*Message_ResetRoute); ok {
		return x.ResetRoute
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_Txs)(nil),
		(*Message_HaveTx)(nil),
		(*Message_ResetRoute)(nil),
	}
}

func init() {
	proto.RegisterType((*Txs)(nil), "cometbft.mempool.v2.Txs")
	proto.RegisterType((*HaveTx)(nil), "cometbft.mempool.v2.HaveTx")
	proto.RegisterType((*ResetRoute)(nil), "cometbft.mempool.v2.ResetRoute")
	proto.RegisterType((*Message)(nil), "cometbft.mempool.v2.Message")
}

func init() { proto.RegisterFile("cometbft/mempool/v2/types.proto", fileDescriptor_f354aa43d1c2a8af) }

var fileDescriptor_f354aa43d1c2a8af = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0x84, 0x50,
	0x14, 0x86, 0xef, 0x4d, 0xc6, 0x81, 0x33, 0xb3, 0x08, 0x23, 0x12, 0x82, 0x3b, 0x83, 0x2b, 0x17,
	0xa1, 0x60, 0xd1, 0x03, 0xb8, 0x12, 0xa2, 0x16, 0x17, 0x57, 0x6d, 0x44, 0xe3, 0x34, 0x0e, 0x25,
	0x57, 0xbc, 0x57, 0xb9, 0xbe, 0x45, 0xcf, 0xd3, 0x13, 0xb4, 0x9c, 0x65, 0xcb, 0xd0, 0x17, 0x09,
	0x9d, 0xc6, 0x36, 0xee, 0xce, 0x81, 0xf3, 0x9d, 0xff, 0xe3, 0x87, 0xcd, 0x8b, 0x28, 0x50, 0x65,
	0xaf, 0xca, 0x2f, 0xb0, 0x28, 0x85, 0x78, 0xf7, 0x9b, 0xc0, 0x57, 0x6d, 0x89, 0xd2, 0x2b, 0x2b,
	0xa1, 0x84, 0x75, 0x71, 0x3a, 0xf0, 0xfe, 0x0e, 0xbc, 0x26, 0x70, 0xae, 0xc0, 0x88, 0xb5, 0xb4,
	0xce, 0xc1, 0x50, 0x5a, 0xda, 0x74, 0x6b, 0xb8, 0x6b, 0x3e, 0x8c, 0xce, 0x06, 0xcc, 0x28, 0x6d,
	0x30, 0xd6, 0xd6, 0x25, 0x98, 0x4a, 0x27, 0x6f, 0xd8, 0xda, 0x74, 0x4b, 0xdd, 0x35, 0x5f, 0x28,
	0xfd, 0x80, 0xad, 0xb3, 0x06, 0xe0, 0x28, 0x51, 0x71, 0x51, 0x2b, 0x74, 0x3e, 0x29, 0x2c, 0x1f,
	0x51, 0xca, 0x74, 0x87, 0xd6, 0xcd, 0xe9, 0x19, 0x75, 0x57, 0x81, 0xed, 0xcd, 0xc4, 0x7a, 0xb1,
	0x96, 0x11, 0x19, 0x83, 0xac, 0x7b, 0x58, 0xe6, 0x69, 0x83, 0x89, 0xd2, 0xf6, 0xd9, 0x48, 0x5c,
	0xcf, 0x12, 0x47, 0x99, 0x88, 0x70, 0x33, 0x3f, 0x6a, 0x85, 0xb0, 0xaa, 0x86, 0xfc, 0xa4, 0x1a,
	0x04, 0x6c, 0x63, 0x64, 0x37, 0xb3, 0xec, 0xbf, 0x67, 0x44, 0x38, 0x54, 0xd3, 0x16, 0x2e, 0xc0,
	0x90, 0x75, 0x11, 0x3e, 0x7d, 0x75, 0x8c, 0x1e, 0x3a, 0x46, 0x7f, 0x3a, 0x46, 0x3f, 0x7a, 0x46,
	0x0e, 0x3d, 0x23, 0xdf, 0x3d, 0x23, 0xcf, 0x77, 0xbb, 0xbd, 0xca, 0xeb, 0x6c, 0xf8, 0xea, 0x4f,
	0xfd, 0x4e, 0x43, 0x5a, 0xee, 0xfd, 0x99, 0xd6, 0x33, 0x73, 0x2c, 0xfc, 0xf6, 0x37, 0x00, 0x00,
	0xff, 0xff, 0xaa, 0x61, 0x3b, 0xc6, 0x93, 0x01, 0x00, 0x00,
}

func (m *Txs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Txs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Txs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for iNdEx := len(m.Txs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Txs[iNdEx])
			copy(dAtA[i:], m.Txs[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Txs[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *HaveTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HaveTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HaveTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxKey) > 0 {
		i -= len(m.TxKey)
		copy(dAtA[i:], m.TxKey)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.TxKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResetRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResetRoute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResetRoute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Message_Txs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_Txs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Txs != nil {
		{
			size, err := m.Txs.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Message_HaveTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_HaveTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.HaveTx != nil {
		{
			size, err := m.HaveTx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Message_ResetRoute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_ResetRoute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ResetRoute != nil {
		{
			size, err := m.ResetRoute.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Txs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for _, b := range m.Txs {
			l = len(b)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *HaveTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxKey)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ResetRoute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *Message_Txs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Txs != nil {
		l = m.Txs.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_HaveTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HaveTx != nil {
		l = m.HaveTx.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Message_ResetRoute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ResetRoute != nil {
		l = m.ResetRoute.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Txs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Txs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Txs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txs = append(m.Txs, make([]byte, postIndex-iNdEx))
			copy(m.Txs[len(m.Txs)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *HaveTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: HaveTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HaveTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxKey = append(m.TxKey[:0], dAtA[iNdEx:postIndex]...)
			if m.TxKey == nil {
				m.TxKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ResetRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResetRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResetRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Txs{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_Txs{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HaveTx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &HaveTx{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_HaveTx{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResetRoute", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResetRoute{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_ResetRoute{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
