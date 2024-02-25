// ログインボーナス受け取りレスポンス

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: login_bonus_receive_response.proto

package loginBonus

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

type LoginBonusReceiveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserLoginBonus           *UserLoginBonus           `protobuf:"bytes,1,opt,name=user_login_bonus,json=userLoginBonus,proto3,oneof" json:"user_login_bonus,omitempty"`
	MasterLoginBonus         *MasterLoginBonus         `protobuf:"bytes,2,opt,name=master_login_bonus,json=masterLoginBonus,proto3,oneof" json:"master_login_bonus,omitempty"`
	MasterLoginBonusEvent    *MasterLoginBonusEvent    `protobuf:"bytes,3,opt,name=master_login_bonus_event,json=masterLoginBonusEvent,proto3,oneof" json:"master_login_bonus_event,omitempty"`
	MasterLoginBonusItems    []*MasterLoginBonusItem   `protobuf:"bytes,4,rep,name=master_login_bonus_items,json=masterLoginBonusItems,proto3" json:"master_login_bonus_items,omitempty"`
	MasterLoginBonusSchedule *MasterLoginBonusSchedule `protobuf:"bytes,5,opt,name=master_login_bonus_schedule,json=masterLoginBonusSchedule,proto3,oneof" json:"master_login_bonus_schedule,omitempty"`
}

func (x *LoginBonusReceiveResponse) Reset() {
	*x = LoginBonusReceiveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_bonus_receive_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginBonusReceiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginBonusReceiveResponse) ProtoMessage() {}

func (x *LoginBonusReceiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_login_bonus_receive_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginBonusReceiveResponse.ProtoReflect.Descriptor instead.
func (*LoginBonusReceiveResponse) Descriptor() ([]byte, []int) {
	return file_login_bonus_receive_response_proto_rawDescGZIP(), []int{0}
}

func (x *LoginBonusReceiveResponse) GetUserLoginBonus() *UserLoginBonus {
	if x != nil {
		return x.UserLoginBonus
	}
	return nil
}

func (x *LoginBonusReceiveResponse) GetMasterLoginBonus() *MasterLoginBonus {
	if x != nil {
		return x.MasterLoginBonus
	}
	return nil
}

func (x *LoginBonusReceiveResponse) GetMasterLoginBonusEvent() *MasterLoginBonusEvent {
	if x != nil {
		return x.MasterLoginBonusEvent
	}
	return nil
}

func (x *LoginBonusReceiveResponse) GetMasterLoginBonusItems() []*MasterLoginBonusItem {
	if x != nil {
		return x.MasterLoginBonusItems
	}
	return nil
}

func (x *LoginBonusReceiveResponse) GetMasterLoginBonusSchedule() *MasterLoginBonusSchedule {
	if x != nil {
		return x.MasterLoginBonusSchedule
	}
	return nil
}

var File_login_bonus_receive_response_proto protoreflect.FileDescriptor

var file_login_bonus_receive_response_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75,
	0x73, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x28, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f,
	0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xad, 0x04, 0x0a, 0x19, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f,
	0x6e, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73,
	0x48, 0x00, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6e,
	0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x4a, 0x0a, 0x12, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x48, 0x01, 0x52, 0x10, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x88, 0x01,
	0x01, 0x12, 0x5a, 0x0a, 0x18, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x48, 0x02, 0x52, 0x15, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x42, 0x6f, 0x6e, 0x75, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x54, 0x0a,
	0x18, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f,
	0x6e, 0x75, 0x73, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x15, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x63, 0x0a, 0x1b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75,
	0x73, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x48, 0x03, 0x52, 0x18, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x42, 0x15, 0x0a,
	0x13, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62,
	0x6f, 0x6e, 0x75, 0x73, 0x42, 0x1b, 0x0a, 0x19, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x42, 0x1e, 0x0a, 0x1c, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x42, 0x29, 0x5a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72,
	0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_login_bonus_receive_response_proto_rawDescOnce sync.Once
	file_login_bonus_receive_response_proto_rawDescData = file_login_bonus_receive_response_proto_rawDesc
)

func file_login_bonus_receive_response_proto_rawDescGZIP() []byte {
	file_login_bonus_receive_response_proto_rawDescOnce.Do(func() {
		file_login_bonus_receive_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_login_bonus_receive_response_proto_rawDescData)
	})
	return file_login_bonus_receive_response_proto_rawDescData
}

var file_login_bonus_receive_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_login_bonus_receive_response_proto_goTypes = []interface{}{
	(*LoginBonusReceiveResponse)(nil), // 0: proto.LoginBonusReceiveResponse
	(*UserLoginBonus)(nil),            // 1: proto.UserLoginBonus
	(*MasterLoginBonus)(nil),          // 2: proto.MasterLoginBonus
	(*MasterLoginBonusEvent)(nil),     // 3: proto.MasterLoginBonusEvent
	(*MasterLoginBonusItem)(nil),      // 4: proto.MasterLoginBonusItem
	(*MasterLoginBonusSchedule)(nil),  // 5: proto.MasterLoginBonusSchedule
}
var file_login_bonus_receive_response_proto_depIdxs = []int32{
	1, // 0: proto.LoginBonusReceiveResponse.user_login_bonus:type_name -> proto.UserLoginBonus
	2, // 1: proto.LoginBonusReceiveResponse.master_login_bonus:type_name -> proto.MasterLoginBonus
	3, // 2: proto.LoginBonusReceiveResponse.master_login_bonus_event:type_name -> proto.MasterLoginBonusEvent
	4, // 3: proto.LoginBonusReceiveResponse.master_login_bonus_items:type_name -> proto.MasterLoginBonusItem
	5, // 4: proto.LoginBonusReceiveResponse.master_login_bonus_schedule:type_name -> proto.MasterLoginBonusSchedule
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_login_bonus_receive_response_proto_init() }
func file_login_bonus_receive_response_proto_init() {
	if File_login_bonus_receive_response_proto != nil {
		return
	}
	file_user_login_bonus_structure_proto_init()
	file_master_login_bonus_structure_proto_init()
	file_master_login_bonus_event_structure_proto_init()
	file_master_login_bonus_item_structure_proto_init()
	file_master_login_bonus_schedule_structure_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_login_bonus_receive_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginBonusReceiveResponse); i {
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
	file_login_bonus_receive_response_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_login_bonus_receive_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_login_bonus_receive_response_proto_goTypes,
		DependencyIndexes: file_login_bonus_receive_response_proto_depIdxs,
		MessageInfos:      file_login_bonus_receive_response_proto_msgTypes,
	}.Build()
	File_login_bonus_receive_response_proto = out.File
	file_login_bonus_receive_response_proto_rawDesc = nil
	file_login_bonus_receive_response_proto_goTypes = nil
	file_login_bonus_receive_response_proto_depIdxs = nil
}
