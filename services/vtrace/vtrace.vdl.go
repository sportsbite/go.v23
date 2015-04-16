// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: vtrace.vdl

// Package vtrace defines an interface to access v.io/v23/vtrace traces, to help
// analyze and debug distributed systems.
package vtrace

import (
	// VDL system imports
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/security/access"
	"v.io/v23/uniqueid"
	"v.io/v23/vtrace"
)

// StoreClientMethods is the client interface
// containing Store methods.
type StoreClientMethods interface {
	// Trace returns the trace that matches the given Id.
	// Will return a NoExists error if no matching trace was found.
	Trace(*context.T, uniqueid.Id, ...rpc.CallOpt) (vtrace.TraceRecord, error)
	// AllTraces returns TraceRecords for all traces the server currently
	// knows about.
	AllTraces(*context.T, ...rpc.CallOpt) (StoreAllTracesClientCall, error)
}

// StoreClientStub adds universal methods to StoreClientMethods.
type StoreClientStub interface {
	StoreClientMethods
	rpc.UniversalServiceMethods
}

// StoreClient returns a client stub for Store.
func StoreClient(name string) StoreClientStub {
	return implStoreClientStub{name}
}

type implStoreClientStub struct {
	name string
}

func (c implStoreClientStub) Trace(ctx *context.T, i0 uniqueid.Id, opts ...rpc.CallOpt) (o0 vtrace.TraceRecord, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Trace", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implStoreClientStub) AllTraces(ctx *context.T, opts ...rpc.CallOpt) (ocall StoreAllTracesClientCall, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "AllTraces", nil, opts...); err != nil {
		return
	}
	ocall = &implStoreAllTracesClientCall{ClientCall: call}
	return
}

// StoreAllTracesClientStream is the client stream for Store.AllTraces.
type StoreAllTracesClientStream interface {
	// RecvStream returns the receiver side of the Store.AllTraces client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() vtrace.TraceRecord
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// StoreAllTracesClientCall represents the call returned from Store.AllTraces.
type StoreAllTracesClientCall interface {
	StoreAllTracesClientStream
	// Finish blocks until the server is done, and returns the positional return
	// values for call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() error
}

type implStoreAllTracesClientCall struct {
	rpc.ClientCall
	valRecv vtrace.TraceRecord
	errRecv error
}

func (c *implStoreAllTracesClientCall) RecvStream() interface {
	Advance() bool
	Value() vtrace.TraceRecord
	Err() error
} {
	return implStoreAllTracesClientCallRecv{c}
}

type implStoreAllTracesClientCallRecv struct {
	c *implStoreAllTracesClientCall
}

func (c implStoreAllTracesClientCallRecv) Advance() bool {
	c.c.valRecv = vtrace.TraceRecord{}
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implStoreAllTracesClientCallRecv) Value() vtrace.TraceRecord {
	return c.c.valRecv
}
func (c implStoreAllTracesClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implStoreAllTracesClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
	return
}

// StoreServerMethods is the interface a server writer
// implements for Store.
type StoreServerMethods interface {
	// Trace returns the trace that matches the given Id.
	// Will return a NoExists error if no matching trace was found.
	Trace(*context.T, rpc.ServerCall, uniqueid.Id) (vtrace.TraceRecord, error)
	// AllTraces returns TraceRecords for all traces the server currently
	// knows about.
	AllTraces(*context.T, StoreAllTracesServerCall) error
}

// StoreServerStubMethods is the server interface containing
// Store methods, as expected by rpc.Server.
// The only difference between this interface and StoreServerMethods
// is the streaming methods.
type StoreServerStubMethods interface {
	// Trace returns the trace that matches the given Id.
	// Will return a NoExists error if no matching trace was found.
	Trace(*context.T, rpc.ServerCall, uniqueid.Id) (vtrace.TraceRecord, error)
	// AllTraces returns TraceRecords for all traces the server currently
	// knows about.
	AllTraces(*context.T, *StoreAllTracesServerCallStub) error
}

// StoreServerStub adds universal methods to StoreServerStubMethods.
type StoreServerStub interface {
	StoreServerStubMethods
	// Describe the Store interfaces.
	Describe__() []rpc.InterfaceDesc
}

// StoreServer returns a server stub for Store.
// It converts an implementation of StoreServerMethods into
// an object that may be used by rpc.Server.
func StoreServer(impl StoreServerMethods) StoreServerStub {
	stub := implStoreServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implStoreServerStub struct {
	impl StoreServerMethods
	gs   *rpc.GlobState
}

func (s implStoreServerStub) Trace(ctx *context.T, call rpc.ServerCall, i0 uniqueid.Id) (vtrace.TraceRecord, error) {
	return s.impl.Trace(ctx, call, i0)
}

func (s implStoreServerStub) AllTraces(ctx *context.T, call *StoreAllTracesServerCallStub) error {
	return s.impl.AllTraces(ctx, call)
}

func (s implStoreServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implStoreServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{StoreDesc}
}

// StoreDesc describes the Store interface.
var StoreDesc rpc.InterfaceDesc = descStore

// descStore hides the desc to keep godoc clean.
var descStore = rpc.InterfaceDesc{
	Name:    "Store",
	PkgPath: "v.io/v23/services/vtrace",
	Methods: []rpc.MethodDesc{
		{
			Name: "Trace",
			Doc:  "// Trace returns the trace that matches the given Id.\n// Will return a NoExists error if no matching trace was found.",
			InArgs: []rpc.ArgDesc{
				{"", ``}, // uniqueid.Id
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // vtrace.TraceRecord
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Debug"))},
		},
		{
			Name: "AllTraces",
			Doc:  "// AllTraces returns TraceRecords for all traces the server currently\n// knows about.",
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Debug"))},
		},
	},
}

// StoreAllTracesServerStream is the server stream for Store.AllTraces.
type StoreAllTracesServerStream interface {
	// SendStream returns the send side of the Store.AllTraces server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item vtrace.TraceRecord) error
	}
}

// StoreAllTracesServerCall represents the context passed to Store.AllTraces.
type StoreAllTracesServerCall interface {
	rpc.ServerCall
	StoreAllTracesServerStream
}

// StoreAllTracesServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements StoreAllTracesServerCall.
type StoreAllTracesServerCallStub struct {
	rpc.StreamServerCall
}

// Init initializes StoreAllTracesServerCallStub from rpc.StreamServerCall.
func (s *StoreAllTracesServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// SendStream returns the send side of the Store.AllTraces server stream.
func (s *StoreAllTracesServerCallStub) SendStream() interface {
	Send(item vtrace.TraceRecord) error
} {
	return implStoreAllTracesServerCallSend{s}
}

type implStoreAllTracesServerCallSend struct {
	s *StoreAllTracesServerCallStub
}

func (s implStoreAllTracesServerCallSend) Send(item vtrace.TraceRecord) error {
	return s.s.Send(item)
}
