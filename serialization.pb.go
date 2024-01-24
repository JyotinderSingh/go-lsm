// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.23.3
// source: serialization.proto

package golsm

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

type Command int32

const (
	Command_PUT    Command = 0
	Command_DELETE Command = 1
)

// Enum value maps for Command.
var (
	Command_name = map[int32]string{
		0: "PUT",
		1: "DELETE",
	}
	Command_value = map[string]int32{
		"PUT":    0,
		"DELETE": 1,
	}
)

func (x Command) Enum() *Command {
	p := new(Command)
	*p = x
	return p
}

func (x Command) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Command) Descriptor() protoreflect.EnumDescriptor {
	return file_serialization_proto_enumTypes[0].Descriptor()
}

func (Command) Type() protoreflect.EnumType {
	return &file_serialization_proto_enumTypes[0]
}

func (x Command) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Command.Descriptor instead.
func (Command) EnumDescriptor() ([]byte, []int) {
	return file_serialization_proto_rawDescGZIP(), []int{0}
}

type SSTableEntryType int32

const (
	SSTableEntryType_BLOOM_FILTER   SSTableEntryType = 0
	SSTableEntryType_INDEX          SSTableEntryType = 1
	SSTableEntryType_MEMTABLE_ENTRY SSTableEntryType = 2
)

// Enum value maps for SSTableEntryType.
var (
	SSTableEntryType_name = map[int32]string{
		0: "BLOOM_FILTER",
		1: "INDEX",
		2: "MEMTABLE_ENTRY",
	}
	SSTableEntryType_value = map[string]int32{
		"BLOOM_FILTER":   0,
		"INDEX":          1,
		"MEMTABLE_ENTRY": 2,
	}
)

func (x SSTableEntryType) Enum() *SSTableEntryType {
	p := new(SSTableEntryType)
	*p = x
	return p
}

func (x SSTableEntryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SSTableEntryType) Descriptor() protoreflect.EnumDescriptor {
	return file_serialization_proto_enumTypes[1].Descriptor()
}

func (SSTableEntryType) Type() protoreflect.EnumType {
	return &file_serialization_proto_enumTypes[1]
}

func (x SSTableEntryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SSTableEntryType.Descriptor instead.
func (SSTableEntryType) EnumDescriptor() ([]byte, []int) {
	return file_serialization_proto_rawDescGZIP(), []int{1}
}

type MemtableEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command Command `protobuf:"varint,1,opt,name=command,proto3,enum=grpcapi.Command" json:"command,omitempty"`
	Value   []byte  `protobuf:"bytes,2,opt,name=value,proto3,oneof" json:"value,omitempty"`
}

func (x *MemtableEntry) Reset() {
	*x = MemtableEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serialization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemtableEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemtableEntry) ProtoMessage() {}

func (x *MemtableEntry) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemtableEntry.ProtoReflect.Descriptor instead.
func (*MemtableEntry) Descriptor() ([]byte, []int) {
	return file_serialization_proto_rawDescGZIP(), []int{0}
}

func (x *MemtableEntry) GetCommand() Command {
	if x != nil {
		return x.Command
	}
	return Command_PUT
}

func (x *MemtableEntry) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type SSTableEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          SSTableEntryType  `protobuf:"varint,1,opt,name=type,proto3,enum=grpcapi.SSTableEntryType" json:"type,omitempty"`
	BloomFilter   *BloomFilter      `protobuf:"bytes,2,opt,name=bloom_filter,json=bloomFilter,proto3,oneof" json:"bloom_filter,omitempty"`
	Index         *Index            `protobuf:"bytes,3,opt,name=index,proto3,oneof" json:"index,omitempty"`
	MemtableEntry *MemtableKeyValue `protobuf:"bytes,4,opt,name=memtable_entry,json=memtableEntry,proto3,oneof" json:"memtable_entry,omitempty"`
}

func (x *SSTableEntry) Reset() {
	*x = SSTableEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serialization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSTableEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSTableEntry) ProtoMessage() {}

func (x *SSTableEntry) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSTableEntry.ProtoReflect.Descriptor instead.
func (*SSTableEntry) Descriptor() ([]byte, []int) {
	return file_serialization_proto_rawDescGZIP(), []int{1}
}

func (x *SSTableEntry) GetType() SSTableEntryType {
	if x != nil {
		return x.Type
	}
	return SSTableEntryType_BLOOM_FILTER
}

func (x *SSTableEntry) GetBloomFilter() *BloomFilter {
	if x != nil {
		return x.BloomFilter
	}
	return nil
}

func (x *SSTableEntry) GetIndex() *Index {
	if x != nil {
		return x.Index
	}
	return nil
}

func (x *SSTableEntry) GetMemtableEntry() *MemtableKeyValue {
	if x != nil {
		return x.MemtableEntry
	}
	return nil
}

type BloomFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hashes []uint32 `protobuf:"varint,1,rep,packed,name=hashes,proto3" json:"hashes,omitempty"`
}

func (x *BloomFilter) Reset() {
	*x = BloomFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serialization_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BloomFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BloomFilter) ProtoMessage() {}

func (x *BloomFilter) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BloomFilter.ProtoReflect.Descriptor instead.
func (*BloomFilter) Descriptor() ([]byte, []int) {
	return file_serialization_proto_rawDescGZIP(), []int{2}
}

func (x *BloomFilter) GetHashes() []uint32 {
	if x != nil {
		return x.Hashes
	}
	return nil
}

type IndexEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Offset uint32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *IndexEntry) Reset() {
	*x = IndexEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serialization_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexEntry) ProtoMessage() {}

func (x *IndexEntry) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexEntry.ProtoReflect.Descriptor instead.
func (*IndexEntry) Descriptor() ([]byte, []int) {
	return file_serialization_proto_rawDescGZIP(), []int{3}
}

func (x *IndexEntry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *IndexEntry) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Index struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*IndexEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *Index) Reset() {
	*x = Index{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serialization_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Index) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Index) ProtoMessage() {}

func (x *Index) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Index.ProtoReflect.Descriptor instead.
func (*Index) Descriptor() ([]byte, []int) {
	return file_serialization_proto_rawDescGZIP(), []int{4}
}

func (x *Index) GetEntries() []*IndexEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type MemtableKeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string         `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *MemtableEntry `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MemtableKeyValue) Reset() {
	*x = MemtableKeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serialization_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemtableKeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemtableKeyValue) ProtoMessage() {}

func (x *MemtableKeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemtableKeyValue.ProtoReflect.Descriptor instead.
func (*MemtableKeyValue) Descriptor() ([]byte, []int) {
	return file_serialization_proto_rawDescGZIP(), []int{5}
}

func (x *MemtableKeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MemtableKeyValue) GetValue() *MemtableEntry {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_serialization_proto protoreflect.FileDescriptor

var file_serialization_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x22, 0x60,
	0x0a, 0x0d, 0x4d, 0x65, 0x6d, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x2a, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x9b, 0x02, 0x0a, 0x0c, 0x53, 0x53, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x53, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x3c, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69,
	0x2e, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0b,
	0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x29,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x48, 0x01, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x0e, 0x6d, 0x65, 0x6d,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x6d, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x02, 0x52, 0x0d,
	0x6d, 0x65, 0x6d, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x88, 0x01, 0x01,
	0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x11, 0x0a, 0x0f, 0x5f,
	0x6d, 0x65, 0x6d, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x25,
	0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x6f, 0x6d, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x68,
	0x61, 0x73, 0x68, 0x65, 0x73, 0x22, 0x36, 0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x36, 0x0a,
	0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2d, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70,
	0x69, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x52, 0x0a, 0x10, 0x4d, 0x65, 0x6d, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x6d, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x1e, 0x0a, 0x07, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x01, 0x2a, 0x43, 0x0a, 0x10, 0x53, 0x53, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a,
	0x0c, 0x42, 0x4c, 0x4f, 0x4f, 0x4d, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x45,
	0x4d, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x02, 0x42, 0x21,
	0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x79, 0x6f,
	0x74, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x68, 0x2f, 0x67, 0x6f, 0x6c, 0x73,
	0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_serialization_proto_rawDescOnce sync.Once
	file_serialization_proto_rawDescData = file_serialization_proto_rawDesc
)

func file_serialization_proto_rawDescGZIP() []byte {
	file_serialization_proto_rawDescOnce.Do(func() {
		file_serialization_proto_rawDescData = protoimpl.X.CompressGZIP(file_serialization_proto_rawDescData)
	})
	return file_serialization_proto_rawDescData
}

var file_serialization_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_serialization_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_serialization_proto_goTypes = []interface{}{
	(Command)(0),             // 0: grpcapi.Command
	(SSTableEntryType)(0),    // 1: grpcapi.SSTableEntryType
	(*MemtableEntry)(nil),    // 2: grpcapi.MemtableEntry
	(*SSTableEntry)(nil),     // 3: grpcapi.SSTableEntry
	(*BloomFilter)(nil),      // 4: grpcapi.BloomFilter
	(*IndexEntry)(nil),       // 5: grpcapi.IndexEntry
	(*Index)(nil),            // 6: grpcapi.Index
	(*MemtableKeyValue)(nil), // 7: grpcapi.MemtableKeyValue
}
var file_serialization_proto_depIdxs = []int32{
	0, // 0: grpcapi.MemtableEntry.command:type_name -> grpcapi.Command
	1, // 1: grpcapi.SSTableEntry.type:type_name -> grpcapi.SSTableEntryType
	4, // 2: grpcapi.SSTableEntry.bloom_filter:type_name -> grpcapi.BloomFilter
	6, // 3: grpcapi.SSTableEntry.index:type_name -> grpcapi.Index
	7, // 4: grpcapi.SSTableEntry.memtable_entry:type_name -> grpcapi.MemtableKeyValue
	5, // 5: grpcapi.Index.entries:type_name -> grpcapi.IndexEntry
	2, // 6: grpcapi.MemtableKeyValue.value:type_name -> grpcapi.MemtableEntry
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_serialization_proto_init() }
func file_serialization_proto_init() {
	if File_serialization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_serialization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemtableEntry); i {
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
		file_serialization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSTableEntry); i {
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
		file_serialization_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BloomFilter); i {
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
		file_serialization_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexEntry); i {
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
		file_serialization_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Index); i {
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
		file_serialization_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemtableKeyValue); i {
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
	file_serialization_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_serialization_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_serialization_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_serialization_proto_goTypes,
		DependencyIndexes: file_serialization_proto_depIdxs,
		EnumInfos:         file_serialization_proto_enumTypes,
		MessageInfos:      file_serialization_proto_msgTypes,
	}.Build()
	File_serialization_proto = out.File
	file_serialization_proto_rawDesc = nil
	file_serialization_proto_goTypes = nil
	file_serialization_proto_depIdxs = nil
}
