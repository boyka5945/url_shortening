// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: proto/url_shortening.proto

package url

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

type CreateUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiDevKey      string `protobuf:"bytes,1,opt,name=api_dev_key,json=apiDevKey,proto3" json:"api_dev_key,omitempty"`
	OriginalUrl    string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
	CustomShortUrl string `protobuf:"bytes,3,opt,name=custom_short_url,json=customShortUrl,proto3" json:"custom_short_url,omitempty"`
	UserName       string `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	ExprireTime    uint32 `protobuf:"varint,5,opt,name=exprire_time,json=exprireTime,proto3" json:"exprire_time,omitempty"`
}

func (x *CreateUrlRequest) Reset() {
	*x = CreateUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_shortening_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUrlRequest) ProtoMessage() {}

func (x *CreateUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_shortening_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUrlRequest.ProtoReflect.Descriptor instead.
func (*CreateUrlRequest) Descriptor() ([]byte, []int) {
	return file_proto_url_shortening_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUrlRequest) GetApiDevKey() string {
	if x != nil {
		return x.ApiDevKey
	}
	return ""
}

func (x *CreateUrlRequest) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

func (x *CreateUrlRequest) GetCustomShortUrl() string {
	if x != nil {
		return x.CustomShortUrl
	}
	return ""
}

func (x *CreateUrlRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *CreateUrlRequest) GetExprireTime() uint32 {
	if x != nil {
		return x.ExprireTime
	}
	return 0
}

type CreateUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl     string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	ErrorCode    uint32 `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *CreateUrlResponse) Reset() {
	*x = CreateUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_shortening_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUrlResponse) ProtoMessage() {}

func (x *CreateUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_shortening_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUrlResponse.ProtoReflect.Descriptor instead.
func (*CreateUrlResponse) Descriptor() ([]byte, []int) {
	return file_proto_url_shortening_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUrlResponse) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *CreateUrlResponse) GetErrorCode() uint32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *CreateUrlResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type AccessUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiDevKey string `protobuf:"bytes,1,opt,name=api_dev_key,json=apiDevKey,proto3" json:"api_dev_key,omitempty"`
	ShortUrl  string `protobuf:"bytes,2,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *AccessUrlRequest) Reset() {
	*x = AccessUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_shortening_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessUrlRequest) ProtoMessage() {}

func (x *AccessUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_shortening_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessUrlRequest.ProtoReflect.Descriptor instead.
func (*AccessUrlRequest) Descriptor() ([]byte, []int) {
	return file_proto_url_shortening_proto_rawDescGZIP(), []int{2}
}

func (x *AccessUrlRequest) GetApiDevKey() string {
	if x != nil {
		return x.ApiDevKey
	}
	return ""
}

func (x *AccessUrlRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type AccessUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalUrl  string `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
	ErrorCode    uint32 `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *AccessUrlResponse) Reset() {
	*x = AccessUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_url_shortening_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessUrlResponse) ProtoMessage() {}

func (x *AccessUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_url_shortening_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessUrlResponse.ProtoReflect.Descriptor instead.
func (*AccessUrlResponse) Descriptor() ([]byte, []int) {
	return file_proto_url_shortening_proto_rawDescGZIP(), []int{3}
}

func (x *AccessUrlResponse) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

func (x *AccessUrlResponse) GetErrorCode() uint32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *AccessUrlResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_proto_url_shortening_proto protoreflect.FileDescriptor

var file_proto_url_shortening_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0xbf, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x64, 0x65,
	0x76, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x69,
	0x44, 0x65, 0x76, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x72, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x72, 0x69, 0x72, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x74, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4f, 0x0a, 0x10, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0b, 0x61, 0x70, 0x69, 0x5f, 0x64, 0x65, 0x76, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x69, 0x44, 0x65, 0x76, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x7a, 0x0a, 0x11, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55,
	0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x8e, 0x01, 0x0a, 0x14, 0x55, 0x72, 0x6c, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3a, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x15, 0x2e, 0x75,
	0x72, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x55, 0x72, 0x6c, 0x12, 0x15, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x75, 0x72, 0x6c, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x55, 0x72, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x3b, 0x75, 0x72, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_url_shortening_proto_rawDescOnce sync.Once
	file_proto_url_shortening_proto_rawDescData = file_proto_url_shortening_proto_rawDesc
)

func file_proto_url_shortening_proto_rawDescGZIP() []byte {
	file_proto_url_shortening_proto_rawDescOnce.Do(func() {
		file_proto_url_shortening_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_url_shortening_proto_rawDescData)
	})
	return file_proto_url_shortening_proto_rawDescData
}

var file_proto_url_shortening_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_url_shortening_proto_goTypes = []interface{}{
	(*CreateUrlRequest)(nil),  // 0: url.CreateUrlRequest
	(*CreateUrlResponse)(nil), // 1: url.CreateUrlResponse
	(*AccessUrlRequest)(nil),  // 2: url.AccessUrlRequest
	(*AccessUrlResponse)(nil), // 3: url.AccessUrlResponse
}
var file_proto_url_shortening_proto_depIdxs = []int32{
	0, // 0: url.UrlShorteningService.CreateUrl:input_type -> url.CreateUrlRequest
	2, // 1: url.UrlShorteningService.AccessUrl:input_type -> url.AccessUrlRequest
	1, // 2: url.UrlShorteningService.CreateUrl:output_type -> url.CreateUrlResponse
	3, // 3: url.UrlShorteningService.AccessUrl:output_type -> url.AccessUrlResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_url_shortening_proto_init() }
func file_proto_url_shortening_proto_init() {
	if File_proto_url_shortening_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_url_shortening_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUrlRequest); i {
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
		file_proto_url_shortening_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUrlResponse); i {
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
		file_proto_url_shortening_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessUrlRequest); i {
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
		file_proto_url_shortening_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessUrlResponse); i {
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
			RawDescriptor: file_proto_url_shortening_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_url_shortening_proto_goTypes,
		DependencyIndexes: file_proto_url_shortening_proto_depIdxs,
		MessageInfos:      file_proto_url_shortening_proto_msgTypes,
	}.Build()
	File_proto_url_shortening_proto = out.File
	file_proto_url_shortening_proto_rawDesc = nil
	file_proto_url_shortening_proto_goTypes = nil
	file_proto_url_shortening_proto_depIdxs = nil
}
