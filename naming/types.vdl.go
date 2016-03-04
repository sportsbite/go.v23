// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: types.vdl

package naming

import (
	// VDL system imports
	"v.io/v23/vdl"
	"v.io/v23/verror"

	// VDL user imports
	"v.io/v23/vdlroot/time"
)

// MountFlag is a bit mask of options to the mount call.
type MountFlag uint32

func (MountFlag) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/naming.MountFlag"`
}) {
}

func (m MountFlag) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromUint(uint64(m), __VDLType_types_v_io_v23_naming_MountFlag); err != nil {
		return err
	}
	return nil
}

func (m MountFlag) MakeVDLTarget() vdl.Target {
	return nil
}

func (m MountFlag) IsZero() bool {

	var1 := (m == MountFlag(0))
	return var1
}

// MountedServer represents a server mounted on a specific name.
type MountedServer struct {
	// Server is the OA that's mounted.
	Server string
	// Deadline before the mount entry expires.
	Deadline time.Deadline
}

func (MountedServer) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/naming.MountedServer"`
}) {
}

func (m *MountedServer) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	__VDLEnsureNativeBuilt_types()
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	var2 := (m.Server == "")
	if !var2 {
		keyTarget3, fieldTarget4, err := fieldsTarget1.StartField("Server")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {
			if err := fieldTarget4.FromString(string(m.Server), vdl.StringType); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget3, fieldTarget4); err != nil {
				return err
			}
		}
	}
	var wireValue5 time.WireDeadline
	if err := time.WireDeadlineFromNative(&wireValue5, m.Deadline); err != nil {
		return err
	}

	var6 := wireValue5.IsZero()
	if !var6 {
		keyTarget7, fieldTarget8, err := fieldsTarget1.StartField("Deadline")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {

			if err := wireValue5.FillVDLTarget(fieldTarget8, __VDLType_types_time_WireDeadline); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget7, fieldTarget8); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *MountedServer) MakeVDLTarget() vdl.Target {
	return nil
}

func (m *MountedServer) IsZero() bool {

	var1 := (*m == MountedServer{})
	return var1
}

// MountEntry represents a given name mounted in the mounttable.
type MountEntry struct {
	// Name is the mounted name.
	Name string
	// Servers (if present) specifies the mounted names.
	Servers []MountedServer
	// ServesMountTable is true if the servers represent mount tables.
	ServesMountTable bool
	// IsLeaf is true if this entry represents a leaf object.
	IsLeaf bool
}

func (MountEntry) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/naming.MountEntry"`
}) {
}

func (m *MountEntry) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	__VDLEnsureNativeBuilt_types()
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	var2 := (m.Name == "")
	if !var2 {
		keyTarget3, fieldTarget4, err := fieldsTarget1.StartField("Name")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {
			if err := fieldTarget4.FromString(string(m.Name), vdl.StringType); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget3, fieldTarget4); err != nil {
				return err
			}
		}
	}
	var var5 bool
	if len(m.Servers) == 0 {
		var5 = true
	}
	if !var5 {
		keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("Servers")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {

			listTarget8, err := fieldTarget7.StartList(__VDLTypetypes2, len(m.Servers))
			if err != nil {
				return err
			}
			for i, elem10 := range m.Servers {
				elemTarget9, err := listTarget8.StartElem(i)
				if err != nil {
					return err
				}

				if err := elem10.FillVDLTarget(elemTarget9, __VDLType_types_v_io_v23_naming_MountedServer); err != nil {
					return err
				}
				if err := listTarget8.FinishElem(elemTarget9); err != nil {
					return err
				}
			}
			if err := fieldTarget7.FinishList(listTarget8); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
				return err
			}
		}
	}
	var11 := (m.ServesMountTable == false)
	if !var11 {
		keyTarget12, fieldTarget13, err := fieldsTarget1.StartField("ServesMountTable")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {
			if err := fieldTarget13.FromBool(bool(m.ServesMountTable), vdl.BoolType); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget12, fieldTarget13); err != nil {
				return err
			}
		}
	}
	var14 := (m.IsLeaf == false)
	if !var14 {
		keyTarget15, fieldTarget16, err := fieldsTarget1.StartField("IsLeaf")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {
			if err := fieldTarget16.FromBool(bool(m.IsLeaf), vdl.BoolType); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget15, fieldTarget16); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *MountEntry) MakeVDLTarget() vdl.Target {
	return nil
}

func (m *MountEntry) IsZero() bool {

	var1 := true
	var2 := (m.Name == "")
	var1 = var1 && var2
	var var3 bool
	if len(m.Servers) == 0 {
		var3 = true
	}
	var1 = var1 && var3
	var4 := (m.ServesMountTable == false)
	var1 = var1 && var4
	var5 := (m.IsLeaf == false)
	var1 = var1 && var5
	return var1
}

