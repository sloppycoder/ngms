// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package api_pb // import "github.com/sloppycoder/ngms/accounts/api"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type Account_Status int32

const (
	Account_ACTIVE  Account_Status = 0
	Account_BLOCKED Account_Status = 1
	Account_DORMANT Account_Status = 2
)

var Account_Status_name = map[int32]string{
	0: "ACTIVE",
	1: "BLOCKED",
	2: "DORMANT",
}
var Account_Status_value = map[string]int32{
	"ACTIVE":  0,
	"BLOCKED": 1,
	"DORMANT": 2,
}

func (x Account_Status) String() string {
	return proto.EnumName(Account_Status_name, int32(x))
}
func (Account_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_account_19bbc98033aa0f7d, []int{0, 0}
}

type Balance_Type int32

const (
	Balance_CURRENT   Balance_Type = 0
	Balance_AVAILABLE Balance_Type = 1
)

var Balance_Type_name = map[int32]string{
	0: "CURRENT",
	1: "AVAILABLE",
}
var Balance_Type_value = map[string]int32{
	"CURRENT":   0,
	"AVAILABLE": 1,
}

func (x Balance_Type) String() string {
	return proto.EnumName(Balance_Type_name, int32(x))
}
func (Balance_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_account_19bbc98033aa0f7d, []int{1, 0}
}

type Account struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountId            string               `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Nickname             string               `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	ProdCode             string               `protobuf:"bytes,4,opt,name=prod_code,json=prodCode,proto3" json:"prod_code,omitempty"`
	ProdName             string               `protobuf:"bytes,5,opt,name=prod_name,json=prodName,proto3" json:"prod_name,omitempty"`
	Currency             string               `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
	Servicer             string               `protobuf:"bytes,7,opt,name=servicer,proto3" json:"servicer,omitempty"`
	Status               Account_Status       `protobuf:"varint,8,opt,name=status,proto3,enum=ngms.accounts.Account_Status" json:"status,omitempty"`
	StatusLastUpdated    *timestamp.Timestamp `protobuf:"bytes,9,opt,name=status_last_updated,json=statusLastUpdated,proto3" json:"status_last_updated,omitempty"`
	Balances             []*Balance           `protobuf:"bytes,10,rep,name=balances,proto3" json:"balances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_19bbc98033aa0f7d, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (dst *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(dst, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Account) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *Account) GetProdCode() string {
	if m != nil {
		return m.ProdCode
	}
	return ""
}

func (m *Account) GetProdName() string {
	if m != nil {
		return m.ProdName
	}
	return ""
}

func (m *Account) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Account) GetServicer() string {
	if m != nil {
		return m.Servicer
	}
	return ""
}

func (m *Account) GetStatus() Account_Status {
	if m != nil {
		return m.Status
	}
	return Account_ACTIVE
}

func (m *Account) GetStatusLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.StatusLastUpdated
	}
	return nil
}

func (m *Account) GetBalances() []*Balance {
	if m != nil {
		return m.Balances
	}
	return nil
}

