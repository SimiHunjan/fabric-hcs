// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/CryptoService.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

func init() {
	proto.RegisterFile("proto/CryptoService.proto", fileDescriptor_5cefbc36ab9ea863)
}

var fileDescriptor_5cefbc36ab9ea863 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4b, 0x6f, 0xe2, 0x30,
	0x10, 0x80, 0x4f, 0x8b, 0xd0, 0x2c, 0xb0, 0x4b, 0xb4, 0xcf, 0x5c, 0xf6, 0xba, 0x3d, 0x90, 0xa8,
	0x54, 0x6a, 0xd5, 0x4a, 0x7d, 0x05, 0x44, 0xc5, 0xb1, 0xc0, 0xa9, 0x37, 0x63, 0x0f, 0x4e, 0x04,
	0xc4, 0x91, 0x3d, 0xa9, 0xc8, 0xcf, 0xea, 0x3f, 0xac, 0x6a, 0x07, 0x5a, 0xe0, 0x90, 0xd2, 0x9e,
	0xac, 0xcc, 0x7c, 0xdf, 0xbc, 0x14, 0xf8, 0x9b, 0x69, 0x45, 0x2a, 0xec, 0xe9, 0x22, 0x23, 0x35,
	0x46, 0xfd, 0x98, 0x70, 0x0c, 0x6c, 0xcc, 0xfb, 0x62, 0x1f, 0xbf, 0xed, 0x88, 0xfb, 0x1c, 0x75,
	0xe1, 0x32, 0xfe, 0x0f, 0x17, 0x1a, 0xa1, 0xc9, 0x54, 0x6a, 0x4a, 0xde, 0xff, 0xe7, 0xa2, 0x13,
	0xcd, 0x52, 0xc3, 0x38, 0x25, 0x2a, 0xdd, 0x01, 0x7e, 0xef, 0x01, 0x2e, 0xd1, 0x7d, 0xaa, 0x41,
	0x73, 0x6b, 0x02, 0xef, 0x1a, 0x9a, 0x5c, 0x23, 0x23, 0xbc, 0xe5, 0x5c, 0xe5, 0x29, 0x79, 0x9e,
	0x43, 0x83, 0x37, 0xb2, 0xef, 0xef, 0xc7, 0xd6, 0x1d, 0x5f, 0x0a, 0xe4, 0x99, 0xf8, 0x44, 0x81,
	0x1b, 0x68, 0x71, 0x3b, 0x92, 0x4d, 0xce, 0x50, 0x1f, 0x5c, 0xe1, 0x0a, 0x1a, 0xae, 0x42, 0x1f,
	0x17, 0x48, 0x78, 0xb0, 0x7f, 0x01, 0x75, 0x26, 0x44, 0x6f, 0xc1, 0x92, 0xe5, 0xc1, 0xee, 0x25,
	0x7c, 0x15, 0xb6, 0xeb, 0xc7, 0xf4, 0x23, 0xa8, 0x4b, 0x24, 0xe7, 0x36, 0x4a, 0xce, 0xfe, 0x00,
	0xfe, 0xb7, 0xf2, 0x6b, 0x83, 0x76, 0xa1, 0x2d, 0x91, 0xca, 0x2b, 0x8f, 0x90, 0x2b, 0x2d, 0x4c,
	0x95, 0x73, 0x0c, 0xdf, 0xdd, 0x65, 0xee, 0x90, 0x22, 0xb6, 0x60, 0x29, 0xc7, 0x2a, 0x25, 0x84,
	0xd6, 0x6b, 0x9b, 0x61, 0x3a, 0x53, 0x55, 0xc2, 0x19, 0xfc, 0x92, 0x48, 0x5b, 0xcb, 0x71, 0x4c,
	0x32, 0xaa, 0x1c, 0xee, 0x1c, 0xfe, 0x48, 0xa4, 0x01, 0x33, 0x3b, 0xb2, 0xd2, 0xe2, 0x7d, 0xb7,
	0x98, 0xac, 0x1c, 0x1d, 0x15, 0x93, 0xd5, 0xb0, 0x5f, 0xe5, 0x9c, 0xc2, 0x4f, 0x89, 0x34, 0x26,
	0x36, 0x47, 0x6d, 0xa2, 0x62, 0xbd, 0x61, 0x95, 0x17, 0x0d, 0xc0, 0xe7, 0x6a, 0x19, 0xc4, 0x28,
	0x50, 0xb3, 0x20, 0x66, 0x26, 0x96, 0x9a, 0x65, 0xb1, 0xc3, 0x1e, 0xfe, 0xcb, 0x84, 0xe2, 0x7c,
	0x1a, 0x70, 0xb5, 0x0c, 0x37, 0xb9, 0xd0, 0xc1, 0x1d, 0x23, 0xe6, 0x1d, 0xa9, 0x42, 0x4b, 0x4e,
	0x6b, 0xf6, 0x39, 0x79, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xe8, 0x87, 0x84, 0x09, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CryptoServiceClient is the client API for CryptoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CryptoServiceClient interface {
	CreateAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	UpdateAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	CryptoTransfer(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	CryptoDelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	AddClaim(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	DeleteClaim(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	GetClaim(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	GetAccountRecords(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	CryptoGetBalance(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	GetAccountInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	GetTransactionReceipts(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	GetFastTransactionRecord(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	GetTxRecordByTxID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	GetStakersByAccountID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
}

type cryptoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCryptoServiceClient(cc grpc.ClientConnInterface) CryptoServiceClient {
	return &cryptoServiceClient{cc}
}

func (c *cryptoServiceClient) CreateAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/createAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) UpdateAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/updateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) CryptoTransfer(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/cryptoTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) CryptoDelete(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/cryptoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) AddClaim(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/addClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) DeleteClaim(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/deleteClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetClaim(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetAccountRecords(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getAccountRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) CryptoGetBalance(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/cryptoGetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetAccountInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getAccountInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetTransactionReceipts(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getTransactionReceipts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetFastTransactionRecord(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getFastTransactionRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetTxRecordByTxID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getTxRecordByTxID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) GetStakersByAccountID(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CryptoService/getStakersByAccountID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoServiceServer is the server API for CryptoService service.
type CryptoServiceServer interface {
	CreateAccount(context.Context, *Transaction) (*TransactionResponse, error)
	UpdateAccount(context.Context, *Transaction) (*TransactionResponse, error)
	CryptoTransfer(context.Context, *Transaction) (*TransactionResponse, error)
	CryptoDelete(context.Context, *Transaction) (*TransactionResponse, error)
	AddClaim(context.Context, *Transaction) (*TransactionResponse, error)
	DeleteClaim(context.Context, *Transaction) (*TransactionResponse, error)
	GetClaim(context.Context, *Query) (*Response, error)
	GetAccountRecords(context.Context, *Query) (*Response, error)
	CryptoGetBalance(context.Context, *Query) (*Response, error)
	GetAccountInfo(context.Context, *Query) (*Response, error)
	GetTransactionReceipts(context.Context, *Query) (*Response, error)
	GetFastTransactionRecord(context.Context, *Query) (*Response, error)
	GetTxRecordByTxID(context.Context, *Query) (*Response, error)
	GetStakersByAccountID(context.Context, *Query) (*Response, error)
}

// UnimplementedCryptoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCryptoServiceServer struct {
}

func (*UnimplementedCryptoServiceServer) CreateAccount(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (*UnimplementedCryptoServiceServer) UpdateAccount(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (*UnimplementedCryptoServiceServer) CryptoTransfer(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CryptoTransfer not implemented")
}
func (*UnimplementedCryptoServiceServer) CryptoDelete(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CryptoDelete not implemented")
}
func (*UnimplementedCryptoServiceServer) AddClaim(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddClaim not implemented")
}
func (*UnimplementedCryptoServiceServer) DeleteClaim(ctx context.Context, req *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClaim not implemented")
}
func (*UnimplementedCryptoServiceServer) GetClaim(ctx context.Context, req *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClaim not implemented")
}
func (*UnimplementedCryptoServiceServer) GetAccountRecords(ctx context.Context, req *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountRecords not implemented")
}
func (*UnimplementedCryptoServiceServer) CryptoGetBalance(ctx context.Context, req *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CryptoGetBalance not implemented")
}
func (*UnimplementedCryptoServiceServer) GetAccountInfo(ctx context.Context, req *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountInfo not implemented")
}
func (*UnimplementedCryptoServiceServer) GetTransactionReceipts(ctx context.Context, req *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionReceipts not implemented")
}
func (*UnimplementedCryptoServiceServer) GetFastTransactionRecord(ctx context.Context, req *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFastTransactionRecord not implemented")
}
func (*UnimplementedCryptoServiceServer) GetTxRecordByTxID(ctx context.Context, req *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxRecordByTxID not implemented")
}
func (*UnimplementedCryptoServiceServer) GetStakersByAccountID(ctx context.Context, req *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStakersByAccountID not implemented")
}

func RegisterCryptoServiceServer(s *grpc.Server, srv CryptoServiceServer) {
	s.RegisterService(&_CryptoService_serviceDesc, srv)
}

func _CryptoService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).CreateAccount(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).UpdateAccount(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_CryptoTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).CryptoTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/CryptoTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).CryptoTransfer(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_CryptoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).CryptoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/CryptoDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).CryptoDelete(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_AddClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).AddClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/AddClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).AddClaim(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_DeleteClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).DeleteClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/DeleteClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).DeleteClaim(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/GetClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetClaim(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetAccountRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetAccountRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/GetAccountRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetAccountRecords(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_CryptoGetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).CryptoGetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/CryptoGetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).CryptoGetBalance(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/GetAccountInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetAccountInfo(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetTransactionReceipts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetTransactionReceipts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/GetTransactionReceipts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetTransactionReceipts(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetFastTransactionRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetFastTransactionRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/GetFastTransactionRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetFastTransactionRecord(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetTxRecordByTxID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetTxRecordByTxID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/GetTxRecordByTxID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetTxRecordByTxID(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_GetStakersByAccountID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).GetStakersByAccountID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CryptoService/GetStakersByAccountID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).GetStakersByAccountID(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

var _CryptoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CryptoService",
	HandlerType: (*CryptoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createAccount",
			Handler:    _CryptoService_CreateAccount_Handler,
		},
		{
			MethodName: "updateAccount",
			Handler:    _CryptoService_UpdateAccount_Handler,
		},
		{
			MethodName: "cryptoTransfer",
			Handler:    _CryptoService_CryptoTransfer_Handler,
		},
		{
			MethodName: "cryptoDelete",
			Handler:    _CryptoService_CryptoDelete_Handler,
		},
		{
			MethodName: "addClaim",
			Handler:    _CryptoService_AddClaim_Handler,
		},
		{
			MethodName: "deleteClaim",
			Handler:    _CryptoService_DeleteClaim_Handler,
		},
		{
			MethodName: "getClaim",
			Handler:    _CryptoService_GetClaim_Handler,
		},
		{
			MethodName: "getAccountRecords",
			Handler:    _CryptoService_GetAccountRecords_Handler,
		},
		{
			MethodName: "cryptoGetBalance",
			Handler:    _CryptoService_CryptoGetBalance_Handler,
		},
		{
			MethodName: "getAccountInfo",
			Handler:    _CryptoService_GetAccountInfo_Handler,
		},
		{
			MethodName: "getTransactionReceipts",
			Handler:    _CryptoService_GetTransactionReceipts_Handler,
		},
		{
			MethodName: "getFastTransactionRecord",
			Handler:    _CryptoService_GetFastTransactionRecord_Handler,
		},
		{
			MethodName: "getTxRecordByTxID",
			Handler:    _CryptoService_GetTxRecordByTxID_Handler,
		},
		{
			MethodName: "getStakersByAccountID",
			Handler:    _CryptoService_GetStakersByAccountID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/CryptoService.proto",
}
