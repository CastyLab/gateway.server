// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
	messages "github.com/CastyLab/grpc.proto/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ChatMsgEvent struct {
	Message              []byte               `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	From                 string               `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   string               `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	User                 *messages.User       `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Emojies              []uint64             `protobuf:"varint,5,rep,packed,name=emojies,proto3" json:"emojies,omitempty"`
	Mentions             []uint64             `protobuf:"varint,6,rep,packed,name=mentions,proto3" json:"mentions,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ChatMsgEvent) Reset()         { *m = ChatMsgEvent{} }
func (m *ChatMsgEvent) String() string { return proto.CompactTextString(m) }
func (*ChatMsgEvent) ProtoMessage()    {}
func (*ChatMsgEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

func (m *ChatMsgEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMsgEvent.Unmarshal(m, b)
}
func (m *ChatMsgEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMsgEvent.Marshal(b, m, deterministic)
}
func (m *ChatMsgEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMsgEvent.Merge(m, src)
}
func (m *ChatMsgEvent) XXX_Size() int {
	return xxx_messageInfo_ChatMsgEvent.Size(m)
}
func (m *ChatMsgEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMsgEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMsgEvent proto.InternalMessageInfo

func (m *ChatMsgEvent) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *ChatMsgEvent) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ChatMsgEvent) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *ChatMsgEvent) GetUser() *messages.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ChatMsgEvent) GetEmojies() []uint64 {
	if m != nil {
		return m.Emojies
	}
	return nil
}

func (m *ChatMsgEvent) GetMentions() []uint64 {
	if m != nil {
		return m.Mentions
	}
	return nil
}

func (m *ChatMsgEvent) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*ChatMsgEvent)(nil), "protobuf.ChatMsgEvent")
}

func init() {
	proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54)
}

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0xe5, 0x34, 0xf4, 0xcf, 0x51, 0x75, 0x30, 0x8b, 0x95, 0x85, 0xa8, 0x53, 0x26, 0x47,
	0x82, 0x89, 0x11, 0x21, 0x46, 0x16, 0x0b, 0x66, 0xe4, 0x96, 0x6b, 0x1a, 0x84, 0x73, 0x95, 0xef,
	0xca, 0xf7, 0xe5, 0x9b, 0xa0, 0x3a, 0x71, 0x27, 0xfb, 0xf9, 0xf7, 0xf4, 0x9e, 0x1f, 0xc0, 0xfe,
	0xe8, 0xc5, 0x9e, 0x22, 0x09, 0xe9, 0x65, 0x3a, 0x76, 0xe7, 0x43, 0x75, 0x17, 0x90, 0xd9, 0x77,
	0xc8, 0xed, 0x99, 0x31, 0x8e, 0xb8, 0xba, 0xef, 0x88, 0xba, 0x1f, 0x6c, 0xb3, 0xab, 0x95, 0x3e,
	0x20, 0x8b, 0x0f, 0xa7, 0xd1, 0xb0, 0xfd, 0x53, 0xb0, 0x7e, 0x39, 0x7a, 0x79, 0xe3, 0xee, 0xf5,
	0x17, 0x07, 0xd1, 0x06, 0x16, 0x53, 0x90, 0x51, 0xb5, 0x6a, 0xd6, 0x2e, 0x4b, 0xad, 0xa1, 0x3c,
	0x44, 0x0a, 0xa6, 0xa8, 0x55, 0xb3, 0x72, 0xe9, 0xae, 0x37, 0x50, 0x08, 0x99, 0x59, 0x7a, 0x29,
	0x84, 0xf4, 0x16, 0xca, 0x4b, 0xbb, 0x29, 0x6b, 0xd5, 0xdc, 0x3e, 0x6c, 0x6c, 0xfe, 0x93, 0xfd,
	0x60, 0x8c, 0x2e, 0xb1, 0x4b, 0x03, 0x06, 0xfa, 0xee, 0x91, 0xcd, 0x4d, 0x3d, 0x6b, 0x4a, 0x97,
	0xa5, 0xae, 0x60, 0x19, 0x70, 0x90, 0x9e, 0x06, 0x36, 0xf3, 0x84, 0xae, 0x5a, 0x3f, 0x01, 0xec,
	0x23, 0x7a, 0xc1, 0xaf, 0x4f, 0x2f, 0x66, 0x91, 0xf2, 0x2b, 0x3b, 0xce, 0xb3, 0x79, 0x9e, 0x7d,
	0xcf, 0xf3, 0xdc, 0x6a, 0x72, 0x3f, 0xcb, 0x6e, 0x9e, 0xf0, 0xe3, 0x7f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x68, 0x09, 0x76, 0x6b, 0x38, 0x01, 0x00, 0x00,
}
