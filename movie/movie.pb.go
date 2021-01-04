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

type Companies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LogoPath      string `protobuf:"bytes,3,opt,name=logo_path,json=logoPath,proto3" json:"logo_path,omitempty"`
	OriginCountry string `protobuf:"bytes,4,opt,name=origin_country,json=originCountry,proto3" json:"origin_country,omitempty"`
}

func (x *Companies) Reset() {
	*x = Companies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Companies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Companies) ProtoMessage() {}

func (x *Companies) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Companies.ProtoReflect.Descriptor instead.
func (*Companies) Descriptor() ([]byte, []int) {
	return file_movie_movie_proto_rawDescGZIP(), []int{1}
}

func (x *Companies) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Companies) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Companies) GetLogoPath() string {
	if x != nil {
		return x.LogoPath
	}
	return ""
}

func (x *Companies) GetOriginCountry() string {
	if x != nil {
		return x.OriginCountry
	}
	return ""
}

type Countries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iso_3166_1 string `protobuf:"bytes,1,opt,name=iso_3166_1,json=iso31661,proto3" json:"iso_3166_1,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Countries) Reset() {
	*x = Countries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Countries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Countries) ProtoMessage() {}

func (x *Countries) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Countries.ProtoReflect.Descriptor instead.
func (*Countries) Descriptor() ([]byte, []int) {
	return file_movie_movie_proto_rawDescGZIP(), []int{2}
}

func (x *Countries) GetIso_3166_1() string {
	if x != nil {
		return x.Iso_3166_1
	}
	return ""
}

func (x *Countries) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Languages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnglishName string `protobuf:"bytes,1,opt,name=english_name,json=englishName,proto3" json:"english_name,omitempty"`
	Iso_639_1   string `protobuf:"bytes,2,opt,name=iso_639_1,json=iso6391,proto3" json:"iso_639_1,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Languages) Reset() {
	*x = Languages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_movie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Languages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Languages) ProtoMessage() {}

func (x *Languages) ProtoReflect() protoreflect.Message {
	mi := &file_movie_movie_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Languages.ProtoReflect.Descriptor instead.
func (*Languages) Descriptor() ([]byte, []int) {
	return file_movie_movie_proto_rawDescGZIP(), []int{3}
}

func (x *Languages) GetEnglishName() string {
	if x != nil {
		return x.EnglishName
	}
	return ""
}

func (x *Languages) GetIso_639_1() string {
	if x != nil {
		return x.Iso_639_1
	}
	return ""
}

func (x *Languages) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Images struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Backdrops []*Image `protobuf:"bytes,1,rep,name=backdrops,proto3" json:"backdrops,omitempty"`
	Posters   []*Image `protobuf:"bytes,2,rep,name=posters,proto3" json:"posters,omitempty"`
}

func (x *Images) Reset() {
	*x = Images{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_movie_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Images) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Images) ProtoMessage() {}

func (x *Images) ProtoReflect() protoreflect.Message {
	mi := &file_movie_movie_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Images.ProtoReflect.Descriptor instead.
func (*Images) Descriptor() ([]byte, []int) {
	return file_movie_movie_proto_rawDescGZIP(), []int{4}
}

func (x *Images) GetBackdrops() []*Image {
	if x != nil {
		return x.Backdrops
	}
	return nil
}

func (x *Images) GetPosters() []*Image {
	if x != nil {
		return x.Posters
	}
	return nil
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_movie_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_movie_movie_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_movie_movie_proto_rawDescGZIP(), []int{5}
}

func (x *Image) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type Credits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cast []*Person `protobuf:"bytes,1,rep,name=cast,proto3" json:"cast,omitempty"`
	Crew []*Person `protobuf:"bytes,2,rep,name=crew,proto3" json:"crew,omitempty"`
}

func (x *Credits) Reset() {
	*x = Credits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_movie_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credits) ProtoMessage() {}

