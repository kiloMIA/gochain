// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: smartContract.proto

package smartcontract

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type SmartContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hashvalue string                 `protobuf:"bytes,1,opt,name=Hashvalue,proto3" json:"Hashvalue,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Pubkey    string                 `protobuf:"bytes,3,opt,name=Pubkey,proto3" json:"Pubkey,omitempty"`
	Payload   *anypb.Any             `protobuf:"bytes,4,opt,name=Payload,proto3" json:"Payload,omitempty"`
}

func (x *SmartContract) Reset() {
	*x = SmartContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smartContract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmartContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartContract) ProtoMessage() {}

func (x *SmartContract) ProtoReflect() protoreflect.Message {
	mi := &file_smartContract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartContract.ProtoReflect.Descriptor instead.
func (*SmartContract) Descriptor() ([]byte, []int) {
	return file_smartContract_proto_rawDescGZIP(), []int{0}
}

func (x *SmartContract) GetHashvalue() string {
	if x != nil {
		return x.Hashvalue
	}
	return ""
}

func (x *SmartContract) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *SmartContract) GetPubkey() string {
	if x != nil {
		return x.Pubkey
	}
	return ""
}

func (x *SmartContract) GetPayload() *anypb.Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

type SC01 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temperature int32 `protobuf:"varint,1,opt,name=Temperature,proto3" json:"Temperature,omitempty"`
}

func (x *SC01) Reset() {
	*x = SC01{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smartContract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC01) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC01) ProtoMessage() {}

func (x *SC01) ProtoReflect() protoreflect.Message {
	mi := &file_smartContract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC01.ProtoReflect.Descriptor instead.
func (*SC01) Descriptor() ([]byte, []int) {
	return file_smartContract_proto_rawDescGZIP(), []int{1}
}

func (x *SC01) GetTemperature() int32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

type NodeIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeID string `protobuf:"bytes,1,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
}

func (x *NodeIdentifier) Reset() {
	*x = NodeIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smartContract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeIdentifier) ProtoMessage() {}

func (x *NodeIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_smartContract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeIdentifier.ProtoReflect.Descriptor instead.
func (*NodeIdentifier) Descriptor() ([]byte, []int) {
	return file_smartContract_proto_rawDescGZIP(), []int{2}
}

func (x *NodeIdentifier) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

type ProposeContractRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contract *SmartContract `protobuf:"bytes,1,opt,name=Contract,proto3" json:"Contract,omitempty"`
}

func (x *ProposeContractRequest) Reset() {
	*x = ProposeContractRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smartContract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposeContractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposeContractRequest) ProtoMessage() {}

func (x *ProposeContractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smartContract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposeContractRequest.ProtoReflect.Descriptor instead.
func (*ProposeContractRequest) Descriptor() ([]byte, []int) {
	return file_smartContract_proto_rawDescGZIP(), []int{3}
}

func (x *ProposeContractRequest) GetContract() *SmartContract {
	if x != nil {
		return x.Contract
	}
	return nil
}

type ProposeContractResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accepted bool `protobuf:"varint,1,opt,name=Accepted,proto3" json:"Accepted,omitempty"`
}

func (x *ProposeContractResponse) Reset() {
	*x = ProposeContractResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smartContract_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposeContractResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposeContractResponse) ProtoMessage() {}

func (x *ProposeContractResponse) ProtoReflect() protoreflect.Message {
	mi := &file_smartContract_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposeContractResponse.ProtoReflect.Descriptor instead.
func (*ProposeContractResponse) Descriptor() ([]byte, []int) {
	return file_smartContract_proto_rawDescGZIP(), []int{4}
}

func (x *ProposeContractResponse) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

type ValidateContractRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contract *SmartContract `protobuf:"bytes,1,opt,name=Contract,proto3" json:"Contract,omitempty"`
}

func (x *ValidateContractRequest) Reset() {
	*x = ValidateContractRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smartContract_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateContractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateContractRequest) ProtoMessage() {}

func (x *ValidateContractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smartContract_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateContractRequest.ProtoReflect.Descriptor instead.
func (*ValidateContractRequest) Descriptor() ([]byte, []int) {
	return file_smartContract_proto_rawDescGZIP(), []int{5}
}

func (x *ValidateContractRequest) GetContract() *SmartContract {
	if x != nil {
		return x.Contract
	}
	return nil
}

type ValidateContractResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=Valid,proto3" json:"Valid,omitempty"`
}

func (x *ValidateContractResponse) Reset() {
	*x = ValidateContractResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smartContract_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateContractResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateContractResponse) ProtoMessage() {}

