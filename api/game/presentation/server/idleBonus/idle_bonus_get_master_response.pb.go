// 放置ボーナスマスター取得レスポンス

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.24.4
// source: idle_bonus_get_master_response.proto

package idleBonus

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

type IdleBonusGetMasterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MasterIdleBonus          *MasterIdleBonus           `protobuf:"bytes,1,opt,name=master_idle_bonus,json=masterIdleBonus,proto3,oneof" json:"master_idle_bonus,omitempty"`
	MasterIdleBonusEvent     *MasterIdleBonusEvent      `protobuf:"bytes,2,opt,name=master_idle_bonus_event,json=masterIdleBonusEvent,proto3,oneof" json:"master_idle_bonus_event,omitempty"`
	MasterIdleBonusItems     []*MasterIdleBonusItem     `protobuf:"bytes,3,rep,name=master_idle_bonus_items,json=masterIdleBonusItems,proto3" json:"master_idle_bonus_items,omitempty"`
	MasterIdleBonusSchedules []*MasterIdleBonusSchedule `protobuf:"bytes,4,rep,name=master_idle_bonus_schedules,json=masterIdleBonusSchedules,proto3" json:"master_idle_bonus_schedules,omitempty"`
}

func (x *IdleBonusGetMasterResponse) Reset() {
	*x = IdleBonusGetMasterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idle_bonus_get_master_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdleBonusGetMasterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdleBonusGetMasterResponse) ProtoMessage() {}

func (x *IdleBonusGetMasterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idle_bonus_get_master_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdleBonusGetMasterResponse.ProtoReflect.Descriptor instead.
func (*IdleBonusGetMasterResponse) Descriptor() ([]byte, []int) {
	return file_idle_bonus_get_master_response_proto_rawDescGZIP(), []int{0}
}

func (x *IdleBonusGetMasterResponse) GetMasterIdleBonus() *MasterIdleBonus {
	if x != nil {
		return x.MasterIdleBonus
	}
	return nil
}

func (x *IdleBonusGetMasterResponse) GetMasterIdleBonusEvent() *MasterIdleBonusEvent {
	if x != nil {
		return x.MasterIdleBonusEvent
	}
	return nil
}

func (x *IdleBonusGetMasterResponse) GetMasterIdleBonusItems() []*MasterIdleBonusItem {
	if x != nil {
		return x.MasterIdleBonusItems
	}
	return nil
}

func (x *IdleBonusGetMasterResponse) GetMasterIdleBonusSchedules() []*MasterIdleBonusSchedule {
	if x != nil {
		return x.MasterIdleBonusSchedules
	}
	return nil
}

var File_idle_bonus_get_master_response_proto protoreflect.FileDescriptor

var file_idle_bonus_get_master_response_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x67, 0x65, 0x74,
	0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73,
	0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x27, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x62, 0x6f,
	0x6e, 0x75, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2a, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x62,
	0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x03,
	0x0a, 0x1a, 0x49, 0x64, 0x6c, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x11,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x62, 0x6f, 0x6e, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x6c, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x48,
	0x00, 0x52, 0x0f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x6c, 0x65, 0x42, 0x6f, 0x6e,
	0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x57, 0x0a, 0x17, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x6c, 0x65, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x6c, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x48, 0x01, 0x52, 0x14, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x6c,
	0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x51,
	0x0a, 0x17, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x62, 0x6f,
	0x6e, 0x75, 0x73, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x6c, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x14, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x6c, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x5d, 0x0a, 0x1b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x6c, 0x65,
	0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x6c, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x18, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x6c, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73,
	0x42, 0x14, 0x0a, 0x12, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x6c, 0x65,
	0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x42, 0x28, 0x5a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x69, 0x64, 0x6c, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idle_bonus_get_master_response_proto_rawDescOnce sync.Once
	file_idle_bonus_get_master_response_proto_rawDescData = file_idle_bonus_get_master_response_proto_rawDesc
)

func file_idle_bonus_get_master_response_proto_rawDescGZIP() []byte {
	file_idle_bonus_get_master_response_proto_rawDescOnce.Do(func() {
		file_idle_bonus_get_master_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_idle_bonus_get_master_response_proto_rawDescData)
	})
	return file_idle_bonus_get_master_response_proto_rawDescData
}

var file_idle_bonus_get_master_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_idle_bonus_get_master_response_proto_goTypes = []interface{}{
	(*IdleBonusGetMasterResponse)(nil), // 0: proto.IdleBonusGetMasterResponse
	(*MasterIdleBonus)(nil),            // 1: proto.MasterIdleBonus
	(*MasterIdleBonusEvent)(nil),       // 2: proto.MasterIdleBonusEvent
	(*MasterIdleBonusItem)(nil),        // 3: proto.MasterIdleBonusItem
	(*MasterIdleBonusSchedule)(nil),    // 4: proto.MasterIdleBonusSchedule
}
var file_idle_bonus_get_master_response_proto_depIdxs = []int32{
	1, // 0: proto.IdleBonusGetMasterResponse.master_idle_bonus:type_name -> proto.MasterIdleBonus
	2, // 1: proto.IdleBonusGetMasterResponse.master_idle_bonus_event:type_name -> proto.MasterIdleBonusEvent
	3, // 2: proto.IdleBonusGetMasterResponse.master_idle_bonus_items:type_name -> proto.MasterIdleBonusItem
	4, // 3: proto.IdleBonusGetMasterResponse.master_idle_bonus_schedules:type_name -> proto.MasterIdleBonusSchedule
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_idle_bonus_get_master_response_proto_init() }
func file_idle_bonus_get_master_response_proto_init() {
	if File_idle_bonus_get_master_response_proto != nil {
		return
	}
	file_master_idle_bonus_structure_proto_init()
	file_master_idle_bonus_event_structure_proto_init()
	file_master_idle_bonus_item_structure_proto_init()
	file_master_idle_bonus_schedule_structure_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_idle_bonus_get_master_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdleBonusGetMasterResponse); i {
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
	file_idle_bonus_get_master_response_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_idle_bonus_get_master_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_idle_bonus_get_master_response_proto_goTypes,
		DependencyIndexes: file_idle_bonus_get_master_response_proto_depIdxs,
		MessageInfos:      file_idle_bonus_get_master_response_proto_msgTypes,
	}.Build()
	File_idle_bonus_get_master_response_proto = out.File
	file_idle_bonus_get_master_response_proto_rawDesc = nil
	file_idle_bonus_get_master_response_proto_goTypes = nil
	file_idle_bonus_get_master_response_proto_depIdxs = nil
}