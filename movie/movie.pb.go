// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: movie/movie.proto

package movie

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

type Genres struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Genres) Reset() {
	*x = Genres{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Genres) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genres) ProtoMessage() {}

func (x *Genres) ProtoReflect() protoreflect.Message {
	mi := &file_movie_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genres.ProtoReflect.Descriptor instead.
func (*Genres) Descriptor() ([]byte, []int) {
	return file_movie_movie_proto_rawDescGZIP(), []int{0}
}

func (x *Genres) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Genres) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RequestDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id",validate:"required"
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id" validate:"required"`
}

func (x *RequestDetails) Reset() {
	*x = RequestDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDetails) ProtoMessage() {}

func (x *RequestDetails) ProtoReflect() protoreflect.Message {
	mi := &file_movie_movie_proto_msgTypes[1]
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
	return file_movie_movie_proto_rawDescGZIP(), []int{1}
}

func (x *RequestDetails) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ResponseDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"title"
	Title string `protobuf:"bytes,2,opt,name=Title,proto3" json:"title"`
	// @inject_tag: json:"backdrop_path"
	BackdropPath string `protobuf:"bytes,3,opt,name=Backdrop_path,json=BackdropPath,proto3" json:"backdrop_path"`
	// @inject_tag: json:"poster_path"
	PosterPath string `protobuf:"bytes,4,opt,name=Poster_path,json=PosterPath,proto3" json:"poster_path"`
	// @inject_tag: json:"genres"
	Genres []*Genres `protobuf:"bytes,5,rep,name=Genres,proto3" json:"genres"`
}

func (x *ResponseDetails) Reset() {
	*x = ResponseDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseDetails) ProtoMessage() {}

func (x *ResponseDetails) ProtoReflect() protoreflect.Message {
	mi := &file_movie_movie_proto_msgTypes[2]
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
	return file_movie_movie_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseDetails) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ResponseDetails) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ResponseDetails) GetBackdropPath() string {
	if x != nil {
		return x.BackdropPath
	}
	return ""
}

func (x *ResponseDetails) GetPosterPath() string {
	if x != nil {
		return x.PosterPath
	}
	return ""
}

func (x *ResponseDetails) GetGenres() []*Genres {
	if x != nil {
		return x.Genres
	}
	return nil
}

var File_movie_movie_proto protoreflect.FileDescriptor

var file_movie_movie_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x2c, 0x0a, 0x06, 0x47, 0x65,
	0x6e, 0x72, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa4, 0x01, 0x0a, 0x0f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x42, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x6f, 0x70,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x42, 0x61, 0x63,
	0x6b, 0x64, 0x72, 0x6f, 0x70, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x50, 0x6f, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x06, 0x47, 0x65,
	0x6e, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x72, 0x65,
	0x73, 0x32, 0x4a, 0x0a, 0x0c, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3a, 0x0a, 0x07, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x15, 0x2e, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x1a, 0x16, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_movie_movie_proto_rawDescOnce sync.Once
	file_movie_movie_proto_rawDescData = file_movie_movie_proto_rawDesc
)

func file_movie_movie_proto_rawDescGZIP() []byte {
	file_movie_movie_proto_rawDescOnce.Do(func() {
		file_movie_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_movie_movie_proto_rawDescData)
	})
	return file_movie_movie_proto_rawDescData
}

var file_movie_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_movie_movie_proto_goTypes = []interface{}{
	(*Genres)(nil),          // 0: movie.Genres
	(*RequestDetails)(nil),  // 1: movie.RequestDetails
	(*ResponseDetails)(nil), // 2: movie.ResponseDetails
}
var file_movie_movie_proto_depIdxs = []int32{
	0, // 0: movie.ResponseDetails.Genres:type_name -> movie.Genres
	1, // 1: movie.movieService.Details:input_type -> movie.RequestDetails
	2, // 2: movie.movieService.Details:output_type -> movie.ResponseDetails
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_movie_movie_proto_init() }
func file_movie_movie_proto_init() {
	if File_movie_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_movie_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Genres); i {
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
		file_movie_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_movie_movie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_movie_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_movie_movie_proto_goTypes,
		DependencyIndexes: file_movie_movie_proto_depIdxs,
		MessageInfos:      file_movie_movie_proto_msgTypes,
	}.Build()
	File_movie_movie_proto = out.File
	file_movie_movie_proto_rawDesc = nil
	file_movie_movie_proto_goTypes = nil
	file_movie_movie_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MovieServiceClient is the client API for MovieService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MovieServiceClient interface {
	Details(ctx context.Context, in *RequestDetails, opts ...grpc.CallOption) (*ResponseDetails, error)
}

type movieServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieServiceClient(cc grpc.ClientConnInterface) MovieServiceClient {
	return &movieServiceClient{cc}
}

func (c *movieServiceClient) Details(ctx context.Context, in *RequestDetails, opts ...grpc.CallOption) (*ResponseDetails, error) {
	out := new(ResponseDetails)
	err := c.cc.Invoke(ctx, "/movie.movieService/Details", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieServiceServer is the server API for MovieService service.
type MovieServiceServer interface {
	Details(context.Context, *RequestDetails) (*ResponseDetails, error)
}

// UnimplementedMovieServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMovieServiceServer struct {
}

func (*UnimplementedMovieServiceServer) Details(context.Context, *RequestDetails) (*ResponseDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Details not implemented")
}

func RegisterMovieServiceServer(s *grpc.Server, srv MovieServiceServer) {
	s.RegisterService(&_MovieService_serviceDesc, srv)
}

func _MovieService_Details_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).Details(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/movie.movieService/Details",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).Details(ctx, req.(*RequestDetails))
	}
	return interceptor(ctx, in, info, handler)
}

var _MovieService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "movie.movieService",
	HandlerType: (*MovieServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Details",
			Handler:    _MovieService_Details_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movie/movie.proto",
}