func (x *Credits) ProtoReflect() protoreflect.Message {
	mi := &file_movie_movie_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credits.ProtoReflect.Descriptor instead.
func (*Credits) Descriptor() ([]byte, []int) {
	return file_movie_movie_proto_rawDescGZIP(), []int{6}
}

func (x *Credits) GetCast() []*Person {
	if x != nil {
		return x.Cast
	}
	return nil
}

func (x *Credits) GetCrew() []*Person {
	if x != nil {
		return x.Crew
	}
	return nil
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Gender       int32  `protobuf:"varint,2,opt,name=gender,proto3" json:"gender,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	OriginalName string `protobuf:"bytes,4,opt,name=original_name,json=originalName,proto3" json:"original_name,omitempty"`
	Image        string `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	// @inject_tag: json:"-"
	ProfilePath string `protobuf:"bytes,6,opt,name=profile_path,json=profilePath,proto3" json:"-"`
	Character   string `protobuf:"bytes,7,opt,name=character,proto3" json:"character,omitempty"`
	Job         string `protobuf:"bytes,8,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_movie_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_movie_movie_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_movie_movie_proto_rawDescGZIP(), []int{7}
}

func (x *Person) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Person) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetOriginalName() string {
	if x != nil {
		return x.OriginalName
	}
	return ""
}

func (x *Person) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Person) GetProfilePath() string {
	if x != nil {
		return x.ProfilePath
	}
	return ""
}

func (x *Person) GetCharacter() string {
	if x != nil {
		return x.Character
	}
	return ""
}

func (x *Person) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

type RequestDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id",validate:"required"
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id" validate:"required"`
	// @inject_tag: json:"full"
	Full bool `protobuf:"varint,2,opt,name=full,proto3" json:"full"`
}

func (x *RequestDetails) Reset() {
	*x = RequestDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_movie_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDetails) ProtoMessage() {}

func (x *RequestDetails) ProtoReflect() protoreflect.Message {
	mi := &file_movie_movie_proto_msgTypes[8]
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
	return file_movie_movie_proto_rawDescGZIP(), []int{8}
}

func (x *RequestDetails) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RequestDetails) GetFull() bool {
	if x != nil {
		return x.Full
	}
	return false
}

type ResponseDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title               string       `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	BackdropPath        string       `protobuf:"bytes,3,opt,name=backdrop_path,json=backdropPath,proto3" json:"backdrop_path,omitempty"`
	PosterPath          string       `protobuf:"bytes,4,opt,name=poster_path,json=posterPath,proto3" json:"poster_path,omitempty"`
	Genres              []*Genres    `protobuf:"bytes,5,rep,name=genres,proto3" json:"genres,omitempty"`
	Budget              int64        `protobuf:"varint,6,opt,name=budget,proto3" json:"budget,omitempty"`
	ImdbId              string       `protobuf:"bytes,7,opt,name=imdb_id,json=imdbId,proto3" json:"imdb_id,omitempty"`
	OriginalLanguage    string       `protobuf:"bytes,8,opt,name=original_language,json=originalLanguage,proto3" json:"original_language,omitempty"`
	OriginalTitle       string       `protobuf:"bytes,9,opt,name=original_title,json=originalTitle,proto3" json:"original_title,omitempty"`
	Overview            string       `protobuf:"bytes,10,opt,name=overview,proto3" json:"overview,omitempty"`
	ProductionCompanies []*Companies `protobuf:"bytes,11,rep,name=production_companies,json=productionCompanies,proto3" json:"production_companies,omitempty"`
	ProductionCountries []*Countries `protobuf:"bytes,12,rep,name=production_countries,json=productionCountries,proto3" json:"production_countries,omitempty"`
	ReleaseDate         string       `protobuf:"bytes,13,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
	Runtime             int32        `protobuf:"varint,14,opt,name=runtime,proto3" json:"runtime,omitempty"`
	SpokenLanguages     []*Languages `protobuf:"bytes,15,rep,name=spoken_languages,json=spokenLanguages,proto3" json:"spoken_languages,omitempty"`
	Status              string       `protobuf:"bytes,16,opt,name=status,proto3" json:"status,omitempty"`
	Tagline             string       `protobuf:"bytes,17,opt,name=tagline,proto3" json:"tagline,omitempty"`
	Images              *Images      `protobuf:"bytes,18,opt,name=images,proto3" json:"images,omitempty"`
	Credits             *Credits     `protobuf:"bytes,19,opt,name=credits,proto3" json:"credits,omitempty"`
}

func (x *ResponseDetails) Reset() {
	*x = ResponseDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_movie_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseDetails) ProtoMessage() {}

func (x *ResponseDetails) ProtoReflect() protoreflect.Message {
	mi := &file_movie_movie_proto_msgTypes[9]
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
	return file_movie_movie_proto_rawDescGZIP(), []int{9}
}

