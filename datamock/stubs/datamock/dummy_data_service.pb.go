// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: dummy_data_service.proto

package datamock

import (
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

type RPCRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RPCRequest) Reset() {
	*x = RPCRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dummy_data_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCRequest) ProtoMessage() {}

func (x *RPCRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dummy_data_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCRequest.ProtoReflect.Descriptor instead.
func (*RPCRequest) Descriptor() ([]byte, []int) {
	return file_dummy_data_service_proto_rawDescGZIP(), []int{0}
}

func (x *RPCRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Class string `protobuf:"bytes,2,opt,name=class,proto3" json:"class,omitempty"`
	Roll  int64  `protobuf:"varint,3,opt,name=roll,proto3" json:"roll,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dummy_data_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_dummy_data_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_dummy_data_service_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

func (x *User) GetRoll() int64 {
	if x != nil {
		return x.Roll
	}
	return 0
}

var File_dummy_data_service_proto protoreflect.FileDescriptor

var file_dummy_data_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x6d, 0x6f, 0x63, 0x6b, 0x22, 0x20, 0x0a, 0x0a, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x44, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x6c, 0x32, 0x4d, 0x0a, 0x10,
	0x44, 0x75, 0x6d, 0x6d, 0x79, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x39, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x14, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x52,
	0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x6d, 0x6f, 0x63, 0x6b, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x6d, 0x6f, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dummy_data_service_proto_rawDescOnce sync.Once
	file_dummy_data_service_proto_rawDescData = file_dummy_data_service_proto_rawDesc
)

func file_dummy_data_service_proto_rawDescGZIP() []byte {
	file_dummy_data_service_proto_rawDescOnce.Do(func() {
		file_dummy_data_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_dummy_data_service_proto_rawDescData)
	})
	return file_dummy_data_service_proto_rawDescData
}

var file_dummy_data_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dummy_data_service_proto_goTypes = []interface{}{
	(*RPCRequest)(nil), // 0: datamock.RPCRequest
	(*User)(nil),       // 1: datamock.User
}
var file_dummy_data_service_proto_depIdxs = []int32{
	0, // 0: datamock.DummyDataService.GetMockUserData:input_type -> datamock.RPCRequest
	1, // 1: datamock.DummyDataService.GetMockUserData:output_type -> datamock.User
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dummy_data_service_proto_init() }
func file_dummy_data_service_proto_init() {
	if File_dummy_data_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dummy_data_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCRequest); i {
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
		file_dummy_data_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_dummy_data_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dummy_data_service_proto_goTypes,
		DependencyIndexes: file_dummy_data_service_proto_depIdxs,
		MessageInfos:      file_dummy_data_service_proto_msgTypes,
	}.Build()
	File_dummy_data_service_proto = out.File
	file_dummy_data_service_proto_rawDesc = nil
	file_dummy_data_service_proto_goTypes = nil
	file_dummy_data_service_proto_depIdxs = nil
}
