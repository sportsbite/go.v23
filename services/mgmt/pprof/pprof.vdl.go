// This file was auto-generated by the vanadium vdl tool.
// Source: pprof.vdl

// Package pprof is used to access runtime profiling data in the format expected
// by the pprof visualization tool. For more information about pprof, see
// http://code.google.com/p/google-perftools/.
package pprof

import (
	// VDL system imports
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/ipc"
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/services/security/access"
)

// PProfClientMethods is the client interface
// containing PProf methods.
type PProfClientMethods interface {
	// CmdLine returns the command-line arguments of the server, including
	// the name of the executable.
	CmdLine(*context.T, ...ipc.CallOpt) ([]string, error)
	// Profiles returns the list of available profiles.
	Profiles(*context.T, ...ipc.CallOpt) ([]string, error)
	// Profile streams the requested profile. The debug parameter enables
	// additional output. Passing debug=0 includes only the hexadecimal
	// addresses that pprof needs. Passing debug=1 adds comments translating
	// addresses to function names and line numbers, so that a programmer
	// can read the profile without tools.
	Profile(ctx *context.T, name string, debug int32, opts ...ipc.CallOpt) (PProfProfileClientCall, error)
	// CPUProfile enables CPU profiling for the requested duration and
	// streams the profile data.
	CPUProfile(ctx *context.T, seconds int32, opts ...ipc.CallOpt) (PProfCPUProfileClientCall, error)
	// Symbol looks up the program counters and returns their respective
	// function names.
	Symbol(ctx *context.T, programCounters []uint64, opts ...ipc.CallOpt) ([]string, error)
}

// PProfClientStub adds universal methods to PProfClientMethods.
type PProfClientStub interface {
	PProfClientMethods
	ipc.UniversalServiceMethods
}

// PProfClient returns a client stub for PProf.
func PProfClient(name string, opts ...ipc.BindOpt) PProfClientStub {
	var client ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(ipc.Client); ok {
			client = clientOpt
		}
	}
	return implPProfClientStub{name, client}
}

type implPProfClientStub struct {
	name   string
	client ipc.Client
}

func (c implPProfClientStub) c(ctx *context.T) ipc.Client {
	if c.client != nil {
		return c.client
	}
	return v23.GetClient(ctx)
}

func (c implPProfClientStub) CmdLine(ctx *context.T, opts ...ipc.CallOpt) (o0 []string, err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "CmdLine", nil, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implPProfClientStub) Profiles(ctx *context.T, opts ...ipc.CallOpt) (o0 []string, err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Profiles", nil, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implPProfClientStub) Profile(ctx *context.T, i0 string, i1 int32, opts ...ipc.CallOpt) (ocall PProfProfileClientCall, err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Profile", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	ocall = &implPProfProfileClientCall{ClientCall: call}
	return
}

func (c implPProfClientStub) CPUProfile(ctx *context.T, i0 int32, opts ...ipc.CallOpt) (ocall PProfCPUProfileClientCall, err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "CPUProfile", []interface{}{i0}, opts...); err != nil {
		return
	}
	ocall = &implPProfCPUProfileClientCall{ClientCall: call}
	return
}

