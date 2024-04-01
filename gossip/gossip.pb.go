// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: gossip.proto

package gossip

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventType int32

const (
	EventType_NoopEventType    EventType = 0
	EventType_MessageEventType EventType = 1
	EventType_UserEventType    EventType = 2
	EventType_DoneEventType    EventType = 3
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "NoopEventType",
		1: "MessageEventType",
		2: "UserEventType",
		3: "DoneEventType",
	}
	EventType_value = map[string]int32{
		"NoopEventType":    0,
		"MessageEventType": 1,
		"UserEventType":    2,
		"DoneEventType":    3,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_gossip_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_gossip_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{0}
}

type NetworkStatus int32

const (
	NetworkStatus_Online  NetworkStatus = 0
	NetworkStatus_Offline NetworkStatus = 1
)

// Enum value maps for NetworkStatus.
var (
	NetworkStatus_name = map[int32]string{
		0: "Online",
		1: "Offline",
	}
	NetworkStatus_value = map[string]int32{
		"Online":  0,
		"Offline": 1,
	}
)

func (x NetworkStatus) Enum() *NetworkStatus {
	p := new(NetworkStatus)
	*p = x
	return p
}

func (x NetworkStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_gossip_proto_enumTypes[1].Descriptor()
}

func (NetworkStatus) Type() protoreflect.EnumType {
	return &file_gossip_proto_enumTypes[1]
}

func (x NetworkStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetworkStatus.Descriptor instead.
func (NetworkStatus) EnumDescriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{1}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      EventType     `protobuf:"varint,1,opt,name=type,proto3,enum=gossip.EventType" json:"type,omitempty"`
	FromId    *string       `protobuf:"bytes,2,opt,name=from_id,json=fromId,proto3,oneof" json:"from_id,omitempty"`
	ToId      *string       `protobuf:"bytes,3,opt,name=to_id,json=toId,proto3,oneof" json:"to_id,omitempty"`
	Timestamp uint64        `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Noop      *NoopEvent    `protobuf:"bytes,5,opt,name=noop,proto3,oneof" json:"noop,omitempty"`
	Message   *MessageEvent `protobuf:"bytes,6,opt,name=message,proto3,oneof" json:"message,omitempty"`
	User      *UserEvent    `protobuf:"bytes,7,opt,name=user,proto3,oneof" json:"user,omitempty"`
	Done      *DoneEvent    `protobuf:"bytes,8,opt,name=done,proto3,oneof" json:"done,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_NoopEventType
}

func (x *Event) GetFromId() string {
	if x != nil && x.FromId != nil {
		return *x.FromId
	}
	return ""
}

func (x *Event) GetToId() string {
	if x != nil && x.ToId != nil {
		return *x.ToId
	}
	return ""
}

func (x *Event) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Event) GetNoop() *NoopEvent {
	if x != nil {
		return x.Noop
	}
	return nil
}

func (x *Event) GetMessage() *MessageEvent {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Event) GetUser() *UserEvent {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Event) GetDone() *DoneEvent {
	if x != nil {
		return x.Done
	}
	return nil
}

type NoopEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoopEvent) Reset() {
	*x = NoopEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoopEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoopEvent) ProtoMessage() {}

func (x *NoopEvent) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoopEvent.ProtoReflect.Descriptor instead.
func (*NoopEvent) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{3}
}

type MessageEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *MessageEvent) Reset() {
	*x = MessageEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageEvent) ProtoMessage() {}

func (x *MessageEvent) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageEvent.ProtoReflect.Descriptor instead.
func (*MessageEvent) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{4}
}

func (x *MessageEvent) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type UserEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string        `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status NetworkStatus `protobuf:"varint,2,opt,name=status,proto3,enum=gossip.NetworkStatus" json:"status,omitempty"`
}

func (x *UserEvent) Reset() {
	*x = UserEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEvent) ProtoMessage() {}

func (x *UserEvent) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEvent.ProtoReflect.Descriptor instead.
func (*UserEvent) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{5}
}

func (x *UserEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserEvent) GetStatus() NetworkStatus {
	if x != nil {
		return x.Status
	}
	return NetworkStatus_Online
}

type DoneEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string        `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status NetworkStatus `protobuf:"varint,2,opt,name=status,proto3,enum=gossip.NetworkStatus" json:"status,omitempty"`
}

func (x *DoneEvent) Reset() {
	*x = DoneEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gossip_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoneEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoneEvent) ProtoMessage() {}

