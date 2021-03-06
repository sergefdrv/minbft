// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

// Package pb defines MinBFT protocol messages in Protobuf format.

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// MessageType represents message type tag.
type MessageType int32

const (
	MessageType_UNKNOWN MessageType = 0
	MessageType_REQUEST MessageType = 1
	MessageType_REPLY   MessageType = 2
	MessageType_PREPARE MessageType = 3
	MessageType_COMMIT  MessageType = 4
)

var MessageType_name = map[int32]string{
	0: "UNKNOWN",
	1: "REQUEST",
	2: "REPLY",
	3: "PREPARE",
	4: "COMMIT",
}

var MessageType_value = map[string]int32{
	"UNKNOWN": 0,
	"REQUEST": 1,
	"REPLY":   2,
	"PREPARE": 3,
	"COMMIT":  4,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

// Message represents arbitrary protocol message.
type Message struct {
	// Types that are valid to be assigned to Typed:
	//	*Message_Request
	//	*Message_Reply
	//	*Message_Prepare
	//	*Message_Commit
	Typed                isMessage_Typed `protobuf_oneof:"typed"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

type isMessage_Typed interface {
	isMessage_Typed()
}

type Message_Request struct {
	Request *Request `protobuf:"bytes,1,opt,name=request,proto3,oneof"`
}

type Message_Reply struct {
	Reply *Reply `protobuf:"bytes,2,opt,name=reply,proto3,oneof"`
}

type Message_Prepare struct {
	Prepare *Prepare `protobuf:"bytes,3,opt,name=prepare,proto3,oneof"`
}

type Message_Commit struct {
	Commit *Commit `protobuf:"bytes,4,opt,name=commit,proto3,oneof"`
}

func (*Message_Request) isMessage_Typed() {}

func (*Message_Reply) isMessage_Typed() {}

func (*Message_Prepare) isMessage_Typed() {}

func (*Message_Commit) isMessage_Typed() {}

func (m *Message) GetTyped() isMessage_Typed {
	if m != nil {
		return m.Typed
	}
	return nil
}

func (m *Message) GetRequest() *Request {
	if x, ok := m.GetTyped().(*Message_Request); ok {
		return x.Request
	}
	return nil
}

func (m *Message) GetReply() *Reply {
	if x, ok := m.GetTyped().(*Message_Reply); ok {
		return x.Reply
	}
	return nil
}

func (m *Message) GetPrepare() *Prepare {
	if x, ok := m.GetTyped().(*Message_Prepare); ok {
		return x.Prepare
	}
	return nil
}

func (m *Message) GetCommit() *Commit {
	if x, ok := m.GetTyped().(*Message_Commit); ok {
		return x.Commit
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_Request)(nil),
		(*Message_Reply)(nil),
		(*Message_Prepare)(nil),
		(*Message_Commit)(nil),
	}
}

// Request represents REQUEST message.
type Request struct {
	// Client identifier
	ClientId uint32 `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// Request identifier (timestamp / sequence number)
	Seq uint64 `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	// Operation to execute on replicated state machine
	Operation []byte `protobuf:"bytes,3,opt,name=operation,proto3" json:"operation,omitempty"`
	// Client's signature
	Signature            []byte   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Request) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Request) GetOperation() []byte {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *Request) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// Reply represents REPLY message.
type Reply struct {
	// Replica identifier
	ReplicaId uint32 `protobuf:"varint,1,opt,name=replica_id,json=replicaId,proto3" json:"replica_id,omitempty"`
	// Client identifier
	ClientId uint32 `protobuf:"varint,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// Request identifier
	Seq uint64 `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`
	// Result of requested operation execution
	Result []byte `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
	// Replica's signature
	Signature            []byte   `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetReplicaId() uint32 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

func (m *Reply) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Reply) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Reply) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Reply) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// Prepare represents PREPARE message.
type Prepare struct {
	// Replica identifier
	ReplicaId uint32 `protobuf:"varint,1,opt,name=replica_id,json=replicaId,proto3" json:"replica_id,omitempty"`
	// View number
	View uint64 `protobuf:"varint,2,opt,name=view,proto3" json:"view,omitempty"`
	// Client's REQUEST
	Request *Request `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	// Replica's UI
	Ui                   []byte   `protobuf:"bytes,4,opt,name=ui,proto3" json:"ui,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Prepare) Reset()         { *m = Prepare{} }
func (m *Prepare) String() string { return proto.CompactTextString(m) }
func (*Prepare) ProtoMessage()    {}
func (*Prepare) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{3}
}

