// フレンド

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: friend.proto

package friend

import (
	reflect "reflect"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_friend_proto protoreflect.FileDescriptor

var file_friend_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x65,
	0x6e, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x66, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x66, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9d, 0x02, 0x0a, 0x06, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x12, 0x3b, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x44, 0x69, 0x73, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x44, 0x69, 0x73, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_friend_proto_goTypes = []interface{}{
	(*FriendSendRequest)(nil),        // 0: proto.FriendSendRequest
	(*FriendApproveRequest)(nil),     // 1: proto.FriendApproveRequest
	(*FriendDisapproveRequest)(nil),  // 2: proto.FriendDisapproveRequest
	(*FriendDeleteRequest)(nil),      // 3: proto.FriendDeleteRequest
	(*FriendSendResponse)(nil),       // 4: proto.FriendSendResponse
	(*FriendApproveResponse)(nil),    // 5: proto.FriendApproveResponse
	(*FriendDisapproveResponse)(nil), // 6: proto.FriendDisapproveResponse
	(*FriendDeleteResponse)(nil),     // 7: proto.FriendDeleteResponse
}
var file_friend_proto_depIdxs = []int32{
	0, // 0: proto.Friend.Send:input_type -> proto.FriendSendRequest
	1, // 1: proto.Friend.Approve:input_type -> proto.FriendApproveRequest
	2, // 2: proto.Friend.Disapprove:input_type -> proto.FriendDisapproveRequest
	3, // 3: proto.Friend.Delete:input_type -> proto.FriendDeleteRequest
	4, // 4: proto.Friend.Send:output_type -> proto.FriendSendResponse
	5, // 5: proto.Friend.Approve:output_type -> proto.FriendApproveResponse
	6, // 6: proto.Friend.Disapprove:output_type -> proto.FriendDisapproveResponse
	7, // 7: proto.Friend.Delete:output_type -> proto.FriendDeleteResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_friend_proto_init() }
func file_friend_proto_init() {
	if File_friend_proto != nil {
		return
	}
	file_friend_send_request_proto_init()
	file_friend_send_response_proto_init()
	file_friend_approve_request_proto_init()
	file_friend_approve_response_proto_init()
	file_friend_disapprove_request_proto_init()
	file_friend_disapprove_response_proto_init()
	file_friend_delete_request_proto_init()
	file_friend_delete_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_friend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_friend_proto_goTypes,
		DependencyIndexes: file_friend_proto_depIdxs,
	}.Build()
	File_friend_proto = out.File
	file_friend_proto_rawDesc = nil
	file_friend_proto_goTypes = nil
	file_friend_proto_depIdxs = nil
}
