// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: api/grpc/v1/push.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PushRequest_Priority int32

const (
	PushRequest_NORMAL PushRequest_Priority = 0
	PushRequest_HIGH   PushRequest_Priority = 1
)

// Enum value maps for PushRequest_Priority.
var (
	PushRequest_Priority_name = map[int32]string{
		0: "NORMAL",
		1: "HIGH",
	}
	PushRequest_Priority_value = map[string]int32{
		"NORMAL": 0,
		"HIGH":   1,
	}
)

func (x PushRequest_Priority) Enum() *PushRequest_Priority {
	p := new(PushRequest_Priority)
	*p = x
	return p
}

func (x PushRequest_Priority) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PushRequest_Priority) Descriptor() protoreflect.EnumDescriptor {
	return file_api_grpc_v1_push_proto_enumTypes[0].Descriptor()
}

func (PushRequest_Priority) Type() protoreflect.EnumType {
	return &file_api_grpc_v1_push_proto_enumTypes[0]
}

func (x PushRequest_Priority) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PushRequest_Priority.Descriptor instead.
func (PushRequest_Priority) EnumDescriptor() ([]byte, []int) {
	return file_api_grpc_v1_push_proto_rawDescGZIP(), []int{1, 0}
}

type Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body         string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Subtitle     string   `protobuf:"bytes,3,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	Action       string   `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	ActionLocKey string   `protobuf:"bytes,5,opt,name=actionLocKey,proto3" json:"actionLocKey,omitempty"`
	LaunchImage  string   `protobuf:"bytes,6,opt,name=launchImage,proto3" json:"launchImage,omitempty"`
	LocKey       string   `protobuf:"bytes,7,opt,name=locKey,proto3" json:"locKey,omitempty"`
	TitleLocKey  string   `protobuf:"bytes,8,opt,name=titleLocKey,proto3" json:"titleLocKey,omitempty"`
	LocArgs      []string `protobuf:"bytes,9,rep,name=locArgs,proto3" json:"locArgs,omitempty"`
	TitleLocArgs []string `protobuf:"bytes,10,rep,name=titleLocArgs,proto3" json:"titleLocArgs,omitempty"`
}

func (x *Alert) Reset() {
	*x = Alert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_push_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_push_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alert.ProtoReflect.Descriptor instead.
func (*Alert) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_push_proto_rawDescGZIP(), []int{0}
}

func (x *Alert) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Alert) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Alert) GetSubtitle() string {
	if x != nil {
		return x.Subtitle
	}
	return ""
}

func (x *Alert) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Alert) GetActionLocKey() string {
	if x != nil {
		return x.ActionLocKey
	}
	return ""
}

func (x *Alert) GetLaunchImage() string {
	if x != nil {
		return x.LaunchImage
	}
	return ""
}

func (x *Alert) GetLocKey() string {
	if x != nil {
		return x.LocKey
	}
	return ""
}

func (x *Alert) GetTitleLocKey() string {
	if x != nil {
		return x.TitleLocKey
	}
	return ""
}

func (x *Alert) GetLocArgs() []string {
	if x != nil {
		return x.LocArgs
	}
	return nil
}

func (x *Alert) GetTitleLocArgs() []string {
	if x != nil {
		return x.TitleLocArgs
	}
	return nil
}

type PushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform         string               `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty"`
	Tokens           []string             `protobuf:"bytes,2,rep,name=tokens,proto3" json:"tokens,omitempty"`
	Title            string               `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Message          string               `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Topic            string               `protobuf:"bytes,5,opt,name=topic,proto3" json:"topic,omitempty"`
	Key              string               `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	Category         string               `protobuf:"bytes,8,opt,name=category,proto3" json:"category,omitempty"`
	Sound            string               `protobuf:"bytes,10,opt,name=sound,proto3" json:"sound,omitempty"`
	Alert            *Alert               `protobuf:"bytes,9,opt,name=alert,proto3" json:"alert,omitempty"`
	Badge            int32                `protobuf:"varint,7,opt,name=badge,proto3" json:"badge,omitempty"`
	ThreadID         string               `protobuf:"bytes,12,opt,name=threadID,proto3" json:"threadID,omitempty"`
	Data             *structpb.Struct     `protobuf:"bytes,14,opt,name=data,proto3" json:"data,omitempty"`
	Image            string               `protobuf:"bytes,15,opt,name=image,proto3" json:"image,omitempty"`
	ID               string               `protobuf:"bytes,17,opt,name=ID,proto3" json:"ID,omitempty"`
	PushType         string               `protobuf:"bytes,18,opt,name=pushType,proto3" json:"pushType,omitempty"`
	AppID            string               `protobuf:"bytes,19,opt,name=appID,proto3" json:"appID,omitempty"`
	Priority         PushRequest_Priority `protobuf:"varint,16,opt,name=priority,proto3,enum=v1.PushRequest_Priority" json:"priority,omitempty"`
	ContentAvailable bool                 `protobuf:"varint,11,opt,name=contentAvailable,proto3" json:"contentAvailable,omitempty"`
	MutableContent   bool                 `protobuf:"varint,13,opt,name=mutableContent,proto3" json:"mutableContent,omitempty"`
	// default is production
	Development bool `protobuf:"varint,20,opt,name=development,proto3" json:"development,omitempty"`
}

func (x *PushRequest) Reset() {
	*x = PushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_push_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRequest) ProtoMessage() {}

func (x *PushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_push_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRequest.ProtoReflect.Descriptor instead.
func (*PushRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_push_proto_rawDescGZIP(), []int{1}
}

func (x *PushRequest) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *PushRequest) GetTokens() []string {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *PushRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PushRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PushRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *PushRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PushRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *PushRequest) GetSound() string {
	if x != nil {
		return x.Sound
	}
	return ""
}

func (x *PushRequest) GetAlert() *Alert {
	if x != nil {
		return x.Alert
	}
	return nil
}

func (x *PushRequest) GetBadge() int32 {
	if x != nil {
		return x.Badge
	}
	return 0
}

func (x *PushRequest) GetThreadID() string {
	if x != nil {
		return x.ThreadID
	}
	return ""
}

func (x *PushRequest) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PushRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *PushRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *PushRequest) GetPushType() string {
	if x != nil {
		return x.PushType
	}
	return ""
}

func (x *PushRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *PushRequest) GetPriority() PushRequest_Priority {
	if x != nil {
		return x.Priority
	}
	return PushRequest_NORMAL
}

func (x *PushRequest) GetContentAvailable() bool {
	if x != nil {
		return x.ContentAvailable
	}
	return false
}

func (x *PushRequest) GetMutableContent() bool {
	if x != nil {
		return x.MutableContent
	}
	return false
}

func (x *PushRequest) GetDevelopment() bool {
	if x != nil {
		return x.Development
	}
	return false
}

type PushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Counts  int32 `protobuf:"varint,2,opt,name=counts,proto3" json:"counts,omitempty"`
}

func (x *PushResponse) Reset() {
	*x = PushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_push_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushResponse) ProtoMessage() {}

func (x *PushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_push_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushResponse.ProtoReflect.Descriptor instead.
func (*PushResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_push_proto_rawDescGZIP(), []int{2}
}

func (x *PushResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PushResponse) GetCounts() int32 {
	if x != nil {
		return x.Counts
	}
	return 0
}

var File_api_grpc_v1_push_proto protoreflect.FileDescriptor

var file_api_grpc_v1_push_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75,
	0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x02, 0x0a, 0x05, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x4b,
	0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x6f, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x75,
	0x6e, 0x63, 0x68, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x4b,
	0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x4b, 0x65, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x4b, 0x65, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x4b,
	0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x63, 0x41, 0x72, 0x67, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x63, 0x41, 0x72, 0x67, 0x73, 0x12, 0x22, 0x0a, 0x0c,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x41, 0x72, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x41, 0x72, 0x67, 0x73,
	0x22, 0xf1, 0x04, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6f, 0x75, 0x6e,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1f,
	0x0a, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x61, 0x64, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x62, 0x61, 0x64, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49,
	0x44, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49,
	0x44, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x12, 0x34, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75,
	0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x10,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x75, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x20, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x0a,
	0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x49,
	0x47, 0x48, 0x10, 0x01, 0x22, 0x40, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x32, 0x3a, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x0f, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x73, 0x73, 0x69, 0x6d, 0x2f, 0x68, 0x69, 0x70, 0x75, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_grpc_v1_push_proto_rawDescOnce sync.Once
	file_api_grpc_v1_push_proto_rawDescData = file_api_grpc_v1_push_proto_rawDesc
)

func file_api_grpc_v1_push_proto_rawDescGZIP() []byte {
	file_api_grpc_v1_push_proto_rawDescOnce.Do(func() {
		file_api_grpc_v1_push_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_v1_push_proto_rawDescData)
	})
	return file_api_grpc_v1_push_proto_rawDescData
}

var file_api_grpc_v1_push_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_grpc_v1_push_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_grpc_v1_push_proto_goTypes = []interface{}{
	(PushRequest_Priority)(0), // 0: v1.PushRequest.Priority
	(*Alert)(nil),             // 1: v1.Alert
	(*PushRequest)(nil),       // 2: v1.PushRequest
	(*PushResponse)(nil),      // 3: v1.PushResponse
	(*structpb.Struct)(nil),   // 4: google.protobuf.Struct
}
var file_api_grpc_v1_push_proto_depIdxs = []int32{
	1, // 0: v1.PushRequest.alert:type_name -> v1.Alert
	4, // 1: v1.PushRequest.data:type_name -> google.protobuf.Struct
	0, // 2: v1.PushRequest.priority:type_name -> v1.PushRequest.Priority
	2, // 3: v1.PushService.Push:input_type -> v1.PushRequest
	3, // 4: v1.PushService.Push:output_type -> v1.PushResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_grpc_v1_push_proto_init() }
func file_api_grpc_v1_push_proto_init() {
	if File_api_grpc_v1_push_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_v1_push_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alert); i {
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
		file_api_grpc_v1_push_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRequest); i {
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
		file_api_grpc_v1_push_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushResponse); i {
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
			RawDescriptor: file_api_grpc_v1_push_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_v1_push_proto_goTypes,
		DependencyIndexes: file_api_grpc_v1_push_proto_depIdxs,
		EnumInfos:         file_api_grpc_v1_push_proto_enumTypes,
		MessageInfos:      file_api_grpc_v1_push_proto_msgTypes,
	}.Build()
	File_api_grpc_v1_push_proto = out.File
	file_api_grpc_v1_push_proto_rawDesc = nil
	file_api_grpc_v1_push_proto_goTypes = nil
	file_api_grpc_v1_push_proto_depIdxs = nil
}