// Code generated by protoc-gen-go.
// source: butler.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

//
// This enumerates the possible contents of published Butler data.
type ButlerMetadata_ContentType int32

const (
	// The published data is a ButlerLogBundle protobuf message.
	ButlerMetadata_ButlerLogBundle ButlerMetadata_ContentType = 1
)

var ButlerMetadata_ContentType_name = map[int32]string{
	1: "ButlerLogBundle",
}
var ButlerMetadata_ContentType_value = map[string]int32{
	"ButlerLogBundle": 1,
}

func (x ButlerMetadata_ContentType) Enum() *ButlerMetadata_ContentType {
	p := new(ButlerMetadata_ContentType)
	*p = x
	return p
}
func (x ButlerMetadata_ContentType) String() string {
	return proto.EnumName(ButlerMetadata_ContentType_name, int32(x))
}
func (x *ButlerMetadata_ContentType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ButlerMetadata_ContentType_value, data, "ButlerMetadata_ContentType")
	if err != nil {
		return err
	}
	*x = ButlerMetadata_ContentType(value)
	return nil
}

//
// ButlerMetadata appears as a frame at the beginning of Butler published data
// to describe the remainder of the contents.
type ButlerMetadata struct {
	// This is the type of data in the subsequent frame.
	Type *ButlerMetadata_ContentType `protobuf:"varint,1,opt,name=type,enum=protocol.ButlerMetadata_ContentType" json:"type,omitempty"`
	// If true, the content data is compressed.
	Compressed       *bool  `protobuf:"varint,2,opt,name=compressed" json:"compressed,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ButlerMetadata) Reset()         { *m = ButlerMetadata{} }
func (m *ButlerMetadata) String() string { return proto.CompactTextString(m) }
func (*ButlerMetadata) ProtoMessage()    {}

func (m *ButlerMetadata) GetType() ButlerMetadata_ContentType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ButlerMetadata_ButlerLogBundle
}

func (m *ButlerMetadata) GetCompressed() bool {
	if m != nil && m.Compressed != nil {
		return *m.Compressed
	}
	return false
}

//
// A message containing log data in transit from the Butler.
//
// The Butler is capable of conserving bandwidth by bundling collected log
// messages together into this protocol buffer. Based on Butler bundling
// settings, this message can represent anything from a single LogRecord to
// multiple LogRecords belonging to several different streams.
//
// Entries in a Log Bundle are fully self-descriptive: no additional information
// is needed to fully associate the contained data with its proper place in
// the source log stream.
type ButlerLogBundle struct {
	//
	// String describing the source of this LogBundle.
	// This is an unstructured field, and is not intended to be parsed. An
	// example would be: "Butler @a33967 (Linux/amd64)".
	//
	// This field will be used for debugging and internal accounting.
	Source *string `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	// The timestamp when this bundle was generated.
	//
	// This field will be used for debugging and internal accounting.
	Timestamp *Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	//
	// The log prefix's secret value (required).
	//
	// The secret is generated by the Butler and unique to this specific log
	// stream. The Coordinator will record the secret associated with a given
	// log Prefix/Stream, but will not share the secret with a client.
	//
	// The Collector will check the secret prior to ingesting logs. If the
	// secret doesn't match the value recorded by the Coordinator, the log
	// will be discarded.
	//
	// This ensures that only the Butler instance that generated the log stream
	// can emit log data for that stream.
	Secret []byte `protobuf:"bytes,3,opt,name=secret" json:"secret,omitempty"`
	// *
	// Each Entry is an individual set of log records for a given log stream.
	Entries          []*ButlerLogBundle_Entry `protobuf:"bytes,4,rep,name=entries" json:"entries,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *ButlerLogBundle) Reset()         { *m = ButlerLogBundle{} }
func (m *ButlerLogBundle) String() string { return proto.CompactTextString(m) }
func (*ButlerLogBundle) ProtoMessage()    {}

func (m *ButlerLogBundle) GetSource() string {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return ""
}

func (m *ButlerLogBundle) GetTimestamp() *Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ButlerLogBundle) GetSecret() []byte {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *ButlerLogBundle) GetEntries() []*ButlerLogBundle_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

//
// A bundle Entry describes a set of LogEntry messages originating from the
// same log stream.
type ButlerLogBundle_Entry struct {
	//
	// The descriptor for this entry's log stream.
	//
	// Each LogEntry in the "logs" field is shares this common descriptor.
	Desc *LogStreamDescriptor `protobuf:"bytes,1,opt,name=desc" json:"desc,omitempty"`
	//
	// Whether this log entry terminates its stream.
	//
	// If present and "true", this field declares that this Entry is the last
	// such entry in the stream. This fact is recorded by the Collector and
	// registered with the Coordinator. The largest stream prefix in this Entry
	// will be bound the stream's LogEntry records to [0:largest_prefix]. Once
	// all messages in that range have been received, the log may be archived.
	//
	// Further log entries belonging to this stream with stream indices
	// exceeding the terminal log's index will be discarded.
	Terminal *bool `protobuf:"varint,3,opt,name=terminal" json:"terminal,omitempty"`
	//
	// If terminal is true, this is the terminal stream index; that is, the last
	// message index in the stream.
	TerminalIndex *uint32 `protobuf:"varint,4,opt,name=terminal_index" json:"terminal_index,omitempty"`
	//
	// Log entries attached to this record. These must be sequential and in
	// order.
	//
	// This is the main log entry content.
	Logs             []*LogEntry `protobuf:"bytes,5,rep,name=logs" json:"logs,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *ButlerLogBundle_Entry) Reset()         { *m = ButlerLogBundle_Entry{} }
func (m *ButlerLogBundle_Entry) String() string { return proto.CompactTextString(m) }
func (*ButlerLogBundle_Entry) ProtoMessage()    {}

func (m *ButlerLogBundle_Entry) GetDesc() *LogStreamDescriptor {
	if m != nil {
		return m.Desc
	}
	return nil
}

func (m *ButlerLogBundle_Entry) GetTerminal() bool {
	if m != nil && m.Terminal != nil {
		return *m.Terminal
	}
	return false
}

func (m *ButlerLogBundle_Entry) GetTerminalIndex() uint32 {
	if m != nil && m.TerminalIndex != nil {
		return *m.TerminalIndex
	}
	return 0
}

func (m *ButlerLogBundle_Entry) GetLogs() []*LogEntry {
	if m != nil {
		return m.Logs
	}
	return nil
}

func init() {
	proto.RegisterEnum("protocol.ButlerMetadata_ContentType", ButlerMetadata_ContentType_name, ButlerMetadata_ContentType_value)
}
