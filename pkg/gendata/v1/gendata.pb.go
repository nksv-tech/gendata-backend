// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.3
// source: gendata/v1/gendata.proto

package gendata

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RenderedFiles []*RenderedFile    `protobuf:"bytes,1,rep,name=renderedFiles,proto3" json:"renderedFiles,omitempty"`
	RenderTime    *duration.Duration `protobuf:"bytes,2,opt,name=renderTime,proto3" json:"renderTime,omitempty"`
}

func (x *GenResponse) Reset() {
	*x = GenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gendata_v1_gendata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenResponse) ProtoMessage() {}

func (x *GenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gendata_v1_gendata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenResponse.ProtoReflect.Descriptor instead.
func (*GenResponse) Descriptor() ([]byte, []int) {
	return file_gendata_v1_gendata_proto_rawDescGZIP(), []int{0}
}

func (x *GenResponse) GetRenderedFiles() []*RenderedFile {
	if x != nil {
		return x.RenderedFiles
	}
	return nil
}

func (x *GenResponse) GetRenderTime() *duration.Duration {
	if x != nil {
		return x.RenderTime
	}
	return nil
}

type RenderedFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content  []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	FileName string `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
}

func (x *RenderedFile) Reset() {
	*x = RenderedFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gendata_v1_gendata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenderedFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenderedFile) ProtoMessage() {}

func (x *RenderedFile) ProtoReflect() protoreflect.Message {
	mi := &file_gendata_v1_gendata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenderedFile.ProtoReflect.Descriptor instead.
func (*RenderedFile) Descriptor() ([]byte, []int) {
	return file_gendata_v1_gendata_proto_rawDescGZIP(), []int{1}
}

func (x *RenderedFile) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *RenderedFile) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type GenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tmpl   []byte  `protobuf:"bytes,1,opt,name=tmpl,proto3" json:"tmpl,omitempty"`
	Data   []byte  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Config *Config `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *GenRequest) Reset() {
	*x = GenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gendata_v1_gendata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenRequest) ProtoMessage() {}

func (x *GenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gendata_v1_gendata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenRequest.ProtoReflect.Descriptor instead.
func (*GenRequest) Descriptor() ([]byte, []int) {
	return file_gendata_v1_gendata_proto_rawDescGZIP(), []int{2}
}

func (x *GenRequest) GetTmpl() []byte {
	if x != nil {
		return x.Tmpl
	}
	return nil
}

func (x *GenRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GenRequest) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LangSettings    *LangSettings `protobuf:"bytes,1,opt,name=langSettings,proto3" json:"langSettings,omitempty"`
	Lang            string        `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	DataFormat      string        `protobuf:"bytes,3,opt,name=dataFormat,proto3" json:"dataFormat,omitempty"`
	RootClassName   string        `protobuf:"bytes,4,opt,name=rootClassName,proto3" json:"rootClassName,omitempty"`
	PrefixClassName string        `protobuf:"bytes,5,opt,name=prefixClassName,proto3" json:"prefixClassName,omitempty"`
	SuffixClassName string        `protobuf:"bytes,6,opt,name=suffixClassName,proto3" json:"suffixClassName,omitempty"`
	SortProperties  bool          `protobuf:"varint,7,opt,name=sort_properties,json=sortProperties,proto3" json:"sort_properties,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gendata_v1_gendata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_gendata_v1_gendata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_gendata_v1_gendata_proto_rawDescGZIP(), []int{3}
}

func (x *Config) GetLangSettings() *LangSettings {
	if x != nil {
		return x.LangSettings
	}
	return nil
}

func (x *Config) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *Config) GetDataFormat() string {
	if x != nil {
		return x.DataFormat
	}
	return ""
}

func (x *Config) GetRootClassName() string {
	if x != nil {
		return x.RootClassName
	}
	return ""
}

func (x *Config) GetPrefixClassName() string {
	if x != nil {
		return x.PrefixClassName
	}
	return ""
}

func (x *Config) GetSuffixClassName() string {
	if x != nil {
		return x.SuffixClassName
	}
	return ""
}

func (x *Config) GetSortProperties() bool {
	if x != nil {
		return x.SortProperties
	}
	return false
}

type LangSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code               string         `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name               string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FileExtension      string         `protobuf:"bytes,3,opt,name=fileExtension,proto3" json:"fileExtension,omitempty"`
	SplitObjectByFiles bool           `protobuf:"varint,4,opt,name=splitObjectByFiles,proto3" json:"splitObjectByFiles,omitempty"`
	ConfigMapping      *ConfigMapping `protobuf:"bytes,5,opt,name=configMapping,proto3" json:"configMapping,omitempty"`
}

func (x *LangSettings) Reset() {
	*x = LangSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gendata_v1_gendata_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LangSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LangSettings) ProtoMessage() {}

func (x *LangSettings) ProtoReflect() protoreflect.Message {
	mi := &file_gendata_v1_gendata_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LangSettings.ProtoReflect.Descriptor instead.
func (*LangSettings) Descriptor() ([]byte, []int) {
	return file_gendata_v1_gendata_proto_rawDescGZIP(), []int{4}
}

func (x *LangSettings) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *LangSettings) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LangSettings) GetFileExtension() string {
	if x != nil {
		return x.FileExtension
	}
	return ""
}

func (x *LangSettings) GetSplitObjectByFiles() bool {
	if x != nil {
		return x.SplitObjectByFiles
	}
	return false
}

func (x *LangSettings) GetConfigMapping() *ConfigMapping {
	if x != nil {
		return x.ConfigMapping
	}
	return nil
}

type ConfigMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeMapping      *TypeMapping `protobuf:"bytes,1,opt,name=typeMapping,proto3" json:"typeMapping,omitempty"`
	TypeDocMapping   *TypeMapping `protobuf:"bytes,2,opt,name=typeDocMapping,proto3" json:"typeDocMapping,omitempty"`
	ClassNameMapping string       `protobuf:"bytes,3,opt,name=classNameMapping,proto3" json:"classNameMapping,omitempty"`
	FileNameMapping  string       `protobuf:"bytes,4,opt,name=fileNameMapping,proto3" json:"fileNameMapping,omitempty"`
}

func (x *ConfigMapping) Reset() {
	*x = ConfigMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gendata_v1_gendata_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigMapping) ProtoMessage() {}

func (x *ConfigMapping) ProtoReflect() protoreflect.Message {
	mi := &file_gendata_v1_gendata_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigMapping.ProtoReflect.Descriptor instead.
func (*ConfigMapping) Descriptor() ([]byte, []int) {
	return file_gendata_v1_gendata_proto_rawDescGZIP(), []int{5}
}

func (x *ConfigMapping) GetTypeMapping() *TypeMapping {
	if x != nil {
		return x.TypeMapping
	}
	return nil
}

func (x *ConfigMapping) GetTypeDocMapping() *TypeMapping {
	if x != nil {
		return x.TypeDocMapping
	}
	return nil
}

func (x *ConfigMapping) GetClassNameMapping() string {
	if x != nil {
		return x.ClassNameMapping
	}
	return ""
}

func (x *ConfigMapping) GetFileNameMapping() string {
	if x != nil {
		return x.FileNameMapping
	}
	return ""
}

type TypeMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Array       string `protobuf:"bytes,1,opt,name=array,proto3" json:"array,omitempty"`
	ArrayBool   string `protobuf:"bytes,2,opt,name=arrayBool,proto3" json:"arrayBool,omitempty"`
	ArrayFloat  string `protobuf:"bytes,3,opt,name=arrayFloat,proto3" json:"arrayFloat,omitempty"`
	ArrayInt    string `protobuf:"bytes,4,opt,name=arrayInt,proto3" json:"arrayInt,omitempty"`
	ArrayObject string `protobuf:"bytes,5,opt,name=arrayObject,proto3" json:"arrayObject,omitempty"`
	ArrayString string `protobuf:"bytes,6,opt,name=arrayString,proto3" json:"arrayString,omitempty"`
	Bool        string `protobuf:"bytes,7,opt,name=bool,proto3" json:"bool,omitempty"`
	Float       string `protobuf:"bytes,8,opt,name=float,proto3" json:"float,omitempty"`
	Int         string `protobuf:"bytes,9,opt,name=int,proto3" json:"int,omitempty"`
	Null        string `protobuf:"bytes,10,opt,name=null,proto3" json:"null,omitempty"`
	Object      string `protobuf:"bytes,11,opt,name=object,proto3" json:"object,omitempty"`
	String_     string `protobuf:"bytes,12,opt,name=string,proto3" json:"string,omitempty"`
}

func (x *TypeMapping) Reset() {
	*x = TypeMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gendata_v1_gendata_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeMapping) ProtoMessage() {}

func (x *TypeMapping) ProtoReflect() protoreflect.Message {
	mi := &file_gendata_v1_gendata_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeMapping.ProtoReflect.Descriptor instead.
func (*TypeMapping) Descriptor() ([]byte, []int) {
	return file_gendata_v1_gendata_proto_rawDescGZIP(), []int{6}
}

func (x *TypeMapping) GetArray() string {
	if x != nil {
		return x.Array
	}
	return ""
}

func (x *TypeMapping) GetArrayBool() string {
	if x != nil {
		return x.ArrayBool
	}
	return ""
}

func (x *TypeMapping) GetArrayFloat() string {
	if x != nil {
		return x.ArrayFloat
	}
	return ""
}

func (x *TypeMapping) GetArrayInt() string {
	if x != nil {
		return x.ArrayInt
	}
	return ""
}

func (x *TypeMapping) GetArrayObject() string {
	if x != nil {
		return x.ArrayObject
	}
	return ""
}

func (x *TypeMapping) GetArrayString() string {
	if x != nil {
		return x.ArrayString
	}
	return ""
}

func (x *TypeMapping) GetBool() string {
	if x != nil {
		return x.Bool
	}
	return ""
}

func (x *TypeMapping) GetFloat() string {
	if x != nil {
		return x.Float
	}
	return ""
}

func (x *TypeMapping) GetInt() string {
	if x != nil {
		return x.Int
	}
	return ""
}

func (x *TypeMapping) GetNull() string {
	if x != nil {
		return x.Null
	}
	return ""
}

func (x *TypeMapping) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *TypeMapping) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

type PredefinedSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PredefinedSettingsRequest) Reset() {
	*x = PredefinedSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gendata_v1_gendata_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredefinedSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredefinedSettingsRequest) ProtoMessage() {}

func (x *PredefinedSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gendata_v1_gendata_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredefinedSettingsRequest.ProtoReflect.Descriptor instead.
func (*PredefinedSettingsRequest) Descriptor() ([]byte, []int) {
	return file_gendata_v1_gendata_proto_rawDescGZIP(), []int{7}
}

type PredefinedSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PredefinedSettingsResponse) Reset() {
	*x = PredefinedSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gendata_v1_gendata_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredefinedSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredefinedSettingsResponse) ProtoMessage() {}

func (x *PredefinedSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gendata_v1_gendata_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredefinedSettingsResponse.ProtoReflect.Descriptor instead.
func (*PredefinedSettingsResponse) Descriptor() ([]byte, []int) {
	return file_gendata_v1_gendata_proto_rawDescGZIP(), []int{8}
}

var File_gendata_v1_gendata_proto protoreflect.FileDescriptor

var file_gendata_v1_gendata_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x0d, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x44, 0x0a, 0x0c, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x60, 0x0a, 0x0a, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6d, 0x70, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x74, 0x6d, 0x70, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x9d, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x3c, 0x0a, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6c, 0x61, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x6f, 0x6f, 0x74, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x6f, 0x6f,
	0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73,
	0x75, 0x66, 0x66, 0x69, 0x78, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0xcd, 0x01, 0x0a, 0x0c, 0x4c, 0x61, 0x6e, 0x67,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x12, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42,
	0x79, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x22, 0xe1, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x39, 0x0a, 0x0b, 0x74, 0x79, 0x70,
	0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x12, 0x3f, 0x0a, 0x0e, 0x74, 0x79, 0x70, 0x65, 0x44, 0x6f, 0x63, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x0e, 0x74, 0x79, 0x70, 0x65, 0x44, 0x6f, 0x63, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61,
	0x6d, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x12, 0x28, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x22, 0xc1, 0x02, 0x0a, 0x0b,
	0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x72, 0x72, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x72, 0x72, 0x61,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x72, 0x61, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x72, 0x61, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x12,
	0x1e, 0x0a, 0x0a, 0x61, 0x72, 0x72, 0x61, 0x79, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x72, 0x72, 0x61, 0x79, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x72, 0x72, 0x61, 0x79, 0x49, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x72, 0x72, 0x61, 0x79, 0x49, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x61,
	0x72, 0x72, 0x61, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x72, 0x72, 0x61, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x61, 0x72, 0x72, 0x61, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x72, 0x72, 0x61, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62,
	0x6f, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x75, 0x6c, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x75, 0x6c, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22,
	0x1b, 0x0a, 0x19, 0x50, 0x72, 0x65, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1c, 0x0a, 0x1a,
	0x50, 0x72, 0x65, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xe3, 0x01, 0x0a, 0x0e, 0x47,
	0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a,
	0x03, 0x47, 0x65, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67,
	0x65, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x22, 0x07, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x84, 0x01, 0x0a, 0x12, 0x50, 0x72,
	0x65, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x25, 0x2e, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x65, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x65,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x2d, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x65, 0x6e, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x61, 0x3b, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_gendata_v1_gendata_proto_rawDescOnce sync.Once
	file_gendata_v1_gendata_proto_rawDescData = file_gendata_v1_gendata_proto_rawDesc
)

func file_gendata_v1_gendata_proto_rawDescGZIP() []byte {
	file_gendata_v1_gendata_proto_rawDescOnce.Do(func() {
		file_gendata_v1_gendata_proto_rawDescData = protoimpl.X.CompressGZIP(file_gendata_v1_gendata_proto_rawDescData)
	})
	return file_gendata_v1_gendata_proto_rawDescData
}

var file_gendata_v1_gendata_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_gendata_v1_gendata_proto_goTypes = []interface{}{
	(*GenResponse)(nil),                // 0: gendata.v1.GenResponse
	(*RenderedFile)(nil),               // 1: gendata.v1.RenderedFile
	(*GenRequest)(nil),                 // 2: gendata.v1.GenRequest
	(*Config)(nil),                     // 3: gendata.v1.Config
	(*LangSettings)(nil),               // 4: gendata.v1.LangSettings
	(*ConfigMapping)(nil),              // 5: gendata.v1.ConfigMapping
	(*TypeMapping)(nil),                // 6: gendata.v1.TypeMapping
	(*PredefinedSettingsRequest)(nil),  // 7: gendata.v1.PredefinedSettingsRequest
	(*PredefinedSettingsResponse)(nil), // 8: gendata.v1.PredefinedSettingsResponse
	(*duration.Duration)(nil),          // 9: google.protobuf.Duration
}
var file_gendata_v1_gendata_proto_depIdxs = []int32{
	1, // 0: gendata.v1.GenResponse.renderedFiles:type_name -> gendata.v1.RenderedFile
	9, // 1: gendata.v1.GenResponse.renderTime:type_name -> google.protobuf.Duration
	3, // 2: gendata.v1.GenRequest.config:type_name -> gendata.v1.Config
	4, // 3: gendata.v1.Config.langSettings:type_name -> gendata.v1.LangSettings
	5, // 4: gendata.v1.LangSettings.configMapping:type_name -> gendata.v1.ConfigMapping
	6, // 5: gendata.v1.ConfigMapping.typeMapping:type_name -> gendata.v1.TypeMapping
	6, // 6: gendata.v1.ConfigMapping.typeDocMapping:type_name -> gendata.v1.TypeMapping
	2, // 7: gendata.v1.GenDataService.Gen:input_type -> gendata.v1.GenRequest
	7, // 8: gendata.v1.GenDataService.PredefinedSettings:input_type -> gendata.v1.PredefinedSettingsRequest
	0, // 9: gendata.v1.GenDataService.Gen:output_type -> gendata.v1.GenResponse
	8, // 10: gendata.v1.GenDataService.PredefinedSettings:output_type -> gendata.v1.PredefinedSettingsResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_gendata_v1_gendata_proto_init() }
func file_gendata_v1_gendata_proto_init() {
	if File_gendata_v1_gendata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gendata_v1_gendata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenResponse); i {
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
		file_gendata_v1_gendata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenderedFile); i {
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
		file_gendata_v1_gendata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenRequest); i {
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
		file_gendata_v1_gendata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_gendata_v1_gendata_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LangSettings); i {
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
		file_gendata_v1_gendata_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigMapping); i {
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
		file_gendata_v1_gendata_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeMapping); i {
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
		file_gendata_v1_gendata_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredefinedSettingsRequest); i {
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
		file_gendata_v1_gendata_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredefinedSettingsResponse); i {
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
			RawDescriptor: file_gendata_v1_gendata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gendata_v1_gendata_proto_goTypes,
		DependencyIndexes: file_gendata_v1_gendata_proto_depIdxs,
		MessageInfos:      file_gendata_v1_gendata_proto_msgTypes,
	}.Build()
	File_gendata_v1_gendata_proto = out.File
	file_gendata_v1_gendata_proto_rawDesc = nil
	file_gendata_v1_gendata_proto_goTypes = nil
	file_gendata_v1_gendata_proto_depIdxs = nil
}
