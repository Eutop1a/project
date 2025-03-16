// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.26.1
// source: project/v1/admin/admin.proto

package adminv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common "helloworld/api/project/v1/common"
	user "helloworld/api/project/v1/user"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetUserListReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int32                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserListReq) Reset() {
	*x = GetUserListReq{}
	mi := &file_project_v1_admin_admin_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListReq) ProtoMessage() {}

func (x *GetUserListReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_admin_admin_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListReq.ProtoReflect.Descriptor instead.
func (*GetUserListReq) Descriptor() ([]byte, []int) {
	return file_project_v1_admin_admin_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserListReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetUserListReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetUserListResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CommonResp    *common.CommonResp     `protobuf:"bytes,1,opt,name=common_resp,json=commonResp,proto3" json:"common_resp,omitempty"`
	Users         []*user.UserInfo       `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Total         int32                  `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserListResp) Reset() {
	*x = GetUserListResp{}
	mi := &file_project_v1_admin_admin_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListResp) ProtoMessage() {}

func (x *GetUserListResp) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_admin_admin_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListResp.ProtoReflect.Descriptor instead.
func (*GetUserListResp) Descriptor() ([]byte, []int) {
	return file_project_v1_admin_admin_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserListResp) GetCommonResp() *common.CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

func (x *GetUserListResp) GetUsers() []*user.UserInfo {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *GetUserListResp) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DeleteUserReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteUserReq) Reset() {
	*x = DeleteUserReq{}
	mi := &file_project_v1_admin_admin_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserReq) ProtoMessage() {}

func (x *DeleteUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_admin_admin_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserReq.ProtoReflect.Descriptor instead.
func (*DeleteUserReq) Descriptor() ([]byte, []int) {
	return file_project_v1_admin_admin_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteUserReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type AddAdminReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddAdminReq) Reset() {
	*x = AddAdminReq{}
	mi := &file_project_v1_admin_admin_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAdminReq) ProtoMessage() {}

func (x *AddAdminReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_admin_admin_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAdminReq.ProtoReflect.Descriptor instead.
func (*AddAdminReq) Descriptor() ([]byte, []int) {
	return file_project_v1_admin_admin_proto_rawDescGZIP(), []int{3}
}

func (x *AddAdminReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type RemoveAdminReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveAdminReq) Reset() {
	*x = RemoveAdminReq{}
	mi := &file_project_v1_admin_admin_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAdminReq) ProtoMessage() {}

func (x *RemoveAdminReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_admin_admin_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAdminReq.ProtoReflect.Descriptor instead.
func (*RemoveAdminReq) Descriptor() ([]byte, []int) {
	return file_project_v1_admin_admin_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveAdminReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type BatchRemoveAdminReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserIdList    []string               `protobuf:"bytes,1,rep,name=user_id_list,json=userIdList,proto3" json:"user_id_list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchRemoveAdminReq) Reset() {
	*x = BatchRemoveAdminReq{}
	mi := &file_project_v1_admin_admin_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchRemoveAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchRemoveAdminReq) ProtoMessage() {}

func (x *BatchRemoveAdminReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_admin_admin_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchRemoveAdminReq.ProtoReflect.Descriptor instead.
func (*BatchRemoveAdminReq) Descriptor() ([]byte, []int) {
	return file_project_v1_admin_admin_proto_rawDescGZIP(), []int{5}
}

func (x *BatchRemoveAdminReq) GetUserIdList() []string {
	if x != nil {
		return x.UserIdList
	}
	return nil
}

var File_project_v1_admin_admin_proto protoreflect.FileDescriptor

var file_project_v1_admin_admin_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x98, 0x01,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x3e, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x2f, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x28, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x26, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x13, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xd9,
	0x04, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x6f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x72, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x2a,
	0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x6c, 0x0a,
	0x08, 0x41, 0x64, 0x64, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x64,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a,
	0x01, 0x2a, 0x22, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x61, 0x64, 0x64, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x75, 0x0a, 0x0b, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x25, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x2d, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x85, 0x01, 0x0a, 0x10, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2d, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x51, 0x0a, 0x16, 0x64, 0x65,
	0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x56, 0x31, 0x50, 0x01, 0x5a, 0x27, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_project_v1_admin_admin_proto_rawDescOnce sync.Once
	file_project_v1_admin_admin_proto_rawDescData []byte
)

func file_project_v1_admin_admin_proto_rawDescGZIP() []byte {
	file_project_v1_admin_admin_proto_rawDescOnce.Do(func() {
		file_project_v1_admin_admin_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_v1_admin_admin_proto_rawDesc), len(file_project_v1_admin_admin_proto_rawDesc)))
	})
	return file_project_v1_admin_admin_proto_rawDescData
}

var file_project_v1_admin_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_project_v1_admin_admin_proto_goTypes = []any{
	(*GetUserListReq)(nil),      // 0: project.v1.admin.GetUserListReq
	(*GetUserListResp)(nil),     // 1: project.v1.admin.GetUserListResp
	(*DeleteUserReq)(nil),       // 2: project.v1.admin.DeleteUserReq
	(*AddAdminReq)(nil),         // 3: project.v1.admin.AddAdminReq
	(*RemoveAdminReq)(nil),      // 4: project.v1.admin.RemoveAdminReq
	(*BatchRemoveAdminReq)(nil), // 5: project.v1.admin.BatchRemoveAdminReq
	(*common.CommonResp)(nil),   // 6: project.v1.common.CommonResp
	(*user.UserInfo)(nil),       // 7: project.v1.user.UserInfo
}
var file_project_v1_admin_admin_proto_depIdxs = []int32{
	6, // 0: project.v1.admin.GetUserListResp.common_resp:type_name -> project.v1.common.CommonResp
	7, // 1: project.v1.admin.GetUserListResp.users:type_name -> project.v1.user.UserInfo
	0, // 2: project.v1.admin.Admin.GetUserList:input_type -> project.v1.admin.GetUserListReq
	2, // 3: project.v1.admin.Admin.DeleteUser:input_type -> project.v1.admin.DeleteUserReq
	3, // 4: project.v1.admin.Admin.AddAdmin:input_type -> project.v1.admin.AddAdminReq
	4, // 5: project.v1.admin.Admin.RemoveAdmin:input_type -> project.v1.admin.RemoveAdminReq
	5, // 6: project.v1.admin.Admin.BatchRemoveAdmin:input_type -> project.v1.admin.BatchRemoveAdminReq
	1, // 7: project.v1.admin.Admin.GetUserList:output_type -> project.v1.admin.GetUserListResp
	6, // 8: project.v1.admin.Admin.DeleteUser:output_type -> project.v1.common.CommonResp
	6, // 9: project.v1.admin.Admin.AddAdmin:output_type -> project.v1.common.CommonResp
	6, // 10: project.v1.admin.Admin.RemoveAdmin:output_type -> project.v1.common.CommonResp
	6, // 11: project.v1.admin.Admin.BatchRemoveAdmin:output_type -> project.v1.common.CommonResp
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_project_v1_admin_admin_proto_init() }
func file_project_v1_admin_admin_proto_init() {
	if File_project_v1_admin_admin_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_v1_admin_admin_proto_rawDesc), len(file_project_v1_admin_admin_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_project_v1_admin_admin_proto_goTypes,
		DependencyIndexes: file_project_v1_admin_admin_proto_depIdxs,
		MessageInfos:      file_project_v1_admin_admin_proto_msgTypes,
	}.Build()
	File_project_v1_admin_admin_proto = out.File
	file_project_v1_admin_admin_proto_goTypes = nil
	file_project_v1_admin_admin_proto_depIdxs = nil
}
