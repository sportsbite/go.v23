// This file was auto-generated by the vanadium vdl tool.
// Source: service.vdl

// Package mounttable defines a set of mount points and how to traverse them.
package mounttable

import (
	// VDL system imports
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/ipc"
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/naming"
	"v.io/v23/security"
	"v.io/v23/services/security/access/object"
)

type Tag string

func (Tag) __VDLReflect(struct {
	Name string "v.io/v23/services/mounttable.Tag"
}) {
}

func init() {
	vdl.Register((*Tag)(nil))
}

// Admin allow the client to SetPermissions or Delete the receiver.  It also subsumes
// all the other tags.
const Admin = Tag("Admin")

// Mount allows the client to Mount or Unmount at the named receiver.
// For example, to Mount onto a/b/c requires Mount or Admin access
// to a/b/c (and Read, Admin, or Resolve to a and a/b).
const Mount = Tag("Mount")

// Read allows the client to Glob any children of the node.  Thus to
// perform a Glob of a/* one must have Read access to a AND any other
// access to each child of a.  It also allows Resolution through the node.
const Read = Tag("Read")

// Create allows the client to create nodes below the receiver.
const Create = Tag("Create")

// Resolve allows one to resolve through the receiver.  Thus to Resolve
// a/b/c, one needs Admin, Resolve, or Read permission on a, a/b,
// and a/b/c.
const Resolve = Tag("Resolve")

// MountTableClientMethods is the client interface
// containing MountTable methods.
//
// MountTable defines the interface to talk to a mounttable.
//
// In all methods of MountTable, the receiver is the name bound to.
type MountTableClientMethods interface {
	// Object provides access control for Vanadium objects.
	//
	// Vanadium services implementing dynamic access control would typically
	// embed this interface and tag additional methods defined by the service
	// with one of Admin, Read, Write, Resolve etc. For example,
	// the VDL definition of the object would be:
	//
	//   package mypackage
	//
	//   import "v.io/v23/security/access"
	//   import "v.io/v23/security/access/object"
	//
	//   type MyObject interface {
	//     object.Object
	//     MyRead() (string, error) {access.Read}
	//     MyWrite(string) error    {access.Write}
	//   }
	//
	// If the set of pre-defined tags is insufficient, services may define their
	// own tag type and annotate all methods with this new type.
	// Instead of embedding this Object interface, define SetPermissions and GetPermissions in
	// their own interface. Authorization policies will typically respect
	// annotations of a single type. For example, the VDL definition of an object
	// would be:
	//
	//  package mypackage
	//
	//  import "v.io/v23/security/access"
	//
	//  type MyTag string
	//
	//  const (
	//    Blue = MyTag("Blue")
	//    Red  = MyTag("Red")
	//  )
	//
	//  type MyObject interface {
	//    MyMethod() (string, error) {Blue}
	//
	//    // Allow clients to change access via the access.Object interface:
	//    SetPermissions(acl access.Permissions, etag string) error         {Red}
	//    GetPermissions() (acl access.Permissions, etag string, err error) {Blue}
	//  }
	object.ObjectClientMethods
	// DEPRECATED: TODO(ashankar): Rename MountX to Mount and remove
	// MountX before the release.
	Mount(ctx *context.T, Server string, TTL uint32, Flags naming.MountFlag, opts ...ipc.CallOpt) error
	// Mount Server (a global name) onto the receiver.
	//
	// Subsequent mounts add to the servers mounted there.  The multiple
	// servers are considered equivalent and are meant solely for
	// availability, i.e., no load balancing is guaranteed.
	//
	// BlessingPatterns is a set of patterns that match the blessings
	// presented by Server to clients that initiate connections with it.
	// If empty, the mounttable makes the conservative assumption that the
	// blessings presented by the client invoking Mount will be the
	// blessings presented by Server.
	//
	// TTL is the number of seconds the mount is to last unless refreshed by
	// another mount of the same server.  A TTL of 0 represents an infinite
	// duration.  A server with an expired TTL should never appear in the
	// results nor affect the operation of any MountTable method, and should
	// act as if it was never present as far as the interface is concerned.
	//
	// Opts represents a bit mask of options.
	MountX(ctx *context.T, Server string, BlessingPatterns []security.BlessingPattern, TTL uint32, Flags naming.MountFlag, opts ...ipc.CallOpt) error
	// Unmount removes Server from the receiver.  If Server is empty, remove
	// all servers mounted there.
	// Returns a non-nil error iff Server remains mounted at the mount point.
	Unmount(ctx *context.T, Server string, opts ...ipc.CallOpt) error
	// Delete removes the receiver.  If the receiver has children, it will not
	// be removed unless DeleteSubtree is true in which case the whole subtree is
	// removed.
	Delete(ctx *context.T, DeleteSubtree bool, opts ...ipc.CallOpt) error
	// ResolveStep takes the next step in resolving a name.  Returns the next
	// servers to query and the suffix at those servers.
	ResolveStep(*context.T, ...ipc.CallOpt) (Entry naming.VDLMountEntry, err error)
	// Obsolete, left for backward compatability until all uses are killed.
	ResolveStepX(*context.T, ...ipc.CallOpt) (Entry naming.VDLMountEntry, err error)
}

