// Copyright 2022 Google LLC
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
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: google/cloud/discoveryengine/v1/purge_config.proto

package discoveryenginepb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for
// [DocumentService.PurgeDocuments][google.cloud.discoveryengine.v1.DocumentService.PurgeDocuments]
// method.
type PurgeDocumentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent resource name, such as
	// `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/branches/{branch}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. Filter matching documents to purge. Only currently supported
	// value is
	// `*` (all items).
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	// Actually performs the purge. If `force` is set to false, return the
	// expected purge count without deleting any documents.
	Force bool `protobuf:"varint,3,opt,name=force,proto3" json:"force,omitempty"`
}

func (x *PurgeDocumentsRequest) Reset() {
	*x = PurgeDocumentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeDocumentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeDocumentsRequest) ProtoMessage() {}

func (x *PurgeDocumentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeDocumentsRequest.ProtoReflect.Descriptor instead.
func (*PurgeDocumentsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1_purge_config_proto_rawDescGZIP(), []int{0}
}

func (x *PurgeDocumentsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *PurgeDocumentsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *PurgeDocumentsRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

// Response message for
// [DocumentService.PurgeDocuments][google.cloud.discoveryengine.v1.DocumentService.PurgeDocuments]
// method. If the long running operation is successfully done, then this message
// is returned by the google.longrunning.Operations.response field.
type PurgeDocumentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The total count of documents purged as a result of the operation.
	PurgeCount int64 `protobuf:"varint,1,opt,name=purge_count,json=purgeCount,proto3" json:"purge_count,omitempty"`
	// A sample of document names that will be deleted. Only populated if `force`
	// is set to false. A max of 100 names will be returned and the names are
	// chosen at random.
	PurgeSample []string `protobuf:"bytes,2,rep,name=purge_sample,json=purgeSample,proto3" json:"purge_sample,omitempty"`
}

func (x *PurgeDocumentsResponse) Reset() {
	*x = PurgeDocumentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeDocumentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeDocumentsResponse) ProtoMessage() {}

func (x *PurgeDocumentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeDocumentsResponse.ProtoReflect.Descriptor instead.
func (*PurgeDocumentsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1_purge_config_proto_rawDescGZIP(), []int{1}
}

func (x *PurgeDocumentsResponse) GetPurgeCount() int64 {
	if x != nil {
		return x.PurgeCount
	}
	return 0
}

func (x *PurgeDocumentsResponse) GetPurgeSample() []string {
	if x != nil {
		return x.PurgeSample
	}
	return nil
}

// Metadata related to the progress of the PurgeDocuments operation.
// This will be returned by the google.longrunning.Operation.metadata field.
type PurgeDocumentsMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Operation create time.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Operation last update time. If the operation is done, this is also the
	// finish time.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Count of entries that were deleted successfully.
	SuccessCount int64 `protobuf:"varint,3,opt,name=success_count,json=successCount,proto3" json:"success_count,omitempty"`
	// Count of entries that encountered errors while processing.
	FailureCount int64 `protobuf:"varint,4,opt,name=failure_count,json=failureCount,proto3" json:"failure_count,omitempty"`
	// Count of entries that were ignored as entries were not found.
	IgnoredCount int64 `protobuf:"varint,5,opt,name=ignored_count,json=ignoredCount,proto3" json:"ignored_count,omitempty"`
}

func (x *PurgeDocumentsMetadata) Reset() {
	*x = PurgeDocumentsMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeDocumentsMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeDocumentsMetadata) ProtoMessage() {}

func (x *PurgeDocumentsMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeDocumentsMetadata.ProtoReflect.Descriptor instead.
func (*PurgeDocumentsMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1_purge_config_proto_rawDescGZIP(), []int{2}
}

func (x *PurgeDocumentsMetadata) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *PurgeDocumentsMetadata) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *PurgeDocumentsMetadata) GetSuccessCount() int64 {
	if x != nil {
		return x.SuccessCount
	}
	return 0
}

func (x *PurgeDocumentsMetadata) GetFailureCount() int64 {
	if x != nil {
		return x.FailureCount
	}
	return 0
}

func (x *PurgeDocumentsMetadata) GetIgnoredCount() int64 {
	if x != nil {
		return x.IgnoredCount
	}
	return 0
}

