// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evmos/claims/v1/claims.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Action defines the list of available actions to claim the airdrop tokens.
type Action int32

const (
	// ACTION_UNSPECIFIED defines an invalid action.
	ActionUnspecified Action = 0
	// ACTION_VOTE defines a proposal vote.
	ActionVote Action = 1
	// ACTION_DELEGATE defines an staking delegation.
	ActionDelegate Action = 2
	// ACTION_EVM defines an EVM transaction.
	ActionEVM Action = 3
	// ACTION_IBC_TRANSFER defines a fungible token transfer transaction via IBC.
	ActionIBCTransfer Action = 4
)

var Action_name = map[int32]string{
	0: "ACTION_UNSPECIFIED",
	1: "ACTION_VOTE",
	2: "ACTION_DELEGATE",
	3: "ACTION_EVM",
	4: "ACTION_IBC_TRANSFER",
}

var Action_value = map[string]int32{
	"ACTION_UNSPECIFIED":  0,
	"ACTION_VOTE":         1,
	"ACTION_DELEGATE":     2,
	"ACTION_EVM":          3,
	"ACTION_IBC_TRANSFER": 4,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a7153f2307523893, []int{0}
}

// Claim defines the action, completed flag and the remaining claimable amount
// for a given user. This is only used during client queries.
type Claim struct {
	// action enum
	Action Action `protobuf:"varint,1,opt,name=action,proto3,enum=evmos.claims.v1.Action" json:"action,omitempty"`
	// completed is true if the action has been completed
	Completed bool `protobuf:"varint,2,opt,name=completed,proto3" json:"completed,omitempty"`
	// claimable_amount of tokens for the action. Zero if completed
	ClaimableAmount cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=claimable_amount,json=claimableAmount,proto3,customtype=cosmossdk.io/math.Int" json:"claimable_amount"`
}

func (m *Claim) Reset()         { *m = Claim{} }
func (m *Claim) String() string { return proto.CompactTextString(m) }
func (*Claim) ProtoMessage()    {}
func (*Claim) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7153f2307523893, []int{0}
}
func (m *Claim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Claim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Claim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Claim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Claim.Merge(m, src)
}
func (m *Claim) XXX_Size() int {
	return m.Size()
}
func (m *Claim) XXX_DiscardUnknown() {
	xxx_messageInfo_Claim.DiscardUnknown(m)
}

var xxx_messageInfo_Claim proto.InternalMessageInfo

func (m *Claim) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return ActionUnspecified
}

func (m *Claim) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

// ClaimsRecordAddress is the claims metadata per address that is used at
// Genesis.
type ClaimsRecordAddress struct {
	// address of claiming user in either bech32 or hex format
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// initial_claimable_amount for the user
	InitialClaimableAmount cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=initial_claimable_amount,json=initialClaimableAmount,proto3,customtype=cosmossdk.io/math.Int" json:"initial_claimable_amount"`
	// actions_completed is a slice that describes which actions were completed
	ActionsCompleted []bool `protobuf:"varint,3,rep,packed,name=actions_completed,json=actionsCompleted,proto3" json:"actions_completed,omitempty"`
}

