// ランキング

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.24.4
// source: master_ranking_structure.proto

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

type MasterRanking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MasterRankingEventId int64            `protobuf:"varint,2,opt,name=master_ranking_event_id,json=masterRankingEventId,proto3" json:"master_ranking_event_id,omitempty"`
	Name                 string           `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	RankingScopeType     RankingScopeType `protobuf:"varint,4,opt,name=ranking_scope_type,json=rankingScopeType,proto3,enum=proto.RankingScopeType" json:"ranking_scope_type,omitempty"`
	Limit                int32            `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *MasterRanking) Reset() {
	*x = MasterRanking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_ranking_structure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterRanking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterRanking) ProtoMessage() {}

func (x *MasterRanking) ProtoReflect() protoreflect.Message {
	mi := &file_master_ranking_structure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterRanking.ProtoReflect.Descriptor instead.
func (*MasterRanking) Descriptor() ([]byte, []int) {
	return file_master_ranking_structure_proto_rawDescGZIP(), []int{0}
}

func (x *MasterRanking) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MasterRanking) GetMasterRankingEventId() int64 {
	if x != nil {
		return x.MasterRankingEventId
	}
	return 0
}

func (x *MasterRanking) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MasterRanking) GetRankingScopeType() RankingScopeType {
	if x != nil {
		return x.RankingScopeType
	}
	return RankingScopeType_Room
}

func (x *MasterRanking) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

var File_master_ranking_structure_proto protoreflect.FileDescriptor

var file_master_ranking_structure_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x65, 0x6e, 0x75, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x0d, 0x4d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x17, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x12, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x10, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x42, 0x26, 0x5a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_master_ranking_structure_proto_rawDescOnce sync.Once
	file_master_ranking_structure_proto_rawDescData = file_master_ranking_structure_proto_rawDesc
)

func file_master_ranking_structure_proto_rawDescGZIP() []byte {
	file_master_ranking_structure_proto_rawDescOnce.Do(func() {
		file_master_ranking_structure_proto_rawDescData = protoimpl.X.CompressGZIP(file_master_ranking_structure_proto_rawDescData)
	})
	return file_master_ranking_structure_proto_rawDescData
}

var file_master_ranking_structure_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_master_ranking_structure_proto_goTypes = []interface{}{
	(*MasterRanking)(nil), // 0: proto.MasterRanking
	(RankingScopeType)(0), // 1: proto.RankingScopeType
}
var file_master_ranking_structure_proto_depIdxs = []int32{
	1, // 0: proto.MasterRanking.ranking_scope_type:type_name -> proto.RankingScopeType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_master_ranking_structure_proto_init() }
func file_master_ranking_structure_proto_init() {
	if File_master_ranking_structure_proto != nil {
		return
	}
	file_ranking_scope_type_enum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_master_ranking_structure_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterRanking); i {
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
			RawDescriptor: file_master_ranking_structure_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_master_ranking_structure_proto_goTypes,
		DependencyIndexes: file_master_ranking_structure_proto_depIdxs,
		MessageInfos:      file_master_ranking_structure_proto_msgTypes,
	}.Build()
	File_master_ranking_structure_proto = out.File
	file_master_ranking_structure_proto_rawDesc = nil
	file_master_ranking_structure_proto_goTypes = nil
	file_master_ranking_structure_proto_depIdxs = nil
}