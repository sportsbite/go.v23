// This file was auto-generated by the veyron idl tool.
// Source: service.idl

// Package mounttable defines a set of mount points and how to traverse them.
package mounttable

import (
	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_idl "veyron2/idl"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt/r"
	_gen_wiretype "veyron2/wiretype"
)

// MountedServer represents a server mounted on a specific name.
type MountedServer struct {
	// Server is the OA that's mounted.
	Server string
	// TTL is the remaining time (in seconds) before the mount entry expires.
	TTL uint32
}

// MountEntry represents a given name mounted in the mounttable.
type MountEntry struct {
	// Name is the mounted name.
	Name string
	// Link (if present) specifies the link name (Servers is nil).
	Link string
	// Servers (if present) specifies the mounted names (Link is empty).
	Servers []MountedServer
}

// Globable is the interface the client binds and uses.
// Globable_InternalNoTagGetter is the interface without the TagGetter
// and UnresolveStep methods (both framework-added, rathern than user-defined),
// to enable embedding without method collisions.  Not to be used directly by
// clients.
type Globable_InternalNoTagGetter interface {

	// Glob returns all matching entries at the given server.
	//
	// If Recursive is true, list all nodes below the matching entries.
	Glob(pattern string, opts ..._gen_ipc.ClientCallOpt) (reply GlobableGlobStream, err error)
}
type Globable interface {
	_gen_idl.TagGetter
	// UnresolveStep returns the names for the remote service, rooted at the
	// service's immediate namespace ancestor.
	UnresolveStep(opts ..._gen_ipc.ClientCallOpt) ([]string, error)
	Globable_InternalNoTagGetter
}

// GlobableService is the interface the server implements.
type GlobableService interface {

	// Glob returns all matching entries at the given server.
	//
	// If Recursive is true, list all nodes below the matching entries.
	Glob(context _gen_ipc.Context, pattern string, stream GlobableServiceGlobStream) (err error)
}

// GlobableGlobStream is the interface for streaming responses of the method
// Glob in the service interface Globable.
type GlobableGlobStream interface {

	// Recv returns the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of input.
	Recv() (item MountEntry, err error)

	// Finish closes the stream and returns the positional return values for
	// call.
	Finish() (err error)

	// Cancel cancels the RPC, notifying the server to stop processing.
	Cancel()
}

// Implementation of the GlobableGlobStream interface that is not exported.
type implGlobableGlobStream struct {
	clientCall _gen_ipc.ClientCall
}

func (c *implGlobableGlobStream) Recv() (item MountEntry, err error) {
	err = c.clientCall.Recv(&item)
	return
}

func (c *implGlobableGlobStream) Finish() (err error) {
	if ierr := c.clientCall.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implGlobableGlobStream) Cancel() {
	c.clientCall.Cancel()
}

// GlobableServiceGlobStream is the interface for streaming responses of the method
// Glob in the service interface Globable.
type GlobableServiceGlobStream interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.
	Send(item MountEntry) error
}

// Implementation of the GlobableServiceGlobStream interface that is not exported.
type implGlobableServiceGlobStream struct {
	serverCall _gen_ipc.ServerCall
}

func (s *implGlobableServiceGlobStream) Send(item MountEntry) error {
	return s.serverCall.Send(item)
}

// BindGlobable returns the client stub implementing the Globable
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindGlobable(name string, opts ..._gen_ipc.BindOpt) (Globable, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		client = _gen_rt.R().Client()
	case 1:
		switch o := opts[0].(type) {
		case _gen_ipc.Runtime:
			client = o.Client()
		case _gen_ipc.Client:
			client = o
		default:
			return nil, _gen_idl.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_idl.ErrTooManyOptionsToBind
	}
	stub := &clientStubGlobable{client: client, name: name}

	return stub, nil
}

// NewServerGlobable creates a new server stub.
//
// It takes a regular server implementing the GlobableService
// interface, and returns a new server stub.
func NewServerGlobable(server GlobableService) interface{} {
	return &ServerStubGlobable{
		service: server,
	}
}

// clientStubGlobable implements Globable.
type clientStubGlobable struct {
	client _gen_ipc.Client
	name   string
}

