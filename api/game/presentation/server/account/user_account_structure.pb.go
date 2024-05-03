// ユーザーアカウント

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v4.24.4
// source: user_account_structure.proto

package account

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

type UserAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name     string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	LoginAt  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=login_at,json=loginAt,proto3" json:"login_at,omitempty"`
	LogoutAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=logout_at,json=logoutAt,proto3" json:"logout_at,omitempty"`
}

func (x *UserAccount) Reset() {
	*x = UserAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_account_structure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAccount) ProtoMessage() {}

func (x *UserAccount) ProtoReflect() protoreflect.Message {
	mi := &file_user_account_structure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAccount.ProtoReflect.Descriptor instead.
func (*UserAccount) Descriptor() ([]byte, []int) {
	return file_user_account_structure_proto_rawDescGZIP(), []int{0}
}

func (x *UserAccount) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserAccount) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserAccount) GetLoginAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LoginAt
	}
	return nil
}

func (x *UserAccount) GetLogoutAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LogoutAt
	}
	return nil
}

var File_user_account_structure_proto protoreflect.FileDescriptor

var file_user_account_structure_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x35, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x41, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x41, 0x74, 0x42,
	0x26, 0x5a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_account_structure_proto_rawDescOnce sync.Once
	file_user_account_structure_proto_rawDescData = file_user_account_structure_proto_rawDesc
)

func file_user_account_structure_proto_rawDescGZIP() []byte {
	file_user_account_structure_proto_rawDescOnce.Do(func() {
		file_user_account_structure_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_account_structure_proto_rawDescData)
	})
	return file_user_account_structure_proto_rawDescData
}

var file_user_account_structure_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_account_structure_proto_goTypes = []interface{}{
	(*UserAccount)(nil),           // 0: proto.UserAccount
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_user_account_structure_proto_depIdxs = []int32{
	1, // 0: proto.UserAccount.login_at:type_name -> google.protobuf.Timestamp
	1, // 1: proto.UserAccount.logout_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_account_structure_proto_init() }
func file_user_account_structure_proto_init() {
	if File_user_account_structure_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_account_structure_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAccount); i {
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
			RawDescriptor: file_user_account_structure_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_account_structure_proto_goTypes,
		DependencyIndexes: file_user_account_structure_proto_depIdxs,
		MessageInfos:      file_user_account_structure_proto_msgTypes,
	}.Build()
	File_user_account_structure_proto = out.File
	file_user_account_structure_proto_rawDesc = nil
	file_user_account_structure_proto_goTypes = nil
	file_user_account_structure_proto_depIdxs = nil
}