// Request message for
// [CompletionService.PurgeSuggestionDenyListEntries][google.cloud.discoveryengine.v1.CompletionService.PurgeSuggestionDenyListEntries]
// method.
type PurgeSuggestionDenyListEntriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent data store resource name for which to import denylist
	// entries. Follows pattern projects/*/locations/*/collections/*/dataStores/*.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (x *PurgeSuggestionDenyListEntriesRequest) Reset() {
	*x = PurgeSuggestionDenyListEntriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeSuggestionDenyListEntriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeSuggestionDenyListEntriesRequest) ProtoMessage() {}

func (x *PurgeSuggestionDenyListEntriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeSuggestionDenyListEntriesRequest.ProtoReflect.Descriptor instead.
func (*PurgeSuggestionDenyListEntriesRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1_purge_config_proto_rawDescGZIP(), []int{3}
}

func (x *PurgeSuggestionDenyListEntriesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

// Response message for
// [CompletionService.PurgeSuggestionDenyListEntries][google.cloud.discoveryengine.v1.CompletionService.PurgeSuggestionDenyListEntries]
// method.
type PurgeSuggestionDenyListEntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of suggestion deny list entries purged.
	PurgeCount int64 `protobuf:"varint,1,opt,name=purge_count,json=purgeCount,proto3" json:"purge_count,omitempty"`
	// A sample of errors encountered while processing the request.
	ErrorSamples []*status.Status `protobuf:"bytes,2,rep,name=error_samples,json=errorSamples,proto3" json:"error_samples,omitempty"`
}

func (x *PurgeSuggestionDenyListEntriesResponse) Reset() {
	*x = PurgeSuggestionDenyListEntriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeSuggestionDenyListEntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeSuggestionDenyListEntriesResponse) ProtoMessage() {}

func (x *PurgeSuggestionDenyListEntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeSuggestionDenyListEntriesResponse.ProtoReflect.Descriptor instead.
func (*PurgeSuggestionDenyListEntriesResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1_purge_config_proto_rawDescGZIP(), []int{4}
}

func (x *PurgeSuggestionDenyListEntriesResponse) GetPurgeCount() int64 {
	if x != nil {
		return x.PurgeCount
	}
	return 0
}

func (x *PurgeSuggestionDenyListEntriesResponse) GetErrorSamples() []*status.Status {
	if x != nil {
		return x.ErrorSamples
	}
	return nil
}

// Metadata related to the progress of the PurgeSuggestionDenyListEntries
// operation. This is returned by the google.longrunning.Operation.metadata
// field.
type PurgeSuggestionDenyListEntriesMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Operation create time.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Operation last update time. If the operation is done, this is also the
	// finish time.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *PurgeSuggestionDenyListEntriesMetadata) Reset() {
	*x = PurgeSuggestionDenyListEntriesMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeSuggestionDenyListEntriesMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeSuggestionDenyListEntriesMetadata) ProtoMessage() {}

func (x *PurgeSuggestionDenyListEntriesMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeSuggestionDenyListEntriesMetadata.ProtoReflect.Descriptor instead.
func (*PurgeSuggestionDenyListEntriesMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1_purge_config_proto_rawDescGZIP(), []int{5}
}

func (x *PurgeSuggestionDenyListEntriesMetadata) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *PurgeSuggestionDenyListEntriesMetadata) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_google_cloud_discoveryengine_v1_purge_config_proto protoreflect.FileDescriptor

var file_google_cloud_discoveryengine_v1_purge_config_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x75, 0x72, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x15,
	0x50, 0x75, 0x72, 0x67, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x27, 0x0a, 0x25, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x22,
	0x8a, 0x01, 0x0a, 0x16, 0x50, 0x75, 0x72, 0x67, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x75,
	0x72, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x70, 0x75, 0x72, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x4f, 0x0a, 0x0c, 0x70,
	0x75, 0x72, 0x67, 0x65, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x2c, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0b, 0x70, 0x75, 0x72, 0x67, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x81, 0x02, 0x0a,
	0x16, 0x50, 0x75, 0x72, 0x67, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x69,
	0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x71, 0x0a, 0x25, 0x50, 0x75, 0x72, 0x67, 0x65, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xe0, 0x41, 0x02, 0xfa, 0x41,
	0x2a, 0x0a, 0x28, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x26, 0x50, 0x75, 0x72, 0x67, 0x65, 0x53, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x75, 0x72, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x75, 0x72, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x37, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x26, 0x50, 0x75, 0x72,
	0x67, 0x65, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6e, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x83, 0x02,
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x50, 0x75, 0x72, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0xa2, 0x02, 0x0f, 0x44, 0x49, 0x53, 0x43, 0x4f,
	0x56, 0x45, 0x52, 0x59, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1f, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0xea, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_discoveryengine_v1_purge_config_proto_rawDescOnce sync.Once
	file_google_cloud_discoveryengine_v1_purge_config_proto_rawDescData = file_google_cloud_discoveryengine_v1_purge_config_proto_rawDesc
)

