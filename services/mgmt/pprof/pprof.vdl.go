// This file was auto-generated by the veyron vdl tool.
// Source: pprof.vdl

// Package pprof is used to access runtime profiling data in the format expected
// by the pprof visualization tool. For more information about pprof, see
// http://code.google.com/p/google-perftools/.
package pprof

import (
	// The non-user imports are prefixed with "__" to prevent collisions.
	__io "io"
	__veyron2 "veyron.io/veyron/veyron2"
	__context "veyron.io/veyron/veyron2/context"
	__ipc "veyron.io/veyron/veyron2/ipc"
	__vdlutil "veyron.io/veyron/veyron2/vdl/vdlutil"
	__wiretype "veyron.io/veyron/veyron2/wiretype"
)

// TODO(toddw): Remove this line once the new signature support is done.
// It corrects a bug where __wiretype is unused in VDL pacakges where only
// bootstrap types are used on interfaces.
const _ = __wiretype.TypeIDInvalid

// PProfClientMethods is the client interface
// containing PProf methods.
type PProfClientMethods interface {
	// CmdLine returns the command-line arguments of the server, including
	// the name of the executable.
	CmdLine(__context.T, ...__ipc.CallOpt) ([]string, error)
	// Profiles returns the list of available profiles.
	Profiles(__context.T, ...__ipc.CallOpt) ([]string, error)
	// Profile streams the requested profile. The debug parameter enables
	// additional output. Passing debug=0 includes only the hexadecimal
	// addresses that pprof needs. Passing debug=1 adds comments translating
	// addresses to function names and line numbers, so that a programmer
	// can read the profile without tools.
	Profile(ctx __context.T, name string, debug int32, opts ...__ipc.CallOpt) (PProfProfileCall, error)
	// CPUProfile enables CPU profiling for the requested duration and
	// streams the profile data.
	CPUProfile(ctx __context.T, seconds int32, opts ...__ipc.CallOpt) (PProfCPUProfileCall, error)
	// Symbol looks up the program counters and returns their respective
	// function names.
	Symbol(ctx __context.T, programCounters []uint64, opts ...__ipc.CallOpt) ([]string, error)
}

// PProfClientStub adds universal methods to PProfClientMethods.
type PProfClientStub interface {
	PProfClientMethods
	__ipc.UniversalServiceMethods
}

// PProfClient returns a client stub for PProf.
func PProfClient(name string, opts ...__ipc.BindOpt) PProfClientStub {
	var client __ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(__ipc.Client); ok {
			client = clientOpt
		}
	}
	return implPProfClientStub{name, client}
}

type implPProfClientStub struct {
	name   string
	client __ipc.Client
}

func (c implPProfClientStub) c(ctx __context.T) __ipc.Client {
	if c.client != nil {
		return c.client
	}
	return __veyron2.RuntimeFromContext(ctx).Client()
}

func (c implPProfClientStub) CmdLine(ctx __context.T, opts ...__ipc.CallOpt) (o0 []string, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "CmdLine", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implPProfClientStub) Profiles(ctx __context.T, opts ...__ipc.CallOpt) (o0 []string, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Profiles", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implPProfClientStub) Profile(ctx __context.T, i0 string, i1 int32, opts ...__ipc.CallOpt) (ocall PProfProfileCall, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Profile", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	ocall = &implPProfProfileCall{Call: call}
	return
}

func (c implPProfClientStub) CPUProfile(ctx __context.T, i0 int32, opts ...__ipc.CallOpt) (ocall PProfCPUProfileCall, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "CPUProfile", []interface{}{i0}, opts...); err != nil {
		return
	}
	ocall = &implPProfCPUProfileCall{Call: call}
	return
}