// GlobError is returned by namespace.Glob to indicate a subtree of the namespace
// that could not be traversed.
type GlobError struct {
	// Root of the subtree.
	Name string
	// The error that occurred fulfilling the request.
	Error error
}

func (GlobError) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/naming.GlobError"`
}) {
}

func (m *GlobError) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	if __VDLType_types_v_io_v23_naming_GlobError == nil || __VDLTypetypes3 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	var2 := (m.Name == "")
	if !var2 {
		keyTarget3, fieldTarget4, err := fieldsTarget1.StartField("Name")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {
			if err := fieldTarget4.FromString(string(m.Name), vdl.StringType); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget3, fieldTarget4); err != nil {
				return err
			}
		}
	}
	var5 := (m.Error == (error)(nil))
	if !var5 {
		keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("Error")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {

			if m.Error == nil {
				if err := fieldTarget7.FromNil(vdl.ErrorType); err != nil {
					return err
				}
			} else {
				var wireError8 vdl.WireError
				if err := verror.WireFromNative(&wireError8, m.Error); err != nil {
					return err
				}
				if err := wireError8.FillVDLTarget(fieldTarget7, vdl.ErrorType); err != nil {
					return err
				}

			}
			if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *GlobError) MakeVDLTarget() vdl.Target {
	return nil
}

func (m *GlobError) IsZero() bool {

	var1 := true
	var2 := (m.Name == "")
	var1 = var1 && var2
	var3 := (m.Error == (error)(nil))
	var1 = var1 && var3
	return var1
}

type (
	// GlobReply represents any single field of the GlobReply union type.
	//
	// GlobReply is the data type returned by Glob__.
	GlobReply interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the GlobReply union type.
		__VDLReflect(__GlobReplyReflect)
		FillVDLTarget(vdl.Target, *vdl.Type) error
		IsZero() bool
	}
	// GlobReplyEntry represents field Entry of the GlobReply union type.
	GlobReplyEntry struct{ Value MountEntry }
	// GlobReplyError represents field Error of the GlobReply union type.
	GlobReplyError struct{ Value GlobError }
	// __GlobReplyReflect describes the GlobReply union type.
	__GlobReplyReflect struct {
		Name  string `vdl:"v.io/v23/naming.GlobReply"`
		Type  GlobReply
		Union struct {
			Entry GlobReplyEntry
			Error GlobReplyError
		}
	}
)

func (x GlobReplyEntry) Index() int                      { return 0 }
func (x GlobReplyEntry) Interface() interface{}          { return x.Value }
func (x GlobReplyEntry) Name() string                    { return "Entry" }
func (x GlobReplyEntry) __VDLReflect(__GlobReplyReflect) {}

func (m GlobReplyEntry) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	fieldsTarget1, err := t.StartFields(__VDLType_types_v_io_v23_naming_GlobReply)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Entry")
	if err != nil {
		return err
	}

	if err := m.Value.FillVDLTarget(fieldTarget3, __VDLType_types_v_io_v23_naming_MountEntry); err != nil {
		return err
	}
	if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
		return err
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}

	return nil
}

func (m GlobReplyEntry) MakeVDLTarget() vdl.Target {
	return nil
}

func (m GlobReplyEntry) IsZero() bool {

	var2 := m.Value.IsZero()
	return var2
}

func (x GlobReplyError) Index() int                      { return 1 }
func (x GlobReplyError) Interface() interface{}          { return x.Value }
func (x GlobReplyError) Name() string                    { return "Error" }
func (x GlobReplyError) __VDLReflect(__GlobReplyReflect) {}

func (m GlobReplyError) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	fieldsTarget1, err := t.StartFields(__VDLType_types_v_io_v23_naming_GlobReply)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Error")
	if err != nil {
		return err
	}

	if err := m.Value.FillVDLTarget(fieldTarget3, __VDLType_types_v_io_v23_naming_GlobError); err != nil {
		return err
	}
	if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
		return err
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}

	return nil
}

func (m GlobReplyError) MakeVDLTarget() vdl.Target {
	return nil
}

func (m GlobReplyError) IsZero() bool {

	unionField2 := false
	return unionField2
}

type (
	// GlobChildrenReply represents any single field of the GlobChildrenReply union type.
	//
	// GlobChildrenReply is the data type returned by GlobChildren__.
	GlobChildrenReply interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the GlobChildrenReply union type.
		__VDLReflect(__GlobChildrenReplyReflect)
		FillVDLTarget(vdl.Target, *vdl.Type) error
		IsZero() bool
	}
	// GlobChildrenReplyName represents field Name of the GlobChildrenReply union type.
	GlobChildrenReplyName struct{ Value string }
	// GlobChildrenReplyError represents field Error of the GlobChildrenReply union type.
	GlobChildrenReplyError struct{ Value GlobError }
	// __GlobChildrenReplyReflect describes the GlobChildrenReply union type.
	__GlobChildrenReplyReflect struct {
		Name  string `vdl:"v.io/v23/naming.GlobChildrenReply"`
		Type  GlobChildrenReply
		Union struct {
			Name  GlobChildrenReplyName
			Error GlobChildrenReplyError
		}
	}
)

