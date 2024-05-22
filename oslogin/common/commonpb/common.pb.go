// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.3
// source: google/cloud/oslogin/common/common.proto

package commonpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The operating system options for account entries.
type OperatingSystemType int32

const (
	// The operating system type associated with the user account information is
	// unspecified.
	OperatingSystemType_OPERATING_SYSTEM_TYPE_UNSPECIFIED OperatingSystemType = 0
	// Linux user account information.
	OperatingSystemType_LINUX OperatingSystemType = 1
	// Windows user account information.
	OperatingSystemType_WINDOWS OperatingSystemType = 2
)

// Enum value maps for OperatingSystemType.
var (
	OperatingSystemType_name = map[int32]string{
		0: "OPERATING_SYSTEM_TYPE_UNSPECIFIED",
		1: "LINUX",
		2: "WINDOWS",
	}
	OperatingSystemType_value = map[string]int32{
		"OPERATING_SYSTEM_TYPE_UNSPECIFIED": 0,
		"LINUX":                             1,
		"WINDOWS":                           2,
	}
)

func (x OperatingSystemType) Enum() *OperatingSystemType {
	p := new(OperatingSystemType)
	*p = x
	return p
}

func (x OperatingSystemType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperatingSystemType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_oslogin_common_common_proto_enumTypes[0].Descriptor()
}

func (OperatingSystemType) Type() protoreflect.EnumType {
	return &file_google_cloud_oslogin_common_common_proto_enumTypes[0]
}

func (x OperatingSystemType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperatingSystemType.Descriptor instead.
func (OperatingSystemType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_oslogin_common_common_proto_rawDescGZIP(), []int{0}
}

// The POSIX account information associated with a Google account.
type PosixAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Only one POSIX account can be marked as primary.
	Primary bool `protobuf:"varint,1,opt,name=primary,proto3" json:"primary,omitempty"`
	// The username of the POSIX account.
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// The user ID.
	Uid int64 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	// The default group ID.
	Gid int64 `protobuf:"varint,4,opt,name=gid,proto3" json:"gid,omitempty"`
	// The path to the home directory for this account.
	HomeDirectory string `protobuf:"bytes,5,opt,name=home_directory,json=homeDirectory,proto3" json:"home_directory,omitempty"`
	// The path to the logic shell for this account.
	Shell string `protobuf:"bytes,6,opt,name=shell,proto3" json:"shell,omitempty"`
	// The GECOS (user information) entry for this account.
	Gecos string `protobuf:"bytes,7,opt,name=gecos,proto3" json:"gecos,omitempty"`
	// System identifier for which account the username or uid applies to.
	// By default, the empty value is used.
	SystemId string `protobuf:"bytes,8,opt,name=system_id,json=systemId,proto3" json:"system_id,omitempty"`
	// Output only. A POSIX account identifier.
	AccountId string `protobuf:"bytes,9,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// The operating system type where this account applies.
	OperatingSystemType OperatingSystemType `protobuf:"varint,10,opt,name=operating_system_type,json=operatingSystemType,proto3,enum=google.cloud.oslogin.common.OperatingSystemType" json:"operating_system_type,omitempty"`
	// Output only. The canonical resource name.
	Name string `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PosixAccount) Reset() {
	*x = PosixAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_oslogin_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PosixAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PosixAccount) ProtoMessage() {}

func (x *PosixAccount) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_oslogin_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PosixAccount.ProtoReflect.Descriptor instead.
func (*PosixAccount) Descriptor() ([]byte, []int) {
	return file_google_cloud_oslogin_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *PosixAccount) GetPrimary() bool {
	if x != nil {
		return x.Primary
	}
	return false
}

func (x *PosixAccount) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PosixAccount) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PosixAccount) GetGid() int64 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *PosixAccount) GetHomeDirectory() string {
	if x != nil {
		return x.HomeDirectory
	}
	return ""
}

func (x *PosixAccount) GetShell() string {
	if x != nil {
		return x.Shell
	}
	return ""
}

func (x *PosixAccount) GetGecos() string {
	if x != nil {
		return x.Gecos
	}
	return ""
}

func (x *PosixAccount) GetSystemId() string {
	if x != nil {
		return x.SystemId
	}
	return ""
}

func (x *PosixAccount) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *PosixAccount) GetOperatingSystemType() OperatingSystemType {
	if x != nil {
		return x.OperatingSystemType
	}
	return OperatingSystemType_OPERATING_SYSTEM_TYPE_UNSPECIFIED
}

