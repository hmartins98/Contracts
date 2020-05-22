// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: github.com/hmartins98/Contracts/Auth/Auth.proto

package Auth

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	CustomTypes "github.com/hmartins98/Contracts/CustomTypes"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AuthenticationLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthenticationLoginRequest) Reset() {
	*x = AuthenticationLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_hmartins98_Contracts_Auth_Auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationLoginRequest) ProtoMessage() {}

func (x *AuthenticationLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_hmartins98_Contracts_Auth_Auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationLoginRequest.ProtoReflect.Descriptor instead.
func (*AuthenticationLoginRequest) Descriptor() ([]byte, []int) {
	return file_github_com_hmartins98_Contracts_Auth_Auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticationLoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthenticationLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthenticationTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidationToken *CustomTypes.GUID `protobuf:"bytes,1,opt,name=validation_token,json=validationToken,proto3" json:"validation_token,omitempty"`
}

func (x *AuthenticationTokenRequest) Reset() {
	*x = AuthenticationTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_hmartins98_Contracts_Auth_Auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationTokenRequest) ProtoMessage() {}

func (x *AuthenticationTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_hmartins98_Contracts_Auth_Auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationTokenRequest.ProtoReflect.Descriptor instead.
func (*AuthenticationTokenRequest) Descriptor() ([]byte, []int) {
	return file_github_com_hmartins98_Contracts_Auth_Auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticationTokenRequest) GetValidationToken() *CustomTypes.GUID {
	if x != nil {
		return x.ValidationToken
	}
	return nil
}

type CreateAuthenticationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateAuthenticationRequest) Reset() {
	*x = CreateAuthenticationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_hmartins98_Contracts_Auth_Auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuthenticationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthenticationRequest) ProtoMessage() {}

func (x *CreateAuthenticationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_hmartins98_Contracts_Auth_Auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthenticationRequest.ProtoReflect.Descriptor instead.
func (*CreateAuthenticationRequest) Descriptor() ([]byte, []int) {
	return file_github_com_hmartins98_Contracts_Auth_Auth_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAuthenticationRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateAuthenticationRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthenticationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          *CustomTypes.GUID `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ValidationToken *CustomTypes.GUID `protobuf:"bytes,2,opt,name=validation_token,json=validationToken,proto3" json:"validation_token,omitempty"`
	SessionKey      *CustomTypes.GUID `protobuf:"bytes,4,opt,name=session_key,json=sessionKey,proto3" json:"session_key,omitempty"`
}

func (x *AuthenticationReply) Reset() {
	*x = AuthenticationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_hmartins98_Contracts_Auth_Auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationReply) ProtoMessage() {}

func (x *AuthenticationReply) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_hmartins98_Contracts_Auth_Auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationReply.ProtoReflect.Descriptor instead.
func (*AuthenticationReply) Descriptor() ([]byte, []int) {
	return file_github_com_hmartins98_Contracts_Auth_Auth_proto_rawDescGZIP(), []int{3}
}

func (x *AuthenticationReply) GetUserId() *CustomTypes.GUID {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *AuthenticationReply) GetValidationToken() *CustomTypes.GUID {
	if x != nil {
		return x.ValidationToken
	}
	return nil
}

func (x *AuthenticationReply) GetSessionKey() *CustomTypes.GUID {
	if x != nil {
		return x.SessionKey
	}
	return nil
}

var File_github_com_hmartins98_Contracts_Auth_Auth_proto protoreflect.FileDescriptor

var file_github_com_hmartins98_Contracts_Auth_Auth_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6d, 0x61,
	0x72, 0x74, 0x69, 0x6e, 0x73, 0x39, 0x38, 0x2f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x2f, 0x41, 0x75, 0x74, 0x68, 0x2f, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x3d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6d, 0x61, 0x72, 0x74,
	0x69, 0x6e, 0x73, 0x39, 0x38, 0x2f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a,
	0x1a, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x4e, 0x0a, 0x1a, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x30, 0x0a, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x47, 0x55,
	0x49, 0x44, 0x52, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x55, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x13, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x1e, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x47, 0x55, 0x49, 0x44, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x30, 0x0a, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x47,
	0x55, 0x49, 0x44, 0x52, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x0a, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x47, 0x55, 0x49, 0x44,
	0x52, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x32, 0xb4, 0x02, 0x0a,
	0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x5e, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x27, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x5e, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x27, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x62, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x68, 0x6d, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x73, 0x39, 0x38, 0x2f, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x41, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_hmartins98_Contracts_Auth_Auth_proto_rawDescOnce sync.Once
	file_github_com_hmartins98_Contracts_Auth_Auth_proto_rawDescData = file_github_com_hmartins98_Contracts_Auth_Auth_proto_rawDesc
)

func file_github_com_hmartins98_Contracts_Auth_Auth_proto_rawDescGZIP() []byte {
	file_github_com_hmartins98_Contracts_Auth_Auth_proto_rawDescOnce.Do(func() {
		file_github_com_hmartins98_Contracts_Auth_Auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_hmartins98_Contracts_Auth_Auth_proto_rawDescData)
	})
	return file_github_com_hmartins98_Contracts_Auth_Auth_proto_rawDescData
}

var file_github_com_hmartins98_Contracts_Auth_Auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_hmartins98_Contracts_Auth_Auth_proto_goTypes = []interface{}{
	(*AuthenticationLoginRequest)(nil),  // 0: AuthPackage.AuthenticationLoginRequest
	(*AuthenticationTokenRequest)(nil),  // 1: AuthPackage.AuthenticationTokenRequest
	(*CreateAuthenticationRequest)(nil), // 2: AuthPackage.CreateAuthenticationRequest
	(*AuthenticationReply)(nil),         // 3: AuthPackage.AuthenticationReply
	(*CustomTypes.GUID)(nil),            // 4: GUID
}
var file_github_com_hmartins98_Contracts_Auth_Auth_proto_depIdxs = []int32{
	4, // 0: AuthPackage.AuthenticationTokenRequest.validation_token:type_name -> GUID
	4, // 1: AuthPackage.AuthenticationReply.user_id:type_name -> GUID
	4, // 2: AuthPackage.AuthenticationReply.validation_token:type_name -> GUID
	4, // 3: AuthPackage.AuthenticationReply.session_key:type_name -> GUID
	0, // 4: AuthPackage.Authentication.AuthenticateLogin:input_type -> AuthPackage.AuthenticationLoginRequest
	1, // 5: AuthPackage.Authentication.AuthenticateToken:input_type -> AuthPackage.AuthenticationTokenRequest
	2, // 6: AuthPackage.Authentication.CreateAuthentication:input_type -> AuthPackage.CreateAuthenticationRequest
	3, // 7: AuthPackage.Authentication.AuthenticateLogin:output_type -> AuthPackage.AuthenticationReply
	3, // 8: AuthPackage.Authentication.AuthenticateToken:output_type -> AuthPackage.AuthenticationReply
	3, // 9: AuthPackage.Authentication.CreateAuthentication:output_type -> AuthPackage.AuthenticationReply
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_hmartins98_Contracts_Auth_Auth_proto_init() }
func file_github_com_hmartins98_Contracts_Auth_Auth_proto_init() {
	if File_github_com_hmartins98_Contracts_Auth_Auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_hmartins98_Contracts_Auth_Auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationLoginRequest); i {
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
		file_github_com_hmartins98_Contracts_Auth_Auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationTokenRequest); i {
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
		file_github_com_hmartins98_Contracts_Auth_Auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuthenticationRequest); i {
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
		file_github_com_hmartins98_Contracts_Auth_Auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationReply); i {
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
			RawDescriptor: file_github_com_hmartins98_Contracts_Auth_Auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_hmartins98_Contracts_Auth_Auth_proto_goTypes,
		DependencyIndexes: file_github_com_hmartins98_Contracts_Auth_Auth_proto_depIdxs,
		MessageInfos:      file_github_com_hmartins98_Contracts_Auth_Auth_proto_msgTypes,
	}.Build()
	File_github_com_hmartins98_Contracts_Auth_Auth_proto = out.File
	file_github_com_hmartins98_Contracts_Auth_Auth_proto_rawDesc = nil
	file_github_com_hmartins98_Contracts_Auth_Auth_proto_goTypes = nil
	file_github_com_hmartins98_Contracts_Auth_Auth_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationClient interface {
	AuthenticateLogin(ctx context.Context, in *AuthenticationLoginRequest, opts ...grpc.CallOption) (*AuthenticationReply, error)
	AuthenticateToken(ctx context.Context, in *AuthenticationTokenRequest, opts ...grpc.CallOption) (*AuthenticationReply, error)
	CreateAuthentication(ctx context.Context, in *CreateAuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationReply, error)
}

type authenticationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationClient(cc grpc.ClientConnInterface) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) AuthenticateLogin(ctx context.Context, in *AuthenticationLoginRequest, opts ...grpc.CallOption) (*AuthenticationReply, error) {
	out := new(AuthenticationReply)
	err := c.cc.Invoke(ctx, "/AuthPackage.Authentication/AuthenticateLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) AuthenticateToken(ctx context.Context, in *AuthenticationTokenRequest, opts ...grpc.CallOption) (*AuthenticationReply, error) {
	out := new(AuthenticationReply)
	err := c.cc.Invoke(ctx, "/AuthPackage.Authentication/AuthenticateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) CreateAuthentication(ctx context.Context, in *CreateAuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationReply, error) {
	out := new(AuthenticationReply)
	err := c.cc.Invoke(ctx, "/AuthPackage.Authentication/CreateAuthentication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
type AuthenticationServer interface {
	AuthenticateLogin(context.Context, *AuthenticationLoginRequest) (*AuthenticationReply, error)
	AuthenticateToken(context.Context, *AuthenticationTokenRequest) (*AuthenticationReply, error)
	CreateAuthentication(context.Context, *CreateAuthenticationRequest) (*AuthenticationReply, error)
}

// UnimplementedAuthenticationServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServer struct {
}

func (*UnimplementedAuthenticationServer) AuthenticateLogin(context.Context, *AuthenticationLoginRequest) (*AuthenticationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticateLogin not implemented")
}
func (*UnimplementedAuthenticationServer) AuthenticateToken(context.Context, *AuthenticationTokenRequest) (*AuthenticationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticateToken not implemented")
}
func (*UnimplementedAuthenticationServer) CreateAuthentication(context.Context, *CreateAuthenticationRequest) (*AuthenticationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuthentication not implemented")
}

func RegisterAuthenticationServer(s *grpc.Server, srv AuthenticationServer) {
	s.RegisterService(&_Authentication_serviceDesc, srv)
}

func _Authentication_AuthenticateLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).AuthenticateLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthPackage.Authentication/AuthenticateLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).AuthenticateLogin(ctx, req.(*AuthenticationLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_AuthenticateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).AuthenticateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthPackage.Authentication/AuthenticateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).AuthenticateToken(ctx, req.(*AuthenticationTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_CreateAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).CreateAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthPackage.Authentication/CreateAuthentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).CreateAuthentication(ctx, req.(*CreateAuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AuthPackage.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthenticateLogin",
			Handler:    _Authentication_AuthenticateLogin_Handler,
		},
		{
			MethodName: "AuthenticateToken",
			Handler:    _Authentication_AuthenticateToken_Handler,
		},
		{
			MethodName: "CreateAuthentication",
			Handler:    _Authentication_CreateAuthentication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/hmartins98/Contracts/Auth/Auth.proto",
}
