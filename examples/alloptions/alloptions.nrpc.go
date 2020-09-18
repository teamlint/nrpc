// Code generated by protoc-gen-nrpc. DO NOT EDIT. 

// Package main generated by protoc-gen-nrpc. sources: alloptions.proto
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/teamlint/nrpc"
)

// SvcCustomSubjectService is the interface that providers of the service
// SvcCustomSubject should implement.
type SvcCustomSubjectService interface {
	MtSimpleReply(ctx context.Context, req *StringArg) (*SimpleStringReply,error)
	MtVoidReply(ctx context.Context, req *StringArg) (error)
	MtStreamedReply(ctx context.Context, req *StringArg, push func(*SimpleStringReply)) (error)
	MtVoidReqStreamedReply(ctx context.Context, push func(*SimpleStringReply)) (error)
	MtRequestNoReply(ctx context.Context, req *StringArg)
}

// SvcCustomSubjectHandler provides a NATS subscription handler that can serve a
// subscription using a given SvcCustomSubjectService implementation.
type SvcCustomSubjectHandler struct {
	ctx       context.Context
	workers   *nrpc.WorkerPool
	nc        nrpc.NatsConn
	service   SvcCustomSubjectService
	encodings []string
}
func NewSvcCustomSubjectHandler(ctx context.Context, nc nrpc.NatsConn, s SvcCustomSubjectService) *SvcCustomSubjectHandler {
	return &SvcCustomSubjectHandler{
		ctx:       ctx,
		nc:        nc,
		service:   s,
		encodings: []string{"protobuf"},
	}
}
func NewSvcCustomSubjectConcurrentHandler(workers *nrpc.WorkerPool, nc nrpc.NatsConn, s SvcCustomSubjectService) *SvcCustomSubjectHandler {
	return &SvcCustomSubjectHandler{
		workers: workers,
		nc:      nc,
		service:  s,
	}
}

// SetEncodings sets the output encodings when using a '*Publish' function
func (h *SvcCustomSubjectHandler) SetEncodings(encodings []string) {
	h.encodings = encodings
}
func (h *SvcCustomSubjectHandler) Subject() string {
	return "root.*.custom_subject.>"
}

func (h *SvcCustomSubjectHandler) MtNoRequestPublish(pkginstance string, msg *SimpleStringReply) error {
	for _, encoding := range h.encodings {
		rawMsg, err := nrpc.Marshal(encoding, msg)
		if err != nil {
			log.Printf("SvcCustomSubjectHandler.MtNoRequestPublish: error marshaling the message: %s", err)
			return err
		}
		subject := "root." + pkginstance + "."+ "custom_subject."+ "mtnorequest"
		if encoding != "protobuf" {
			subject += "." + encoding
		}
		if err := h.nc.Publish(subject, rawMsg); err != nil {
			return err
		}
	}
	return nil
}

