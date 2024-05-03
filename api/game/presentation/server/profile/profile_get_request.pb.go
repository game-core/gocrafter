// プロフィール取得リクエスト

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v4.24.4
// source: profile_get_request.proto

package profile

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

type ProfileGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ProfileGetRequest) Reset() {
	*x = ProfileGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_get_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileGetRequest) ProtoMessage() {}

func (x *ProfileGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_get_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileGetRequest.ProtoReflect.Descriptor instead.
func (*ProfileGetRequest) Descriptor() ([]byte, []int) {
	return file_profile_get_request_proto_rawDescGZIP(), []int{0}
}

func (x *ProfileGetRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_profile_get_request_proto protoreflect.FileDescriptor

var file_profile_get_request_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x42, 0x26, 0x5a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_profile_get_request_proto_rawDescOnce sync.Once
	file_profile_get_request_proto_rawDescData = file_profile_get_request_proto_rawDesc
)

func file_profile_get_request_proto_rawDescGZIP() []byte {
	file_profile_get_request_proto_rawDescOnce.Do(func() {
		file_profile_get_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_profile_get_request_proto_rawDescData)
	})
	return file_profile_get_request_proto_rawDescData
}

var file_profile_get_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_profile_get_request_proto_goTypes = []interface{}{
	(*ProfileGetRequest)(nil), // 0: proto.ProfileGetRequest
}
var file_profile_get_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_profile_get_request_proto_init() }
func file_profile_get_request_proto_init() {
	if File_profile_get_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_profile_get_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileGetRequest); i {
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
			RawDescriptor: file_profile_get_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_profile_get_request_proto_goTypes,
		DependencyIndexes: file_profile_get_request_proto_depIdxs,
		MessageInfos:      file_profile_get_request_proto_msgTypes,
	}.Build()
	File_profile_get_request_proto = out.File
	file_profile_get_request_proto_rawDesc = nil
	file_profile_get_request_proto_goTypes = nil
	file_profile_get_request_proto_depIdxs = nil
}
