// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: cmdService.proto

package pbs

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{0}
}

type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{1}
}

func (x *CommonResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type LogLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module string `protobuf:"bytes,1,opt,name=Module,proto3" json:"Module,omitempty"`
	Level  int32  `protobuf:"varint,2,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (x *LogLevel) Reset() {
	*x = LogLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogLevel) ProtoMessage() {}

func (x *LogLevel) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogLevel.ProtoReflect.Descriptor instead.
func (*LogLevel) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{2}
}

func (x *LogLevel) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *LogLevel) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{3}
}

func (x *Address) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type VpnSrvParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pool  string `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool,omitempty"`
	Miner string `protobuf:"bytes,2,opt,name=miner,proto3" json:"miner,omitempty"`
}

func (x *VpnSrvParam) Reset() {
	*x = VpnSrvParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VpnSrvParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VpnSrvParam) ProtoMessage() {}

func (x *VpnSrvParam) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VpnSrvParam.ProtoReflect.Descriptor instead.
func (*VpnSrvParam) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{4}
}

func (x *VpnSrvParam) GetPool() string {
	if x != nil {
		return x.Pool
	}
	return ""
}

func (x *VpnSrvParam) GetMiner() string {
	if x != nil {
		return x.Miner
	}
	return ""
}

var File_cmdService_proto protoreflect.FileDescriptor

var file_cmdService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6d, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x70, 0x62, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x22, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x38, 0x0a, 0x08, 0x4c,
	0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x23, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x37, 0x0a, 0x0b, 0x56, 0x70,
	0x6e, 0x53, 0x72, 0x76, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6f,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x69,
	0x6e, 0x65, 0x72, 0x32, 0xa0, 0x02, 0x0a, 0x0a, 0x43, 0x6d, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x77, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x12, 0x11, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x08, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x56, 0x70, 0x6e, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x56, 0x70, 0x6e,
	0x53, 0x72, 0x76, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x33, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x70, 0x56, 0x70, 0x6e, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x73,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x70, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x77, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x56, 0x70, 0x6e, 0x53, 0x72,
	0x76, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x0b, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x75, 0x62, 0x70, 0x6f, 0x6f, 0x6c, 0x12, 0x11, 0x2e, 0x70,
	0x62, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x70, 0x62, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmdService_proto_rawDescOnce sync.Once
	file_cmdService_proto_rawDescData = file_cmdService_proto_rawDesc
)

func file_cmdService_proto_rawDescGZIP() []byte {
	file_cmdService_proto_rawDescOnce.Do(func() {
		file_cmdService_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmdService_proto_rawDescData)
	})
	return file_cmdService_proto_rawDescData
}

