// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/campaign_label_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v1/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import status "google.golang.org/genproto/googleapis/rpc/status"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for [CampaignLabelService.GetCampaignLabel][google.ads.googleads.v1.services.CampaignLabelService.GetCampaignLabel].
type GetCampaignLabelRequest struct {
	// The resource name of the campaign-label relationship to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignLabelRequest) Reset()         { *m = GetCampaignLabelRequest{} }
func (m *GetCampaignLabelRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignLabelRequest) ProtoMessage()    {}
func (*GetCampaignLabelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_campaign_label_service_72c5db67fe58a597, []int{0}
}
func (m *GetCampaignLabelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignLabelRequest.Unmarshal(m, b)
}
func (m *GetCampaignLabelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignLabelRequest.Marshal(b, m, deterministic)
}
func (dst *GetCampaignLabelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignLabelRequest.Merge(dst, src)
}
func (m *GetCampaignLabelRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignLabelRequest.Size(m)
}
func (m *GetCampaignLabelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignLabelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignLabelRequest proto.InternalMessageInfo

func (m *GetCampaignLabelRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CampaignLabelService.MutateCampaignLabels][google.ads.googleads.v1.services.CampaignLabelService.MutateCampaignLabels].
type MutateCampaignLabelsRequest struct {
	// ID of the customer whose campaign-label relationships are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on campaign-label relationships.
	Operations []*CampaignLabelOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignLabelsRequest) Reset()         { *m = MutateCampaignLabelsRequest{} }
func (m *MutateCampaignLabelsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignLabelsRequest) ProtoMessage()    {}
func (*MutateCampaignLabelsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_campaign_label_service_72c5db67fe58a597, []int{1}
}
func (m *MutateCampaignLabelsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignLabelsRequest.Unmarshal(m, b)
}
func (m *MutateCampaignLabelsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignLabelsRequest.Marshal(b, m, deterministic)
}
func (dst *MutateCampaignLabelsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignLabelsRequest.Merge(dst, src)
}
func (m *MutateCampaignLabelsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignLabelsRequest.Size(m)
}
func (m *MutateCampaignLabelsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignLabelsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignLabelsRequest proto.InternalMessageInfo

func (m *MutateCampaignLabelsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCampaignLabelsRequest) GetOperations() []*CampaignLabelOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCampaignLabelsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCampaignLabelsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove) on a campaign-label relationship.
type CampaignLabelOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CampaignLabelOperation_Create
	//	*CampaignLabelOperation_Remove
	Operation            isCampaignLabelOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *CampaignLabelOperation) Reset()         { *m = CampaignLabelOperation{} }
func (m *CampaignLabelOperation) String() string { return proto.CompactTextString(m) }
func (*CampaignLabelOperation) ProtoMessage()    {}
func (*CampaignLabelOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_campaign_label_service_72c5db67fe58a597, []int{2}
}
func (m *CampaignLabelOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignLabelOperation.Unmarshal(m, b)
}
func (m *CampaignLabelOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignLabelOperation.Marshal(b, m, deterministic)
}
func (dst *CampaignLabelOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignLabelOperation.Merge(dst, src)
}
func (m *CampaignLabelOperation) XXX_Size() int {
	return xxx_messageInfo_CampaignLabelOperation.Size(m)
}
func (m *CampaignLabelOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignLabelOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignLabelOperation proto.InternalMessageInfo

type isCampaignLabelOperation_Operation interface {
	isCampaignLabelOperation_Operation()
}

type CampaignLabelOperation_Create struct {
	Create *resources.CampaignLabel `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignLabelOperation_Remove struct {
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*CampaignLabelOperation_Create) isCampaignLabelOperation_Operation() {}

func (*CampaignLabelOperation_Remove) isCampaignLabelOperation_Operation() {}

func (m *CampaignLabelOperation) GetOperation() isCampaignLabelOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CampaignLabelOperation) GetCreate() *resources.CampaignLabel {
	if x, ok := m.GetOperation().(*CampaignLabelOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CampaignLabelOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CampaignLabelOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CampaignLabelOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CampaignLabelOperation_OneofMarshaler, _CampaignLabelOperation_OneofUnmarshaler, _CampaignLabelOperation_OneofSizer, []interface{}{
		(*CampaignLabelOperation_Create)(nil),
		(*CampaignLabelOperation_Remove)(nil),
	}
}

func _CampaignLabelOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CampaignLabelOperation)
	// operation
	switch x := m.Operation.(type) {
	case *CampaignLabelOperation_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *CampaignLabelOperation_Remove:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Remove)
	case nil:
	default:
		return fmt.Errorf("CampaignLabelOperation.Operation has unexpected type %T", x)
	}
	return nil
}

func _CampaignLabelOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CampaignLabelOperation)
	switch tag {
	case 1: // operation.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.CampaignLabel)
		err := b.DecodeMessage(msg)
		m.Operation = &CampaignLabelOperation_Create{msg}
		return true, err
	case 2: // operation.remove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operation = &CampaignLabelOperation_Remove{x}
		return true, err
	default:
		return false, nil
	}
}

func _CampaignLabelOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CampaignLabelOperation)
	// operation
	switch x := m.Operation.(type) {
	case *CampaignLabelOperation_Create:
		s := proto.Size(x.Create)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignLabelOperation_Remove:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Remove)))
		n += len(x.Remove)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for a campaign labels mutate.
type MutateCampaignLabelsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCampaignLabelResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *MutateCampaignLabelsResponse) Reset()         { *m = MutateCampaignLabelsResponse{} }
func (m *MutateCampaignLabelsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignLabelsResponse) ProtoMessage()    {}
func (*MutateCampaignLabelsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_campaign_label_service_72c5db67fe58a597, []int{3}
}
func (m *MutateCampaignLabelsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignLabelsResponse.Unmarshal(m, b)
}
func (m *MutateCampaignLabelsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignLabelsResponse.Marshal(b, m, deterministic)
}
func (dst *MutateCampaignLabelsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignLabelsResponse.Merge(dst, src)
}
func (m *MutateCampaignLabelsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignLabelsResponse.Size(m)
}
func (m *MutateCampaignLabelsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignLabelsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignLabelsResponse proto.InternalMessageInfo

func (m *MutateCampaignLabelsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCampaignLabelsResponse) GetResults() []*MutateCampaignLabelResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for a campaign label mutate.
type MutateCampaignLabelResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignLabelResult) Reset()         { *m = MutateCampaignLabelResult{} }
func (m *MutateCampaignLabelResult) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignLabelResult) ProtoMessage()    {}
func (*MutateCampaignLabelResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_campaign_label_service_72c5db67fe58a597, []int{4}
}
func (m *MutateCampaignLabelResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignLabelResult.Unmarshal(m, b)
}
func (m *MutateCampaignLabelResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignLabelResult.Marshal(b, m, deterministic)
}
func (dst *MutateCampaignLabelResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignLabelResult.Merge(dst, src)
}
func (m *MutateCampaignLabelResult) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignLabelResult.Size(m)
}
func (m *MutateCampaignLabelResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignLabelResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignLabelResult proto.InternalMessageInfo

func (m *MutateCampaignLabelResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignLabelRequest)(nil), "google.ads.googleads.v1.services.GetCampaignLabelRequest")
	proto.RegisterType((*MutateCampaignLabelsRequest)(nil), "google.ads.googleads.v1.services.MutateCampaignLabelsRequest")
	proto.RegisterType((*CampaignLabelOperation)(nil), "google.ads.googleads.v1.services.CampaignLabelOperation")
	proto.RegisterType((*MutateCampaignLabelsResponse)(nil), "google.ads.googleads.v1.services.MutateCampaignLabelsResponse")
	proto.RegisterType((*MutateCampaignLabelResult)(nil), "google.ads.googleads.v1.services.MutateCampaignLabelResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CampaignLabelServiceClient is the client API for CampaignLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignLabelServiceClient interface {
	// Returns the requested campaign-label relationship in full detail.
	GetCampaignLabel(ctx context.Context, in *GetCampaignLabelRequest, opts ...grpc.CallOption) (*resources.CampaignLabel, error)
	// Creates and removes campaign-label relationships.
	// Operation statuses are returned.
	MutateCampaignLabels(ctx context.Context, in *MutateCampaignLabelsRequest, opts ...grpc.CallOption) (*MutateCampaignLabelsResponse, error)
}

type campaignLabelServiceClient struct {
	cc *grpc.ClientConn
}

func NewCampaignLabelServiceClient(cc *grpc.ClientConn) CampaignLabelServiceClient {
	return &campaignLabelServiceClient{cc}
}

func (c *campaignLabelServiceClient) GetCampaignLabel(ctx context.Context, in *GetCampaignLabelRequest, opts ...grpc.CallOption) (*resources.CampaignLabel, error) {
	out := new(resources.CampaignLabel)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignLabelService/GetCampaignLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignLabelServiceClient) MutateCampaignLabels(ctx context.Context, in *MutateCampaignLabelsRequest, opts ...grpc.CallOption) (*MutateCampaignLabelsResponse, error) {
	out := new(MutateCampaignLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignLabelService/MutateCampaignLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignLabelServiceServer is the server API for CampaignLabelService service.
type CampaignLabelServiceServer interface {
	// Returns the requested campaign-label relationship in full detail.
	GetCampaignLabel(context.Context, *GetCampaignLabelRequest) (*resources.CampaignLabel, error)
	// Creates and removes campaign-label relationships.
	// Operation statuses are returned.
	MutateCampaignLabels(context.Context, *MutateCampaignLabelsRequest) (*MutateCampaignLabelsResponse, error)
}

func RegisterCampaignLabelServiceServer(s *grpc.Server, srv CampaignLabelServiceServer) {
	s.RegisterService(&_CampaignLabelService_serviceDesc, srv)
}

func _CampaignLabelService_GetCampaignLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignLabelServiceServer).GetCampaignLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignLabelService/GetCampaignLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignLabelServiceServer).GetCampaignLabel(ctx, req.(*GetCampaignLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignLabelService_MutateCampaignLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignLabelServiceServer).MutateCampaignLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignLabelService/MutateCampaignLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignLabelServiceServer).MutateCampaignLabels(ctx, req.(*MutateCampaignLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignLabelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.CampaignLabelService",
	HandlerType: (*CampaignLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignLabel",
			Handler:    _CampaignLabelService_GetCampaignLabel_Handler,
		},
		{
			MethodName: "MutateCampaignLabels",
			Handler:    _CampaignLabelService_MutateCampaignLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/campaign_label_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/campaign_label_service.proto", fileDescriptor_campaign_label_service_72c5db67fe58a597)
}

var fileDescriptor_campaign_label_service_72c5db67fe58a597 = []byte{
	// 668 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6a, 0xd4, 0x4e,
	0x14, 0xff, 0x27, 0xfb, 0xa7, 0xda, 0xd9, 0xfa, 0xc1, 0x58, 0x6d, 0xba, 0x16, 0x5d, 0x62, 0xc1,
	0xb2, 0x17, 0x49, 0x77, 0x0b, 0x52, 0x53, 0xb6, 0xb8, 0x15, 0xdb, 0x2a, 0x6a, 0x4b, 0x8a, 0x45,
	0x64, 0x21, 0x4c, 0x93, 0x69, 0x08, 0x24, 0x99, 0x38, 0x33, 0x59, 0x29, 0xa5, 0x20, 0xde, 0x7a,
	0xe9, 0x1b, 0x78, 0xd9, 0x37, 0xf0, 0xc2, 0x17, 0xf0, 0x56, 0x7c, 0x03, 0xf1, 0xc2, 0xa7, 0x90,
	0x64, 0x32, 0xb1, 0x59, 0x77, 0x59, 0xed, 0xdd, 0x99, 0xf3, 0xf1, 0x3b, 0xe7, 0x77, 0x3e, 0x06,
	0x74, 0x7d, 0x42, 0xfc, 0x10, 0x9b, 0xc8, 0x63, 0xa6, 0x10, 0x33, 0x69, 0xd0, 0x36, 0x19, 0xa6,
	0x83, 0xc0, 0xc5, 0xcc, 0x74, 0x51, 0x94, 0xa0, 0xc0, 0x8f, 0x9d, 0x10, 0x1d, 0xe0, 0xd0, 0x29,
	0xf4, 0x46, 0x42, 0x09, 0x27, 0xb0, 0x29, 0x62, 0x0c, 0xe4, 0x31, 0xa3, 0x0c, 0x37, 0x06, 0x6d,
	0x43, 0x86, 0x37, 0xee, 0x8d, 0x4b, 0x40, 0x31, 0x23, 0x29, 0xfd, 0x33, 0x83, 0x40, 0x6e, 0x2c,
	0xc8, 0xb8, 0x24, 0x30, 0x51, 0x1c, 0x13, 0x8e, 0x78, 0x40, 0x62, 0x56, 0x58, 0x6f, 0x15, 0xd6,
	0xfc, 0x75, 0x90, 0x1e, 0x9a, 0x6f, 0x28, 0x4a, 0x12, 0x4c, 0xa5, 0x7d, 0xae, 0xb0, 0xd3, 0xc4,
	0x35, 0x19, 0x47, 0x3c, 0x2d, 0x0c, 0xfa, 0x3a, 0x98, 0xdb, 0xc2, 0xfc, 0x61, 0x91, 0xf1, 0x69,
	0x96, 0xd0, 0xc6, 0xaf, 0x53, 0xcc, 0x38, 0xbc, 0x03, 0x2e, 0xc9, 0x9a, 0x9c, 0x18, 0x45, 0x58,
	0x53, 0x9a, 0xca, 0xd2, 0xb4, 0x3d, 0x23, 0x95, 0xcf, 0x51, 0x84, 0xf5, 0x1f, 0x0a, 0xb8, 0xf9,
	0x2c, 0xe5, 0x88, 0xe3, 0x0a, 0x06, 0x93, 0x20, 0xb7, 0x41, 0xdd, 0x4d, 0x19, 0x27, 0x11, 0xa6,
	0x4e, 0xe0, 0x15, 0x10, 0x40, 0xaa, 0x1e, 0x7b, 0xf0, 0x25, 0x00, 0x24, 0xc1, 0x54, 0xb0, 0xd1,
	0xd4, 0x66, 0x6d, 0xa9, 0xde, 0x59, 0x35, 0x26, 0xb5, 0xd1, 0xa8, 0x64, 0xdb, 0x91, 0x00, 0xf6,
	0x19, 0x2c, 0x78, 0x17, 0x5c, 0x49, 0x10, 0xe5, 0x01, 0x0a, 0x9d, 0x43, 0x14, 0x84, 0x29, 0xc5,
	0x5a, 0xad, 0xa9, 0x2c, 0x5d, 0xb4, 0x2f, 0x17, 0xea, 0x4d, 0xa1, 0xcd, 0x88, 0x0e, 0x50, 0x18,
	0x78, 0x88, 0x63, 0x87, 0xc4, 0xe1, 0x91, 0xf6, 0x7f, 0xee, 0x36, 0x23, 0x95, 0x3b, 0x71, 0x78,
	0xa4, 0xbf, 0x57, 0xc0, 0x8d, 0xd1, 0x49, 0xe1, 0x13, 0x30, 0xe5, 0x52, 0x8c, 0xb8, 0xe8, 0x50,
	0xbd, 0xb3, 0x3c, 0xb6, 0xfc, 0x72, 0xc6, 0xd5, 0xfa, 0xb7, 0xff, 0xb3, 0x0b, 0x04, 0xa8, 0x81,
	0x29, 0x8a, 0x23, 0x32, 0xc0, 0x9a, 0x9a, 0xb5, 0x2a, 0xb3, 0x88, 0xf7, 0x46, 0x1d, 0x4c, 0x97,
	0xe4, 0xf4, 0xcf, 0x0a, 0x58, 0x18, 0xdd, 0x76, 0x96, 0x90, 0x98, 0x61, 0xb8, 0x09, 0xae, 0x0f,
	0x91, 0x77, 0x30, 0xa5, 0x84, 0xe6, 0x2d, 0xa8, 0x77, 0xa0, 0x2c, 0x91, 0x26, 0xae, 0xb1, 0x97,
	0x2f, 0x84, 0x7d, 0xad, 0xda, 0x96, 0x47, 0x99, 0x3b, 0x7c, 0x01, 0x2e, 0x50, 0xcc, 0xd2, 0x90,
	0xcb, 0xd9, 0xac, 0x4d, 0x9e, 0xcd, 0x88, 0xc2, 0xec, 0x1c, 0xc3, 0x96, 0x58, 0xfa, 0x03, 0x30,
	0x3f, 0xd6, 0xeb, 0xaf, 0x16, 0xaf, 0x73, 0x5a, 0x03, 0xb3, 0x95, 0xe0, 0x3d, 0x91, 0x1e, 0x7e,
	0x52, 0xc0, 0xd5, 0xe1, 0x95, 0x86, 0xf7, 0x27, 0x57, 0x3d, 0xe6, 0x0c, 0x1a, 0xff, 0x3c, 0x4d,
	0x7d, 0xf5, 0xdd, 0xd7, 0xef, 0x1f, 0xd4, 0x0e, 0x5c, 0xce, 0xce, 0xfa, 0xb8, 0x42, 0xa5, 0x2b,
	0x37, 0x9f, 0x99, 0xad, 0xf2, 0xce, 0xc5, 0xe8, 0xcc, 0xd6, 0x09, 0xfc, 0xa6, 0x80, 0xd9, 0x51,
	0x63, 0x85, 0xdd, 0x73, 0x75, 0x5d, 0x5e, 0x61, 0x63, 0xfd, 0xbc, 0xe1, 0x62, 0x9b, 0xf4, 0xf5,
	0x9c, 0xd1, 0xaa, 0xbe, 0x92, 0x31, 0xfa, 0x4d, 0xe1, 0xf8, 0xcc, 0x69, 0x77, 0x5b, 0x27, 0x43,
	0x84, 0xac, 0x28, 0x87, 0xb4, 0x94, 0xd6, 0xc6, 0x5b, 0x15, 0x2c, 0xba, 0x24, 0x9a, 0x58, 0xc5,
	0xc6, 0xfc, 0xa8, 0x91, 0xee, 0x66, 0x3f, 0xd5, 0xae, 0xf2, 0x6a, 0xbb, 0x08, 0xf7, 0x49, 0x88,
	0x62, 0xdf, 0x20, 0xd4, 0x37, 0x7d, 0x1c, 0xe7, 0xff, 0x98, 0xfc, 0x4a, 0x93, 0x80, 0x8d, 0xff,
	0xba, 0xd7, 0xa4, 0xf0, 0x51, 0xad, 0x6d, 0xf5, 0x7a, 0xa7, 0x6a, 0x73, 0x4b, 0x00, 0xf6, 0x3c,
	0x66, 0x08, 0x31, 0x93, 0xf6, 0xdb, 0x46, 0x91, 0x98, 0x7d, 0x91, 0x2e, 0xfd, 0x9e, 0xc7, 0xfa,
	0xa5, 0x4b, 0x7f, 0xbf, 0xdd, 0x97, 0x2e, 0x3f, 0xd5, 0x45, 0xa1, 0xb7, 0xac, 0x9e, 0xc7, 0x2c,
	0xab, 0x74, 0xb2, 0xac, 0xfd, 0xb6, 0x65, 0x49, 0xb7, 0x83, 0xa9, 0xbc, 0xce, 0x95, 0x5f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x6a, 0xf6, 0xcc, 0x51, 0x61, 0x06, 0x00, 0x00,
}
