// Code generated by protoc-gen-go.
// source: storage.proto
// DO NOT EDIT!

package services

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Storage is the in-transit storage configuration.
type Storage struct {
	// Type is the transport configuration that is being used.
	//
	// Types that are valid to be assigned to Type:
	//	*Storage_Bigtable
	Type isStorage_Type `protobuf_oneof:"Type"`
}

func (m *Storage) Reset()                    { *m = Storage{} }
func (m *Storage) String() string            { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()               {}
func (*Storage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type isStorage_Type interface {
	isStorage_Type()
}

type Storage_Bigtable struct {
	Bigtable *Storage_BigTable `protobuf:"bytes,1,opt,name=bigtable,oneof"`
}

func (*Storage_Bigtable) isStorage_Type() {}

func (m *Storage) GetType() isStorage_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Storage) GetBigtable() *Storage_BigTable {
	if x, ok := m.GetType().(*Storage_Bigtable); ok {
		return x.Bigtable
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Storage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _Storage_OneofMarshaler, _Storage_OneofUnmarshaler, []interface{}{
		(*Storage_Bigtable)(nil),
	}
}

func _Storage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Storage)
	// Type
	switch x := m.Type.(type) {
	case *Storage_Bigtable:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Bigtable); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Storage.Type has unexpected type %T", x)
	}
	return nil
}

func _Storage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Storage)
	switch tag {
	case 1: // Type.bigtable
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Storage_BigTable)
		err := b.DecodeMessage(msg)
		m.Type = &Storage_Bigtable{msg}
		return true, err
	default:
		return false, nil
	}
}

// BigTable is the set of BigTable configuration parameters.
type Storage_BigTable struct {
	// The name of the BigTable instance project.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// The name of the BigTable instance zone.
	Zone string `protobuf:"bytes,2,opt,name=zone" json:"zone,omitempty"`
	// The name of the BigTable instance cluster.
	Cluster string `protobuf:"bytes,3,opt,name=cluster" json:"cluster,omitempty"`
	// The name of the BigTable instance's log table.
	LogTableName string `protobuf:"bytes,4,opt,name=log_table_name" json:"log_table_name,omitempty"`
}

func (m *Storage_BigTable) Reset()                    { *m = Storage_BigTable{} }
func (m *Storage_BigTable) String() string            { return proto.CompactTextString(m) }
func (*Storage_BigTable) ProtoMessage()               {}
func (*Storage_BigTable) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func init() {
	proto.RegisterType((*Storage)(nil), "services.Storage")
	proto.RegisterType((*Storage_BigTable)(nil), "services.Storage.BigTable")
}

var fileDescriptor1 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0x2d, 0x56, 0x9a, 0xca, 0xc8, 0xc5, 0x1e, 0x0c, 0x91, 0x13, 0xd2, 0xe3, 0xe2, 0x48,
	0xca, 0x4c, 0x2f, 0x49, 0x4c, 0xca, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd2,
	0x83, 0x29, 0xd4, 0x83, 0x2a, 0xd2, 0x73, 0xca, 0x4c, 0x0f, 0x01, 0xa9, 0xf0, 0x60, 0x90, 0x0a,
	0xe2, 0xe2, 0x80, 0xf1, 0x84, 0xf8, 0xb9, 0xd8, 0x81, 0x46, 0x67, 0xa5, 0x26, 0x97, 0x80, 0xb5,
	0x72, 0x0a, 0xf1, 0x70, 0xb1, 0x54, 0xe5, 0xe7, 0xa5, 0x4a, 0x30, 0x81, 0x79, 0x40, 0xe9, 0xe4,
	0x9c, 0xd2, 0xe2, 0x92, 0xd4, 0x22, 0x09, 0x66, 0xb0, 0x80, 0x18, 0x17, 0x5f, 0x4e, 0x7e, 0x7a,
	0x3c, 0xd8, 0xb2, 0xf8, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x16, 0x90, 0xb8, 0x13, 0x1b, 0x17, 0x4b,
	0x48, 0x65, 0x41, 0x6a, 0x12, 0x1b, 0xd8, 0xa1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc5,
	0x68, 0x26, 0xb8, 0xb9, 0x00, 0x00, 0x00,
}
