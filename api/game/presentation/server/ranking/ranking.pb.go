// ランキング

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.24.4
// source: ranking.proto

package ranking

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

var File_ranking_proto protoreflect.FileDescriptor

var file_ranking_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x5f,
	0x67, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x57, 0x0a, 0x07, 0x52,
	0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x4c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61, 0x6e, 0x6b,
	0x69, 0x6e, 0x67, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61, 0x6e, 0x6b,
	0x69, 0x6e, 0x67, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x6d, 0x65,
	0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_ranking_proto_goTypes = []interface{}{
	(*RankingGetMasterRequest)(nil),  // 0: proto.RankingGetMasterRequest
	(*RankingGetMasterResponse)(nil), // 1: proto.RankingGetMasterResponse
}
var file_ranking_proto_depIdxs = []int32{
	0, // 0: proto.Ranking.GetMaster:input_type -> proto.RankingGetMasterRequest
	1, // 1: proto.Ranking.GetMaster:output_type -> proto.RankingGetMasterResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ranking_proto_init() }
func file_ranking_proto_init() {
	if File_ranking_proto != nil {
		return
	}
	file_ranking_get_master_request_proto_init()
	file_ranking_get_master_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ranking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ranking_proto_goTypes,
		DependencyIndexes: file_ranking_proto_depIdxs,
	}.Build()
	File_ranking_proto = out.File
	file_ranking_proto_rawDesc = nil
	file_ranking_proto_goTypes = nil
	file_ranking_proto_depIdxs = nil
}