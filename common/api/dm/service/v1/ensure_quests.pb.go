// Code generated by protoc-gen-go.
// source: ensure_quests.proto
// DO NOT EDIT!

package dm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// EnsureQuestsReq allows either a human or a client to ensure that various
// Quests exist in DM.
type EnsureQuestsReq struct {
	// optional: Only needed if this is being called from an execution.
	Auth *ExecutionAuth `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
	// ToEnsure is a list of quest descriptors. DM will ensure that the
	// corresponding Quests exist. If they don't, they'll be created.
	ToEnsure []*Quest_Descriptor `protobuf:"bytes,2,rep,name=to_ensure" json:"to_ensure,omitempty"`
}

func (m *EnsureQuestsReq) Reset()                    { *m = EnsureQuestsReq{} }
func (m *EnsureQuestsReq) String() string            { return proto.CompactTextString(m) }
func (*EnsureQuestsReq) ProtoMessage()               {}
func (*EnsureQuestsReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *EnsureQuestsReq) GetAuth() *ExecutionAuth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *EnsureQuestsReq) GetToEnsure() []*Quest_Descriptor {
	if m != nil {
		return m.ToEnsure
	}
	return nil
}

type EnsureQuestsRsp struct {
	// QuestIds is a list that matches 1:1 with the QuestDescriptors listed in the
	// request. These IDs are the canonical ids of the Quests.
	QuestIds []*Quest_ID `protobuf:"bytes,1,rep,name=quest_ids" json:"quest_ids,omitempty"`
}

func (m *EnsureQuestsRsp) Reset()                    { *m = EnsureQuestsRsp{} }
func (m *EnsureQuestsRsp) String() string            { return proto.CompactTextString(m) }
func (*EnsureQuestsRsp) ProtoMessage()               {}
func (*EnsureQuestsRsp) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *EnsureQuestsRsp) GetQuestIds() []*Quest_ID {
	if m != nil {
		return m.QuestIds
	}
	return nil
}

func init() {
	proto.RegisterType((*EnsureQuestsReq)(nil), "dm.EnsureQuestsReq")
	proto.RegisterType((*EnsureQuestsRsp)(nil), "dm.EnsureQuestsRsp")
}

var fileDescriptor3 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xcd, 0x2b, 0x2e,
	0x2d, 0x4a, 0x8d, 0x2f, 0x2c, 0x4d, 0x2d, 0x2e, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x4a, 0xc9, 0x95, 0x12, 0x48, 0x2f, 0x4a, 0x2c, 0xc8, 0x88, 0x4f, 0x49, 0x2c, 0x49, 0x84,
	0x88, 0x4a, 0x71, 0x97, 0x54, 0x16, 0xa4, 0x42, 0x95, 0x28, 0x45, 0x73, 0xf1, 0xbb, 0x82, 0x75,
	0x06, 0x82, 0x35, 0x06, 0xa5, 0x16, 0x0a, 0xc9, 0x73, 0xb1, 0x24, 0x96, 0x96, 0x64, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x09, 0xea, 0xa5, 0xe4, 0xea, 0xb9, 0x56, 0xa4, 0x26, 0x97, 0x96,
	0x64, 0xe6, 0xe7, 0x39, 0x02, 0x25, 0x84, 0xd4, 0xb9, 0x38, 0x4b, 0xf2, 0xe3, 0x21, 0x16, 0x4a,
	0x30, 0x29, 0x30, 0x03, 0x55, 0x89, 0x80, 0x54, 0x81, 0x8d, 0xd0, 0x73, 0x49, 0x2d, 0x4e, 0x2e,
	0xca, 0x2c, 0x28, 0xc9, 0x2f, 0x52, 0x32, 0x42, 0x33, 0xbc, 0xb8, 0x00, 0x68, 0x38, 0x27, 0xd8,
	0x89, 0xf1, 0x99, 0x29, 0xc5, 0x40, 0x1b, 0x40, 0x7a, 0x79, 0x10, 0x7a, 0x3d, 0x5d, 0x92, 0xd8,
	0xc0, 0xee, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x28, 0x57, 0xbc, 0xd1, 0x00, 0x00,
	0x00,
}