func (x *ValidateContractResponse) ProtoReflect() protoreflect.Message {
	mi := &file_smartContract_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateContractResponse.ProtoReflect.Descriptor instead.
func (*ValidateContractResponse) Descriptor() ([]byte, []int) {
	return file_smartContract_proto_rawDescGZIP(), []int{6}
}

func (x *ValidateContractResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

type ConsensusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeID   *NodeIdentifier `protobuf:"bytes,1,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
	Contract *SmartContract  `protobuf:"bytes,2,opt,name=Contract,proto3" json:"Contract,omitempty"`
}

func (x *ConsensusRequest) Reset() {
	*x = ConsensusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smartContract_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusRequest) ProtoMessage() {}

func (x *ConsensusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smartContract_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusRequest.ProtoReflect.Descriptor instead.
func (*ConsensusRequest) Descriptor() ([]byte, []int) {
	return file_smartContract_proto_rawDescGZIP(), []int{7}
}

func (x *ConsensusRequest) GetNodeID() *NodeIdentifier {
	if x != nil {
		return x.NodeID
	}
	return nil
}

func (x *ConsensusRequest) GetContract() *SmartContract {
	if x != nil {
		return x.Contract
	}
	return nil
}

type ConsensusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agreed bool `protobuf:"varint,1,opt,name=Agreed,proto3" json:"Agreed,omitempty"`
}

func (x *ConsensusResponse) Reset() {
	*x = ConsensusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smartContract_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusResponse) ProtoMessage() {}

func (x *ConsensusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_smartContract_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusResponse.ProtoReflect.Descriptor instead.
func (*ConsensusResponse) Descriptor() ([]byte, []int) {
	return file_smartContract_proto_rawDescGZIP(), []int{8}
}

func (x *ConsensusResponse) GetAgreed() bool {
	if x != nil {
		return x.Agreed
	}
	return false
}

var File_smartContract_proto protoreflect.FileDescriptor

var file_smartContract_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x0d, 0x53, 0x6d, 0x61,
	0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x48, 0x61,
	0x73, 0x68, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x48,
	0x61, 0x73, 0x68, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x07, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x28, 0x0a, 0x04, 0x53, 0x43,
	0x30, 0x31, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x22, 0x4a,
	0x0a, 0x16, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x52, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x22, 0x35, 0x0a, 0x17, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65,
	0x64, 0x22, 0x4b, 0x0a, 0x17, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x08,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x22, 0x30,
	0x0a, 0x18, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x22, 0x73, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x06, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6d,
	0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x08, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x22, 0x2b, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x67,
	0x72, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x41, 0x67, 0x72, 0x65,
	0x65, 0x64, 0x32, 0xfb, 0x01, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72,
	0x6f, 0x70, 0x6f, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x10, 0x41, 0x63, 0x68,
	0x69, 0x65, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x12, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x15, 0x5a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_smartContract_proto_rawDescOnce sync.Once
	file_smartContract_proto_rawDescData = file_smartContract_proto_rawDesc
)

func file_smartContract_proto_rawDescGZIP() []byte {
	file_smartContract_proto_rawDescOnce.Do(func() {
		file_smartContract_proto_rawDescData = protoimpl.X.CompressGZIP(file_smartContract_proto_rawDescData)
	})
	return file_smartContract_proto_rawDescData
}

var file_smartContract_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_smartContract_proto_goTypes = []interface{}{
	(*SmartContract)(nil),            // 0: proto.SmartContract
	(*SC01)(nil),                     // 1: proto.SC01
	(*NodeIdentifier)(nil),           // 2: proto.NodeIdentifier
	(*ProposeContractRequest)(nil),   // 3: proto.ProposeContractRequest
	(*ProposeContractResponse)(nil),  // 4: proto.ProposeContractResponse
	(*ValidateContractRequest)(nil),  // 5: proto.ValidateContractRequest
	(*ValidateContractResponse)(nil), // 6: proto.ValidateContractResponse
	(*ConsensusRequest)(nil),         // 7: proto.ConsensusRequest
	(*ConsensusResponse)(nil),        // 8: proto.ConsensusResponse
	(*timestamppb.Timestamp)(nil),    // 9: google.protobuf.Timestamp
	(*anypb.Any)(nil),                // 10: google.protobuf.Any
}
var file_smartContract_proto_depIdxs = []int32{
	9,  // 0: proto.SmartContract.Timestamp:type_name -> google.protobuf.Timestamp
	10, // 1: proto.SmartContract.Payload:type_name -> google.protobuf.Any
	0,  // 2: proto.ProposeContractRequest.Contract:type_name -> proto.SmartContract
	0,  // 3: proto.ValidateContractRequest.Contract:type_name -> proto.SmartContract
	2,  // 4: proto.ConsensusRequest.NodeID:type_name -> proto.NodeIdentifier
	0,  // 5: proto.ConsensusRequest.Contract:type_name -> proto.SmartContract
	3,  // 6: proto.NodeService.ProposeContract:input_type -> proto.ProposeContractRequest
	5,  // 7: proto.NodeService.ValidateContract:input_type -> proto.ValidateContractRequest
	7,  // 8: proto.NodeService.AchieveConsensus:input_type -> proto.ConsensusRequest
	4,  // 9: proto.NodeService.ProposeContract:output_type -> proto.ProposeContractResponse
	6,  // 10: proto.NodeService.ValidateContract:output_type -> proto.ValidateContractResponse
	8,  // 11: proto.NodeService.AchieveConsensus:output_type -> proto.ConsensusResponse
	9,  // [9:12] is the sub-list for method output_type
	6,  // [6:9] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_smartContract_proto_init() }
func file_smartContract_proto_init() {
	if File_smartContract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_smartContract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmartContract); i {
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
		file_smartContract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC01); i {
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
		file_smartContract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeIdentifier); i {
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
		file_smartContract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposeContractRequest); i {
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
		file_smartContract_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposeContractResponse); i {
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
		file_smartContract_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateContractRequest); i {
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
		file_smartContract_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateContractResponse); i {
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
		file_smartContract_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusRequest); i {
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
		file_smartContract_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusResponse); i {
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
			RawDescriptor: file_smartContract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_smartContract_proto_goTypes,
		DependencyIndexes: file_smartContract_proto_depIdxs,
		MessageInfos:      file_smartContract_proto_msgTypes,
	}.Build()
	File_smartContract_proto = out.File
	file_smartContract_proto_rawDesc = nil
	file_smartContract_proto_goTypes = nil
	file_smartContract_proto_depIdxs = nil
}