func (x GlobChildrenReplyName) Index() int                              { return 0 }
func (x GlobChildrenReplyName) Interface() interface{}                  { return x.Value }
func (x GlobChildrenReplyName) Name() string                            { return "Name" }
func (x GlobChildrenReplyName) __VDLReflect(__GlobChildrenReplyReflect) {}

func (m GlobChildrenReplyName) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	fieldsTarget1, err := t.StartFields(__VDLType_types_v_io_v23_naming_GlobChildrenReply)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Name")
	if err != nil {
		return err
	}
	if err := fieldTarget3.FromString(string(m.Value), vdl.StringType); err != nil {
		return err
	}
	if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
		return err
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}

	return nil
}

func (m GlobChildrenReplyName) MakeVDLTarget() vdl.Target {
	return nil
}

func (m GlobChildrenReplyName) IsZero() bool {

	var2 := (m.Value == "")
	return var2
}

func (x GlobChildrenReplyError) Index() int                              { return 1 }
func (x GlobChildrenReplyError) Interface() interface{}                  { return x.Value }
func (x GlobChildrenReplyError) Name() string                            { return "Error" }
func (x GlobChildrenReplyError) __VDLReflect(__GlobChildrenReplyReflect) {}

func (m GlobChildrenReplyError) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	fieldsTarget1, err := t.StartFields(__VDLType_types_v_io_v23_naming_GlobChildrenReply)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Error")
	if err != nil {
		return err
	}

	if err := m.Value.FillVDLTarget(fieldTarget3, __VDLType_types_v_io_v23_naming_GlobError); err != nil {
		return err
	}
	if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
		return err
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}

	return nil
}

func (m GlobChildrenReplyError) MakeVDLTarget() vdl.Target {
	return nil
}

func (m GlobChildrenReplyError) IsZero() bool {

	unionField2 := false
	return unionField2
}

func init() {
	vdl.Register((*MountFlag)(nil))
	vdl.Register((*MountedServer)(nil))
	vdl.Register((*MountEntry)(nil))
	vdl.Register((*GlobError)(nil))
	vdl.Register((*GlobReply)(nil))
	vdl.Register((*GlobChildrenReply)(nil))
}

var __VDLTypetypes3 *vdl.Type = vdl.TypeOf((*GlobError)(nil))
var __VDLTypetypes1 *vdl.Type

func __VDLTypetypes1_gen() *vdl.Type {
	__VDLTypetypes1Builder := vdl.TypeBuilder{}

	__VDLTypetypes11 := __VDLTypetypes1Builder.Optional()
	__VDLTypetypes12 := __VDLTypetypes1Builder.Struct()
	__VDLTypetypes13 := __VDLTypetypes1Builder.Named("v.io/v23/naming.MountEntry").AssignBase(__VDLTypetypes12)
	__VDLTypetypes14 := vdl.StringType
	__VDLTypetypes12.AppendField("Name", __VDLTypetypes14)
	__VDLTypetypes15 := __VDLTypetypes1Builder.List()
	__VDLTypetypes16 := __VDLTypetypes1Builder.Struct()
	__VDLTypetypes17 := __VDLTypetypes1Builder.Named("v.io/v23/naming.MountedServer").AssignBase(__VDLTypetypes16)
	__VDLTypetypes16.AppendField("Server", __VDLTypetypes14)
	__VDLTypetypes18 := __VDLTypetypes1Builder.Struct()
	__VDLTypetypes19 := __VDLTypetypes1Builder.Named("time.WireDeadline").AssignBase(__VDLTypetypes18)
	__VDLTypetypes110 := __VDLTypetypes1Builder.Struct()
	__VDLTypetypes111 := __VDLTypetypes1Builder.Named("time.Duration").AssignBase(__VDLTypetypes110)
	__VDLTypetypes112 := vdl.Int64Type
	__VDLTypetypes110.AppendField("Seconds", __VDLTypetypes112)
	__VDLTypetypes113 := vdl.Int32Type
	__VDLTypetypes110.AppendField("Nanos", __VDLTypetypes113)
	__VDLTypetypes18.AppendField("FromNow", __VDLTypetypes111)
	__VDLTypetypes114 := vdl.BoolType
	__VDLTypetypes18.AppendField("NoDeadline", __VDLTypetypes114)
	__VDLTypetypes16.AppendField("Deadline", __VDLTypetypes19)
	__VDLTypetypes15.AssignElem(__VDLTypetypes17)
	__VDLTypetypes12.AppendField("Servers", __VDLTypetypes15)
	__VDLTypetypes12.AppendField("ServesMountTable", __VDLTypetypes114)
	__VDLTypetypes12.AppendField("IsLeaf", __VDLTypetypes114)
	__VDLTypetypes11.AssignElem(__VDLTypetypes13)
	__VDLTypetypes1Builder.Build()
	__VDLTypetypes1v, err := __VDLTypetypes11.Built()
	if err != nil {
		panic(err)
	}
	return __VDLTypetypes1v
}
func init() {
	__VDLTypetypes1 = __VDLTypetypes1_gen()
}

