// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: proto/order_management.proto

package proto

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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId         string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	UserId          string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Timestamp       string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ItemId          string `protobuf:"bytes,4,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	RiderId         string `protobuf:"bytes,5,opt,name=rider_id,json=riderId,proto3" json:"rider_id,omitempty"`
	RestaurantId    string `protobuf:"bytes,6,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	DeliveryAddress string `protobuf:"bytes,7,opt,name=delivery_address,json=deliveryAddress,proto3" json:"delivery_address,omitempty"`
	Assigned        bool   `protobuf:"varint,8,opt,name=assigned,proto3" json:"assigned,omitempty"`
	Delivered       bool   `protobuf:"varint,9,opt,name=delivered,proto3" json:"delivered,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_management_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_management_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_proto_order_management_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Order) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Order) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *Order) GetRiderId() string {
	if x != nil {
		return x.RiderId
	}
	return ""
}

func (x *Order) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *Order) GetDeliveryAddress() string {
	if x != nil {
		return x.DeliveryAddress
	}
	return ""
}

func (x *Order) GetAssigned() bool {
	if x != nil {
		return x.Assigned
	}
	return false
}

func (x *Order) GetDelivered() bool {
	if x != nil {
		return x.Delivered
	}
	return false
}

type GetAllOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllOrdersRequest) Reset() {
	*x = GetAllOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_management_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllOrdersRequest) ProtoMessage() {}

func (x *GetAllOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_management_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllOrdersRequest.ProtoReflect.Descriptor instead.
func (*GetAllOrdersRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_management_proto_rawDescGZIP(), []int{1}
}

type GetAllOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *GetAllOrdersResponse) Reset() {
	*x = GetAllOrdersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_management_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllOrdersResponse) ProtoMessage() {}

func (x *GetAllOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_management_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllOrdersResponse.ProtoReflect.Descriptor instead.
func (*GetAllOrdersResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_management_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllOrdersResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ItemId          string `protobuf:"bytes,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	RestaurantId    string `protobuf:"bytes,3,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	DeliveryAddress string `protobuf:"bytes,4,opt,name=delivery_address,json=deliveryAddress,proto3" json:"delivery_address,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_management_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_management_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_management_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOrderRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateOrderRequest) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *CreateOrderRequest) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *CreateOrderRequest) GetDeliveryAddress() string {
	if x != nil {
		return x.DeliveryAddress
	}
	return ""
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_management_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_management_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_management_proto_rawDescGZIP(), []int{4}
}

func (x *CreateOrderResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CreateOrderResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type AssignRiderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *AssignRiderRequest) Reset() {
	*x = AssignRiderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_management_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignRiderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignRiderRequest) ProtoMessage() {}

func (x *AssignRiderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_management_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignRiderRequest.ProtoReflect.Descriptor instead.
func (*AssignRiderRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_management_proto_rawDescGZIP(), []int{5}
}

func (x *AssignRiderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type AssignRiderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	RiderId string `protobuf:"bytes,2,opt,name=rider_id,json=riderId,proto3" json:"rider_id,omitempty"`
}

func (x *AssignRiderResponse) Reset() {
	*x = AssignRiderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_management_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignRiderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignRiderResponse) ProtoMessage() {}

func (x *AssignRiderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_management_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignRiderResponse.ProtoReflect.Descriptor instead.
func (*AssignRiderResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_management_proto_rawDescGZIP(), []int{6}
}

func (x *AssignRiderResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AssignRiderResponse) GetRiderId() string {
	if x != nil {
		return x.RiderId
	}
	return ""
}

type MarkOrderDeliveredRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *MarkOrderDeliveredRequest) Reset() {
	*x = MarkOrderDeliveredRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_management_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkOrderDeliveredRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkOrderDeliveredRequest) ProtoMessage() {}

func (x *MarkOrderDeliveredRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_management_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkOrderDeliveredRequest.ProtoReflect.Descriptor instead.
func (*MarkOrderDeliveredRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_management_proto_rawDescGZIP(), []int{7}
}

func (x *MarkOrderDeliveredRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type MarkOrderDeliveredResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *MarkOrderDeliveredResponse) Reset() {
	*x = MarkOrderDeliveredResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_management_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkOrderDeliveredResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkOrderDeliveredResponse) ProtoMessage() {}

func (x *MarkOrderDeliveredResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_management_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkOrderDeliveredResponse.ProtoReflect.Descriptor instead.
func (*MarkOrderDeliveredResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_management_proto_rawDescGZIP(), []int{8}
}

func (x *MarkOrderDeliveredResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_order_management_proto protoreflect.FileDescriptor

var file_proto_order_management_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x97, 0x02, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x46, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x96, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x4a, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2f, 0x0a,
	0x12, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4a,
	0x0a, 0x13, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x19, 0x0a, 0x08, 0x72, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x19, 0x4d, 0x61,
	0x72, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x36, 0x0a, 0x1a, 0x4d, 0x61, 0x72, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x91, 0x03, 0x0a, 0x0f, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x5b,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x24,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0b, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x52,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x52, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x6d, 0x0a, 0x12, 0x4d, 0x61, 0x72, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x65, 0x64, 0x12, 0x2a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08,
	0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_order_management_proto_rawDescOnce sync.Once
	file_proto_order_management_proto_rawDescData = file_proto_order_management_proto_rawDesc
)

func file_proto_order_management_proto_rawDescGZIP() []byte {
	file_proto_order_management_proto_rawDescOnce.Do(func() {
		file_proto_order_management_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_order_management_proto_rawDescData)
	})
	return file_proto_order_management_proto_rawDescData
}

var file_proto_order_management_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_order_management_proto_goTypes = []any{
	(*Order)(nil),                      // 0: ordermanagement.Order
	(*GetAllOrdersRequest)(nil),        // 1: ordermanagement.GetAllOrdersRequest
	(*GetAllOrdersResponse)(nil),       // 2: ordermanagement.GetAllOrdersResponse
	(*CreateOrderRequest)(nil),         // 3: ordermanagement.CreateOrderRequest
	(*CreateOrderResponse)(nil),        // 4: ordermanagement.CreateOrderResponse
	(*AssignRiderRequest)(nil),         // 5: ordermanagement.AssignRiderRequest
	(*AssignRiderResponse)(nil),        // 6: ordermanagement.AssignRiderResponse
	(*MarkOrderDeliveredRequest)(nil),  // 7: ordermanagement.MarkOrderDeliveredRequest
	(*MarkOrderDeliveredResponse)(nil), // 8: ordermanagement.MarkOrderDeliveredResponse
}
var file_proto_order_management_proto_depIdxs = []int32{
	0, // 0: ordermanagement.GetAllOrdersResponse.orders:type_name -> ordermanagement.Order
	1, // 1: ordermanagement.OrderManagement.GetAllOrders:input_type -> ordermanagement.GetAllOrdersRequest
	3, // 2: ordermanagement.OrderManagement.CreateOrder:input_type -> ordermanagement.CreateOrderRequest
	5, // 3: ordermanagement.OrderManagement.AssignRider:input_type -> ordermanagement.AssignRiderRequest
	7, // 4: ordermanagement.OrderManagement.MarkOrderDelivered:input_type -> ordermanagement.MarkOrderDeliveredRequest
	2, // 5: ordermanagement.OrderManagement.GetAllOrders:output_type -> ordermanagement.GetAllOrdersResponse
	4, // 6: ordermanagement.OrderManagement.CreateOrder:output_type -> ordermanagement.CreateOrderResponse
	6, // 7: ordermanagement.OrderManagement.AssignRider:output_type -> ordermanagement.AssignRiderResponse
	8, // 8: ordermanagement.OrderManagement.MarkOrderDelivered:output_type -> ordermanagement.MarkOrderDeliveredResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_order_management_proto_init() }
func file_proto_order_management_proto_init() {
	if File_proto_order_management_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_order_management_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Order); i {
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
		file_proto_order_management_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllOrdersRequest); i {
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
		file_proto_order_management_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllOrdersResponse); i {
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
		file_proto_order_management_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateOrderRequest); i {
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
		file_proto_order_management_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CreateOrderResponse); i {
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
		file_proto_order_management_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AssignRiderRequest); i {
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
		file_proto_order_management_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*AssignRiderResponse); i {
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
		file_proto_order_management_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*MarkOrderDeliveredRequest); i {
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
		file_proto_order_management_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*MarkOrderDeliveredResponse); i {
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
			RawDescriptor: file_proto_order_management_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_order_management_proto_goTypes,
		DependencyIndexes: file_proto_order_management_proto_depIdxs,
		MessageInfos:      file_proto_order_management_proto_msgTypes,
	}.Build()
	File_proto_order_management_proto = out.File
	file_proto_order_management_proto_rawDesc = nil
	file_proto_order_management_proto_goTypes = nil
	file_proto_order_management_proto_depIdxs = nil
}
