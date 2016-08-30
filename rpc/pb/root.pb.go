// Code generated by protoc-gen-go.
// source: root.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	root.proto

It has these top-level messages:
	LocationRequest
	ExecRequest
	ContentResponse
	StatusResponse
	DiffResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

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

// the stage from which this status response is being sent
type StatusResponse_Stage int32

const (
	StatusResponse_UNSPECIFIED_STAGE StatusResponse_Stage = 0
	StatusResponse_PLAN              StatusResponse_Stage = 1
	StatusResponse_APPLY             StatusResponse_Stage = 2
)

var StatusResponse_Stage_name = map[int32]string{
	0: "UNSPECIFIED_STAGE",
	1: "PLAN",
	2: "APPLY",
}
var StatusResponse_Stage_value = map[string]int32{
	"UNSPECIFIED_STAGE": 0,
	"PLAN":              1,
	"APPLY":             2,
}

func (x StatusResponse_Stage) String() string {
	return proto.EnumName(StatusResponse_Stage_name, int32(x))
}
func (StatusResponse_Stage) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

// when is this status response being sent?
type StatusResponse_Run int32

const (
	StatusResponse_UNSPECIFIED_RUN StatusResponse_Run = 0
	StatusResponse_STARTED         StatusResponse_Run = 1
	StatusResponse_FINISHED        StatusResponse_Run = 2
)

var StatusResponse_Run_name = map[int32]string{
	0: "UNSPECIFIED_RUN",
	1: "STARTED",
	2: "FINISHED",
}
var StatusResponse_Run_value = map[string]int32{
	"UNSPECIFIED_RUN": 0,
	"STARTED":         1,
	"FINISHED":        2,
}

func (x StatusResponse_Run) String() string {
	return proto.EnumName(StatusResponse_Run_name, int32(x))
}
func (StatusResponse_Run) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 1} }

type LocationRequest struct {
	Location string `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
}

func (m *LocationRequest) Reset()                    { *m = LocationRequest{} }
func (m *LocationRequest) String() string            { return proto.CompactTextString(m) }
func (*LocationRequest) ProtoMessage()               {}
func (*LocationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ExecRequest struct {
	Location   string            `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	Parameters map[string]string `protobuf:"bytes,2,rep,name=parameters" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ExecRequest) Reset()                    { *m = ExecRequest{} }
func (m *ExecRequest) String() string            { return proto.CompactTextString(m) }
func (*ExecRequest) ProtoMessage()               {}
func (*ExecRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ExecRequest) GetParameters() map[string]string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type ContentResponse struct {
	Content string `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
}

