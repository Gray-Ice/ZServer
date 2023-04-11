// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: clipboard.proto

package clipboard

import (
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

type ClipboardContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ClipboardContent) Reset() {
	*x = ClipboardContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clipboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClipboardContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClipboardContent) ProtoMessage() {}

func (x *ClipboardContent) ProtoReflect() protoreflect.Message {
	mi := &file_clipboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClipboardContent.ProtoReflect.Descriptor instead.
func (*ClipboardContent) Descriptor() ([]byte, []int) {
	return file_clipboard_proto_rawDescGZIP(), []int{0}
}

func (x *ClipboardContent) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_clipboard_proto protoreflect.FileDescriptor

var file_clipboard_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x72, 0x70, 0x63,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a,
	0x10, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x32, 0x58, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x12, 0x4b, 0x0a, 0x0e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x43, 0x6c, 0x69, 0x70, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x12, 0x1f, 0x2e, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42,
	0x13, 0x5a, 0x11, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x70, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_clipboard_proto_rawDescOnce sync.Once
	file_clipboard_proto_rawDescData = file_clipboard_proto_rawDesc
)

func file_clipboard_proto_rawDescGZIP() []byte {
	file_clipboard_proto_rawDescOnce.Do(func() {
		file_clipboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_clipboard_proto_rawDescData)
	})
	return file_clipboard_proto_rawDescData
}

var file_clipboard_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_clipboard_proto_goTypes = []interface{}{
	(*ClipboardContent)(nil), // 0: clipboard_rpc.ClipboardContent
	(*emptypb.Empty)(nil),    // 1: google.protobuf.Empty
}
var file_clipboard_proto_depIdxs = []int32{
	0, // 0: clipboard_rpc.Clipboard.ShareClipboard:input_type -> clipboard_rpc.ClipboardContent
	1, // 1: clipboard_rpc.Clipboard.ShareClipboard:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_clipboard_proto_init() }
func file_clipboard_proto_init() {
	if File_clipboard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_clipboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClipboardContent); i {
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
			RawDescriptor: file_clipboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_clipboard_proto_goTypes,
		DependencyIndexes: file_clipboard_proto_depIdxs,
		MessageInfos:      file_clipboard_proto_msgTypes,
	}.Build()
	File_clipboard_proto = out.File
	file_clipboard_proto_rawDesc = nil
	file_clipboard_proto_goTypes = nil
	file_clipboard_proto_depIdxs = nil
}