var __VDLTypetypes0 *vdl.Type

func __VDLTypetypes0_gen() *vdl.Type {
	__VDLTypetypes0Builder := vdl.TypeBuilder{}

	__VDLTypetypes01 := __VDLTypetypes0Builder.Optional()
	__VDLTypetypes02 := __VDLTypetypes0Builder.Struct()
	__VDLTypetypes03 := __VDLTypetypes0Builder.Named("v.io/v23/naming.MountedServer").AssignBase(__VDLTypetypes02)
	__VDLTypetypes04 := vdl.StringType
	__VDLTypetypes02.AppendField("Server", __VDLTypetypes04)
	__VDLTypetypes05 := __VDLTypetypes0Builder.Struct()
	__VDLTypetypes06 := __VDLTypetypes0Builder.Named("time.WireDeadline").AssignBase(__VDLTypetypes05)
	__VDLTypetypes07 := __VDLTypetypes0Builder.Struct()
	__VDLTypetypes08 := __VDLTypetypes0Builder.Named("time.Duration").AssignBase(__VDLTypetypes07)
	__VDLTypetypes09 := vdl.Int64Type
	__VDLTypetypes07.AppendField("Seconds", __VDLTypetypes09)
	__VDLTypetypes010 := vdl.Int32Type
	__VDLTypetypes07.AppendField("Nanos", __VDLTypetypes010)
	__VDLTypetypes05.AppendField("FromNow", __VDLTypetypes08)
	__VDLTypetypes011 := vdl.BoolType
	__VDLTypetypes05.AppendField("NoDeadline", __VDLTypetypes011)
	__VDLTypetypes02.AppendField("Deadline", __VDLTypetypes06)
	__VDLTypetypes01.AssignElem(__VDLTypetypes03)
	__VDLTypetypes0Builder.Build()
	__VDLTypetypes0v, err := __VDLTypetypes01.Built()
	if err != nil {
		panic(err)
	}
	return __VDLTypetypes0v
}
func init() {
	__VDLTypetypes0 = __VDLTypetypes0_gen()
}

var __VDLTypetypes2 *vdl.Type

func __VDLTypetypes2_gen() *vdl.Type {
	__VDLTypetypes2Builder := vdl.TypeBuilder{}

	__VDLTypetypes21 := __VDLTypetypes2Builder.List()
	__VDLTypetypes22 := __VDLTypetypes2Builder.Struct()
	__VDLTypetypes23 := __VDLTypetypes2Builder.Named("v.io/v23/naming.MountedServer").AssignBase(__VDLTypetypes22)
	__VDLTypetypes24 := vdl.StringType
	__VDLTypetypes22.AppendField("Server", __VDLTypetypes24)
	__VDLTypetypes25 := __VDLTypetypes2Builder.Struct()
	__VDLTypetypes26 := __VDLTypetypes2Builder.Named("time.WireDeadline").AssignBase(__VDLTypetypes25)
	__VDLTypetypes27 := __VDLTypetypes2Builder.Struct()
	__VDLTypetypes28 := __VDLTypetypes2Builder.Named("time.Duration").AssignBase(__VDLTypetypes27)
	__VDLTypetypes29 := vdl.Int64Type
	__VDLTypetypes27.AppendField("Seconds", __VDLTypetypes29)
	__VDLTypetypes210 := vdl.Int32Type
	__VDLTypetypes27.AppendField("Nanos", __VDLTypetypes210)
	__VDLTypetypes25.AppendField("FromNow", __VDLTypetypes28)
	__VDLTypetypes211 := vdl.BoolType
	__VDLTypetypes25.AppendField("NoDeadline", __VDLTypetypes211)
	__VDLTypetypes22.AppendField("Deadline", __VDLTypetypes26)
	__VDLTypetypes21.AssignElem(__VDLTypetypes23)
	__VDLTypetypes2Builder.Build()
	__VDLTypetypes2v, err := __VDLTypetypes21.Built()
	if err != nil {
		panic(err)
	}
	return __VDLTypetypes2v
}
func init() {
	__VDLTypetypes2 = __VDLTypetypes2_gen()
}

var __VDLType_types_time_WireDeadline *vdl.Type