func (m *ClaimsRecordAddress) Reset()         { *m = ClaimsRecordAddress{} }
func (m *ClaimsRecordAddress) String() string { return proto.CompactTextString(m) }
func (*ClaimsRecordAddress) ProtoMessage()    {}
func (*ClaimsRecordAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7153f2307523893, []int{1}
}
func (m *ClaimsRecordAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimsRecordAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimsRecordAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimsRecordAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimsRecordAddress.Merge(m, src)
}
func (m *ClaimsRecordAddress) XXX_Size() int {
	return m.Size()
}
func (m *ClaimsRecordAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimsRecordAddress.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimsRecordAddress proto.InternalMessageInfo

func (m *ClaimsRecordAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClaimsRecordAddress) GetActionsCompleted() []bool {
	if m != nil {
		return m.ActionsCompleted
	}
	return nil
}

// ClaimsRecord defines the initial claimable airdrop amount and the list of
// completed actions to claim the tokens.
type ClaimsRecord struct {
	// initial_claimable_amount for the user
	InitialClaimableAmount cosmossdk_io_math.Int `protobuf:"bytes,1,opt,name=initial_claimable_amount,json=initialClaimableAmount,proto3,customtype=cosmossdk.io/math.Int" json:"initial_claimable_amount"`
	// actions_completed is a slice that describes which actions were completed
	ActionsCompleted []bool `protobuf:"varint,2,rep,packed,name=actions_completed,json=actionsCompleted,proto3" json:"actions_completed,omitempty"`
}

func (m *ClaimsRecord) Reset()         { *m = ClaimsRecord{} }
func (m *ClaimsRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimsRecord) ProtoMessage()    {}
func (*ClaimsRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7153f2307523893, []int{2}
}
func (m *ClaimsRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimsRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimsRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimsRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimsRecord.Merge(m, src)
}
func (m *ClaimsRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimsRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimsRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimsRecord proto.InternalMessageInfo

func (m *ClaimsRecord) GetActionsCompleted() []bool {
	if m != nil {
		return m.ActionsCompleted
	}
	return nil
}

func init() {
	proto.RegisterEnum("evmos.claims.v1.Action", Action_name, Action_value)
	proto.RegisterType((*Claim)(nil), "evmos.claims.v1.Claim")
	proto.RegisterType((*ClaimsRecordAddress)(nil), "evmos.claims.v1.ClaimsRecordAddress")
	proto.RegisterType((*ClaimsRecord)(nil), "evmos.claims.v1.ClaimsRecord")
}

func init() { proto.RegisterFile("evmos/claims/v1/claims.proto", fileDescriptor_a7153f2307523893) }

var fileDescriptor_a7153f2307523893 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x6e, 0xd3, 0x40,
	0x14, 0x87, 0x3d, 0x49, 0x09, 0xcd, 0x00, 0x89, 0x3b, 0xa5, 0x10, 0x59, 0xad, 0x63, 0x65, 0x43,
	0x04, 0xaa, 0xad, 0x96, 0x13, 0x38, 0x8e, 0x0b, 0x96, 0xda, 0x14, 0xb9, 0x6e, 0x90, 0xd8, 0x58,
	0x8e, 0x3d, 0x4d, 0x0c, 0xb6, 0xc7, 0xf2, 0x4c, 0x22, 0xb8, 0x01, 0xca, 0x8a, 0x05, 0xdb, 0x48,
	0x48, 0x9c, 0x81, 0x3b, 0x74, 0xd9, 0x15, 0x42, 0x2c, 0x2a, 0x94, 0x5c, 0x04, 0xc5, 0x9e, 0xd0,
	0xf2, 0x47, 0x88, 0x05, 0xbb, 0x37, 0xef, 0xf7, 0x3d, 0xeb, 0x7b, 0xb2, 0x1e, 0xdc, 0xc6, 0x93,
	0x98, 0x50, 0xcd, 0x8f, 0xbc, 0x30, 0xa6, 0xda, 0x64, 0x8f, 0x57, 0x6a, 0x9a, 0x11, 0x46, 0x50,
	0x3d, 0x4f, 0x55, 0xde, 0x9b, 0xec, 0x49, 0x77, 0x87, 0x64, 0x48, 0xf2, 0x4c, 0x5b, 0x56, 0x05,
	0xd6, 0xfa, 0x00, 0xe0, 0x0d, 0x63, 0xc9, 0x20, 0x0d, 0x56, 0x3c, 0x9f, 0x85, 0x24, 0x69, 0x00,
	0x05, 0xb4, 0x6b, 0xfb, 0xf7, 0xd5, 0x5f, 0xbe, 0xa0, 0xea, 0x79, 0x6c, 0x73, 0x0c, 0x6d, 0xc3,
	0xaa, 0x4f, 0xe2, 0x34, 0xc2, 0x0c, 0x07, 0x8d, 0x92, 0x02, 0xda, 0xeb, 0xf6, 0x55, 0x03, 0x3d,
	0x85, 0x62, 0x3e, 0xe9, 0x0d, 0x22, 0xec, 0x7a, 0x31, 0x19, 0x27, 0xac, 0x51, 0x56, 0x40, 0xbb,
	0xda, 0xd9, 0x39, 0xbf, 0x6c, 0x0a, 0x5f, 0x2f, 0x9b, 0x5b, 0x3e, 0xa1, 0x31, 0xa1, 0x34, 0x78,
	0xa5, 0x86, 0x44, 0x8b, 0x3d, 0x36, 0x52, 0xad, 0x84, 0xd9, 0xf5, 0x1f, 0x63, 0x7a, 0x3e, 0xd5,
	0xfa, 0x04, 0xe0, 0x66, 0xae, 0x48, 0x6d, 0xec, 0x93, 0x2c, 0xd0, 0x83, 0x20, 0xc3, 0x94, 0xa2,
	0x06, 0xbc, 0xe9, 0x15, 0x65, 0x6e, 0x5c, 0xb5, 0x57, 0x4f, 0xf4, 0x1c, 0x36, 0xc2, 0x24, 0x64,
	0xa1, 0x17, 0xb9, 0xbf, 0x39, 0x94, 0xfe, 0xc5, 0xe1, 0x1e, 0x1f, 0x37, 0x7e, 0x56, 0x41, 0x8f,
	0xe0, 0x46, 0xb1, 0x3c, 0x75, 0xaf, 0x56, 0x2f, 0x2b, 0xe5, 0xf6, 0xba, 0x2d, 0xf2, 0xc0, 0x58,
	0xf5, 0x5b, 0xef, 0x01, 0xbc, 0x7d, 0xdd, 0xfb, 0xaf, 0x5a, 0xe0, 0xbf, 0x6b, 0x95, 0xfe, 0xac,
	0xf5, 0xf0, 0x33, 0x80, 0x95, 0xe2, 0x4f, 0xa2, 0x5d, 0x88, 0x74, 0xc3, 0xb1, 0x8e, 0x7b, 0xee,
	0x69, 0xef, 0xe4, 0x99, 0x69, 0x58, 0x07, 0x96, 0xd9, 0x15, 0x05, 0x69, 0x6b, 0x3a, 0x53, 0x36,
	0x0a, 0xe6, 0x34, 0xa1, 0x29, 0xf6, 0xc3, 0xb3, 0x10, 0x07, 0xa8, 0x09, 0x6f, 0x71, 0xbc, 0x7f,
	0xec, 0x98, 0x22, 0x90, 0x6a, 0xd3, 0x99, 0x02, 0x0b, 0xae, 0x4f, 0x18, 0x46, 0x0f, 0x60, 0x9d,
	0x03, 0x5d, 0xf3, 0xd0, 0x7c, 0xa2, 0x3b, 0xa6, 0x58, 0x92, 0xd0, 0x74, 0xa6, 0xd4, 0x0a, 0xa8,
	0x8b, 0x23, 0x3c, 0xf4, 0x18, 0x46, 0x3b, 0x10, 0x72, 0xd0, 0xec, 0x1f, 0x89, 0x65, 0xe9, 0xce,
	0x74, 0xa6, 0x54, 0x0b, 0xc6, 0xec, 0x1f, 0x21, 0x15, 0x6e, 0xf2, 0xd8, 0xea, 0x18, 0xae, 0x63,
	0xeb, 0xbd, 0x93, 0x03, 0xd3, 0x16, 0xd7, 0xae, 0x8b, 0x59, 0x1d, 0xc3, 0xc9, 0xbc, 0x84, 0x9e,
	0xe1, 0x4c, 0x5a, 0x7b, 0xfb, 0x51, 0x16, 0x3a, 0x87, 0xe7, 0x73, 0x19, 0x5c, 0xcc, 0x65, 0xf0,
	0x6d, 0x2e, 0x83, 0x77, 0x0b, 0x59, 0xb8, 0x58, 0xc8, 0xc2, 0x97, 0x85, 0x2c, 0xbc, 0xd8, 0x1f,
	0x86, 0x6c, 0x34, 0x1e, 0xa8, 0x3e, 0x89, 0xb5, 0x11, 0x66, 0xe3, 0xdd, 0x34, 0x23, 0x2f, 0xb1,
	0xcf, 0x8a, 0xc7, 0x68, 0x3c, 0x58, 0x5e, 0xcf, 0xeb, 0xd5, 0x25, 0xb1, 0x37, 0x29, 0xa6, 0x83,
	0x4a, 0x7e, 0x1f, 0x8f, 0xbf, 0x07, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x5b, 0x16, 0x68, 0x66, 0x03,
	0x00, 0x00,
}

func (m *Claim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Claim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Claim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ClaimableAmount.Size()
		i -= size
		if _, err := m.ClaimableAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintClaims(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Completed {
		i--
		if m.Completed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Action != 0 {
		i = encodeVarintClaims(dAtA, i, uint64(m.Action))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ClaimsRecordAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimsRecordAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimsRecordAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ActionsCompleted) > 0 {
		for iNdEx := len(m.ActionsCompleted) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.ActionsCompleted[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintClaims(dAtA, i, uint64(len(m.ActionsCompleted)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.InitialClaimableAmount.Size()
		i -= size
		if _, err := m.InitialClaimableAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintClaims(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClaimsRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimsRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimsRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ActionsCompleted) > 0 {
		for iNdEx := len(m.ActionsCompleted) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.ActionsCompleted[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintClaims(dAtA, i, uint64(len(m.ActionsCompleted)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.InitialClaimableAmount.Size()
		i -= size
		if _, err := m.InitialClaimableAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintClaims(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintClaims(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaims(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Claim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Action != 0 {
		n += 1 + sovClaims(uint64(m.Action))
	}
	if m.Completed {
		n += 2
	}
	l = m.ClaimableAmount.Size()
	n += 1 + l + sovClaims(uint64(l))
	return n
}

func (m *ClaimsRecordAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	l = m.InitialClaimableAmount.Size()
	n += 1 + l + sovClaims(uint64(l))
	if len(m.ActionsCompleted) > 0 {
		n += 1 + sovClaims(uint64(len(m.ActionsCompleted))) + len(m.ActionsCompleted)*1
	}
	return n
}

func (m *ClaimsRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.InitialClaimableAmount.Size()
	n += 1 + l + sovClaims(uint64(l))
	if len(m.ActionsCompleted) > 0 {
		n += 1 + sovClaims(uint64(len(m.ActionsCompleted))) + len(m.ActionsCompleted)*1
	}
	return n
}

func sovClaims(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaims(x uint64) (n int) {
	return sovClaims(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Claim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaims
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
			return fmt.Errorf("proto: Claim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Claim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Action |= Action(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Completed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
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
			m.Completed = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimableAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
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
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClaimableAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaims(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaims
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
func (m *ClaimsRecordAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaims
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
			return fmt.Errorf("proto: ClaimsRecordAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimsRecordAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
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
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialClaimableAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
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
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialClaimableAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaims
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
				m.ActionsCompleted = append(m.ActionsCompleted, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaims
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaims
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaims
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.ActionsCompleted) == 0 {
					m.ActionsCompleted = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaims
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
					m.ActionsCompleted = append(m.ActionsCompleted, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionsCompleted", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaims(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaims
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
func (m *ClaimsRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaims
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
			return fmt.Errorf("proto: ClaimsRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimsRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialClaimableAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
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
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialClaimableAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaims
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
				m.ActionsCompleted = append(m.ActionsCompleted, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaims
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaims
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaims
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.ActionsCompleted) == 0 {
					m.ActionsCompleted = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaims
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
					m.ActionsCompleted = append(m.ActionsCompleted, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionsCompleted", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaims(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaims
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
func skipClaims(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaims
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
					return 0, ErrIntOverflowClaims
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
					return 0, ErrIntOverflowClaims
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
				return 0, ErrInvalidLengthClaims
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaims
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaims
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaims        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaims          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaims = fmt.Errorf("proto: unexpected end of group")
)
