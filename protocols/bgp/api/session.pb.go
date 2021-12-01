// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protocols/bgp/api/session.proto

package api

import (
	api "github.com/bio-routing/bio-rd/net/api"
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

type Session_State int32

const (
	Session_Disabled      Session_State = 0
	Session_Idle          Session_State = 1
	Session_Connect       Session_State = 2
	Session_Active        Session_State = 3
	Session_OpenSent      Session_State = 4
	Session_OpenConfirmed Session_State = 5
	Session_Established   Session_State = 6
)

// Enum value maps for Session_State.
var (
	Session_State_name = map[int32]string{
		0: "Disabled",
		1: "Idle",
		2: "Connect",
		3: "Active",
		4: "OpenSent",
		5: "OpenConfirmed",
		6: "Established",
	}
	Session_State_value = map[string]int32{
		"Disabled":      0,
		"Idle":          1,
		"Connect":       2,
		"Active":        3,
		"OpenSent":      4,
		"OpenConfirmed": 5,
		"Established":   6,
	}
)

func (x Session_State) Enum() *Session_State {
	p := new(Session_State)
	*p = x
	return p
}

func (x Session_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Session_State) Descriptor() protoreflect.EnumDescriptor {
	return file_protocols_bgp_api_session_proto_enumTypes[0].Descriptor()
}

func (Session_State) Type() protoreflect.EnumType {
	return &file_protocols_bgp_api_session_proto_enumTypes[0]
}

func (x Session_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Session_State.Descriptor instead.
func (Session_State) EnumDescriptor() ([]byte, []int) {
	return file_protocols_bgp_api_session_proto_rawDescGZIP(), []int{0, 0}
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalAddress     *api.IP       `protobuf:"bytes,1,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	NeighborAddress  *api.IP       `protobuf:"bytes,2,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	LocalAsn         uint32        `protobuf:"varint,3,opt,name=local_asn,json=localAsn,proto3" json:"local_asn,omitempty"`
	PeerAsn          uint32        `protobuf:"varint,4,opt,name=peer_asn,json=peerAsn,proto3" json:"peer_asn,omitempty"`
	Status           Session_State `protobuf:"varint,5,opt,name=status,proto3,enum=bio.bgp.Session_State" json:"status,omitempty"`
	Stats            *SessionStats `protobuf:"bytes,6,opt,name=stats,proto3" json:"stats,omitempty"`
	EstablishedSince uint64        `protobuf:"varint,7,opt,name=established_since,json=establishedSince,proto3" json:"established_since,omitempty"`
	Description      string        `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_bgp_api_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_bgp_api_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_protocols_bgp_api_session_proto_rawDescGZIP(), []int{0}
}

func (x *Session) GetLocalAddress() *api.IP {
	if x != nil {
		return x.LocalAddress
	}
	return nil
}

func (x *Session) GetNeighborAddress() *api.IP {
	if x != nil {
		return x.NeighborAddress
	}
	return nil
}

func (x *Session) GetLocalAsn() uint32 {
	if x != nil {
		return x.LocalAsn
	}
	return 0
}

func (x *Session) GetPeerAsn() uint32 {
	if x != nil {
		return x.PeerAsn
	}
	return 0
}

func (x *Session) GetStatus() Session_State {
	if x != nil {
		return x.Status
	}
	return Session_Disabled
}

func (x *Session) GetStats() *SessionStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *Session) GetEstablishedSince() uint64 {
	if x != nil {
		return x.EstablishedSince
	}
	return 0
}

func (x *Session) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type SessionStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessagesIn     uint64 `protobuf:"varint,1,opt,name=messages_in,json=messagesIn,proto3" json:"messages_in,omitempty"`
	MessagesOut    uint64 `protobuf:"varint,2,opt,name=messages_out,json=messagesOut,proto3" json:"messages_out,omitempty"`
	Flaps          uint64 `protobuf:"varint,3,opt,name=flaps,proto3" json:"flaps,omitempty"`
	RoutesReceived uint64 `protobuf:"varint,4,opt,name=routes_received,json=routesReceived,proto3" json:"routes_received,omitempty"`
	RoutesImported uint64 `protobuf:"varint,5,opt,name=routes_imported,json=routesImported,proto3" json:"routes_imported,omitempty"`
	RoutesExported uint64 `protobuf:"varint,6,opt,name=routes_exported,json=routesExported,proto3" json:"routes_exported,omitempty"`
}

func (x *SessionStats) Reset() {
	*x = SessionStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_bgp_api_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionStats) ProtoMessage() {}

func (x *SessionStats) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_bgp_api_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionStats.ProtoReflect.Descriptor instead.
func (*SessionStats) Descriptor() ([]byte, []int) {
	return file_protocols_bgp_api_session_proto_rawDescGZIP(), []int{1}
}

func (x *SessionStats) GetMessagesIn() uint64 {
	if x != nil {
		return x.MessagesIn
	}
	return 0
}

func (x *SessionStats) GetMessagesOut() uint64 {
	if x != nil {
		return x.MessagesOut
	}
	return 0
}

func (x *SessionStats) GetFlaps() uint64 {
	if x != nil {
		return x.Flaps
	}
	return 0
}

func (x *SessionStats) GetRoutesReceived() uint64 {
	if x != nil {
		return x.RoutesReceived
	}
	return 0
}

func (x *SessionStats) GetRoutesImported() uint64 {
	if x != nil {
		return x.RoutesImported
	}
	return 0
}

func (x *SessionStats) GetRoutesExported() uint64 {
	if x != nil {
		return x.RoutesExported
	}
	return 0
}

var File_protocols_bgp_api_session_proto protoreflect.FileDescriptor

var file_protocols_bgp_api_session_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x62, 0x67, 0x70, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x62, 0x69, 0x6f, 0x2e, 0x62, 0x67, 0x70, 0x1a, 0x11, 0x6e, 0x65, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x03,
	0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x0d, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x62, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x49, 0x50, 0x52, 0x0c, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x36, 0x0a, 0x10, 0x6e,
	0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x2e,
	0x49, 0x50, 0x52, 0x0f, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x73, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x73, 0x6e,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x61, 0x73, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x70, 0x65, 0x65, 0x72, 0x41, 0x73, 0x6e, 0x12, 0x2e, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x62, 0x69,
	0x6f, 0x2e, 0x62, 0x67, 0x70, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x69, 0x6f,
	0x2e, 0x62, 0x67, 0x70, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x73, 0x74, 0x61,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x10, 0x65, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x53, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6a, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x0c, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x49, 0x64, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10,
	0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x65, 0x6e, 0x74, 0x10, 0x04, 0x12,
	0x11, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64,
	0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x10, 0x06, 0x22, 0xe3, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x5f, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x49, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x70,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x70, 0x73, 0x12, 0x27,
	0x0a, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x5f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64,
	0x12, 0x27, 0x0a, 0x0f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x6f, 0x2d, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x2f, 0x62, 0x69, 0x6f, 0x2d, 0x72, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x62, 0x67, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocols_bgp_api_session_proto_rawDescOnce sync.Once
	file_protocols_bgp_api_session_proto_rawDescData = file_protocols_bgp_api_session_proto_rawDesc
)

func file_protocols_bgp_api_session_proto_rawDescGZIP() []byte {
	file_protocols_bgp_api_session_proto_rawDescOnce.Do(func() {
		file_protocols_bgp_api_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocols_bgp_api_session_proto_rawDescData)
	})
	return file_protocols_bgp_api_session_proto_rawDescData
}

var file_protocols_bgp_api_session_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protocols_bgp_api_session_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protocols_bgp_api_session_proto_goTypes = []interface{}{
	(Session_State)(0),   // 0: bio.bgp.Session.State
	(*Session)(nil),      // 1: bio.bgp.Session
	(*SessionStats)(nil), // 2: bio.bgp.SessionStats
	(*api.IP)(nil),       // 3: bio.net.IP
}
var file_protocols_bgp_api_session_proto_depIdxs = []int32{
	3, // 0: bio.bgp.Session.local_address:type_name -> bio.net.IP
	3, // 1: bio.bgp.Session.neighbor_address:type_name -> bio.net.IP
	0, // 2: bio.bgp.Session.status:type_name -> bio.bgp.Session.State
	2, // 3: bio.bgp.Session.stats:type_name -> bio.bgp.SessionStats
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protocols_bgp_api_session_proto_init() }
func file_protocols_bgp_api_session_proto_init() {
	if File_protocols_bgp_api_session_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocols_bgp_api_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_protocols_bgp_api_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionStats); i {
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
			RawDescriptor: file_protocols_bgp_api_session_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocols_bgp_api_session_proto_goTypes,
		DependencyIndexes: file_protocols_bgp_api_session_proto_depIdxs,
		EnumInfos:         file_protocols_bgp_api_session_proto_enumTypes,
		MessageInfos:      file_protocols_bgp_api_session_proto_msgTypes,
	}.Build()
	File_protocols_bgp_api_session_proto = out.File
	file_protocols_bgp_api_session_proto_rawDesc = nil
	file_protocols_bgp_api_session_proto_goTypes = nil
	file_protocols_bgp_api_session_proto_depIdxs = nil
}