func (h *SvcCustomSubjectHandler) MsgHandler(msg *nats.Msg) {
	var ctx context.Context
	if h.workers != nil {
		ctx = h.workers.Context
	} else {
		ctx = h.ctx
	}
	request := nrpc.NewRequest(ctx, h.nc, msg.Subject, msg.Reply)
	// extract method name & encoding from subject
	pkgParams, _, name, tail, err := nrpc.ParseSubject(
		"root", 1, "custom_subject", 0, msg.Subject)
	if err != nil {
		log.Printf("SvcCustomSubjectHanlder: SvcCustomSubject subject parsing failed: %v", err)
		return
	}

	request.MethodName = name
	request.SubjectTail = tail
	request.SetPackageParam("instance", pkgParams[0])

	// call handler and form response
	var immediateError *nrpc.Error
	switch name {
		
	case "mt_simple_reply":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("MtSimpleReplyHanlder: MtSimpleReply subject parsing failed: %v", err)
			break
		}
		
		req := new(StringArg)
		
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, req); err != nil {
			log.Printf("MtSimpleReplyHandler: MtSimpleReply request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.service.MtSimpleReply(ctx, req)
				if err != nil {
					return nil, err
				}
				return innerResp, nil
			}
		}
		
	case "mtvoidreply":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("MtVoidReplyHanlder: MtVoidReply subject parsing failed: %v", err)
			break
		}
		
		req := new(StringArg)
		
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, req); err != nil {
			log.Printf("MtVoidReplyHandler: MtVoidReply request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp := new(nrpc.Void)
				err := h.service.MtVoidReply(ctx, req)
				if err != nil {
					return nil, err
				}
				return innerResp, nil
			}
		}
		
	case "mtnorequest":
		// MtNoRequest is a no-request method. Ignore it.
		return
		
	case "mtstreamedreply":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("MtStreamedReplyHanlder: MtStreamedReply subject parsing failed: %v", err)
			break
		}
		
		req := new(StringArg)
		
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, req); err != nil {
			log.Printf("MtStreamedReplyHandler: MtStreamedReply request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.EnableStreamedReply()
			request.Handler = func(ctx context.Context)(proto.Message, error){
				err := h.service.MtStreamedReply(ctx, req, func(rep *SimpleStringReply){
					request.SendStreamReply(rep)
				})
				return nil, err
			}
		}
		
	case "mtvoidreqstreamedreply":
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("MtVoidReqStreamedReplyHanlder: MtVoidReqStreamedReply subject parsing failed: %v", err)
			break
		}
		
		req := new(nrpc.Void)
		
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, req); err != nil {
			log.Printf("MtVoidReqStreamedReplyHandler: MtVoidReqStreamedReply request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.EnableStreamedReply()
			request.Handler = func(ctx context.Context)(proto.Message, error){
				err := h.service.MtVoidReqStreamedReply(ctx, func(rep *SimpleStringReply){
					request.SendStreamReply(rep)
				})
				return nil, err
			}
		}
		
	case "mtrequestnoreply":
		request.NoReply = true
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("MtRequestNoReplyHanlder: MtRequestNoReply subject parsing failed: %v", err)
			break
		}
		
		req := new(StringArg)
		
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, req); err != nil {
			log.Printf("MtRequestNoReplyHandler: MtRequestNoReply request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp := new(nrpc.NoReply)
				h.service.MtRequestNoReply(ctx, req)
				if err != nil {
					return nil, err
				}
				return innerResp, nil
			}
		}
	default:
		log.Printf("SvcCustomSubjectHandler: unknown name %q", name)
		immediateError = &nrpc.Error{
			Type: nrpc.Error_CLIENT,
			Message: "unknown name: " + name,
		}
	}

	if immediateError == nil {
		if h.workers != nil {
			// Try queuing the request
			if err := h.workers.QueueRequest(request); err != nil {
				log.Printf("nrpc: Error queuing the request: %s", err)
			}
		} else {
			// Run the handler synchronously
			request.RunAndReply()
		}
	}

	if immediateError != nil {
		if err := request.SendReply(nil, immediateError); err != nil {
			log.Printf("SvcCustomSubjectHandler: SvcCustomSubject handler failed to publish the response: %s", err)
		}
	}
}

type SvcCustomSubjectClient struct {
	nc                nrpc.NatsConn
	PkgSubject       string
	PkgParaminstance string
	Subject          string
	Encoding          string
	Timeout           time.Duration
}

func NewSvcCustomSubjectClient(nc nrpc.NatsConn, pkgParaminstance string) *SvcCustomSubjectClient {
	return &SvcCustomSubjectClient{
		nc:      nc,
		PkgSubject: "root",
		PkgParaminstance: pkgParaminstance,
		Subject: "custom_subject",
		Encoding: "protobuf",
		Timeout: 5 * time.Second,
	}
}

func (c *SvcCustomSubjectClient) SetEncoding(encoding string) {
	c.Encoding = encoding
}