func (x *ResponseDetails) GetId() int32 {
	if x != nil {
		return x.Id
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

func (x *ResponseDetails) GetBudget() int64 {
	if x != nil {
		return x.Budget
	}
	return 0
}

func (x *ResponseDetails) GetImdbId() string {
	if x != nil {
		return x.ImdbId
	}
	return ""
}

func (x *ResponseDetails) GetOriginalLanguage() string {
	if x != nil {
		return x.OriginalLanguage
	}
	return ""
}

func (x *ResponseDetails) GetOriginalTitle() string {
	if x != nil {
		return x.OriginalTitle
	}
	return ""
}

func (x *ResponseDetails) GetOverview() string {
	if x != nil {
		return x.Overview
	}
	return ""
}

func (x *ResponseDetails) GetProductionCompanies() []*Companies {
	if x != nil {
		return x.ProductionCompanies
	}
	return nil
}

func (x *ResponseDetails) GetProductionCountries() []*Countries {
	if x != nil {
		return x.ProductionCountries
	}
	return nil
}

func (x *ResponseDetails) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

func (x *ResponseDetails) GetRuntime() int32 {
	if x != nil {
		return x.Runtime
	}
	return 0
}

func (x *ResponseDetails) GetSpokenLanguages() []*Languages {
	if x != nil {
		return x.SpokenLanguages
	}
	return nil
}

func (x *ResponseDetails) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ResponseDetails) GetTagline() string {
	if x != nil {
		return x.Tagline
	}
	return ""
}

func (x *ResponseDetails) GetImages() *Images {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *ResponseDetails) GetCredits() *Credits {
	if x != nil {
		return x.Credits
	}
	return nil
}

var File_movie_movie_proto protoreflect.FileDescriptor

var file_movie_movie_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x2c, 0x0a, 0x06, 0x47, 0x65,
	0x6e, 0x72, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x73, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x69, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67,
	0x6f, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x67, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x3d, 0x0a,
	0x09, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x0a, 0x69, 0x73,
	0x6f, 0x5f, 0x33, 0x31, 0x36, 0x36, 0x5f, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x73, 0x6f, 0x33, 0x31, 0x36, 0x36, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5e, 0x0a, 0x09,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6e, 0x67,
	0x6c, 0x69, 0x73, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x09,
	0x69, 0x73, 0x6f, 0x5f, 0x36, 0x33, 0x39, 0x5f, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x69, 0x73, 0x6f, 0x36, 0x33, 0x39, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5c, 0x0a, 0x06,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x64, 0x72,
	0x6f, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x6f,
	0x70, 0x73, 0x12, 0x26, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x24, 0x0a, 0x05, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x22, 0x4f, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x04, 0x63,
	0x61, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x04, 0x63, 0x61, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x04, 0x63, 0x72, 0x65, 0x77, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x04, 0x63, 0x72, 0x65,
	0x77, 0x22, 0xd2, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x22, 0x34, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x75, 0x6c, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x66, 0x75, 0x6c, 0x6c, 0x22, 0xcc, 0x05, 0x0a,
	0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x64, 0x72,
	0x6f, 0x70, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62,
	0x61, 0x63, 0x6b, 0x64, 0x72, 0x6f, 0x70, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x6f, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x72, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x6d, 0x64, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d,
	0x64, 0x62, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x76, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x12, 0x43, 0x0a, 0x14, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x69, 0x65, 0x73, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x14, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x10, 0x73,
	0x70, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52, 0x0f, 0x73, 0x70, 0x6f, 0x6b, 0x65, 0x6e, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x67, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x28, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x73, 0x52, 0x07, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x32, 0x4a, 0x0a, 0x0c, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x15, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x16, 0x2e,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_movie_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_movie_movie_proto_goTypes = []interface{}{
	(*Genres)(nil),          // 0: movie.Genres
	(*Companies)(nil),       // 1: movie.Companies
	(*Countries)(nil),       // 2: movie.Countries
	(*Languages)(nil),       // 3: movie.Languages
	(*Images)(nil),          // 4: movie.Images
	(*Image)(nil),           // 5: movie.Image
	(*Credits)(nil),         // 6: movie.Credits
	(*Person)(nil),          // 7: movie.Person
	(*RequestDetails)(nil),  // 8: movie.RequestDetails
	(*ResponseDetails)(nil), // 9: movie.ResponseDetails
}
var file_movie_movie_proto_depIdxs = []int32{
	5,  // 0: movie.Images.backdrops:type_name -> movie.Image
	5,  // 1: movie.Images.posters:type_name -> movie.Image
	7,  // 2: movie.Credits.cast:type_name -> movie.Person
	7,  // 3: movie.Credits.crew:type_name -> movie.Person
	0,  // 4: movie.ResponseDetails.genres:type_name -> movie.Genres
	1,  // 5: movie.ResponseDetails.production_companies:type_name -> movie.Companies
	2,  // 6: movie.ResponseDetails.production_countries:type_name -> movie.Countries
	3,  // 7: movie.ResponseDetails.spoken_languages:type_name -> movie.Languages
	4,  // 8: movie.ResponseDetails.images:type_name -> movie.Images
	6,  // 9: movie.ResponseDetails.credits:type_name -> movie.Credits
	8,  // 10: movie.movieService.Details:input_type -> movie.RequestDetails
	9,  // 11: movie.movieService.Details:output_type -> movie.ResponseDetails
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
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
			switch v := v.(*Companies); i {
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
			switch v := v.(*Countries); i {
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
		file_movie_movie_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Languages); i {
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
		file_movie_movie_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Images); i {
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
		file_movie_movie_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_movie_movie_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credits); i {
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
		file_movie_movie_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_movie_movie_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_movie_movie_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   10,
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
