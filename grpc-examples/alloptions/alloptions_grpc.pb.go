// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package main

import (
	context "context"
	nrpc "github.com/teamlint/nrpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SvcCustomSubjectClient is the client API for SvcCustomSubject service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SvcCustomSubjectClient interface {
	MtSimpleReply(ctx context.Context, in *StringArg, opts ...grpc.CallOption) (*SimpleStringReply, error)
	MtVoidReply(ctx context.Context, in *StringArg, opts ...grpc.CallOption) (*nrpc.Void, error)
	MtNoRequest(ctx context.Context, in *nrpc.NoRequest, opts ...grpc.CallOption) (*SimpleStringReply, error)
	MtStreamedReply(ctx context.Context, in *StringArg, opts ...grpc.CallOption) (*SimpleStringReply, error)
	MtVoidReqStreamedReply(ctx context.Context, in *nrpc.Void, opts ...grpc.CallOption) (*SimpleStringReply, error)
}

type svcCustomSubjectClient struct {
	cc grpc.ClientConnInterface
}

func NewSvcCustomSubjectClient(cc grpc.ClientConnInterface) SvcCustomSubjectClient {
	return &svcCustomSubjectClient{cc}
}

var svcCustomSubjectMtSimpleReplyStreamDesc = &grpc.StreamDesc{
	StreamName: "MtSimpleReply",
}

