// ルームユーザー

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: common_room_user_structure.proto

package room

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

type CommonRoomUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId               string               `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	UserId               string               `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoomUserPositionType RoomUserPositionType `protobuf:"varint,3,opt,name=room_user_position_type,json=roomUserPositionType,proto3,enum=proto.RoomUserPositionType" json:"room_user_position_type,omitempty"`
}

func (x *CommonRoomUser) Reset() {
	*x = CommonRoomUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_room_user_structure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonRoomUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonRoomUser) ProtoMessage() {}

func (x *CommonRoomUser) ProtoReflect() protoreflect.Message {
	mi := &file_common_room_user_structure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonRoomUser.ProtoReflect.Descriptor instead.
func (*CommonRoomUser) Descriptor() ([]byte, []int) {
	return file_common_room_user_structure_proto_rawDescGZIP(), []int{0}
}

func (x *CommonRoomUser) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *CommonRoomUser) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CommonRoomUser) GetRoomUserPositionType() RoomUserPositionType {
	if x != nil {
		return x.RoomUserPositionType
	}
	return RoomUserPositionType_Leader
}

var File_common_room_user_structure_proto protoreflect.FileDescriptor

var file_common_room_user_structure_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01,
	0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x52, 0x0a, 0x17, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6f, 0x6d,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x14, 0x72, 0x6f, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x23, 0x5a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61,
	0x6d, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_common_room_user_structure_proto_rawDescOnce sync.Once
	file_common_room_user_structure_proto_rawDescData = file_common_room_user_structure_proto_rawDesc
)

func file_common_room_user_structure_proto_rawDescGZIP() []byte {
	file_common_room_user_structure_proto_rawDescOnce.Do(func() {
		file_common_room_user_structure_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_room_user_structure_proto_rawDescData)
	})
	return file_common_room_user_structure_proto_rawDescData
}

var file_common_room_user_structure_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_room_user_structure_proto_goTypes = []interface{}{
	(*CommonRoomUser)(nil),    // 0: proto.CommonRoomUser
	(RoomUserPositionType)(0), // 1: proto.RoomUserPositionType
}
var file_common_room_user_structure_proto_depIdxs = []int32{
	1, // 0: proto.CommonRoomUser.room_user_position_type:type_name -> proto.RoomUserPositionType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_room_user_structure_proto_init() }
func file_common_room_user_structure_proto_init() {
	if File_common_room_user_structure_proto != nil {
		return
	}
	file_room_user_position_type_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_room_user_structure_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonRoomUser); i {
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
			RawDescriptor: file_common_room_user_structure_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_room_user_structure_proto_goTypes,
		DependencyIndexes: file_common_room_user_structure_proto_depIdxs,
		MessageInfos:      file_common_room_user_structure_proto_msgTypes,
	}.Build()
	File_common_room_user_structure_proto = out.File
	file_common_room_user_structure_proto_rawDesc = nil
	file_common_room_user_structure_proto_goTypes = nil
	file_common_room_user_structure_proto_depIdxs = nil
}
