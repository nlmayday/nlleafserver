// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: msgproto/serviceagent.proto

package msgproto

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

type ServerRegister struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Addr          string                 `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"` // 自己所在的服务器IP:PROT
	Tag           string                 `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`   // 服务器标识
	Id            int32                  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`    // 唯一ID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServerRegister) Reset() {
	*x = ServerRegister{}
	mi := &file_msgproto_serviceagent_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServerRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerRegister) ProtoMessage() {}

func (x *ServerRegister) ProtoReflect() protoreflect.Message {
	mi := &file_msgproto_serviceagent_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerRegister.ProtoReflect.Descriptor instead.
func (*ServerRegister) Descriptor() ([]byte, []int) {
	return file_msgproto_serviceagent_proto_rawDescGZIP(), []int{0}
}

func (x *ServerRegister) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *ServerRegister) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *ServerRegister) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CommonMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Gid           uint32                 `protobuf:"varint,1,opt,name=gid,proto3" json:"gid,omitempty"`  // 消息 ID
	Sid           uint32                 `protobuf:"varint,2,opt,name=sid,proto3" json:"sid,omitempty"`  // 子消息
	Uid           uint64                 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`  // 用户 ID
	Body          string                 `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"` // jsoni string 游戏协议内容
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CommonMessage) Reset() {
	*x = CommonMessage{}
	mi := &file_msgproto_serviceagent_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommonMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonMessage) ProtoMessage() {}

func (x *CommonMessage) ProtoReflect() protoreflect.Message {
	mi := &file_msgproto_serviceagent_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonMessage.ProtoReflect.Descriptor instead.
func (*CommonMessage) Descriptor() ([]byte, []int) {
	return file_msgproto_serviceagent_proto_rawDescGZIP(), []int{1}
}

func (x *CommonMessage) GetGid() uint32 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *CommonMessage) GetSid() uint32 {
	if x != nil {
		return x.Sid
	}
	return 0
}

func (x *CommonMessage) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *CommonMessage) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

