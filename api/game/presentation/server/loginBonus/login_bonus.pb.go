// ログインボーナス

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: login_bonus.proto

package loginBonus

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

var File_login_bonus_proto protoreflect.FileDescriptor

var file_login_bonus_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x67, 0x65, 0x74, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73,
	0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73,
	0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfc, 0x01, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x12, 0x4c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42,
	0x6f, 0x6e, 0x75, 0x73, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x42, 0x6f, 0x6e, 0x75, 0x73, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42,
	0x6f, 0x6e, 0x75, 0x73, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61,
	0x6d, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_login_bonus_proto_goTypes = []interface{}{
	(*LoginBonusGetUserRequest)(nil),    // 0: proto.LoginBonusGetUserRequest
	(*LoginBonusGetMasterRequest)(nil),  // 1: proto.LoginBonusGetMasterRequest
	(*LoginBonusReceiveRequest)(nil),    // 2: proto.LoginBonusReceiveRequest
	(*LoginBonusGetUserResponse)(nil),   // 3: proto.LoginBonusGetUserResponse
	(*LoginBonusGetMasterResponse)(nil), // 4: proto.LoginBonusGetMasterResponse
	(*LoginBonusReceiveResponse)(nil),   // 5: proto.LoginBonusReceiveResponse
}
var file_login_bonus_proto_depIdxs = []int32{
	0, // 0: proto.LoginBonus.GetUser:input_type -> proto.LoginBonusGetUserRequest
	1, // 1: proto.LoginBonus.GetMaster:input_type -> proto.LoginBonusGetMasterRequest
	2, // 2: proto.LoginBonus.Receive:input_type -> proto.LoginBonusReceiveRequest
	3, // 3: proto.LoginBonus.GetUser:output_type -> proto.LoginBonusGetUserResponse
	4, // 4: proto.LoginBonus.GetMaster:output_type -> proto.LoginBonusGetMasterResponse
	5, // 5: proto.LoginBonus.Receive:output_type -> proto.LoginBonusReceiveResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_login_bonus_proto_init() }
func file_login_bonus_proto_init() {
	if File_login_bonus_proto != nil {
		return
	}
	file_login_bonus_get_user_request_proto_init()
	file_login_bonus_get_user_response_proto_init()
	file_login_bonus_get_master_request_proto_init()
	file_login_bonus_get_master_response_proto_init()
	file_login_bonus_receive_request_proto_init()
	file_login_bonus_receive_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_login_bonus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_login_bonus_proto_goTypes,
		DependencyIndexes: file_login_bonus_proto_depIdxs,
	}.Build()
	File_login_bonus_proto = out.File
	file_login_bonus_proto_rawDesc = nil
	file_login_bonus_proto_goTypes = nil
	file_login_bonus_proto_depIdxs = nil
}
