/*
Copyright 2024 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: singer.proto

package protos

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Genre int32

const (
	Genre_POP  Genre = 0
	Genre_JAZZ Genre = 1
	Genre_FOLK Genre = 2
	Genre_ROCK Genre = 3
)

// Enum value maps for Genre.
var (
	Genre_name = map[int32]string{
		0: "POP",
		1: "JAZZ",
		2: "FOLK",
		3: "ROCK",
	}
	Genre_value = map[string]int32{
		"POP":  0,
		"JAZZ": 1,
		"FOLK": 2,
		"ROCK": 3,
	}
)

func (x Genre) Enum() *Genre {
	p := new(Genre)
	*p = x
	return p
}

func (x Genre) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Genre) Descriptor() protoreflect.EnumDescriptor {
	return file_singer_proto_enumTypes[0].Descriptor()
}

func (Genre) Type() protoreflect.EnumType {
	return &file_singer_proto_enumTypes[0]
}

func (x Genre) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Genre) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Genre(num)
	return nil
}

// Deprecated: Use Genre.Descriptor instead.
func (Genre) EnumDescriptor() ([]byte, []int) {
	return file_singer_proto_rawDescGZIP(), []int{0}
}

type SingerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SingerId    *int64  `protobuf:"varint,1,opt,name=singer_id,json=singerId" json:"singer_id,omitempty"`
	BirthDate   *string `protobuf:"bytes,2,opt,name=birth_date,json=birthDate" json:"birth_date,omitempty"`
	Nationality *string `protobuf:"bytes,3,opt,name=nationality" json:"nationality,omitempty"`
	Genre       *Genre  `protobuf:"varint,4,opt,name=genre,enum=examples.spanner.music.Genre" json:"genre,omitempty"`
}

func (x *SingerInfo) Reset() {
	*x = SingerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_singer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingerInfo) ProtoMessage() {}

func (x *SingerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_singer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingerInfo.ProtoReflect.Descriptor instead.
func (*SingerInfo) Descriptor() ([]byte, []int) {
	return file_singer_proto_rawDescGZIP(), []int{0}
}

func (x *SingerInfo) GetSingerId() int64 {
	if x != nil && x.SingerId != nil {
		return *x.SingerId
	}
	return 0
}

func (x *SingerInfo) GetBirthDate() string {
	if x != nil && x.BirthDate != nil {
		return *x.BirthDate
	}
	return ""
}

func (x *SingerInfo) GetNationality() string {
	if x != nil && x.Nationality != nil {
		return *x.Nationality
	}
	return ""
}

func (x *SingerInfo) GetGenre() Genre {
	if x != nil && x.Genre != nil {
		return *x.Genre
	}
	return Genre_POP
}

var File_singer_proto protoreflect.FileDescriptor

var file_singer_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x22, 0x9f, 0x01, 0x0a, 0x0a, 0x53, 0x69, 0x6e, 0x67, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x69, 0x6e, 0x67, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x70,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x6e, 0x72,
	0x65, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x2a, 0x2e, 0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72,
	0x65, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x4f, 0x50, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x41,
	0x5a, 0x5a, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x4f, 0x4c, 0x4b, 0x10, 0x02, 0x12, 0x08,
	0x0a, 0x04, 0x52, 0x4f, 0x43, 0x4b, 0x10, 0x03, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73,
}

var (
	file_singer_proto_rawDescOnce sync.Once
	file_singer_proto_rawDescData = file_singer_proto_rawDesc
)

func file_singer_proto_rawDescGZIP() []byte {
	file_singer_proto_rawDescOnce.Do(func() {
		file_singer_proto_rawDescData = protoimpl.X.CompressGZIP(file_singer_proto_rawDescData)
	})
	return file_singer_proto_rawDescData
}

var file_singer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_singer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_singer_proto_goTypes = []interface{}{
	(Genre)(0),         // 0: examples.spanner.music.Genre
	(*SingerInfo)(nil), // 1: examples.spanner.music.SingerInfo
}
var file_singer_proto_depIdxs = []int32{
	0, // 0: examples.spanner.music.SingerInfo.genre:type_name -> examples.spanner.music.Genre
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_singer_proto_init() }
func file_singer_proto_init() {
	if File_singer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_singer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingerInfo); i {
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
			RawDescriptor: file_singer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_singer_proto_goTypes,
		DependencyIndexes: file_singer_proto_depIdxs,
		EnumInfos:         file_singer_proto_enumTypes,
		MessageInfos:      file_singer_proto_msgTypes,
	}.Build()
	File_singer_proto = out.File
	file_singer_proto_rawDesc = nil
	file_singer_proto_goTypes = nil
	file_singer_proto_depIdxs = nil
}
