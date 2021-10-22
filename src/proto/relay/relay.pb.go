// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.1
// source: relay/relay.proto

package pbRelay

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

type MsgToUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendID         string `protobuf:"bytes,1,opt,name=SendID,proto3" json:"SendID,omitempty"`
	RecvID         string `protobuf:"bytes,2,opt,name=RecvID,proto3" json:"RecvID,omitempty"`
	Content        string `protobuf:"bytes,5,opt,name=Content,proto3" json:"Content,omitempty"`
	RecvSeq        int64  `protobuf:"varint,6,opt,name=RecvSeq,proto3" json:"RecvSeq,omitempty"`
	SendTime       int64  `protobuf:"varint,7,opt,name=SendTime,proto3" json:"SendTime,omitempty"`
	MsgFrom        int32  `protobuf:"varint,8,opt,name=MsgFrom,proto3" json:"MsgFrom,omitempty"`
	ContentType    int32  `protobuf:"varint,9,opt,name=ContentType,proto3" json:"ContentType,omitempty"`
	SessionType    int32  `protobuf:"varint,10,opt,name=SessionType,proto3" json:"SessionType,omitempty"`
	OperationID    string `protobuf:"bytes,11,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
	ServerMsgID    string `protobuf:"bytes,12,opt,name=ServerMsgID,proto3" json:"ServerMsgID,omitempty"`
	PlatformID     int32  `protobuf:"varint,13,opt,name=PlatformID,proto3" json:"PlatformID,omitempty"`
	SenderNickName string `protobuf:"bytes,14,opt,name=SenderNickName,proto3" json:"SenderNickName,omitempty"`
	SenderFaceURL  string `protobuf:"bytes,15,opt,name=SenderFaceURL,proto3" json:"SenderFaceURL,omitempty"`
	ClientMsgID    string `protobuf:"bytes,16,opt,name=ClientMsgID,proto3" json:"ClientMsgID,omitempty"`
}

func (x *MsgToUserReq) Reset() {
	*x = MsgToUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_relay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgToUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgToUserReq) ProtoMessage() {}

func (x *MsgToUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_relay_relay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgToUserReq.ProtoReflect.Descriptor instead.
func (*MsgToUserReq) Descriptor() ([]byte, []int) {
	return file_relay_relay_proto_rawDescGZIP(), []int{0}
}

func (x *MsgToUserReq) GetSendID() string {
	if x != nil {
		return x.SendID
	}
	return ""
}

func (x *MsgToUserReq) GetRecvID() string {
	if x != nil {
		return x.RecvID
	}
	return ""
}

func (x *MsgToUserReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *MsgToUserReq) GetRecvSeq() int64 {
	if x != nil {
		return x.RecvSeq
	}
	return 0
}

func (x *MsgToUserReq) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

func (x *MsgToUserReq) GetMsgFrom() int32 {
	if x != nil {
		return x.MsgFrom
	}
	return 0
}

func (x *MsgToUserReq) GetContentType() int32 {
	if x != nil {
		return x.ContentType
	}
	return 0
}

func (x *MsgToUserReq) GetSessionType() int32 {
	if x != nil {
		return x.SessionType
	}
	return 0
}

func (x *MsgToUserReq) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

func (x *MsgToUserReq) GetServerMsgID() string {
	if x != nil {
		return x.ServerMsgID
	}
	return ""
}

func (x *MsgToUserReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

func (x *MsgToUserReq) GetSenderNickName() string {
	if x != nil {
		return x.SenderNickName
	}
	return ""
}

func (x *MsgToUserReq) GetSenderFaceURL() string {
	if x != nil {
		return x.SenderFaceURL
	}
	return ""
}

func (x *MsgToUserReq) GetClientMsgID() string {
	if x != nil {
		return x.ClientMsgID
	}
	return ""
}

type MsgToUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resp []*SingleMsgToUser `protobuf:"bytes,1,rep,name=resp,proto3" json:"resp,omitempty"`
}

func (x *MsgToUserResp) Reset() {
	*x = MsgToUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_relay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgToUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgToUserResp) ProtoMessage() {}

func (x *MsgToUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_relay_relay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgToUserResp.ProtoReflect.Descriptor instead.
func (*MsgToUserResp) Descriptor() ([]byte, []int) {
	return file_relay_relay_proto_rawDescGZIP(), []int{1}
}

func (x *MsgToUserResp) GetResp() []*SingleMsgToUser {
	if x != nil {
		return x.Resp
	}
	return nil
}

//message SendMsgByWSReq{
//  string SendID = 1;
//  string RecvID = 2;
//  string Content = 3;
//  int64 SendTime = 4;
//  int64  MsgFrom = 5;
//  int64  ContentType = 6;
//  int64  SessionType = 7;
//  string OperationID = 8;
//  int64  PlatformID = 9;
//}
type SingleMsgToUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultCode     int64  `protobuf:"varint,1,opt,name=ResultCode,proto3" json:"ResultCode,omitempty"`
	RecvID         string `protobuf:"bytes,2,opt,name=RecvID,proto3" json:"RecvID,omitempty"`
	RecvPlatFormID int32  `protobuf:"varint,3,opt,name=RecvPlatFormID,proto3" json:"RecvPlatFormID,omitempty"`
}

func (x *SingleMsgToUser) Reset() {
	*x = SingleMsgToUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_relay_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleMsgToUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleMsgToUser) ProtoMessage() {}

func (x *SingleMsgToUser) ProtoReflect() protoreflect.Message {
	mi := &file_relay_relay_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleMsgToUser.ProtoReflect.Descriptor instead.
func (*SingleMsgToUser) Descriptor() ([]byte, []int) {
	return file_relay_relay_proto_rawDescGZIP(), []int{2}
}

func (x *SingleMsgToUser) GetResultCode() int64 {
	if x != nil {
		return x.ResultCode
	}
	return 0
}

func (x *SingleMsgToUser) GetRecvID() string {
	if x != nil {
		return x.RecvID
	}
	return ""
}

func (x *SingleMsgToUser) GetRecvPlatFormID() int32 {
	if x != nil {
		return x.RecvPlatFormID
	}
	return 0
}

var File_relay_relay_proto protoreflect.FileDescriptor

var file_relay_relay_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x22, 0xc0, 0x03, 0x0a, 0x0c, 0x4d,
	0x73, 0x67, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x65, 0x6e, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x65, 0x6e,
	0x64, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x76, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x63, 0x76, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x76, 0x53, 0x65, 0x71,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x52, 0x65, 0x63, 0x76, 0x53, 0x65, 0x71, 0x12,
	0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d,
	0x73, 0x67, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4d, 0x73,
	0x67, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x1e, 0x0a,
	0x0a, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x12, 0x26, 0x0a,
	0x0e, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x69, 0x63,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x46,
	0x61, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x53, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x46, 0x61, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x22, 0x3b, 0x0a,
	0x0d, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a,
	0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x54, 0x6f,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x72, 0x65, 0x73, 0x70, 0x22, 0x71, 0x0a, 0x0f, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a,
	0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x65, 0x63, 0x76, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52,
	0x65, 0x63, 0x76, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x76, 0x50, 0x6c, 0x61,
	0x74, 0x46, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x52,
	0x65, 0x63, 0x76, 0x50, 0x6c, 0x61, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x32, 0x53, 0x0a,
	0x19, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x6c, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x4d, 0x73,
	0x67, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e,
	0x4d, 0x73, 0x67, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x3b, 0x70, 0x62,
	0x52, 0x65, 0x6c, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relay_relay_proto_rawDescOnce sync.Once
	file_relay_relay_proto_rawDescData = file_relay_relay_proto_rawDesc
)

func file_relay_relay_proto_rawDescGZIP() []byte {
	file_relay_relay_proto_rawDescOnce.Do(func() {
		file_relay_relay_proto_rawDescData = protoimpl.X.CompressGZIP(file_relay_relay_proto_rawDescData)
	})
	return file_relay_relay_proto_rawDescData
}

var file_relay_relay_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_relay_relay_proto_goTypes = []interface{}{
	(*MsgToUserReq)(nil),    // 0: relay.MsgToUserReq
	(*MsgToUserResp)(nil),   // 1: relay.MsgToUserResp
	(*SingleMsgToUser)(nil), // 2: relay.SingleMsgToUser
}
var file_relay_relay_proto_depIdxs = []int32{
	2, // 0: relay.MsgToUserResp.resp:type_name -> relay.SingleMsgToUser
	0, // 1: relay.OnlineMessageRelayService.MsgToUser:input_type -> relay.MsgToUserReq
	1, // 2: relay.OnlineMessageRelayService.MsgToUser:output_type -> relay.MsgToUserResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_relay_relay_proto_init() }
func file_relay_relay_proto_init() {
	if File_relay_relay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relay_relay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgToUserReq); i {
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
		file_relay_relay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgToUserResp); i {
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
		file_relay_relay_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleMsgToUser); i {
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
			RawDescriptor: file_relay_relay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_relay_relay_proto_goTypes,
		DependencyIndexes: file_relay_relay_proto_depIdxs,
		MessageInfos:      file_relay_relay_proto_msgTypes,
	}.Build()
	File_relay_relay_proto = out.File
	file_relay_relay_proto_rawDesc = nil
	file_relay_relay_proto_goTypes = nil
	file_relay_relay_proto_depIdxs = nil
}
