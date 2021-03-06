// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: server/options/transaction/options.proto

package transaction

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageOption_AccessorType int32

const (
	MessageOption_Unknown         MessageOption_AccessorType = 0
	MessageOption_OnlyServer      MessageOption_AccessorType = 1
	MessageOption_ServerAndClient MessageOption_AccessorType = 2
)

// Enum value maps for MessageOption_AccessorType.
var (
	MessageOption_AccessorType_name = map[int32]string{
		0: "Unknown",
		1: "OnlyServer",
		2: "ServerAndClient",
	}
	MessageOption_AccessorType_value = map[string]int32{
		"Unknown":         0,
		"OnlyServer":      1,
		"ServerAndClient": 2,
	}
)

func (x MessageOption_AccessorType) Enum() *MessageOption_AccessorType {
	p := new(MessageOption_AccessorType)
	*p = x
	return p
}

func (x MessageOption_AccessorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageOption_AccessorType) Descriptor() protoreflect.EnumDescriptor {
	return file_server_options_transaction_options_proto_enumTypes[0].Descriptor()
}

func (MessageOption_AccessorType) Type() protoreflect.EnumType {
	return &file_server_options_transaction_options_proto_enumTypes[0]
}

func (x MessageOption_AccessorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageOption_AccessorType.Descriptor instead.
func (MessageOption_AccessorType) EnumDescriptor() ([]byte, []int) {
	return file_server_options_transaction_options_proto_rawDescGZIP(), []int{0, 0}
}

type FieldOption_AccessorType int32

const (
	FieldOption_All        FieldOption_AccessorType = 0
	FieldOption_OnlyServer FieldOption_AccessorType = 1
	FieldOption_OnlyClient FieldOption_AccessorType = 2
)

// Enum value maps for FieldOption_AccessorType.
var (
	FieldOption_AccessorType_name = map[int32]string{
		0: "All",
		1: "OnlyServer",
		2: "OnlyClient",
	}
	FieldOption_AccessorType_value = map[string]int32{
		"All":        0,
		"OnlyServer": 1,
		"OnlyClient": 2,
	}
)

func (x FieldOption_AccessorType) Enum() *FieldOption_AccessorType {
	p := new(FieldOption_AccessorType)
	*p = x
	return p
}

func (x FieldOption_AccessorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldOption_AccessorType) Descriptor() protoreflect.EnumDescriptor {
	return file_server_options_transaction_options_proto_enumTypes[1].Descriptor()
}

func (FieldOption_AccessorType) Type() protoreflect.EnumType {
	return &file_server_options_transaction_options_proto_enumTypes[1]
}

func (x FieldOption_AccessorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FieldOption_AccessorType.Descriptor instead.
func (FieldOption_AccessorType) EnumDescriptor() ([]byte, []int) {
	return file_server_options_transaction_options_proto_rawDescGZIP(), []int{1, 0}
}

type MessageOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessorType MessageOption_AccessorType `protobuf:"varint,1,opt,name=accessor_type,json=accessorType,proto3,enum=server.options.transaction.MessageOption_AccessorType" json:"accessor_type,omitempty"`
	Ddl          *MessageOption_DDL         `protobuf:"bytes,2,opt,name=ddl,proto3" json:"ddl,omitempty"`
}

func (x *MessageOption) Reset() {
	*x = MessageOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_options_transaction_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOption) ProtoMessage() {}

func (x *MessageOption) ProtoReflect() protoreflect.Message {
	mi := &file_server_options_transaction_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageOption.ProtoReflect.Descriptor instead.
func (*MessageOption) Descriptor() ([]byte, []int) {
	return file_server_options_transaction_options_proto_rawDescGZIP(), []int{0}
}

func (x *MessageOption) GetAccessorType() MessageOption_AccessorType {
	if x != nil {
		return x.AccessorType
	}
	return MessageOption_Unknown
}

func (x *MessageOption) GetDdl() *MessageOption_DDL {
	if x != nil {
		return x.Ddl
	}
	return nil
}

type FieldOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessorType FieldOption_AccessorType `protobuf:"varint,1,opt,name=accessor_type,json=accessorType,proto3,enum=server.options.transaction.FieldOption_AccessorType" json:"accessor_type,omitempty"`
	Ddl          *FieldOption_DDL         `protobuf:"bytes,2,opt,name=ddl,proto3" json:"ddl,omitempty"`
}

func (x *FieldOption) Reset() {
	*x = FieldOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_options_transaction_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldOption) ProtoMessage() {}

func (x *FieldOption) ProtoReflect() protoreflect.Message {
	mi := &file_server_options_transaction_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldOption.ProtoReflect.Descriptor instead.
func (*FieldOption) Descriptor() ([]byte, []int) {
	return file_server_options_transaction_options_proto_rawDescGZIP(), []int{1}
}

func (x *FieldOption) GetAccessorType() FieldOption_AccessorType {
	if x != nil {
		return x.AccessorType
	}
	return FieldOption_All
}

func (x *FieldOption) GetDdl() *FieldOption_DDL {
	if x != nil {
		return x.Ddl
	}
	return nil
}

type MessageOption_DDL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Indexes    []*MessageOption_DDL_Index    `protobuf:"bytes,1,rep,name=indexes,proto3" json:"indexes,omitempty"`
	Interleave *MessageOption_DDL_Interleave `protobuf:"bytes,2,opt,name=interleave,proto3" json:"interleave,omitempty"`
}

func (x *MessageOption_DDL) Reset() {
	*x = MessageOption_DDL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_options_transaction_options_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageOption_DDL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOption_DDL) ProtoMessage() {}

func (x *MessageOption_DDL) ProtoReflect() protoreflect.Message {
	mi := &file_server_options_transaction_options_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageOption_DDL.ProtoReflect.Descriptor instead.
func (*MessageOption_DDL) Descriptor() ([]byte, []int) {
	return file_server_options_transaction_options_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MessageOption_DDL) GetIndexes() []*MessageOption_DDL_Index {
	if x != nil {
		return x.Indexes
	}
	return nil
}

func (x *MessageOption_DDL) GetInterleave() *MessageOption_DDL_Interleave {
	if x != nil {
		return x.Interleave
	}
	return nil
}

type MessageOption_DDL_Index struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys    []*MessageOption_DDL_Index_Key `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Unique  bool                           `protobuf:"varint,2,opt,name=unique,proto3" json:"unique,omitempty"`
	Storing []string                       `protobuf:"bytes,3,rep,name=storing,proto3" json:"storing,omitempty"`
}

func (x *MessageOption_DDL_Index) Reset() {
	*x = MessageOption_DDL_Index{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_options_transaction_options_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageOption_DDL_Index) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOption_DDL_Index) ProtoMessage() {}

func (x *MessageOption_DDL_Index) ProtoReflect() protoreflect.Message {
	mi := &file_server_options_transaction_options_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageOption_DDL_Index.ProtoReflect.Descriptor instead.
func (*MessageOption_DDL_Index) Descriptor() ([]byte, []int) {
	return file_server_options_transaction_options_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *MessageOption_DDL_Index) GetKeys() []*MessageOption_DDL_Index_Key {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *MessageOption_DDL_Index) GetUnique() bool {
	if x != nil {
		return x.Unique
	}
	return false
}

func (x *MessageOption_DDL_Index) GetStoring() []string {
	if x != nil {
		return x.Storing
	}
	return nil
}

type MessageOption_DDL_Interleave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table string `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
}

func (x *MessageOption_DDL_Interleave) Reset() {
	*x = MessageOption_DDL_Interleave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_options_transaction_options_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageOption_DDL_Interleave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOption_DDL_Interleave) ProtoMessage() {}

func (x *MessageOption_DDL_Interleave) ProtoReflect() protoreflect.Message {
	mi := &file_server_options_transaction_options_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageOption_DDL_Interleave.ProtoReflect.Descriptor instead.
func (*MessageOption_DDL_Interleave) Descriptor() ([]byte, []int) {
	return file_server_options_transaction_options_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *MessageOption_DDL_Interleave) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

type MessageOption_DDL_Index_Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Column string `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`
	Desc   bool   `protobuf:"varint,2,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *MessageOption_DDL_Index_Key) Reset() {
	*x = MessageOption_DDL_Index_Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_options_transaction_options_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageOption_DDL_Index_Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOption_DDL_Index_Key) ProtoMessage() {}

func (x *MessageOption_DDL_Index_Key) ProtoReflect() protoreflect.Message {
	mi := &file_server_options_transaction_options_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageOption_DDL_Index_Key.ProtoReflect.Descriptor instead.
func (*MessageOption_DDL_Index_Key) Descriptor() ([]byte, []int) {
	return file_server_options_transaction_options_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

func (x *MessageOption_DDL_Index_Key) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

func (x *MessageOption_DDL_Index_Key) GetDesc() bool {
	if x != nil {
		return x.Desc
	}
	return false
}

type FieldOption_DDL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pk        bool                       `protobuf:"varint,1,opt,name=pk,proto3" json:"pk,omitempty"`
	MasterRef *FieldOption_DDL_MasterRef `protobuf:"bytes,3,opt,name=master_ref,json=masterRef,proto3" json:"master_ref,omitempty"`
}

func (x *FieldOption_DDL) Reset() {
	*x = FieldOption_DDL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_options_transaction_options_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldOption_DDL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldOption_DDL) ProtoMessage() {}

func (x *FieldOption_DDL) ProtoReflect() protoreflect.Message {
	mi := &file_server_options_transaction_options_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldOption_DDL.ProtoReflect.Descriptor instead.
func (*FieldOption_DDL) Descriptor() ([]byte, []int) {
	return file_server_options_transaction_options_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FieldOption_DDL) GetPk() bool {
	if x != nil {
		return x.Pk
	}
	return false
}

func (x *FieldOption_DDL) GetMasterRef() *FieldOption_DDL_MasterRef {
	if x != nil {
		return x.MasterRef
	}
	return nil
}

type FieldOption_DDL_MasterRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table  string `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Column string `protobuf:"bytes,2,opt,name=column,proto3" json:"column,omitempty"`
}

func (x *FieldOption_DDL_MasterRef) Reset() {
	*x = FieldOption_DDL_MasterRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_options_transaction_options_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldOption_DDL_MasterRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldOption_DDL_MasterRef) ProtoMessage() {}

func (x *FieldOption_DDL_MasterRef) ProtoReflect() protoreflect.Message {
	mi := &file_server_options_transaction_options_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldOption_DDL_MasterRef.ProtoReflect.Descriptor instead.
func (*FieldOption_DDL_MasterRef) Descriptor() ([]byte, []int) {
	return file_server_options_transaction_options_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *FieldOption_DDL_MasterRef) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *FieldOption_DDL_MasterRef) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

var file_server_options_transaction_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*MessageOption)(nil),
		Field:         53001,
		Name:          "server.options.transaction.message",
		Tag:           "bytes,53001,opt,name=message",
		Filename:      "server/options/transaction/options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldOption)(nil),
		Field:         53002,
		Name:          "server.options.transaction.field",
		Tag:           "bytes,53002,opt,name=field",
		Filename:      "server/options/transaction/options.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional server.options.transaction.MessageOption message = 53001;
	E_Message = &file_server_options_transaction_options_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional server.options.transaction.FieldOption field = 53002;
	E_Field = &file_server_options_transaction_options_proto_extTypes[1]
)