func (x *DoneEvent) ProtoReflect() protoreflect.Message {
	mi := &file_gossip_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoneEvent.ProtoReflect.Descriptor instead.
func (*DoneEvent) Descriptor() ([]byte, []int) {
	return file_gossip_proto_rawDescGZIP(), []int{6}
}

func (x *DoneEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DoneEvent) GetStatus() NetworkStatus {
	if x != nil {
		return x.Status
	}
	return NetworkStatus_Online
}

var File_gossip_proto protoreflect.FileDescriptor

var file_gossip_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x16, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xfa, 0x02, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x11, 0x2e, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x07, 0x66, 0x72, 0x6f, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x66, 0x72, 0x6f,
	0x6d, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x05, 0x74, 0x6f, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2a,
	0x0a, 0x04, 0x6e, 0x6f, 0x6f, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67,
	0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x4e, 0x6f, 0x6f, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48,
	0x02, 0x52, 0x04, 0x6e, 0x6f, 0x6f, 0x70, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x73, 0x73, 0x69, 0x70, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x48, 0x03, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x2a, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x48, 0x04, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x04, 0x64,
	0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x73, 0x73,
	0x69, 0x70, 0x2e, 0x44, 0x6f, 0x6e, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x05, 0x52, 0x04,
	0x64, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x66, 0x72, 0x6f, 0x6d,
	0x5f, 0x69, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x6e, 0x6f, 0x6f, 0x70, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x64, 0x6f, 0x6e, 0x65, 0x22, 0x0b, 0x0a, 0x09, 0x4e, 0x6f, 0x6f, 0x70, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x28, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x53, 0x0a, 0x09, 0x55,
	0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x53, 0x0a, 0x09, 0x44, 0x6f, 0x6e, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x5a, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x6f, 0x6f, 0x70, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x55,
	0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x02, 0x12, 0x11,
	0x0a, 0x0d, 0x44, 0x6f, 0x6e, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10,
	0x03, 0x2a, 0x28, 0x0a, 0x0d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x01, 0x32, 0xa6, 0x01, 0x0a, 0x09,
	0x47, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x41, 0x70, 0x69, 0x12, 0x34, 0x0a, 0x11, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0c,
	0x2e, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0d, 0x2e, 0x67,
	0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x2b, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0d, 0x2e, 0x67,
	0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x0d, 0x2e, 0x67, 0x6f,
	0x73, 0x73, 0x69, 0x70, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x15,
	0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0c, 0x2e, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x1a, 0x0d, 0x2e, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x67, 0x6f, 0x73, 0x73, 0x69,
	0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gossip_proto_rawDescOnce sync.Once
	file_gossip_proto_rawDescData = file_gossip_proto_rawDesc
)

func file_gossip_proto_rawDescGZIP() []byte {
	file_gossip_proto_rawDescOnce.Do(func() {
		file_gossip_proto_rawDescData = protoimpl.X.CompressGZIP(file_gossip_proto_rawDescData)
	})
	return file_gossip_proto_rawDescData
}

var file_gossip_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_gossip_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_gossip_proto_goTypes = []interface{}{
	(EventType)(0),       // 0: gossip.EventType
	(NetworkStatus)(0),   // 1: gossip.NetworkStatus
	(*Empty)(nil),        // 2: gossip.Empty
	(*User)(nil),         // 3: gossip.User
	(*Event)(nil),        // 4: gossip.Event
	(*NoopEvent)(nil),    // 5: gossip.NoopEvent
	(*MessageEvent)(nil), // 6: gossip.MessageEvent
	(*UserEvent)(nil),    // 7: gossip.UserEvent
	(*DoneEvent)(nil),    // 8: gossip.DoneEvent
}
var file_gossip_proto_depIdxs = []int32{
	0,  // 0: gossip.Event.type:type_name -> gossip.EventType
	5,  // 1: gossip.Event.noop:type_name -> gossip.NoopEvent
	6,  // 2: gossip.Event.message:type_name -> gossip.MessageEvent
	7,  // 3: gossip.Event.user:type_name -> gossip.UserEvent
	8,  // 4: gossip.Event.done:type_name -> gossip.DoneEvent
	1,  // 5: gossip.UserEvent.status:type_name -> gossip.NetworkStatus
	1,  // 6: gossip.DoneEvent.status:type_name -> gossip.NetworkStatus
	3,  // 7: gossip.GossipApi.SubscribeToEvents:input_type -> gossip.User
	4,  // 8: gossip.GossipApi.SendEvent:input_type -> gossip.Event
	3,  // 9: gossip.GossipApi.UnsubscribeFromEvents:input_type -> gossip.User
	4,  // 10: gossip.GossipApi.SubscribeToEvents:output_type -> gossip.Event
	2,  // 11: gossip.GossipApi.SendEvent:output_type -> gossip.Empty
	2,  // 12: gossip.GossipApi.UnsubscribeFromEvents:output_type -> gossip.Empty
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_gossip_proto_init() }
func file_gossip_proto_init() {
	if File_gossip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gossip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gossip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gossip_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gossip_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoopEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gossip_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gossip_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gossip_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoneEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_gossip_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gossip_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gossip_proto_goTypes,
		DependencyIndexes: file_gossip_proto_depIdxs,
		EnumInfos:         file_gossip_proto_enumTypes,
		MessageInfos:      file_gossip_proto_msgTypes,
	}.Build()
	File_gossip_proto = out.File
	file_gossip_proto_rawDesc = nil
	file_gossip_proto_goTypes = nil
	file_gossip_proto_depIdxs = nil
}
