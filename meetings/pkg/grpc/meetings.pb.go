// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.1
// source: pkg/grpc/meetings.proto

package meetings

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Meeting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	EstateID     string                 `protobuf:"bytes,3,opt,name=estateID,proto3" json:"estateID,omitempty"`
	VisitorPhone string                 `protobuf:"bytes,4,opt,name=visitorPhone,proto3" json:"visitorPhone,omitempty"`
}

func (x *Meeting) Reset() {
	*x = Meeting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_meetings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meeting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meeting) ProtoMessage() {}

func (x *Meeting) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_meetings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meeting.ProtoReflect.Descriptor instead.
func (*Meeting) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_meetings_proto_rawDescGZIP(), []int{0}
}

func (x *Meeting) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Meeting) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Meeting) GetEstateID() string {
	if x != nil {
		return x.EstateID
	}
	return ""
}

func (x *Meeting) GetVisitorPhone() string {
	if x != nil {
		return x.VisitorPhone
	}
	return ""
}

type Meetings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meetings []*Meeting `protobuf:"bytes,1,rep,name=meetings,proto3" json:"meetings,omitempty"`
}

func (x *Meetings) Reset() {
	*x = Meetings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_meetings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meetings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meetings) ProtoMessage() {}

func (x *Meetings) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_meetings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meetings.ProtoReflect.Descriptor instead.
func (*Meetings) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_meetings_proto_rawDescGZIP(), []int{1}
}

func (x *Meetings) GetMeetings() []*Meeting {
	if x != nil {
		return x.Meetings
	}
	return nil
}

type GetMeetingsByEstateIDParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMeetingsByEstateIDParameter) Reset() {
	*x = GetMeetingsByEstateIDParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_meetings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeetingsByEstateIDParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeetingsByEstateIDParameter) ProtoMessage() {}

func (x *GetMeetingsByEstateIDParameter) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_meetings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeetingsByEstateIDParameter.ProtoReflect.Descriptor instead.
func (*GetMeetingsByEstateIDParameter) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_meetings_proto_rawDescGZIP(), []int{2}
}

