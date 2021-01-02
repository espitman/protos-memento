// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: merchandising/merchandising.proto

package merchandising

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

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"title",validate:"required"
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title" validate:"required"`
	// @inject_tag: json:"type",validate:"required,oneof=carousel movies videos"
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type" validate:"required,oneof=carousel movies videos"`
	// @inject_tag: json:"items",validate:"required,dive"
	Items []*RowItems `protobuf:"bytes,3,rep,name=items,proto3" json:"items" validate:"required,dive"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchandising_merchandising_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_merchandising_merchandising_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_merchandising_merchandising_proto_rawDescGZIP(), []int{0}
}

func (x *Row) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Row) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Row) GetItems() []*RowItems {
	if x != nil {
		return x.Items
	}
	return nil
}

type RowItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"type",validate:"required,oneof=movie person"
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type" validate:"required,oneof=movie person"`
	// @inject_tag: json:"id",validate:"required"
	Id int32 `protobuf:"varint,2,opt,name=id,proto3" json:"id" validate:"required"`
}

func (x *RowItems) Reset() {
	*x = RowItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchandising_merchandising_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RowItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RowItems) ProtoMessage() {}

func (x *RowItems) ProtoReflect() protoreflect.Message {
	mi := &file_merchandising_merchandising_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RowItems.ProtoReflect.Descriptor instead.
func (*RowItems) Descriptor() ([]byte, []int) {
	return file_merchandising_merchandising_proto_rawDescGZIP(), []int{1}
}

func (x *RowItems) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RowItems) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RequestCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"name",validate:"required"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" validate:"required"`
	// @inject_tag: json:"rows",validate:"required,dive"
	Rows []*Row `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows" validate:"required,dive"`
}

func (x *RequestCreate) Reset() {
	*x = RequestCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchandising_merchandising_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCreate) ProtoMessage() {}

func (x *RequestCreate) ProtoReflect() protoreflect.Message {
	mi := &file_merchandising_merchandising_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCreate.ProtoReflect.Descriptor instead.
func (*RequestCreate) Descriptor() ([]byte, []int) {
	return file_merchandising_merchandising_proto_rawDescGZIP(), []int{2}
}

func (x *RequestCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RequestCreate) GetRows() []*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

type RequestDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"name",validate:"required"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" validate:"required"`
}

func (x *RequestDetails) Reset() {
	*x = RequestDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchandising_merchandising_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDetails) ProtoMessage() {}

