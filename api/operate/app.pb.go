// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: operate/app.proto

package operate

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

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 内容ID
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 内容标题
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// 视频播放URL
	VideoUrl string `protobuf:"bytes,3,opt,name=video_url,json=videoUrl,proto3" json:"video_url,omitempty"`
	// 作者
	Author string `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	// 内容描述
	Description string `protobuf:"bytes,5,opt,name=Description,proto3" json:"Description,omitempty"`
	// 封面URL
	Thumbnail string `protobuf:"bytes,6,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	// 内容分类
	Category string `protobuf:"bytes,7,opt,name=category,proto3" json:"category,omitempty"`
	// 内容时长
	Duration int64 `protobuf:"varint,8,opt,name=duration,proto3" json:"duration,omitempty"`
	// 分辨率
	Resolution string `protobuf:"bytes,9,opt,name=resolution,proto3" json:"resolution,omitempty"`
	// 文件大小
	FileSize int64 `protobuf:"varint,10,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	// 文件格式
	Format string `protobuf:"bytes,11,opt,name=format,proto3" json:"format,omitempty"`
	// 视频质量 1-高清 2-标清
	Quality int32 `protobuf:"varint,12,opt,name=quality,proto3" json:"quality,omitempty"`
	// 审核状态
	ApprovalStatus int32 `protobuf:"varint,13,opt,name=approval_status,json=approvalStatus,proto3" json:"approval_status,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operate_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{0}
}

func (x *Content) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Content) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Content) GetVideoUrl() string {
	if x != nil {
		return x.VideoUrl
	}
	return ""
}

func (x *Content) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Content) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Content) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

func (x *Content) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Content) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Content) GetResolution() string {
	if x != nil {
		return x.Resolution
	}
	return ""
}

func (x *Content) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *Content) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *Content) GetQuality() int32 {
	if x != nil {
		return x.Quality
	}
	return 0
}

func (x *Content) GetApprovalStatus() int32 {
	if x != nil {
		return x.ApprovalStatus
	}
	return 0
}

type CreateContentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 内容
	Content *Content `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreateContentReq) Reset() {
	*x = CreateContentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operate_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContentReq) ProtoMessage() {}

func (x *CreateContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContentReq.ProtoReflect.Descriptor instead.
func (*CreateContentReq) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{1}
}

func (x *CreateContentReq) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}

type CreateContentRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateContentRsp) Reset() {
	*x = CreateContentRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operate_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContentRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContentRsp) ProtoMessage() {}

func (x *CreateContentRsp) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContentRsp.ProtoReflect.Descriptor instead.
func (*CreateContentRsp) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{2}
}

type UpdateContentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content *Content `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UpdateContentReq) Reset() {
	*x = UpdateContentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operate_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateContentReq) ProtoMessage() {}

func (x *UpdateContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateContentReq.ProtoReflect.Descriptor instead.
func (*UpdateContentReq) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateContentReq) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}

type UpdateContentRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateContentRsp) Reset() {
	*x = UpdateContentRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operate_app_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateContentRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateContentRsp) ProtoMessage() {}

func (x *UpdateContentRsp) ProtoReflect() protoreflect.Message {
	mi := &file_operate_app_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateContentRsp.ProtoReflect.Descriptor instead.
func (*UpdateContentRsp) Descriptor() ([]byte, []int) {
	return file_operate_app_proto_rawDescGZIP(), []int{4}
}

var File_operate_app_proto protoreflect.FileDescriptor

var file_operate_app_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x22, 0xf4, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x72, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x27,
	0x0a, 0x0f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x42, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2e, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x22,
	0x42, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x32, 0xa3, 0x01, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x12,
	0x4d, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x12, 0x4d,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x42, 0x32, 0x0a,
	0x0b, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x01, 0x5a, 0x21,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x3b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_operate_app_proto_rawDescOnce sync.Once
	file_operate_app_proto_rawDescData = file_operate_app_proto_rawDesc
)

func file_operate_app_proto_rawDescGZIP() []byte {
	file_operate_app_proto_rawDescOnce.Do(func() {
		file_operate_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_operate_app_proto_rawDescData)
	})
	return file_operate_app_proto_rawDescData
}

var file_operate_app_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_operate_app_proto_goTypes = []interface{}{
	(*Content)(nil),          // 0: api.operate.Content
	(*CreateContentReq)(nil), // 1: api.operate.CreateContentReq
	(*CreateContentRsp)(nil), // 2: api.operate.CreateContentRsp
	(*UpdateContentReq)(nil), // 3: api.operate.UpdateContentReq
	(*UpdateContentRsp)(nil), // 4: api.operate.UpdateContentRsp
}
var file_operate_app_proto_depIdxs = []int32{
	0, // 0: api.operate.CreateContentReq.content:type_name -> api.operate.Content
	0, // 1: api.operate.UpdateContentReq.content:type_name -> api.operate.Content
	1, // 2: api.operate.App.CreateContent:input_type -> api.operate.CreateContentReq
	3, // 3: api.operate.App.UpdateContent:input_type -> api.operate.UpdateContentReq
	2, // 4: api.operate.App.CreateContent:output_type -> api.operate.CreateContentRsp
	4, // 5: api.operate.App.UpdateContent:output_type -> api.operate.UpdateContentRsp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_operate_app_proto_init() }
func file_operate_app_proto_init() {
	if File_operate_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_operate_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
		file_operate_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContentReq); i {
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
		file_operate_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContentRsp); i {
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
		file_operate_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateContentReq); i {
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
		file_operate_app_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateContentRsp); i {
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
			RawDescriptor: file_operate_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_operate_app_proto_goTypes,
		DependencyIndexes: file_operate_app_proto_depIdxs,
		MessageInfos:      file_operate_app_proto_msgTypes,
	}.Build()
	File_operate_app_proto = out.File
	file_operate_app_proto_rawDesc = nil
	file_operate_app_proto_goTypes = nil
	file_operate_app_proto_depIdxs = nil
}