var File_server_options_transaction_options_proto protoreflect.FileDescriptor

var file_server_options_transaction_options_proto_rawDesc = []byte{
	0x0a, 0x28, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x05, 0x0a, 0x0d, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5b, 0x0a, 0x0d, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3f, 0x0a, 0x03, 0x64, 0x64, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x44, 0x44, 0x4c, 0x52, 0x03, 0x64, 0x64, 0x6c, 0x1a, 0x8e, 0x03, 0x0a, 0x03, 0x44, 0x44, 0x4c,
	0x12, 0x4d, 0x0a, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x33, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x44, 0x4c,
	0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x12,
	0x58, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44,
	0x44, 0x4c, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x0a, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x1a, 0xb9, 0x01, 0x0a, 0x05, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x4b, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x44, 0x4c,
	0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x1a, 0x31, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x1a, 0x22, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6c, 0x65,
	0x61, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x40, 0x0a, 0x0c, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x6e, 0x6c, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x41, 0x6e, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x10, 0x02, 0x22, 0x89, 0x03, 0x0a, 0x0b,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x59, 0x0a, 0x0d, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x34, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x03, 0x64, 0x64, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x44, 0x4c,
	0x52, 0x03, 0x64, 0x64, 0x6c, 0x1a, 0xa6, 0x01, 0x0a, 0x03, 0x44, 0x44, 0x4c, 0x12, 0x0e, 0x0a,
	0x02, 0x70, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x70, 0x6b, 0x12, 0x54, 0x0a,
	0x0a, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x35, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x44, 0x4c, 0x2e, 0x4d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x66, 0x52, 0x09, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x66, 0x1a, 0x39, 0x0a, 0x09, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x66,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x22, 0x37,
	0x0a, 0x0c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07,
	0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x6e, 0x6c, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x6e, 0x6c, 0x79, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x10, 0x02, 0x3a, 0x66, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x89, 0x9e, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3a,
	0x5e, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8a, 0x9e, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42,
	0x61, 0x5a, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69,
	0x6b, 0x79, 0x61, 0x72, 0x75, 0x2d, 0x73, 0x75, 0x7a, 0x75, 0x6b, 0x69, 0x2f, 0x67, 0x6f, 0x2d,
	0x63, 0x6f, 0x6e, 0x2d, 0x32, 0x30, 0x32, 0x32, 0x2d, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x2d,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_options_transaction_options_proto_rawDescOnce sync.Once
	file_server_options_transaction_options_proto_rawDescData = file_server_options_transaction_options_proto_rawDesc
)

