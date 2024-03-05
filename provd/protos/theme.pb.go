// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: protos/theme.proto

package proto

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

type Theme int32

const (
	Theme_LIGHT Theme = 0
	Theme_DARK  Theme = 1
)

// Enum value maps for Theme.
var (
	Theme_name = map[int32]string{
		0: "LIGHT",
		1: "DARK",
	}
	Theme_value = map[string]int32{
		"LIGHT": 0,
		"DARK":  1,
	}
)

func (x Theme) Enum() *Theme {
	p := new(Theme)
	*p = x
	return p
}

func (x Theme) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Theme) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_theme_proto_enumTypes[0].Descriptor()
}

func (Theme) Type() protoreflect.EnumType {
	return &file_protos_theme_proto_enumTypes[0]
}

func (x Theme) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Theme.Descriptor instead.
func (Theme) EnumDescriptor() ([]byte, []int) {
	return file_protos_theme_proto_rawDescGZIP(), []int{0}
}

type AccentColor int32

const (
	AccentColor_ORANGE         AccentColor = 0
	AccentColor_BARK           AccentColor = 1
	AccentColor_SAGE           AccentColor = 2
	AccentColor_OLIVE          AccentColor = 3
	AccentColor_VIRIDIAN       AccentColor = 4
	AccentColor_PRUSSIAN_GREEN AccentColor = 5
	AccentColor_BLUE           AccentColor = 6
	AccentColor_PURPLE         AccentColor = 7
	AccentColor_MAGENTA        AccentColor = 8
	AccentColor_RED            AccentColor = 9
)

// Enum value maps for AccentColor.
var (
	AccentColor_name = map[int32]string{
		0: "ORANGE",
		1: "BARK",
		2: "SAGE",
		3: "OLIVE",
		4: "VIRIDIAN",
		5: "PRUSSIAN_GREEN",
		6: "BLUE",
		7: "PURPLE",
		8: "MAGENTA",
		9: "RED",
	}
	AccentColor_value = map[string]int32{
		"ORANGE":         0,
		"BARK":           1,
		"SAGE":           2,
		"OLIVE":          3,
		"VIRIDIAN":       4,
		"PRUSSIAN_GREEN": 5,
		"BLUE":           6,
		"PURPLE":         7,
		"MAGENTA":        8,
		"RED":            9,
	}
)

func (x AccentColor) Enum() *AccentColor {
	p := new(AccentColor)
	*p = x
	return p
}

func (x AccentColor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccentColor) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_theme_proto_enumTypes[1].Descriptor()
}

func (AccentColor) Type() protoreflect.EnumType {
	return &file_protos_theme_proto_enumTypes[1]
}

func (x AccentColor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccentColor.Descriptor instead.
func (AccentColor) EnumDescriptor() ([]byte, []int) {
	return file_protos_theme_proto_rawDescGZIP(), []int{1}
}

type SetThemeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Theme Theme `protobuf:"varint,1,opt,name=theme,proto3,enum=theme.Theme" json:"theme,omitempty"`
}

func (x *SetThemeRequest) Reset() {
	*x = SetThemeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_theme_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetThemeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetThemeRequest) ProtoMessage() {}

func (x *SetThemeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_theme_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetThemeRequest.ProtoReflect.Descriptor instead.
func (*SetThemeRequest) Descriptor() ([]byte, []int) {
	return file_protos_theme_proto_rawDescGZIP(), []int{0}
}

func (x *SetThemeRequest) GetTheme() Theme {
	if x != nil {
		return x.Theme
	}
	return Theme_LIGHT
}

type GetThemeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brightness Theme `protobuf:"varint,1,opt,name=brightness,proto3,enum=theme.Theme" json:"brightness,omitempty"`
}

func (x *GetThemeResponse) Reset() {
	*x = GetThemeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_theme_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetThemeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetThemeResponse) ProtoMessage() {}

func (x *GetThemeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_theme_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetThemeResponse.ProtoReflect.Descriptor instead.
func (*GetThemeResponse) Descriptor() ([]byte, []int) {
	return file_protos_theme_proto_rawDescGZIP(), []int{1}
}

func (x *GetThemeResponse) GetBrightness() Theme {
	if x != nil {
		return x.Brightness
	}
	return Theme_LIGHT
}

type SetAccentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccentColor AccentColor `protobuf:"varint,1,opt,name=accent_color,json=accentColor,proto3,enum=theme.AccentColor" json:"accent_color,omitempty"`
}

func (x *SetAccentRequest) Reset() {
	*x = SetAccentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_theme_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAccentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAccentRequest) ProtoMessage() {}

func (x *SetAccentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_theme_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAccentRequest.ProtoReflect.Descriptor instead.
func (*SetAccentRequest) Descriptor() ([]byte, []int) {
	return file_protos_theme_proto_rawDescGZIP(), []int{2}
}

func (x *SetAccentRequest) GetAccentColor() AccentColor {
	if x != nil {
		return x.AccentColor
	}
	return AccentColor_ORANGE
}

type GetAccentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccentColor AccentColor `protobuf:"varint,1,opt,name=accent_color,json=accentColor,proto3,enum=theme.AccentColor" json:"accent_color,omitempty"`
}

func (x *GetAccentResponse) Reset() {
	*x = GetAccentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_theme_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccentResponse) ProtoMessage() {}

