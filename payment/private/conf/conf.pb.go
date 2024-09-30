// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: payment/private/conf/conf.proto

package conf

import (
	_ "github.com/go-saas/kit/pkg/blob"
	conf "github.com/go-saas/kit/pkg/conf"
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

type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     *conf.Data      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Security *conf.Security  `protobuf:"bytes,3,opt,name=security,proto3" json:"security,omitempty"`
	Services *conf.Services  `protobuf:"bytes,4,opt,name=services,proto3" json:"services,omitempty"`
	Logging  *conf.Logging   `protobuf:"bytes,6,opt,name=logging,proto3" json:"logging,omitempty"`
	Tracing  *conf.Tracers   `protobuf:"bytes,7,opt,name=tracing,proto3" json:"tracing,omitempty"`
	App      *conf.AppConfig `protobuf:"bytes,8,opt,name=app,proto3" json:"app,omitempty"`
	Payment  *PaymentConf    `protobuf:"bytes,500,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_private_conf_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_payment_private_conf_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_payment_private_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetData() *conf.Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetSecurity() *conf.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *Bootstrap) GetServices() *conf.Services {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *Bootstrap) GetLogging() *conf.Logging {
	if x != nil {
		return x.Logging
	}
	return nil
}

func (x *Bootstrap) GetTracing() *conf.Tracers {
	if x != nil {
		return x.Tracing
	}
	return nil
}

func (x *Bootstrap) GetApp() *conf.AppConfig {
	if x != nil {
		return x.App
	}
	return nil
}

func (x *Bootstrap) GetPayment() *PaymentConf {
	if x != nil {
		return x.Payment
	}
	return nil
}

type PaymentConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Methods map[string]*PaymentMethod `protobuf:"bytes,1,rep,name=methods,proto3" json:"methods,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PaymentConf) Reset() {
	*x = PaymentConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_private_conf_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentConf) ProtoMessage() {}

func (x *PaymentConf) ProtoReflect() protoreflect.Message {
	mi := &file_payment_private_conf_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentConf.ProtoReflect.Descriptor instead.
func (*PaymentConf) Descriptor() ([]byte, []int) {
	return file_payment_private_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentConf) GetMethods() map[string]*PaymentMethod {
	if x != nil {
		return x.Methods
	}
	return nil
}

type PaymentMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stripe *StripePayment `protobuf:"bytes,1,opt,name=stripe,proto3" json:"stripe,omitempty"`
}

func (x *PaymentMethod) Reset() {
	*x = PaymentMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_private_conf_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentMethod) ProtoMessage() {}

