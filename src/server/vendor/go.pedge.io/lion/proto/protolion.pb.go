// Code generated by protoc-gen-go.
// source: proto/protolion.proto
// DO NOT EDIT!

package protolion

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Level is a logging level.
type Level int32

const (
	Level_LEVEL_DEBUG Level = 0
	Level_LEVEL_INFO  Level = 1
	Level_LEVEL_WARN  Level = 2
	Level_LEVEL_ERROR Level = 3
	Level_LEVEL_FATAL Level = 4
	Level_LEVEL_PANIC Level = 5
	Level_LEVEL_NONE  Level = 6
)

var Level_name = map[int32]string{
	0: "LEVEL_DEBUG",
	1: "LEVEL_INFO",
	2: "LEVEL_WARN",
	3: "LEVEL_ERROR",
	4: "LEVEL_FATAL",
	5: "LEVEL_PANIC",
	6: "LEVEL_NONE",
}
var Level_value = map[string]int32{
	"LEVEL_DEBUG": 0,
	"LEVEL_INFO":  1,
	"LEVEL_WARN":  2,
	"LEVEL_ERROR": 3,
	"LEVEL_FATAL": 4,
	"LEVEL_PANIC": 5,
	"LEVEL_NONE":  6,
}

func (x Level) String() string {
	return proto.EnumName(Level_name, int32(x))
}
func (Level) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Entry is the object serialized for logging.
type Entry struct {
	// id may not be set depending on logger options
	// it is up to the user to determine if id is required
	Id        string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Level     Level                      `protobuf:"varint,2,opt,name=level,enum=lion.Level" json:"level,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	// both context and fields may be set
	Context []*Entry_Message  `protobuf:"bytes,4,rep,name=context" json:"context,omitempty"`
	Fields  map[string]string `protobuf:"bytes,5,rep,name=fields" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// one of event, message, writer_output will be set
	Event        *Entry_Message `protobuf:"bytes,6,opt,name=event" json:"event,omitempty"`
	Message      string         `protobuf:"bytes,7,opt,name=message" json:"message,omitempty"`
	WriterOutput []byte         `protobuf:"bytes,8,opt,name=writer_output,json=writerOutput,proto3" json:"writer_output,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Entry) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Entry) GetContext() []*Entry_Message {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Entry) GetFields() map[string]string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Entry) GetEvent() *Entry_Message {
	if m != nil {
		return m.Event
	}
	return nil
}

// Message represents a serialized protobuf message.
// The name is the name registered with lion.
type Entry_Message struct {
	Encoding string `protobuf:"bytes,1,opt,name=encoding" json:"encoding,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Value    []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Entry_Message) Reset()                    { *m = Entry_Message{} }
func (m *Entry_Message) String() string            { return proto.CompactTextString(m) }
func (*Entry_Message) ProtoMessage()               {}
func (*Entry_Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func init() {
	proto.RegisterType((*Entry)(nil), "lion.Entry")
	proto.RegisterType((*Entry_Message)(nil), "lion.Entry.Message")
	proto.RegisterEnum("lion.Level", Level_name, Level_value)
}

var fileDescriptor0 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0x5d, 0x8f, 0x93, 0x40,
	0x14, 0x95, 0xef, 0xe5, 0x52, 0xd7, 0xc9, 0xa8, 0x71, 0xc2, 0x8b, 0xab, 0xbe, 0xa8, 0x89, 0x34,
	0x59, 0x5f, 0x56, 0xdf, 0x58, 0xa5, 0x66, 0x13, 0x04, 0x33, 0x59, 0x35, 0xf1, 0x65, 0x43, 0xdb,
	0x29, 0x21, 0x02, 0xd3, 0xc0, 0x50, 0x6d, 0xfc, 0x87, 0xfe, 0x2a, 0xe9, 0x0c, 0x94, 0x3e, 0xf8,
	0x42, 0xce, 0x39, 0xf7, 0xdc, 0x33, 0x77, 0xee, 0x00, 0x8f, 0xb7, 0x0d, 0x17, 0x7c, 0x2e, 0xbf,
	0x65, 0xc1, 0xeb, 0x40, 0x22, 0x6c, 0x1e, 0xb0, 0xff, 0x34, 0xe7, 0x3c, 0x2f, 0x99, 0xaa, 0x2e,
	0xbb, 0xcd, 0x5c, 0x14, 0x15, 0x6b, 0x45, 0x56, 0x6d, 0x95, 0xed, 0xf9, 0x5f, 0x03, 0xac, 0xa8,
	0x16, 0xcd, 0x1e, 0x9f, 0x83, 0x5e, 0xac, 0x89, 0x76, 0xa1, 0xbd, 0x74, 0x69, 0x8f, 0xf0, 0x33,
	0xb0, 0x4a, 0xb6, 0x63, 0x25, 0xd1, 0x7b, 0xe9, 0xfc, 0xd2, 0x0b, 0x64, 0x78, 0x7c, 0x90, 0xa8,
	0xaa, 0xe0, 0x2b, 0x70, 0x8f, 0x79, 0xc4, 0xe8, 0x6d, 0xde, 0xa5, 0x1f, 0xa8, 0x13, 0x83, 0xf1,
	0xc4, 0xe0, 0x76, 0x74, 0xd0, 0xc9, 0x8c, 0xdf, 0x80, 0xb3, 0xe2, 0xb5, 0x60, 0xbf, 0x05, 0x31,
	0x2f, 0x8c, 0xbe, 0xef, 0xa1, 0x8a, 0x97, 0xa3, 0x04, 0x9f, 0x59, 0xdb, 0x66, 0x39, 0xa3, 0xa3,
	0x07, 0xcf, 0xc1, 0xde, 0x14, 0xac, 0x5c, 0xb7, 0xc4, 0x92, 0xee, 0x27, 0xa7, 0xee, 0x85, 0xac,
	0x48, 0x4c, 0x07, 0x1b, 0x7e, 0x05, 0x56, 0x3f, 0x61, 0x2d, 0x88, 0x2d, 0xa7, 0xfa, 0x6f, 0xba,
	0x72, 0x60, 0x02, 0x4e, 0xa5, 0x14, 0xe2, 0xc8, 0xcb, 0x8f, 0x14, 0xbf, 0x80, 0xfb, 0xbf, 0x9a,
	0x42, 0xb0, 0xe6, 0x8e, 0x77, 0x62, 0xdb, 0x09, 0x72, 0xd6, 0xd7, 0x67, 0x74, 0xa6, 0xc4, 0x54,
	0x6a, 0x7e, 0x0a, 0xce, 0x10, 0x88, 0x7d, 0x38, 0x63, 0xf5, 0x8a, 0xaf, 0x8b, 0x3a, 0x1f, 0xf6,
	0x78, 0xe4, 0x18, 0x83, 0x59, 0x67, 0x15, 0x93, 0xcb, 0x74, 0xa9, 0xc4, 0xf8, 0x11, 0x58, 0xbb,
	0xac, 0xec, 0x98, 0x5c, 0xdd, 0x8c, 0x2a, 0xe2, 0xbf, 0x03, 0xef, 0xe4, 0x46, 0x18, 0x81, 0xf1,
	0x93, 0xed, 0x87, 0xbc, 0x03, 0x9c, 0xda, 0x54, 0x96, 0x22, 0xef, 0xf5, 0x2b, 0xed, 0xf5, 0x1f,
	0xb0, 0xe4, 0xfb, 0xe0, 0x07, 0xe0, 0xc5, 0xd1, 0xb7, 0x28, 0xbe, 0xfb, 0x18, 0x5d, 0x7f, 0xfd,
	0x84, 0xee, 0xf5, 0x8f, 0x0b, 0x4a, 0xb8, 0x49, 0x16, 0x29, 0xd2, 0x26, 0xfe, 0x3d, 0xa4, 0x09,
	0xd2, 0xa7, 0x86, 0x88, 0xd2, 0x94, 0x22, 0x63, 0x12, 0x16, 0xe1, 0x6d, 0x18, 0x23, 0x73, 0x12,
	0xbe, 0x84, 0xc9, 0xcd, 0x07, 0x64, 0x4d, 0x11, 0x49, 0x9a, 0x44, 0xc8, 0xbe, 0xf6, 0x7e, 0xb8,
	0xc7, 0x7f, 0x70, 0x69, 0x4b, 0xf8, 0xf6, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xd0, 0x4d,
	0xc7, 0x9d, 0x02, 0x00, 0x00,
}