func __VDLType_types_time_WireDeadline_gen() *vdl.Type {
	__VDLType_types_time_WireDeadlineBuilder := vdl.TypeBuilder{}

	__VDLType_types_time_WireDeadline1 := __VDLType_types_time_WireDeadlineBuilder.Struct()
	__VDLType_types_time_WireDeadline2 := __VDLType_types_time_WireDeadlineBuilder.Named("time.WireDeadline").AssignBase(__VDLType_types_time_WireDeadline1)
	__VDLType_types_time_WireDeadline3 := __VDLType_types_time_WireDeadlineBuilder.Struct()
	__VDLType_types_time_WireDeadline4 := __VDLType_types_time_WireDeadlineBuilder.Named("time.Duration").AssignBase(__VDLType_types_time_WireDeadline3)
	__VDLType_types_time_WireDeadline5 := vdl.Int64Type
	__VDLType_types_time_WireDeadline3.AppendField("Seconds", __VDLType_types_time_WireDeadline5)
	__VDLType_types_time_WireDeadline6 := vdl.Int32Type
	__VDLType_types_time_WireDeadline3.AppendField("Nanos", __VDLType_types_time_WireDeadline6)
	__VDLType_types_time_WireDeadline1.AppendField("FromNow", __VDLType_types_time_WireDeadline4)
	__VDLType_types_time_WireDeadline7 := vdl.BoolType
	__VDLType_types_time_WireDeadline1.AppendField("NoDeadline", __VDLType_types_time_WireDeadline7)
	__VDLType_types_time_WireDeadlineBuilder.Build()
	__VDLType_types_time_WireDeadlinev, err := __VDLType_types_time_WireDeadline2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_types_time_WireDeadlinev
}
func init() {
	__VDLType_types_time_WireDeadline = __VDLType_types_time_WireDeadline_gen()
}

var __VDLType_types_v_io_v23_naming_GlobChildrenReply *vdl.Type = vdl.TypeOf(GlobChildrenReply(GlobChildrenReplyName{""}))
var __VDLType_types_v_io_v23_naming_GlobError *vdl.Type = vdl.TypeOf(GlobError{})
var __VDLType_types_v_io_v23_naming_GlobReply *vdl.Type

