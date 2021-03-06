// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: chat.proto

package protobuf

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

type IdentifierPrefix int32

const (
	IdentifierPrefix_USER     IdentifierPrefix = 0
	IdentifierPrefix_MEG      IdentifierPrefix = 1
	IdentifierPrefix_ROOM     IdentifierPrefix = 2
	IdentifierPrefix_GROUP    IdentifierPrefix = 3
	IdentifierPrefix_REACTION IdentifierPrefix = 4
)

// Enum value maps for IdentifierPrefix.
var (
	IdentifierPrefix_name = map[int32]string{
		0: "USER",
		1: "MEG",
		2: "ROOM",
		3: "GROUP",
		4: "REACTION",
	}
	IdentifierPrefix_value = map[string]int32{
		"USER":     0,
		"MEG":      1,
		"ROOM":     2,
		"GROUP":    3,
		"REACTION": 4,
	}
)

func (x IdentifierPrefix) Enum() *IdentifierPrefix {
	p := new(IdentifierPrefix)
	*p = x
	return p
}

func (x IdentifierPrefix) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IdentifierPrefix) Descriptor() protoreflect.EnumDescriptor {
	return file_chat_proto_enumTypes[0].Descriptor()
}

func (IdentifierPrefix) Type() protoreflect.EnumType {
	return &file_chat_proto_enumTypes[0]
}

func (x IdentifierPrefix) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IdentifierPrefix.Descriptor instead.
func (IdentifierPrefix) EnumDescriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

type ReactionType int32

const (
	ReactionType_LIKE    ReactionType = 0
	ReactionType_DISLIKE ReactionType = 1
	ReactionType_LOVE    ReactionType = 2
)

// Enum value maps for ReactionType.
var (
	ReactionType_name = map[int32]string{
		0: "LIKE",
		1: "DISLIKE",
		2: "LOVE",
	}
	ReactionType_value = map[string]int32{
		"LIKE":    0,
		"DISLIKE": 1,
		"LOVE":    2,
	}
)

func (x ReactionType) Enum() *ReactionType {
	p := new(ReactionType)
	*p = x
	return p
}

func (x ReactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_chat_proto_enumTypes[1].Descriptor()
}

func (ReactionType) Type() protoreflect.EnumType {
	return &file_chat_proto_enumTypes[1]
}

func (x ReactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReactionType.Descriptor instead.
func (ReactionType) EnumDescriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

type Identifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefix IdentifierPrefix `protobuf:"varint,1,opt,name=prefix,proto3,enum=models.IdentifierPrefix" json:"prefix,omitempty"`
	Id     string           `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Identifier) Reset() {
	*x = Identifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identifier) ProtoMessage() {}

func (x *Identifier) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identifier.ProtoReflect.Descriptor instead.
func (*Identifier) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Identifier) GetPrefix() IdentifierPrefix {
	if x != nil {
		return x.Prefix
	}
	return IdentifierPrefix_USER
}

func (x *Identifier) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string      `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string      `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Age       int32       `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	EMailId   string      `protobuf:"bytes,5,opt,name=e_mail_id,json=eMailId,proto3" json:"e_mail_id,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
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
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() *Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *User) GetEMailId() string {
	if x != nil {
		return x.EMailId
	}
	return ""
}

type Reaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type ReactionType `protobuf:"varint,1,opt,name=type,proto3,enum=models.ReactionType" json:"type,omitempty"`
}

func (x *Reaction) Reset() {
	*x = Reaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reaction) ProtoMessage() {}

func (x *Reaction) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reaction.ProtoReflect.Descriptor instead.
func (*Reaction) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *Reaction) GetType() ReactionType {
	if x != nil {
		return x.Type
	}
	return ReactionType_LIKE
}

type GroupInChatRoom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *Identifier   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Admin     *Identifier   `protobuf:"bytes,3,opt,name=admin,proto3" json:"admin,omitempty"`
	CreatedBy *Identifier   `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Users     []*Identifier `protobuf:"bytes,5,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GroupInChatRoom) Reset() {
	*x = GroupInChatRoom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupInChatRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupInChatRoom) ProtoMessage() {}

func (x *GroupInChatRoom) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupInChatRoom.ProtoReflect.Descriptor instead.
func (*GroupInChatRoom) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{3}
}

func (x *GroupInChatRoom) GetId() *Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *GroupInChatRoom) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GroupInChatRoom) GetAdmin() *Identifier {
	if x != nil {
		return x.Admin
	}
	return nil
}

func (x *GroupInChatRoom) GetCreatedBy() *Identifier {
	if x != nil {
		return x.CreatedBy
	}
	return nil
}

func (x *GroupInChatRoom) GetUsers() []*Identifier {
	if x != nil {
		return x.Users
	}
	return nil
}

type ChatRoom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *Identifier   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Admin     *Identifier   `protobuf:"bytes,3,opt,name=admin,proto3" json:"admin,omitempty"`
	CreatedBy *Identifier   `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	Users     []*Identifier `protobuf:"bytes,5,rep,name=users,proto3" json:"users,omitempty"`
	Groups    []*Identifier `protobuf:"bytes,6,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *ChatRoom) Reset() {
	*x = ChatRoom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRoom) ProtoMessage() {}

func (x *ChatRoom) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRoom.ProtoReflect.Descriptor instead.
func (*ChatRoom) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{4}
}

func (x *ChatRoom) GetId() *Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ChatRoom) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChatRoom) GetAdmin() *Identifier {
	if x != nil {
		return x.Admin
	}
	return nil
}

func (x *ChatRoom) GetCreatedBy() *Identifier {
	if x != nil {
		return x.CreatedBy
	}
	return nil
}

func (x *ChatRoom) GetUsers() []*Identifier {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *ChatRoom) GetGroups() []*Identifier {
	if x != nil {
		return x.Groups
	}
	return nil
}

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              *Identifier   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender          *Identifier   `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Data            string        `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	RepliedTo       *Identifier   `protobuf:"bytes,4,opt,name=replied_to,json=repliedTo,proto3" json:"replied_to,omitempty"`
	ChatRoom        *Identifier   `protobuf:"bytes,5,opt,name=chat_room,json=chatRoom,proto3" json:"chat_room,omitempty"`
	GroupInChatRoom *Identifier   `protobuf:"bytes,6,opt,name=group_in_chat_room,json=groupInChatRoom,proto3" json:"group_in_chat_room,omitempty"`
	Reactions       []*Identifier `protobuf:"bytes,7,rep,name=reactions,proto3" json:"reactions,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{5}
}