// MountTableClientStub adds universal methods to MountTableClientMethods.
type MountTableClientStub interface {
	MountTableClientMethods
	ipc.UniversalServiceMethods
}

// MountTableClient returns a client stub for MountTable.
func MountTableClient(name string, opts ...ipc.BindOpt) MountTableClientStub {
	var client ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(ipc.Client); ok {
			client = clientOpt
		}
	}
	return implMountTableClientStub{name, client, object.ObjectClient(name, client)}
}

type implMountTableClientStub struct {
	name   string
	client ipc.Client

	object.ObjectClientStub
}

func (c implMountTableClientStub) c(ctx *context.T) ipc.Client {
	if c.client != nil {
		return c.client
	}
	return v23.GetClient(ctx)
}

func (c implMountTableClientStub) Mount(ctx *context.T, i0 string, i1 uint32, i2 naming.MountFlag, opts ...ipc.CallOpt) (err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Mount", []interface{}{i0, i1, i2}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implMountTableClientStub) MountX(ctx *context.T, i0 string, i1 []security.BlessingPattern, i2 uint32, i3 naming.MountFlag, opts ...ipc.CallOpt) (err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "MountX", []interface{}{i0, i1, i2, i3}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implMountTableClientStub) Unmount(ctx *context.T, i0 string, opts ...ipc.CallOpt) (err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Unmount", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implMountTableClientStub) Delete(ctx *context.T, i0 bool, opts ...ipc.CallOpt) (err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Delete", []interface{}{i0}, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

func (c implMountTableClientStub) ResolveStep(ctx *context.T, opts ...ipc.CallOpt) (o0 naming.VDLMountEntry, err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "ResolveStep", nil, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

func (c implMountTableClientStub) ResolveStepX(ctx *context.T, opts ...ipc.CallOpt) (o0 naming.VDLMountEntry, err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "ResolveStepX", nil, opts...); err != nil {
		return
	}
	err = call.Finish(&o0)
	return
}

// MountTableServerMethods is the interface a server writer
// implements for MountTable.
//
// MountTable defines the interface to talk to a mounttable.
//
// In all methods of MountTable, the receiver is the name bound to.
type MountTableServerMethods interface {
	// Object provides access control for Vanadium objects.
	//
	// Vanadium services implementing dynamic access control would typically
	// embed this interface and tag additional methods defined by the service
	// with one of Admin, Read, Write, Resolve etc. For example,
	// the VDL definition of the object would be:
	//
	//   package mypackage
	//
	//   import "v.io/v23/security/access"
	//   import "v.io/v23/security/access/object"
	//
	//   type MyObject interface {
	//     object.Object
	//     MyRead() (string, error) {access.Read}
	//     MyWrite(string) error    {access.Write}
	//   }
	//
	// If the set of pre-defined tags is insufficient, services may define their
	// own tag type and annotate all methods with this new type.
	// Instead of embedding this Object interface, define SetPermissions and GetPermissions in
	// their own interface. Authorization policies will typically respect
	// annotations of a single type. For example, the VDL definition of an object
	// would be:
	//
	//  package mypackage
	//
	//  import "v.io/v23/security/access"
	//
	//  type MyTag string
	//
	//  const (
	//    Blue = MyTag("Blue")
	//    Red  = MyTag("Red")
	//  )
	//
	//  type MyObject interface {
	//    MyMethod() (string, error) {Blue}
	//
	//    // Allow clients to change access via the access.Object interface:
	//    SetPermissions(acl access.Permissions, etag string) error         {Red}
	//    GetPermissions() (acl access.Permissions, etag string, err error) {Blue}
	//  }
	object.ObjectServerMethods
	// DEPRECATED: TODO(ashankar): Rename MountX to Mount and remove
	// MountX before the release.
	Mount(call ipc.ServerCall, Server string, TTL uint32, Flags naming.MountFlag) error
	// Mount Server (a global name) onto the receiver.
	//
	// Subsequent mounts add to the servers mounted there.  The multiple
	// servers are considered equivalent and are meant solely for
	// availability, i.e., no load balancing is guaranteed.
	//
	// BlessingPatterns is a set of patterns that match the blessings
	// presented by Server to clients that initiate connections with it.
	// If empty, the mounttable makes the conservative assumption that the
	// blessings presented by the client invoking Mount will be the
	// blessings presented by Server.
	//
	// TTL is the number of seconds the mount is to last unless refreshed by
	// another mount of the same server.  A TTL of 0 represents an infinite
	// duration.  A server with an expired TTL should never appear in the
	// results nor affect the operation of any MountTable method, and should
	// act as if it was never present as far as the interface is concerned.
	//
	// Opts represents a bit mask of options.
	MountX(call ipc.ServerCall, Server string, BlessingPatterns []security.BlessingPattern, TTL uint32, Flags naming.MountFlag) error
	// Unmount removes Server from the receiver.  If Server is empty, remove
	// all servers mounted there.
	// Returns a non-nil error iff Server remains mounted at the mount point.
	Unmount(call ipc.ServerCall, Server string) error
	// Delete removes the receiver.  If the receiver has children, it will not
	// be removed unless DeleteSubtree is true in which case the whole subtree is
	// removed.
	Delete(call ipc.ServerCall, DeleteSubtree bool) error
	// ResolveStep takes the next step in resolving a name.  Returns the next
	// servers to query and the suffix at those servers.
	ResolveStep(ipc.ServerCall) (Entry naming.VDLMountEntry, err error)
	// Obsolete, left for backward compatability until all uses are killed.
	ResolveStepX(ipc.ServerCall) (Entry naming.VDLMountEntry, err error)
}

// MountTableServerStubMethods is the server interface containing
// MountTable methods, as expected by ipc.Server.
// There is no difference between this interface and MountTableServerMethods
// since there are no streaming methods.
type MountTableServerStubMethods MountTableServerMethods

// MountTableServerStub adds universal methods to MountTableServerStubMethods.
type MountTableServerStub interface {
	MountTableServerStubMethods
	// Describe the MountTable interfaces.
	Describe__() []ipc.InterfaceDesc
}

// MountTableServer returns a server stub for MountTable.
// It converts an implementation of MountTableServerMethods into
// an object that may be used by ipc.Server.
func MountTableServer(impl MountTableServerMethods) MountTableServerStub {
	stub := implMountTableServerStub{
		impl:             impl,
		ObjectServerStub: object.ObjectServer(impl),
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

type implMountTableServerStub struct {
	impl MountTableServerMethods
	object.ObjectServerStub
	gs *ipc.GlobState
}

func (s implMountTableServerStub) Mount(call ipc.ServerCall, i0 string, i1 uint32, i2 naming.MountFlag) error {
	return s.impl.Mount(call, i0, i1, i2)
}

func (s implMountTableServerStub) MountX(call ipc.ServerCall, i0 string, i1 []security.BlessingPattern, i2 uint32, i3 naming.MountFlag) error {
	return s.impl.MountX(call, i0, i1, i2, i3)
}

func (s implMountTableServerStub) Unmount(call ipc.ServerCall, i0 string) error {
	return s.impl.Unmount(call, i0)
}

func (s implMountTableServerStub) Delete(call ipc.ServerCall, i0 bool) error {
	return s.impl.Delete(call, i0)
}

func (s implMountTableServerStub) ResolveStep(call ipc.ServerCall) (naming.VDLMountEntry, error) {
	return s.impl.ResolveStep(call)
}

func (s implMountTableServerStub) ResolveStepX(call ipc.ServerCall) (naming.VDLMountEntry, error) {
	return s.impl.ResolveStepX(call)
}

func (s implMountTableServerStub) Globber() *ipc.GlobState {
	return s.gs
}

func (s implMountTableServerStub) Describe__() []ipc.InterfaceDesc {
	return []ipc.InterfaceDesc{MountTableDesc, object.ObjectDesc}
}

// MountTableDesc describes the MountTable interface.
var MountTableDesc ipc.InterfaceDesc = descMountTable

// descMountTable hides the desc to keep godoc clean.
var descMountTable = ipc.InterfaceDesc{
	Name:    "MountTable",
	PkgPath: "v.io/v23/services/mounttable",
	Doc:     "// MountTable defines the interface to talk to a mounttable.\n//\n// In all methods of MountTable, the receiver is the name bound to.",
	Embeds: []ipc.EmbedDesc{
		{"Object", "v.io/v23/services/security/access/object", "// Object provides access control for Vanadium objects.\n//\n// Vanadium services implementing dynamic access control would typically\n// embed this interface and tag additional methods defined by the service\n// with one of Admin, Read, Write, Resolve etc. For example,\n// the VDL definition of the object would be:\n//\n//   package mypackage\n//\n//   import \"v.io/v23/security/access\"\n//   import \"v.io/v23/security/access/object\"\n//\n//   type MyObject interface {\n//     object.Object\n//     MyRead() (string, error) {access.Read}\n//     MyWrite(string) error    {access.Write}\n//   }\n//\n// If the set of pre-defined tags is insufficient, services may define their\n// own tag type and annotate all methods with this new type.\n// Instead of embedding this Object interface, define SetPermissions and GetPermissions in\n// their own interface. Authorization policies will typically respect\n// annotations of a single type. For example, the VDL definition of an object\n// would be:\n//\n//  package mypackage\n//\n//  import \"v.io/v23/security/access\"\n//\n//  type MyTag string\n//\n//  const (\n//    Blue = MyTag(\"Blue\")\n//    Red  = MyTag(\"Red\")\n//  )\n//\n//  type MyObject interface {\n//    MyMethod() (string, error) {Blue}\n//\n//    // Allow clients to change access via the access.Object interface:\n//    SetPermissions(acl access.Permissions, etag string) error         {Red}\n//    GetPermissions() (acl access.Permissions, etag string, err error) {Blue}\n//  }"},
	},
	Methods: []ipc.MethodDesc{
		{
			Name: "Mount",
			Doc:  "// DEPRECATED: TODO(ashankar): Rename MountX to Mount and remove\n// MountX before the release.",
			InArgs: []ipc.ArgDesc{
				{"Server", ``}, // string
				{"TTL", ``},    // uint32
				{"Flags", ``},  // naming.MountFlag
			},
		},
		{
			Name: "MountX",
			Doc:  "// Mount Server (a global name) onto the receiver.\n//\n// Subsequent mounts add to the servers mounted there.  The multiple\n// servers are considered equivalent and are meant solely for\n// availability, i.e., no load balancing is guaranteed.\n//\n// BlessingPatterns is a set of patterns that match the blessings\n// presented by Server to clients that initiate connections with it.\n// If empty, the mounttable makes the conservative assumption that the\n// blessings presented by the client invoking Mount will be the\n// blessings presented by Server.\n//\n// TTL is the number of seconds the mount is to last unless refreshed by\n// another mount of the same server.  A TTL of 0 represents an infinite\n// duration.  A server with an expired TTL should never appear in the\n// results nor affect the operation of any MountTable method, and should\n// act as if it was never present as far as the interface is concerned.\n//\n// Opts represents a bit mask of options.",
			InArgs: []ipc.ArgDesc{
				{"Server", ``},           // string
				{"BlessingPatterns", ``}, // []security.BlessingPattern
				{"TTL", ``},              // uint32
				{"Flags", ``},            // naming.MountFlag
			},
		},
		{
			Name: "Unmount",
			Doc:  "// Unmount removes Server from the receiver.  If Server is empty, remove\n// all servers mounted there.\n// Returns a non-nil error iff Server remains mounted at the mount point.",
			InArgs: []ipc.ArgDesc{
				{"Server", ``}, // string
			},
		},
		{
			Name: "Delete",
			Doc:  "// Delete removes the receiver.  If the receiver has children, it will not\n// be removed unless DeleteSubtree is true in which case the whole subtree is\n// removed.",
			InArgs: []ipc.ArgDesc{
				{"DeleteSubtree", ``}, // bool
			},
		},
		{
			Name: "ResolveStep",
			Doc:  "// ResolveStep takes the next step in resolving a name.  Returns the next\n// servers to query and the suffix at those servers.",
			OutArgs: []ipc.ArgDesc{
				{"Entry", ``}, // naming.VDLMountEntry
			},
		},
		{
			Name: "ResolveStepX",
			Doc:  "// Obsolete, left for backward compatability until all uses are killed.",
			OutArgs: []ipc.ArgDesc{
				{"Entry", ``}, // naming.VDLMountEntry
			},
		},
	},
}