func (c *SvcCustomSubjectClient) SetTimeout(t time.Duration) {
	c.Timeout = t
}

func (c *SvcCustomSubjectClient) MtSimpleReply(req *StringArg) (*SimpleStringReply, error) {
	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + "mt_simple_reply"
	resp := new(SimpleStringReply)
	err := nrpc.Call(req, resp, c.nc, subject, c.Encoding, c.Timeout)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *SvcCustomSubjectClient) MtVoidReply(req *StringArg) (error) {
	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + "mtvoidreply"
	resp := new(nrpc.Void)
	err := nrpc.Call(req, resp, c.nc, subject, c.Encoding, c.Timeout)
	if err != nil {
		return err
	}
	return nil
}

func (c *SvcCustomSubjectClient) MtNoRequestSubject(
	
) string {
	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + "mtnorequest"
	if c.Encoding != "protobuf" {
		subject += "." + c.Encoding
	}
	return subject
}

type SvcCustomSubjectMtNoRequestSubscription struct {
	*nats.Subscription
	encoding string
}

func (s *SvcCustomSubjectMtNoRequestSubscription) Next(timeout time.Duration) (*SimpleStringReply, error) {
	reply := new(SimpleStringReply)
	msg, err := s.Subscription.NextMsg(timeout)
	if err != nil {
		return nil, err
	}
	err = nrpc.Unmarshal(s.encoding, msg.Data, reply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (c *SvcCustomSubjectClient) MtNoRequestSubscribeSync(
	
) (*SvcCustomSubjectMtNoRequestSubscription, error) {
	subject := c.MtNoRequestSubject(
		
	)
	natsSub, err := c.nc.SubscribeSync(subject)
	if err != nil {
		return nil, err
	}
	return &SvcCustomSubjectMtNoRequestSubscription{natsSub, c.Encoding}, nil
}

func (c *SvcCustomSubjectClient) MtNoRequestSubscribe(handler func (*SimpleStringReply)) (*nats.Subscription, error) {
	subject := c.MtNoRequestSubject(
		
	)
	sub, err := c.nc.Subscribe(subject, func(msg *nats.Msg){
		reply := new(SimpleStringReply)
		err := nrpc.Unmarshal(c.Encoding, msg.Data, reply)
		if err != nil {
			log.Printf("SvcCustomSubjectClient.MtNoRequestSubscribe: Error decoding, %s", err)
			return 
		}
		handler(reply)
	})
	return sub, err
}

func (c *SvcCustomSubjectClient) MtNoRequestSubscribeChan(
	
) (<-chan *SimpleStringReply, *nats.Subscription, error) {
	ch := make(chan *SimpleStringReply)
	sub, err := c.MtNoRequestSubscribe(func (reply *SimpleStringReply) {
		ch <- reply
	})
	return ch, sub, err
}


func (c *SvcCustomSubjectClient) MtStreamedReply(
	ctx context.Context,
	req *StringArg,
	cb func (context.Context, *SimpleStringReply),
) error {
	
	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + "mtstreamedreply"

	sub, err := nrpc.StreamCall(ctx, c.nc, subject, req, c.Encoding, c.Timeout)
	if err != nil {
		return err
	}

	reply := new(SimpleStringReply)
	for {
		err = sub.Next(reply)
		if err != nil {
			break
		}
		cb(ctx, reply)
	}
	if err == nrpc.ErrEOS {
		err = nil
	}
	return err
}


func (c *SvcCustomSubjectClient) MtVoidReqStreamedReply(
	ctx context.Context,
	cb func (context.Context, *SimpleStringReply),
) error {
	
	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + "mtvoidreqstreamedreply"

	sub, err := nrpc.StreamCall(ctx, c.nc, subject, &nrpc.Void{}, c.Encoding, c.Timeout)
	if err != nil {
		return err
	}

	reply := new(SimpleStringReply)
	for {
		err = sub.Next(reply)
		if err != nil {
			break
		}
		cb(ctx, reply)
	}
	if err == nrpc.ErrEOS {
		err = nil
	}
	return err
}


func (c *SvcCustomSubjectClient) MtRequestNoReply(req *StringArg) (error) {
	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + "mtrequestnoreply"
	resp := new(nrpc.NoReply)
	err := nrpc.Call(req, resp, c.nc, subject, c.Encoding, c.Timeout)
	if err != nil {
		return err
	}
	return nil
}

// SvcSubjectParamsService is the interface that providers of the service
// SvcSubjectParams should implement.
type SvcSubjectParamsService interface {
	MtWithSubjectParams(ctx context.Context, mp1 string, mp2 string) (*SimpleStringReply,error)
	MtStreamedReplyWithSubjectParams(ctx context.Context, mp1 string, mp2 string, push func(*SimpleStringReply)) (error)
	MtNoReply(ctx context.Context)
}

// SvcSubjectParamsHandler provides a NATS subscription handler that can serve a
// subscription using a given SvcSubjectParamsService implementation.
type SvcSubjectParamsHandler struct {
	ctx       context.Context
	workers   *nrpc.WorkerPool
	nc        nrpc.NatsConn
	service   SvcSubjectParamsService
	encodings []string
}
func NewSvcSubjectParamsHandler(ctx context.Context, nc nrpc.NatsConn, s SvcSubjectParamsService) *SvcSubjectParamsHandler {
	return &SvcSubjectParamsHandler{
		ctx:       ctx,
		nc:        nc,
		service:   s,
		encodings: []string{"protobuf"},
	}
}
func NewSvcSubjectParamsConcurrentHandler(workers *nrpc.WorkerPool, nc nrpc.NatsConn, s SvcSubjectParamsService) *SvcSubjectParamsHandler {
	return &SvcSubjectParamsHandler{
		workers: workers,
		nc:      nc,
		service:  s,
	}
}

// SetEncodings sets the output encodings when using a '*Publish' function
func (h *SvcSubjectParamsHandler) SetEncodings(encodings []string) {
	h.encodings = encodings
}
func (h *SvcSubjectParamsHandler) Subject() string {
	return "root.*.svcsubjectparams.*.>"
}

func (h *SvcSubjectParamsHandler) MtNoRequestWParamsPublish(pkginstance string, svcclientid string, mtmp1 string, msg *SimpleStringReply) error {
	for _, encoding := range h.encodings {
		rawMsg, err := nrpc.Marshal(encoding, msg)
		if err != nil {
			log.Printf("SvcSubjectParamsHandler.MtNoRequestWParamsPublish: error marshaling the message: %s", err)
			return err
		}
		subject := "root." + pkginstance + "."+ "svcsubjectparams." + svcclientid + "."+ "mtnorequestwparams" + "." + mtmp1
		if encoding != "protobuf" {
			subject += "." + encoding
		}
		if err := h.nc.Publish(subject, rawMsg); err != nil {
			return err
		}
	}
	return nil
}

func (h *SvcSubjectParamsHandler) MsgHandler(msg *nats.Msg) {
	var ctx context.Context
	if h.workers != nil {
		ctx = h.workers.Context
	} else {
		ctx = h.ctx
	}
	request := nrpc.NewRequest(ctx, h.nc, msg.Subject, msg.Reply)
	// extract method name & encoding from subject
	pkgParams, svcParams, name, tail, err := nrpc.ParseSubject(
		"root", 1, "svcsubjectparams", 1, msg.Subject)
	if err != nil {
		log.Printf("SvcSubjectParamsHanlder: SvcSubjectParams subject parsing failed: %v", err)
		return
	}

	request.MethodName = name
	request.SubjectTail = tail
	request.SetPackageParam("instance", pkgParams[0])
	request.SetServiceParam("clientid", svcParams[0])

	// call handler and form response
	var immediateError *nrpc.Error
	switch name {
		
	case "mtwithsubjectparams":
		var mtParams []string
		mtParams, request.Encoding, err = nrpc.ParseSubjectTail(2, request.SubjectTail)
		if err != nil {
			log.Printf("MtWithSubjectParamsHanlder: MtWithSubjectParams subject parsing failed: %v", err)
			break
		}
		
		req := new(nrpc.Void)
		
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, req); err != nil {
			log.Printf("MtWithSubjectParamsHandler: MtWithSubjectParams request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp, err := h.service.MtWithSubjectParams(ctx, mtParams[0], mtParams[1])
				if err != nil {
					return nil, err
				}
				return innerResp, nil
			}
		}
		
	case "mtstreamedreplywithsubjectparams":
		var mtParams []string
		mtParams, request.Encoding, err = nrpc.ParseSubjectTail(2, request.SubjectTail)
		if err != nil {
			log.Printf("MtStreamedReplyWithSubjectParamsHanlder: MtStreamedReplyWithSubjectParams subject parsing failed: %v", err)
			break
		}
		
		req := new(nrpc.Void)
		
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, req); err != nil {
			log.Printf("MtStreamedReplyWithSubjectParamsHandler: MtStreamedReplyWithSubjectParams request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.EnableStreamedReply()
			request.Handler = func(ctx context.Context)(proto.Message, error){
				err := h.service.MtStreamedReplyWithSubjectParams(ctx, mtParams[0], mtParams[1], func(rep *SimpleStringReply){
					request.SendStreamReply(rep)
				})
				return nil, err
			}
		}
		
	case "mtnoreply":
		request.NoReply = true
		_, request.Encoding, err = nrpc.ParseSubjectTail(0, request.SubjectTail)
		if err != nil {
			log.Printf("MtNoReplyHanlder: MtNoReply subject parsing failed: %v", err)
			break
		}
		
		req := new(nrpc.Void)
		
		if err := nrpc.Unmarshal(request.Encoding, msg.Data, req); err != nil {
			log.Printf("MtNoReplyHandler: MtNoReply request unmarshal failed: %v", err)
			immediateError = &nrpc.Error{
				Type: nrpc.Error_CLIENT,
				Message: "bad request received: " + err.Error(),
			}
		} else {
			request.Handler = func(ctx context.Context)(proto.Message, error){
				innerResp := new(nrpc.NoReply)
				h.service.MtNoReply(ctx)
				if err != nil {
					return nil, err
				}
				return innerResp, nil
			}
		}
		
	case "mtnorequestwparams":
		// MtNoRequestWParams is a no-request method. Ignore it.
		return
	default:
		log.Printf("SvcSubjectParamsHandler: unknown name %q", name)
		immediateError = &nrpc.Error{
			Type: nrpc.Error_CLIENT,
			Message: "unknown name: " + name,
		}
	}

	if immediateError == nil {
		if h.workers != nil {
			// Try queuing the request
			if err := h.workers.QueueRequest(request); err != nil {
				log.Printf("nrpc: Error queuing the request: %s", err)
			}
		} else {
			// Run the handler synchronously
			request.RunAndReply()
		}
	}

	if immediateError != nil {
		if err := request.SendReply(nil, immediateError); err != nil {
			log.Printf("SvcSubjectParamsHandler: SvcSubjectParams handler failed to publish the response: %s", err)
		}
	}
}

type SvcSubjectParamsClient struct {
	nc                nrpc.NatsConn
	PkgSubject       string
	PkgParaminstance string
	Subject          string
	SvcParamclientid string
	Encoding          string
	Timeout           time.Duration
}

func NewSvcSubjectParamsClient(nc nrpc.NatsConn, pkgParaminstance string, svcParamclientid string) *SvcSubjectParamsClient {
	return &SvcSubjectParamsClient{
		nc:      nc,
		PkgSubject: "root",
		PkgParaminstance: pkgParaminstance,
		Subject: "svcsubjectparams",
		SvcParamclientid: svcParamclientid,
		Encoding: "protobuf",
		Timeout: 5 * time.Second,
	}
}

func (c *SvcSubjectParamsClient) SetEncoding(encoding string) {
	c.Encoding = encoding
}

func (c *SvcSubjectParamsClient) SetTimeout(t time.Duration) {
	c.Timeout = t
}

func (c *SvcSubjectParamsClient) MtWithSubjectParams(mp1 string, mp2 string, ) (*SimpleStringReply, error) {
	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + c.SvcParamclientid + "." + "mtwithsubjectparams" + "." + mp1 + "." + mp2
	req := new(nrpc.Void)
	resp := new(SimpleStringReply)
	err := nrpc.Call(req, resp, c.nc, subject, c.Encoding, c.Timeout)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *SvcSubjectParamsClient) MtStreamedReplyWithSubjectParams(
	ctx context.Context,mp1 string,mp2 string,
	cb func (context.Context, *SimpleStringReply),
) error {
	
	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + c.SvcParamclientid + "." + "mtstreamedreplywithsubjectparams" + "." + mp1 + "." + mp2

	sub, err := nrpc.StreamCall(ctx, c.nc, subject, &nrpc.Void{}, c.Encoding, c.Timeout)
	if err != nil {
		return err
	}

	reply := new(SimpleStringReply)
	for {
		err = sub.Next(reply)
		if err != nil {
			break
		}
		cb(ctx, reply)
	}
	if err == nrpc.ErrEOS {
		err = nil
	}
	return err
}


func (c *SvcSubjectParamsClient) MtNoReply() (error) {
	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + c.SvcParamclientid + "." + "mtnoreply"
	req := new(nrpc.Void)
	resp := new(nrpc.NoReply)
	err := nrpc.Call(req, resp, c.nc, subject, c.Encoding, c.Timeout)
	if err != nil {
		return err
	}
	return nil
}

func (c *SvcSubjectParamsClient) MtNoRequestWParamsSubject(
	mtmp1 string,
) string {
	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + c.SvcParamclientid + "." + "mtnorequestwparams" + "." + mtmp1
	if c.Encoding != "protobuf" {
		subject += "." + c.Encoding
	}
	return subject
}

type SvcSubjectParamsMtNoRequestWParamsSubscription struct {
	*nats.Subscription
	encoding string
}

func (s *SvcSubjectParamsMtNoRequestWParamsSubscription) Next(timeout time.Duration) (*SimpleStringReply, error) {
	reply := new(SimpleStringReply)
	msg, err := s.Subscription.NextMsg(timeout)
	if err != nil {
		return nil, err
	}
	err = nrpc.Unmarshal(s.encoding, msg.Data, reply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (c *SvcSubjectParamsClient) MtNoRequestWParamsSubscribeSync(
	mtmp1 string,
) (*SvcSubjectParamsMtNoRequestWParamsSubscription, error) {
	subject := c.MtNoRequestWParamsSubject(
		mtmp1,
	)
	natsSub, err := c.nc.SubscribeSync(subject)
	if err != nil {
		return nil, err
	}
	return &SvcSubjectParamsMtNoRequestWParamsSubscription{natsSub, c.Encoding}, nil
}

func (c *SvcSubjectParamsClient) MtNoRequestWParamsSubscribe(mtmp1 string,handler func (*SimpleStringReply)) (*nats.Subscription, error) {
	subject := c.MtNoRequestWParamsSubject(
		mtmp1,
	)
	sub, err := c.nc.Subscribe(subject, func(msg *nats.Msg){
		reply := new(SimpleStringReply)
		err := nrpc.Unmarshal(c.Encoding, msg.Data, reply)
		if err != nil {
			log.Printf("SvcSubjectParamsClient.MtNoRequestWParamsSubscribe: Error decoding, %s", err)
			return 
		}
		handler(reply)
	})
	return sub, err
}

func (c *SvcSubjectParamsClient) MtNoRequestWParamsSubscribeChan(
	mtmp1 string,
) (<-chan *SimpleStringReply, *nats.Subscription, error) {
	ch := make(chan *SimpleStringReply)
	sub, err := c.MtNoRequestWParamsSubscribe(mtmp1, func (reply *SimpleStringReply) {
		ch <- reply
	})
	return ch, sub, err
}


// NoRequestServiceService is the interface that providers of the service
// NoRequestService should implement.
type NoRequestServiceService interface {
}

// NoRequestServiceHandler provides a NATS subscription handler that can serve a
// subscription using a given NoRequestServiceService implementation.
type NoRequestServiceHandler struct {
	ctx       context.Context
	workers   *nrpc.WorkerPool
	nc        nrpc.NatsConn
	service   NoRequestServiceService
	encodings []string
}
func NewNoRequestServiceHandler(ctx context.Context, nc nrpc.NatsConn, s NoRequestServiceService) *NoRequestServiceHandler {
	return &NoRequestServiceHandler{
		ctx:       ctx,
		nc:        nc,
		service:   s,
		encodings: []string{"protobuf"},
	}
}
func NewNoRequestServiceConcurrentHandler(workers *nrpc.WorkerPool, nc nrpc.NatsConn, s NoRequestServiceService) *NoRequestServiceHandler {
	return &NoRequestServiceHandler{
		workers: workers,
		nc:      nc,
		service:  s,
	}
}

// SetEncodings sets the output encodings when using a '*Publish' function
func (h *NoRequestServiceHandler) SetEncodings(encodings []string) {
	h.encodings = encodings
}
func (h *NoRequestServiceHandler) Subject() string {
	return "root.*.norequestservice.>"
}

func (h *NoRequestServiceHandler) MtNoRequestPublish(pkginstance string, msg *SimpleStringReply) error {
	for _, encoding := range h.encodings {
		rawMsg, err := nrpc.Marshal(encoding, msg)
		if err != nil {
			log.Printf("NoRequestServiceHandler.MtNoRequestPublish: error marshaling the message: %s", err)
			return err
		}
		subject := "root." + pkginstance + "."+ "norequestservice."+ "mtnorequest"
		if encoding != "protobuf" {
			subject += "." + encoding
		}
		if err := h.nc.Publish(subject, rawMsg); err != nil {
			return err
		}
	}
	return nil
}

type NoRequestServiceClient struct {
	nc                nrpc.NatsConn
	PkgSubject       string
	PkgParaminstance string
	Subject          string
	Encoding          string
	Timeout           time.Duration
}

func NewNoRequestServiceClient(nc nrpc.NatsConn, pkgParaminstance string) *NoRequestServiceClient {
	return &NoRequestServiceClient{
		nc:      nc,
		PkgSubject: "root",
		PkgParaminstance: pkgParaminstance,
		Subject: "norequestservice",
		Encoding: "protobuf",
		Timeout: 5 * time.Second,
	}
}

func (c *NoRequestServiceClient) SetEncoding(encoding string) {
	c.Encoding = encoding
}

func (c *NoRequestServiceClient) SetTimeout(t time.Duration) {
	c.Timeout = t
}

func (c *NoRequestServiceClient) MtNoRequestSubject(
	
) string {
	subject := c.PkgSubject + "." + c.PkgParaminstance + "." + c.Subject + "." + "mtnorequest"
	if c.Encoding != "protobuf" {
		subject += "." + c.Encoding
	}
	return subject
}

type NoRequestServiceMtNoRequestSubscription struct {
	*nats.Subscription
	encoding string
}

func (s *NoRequestServiceMtNoRequestSubscription) Next(timeout time.Duration) (*SimpleStringReply, error) {
	reply := new(SimpleStringReply)
	msg, err := s.Subscription.NextMsg(timeout)
	if err != nil {
		return nil, err
	}
	err = nrpc.Unmarshal(s.encoding, msg.Data, reply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (c *NoRequestServiceClient) MtNoRequestSubscribeSync(
	
) (*NoRequestServiceMtNoRequestSubscription, error) {
	subject := c.MtNoRequestSubject(
		
	)
	natsSub, err := c.nc.SubscribeSync(subject)
	if err != nil {
		return nil, err
	}
	return &NoRequestServiceMtNoRequestSubscription{natsSub, c.Encoding}, nil
}

func (c *NoRequestServiceClient) MtNoRequestSubscribe(handler func (*SimpleStringReply)) (*nats.Subscription, error) {
	subject := c.MtNoRequestSubject(
		
	)
	sub, err := c.nc.Subscribe(subject, func(msg *nats.Msg){
		reply := new(SimpleStringReply)
		err := nrpc.Unmarshal(c.Encoding, msg.Data, reply)
		if err != nil {
			log.Printf("NoRequestServiceClient.MtNoRequestSubscribe: Error decoding, %s", err)
			return 
		}
		handler(reply)
	})
	return sub, err
}

func (c *NoRequestServiceClient) MtNoRequestSubscribeChan(
	
) (<-chan *SimpleStringReply, *nats.Subscription, error) {
	ch := make(chan *SimpleStringReply)
	sub, err := c.MtNoRequestSubscribe(func (reply *SimpleStringReply) {
		ch <- reply
	})
	return ch, sub, err
}


type Client struct {
	nc               nrpc.NatsConn
	defaultEncoding  string
	defaultTimeout   time.Duration
	pkgSubject       string
	pkgParaminstance string
	SvcCustomSubject *SvcCustomSubjectClient
	SvcSubjectParams *SvcSubjectParamsClient
	NoRequestService *NoRequestServiceClient
}

func NewClient(nc nrpc.NatsConn, pkgParaminstance string) *Client {
	c := Client{
		nc:               nc,
		defaultEncoding:  "protobuf",
		defaultTimeout:   5*time.Second,
		pkgSubject:       "root",
		pkgParaminstance: pkgParaminstance,
	}
	c.SvcCustomSubject = NewSvcCustomSubjectClient(nc, c.pkgParaminstance)
	c.NoRequestService = NewNoRequestServiceClient(nc, c.pkgParaminstance)
	return &c
}

func (c *Client) SetEncoding(encoding string) {
	c.defaultEncoding = encoding
	if c.SvcCustomSubject != nil {
		c.SvcCustomSubject.Encoding = encoding
	}
	if c.SvcSubjectParams != nil {
		c.SvcSubjectParams.Encoding = encoding
	}
	if c.NoRequestService != nil {
		c.NoRequestService.Encoding = encoding
	}
}

func (c *Client) SetTimeout(t time.Duration) {
	c.defaultTimeout = t
	if c.SvcCustomSubject != nil {
		c.SvcCustomSubject.Timeout = t
	}
	if c.SvcSubjectParams != nil {
		c.SvcSubjectParams.Timeout = t
	}
	if c.NoRequestService != nil {
		c.NoRequestService.Timeout = t
	}
}

func (c *Client) SetSvcSubjectParams(
	clientid string,
) {
	c.SvcSubjectParams = NewSvcSubjectParamsClient(
		c.nc,
		c.pkgParaminstance,
		clientid,
	)
	c.SvcSubjectParams.Encoding = c.defaultEncoding
	c.SvcSubjectParams.Timeout = c.defaultTimeout
}

func (c *Client) NewSvcSubjectParams(
	clientid string,
) *SvcSubjectParamsClient {
	client := NewSvcSubjectParamsClient(
		c.nc,
		c.pkgParaminstance,
		clientid,
	)
	client.Encoding = c.defaultEncoding
	client.Timeout = c.defaultTimeout
	return client
}