func __VDLType_types_v_io_v23_naming_GlobReply_gen() *vdl.Type {
	__VDLType_types_v_io_v23_naming_GlobReplyBuilder := vdl.TypeBuilder{}

	__VDLType_types_v_io_v23_naming_GlobReply1 := __VDLType_types_v_io_v23_naming_GlobReplyBuilder.Union()
	__VDLType_types_v_io_v23_naming_GlobReply2 := __VDLType_types_v_io_v23_naming_GlobReplyBuilder.Named("v.io/v23/naming.GlobReply").AssignBase(__VDLType_types_v_io_v23_naming_GlobReply1)
	__VDLType_types_v_io_v23_naming_GlobReply3 := __VDLType_types_v_io_v23_naming_GlobReplyBuilder.Struct()
	__VDLType_types_v_io_v23_naming_GlobReply4 := __VDLType_types_v_io_v23_naming_GlobReplyBuilder.Named("v.io/v23/naming.MountEntry").AssignBase(__VDLType_types_v_io_v23_naming_GlobReply3)
	__VDLType_types_v_io_v23_naming_GlobReply5 := vdl.StringType
	__VDLType_types_v_io_v23_naming_GlobReply3.AppendField("Name", __VDLType_types_v_io_v23_naming_GlobReply5)
	__VDLType_types_v_io_v23_naming_GlobReply6 := __VDLType_types_v_io_v23_naming_GlobReplyBuilder.List()
	__VDLType_types_v_io_v23_naming_GlobReply7 := __VDLType_types_v_io_v23_naming_GlobReplyBuilder.Struct()
	__VDLType_types_v_io_v23_naming_GlobReply8 := __VDLType_types_v_io_v23_naming_GlobReplyBuilder.Named("v.io/v23/naming.MountedServer").AssignBase(__VDLType_types_v_io_v23_naming_GlobReply7)
	__VDLType_types_v_io_v23_naming_GlobReply7.AppendField("Server", __VDLType_types_v_io_v23_naming_GlobReply5)
	__VDLType_types_v_io_v23_naming_GlobReply9 := __VDLType_types_v_io_v23_naming_GlobReplyBuilder.Struct()
	__VDLType_types_v_io_v23_naming_GlobReply10 := __VDLType_types_v_io_v23_naming_GlobReplyBuilder.Named("time.WireDeadline").AssignBase(__VDLType_types_v_io_v23_naming_GlobReply9)
	__VDLType_types_v_io_v23_naming_GlobReply11 := __VDLType_types_v_io_v23_naming_GlobReplyBuilder.Struct()
	__VDLType_types_v_io_v23_naming_GlobReply12 := __VDLType_types_v_io_v23_naming_GlobReplyBuilder.Named("time.Duration").AssignBase(__VDLType_types_v_io_v23_naming_GlobReply11)
	__VDLType_types_v_io_v23_naming_GlobReply13 := vdl.Int64Type
	__VDLType_types_v_io_v23_naming_GlobReply11.AppendField("Seconds", __VDLType_types_v_io_v23_naming_GlobReply13)
	__VDLType_types_v_io_v23_naming_GlobReply14 := vdl.Int32Type
	__VDLType_types_v_io_v23_naming_GlobReply11.AppendField("Nanos", __VDLType_types_v_io_v23_naming_GlobReply14)
	__VDLType_types_v_io_v23_naming_GlobReply9.AppendField("FromNow", __VDLType_types_v_io_v23_naming_GlobReply12)
	__VDLType_types_v_io_v23_naming_GlobReply15 := vdl.BoolType
	__VDLType_types_v_io_v23_naming_GlobReply9.AppendField("NoDeadline", __VDLType_types_v_io_v23_naming_GlobReply15)
	__VDLType_types_v_io_v23_naming_GlobReply7.AppendField("Deadline", __VDLType_types_v_io_v23_naming_GlobReply10)
	__VDLType_types_v_io_v23_naming_GlobReply6.AssignElem(__VDLType_types_v_io_v23_naming_GlobReply8)
	__VDLType_types_v_io_v23_naming_GlobReply3.AppendField("Servers", __VDLType_types_v_io_v23_naming_GlobReply6)
	__VDLType_types_v_io_v23_naming_GlobReply3.AppendField("ServesMountTable", __VDLType_types_v_io_v23_naming_GlobReply15)
	__VDLType_types_v_io_v23_naming_GlobReply3.AppendField("IsLeaf", __VDLType_types_v_io_v23_naming_GlobReply15)
	__VDLType_types_v_io_v23_naming_GlobReply1.AppendField("Entry", __VDLType_types_v_io_v23_naming_GlobReply4)
	__VDLType_types_v_io_v23_naming_GlobReply16 := __VDLType_types_v_io_v23_naming_GlobReplyBuilder.Struct()
	__VDLType_types_v_io_v23_naming_GlobReply17 := __VDLType_types_v_io_v23_naming_GlobReplyBuilder.Named("v.io/v23/naming.GlobError").AssignBase(__VDLType_types_v_io_v23_naming_GlobReply16)
	__VDLType_types_v_io_v23_naming_GlobReply16.AppendField("Name", __VDLType_types_v_io_v23_naming_GlobReply5)
	__VDLType_types_v_io_v23_naming_GlobReply18 := __VDLType_types_v_io_v23_naming_GlobReplyBuilder.Optional()
	__VDLType_types_v_io_v23_naming_GlobReply19 := __VDLType_types_v_io_v23_naming_GlobReplyBuilder.Struct()
	__VDLType_types_v_io_v23_naming_GlobReply20 := __VDLType_types_v_io_v23_naming_GlobReplyBuilder.Named("error").AssignBase(__VDLType_types_v_io_v23_naming_GlobReply19)
	__VDLType_types_v_io_v23_naming_GlobReply19.AppendField("Id", __VDLType_types_v_io_v23_naming_GlobReply5)
	__VDLType_types_v_io_v23_naming_GlobReply21 := __VDLType_types_v_io_v23_naming_GlobReplyBuilder.Enum()
	__VDLType_types_v_io_v23_naming_GlobReply21.AppendLabel("NoRetry")
	__VDLType_types_v_io_v23_naming_GlobReply21.AppendLabel("RetryConnection")
	__VDLType_types_v_io_v23_naming_GlobReply21.AppendLabel("RetryRefetch")
	__VDLType_types_v_io_v23_naming_GlobReply21.AppendLabel("RetryBackoff")
	__VDLType_types_v_io_v23_naming_GlobReply19.AppendField("RetryCode", __VDLType_types_v_io_v23_naming_GlobReply21)
	__VDLType_types_v_io_v23_naming_GlobReply19.AppendField("Msg", __VDLType_types_v_io_v23_naming_GlobReply5)
	__VDLType_types_v_io_v23_naming_GlobReply22 := __VDLType_types_v_io_v23_naming_GlobReplyBuilder.List()
	__VDLType_types_v_io_v23_naming_GlobReply23 := vdl.AnyType
	__VDLType_types_v_io_v23_naming_GlobReply22.AssignElem(__VDLType_types_v_io_v23_naming_GlobReply23)
	__VDLType_types_v_io_v23_naming_GlobReply19.AppendField("ParamList", __VDLType_types_v_io_v23_naming_GlobReply22)
	__VDLType_types_v_io_v23_naming_GlobReply18.AssignElem(__VDLType_types_v_io_v23_naming_GlobReply20)
	__VDLType_types_v_io_v23_naming_GlobReply16.AppendField("Error", __VDLType_types_v_io_v23_naming_GlobReply18)
	__VDLType_types_v_io_v23_naming_GlobReply1.AppendField("Error", __VDLType_types_v_io_v23_naming_GlobReply17)
	__VDLType_types_v_io_v23_naming_GlobReplyBuilder.Build()
	__VDLType_types_v_io_v23_naming_GlobReplyv, err := __VDLType_types_v_io_v23_naming_GlobReply2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_types_v_io_v23_naming_GlobReplyv
}
func init() {
	__VDLType_types_v_io_v23_naming_GlobReply = __VDLType_types_v_io_v23_naming_GlobReply_gen()
}