// 心跳
type HeartbeatMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Times         uint64                 `protobuf:"varint,1,opt,name=times,proto3" json:"times,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HeartbeatMessage) Reset() {
	*x = HeartbeatMessage{}
	mi := &file_msgproto_serviceagent_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartbeatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatMessage) ProtoMessage() {}

func (x *HeartbeatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_msgproto_serviceagent_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatMessage.ProtoReflect.Descriptor instead.
func (*HeartbeatMessage) Descriptor() ([]byte, []int) {
	return file_msgproto_serviceagent_proto_rawDescGZIP(), []int{2}
}

func (x *HeartbeatMessage) GetTimes() uint64 {
	if x != nil {
		return x.Times
	}
	return 0
}

type JoinMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uid           uint64                 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JoinMessage) Reset() {
	*x = JoinMessage{}
	mi := &file_msgproto_serviceagent_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinMessage) ProtoMessage() {}

func (x *JoinMessage) ProtoReflect() protoreflect.Message {
	mi := &file_msgproto_serviceagent_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinMessage.ProtoReflect.Descriptor instead.
func (*JoinMessage) Descriptor() ([]byte, []int) {
	return file_msgproto_serviceagent_proto_rawDescGZIP(), []int{3}
}

func (x *JoinMessage) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type LeaveMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uid           uint64                 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LeaveMessage) Reset() {
	*x = LeaveMessage{}
	mi := &file_msgproto_serviceagent_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LeaveMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveMessage) ProtoMessage() {}

func (x *LeaveMessage) ProtoReflect() protoreflect.Message {
	mi := &file_msgproto_serviceagent_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveMessage.ProtoReflect.Descriptor instead.
func (*LeaveMessage) Descriptor() ([]byte, []int) {
	return file_msgproto_serviceagent_proto_rawDescGZIP(), []int{4}
}

func (x *LeaveMessage) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

var File_msgproto_serviceagent_proto protoreflect.FileDescriptor

var file_msgproto_serviceagent_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x73, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d,
	0x73, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x59, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x67,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x73, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x28, 0x0a, 0x10, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x22, 0x1f, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x32, 0xfe, 0x02, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x6d, 0x73, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x18,
	0x2e, 0x6d, 0x73, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x12, 0x17, 0x2e, 0x6d, 0x73, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x17, 0x2e, 0x6d,
	0x73, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x34, 0x0a, 0x04, 0x4a, 0x6f, 0x69,
	0x6e, 0x12, 0x15, 0x2e, 0x6d, 0x73, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x69,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x15, 0x2e, 0x6d, 0x73, 0x67, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x37, 0x0a, 0x05, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x6d, 0x73, 0x67, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x16, 0x2e, 0x6d, 0x73, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x65, 0x61, 0x76,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x52, 0x65, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x15, 0x2e, 0x6d, 0x73, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x15, 0x2e, 0x6d,
	0x73, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74,
	0x12, 0x1a, 0x2e, 0x6d, 0x73, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1a, 0x2e, 0x6d,
	0x73, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x73, 0x72, 0x63, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6d, 0x73, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msgproto_serviceagent_proto_rawDescOnce sync.Once
	file_msgproto_serviceagent_proto_rawDescData = file_msgproto_serviceagent_proto_rawDesc
)

func file_msgproto_serviceagent_proto_rawDescGZIP() []byte {
	file_msgproto_serviceagent_proto_rawDescOnce.Do(func() {
		file_msgproto_serviceagent_proto_rawDescData = protoimpl.X.CompressGZIP(file_msgproto_serviceagent_proto_rawDescData)
	})
	return file_msgproto_serviceagent_proto_rawDescData
}

var file_msgproto_serviceagent_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_msgproto_serviceagent_proto_goTypes = []any{
	(*ServerRegister)(nil),   // 0: msgproto.ServerRegister
	(*CommonMessage)(nil),    // 1: msgproto.CommonMessage
	(*HeartbeatMessage)(nil), // 2: msgproto.HeartbeatMessage
	(*JoinMessage)(nil),      // 3: msgproto.JoinMessage
	(*LeaveMessage)(nil),     // 4: msgproto.LeaveMessage
}
var file_msgproto_serviceagent_proto_depIdxs = []int32{
	0, // 0: msgproto.Serviceagent.Register:input_type -> msgproto.ServerRegister
	1, // 1: msgproto.Serviceagent.Channel:input_type -> msgproto.CommonMessage
	3, // 2: msgproto.Serviceagent.Join:input_type -> msgproto.JoinMessage
	4, // 3: msgproto.Serviceagent.Leave:input_type -> msgproto.LeaveMessage
	3, // 4: msgproto.Serviceagent.ReConnect:input_type -> msgproto.JoinMessage
	2, // 5: msgproto.Serviceagent.Heartbeat:input_type -> msgproto.HeartbeatMessage
	0, // 6: msgproto.Serviceagent.Register:output_type -> msgproto.ServerRegister
	1, // 7: msgproto.Serviceagent.Channel:output_type -> msgproto.CommonMessage
	3, // 8: msgproto.Serviceagent.Join:output_type -> msgproto.JoinMessage
	4, // 9: msgproto.Serviceagent.Leave:output_type -> msgproto.LeaveMessage
	3, // 10: msgproto.Serviceagent.ReConnect:output_type -> msgproto.JoinMessage
	2, // 11: msgproto.Serviceagent.Heartbeat:output_type -> msgproto.HeartbeatMessage
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msgproto_serviceagent_proto_init() }
func file_msgproto_serviceagent_proto_init() {
	if File_msgproto_serviceagent_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msgproto_serviceagent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_msgproto_serviceagent_proto_goTypes,
		DependencyIndexes: file_msgproto_serviceagent_proto_depIdxs,
		MessageInfos:      file_msgproto_serviceagent_proto_msgTypes,
	}.Build()
	File_msgproto_serviceagent_proto = out.File
	file_msgproto_serviceagent_proto_rawDesc = nil
	file_msgproto_serviceagent_proto_goTypes = nil
	file_msgproto_serviceagent_proto_depIdxs = nil
}
