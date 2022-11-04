// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: alert.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Summary string `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_alert_proto_rawDescGZIP(), []int{0}
}

func (x *Notification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notification) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Notification) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SendAlertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notification *Notification `protobuf:"bytes,1,opt,name=notification,proto3" json:"notification,omitempty"`
}

func (x *SendAlertRequest) Reset() {
	*x = SendAlertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendAlertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendAlertRequest) ProtoMessage() {}

func (x *SendAlertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendAlertRequest.ProtoReflect.Descriptor instead.
func (*SendAlertRequest) Descriptor() ([]byte, []int) {
	return file_alert_proto_rawDescGZIP(), []int{1}
}

func (x *SendAlertRequest) GetNotification() *Notification {
	if x != nil {
		return x.Notification
	}
	return nil
}

type SendAlertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SendAlertResponse) Reset() {
	*x = SendAlertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendAlertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendAlertResponse) ProtoMessage() {}

func (x *SendAlertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendAlertResponse.ProtoReflect.Descriptor instead.
func (*SendAlertResponse) Descriptor() ([]byte, []int) {
	return file_alert_proto_rawDescGZIP(), []int{2}
}

func (x *SendAlertResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_alert_proto protoreflect.FileDescriptor

var file_alert_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x50, 0x6f, 0x63, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x0c, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x55, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x50, 0x6f, 0x63, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x23, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xae, 0x01, 0x0a, 0x0c,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x09,
	0x53, 0x65, 0x6e, 0x64, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x21, 0x2e, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x50, 0x6f, 0x63, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x50, 0x6f, 0x63, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x48, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x50, 0x6f, 0x63, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x42, 0xad, 0x01, 0x0a,
	0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x50, 0x6f, 0x63, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x42, 0x0a, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52,
	0x63, 0x6f, 0x6e, 0x64, 0x75, 0x72, 0x75, 0x2f, 0x52, 0x65, 0x61, 0x63, 0x74, 0x2d, 0x6e, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x77, 0x65, 0x62, 0x2f, 0x67, 0x65,
	0x6e, 0xa2, 0x02, 0x03, 0x4d, 0x41, 0x58, 0xaa, 0x02, 0x0f, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x50, 0x6f, 0x63, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0xca, 0x02, 0x0f, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x50, 0x6f, 0x63, 0x5c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0xe2, 0x02, 0x1b, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x50, 0x6f, 0x63, 0x5c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x50, 0x6f, 0x63, 0x3a, 0x3a, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alert_proto_rawDescOnce sync.Once
	file_alert_proto_rawDescData = file_alert_proto_rawDesc
)

func file_alert_proto_rawDescGZIP() []byte {
	file_alert_proto_rawDescOnce.Do(func() {
		file_alert_proto_rawDescData = protoimpl.X.CompressGZIP(file_alert_proto_rawDescData)
	})
	return file_alert_proto_rawDescData
}

var file_alert_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_alert_proto_goTypes = []interface{}{
	(*Notification)(nil),      // 0: mobilePoc.alert.Notification
	(*SendAlertRequest)(nil),  // 1: mobilePoc.alert.SendAlertRequest
	(*SendAlertResponse)(nil), // 2: mobilePoc.alert.SendAlertResponse
	(*emptypb.Empty)(nil),     // 3: google.protobuf.Empty
}
var file_alert_proto_depIdxs = []int32{
	0, // 0: mobilePoc.alert.SendAlertRequest.notification:type_name -> mobilePoc.alert.Notification
	1, // 1: mobilePoc.alert.AlertService.SendAlert:input_type -> mobilePoc.alert.SendAlertRequest
	3, // 2: mobilePoc.alert.AlertService.ListenAlert:input_type -> google.protobuf.Empty
	2, // 3: mobilePoc.alert.AlertService.SendAlert:output_type -> mobilePoc.alert.SendAlertResponse
	0, // 4: mobilePoc.alert.AlertService.ListenAlert:output_type -> mobilePoc.alert.Notification
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_alert_proto_init() }
func file_alert_proto_init() {
	if File_alert_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
		file_alert_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendAlertRequest); i {
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
		file_alert_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendAlertResponse); i {
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
			RawDescriptor: file_alert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_alert_proto_goTypes,
		DependencyIndexes: file_alert_proto_depIdxs,
		MessageInfos:      file_alert_proto_msgTypes,
	}.Build()
	File_alert_proto = out.File
	file_alert_proto_rawDesc = nil
	file_alert_proto_goTypes = nil
	file_alert_proto_depIdxs = nil
}
