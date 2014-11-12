// This file was auto-generated by the veyron vdl tool.
// Source: service.vdl

// Package access defines the service for dynamic access control
// in Veyron.  Examples: "allow app to read this photo", "prevent user
// from modifying this file".
//
// Target Developers
//
// Developers creating functionality to share data or services between
// multiple users/devices/apps.
//
// Overview
//
// Every Veyron object supports GetACL and SetACL methods.  An ACL (Access
// Control List) contains blessings, groups, and the labels that principals
// with these blessings and group memberships can access for that object.
//
// An object can have multiple names, so GetACL and SetACL can be invoked on
// any of these names, but the object itself has a single ACL.
//
// SetACL completely replaces the ACL. To perform an atomic read-modify-write
// of the ACL, use the etag parameter.
//   client := access.ObjectClient(name)
//   for {
//     acl, etag, err := client.GetACL()
//     if err != nil {
//       return err
//     }
//     // Add newLabel to the LabelSet.
//     // TODO(kash): Update when we switch labels to strings instead of ints.
//     acl.In[newPattern] = acl.In[newPattern] | newLabel
//     // Use the same etag with the modified acl to ensure that no other client
//     // has modified the acl since GetACL returned.
//     if err := client.SetACL(acl, etag); err != nil {
//       if verror.Is(err, access.ErrBadEtag) {
//         // Another client replaced the ACL after our GetACL returned.
//         // Try again.
//         continue
//       }
//       return err
//     }
//   }
//
// Conventions
//
// Service implementors should follow the conventions below to be consistent
// with other parts of Veyron and with each other.
//
// All methods that create an object (e.g. Put, Mount, Link) should take an
// optional ACL parameter.  If the ACL is not specified, the new object, O,
// copies its ACL from the parent.  Subsequent changes to the parent ACL are
// not automatically propagated to O.  Instead, a client library could do
// recursive ACL changes if desired.
//
// security.ResolveLabel is required on all components of a name, except the
// last one, in order to access the object referenced by that name.  For
// example, for principal P to access the name "a/b/c", P must have resolve
// access to "a" and "a/b".
//
// security.ResolveLabel means that a principal can traverse that component of
// the name to access the child.  It does not give the principal permission to
// list the children via Glob or a similar method.  For example, a server
// might have an object named "home" with a child for each user of the system.
// If these users were allowed to list the contents of "home", they could
// discover the other users of the system.  That could be a privacy violation.
// Without ResolveLabel, every user of the system would need read access to
// "home" to access "home/<user>".  If the user called Glob("home/*"), it
// would then be up to the server to filter out the names that the user could
// not access.  That could be a very expensive operation if there were a lot
// of children of "home".  ResolveLabel protects these servers against
// potential denial of service attacks on these large, shared directories.
//
// Groups and blessings allow for sweeping access changes.  A group is
// suitable for saying that the same set of principals have access to a set of
// unrelated resources (e.g. docs, VMs, images).  See the Group API for a
// complete description.  A blessing is useful for controlling access to objects
// that are always accessed together.  For example, a document may have
// embedded images and comments, each with a unique name.  When accessing a
// document, the server would generate a blessing that the client would use to
// fetch the images and comments; the images and comments would have this
// blessed identity in their ACLs.  Changes to the document’s ACL are
// therefore “propagated” to the images and comments.
//
// Some services will want a concept of implicit access control.  They are
// free to implement this as is best for their service.  However, GetACL
// should respond with the correct ACL.  For example, a corporate file server
// would allow all employees to create their own directory and have full
// control within that directory.  Employees should not be allowed to modify
// other employee directories.  In other words, within the directory "home",
// employee E should be allowed to modify only "home/E".  The file server
// doesn't know the list of all employees a priori, so it uses an
// implementation-specific rule to map employee identities to their home
// directory.
package access

import (
	"veyron.io/veyron/veyron2/security"

	// The non-user imports are prefixed with "__" to prevent collisions.
	__veyron2 "veyron.io/veyron/veyron2"
	__context "veyron.io/veyron/veyron2/context"
	__ipc "veyron.io/veyron/veyron2/ipc"
	__vdlutil "veyron.io/veyron/veyron2/vdl/vdlutil"
	__verror "veyron.io/veyron/veyron2/verror"
	__wiretype "veyron.io/veyron/veyron2/wiretype"
)

