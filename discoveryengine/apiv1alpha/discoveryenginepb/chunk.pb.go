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
// source: google/cloud/discoveryengine/v1alpha/chunk.proto

package discoveryenginepb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Chunk captures all raw metadata information of items to be recommended or
// searched in the chunk mode.
type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The full resource name of the chunk.
	// Format:
	// `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/branches/{branch}/documents/{document_id}/chunks/{chunk_id}`.
	//
	// This field must be a UTF-8 encoded string with a length limit of 1024
	// characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Unique chunk ID of the current chunk.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Content is a string from a document (parsed content).
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// Metadata of the document from the current chunk.
	DocumentMetadata *Chunk_DocumentMetadata `protobuf:"bytes,5,opt,name=document_metadata,json=documentMetadata,proto3" json:"document_metadata,omitempty"`
	// Output only. This field is OUTPUT_ONLY.
	// It contains derived data that are not in the original input document.
	DerivedStructData *structpb.Struct `protobuf:"bytes,4,opt,name=derived_struct_data,json=derivedStructData,proto3" json:"derived_struct_data,omitempty"`
	// Page span of the chunk.
	PageSpan *Chunk_PageSpan `protobuf:"bytes,6,opt,name=page_span,json=pageSpan,proto3" json:"page_span,omitempty"`
	// Output only. Metadata of the current chunk.
	ChunkMetadata *Chunk_ChunkMetadata `protobuf:"bytes,7,opt,name=chunk_metadata,json=chunkMetadata,proto3" json:"chunk_metadata,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_chunk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_chunk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_chunk_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Chunk) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Chunk) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Chunk) GetDocumentMetadata() *Chunk_DocumentMetadata {
	if x != nil {
		return x.DocumentMetadata
	}
	return nil
}

func (x *Chunk) GetDerivedStructData() *structpb.Struct {
	if x != nil {
		return x.DerivedStructData
	}
	return nil
}

func (x *Chunk) GetPageSpan() *Chunk_PageSpan {
	if x != nil {
		return x.PageSpan
	}
	return nil
}

func (x *Chunk) GetChunkMetadata() *Chunk_ChunkMetadata {
	if x != nil {
		return x.ChunkMetadata
	}
	return nil
}

// Document metadata contains the information of the document of the current
// chunk.
type Chunk_DocumentMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Uri of the document.
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Title of the document.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Data representation.
	// The structured JSON data for the document. It should conform to the
	// registered [Schema][google.cloud.discoveryengine.v1alpha.Schema] or an
	// `INVALID_ARGUMENT` error is thrown.
	StructData *structpb.Struct `protobuf:"bytes,3,opt,name=struct_data,json=structData,proto3" json:"struct_data,omitempty"`
}

func (x *Chunk_DocumentMetadata) Reset() {
	*x = Chunk_DocumentMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_chunk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk_DocumentMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk_DocumentMetadata) ProtoMessage() {}

func (x *Chunk_DocumentMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_chunk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk_DocumentMetadata.ProtoReflect.Descriptor instead.
func (*Chunk_DocumentMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_chunk_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Chunk_DocumentMetadata) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Chunk_DocumentMetadata) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Chunk_DocumentMetadata) GetStructData() *structpb.Struct {
	if x != nil {
		return x.StructData
	}
	return nil
}

// Page span of the chunk.
type Chunk_PageSpan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The start page of the chunk.
	PageStart int32 `protobuf:"varint,1,opt,name=page_start,json=pageStart,proto3" json:"page_start,omitempty"`
	// The end page of the chunk.
	PageEnd int32 `protobuf:"varint,2,opt,name=page_end,json=pageEnd,proto3" json:"page_end,omitempty"`
}

func (x *Chunk_PageSpan) Reset() {
	*x = Chunk_PageSpan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_chunk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk_PageSpan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk_PageSpan) ProtoMessage() {}

func (x *Chunk_PageSpan) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_chunk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk_PageSpan.ProtoReflect.Descriptor instead.
func (*Chunk_PageSpan) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_chunk_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Chunk_PageSpan) GetPageStart() int32 {
	if x != nil {
		return x.PageStart
	}
	return 0
}

func (x *Chunk_PageSpan) GetPageEnd() int32 {
	if x != nil {
		return x.PageEnd
	}
	return 0
}

// Metadata of the current chunk. This field is only populated on
// [SearchService.Search][google.cloud.discoveryengine.v1alpha.SearchService.Search]
// API.
type Chunk_ChunkMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The previous chunks of the current chunk. The number is controlled by
	// [SearchRequest.ContentSearchSpec.ChunkSpec.num_previous_chunks][google.cloud.discoveryengine.v1alpha.SearchRequest.ContentSearchSpec.ChunkSpec.num_previous_chunks].
	// This field is only populated on
	// [SearchService.Search][google.cloud.discoveryengine.v1alpha.SearchService.Search]
	// API.
	PreviousChunks []*Chunk `protobuf:"bytes,1,rep,name=previous_chunks,json=previousChunks,proto3" json:"previous_chunks,omitempty"`
	// The next chunks of the current chunk. The number is controlled by
	// [SearchRequest.ContentSearchSpec.ChunkSpec.num_next_chunks][google.cloud.discoveryengine.v1alpha.SearchRequest.ContentSearchSpec.ChunkSpec.num_next_chunks].
	// This field is only populated on
	// [SearchService.Search][google.cloud.discoveryengine.v1alpha.SearchService.Search]
	// API.
	NextChunks []*Chunk `protobuf:"bytes,2,rep,name=next_chunks,json=nextChunks,proto3" json:"next_chunks,omitempty"`
}

func (x *Chunk_ChunkMetadata) Reset() {
	*x = Chunk_ChunkMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_chunk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk_ChunkMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk_ChunkMetadata) ProtoMessage() {}

func (x *Chunk_ChunkMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_chunk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk_ChunkMetadata.ProtoReflect.Descriptor instead.
func (*Chunk_ChunkMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_chunk_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Chunk_ChunkMetadata) GetPreviousChunks() []*Chunk {
	if x != nil {
		return x.PreviousChunks
	}
	return nil
}

func (x *Chunk_ChunkMetadata) GetNextChunks() []*Chunk {
	if x != nil {
		return x.NextChunks
	}
	return nil
}

var File_google_cloud_discoveryengine_v1alpha_chunk_proto protoreflect.FileDescriptor

var file_google_cloud_discoveryengine_v1alpha_chunk_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xdf, 0x08, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x69, 0x0a, 0x11, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x10, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x4c, 0x0a, 0x13, 0x64, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64,
	0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x11, 0x64, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x51, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x61, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x12, 0x65, 0x0a, 0x0e, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x2e, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x74, 0x0a,
	0x10, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x69, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x44, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x70, 0x61, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x64, 0x1a, 0xb3, 0x01, 0x0a, 0x0d, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x54, 0x0a, 0x0f, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x73, 0x12, 0x4c, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x3a,
	0xb2, 0x02, 0xea, 0x41, 0xae, 0x02, 0x0a, 0x24, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x75, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73,
	0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x2f, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2f, 0x7b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x7d, 0x2f,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x7d, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x7d, 0x12, 0x8e, 0x01, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x2f, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2f, 0x7b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x7d,
	0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x7d, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x2f, 0x7b, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x7d, 0x42, 0x96, 0x02, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x42, 0x0a, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x52, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70,
	0x62, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x70, 0x62, 0xa2, 0x02, 0x0f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x45,
	0x4e, 0x47, 0x49, 0x4e, 0x45, 0xaa, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0x41, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x24, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0xea, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_discoveryengine_v1alpha_chunk_proto_rawDescOnce sync.Once
	file_google_cloud_discoveryengine_v1alpha_chunk_proto_rawDescData = file_google_cloud_discoveryengine_v1alpha_chunk_proto_rawDesc
)

func file_google_cloud_discoveryengine_v1alpha_chunk_proto_rawDescGZIP() []byte {
	file_google_cloud_discoveryengine_v1alpha_chunk_proto_rawDescOnce.Do(func() {
		file_google_cloud_discoveryengine_v1alpha_chunk_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_discoveryengine_v1alpha_chunk_proto_rawDescData)
	})
	return file_google_cloud_discoveryengine_v1alpha_chunk_proto_rawDescData
}

var file_google_cloud_discoveryengine_v1alpha_chunk_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_discoveryengine_v1alpha_chunk_proto_goTypes = []interface{}{
	(*Chunk)(nil),                  // 0: google.cloud.discoveryengine.v1alpha.Chunk
	(*Chunk_DocumentMetadata)(nil), // 1: google.cloud.discoveryengine.v1alpha.Chunk.DocumentMetadata
	(*Chunk_PageSpan)(nil),         // 2: google.cloud.discoveryengine.v1alpha.Chunk.PageSpan
	(*Chunk_ChunkMetadata)(nil),    // 3: google.cloud.discoveryengine.v1alpha.Chunk.ChunkMetadata
	(*structpb.Struct)(nil),        // 4: google.protobuf.Struct
}
var file_google_cloud_discoveryengine_v1alpha_chunk_proto_depIdxs = []int32{
	1, // 0: google.cloud.discoveryengine.v1alpha.Chunk.document_metadata:type_name -> google.cloud.discoveryengine.v1alpha.Chunk.DocumentMetadata
	4, // 1: google.cloud.discoveryengine.v1alpha.Chunk.derived_struct_data:type_name -> google.protobuf.Struct
	2, // 2: google.cloud.discoveryengine.v1alpha.Chunk.page_span:type_name -> google.cloud.discoveryengine.v1alpha.Chunk.PageSpan
	3, // 3: google.cloud.discoveryengine.v1alpha.Chunk.chunk_metadata:type_name -> google.cloud.discoveryengine.v1alpha.Chunk.ChunkMetadata
	4, // 4: google.cloud.discoveryengine.v1alpha.Chunk.DocumentMetadata.struct_data:type_name -> google.protobuf.Struct
	0, // 5: google.cloud.discoveryengine.v1alpha.Chunk.ChunkMetadata.previous_chunks:type_name -> google.cloud.discoveryengine.v1alpha.Chunk
	0, // 6: google.cloud.discoveryengine.v1alpha.Chunk.ChunkMetadata.next_chunks:type_name -> google.cloud.discoveryengine.v1alpha.Chunk
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_cloud_discoveryengine_v1alpha_chunk_proto_init() }
func file_google_cloud_discoveryengine_v1alpha_chunk_proto_init() {
	if File_google_cloud_discoveryengine_v1alpha_chunk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_discoveryengine_v1alpha_chunk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_google_cloud_discoveryengine_v1alpha_chunk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk_DocumentMetadata); i {
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
		file_google_cloud_discoveryengine_v1alpha_chunk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk_PageSpan); i {
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
		file_google_cloud_discoveryengine_v1alpha_chunk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk_ChunkMetadata); i {
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
			RawDescriptor: file_google_cloud_discoveryengine_v1alpha_chunk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_discoveryengine_v1alpha_chunk_proto_goTypes,
		DependencyIndexes: file_google_cloud_discoveryengine_v1alpha_chunk_proto_depIdxs,
		MessageInfos:      file_google_cloud_discoveryengine_v1alpha_chunk_proto_msgTypes,
	}.Build()
	File_google_cloud_discoveryengine_v1alpha_chunk_proto = out.File
	file_google_cloud_discoveryengine_v1alpha_chunk_proto_rawDesc = nil
	file_google_cloud_discoveryengine_v1alpha_chunk_proto_goTypes = nil
	file_google_cloud_discoveryengine_v1alpha_chunk_proto_depIdxs = nil
}
