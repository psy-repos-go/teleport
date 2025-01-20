// Copyright 2022 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: teleport/users/v1/users_service.proto

package usersv1

import (
	types "github.com/gravitational/teleport/api/types"
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

// Request for GetUser.
type GetUserRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name of the user to retrieve, this take priority over current_user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Whether to return the current user. If the name is provided
	// then this field is ignored.
	CurrentUser bool `protobuf:"varint,2,opt,name=current_user,json=currentUser,proto3" json:"current_user,omitempty"`
	// Specifies whether to load associated secrets(password, mfa devices, etc.).
	WithSecrets   bool `protobuf:"varint,3,opt,name=with_secrets,json=withSecrets,proto3" json:"with_secrets,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_teleport_users_v1_users_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetUserRequest) GetCurrentUser() bool {
	if x != nil {
		return x.CurrentUser
	}
	return false
}

func (x *GetUserRequest) GetWithSecrets() bool {
	if x != nil {
		return x.WithSecrets
	}
	return false
}

// Response for GetUser.
type GetUserResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The user matching the request filters.
	User          *types.UserV2 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_teleport_users_v1_users_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserResponse) GetUser() *types.UserV2 {
	if x != nil {
		return x.User
	}
	return nil
}

// Request for ListUsers.
//
// Follows the pagination semantics of
// https://cloud.google.com/apis/design/standard_methods#list.
type ListUsersRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The maximum number of items to return.
	// The server may impose a different page size at its discretion.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Specifies whether to load associated secrets(password, mfa devices, etc.).
	WithSecrets bool `protobuf:"varint,3,opt,name=with_secrets,json=withSecrets,proto3" json:"with_secrets,omitempty"`
	// Filter matches users.
	Filter        *types.UserFilter `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUsersRequest) Reset() {
	*x = ListUsersRequest{}
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersRequest) ProtoMessage() {}

func (x *ListUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersRequest.ProtoReflect.Descriptor instead.
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return file_teleport_users_v1_users_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListUsersRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListUsersRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListUsersRequest) GetWithSecrets() bool {
	if x != nil {
		return x.WithSecrets
	}
	return false
}

func (x *ListUsersRequest) GetFilter() *types.UserFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

// Response for ListUsers.
type ListUsersResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Users that matched the search.
	Users []*types.UserV2 `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUsersResponse) Reset() {
	*x = ListUsersResponse{}
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersResponse) ProtoMessage() {}

func (x *ListUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersResponse.ProtoReflect.Descriptor instead.
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return file_teleport_users_v1_users_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListUsersResponse) GetUsers() []*types.UserV2 {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *ListUsersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Request for CreateUser.
type CreateUserRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The user resource to create.
	User          *types.UserV2 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_teleport_users_v1_users_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateUserRequest) GetUser() *types.UserV2 {
	if x != nil {
		return x.User
	}
	return nil
}

// Response for CreateUser.
type CreateUserResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The created user with any server side generated fields populated.
	User          *types.UserV2 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_teleport_users_v1_users_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserResponse) GetUser() *types.UserV2 {
	if x != nil {
		return x.User
	}
	return nil
}

// Request for UpdateUser.
type UpdateUserRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The user resource to update.
	User          *types.UserV2 `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_teleport_users_v1_users_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUserRequest) GetUser() *types.UserV2 {
	if x != nil {
		return x.User
	}
	return nil
}

// Response for UpdateUser.
type UpdateUserResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The updated user with any server side generated fields populated.
	User          *types.UserV2 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserResponse) Reset() {
	*x = UpdateUserResponse{}
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserResponse) ProtoMessage() {}

func (x *UpdateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return file_teleport_users_v1_users_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateUserResponse) GetUser() *types.UserV2 {
	if x != nil {
		return x.User
	}
	return nil
}