func (m *ContentResponse) Reset()                    { *m = ContentResponse{} }
func (m *ContentResponse) String() string            { return proto.CompactTextString(m) }
func (*ContentResponse) ProtoMessage()               {}
func (*ContentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type StatusResponse struct {
	Id      string                  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Stage   StatusResponse_Stage    `protobuf:"varint,2,opt,name=stage,enum=pb.StatusResponse_Stage" json:"stage,omitempty"`
	Run     StatusResponse_Run      `protobuf:"varint,3,opt,name=run,enum=pb.StatusResponse_Run" json:"run,omitempty"`
	Details *StatusResponse_Details `protobuf:"bytes,4,opt,name=details" json:"details,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StatusResponse) GetDetails() *StatusResponse_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

// the informational message, if present
type StatusResponse_Details struct {
	Messages   []string                 `protobuf:"bytes,1,rep,name=messages" json:"messages,omitempty"`
	Changes    map[string]*DiffResponse `protobuf:"bytes,2,rep,name=changes" json:"changes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	HasChanges bool                     `protobuf:"varint,3,opt,name=hasChanges" json:"hasChanges,omitempty"`
	Error      string                   `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
}

func (m *StatusResponse_Details) Reset()                    { *m = StatusResponse_Details{} }
func (m *StatusResponse_Details) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse_Details) ProtoMessage()               {}
func (*StatusResponse_Details) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *StatusResponse_Details) GetChanges() map[string]*DiffResponse {
	if m != nil {
		return m.Changes
	}
	return nil
}

type DiffResponse struct {
	Original string `protobuf:"bytes,1,opt,name=original" json:"original,omitempty"`
	Current  string `protobuf:"bytes,2,opt,name=current" json:"current,omitempty"`
	Changes  bool   `protobuf:"varint,3,opt,name=changes" json:"changes,omitempty"`
}

func (m *DiffResponse) Reset()                    { *m = DiffResponse{} }
func (m *DiffResponse) String() string            { return proto.CompactTextString(m) }
func (*DiffResponse) ProtoMessage()               {}
func (*DiffResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*LocationRequest)(nil), "pb.LocationRequest")
	proto.RegisterType((*ExecRequest)(nil), "pb.ExecRequest")
	proto.RegisterType((*ContentResponse)(nil), "pb.ContentResponse")
	proto.RegisterType((*StatusResponse)(nil), "pb.StatusResponse")
	proto.RegisterType((*StatusResponse_Details)(nil), "pb.StatusResponse.Details")
	proto.RegisterType((*DiffResponse)(nil), "pb.DiffResponse")
	proto.RegisterEnum("pb.StatusResponse_Stage", StatusResponse_Stage_name, StatusResponse_Stage_value)
	proto.RegisterEnum("pb.StatusResponse_Run", StatusResponse_Run_name, StatusResponse_Run_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Executor service

type ExecutorClient interface {
	// Plan out the execution of a module given by the location
	Plan(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (Executor_PlanClient, error)
	// Apply a module given by the location
	Apply(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (Executor_ApplyClient, error)
}

type executorClient struct {
	cc *grpc.ClientConn
}

func NewExecutorClient(cc *grpc.ClientConn) ExecutorClient {
	return &executorClient{cc}
}

func (c *executorClient) Plan(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (Executor_PlanClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Executor_serviceDesc.Streams[0], c.cc, "/pb.Executor/Plan", opts...)
	if err != nil {
		return nil, err
	}
	x := &executorPlanClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Executor_PlanClient interface {
	Recv() (*StatusResponse, error)
	grpc.ClientStream
}

type executorPlanClient struct {
	grpc.ClientStream
}

func (x *executorPlanClient) Recv() (*StatusResponse, error) {
	m := new(StatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *executorClient) Apply(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (Executor_ApplyClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Executor_serviceDesc.Streams[1], c.cc, "/pb.Executor/Apply", opts...)
	if err != nil {
		return nil, err
	}
	x := &executorApplyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Executor_ApplyClient interface {
	Recv() (*StatusResponse, error)
	grpc.ClientStream
}

type executorApplyClient struct {
	grpc.ClientStream
}

func (x *executorApplyClient) Recv() (*StatusResponse, error) {
	m := new(StatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Executor service

type ExecutorServer interface {
	// Plan out the execution of a module given by the location
	Plan(*ExecRequest, Executor_PlanServer) error
	// Apply a module given by the location
	Apply(*ExecRequest, Executor_ApplyServer) error
}

func RegisterExecutorServer(s *grpc.Server, srv ExecutorServer) {
	s.RegisterService(&_Executor_serviceDesc, srv)
}

func _Executor_Plan_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExecRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExecutorServer).Plan(m, &executorPlanServer{stream})
}

type Executor_PlanServer interface {
	Send(*StatusResponse) error
	grpc.ServerStream
}

type executorPlanServer struct {
	grpc.ServerStream
}

func (x *executorPlanServer) Send(m *StatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Executor_Apply_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExecRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExecutorServer).Apply(m, &executorApplyServer{stream})
}

type Executor_ApplyServer interface {
	Send(*StatusResponse) error
	grpc.ServerStream
}

type executorApplyServer struct {
	grpc.ServerStream
}

func (x *executorApplyServer) Send(m *StatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Executor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Executor",
	HandlerType: (*ExecutorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Plan",
			Handler:       _Executor_Plan_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Apply",
			Handler:       _Executor_Apply_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

// Client API for ResourceHost service

type ResourceHostClient interface {
	// GetBinary returns the converge binary itself
	GetBinary(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*ContentResponse, error)
	// GetModule gets the content of a module at the given path
	GetModule(ctx context.Context, in *LocationRequest, opts ...grpc.CallOption) (*ContentResponse, error)
}

type resourceHostClient struct {
	cc *grpc.ClientConn
}

func NewResourceHostClient(cc *grpc.ClientConn) ResourceHostClient {
	return &resourceHostClient{cc}
}

func (c *resourceHostClient) GetBinary(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*ContentResponse, error) {
	out := new(ContentResponse)
	err := grpc.Invoke(ctx, "/pb.ResourceHost/GetBinary", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceHostClient) GetModule(ctx context.Context, in *LocationRequest, opts ...grpc.CallOption) (*ContentResponse, error) {
	out := new(ContentResponse)
	err := grpc.Invoke(ctx, "/pb.ResourceHost/GetModule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ResourceHost service

type ResourceHostServer interface {
	// GetBinary returns the converge binary itself
	GetBinary(context.Context, *google_protobuf1.Empty) (*ContentResponse, error)
	// GetModule gets the content of a module at the given path
	GetModule(context.Context, *LocationRequest) (*ContentResponse, error)
}

func RegisterResourceHostServer(s *grpc.Server, srv ResourceHostServer) {
	s.RegisterService(&_ResourceHost_serviceDesc, srv)
}

func _ResourceHost_GetBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceHostServer).GetBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ResourceHost/GetBinary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceHostServer).GetBinary(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceHost_GetModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceHostServer).GetModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ResourceHost/GetModule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceHostServer).GetModule(ctx, req.(*LocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceHost_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ResourceHost",
	HandlerType: (*ResourceHostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBinary",
			Handler:    _ResourceHost_GetBinary_Handler,
		},
		{
			MethodName: "GetModule",
			Handler:    _ResourceHost_GetModule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("root.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 702 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xae, 0x9d, 0x84, 0x24, 0x27, 0x51, 0x63, 0xa6, 0x3f, 0x32, 0x06, 0xd1, 0xc8, 0x17, 0x34,
	0x14, 0x61, 0x43, 0x0a, 0x12, 0xaa, 0x54, 0xa1, 0xd0, 0xb8, 0x6d, 0xa4, 0x10, 0x45, 0x4e, 0x7b,
	0x01, 0x17, 0xa0, 0x49, 0x32, 0x49, 0x2d, 0x9c, 0x19, 0x33, 0x33, 0xae, 0x88, 0x10, 0x37, 0xdc,
	0xee, 0xe5, 0x3e, 0xc5, 0x4a, 0xfb, 0x34, 0xbb, 0xaf, 0xb0, 0x97, 0xfb, 0x10, 0x2b, 0x8f, 0xed,
	0xca, 0x4d, 0xb3, 0x52, 0xef, 0xfc, 0x9d, 0xf3, 0xcd, 0x77, 0x8e, 0xe7, 0x9b, 0x73, 0x00, 0x38,
	0x63, 0xd2, 0x89, 0x38, 0x93, 0x0c, 0xe9, 0xd1, 0xd4, 0xfa, 0x62, 0xc9, 0xd8, 0x32, 0x24, 0x2e,
	0x8e, 0x02, 0x17, 0x53, 0xca, 0x24, 0x96, 0x01, 0xa3, 0x22, 0x65, 0x58, 0x9f, 0x67, 0x59, 0x85,
	0xa6, 0xf1, 0xc2, 0x25, 0xab, 0x48, 0xae, 0xd3, 0xa4, 0xfd, 0x2d, 0xb4, 0x86, 0x6c, 0xa6, 0xf8,
	0x3e, 0xf9, 0x3b, 0x26, 0x42, 0x22, 0x0b, 0x6a, 0x61, 0x16, 0x32, 0xb5, 0xb6, 0xd6, 0xa9, 0xfb,
	0x0f, 0xd8, 0x7e, 0xa5, 0x41, 0xc3, 0xfb, 0x87, 0xcc, 0x9e, 0xc1, 0x45, 0x3f, 0x03, 0x44, 0x98,
	0xe3, 0x15, 0x91, 0x84, 0x0b, 0x53, 0x6f, 0x97, 0x3a, 0x8d, 0xee, 0x91, 0x13, 0x4d, 0x9d, 0x82,
	0x80, 0x33, 0x7e, 0x60, 0x78, 0x54, 0xf2, 0xb5, 0x5f, 0x38, 0x62, 0x9d, 0x43, 0x6b, 0x23, 0x8d,
	0x0c, 0x28, 0xfd, 0x45, 0xd6, 0x59, 0xa9, 0xe4, 0x13, 0xed, 0x43, 0xe5, 0x1e, 0x87, 0x31, 0x31,
	0x75, 0x15, 0x4b, 0xc1, 0x99, 0xfe, 0x93, 0x66, 0x7f, 0x03, 0xad, 0x0b, 0x46, 0x25, 0xa1, 0xd2,
	0x27, 0x22, 0x62, 0x54, 0x10, 0x64, 0x42, 0x75, 0x96, 0x86, 0x32, 0x89, 0x1c, 0xda, 0x2f, 0xca,
	0xb0, 0x3b, 0x91, 0x58, 0xc6, 0xe2, 0x81, 0xbc, 0x0b, 0x7a, 0x30, 0xcf, 0x78, 0x7a, 0x30, 0x47,
	0x0e, 0x54, 0x84, 0xc4, 0xcb, 0xb4, 0xd2, 0x6e, 0xd7, 0x4c, 0x7e, 0xe5, 0xf1, 0x91, 0x04, 0x2e,
	0x89, 0x9f, 0xd2, 0x50, 0x07, 0x4a, 0x3c, 0xa6, 0x66, 0x49, 0xb1, 0x0f, 0xb7, 0xb0, 0xfd, 0x98,
	0xfa, 0x09, 0x05, 0xfd, 0x00, 0xd5, 0x39, 0x91, 0x38, 0x08, 0x85, 0x59, 0x6e, 0x6b, 0x9d, 0x46,
	0xd7, 0xda, 0xc2, 0xee, 0xa7, 0x0c, 0x3f, 0xa7, 0x5a, 0xef, 0x35, 0xa8, 0x66, 0xc1, 0xc4, 0x87,
	0x15, 0x11, 0x02, 0x2f, 0x89, 0x30, 0xb5, 0x76, 0x29, 0xf1, 0x21, 0xc7, 0xa8, 0x07, 0xd5, 0xd9,
	0x1d, 0xa6, 0x49, 0x2a, 0x35, 0xe1, 0xf8, 0xe3, 0xea, 0xce, 0x45, 0xca, 0x4c, 0xcd, 0xc8, 0xcf,
	0xa1, 0x2f, 0x01, 0xee, 0xb0, 0xc8, 0x72, 0xea, 0x8f, 0x6a, 0x7e, 0x21, 0x92, 0x98, 0x40, 0x38,
	0x67, 0x5c, 0xb5, 0x5f, 0xf7, 0x53, 0x60, 0x0d, 0xa1, 0x59, 0x94, 0xdb, 0x62, 0xde, 0x57, 0x45,
	0xf3, 0x1a, 0x5d, 0x23, 0x69, 0xac, 0x1f, 0x2c, 0x16, 0x79, 0x5b, 0x45, 0x3b, 0x4f, 0xa1, 0xa2,
	0xae, 0x17, 0x1d, 0xc0, 0xa7, 0xb7, 0xa3, 0xc9, 0xd8, 0xbb, 0x18, 0x5c, 0x0e, 0xbc, 0xfe, 0x9f,
	0x93, 0x9b, 0xde, 0x95, 0x67, 0xec, 0xa0, 0x1a, 0x94, 0xc7, 0xc3, 0xde, 0xc8, 0xd0, 0x50, 0x1d,
	0x2a, 0xbd, 0xf1, 0x78, 0xf8, 0x9b, 0xa1, 0xdb, 0x3f, 0x42, 0xc9, 0x8f, 0x29, 0xda, 0x83, 0x56,
	0xf1, 0x88, 0x7f, 0x3b, 0x32, 0x76, 0x50, 0x03, 0xaa, 0x93, 0x9b, 0x9e, 0x7f, 0xe3, 0xf5, 0x0d,
	0x0d, 0x35, 0xa1, 0x76, 0x39, 0x18, 0x0d, 0x26, 0xd7, 0x5e, 0xdf, 0xd0, 0xed, 0x3f, 0xa0, 0x59,
	0x6c, 0x23, 0xb9, 0x5e, 0xc6, 0x83, 0x65, 0x40, 0x71, 0x98, 0x3f, 0xf3, 0x1c, 0xab, 0x37, 0x15,
	0x73, 0x9e, 0xbc, 0x29, 0x3d, 0x7b, 0x53, 0x29, 0x54, 0x99, 0x47, 0x57, 0x96, 0xc3, 0xee, 0x6b,
	0x0d, 0x6a, 0xc9, 0x14, 0xc4, 0x92, 0x71, 0x34, 0x82, 0xf2, 0x38, 0xc4, 0x14, 0xb5, 0x36, 0x66,
	0xc3, 0x42, 0x4f, 0x7d, 0xb2, 0x8f, 0xfe, 0x7f, 0xfb, 0xee, 0xa5, 0xfe, 0x99, 0xbd, 0xaf, 0x86,
	0xfd, 0xfe, 0x7b, 0x77, 0x85, 0x67, 0x77, 0x01, 0x25, 0x6e, 0x14, 0x62, 0x7a, 0xa6, 0x9d, 0x7c,
	0xa7, 0xa1, 0x31, 0x54, 0x7a, 0x51, 0x14, 0xae, 0x9f, 0x27, 0xd8, 0x56, 0x82, 0x96, 0x7d, 0xb0,
	0x29, 0x88, 0x13, 0x0d, 0xa5, 0xd8, 0x7d, 0xa3, 0x41, 0xd3, 0x27, 0x82, 0xc5, 0x7c, 0x46, 0xae,
	0x99, 0x90, 0xe8, 0x77, 0xa8, 0x5f, 0x11, 0xf9, 0x4b, 0x40, 0x31, 0x5f, 0xa3, 0x43, 0x27, 0x5d,
	0x30, 0x4e, 0xbe, 0x60, 0x1c, 0x2f, 0x59, 0x30, 0xd6, 0x5e, 0x52, 0x6d, 0x63, 0x02, 0xf3, 0x72,
	0xc8, 0xcc, 0xcb, 0xf1, 0x4c, 0x57, 0xb8, 0xd3, 0x54, 0x6e, 0xa1, 0xb4, 0x7f, 0x65, 0xf3, 0x38,
	0x24, 0x48, 0x69, 0x6c, 0x2c, 0xa8, 0xed, 0xc2, 0xae, 0x12, 0xfe, 0x1a, 0x1d, 0x3f, 0x15, 0x5e,
	0x29, 0x2d, 0xe1, 0xfe, 0x9b, 0xaf, 0xa6, 0xf3, 0x93, 0x93, 0xff, 0xa6, 0x9f, 0xa8, 0x76, 0x4f,
	0x3f, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xf3, 0x2c, 0xd6, 0x4d, 0x05, 0x00, 0x00,
}
