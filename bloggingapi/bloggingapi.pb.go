// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: bloggingapi/bloggingapi.proto

package bloggingapi

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

// The request message containing the blog's details.
type BlogPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostID  uint64 `protobuf:"varint,1,opt,name=postID,proto3" json:"postID,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Author  string `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *BlogPost) Reset() {
	*x = BlogPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bloggingapi_bloggingapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogPost) ProtoMessage() {}

func (x *BlogPost) ProtoReflect() protoreflect.Message {
	mi := &file_bloggingapi_bloggingapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogPost.ProtoReflect.Descriptor instead.
func (*BlogPost) Descriptor() ([]byte, []int) {
	return file_bloggingapi_bloggingapi_proto_rawDescGZIP(), []int{0}
}

func (x *BlogPost) GetPostID() uint64 {
	if x != nil {
		return x.PostID
	}
	return 0
}

func (x *BlogPost) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BlogPost) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *BlogPost) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

// The response message containing the greetings
type BlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Postid uint64 `protobuf:"varint,1,opt,name=postid,proto3" json:"postid,omitempty"`
}

func (x *BlogRequest) Reset() {
	*x = BlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bloggingapi_bloggingapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogRequest) ProtoMessage() {}

func (x *BlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bloggingapi_bloggingapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogRequest.ProtoReflect.Descriptor instead.
func (*BlogRequest) Descriptor() ([]byte, []int) {
	return file_bloggingapi_bloggingapi_proto_rawDescGZIP(), []int{1}
}

func (x *BlogRequest) GetPostid() uint64 {
	if x != nil {
		return x.Postid
	}
	return 0
}

var File_bloggingapi_bloggingapi_proto protoreflect.FileDescriptor

var file_bloggingapi_bloggingapi_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x22, 0x6a, 0x0a, 0x08,
	0x42, 0x6c, 0x6f, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x25, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x64, 0x32,
	0x8a, 0x01, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x12,
	0x3c, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x15, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6c, 0x6f, 0x67,
	0x50, 0x6f, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x61,
	0x70, 0x69, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x08, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x18, 0x2e, 0x62, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x61, 0x70,
	0x69, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b,
	0x67, 0x6f, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2f,
	0x62, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_bloggingapi_bloggingapi_proto_rawDescOnce sync.Once
	file_bloggingapi_bloggingapi_proto_rawDescData = file_bloggingapi_bloggingapi_proto_rawDesc
)

func file_bloggingapi_bloggingapi_proto_rawDescGZIP() []byte {
	file_bloggingapi_bloggingapi_proto_rawDescOnce.Do(func() {
		file_bloggingapi_bloggingapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_bloggingapi_bloggingapi_proto_rawDescData)
	})
	return file_bloggingapi_bloggingapi_proto_rawDescData
}

var file_bloggingapi_bloggingapi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bloggingapi_bloggingapi_proto_goTypes = []interface{}{
	(*BlogPost)(nil),    // 0: bloggingapi.BlogPost
	(*BlogRequest)(nil), // 1: bloggingapi.BlogRequest
}
var file_bloggingapi_bloggingapi_proto_depIdxs = []int32{
	0, // 0: bloggingapi.bloggingapi.CreateBlog:input_type -> bloggingapi.BlogPost
	1, // 1: bloggingapi.bloggingapi.ReadBlog:input_type -> bloggingapi.BlogRequest
	0, // 2: bloggingapi.bloggingapi.CreateBlog:output_type -> bloggingapi.BlogPost
	0, // 3: bloggingapi.bloggingapi.ReadBlog:output_type -> bloggingapi.BlogPost
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bloggingapi_bloggingapi_proto_init() }
func file_bloggingapi_bloggingapi_proto_init() {
	if File_bloggingapi_bloggingapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bloggingapi_bloggingapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogPost); i {
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
		file_bloggingapi_bloggingapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogRequest); i {
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
			RawDescriptor: file_bloggingapi_bloggingapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bloggingapi_bloggingapi_proto_goTypes,
		DependencyIndexes: file_bloggingapi_bloggingapi_proto_depIdxs,
		MessageInfos:      file_bloggingapi_bloggingapi_proto_msgTypes,
	}.Build()
	File_bloggingapi_bloggingapi_proto = out.File
	file_bloggingapi_bloggingapi_proto_rawDesc = nil
	file_bloggingapi_bloggingapi_proto_goTypes = nil
	file_bloggingapi_bloggingapi_proto_depIdxs = nil
}