var __VDLType_types_v_io_v23_naming_MountEntry *vdl.Type

func __VDLType_types_v_io_v23_naming_MountEntry_gen() *vdl.Type {
	__VDLType_types_v_io_v23_naming_MountEntryBuilder := vdl.TypeBuilder{}

	__VDLType_types_v_io_v23_naming_MountEntry1 := __VDLType_types_v_io_v23_naming_MountEntryBuilder.Struct()
	__VDLType_types_v_io_v23_naming_MountEntry2 := __VDLType_types_v_io_v23_naming_MountEntryBuilder.Named("v.io/v23/naming.MountEntry").AssignBase(__VDLType_types_v_io_v23_naming_MountEntry1)
	__VDLType_types_v_io_v23_naming_MountEntry3 := vdl.StringType
	__VDLType_types_v_io_v23_naming_MountEntry1.AppendField("Name", __VDLType_types_v_io_v23_naming_MountEntry3)
	__VDLType_types_v_io_v23_naming_MountEntry4 := __VDLType_types_v_io_v23_naming_MountEntryBuilder.List()
	__VDLType_types_v_io_v23_naming_MountEntry5 := __VDLType_types_v_io_v23_naming_MountEntryBuilder.Struct()
	__VDLType_types_v_io_v23_naming_MountEntry6 := __VDLType_types_v_io_v23_naming_MountEntryBuilder.Named("v.io/v23/naming.MountedServer").AssignBase(__VDLType_types_v_io_v23_naming_MountEntry5)
	__VDLType_types_v_io_v23_naming_MountEntry5.AppendField("Server", __VDLType_types_v_io_v23_naming_MountEntry3)
	__VDLType_types_v_io_v23_naming_MountEntry7 := __VDLType_types_v_io_v23_naming_MountEntryBuilder.Struct()
	__VDLType_types_v_io_v23_naming_MountEntry8 := __VDLType_types_v_io_v23_naming_MountEntryBuilder.Named("time.WireDeadline").AssignBase(__VDLType_types_v_io_v23_naming_MountEntry7)
	__VDLType_types_v_io_v23_naming_MountEntry9 := __VDLType_types_v_io_v23_naming_MountEntryBuilder.Struct()
	__VDLType_types_v_io_v23_naming_MountEntry10 := __VDLType_types_v_io_v23_naming_MountEntryBuilder.Named("time.Duration").AssignBase(__VDLType_types_v_io_v23_naming_MountEntry9)
	__VDLType_types_v_io_v23_naming_MountEntry11 := vdl.Int64Type
	__VDLType_types_v_io_v23_naming_MountEntry9.AppendField("Seconds", __VDLType_types_v_io_v23_naming_MountEntry11)
	__VDLType_types_v_io_v23_naming_MountEntry12 := vdl.Int32Type
	__VDLType_types_v_io_v23_naming_MountEntry9.AppendField("Nanos", __VDLType_types_v_io_v23_naming_MountEntry12)
	__VDLType_types_v_io_v23_naming_MountEntry7.AppendField("FromNow", __VDLType_types_v_io_v23_naming_MountEntry10)
	__VDLType_types_v_io_v23_naming_MountEntry13 := vdl.BoolType
	__VDLType_types_v_io_v23_naming_MountEntry7.AppendField("NoDeadline", __VDLType_types_v_io_v23_naming_MountEntry13)
	__VDLType_types_v_io_v23_naming_MountEntry5.AppendField("Deadline", __VDLType_types_v_io_v23_naming_MountEntry8)
	__VDLType_types_v_io_v23_naming_MountEntry4.AssignElem(__VDLType_types_v_io_v23_naming_MountEntry6)
	__VDLType_types_v_io_v23_naming_MountEntry1.AppendField("Servers", __VDLType_types_v_io_v23_naming_MountEntry4)
	__VDLType_types_v_io_v23_naming_MountEntry1.AppendField("ServesMountTable", __VDLType_types_v_io_v23_naming_MountEntry13)
	__VDLType_types_v_io_v23_naming_MountEntry1.AppendField("IsLeaf", __VDLType_types_v_io_v23_naming_MountEntry13)
	__VDLType_types_v_io_v23_naming_MountEntryBuilder.Build()
	__VDLType_types_v_io_v23_naming_MountEntryv, err := __VDLType_types_v_io_v23_naming_MountEntry2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_types_v_io_v23_naming_MountEntryv
}
func init() {
	__VDLType_types_v_io_v23_naming_MountEntry = __VDLType_types_v_io_v23_naming_MountEntry_gen()
}