func (x *PosixAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The SSH public key information associated with a Google account.
type SshPublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Public key text in SSH format, defined by
	// <a href="https://www.ietf.org/rfc/rfc4253.txt" target="_blank">RFC4253</a>
	// section 6.6.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// An expiration time in microseconds since epoch.
	ExpirationTimeUsec int64 `protobuf:"varint,2,opt,name=expiration_time_usec,json=expirationTimeUsec,proto3" json:"expiration_time_usec,omitempty"`
	// Output only. The SHA-256 fingerprint of the SSH public key.
	Fingerprint string `protobuf:"bytes,3,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// Output only. The canonical resource name.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SshPublicKey) Reset() {
	*x = SshPublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_oslogin_common_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SshPublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SshPublicKey) ProtoMessage() {}

func (x *SshPublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_oslogin_common_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SshPublicKey.ProtoReflect.Descriptor instead.
func (*SshPublicKey) Descriptor() ([]byte, []int) {
	return file_google_cloud_oslogin_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *SshPublicKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SshPublicKey) GetExpirationTimeUsec() int64 {
	if x != nil {
		return x.ExpirationTimeUsec
	}
	return 0
}

func (x *SshPublicKey) GetFingerprint() string {
	if x != nil {
		return x.Fingerprint
	}
	return ""
}

func (x *SshPublicKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_google_cloud_oslogin_common_common_proto protoreflect.FileDescriptor

var file_google_cloud_oslogin_common_common_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f,
	0x73, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x73, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x03, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x69, 0x78, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x67, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x6f, 0x6d, 0x65, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x67,
	0x65, 0x63, 0x6f, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x65, 0x63, 0x6f,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x64, 0x0a, 0x15, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6f, 0x73, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x13, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x3a, 0x49, 0xea, 0x41, 0x46, 0x0a, 0x23, 0x6f, 0x73, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50,
	0x6f, 0x73, 0x69, 0x78, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x22, 0xe6, 0x01, 0x0a,
	0x0c, 0x53, 0x73, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x30, 0x0a, 0x14, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x65,
	0x63, 0x12, 0x25, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x66, 0x69, 0x6e,
	0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x3a, 0x52, 0xea, 0x41, 0x4f, 0x0a, 0x23, 0x6f, 0x73, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53,
	0x73, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x7d, 0x2f, 0x73, 0x73, 0x68, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x2f, 0x7b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x7d, 0x2a, 0x54, 0x0a, 0x13, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x21,
	0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x49, 0x4e, 0x55, 0x58, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x53, 0x10, 0x02, 0x42, 0xf0, 0x01, 0xea, 0x41,
	0x2b, 0x0a, 0x1b, 0x6f, 0x73, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x7d, 0x0a, 0x1f, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f,
	0x73, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x0c, 0x4f,
	0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x34, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x6f, 0x73, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70,
	0x62, 0xaa, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x4f, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca,
	0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4f,
	0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xea, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4f,
	0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_oslogin_common_common_proto_rawDescOnce sync.Once
	file_google_cloud_oslogin_common_common_proto_rawDescData = file_google_cloud_oslogin_common_common_proto_rawDesc
)

func file_google_cloud_oslogin_common_common_proto_rawDescGZIP() []byte {
	file_google_cloud_oslogin_common_common_proto_rawDescOnce.Do(func() {
		file_google_cloud_oslogin_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_oslogin_common_common_proto_rawDescData)
	})
	return file_google_cloud_oslogin_common_common_proto_rawDescData
}

var file_google_cloud_oslogin_common_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_oslogin_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_oslogin_common_common_proto_goTypes = []interface{}{
	(OperatingSystemType)(0), // 0: google.cloud.oslogin.common.OperatingSystemType
	(*PosixAccount)(nil),     // 1: google.cloud.oslogin.common.PosixAccount
	(*SshPublicKey)(nil),     // 2: google.cloud.oslogin.common.SshPublicKey
}
var file_google_cloud_oslogin_common_common_proto_depIdxs = []int32{
	0, // 0: google.cloud.oslogin.common.PosixAccount.operating_system_type:type_name -> google.cloud.oslogin.common.OperatingSystemType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_oslogin_common_common_proto_init() }
func file_google_cloud_oslogin_common_common_proto_init() {
	if File_google_cloud_oslogin_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_oslogin_common_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PosixAccount); i {
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
		file_google_cloud_oslogin_common_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SshPublicKey); i {
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
			RawDescriptor: file_google_cloud_oslogin_common_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_oslogin_common_common_proto_goTypes,
		DependencyIndexes: file_google_cloud_oslogin_common_common_proto_depIdxs,
		EnumInfos:         file_google_cloud_oslogin_common_common_proto_enumTypes,
		MessageInfos:      file_google_cloud_oslogin_common_common_proto_msgTypes,
	}.Build()
	File_google_cloud_oslogin_common_common_proto = out.File
	file_google_cloud_oslogin_common_common_proto_rawDesc = nil
	file_google_cloud_oslogin_common_common_proto_goTypes = nil
	file_google_cloud_oslogin_common_common_proto_depIdxs = nil
}
