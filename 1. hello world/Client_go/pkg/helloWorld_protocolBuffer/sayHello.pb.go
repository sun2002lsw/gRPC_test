// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: sayHello.proto

package helloWorld_protocolBuffer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sayHello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_sayHello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_sayHello_proto_rawDescGZIP(), []int{0}
}

type HelloMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloMsg) Reset() {
	*x = HelloMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sayHello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloMsg) ProtoMessage() {}

func (x *HelloMsg) ProtoReflect() protoreflect.Message {
	mi := &file_sayHello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloMsg.ProtoReflect.Descriptor instead.
func (*HelloMsg) Descriptor() ([]byte, []int) {
	return file_sayHello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_sayHello_proto protoreflect.FileDescriptor

var file_sayHello_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1c, 0x0a, 0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4d, 0x73,
	0x67, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x32, 0x4e, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12,
	0x42, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x4d, 0x73, 0x67, 0x42, 0x1c, 0x5a, 0x1a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c,
	0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sayHello_proto_rawDescOnce sync.Once
	file_sayHello_proto_rawDescData = file_sayHello_proto_rawDesc
)

func file_sayHello_proto_rawDescGZIP() []byte {
	file_sayHello_proto_rawDescOnce.Do(func() {
		file_sayHello_proto_rawDescData = protoimpl.X.CompressGZIP(file_sayHello_proto_rawDescData)
	})
	return file_sayHello_proto_rawDescData
}

var file_sayHello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sayHello_proto_goTypes = []interface{}{
	(*Empty)(nil),    // 0: helloWorld.Empty
	(*HelloMsg)(nil), // 1: helloWorld.HelloMsg
}
var file_sayHello_proto_depIdxs = []int32{
	0, // 0: helloWorld.SayHello.GetHelloWorldFromServer:input_type -> helloWorld.Empty
	1, // 1: helloWorld.SayHello.GetHelloWorldFromServer:output_type -> helloWorld.HelloMsg
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sayHello_proto_init() }
func file_sayHello_proto_init() {
	if File_sayHello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sayHello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_sayHello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloMsg); i {
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
			RawDescriptor: file_sayHello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sayHello_proto_goTypes,
		DependencyIndexes: file_sayHello_proto_depIdxs,
		MessageInfos:      file_sayHello_proto_msgTypes,
	}.Build()
	File_sayHello_proto = out.File
	file_sayHello_proto_rawDesc = nil
	file_sayHello_proto_goTypes = nil
	file_sayHello_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SayHelloClient is the client API for SayHello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SayHelloClient interface {
	GetHelloWorldFromServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HelloMsg, error)
}

type sayHelloClient struct {
	cc grpc.ClientConnInterface
}

func NewSayHelloClient(cc grpc.ClientConnInterface) SayHelloClient {
	return &sayHelloClient{cc}
}

func (c *sayHelloClient) GetHelloWorldFromServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HelloMsg, error) {
	out := new(HelloMsg)
	err := c.cc.Invoke(ctx, "/helloWorld.SayHello/GetHelloWorldFromServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SayHelloServer is the server API for SayHello service.
type SayHelloServer interface {
	GetHelloWorldFromServer(context.Context, *Empty) (*HelloMsg, error)
}

// UnimplementedSayHelloServer can be embedded to have forward compatible implementations.
type UnimplementedSayHelloServer struct {
}

func (*UnimplementedSayHelloServer) GetHelloWorldFromServer(context.Context, *Empty) (*HelloMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHelloWorldFromServer not implemented")
}

func RegisterSayHelloServer(s *grpc.Server, srv SayHelloServer) {
	s.RegisterService(&_SayHello_serviceDesc, srv)
}

func _SayHello_GetHelloWorldFromServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SayHelloServer).GetHelloWorldFromServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloWorld.SayHello/GetHelloWorldFromServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SayHelloServer).GetHelloWorldFromServer(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _SayHello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloWorld.SayHello",
	HandlerType: (*SayHelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHelloWorldFromServer",
			Handler:    _SayHello_GetHelloWorldFromServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sayHello.proto",
}