var __VDLType_types_v_io_v23_naming_MountFlag *vdl.Type = vdl.TypeOf(MountFlag(0))
var __VDLType_types_v_io_v23_naming_MountedServer *vdl.Type

func __VDLType_types_v_io_v23_naming_MountedServer_gen() *vdl.Type {
	__VDLType_types_v_io_v23_naming_MountedServerBuilder := vdl.TypeBuilder{}

	__VDLType_types_v_io_v23_naming_MountedServer1 := __VDLType_types_v_io_v23_naming_MountedServerBuilder.Struct()
	__VDLType_types_v_io_v23_naming_MountedServer2 := __VDLType_types_v_io_v23_naming_MountedServerBuilder.Named("v.io/v23/naming.MountedServer").AssignBase(__VDLType_types_v_io_v23_naming_MountedServer1)
	__VDLType_types_v_io_v23_naming_MountedServer3 := vdl.StringType
	__VDLType_types_v_io_v23_naming_MountedServer1.AppendField("Server", __VDLType_types_v_io_v23_naming_MountedServer3)
	__VDLType_types_v_io_v23_naming_MountedServer4 := __VDLType_types_v_io_v23_naming_MountedServerBuilder.Struct()
	__VDLType_types_v_io_v23_naming_MountedServer5 := __VDLType_types_v_io_v23_naming_MountedServerBuilder.Named("time.WireDeadline").AssignBase(__VDLType_types_v_io_v23_naming_MountedServer4)
	__VDLType_types_v_io_v23_naming_MountedServer6 := __VDLType_types_v_io_v23_naming_MountedServerBuilder.Struct()
	__VDLType_types_v_io_v23_naming_MountedServer7 := __VDLType_types_v_io_v23_naming_MountedServerBuilder.Named("time.Duration").AssignBase(__VDLType_types_v_io_v23_naming_MountedServer6)
	__VDLType_types_v_io_v23_naming_MountedServer8 := vdl.Int64Type
	__VDLType_types_v_io_v23_naming_MountedServer6.AppendField("Seconds", __VDLType_types_v_io_v23_naming_MountedServer8)
	__VDLType_types_v_io_v23_naming_MountedServer9 := vdl.Int32Type
	__VDLType_types_v_io_v23_naming_MountedServer6.AppendField("Nanos", __VDLType_types_v_io_v23_naming_MountedServer9)
	__VDLType_types_v_io_v23_naming_MountedServer4.AppendField("FromNow", __VDLType_types_v_io_v23_naming_MountedServer7)
	__VDLType_types_v_io_v23_naming_MountedServer10 := vdl.BoolType
	__VDLType_types_v_io_v23_naming_MountedServer4.AppendField("NoDeadline", __VDLType_types_v_io_v23_naming_MountedServer10)
	__VDLType_types_v_io_v23_naming_MountedServer1.AppendField("Deadline", __VDLType_types_v_io_v23_naming_MountedServer5)
	__VDLType_types_v_io_v23_naming_MountedServerBuilder.Build()
	__VDLType_types_v_io_v23_naming_MountedServerv, err := __VDLType_types_v_io_v23_naming_MountedServer2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_types_v_io_v23_naming_MountedServerv
}
func init() {
	__VDLType_types_v_io_v23_naming_MountedServer = __VDLType_types_v_io_v23_naming_MountedServer_gen()
}
func __VDLEnsureNativeBuilt_types() {
	if __VDLTypetypes1 == nil {
		__VDLTypetypes1 = __VDLTypetypes1_gen()
	}
	if __VDLTypetypes0 == nil {
		__VDLTypetypes0 = __VDLTypetypes0_gen()
	}
	if __VDLTypetypes2 == nil {
		__VDLTypetypes2 = __VDLTypetypes2_gen()
	}
	if __VDLType_types_time_WireDeadline == nil {
		__VDLType_types_time_WireDeadline = __VDLType_types_time_WireDeadline_gen()
	}
	if __VDLType_types_v_io_v23_naming_GlobReply == nil {
		__VDLType_types_v_io_v23_naming_GlobReply = __VDLType_types_v_io_v23_naming_GlobReply_gen()
	}
	if __VDLType_types_v_io_v23_naming_MountEntry == nil {
		__VDLType_types_v_io_v23_naming_MountEntry = __VDLType_types_v_io_v23_naming_MountEntry_gen()
	}
	if __VDLType_types_v_io_v23_naming_MountedServer == nil {
		__VDLType_types_v_io_v23_naming_MountedServer = __VDLType_types_v_io_v23_naming_MountedServer_gen()
	}
}

const Replace = MountFlag(1) // Replace means the mount should replace what is currently at the mount point

const MT = MountFlag(2) // MT means that the target server is a mount table.

const Leaf = MountFlag(4) // Leaf means that the target server is a leaf.
