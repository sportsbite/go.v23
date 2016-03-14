// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: tidyable

// Package tidyable defines an interface for services that can be
// requested to clean up transient resource use (such as logs or caches.)
package tidyable

import (
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/security/access"
	"v.io/v23/vdl"
)

func __VDLEnsureNativeBuilt() {
}

// TidyableClientMethods is the client interface
// containing Tidyable methods.
//
// Tidyable specifies that a service can be tidied.
type TidyableClientMethods interface {
	// Request the implementing service to perform regularly scheduled cleanup
	//  actions such as shrinking caches or rolling logs immediately.
	TidyNow(*context.T, ...rpc.CallOpt) error
}

// TidyableClientStub adds universal methods to TidyableClientMethods.
type TidyableClientStub interface {
	TidyableClientMethods
	rpc.UniversalServiceMethods
}

// TidyableClient returns a client stub for Tidyable.
func TidyableClient(name string) TidyableClientStub {
	return implTidyableClientStub{name}
}

type implTidyableClientStub struct {
	name string
}

func (c implTidyableClientStub) TidyNow(ctx *context.T, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "TidyNow", nil, nil, opts...)
	return
}

// TidyableServerMethods is the interface a server writer
// implements for Tidyable.
//
// Tidyable specifies that a service can be tidied.
type TidyableServerMethods interface {
	// Request the implementing service to perform regularly scheduled cleanup
	//  actions such as shrinking caches or rolling logs immediately.
	TidyNow(*context.T, rpc.ServerCall) error
}

// TidyableServerStubMethods is the server interface containing
// Tidyable methods, as expected by rpc.Server.
// There is no difference between this interface and TidyableServerMethods
// since there are no streaming methods.
type TidyableServerStubMethods TidyableServerMethods

// TidyableServerStub adds universal methods to TidyableServerStubMethods.
type TidyableServerStub interface {
	TidyableServerStubMethods
	// Describe the Tidyable interfaces.
	Describe__() []rpc.InterfaceDesc
}

// TidyableServer returns a server stub for Tidyable.
// It converts an implementation of TidyableServerMethods into
// an object that may be used by rpc.Server.
func TidyableServer(impl TidyableServerMethods) TidyableServerStub {
	stub := implTidyableServerStub{
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

type implTidyableServerStub struct {
	impl TidyableServerMethods
	gs   *rpc.GlobState
}

func (s implTidyableServerStub) TidyNow(ctx *context.T, call rpc.ServerCall) error {
	return s.impl.TidyNow(ctx, call)
}

func (s implTidyableServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implTidyableServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{TidyableDesc}
}

// TidyableDesc describes the Tidyable interface.
var TidyableDesc rpc.InterfaceDesc = descTidyable

// descTidyable hides the desc to keep godoc clean.
var descTidyable = rpc.InterfaceDesc{
	Name:    "Tidyable",
	PkgPath: "v.io/v23/services/tidyable",
	Doc:     "// Tidyable specifies that a service can be tidied.",
	Methods: []rpc.MethodDesc{
		{
			Name: "TidyNow",
			Doc:  "// Request the implementing service to perform regularly scheduled cleanup\n//  actions such as shrinking caches or rolling logs immediately.",
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Admin"))},
		},
	},
}