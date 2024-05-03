// ユーザーログインボーナス

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v4.24.4
// source: user_login_bonus_structure.proto

package loginBonus

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

type UserLoginBonus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId             string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MasterLoginBonusId int64                  `protobuf:"varint,2,opt,name=master_login_bonus_id,json=masterLoginBonusId,proto3" json:"master_login_bonus_id,omitempty"`
	ReceivedAt         *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=received_at,json=receivedAt,proto3" json:"received_at,omitempty"`
}

func (x *UserLoginBonus) Reset() {
	*x = UserLoginBonus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_login_bonus_structure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginBonus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginBonus) ProtoMessage() {}

func (x *UserLoginBonus) ProtoReflect() protoreflect.Message {
	mi := &file_user_login_bonus_structure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginBonus.ProtoReflect.Descriptor instead.
func (*UserLoginBonus) Descriptor() ([]byte, []int) {
	return file_user_login_bonus_structure_proto_rawDescGZIP(), []int{0}
}

func (x *UserLoginBonus) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserLoginBonus) GetMasterLoginBonusId() int64 {
	if x != nil {
		return x.MasterLoginBonusId
	}
	return 0
}

func (x *UserLoginBonus) GetReceivedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ReceivedAt
	}
	return nil
}

var File_user_login_bonus_structure_proto protoreflect.FileDescriptor

var file_user_login_bonus_structure_proto_rawDesc = []byte{
	0x0a, 0x20, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e,
	0x75, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x0e, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x15, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x41, 0x74, 0x42, 0x29, 0x5a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61,
	0x6d, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_login_bonus_structure_proto_rawDescOnce sync.Once
	file_user_login_bonus_structure_proto_rawDescData = file_user_login_bonus_structure_proto_rawDesc
)

func file_user_login_bonus_structure_proto_rawDescGZIP() []byte {
	file_user_login_bonus_structure_proto_rawDescOnce.Do(func() {
		file_user_login_bonus_structure_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_login_bonus_structure_proto_rawDescData)
	})
	return file_user_login_bonus_structure_proto_rawDescData
}

var file_user_login_bonus_structure_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_login_bonus_structure_proto_goTypes = []interface{}{
	(*UserLoginBonus)(nil),        // 0: proto.UserLoginBonus
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_user_login_bonus_structure_proto_depIdxs = []int32{
	1, // 0: proto.UserLoginBonus.received_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_login_bonus_structure_proto_init() }
func file_user_login_bonus_structure_proto_init() {
	if File_user_login_bonus_structure_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_login_bonus_structure_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginBonus); i {
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
			RawDescriptor: file_user_login_bonus_structure_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_login_bonus_structure_proto_goTypes,
		DependencyIndexes: file_user_login_bonus_structure_proto_depIdxs,
		MessageInfos:      file_user_login_bonus_structure_proto_msgTypes,
	}.Build()
	File_user_login_bonus_structure_proto = out.File
	file_user_login_bonus_structure_proto_rawDesc = nil
	file_user_login_bonus_structure_proto_goTypes = nil
	file_user_login_bonus_structure_proto_depIdxs = nil
}