func (c *svcCustomSubjectClient) MtSimpleReply(ctx context.Context, in *StringArg, opts ...grpc.CallOption) (*SimpleStringReply, error) {
	out := new(SimpleStringReply)
	err := c.cc.Invoke(ctx, "/main.SvcCustomSubject/MtSimpleReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var svcCustomSubjectMtVoidReplyStreamDesc = &grpc.StreamDesc{
	StreamName: "MtVoidReply",
}

func (c *svcCustomSubjectClient) MtVoidReply(ctx context.Context, in *StringArg, opts ...grpc.CallOption) (*nrpc.Void, error) {
	out := new(nrpc.Void)
	err := c.cc.Invoke(ctx, "/main.SvcCustomSubject/MtVoidReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var svcCustomSubjectMtNoRequestStreamDesc = &grpc.StreamDesc{
	StreamName: "MtNoRequest",
}

func (c *svcCustomSubjectClient) MtNoRequest(ctx context.Context, in *nrpc.NoRequest, opts ...grpc.CallOption) (*SimpleStringReply, error) {
	out := new(SimpleStringReply)
	err := c.cc.Invoke(ctx, "/main.SvcCustomSubject/MtNoRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var svcCustomSubjectMtStreamedReplyStreamDesc = &grpc.StreamDesc{
	StreamName: "MtStreamedReply",
}

func (c *svcCustomSubjectClient) MtStreamedReply(ctx context.Context, in *StringArg, opts ...grpc.CallOption) (*SimpleStringReply, error) {
	out := new(SimpleStringReply)
	err := c.cc.Invoke(ctx, "/main.SvcCustomSubject/MtStreamedReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var svcCustomSubjectMtVoidReqStreamedReplyStreamDesc = &grpc.StreamDesc{
	StreamName: "MtVoidReqStreamedReply",
}

func (c *svcCustomSubjectClient) MtVoidReqStreamedReply(ctx context.Context, in *nrpc.Void, opts ...grpc.CallOption) (*SimpleStringReply, error) {
	out := new(SimpleStringReply)
	err := c.cc.Invoke(ctx, "/main.SvcCustomSubject/MtVoidReqStreamedReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SvcCustomSubjectService is the service API for SvcCustomSubject service.
// Fields should be assigned to their respective handler implementations only before
// RegisterSvcCustomSubjectService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type SvcCustomSubjectService struct {
	MtSimpleReply          func(context.Context, *StringArg) (*SimpleStringReply, error)
	MtVoidReply            func(context.Context, *StringArg) (*nrpc.Void, error)
	MtNoRequest            func(context.Context, *nrpc.NoRequest) (*SimpleStringReply, error)
	MtStreamedReply        func(context.Context, *StringArg) (*SimpleStringReply, error)
	MtVoidReqStreamedReply func(context.Context, *nrpc.Void) (*SimpleStringReply, error)
}

func (s *SvcCustomSubjectService) mtSimpleReply(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.MtSimpleReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/main.SvcCustomSubject/MtSimpleReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.MtSimpleReply(ctx, req.(*StringArg))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SvcCustomSubjectService) mtVoidReply(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.MtVoidReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/main.SvcCustomSubject/MtVoidReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.MtVoidReply(ctx, req.(*StringArg))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SvcCustomSubjectService) mtNoRequest(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(nrpc.NoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.MtNoRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/main.SvcCustomSubject/MtNoRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.MtNoRequest(ctx, req.(*nrpc.NoRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SvcCustomSubjectService) mtStreamedReply(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.MtStreamedReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/main.SvcCustomSubject/MtStreamedReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.MtStreamedReply(ctx, req.(*StringArg))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SvcCustomSubjectService) mtVoidReqStreamedReply(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(nrpc.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.MtVoidReqStreamedReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/main.SvcCustomSubject/MtVoidReqStreamedReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.MtVoidReqStreamedReply(ctx, req.(*nrpc.Void))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterSvcCustomSubjectService registers a service implementation with a gRPC server.
func RegisterSvcCustomSubjectService(s grpc.ServiceRegistrar, srv *SvcCustomSubjectService) {
	srvCopy := *srv
	if srvCopy.MtSimpleReply == nil {
		srvCopy.MtSimpleReply = func(context.Context, *StringArg) (*SimpleStringReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method MtSimpleReply not implemented")
		}
	}
	if srvCopy.MtVoidReply == nil {
		srvCopy.MtVoidReply = func(context.Context, *StringArg) (*nrpc.Void, error) {
			return nil, status.Errorf(codes.Unimplemented, "method MtVoidReply not implemented")
		}
	}
	if srvCopy.MtNoRequest == nil {
		srvCopy.MtNoRequest = func(context.Context, *nrpc.NoRequest) (*SimpleStringReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method MtNoRequest not implemented")
		}
	}
	if srvCopy.MtStreamedReply == nil {
		srvCopy.MtStreamedReply = func(context.Context, *StringArg) (*SimpleStringReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method MtStreamedReply not implemented")
		}
	}
	if srvCopy.MtVoidReqStreamedReply == nil {
		srvCopy.MtVoidReqStreamedReply = func(context.Context, *nrpc.Void) (*SimpleStringReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method MtVoidReqStreamedReply not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "main.SvcCustomSubject",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "MtSimpleReply",
				Handler:    srvCopy.mtSimpleReply,
			},
			{
				MethodName: "MtVoidReply",
				Handler:    srvCopy.mtVoidReply,
			},
			{
				MethodName: "MtNoRequest",
				Handler:    srvCopy.mtNoRequest,
			},
			{
				MethodName: "MtStreamedReply",
				Handler:    srvCopy.mtStreamedReply,
			},
			{
				MethodName: "MtVoidReqStreamedReply",
				Handler:    srvCopy.mtVoidReqStreamedReply,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "examples/alloptions/alloptions.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewSvcCustomSubjectService creates a new SvcCustomSubjectService containing the
// implemented methods of the SvcCustomSubject service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewSvcCustomSubjectService(s interface{}) *SvcCustomSubjectService {
	ns := &SvcCustomSubjectService{}
	if h, ok := s.(interface {
		MtSimpleReply(context.Context, *StringArg) (*SimpleStringReply, error)
	}); ok {
		ns.MtSimpleReply = h.MtSimpleReply
	}
	if h, ok := s.(interface {
		MtVoidReply(context.Context, *StringArg) (*nrpc.Void, error)
	}); ok {
		ns.MtVoidReply = h.MtVoidReply
	}
	if h, ok := s.(interface {
		MtNoRequest(context.Context, *nrpc.NoRequest) (*SimpleStringReply, error)
	}); ok {
		ns.MtNoRequest = h.MtNoRequest
	}
	if h, ok := s.(interface {
		MtStreamedReply(context.Context, *StringArg) (*SimpleStringReply, error)
	}); ok {
		ns.MtStreamedReply = h.MtStreamedReply
	}
	if h, ok := s.(interface {
		MtVoidReqStreamedReply(context.Context, *nrpc.Void) (*SimpleStringReply, error)
	}); ok {
		ns.MtVoidReqStreamedReply = h.MtVoidReqStreamedReply
	}
	return ns
}

// UnstableSvcCustomSubjectService is the service API for SvcCustomSubject service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableSvcCustomSubjectService interface {
	MtSimpleReply(context.Context, *StringArg) (*SimpleStringReply, error)
	MtVoidReply(context.Context, *StringArg) (*nrpc.Void, error)
	MtNoRequest(context.Context, *nrpc.NoRequest) (*SimpleStringReply, error)
	MtStreamedReply(context.Context, *StringArg) (*SimpleStringReply, error)
	MtVoidReqStreamedReply(context.Context, *nrpc.Void) (*SimpleStringReply, error)
}

// SvcSubjectParamsClient is the client API for SvcSubjectParams service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SvcSubjectParamsClient interface {
	MtWithSubjectParams(ctx context.Context, in *nrpc.Void, opts ...grpc.CallOption) (*SimpleStringReply, error)
	MtStreamedReplyWithSubjectParams(ctx context.Context, in *nrpc.Void, opts ...grpc.CallOption) (*SimpleStringReply, error)
	MtNoReply(ctx context.Context, in *nrpc.Void, opts ...grpc.CallOption) (*nrpc.NoReply, error)
	MtNoRequestWParams(ctx context.Context, in *nrpc.NoRequest, opts ...grpc.CallOption) (*SimpleStringReply, error)
}

type svcSubjectParamsClient struct {
	cc grpc.ClientConnInterface
}

func NewSvcSubjectParamsClient(cc grpc.ClientConnInterface) SvcSubjectParamsClient {
	return &svcSubjectParamsClient{cc}
}

var svcSubjectParamsMtWithSubjectParamsStreamDesc = &grpc.StreamDesc{
	StreamName: "MtWithSubjectParams",
}

func (c *svcSubjectParamsClient) MtWithSubjectParams(ctx context.Context, in *nrpc.Void, opts ...grpc.CallOption) (*SimpleStringReply, error) {
	out := new(SimpleStringReply)
	err := c.cc.Invoke(ctx, "/main.SvcSubjectParams/MtWithSubjectParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var svcSubjectParamsMtStreamedReplyWithSubjectParamsStreamDesc = &grpc.StreamDesc{
	StreamName: "MtStreamedReplyWithSubjectParams",
}

func (c *svcSubjectParamsClient) MtStreamedReplyWithSubjectParams(ctx context.Context, in *nrpc.Void, opts ...grpc.CallOption) (*SimpleStringReply, error) {
	out := new(SimpleStringReply)
	err := c.cc.Invoke(ctx, "/main.SvcSubjectParams/MtStreamedReplyWithSubjectParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var svcSubjectParamsMtNoReplyStreamDesc = &grpc.StreamDesc{
	StreamName: "MtNoReply",
}

func (c *svcSubjectParamsClient) MtNoReply(ctx context.Context, in *nrpc.Void, opts ...grpc.CallOption) (*nrpc.NoReply, error) {
	out := new(nrpc.NoReply)
	err := c.cc.Invoke(ctx, "/main.SvcSubjectParams/MtNoReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var svcSubjectParamsMtNoRequestWParamsStreamDesc = &grpc.StreamDesc{
	StreamName: "MtNoRequestWParams",
}

func (c *svcSubjectParamsClient) MtNoRequestWParams(ctx context.Context, in *nrpc.NoRequest, opts ...grpc.CallOption) (*SimpleStringReply, error) {
	out := new(SimpleStringReply)
	err := c.cc.Invoke(ctx, "/main.SvcSubjectParams/MtNoRequestWParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SvcSubjectParamsService is the service API for SvcSubjectParams service.
// Fields should be assigned to their respective handler implementations only before
// RegisterSvcSubjectParamsService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type SvcSubjectParamsService struct {
	MtWithSubjectParams              func(context.Context, *nrpc.Void) (*SimpleStringReply, error)
	MtStreamedReplyWithSubjectParams func(context.Context, *nrpc.Void) (*SimpleStringReply, error)
	MtNoReply                        func(context.Context, *nrpc.Void) (*nrpc.NoReply, error)
	MtNoRequestWParams               func(context.Context, *nrpc.NoRequest) (*SimpleStringReply, error)
}

func (s *SvcSubjectParamsService) mtWithSubjectParams(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(nrpc.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.MtWithSubjectParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/main.SvcSubjectParams/MtWithSubjectParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.MtWithSubjectParams(ctx, req.(*nrpc.Void))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SvcSubjectParamsService) mtStreamedReplyWithSubjectParams(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(nrpc.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.MtStreamedReplyWithSubjectParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/main.SvcSubjectParams/MtStreamedReplyWithSubjectParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.MtStreamedReplyWithSubjectParams(ctx, req.(*nrpc.Void))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SvcSubjectParamsService) mtNoReply(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(nrpc.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.MtNoReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/main.SvcSubjectParams/MtNoReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.MtNoReply(ctx, req.(*nrpc.Void))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SvcSubjectParamsService) mtNoRequestWParams(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(nrpc.NoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.MtNoRequestWParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/main.SvcSubjectParams/MtNoRequestWParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.MtNoRequestWParams(ctx, req.(*nrpc.NoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterSvcSubjectParamsService registers a service implementation with a gRPC server.
func RegisterSvcSubjectParamsService(s grpc.ServiceRegistrar, srv *SvcSubjectParamsService) {
	srvCopy := *srv
	if srvCopy.MtWithSubjectParams == nil {
		srvCopy.MtWithSubjectParams = func(context.Context, *nrpc.Void) (*SimpleStringReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method MtWithSubjectParams not implemented")
		}
	}
	if srvCopy.MtStreamedReplyWithSubjectParams == nil {
		srvCopy.MtStreamedReplyWithSubjectParams = func(context.Context, *nrpc.Void) (*SimpleStringReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method MtStreamedReplyWithSubjectParams not implemented")
		}
	}
	if srvCopy.MtNoReply == nil {
		srvCopy.MtNoReply = func(context.Context, *nrpc.Void) (*nrpc.NoReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method MtNoReply not implemented")
		}
	}
	if srvCopy.MtNoRequestWParams == nil {
		srvCopy.MtNoRequestWParams = func(context.Context, *nrpc.NoRequest) (*SimpleStringReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method MtNoRequestWParams not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "main.SvcSubjectParams",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "MtWithSubjectParams",
				Handler:    srvCopy.mtWithSubjectParams,
			},
			{
				MethodName: "MtStreamedReplyWithSubjectParams",
				Handler:    srvCopy.mtStreamedReplyWithSubjectParams,
			},
			{
				MethodName: "MtNoReply",
				Handler:    srvCopy.mtNoReply,
			},
			{
				MethodName: "MtNoRequestWParams",
				Handler:    srvCopy.mtNoRequestWParams,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "examples/alloptions/alloptions.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewSvcSubjectParamsService creates a new SvcSubjectParamsService containing the
// implemented methods of the SvcSubjectParams service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewSvcSubjectParamsService(s interface{}) *SvcSubjectParamsService {
	ns := &SvcSubjectParamsService{}
	if h, ok := s.(interface {
		MtWithSubjectParams(context.Context, *nrpc.Void) (*SimpleStringReply, error)
	}); ok {
		ns.MtWithSubjectParams = h.MtWithSubjectParams
	}
	if h, ok := s.(interface {
		MtStreamedReplyWithSubjectParams(context.Context, *nrpc.Void) (*SimpleStringReply, error)
	}); ok {
		ns.MtStreamedReplyWithSubjectParams = h.MtStreamedReplyWithSubjectParams
	}
	if h, ok := s.(interface {
		MtNoReply(context.Context, *nrpc.Void) (*nrpc.NoReply, error)
	}); ok {
		ns.MtNoReply = h.MtNoReply
	}
	if h, ok := s.(interface {
		MtNoRequestWParams(context.Context, *nrpc.NoRequest) (*SimpleStringReply, error)
	}); ok {
		ns.MtNoRequestWParams = h.MtNoRequestWParams
	}
	return ns
}

// UnstableSvcSubjectParamsService is the service API for SvcSubjectParams service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableSvcSubjectParamsService interface {
	MtWithSubjectParams(context.Context, *nrpc.Void) (*SimpleStringReply, error)
	MtStreamedReplyWithSubjectParams(context.Context, *nrpc.Void) (*SimpleStringReply, error)
	MtNoReply(context.Context, *nrpc.Void) (*nrpc.NoReply, error)
	MtNoRequestWParams(context.Context, *nrpc.NoRequest) (*SimpleStringReply, error)
}

// NoRequestServiceClient is the client API for NoRequestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NoRequestServiceClient interface {
	MtNoRequest(ctx context.Context, in *nrpc.NoRequest, opts ...grpc.CallOption) (*SimpleStringReply, error)
}

type noRequestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNoRequestServiceClient(cc grpc.ClientConnInterface) NoRequestServiceClient {
	return &noRequestServiceClient{cc}
}

var noRequestServiceMtNoRequestStreamDesc = &grpc.StreamDesc{
	StreamName: "MtNoRequest",
}

func (c *noRequestServiceClient) MtNoRequest(ctx context.Context, in *nrpc.NoRequest, opts ...grpc.CallOption) (*SimpleStringReply, error) {
	out := new(SimpleStringReply)
	err := c.cc.Invoke(ctx, "/main.NoRequestService/MtNoRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoRequestServiceService is the service API for NoRequestService service.
// Fields should be assigned to their respective handler implementations only before
// RegisterNoRequestServiceService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type NoRequestServiceService struct {
	MtNoRequest func(context.Context, *nrpc.NoRequest) (*SimpleStringReply, error)
}

func (s *NoRequestServiceService) mtNoRequest(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(nrpc.NoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.MtNoRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/main.NoRequestService/MtNoRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.MtNoRequest(ctx, req.(*nrpc.NoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterNoRequestServiceService registers a service implementation with a gRPC server.
func RegisterNoRequestServiceService(s grpc.ServiceRegistrar, srv *NoRequestServiceService) {
	srvCopy := *srv
	if srvCopy.MtNoRequest == nil {
		srvCopy.MtNoRequest = func(context.Context, *nrpc.NoRequest) (*SimpleStringReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method MtNoRequest not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "main.NoRequestService",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "MtNoRequest",
				Handler:    srvCopy.mtNoRequest,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "examples/alloptions/alloptions.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewNoRequestServiceService creates a new NoRequestServiceService containing the
// implemented methods of the NoRequestService service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewNoRequestServiceService(s interface{}) *NoRequestServiceService {
	ns := &NoRequestServiceService{}
	if h, ok := s.(interface {
		MtNoRequest(context.Context, *nrpc.NoRequest) (*SimpleStringReply, error)
	}); ok {
		ns.MtNoRequest = h.MtNoRequest
	}
	return ns
}

// UnstableNoRequestServiceService is the service API for NoRequestService service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableNoRequestServiceService interface {
	MtNoRequest(context.Context, *nrpc.NoRequest) (*SimpleStringReply, error)
}