func file_server_options_transaction_options_proto_rawDescGZIP() []byte {
	file_server_options_transaction_options_proto_rawDescOnce.Do(func() {
		file_server_options_transaction_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_options_transaction_options_proto_rawDescData)
	})
	return file_server_options_transaction_options_proto_rawDescData
}

var file_server_options_transaction_options_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_server_options_transaction_options_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_server_options_transaction_options_proto_goTypes = []interface{}{
	(MessageOption_AccessorType)(0),      // 0: server.options.transaction.MessageOption.AccessorType
	(FieldOption_AccessorType)(0),        // 1: server.options.transaction.FieldOption.AccessorType
	(*MessageOption)(nil),                // 2: server.options.transaction.MessageOption
	(*FieldOption)(nil),                  // 3: server.options.transaction.FieldOption
	(*MessageOption_DDL)(nil),            // 4: server.options.transaction.MessageOption.DDL
	(*MessageOption_DDL_Index)(nil),      // 5: server.options.transaction.MessageOption.DDL.Index
	(*MessageOption_DDL_Interleave)(nil), // 6: server.options.transaction.MessageOption.DDL.Interleave
	(*MessageOption_DDL_Index_Key)(nil),  // 7: server.options.transaction.MessageOption.DDL.Index.Key
	(*FieldOption_DDL)(nil),              // 8: server.options.transaction.FieldOption.DDL
	(*FieldOption_DDL_MasterRef)(nil),    // 9: server.options.transaction.FieldOption.DDL.MasterRef
	(*descriptorpb.MessageOptions)(nil),  // 10: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),    // 11: google.protobuf.FieldOptions
}
var file_server_options_transaction_options_proto_depIdxs = []int32{
	0,  // 0: server.options.transaction.MessageOption.accessor_type:type_name -> server.options.transaction.MessageOption.AccessorType
	4,  // 1: server.options.transaction.MessageOption.ddl:type_name -> server.options.transaction.MessageOption.DDL
	1,  // 2: server.options.transaction.FieldOption.accessor_type:type_name -> server.options.transaction.FieldOption.AccessorType
	8,  // 3: server.options.transaction.FieldOption.ddl:type_name -> server.options.transaction.FieldOption.DDL
	5,  // 4: server.options.transaction.MessageOption.DDL.indexes:type_name -> server.options.transaction.MessageOption.DDL.Index
	6,  // 5: server.options.transaction.MessageOption.DDL.interleave:type_name -> server.options.transaction.MessageOption.DDL.Interleave
	7,  // 6: server.options.transaction.MessageOption.DDL.Index.keys:type_name -> server.options.transaction.MessageOption.DDL.Index.Key
	9,  // 7: server.options.transaction.FieldOption.DDL.master_ref:type_name -> server.options.transaction.FieldOption.DDL.MasterRef
	10, // 8: server.options.transaction.message:extendee -> google.protobuf.MessageOptions
	11, // 9: server.options.transaction.field:extendee -> google.protobuf.FieldOptions
	2,  // 10: server.options.transaction.message:type_name -> server.options.transaction.MessageOption
	3,  // 11: server.options.transaction.field:type_name -> server.options.transaction.FieldOption
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	10, // [10:12] is the sub-list for extension type_name
	8,  // [8:10] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_server_options_transaction_options_proto_init() }
func file_server_options_transaction_options_proto_init() {
	if File_server_options_transaction_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_options_transaction_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageOption); i {
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
		file_server_options_transaction_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldOption); i {
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
		file_server_options_transaction_options_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageOption_DDL); i {
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
		file_server_options_transaction_options_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageOption_DDL_Index); i {
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
		file_server_options_transaction_options_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageOption_DDL_Interleave); i {
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
		file_server_options_transaction_options_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageOption_DDL_Index_Key); i {
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
		file_server_options_transaction_options_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldOption_DDL); i {
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
		file_server_options_transaction_options_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldOption_DDL_MasterRef); i {
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
			RawDescriptor: file_server_options_transaction_options_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_server_options_transaction_options_proto_goTypes,
		DependencyIndexes: file_server_options_transaction_options_proto_depIdxs,
		EnumInfos:         file_server_options_transaction_options_proto_enumTypes,
		MessageInfos:      file_server_options_transaction_options_proto_msgTypes,
		ExtensionInfos:    file_server_options_transaction_options_proto_extTypes,
	}.Build()
	File_server_options_transaction_options_proto = out.File
	file_server_options_transaction_options_proto_rawDesc = nil
	file_server_options_transaction_options_proto_goTypes = nil
	file_server_options_transaction_options_proto_depIdxs = nil
}