func (c implPProfClientStub) Symbol(ctx __context.T, i0 []uint64, opts ...__ipc.CallOpt) (o0 []string, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Symbol", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implPProfClientStub) Signature(ctx __context.T, opts ...__ipc.CallOpt) (o0 __ipc.ServiceSignature, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// PProfProfileClientStream is the client stream for PProf.Profile.
type PProfProfileClientStream interface {
	// RecvStream returns the receiver side of the PProf.Profile client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() []byte
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// PProfProfileCall represents the call returned from PProf.Profile.
type PProfProfileCall interface {
	PProfProfileClientStream
	// Finish blocks until the server is done, and returns the positional return
	// values for call.
	//
	// Finish returns immediately if Cancel has been called; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless Cancel
	// has been called or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() error
	// Cancel cancels the RPC, notifying the server to stop processing.  It is
	// safe to call Cancel concurrently with any of the other stream methods.
	// Calling Cancel after Finish has returned is a no-op.
	Cancel()
}

type implPProfProfileCall struct {
	__ipc.Call
	valRecv []byte
	errRecv error
}

func (c *implPProfProfileCall) RecvStream() interface {
	Advance() bool
	Value() []byte
	Err() error
} {
	return implPProfProfileCallRecv{c}
}

type implPProfProfileCallRecv struct {
	c *implPProfProfileCall
}

func (c implPProfProfileCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implPProfProfileCallRecv) Value() []byte {
	return c.c.valRecv
}
func (c implPProfProfileCallRecv) Err() error {
	if c.c.errRecv == __io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implPProfProfileCall) Finish() (err error) {
	if ierr := c.Call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

// PProfCPUProfileClientStream is the client stream for PProf.CPUProfile.
type PProfCPUProfileClientStream interface {
	// RecvStream returns the receiver side of the PProf.CPUProfile client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() []byte
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// PProfCPUProfileCall represents the call returned from PProf.CPUProfile.
type PProfCPUProfileCall interface {
	PProfCPUProfileClientStream
	// Finish blocks until the server is done, and returns the positional return
	// values for call.
	//
	// Finish returns immediately if Cancel has been called; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless Cancel
	// has been called or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() error
	// Cancel cancels the RPC, notifying the server to stop processing.  It is
	// safe to call Cancel concurrently with any of the other stream methods.
	// Calling Cancel after Finish has returned is a no-op.
	Cancel()
}

type implPProfCPUProfileCall struct {
	__ipc.Call
	valRecv []byte
	errRecv error
}

func (c *implPProfCPUProfileCall) RecvStream() interface {
	Advance() bool
	Value() []byte
	Err() error
} {
	return implPProfCPUProfileCallRecv{c}
}

type implPProfCPUProfileCallRecv struct {
	c *implPProfCPUProfileCall
}

func (c implPProfCPUProfileCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implPProfCPUProfileCallRecv) Value() []byte {
	return c.c.valRecv
}
func (c implPProfCPUProfileCallRecv) Err() error {
	if c.c.errRecv == __io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implPProfCPUProfileCall) Finish() (err error) {
	if ierr := c.Call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

// PProfServerMethods is the interface a server writer
// implements for PProf.
type PProfServerMethods interface {
	// CmdLine returns the command-line arguments of the server, including
	// the name of the executable.
	CmdLine(__ipc.ServerContext) ([]string, error)
	// Profiles returns the list of available profiles.
	Profiles(__ipc.ServerContext) ([]string, error)
	// Profile streams the requested profile. The debug parameter enables
	// additional output. Passing debug=0 includes only the hexadecimal
	// addresses that pprof needs. Passing debug=1 adds comments translating
	// addresses to function names and line numbers, so that a programmer
	// can read the profile without tools.
	Profile(ctx PProfProfileContext, name string, debug int32) error
	// CPUProfile enables CPU profiling for the requested duration and
	// streams the profile data.
	CPUProfile(ctx PProfCPUProfileContext, seconds int32) error
	// Symbol looks up the program counters and returns their respective
	// function names.
	Symbol(ctx __ipc.ServerContext, programCounters []uint64) ([]string, error)
}

// PProfServerStubMethods is the server interface containing
// PProf methods, as expected by ipc.Server.
// The only difference between this interface and PProfServerMethods
// is the streaming methods.
type PProfServerStubMethods interface {
	// CmdLine returns the command-line arguments of the server, including
	// the name of the executable.
	CmdLine(__ipc.ServerContext) ([]string, error)
	// Profiles returns the list of available profiles.
	Profiles(__ipc.ServerContext) ([]string, error)
	// Profile streams the requested profile. The debug parameter enables
	// additional output. Passing debug=0 includes only the hexadecimal
	// addresses that pprof needs. Passing debug=1 adds comments translating
	// addresses to function names and line numbers, so that a programmer
	// can read the profile without tools.
	Profile(ctx *PProfProfileContextStub, name string, debug int32) error
	// CPUProfile enables CPU profiling for the requested duration and
	// streams the profile data.
	CPUProfile(ctx *PProfCPUProfileContextStub, seconds int32) error
	// Symbol looks up the program counters and returns their respective
	// function names.
	Symbol(ctx __ipc.ServerContext, programCounters []uint64) ([]string, error)
}

// PProfServerStub adds universal methods to PProfServerStubMethods.
type PProfServerStub interface {
	PProfServerStubMethods
	// Describe the PProf interfaces.
	Describe__() []__ipc.InterfaceDesc
	// Signature will be replaced with Describe__.
	Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error)
}

// PProfServer returns a server stub for PProf.
// It converts an implementation of PProfServerMethods into
// an object that may be used by ipc.Server.
func PProfServer(impl PProfServerMethods) PProfServerStub {
	stub := implPProfServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := __ipc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := __ipc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implPProfServerStub struct {
	impl PProfServerMethods
	gs   *__ipc.GlobState
}

func (s implPProfServerStub) CmdLine(ctx __ipc.ServerContext) ([]string, error) {
	return s.impl.CmdLine(ctx)
}

func (s implPProfServerStub) Profiles(ctx __ipc.ServerContext) ([]string, error) {
	return s.impl.Profiles(ctx)
}

func (s implPProfServerStub) Profile(ctx *PProfProfileContextStub, i0 string, i1 int32) error {
	return s.impl.Profile(ctx, i0, i1)
}

func (s implPProfServerStub) CPUProfile(ctx *PProfCPUProfileContextStub, i0 int32) error {
	return s.impl.CPUProfile(ctx, i0)
}

func (s implPProfServerStub) Symbol(ctx __ipc.ServerContext, i0 []uint64) ([]string, error) {
	return s.impl.Symbol(ctx, i0)
}

func (s implPProfServerStub) Globber() *__ipc.GlobState {
	return s.gs
}

func (s implPProfServerStub) Describe__() []__ipc.InterfaceDesc {
	return []__ipc.InterfaceDesc{PProfDesc}
}

// PProfDesc describes the PProf interface.
var PProfDesc __ipc.InterfaceDesc = descPProf

// descPProf hides the desc to keep godoc clean.
var descPProf = __ipc.InterfaceDesc{
	Name:    "PProf",
	PkgPath: "veyron.io/veyron/veyron2/services/mgmt/pprof",
	Methods: []__ipc.MethodDesc{
		{
			Name: "CmdLine",
			Doc:  "// CmdLine returns the command-line arguments of the server, including\n// the name of the executable.",
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // []string
				{"", ``}, // error
			},
		},
		{
			Name: "Profiles",
			Doc:  "// Profiles returns the list of available profiles.",
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // []string
				{"", ``}, // error
			},
		},
		{
			Name: "Profile",
			Doc:  "// Profile streams the requested profile. The debug parameter enables\n// additional output. Passing debug=0 includes only the hexadecimal\n// addresses that pprof needs. Passing debug=1 adds comments translating\n// addresses to function names and line numbers, so that a programmer\n// can read the profile without tools.",
			InArgs: []__ipc.ArgDesc{
				{"name", ``},  // string
				{"debug", ``}, // int32
			},
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // error
			},
		},
		{
			Name: "CPUProfile",
			Doc:  "// CPUProfile enables CPU profiling for the requested duration and\n// streams the profile data.",
			InArgs: []__ipc.ArgDesc{
				{"seconds", ``}, // int32
			},
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // error
			},
		},
		{
			Name: "Symbol",
			Doc:  "// Symbol looks up the program counters and returns their respective\n// function names.",
			InArgs: []__ipc.ArgDesc{
				{"programCounters", ``}, // []uint64
			},
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // []string
				{"", ``}, // error
			},
		},
	},
}

