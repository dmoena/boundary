// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.11.4
// source: controller/api/resources/hosts/v1/host.proto

package hosts

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Host contains all fields related to a Host resource
type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the host
	// Output only.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The type of the resource, to help differentiate schemas
	Type *wrappers.StringValue `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Optional name for identification purposes
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Optional user-set description for identification purposes
	Description *wrappers.StringValue `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// The time this resource was created
	// Output only.
	CreatedTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_time,proto3" json:"created_time,omitempty"`
	// The time this resource was last updated
	// Output only.
	UpdatedTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_time,proto3" json:"updated_time,omitempty"`
	// Whether the host is disabled
	Disabled *wrappers.BoolValue `protobuf:"bytes,7,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// The address (DNS or IP name) used to reach the host
	Address *wrappers.StringValue `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resources_hosts_v1_host_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_hosts_v1_host_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_hosts_v1_host_proto_rawDescGZIP(), []int{0}
}

func (x *Host) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Host) GetType() *wrappers.StringValue {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Host) GetName() *wrappers.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Host) GetDescription() *wrappers.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Host) GetCreatedTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *Host) GetUpdatedTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *Host) GetDisabled() *wrappers.BoolValue {
	if x != nil {
		return x.Disabled
	}
	return nil
}

func (x *Host) GetAddress() *wrappers.StringValue {
	if x != nil {
		return x.Address
	}
	return nil
}

var File_controller_api_resources_hosts_v1_host_proto protoreflect.FileDescriptor

var file_controller_api_resources_hosts_v1_host_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xaa, 0x03, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x3e, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x3e, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x3e, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x36, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x36, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42,
	0x53, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x74, 0x6f, 0x77,
	0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x3b, 0x68,
	0x6f, 0x73, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_api_resources_hosts_v1_host_proto_rawDescOnce sync.Once
	file_controller_api_resources_hosts_v1_host_proto_rawDescData = file_controller_api_resources_hosts_v1_host_proto_rawDesc
)

func file_controller_api_resources_hosts_v1_host_proto_rawDescGZIP() []byte {
	file_controller_api_resources_hosts_v1_host_proto_rawDescOnce.Do(func() {
		file_controller_api_resources_hosts_v1_host_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_api_resources_hosts_v1_host_proto_rawDescData)
	})
	return file_controller_api_resources_hosts_v1_host_proto_rawDescData
}

var file_controller_api_resources_hosts_v1_host_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_controller_api_resources_hosts_v1_host_proto_goTypes = []interface{}{
	(*Host)(nil),                 // 0: controller.api.resources.hosts.v1.Host
	(*wrappers.StringValue)(nil), // 1: google.protobuf.StringValue
	(*timestamp.Timestamp)(nil),  // 2: google.protobuf.Timestamp
	(*wrappers.BoolValue)(nil),   // 3: google.protobuf.BoolValue
}
var file_controller_api_resources_hosts_v1_host_proto_depIdxs = []int32{
	1, // 0: controller.api.resources.hosts.v1.Host.type:type_name -> google.protobuf.StringValue
	1, // 1: controller.api.resources.hosts.v1.Host.name:type_name -> google.protobuf.StringValue
	1, // 2: controller.api.resources.hosts.v1.Host.description:type_name -> google.protobuf.StringValue
	2, // 3: controller.api.resources.hosts.v1.Host.created_time:type_name -> google.protobuf.Timestamp
	2, // 4: controller.api.resources.hosts.v1.Host.updated_time:type_name -> google.protobuf.Timestamp
	3, // 5: controller.api.resources.hosts.v1.Host.disabled:type_name -> google.protobuf.BoolValue
	1, // 6: controller.api.resources.hosts.v1.Host.address:type_name -> google.protobuf.StringValue
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_controller_api_resources_hosts_v1_host_proto_init() }
func file_controller_api_resources_hosts_v1_host_proto_init() {
	if File_controller_api_resources_hosts_v1_host_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_api_resources_hosts_v1_host_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host); i {
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
			RawDescriptor: file_controller_api_resources_hosts_v1_host_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_api_resources_hosts_v1_host_proto_goTypes,
		DependencyIndexes: file_controller_api_resources_hosts_v1_host_proto_depIdxs,
		MessageInfos:      file_controller_api_resources_hosts_v1_host_proto_msgTypes,
	}.Build()
	File_controller_api_resources_hosts_v1_host_proto = out.File
	file_controller_api_resources_hosts_v1_host_proto_rawDesc = nil
	file_controller_api_resources_hosts_v1_host_proto_goTypes = nil
	file_controller_api_resources_hosts_v1_host_proto_depIdxs = nil
}
