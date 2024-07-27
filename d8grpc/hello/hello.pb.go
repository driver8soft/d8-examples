// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: hello/hello.proto

package hello

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

// The request message containing the user's name.
type MsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HelloName string `protobuf:"bytes,1,opt,name=hello_name,json=helloName,proto3" json:"hello_name,omitempty"`
}

func (x *MsgReq) Reset() {
	*x = MsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgReq) ProtoMessage() {}

func (x *MsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_hello_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgReq.ProtoReflect.Descriptor instead.
func (*MsgReq) Descriptor() ([]byte, []int) {
	return file_hello_hello_proto_rawDescGZIP(), []int{0}
}

func (x *MsgReq) GetHelloName() string {
	if x != nil {
		return x.HelloName
	}
	return ""
}

// The response message containing the greetings
type MsgRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *MsgRes) Reset() {
	*x = MsgRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgRes) ProtoMessage() {}

func (x *MsgRes) ProtoReflect() protoreflect.Message {
	mi := &file_hello_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgRes.ProtoReflect.Descriptor instead.
func (*MsgRes) Descriptor() ([]byte, []int) {
	return file_hello_hello_proto_rawDescGZIP(), []int{1}
}

func (x *MsgRes) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_hello_hello_proto protoreflect.FileDescriptor

var file_hello_hello_proto_rawDesc = []byte{
	0x0a, 0x11, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x22, 0x27, 0x0a, 0x06, 0x4d, 0x73,
	0x67, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x06, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x31, 0x0a, 0x06, 0x44, 0x38, 0x67,
	0x72, 0x70, 0x63, 0x12, 0x27, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0d, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12,
	0x64, 0x38, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hello_hello_proto_rawDescOnce sync.Once
	file_hello_hello_proto_rawDescData = file_hello_hello_proto_rawDesc
)

func file_hello_hello_proto_rawDescGZIP() []byte {
	file_hello_hello_proto_rawDescOnce.Do(func() {
		file_hello_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_hello_proto_rawDescData)
	})
	return file_hello_hello_proto_rawDescData
}

var file_hello_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_hello_hello_proto_goTypes = []interface{}{
	(*MsgReq)(nil), // 0: hello.MsgReq
	(*MsgRes)(nil), // 1: hello.MsgRes
}
var file_hello_hello_proto_depIdxs = []int32{
	0, // 0: hello.D8grpc.Hello:input_type -> hello.MsgReq
	1, // 1: hello.D8grpc.Hello:output_type -> hello.MsgRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hello_hello_proto_init() }
func file_hello_hello_proto_init() {
	if File_hello_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgReq); i {
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
		file_hello_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgRes); i {
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
			RawDescriptor: file_hello_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_hello_proto_goTypes,
		DependencyIndexes: file_hello_hello_proto_depIdxs,
		MessageInfos:      file_hello_hello_proto_msgTypes,
	}.Build()
	File_hello_hello_proto = out.File
	file_hello_hello_proto_rawDesc = nil
	file_hello_hello_proto_goTypes = nil
	file_hello_hello_proto_depIdxs = nil
}