// TODO(toddw): Remove this line once the new signature support is done.
// It corrects a bug where __wiretype is unused in VDL pacakges where only
// bootstrap types are used on interfaces.
const _ = __wiretype.TypeIDInvalid

// The etag passed to SetACL is invalid.  Likely, another client set
// the ACL already and invalidated the etag.  Use GetACL to fetch a
// fresh etag.
const ErrBadEtag = __verror.ID("veyron.io/veyron/veyron2/services/security/access.ErrBadEtag")

// The ACL is too big.  Use groups to represent large sets of principals.
const ErrTooBig = __verror.ID("veyron.io/veyron/veyron2/services/security/access.ErrTooBig")

// ObjectClientMethods is the client interface
// containing Object methods.
//
// Object provides access control for Veyron objects.
type ObjectClientMethods interface {
	// SetACL replaces the current ACL for an object.  etag allows for optional,
	// optimistic concurrency control.  If non-empty, etag's value must come
	// from GetACL.  If any client has successfully called SetACL in the
	// meantime, the etag will be stale and SetACL will fail.
	//
	// ACL objects are expected to be small.  It is up to the implementation to
	// define the exact limit, though it should probably be around 100KB.  Large
	// lists of principals should use the Group API or blessings.
	//
	// There is some ambiguity when calling SetACL on a mount point.  Does it
	// affect the mount itself or does it affect the service endpoint that the
	// mount points to?  The chosen behavior is that it affects the service
	// endpoint.  To modify the mount point's ACL, use ResolveToMountTable
	// to get an endpoint and call SetACL on that.  This means that clients
	// must know when a name refers to a mount point to change its ACL.
	SetACL(ctx __context.T, acl security.ACL, etag string, opts ...__ipc.CallOpt) error
	// GetACL returns the complete, current ACL for an object.  The returned etag
	// can be passed to a subsequent call to SetACL for optimistic concurrency
	// control. A successful call to SetACL will invalidate etag, and the client
	// must call GetACL again to get the current etag.
	GetACL(__context.T, ...__ipc.CallOpt) (acl security.ACL, etag string, err error)
}

// ObjectClientStub adds universal methods to ObjectClientMethods.
type ObjectClientStub interface {
	ObjectClientMethods
	__ipc.UniversalServiceMethods
}

// ObjectClient returns a client stub for Object.
func ObjectClient(name string, opts ...__ipc.BindOpt) ObjectClientStub {
	var client __ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(__ipc.Client); ok {
			client = clientOpt
		}
	}
	return implObjectClientStub{name, client}
}

type implObjectClientStub struct {
	name   string
	client __ipc.Client
}

func (c implObjectClientStub) c(ctx __context.T) __ipc.Client {
	if c.client != nil {
		return c.client
	}
	return __veyron2.RuntimeFromContext(ctx).Client()
}

