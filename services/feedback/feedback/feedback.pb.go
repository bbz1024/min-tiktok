// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.19.4
// source: feedback.proto

package feedback

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

type FeedbackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoIds []uint32 `protobuf:"varint,1,rep,packed,name=video_ids,json=videoIds,proto3" json:"video_ids,omitempty"`
	UserId   uint32   `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Type     string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"` // read comment favorite
}

func (x *FeedbackRequest) Reset() {
	*x = FeedbackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedbackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedbackRequest) ProtoMessage() {}

func (x *FeedbackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedbackRequest.ProtoReflect.Descriptor instead.
func (*FeedbackRequest) Descriptor() ([]byte, []int) {
	return file_feedback_proto_rawDescGZIP(), []int{0}
}

func (x *FeedbackRequest) GetVideoIds() []uint32 {
	if x != nil {
		return x.VideoIds
	}
	return nil
}

func (x *FeedbackRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FeedbackRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type FeedbackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode uint32 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
}

func (x *FeedbackResponse) Reset() {
	*x = FeedbackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedbackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedbackResponse) ProtoMessage() {}

func (x *FeedbackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedbackResponse.ProtoReflect.Descriptor instead.
func (*FeedbackResponse) Descriptor() ([]byte, []int) {
	return file_feedback_proto_rawDescGZIP(), []int{1}
}

func (x *FeedbackResponse) GetStatusCode() uint32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FeedbackResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type ListRecommendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorId uint32 `protobuf:"varint,1,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"`
	Count   uint32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListRecommendRequest) Reset() {
	*x = ListRecommendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRecommendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRecommendRequest) ProtoMessage() {}

func (x *ListRecommendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRecommendRequest.ProtoReflect.Descriptor instead.
func (*ListRecommendRequest) Descriptor() ([]byte, []int) {
	return file_feedback_proto_rawDescGZIP(), []int{2}
}

func (x *ListRecommendRequest) GetActorId() uint32 {
	if x != nil {
		return x.ActorId
	}
	return 0
}

func (x *ListRecommendRequest) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListRecommendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode uint32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string   `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	VideoIds   []string `protobuf:"bytes,3,rep,name=video_ids,json=videoIds,proto3" json:"video_ids,omitempty"`
}

func (x *ListRecommendResponse) Reset() {
	*x = ListRecommendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feedback_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRecommendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRecommendResponse) ProtoMessage() {}

func (x *ListRecommendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feedback_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRecommendResponse.ProtoReflect.Descriptor instead.
func (*ListRecommendResponse) Descriptor() ([]byte, []int) {
	return file_feedback_proto_rawDescGZIP(), []int{3}
}

func (x *ListRecommendResponse) GetStatusCode() uint32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ListRecommendResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *ListRecommendResponse) GetVideoIds() []string {
	if x != nil {
		return x.VideoIds
	}
	return nil
}

var File_feedback_proto protoreflect.FileDescriptor

var file_feedback_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x22, 0x5b, 0x0a, 0x0f, 0x46, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x52, 0x0a, 0x10, 0x46, 0x65, 0x65, 0x64, 0x62,
	0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x47, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x74, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x1b, 0x0a,
	0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x73, 0x32, 0xab, 0x01, 0x0a, 0x08, 0x46,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x41, 0x0a, 0x08, 0x46, 0x65, 0x65, 0x64, 0x62,
	0x61, 0x63, 0x6b, 0x12, 0x19, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x46,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x19, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x57, 0x69, 0x74, 0x68, 0x46,
	0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x1e, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61,
	0x63, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x66, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feedback_proto_rawDescOnce sync.Once
	file_feedback_proto_rawDescData = file_feedback_proto_rawDesc
)

func file_feedback_proto_rawDescGZIP() []byte {
	file_feedback_proto_rawDescOnce.Do(func() {
		file_feedback_proto_rawDescData = protoimpl.X.CompressGZIP(file_feedback_proto_rawDescData)
	})
	return file_feedback_proto_rawDescData
}

var file_feedback_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_feedback_proto_goTypes = []any{
	(*FeedbackRequest)(nil),       // 0: feedback.FeedbackRequest
	(*FeedbackResponse)(nil),      // 1: feedback.FeedbackResponse
	(*ListRecommendRequest)(nil),  // 2: feedback.ListRecommendRequest
	(*ListRecommendResponse)(nil), // 3: feedback.ListRecommendResponse
}
var file_feedback_proto_depIdxs = []int32{
	0, // 0: feedback.Feedback.Feedback:input_type -> feedback.FeedbackRequest
	2, // 1: feedback.Feedback.ListRecommendWithFeedback:input_type -> feedback.ListRecommendRequest
	1, // 2: feedback.Feedback.Feedback:output_type -> feedback.FeedbackResponse
	3, // 3: feedback.Feedback.ListRecommendWithFeedback:output_type -> feedback.ListRecommendResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_feedback_proto_init() }
func file_feedback_proto_init() {
	if File_feedback_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feedback_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FeedbackRequest); i {
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
		file_feedback_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FeedbackResponse); i {
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
		file_feedback_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListRecommendRequest); i {
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
		file_feedback_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListRecommendResponse); i {
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
			RawDescriptor: file_feedback_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feedback_proto_goTypes,
		DependencyIndexes: file_feedback_proto_depIdxs,
		MessageInfos:      file_feedback_proto_msgTypes,
	}.Build()
	File_feedback_proto = out.File
	file_feedback_proto_rawDesc = nil
	file_feedback_proto_goTypes = nil
	file_feedback_proto_depIdxs = nil
}