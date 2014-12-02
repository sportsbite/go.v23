// This file was auto-generated by the veyron vdl tool.
// Source: build.vdl

// Package build supports building and describing Veyron binaries.
//
// TODO(jsimsa): Switch Architecture, Format, and OperatingSystem type
// to enum when supported.
package build

import (
	"veyron.io/veyron/veyron2/services/mgmt/binary"

	// The non-user imports are prefixed with "__" to prevent collisions.
	__io "io"
	__veyron2 "veyron.io/veyron/veyron2"
	__context "veyron.io/veyron/veyron2/context"
	__ipc "veyron.io/veyron/veyron2/ipc"
	__vdl "veyron.io/veyron/veyron2/vdl"
	__vdlutil "veyron.io/veyron/veyron2/vdl/vdlutil"
	__wiretype "veyron.io/veyron/veyron2/wiretype"
)

// TODO(toddw): Remove this line once the new signature support is done.
// It corrects a bug where __wiretype is unused in VDL pacakges where only
// bootstrap types are used on interfaces.
const _ = __wiretype.TypeIDInvalid

// Architecture specifies the hardware architecture of a host.
type Architecture string

func (Architecture) __VDLReflect(struct {
	Name string "veyron.io/veyron/veyron2/services/mgmt/build.Architecture"
}) {
}

// Format specifies the file format of a host.
type Format string

func (Format) __VDLReflect(struct {
	Name string "veyron.io/veyron/veyron2/services/mgmt/build.Format"
}) {
}

// OperatingSystem specifies the operating system of a host.
type OperatingSystem string

func (OperatingSystem) __VDLReflect(struct {
	Name string "veyron.io/veyron/veyron2/services/mgmt/build.OperatingSystem"
}) {
}

// File records the name and contents of a file.
type File struct {
	Name     string
	Contents []byte
}

func (File) __VDLReflect(struct {
	Name string "veyron.io/veyron/veyron2/services/mgmt/build.File"
}) {
}

func init() {
	__vdl.Register(Architecture(""))
	__vdl.Register(Format(""))
	__vdl.Register(OperatingSystem(""))
	__vdl.Register(File{})
}

const X86 = Architecture("386")

const AMD64 = Architecture("amd64")

const ARM = Architecture("arm")

const UnsupportedArchitecture = Architecture("unsupported")

const ELF = Format("ELF")

const MACH = Format("MACH")

const PE = Format("PE")

const UnsupportedFormat = Format("unsupported")

const Darwin = OperatingSystem("darwin")

const Linux = OperatingSystem("linux")

const Windows = OperatingSystem("windows")

const UnsupportedOS = OperatingSystem("unsupported")

// BuilderClientMethods is the client interface
// containing Builder methods.
//
// Builder describes an interface for building binaries from source.
type BuilderClientMethods interface {
	// Build streams sources to the build server, which then attempts to
	// build the sources and streams back the compiled binaries.
	Build(ctx __context.T, Arch Architecture, OS OperatingSystem, opts ...__ipc.CallOpt) (BuilderBuildCall, error)
	// Describe generates a description for a binary identified by
	// the given Object name.
	Describe(ctx __context.T, Name string, opts ...__ipc.CallOpt) (binary.Description, error)
}

// BuilderClientStub adds universal methods to BuilderClientMethods.
type BuilderClientStub interface {
	BuilderClientMethods
	__ipc.UniversalServiceMethods
}

// BuilderClient returns a client stub for Builder.
func BuilderClient(name string, opts ...__ipc.BindOpt) BuilderClientStub {
	var client __ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(__ipc.Client); ok {
			client = clientOpt
		}
	}
	return implBuilderClientStub{name, client}
}

type implBuilderClientStub struct {
	name   string
	client __ipc.Client
}

func (c implBuilderClientStub) c(ctx __context.T) __ipc.Client {
	if c.client != nil {
		return c.client
	}
	return __veyron2.RuntimeFromContext(ctx).Client()
}

func (c implBuilderClientStub) Build(ctx __context.T, i0 Architecture, i1 OperatingSystem, opts ...__ipc.CallOpt) (ocall BuilderBuildCall, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Build", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	ocall = &implBuilderBuildCall{Call: call}
	return
}

func (c implBuilderClientStub) Describe(ctx __context.T, i0 string, opts ...__ipc.CallOpt) (o0 binary.Description, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Describe", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implBuilderClientStub) Signature(ctx __context.T, opts ...__ipc.CallOpt) (o0 __ipc.ServiceSignature, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// BuilderBuildClientStream is the client stream for Builder.Build.
type BuilderBuildClientStream interface {
	// RecvStream returns the receiver side of the Builder.Build client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() File
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Builder.Build client stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending, or if Send is called after Close or Cancel.  Blocks if
		// there is no buffer space; will unblock when buffer space is available or
		// after Cancel.
		Send(item File) error
		// Close indicates to the server that no more items will be sent; server
		// Recv calls will receive io.EOF after all sent items.  This is an optional
		// call - e.g. a client might call Close if it needs to continue receiving
		// items from the server after it's done sending.  Returns errors
		// encountered while closing, or if Close is called after Cancel.  Like
		// Send, blocks if there is no buffer space available.
		Close() error
	}
}