func (c implPProfClientStub) Symbol(ctx *context.T, i0 []uint64, opts ...ipc.CallOpt) (o0 []string, err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Symbol", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
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

// PProfProfileClientCall represents the call returned from PProf.Profile.
type PProfProfileClientCall interface {
	PProfProfileClientStream
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

type implPProfProfileClientCall struct {
	ipc.ClientCall
	valRecv []byte
	errRecv error
}

func (c *implPProfProfileClientCall) RecvStream() interface {
	Advance() bool
	Value() []byte
	Err() error
} {
	return implPProfProfileClientCallRecv{c}
}

type implPProfProfileClientCallRecv struct {
	c *implPProfProfileClientCall
}

func (c implPProfProfileClientCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implPProfProfileClientCallRecv) Value() []byte {
	return c.c.valRecv
}
func (c implPProfProfileClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implPProfProfileClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
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

// PProfCPUProfileClientCall represents the call returned from PProf.CPUProfile.
type PProfCPUProfileClientCall interface {
	PProfCPUProfileClientStream
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

type implPProfCPUProfileClientCall struct {
	ipc.ClientCall
	valRecv []byte
	errRecv error
}

func (c *implPProfCPUProfileClientCall) RecvStream() interface {
	Advance() bool
	Value() []byte
	Err() error
} {
	return implPProfCPUProfileClientCallRecv{c}
}

type implPProfCPUProfileClientCallRecv struct {
	c *implPProfCPUProfileClientCall
}

func (c implPProfCPUProfileClientCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implPProfCPUProfileClientCallRecv) Value() []byte {
	return c.c.valRecv
}
func (c implPProfCPUProfileClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implPProfCPUProfileClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
	return
}

// PProfServerMethods is the interface a server writer
// implements for PProf.
type PProfServerMethods interface {
	// CmdLine returns the command-line arguments of the server, including
	// the name of the executable.
	CmdLine(ipc.ServerCall) ([]string, error)
	// Profiles returns the list of available profiles.
	Profiles(ipc.ServerCall) ([]string, error)
	// Profile streams the requested profile. The debug parameter enables
	// additional output. Passing debug=0 includes only the hexadecimal
	// addresses that pprof needs. Passing debug=1 adds comments translating
	// addresses to function names and line numbers, so that a programmer
	// can read the profile without tools.
	Profile(call PProfProfileServerCall, name string, debug int32) error
	// CPUProfile enables CPU profiling for the requested duration and
	// streams the profile data.
	CPUProfile(call PProfCPUProfileServerCall, seconds int32) error
	// Symbol looks up the program counters and returns their respective
	// function names.
	Symbol(call ipc.ServerCall, programCounters []uint64) ([]string, error)
}

// PProfServerStubMethods is the server interface containing
// PProf methods, as expected by ipc.Server.
// The only difference between this interface and PProfServerMethods
// is the streaming methods.
type PProfServerStubMethods interface {
	// CmdLine returns the command-line arguments of the server, including
	// the name of the executable.
	CmdLine(ipc.ServerCall) ([]string, error)
	// Profiles returns the list of available profiles.
	Profiles(ipc.ServerCall) ([]string, error)
	// Profile streams the requested profile. The debug parameter enables
	// additional output. Passing debug=0 includes only the hexadecimal
	// addresses that pprof needs. Passing debug=1 adds comments translating
	// addresses to function names and line numbers, so that a programmer
	// can read the profile without tools.
	Profile(call *PProfProfileServerCallStub, name string, debug int32) error
	// CPUProfile enables CPU profiling for the requested duration and
	// streams the profile data.
	CPUProfile(call *PProfCPUProfileServerCallStub, seconds int32) error
	// Symbol looks up the program counters and returns their respective
	// function names.
	Symbol(call ipc.ServerCall, programCounters []uint64) ([]string, error)
}

// PProfServerStub adds universal methods to PProfServerStubMethods.
type PProfServerStub interface {
	PProfServerStubMethods
	// Describe the PProf interfaces.
	Describe__() []ipc.InterfaceDesc
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
	if gs := ipc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := ipc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implPProfServerStub struct {
	impl PProfServerMethods
	gs   *ipc.GlobState
}

func (s implPProfServerStub) CmdLine(call ipc.ServerCall) ([]string, error) {
	return s.impl.CmdLine(call)
}

func (s implPProfServerStub) Profiles(call ipc.ServerCall) ([]string, error) {
	return s.impl.Profiles(call)
}

func (s implPProfServerStub) Profile(call *PProfProfileServerCallStub, i0 string, i1 int32) error {
	return s.impl.Profile(call, i0, i1)
}

func (s implPProfServerStub) CPUProfile(call *PProfCPUProfileServerCallStub, i0 int32) error {
	return s.impl.CPUProfile(call, i0)
}

func (s implPProfServerStub) Symbol(call ipc.ServerCall, i0 []uint64) ([]string, error) {
	return s.impl.Symbol(call, i0)
}

func (s implPProfServerStub) Globber() *ipc.GlobState {
	return s.gs
}

func (s implPProfServerStub) Describe__() []ipc.InterfaceDesc {
	return []ipc.InterfaceDesc{PProfDesc}
}

// PProfDesc describes the PProf interface.
var PProfDesc ipc.InterfaceDesc = descPProf

// descPProf hides the desc to keep godoc clean.
var descPProf = ipc.InterfaceDesc{
	Name:    "PProf",
	PkgPath: "v.io/v23/services/mgmt/pprof",
	Methods: []ipc.MethodDesc{
		{
			Name: "CmdLine",
			Doc:  "// CmdLine returns the command-line arguments of the server, including\n// the name of the executable.",
			OutArgs: []ipc.ArgDesc{
				{"", ``}, // []string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Debug"))},
		},
		{
			Name: "Profiles",
			Doc:  "// Profiles returns the list of available profiles.",
			OutArgs: []ipc.ArgDesc{
				{"", ``}, // []string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Debug"))},
		},
		{
			Name: "Profile",
			Doc:  "// Profile streams the requested profile. The debug parameter enables\n// additional output. Passing debug=0 includes only the hexadecimal\n// addresses that pprof needs. Passing debug=1 adds comments translating\n// addresses to function names and line numbers, so that a programmer\n// can read the profile without tools.",
			InArgs: []ipc.ArgDesc{
				{"name", ``},  // string
				{"debug", ``}, // int32
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Debug"))},
		},
		{
			Name: "CPUProfile",
			Doc:  "// CPUProfile enables CPU profiling for the requested duration and\n// streams the profile data.",
			InArgs: []ipc.ArgDesc{
				{"seconds", ``}, // int32
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Debug"))},
		},
		{
			Name: "Symbol",
			Doc:  "// Symbol looks up the program counters and returns their respective\n// function names.",
			InArgs: []ipc.ArgDesc{
				{"programCounters", ``}, // []uint64
			},
			OutArgs: []ipc.ArgDesc{
				{"", ``}, // []string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Debug"))},
		},
	},
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

// PProfProfileServerCall represents the context passed to PProf.Profile.
type PProfProfileServerCall interface {
	ipc.ServerCall
	PProfProfileServerStream
}

// PProfProfileServerCallStub is a wrapper that converts ipc.StreamServerCall into
// a typesafe stub that implements PProfProfileServerCall.
type PProfProfileServerCallStub struct {
	ipc.StreamServerCall
}

// Init initializes PProfProfileServerCallStub from ipc.StreamServerCall.
func (s *PProfProfileServerCallStub) Init(call ipc.StreamServerCall) {
	s.StreamServerCall = call
}

// SendStream returns the send side of the PProf.Profile server stream.
func (s *PProfProfileServerCallStub) SendStream() interface {
	Send(item []byte) error
} {
	return implPProfProfileServerCallSend{s}
}

type implPProfProfileServerCallSend struct {
	s *PProfProfileServerCallStub
}

func (s implPProfProfileServerCallSend) Send(item []byte) error {
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

// PProfCPUProfileServerCall represents the context passed to PProf.CPUProfile.
type PProfCPUProfileServerCall interface {
	ipc.ServerCall
	PProfCPUProfileServerStream
}

// PProfCPUProfileServerCallStub is a wrapper that converts ipc.StreamServerCall into
// a typesafe stub that implements PProfCPUProfileServerCall.
type PProfCPUProfileServerCallStub struct {
	ipc.StreamServerCall
}

// Init initializes PProfCPUProfileServerCallStub from ipc.StreamServerCall.
func (s *PProfCPUProfileServerCallStub) Init(call ipc.StreamServerCall) {
	s.StreamServerCall = call
}

// SendStream returns the send side of the PProf.CPUProfile server stream.
func (s *PProfCPUProfileServerCallStub) SendStream() interface {
	Send(item []byte) error
} {
	return implPProfCPUProfileServerCallSend{s}
}

type implPProfCPUProfileServerCallSend struct {
	s *PProfCPUProfileServerCallStub
}

func (s implPProfCPUProfileServerCallSend) Send(item []byte) error {
	return s.s.Send(item)
}
