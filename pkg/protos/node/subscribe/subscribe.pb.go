// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.1
// source: node/subscribe/subscribe.proto

package subscribe

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

type Type int32

const (
	Type_reserve      Type = 0
	Type_trojan       Type = 1
	Type_vmess        Type = 2
	Type_shadowsocks  Type = 3
	Type_shadowsocksr Type = 4
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "reserve",
		1: "trojan",
		2: "vmess",
		3: "shadowsocks",
		4: "shadowsocksr",
	}
	Type_value = map[string]int32{
		"reserve":      0,
		"trojan":       1,
		"vmess":        2,
		"shadowsocks":  3,
		"shadowsocksr": 4,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_node_subscribe_subscribe_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_node_subscribe_subscribe_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_node_subscribe_subscribe_proto_rawDescGZIP(), []int{0}
}

type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type Type   `protobuf:"varint,2,opt,name=type,proto3,enum=yuhaiin.subscribe.Type" json:"type,omitempty"`
	Url  string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_subscribe_subscribe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_node_subscribe_subscribe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_node_subscribe_subscribe_proto_rawDescGZIP(), []int{0}
}

func (x *Link) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Link) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_reserve
}

func (x *Link) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_node_subscribe_subscribe_proto protoreflect.FileDescriptor

var file_node_subscribe_subscribe_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x22, 0x59, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e,
	0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x2a, 0x4d,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x74, 0x72, 0x6f, 0x6a, 0x61, 0x6e, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x76, 0x6d, 0x65, 0x73, 0x73, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x73, 0x68,
	0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x73,
	0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x72, 0x10, 0x04, 0x42, 0x38, 0x5a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x75, 0x74,
	0x6f, 0x72, 0x75, 0x66, 0x61, 0x2f, 0x79, 0x75, 0x68, 0x61, 0x69, 0x69, 0x6e, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_subscribe_subscribe_proto_rawDescOnce sync.Once
	file_node_subscribe_subscribe_proto_rawDescData = file_node_subscribe_subscribe_proto_rawDesc
)

func file_node_subscribe_subscribe_proto_rawDescGZIP() []byte {
	file_node_subscribe_subscribe_proto_rawDescOnce.Do(func() {
		file_node_subscribe_subscribe_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_subscribe_subscribe_proto_rawDescData)
	})
	return file_node_subscribe_subscribe_proto_rawDescData
}

var file_node_subscribe_subscribe_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_node_subscribe_subscribe_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_node_subscribe_subscribe_proto_goTypes = []interface{}{
	(Type)(0),    // 0: yuhaiin.subscribe.type
	(*Link)(nil), // 1: yuhaiin.subscribe.link
}
var file_node_subscribe_subscribe_proto_depIdxs = []int32{
	0, // 0: yuhaiin.subscribe.link.type:type_name -> yuhaiin.subscribe.type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_node_subscribe_subscribe_proto_init() }
func file_node_subscribe_subscribe_proto_init() {
	if File_node_subscribe_subscribe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_subscribe_subscribe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
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
			RawDescriptor: file_node_subscribe_subscribe_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_node_subscribe_subscribe_proto_goTypes,
		DependencyIndexes: file_node_subscribe_subscribe_proto_depIdxs,
		EnumInfos:         file_node_subscribe_subscribe_proto_enumTypes,
		MessageInfos:      file_node_subscribe_subscribe_proto_msgTypes,
	}.Build()
	File_node_subscribe_subscribe_proto = out.File
	file_node_subscribe_subscribe_proto_rawDesc = nil
	file_node_subscribe_subscribe_proto_goTypes = nil
	file_node_subscribe_subscribe_proto_depIdxs = nil
}