func (x *GetAccentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_theme_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccentResponse.ProtoReflect.Descriptor instead.
func (*GetAccentResponse) Descriptor() ([]byte, []int) {
	return file_protos_theme_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccentResponse) GetAccentColor() AccentColor {
	if x != nil {
		return x.AccentColor
	}
	return AccentColor_ORANGE
}

var File_protos_theme_proto protoreflect.FileDescriptor

var file_protos_theme_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x54,
	0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x74,
	0x68, 0x65, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x74, 0x68, 0x65,
	0x6d, 0x65, 0x2e, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x22,
	0x40, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x0a, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x2e,
	0x54, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x0a, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x65, 0x73,
	0x73, 0x22, 0x49, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x6e, 0x74, 0x5f,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74, 0x68,
	0x65, 0x6d, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x4a, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x35, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x2e,
	0x41, 0x63, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x2a, 0x1c, 0x0a, 0x05, 0x54, 0x68, 0x65, 0x6d,
	0x65, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x44, 0x41, 0x52, 0x4b, 0x10, 0x01, 0x2a, 0x86, 0x01, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x52, 0x41, 0x4e, 0x47, 0x45,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x41, 0x52, 0x4b, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x53, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x4c, 0x49, 0x56, 0x45, 0x10,
	0x03, 0x12, 0x0c, 0x0a, 0x08, 0x56, 0x49, 0x52, 0x49, 0x44, 0x49, 0x41, 0x4e, 0x10, 0x04, 0x12,
	0x12, 0x0a, 0x0e, 0x50, 0x52, 0x55, 0x53, 0x53, 0x49, 0x41, 0x4e, 0x5f, 0x47, 0x52, 0x45, 0x45,
	0x4e, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4c, 0x55, 0x45, 0x10, 0x06, 0x12, 0x0a, 0x0a,
	0x06, 0x50, 0x55, 0x52, 0x50, 0x4c, 0x45, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x41, 0x47,
	0x45, 0x4e, 0x54, 0x41, 0x10, 0x08, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x45, 0x44, 0x10, 0x09, 0x32,
	0x84, 0x02, 0x0a, 0x0c, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3a, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x74,
	0x68, 0x65, 0x6d, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x17, 0x2e, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x68, 0x65, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x53, 0x65, 0x74,
	0x41, 0x63, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x53,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x74,
	0x68, 0x65, 0x6d, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x2f, 0x75,
	0x62, 0x75, 0x6e, 0x74, 0x75, 0x2d, 0x64, 0x65, 0x73, 0x6b, 0x74, 0x6f, 0x70, 0x2d, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x64, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_theme_proto_rawDescOnce sync.Once
	file_protos_theme_proto_rawDescData = file_protos_theme_proto_rawDesc
)

func file_protos_theme_proto_rawDescGZIP() []byte {
	file_protos_theme_proto_rawDescOnce.Do(func() {
		file_protos_theme_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_theme_proto_rawDescData)
	})
	return file_protos_theme_proto_rawDescData
}

var file_protos_theme_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_protos_theme_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_theme_proto_goTypes = []interface{}{
	(Theme)(0),                // 0: theme.Theme
	(AccentColor)(0),          // 1: theme.AccentColor
	(*SetThemeRequest)(nil),   // 2: theme.SetThemeRequest
	(*GetThemeResponse)(nil),  // 3: theme.GetThemeResponse
	(*SetAccentRequest)(nil),  // 4: theme.SetAccentRequest
	(*GetAccentResponse)(nil), // 5: theme.GetAccentResponse
	(*emptypb.Empty)(nil),     // 6: google.protobuf.Empty
}
var file_protos_theme_proto_depIdxs = []int32{
	0, // 0: theme.SetThemeRequest.theme:type_name -> theme.Theme
	0, // 1: theme.GetThemeResponse.brightness:type_name -> theme.Theme
	1, // 2: theme.SetAccentRequest.accent_color:type_name -> theme.AccentColor
	1, // 3: theme.GetAccentResponse.accent_color:type_name -> theme.AccentColor
	2, // 4: theme.ThemeService.SetTheme:input_type -> theme.SetThemeRequest
	6, // 5: theme.ThemeService.GetTheme:input_type -> google.protobuf.Empty
	4, // 6: theme.ThemeService.SetAccent:input_type -> theme.SetAccentRequest
	6, // 7: theme.ThemeService.GetAccent:input_type -> google.protobuf.Empty
	6, // 8: theme.ThemeService.SetTheme:output_type -> google.protobuf.Empty
	3, // 9: theme.ThemeService.GetTheme:output_type -> theme.GetThemeResponse
	6, // 10: theme.ThemeService.SetAccent:output_type -> google.protobuf.Empty
	5, // 11: theme.ThemeService.GetAccent:output_type -> theme.GetAccentResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protos_theme_proto_init() }
func file_protos_theme_proto_init() {
	if File_protos_theme_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_theme_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetThemeRequest); i {
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
		file_protos_theme_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetThemeResponse); i {
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
		file_protos_theme_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAccentRequest); i {
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
		file_protos_theme_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccentResponse); i {
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
			RawDescriptor: file_protos_theme_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_theme_proto_goTypes,
		DependencyIndexes: file_protos_theme_proto_depIdxs,
		EnumInfos:         file_protos_theme_proto_enumTypes,
		MessageInfos:      file_protos_theme_proto_msgTypes,
	}.Build()
	File_protos_theme_proto = out.File
	file_protos_theme_proto_rawDesc = nil
	file_protos_theme_proto_goTypes = nil
	file_protos_theme_proto_depIdxs = nil
}