// Request for UpsertUser.
type UpsertUserRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The user resource to upsert.
	User          *types.UserV2 `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertUserRequest) Reset() {
	*x = UpsertUserRequest{}
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertUserRequest) ProtoMessage() {}

func (x *UpsertUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertUserRequest.ProtoReflect.Descriptor instead.
func (*UpsertUserRequest) Descriptor() ([]byte, []int) {
	return file_teleport_users_v1_users_service_proto_rawDescGZIP(), []int{8}
}

func (x *UpsertUserRequest) GetUser() *types.UserV2 {
	if x != nil {
		return x.User
	}
	return nil
}

// Response for UpsertUser.
type UpsertUserResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The upserted user with any server side generated fields populated.
	User          *types.UserV2 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertUserResponse) Reset() {
	*x = UpsertUserResponse{}
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertUserResponse) ProtoMessage() {}

func (x *UpsertUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertUserResponse.ProtoReflect.Descriptor instead.
func (*UpsertUserResponse) Descriptor() ([]byte, []int) {
	return file_teleport_users_v1_users_service_proto_rawDescGZIP(), []int{9}
}

func (x *UpsertUserResponse) GetUser() *types.UserV2 {
	if x != nil {
		return x.User
	}
	return nil
}

// Request for DeleteUser.
type DeleteUserRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name of the user to remove.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_users_v1_users_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_teleport_users_v1_users_service_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_teleport_users_v1_users_service_proto protoreflect.FileDescriptor

var file_teleport_users_v1_users_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x77, 0x69, 0x74, 0x68, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x22, 0x34, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x56, 0x32, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x9c, 0x01, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a,
	0x0c, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x77, 0x69, 0x74, 0x68, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x12, 0x29, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x60, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x23, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x32, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x36, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x32, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x37, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x32, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x36,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x32,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x37, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x32, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x36, 0x0a, 0x11, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x56,
	0x32, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x37, 0x0a, 0x12, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x32, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x27, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x95, 0x04, 0x0a, 0x0c, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x09,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x23, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x24, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x59, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x24, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0a, 0x55, 0x70,
	0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_users_v1_users_service_proto_rawDescOnce sync.Once
	file_teleport_users_v1_users_service_proto_rawDescData = file_teleport_users_v1_users_service_proto_rawDesc
)

func file_teleport_users_v1_users_service_proto_rawDescGZIP() []byte {
	file_teleport_users_v1_users_service_proto_rawDescOnce.Do(func() {
		file_teleport_users_v1_users_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_users_v1_users_service_proto_rawDescData)
	})
	return file_teleport_users_v1_users_service_proto_rawDescData
}

var file_teleport_users_v1_users_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_teleport_users_v1_users_service_proto_goTypes = []any{
	(*GetUserRequest)(nil),     // 0: teleport.users.v1.GetUserRequest
	(*GetUserResponse)(nil),    // 1: teleport.users.v1.GetUserResponse
	(*ListUsersRequest)(nil),   // 2: teleport.users.v1.ListUsersRequest
	(*ListUsersResponse)(nil),  // 3: teleport.users.v1.ListUsersResponse
	(*CreateUserRequest)(nil),  // 4: teleport.users.v1.CreateUserRequest
	(*CreateUserResponse)(nil), // 5: teleport.users.v1.CreateUserResponse
	(*UpdateUserRequest)(nil),  // 6: teleport.users.v1.UpdateUserRequest
	(*UpdateUserResponse)(nil), // 7: teleport.users.v1.UpdateUserResponse
	(*UpsertUserRequest)(nil),  // 8: teleport.users.v1.UpsertUserRequest
	(*UpsertUserResponse)(nil), // 9: teleport.users.v1.UpsertUserResponse
	(*DeleteUserRequest)(nil),  // 10: teleport.users.v1.DeleteUserRequest
	(*types.UserV2)(nil),       // 11: types.UserV2
	(*types.UserFilter)(nil),   // 12: types.UserFilter
	(*emptypb.Empty)(nil),      // 13: google.protobuf.Empty
}
var file_teleport_users_v1_users_service_proto_depIdxs = []int32{
	11, // 0: teleport.users.v1.GetUserResponse.user:type_name -> types.UserV2
	12, // 1: teleport.users.v1.ListUsersRequest.filter:type_name -> types.UserFilter
	11, // 2: teleport.users.v1.ListUsersResponse.users:type_name -> types.UserV2
	11, // 3: teleport.users.v1.CreateUserRequest.user:type_name -> types.UserV2
	11, // 4: teleport.users.v1.CreateUserResponse.user:type_name -> types.UserV2
	11, // 5: teleport.users.v1.UpdateUserRequest.user:type_name -> types.UserV2
	11, // 6: teleport.users.v1.UpdateUserResponse.user:type_name -> types.UserV2
	11, // 7: teleport.users.v1.UpsertUserRequest.user:type_name -> types.UserV2
	11, // 8: teleport.users.v1.UpsertUserResponse.user:type_name -> types.UserV2
	0,  // 9: teleport.users.v1.UsersService.GetUser:input_type -> teleport.users.v1.GetUserRequest
	2,  // 10: teleport.users.v1.UsersService.ListUsers:input_type -> teleport.users.v1.ListUsersRequest
	4,  // 11: teleport.users.v1.UsersService.CreateUser:input_type -> teleport.users.v1.CreateUserRequest
	6,  // 12: teleport.users.v1.UsersService.UpdateUser:input_type -> teleport.users.v1.UpdateUserRequest
	8,  // 13: teleport.users.v1.UsersService.UpsertUser:input_type -> teleport.users.v1.UpsertUserRequest
	10, // 14: teleport.users.v1.UsersService.DeleteUser:input_type -> teleport.users.v1.DeleteUserRequest
	1,  // 15: teleport.users.v1.UsersService.GetUser:output_type -> teleport.users.v1.GetUserResponse
	3,  // 16: teleport.users.v1.UsersService.ListUsers:output_type -> teleport.users.v1.ListUsersResponse
	5,  // 17: teleport.users.v1.UsersService.CreateUser:output_type -> teleport.users.v1.CreateUserResponse
	7,  // 18: teleport.users.v1.UsersService.UpdateUser:output_type -> teleport.users.v1.UpdateUserResponse
	9,  // 19: teleport.users.v1.UsersService.UpsertUser:output_type -> teleport.users.v1.UpsertUserResponse
	13, // 20: teleport.users.v1.UsersService.DeleteUser:output_type -> google.protobuf.Empty
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_teleport_users_v1_users_service_proto_init() }
func file_teleport_users_v1_users_service_proto_init() {
	if File_teleport_users_v1_users_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_users_v1_users_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_users_v1_users_service_proto_goTypes,
		DependencyIndexes: file_teleport_users_v1_users_service_proto_depIdxs,
		MessageInfos:      file_teleport_users_v1_users_service_proto_msgTypes,
	}.Build()
	File_teleport_users_v1_users_service_proto = out.File
	file_teleport_users_v1_users_service_proto_rawDesc = nil
	file_teleport_users_v1_users_service_proto_goTypes = nil
	file_teleport_users_v1_users_service_proto_depIdxs = nil
}