// BuilderBuildCall represents the call returned from Builder.Build.
type BuilderBuildCall interface {
	BuilderBuildClientStream
	// Finish performs the equivalent of SendStream().Close, then blocks until
	// the server is done, and returns the positional return values for the call.
	//
	// Finish returns immediately if Cancel has been called; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless Cancel
	// has been called or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() ([]byte, error)
	// Cancel cancels the RPC, notifying the server to stop processing.  It is
	// safe to call Cancel concurrently with any of the other stream methods.
	// Calling Cancel after Finish has returned is a no-op.
	Cancel()
}

type implBuilderBuildCall struct {
	__ipc.Call
	valRecv File
	errRecv error
}

func (c *implBuilderBuildCall) RecvStream() interface {
	Advance() bool
	Value() File
	Err() error
} {
	return implBuilderBuildCallRecv{c}
}

type implBuilderBuildCallRecv struct {
	c *implBuilderBuildCall
}

func (c implBuilderBuildCallRecv) Advance() bool {
	c.c.valRecv = File{}
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implBuilderBuildCallRecv) Value() File {
	return c.c.valRecv
}
func (c implBuilderBuildCallRecv) Err() error {
	if c.c.errRecv == __io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implBuilderBuildCall) SendStream() interface {
	Send(item File) error
	Close() error
} {
	return implBuilderBuildCallSend{c}
}

type implBuilderBuildCallSend struct {
	c *implBuilderBuildCall
}

func (c implBuilderBuildCallSend) Send(item File) error {
	return c.c.Send(item)
}
func (c implBuilderBuildCallSend) Close() error {
	return c.c.CloseSend()
}
func (c *implBuilderBuildCall) Finish() (o0 []byte, err error) {
	if ierr := c.Call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// BuilderServerMethods is the interface a server writer
// implements for Builder.
//
// Builder describes an interface for building binaries from source.
type BuilderServerMethods interface {
	// Build streams sources to the build server, which then attempts to
	// build the sources and streams back the compiled binaries.
	Build(ctx BuilderBuildContext, Arch Architecture, OS OperatingSystem) ([]byte, error)
	// Describe generates a description for a binary identified by
	// the given Object name.
	Describe(ctx __ipc.ServerContext, Name string) (binary.Description, error)
}

// BuilderServerStubMethods is the server interface containing
// Builder methods, as expected by ipc.Server.
// The only difference between this interface and BuilderServerMethods
// is the streaming methods.
type BuilderServerStubMethods interface {
	// Build streams sources to the build server, which then attempts to
	// build the sources and streams back the compiled binaries.
	Build(ctx *BuilderBuildContextStub, Arch Architecture, OS OperatingSystem) ([]byte, error)
	// Describe generates a description for a binary identified by
	// the given Object name.
	Describe(ctx __ipc.ServerContext, Name string) (binary.Description, error)
}

// BuilderServerStub adds universal methods to BuilderServerStubMethods.
type BuilderServerStub interface {
	BuilderServerStubMethods
	// Describe the Builder interfaces.
	Describe__() []__ipc.InterfaceDesc
	// Signature will be replaced with Describe__.
	Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error)
}

// BuilderServer returns a server stub for Builder.
// It converts an implementation of BuilderServerMethods into
// an object that may be used by ipc.Server.
func BuilderServer(impl BuilderServerMethods) BuilderServerStub {
	stub := implBuilderServerStub{
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

type implBuilderServerStub struct {
	impl BuilderServerMethods
	gs   *__ipc.GlobState
}

func (s implBuilderServerStub) Build(ctx *BuilderBuildContextStub, i0 Architecture, i1 OperatingSystem) ([]byte, error) {
	return s.impl.Build(ctx, i0, i1)
}

func (s implBuilderServerStub) Describe(ctx __ipc.ServerContext, i0 string) (binary.Description, error) {
	return s.impl.Describe(ctx, i0)
}

func (s implBuilderServerStub) Globber() *__ipc.GlobState {
	return s.gs
}

func (s implBuilderServerStub) Describe__() []__ipc.InterfaceDesc {
	return []__ipc.InterfaceDesc{BuilderDesc}
}

// BuilderDesc describes the Builder interface.
var BuilderDesc __ipc.InterfaceDesc = descBuilder

// descBuilder hides the desc to keep godoc clean.
var descBuilder = __ipc.InterfaceDesc{
	Name:    "Builder",
	PkgPath: "veyron.io/veyron/veyron2/services/mgmt/build",
	Doc:     "// Builder describes an interface for building binaries from source.",
	Methods: []__ipc.MethodDesc{
		{
			Name: "Build",
			Doc:  "// Build streams sources to the build server, which then attempts to\n// build the sources and streams back the compiled binaries.",
			InArgs: []__ipc.ArgDesc{
				{"Arch", ``}, // Architecture
				{"OS", ``},   // OperatingSystem
			},
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // []byte
				{"", ``}, // error
			},
		},
		{
			Name: "Describe",
			Doc:  "// Describe generates a description for a binary identified by\n// the given Object name.",
			InArgs: []__ipc.ArgDesc{
				{"Name", ``}, // string
			},
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // binary.Description
				{"", ``}, // error
			},
		},
	},
}

