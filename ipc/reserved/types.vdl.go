// This file was auto-generated by the vanadium vdl tool.
// Source: types.vdl

package reserved

import (
	// VDL system imports
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/verror"
)

var (
	// GlobMaxRecursionReached indicates that the Glob request exceeded the
	// max recursion level.
	ErrGlobMaxRecursionReached = verror.Register("v.io/v23/ipc/reserved.GlobMaxRecursionReached", verror.NoRetry, "{1:}{2:} max recursion level reached{:_}")
	// GlobMatchesOmitted indicates that some of the Glob results might
	// have been omitted due to access restrictions.
	ErrGlobMatchesOmitted = verror.Register("v.io/v23/ipc/reserved.GlobMatchesOmitted", verror.NoRetry, "{1:}{2:} some matches might have been omitted")
	// GlobNotImplemented indicates that Glob is not implemented by the
	// object.
	ErrGlobNotImplemented = verror.Register("v.io/v23/ipc/reserved.GlobNotImplemented", verror.NoRetry, "{1:}{2:} Glob not implemented")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrGlobMaxRecursionReached.ID), "{1:}{2:} max recursion level reached{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrGlobMatchesOmitted.ID), "{1:}{2:} some matches might have been omitted")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrGlobNotImplemented.ID), "{1:}{2:} Glob not implemented")
}

// NewErrGlobMaxRecursionReached returns an error with the ErrGlobMaxRecursionReached ID.
func NewErrGlobMaxRecursionReached(ctx *context.T) error {
	return verror.New(ErrGlobMaxRecursionReached, ctx)
}

// NewErrGlobMatchesOmitted returns an error with the ErrGlobMatchesOmitted ID.
func NewErrGlobMatchesOmitted(ctx *context.T) error {
	return verror.New(ErrGlobMatchesOmitted, ctx)
}

// NewErrGlobNotImplemented returns an error with the ErrGlobNotImplemented ID.
func NewErrGlobNotImplemented(ctx *context.T) error {
	return verror.New(ErrGlobNotImplemented, ctx)
}
