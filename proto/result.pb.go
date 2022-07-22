// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: result.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type AR_Status int32

const (
	AR_Status_FAILURE AR_Status = 0
	AR_Status_SUCCESS AR_Status = 1
	AR_Status_UNKNOWN AR_Status = 2
	AR_Status_INVALID AR_Status = 3
)

// Enum value maps for AR_Status.
var (
	AR_Status_name = map[int32]string{
		0: "FAILURE",
		1: "SUCCESS",
		2: "UNKNOWN",
		3: "INVALID",
	}
	AR_Status_value = map[string]int32{
		"FAILURE": 0,
		"SUCCESS": 1,
		"UNKNOWN": 2,
		"INVALID": 3,
	}
)

func (x AR_Status) Enum() *AR_Status {
	p := new(AR_Status)
	*p = x
	return p
}

func (x AR_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AR_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_result_proto_enumTypes[0].Descriptor()
}

func (AR_Status) Type() protoreflect.EnumType {
	return &file_result_proto_enumTypes[0]
}

func (x AR_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AR_Status.Descriptor instead.
func (AR_Status) EnumDescriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{0}
}

type TrustVector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HardwareAuthenticity AR_Status `protobuf:"varint,1,opt,name=hardware_authenticity,json=hw-authenticity,proto3,enum=proto.AR_Status" json:"hardware_authenticity,omitempty"`
	SoftwareIntegrity    AR_Status `protobuf:"varint,2,opt,name=software_integrity,json=sw-integrity,proto3,enum=proto.AR_Status" json:"software_integrity,omitempty"`
	SoftwareUpToDateness AR_Status `protobuf:"varint,3,opt,name=software_up_to_dateness,json=sw-up-to-dateness,proto3,enum=proto.AR_Status" json:"software_up_to_dateness,omitempty"`
	ConfigIntegrity      AR_Status `protobuf:"varint,4,opt,name=config_integrity,json=config-integrity,proto3,enum=proto.AR_Status" json:"config_integrity,omitempty"`
	RuntimeIntegrity     AR_Status `protobuf:"varint,5,opt,name=runtime_integrity,json=runtime-integrity,proto3,enum=proto.AR_Status" json:"runtime_integrity,omitempty"`
	CertificationStatus  AR_Status `protobuf:"varint,6,opt,name=certification_status,json=certification-status,proto3,enum=proto.AR_Status" json:"certification_status,omitempty"`
}

func (x *TrustVector) Reset() {
	*x = TrustVector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrustVector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrustVector) ProtoMessage() {}

func (x *TrustVector) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrustVector.ProtoReflect.Descriptor instead.
func (*TrustVector) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{0}
}

func (x *TrustVector) GetHardwareAuthenticity() AR_Status {
	if x != nil {
		return x.HardwareAuthenticity
	}
	return AR_Status_FAILURE
}

func (x *TrustVector) GetSoftwareIntegrity() AR_Status {
	if x != nil {
		return x.SoftwareIntegrity
	}
	return AR_Status_FAILURE
}

func (x *TrustVector) GetSoftwareUpToDateness() AR_Status {
	if x != nil {
		return x.SoftwareUpToDateness
	}
	return AR_Status_FAILURE
}

func (x *TrustVector) GetConfigIntegrity() AR_Status {
	if x != nil {
		return x.ConfigIntegrity
	}
	return AR_Status_FAILURE
}

func (x *TrustVector) GetRuntimeIntegrity() AR_Status {
	if x != nil {
		return x.RuntimeIntegrity
	}
	return AR_Status_FAILURE
}

func (x *TrustVector) GetCertificationStatus() AR_Status {
	if x != nil {
		return x.CertificationStatus
	}
	return AR_Status_FAILURE
}

type EndorsedClaims struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HardwareDetails      *structpb.Struct `protobuf:"bytes,1,opt,name=hardware_details,json=hw-details,proto3" json:"hardware_details,omitempty"`
	SoftwareDetails      *structpb.Struct `protobuf:"bytes,2,opt,name=software_details,json=sw-details,proto3" json:"software_details,omitempty"`
	CertificationDetails *structpb.Struct `protobuf:"bytes,3,opt,name=certification_details,json=certification-details,proto3" json:"certification_details,omitempty"`
	ConfigDetails        *structpb.Struct `protobuf:"bytes,4,opt,name=config_details,json=config-details,proto3" json:"config_details,omitempty"`
}

