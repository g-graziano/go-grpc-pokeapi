// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/pokemon.proto

package lionparcel

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

type GetPokemonResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Height               int32    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Weight               int32    `protobuf:"varint,4,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPokemonResponse) Reset()         { *m = GetPokemonResponse{} }
func (m *GetPokemonResponse) String() string { return proto.CompactTextString(m) }
func (*GetPokemonResponse) ProtoMessage()    {}
func (*GetPokemonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a3d9abfa707192f, []int{0}
}

func (m *GetPokemonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPokemonResponse.Unmarshal(m, b)
}
func (m *GetPokemonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPokemonResponse.Marshal(b, m, deterministic)
}
func (m *GetPokemonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPokemonResponse.Merge(m, src)
}
func (m *GetPokemonResponse) XXX_Size() int {
	return xxx_messageInfo_GetPokemonResponse.Size(m)
}
func (m *GetPokemonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPokemonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPokemonResponse proto.InternalMessageInfo

func (m *GetPokemonResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetPokemonResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetPokemonResponse) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *GetPokemonResponse) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type GetPokemonInput struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPokemonInput) Reset()         { *m = GetPokemonInput{} }
func (m *GetPokemonInput) String() string { return proto.CompactTextString(m) }
func (*GetPokemonInput) ProtoMessage()    {}
func (*GetPokemonInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a3d9abfa707192f, []int{1}
}

func (m *GetPokemonInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPokemonInput.Unmarshal(m, b)
}
func (m *GetPokemonInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPokemonInput.Marshal(b, m, deterministic)
}
func (m *GetPokemonInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPokemonInput.Merge(m, src)
}
func (m *GetPokemonInput) XXX_Size() int {
	return xxx_messageInfo_GetPokemonInput.Size(m)
}
func (m *GetPokemonInput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPokemonInput.DiscardUnknown(m)
}

var xxx_messageInfo_GetPokemonInput proto.InternalMessageInfo

func (m *GetPokemonInput) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*GetPokemonResponse)(nil), "lionparcel.GetPokemonResponse")
	proto.RegisterType((*GetPokemonInput)(nil), "lionparcel.GetPokemonInput")
}

func init() { proto.RegisterFile("proto/pokemon.proto", fileDescriptor_4a3d9abfa707192f) }

var fileDescriptor_4a3d9abfa707192f = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xc8, 0xcf, 0x4e, 0xcd, 0xcd, 0xcf, 0xd3, 0x03, 0xf3, 0x84, 0xb8, 0x72, 0x32,
	0xf3, 0xf3, 0x0a, 0x12, 0x8b, 0x92, 0x53, 0x73, 0x94, 0x32, 0xb8, 0x84, 0xdc, 0x53, 0x4b, 0x02,
	0x20, 0xf2, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x4c, 0x99, 0x29, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89,
	0xb9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x18, 0x17, 0x5b, 0x46,
	0x6a, 0x66, 0x7a, 0x46, 0x89, 0x04, 0x33, 0x58, 0x1d, 0x94, 0x07, 0x12, 0x2f, 0x87, 0x88, 0xb3,
	0x40, 0xc4, 0x21, 0x3c, 0x25, 0x45, 0x2e, 0x7e, 0x84, 0x4d, 0x9e, 0x79, 0x05, 0xa5, 0x25, 0xe8,
	0xd6, 0x18, 0x45, 0x71, 0xb1, 0x43, 0xe5, 0x85, 0xfc, 0xb9, 0xb8, 0x10, 0xaa, 0x85, 0xa4, 0xf5,
	0x10, 0x4e, 0xd6, 0x43, 0x33, 0x45, 0x4a, 0x0e, 0xbb, 0x24, 0xcc, 0x33, 0x4a, 0x0c, 0x1a, 0x8c,
	0x06, 0x8c, 0x49, 0x6c, 0x60, 0xbf, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x70, 0x71, 0x2e,
	0x0f, 0x12, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PokemonClient is the client API for Pokemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PokemonClient interface {
	GetPokemon(ctx context.Context, opts ...grpc.CallOption) (Pokemon_GetPokemonClient, error)
}

type pokemonClient struct {
	cc grpc.ClientConnInterface
}

func NewPokemonClient(cc grpc.ClientConnInterface) PokemonClient {
	return &pokemonClient{cc}
}

func (c *pokemonClient) GetPokemon(ctx context.Context, opts ...grpc.CallOption) (Pokemon_GetPokemonClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Pokemon_serviceDesc.Streams[0], "/lionparcel.Pokemon/GetPokemon", opts...)
	if err != nil {
		return nil, err
	}
	x := &pokemonGetPokemonClient{stream}
	return x, nil
}

type Pokemon_GetPokemonClient interface {
	Send(*GetPokemonInput) error
	Recv() (*GetPokemonResponse, error)
	grpc.ClientStream
}

type pokemonGetPokemonClient struct {
	grpc.ClientStream
}

func (x *pokemonGetPokemonClient) Send(m *GetPokemonInput) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pokemonGetPokemonClient) Recv() (*GetPokemonResponse, error) {
	m := new(GetPokemonResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PokemonServer is the server API for Pokemon service.
type PokemonServer interface {
	GetPokemon(Pokemon_GetPokemonServer) error
}

// UnimplementedPokemonServer can be embedded to have forward compatible implementations.
type UnimplementedPokemonServer struct {
}

func (*UnimplementedPokemonServer) GetPokemon(srv Pokemon_GetPokemonServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPokemon not implemented")
}

func RegisterPokemonServer(s *grpc.Server, srv PokemonServer) {
	s.RegisterService(&_Pokemon_serviceDesc, srv)
}

func _Pokemon_GetPokemon_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PokemonServer).GetPokemon(&pokemonGetPokemonServer{stream})
}

type Pokemon_GetPokemonServer interface {
	Send(*GetPokemonResponse) error
	Recv() (*GetPokemonInput, error)
	grpc.ServerStream
}

type pokemonGetPokemonServer struct {
	grpc.ServerStream
}

func (x *pokemonGetPokemonServer) Send(m *GetPokemonResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pokemonGetPokemonServer) Recv() (*GetPokemonInput, error) {
	m := new(GetPokemonInput)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Pokemon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lionparcel.Pokemon",
	HandlerType: (*PokemonServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPokemon",
			Handler:       _Pokemon_GetPokemon_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/pokemon.proto",
}