// ワールドランキング

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v4.24.4
// source: common_ranking_world_structure.proto

package ranking

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CommonRankingWorld struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MasterRankingId int64                  `protobuf:"varint,1,opt,name=master_ranking_id,json=masterRankingId,proto3" json:"master_ranking_id,omitempty"`
	UserId          string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Score           int32                  `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
	RankedAt        *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=ranked_at,json=rankedAt,proto3" json:"ranked_at,omitempty"`
}

func (x *CommonRankingWorld) Reset() {
	*x = CommonRankingWorld{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_ranking_world_structure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonRankingWorld) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonRankingWorld) ProtoMessage() {}

func (x *CommonRankingWorld) ProtoReflect() protoreflect.Message {
	mi := &file_common_ranking_world_structure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonRankingWorld.ProtoReflect.Descriptor instead.
func (*CommonRankingWorld) Descriptor() ([]byte, []int) {
	return file_common_ranking_world_structure_proto_rawDescGZIP(), []int{0}
}

func (x *CommonRankingWorld) GetMasterRankingId() int64 {
	if x != nil {
		return x.MasterRankingId
	}
	return 0
}

func (x *CommonRankingWorld) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CommonRankingWorld) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *CommonRankingWorld) GetRankedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RankedAt
	}
	return nil
}

var File_common_ranking_world_structure_proto protoreflect.FileDescriptor

var file_common_ranking_world_structure_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8,
	0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x37, 0x0a, 0x09, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x08, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x64, 0x41, 0x74, 0x42, 0x26, 0x5a, 0x24, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_ranking_world_structure_proto_rawDescOnce sync.Once
	file_common_ranking_world_structure_proto_rawDescData = file_common_ranking_world_structure_proto_rawDesc
)

func file_common_ranking_world_structure_proto_rawDescGZIP() []byte {
	file_common_ranking_world_structure_proto_rawDescOnce.Do(func() {
		file_common_ranking_world_structure_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_ranking_world_structure_proto_rawDescData)
	})
	return file_common_ranking_world_structure_proto_rawDescData
}

var file_common_ranking_world_structure_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_ranking_world_structure_proto_goTypes = []interface{}{
	(*CommonRankingWorld)(nil),    // 0: proto.CommonRankingWorld
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_common_ranking_world_structure_proto_depIdxs = []int32{
	1, // 0: proto.CommonRankingWorld.ranked_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_ranking_world_structure_proto_init() }
func file_common_ranking_world_structure_proto_init() {
	if File_common_ranking_world_structure_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_ranking_world_structure_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonRankingWorld); i {
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
			RawDescriptor: file_common_ranking_world_structure_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_ranking_world_structure_proto_goTypes,
		DependencyIndexes: file_common_ranking_world_structure_proto_depIdxs,
		MessageInfos:      file_common_ranking_world_structure_proto_msgTypes,
	}.Build()
	File_common_ranking_world_structure_proto = out.File
	file_common_ranking_world_structure_proto_rawDesc = nil
	file_common_ranking_world_structure_proto_goTypes = nil
	file_common_ranking_world_structure_proto_depIdxs = nil
}