func (s implPProfServerStub) Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error) {
	// TODO(toddw): Replace with new Describe__ implementation.
	result := __ipc.ServiceSignature{Methods: make(map[string]__ipc.MethodSignature)}
	result.Methods["CPUProfile"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "seconds", Type: 36},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 65},
		},

		OutStream: 67,
	}
	result.Methods["CmdLine"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 61},
			{Name: "", Type: 65},
		},
	}
	result.Methods["Profile"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "name", Type: 3},
			{Name: "debug", Type: 36},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 65},
		},

		OutStream: 67,
	}
	result.Methods["Profiles"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 61},
			{Name: "", Type: 65},
		},
	}
	result.Methods["Symbol"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "programCounters", Type: 68},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 61},
			{Name: "", Type: 65},
		},
	}

	result.TypeDefs = []__vdlutil.Any{
		__wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, __wiretype.NamedPrimitiveType{Type: 0x32, Name: "byte", Tags: []string(nil)}, __wiretype.SliceType{Elem: 0x42, Name: "", Tags: []string(nil)}, __wiretype.SliceType{Elem: 0x35, Name: "", Tags: []string(nil)}}

	return result, nil
}

// PProfProfileServerStream is the server stream for PProf.Profile.
type PProfProfileServerStream interface {
	// SendStream returns the send side of the PProf.Profile server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item []byte) error
	}
}

