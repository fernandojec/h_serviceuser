// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: domain/notifications/proto/mesaging.proto

package notifications

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

type EmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From        string        `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To          []string      `protobuf:"bytes,2,rep,name=to,proto3" json:"to,omitempty"`
	Cc          []string      `protobuf:"bytes,3,rep,name=cc,proto3" json:"cc,omitempty"`
	Bcc         []string      `protobuf:"bytes,4,rep,name=bcc,proto3" json:"bcc,omitempty"`
	Subject     string        `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject,omitempty"`
	Body        string        `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	Attachments []*Attachment `protobuf:"bytes,7,rep,name=attachments,proto3" json:"attachments,omitempty"`
}

func (x *EmailRequest) Reset() {
	*x = EmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_notifications_proto_mesaging_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailRequest) ProtoMessage() {}

func (x *EmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_domain_notifications_proto_mesaging_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailRequest.ProtoReflect.Descriptor instead.
func (*EmailRequest) Descriptor() ([]byte, []int) {
	return file_domain_notifications_proto_mesaging_proto_rawDescGZIP(), []int{0}
}

func (x *EmailRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *EmailRequest) GetTo() []string {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *EmailRequest) GetCc() []string {
	if x != nil {
		return x.Cc
	}
	return nil
}

func (x *EmailRequest) GetBcc() []string {
	if x != nil {
		return x.Bcc
	}
	return nil
}

func (x *EmailRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *EmailRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *EmailRequest) GetAttachments() []*Attachment {
	if x != nil {
		return x.Attachments
	}
	return nil
}

type Attachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Attachment) Reset() {
	*x = Attachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_notifications_proto_mesaging_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attachment) ProtoMessage() {}

func (x *Attachment) ProtoReflect() protoreflect.Message {
	mi := &file_domain_notifications_proto_mesaging_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attachment.ProtoReflect.Descriptor instead.
func (*Attachment) Descriptor() ([]byte, []int) {
	return file_domain_notifications_proto_mesaging_proto_rawDescGZIP(), []int{1}
}

func (x *Attachment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Attachment) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type EmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EmailResponse) Reset() {
	*x = EmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_notifications_proto_mesaging_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailResponse) ProtoMessage() {}

func (x *EmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_domain_notifications_proto_mesaging_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailResponse.ProtoReflect.Descriptor instead.
func (*EmailResponse) Descriptor() ([]byte, []int) {
	return file_domain_notifications_proto_mesaging_proto_rawDescGZIP(), []int{2}
}

func (x *EmailResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *EmailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type WhatsAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To      string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *WhatsAppRequest) Reset() {
	*x = WhatsAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_notifications_proto_mesaging_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhatsAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhatsAppRequest) ProtoMessage() {}

func (x *WhatsAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_domain_notifications_proto_mesaging_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhatsAppRequest.ProtoReflect.Descriptor instead.
func (*WhatsAppRequest) Descriptor() ([]byte, []int) {
	return file_domain_notifications_proto_mesaging_proto_rawDescGZIP(), []int{3}
}

func (x *WhatsAppRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *WhatsAppRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type WhatsAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *WhatsAppResponse) Reset() {
	*x = WhatsAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_notifications_proto_mesaging_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhatsAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhatsAppResponse) ProtoMessage() {}

func (x *WhatsAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_domain_notifications_proto_mesaging_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhatsAppResponse.ProtoReflect.Descriptor instead.
func (*WhatsAppResponse) Descriptor() ([]byte, []int) {
	return file_domain_notifications_proto_mesaging_proto_rawDescGZIP(), []int{4}
}

func (x *WhatsAppResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *WhatsAppResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_domain_notifications_proto_mesaging_proto protoreflect.FileDescriptor

var file_domain_notifications_proto_mesaging_proto_rawDesc = []byte{
	0x0a, 0x29, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbf, 0x01, 0x0a, 0x0c, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x63, 0x63, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x63, 0x63, 0x12,
	0x10, 0x0a, 0x03, 0x62, 0x63, 0x63, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x62, 0x63,
	0x63, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12,
	0x3b, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x3a, 0x0a, 0x0a,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x43, 0x0a, 0x0d, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3b, 0x0a,
	0x0f, 0x57, 0x68, 0x61, 0x74, 0x73, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x46, 0x0a, 0x10, 0x57, 0x68,
	0x61, 0x74, 0x73, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0xab, 0x01, 0x0a, 0x10, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4f, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x57, 0x68, 0x61, 0x74, 0x73, 0x41, 0x70, 0x70, 0x12,
	0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x57, 0x68, 0x61, 0x74, 0x73, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x57, 0x68, 0x61, 0x74, 0x73, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x18, 0x5a, 0x16, 0x2e, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_domain_notifications_proto_mesaging_proto_rawDescOnce sync.Once
	file_domain_notifications_proto_mesaging_proto_rawDescData = file_domain_notifications_proto_mesaging_proto_rawDesc
)

func file_domain_notifications_proto_mesaging_proto_rawDescGZIP() []byte {
	file_domain_notifications_proto_mesaging_proto_rawDescOnce.Do(func() {
		file_domain_notifications_proto_mesaging_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_notifications_proto_mesaging_proto_rawDescData)
	})
	return file_domain_notifications_proto_mesaging_proto_rawDescData
}

var file_domain_notifications_proto_mesaging_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_domain_notifications_proto_mesaging_proto_goTypes = []interface{}{
	(*EmailRequest)(nil),     // 0: communication.EmailRequest
	(*Attachment)(nil),       // 1: communication.Attachment
	(*EmailResponse)(nil),    // 2: communication.EmailResponse
	(*WhatsAppRequest)(nil),  // 3: communication.WhatsAppRequest
	(*WhatsAppResponse)(nil), // 4: communication.WhatsAppResponse
}
var file_domain_notifications_proto_mesaging_proto_depIdxs = []int32{
	1, // 0: communication.EmailRequest.attachments:type_name -> communication.Attachment
	0, // 1: communication.MessagingService.SendEmail:input_type -> communication.EmailRequest
	3, // 2: communication.MessagingService.SendWhatsApp:input_type -> communication.WhatsAppRequest
	2, // 3: communication.MessagingService.SendEmail:output_type -> communication.EmailResponse
	4, // 4: communication.MessagingService.SendWhatsApp:output_type -> communication.WhatsAppResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_domain_notifications_proto_mesaging_proto_init() }
func file_domain_notifications_proto_mesaging_proto_init() {
	if File_domain_notifications_proto_mesaging_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_notifications_proto_mesaging_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailRequest); i {
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
		file_domain_notifications_proto_mesaging_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attachment); i {
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
		file_domain_notifications_proto_mesaging_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailResponse); i {
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
		file_domain_notifications_proto_mesaging_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhatsAppRequest); i {
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
		file_domain_notifications_proto_mesaging_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhatsAppResponse); i {
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
			RawDescriptor: file_domain_notifications_proto_mesaging_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_domain_notifications_proto_mesaging_proto_goTypes,
		DependencyIndexes: file_domain_notifications_proto_mesaging_proto_depIdxs,
		MessageInfos:      file_domain_notifications_proto_mesaging_proto_msgTypes,
	}.Build()
	File_domain_notifications_proto_mesaging_proto = out.File
	file_domain_notifications_proto_mesaging_proto_rawDesc = nil
	file_domain_notifications_proto_mesaging_proto_goTypes = nil
	file_domain_notifications_proto_mesaging_proto_depIdxs = nil
}
