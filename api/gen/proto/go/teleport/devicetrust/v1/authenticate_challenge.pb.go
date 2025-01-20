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
// source: teleport/devicetrust/v1/authenticate_challenge.proto

package devicetrustv1

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

// AuthenticateDeviceChallenge carries the authentication challenge.
type AuthenticateDeviceChallenge struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Randomly-generated, opaque challenge to be signed using the device key.
	Challenge     []byte `protobuf:"bytes,1,opt,name=challenge,proto3" json:"challenge,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthenticateDeviceChallenge) Reset() {
	*x = AuthenticateDeviceChallenge{}
	mi := &file_teleport_devicetrust_v1_authenticate_challenge_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthenticateDeviceChallenge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateDeviceChallenge) ProtoMessage() {}

func (x *AuthenticateDeviceChallenge) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_devicetrust_v1_authenticate_challenge_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateDeviceChallenge.ProtoReflect.Descriptor instead.
func (*AuthenticateDeviceChallenge) Descriptor() ([]byte, []int) {
	return file_teleport_devicetrust_v1_authenticate_challenge_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticateDeviceChallenge) GetChallenge() []byte {
	if x != nil {
		return x.Challenge
	}
	return nil
}

// AuthenticateDeviceChallengeResponse carries the authentication challenge
// response.
type AuthenticateDeviceChallengeResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Signature over the challenge, using the device key.
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// Signature over the challenge, using the SSH key. This is required when the
	// SSH and TLS public keys do not match, to prove ownership of the private key
	// associated with the SSH certificate being augmented.
	SshSignature  []byte `protobuf:"bytes,2,opt,name=ssh_signature,json=sshSignature,proto3" json:"ssh_signature,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthenticateDeviceChallengeResponse) Reset() {
	*x = AuthenticateDeviceChallengeResponse{}
	mi := &file_teleport_devicetrust_v1_authenticate_challenge_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthenticateDeviceChallengeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateDeviceChallengeResponse) ProtoMessage() {}

func (x *AuthenticateDeviceChallengeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_devicetrust_v1_authenticate_challenge_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateDeviceChallengeResponse.ProtoReflect.Descriptor instead.
func (*AuthenticateDeviceChallengeResponse) Descriptor() ([]byte, []int) {
	return file_teleport_devicetrust_v1_authenticate_challenge_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticateDeviceChallengeResponse) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *AuthenticateDeviceChallengeResponse) GetSshSignature() []byte {
	if x != nil {
		return x.SshSignature
	}
	return nil
}

// TPMAuthenticateDeviceChallenge carries the authentication challenge
// specific to TPMs.
type TPMAuthenticateDeviceChallenge struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Randomly-generated nonce to be used during platform attestation by the
	// TPM.
	AttestationNonce []byte `protobuf:"bytes,1,opt,name=attestation_nonce,json=attestationNonce,proto3" json:"attestation_nonce,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *TPMAuthenticateDeviceChallenge) Reset() {
	*x = TPMAuthenticateDeviceChallenge{}
	mi := &file_teleport_devicetrust_v1_authenticate_challenge_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TPMAuthenticateDeviceChallenge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TPMAuthenticateDeviceChallenge) ProtoMessage() {}

func (x *TPMAuthenticateDeviceChallenge) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_devicetrust_v1_authenticate_challenge_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TPMAuthenticateDeviceChallenge.ProtoReflect.Descriptor instead.
func (*TPMAuthenticateDeviceChallenge) Descriptor() ([]byte, []int) {
	return file_teleport_devicetrust_v1_authenticate_challenge_proto_rawDescGZIP(), []int{2}
}

func (x *TPMAuthenticateDeviceChallenge) GetAttestationNonce() []byte {
	if x != nil {
		return x.AttestationNonce
	}
	return nil
}

// TPMAuthenticateDeviceChallengeResponse carries the authentication challenge
// response specific to TPMs.
type TPMAuthenticateDeviceChallengeResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The result of the client's platform attestation with the nonce provided
	// in `TPMAuthenticateDeviceChallenge`.
	PlatformParameters *TPMPlatformParameters `protobuf:"bytes,1,opt,name=platform_parameters,json=platformParameters,proto3" json:"platform_parameters,omitempty"`
	// Signature over the attestation_nonce, using the SSH key. This is required
	// when the SSH and TLS public keys do not match, to prove ownership of the
	// private key associated with the SSH certificate being augmented.
	SshSignature  []byte `protobuf:"bytes,2,opt,name=ssh_signature,json=sshSignature,proto3" json:"ssh_signature,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TPMAuthenticateDeviceChallengeResponse) Reset() {
	*x = TPMAuthenticateDeviceChallengeResponse{}
	mi := &file_teleport_devicetrust_v1_authenticate_challenge_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TPMAuthenticateDeviceChallengeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TPMAuthenticateDeviceChallengeResponse) ProtoMessage() {}

func (x *TPMAuthenticateDeviceChallengeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_devicetrust_v1_authenticate_challenge_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TPMAuthenticateDeviceChallengeResponse.ProtoReflect.Descriptor instead.
func (*TPMAuthenticateDeviceChallengeResponse) Descriptor() ([]byte, []int) {
	return file_teleport_devicetrust_v1_authenticate_challenge_proto_rawDescGZIP(), []int{3}
}

func (x *TPMAuthenticateDeviceChallengeResponse) GetPlatformParameters() *TPMPlatformParameters {
	if x != nil {
		return x.PlatformParameters
	}
	return nil
}

func (x *TPMAuthenticateDeviceChallengeResponse) GetSshSignature() []byte {
	if x != nil {
		return x.SshSignature
	}
	return nil
}

var File_teleport_devicetrust_v1_authenticate_challenge_proto protoreflect.FileDescriptor

var file_teleport_devicetrust_v1_authenticate_challenge_proto_rawDesc = []byte{
	0x0a, 0x34, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a,
	0x21, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x70, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x1b, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x22,
	0x68, 0x0a, 0x23, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x73, 0x68, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x73, 0x73, 0x68,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x4d, 0x0a, 0x1e, 0x54, 0x50, 0x4d,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0xae, 0x01, 0x0a, 0x26, 0x54, 0x50, 0x4d,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x13, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x50, 0x4d, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x12, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x73, 0x68, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x73, 0x73, 0x68,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x5a, 0x5a, 0x58, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74,
	0x72, 0x75, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_devicetrust_v1_authenticate_challenge_proto_rawDescOnce sync.Once
	file_teleport_devicetrust_v1_authenticate_challenge_proto_rawDescData = file_teleport_devicetrust_v1_authenticate_challenge_proto_rawDesc
)

func file_teleport_devicetrust_v1_authenticate_challenge_proto_rawDescGZIP() []byte {
	file_teleport_devicetrust_v1_authenticate_challenge_proto_rawDescOnce.Do(func() {
		file_teleport_devicetrust_v1_authenticate_challenge_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_devicetrust_v1_authenticate_challenge_proto_rawDescData)
	})
	return file_teleport_devicetrust_v1_authenticate_challenge_proto_rawDescData
}

var file_teleport_devicetrust_v1_authenticate_challenge_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_teleport_devicetrust_v1_authenticate_challenge_proto_goTypes = []any{
	(*AuthenticateDeviceChallenge)(nil),            // 0: teleport.devicetrust.v1.AuthenticateDeviceChallenge
	(*AuthenticateDeviceChallengeResponse)(nil),    // 1: teleport.devicetrust.v1.AuthenticateDeviceChallengeResponse
	(*TPMAuthenticateDeviceChallenge)(nil),         // 2: teleport.devicetrust.v1.TPMAuthenticateDeviceChallenge
	(*TPMAuthenticateDeviceChallengeResponse)(nil), // 3: teleport.devicetrust.v1.TPMAuthenticateDeviceChallengeResponse
	(*TPMPlatformParameters)(nil),                  // 4: teleport.devicetrust.v1.TPMPlatformParameters
}
var file_teleport_devicetrust_v1_authenticate_challenge_proto_depIdxs = []int32{
	4, // 0: teleport.devicetrust.v1.TPMAuthenticateDeviceChallengeResponse.platform_parameters:type_name -> teleport.devicetrust.v1.TPMPlatformParameters
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_teleport_devicetrust_v1_authenticate_challenge_proto_init() }
func file_teleport_devicetrust_v1_authenticate_challenge_proto_init() {
	if File_teleport_devicetrust_v1_authenticate_challenge_proto != nil {
		return
	}
	file_teleport_devicetrust_v1_tpm_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_devicetrust_v1_authenticate_challenge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_devicetrust_v1_authenticate_challenge_proto_goTypes,
		DependencyIndexes: file_teleport_devicetrust_v1_authenticate_challenge_proto_depIdxs,
		MessageInfos:      file_teleport_devicetrust_v1_authenticate_challenge_proto_msgTypes,
	}.Build()
	File_teleport_devicetrust_v1_authenticate_challenge_proto = out.File
	file_teleport_devicetrust_v1_authenticate_challenge_proto_rawDesc = nil
	file_teleport_devicetrust_v1_authenticate_challenge_proto_goTypes = nil
	file_teleport_devicetrust_v1_authenticate_challenge_proto_depIdxs = nil
}