func (x *PaymentMethod) ProtoReflect() protoreflect.Message {
	mi := &file_payment_private_conf_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentMethod.ProtoReflect.Descriptor instead.
func (*PaymentMethod) Descriptor() ([]byte, []int) {
	return file_payment_private_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentMethod) GetStripe() *StripePayment {
	if x != nil {
		return x.Stripe
	}
	return nil
}

type StripePayment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsTest     bool   `protobuf:"varint,1,opt,name=is_test,json=isTest,proto3" json:"is_test,omitempty"`
	PublishKey string `protobuf:"bytes,2,opt,name=publish_key,json=publishKey,proto3" json:"publish_key,omitempty"`
	PrivateKey string `protobuf:"bytes,3,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	WebhookKey string `protobuf:"bytes,4,opt,name=webhook_key,json=webhookKey,proto3" json:"webhook_key,omitempty"`
}

func (x *StripePayment) Reset() {
	*x = StripePayment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_private_conf_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StripePayment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StripePayment) ProtoMessage() {}

func (x *StripePayment) ProtoReflect() protoreflect.Message {
	mi := &file_payment_private_conf_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StripePayment.ProtoReflect.Descriptor instead.
func (*StripePayment) Descriptor() ([]byte, []int) {
	return file_payment_private_conf_conf_proto_rawDescGZIP(), []int{3}
}

func (x *StripePayment) GetIsTest() bool {
	if x != nil {
		return x.IsTest
	}
	return false
}

func (x *StripePayment) GetPublishKey() string {
	if x != nil {
		return x.PublishKey
	}
	return ""
}

func (x *StripePayment) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *StripePayment) GetWebhookKey() string {
	if x != nil {
		return x.WebhookKey
	}
	return ""
}

var File_payment_private_conf_conf_proto protoreflect.FileDescriptor

var file_payment_private_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x1a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x62,
	0x6c, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x02, 0x0a, 0x09, 0x42, 0x6f,
	0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x27, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x52,
	0x07, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63,
	0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x72, 0x73, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x12, 0x21, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x03, 0x61, 0x70, 0x70, 0x12, 0x3c, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0xf4, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0xb8, 0x01, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x12, 0x48, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x1a, 0x5f, 0x0a, 0x0c,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x39,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4c, 0x0a,
	0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x3b,
	0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x70, 0x65, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x0d,
	0x53, 0x74, 0x72, 0x69, 0x70, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x69, 0x73, 0x54, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x4b, 0x65, 0x79, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f,
	0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_private_conf_conf_proto_rawDescOnce sync.Once
	file_payment_private_conf_conf_proto_rawDescData = file_payment_private_conf_conf_proto_rawDesc
)

func file_payment_private_conf_conf_proto_rawDescGZIP() []byte {
	file_payment_private_conf_conf_proto_rawDescOnce.Do(func() {
		file_payment_private_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_private_conf_conf_proto_rawDescData)
	})
	return file_payment_private_conf_conf_proto_rawDescData
}

var file_payment_private_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_payment_private_conf_conf_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),      // 0: payment.private.conf.Bootstrap
	(*PaymentConf)(nil),    // 1: payment.private.conf.PaymentConf
	(*PaymentMethod)(nil),  // 2: payment.private.conf.PaymentMethod
	(*StripePayment)(nil),  // 3: payment.private.conf.StripePayment
	nil,                    // 4: payment.private.conf.PaymentConf.MethodsEntry
	(*conf.Data)(nil),      // 5: conf.Data
	(*conf.Security)(nil),  // 6: conf.Security
	(*conf.Services)(nil),  // 7: conf.Services
	(*conf.Logging)(nil),   // 8: conf.Logging
	(*conf.Tracers)(nil),   // 9: conf.Tracers
	(*conf.AppConfig)(nil), // 10: conf.AppConfig
}
var file_payment_private_conf_conf_proto_depIdxs = []int32{
	5,  // 0: payment.private.conf.Bootstrap.data:type_name -> conf.Data
	6,  // 1: payment.private.conf.Bootstrap.security:type_name -> conf.Security
	7,  // 2: payment.private.conf.Bootstrap.services:type_name -> conf.Services
	8,  // 3: payment.private.conf.Bootstrap.logging:type_name -> conf.Logging
	9,  // 4: payment.private.conf.Bootstrap.tracing:type_name -> conf.Tracers
	10, // 5: payment.private.conf.Bootstrap.app:type_name -> conf.AppConfig
	1,  // 6: payment.private.conf.Bootstrap.payment:type_name -> payment.private.conf.PaymentConf
	4,  // 7: payment.private.conf.PaymentConf.methods:type_name -> payment.private.conf.PaymentConf.MethodsEntry
	3,  // 8: payment.private.conf.PaymentMethod.stripe:type_name -> payment.private.conf.StripePayment
	2,  // 9: payment.private.conf.PaymentConf.MethodsEntry.value:type_name -> payment.private.conf.PaymentMethod
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_payment_private_conf_conf_proto_init() }
func file_payment_private_conf_conf_proto_init() {
	if File_payment_private_conf_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_private_conf_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
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
		file_payment_private_conf_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentConf); i {
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
		file_payment_private_conf_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentMethod); i {
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
		file_payment_private_conf_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StripePayment); i {
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
			RawDescriptor: file_payment_private_conf_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payment_private_conf_conf_proto_goTypes,
		DependencyIndexes: file_payment_private_conf_conf_proto_depIdxs,
		MessageInfos:      file_payment_private_conf_conf_proto_msgTypes,
	}.Build()
	File_payment_private_conf_conf_proto = out.File
	file_payment_private_conf_conf_proto_rawDesc = nil
	file_payment_private_conf_conf_proto_goTypes = nil
	file_payment_private_conf_conf_proto_depIdxs = nil
}