func (c *clientStubGlobable) GetMethodTags(method string) []interface{} {
	return GetGlobableMethodTags(method)
}

func (__gen_c *clientStubGlobable) Glob(pattern string, opts ..._gen_ipc.ClientCallOpt) (reply GlobableGlobStream, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Glob", []interface{}{pattern}, opts...); err != nil {
		return
	}
	reply = &implGlobableGlobStream{clientCall: call}
	return
}

func (c *clientStubGlobable) UnresolveStep(opts ..._gen_ipc.ClientCallOpt) (reply []string, err error) {
	var call _gen_ipc.ClientCall
	if call, err = c.client.StartCall(c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubGlobable wraps a server that implements
// GlobableService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubGlobable struct {
	service GlobableService
}

func (s *ServerStubGlobable) GetMethodTags(method string) []interface{} {
	return GetGlobableMethodTags(method)
}

func (s *ServerStubGlobable) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Glob"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "pattern", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},

		OutStream: 68,
	}

	result.TypeDefs = []_gen_idl.AnyData{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "Server"},
				_gen_wiretype.FieldType{Type: 0x34, Name: "TTL"},
			},
			"MountedServer", []string(nil)},
		_gen_wiretype.SliceType{Elem: 0x42, Name: "", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "Name"},
				_gen_wiretype.FieldType{Type: 0x3, Name: "Link"},
				_gen_wiretype.FieldType{Type: 0x43, Name: "Servers"},
			},
			"MountEntry", []string(nil)},
	}

	return result, nil
}

func (s *ServerStubGlobable) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := s.service.(_gen_ipc.Unresolver); ok {
		return unresolver.UnresolveStep(call)
	}
	if call.Server() == nil {
		return
	}
	var published []string
	if published, err = call.Server().Published(); err != nil || published == nil {
		return
	}
	reply = make([]string, len(published))
	for i, p := range published {
		reply[i] = _gen_naming.Join(p, call.Name())
	}
	return
}

func (__gen_s *ServerStubGlobable) Glob(call _gen_ipc.ServerCall, pattern string) (err error) {
	stream := &implGlobableServiceGlobStream{serverCall: call}
	err = __gen_s.service.Glob(call, pattern, stream)
	return
}

func GetGlobableMethodTags(method string) []interface{} {
	switch method {
	case "Glob":
		return []interface{}{}
	default:
		return nil
	}
}

// MountTable defines the interface to talk to a mounttable.
// MountTable is the interface the client binds and uses.
// MountTable_InternalNoTagGetter is the interface without the TagGetter
// and UnresolveStep methods (both framework-added, rathern than user-defined),
// to enable embedding without method collisions.  Not to be used directly by
// clients.
type MountTable_InternalNoTagGetter interface {

	// Mount Server (a global name) onto the receiver.
	// Subsequent mounts add to the servers mounted there.  The multiple
	// servers are considered equivalent and are meant solely for
	// availability, i.e., no load balancing is guaranteed.
	//
	// TTL is the number of seconds the mount is to last unless refreshed by
	// another mount of the same server.  A TTL of 0 represents an infinite
	// duration.  A server with an exipred TTL should never appear in the
	// results nor affect the operation of any MountTable method, and should
	// act as if it was never present as far as the interface is concerned.
	Mount(Server string, TTL uint32, opts ..._gen_ipc.ClientCallOpt) (err error)

	// Unmount removes Server from the mount point.  If Server is empty, remove all
	// servers mounted there.
	Unmount(Server string, opts ..._gen_ipc.ClientCallOpt) (err error)

	// Link creates a link in the tree from the receiver to LinkName.
	Link(LinkName string, opts ..._gen_ipc.ClientCallOpt) (err error)

	// Unlink removes a link at the receiver.
	Unlink(opts ..._gen_ipc.ClientCallOpt) (err error)

	// ResolveStep takes the next step in resolving a name.  Returns the next
	// servers to query and the suffix at those servers.
	ResolveStep(opts ..._gen_ipc.ClientCallOpt) (Servers []MountedServer, Suffix string, err error)

	Globable_InternalNoTagGetter
}
type MountTable interface {
	_gen_idl.TagGetter
	// UnresolveStep returns the names for the remote service, rooted at the
	// service's immediate namespace ancestor.
	UnresolveStep(opts ..._gen_ipc.ClientCallOpt) ([]string, error)
	MountTable_InternalNoTagGetter
}

