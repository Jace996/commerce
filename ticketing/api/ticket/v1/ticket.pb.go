// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: ticketing/api/ticket/v1/ticket.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v11 "github.com/go-saas/commerce/ticketing/api/activity/v1"
	v1 "github.com/go-saas/commerce/ticketing/api/location/v1"
	v12 "github.com/go-saas/commerce/ticketing/api/show/v1"
	query "github.com/go-saas/kit/pkg/query"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

type GetTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTicketRequest) Reset() {
	*x = GetTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticketing_api_ticket_v1_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTicketRequest) ProtoMessage() {}

func (x *GetTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticketing_api_ticket_v1_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTicketRequest.ProtoReflect.Descriptor instead.
func (*GetTicketRequest) Descriptor() ([]byte, []int) {
	return file_ticketing_api_ticket_v1_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *GetTicketRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TicketFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *query.StringFilterOperation `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LocationId *query.StringFilterOperation `protobuf:"bytes,2,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	HallId     *query.StringFilterOperation `protobuf:"bytes,3,opt,name=hall_id,json=hallId,proto3" json:"hall_id,omitempty"`
	ActivityId *query.StringFilterOperation `protobuf:"bytes,4,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	UserId     *query.StringFilterOperation `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status     *query.StringFilterOperation `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *TicketFilter) Reset() {
	*x = TicketFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticketing_api_ticket_v1_ticket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketFilter) ProtoMessage() {}

func (x *TicketFilter) ProtoReflect() protoreflect.Message {
	mi := &file_ticketing_api_ticket_v1_ticket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketFilter.ProtoReflect.Descriptor instead.
func (*TicketFilter) Descriptor() ([]byte, []int) {
	return file_ticketing_api_ticket_v1_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *TicketFilter) GetId() *query.StringFilterOperation {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *TicketFilter) GetLocationId() *query.StringFilterOperation {
	if x != nil {
		return x.LocationId
	}
	return nil
}

func (x *TicketFilter) GetHallId() *query.StringFilterOperation {
	if x != nil {
		return x.HallId
	}
	return nil
}

func (x *TicketFilter) GetActivityId() *query.StringFilterOperation {
	if x != nil {
		return x.ActivityId
	}
	return nil
}

func (x *TicketFilter) GetUserId() *query.StringFilterOperation {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *TicketFilter) GetStatus() *query.StringFilterOperation {
	if x != nil {
		return x.Status
	}
	return nil
}

type ListTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageOffset      int32                  `protobuf:"varint,1,opt,name=page_offset,json=pageOffset,proto3" json:"page_offset,omitempty"`
	PageSize        int32                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Search          string                 `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	Sort            []string               `protobuf:"bytes,4,rep,name=sort,proto3" json:"sort,omitempty"`
	Fields          *fieldmaskpb.FieldMask `protobuf:"bytes,5,opt,name=fields,proto3" json:"fields,omitempty"`
	Filter          *TicketFilter          `protobuf:"bytes,6,opt,name=filter,proto3" json:"filter,omitempty"`
	AfterPageToken  string                 `protobuf:"bytes,10,opt,name=after_page_token,json=afterPageToken,proto3" json:"after_page_token,omitempty"`
	BeforePageToken string                 `protobuf:"bytes,11,opt,name=before_page_token,json=beforePageToken,proto3" json:"before_page_token,omitempty"`
}

func (x *ListTicketRequest) Reset() {
	*x = ListTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticketing_api_ticket_v1_ticket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTicketRequest) ProtoMessage() {}

func (x *ListTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticketing_api_ticket_v1_ticket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTicketRequest.ProtoReflect.Descriptor instead.
func (*ListTicketRequest) Descriptor() ([]byte, []int) {
	return file_ticketing_api_ticket_v1_ticket_proto_rawDescGZIP(), []int{2}
}

func (x *ListTicketRequest) GetPageOffset() int32 {
	if x != nil {
		return x.PageOffset
	}
	return 0
}

func (x *ListTicketRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListTicketRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *ListTicketRequest) GetSort() []string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *ListTicketRequest) GetFields() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *ListTicketRequest) GetFilter() *TicketFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ListTicketRequest) GetAfterPageToken() string {
	if x != nil {
		return x.AfterPageToken
	}
	return ""
}

