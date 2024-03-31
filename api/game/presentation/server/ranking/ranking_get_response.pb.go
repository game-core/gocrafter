// ランキング取得レスポンス

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.24.4
// source: ranking_get_response.proto

package ranking

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

type RankingGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonRankingRooms  []*CommonRankingRoom  `protobuf:"bytes,1,rep,name=common_ranking_rooms,json=commonRankingRooms,proto3" json:"common_ranking_rooms,omitempty"`
	CommonRankingWorlds []*CommonRankingWorld `protobuf:"bytes,2,rep,name=common_ranking_worlds,json=commonRankingWorlds,proto3" json:"common_ranking_worlds,omitempty"`
}

func (x *RankingGetResponse) Reset() {
	*x = RankingGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ranking_get_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankingGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankingGetResponse) ProtoMessage() {}

func (x *RankingGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ranking_get_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankingGetResponse.ProtoReflect.Descriptor instead.
func (*RankingGetResponse) Descriptor() ([]byte, []int) {
	return file_ranking_get_response_proto_rawDescGZIP(), []int{0}
}

func (x *RankingGetResponse) GetCommonRankingRooms() []*CommonRankingRoom {
	if x != nil {
		return x.CommonRankingRooms
	}
	return nil
}

func (x *RankingGetResponse) GetCommonRankingWorlds() []*CommonRankingWorld {
	if x != nil {
		return x.CommonRankingWorlds
	}
	return nil
}

var File_ranking_get_response_proto protoreflect.FileDescriptor

var file_ranking_get_response_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x6e, 0x6b,
	0x69, 0x6e, 0x67, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf,
	0x01, 0x0a, 0x12, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f,
	0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x12, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x6f, 0x6f, 0x6d,
	0x73, 0x12, 0x4d, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x6e, 0x6b,
	0x69, 0x6e, 0x67, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52,
	0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x13, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x73,
	0x42, 0x26, 0x5a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ranking_get_response_proto_rawDescOnce sync.Once
	file_ranking_get_response_proto_rawDescData = file_ranking_get_response_proto_rawDesc
)

func file_ranking_get_response_proto_rawDescGZIP() []byte {
	file_ranking_get_response_proto_rawDescOnce.Do(func() {
		file_ranking_get_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_ranking_get_response_proto_rawDescData)
	})
	return file_ranking_get_response_proto_rawDescData
}

var file_ranking_get_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ranking_get_response_proto_goTypes = []interface{}{
	(*RankingGetResponse)(nil), // 0: proto.RankingGetResponse
	(*CommonRankingRoom)(nil),  // 1: proto.CommonRankingRoom
	(*CommonRankingWorld)(nil), // 2: proto.CommonRankingWorld
}
var file_ranking_get_response_proto_depIdxs = []int32{
	1, // 0: proto.RankingGetResponse.common_ranking_rooms:type_name -> proto.CommonRankingRoom
	2, // 1: proto.RankingGetResponse.common_ranking_worlds:type_name -> proto.CommonRankingWorld
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ranking_get_response_proto_init() }
func file_ranking_get_response_proto_init() {
	if File_ranking_get_response_proto != nil {
		return
	}
	file_common_ranking_room_structure_proto_init()
	file_common_ranking_world_structure_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ranking_get_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankingGetResponse); i {
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
			RawDescriptor: file_ranking_get_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ranking_get_response_proto_goTypes,
		DependencyIndexes: file_ranking_get_response_proto_depIdxs,
		MessageInfos:      file_ranking_get_response_proto_msgTypes,
	}.Build()
	File_ranking_get_response_proto = out.File
	file_ranking_get_response_proto_rawDesc = nil
	file_ranking_get_response_proto_goTypes = nil
	file_ranking_get_response_proto_depIdxs = nil
}
