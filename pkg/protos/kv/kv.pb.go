// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.2
// source: kv/kv.proto

// this is for android multiple process access bboltdb only

package kv

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Element struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Buckets     []string               `protobuf:"bytes,1,rep,name=buckets"`
	xxx_hidden_Key         []byte                 `protobuf:"bytes,2,opt,name=key"`
	xxx_hidden_Value       []byte                 `protobuf:"bytes,3,opt,name=value"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Element) Reset() {
	*x = Element{}
	mi := &file_kv_kv_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Element) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Element) ProtoMessage() {}

func (x *Element) ProtoReflect() protoreflect.Message {
	mi := &file_kv_kv_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Element) GetBuckets() []string {
	if x != nil {
		return x.xxx_hidden_Buckets
	}
	return nil
}

func (x *Element) GetKey() []byte {
	if x != nil {
		return x.xxx_hidden_Key
	}
	return nil
}

func (x *Element) GetValue() []byte {
	if x != nil {
		return x.xxx_hidden_Value
	}
	return nil
}

func (x *Element) SetBuckets(v []string) {
	x.xxx_hidden_Buckets = v
}

func (x *Element) SetKey(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Key = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 3)
}

func (x *Element) SetValue(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Value = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 3)
}

func (x *Element) HasKey() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *Element) HasValue() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *Element) ClearKey() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Key = nil
}

func (x *Element) ClearValue() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_Value = nil
}

type Element_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Buckets []string
	Key     []byte
	Value   []byte
}

func (b0 Element_builder) Build() *Element {
	m0 := &Element{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Buckets = b.Buckets
	if b.Key != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 3)
		x.xxx_hidden_Key = b.Key
	}
	if b.Value != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 3)
		x.xxx_hidden_Value = b.Value
	}
	return m0
}

type Keys struct {
	state              protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Buckets []string               `protobuf:"bytes,1,rep,name=buckets"`
	xxx_hidden_Keys    [][]byte               `protobuf:"bytes,2,rep,name=keys"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *Keys) Reset() {
	*x = Keys{}
	mi := &file_kv_kv_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Keys) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keys) ProtoMessage() {}

func (x *Keys) ProtoReflect() protoreflect.Message {
	mi := &file_kv_kv_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Keys) GetBuckets() []string {
	if x != nil {
		return x.xxx_hidden_Buckets
	}
	return nil
}

func (x *Keys) GetKeys() [][]byte {
	if x != nil {
		return x.xxx_hidden_Keys
	}
	return nil
}

func (x *Keys) SetBuckets(v []string) {
	x.xxx_hidden_Buckets = v
}

func (x *Keys) SetKeys(v [][]byte) {
	x.xxx_hidden_Keys = v
}

type Keys_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Buckets []string
	Keys    [][]byte
}

func (b0 Keys_builder) Build() *Keys {
	m0 := &Keys{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Buckets = b.Buckets
	x.xxx_hidden_Keys = b.Keys
	return m0
}

var File_kv_kv_proto protoreflect.FileDescriptor

var file_kv_kv_proto_rawDesc = string([]byte{
	0x0a, 0x0b, 0x6b, 0x76, 0x2f, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x79,
	0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x5f,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b,
	0x0a, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x34, 0x0a, 0x04, 0x4b,
	0x65, 0x79, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x6b, 0x65, 0x79,
	0x73, 0x32, 0xad, 0x02, 0x0a, 0x07, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x39, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x6b,
	0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x18,
	0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12,
	0x18, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x37, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x79, 0x75,
	0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4b, 0x65,
	0x79, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x05, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x18, 0x2e, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x6b, 0x76,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x18, 0x2e,
	0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x6b, 0x76, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x30, 0x01, 0x12, 0x36, 0x0a, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x34, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x41, 0x73, 0x75, 0x74, 0x6f, 0x72, 0x75, 0x66, 0x61, 0x2f, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69,
	0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6b, 0x76, 0x92,
	0x03, 0x05, 0xd2, 0x3e, 0x02, 0x10, 0x03, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x70, 0xe8, 0x07,
})

var file_kv_kv_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kv_kv_proto_goTypes = []any{
	(*Element)(nil),       // 0: yuhaiin.kvstore.element
	(*Keys)(nil),          // 1: yuhaiin.kvstore.Keys
	(*emptypb.Empty)(nil), // 2: google.protobuf.Empty
}
var file_kv_kv_proto_depIdxs = []int32{
	0, // 0: yuhaiin.kvstore.kvstore.Get:input_type -> yuhaiin.kvstore.element
	0, // 1: yuhaiin.kvstore.kvstore.Set:input_type -> yuhaiin.kvstore.element
	1, // 2: yuhaiin.kvstore.kvstore.Delete:input_type -> yuhaiin.kvstore.Keys
	0, // 3: yuhaiin.kvstore.kvstore.Range:input_type -> yuhaiin.kvstore.element
	2, // 4: yuhaiin.kvstore.kvstore.Ping:input_type -> google.protobuf.Empty
	0, // 5: yuhaiin.kvstore.kvstore.Get:output_type -> yuhaiin.kvstore.element
	2, // 6: yuhaiin.kvstore.kvstore.Set:output_type -> google.protobuf.Empty
	2, // 7: yuhaiin.kvstore.kvstore.Delete:output_type -> google.protobuf.Empty
	0, // 8: yuhaiin.kvstore.kvstore.Range:output_type -> yuhaiin.kvstore.element
	2, // 9: yuhaiin.kvstore.kvstore.Ping:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kv_kv_proto_init() }
func file_kv_kv_proto_init() {
	if File_kv_kv_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_kv_kv_proto_rawDesc), len(file_kv_kv_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kv_kv_proto_goTypes,
		DependencyIndexes: file_kv_kv_proto_depIdxs,
		MessageInfos:      file_kv_kv_proto_msgTypes,
	}.Build()
	File_kv_kv_proto = out.File
	file_kv_kv_proto_goTypes = nil
	file_kv_kv_proto_depIdxs = nil
}