func (s implBuilderServerStub) Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error) {
	// TODO(toddw): Replace with new Describe__ implementation.
	result := __ipc.ServiceSignature{Methods: make(map[string]__ipc.MethodSignature)}
	result.Methods["Build"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "Arch", Type: 65},
			{Name: "OS", Type: 66},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 68},
			{Name: "", Type: 69},
		},
		InStream:  70,
		OutStream: 70,
	}
	result.Methods["Describe"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "Name", Type: 3},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 72},
			{Name: "", Type: 69},
		},
	}

	result.TypeDefs = []__vdlutil.Any{
		__wiretype.NamedPrimitiveType{Type: 0x3, Name: "veyron.io/veyron/veyron2/services/mgmt/build.Architecture", Tags: []string(nil)}, __wiretype.NamedPrimitiveType{Type: 0x3, Name: "veyron.io/veyron/veyron2/services/mgmt/build.OperatingSystem", Tags: []string(nil)}, __wiretype.NamedPrimitiveType{Type: 0x32, Name: "byte", Tags: []string(nil)}, __wiretype.SliceType{Elem: 0x43, Name: "", Tags: []string(nil)}, __wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, __wiretype.StructType{
			[]__wiretype.FieldType{
				__wiretype.FieldType{Type: 0x3, Name: "Name"},
				__wiretype.FieldType{Type: 0x44, Name: "Contents"},
			},
			"veyron.io/veyron/veyron2/services/mgmt/build.File", []string(nil)},
		__wiretype.MapType{Key: 0x3, Elem: 0x2, Name: "", Tags: []string(nil)}, __wiretype.StructType{
			[]__wiretype.FieldType{
				__wiretype.FieldType{Type: 0x3, Name: "Name"},
				__wiretype.FieldType{Type: 0x47, Name: "Profiles"},
			},
			"veyron.io/veyron/veyron2/services/mgmt/binary.Description", []string(nil)},
	}

	return result, nil
}

// BuilderBuildServerStream is the server stream for Builder.Build.
type BuilderBuildServerStream interface {
	// RecvStream returns the receiver side of the Builder.Build server stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() File
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Builder.Build server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item File) error
	}
}

// BuilderBuildContext represents the context passed to Builder.Build.
type BuilderBuildContext interface {
	__ipc.ServerContext
	BuilderBuildServerStream
}

// BuilderBuildContextStub is a wrapper that converts ipc.ServerCall into
// a typesafe stub that implements BuilderBuildContext.
type BuilderBuildContextStub struct {
	__ipc.ServerCall
	valRecv File
	errRecv error
}

// Init initializes BuilderBuildContextStub from ipc.ServerCall.
func (s *BuilderBuildContextStub) Init(call __ipc.ServerCall) {
	s.ServerCall = call
}

// RecvStream returns the receiver side of the Builder.Build server stream.
func (s *BuilderBuildContextStub) RecvStream() interface {
	Advance() bool
	Value() File
	Err() error
} {
	return implBuilderBuildContextRecv{s}
}

type implBuilderBuildContextRecv struct {
	s *BuilderBuildContextStub
}

func (s implBuilderBuildContextRecv) Advance() bool {
	s.s.valRecv = File{}
	s.s.errRecv = s.s.Recv(&s.s.valRecv)
	return s.s.errRecv == nil
}
func (s implBuilderBuildContextRecv) Value() File {
	return s.s.valRecv
}
func (s implBuilderBuildContextRecv) Err() error {
	if s.s.errRecv == __io.EOF {
		return nil
	}
	return s.s.errRecv
}

// SendStream returns the send side of the Builder.Build server stream.
func (s *BuilderBuildContextStub) SendStream() interface {
	Send(item File) error
} {
	return implBuilderBuildContextSend{s}
}

type implBuilderBuildContextSend struct {
	s *BuilderBuildContextStub
}

func (s implBuilderBuildContextSend) Send(item File) error {
	return s.s.Send(item)
}
