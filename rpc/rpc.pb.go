// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Block struct {
	Version              uint32   `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Flags                []string `protobuf:"bytes,2,rep,name=flags,proto3" json:"flags,omitempty"`
	Hash                 string   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	HashPrevBlock        string   `protobuf:"bytes,4,opt,name=hashPrevBlock,proto3" json:"hashPrevBlock,omitempty"`
	HashMerkleRoot       string   `protobuf:"bytes,5,opt,name=hashMerkleRoot,proto3" json:"hashMerkleRoot,omitempty"`
	Timestamp            int64    `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Height               uint32   `protobuf:"varint,7,opt,name=height,proto3" json:"height,omitempty"`
	Bits                 uint32   `protobuf:"varint,8,opt,name=bits,proto3" json:"bits,omitempty"`
	Nonce                uint64   `protobuf:"varint,9,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_6b9bdfdf35343c49, []int{0}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Block) GetFlags() []string {
	if m != nil {
		return m.Flags
	}
	return nil
}

func (m *Block) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Block) GetHashPrevBlock() string {
	if m != nil {
		return m.HashPrevBlock
	}
	return ""
}

func (m *Block) GetHashMerkleRoot() string {
	if m != nil {
		return m.HashMerkleRoot
	}
	return ""
}

func (m *Block) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Block) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Block) GetBits() uint32 {
	if m != nil {
		return m.Bits
	}
	return 0
}

func (m *Block) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type GetBlockHashAtHeightRequest struct {
	Height               uint32   `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockHashAtHeightRequest) Reset()         { *m = GetBlockHashAtHeightRequest{} }
func (m *GetBlockHashAtHeightRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockHashAtHeightRequest) ProtoMessage()    {}
func (*GetBlockHashAtHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_6b9bdfdf35343c49, []int{1}
}
func (m *GetBlockHashAtHeightRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockHashAtHeightRequest.Unmarshal(m, b)
}
func (m *GetBlockHashAtHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockHashAtHeightRequest.Marshal(b, m, deterministic)
}
func (dst *GetBlockHashAtHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockHashAtHeightRequest.Merge(dst, src)
}
func (m *GetBlockHashAtHeightRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlockHashAtHeightRequest.Size(m)
}
func (m *GetBlockHashAtHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockHashAtHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockHashAtHeightRequest proto.InternalMessageInfo

func (m *GetBlockHashAtHeightRequest) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

type GetBlockHashAtHeightReply struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockHashAtHeightReply) Reset()         { *m = GetBlockHashAtHeightReply{} }
func (m *GetBlockHashAtHeightReply) String() string { return proto.CompactTextString(m) }
func (*GetBlockHashAtHeightReply) ProtoMessage()    {}
func (*GetBlockHashAtHeightReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_6b9bdfdf35343c49, []int{2}
}
func (m *GetBlockHashAtHeightReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockHashAtHeightReply.Unmarshal(m, b)
}
func (m *GetBlockHashAtHeightReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockHashAtHeightReply.Marshal(b, m, deterministic)
}
func (dst *GetBlockHashAtHeightReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockHashAtHeightReply.Merge(dst, src)
}
func (m *GetBlockHashAtHeightReply) XXX_Size() int {
	return xxx_messageInfo_GetBlockHashAtHeightReply.Size(m)
}
func (m *GetBlockHashAtHeightReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockHashAtHeightReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockHashAtHeightReply proto.InternalMessageInfo

func (m *GetBlockHashAtHeightReply) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type GetBlockByHashRequest struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockByHashRequest) Reset()         { *m = GetBlockByHashRequest{} }
func (m *GetBlockByHashRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockByHashRequest) ProtoMessage()    {}
func (*GetBlockByHashRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_6b9bdfdf35343c49, []int{3}
}
func (m *GetBlockByHashRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockByHashRequest.Unmarshal(m, b)
}
func (m *GetBlockByHashRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockByHashRequest.Marshal(b, m, deterministic)
}
func (dst *GetBlockByHashRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockByHashRequest.Merge(dst, src)
}
func (m *GetBlockByHashRequest) XXX_Size() int {
	return xxx_messageInfo_GetBlockByHashRequest.Size(m)
}
func (m *GetBlockByHashRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockByHashRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockByHashRequest proto.InternalMessageInfo

func (m *GetBlockByHashRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type GetBlockByHashReply struct {
	Block                *Block   `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockByHashReply) Reset()         { *m = GetBlockByHashReply{} }
func (m *GetBlockByHashReply) String() string { return proto.CompactTextString(m) }
func (*GetBlockByHashReply) ProtoMessage()    {}
func (*GetBlockByHashReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_6b9bdfdf35343c49, []int{4}
}
func (m *GetBlockByHashReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBlockByHashReply.Unmarshal(m, b)
}
func (m *GetBlockByHashReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBlockByHashReply.Marshal(b, m, deterministic)
}
func (dst *GetBlockByHashReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockByHashReply.Merge(dst, src)
}
func (m *GetBlockByHashReply) XXX_Size() int {
	return xxx_messageInfo_GetBlockByHashReply.Size(m)
}
func (m *GetBlockByHashReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockByHashReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockByHashReply proto.InternalMessageInfo

func (m *GetBlockByHashReply) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type GetMainChainHeightRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMainChainHeightRequest) Reset()         { *m = GetMainChainHeightRequest{} }
func (m *GetMainChainHeightRequest) String() string { return proto.CompactTextString(m) }
func (*GetMainChainHeightRequest) ProtoMessage()    {}
func (*GetMainChainHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_6b9bdfdf35343c49, []int{5}
}
func (m *GetMainChainHeightRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMainChainHeightRequest.Unmarshal(m, b)
}
func (m *GetMainChainHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMainChainHeightRequest.Marshal(b, m, deterministic)
}
func (dst *GetMainChainHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMainChainHeightRequest.Merge(dst, src)
}
func (m *GetMainChainHeightRequest) XXX_Size() int {
	return xxx_messageInfo_GetMainChainHeightRequest.Size(m)
}
func (m *GetMainChainHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMainChainHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMainChainHeightRequest proto.InternalMessageInfo

type GetMainChainHeightReply struct {
	Height               uint32   `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMainChainHeightReply) Reset()         { *m = GetMainChainHeightReply{} }
func (m *GetMainChainHeightReply) String() string { return proto.CompactTextString(m) }
func (*GetMainChainHeightReply) ProtoMessage()    {}
func (*GetMainChainHeightReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_6b9bdfdf35343c49, []int{6}
}
func (m *GetMainChainHeightReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMainChainHeightReply.Unmarshal(m, b)
}
func (m *GetMainChainHeightReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMainChainHeightReply.Marshal(b, m, deterministic)
}
func (dst *GetMainChainHeightReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMainChainHeightReply.Merge(dst, src)
}
func (m *GetMainChainHeightReply) XXX_Size() int {
	return xxx_messageInfo_GetMainChainHeightReply.Size(m)
}
func (m *GetMainChainHeightReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMainChainHeightReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetMainChainHeightReply proto.InternalMessageInfo

func (m *GetMainChainHeightReply) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

type GetMainChainRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMainChainRequest) Reset()         { *m = GetMainChainRequest{} }
func (m *GetMainChainRequest) String() string { return proto.CompactTextString(m) }
func (*GetMainChainRequest) ProtoMessage()    {}
func (*GetMainChainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_6b9bdfdf35343c49, []int{7}
}
func (m *GetMainChainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMainChainRequest.Unmarshal(m, b)
}
func (m *GetMainChainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMainChainRequest.Marshal(b, m, deterministic)
}
func (dst *GetMainChainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMainChainRequest.Merge(dst, src)
}
func (m *GetMainChainRequest) XXX_Size() int {
	return xxx_messageInfo_GetMainChainRequest.Size(m)
}
func (m *GetMainChainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMainChainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMainChainRequest proto.InternalMessageInfo

type GetMainChainReply struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMainChainReply) Reset()         { *m = GetMainChainReply{} }
func (m *GetMainChainReply) String() string { return proto.CompactTextString(m) }
func (*GetMainChainReply) ProtoMessage()    {}
func (*GetMainChainReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_6b9bdfdf35343c49, []int{8}
}
func (m *GetMainChainReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMainChainReply.Unmarshal(m, b)
}
func (m *GetMainChainReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMainChainReply.Marshal(b, m, deterministic)
}
func (dst *GetMainChainReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMainChainReply.Merge(dst, src)
}
func (m *GetMainChainReply) XXX_Size() int {
	return xxx_messageInfo_GetMainChainReply.Size(m)
}
func (m *GetMainChainReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMainChainReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetMainChainReply proto.InternalMessageInfo

func (m *GetMainChainReply) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*Block)(nil), "rpc.Block")
	proto.RegisterType((*GetBlockHashAtHeightRequest)(nil), "rpc.GetBlockHashAtHeightRequest")
	proto.RegisterType((*GetBlockHashAtHeightReply)(nil), "rpc.GetBlockHashAtHeightReply")
	proto.RegisterType((*GetBlockByHashRequest)(nil), "rpc.GetBlockByHashRequest")
	proto.RegisterType((*GetBlockByHashReply)(nil), "rpc.GetBlockByHashReply")
	proto.RegisterType((*GetMainChainHeightRequest)(nil), "rpc.GetMainChainHeightRequest")
	proto.RegisterType((*GetMainChainHeightReply)(nil), "rpc.GetMainChainHeightReply")
	proto.RegisterType((*GetMainChainRequest)(nil), "rpc.GetMainChainRequest")
	proto.RegisterType((*GetMainChainReply)(nil), "rpc.GetMainChainReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlockchainClient is the client API for Blockchain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockchainClient interface {
	GetBlockHashAtHeight(ctx context.Context, in *GetBlockHashAtHeightRequest, opts ...grpc.CallOption) (*GetBlockHashAtHeightReply, error)
	GetBlockByHash(ctx context.Context, in *GetBlockByHashRequest, opts ...grpc.CallOption) (*GetBlockByHashReply, error)
	GetMainChainHeight(ctx context.Context, in *GetMainChainHeightRequest, opts ...grpc.CallOption) (*GetMainChainHeightReply, error)
	GetMainChain(ctx context.Context, in *GetMainChainRequest, opts ...grpc.CallOption) (*GetMainChainReply, error)
}

type blockchainClient struct {
	cc *grpc.ClientConn
}

func NewBlockchainClient(cc *grpc.ClientConn) BlockchainClient {
	return &blockchainClient{cc}
}

func (c *blockchainClient) GetBlockHashAtHeight(ctx context.Context, in *GetBlockHashAtHeightRequest, opts ...grpc.CallOption) (*GetBlockHashAtHeightReply, error) {
	out := new(GetBlockHashAtHeightReply)
	err := c.cc.Invoke(ctx, "/rpc.Blockchain/GetBlockHashAtHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetBlockByHash(ctx context.Context, in *GetBlockByHashRequest, opts ...grpc.CallOption) (*GetBlockByHashReply, error) {
	out := new(GetBlockByHashReply)
	err := c.cc.Invoke(ctx, "/rpc.Blockchain/GetBlockByHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetMainChainHeight(ctx context.Context, in *GetMainChainHeightRequest, opts ...grpc.CallOption) (*GetMainChainHeightReply, error) {
	out := new(GetMainChainHeightReply)
	err := c.cc.Invoke(ctx, "/rpc.Blockchain/GetMainChainHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetMainChain(ctx context.Context, in *GetMainChainRequest, opts ...grpc.CallOption) (*GetMainChainReply, error) {
	out := new(GetMainChainReply)
	err := c.cc.Invoke(ctx, "/rpc.Blockchain/GetMainChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockchainServer is the server API for Blockchain service.
type BlockchainServer interface {
	GetBlockHashAtHeight(context.Context, *GetBlockHashAtHeightRequest) (*GetBlockHashAtHeightReply, error)
	GetBlockByHash(context.Context, *GetBlockByHashRequest) (*GetBlockByHashReply, error)
	GetMainChainHeight(context.Context, *GetMainChainHeightRequest) (*GetMainChainHeightReply, error)
	GetMainChain(context.Context, *GetMainChainRequest) (*GetMainChainReply, error)
}

func RegisterBlockchainServer(s *grpc.Server, srv BlockchainServer) {
	s.RegisterService(&_Blockchain_serviceDesc, srv)
}

func _Blockchain_GetBlockHashAtHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockHashAtHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetBlockHashAtHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Blockchain/GetBlockHashAtHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetBlockHashAtHeight(ctx, req.(*GetBlockHashAtHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetBlockByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetBlockByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Blockchain/GetBlockByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetBlockByHash(ctx, req.(*GetBlockByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetMainChainHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMainChainHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetMainChainHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Blockchain/GetMainChainHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetMainChainHeight(ctx, req.(*GetMainChainHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetMainChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMainChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetMainChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Blockchain/GetMainChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetMainChain(ctx, req.(*GetMainChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Blockchain_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Blockchain",
	HandlerType: (*BlockchainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlockHashAtHeight",
			Handler:    _Blockchain_GetBlockHashAtHeight_Handler,
		},
		{
			MethodName: "GetBlockByHash",
			Handler:    _Blockchain_GetBlockByHash_Handler,
		},
		{
			MethodName: "GetMainChainHeight",
			Handler:    _Blockchain_GetMainChainHeight_Handler,
		},
		{
			MethodName: "GetMainChain",
			Handler:    _Blockchain_GetMainChain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_rpc_6b9bdfdf35343c49) }

var fileDescriptor_rpc_6b9bdfdf35343c49 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x41, 0xaf, 0xd2, 0x40,
	0x10, 0x76, 0x29, 0x05, 0x3b, 0x0a, 0x89, 0x23, 0xe0, 0x5a, 0x88, 0x69, 0x1a, 0xa3, 0x4d, 0x4c,
	0x30, 0x62, 0x8c, 0x67, 0xf1, 0x20, 0x17, 0x12, 0xb3, 0xf1, 0xe4, 0xad, 0x34, 0x2b, 0x6d, 0x28,
	0x6d, 0x6d, 0xf7, 0x91, 0xf0, 0x1f, 0xdf, 0x1f, 0x7a, 0xb7, 0x97, 0x9d, 0xbe, 0x02, 0x85, 0xc2,
	0xa9, 0x3b, 0xf3, 0x7d, 0xf3, 0xed, 0x37, 0xdf, 0xa6, 0x60, 0xe5, 0x59, 0x30, 0xcd, 0xf2, 0x54,
	0xa5, 0x68, 0xe4, 0x59, 0xe0, 0x3e, 0x30, 0x30, 0xe7, 0x71, 0x1a, 0x6c, 0x90, 0x43, 0x77, 0x27,
	0xf3, 0x22, 0x4a, 0x13, 0xce, 0x1c, 0xe6, 0xf5, 0x44, 0x55, 0xe2, 0x00, 0xcc, 0x7f, 0xb1, 0xbf,
	0x2e, 0x78, 0xcb, 0x31, 0x3c, 0x4b, 0x94, 0x05, 0x22, 0xb4, 0x43, 0xbf, 0x08, 0xb9, 0xe1, 0x30,
	0xcf, 0x12, 0x74, 0xc6, 0xf7, 0xd0, 0xd3, 0xdf, 0xdf, 0xb9, 0xdc, 0x91, 0x28, 0x6f, 0x13, 0x58,
	0x6f, 0xe2, 0x07, 0xe8, 0xeb, 0xc6, 0x52, 0xe6, 0x9b, 0x58, 0x8a, 0x34, 0x55, 0xdc, 0x24, 0xda,
	0x59, 0x17, 0x27, 0x60, 0xa9, 0x68, 0x2b, 0x0b, 0xe5, 0x6f, 0x33, 0xde, 0x71, 0x98, 0x67, 0x88,
	0x63, 0x03, 0x47, 0xd0, 0x09, 0x65, 0xb4, 0x0e, 0x15, 0xef, 0x92, 0xdd, 0xa7, 0x4a, 0xfb, 0x5a,
	0x45, 0xaa, 0xe0, 0xcf, 0xa9, 0x4b, 0x67, 0xbd, 0x41, 0x92, 0x26, 0x81, 0xe4, 0x96, 0xc3, 0xbc,
	0xb6, 0x28, 0x0b, 0xf7, 0x1b, 0x8c, 0x7f, 0x49, 0x45, 0x9e, 0x16, 0x7e, 0x11, 0xfe, 0x50, 0x0b,
	0x52, 0x10, 0xf2, 0xff, 0x9d, 0x2c, 0xd4, 0xc9, 0x05, 0xec, 0xf4, 0x02, 0xf7, 0x33, 0xbc, 0x6d,
	0x1e, 0xcb, 0xe2, 0xfd, 0x21, 0x15, 0x76, 0x4c, 0xc5, 0xfd, 0x04, 0xc3, 0x6a, 0x60, 0xbe, 0xd7,
	0x23, 0xd5, 0x0d, 0x4d, 0xe4, 0xef, 0xf0, 0xfa, 0x9c, 0xac, 0x75, 0x1d, 0x30, 0x57, 0x94, 0xa8,
	0xe6, 0xbe, 0x98, 0xc1, 0x54, 0xbf, 0x23, 0xb1, 0x44, 0x09, 0xb8, 0x63, 0xb2, 0xb5, 0xf4, 0xa3,
	0xe4, 0x67, 0xe8, 0x47, 0x49, 0x6d, 0x17, 0xf7, 0x0b, 0xbc, 0x69, 0x02, 0xb5, 0xf2, 0xb5, 0x35,
	0x87, 0x64, 0xe4, 0x30, 0x52, 0x29, 0x7d, 0x84, 0x57, 0xf5, 0xf6, 0x95, 0xad, 0x67, 0xf7, 0x2d,
	0x00, 0x32, 0x18, 0x68, 0x1e, 0xfe, 0x85, 0x41, 0x53, 0x6a, 0xe8, 0xd0, 0x26, 0x37, 0xde, 0xc1,
	0x7e, 0x77, 0x83, 0x91, 0xc5, 0x7b, 0xf7, 0x19, 0x2e, 0xa0, 0x5f, 0xcf, 0x0c, 0xed, 0xda, 0x4c,
	0x2d, 0x75, 0x9b, 0x37, 0x62, 0xa5, 0xd2, 0x1f, 0xc0, 0xcb, 0x9c, 0xf0, 0xe0, 0xa0, 0x39, 0x5d,
	0x7b, 0x72, 0x15, 0x2f, 0x55, 0xe7, 0xf0, 0xf2, 0x14, 0x44, 0x7e, 0xc1, 0xaf, 0x94, 0x46, 0x0d,
	0x08, 0x69, 0xac, 0x3a, 0xf4, 0xd3, 0x7e, 0x7d, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x09, 0x41, 0x55,
	0x08, 0xc1, 0x03, 0x00, 0x00,
}