func (m *Prepare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Prepare.Unmarshal(m, b)
}
func (m *Prepare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Prepare.Marshal(b, m, deterministic)
}
func (m *Prepare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prepare.Merge(m, src)
}
func (m *Prepare) XXX_Size() int {
	return xxx_messageInfo_Prepare.Size(m)
}
func (m *Prepare) XXX_DiscardUnknown() {
	xxx_messageInfo_Prepare.DiscardUnknown(m)
}

var xxx_messageInfo_Prepare proto.InternalMessageInfo

func (m *Prepare) GetReplicaId() uint32 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

func (m *Prepare) GetView() uint64 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *Prepare) GetRequest() *Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Prepare) GetUi() []byte {
	if m != nil {
		return m.Ui
	}
	return nil
}

// Commit represents COMMIT message.
type Commit struct {
	// Replica identifier
	ReplicaId uint32 `protobuf:"varint,1,opt,name=replica_id,json=replicaId,proto3" json:"replica_id,omitempty"`
	// Primary's PREPARE
	Prepare *Prepare `protobuf:"bytes,2,opt,name=prepare,proto3" json:"prepare,omitempty"`
	// Replica's UI
	Ui                   []byte   `protobuf:"bytes,3,opt,name=ui,proto3" json:"ui,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Commit) Reset()         { *m = Commit{} }
func (m *Commit) String() string { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()    {}
func (*Commit) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{4}
}

func (m *Commit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commit.Unmarshal(m, b)
}
func (m *Commit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commit.Marshal(b, m, deterministic)
}
func (m *Commit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commit.Merge(m, src)
}
func (m *Commit) XXX_Size() int {
	return xxx_messageInfo_Commit.Size(m)
}
func (m *Commit) XXX_DiscardUnknown() {
	xxx_messageInfo_Commit.DiscardUnknown(m)
}

var xxx_messageInfo_Commit proto.InternalMessageInfo

func (m *Commit) GetReplicaId() uint32 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

func (m *Commit) GetPrepare() *Prepare {
	if m != nil {
		return m.Prepare
	}
	return nil
}

func (m *Commit) GetUi() []byte {
	if m != nil {
		return m.Ui
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.MessageType", MessageType_name, MessageType_value)
	proto.RegisterType((*Message)(nil), "pb.Message")
	proto.RegisterType((*Request)(nil), "pb.Request")
	proto.RegisterType((*Reply)(nil), "pb.Reply")
	proto.RegisterType((*Prepare)(nil), "pb.Prepare")
	proto.RegisterType((*Commit)(nil), "pb.Commit")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xdd, 0x8a, 0xd4, 0x30,
	0x1c, 0xc5, 0xfb, 0x5d, 0xfb, 0xef, 0xba, 0x94, 0x5c, 0x48, 0x41, 0x05, 0x2d, 0x2e, 0x8a, 0x17,
	0x73, 0xa1, 0x4f, 0xa0, 0x4b, 0x61, 0x86, 0x75, 0x66, 0x6b, 0x9c, 0x45, 0xbc, 0x51, 0x3a, 0xd3,
	0xb0, 0x04, 0x3a, 0xd3, 0x6c, 0x92, 0x2a, 0x7d, 0x06, 0x9f, 0xc5, 0x77, 0x94, 0x7c, 0x0c, 0x61,
	0x1c, 0x61, 0xee, 0x92, 0x73, 0xfe, 0x9c, 0xfc, 0x92, 0x13, 0xb8, 0xdc, 0x11, 0x21, 0xda, 0x7b,
	0x22, 0x66, 0x8c, 0x0f, 0x72, 0x40, 0x01, 0xdb, 0x54, 0x7f, 0x7c, 0x48, 0x97, 0x46, 0x46, 0xaf,
	0x21, 0xe5, 0xe4, 0x61, 0x24, 0x42, 0x96, 0xfe, 0x0b, 0xff, 0x4d, 0xfe, 0x2e, 0x9f, 0xb1, 0xcd,
	0x0c, 0x1b, 0x69, 0xee, 0xe1, 0x83, 0x8b, 0x5e, 0x42, 0xcc, 0x09, 0xeb, 0xa7, 0x32, 0xd0, 0x63,
	0x99, 0x19, 0x63, 0xfd, 0x34, 0xf7, 0xb0, 0x71, 0x54, 0x16, 0xe3, 0x84, 0xb5, 0x9c, 0x94, 0xa1,
	0xcb, 0x6a, 0x8c, 0xa4, 0xb2, 0xac, 0x8b, 0x5e, 0x41, 0xb2, 0x1d, 0x76, 0x3b, 0x2a, 0xcb, 0x48,
	0xcf, 0x81, 0x9a, 0xbb, 0xd6, 0xca, 0xdc, 0xc3, 0xd6, 0xfb, 0x98, 0x42, 0x2c, 0x27, 0x46, 0xba,
	0x4a, 0x42, 0x6a, 0x81, 0xd0, 0x53, 0xc8, 0xb6, 0x3d, 0x25, 0x7b, 0xf9, 0x83, 0x76, 0x1a, 0xf8,
	0x31, 0x7e, 0x64, 0x84, 0x45, 0x87, 0x0a, 0x08, 0x05, 0x79, 0xd0, 0x80, 0x11, 0x56, 0x4b, 0xf4,
	0x0c, 0xb2, 0x81, 0x11, 0xde, 0x4a, 0x3a, 0xec, 0x35, 0xd3, 0x05, 0x76, 0x82, 0x72, 0x05, 0xbd,
	0xdf, 0xb7, 0x72, 0xe4, 0x44, 0x93, 0x5c, 0x60, 0x27, 0x54, 0xbf, 0x7d, 0x88, 0xf5, 0x05, 0xd1,
	0x73, 0x00, 0x75, 0x41, 0xba, 0x6d, 0xdd, 0xa9, 0x99, 0x55, 0x16, 0xdd, 0x31, 0x53, 0xf0, 0x7f,
	0xa6, 0xd0, 0x31, 0x3d, 0x81, 0x84, 0x13, 0x31, 0xf6, 0xd2, 0x1e, 0x69, 0x77, 0xc7, 0x34, 0xf1,
	0xbf, 0x34, 0x02, 0x52, 0xfb, 0x90, 0xe7, 0x70, 0x10, 0x44, 0x3f, 0x29, 0xf9, 0x65, 0x9f, 0x41,
	0xaf, 0xd1, 0x95, 0x6b, 0x39, 0x3c, 0x69, 0xd9, 0x75, 0x7c, 0x09, 0xc1, 0x48, 0x2d, 0x56, 0x30,
	0xd2, 0xea, 0x3b, 0x24, 0xa6, 0x95, 0x73, 0x67, 0x5e, 0xb9, 0xe6, 0x83, 0x93, 0xe6, 0x5d, 0xef,
	0x26, 0x3f, 0x3c, 0xe4, 0xbf, 0xbd, 0x81, 0xdc, 0xfe, 0xc3, 0xf5, 0xc4, 0x08, 0xca, 0x21, 0xbd,
	0x5b, 0xdd, 0xac, 0x6e, 0xbf, 0xae, 0x0a, 0x4f, 0x6d, 0x70, 0xfd, 0xf9, 0xae, 0xfe, 0xb2, 0x2e,
	0x7c, 0x94, 0x41, 0x8c, 0xeb, 0xe6, 0xd3, 0xb7, 0x22, 0x50, 0x7a, 0x83, 0xeb, 0xe6, 0x03, 0xae,
	0x8b, 0x10, 0x01, 0x24, 0xd7, 0xb7, 0xcb, 0xe5, 0x62, 0x5d, 0x44, 0x9b, 0x44, 0x7f, 0xf0, 0xf7,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x39, 0xf4, 0xd1, 0xe0, 0xf2, 0x02, 0x00, 0x00,
}