func (x *GetMeetingsByEstateIDParameter) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetMeetingsByPhoneNumberParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string `protobuf:"bytes,1,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
}

func (x *GetMeetingsByPhoneNumberParameter) Reset() {
	*x = GetMeetingsByPhoneNumberParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_meetings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeetingsByPhoneNumberParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeetingsByPhoneNumberParameter) ProtoMessage() {}

func (x *GetMeetingsByPhoneNumberParameter) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_meetings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeetingsByPhoneNumberParameter.ProtoReflect.Descriptor instead.
func (*GetMeetingsByPhoneNumberParameter) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_meetings_proto_rawDescGZIP(), []int{3}
}

func (x *GetMeetingsByPhoneNumberParameter) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type CancelMeetingParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VisitorPhone string `protobuf:"bytes,2,opt,name=visitorPhone,proto3" json:"visitorPhone,omitempty"`
}

func (x *CancelMeetingParameter) Reset() {
	*x = CancelMeetingParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_meetings_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelMeetingParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelMeetingParameter) ProtoMessage() {}

func (x *CancelMeetingParameter) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_meetings_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelMeetingParameter.ProtoReflect.Descriptor instead.
func (*CancelMeetingParameter) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_meetings_proto_rawDescGZIP(), []int{4}
}

func (x *CancelMeetingParameter) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CancelMeetingParameter) GetVisitorPhone() string {
	if x != nil {
		return x.VisitorPhone
	}
	return ""
}

type GetAvailableTimeForMeetingParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EstateID string `protobuf:"bytes,1,opt,name=estateID,proto3" json:"estateID,omitempty"`
}

func (x *GetAvailableTimeForMeetingParameter) Reset() {
	*x = GetAvailableTimeForMeetingParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_meetings_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvailableTimeForMeetingParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailableTimeForMeetingParameter) ProtoMessage() {}

func (x *GetAvailableTimeForMeetingParameter) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_meetings_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailableTimeForMeetingParameter.ProtoReflect.Descriptor instead.
func (*GetAvailableTimeForMeetingParameter) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_meetings_proto_rawDescGZIP(), []int{5}
}

func (x *GetAvailableTimeForMeetingParameter) GetEstateID() string {
	if x != nil {
		return x.EstateID
	}
	return ""
}

type AvailableTimeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamps []*timestamppb.Timestamp `protobuf:"bytes,1,rep,name=timestamps,proto3" json:"timestamps,omitempty"`
}

func (x *AvailableTimeList) Reset() {
	*x = AvailableTimeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_meetings_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailableTimeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailableTimeList) ProtoMessage() {}

func (x *AvailableTimeList) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_meetings_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailableTimeList.ProtoReflect.Descriptor instead.
func (*AvailableTimeList) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_meetings_proto_rawDescGZIP(), []int{6}
}

func (x *AvailableTimeList) GetTimestamps() []*timestamppb.Timestamp {
	if x != nil {
		return x.Timestamps
	}
	return nil
}

var File_pkg_grpc_meetings_proto protoreflect.FileDescriptor

var file_pkg_grpc_meetings_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x93, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x76, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x76, 0x69, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x39, 0x0a, 0x08, 0x4d, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x22, 0x30, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x42, 0x79, 0x45, 0x73, 0x74, 0x61, 0x74, 0x65, 0x49, 0x44, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x42, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x4c, 0x0a, 0x16, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x76, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x76, 0x69, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x41, 0x0a, 0x23, 0x47, 0x65, 0x74,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x6f, 0x72,
	0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x49, 0x44, 0x22, 0x4f, 0x0a, 0x11,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x32, 0xbc, 0x03,
	0x0a, 0x0f, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x11, 0x2e, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x4d,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x49,
	0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x20, 0x2e, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x68, 0x0a, 0x1a, 0x47, 0x65, 0x74,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x6f, 0x72,
	0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2d, 0x2e, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x46, 0x6f, 0x72, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x1a, 0x1b, 0x2e, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x5a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x42, 0x79, 0x45, 0x73, 0x74, 0x61, 0x74, 0x65, 0x49, 0x44, 0x12, 0x2d, 0x2e, 0x6d,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x6f, 0x72, 0x4d, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x6d, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x5b, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x79,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2b, 0x2e, 0x6d, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x42, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x6d, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x3b, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x73, 0x65, 0x72,
	0x6f, 0x76, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67,
	0x3b, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_grpc_meetings_proto_rawDescOnce sync.Once
	file_pkg_grpc_meetings_proto_rawDescData = file_pkg_grpc_meetings_proto_rawDesc
)

func file_pkg_grpc_meetings_proto_rawDescGZIP() []byte {
	file_pkg_grpc_meetings_proto_rawDescOnce.Do(func() {
		file_pkg_grpc_meetings_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_grpc_meetings_proto_rawDescData)
	})
	return file_pkg_grpc_meetings_proto_rawDescData
}

var file_pkg_grpc_meetings_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_grpc_meetings_proto_goTypes = []interface{}{
	(*Meeting)(nil),                             // 0: meetings.Meeting
	(*Meetings)(nil),                            // 1: meetings.Meetings
	(*GetMeetingsByEstateIDParameter)(nil),      // 2: meetings.GetMeetingsByEstateIDParameter
	(*GetMeetingsByPhoneNumberParameter)(nil),   // 3: meetings.GetMeetingsByPhoneNumberParameter
	(*CancelMeetingParameter)(nil),              // 4: meetings.CancelMeetingParameter
	(*GetAvailableTimeForMeetingParameter)(nil), // 5: meetings.GetAvailableTimeForMeetingParameter
	(*AvailableTimeList)(nil),                   // 6: meetings.AvailableTimeList
	(*timestamppb.Timestamp)(nil),               // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                       // 8: google.protobuf.Empty
}
var file_pkg_grpc_meetings_proto_depIdxs = []int32{
	7, // 0: meetings.Meeting.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: meetings.Meetings.meetings:type_name -> meetings.Meeting
	7, // 2: meetings.AvailableTimeList.timestamps:type_name -> google.protobuf.Timestamp
	0, // 3: meetings.MeetingsService.ArrangeMeeting:input_type -> meetings.Meeting
	4, // 4: meetings.MeetingsService.CancelMeeting:input_type -> meetings.CancelMeetingParameter
	5, // 5: meetings.MeetingsService.GetAvailableTimeForMeeting:input_type -> meetings.GetAvailableTimeForMeetingParameter
	5, // 6: meetings.MeetingsService.GetMeetingsByEstateID:input_type -> meetings.GetAvailableTimeForMeetingParameter
	3, // 7: meetings.MeetingsService.GetMeetingsByPhoneNumber:input_type -> meetings.GetMeetingsByPhoneNumberParameter
	8, // 8: meetings.MeetingsService.ArrangeMeeting:output_type -> google.protobuf.Empty
	8, // 9: meetings.MeetingsService.CancelMeeting:output_type -> google.protobuf.Empty
	6, // 10: meetings.MeetingsService.GetAvailableTimeForMeeting:output_type -> meetings.AvailableTimeList
	1, // 11: meetings.MeetingsService.GetMeetingsByEstateID:output_type -> meetings.Meetings
	1, // 12: meetings.MeetingsService.GetMeetingsByPhoneNumber:output_type -> meetings.Meetings
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_grpc_meetings_proto_init() }
func file_pkg_grpc_meetings_proto_init() {
	if File_pkg_grpc_meetings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_grpc_meetings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meeting); i {
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
		file_pkg_grpc_meetings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meetings); i {
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
		file_pkg_grpc_meetings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMeetingsByEstateIDParameter); i {
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
		file_pkg_grpc_meetings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMeetingsByPhoneNumberParameter); i {
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
		file_pkg_grpc_meetings_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelMeetingParameter); i {
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
		file_pkg_grpc_meetings_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAvailableTimeForMeetingParameter); i {
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
		file_pkg_grpc_meetings_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvailableTimeList); i {
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
			RawDescriptor: file_pkg_grpc_meetings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_grpc_meetings_proto_goTypes,
		DependencyIndexes: file_pkg_grpc_meetings_proto_depIdxs,
		MessageInfos:      file_pkg_grpc_meetings_proto_msgTypes,
	}.Build()
	File_pkg_grpc_meetings_proto = out.File
	file_pkg_grpc_meetings_proto_rawDesc = nil
	file_pkg_grpc_meetings_proto_goTypes = nil
	file_pkg_grpc_meetings_proto_depIdxs = nil
}