var file_cmdService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cmdService_proto_goTypes = []interface{}{
	(*EmptyRequest)(nil),   // 0: pbs.EmptyRequest
	(*CommonResponse)(nil), // 1: pbs.CommonResponse
	(*LogLevel)(nil),       // 2: pbs.LogLevel
	(*Address)(nil),        // 3: pbs.Address
	(*VpnSrvParam)(nil),    // 4: pbs.VpnSrvParam
}
var file_cmdService_proto_depIdxs = []int32{
	0, // 0: pbs.CmdService.ShowWallet:input_type -> pbs.EmptyRequest
	4, // 1: pbs.CmdService.StartVpn:input_type -> pbs.VpnSrvParam
	0, // 2: pbs.CmdService.StopVpn:input_type -> pbs.EmptyRequest
	4, // 3: pbs.CmdService.ShowUserData:input_type -> pbs.VpnSrvParam
	0, // 4: pbs.CmdService.ShowSubpool:input_type -> pbs.EmptyRequest
	1, // 5: pbs.CmdService.ShowWallet:output_type -> pbs.CommonResponse
	1, // 6: pbs.CmdService.StartVpn:output_type -> pbs.CommonResponse
	1, // 7: pbs.CmdService.StopVpn:output_type -> pbs.CommonResponse
	1, // 8: pbs.CmdService.ShowUserData:output_type -> pbs.CommonResponse
	1, // 9: pbs.CmdService.ShowSubpool:output_type -> pbs.CommonResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmdService_proto_init() }
func file_cmdService_proto_init() {
	if File_cmdService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmdService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_cmdService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResponse); i {
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
		file_cmdService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogLevel); i {
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
		file_cmdService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_cmdService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VpnSrvParam); i {
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
			RawDescriptor: file_cmdService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmdService_proto_goTypes,
		DependencyIndexes: file_cmdService_proto_depIdxs,
		MessageInfos:      file_cmdService_proto_msgTypes,
	}.Build()
	File_cmdService_proto = out.File
	file_cmdService_proto_rawDesc = nil
	file_cmdService_proto_goTypes = nil
	file_cmdService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CmdServiceClient is the client API for CmdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CmdServiceClient interface {
	ShowWallet(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	StartVpn(ctx context.Context, in *VpnSrvParam, opts ...grpc.CallOption) (*CommonResponse, error)
	StopVpn(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	ShowUserData(ctx context.Context, in *VpnSrvParam, opts ...grpc.CallOption) (*CommonResponse, error)
	ShowSubpool(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CommonResponse, error)
}

type cmdServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCmdServiceClient(cc grpc.ClientConnInterface) CmdServiceClient {
	return &cmdServiceClient{cc}
}

func (c *cmdServiceClient) ShowWallet(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pbs.CmdService/ShowWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdServiceClient) StartVpn(ctx context.Context, in *VpnSrvParam, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pbs.CmdService/StartVpn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdServiceClient) StopVpn(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pbs.CmdService/StopVpn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdServiceClient) ShowUserData(ctx context.Context, in *VpnSrvParam, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pbs.CmdService/ShowUserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdServiceClient) ShowSubpool(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pbs.CmdService/ShowSubpool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CmdServiceServer is the server API for CmdService service.
type CmdServiceServer interface {
	ShowWallet(context.Context, *EmptyRequest) (*CommonResponse, error)
	StartVpn(context.Context, *VpnSrvParam) (*CommonResponse, error)
	StopVpn(context.Context, *EmptyRequest) (*CommonResponse, error)
}

// UnimplementedCmdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCmdServiceServer struct {
}

func (*UnimplementedCmdServiceServer) ShowWallet(context.Context, *EmptyRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowWallet not implemented")
}
func (*UnimplementedCmdServiceServer) StartVpn(context.Context, *VpnSrvParam) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartVpn not implemented")
}
func (*UnimplementedCmdServiceServer) StopVpn(context.Context, *EmptyRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopVpn not implemented")
}
func (*UnimplementedCmdServiceServer) ShowUserData(context.Context, *VpnSrvParam) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowUserData not implemented")
}
func (*UnimplementedCmdServiceServer) ShowSubpool(context.Context, *EmptyRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowSubpool not implemented")
}

func RegisterCmdServiceServer(s *grpc.Server, srv CmdServiceServer) {
	s.RegisterService(&_CmdService_serviceDesc, srv)
}

func _CmdService_ShowWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServiceServer).ShowWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbs.CmdService/ShowWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServiceServer).ShowWallet(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmdService_StartVpn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VpnSrvParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServiceServer).StartVpn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbs.CmdService/StartVpn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServiceServer).StartVpn(ctx, req.(*VpnSrvParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmdService_StopVpn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServiceServer).StopVpn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbs.CmdService/StopVpn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServiceServer).StopVpn(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmdService_ShowUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VpnSrvParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServiceServer).ShowUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbs.CmdService/ShowUserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServiceServer).ShowUserData(ctx, req.(*VpnSrvParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmdService_ShowSubpool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServiceServer).ShowSubpool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbs.CmdService/ShowSubpool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServiceServer).ShowSubpool(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CmdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbs.CmdService",
	HandlerType: (*CmdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShowWallet",
			Handler:    _CmdService_ShowWallet_Handler,
		},
		{
			MethodName: "StartVpn",
			Handler:    _CmdService_StartVpn_Handler,
		},
		{
			MethodName: "StopVpn",
			Handler:    _CmdService_StopVpn_Handler,
		},
		{
			MethodName: "ShowUserData",
			Handler:    _CmdService_ShowUserData_Handler,
		},
		{
			MethodName: "ShowSubpool",
			Handler:    _CmdService_ShowSubpool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmdService.proto",
}
