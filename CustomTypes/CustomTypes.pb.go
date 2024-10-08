// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: github.com/hmartins98/Contracts/CustomTypes/CustomTypes.proto

package CustomTypes

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DECIMAL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Units int64 `protobuf:"varint,1,opt,name=units,proto3" json:"units,omitempty"`
	Nanos int32 `protobuf:"fixed32,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *DECIMAL) Reset() {
	*x = DECIMAL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DECIMAL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DECIMAL) ProtoMessage() {}

func (x *DECIMAL) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DECIMAL.ProtoReflect.Descriptor instead.
func (*DECIMAL) Descriptor() ([]byte, []int) {
	return file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_rawDescGZIP(), []int{0}
}

func (x *DECIMAL) GetUnits() int64 {
	if x != nil {
		return x.Units
	}
	return 0
}

func (x *DECIMAL) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

type GUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GUID) Reset() {
	*x = GUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GUID) ProtoMessage() {}

func (x *GUID) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GUID.ProtoReflect.Descriptor instead.
func (*GUID) Descriptor() ([]byte, []int) {
	return file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_rawDescGZIP(), []int{1}
}

func (x *GUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type DATETIME struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DATETIME) Reset() {
	*x = DATETIME{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DATETIME) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DATETIME) ProtoMessage() {}

func (x *DATETIME) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DATETIME.ProtoReflect.Descriptor instead.
func (*DATETIME) Descriptor() ([]byte, []int) {
	return file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_rawDescGZIP(), []int{2}
}

func (x *DATETIME) GetValue() *timestamppb.Timestamp {
	if x != nil {
		return x.Value
	}
	return nil
}

type BOOL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BOOL) Reset() {
	*x = BOOL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BOOL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BOOL) ProtoMessage() {}

func (x *BOOL) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BOOL.ProtoReflect.Descriptor instead.
func (*BOOL) Descriptor() ([]byte, []int) {
	return file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_rawDescGZIP(), []int{3}
}

func (x *BOOL) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

var File_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto protoreflect.FileDescriptor

var file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6d, 0x61,
	0x72, 0x74, 0x69, 0x6e, 0x73, 0x39, 0x38, 0x2f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x35, 0x0a, 0x07, 0x44, 0x45, 0x43, 0x49, 0x4d, 0x41, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x75,
	0x6e, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0f,
	0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x22, 0x1c, 0x0a, 0x04, 0x47, 0x55, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3c, 0x0a, 0x08, 0x44, 0x41, 0x54, 0x45, 0x54, 0x49, 0x4d,
	0x45, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x1c, 0x0a, 0x04, 0x42, 0x4f, 0x4f, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x6d, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x73, 0x39, 0x38, 0x2f, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x73, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_rawDescOnce sync.Once
	file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_rawDescData = file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_rawDesc
)

func file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_rawDescGZIP() []byte {
	file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_rawDescOnce.Do(func() {
		file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_rawDescData)
	})
	return file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_rawDescData
}

var file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_goTypes = []interface{}{
	(*DECIMAL)(nil),               // 0: DECIMAL
	(*GUID)(nil),                  // 1: GUID
	(*DATETIME)(nil),              // 2: DATETIME
	(*BOOL)(nil),                  // 3: BOOL
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_depIdxs = []int32{
	4, // 0: DATETIME.value:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_init() }
func file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_init() {
	if File_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DECIMAL); i {
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
		file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GUID); i {
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
		file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DATETIME); i {
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
		file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BOOL); i {
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
			RawDescriptor: file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_goTypes,
		DependencyIndexes: file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_depIdxs,
		MessageInfos:      file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_msgTypes,
	}.Build()
	File_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto = out.File
	file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_rawDesc = nil
	file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_goTypes = nil
	file_github_com_hmartins98_Contracts_CustomTypes_CustomTypes_proto_depIdxs = nil
}
