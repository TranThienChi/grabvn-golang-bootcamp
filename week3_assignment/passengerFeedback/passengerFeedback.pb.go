// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passengerFeedback.proto

package passengerFeedback

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// PassengerFeedback contains a booking code, a passenger id and a feedback
type PassengerFeedback struct {
	BookingCode          string   `protobuf:"bytes,1,opt,name=bookingCode,proto3" json:"bookingCode,omitempty"`
	PassengerID          int32    `protobuf:"varint,2,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	Feedback             string   `protobuf:"bytes,3,opt,name=feedback,proto3" json:"feedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PassengerFeedback) Reset()         { *m = PassengerFeedback{} }
func (m *PassengerFeedback) String() string { return proto.CompactTextString(m) }
func (*PassengerFeedback) ProtoMessage()    {}
func (*PassengerFeedback) Descriptor() ([]byte, []int) {
	return fileDescriptor_8724f3b5de75954a, []int{0}
}

func (m *PassengerFeedback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PassengerFeedback.Unmarshal(m, b)
}
func (m *PassengerFeedback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PassengerFeedback.Marshal(b, m, deterministic)
}
func (m *PassengerFeedback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PassengerFeedback.Merge(m, src)
}
func (m *PassengerFeedback) XXX_Size() int {
	return xxx_messageInfo_PassengerFeedback.Size(m)
}
func (m *PassengerFeedback) XXX_DiscardUnknown() {
	xxx_messageInfo_PassengerFeedback.DiscardUnknown(m)
}

var xxx_messageInfo_PassengerFeedback proto.InternalMessageInfo

func (m *PassengerFeedback) GetBookingCode() string {
	if m != nil {
		return m.BookingCode
	}
	return ""
}

func (m *PassengerFeedback) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

func (m *PassengerFeedback) GetFeedback() string {
	if m != nil {
		return m.Feedback
	}
	return ""
}

// AddPassengerFeedbackRequest the request passenger feedback
type AddPassengerFeedbackRequest struct {
	PassengerFeedback    *PassengerFeedback `protobuf:"bytes,1,opt,name=passengerFeedback,proto3" json:"passengerFeedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AddPassengerFeedbackRequest) Reset()         { *m = AddPassengerFeedbackRequest{} }
func (m *AddPassengerFeedbackRequest) String() string { return proto.CompactTextString(m) }
func (*AddPassengerFeedbackRequest) ProtoMessage()    {}
func (*AddPassengerFeedbackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8724f3b5de75954a, []int{1}
}

func (m *AddPassengerFeedbackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPassengerFeedbackRequest.Unmarshal(m, b)
}
func (m *AddPassengerFeedbackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPassengerFeedbackRequest.Marshal(b, m, deterministic)
}
func (m *AddPassengerFeedbackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPassengerFeedbackRequest.Merge(m, src)
}
func (m *AddPassengerFeedbackRequest) XXX_Size() int {
	return xxx_messageInfo_AddPassengerFeedbackRequest.Size(m)
}
func (m *AddPassengerFeedbackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPassengerFeedbackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddPassengerFeedbackRequest proto.InternalMessageInfo

func (m *AddPassengerFeedbackRequest) GetPassengerFeedback() *PassengerFeedback {
	if m != nil {
		return m.PassengerFeedback
	}
	return nil
}

// AddPassengerFeedbackResponse the response passenger feedback
type AddPassengerFeedbackResponse struct {
	PassengerFeedback    *PassengerFeedback `protobuf:"bytes,1,opt,name=passengerFeedback,proto3" json:"passengerFeedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AddPassengerFeedbackResponse) Reset()         { *m = AddPassengerFeedbackResponse{} }
func (m *AddPassengerFeedbackResponse) String() string { return proto.CompactTextString(m) }
func (*AddPassengerFeedbackResponse) ProtoMessage()    {}
func (*AddPassengerFeedbackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8724f3b5de75954a, []int{2}
}

func (m *AddPassengerFeedbackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPassengerFeedbackResponse.Unmarshal(m, b)
}
func (m *AddPassengerFeedbackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPassengerFeedbackResponse.Marshal(b, m, deterministic)
}
func (m *AddPassengerFeedbackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPassengerFeedbackResponse.Merge(m, src)
}
func (m *AddPassengerFeedbackResponse) XXX_Size() int {
	return xxx_messageInfo_AddPassengerFeedbackResponse.Size(m)
}
func (m *AddPassengerFeedbackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPassengerFeedbackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddPassengerFeedbackResponse proto.InternalMessageInfo

func (m *AddPassengerFeedbackResponse) GetPassengerFeedback() *PassengerFeedback {
	if m != nil {
		return m.PassengerFeedback
	}
	return nil
}

// GetByPassengerIdRequest the request passenger id
type GetByPassengerIdRequest struct {
	PassengerID          int32    `protobuf:"varint,1,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByPassengerIdRequest) Reset()         { *m = GetByPassengerIdRequest{} }
func (m *GetByPassengerIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetByPassengerIdRequest) ProtoMessage()    {}
func (*GetByPassengerIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8724f3b5de75954a, []int{3}
}

func (m *GetByPassengerIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByPassengerIdRequest.Unmarshal(m, b)
}
func (m *GetByPassengerIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByPassengerIdRequest.Marshal(b, m, deterministic)
}
func (m *GetByPassengerIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByPassengerIdRequest.Merge(m, src)
}
func (m *GetByPassengerIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetByPassengerIdRequest.Size(m)
}
func (m *GetByPassengerIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByPassengerIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetByPassengerIdRequest proto.InternalMessageInfo

func (m *GetByPassengerIdRequest) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

// GetByPassengerIdResponse the response passenger id
type GetByPassengerIdResponse struct {
	PassengerFeedbacks   []*PassengerFeedback `protobuf:"bytes,1,rep,name=passengerFeedbacks,proto3" json:"passengerFeedbacks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetByPassengerIdResponse) Reset()         { *m = GetByPassengerIdResponse{} }
func (m *GetByPassengerIdResponse) String() string { return proto.CompactTextString(m) }
func (*GetByPassengerIdResponse) ProtoMessage()    {}
func (*GetByPassengerIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8724f3b5de75954a, []int{4}
}

func (m *GetByPassengerIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByPassengerIdResponse.Unmarshal(m, b)
}
func (m *GetByPassengerIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByPassengerIdResponse.Marshal(b, m, deterministic)
}
func (m *GetByPassengerIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByPassengerIdResponse.Merge(m, src)
}
func (m *GetByPassengerIdResponse) XXX_Size() int {
	return xxx_messageInfo_GetByPassengerIdResponse.Size(m)
}
func (m *GetByPassengerIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByPassengerIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetByPassengerIdResponse proto.InternalMessageInfo

func (m *GetByPassengerIdResponse) GetPassengerFeedbacks() []*PassengerFeedback {
	if m != nil {
		return m.PassengerFeedbacks
	}
	return nil
}

// GetByBookingCodeRequest the request booking code
type GetByBookingCodeRequest struct {
	BookingCode          string   `protobuf:"bytes,1,opt,name=bookingCode,proto3" json:"bookingCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByBookingCodeRequest) Reset()         { *m = GetByBookingCodeRequest{} }
func (m *GetByBookingCodeRequest) String() string { return proto.CompactTextString(m) }
func (*GetByBookingCodeRequest) ProtoMessage()    {}
func (*GetByBookingCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8724f3b5de75954a, []int{5}
}

func (m *GetByBookingCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByBookingCodeRequest.Unmarshal(m, b)
}
func (m *GetByBookingCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByBookingCodeRequest.Marshal(b, m, deterministic)
}
func (m *GetByBookingCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByBookingCodeRequest.Merge(m, src)
}
func (m *GetByBookingCodeRequest) XXX_Size() int {
	return xxx_messageInfo_GetByBookingCodeRequest.Size(m)
}
func (m *GetByBookingCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByBookingCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetByBookingCodeRequest proto.InternalMessageInfo

func (m *GetByBookingCodeRequest) GetBookingCode() string {
	if m != nil {
		return m.BookingCode
	}
	return ""
}

// GetByBookingCodeResponse the response booking code
type GetByBookingCodeResponse struct {
	PassengerFeedback    *PassengerFeedback `protobuf:"bytes,1,opt,name=passengerFeedback,proto3" json:"passengerFeedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetByBookingCodeResponse) Reset()         { *m = GetByBookingCodeResponse{} }
func (m *GetByBookingCodeResponse) String() string { return proto.CompactTextString(m) }
func (*GetByBookingCodeResponse) ProtoMessage()    {}
func (*GetByBookingCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8724f3b5de75954a, []int{6}
}

func (m *GetByBookingCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByBookingCodeResponse.Unmarshal(m, b)
}
func (m *GetByBookingCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByBookingCodeResponse.Marshal(b, m, deterministic)
}
func (m *GetByBookingCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByBookingCodeResponse.Merge(m, src)
}
func (m *GetByBookingCodeResponse) XXX_Size() int {
	return xxx_messageInfo_GetByBookingCodeResponse.Size(m)
}
func (m *GetByBookingCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByBookingCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetByBookingCodeResponse proto.InternalMessageInfo

func (m *GetByBookingCodeResponse) GetPassengerFeedback() *PassengerFeedback {
	if m != nil {
		return m.PassengerFeedback
	}
	return nil
}

// DeleteByPassengerIdRequest the request delete feedback by passenger id
type DeleteByPassengerIdRequest struct {
	PassengerID          int32    `protobuf:"varint,1,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteByPassengerIdRequest) Reset()         { *m = DeleteByPassengerIdRequest{} }
func (m *DeleteByPassengerIdRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteByPassengerIdRequest) ProtoMessage()    {}
func (*DeleteByPassengerIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8724f3b5de75954a, []int{7}
}

func (m *DeleteByPassengerIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteByPassengerIdRequest.Unmarshal(m, b)
}
func (m *DeleteByPassengerIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteByPassengerIdRequest.Marshal(b, m, deterministic)
}
func (m *DeleteByPassengerIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteByPassengerIdRequest.Merge(m, src)
}
func (m *DeleteByPassengerIdRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteByPassengerIdRequest.Size(m)
}
func (m *DeleteByPassengerIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteByPassengerIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteByPassengerIdRequest proto.InternalMessageInfo

func (m *DeleteByPassengerIdRequest) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

// DeleteByPassengerIdResponse the response delete feedback by passenger id
type DeleteByPassengerIdResponse struct {
	TotalDeletedFeedback int32    `protobuf:"varint,1,opt,name=totalDeletedFeedback,proto3" json:"totalDeletedFeedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteByPassengerIdResponse) Reset()         { *m = DeleteByPassengerIdResponse{} }
func (m *DeleteByPassengerIdResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteByPassengerIdResponse) ProtoMessage()    {}
func (*DeleteByPassengerIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8724f3b5de75954a, []int{8}
}

func (m *DeleteByPassengerIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteByPassengerIdResponse.Unmarshal(m, b)
}
func (m *DeleteByPassengerIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteByPassengerIdResponse.Marshal(b, m, deterministic)
}
func (m *DeleteByPassengerIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteByPassengerIdResponse.Merge(m, src)
}
func (m *DeleteByPassengerIdResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteByPassengerIdResponse.Size(m)
}
func (m *DeleteByPassengerIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteByPassengerIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteByPassengerIdResponse proto.InternalMessageInfo

func (m *DeleteByPassengerIdResponse) GetTotalDeletedFeedback() int32 {
	if m != nil {
		return m.TotalDeletedFeedback
	}
	return 0
}

func init() {
	proto.RegisterType((*PassengerFeedback)(nil), "passengerFeedback.PassengerFeedback")
	proto.RegisterType((*AddPassengerFeedbackRequest)(nil), "passengerFeedback.AddPassengerFeedbackRequest")
	proto.RegisterType((*AddPassengerFeedbackResponse)(nil), "passengerFeedback.AddPassengerFeedbackResponse")
	proto.RegisterType((*GetByPassengerIdRequest)(nil), "passengerFeedback.GetByPassengerIdRequest")
	proto.RegisterType((*GetByPassengerIdResponse)(nil), "passengerFeedback.GetByPassengerIdResponse")
	proto.RegisterType((*GetByBookingCodeRequest)(nil), "passengerFeedback.GetByBookingCodeRequest")
	proto.RegisterType((*GetByBookingCodeResponse)(nil), "passengerFeedback.GetByBookingCodeResponse")
	proto.RegisterType((*DeleteByPassengerIdRequest)(nil), "passengerFeedback.DeleteByPassengerIdRequest")
	proto.RegisterType((*DeleteByPassengerIdResponse)(nil), "passengerFeedback.DeleteByPassengerIdResponse")
}

func init() { proto.RegisterFile("passengerFeedback.proto", fileDescriptor_8724f3b5de75954a) }

var fileDescriptor_8724f3b5de75954a = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x51, 0x4b, 0xf3, 0x30,
	0x14, 0x5d, 0xbf, 0xb1, 0x0f, 0xbd, 0x7b, 0x71, 0x71, 0xb0, 0xd0, 0xf9, 0x50, 0x82, 0x0f, 0x43,
	0xb1, 0x42, 0x7d, 0x14, 0x04, 0xe7, 0x50, 0xf6, 0xa6, 0xd5, 0x3f, 0xd0, 0x2e, 0xd7, 0x31, 0x36,
	0x9b, 0xae, 0x89, 0x83, 0xfd, 0x66, 0xff, 0x84, 0xd8, 0x75, 0xb5, 0x6b, 0xd2, 0x51, 0x84, 0x3d,
	0xf6, 0xe6, 0xdc, 0x73, 0xce, 0xbd, 0x39, 0x0d, 0xf4, 0xe2, 0x40, 0x4a, 0x8c, 0xa6, 0x98, 0x3c,
	0x22, 0xf2, 0x30, 0x98, 0xcc, 0xdd, 0x38, 0x11, 0x4a, 0x90, 0x8e, 0x76, 0xc0, 0x24, 0x74, 0x9e,
	0xcb, 0x45, 0xe2, 0x40, 0x3b, 0x14, 0x62, 0x3e, 0x8b, 0xa6, 0x0f, 0x82, 0x23, 0xb5, 0x1c, 0x6b,
	0x70, 0xec, 0x17, 0x4b, 0x3f, 0x88, 0x9c, 0x6b, 0x3c, 0xa2, 0xff, 0x1c, 0x6b, 0xd0, 0xf2, 0x8b,
	0x25, 0x62, 0xc3, 0xd1, 0x7b, 0xc6, 0x47, 0x9b, 0x29, 0x41, 0xfe, 0xcd, 0x96, 0xd0, 0xbf, 0xe7,
	0x5c, 0xd3, 0xf5, 0x71, 0xf9, 0x89, 0x52, 0x11, 0x1f, 0x74, 0xa3, 0xa9, 0x89, 0xb6, 0x77, 0xee,
	0xea, 0xb3, 0xe9, 0x3c, 0x86, 0x39, 0x13, 0x38, 0x33, 0x4b, 0xca, 0x58, 0x44, 0x12, 0x0f, 0xa2,
	0x79, 0x0b, 0xbd, 0x27, 0x54, 0xc3, 0x75, 0x0e, 0x1e, 0xf3, 0xed, 0x88, 0xa5, 0xfd, 0x59, 0xda,
	0xfe, 0x58, 0x0c, 0x54, 0x6f, 0xce, 0xcc, 0xbe, 0x01, 0xd1, 0xd4, 0x24, 0xb5, 0x9c, 0x66, 0x6d,
	0xb7, 0x86, 0xfe, 0xdc, 0xee, 0xf0, 0xf7, 0x9e, 0x0b, 0x76, 0xf7, 0x07, 0x82, 0x45, 0x99, 0xdd,
	0x9d, 0xe6, 0x03, 0xee, 0xf6, 0x0e, 0xec, 0x11, 0x2e, 0x50, 0xe1, 0x1f, 0xd7, 0xfb, 0x02, 0x7d,
	0x63, 0x7f, 0x66, 0xd9, 0x83, 0xae, 0x12, 0x2a, 0x58, 0x6c, 0x30, 0x7c, 0xc7, 0x75, 0xcb, 0x37,
	0x9e, 0x79, 0x5f, 0x4d, 0xa0, 0x9a, 0xf7, 0x57, 0x4c, 0x56, 0xb3, 0x09, 0x92, 0x35, 0x74, 0x4d,
	0xf9, 0x23, 0xae, 0x61, 0x01, 0x7b, 0xfe, 0x0d, 0xfb, 0xba, 0x36, 0x7e, 0x33, 0x09, 0x6b, 0x90,
	0x0f, 0x38, 0x29, 0x27, 0x89, 0x5c, 0x18, 0x68, 0x2a, 0xb2, 0x6a, 0x5f, 0xd6, 0xc2, 0x6a, 0x72,
	0x85, 0x24, 0x54, 0xcb, 0xe9, 0x59, 0xab, 0x96, 0x33, 0x44, 0x8b, 0x35, 0xc8, 0x0a, 0x4e, 0x0d,
	0x17, 0x49, 0xae, 0x0c, 0x2c, 0xd5, 0x81, 0xb1, 0xdd, 0xba, 0xf0, 0xad, 0x6e, 0xf8, 0x3f, 0x7d,
	0x52, 0x6f, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x49, 0x72, 0xe7, 0xc9, 0x6d, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PassengerFeedbackServiceClient is the client API for PassengerFeedbackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PassengerFeedbackServiceClient interface {
	AddPassengerFeedback(ctx context.Context, in *AddPassengerFeedbackRequest, opts ...grpc.CallOption) (*AddPassengerFeedbackResponse, error)
	GetByPassengerId(ctx context.Context, in *GetByPassengerIdRequest, opts ...grpc.CallOption) (*GetByPassengerIdResponse, error)
	GetByBookingCode(ctx context.Context, in *GetByBookingCodeRequest, opts ...grpc.CallOption) (*GetByBookingCodeResponse, error)
	DeleteByPassengerId(ctx context.Context, in *DeleteByPassengerIdRequest, opts ...grpc.CallOption) (*DeleteByPassengerIdResponse, error)
}

type passengerFeedbackServiceClient struct {
	cc *grpc.ClientConn
}

func NewPassengerFeedbackServiceClient(cc *grpc.ClientConn) PassengerFeedbackServiceClient {
	return &passengerFeedbackServiceClient{cc}
}

func (c *passengerFeedbackServiceClient) AddPassengerFeedback(ctx context.Context, in *AddPassengerFeedbackRequest, opts ...grpc.CallOption) (*AddPassengerFeedbackResponse, error) {
	out := new(AddPassengerFeedbackResponse)
	err := c.cc.Invoke(ctx, "/passengerFeedback.PassengerFeedbackService/AddPassengerFeedback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerFeedbackServiceClient) GetByPassengerId(ctx context.Context, in *GetByPassengerIdRequest, opts ...grpc.CallOption) (*GetByPassengerIdResponse, error) {
	out := new(GetByPassengerIdResponse)
	err := c.cc.Invoke(ctx, "/passengerFeedback.PassengerFeedbackService/GetByPassengerId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerFeedbackServiceClient) GetByBookingCode(ctx context.Context, in *GetByBookingCodeRequest, opts ...grpc.CallOption) (*GetByBookingCodeResponse, error) {
	out := new(GetByBookingCodeResponse)
	err := c.cc.Invoke(ctx, "/passengerFeedback.PassengerFeedbackService/GetByBookingCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerFeedbackServiceClient) DeleteByPassengerId(ctx context.Context, in *DeleteByPassengerIdRequest, opts ...grpc.CallOption) (*DeleteByPassengerIdResponse, error) {
	out := new(DeleteByPassengerIdResponse)
	err := c.cc.Invoke(ctx, "/passengerFeedback.PassengerFeedbackService/DeleteByPassengerId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PassengerFeedbackServiceServer is the server API for PassengerFeedbackService service.
type PassengerFeedbackServiceServer interface {
	AddPassengerFeedback(context.Context, *AddPassengerFeedbackRequest) (*AddPassengerFeedbackResponse, error)
	GetByPassengerId(context.Context, *GetByPassengerIdRequest) (*GetByPassengerIdResponse, error)
	GetByBookingCode(context.Context, *GetByBookingCodeRequest) (*GetByBookingCodeResponse, error)
	DeleteByPassengerId(context.Context, *DeleteByPassengerIdRequest) (*DeleteByPassengerIdResponse, error)
}

// UnimplementedPassengerFeedbackServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPassengerFeedbackServiceServer struct {
}

func (*UnimplementedPassengerFeedbackServiceServer) AddPassengerFeedback(ctx context.Context, req *AddPassengerFeedbackRequest) (*AddPassengerFeedbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPassengerFeedback not implemented")
}
func (*UnimplementedPassengerFeedbackServiceServer) GetByPassengerId(ctx context.Context, req *GetByPassengerIdRequest) (*GetByPassengerIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByPassengerId not implemented")
}
func (*UnimplementedPassengerFeedbackServiceServer) GetByBookingCode(ctx context.Context, req *GetByBookingCodeRequest) (*GetByBookingCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByBookingCode not implemented")
}
func (*UnimplementedPassengerFeedbackServiceServer) DeleteByPassengerId(ctx context.Context, req *DeleteByPassengerIdRequest) (*DeleteByPassengerIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByPassengerId not implemented")
}

func RegisterPassengerFeedbackServiceServer(s *grpc.Server, srv PassengerFeedbackServiceServer) {
	s.RegisterService(&_PassengerFeedbackService_serviceDesc, srv)
}

func _PassengerFeedbackService_AddPassengerFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPassengerFeedbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerFeedbackServiceServer).AddPassengerFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passengerFeedback.PassengerFeedbackService/AddPassengerFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerFeedbackServiceServer).AddPassengerFeedback(ctx, req.(*AddPassengerFeedbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassengerFeedbackService_GetByPassengerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByPassengerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerFeedbackServiceServer).GetByPassengerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passengerFeedback.PassengerFeedbackService/GetByPassengerId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerFeedbackServiceServer).GetByPassengerId(ctx, req.(*GetByPassengerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassengerFeedbackService_GetByBookingCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByBookingCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerFeedbackServiceServer).GetByBookingCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passengerFeedback.PassengerFeedbackService/GetByBookingCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerFeedbackServiceServer).GetByBookingCode(ctx, req.(*GetByBookingCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassengerFeedbackService_DeleteByPassengerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteByPassengerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerFeedbackServiceServer).DeleteByPassengerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passengerFeedback.PassengerFeedbackService/DeleteByPassengerId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerFeedbackServiceServer).DeleteByPassengerId(ctx, req.(*DeleteByPassengerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PassengerFeedbackService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "passengerFeedback.PassengerFeedbackService",
	HandlerType: (*PassengerFeedbackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPassengerFeedback",
			Handler:    _PassengerFeedbackService_AddPassengerFeedback_Handler,
		},
		{
			MethodName: "GetByPassengerId",
			Handler:    _PassengerFeedbackService_GetByPassengerId_Handler,
		},
		{
			MethodName: "GetByBookingCode",
			Handler:    _PassengerFeedbackService_GetByBookingCode_Handler,
		},
		{
			MethodName: "DeleteByPassengerId",
			Handler:    _PassengerFeedbackService_DeleteByPassengerId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "passengerFeedback.proto",
}