func (x *EndorsedClaims) Reset() {
	*x = EndorsedClaims{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndorsedClaims) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndorsedClaims) ProtoMessage() {}

func (x *EndorsedClaims) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndorsedClaims.ProtoReflect.Descriptor instead.
func (*EndorsedClaims) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{1}
}

func (x *EndorsedClaims) GetHardwareDetails() *structpb.Struct {
	if x != nil {
		return x.HardwareDetails
	}
	return nil
}

func (x *EndorsedClaims) GetSoftwareDetails() *structpb.Struct {
	if x != nil {
		return x.SoftwareDetails
	}
	return nil
}

func (x *EndorsedClaims) GetCertificationDetails() *structpb.Struct {
	if x != nil {
		return x.CertificationDetails
	}
	return nil
}

func (x *EndorsedClaims) GetConfigDetails() *structpb.Struct {
	if x != nil {
		return x.ConfigDetails
	}
	return nil
}

type AttestationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status            AR_Status              `protobuf:"varint,1,opt,name=status,proto3,enum=proto.AR_Status" json:"status,omitempty"`
	TrustVector       *TrustVector           `protobuf:"bytes,2,opt,name=trust_vector,json=trust-vector,proto3" json:"trust_vector,omitempty"`
	RawEvidence       []byte                 `protobuf:"bytes,3,opt,name=raw_evidence,json=raw-evidence,proto3" json:"raw_evidence,omitempty"`
	Timestamp         *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	EndorsedClaims    *EndorsedClaims        `protobuf:"bytes,5,opt,name=endorsed_claims,json=endorsed-claims,proto3" json:"endorsed_claims,omitempty"`
	AppraisalPolicyID string                 `protobuf:"bytes,6,opt,name=AppraisalPolicyID,json=appraisal-policy-id,proto3" json:"AppraisalPolicyID,omitempty"`
	// Extension
	ProcessedEvidence *structpb.Struct `protobuf:"bytes,7,opt,name=processed_evidence,json=veraison-procesed-evidence,proto3" json:"processed_evidence,omitempty"`
}

func (x *AttestationResult) Reset() {
	*x = AttestationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestationResult) ProtoMessage() {}

func (x *AttestationResult) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestationResult.ProtoReflect.Descriptor instead.
func (*AttestationResult) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{2}
}

func (x *AttestationResult) GetStatus() AR_Status {
	if x != nil {
		return x.Status
	}
	return AR_Status_FAILURE
}

func (x *AttestationResult) GetTrustVector() *TrustVector {
	if x != nil {
		return x.TrustVector
	}
	return nil
}

func (x *AttestationResult) GetRawEvidence() []byte {
	if x != nil {
		return x.RawEvidence
	}
	return nil
}

func (x *AttestationResult) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *AttestationResult) GetEndorsedClaims() *EndorsedClaims {
	if x != nil {
		return x.EndorsedClaims
	}
	return nil
}

func (x *AttestationResult) GetAppraisalPolicyID() string {
	if x != nil {
		return x.AppraisalPolicyID
	}
	return ""
}

func (x *AttestationResult) GetProcessedEvidence() *structpb.Struct {
	if x != nil {
		return x.ProcessedEvidence
	}
	return nil
}

var File_result_proto protoreflect.FileDescriptor

