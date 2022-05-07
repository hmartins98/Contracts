// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: github.com/hmartins98/Contracts/Products/Products.proto

package Products

import (
	CustomTypes "github.com/hmartins98/Contracts/CustomTypes"
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

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TypeId      int64   `protobuf:"varint,4,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	Price       float32 `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	Description string  `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	ReviewScore float32 `protobuf:"fixed32,7,opt,name=review_score,json=reviewScore,proto3" json:"review_score,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_hmartins98_Contracts_Products_Products_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_hmartins98_Contracts_Products_Products_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_github_com_hmartins98_Contracts_Products_Products_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetTypeId() int64 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *Product) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetReviewScore() float32 {
	if x != nil {
		return x.ReviewScore
	}
	return 0
}

type ProductId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProductId) Reset() {
	*x = ProductId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_hmartins98_Contracts_Products_Products_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductId) ProtoMessage() {}

func (x *ProductId) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_hmartins98_Contracts_Products_Products_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductId.ProtoReflect.Descriptor instead.
func (*ProductId) Descriptor() ([]byte, []int) {
	return file_github_com_hmartins98_Contracts_Products_Products_proto_rawDescGZIP(), []int{1}
}

func (x *ProductId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_github_com_hmartins98_Contracts_Products_Products_proto protoreflect.FileDescriptor

var file_github_com_hmartins98_Contracts_Products_Products_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6d, 0x61,
	0x72, 0x74, 0x69, 0x6e, 0x73, 0x39, 0x38, 0x2f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x3d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6d, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x73, 0x39,
	0x38, 0x2f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0b, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x1b, 0x0a,
	0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xef, 0x01, 0x0a, 0x10, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12,
	0x30, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x18, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x05, 0x2e, 0x42, 0x4f, 0x4f,
	0x4c, 0x12, 0x43, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x1a, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x18, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x30, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x18, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x1a, 0x05, 0x2e, 0x42, 0x4f, 0x4f, 0x4c, 0x12, 0x32, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x1a, 0x05, 0x2e, 0x42, 0x4f, 0x4f, 0x4c, 0x42, 0x14, 0x5a, 0x12,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_hmartins98_Contracts_Products_Products_proto_rawDescOnce sync.Once
	file_github_com_hmartins98_Contracts_Products_Products_proto_rawDescData = file_github_com_hmartins98_Contracts_Products_Products_proto_rawDesc
)

func file_github_com_hmartins98_Contracts_Products_Products_proto_rawDescGZIP() []byte {
	file_github_com_hmartins98_Contracts_Products_Products_proto_rawDescOnce.Do(func() {
		file_github_com_hmartins98_Contracts_Products_Products_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_hmartins98_Contracts_Products_Products_proto_rawDescData)
	})
	return file_github_com_hmartins98_Contracts_Products_Products_proto_rawDescData
}

var file_github_com_hmartins98_Contracts_Products_Products_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_hmartins98_Contracts_Products_Products_proto_goTypes = []interface{}{
	(*Product)(nil),          // 0: ProductsPackage.Product
	(*ProductId)(nil),        // 1: ProductsPackage.ProductId
	(*CustomTypes.BOOL)(nil), // 2: BOOL
}
var file_github_com_hmartins98_Contracts_Products_Products_proto_depIdxs = []int32{
	0, // 0: ProductsPackage.ProductsContract.CreateProduct:input_type -> ProductsPackage.Product
	1, // 1: ProductsPackage.ProductsContract.ReadProduct:input_type -> ProductsPackage.ProductId
	0, // 2: ProductsPackage.ProductsContract.UpdateProduct:input_type -> ProductsPackage.Product
	1, // 3: ProductsPackage.ProductsContract.DeleteProduct:input_type -> ProductsPackage.ProductId
	2, // 4: ProductsPackage.ProductsContract.CreateProduct:output_type -> BOOL
	0, // 5: ProductsPackage.ProductsContract.ReadProduct:output_type -> ProductsPackage.Product
	2, // 6: ProductsPackage.ProductsContract.UpdateProduct:output_type -> BOOL
	2, // 7: ProductsPackage.ProductsContract.DeleteProduct:output_type -> BOOL
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_hmartins98_Contracts_Products_Products_proto_init() }
func file_github_com_hmartins98_Contracts_Products_Products_proto_init() {
	if File_github_com_hmartins98_Contracts_Products_Products_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_hmartins98_Contracts_Products_Products_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_github_com_hmartins98_Contracts_Products_Products_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductId); i {
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
			RawDescriptor: file_github_com_hmartins98_Contracts_Products_Products_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_hmartins98_Contracts_Products_Products_proto_goTypes,
		DependencyIndexes: file_github_com_hmartins98_Contracts_Products_Products_proto_depIdxs,
		MessageInfos:      file_github_com_hmartins98_Contracts_Products_Products_proto_msgTypes,
	}.Build()
	File_github_com_hmartins98_Contracts_Products_Products_proto = out.File
	file_github_com_hmartins98_Contracts_Products_Products_proto_rawDesc = nil
	file_github_com_hmartins98_Contracts_Products_Products_proto_goTypes = nil
	file_github_com_hmartins98_Contracts_Products_Products_proto_depIdxs = nil
}