func (c implObjectClientStub) SetACL(ctx __context.T, i0 security.ACL, i1 string, opts ...__ipc.CallOpt) (err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "SetACL", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c implObjectClientStub) GetACL(ctx __context.T, opts ...__ipc.CallOpt) (o0 security.ACL, o1 string, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "GetACL", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &o1, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implObjectClientStub) Signature(ctx __context.T, opts ...__ipc.CallOpt) (o0 __ipc.ServiceSignature, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implObjectClientStub) GetMethodTags(ctx __context.T, method string, opts ...__ipc.CallOpt) (o0 []interface{}, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// ObjectServerMethods is the interface a server writer
// implements for Object.
//
// Object provides access control for Veyron objects.
type ObjectServerMethods interface {
	// SetACL replaces the current ACL for an object.  etag allows for optional,
	// optimistic concurrency control.  If non-empty, etag's value must come
	// from GetACL.  If any client has successfully called SetACL in the
	// meantime, the etag will be stale and SetACL will fail.
	//
	// ACL objects are expected to be small.  It is up to the implementation to
	// define the exact limit, though it should probably be around 100KB.  Large
	// lists of principals should use the Group API or blessings.
	//
	// There is some ambiguity when calling SetACL on a mount point.  Does it
	// affect the mount itself or does it affect the service endpoint that the
	// mount points to?  The chosen behavior is that it affects the service
	// endpoint.  To modify the mount point's ACL, use ResolveToMountTable
	// to get an endpoint and call SetACL on that.  This means that clients
	// must know when a name refers to a mount point to change its ACL.
	SetACL(ctx __ipc.ServerContext, acl security.ACL, etag string) error
	// GetACL returns the complete, current ACL for an object.  The returned etag
	// can be passed to a subsequent call to SetACL for optimistic concurrency
	// control. A successful call to SetACL will invalidate etag, and the client
	// must call GetACL again to get the current etag.
	GetACL(__ipc.ServerContext) (acl security.ACL, etag string, err error)
}

// ObjectServerStubMethods is the server interface containing
// Object methods, as expected by ipc.Server.
// There is no difference between this interface and ObjectServerMethods
// since there are no streaming methods.
type ObjectServerStubMethods ObjectServerMethods

// ObjectServerStub adds universal methods to ObjectServerStubMethods.
type ObjectServerStub interface {
	ObjectServerStubMethods
	// GetMethodTags will be replaced with DescribeInterfaces.
	GetMethodTags(ctx __ipc.ServerContext, method string) ([]interface{}, error)
	// Signature will be replaced with DescribeInterfaces.
	Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error)
}

// ObjectServer returns a server stub for Object.
// It converts an implementation of ObjectServerMethods into
// an object that may be used by ipc.Server.
func ObjectServer(impl ObjectServerMethods) ObjectServerStub {
	stub := implObjectServerStub{
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

type implObjectServerStub struct {
	impl ObjectServerMethods
	gs   *__ipc.GlobState
}

func (s implObjectServerStub) SetACL(ctx __ipc.ServerContext, i0 security.ACL, i1 string) error {
	return s.impl.SetACL(ctx, i0, i1)
}

func (s implObjectServerStub) GetACL(ctx __ipc.ServerContext) (security.ACL, string, error) {
	return s.impl.GetACL(ctx)
}

func (s implObjectServerStub) VGlob() *__ipc.GlobState {
	return s.gs
}

func (s implObjectServerStub) GetMethodTags(ctx __ipc.ServerContext, method string) ([]interface{}, error) {
	// TODO(toddw): Replace with new DescribeInterfaces implementation.
	switch method {
	case "SetACL":
		return []interface{}{security.Label(8)}, nil
	case "GetACL":
		return []interface{}{security.Label(8)}, nil
	default:
		return nil, nil
	}
}

func (s implObjectServerStub) Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error) {
	// TODO(toddw) Replace with new DescribeInterfaces implementation.
	result := __ipc.ServiceSignature{Methods: make(map[string]__ipc.MethodSignature)}
	result.Methods["GetACL"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{},
		OutArgs: []__ipc.MethodArgument{
			{Name: "acl", Type: 69},
			{Name: "etag", Type: 3},
			{Name: "err", Type: 70},
		},
	}
	result.Methods["SetACL"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "acl", Type: 69},
			{Name: "etag", Type: 3},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 70},
		},
	}

	result.TypeDefs = []__vdlutil.Any{
		__wiretype.NamedPrimitiveType{Type: 0x3, Name: "veyron.io/veyron/veyron2/security.BlessingPattern", Tags: []string(nil)}, __wiretype.NamedPrimitiveType{Type: 0x34, Name: "veyron.io/veyron/veyron2/security.LabelSet", Tags: []string(nil)}, __wiretype.MapType{Key: 0x41, Elem: 0x42, Name: "", Tags: []string(nil)}, __wiretype.MapType{Key: 0x3, Elem: 0x42, Name: "", Tags: []string(nil)}, __wiretype.StructType{
			[]__wiretype.FieldType{
				__wiretype.FieldType{Type: 0x43, Name: "In"},
				__wiretype.FieldType{Type: 0x44, Name: "NotIn"},
			},
			"veyron.io/veyron/veyron2/security.ACL", []string(nil)},
		__wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}
