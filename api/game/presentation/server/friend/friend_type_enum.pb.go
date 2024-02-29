// フレンドタイプ

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: friend_type_enum.proto

package friend

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FriendType int32

const (
	FriendType_Applying    FriendType = 0
	FriendType_NotApproved FriendType = 1
	FriendType_Approved    FriendType = 2
	FriendType_Disapproved FriendType = 3
)

// Enum value maps for FriendType.
var (
	FriendType_name = map[int32]string{
		0: "Applying",
		1: "NotApproved",
		2: "Approved",
		3: "Disapproved",
	}
	FriendType_value = map[string]int32{
		"Applying":    0,
		"NotApproved": 1,
		"Approved":    2,
		"Disapproved": 3,
	}
)

func (x FriendType) Enum() *FriendType {
	p := new(FriendType)
	*p = x
	return p
}

func (x FriendType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FriendType) Descriptor() protoreflect.EnumDescriptor {
	return file_friend_type_enum_proto_enumTypes[0].Descriptor()
}

func (FriendType) Type() protoreflect.EnumType {
	return &file_friend_type_enum_proto_enumTypes[0]
}

func (x FriendType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FriendType.Descriptor instead.
func (FriendType) EnumDescriptor() ([]byte, []int) {
	return file_friend_type_enum_proto_rawDescGZIP(), []int{0}
}

var File_friend_type_enum_proto protoreflect.FileDescriptor

var file_friend_type_enum_proto_rawDesc = []byte{
	0x0a, 0x16, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x65, 0x6e,
	0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a,
	0x4a, 0x0a, 0x0a, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a,
	0x08, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4e,
	0x6f, 0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x69,
	0x73, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x10, 0x03, 0x42, 0x25, 0x5a, 0x23, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_friend_type_enum_proto_rawDescOnce sync.Once
	file_friend_type_enum_proto_rawDescData = file_friend_type_enum_proto_rawDesc
)

func file_friend_type_enum_proto_rawDescGZIP() []byte {
	file_friend_type_enum_proto_rawDescOnce.Do(func() {
		file_friend_type_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_friend_type_enum_proto_rawDescData)
	})
	return file_friend_type_enum_proto_rawDescData
}

var file_friend_type_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_friend_type_enum_proto_goTypes = []interface{}{
	(FriendType)(0), // 0: proto.FriendType
}
var file_friend_type_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_friend_type_enum_proto_init() }
func file_friend_type_enum_proto_init() {
	if File_friend_type_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_friend_type_enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_friend_type_enum_proto_goTypes,
		DependencyIndexes: file_friend_type_enum_proto_depIdxs,
		EnumInfos:         file_friend_type_enum_proto_enumTypes,
	}.Build()
	File_friend_type_enum_proto = out.File
	file_friend_type_enum_proto_rawDesc = nil
	file_friend_type_enum_proto_goTypes = nil
	file_friend_type_enum_proto_depIdxs = nil
}