// MountTableService is the interface the server implements.
type MountTableService interface {

	// Mount Server (a global name) onto the receiver.
	// Subsequent mounts add to the servers mounted there.  The multiple
	// servers are considered equivalent and are meant solely for
	// availability, i.e., no load balancing is guaranteed.
	//
	// TTL is the number of seconds the mount is to last unless refreshed by
	// another mount of the same server.  A TTL of 0 represents an infinite
	// duration.  A server with an exipred TTL should never appear in the
	// results nor affect the operation of any MountTable method, and should
	// act as if it was never present as far as the interface is concerned.
	Mount(context _gen_ipc.Context, Server string, TTL uint32) (err error)

	// Unmount removes Server from the mount point.  If Server is empty, remove all
	// servers mounted there.
	Unmount(context _gen_ipc.Context, Server string) (err error)

	// Link creates a link in the tree from the receiver to LinkName.
	Link(context _gen_ipc.Context, LinkName string) (err error)

	// Unlink removes a link at the receiver.
	Unlink(context _gen_ipc.Context) (err error)

	// ResolveStep takes the next step in resolving a name.  Returns the next
	// servers to query and the suffix at those servers.
	ResolveStep(context _gen_ipc.Context) (Servers []MountedServer, Suffix string, err error)

	GlobableService
}

// BindMountTable returns the client stub implementing the MountTable
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindMountTable(name string, opts ..._gen_ipc.BindOpt) (MountTable, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		client = _gen_rt.R().Client()
	case 1:
		switch o := opts[0].(type) {
		case _gen_ipc.Runtime:
			client = o.Client()
		case _gen_ipc.Client:
			client = o
		default:
			return nil, _gen_idl.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_idl.ErrTooManyOptionsToBind
	}
	stub := &clientStubMountTable{client: client, name: name}
	stub.Globable_InternalNoTagGetter, _ = BindGlobable(name, client)

	return stub, nil
}

// NewServerMountTable creates a new server stub.
//
// It takes a regular server implementing the MountTableService
// interface, and returns a new server stub.
func NewServerMountTable(server MountTableService) interface{} {
	return &ServerStubMountTable{
		ServerStubGlobable: *NewServerGlobable(server).(*ServerStubGlobable),
		service:            server,
	}
}

// clientStubMountTable implements MountTable.
type clientStubMountTable struct {
	Globable_InternalNoTagGetter

	client _gen_ipc.Client
	name   string
}

func (c *clientStubMountTable) GetMethodTags(method string) []interface{} {
	return GetMountTableMethodTags(method)
}