var file_result_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x03, 0x0a, 0x0b, 0x54, 0x72, 0x75, 0x73, 0x74, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x40, 0x0a, 0x15, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x52, 0x5f, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0f, 0x68, 0x77, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x69, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x12, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61,
	0x72, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x52, 0x5f, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x73, 0x77, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x44, 0x0a, 0x17, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x75,
	0x70, 0x5f, 0x74, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x52, 0x5f, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x11, 0x73, 0x77, 0x2d, 0x75, 0x70, 0x2d, 0x74, 0x6f, 0x2d,
	0x64, 0x61, 0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x3c, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x52, 0x5f, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2d, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x69, 0x74, 0x79, 0x12, 0x3e, 0x0a, 0x11, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x52, 0x5f, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x11, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x69, 0x74, 0x79, 0x12, 0x44, 0x0a, 0x14, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x52, 0x5f,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x14, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x9e, 0x02, 0x0a,
	0x0e, 0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x64, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12,
	0x3d, 0x0a, 0x10, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x0a, 0x68, 0x77, 0x2d, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x3d,
	0x0a, 0x10, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x0a, 0x73, 0x77, 0x2d, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x4d, 0x0a,
	0x15, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x15, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x3f, 0x0a, 0x0e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2d, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x95, 0x03,
	0x0a, 0x11, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x52, 0x5f, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x36, 0x0a,
	0x0c, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x75, 0x73,
	0x74, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0c, 0x74, 0x72, 0x75, 0x73, 0x74, 0x2d, 0x76,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x61, 0x77, 0x5f, 0x65, 0x76, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x72, 0x61, 0x77,
	0x2d, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x3f, 0x0a, 0x0f, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x64, 0x5f,
	0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x64, 0x43, 0x6c, 0x61,
	0x69, 0x6d, 0x73, 0x52, 0x0f, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x64, 0x2d, 0x63, 0x6c,
	0x61, 0x69, 0x6d, 0x73, 0x12, 0x2e, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x72, 0x61, 0x69, 0x73, 0x61,
	0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x61, 0x70, 0x70, 0x72, 0x61, 0x69, 0x73, 0x61, 0x6c, 0x2d, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2d, 0x69, 0x64, 0x12, 0x4f, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x64, 0x5f, 0x65, 0x76, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x1a, 0x76, 0x65, 0x72, 0x61, 0x69,
	0x73, 0x6f, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x65, 0x64, 0x2d, 0x65, 0x76, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x2a, 0x3f, 0x0a, 0x09, 0x41, 0x52, 0x5f, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x03, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x65, 0x72, 0x61, 0x69, 0x73, 0x6f, 0x6e, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_result_proto_rawDescOnce sync.Once
	file_result_proto_rawDescData = file_result_proto_rawDesc
)

func file_result_proto_rawDescGZIP() []byte {
	file_result_proto_rawDescOnce.Do(func() {
		file_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_result_proto_rawDescData)
	})
	return file_result_proto_rawDescData
}

var file_result_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_result_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_result_proto_goTypes = []interface{}{
	(AR_Status)(0),                // 0: proto.AR_Status
	(*TrustVector)(nil),           // 1: proto.TrustVector
	(*EndorsedClaims)(nil),        // 2: proto.EndorsedClaims
	(*AttestationResult)(nil),     // 3: proto.AttestationResult
	(*structpb.Struct)(nil),       // 4: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_result_proto_depIdxs = []int32{
	0,  // 0: proto.TrustVector.hardware_authenticity:type_name -> proto.AR_Status
	0,  // 1: proto.TrustVector.software_integrity:type_name -> proto.AR_Status
	0,  // 2: proto.TrustVector.software_up_to_dateness:type_name -> proto.AR_Status
	0,  // 3: proto.TrustVector.config_integrity:type_name -> proto.AR_Status
	0,  // 4: proto.TrustVector.runtime_integrity:type_name -> proto.AR_Status
	0,  // 5: proto.TrustVector.certification_status:type_name -> proto.AR_Status
	4,  // 6: proto.EndorsedClaims.hardware_details:type_name -> google.protobuf.Struct
	4,  // 7: proto.EndorsedClaims.software_details:type_name -> google.protobuf.Struct
	4,  // 8: proto.EndorsedClaims.certification_details:type_name -> google.protobuf.Struct
	4,  // 9: proto.EndorsedClaims.config_details:type_name -> google.protobuf.Struct
	0,  // 10: proto.AttestationResult.status:type_name -> proto.AR_Status
	1,  // 11: proto.AttestationResult.trust_vector:type_name -> proto.TrustVector
	5,  // 12: proto.AttestationResult.timestamp:type_name -> google.protobuf.Timestamp
	2,  // 13: proto.AttestationResult.endorsed_claims:type_name -> proto.EndorsedClaims
	4,  // 14: proto.AttestationResult.processed_evidence:type_name -> google.protobuf.Struct
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_result_proto_init() }
func file_result_proto_init() {
	if File_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrustVector); i {
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
		file_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndorsedClaims); i {
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
		file_result_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestationResult); i {
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
			RawDescriptor: file_result_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_result_proto_goTypes,
		DependencyIndexes: file_result_proto_depIdxs,
		EnumInfos:         file_result_proto_enumTypes,
		MessageInfos:      file_result_proto_msgTypes,
	}.Build()
	File_result_proto = out.File
	file_result_proto_rawDesc = nil
	file_result_proto_goTypes = nil
	file_result_proto_depIdxs = nil
}
