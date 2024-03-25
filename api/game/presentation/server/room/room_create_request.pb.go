// ルーム作成リクエスト

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.24.4
// source: room_create_request.proto

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

type RoomCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string          `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name            string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RoomReleaseType RoomReleaseType `protobuf:"varint,3,opt,name=room_release_type,json=roomReleaseType,proto3,enum=proto.RoomReleaseType" json:"room_release_type,omitempty"`
}

func (x *RoomCreateRequest) Reset() {
	*x = RoomCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_create_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomCreateRequest) ProtoMessage() {}

func (x *RoomCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_room_create_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomCreateRequest.ProtoReflect.Descriptor instead.
func (*RoomCreateRequest) Descriptor() ([]byte, []int) {
	return file_room_create_request_proto_rawDescGZIP(), []int{0}
}

func (x *RoomCreateRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RoomCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoomCreateRequest) GetRoomReleaseType() RoomReleaseType {
	if x != nil {
		return x.RoomReleaseType
	}
	return RoomReleaseType_Private
}

var File_room_create_request_proto protoreflect.FileDescriptor

var file_room_create_request_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x84, 0x01, 0x0a, 0x11, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x11, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x72, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x23, 0x5a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_room_create_request_proto_rawDescOnce sync.Once
	file_room_create_request_proto_rawDescData = file_room_create_request_proto_rawDesc
)

func file_room_create_request_proto_rawDescGZIP() []byte {
	file_room_create_request_proto_rawDescOnce.Do(func() {
		file_room_create_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_room_create_request_proto_rawDescData)
	})
	return file_room_create_request_proto_rawDescData
}

var file_room_create_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_room_create_request_proto_goTypes = []interface{}{
	(*RoomCreateRequest)(nil), // 0: proto.RoomCreateRequest
	(RoomReleaseType)(0),      // 1: proto.RoomReleaseType
}
var file_room_create_request_proto_depIdxs = []int32{
	1, // 0: proto.RoomCreateRequest.room_release_type:type_name -> proto.RoomReleaseType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_room_create_request_proto_init() }
func file_room_create_request_proto_init() {
	if File_room_create_request_proto != nil {
		return
	}
	file_room_release_type_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_room_create_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomCreateRequest); i {
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
			RawDescriptor: file_room_create_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_room_create_request_proto_goTypes,
		DependencyIndexes: file_room_create_request_proto_depIdxs,
		MessageInfos:      file_room_create_request_proto_msgTypes,
	}.Build()
	File_room_create_request_proto = out.File
	file_room_create_request_proto_rawDesc = nil
	file_room_create_request_proto_goTypes = nil
	file_room_create_request_proto_depIdxs = nil
}