func (__gen_c *clientStubMountTable) Mount(Server string, TTL uint32, opts ..._gen_ipc.ClientCallOpt) (err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Mount", []interface{}{Server, TTL}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubMountTable) Unmount(Server string, opts ..._gen_ipc.ClientCallOpt) (err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Unmount", []interface{}{Server}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubMountTable) Link(LinkName string, opts ..._gen_ipc.ClientCallOpt) (err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Link", []interface{}{LinkName}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubMountTable) Unlink(opts ..._gen_ipc.ClientCallOpt) (err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Unlink", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubMountTable) ResolveStep(opts ..._gen_ipc.ClientCallOpt) (Servers []MountedServer, Suffix string, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "ResolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&Servers, &Suffix, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c *clientStubMountTable) UnresolveStep(opts ..._gen_ipc.ClientCallOpt) (reply []string, err error) {
	var call _gen_ipc.ClientCall
	if call, err = c.client.StartCall(c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubMountTable wraps a server that implements
// MountTableService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubMountTable struct {
	ServerStubGlobable

	service MountTableService
}

func (s *ServerStubMountTable) GetMethodTags(method string) []interface{} {
	return GetMountTableMethodTags(method)
}

func (s *ServerStubMountTable) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Link"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "LinkName", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Mount"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Server", Type: 3},
			{Name: "TTL", Type: 52},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["ResolveStep"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "Servers", Type: 67},
			{Name: "Suffix", Type: 3},
			{Name: "Error", Type: 65},
		},
	}
	result.Methods["Unlink"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Unmount"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "Server", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}

	result.TypeDefs = []_gen_idl.AnyData{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "Server"},
				_gen_wiretype.FieldType{Type: 0x34, Name: "TTL"},
			},
			"MountedServer", []string(nil)},
		_gen_wiretype.SliceType{Elem: 0x42, Name: "", Tags: []string(nil)}}
	var ss _gen_ipc.ServiceSignature
	var firstAdded int
	ss, _ = s.ServerStubGlobable.Signature(call)
	firstAdded = len(result.TypeDefs)
	for k, v := range ss.Methods {
		for i, _ := range v.InArgs {
			if v.InArgs[i].Type >= _gen_wiretype.TypeIDFirst {
				v.InArgs[i].Type += _gen_wiretype.TypeID(firstAdded)
			}
		}
		for i, _ := range v.OutArgs {
			if v.OutArgs[i].Type >= _gen_wiretype.TypeIDFirst {
				v.OutArgs[i].Type += _gen_wiretype.TypeID(firstAdded)
			}
		}
		if v.InStream >= _gen_wiretype.TypeIDFirst {
			v.InStream += _gen_wiretype.TypeID(firstAdded)
		}
		if v.OutStream >= _gen_wiretype.TypeIDFirst {
			v.OutStream += _gen_wiretype.TypeID(firstAdded)
		}
		result.Methods[k] = v
	}
	//TODO(bprosnitz) combine type definitions from embeded interfaces in a way that doesn't cause duplication.
	for _, d := range ss.TypeDefs {
		switch wt := d.(type) {
		case _gen_wiretype.SliceType:
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.ArrayType:
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.MapType:
			if wt.Key >= _gen_wiretype.TypeIDFirst {
				wt.Key += _gen_wiretype.TypeID(firstAdded)
			}
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.StructType:
			for _, fld := range wt.Fields {
				if fld.Type >= _gen_wiretype.TypeIDFirst {
					fld.Type += _gen_wiretype.TypeID(firstAdded)
				}
			}
			d = wt
		}
		result.TypeDefs = append(result.TypeDefs, d)
	}

	return result, nil
}

func (s *ServerStubMountTable) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := s.service.(_gen_ipc.Unresolver); ok {
		return unresolver.UnresolveStep(call)
	}
	if call.Server() == nil {
		return
	}
	var published []string
	if published, err = call.Server().Published(); err != nil || published == nil {
		return
	}
	reply = make([]string, len(published))
	for i, p := range published {
		reply[i] = _gen_naming.Join(p, call.Name())
	}
	return
}

func (__gen_s *ServerStubMountTable) Mount(call _gen_ipc.ServerCall, Server string, TTL uint32) (err error) {
	err = __gen_s.service.Mount(call, Server, TTL)
	return
}

func (__gen_s *ServerStubMountTable) Unmount(call _gen_ipc.ServerCall, Server string) (err error) {
	err = __gen_s.service.Unmount(call, Server)
	return
}

func (__gen_s *ServerStubMountTable) Link(call _gen_ipc.ServerCall, LinkName string) (err error) {
	err = __gen_s.service.Link(call, LinkName)
	return
}

func (__gen_s *ServerStubMountTable) Unlink(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.Unlink(call)
	return
}

func (__gen_s *ServerStubMountTable) ResolveStep(call _gen_ipc.ServerCall) (Servers []MountedServer, Suffix string, err error) {
	Servers, Suffix, err = __gen_s.service.ResolveStep(call)
	return
}

func GetMountTableMethodTags(method string) []interface{} {
	if resp := GetGlobableMethodTags(method); resp != nil {
		return resp
	}
	switch method {
	case "Mount":
		return []interface{}{}
	case "Unmount":
		return []interface{}{}
	case "Link":
		return []interface{}{}
	case "Unlink":
		return []interface{}{}
	case "ResolveStep":
		return []interface{}{}
	default:
		return nil
	}
}
