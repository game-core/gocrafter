// ランキング更新リクエスト

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v4.24.4
// source: ranking_update_request.proto

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

type RankingUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MasterRankingId int64  `protobuf:"varint,2,opt,name=master_ranking_id,json=masterRankingId,proto3" json:"master_ranking_id,omitempty"`
	RoomId          string `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Score           int32  `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *RankingUpdateRequest) Reset() {
	*x = RankingUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ranking_update_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankingUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankingUpdateRequest) ProtoMessage() {}

func (x *RankingUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ranking_update_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankingUpdateRequest.ProtoReflect.Descriptor instead.
func (*RankingUpdateRequest) Descriptor() ([]byte, []int) {
	return file_ranking_update_request_proto_rawDescGZIP(), []int{0}
}

func (x *RankingUpdateRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RankingUpdateRequest) GetMasterRankingId() int64 {
	if x != nil {
		return x.MasterRankingId
	}
	return 0
}

func (x *RankingUpdateRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *RankingUpdateRequest) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

var File_ranking_update_request_proto protoreflect.FileDescriptor

var file_ranking_update_request_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x14, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ranking_update_request_proto_rawDescOnce sync.Once
	file_ranking_update_request_proto_rawDescData = file_ranking_update_request_proto_rawDesc
)

func file_ranking_update_request_proto_rawDescGZIP() []byte {
	file_ranking_update_request_proto_rawDescOnce.Do(func() {
		file_ranking_update_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_ranking_update_request_proto_rawDescData)
	})
	return file_ranking_update_request_proto_rawDescData
}

var file_ranking_update_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ranking_update_request_proto_goTypes = []interface{}{
	(*RankingUpdateRequest)(nil), // 0: proto.RankingUpdateRequest
}
var file_ranking_update_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ranking_update_request_proto_init() }
func file_ranking_update_request_proto_init() {
	if File_ranking_update_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ranking_update_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankingUpdateRequest); i {
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
			RawDescriptor: file_ranking_update_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ranking_update_request_proto_goTypes,
		DependencyIndexes: file_ranking_update_request_proto_depIdxs,
		MessageInfos:      file_ranking_update_request_proto_msgTypes,
	}.Build()
	File_ranking_update_request_proto = out.File
	file_ranking_update_request_proto_rawDesc = nil
	file_ranking_update_request_proto_goTypes = nil
	file_ranking_update_request_proto_depIdxs = nil
}
