// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.7
// source: state.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ServiceStatus int32

const (
	ServiceStatus_DOWN         ServiceStatus = 0
	ServiceStatus_INITIALIZING ServiceStatus = 1
	ServiceStatus_READY        ServiceStatus = 2
	ServiceStatus_TERMINATING  ServiceStatus = 3
)

// Enum value maps for ServiceStatus.
var (
	ServiceStatus_name = map[int32]string{
		0: "DOWN",
		1: "INITIALIZING",
		2: "READY",
		3: "TERMINATING",
	}
	ServiceStatus_value = map[string]int32{
		"DOWN":         0,
		"INITIALIZING": 1,
		"READY":        2,
		"TERMINATING":  3,
	}
)

func (x ServiceStatus) Enum() *ServiceStatus {
	p := new(ServiceStatus)
	*p = x
	return p
}

func (x ServiceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_state_proto_enumTypes[0].Descriptor()
}

func (ServiceStatus) Type() protoreflect.EnumType {
	return &file_state_proto_enumTypes[0]
}

func (x ServiceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceStatus.Descriptor instead.
func (ServiceStatus) EnumDescriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{0}
}

type ServiceState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status              ServiceStatus                  `protobuf:"varint,1,opt,name=status,proto3,enum=proto.ServiceStatus" json:"status,omitempty"`
	ServerVersion       string                         `protobuf:"bytes,2,opt,name=server_version,json=server-version,proto3" json:"server_version,omitempty"`
	SupportedMediaTypes map[string]*structpb.ListValue `protobuf:"bytes,3,rep,name=supported_media_types,json=supported-media-types,proto3" json:"supported_media_types,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ServiceState) Reset() {
	*x = ServiceState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceState) ProtoMessage() {}

func (x *ServiceState) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceState.ProtoReflect.Descriptor instead.
func (*ServiceState) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceState) GetStatus() ServiceStatus {
	if x != nil {
		return x.Status
	}
	return ServiceStatus_DOWN
}

func (x *ServiceState) GetServerVersion() string {
	if x != nil {
		return x.ServerVersion
	}
	return ""
}

func (x *ServiceState) GetSupportedMediaTypes() map[string]*structpb.ListValue {
	if x != nil {
		return x.SupportedMediaTypes
	}
	return nil
}

var File_state_proto protoreflect.FileDescriptor

var file_state_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xac, 0x02, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x62, 0x0a, 0x15, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x15, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x2d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x62, 0x0a,
	0x18, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x2a, 0x47, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x45, 0x52,
	0x4d, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x65, 0x72, 0x61, 0x69, 0x73, 0x6f,
	0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_state_proto_rawDescOnce sync.Once
	file_state_proto_rawDescData = file_state_proto_rawDesc
)

func file_state_proto_rawDescGZIP() []byte {
	file_state_proto_rawDescOnce.Do(func() {
		file_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_state_proto_rawDescData)
	})
	return file_state_proto_rawDescData
}

var file_state_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_state_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_state_proto_goTypes = []interface{}{
	(ServiceStatus)(0),         // 0: proto.ServiceStatus
	(*ServiceState)(nil),       // 1: proto.ServiceState
	nil,                        // 2: proto.ServiceState.SupportedMediaTypesEntry
	(*structpb.ListValue)(nil), // 3: google.protobuf.ListValue
}
var file_state_proto_depIdxs = []int32{
	0, // 0: proto.ServiceState.status:type_name -> proto.ServiceStatus
	2, // 1: proto.ServiceState.supported_media_types:type_name -> proto.ServiceState.SupportedMediaTypesEntry
	3, // 2: proto.ServiceState.SupportedMediaTypesEntry.value:type_name -> google.protobuf.ListValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_state_proto_init() }
func file_state_proto_init() {
	if File_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceState); i {
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
			RawDescriptor: file_state_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_state_proto_goTypes,
		DependencyIndexes: file_state_proto_depIdxs,
		EnumInfos:         file_state_proto_enumTypes,
		MessageInfos:      file_state_proto_msgTypes,
	}.Build()
	File_state_proto = out.File
	file_state_proto_rawDesc = nil
	file_state_proto_goTypes = nil
	file_state_proto_depIdxs = nil
}