func file_google_cloud_discoveryengine_v1_purge_config_proto_rawDescGZIP() []byte {
	file_google_cloud_discoveryengine_v1_purge_config_proto_rawDescOnce.Do(func() {
		file_google_cloud_discoveryengine_v1_purge_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_discoveryengine_v1_purge_config_proto_rawDescData)
	})
	return file_google_cloud_discoveryengine_v1_purge_config_proto_rawDescData
}

var file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_discoveryengine_v1_purge_config_proto_goTypes = []interface{}{
	(*PurgeDocumentsRequest)(nil),                  // 0: google.cloud.discoveryengine.v1.PurgeDocumentsRequest
	(*PurgeDocumentsResponse)(nil),                 // 1: google.cloud.discoveryengine.v1.PurgeDocumentsResponse
	(*PurgeDocumentsMetadata)(nil),                 // 2: google.cloud.discoveryengine.v1.PurgeDocumentsMetadata
	(*PurgeSuggestionDenyListEntriesRequest)(nil),  // 3: google.cloud.discoveryengine.v1.PurgeSuggestionDenyListEntriesRequest
	(*PurgeSuggestionDenyListEntriesResponse)(nil), // 4: google.cloud.discoveryengine.v1.PurgeSuggestionDenyListEntriesResponse
	(*PurgeSuggestionDenyListEntriesMetadata)(nil), // 5: google.cloud.discoveryengine.v1.PurgeSuggestionDenyListEntriesMetadata
	(*timestamppb.Timestamp)(nil),                  // 6: google.protobuf.Timestamp
	(*status.Status)(nil),                          // 7: google.rpc.Status
}
var file_google_cloud_discoveryengine_v1_purge_config_proto_depIdxs = []int32{
	6, // 0: google.cloud.discoveryengine.v1.PurgeDocumentsMetadata.create_time:type_name -> google.protobuf.Timestamp
	6, // 1: google.cloud.discoveryengine.v1.PurgeDocumentsMetadata.update_time:type_name -> google.protobuf.Timestamp
	7, // 2: google.cloud.discoveryengine.v1.PurgeSuggestionDenyListEntriesResponse.error_samples:type_name -> google.rpc.Status
	6, // 3: google.cloud.discoveryengine.v1.PurgeSuggestionDenyListEntriesMetadata.create_time:type_name -> google.protobuf.Timestamp
	6, // 4: google.cloud.discoveryengine.v1.PurgeSuggestionDenyListEntriesMetadata.update_time:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_discoveryengine_v1_purge_config_proto_init() }
func file_google_cloud_discoveryengine_v1_purge_config_proto_init() {
	if File_google_cloud_discoveryengine_v1_purge_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeDocumentsRequest); i {
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
		file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeDocumentsResponse); i {
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
		file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeDocumentsMetadata); i {
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
		file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeSuggestionDenyListEntriesRequest); i {
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
		file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeSuggestionDenyListEntriesResponse); i {
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
		file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeSuggestionDenyListEntriesMetadata); i {
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
			RawDescriptor: file_google_cloud_discoveryengine_v1_purge_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_discoveryengine_v1_purge_config_proto_goTypes,
		DependencyIndexes: file_google_cloud_discoveryengine_v1_purge_config_proto_depIdxs,
		MessageInfos:      file_google_cloud_discoveryengine_v1_purge_config_proto_msgTypes,
	}.Build()
	File_google_cloud_discoveryengine_v1_purge_config_proto = out.File
	file_google_cloud_discoveryengine_v1_purge_config_proto_rawDesc = nil
	file_google_cloud_discoveryengine_v1_purge_config_proto_goTypes = nil
	file_google_cloud_discoveryengine_v1_purge_config_proto_depIdxs = nil
}
