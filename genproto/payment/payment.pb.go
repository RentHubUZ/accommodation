// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.18.0
// source: protos/payment/payment.proto

package payment

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

type GetPaymentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // Payment ID
}

func (x *GetPaymentReq) Reset() {
	*x = GetPaymentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_payment_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentReq) ProtoMessage() {}

func (x *GetPaymentReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_payment_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentReq.ProtoReflect.Descriptor instead.
func (*GetPaymentReq) Descriptor() ([]byte, []int) {
	return file_protos_payment_payment_proto_rawDescGZIP(), []int{0}
}

func (x *GetPaymentReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPaymentRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payment *Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *GetPaymentRes) Reset() {
	*x = GetPaymentRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_payment_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentRes) ProtoMessage() {}

func (x *GetPaymentRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_payment_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentRes.ProtoReflect.Descriptor instead.
func (*GetPaymentRes) Descriptor() ([]byte, []int) {
	return file_protos_payment_payment_proto_rawDescGZIP(), []int{1}
}

func (x *GetPaymentRes) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

type GetAllPaymentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetAllPaymentReq) Reset() {
	*x = GetAllPaymentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_payment_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPaymentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPaymentReq) ProtoMessage() {}

func (x *GetAllPaymentReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_payment_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPaymentReq.ProtoReflect.Descriptor instead.
func (*GetAllPaymentReq) Descriptor() ([]byte, []int) {
	return file_protos_payment_payment_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllPaymentReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllPaymentReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAllPaymentRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payments []*Payment `protobuf:"bytes,1,rep,name=payments,proto3" json:"payments,omitempty"`
}

func (x *GetAllPaymentRes) Reset() {
	*x = GetAllPaymentRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_payment_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPaymentRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPaymentRes) ProtoMessage() {}

func (x *GetAllPaymentRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_payment_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPaymentRes.ProtoReflect.Descriptor instead.
func (*GetAllPaymentRes) Descriptor() ([]byte, []int) {
	return file_protos_payment_payment_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllPaymentRes) GetPayments() []*Payment {
	if x != nil {
		return x.Payments
	}
	return nil
}

type CreatePaymentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"` // Decimal as a string to avoid precision loss
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"` // 'pending', 'completed', etc.
}

func (x *CreatePaymentReq) Reset() {
	*x = CreatePaymentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_payment_payment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePaymentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentReq) ProtoMessage() {}

func (x *CreatePaymentReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_payment_payment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentReq.ProtoReflect.Descriptor instead.
func (*CreatePaymentReq) Descriptor() ([]byte, []int) {
	return file_protos_payment_payment_proto_rawDescGZIP(), []int{4}
}

func (x *CreatePaymentReq) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *CreatePaymentReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CreatePaymentRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // Payment ID
}

func (x *CreatePaymentRes) Reset() {
	*x = CreatePaymentRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_payment_payment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePaymentRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentRes) ProtoMessage() {}

func (x *CreatePaymentRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_payment_payment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentRes.ProtoReflect.Descriptor instead.
func (*CreatePaymentRes) Descriptor() ([]byte, []int) {
	return file_protos_payment_payment_proto_rawDescGZIP(), []int{5}
}

func (x *CreatePaymentRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeletePaymentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePaymentReq) Reset() {
	*x = DeletePaymentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_payment_payment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePaymentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePaymentReq) ProtoMessage() {}

func (x *DeletePaymentReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_payment_payment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePaymentReq.ProtoReflect.Descriptor instead.
func (*DeletePaymentReq) Descriptor() ([]byte, []int) {
	return file_protos_payment_payment_proto_rawDescGZIP(), []int{6}
}

func (x *DeletePaymentReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeletePaymentRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeletePaymentRes) Reset() {
	*x = DeletePaymentRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_payment_payment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePaymentRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePaymentRes) ProtoMessage() {}

func (x *DeletePaymentRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_payment_payment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePaymentRes.ProtoReflect.Descriptor instead.
func (*DeletePaymentRes) Descriptor() ([]byte, []int) {
	return file_protos_payment_payment_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePaymentRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount          string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"` // Decimal as a string
	Status          string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"` // 'pending', 'completed', etc.
	TransactionDate string `protobuf:"bytes,4,opt,name=transaction_date,json=transactionDate,proto3" json:"transaction_date,omitempty"`
}

func (x *Payment) Reset() {
	*x = Payment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_payment_payment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_protos_payment_payment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_protos_payment_payment_proto_rawDescGZIP(), []int{8}
}

func (x *Payment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Payment) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Payment) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Payment) GetTransactionDate() string {
	if x != nil {
		return x.TransactionDate
	}
	return ""
}

var File_protos_payment_payment_proto protoreflect.FileDescriptor

var file_protos_payment_payment_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x3c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x40, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x42, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x22, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x2c, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x74, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x32, 0x87, 0x02, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x16, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12,
	0x3e, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12,
	0x3e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12,
	0x3e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x42,
	0x12, 0x5a, 0x10, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_payment_payment_proto_rawDescOnce sync.Once
	file_protos_payment_payment_proto_rawDescData = file_protos_payment_payment_proto_rawDesc
)

func file_protos_payment_payment_proto_rawDescGZIP() []byte {
	file_protos_payment_payment_proto_rawDescOnce.Do(func() {
		file_protos_payment_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_payment_payment_proto_rawDescData)
	})
	return file_protos_payment_payment_proto_rawDescData
}

var file_protos_payment_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protos_payment_payment_proto_goTypes = []any{
	(*GetPaymentReq)(nil),    // 0: payment.GetPaymentReq
	(*GetPaymentRes)(nil),    // 1: payment.GetPaymentRes
	(*GetAllPaymentReq)(nil), // 2: payment.GetAllPaymentReq
	(*GetAllPaymentRes)(nil), // 3: payment.GetAllPaymentRes
	(*CreatePaymentReq)(nil), // 4: payment.CreatePaymentReq
	(*CreatePaymentRes)(nil), // 5: payment.CreatePaymentRes
	(*DeletePaymentReq)(nil), // 6: payment.DeletePaymentReq
	(*DeletePaymentRes)(nil), // 7: payment.DeletePaymentRes
	(*Payment)(nil),          // 8: payment.Payment
}
var file_protos_payment_payment_proto_depIdxs = []int32{
	8, // 0: payment.GetPaymentRes.payment:type_name -> payment.Payment
	8, // 1: payment.GetAllPaymentRes.payments:type_name -> payment.Payment
	0, // 2: payment.PaymentService.Get:input_type -> payment.GetPaymentReq
	2, // 3: payment.PaymentService.GetAll:input_type -> payment.GetAllPaymentReq
	4, // 4: payment.PaymentService.Create:input_type -> payment.CreatePaymentReq
	6, // 5: payment.PaymentService.Delete:input_type -> payment.DeletePaymentReq
	1, // 6: payment.PaymentService.Get:output_type -> payment.GetPaymentRes
	3, // 7: payment.PaymentService.GetAll:output_type -> payment.GetAllPaymentRes
	5, // 8: payment.PaymentService.Create:output_type -> payment.CreatePaymentRes
	7, // 9: payment.PaymentService.Delete:output_type -> payment.DeletePaymentRes
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_payment_payment_proto_init() }
func file_protos_payment_payment_proto_init() {
	if File_protos_payment_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_payment_payment_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetPaymentReq); i {
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
		file_protos_payment_payment_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetPaymentRes); i {
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
		file_protos_payment_payment_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllPaymentReq); i {
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
		file_protos_payment_payment_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllPaymentRes); i {
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
		file_protos_payment_payment_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePaymentReq); i {
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
		file_protos_payment_payment_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePaymentRes); i {
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
		file_protos_payment_payment_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeletePaymentReq); i {
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
		file_protos_payment_payment_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeletePaymentRes); i {
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
		file_protos_payment_payment_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Payment); i {
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
			RawDescriptor: file_protos_payment_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_payment_payment_proto_goTypes,
		DependencyIndexes: file_protos_payment_payment_proto_depIdxs,
		MessageInfos:      file_protos_payment_payment_proto_msgTypes,
	}.Build()
	File_protos_payment_payment_proto = out.File
	file_protos_payment_payment_proto_rawDesc = nil
	file_protos_payment_payment_proto_goTypes = nil
	file_protos_payment_payment_proto_depIdxs = nil
}