func (x *ListTicketRequest) GetBeforePageToken() string {
	if x != nil {
		return x.BeforePageToken
	}
	return ""
}

type ListTicketReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalSize           int32     `protobuf:"varint,1,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	FilterSize          int32     `protobuf:"varint,2,opt,name=filter_size,json=filterSize,proto3" json:"filter_size,omitempty"`
	Items               []*Ticket `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	NextAfterPageToken  *string   `protobuf:"bytes,4,opt,name=next_after_page_token,json=nextAfterPageToken,proto3,oneof" json:"next_after_page_token,omitempty"`
	NextBeforePageToken *string   `protobuf:"bytes,5,opt,name=next_before_page_token,json=nextBeforePageToken,proto3,oneof" json:"next_before_page_token,omitempty"`
}

func (x *ListTicketReply) Reset() {
	*x = ListTicketReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticketing_api_ticket_v1_ticket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTicketReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTicketReply) ProtoMessage() {}

func (x *ListTicketReply) ProtoReflect() protoreflect.Message {
	mi := &file_ticketing_api_ticket_v1_ticket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTicketReply.ProtoReflect.Descriptor instead.
func (*ListTicketReply) Descriptor() ([]byte, []int) {
	return file_ticketing_api_ticket_v1_ticket_proto_rawDescGZIP(), []int{3}
}

func (x *ListTicketReply) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

func (x *ListTicketReply) GetFilterSize() int32 {
	if x != nil {
		return x.FilterSize
	}
	return 0
}

func (x *ListTicketReply) GetItems() []*Ticket {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListTicketReply) GetNextAfterPageToken() string {
	if x != nil && x.NextAfterPageToken != nil {
		return *x.NextAfterPageToken
	}
	return ""
}

func (x *ListTicketReply) GetNextBeforePageToken() string {
	if x != nil && x.NextBeforePageToken != nil {
		return *x.NextBeforePageToken
	}
	return ""
}

type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Location      *v1.Location           `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Hall          *v1.LocationHall       `protobuf:"bytes,3,opt,name=hall,proto3" json:"hall,omitempty"`
	Activity      *v11.Activity          `protobuf:"bytes,4,opt,name=activity,proto3" json:"activity,omitempty"`
	Show          *v12.Show              `protobuf:"bytes,5,opt,name=show,proto3" json:"show,omitempty"`
	ShowSalesType *v12.ShowSalesType     `protobuf:"bytes,6,opt,name=show_sales_type,json=showSalesType,proto3" json:"show_sales_type,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Status        string                 `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	OrderId       string                 `protobuf:"bytes,12,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	UserId        string                 `protobuf:"bytes,30,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Notice        string                 `protobuf:"bytes,40,opt,name=notice,proto3" json:"notice,omitempty"`
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ticketing_api_ticket_v1_ticket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_ticketing_api_ticket_v1_ticket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_ticketing_api_ticket_v1_ticket_proto_rawDescGZIP(), []int{4}
}

func (x *Ticket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ticket) GetLocation() *v1.Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Ticket) GetHall() *v1.LocationHall {
	if x != nil {
		return x.Hall
	}
	return nil
}

func (x *Ticket) GetActivity() *v11.Activity {
	if x != nil {
		return x.Activity
	}
	return nil
}

func (x *Ticket) GetShow() *v12.Show {
	if x != nil {
		return x.Show
	}
	return nil
}

func (x *Ticket) GetShowSalesType() *v12.ShowSalesType {
	if x != nil {
		return x.ShowSalesType
	}
	return nil
}

func (x *Ticket) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Ticket) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Ticket) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Ticket) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Ticket) GetNotice() string {
	if x != nil {
		return x.Notice
	}
	return ""
}

var File_ticketing_api_ticket_v1_ticket_proto protoreflect.FileDescriptor

var file_ticketing_api_ticket_v1_ticket_proto_rawDesc = []byte{
	0x0a, 0x24, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x28, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f,
	0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xe0, 0x41, 0x02, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9a, 0x03, 0x0a, 0x0c, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x47, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x07, 0x68, 0x61, 0x6c,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x68, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x0b, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0xc6, 0x02, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x6f, 0x72, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x3d, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x61, 0x66, 0x74, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x65,
	0x66, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xaf, 0x02,
	0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x35, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x36, 0x0a, 0x15, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x12, 0x6e, 0x65, 0x78, 0x74, 0x41,
	0x66, 0x74, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x38, 0x0a, 0x16, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x13, 0x6e, 0x65, 0x78, 0x74, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x62, 0x65,
	0x66, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0xf5, 0x03, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3f, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x04, 0x68,
	0x61, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61,
	0x6c, 0x6c, 0x52, 0x04, 0x68, 0x61, 0x6c, 0x6c, 0x12, 0x3f, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52,
	0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x2f, 0x0a, 0x04, 0x73, 0x68, 0x6f,
	0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x68, 0x6f, 0x77, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x77, 0x12, 0x4c, 0x0a, 0x0f, 0x73, 0x68,
	0x6f, 0x77, 0x5f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x6f, 0x77,
	0x53, 0x61, 0x6c, 0x65, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x73, 0x68, 0x6f, 0x77, 0x53,
	0x61, 0x6c, 0x65, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x32, 0xae, 0x02, 0x0a, 0x0d, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa0, 0x01, 0x0a, 0x0a, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x2a, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x3c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x5a, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x7a, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x29, 0x2e, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ticketing_api_ticket_v1_ticket_proto_rawDescOnce sync.Once
	file_ticketing_api_ticket_v1_ticket_proto_rawDescData = file_ticketing_api_ticket_v1_ticket_proto_rawDesc
)

func file_ticketing_api_ticket_v1_ticket_proto_rawDescGZIP() []byte {
	file_ticketing_api_ticket_v1_ticket_proto_rawDescOnce.Do(func() {
		file_ticketing_api_ticket_v1_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_ticketing_api_ticket_v1_ticket_proto_rawDescData)
	})
	return file_ticketing_api_ticket_v1_ticket_proto_rawDescData
}

var file_ticketing_api_ticket_v1_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ticketing_api_ticket_v1_ticket_proto_goTypes = []interface{}{
	(*GetTicketRequest)(nil),            // 0: ticketing.api.ticket.v1.GetTicketRequest
	(*TicketFilter)(nil),                // 1: ticketing.api.ticket.v1.TicketFilter
	(*ListTicketRequest)(nil),           // 2: ticketing.api.ticket.v1.ListTicketRequest
	(*ListTicketReply)(nil),             // 3: ticketing.api.ticket.v1.ListTicketReply
	(*Ticket)(nil),                      // 4: ticketing.api.ticket.v1.Ticket
	(*query.StringFilterOperation)(nil), // 5: query.operation.StringFilterOperation
	(*fieldmaskpb.FieldMask)(nil),       // 6: google.protobuf.FieldMask
	(*v1.Location)(nil),                 // 7: ticketing.api.location.v1.Location
	(*v1.LocationHall)(nil),             // 8: ticketing.api.location.v1.LocationHall
	(*v11.Activity)(nil),                // 9: ticketing.api.activity.v1.Activity
	(*v12.Show)(nil),                    // 10: ticketing.api.show.v1.Show
	(*v12.ShowSalesType)(nil),           // 11: ticketing.api.show.v1.ShowSalesType
	(*timestamppb.Timestamp)(nil),       // 12: google.protobuf.Timestamp
}
var file_ticketing_api_ticket_v1_ticket_proto_depIdxs = []int32{
	5,  // 0: ticketing.api.ticket.v1.TicketFilter.id:type_name -> query.operation.StringFilterOperation
	5,  // 1: ticketing.api.ticket.v1.TicketFilter.location_id:type_name -> query.operation.StringFilterOperation
	5,  // 2: ticketing.api.ticket.v1.TicketFilter.hall_id:type_name -> query.operation.StringFilterOperation
	5,  // 3: ticketing.api.ticket.v1.TicketFilter.activity_id:type_name -> query.operation.StringFilterOperation
	5,  // 4: ticketing.api.ticket.v1.TicketFilter.user_id:type_name -> query.operation.StringFilterOperation
	5,  // 5: ticketing.api.ticket.v1.TicketFilter.status:type_name -> query.operation.StringFilterOperation
	6,  // 6: ticketing.api.ticket.v1.ListTicketRequest.fields:type_name -> google.protobuf.FieldMask
	1,  // 7: ticketing.api.ticket.v1.ListTicketRequest.filter:type_name -> ticketing.api.ticket.v1.TicketFilter
	4,  // 8: ticketing.api.ticket.v1.ListTicketReply.items:type_name -> ticketing.api.ticket.v1.Ticket
	7,  // 9: ticketing.api.ticket.v1.Ticket.location:type_name -> ticketing.api.location.v1.Location
	8,  // 10: ticketing.api.ticket.v1.Ticket.hall:type_name -> ticketing.api.location.v1.LocationHall
	9,  // 11: ticketing.api.ticket.v1.Ticket.activity:type_name -> ticketing.api.activity.v1.Activity
	10, // 12: ticketing.api.ticket.v1.Ticket.show:type_name -> ticketing.api.show.v1.Show
	11, // 13: ticketing.api.ticket.v1.Ticket.show_sales_type:type_name -> ticketing.api.show.v1.ShowSalesType
	12, // 14: ticketing.api.ticket.v1.Ticket.created_at:type_name -> google.protobuf.Timestamp
	2,  // 15: ticketing.api.ticket.v1.TicketService.ListTicket:input_type -> ticketing.api.ticket.v1.ListTicketRequest
	0,  // 16: ticketing.api.ticket.v1.TicketService.GetTicket:input_type -> ticketing.api.ticket.v1.GetTicketRequest
	3,  // 17: ticketing.api.ticket.v1.TicketService.ListTicket:output_type -> ticketing.api.ticket.v1.ListTicketReply
	4,  // 18: ticketing.api.ticket.v1.TicketService.GetTicket:output_type -> ticketing.api.ticket.v1.Ticket
	17, // [17:19] is the sub-list for method output_type
	15, // [15:17] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_ticketing_api_ticket_v1_ticket_proto_init() }
func file_ticketing_api_ticket_v1_ticket_proto_init() {
	if File_ticketing_api_ticket_v1_ticket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ticketing_api_ticket_v1_ticket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTicketRequest); i {
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
		file_ticketing_api_ticket_v1_ticket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketFilter); i {
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
		file_ticketing_api_ticket_v1_ticket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTicketRequest); i {
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
		file_ticketing_api_ticket_v1_ticket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTicketReply); i {
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
		file_ticketing_api_ticket_v1_ticket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ticket); i {
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
	file_ticketing_api_ticket_v1_ticket_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ticketing_api_ticket_v1_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ticketing_api_ticket_v1_ticket_proto_goTypes,
		DependencyIndexes: file_ticketing_api_ticket_v1_ticket_proto_depIdxs,
		MessageInfos:      file_ticketing_api_ticket_v1_ticket_proto_msgTypes,
	}.Build()
	File_ticketing_api_ticket_v1_ticket_proto = out.File
	file_ticketing_api_ticket_v1_ticket_proto_rawDesc = nil
	file_ticketing_api_ticket_v1_ticket_proto_goTypes = nil
	file_ticketing_api_ticket_v1_ticket_proto_depIdxs = nil
}
