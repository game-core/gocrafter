// フレンド取得レスポンス

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: friend_get_response.proto

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

type FriendGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserFriends []*UserFriend `protobuf:"bytes,1,rep,name=user_friends,json=userFriends,proto3" json:"user_friends,omitempty"`
}

func (x *FriendGetResponse) Reset() {
	*x = FriendGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_get_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendGetResponse) ProtoMessage() {}

func (x *FriendGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_friend_get_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendGetResponse.ProtoReflect.Descriptor instead.
func (*FriendGetResponse) Descriptor() ([]byte, []int) {
	return file_friend_get_response_proto_rawDescGZIP(), []int{0}
}

func (x *FriendGetResponse) GetUserFriends() []*UserFriend {
	if x != nil {
		return x.UserFriends
	}
	return nil
}

var File_friend_get_response_proto protoreflect.FileDescriptor

var file_friend_get_response_proto_rawDesc = []byte{
	0x0a, 0x19, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x49, 0x0a, 0x11, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x0b, 0x75,
	0x73, 0x65, 0x72, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x42, 0x25, 0x5a, 0x23, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_friend_get_response_proto_rawDescOnce sync.Once
	file_friend_get_response_proto_rawDescData = file_friend_get_response_proto_rawDesc
)

func file_friend_get_response_proto_rawDescGZIP() []byte {
	file_friend_get_response_proto_rawDescOnce.Do(func() {
		file_friend_get_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_friend_get_response_proto_rawDescData)
	})
	return file_friend_get_response_proto_rawDescData
}

var file_friend_get_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_friend_get_response_proto_goTypes = []interface{}{
	(*FriendGetResponse)(nil), // 0: proto.FriendGetResponse
	(*UserFriend)(nil),        // 1: proto.UserFriend
}
var file_friend_get_response_proto_depIdxs = []int32{
	1, // 0: proto.FriendGetResponse.user_friends:type_name -> proto.UserFriend
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_friend_get_response_proto_init() }
func file_friend_get_response_proto_init() {
	if File_friend_get_response_proto != nil {
		return
	}
	file_user_friend_structure_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_friend_get_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendGetResponse); i {
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
			RawDescriptor: file_friend_get_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_friend_get_response_proto_goTypes,
		DependencyIndexes: file_friend_get_response_proto_depIdxs,
		MessageInfos:      file_friend_get_response_proto_msgTypes,
	}.Build()
	File_friend_get_response_proto = out.File
	file_friend_get_response_proto_rawDesc = nil
	file_friend_get_response_proto_goTypes = nil
	file_friend_get_response_proto_depIdxs = nil
}