// PProfProfileContext represents the context passed to PProf.Profile.
type PProfProfileContext interface {
	__ipc.ServerContext
	PProfProfileServerStream
}

// PProfProfileContextStub is a wrapper that converts ipc.ServerCall into
// a typesafe stub that implements PProfProfileContext.
type PProfProfileContextStub struct {
	__ipc.ServerCall
}

// Init initializes PProfProfileContextStub from ipc.ServerCall.
func (s *PProfProfileContextStub) Init(call __ipc.ServerCall) {
	s.ServerCall = call
}

// SendStream returns the send side of the PProf.Profile server stream.
func (s *PProfProfileContextStub) SendStream() interface {
	Send(item []byte) error
} {
	return implPProfProfileContextSend{s}
}

type implPProfProfileContextSend struct {
	s *PProfProfileContextStub
}

func (s implPProfProfileContextSend) Send(item []byte) error {
	return s.s.Send(item)
}

// PProfCPUProfileServerStream is the server stream for PProf.CPUProfile.
type PProfCPUProfileServerStream interface {
	// SendStream returns the send side of the PProf.CPUProfile server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item []byte) error
	}
}

// PProfCPUProfileContext represents the context passed to PProf.CPUProfile.
type PProfCPUProfileContext interface {
	__ipc.ServerContext
	PProfCPUProfileServerStream
}

// PProfCPUProfileContextStub is a wrapper that converts ipc.ServerCall into
// a typesafe stub that implements PProfCPUProfileContext.
type PProfCPUProfileContextStub struct {
	__ipc.ServerCall
}

// Init initializes PProfCPUProfileContextStub from ipc.ServerCall.
func (s *PProfCPUProfileContextStub) Init(call __ipc.ServerCall) {
	s.ServerCall = call
}

// SendStream returns the send side of the PProf.CPUProfile server stream.
func (s *PProfCPUProfileContextStub) SendStream() interface {
	Send(item []byte) error
} {
	return implPProfCPUProfileContextSend{s}
}

type implPProfCPUProfileContextSend struct {
	s *PProfCPUProfileContextStub
}

func (s implPProfCPUProfileContextSend) Send(item []byte) error {
	return s.s.Send(item)
}