func (x *ChatMessage) GetId() *Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ChatMessage) GetSender() *Identifier {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *ChatMessage) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *ChatMessage) GetRepliedTo() *Identifier {
	if x != nil {
		return x.RepliedTo
	}
	return nil
}

func (x *ChatMessage) GetChatRoom() *Identifier {
	if x != nil {
		return x.ChatRoom
	}
	return nil
}

func (x *ChatMessage) GetGroupInChatRoom() *Identifier {
	if x != nil {
		return x.GroupInChatRoom
	}
	return nil
}

func (x *ChatMessage) GetReactions() []*Identifier {
	if x != nil {
		return x.Reactions
	}
	return nil
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x22, 0x4e, 0x0a, 0x0a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x06, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x22, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x09, 0x65, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x08, 0x52,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0xd0, 0x01, 0x0a, 0x0f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x22, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a,
	0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x22, 0xf5, 0x01, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f,
	0x6d, 0x12, 0x22, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x05, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x2a, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0xc8, 0x02, 0x0a,
	0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2a, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x31, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x65,
	0x64, 0x54, 0x6f, 0x12, 0x2f, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x6d,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x3f, 0x0a, 0x12, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x6e,
	0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x52, 0x0f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x30, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x09, 0x72, 0x65,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a, 0x48, 0x0a, 0x10, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x08, 0x0a, 0x04, 0x55,
	0x53, 0x45, 0x52, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x45, 0x47, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x52, 0x4f, 0x55,
	0x50, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x04, 0x2a, 0x2f, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44,
	0x49, 0x53, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x4f, 0x56, 0x45,
	0x10, 0x02, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_chat_proto_goTypes = []interface{}{
	(IdentifierPrefix)(0),   // 0: models.IdentifierPrefix
	(ReactionType)(0),       // 1: models.ReactionType
	(*Identifier)(nil),      // 2: models.Identifier
	(*User)(nil),            // 3: models.User
	(*Reaction)(nil),        // 4: models.Reaction
	(*GroupInChatRoom)(nil), // 5: models.GroupInChatRoom
	(*ChatRoom)(nil),        // 6: models.ChatRoom
	(*ChatMessage)(nil),     // 7: models.ChatMessage
}
var file_chat_proto_depIdxs = []int32{
	0,  // 0: models.Identifier.prefix:type_name -> models.IdentifierPrefix
	2,  // 1: models.User.id:type_name -> models.Identifier
	1,  // 2: models.Reaction.type:type_name -> models.ReactionType
	2,  // 3: models.GroupInChatRoom.id:type_name -> models.Identifier
	2,  // 4: models.GroupInChatRoom.admin:type_name -> models.Identifier
	2,  // 5: models.GroupInChatRoom.created_by:type_name -> models.Identifier
	2,  // 6: models.GroupInChatRoom.users:type_name -> models.Identifier
	2,  // 7: models.ChatRoom.id:type_name -> models.Identifier
	2,  // 8: models.ChatRoom.admin:type_name -> models.Identifier
	2,  // 9: models.ChatRoom.created_by:type_name -> models.Identifier
	2,  // 10: models.ChatRoom.users:type_name -> models.Identifier
	2,  // 11: models.ChatRoom.groups:type_name -> models.Identifier
	2,  // 12: models.ChatMessage.id:type_name -> models.Identifier
	2,  // 13: models.ChatMessage.sender:type_name -> models.Identifier
	2,  // 14: models.ChatMessage.replied_to:type_name -> models.Identifier
	2,  // 15: models.ChatMessage.chat_room:type_name -> models.Identifier
	2,  // 16: models.ChatMessage.group_in_chat_room:type_name -> models.Identifier
	2,  // 17: models.ChatMessage.reactions:type_name -> models.Identifier
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identifier); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reaction); i {
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
		file_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupInChatRoom); i {
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
		file_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRoom); i {
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
		file_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		EnumInfos:         file_chat_proto_enumTypes,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}
