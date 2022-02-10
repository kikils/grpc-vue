// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.2
// source: count.proto

package proto

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

type GetTotalNumParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTotalNumParams) Reset() {
	*x = GetTotalNumParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_count_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTotalNumParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTotalNumParams) ProtoMessage() {}

func (x *GetTotalNumParams) ProtoReflect() protoreflect.Message {
	mi := &file_count_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTotalNumParams.ProtoReflect.Descriptor instead.
func (*GetTotalNumParams) Descriptor() ([]byte, []int) {
	return file_count_proto_rawDescGZIP(), []int{0}
}

type AddNumParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *AddNumParams) Reset() {
	*x = AddNumParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_count_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNumParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNumParams) ProtoMessage() {}

func (x *AddNumParams) ProtoReflect() protoreflect.Message {
	mi := &file_count_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNumParams.ProtoReflect.Descriptor instead.
func (*AddNumParams) Descriptor() ([]byte, []int) {
	return file_count_proto_rawDescGZIP(), []int{1}
}

func (x *AddNumParams) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type TotalNum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *TotalNum) Reset() {
	*x = TotalNum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_count_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TotalNum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TotalNum) ProtoMessage() {}

func (x *TotalNum) ProtoReflect() protoreflect.Message {
	mi := &file_count_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TotalNum.ProtoReflect.Descriptor instead.
func (*TotalNum) Descriptor() ([]byte, []int) {
	return file_count_proto_rawDescGZIP(), []int{2}
}

func (x *TotalNum) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_count_proto protoreflect.FileDescriptor

var file_count_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x13, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x4e, 0x75, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x26, 0x0a, 0x0c, 0x61, 0x64,
	0x64, 0x4e, 0x75, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x20, 0x0a, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x32, 0x81, 0x01, 0x0a, 0x0d, 0x61, 0x64, 0x64, 0x4e, 0x75, 0x6d, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x61, 0x64, 0x64, 0x4e, 0x75, 0x6d,
	0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x64, 0x4e, 0x75, 0x6d,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b, 0x67, 0x65,
	0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x1a, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_count_proto_rawDescOnce sync.Once
	file_count_proto_rawDescData = file_count_proto_rawDesc
)

func file_count_proto_rawDescGZIP() []byte {
	file_count_proto_rawDescOnce.Do(func() {
		file_count_proto_rawDescData = protoimpl.X.CompressGZIP(file_count_proto_rawDescData)
	})
	return file_count_proto_rawDescData
}

var file_count_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_count_proto_goTypes = []interface{}{
	(*GetTotalNumParams)(nil), // 0: server.getTotalNumParams
	(*AddNumParams)(nil),      // 1: server.addNumParams
	(*TotalNum)(nil),          // 2: server.totalNum
}
var file_count_proto_depIdxs = []int32{
	1, // 0: server.addNumService.addNum:input_type -> server.addNumParams
	0, // 1: server.addNumService.getTotalNum:input_type -> server.getTotalNumParams
	2, // 2: server.addNumService.addNum:output_type -> server.totalNum
	2, // 3: server.addNumService.getTotalNum:output_type -> server.totalNum
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_count_proto_init() }
func file_count_proto_init() {
	if File_count_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_count_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTotalNumParams); i {
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
		file_count_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNumParams); i {
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
		file_count_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TotalNum); i {
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
			RawDescriptor: file_count_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_count_proto_goTypes,
		DependencyIndexes: file_count_proto_depIdxs,
		MessageInfos:      file_count_proto_msgTypes,
	}.Build()
	File_count_proto = out.File
	file_count_proto_rawDesc = nil
	file_count_proto_goTypes = nil
	file_count_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AddNumServiceClient is the client API for AddNumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddNumServiceClient interface {
	AddNum(ctx context.Context, in *AddNumParams, opts ...grpc.CallOption) (*TotalNum, error)
	GetTotalNum(ctx context.Context, in *GetTotalNumParams, opts ...grpc.CallOption) (*TotalNum, error)
}

type addNumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAddNumServiceClient(cc grpc.ClientConnInterface) AddNumServiceClient {
	return &addNumServiceClient{cc}
}

func (c *addNumServiceClient) AddNum(ctx context.Context, in *AddNumParams, opts ...grpc.CallOption) (*TotalNum, error) {
	out := new(TotalNum)
	err := c.cc.Invoke(ctx, "/server.addNumService/addNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addNumServiceClient) GetTotalNum(ctx context.Context, in *GetTotalNumParams, opts ...grpc.CallOption) (*TotalNum, error) {
	out := new(TotalNum)
	err := c.cc.Invoke(ctx, "/server.addNumService/getTotalNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddNumServiceServer is the server API for AddNumService service.
type AddNumServiceServer interface {
	AddNum(context.Context, *AddNumParams) (*TotalNum, error)
	GetTotalNum(context.Context, *GetTotalNumParams) (*TotalNum, error)
}

// UnimplementedAddNumServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAddNumServiceServer struct {
}

func (*UnimplementedAddNumServiceServer) AddNum(context.Context, *AddNumParams) (*TotalNum, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNum not implemented")
}
func (*UnimplementedAddNumServiceServer) GetTotalNum(context.Context, *GetTotalNumParams) (*TotalNum, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalNum not implemented")
}

func RegisterAddNumServiceServer(s *grpc.Server, srv AddNumServiceServer) {
	s.RegisterService(&_AddNumService_serviceDesc, srv)
}

func _AddNumService_AddNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNumParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddNumServiceServer).AddNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.addNumService/AddNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddNumServiceServer).AddNum(ctx, req.(*AddNumParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddNumService_GetTotalNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTotalNumParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddNumServiceServer).GetTotalNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.addNumService/GetTotalNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddNumServiceServer).GetTotalNum(ctx, req.(*GetTotalNumParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddNumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.addNumService",
	HandlerType: (*AddNumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addNum",
			Handler:    _AddNumService_AddNum_Handler,
		},
		{
			MethodName: "getTotalNum",
			Handler:    _AddNumService_GetTotalNum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "count.proto",
}