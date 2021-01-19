// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: checkpoint/v1beta/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/duration"
	types "github.com/maticnetwork/heimdall/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Params struct {
	CheckpointBufferTime time.Duration `protobuf:"bytes,1,opt,name=checkpoint_buffer_time,json=checkpointBufferTime,proto3,stdduration" json:"checkpoint_buffer_time"`
	AvgCheckpointLength  uint64        `protobuf:"varint,2,opt,name=avg_checkpoint_length,json=avgCheckpointLength,proto3" json:"avg_checkpoint_length,omitempty"`
	MaxCheckpointLength  uint64        `protobuf:"varint,3,opt,name=max_checkpoint_length,json=maxCheckpointLength,proto3" json:"max_checkpoint_length,omitempty"`
	ChildBlockInterval   uint64        `protobuf:"varint,4,opt,name=child_block_interval,json=childBlockInterval,proto3" json:"child_block_interval,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_99a783d62b74a2d0, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

// GenesisState defines the checkpoint module's genesis state.
type GenesisState struct {
	Params             Params              `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	BufferedCheckpoint *types.Checkpoint   `protobuf:"bytes,2,opt,name=buffered_checkpoint,json=bufferedCheckpoint,proto3" json:"buffered_checkpoint,omitempty"`
	LastNoACK          uint64              `protobuf:"varint,3,opt,name=last_no_ack,json=lastNoAck,proto3" json:"last_no_ack,omitempty"`
	AckCount           uint64              `protobuf:"varint,4,opt,name=ack_count,json=ackCount,proto3" json:"ack_count,omitempty"`
	Checkpoints        []*types.Checkpoint `protobuf:"bytes,5,rep,name=checkpoints,proto3" json:"checkpoints,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_99a783d62b74a2d0, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "heimdall.checkpoint.v1beta1.Params")
	proto.RegisterType((*GenesisState)(nil), "heimdall.checkpoint.v1beta1.GenesisState")
}

func init() { proto.RegisterFile("checkpoint/v1beta/genesis.proto", fileDescriptor_99a783d62b74a2d0) }

var fileDescriptor_99a783d62b74a2d0 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x3f, 0x8f, 0xd3, 0x30,
	0x1c, 0x4d, 0x7a, 0xa5, 0x6a, 0x5d, 0x58, 0x7c, 0x05, 0x95, 0x1e, 0x4a, 0xaa, 0x63, 0xb9, 0x05,
	0x9b, 0xeb, 0x6d, 0x88, 0xe5, 0x52, 0x24, 0x84, 0xee, 0x40, 0x28, 0xb0, 0xc0, 0x12, 0x39, 0x8e,
	0x9b, 0x58, 0xf9, 0xe3, 0x2a, 0x71, 0x4a, 0xf9, 0x04, 0x30, 0x32, 0xde, 0x78, 0x9f, 0x81, 0x4f,
	0x71, 0xe3, 0x8d, 0x4c, 0x05, 0xb5, 0x5f, 0x04, 0xc5, 0x4e, 0x9b, 0x4a, 0x20, 0xd8, 0xec, 0xbc,
	0xdf, 0x7b, 0xcf, 0xbf, 0xf7, 0xfb, 0x05, 0xd8, 0x34, 0x62, 0x34, 0x9e, 0x0b, 0x9e, 0x49, 0xbc,
	0x38, 0xf5, 0x99, 0x24, 0x38, 0x64, 0x19, 0x2b, 0x78, 0x81, 0xe6, 0xb9, 0x90, 0x02, 0x1e, 0x45,
	0x8c, 0xa7, 0x01, 0x49, 0x12, 0xd4, 0x54, 0x22, 0x5d, 0x79, 0x3a, 0x1a, 0x84, 0x22, 0x14, 0xaa,
	0x0e, 0x57, 0x27, 0x4d, 0x19, 0x59, 0xa1, 0x10, 0x61, 0xc2, 0xb0, 0xba, 0xf9, 0xe5, 0x0c, 0x07,
	0x65, 0x4e, 0x24, 0x17, 0x59, 0x8d, 0x3f, 0xda, 0x4a, 0x62, 0xf9, 0x79, 0xce, 0x0a, 0x1c, 0x31,
	0x12, 0xb0, 0xbc, 0x36, 0x3c, 0xfe, 0xd2, 0x02, 0x9d, 0xb7, 0x24, 0x27, 0x69, 0x01, 0x3f, 0x80,
	0x07, 0x8d, 0xa9, 0xe7, 0x97, 0xb3, 0x19, 0xcb, 0x3d, 0xc9, 0x53, 0x36, 0x34, 0xc7, 0xe6, 0x49,
	0x7f, 0xf2, 0x10, 0x69, 0x27, 0xb4, 0x75, 0x42, 0x2f, 0x6a, 0x27, 0xa7, 0x7b, 0xb3, 0xb2, 0x8d,
	0xab, 0x9f, 0xb6, 0xe9, 0x0e, 0x1a, 0x09, 0x47, 0x29, 0xbc, 0xe7, 0x29, 0x83, 0x13, 0x70, 0x9f,
	0x2c, 0x42, 0x6f, 0x4f, 0x3e, 0x61, 0x59, 0x28, 0xa3, 0x61, 0x6b, 0x6c, 0x9e, 0xb4, 0xdd, 0x43,
	0xb2, 0x08, 0xa7, 0x3b, 0xec, 0x52, 0x41, 0x15, 0x27, 0x25, 0xcb, 0xbf, 0x70, 0x0e, 0x34, 0x27,
	0x25, 0xcb, 0x3f, 0x38, 0x4f, 0xc1, 0x80, 0x46, 0x3c, 0x09, 0x3c, 0x3f, 0x11, 0x34, 0xf6, 0x78,
	0x26, 0x59, 0xbe, 0x20, 0xc9, 0xb0, 0xad, 0x28, 0x50, 0x61, 0x4e, 0x05, 0xbd, 0xaa, 0x91, 0x67,
	0xdd, 0xaf, 0xd7, 0xb6, 0x71, 0x75, 0x6d, 0x1b, 0xc7, 0xdf, 0x5b, 0xe0, 0xee, 0x4b, 0x3d, 0x8c,
	0x77, 0x92, 0x48, 0x06, 0xcf, 0x41, 0x67, 0xae, 0x92, 0xa9, 0xfb, 0x7f, 0x8c, 0xfe, 0x31, 0x1c,
	0xa4, 0x43, 0x74, 0xda, 0x55, 0x12, 0x6e, 0x4d, 0x84, 0x17, 0xe0, 0x50, 0xe7, 0xc8, 0x82, 0xbd,
	0x46, 0x54, 0xd7, 0xfd, 0xc9, 0xa8, 0xd1, 0x53, 0x93, 0x41, 0x4d, 0x3b, 0x2e, 0xdc, 0xd2, 0x9a,
	0x6f, 0xf0, 0x09, 0xe8, 0x27, 0xa4, 0x90, 0x5e, 0x26, 0x3c, 0x42, 0x63, 0x1d, 0x83, 0x73, 0x6f,
	0xbd, 0xb2, 0x7b, 0x97, 0xa4, 0x90, 0x6f, 0xc4, 0xf9, 0xf4, 0xc2, 0xed, 0x25, 0xfa, 0x48, 0x63,
	0x78, 0x04, 0x7a, 0x84, 0xc6, 0x1e, 0x15, 0x65, 0x26, 0xeb, 0x00, 0xba, 0x84, 0xc6, 0xd3, 0xea,
	0x0e, 0x9f, 0x83, 0x7e, 0xf3, 0x9e, 0x62, 0x78, 0x67, 0x7c, 0xf0, 0x9f, 0x07, 0xed, 0x97, 0xef,
	0x42, 0x33, 0x9d, 0xd7, 0x37, 0x6b, 0xcb, 0xbc, 0x5d, 0x5b, 0xe6, 0xaf, 0xb5, 0x65, 0x7e, 0xdb,
	0x58, 0xc6, 0xed, 0xc6, 0x32, 0x7e, 0x6c, 0x2c, 0xe3, 0xe3, 0x59, 0xc8, 0x65, 0x54, 0xfa, 0x88,
	0x8a, 0x14, 0xa7, 0x44, 0x72, 0x9a, 0x31, 0xf9, 0x49, 0xe4, 0x31, 0xde, 0xad, 0xe3, 0x12, 0xef,
	0xfd, 0x0d, 0xca, 0xd0, 0xef, 0xa8, 0xd5, 0x3a, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x23, 0x53,
	0xb4, 0x29, 0x28, 0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ChildBlockInterval != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ChildBlockInterval))
		i--
		dAtA[i] = 0x20
	}
	if m.MaxCheckpointLength != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaxCheckpointLength))
		i--
		dAtA[i] = 0x18
	}
	if m.AvgCheckpointLength != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AvgCheckpointLength))
		i--
		dAtA[i] = 0x10
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.CheckpointBufferTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.CheckpointBufferTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGenesis(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Checkpoints) > 0 {
		for iNdEx := len(m.Checkpoints) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Checkpoints[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.AckCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AckCount))
		i--
		dAtA[i] = 0x20
	}
	if m.LastNoACK != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastNoACK))
		i--
		dAtA[i] = 0x18
	}
	if m.BufferedCheckpoint != nil {
		{
			size, err := m.BufferedCheckpoint.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.CheckpointBufferTime)
	n += 1 + l + sovGenesis(uint64(l))
	if m.AvgCheckpointLength != 0 {
		n += 1 + sovGenesis(uint64(m.AvgCheckpointLength))
	}
	if m.MaxCheckpointLength != 0 {
		n += 1 + sovGenesis(uint64(m.MaxCheckpointLength))
	}
	if m.ChildBlockInterval != 0 {
		n += 1 + sovGenesis(uint64(m.ChildBlockInterval))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.BufferedCheckpoint != nil {
		l = m.BufferedCheckpoint.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.LastNoACK != 0 {
		n += 1 + sovGenesis(uint64(m.LastNoACK))
	}
	if m.AckCount != 0 {
		n += 1 + sovGenesis(uint64(m.AckCount))
	}
	if len(m.Checkpoints) > 0 {
		for _, e := range m.Checkpoints {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CheckpointBufferTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.CheckpointBufferTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvgCheckpointLength", wireType)
			}
			m.AvgCheckpointLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AvgCheckpointLength |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxCheckpointLength", wireType)
			}
			m.MaxCheckpointLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxCheckpointLength |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChildBlockInterval", wireType)
			}
			m.ChildBlockInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChildBlockInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BufferedCheckpoint", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BufferedCheckpoint == nil {
				m.BufferedCheckpoint = &types.Checkpoint{}
			}
			if err := m.BufferedCheckpoint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastNoACK", wireType)
			}
			m.LastNoACK = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastNoACK |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckCount", wireType)
			}
			m.AckCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AckCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checkpoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checkpoints = append(m.Checkpoints, &types.Checkpoint{})
			if err := m.Checkpoints[len(m.Checkpoints)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