func (x *RequestDetails) ProtoReflect() protoreflect.Message {
	mi := &file_merchandising_merchandising_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDetails.ProtoReflect.Descriptor instead.
func (*RequestDetails) Descriptor() ([]byte, []int) {
	return file_merchandising_merchandising_proto_rawDescGZIP(), []int{3}
}

func (x *RequestDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ResponseDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ResponseDetails) Reset() {
	*x = ResponseDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchandising_merchandising_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseDetails) ProtoMessage() {}

func (x *ResponseDetails) ProtoReflect() protoreflect.Message {
	mi := &file_merchandising_merchandising_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseDetails.ProtoReflect.Descriptor instead.
func (*ResponseDetails) Descriptor() ([]byte, []int) {
	return file_merchandising_merchandising_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_merchandising_merchandising_proto protoreflect.FileDescriptor

var file_merchandising_merchandising_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x64, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x2f,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x64, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x64, 0x69, 0x73, 0x69,
	0x6e, 0x67, 0x22, 0x5e, 0x0a, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x64, 0x69, 0x73, 0x69,
	0x6e, 0x67, 0x2e, 0x52, 0x6f, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x2e, 0x0a, 0x08, 0x52, 0x6f, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x4b, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x64,
	0x69, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22,
	0x24, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xac, 0x01, 0x0a, 0x14, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x64, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x48, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x6d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x64, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x64, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x07, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x64, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x64,
	0x69, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_merchandising_merchandising_proto_rawDescOnce sync.Once
	file_merchandising_merchandising_proto_rawDescData = file_merchandising_merchandising_proto_rawDesc
)

func file_merchandising_merchandising_proto_rawDescGZIP() []byte {
	file_merchandising_merchandising_proto_rawDescOnce.Do(func() {
		file_merchandising_merchandising_proto_rawDescData = protoimpl.X.CompressGZIP(file_merchandising_merchandising_proto_rawDescData)
	})
	return file_merchandising_merchandising_proto_rawDescData
}

var file_merchandising_merchandising_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_merchandising_merchandising_proto_goTypes = []interface{}{
	(*Row)(nil),             // 0: merchandising.Row
	(*RowItems)(nil),        // 1: merchandising.RowItems
	(*RequestCreate)(nil),   // 2: merchandising.RequestCreate
	(*RequestDetails)(nil),  // 3: merchandising.RequestDetails
	(*ResponseDetails)(nil), // 4: merchandising.ResponseDetails
}
var file_merchandising_merchandising_proto_depIdxs = []int32{
	1, // 0: merchandising.Row.items:type_name -> merchandising.RowItems
	0, // 1: merchandising.RequestCreate.rows:type_name -> merchandising.Row
	2, // 2: merchandising.merchandisingService.Create:input_type -> merchandising.RequestCreate
	3, // 3: merchandising.merchandisingService.Details:input_type -> merchandising.RequestDetails
	4, // 4: merchandising.merchandisingService.Create:output_type -> merchandising.ResponseDetails
	4, // 5: merchandising.merchandisingService.Details:output_type -> merchandising.ResponseDetails
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_merchandising_merchandising_proto_init() }
func file_merchandising_merchandising_proto_init() {
	if File_merchandising_merchandising_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_merchandising_merchandising_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
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
		file_merchandising_merchandising_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RowItems); i {
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
		file_merchandising_merchandising_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCreate); i {
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
		file_merchandising_merchandising_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestDetails); i {
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
		file_merchandising_merchandising_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseDetails); i {
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
			RawDescriptor: file_merchandising_merchandising_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_merchandising_merchandising_proto_goTypes,
		DependencyIndexes: file_merchandising_merchandising_proto_depIdxs,
		MessageInfos:      file_merchandising_merchandising_proto_msgTypes,
	}.Build()
	File_merchandising_merchandising_proto = out.File
	file_merchandising_merchandising_proto_rawDesc = nil
	file_merchandising_merchandising_proto_goTypes = nil
	file_merchandising_merchandising_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MerchandisingServiceClient is the client API for MerchandisingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MerchandisingServiceClient interface {
	Create(ctx context.Context, in *RequestCreate, opts ...grpc.CallOption) (*ResponseDetails, error)
	Details(ctx context.Context, in *RequestDetails, opts ...grpc.CallOption) (*ResponseDetails, error)
}

type merchandisingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMerchandisingServiceClient(cc grpc.ClientConnInterface) MerchandisingServiceClient {
	return &merchandisingServiceClient{cc}
}

func (c *merchandisingServiceClient) Create(ctx context.Context, in *RequestCreate, opts ...grpc.CallOption) (*ResponseDetails, error) {
	out := new(ResponseDetails)
	err := c.cc.Invoke(ctx, "/merchandising.merchandisingService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchandisingServiceClient) Details(ctx context.Context, in *RequestDetails, opts ...grpc.CallOption) (*ResponseDetails, error) {
	out := new(ResponseDetails)
	err := c.cc.Invoke(ctx, "/merchandising.merchandisingService/Details", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MerchandisingServiceServer is the server API for MerchandisingService service.
type MerchandisingServiceServer interface {
	Create(context.Context, *RequestCreate) (*ResponseDetails, error)
	Details(context.Context, *RequestDetails) (*ResponseDetails, error)
}

// UnimplementedMerchandisingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMerchandisingServiceServer struct {
}

func (*UnimplementedMerchandisingServiceServer) Create(context.Context, *RequestCreate) (*ResponseDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedMerchandisingServiceServer) Details(context.Context, *RequestDetails) (*ResponseDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Details not implemented")
}

func RegisterMerchandisingServiceServer(s *grpc.Server, srv MerchandisingServiceServer) {
	s.RegisterService(&_MerchandisingService_serviceDesc, srv)
}

func _MerchandisingService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchandisingServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/merchandising.merchandisingService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchandisingServiceServer).Create(ctx, req.(*RequestCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchandisingService_Details_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchandisingServiceServer).Details(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/merchandising.merchandisingService/Details",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchandisingServiceServer).Details(ctx, req.(*RequestDetails))
	}
	return interceptor(ctx, in, info, handler)
}

var _MerchandisingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "merchandising.merchandisingService",
	HandlerType: (*MerchandisingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _MerchandisingService_Create_Handler,
		},
		{
			MethodName: "Details",
			Handler:    _MerchandisingService_Details_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "merchandising/merchandising.proto",
}