type Balance struct {
	Amount               float64              `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Type                 Balance_Type         `protobuf:"varint,2,opt,name=type,proto3,enum=ngms.accounts.Balance_Type" json:"type,omitempty"`
	CreditFlag           bool                 `protobuf:"varint,3,opt,name=credit_flag,json=creditFlag,proto3" json:"credit_flag,omitempty"`
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,4,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Balance) Reset()         { *m = Balance{} }
func (m *Balance) String() string { return proto.CompactTextString(m) }
func (*Balance) ProtoMessage()    {}
func (*Balance) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_19bbc98033aa0f7d, []int{1}
}
func (m *Balance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Balance.Unmarshal(m, b)
}
func (m *Balance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Balance.Marshal(b, m, deterministic)
}
func (dst *Balance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Balance.Merge(dst, src)
}
func (m *Balance) XXX_Size() int {
	return xxx_messageInfo_Balance.Size(m)
}
func (m *Balance) XXX_DiscardUnknown() {
	xxx_messageInfo_Balance.DiscardUnknown(m)
}

var xxx_messageInfo_Balance proto.InternalMessageInfo

func (m *Balance) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Balance) GetType() Balance_Type {
	if m != nil {
		return m.Type
	}
	return Balance_CURRENT
}

func (m *Balance) GetCreditFlag() bool {
	if m != nil {
		return m.CreditFlag
	}
	return false
}

func (m *Balance) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type GetAccountRequest struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountRequest) Reset()         { *m = GetAccountRequest{} }
func (m *GetAccountRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccountRequest) ProtoMessage()    {}
func (*GetAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_19bbc98033aa0f7d, []int{2}
}
func (m *GetAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountRequest.Unmarshal(m, b)
}
func (m *GetAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountRequest.Marshal(b, m, deterministic)
}
func (dst *GetAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountRequest.Merge(dst, src)
}
func (m *GetAccountRequest) XXX_Size() int {
	return xxx_messageInfo_GetAccountRequest.Size(m)
}
func (m *GetAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountRequest proto.InternalMessageInfo

func (m *GetAccountRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func init() {
	proto.RegisterType((*Account)(nil), "ngms.accounts.Account")
	proto.RegisterType((*Balance)(nil), "ngms.accounts.Balance")
	proto.RegisterType((*GetAccountRequest)(nil), "ngms.accounts.GetAccountRequest")
	proto.RegisterEnum("ngms.accounts.Account_Status", Account_Status_name, Account_Status_value)
	proto.RegisterEnum("ngms.accounts.Balance_Type", Balance_Type_name, Balance_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServiceClient interface {
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*Account, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/ngms.accounts.AccountService/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	GetAccount(context.Context, *GetAccountRequest) (*Account, error)
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ngms.accounts.AccountService/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ngms.accounts.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccount",
			Handler:    _AccountService_GetAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_account_19bbc98033aa0f7d) }

var fileDescriptor_account_19bbc98033aa0f7d = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0xd3, 0xe0, 0x24, 0x13, 0x1a, 0xa5, 0x8b, 0x14, 0x59, 0x2e, 0xa5, 0x91, 0x4f, 0x39,
	0xd9, 0xc2, 0x88, 0x13, 0xe2, 0xe0, 0xa4, 0x01, 0x15, 0x42, 0x2a, 0xb9, 0x69, 0x0f, 0x5c, 0xac,
	0x8d, 0xbd, 0x31, 0x2b, 0x6c, 0xef, 0xe2, 0x5d, 0x23, 0x22, 0xc4, 0x85, 0x5f, 0xe0, 0x3b, 0xf8,
	0x1a, 0x24, 0xbe, 0x80, 0x0f, 0x41, 0xf6, 0x3a, 0xb5, 0x1a, 0x2a, 0x71, 0xdb, 0x99, 0xf7, 0xde,
	0xec, 0xcc, 0x9b, 0x81, 0x23, 0x1c, 0x86, 0xac, 0xc8, 0xa4, 0xcd, 0x73, 0x26, 0x19, 0x3a, 0xca,
	0xe2, 0x54, 0xd8, 0x75, 0x4e, 0x98, 0x67, 0x31, 0x63, 0x71, 0x42, 0x9c, 0x0a, 0x5c, 0x17, 0x1b,
	0x47, 0xd2, 0x94, 0x08, 0x89, 0x53, 0xae, 0xf8, 0xe6, 0xe3, 0x9a, 0x80, 0x39, 0x75, 0x70, 0x96,
	0x31, 0x89, 0x25, 0x65, 0x99, 0xa8, 0xd1, 0x93, 0x7d, 0x39, 0x49, 0xb9, 0xdc, 0x2a, 0xd0, 0xfa,
	0x79, 0x08, 0x1d, 0x4f, 0x7d, 0x84, 0x06, 0xd0, 0xa2, 0x91, 0xa1, 0x8d, 0xb5, 0x49, 0xcf, 0x6f,
	0xd1, 0x08, 0x9d, 0x02, 0xd4, 0x3d, 0x04, 0x34, 0x32, 0x5a, 0x55, 0xbe, 0x57, 0x67, 0x2e, 0x22,
	0x64, 0x42, 0x37, 0xa3, 0xe1, 0xc7, 0x0c, 0xa7, 0xc4, 0x38, 0xac, 0xc0, 0xdb, 0x18, 0x9d, 0x40,
	0x8f, 0xe7, 0x2c, 0x0a, 0x42, 0x16, 0x11, 0xa3, 0xad, 0xc0, 0x32, 0x31, 0x63, 0x51, 0x03, 0x56,
	0xca, 0x07, 0x0d, 0xb8, 0x2c, 0x95, 0x26, 0x74, 0xc3, 0x22, 0xcf, 0x49, 0x16, 0x6e, 0x0d, 0x5d,
	0x61, 0xbb, 0xb8, 0xc4, 0x04, 0xc9, 0x3f, 0xd3, 0x90, 0xe4, 0x46, 0x47, 0x61, 0xbb, 0x18, 0x3d,
	0x07, 0x5d, 0x48, 0x2c, 0x0b, 0x61, 0x74, 0xc7, 0xda, 0x64, 0xe0, 0x9e, 0xda, 0x77, 0x4c, 0xb4,
	0xeb, 0x21, 0xed, 0xab, 0x8a, 0xe4, 0xd7, 0x64, 0xf4, 0x06, 0x1e, 0xa9, 0x57, 0x90, 0x60, 0x21,
	0x83, 0x82, 0x47, 0x58, 0x92, 0xc8, 0xe8, 0x8d, 0xb5, 0x49, 0xdf, 0x35, 0x6d, 0x65, 0x9d, 0xbd,
	0xb3, 0xce, 0x5e, 0xed, 0x9c, 0xf7, 0x8f, 0x95, 0x6c, 0x81, 0x85, 0xbc, 0x56, 0x22, 0xe4, 0x42,
	0x77, 0x8d, 0x13, 0x9c, 0x85, 0x44, 0x18, 0x30, 0x3e, 0x9c, 0xf4, 0xdd, 0xd1, 0x5e, 0x13, 0x53,
	0x05, 0xfb, 0xb7, 0x3c, 0xcb, 0x06, 0x5d, 0x75, 0x84, 0x00, 0x74, 0x6f, 0xb6, 0xba, 0xb8, 0x99,
	0x0f, 0x0f, 0x50, 0x1f, 0x3a, 0xd3, 0xc5, 0xe5, 0xec, 0xed, 0xfc, 0x7c, 0xa8, 0x95, 0xc1, 0xf9,
	0xa5, 0xff, 0xce, 0x5b, 0xae, 0x86, 0x2d, 0xeb, 0xb7, 0x06, 0x9d, 0xba, 0x0a, 0x1a, 0x81, 0x8e,
	0xd3, 0xb2, 0x70, 0xb5, 0x33, 0xcd, 0xaf, 0x23, 0xe4, 0x40, 0x5b, 0x6e, 0x39, 0xa9, 0x36, 0x36,
	0x70, 0x4f, 0xee, 0xef, 0xc1, 0x5e, 0x6d, 0x39, 0xf1, 0x2b, 0x22, 0x3a, 0x83, 0x7e, 0x98, 0x93,
	0x88, 0xca, 0x60, 0x93, 0xe0, 0xb8, 0x5a, 0x66, 0xd7, 0x07, 0x95, 0x7a, 0x95, 0xe0, 0x18, 0xbd,
	0x84, 0x87, 0x77, 0xec, 0x69, 0xff, 0xd7, 0x9e, 0x7e, 0xd2, 0x18, 0x63, 0x59, 0xd0, 0x2e, 0x7f,
	0x2b, 0x27, 0x99, 0x5d, 0xfb, 0xfe, 0x7c, 0xb9, 0x1a, 0x1e, 0xa0, 0x23, 0xe8, 0x79, 0x37, 0xde,
	0xc5, 0xc2, 0x9b, 0x2e, 0xe6, 0x43, 0xcd, 0x72, 0xe1, 0xf8, 0x35, 0x91, 0xf5, 0x96, 0x7c, 0xf2,
	0xa9, 0x20, 0x42, 0xee, 0x5d, 0xa0, 0xb6, 0x77, 0x81, 0xee, 0x17, 0x18, 0xd4, 0x82, 0x2b, 0x75,
	0x06, 0x68, 0x03, 0xd0, 0x54, 0x41, 0xe3, 0xbd, 0xd1, 0xff, 0xf9, 0xc0, 0x1c, 0xdd, 0x7f, 0x25,
	0xd6, 0x93, 0xef, 0xbf, 0xfe, 0xfc, 0x68, 0x19, 0x68, 0xe4, 0xec, 0x20, 0xe7, 0x6b, 0xd3, 0xc9,
	0xb7, 0xe9, 0xd3, 0xf7, 0x4e, 0x4c, 0xe5, 0x87, 0x62, 0x6d, 0x87, 0x2c, 0x75, 0x44, 0xc2, 0x38,
	0xdf, 0x96, 0xb7, 0x9e, 0x3b, 0x65, 0xbd, 0x46, 0x84, 0x39, 0x7d, 0x81, 0x39, 0x0d, 0xf8, 0x7a,
	0xad, 0x57, 0x2e, 0x3d, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x88, 0xa2, 0x98, 0x6e, 0xec, 0x03,
	0x00, 0x00,
}
