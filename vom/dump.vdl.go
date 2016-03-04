// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: dump.vdl

package vom

import (
	// VDL system imports
	"fmt"
	"v.io/v23/vdl"
)

type (
	// Primitive represents any single field of the Primitive union type.
	//
	// Primitive represents one of the primitive vom values.  All vom values are
	// composed of combinations of these primitives.
	Primitive interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the Primitive union type.
		__VDLReflect(__PrimitiveReflect)
		FillVDLTarget(vdl.Target, *vdl.Type) error
		IsZero() bool
	}
	// PrimitivePBool represents field PBool of the Primitive union type.
	PrimitivePBool struct{ Value bool }
	// PrimitivePByte represents field PByte of the Primitive union type.
	PrimitivePByte struct{ Value byte }
	// PrimitivePUint represents field PUint of the Primitive union type.
	PrimitivePUint struct{ Value uint64 }
	// PrimitivePInt represents field PInt of the Primitive union type.
	PrimitivePInt struct{ Value int64 }
	// PrimitivePFloat represents field PFloat of the Primitive union type.
	PrimitivePFloat struct{ Value float64 }
	// PrimitivePString represents field PString of the Primitive union type.
	PrimitivePString struct{ Value string }
	// PrimitivePControl represents field PControl of the Primitive union type.
	PrimitivePControl struct{ Value ControlKind }
	// __PrimitiveReflect describes the Primitive union type.
	__PrimitiveReflect struct {
		Name  string `vdl:"v.io/v23/vom.Primitive"`
		Type  Primitive
		Union struct {
			PBool    PrimitivePBool
			PByte    PrimitivePByte
			PUint    PrimitivePUint
			PInt     PrimitivePInt
			PFloat   PrimitivePFloat
			PString  PrimitivePString
			PControl PrimitivePControl
		}
	}
)

func (x PrimitivePBool) Index() int                      { return 0 }
func (x PrimitivePBool) Interface() interface{}          { return x.Value }
func (x PrimitivePBool) Name() string                    { return "PBool" }
func (x PrimitivePBool) __VDLReflect(__PrimitiveReflect) {}

func (m PrimitivePBool) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	fieldsTarget1, err := t.StartFields(__VDLType_dump_v_io_v23_vom_Primitive)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("PBool")
	if err != nil {
		return err
	}
	if err := fieldTarget3.FromBool(bool(m.Value), vdl.BoolType); err != nil {
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

func (m PrimitivePBool) MakeVDLTarget() vdl.Target {
	return nil
}

func (m PrimitivePBool) IsZero() bool {

	var2 := (m.Value == false)
	return var2
}

func (x PrimitivePByte) Index() int                      { return 1 }
func (x PrimitivePByte) Interface() interface{}          { return x.Value }
func (x PrimitivePByte) Name() string                    { return "PByte" }
func (x PrimitivePByte) __VDLReflect(__PrimitiveReflect) {}

func (m PrimitivePByte) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	fieldsTarget1, err := t.StartFields(__VDLType_dump_v_io_v23_vom_Primitive)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("PByte")
	if err != nil {
		return err
	}
	if err := fieldTarget3.FromUint(uint64(m.Value), vdl.ByteType); err != nil {
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

func (m PrimitivePByte) MakeVDLTarget() vdl.Target {
	return nil
}

func (m PrimitivePByte) IsZero() bool {

	unionField2 := false
	return unionField2
}

func (x PrimitivePUint) Index() int                      { return 2 }
func (x PrimitivePUint) Interface() interface{}          { return x.Value }
func (x PrimitivePUint) Name() string                    { return "PUint" }
func (x PrimitivePUint) __VDLReflect(__PrimitiveReflect) {}

func (m PrimitivePUint) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	fieldsTarget1, err := t.StartFields(__VDLType_dump_v_io_v23_vom_Primitive)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("PUint")
	if err != nil {
		return err
	}
	if err := fieldTarget3.FromUint(uint64(m.Value), vdl.Uint64Type); err != nil {
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

func (m PrimitivePUint) MakeVDLTarget() vdl.Target {
	return nil
}

func (m PrimitivePUint) IsZero() bool {

	unionField2 := false
	return unionField2
}

func (x PrimitivePInt) Index() int                      { return 3 }
func (x PrimitivePInt) Interface() interface{}          { return x.Value }
func (x PrimitivePInt) Name() string                    { return "PInt" }
func (x PrimitivePInt) __VDLReflect(__PrimitiveReflect) {}

func (m PrimitivePInt) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	fieldsTarget1, err := t.StartFields(__VDLType_dump_v_io_v23_vom_Primitive)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("PInt")
	if err != nil {
		return err
	}
	if err := fieldTarget3.FromInt(int64(m.Value), vdl.Int64Type); err != nil {
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

func (m PrimitivePInt) MakeVDLTarget() vdl.Target {
	return nil
}

func (m PrimitivePInt) IsZero() bool {

	unionField2 := false
	return unionField2
}

func (x PrimitivePFloat) Index() int                      { return 4 }
func (x PrimitivePFloat) Interface() interface{}          { return x.Value }
func (x PrimitivePFloat) Name() string                    { return "PFloat" }
func (x PrimitivePFloat) __VDLReflect(__PrimitiveReflect) {}

func (m PrimitivePFloat) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	fieldsTarget1, err := t.StartFields(__VDLType_dump_v_io_v23_vom_Primitive)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("PFloat")
	if err != nil {
		return err
	}
	if err := fieldTarget3.FromFloat(float64(m.Value), vdl.Float64Type); err != nil {
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

func (m PrimitivePFloat) MakeVDLTarget() vdl.Target {
	return nil
}

func (m PrimitivePFloat) IsZero() bool {

	unionField2 := false
	return unionField2
}

func (x PrimitivePString) Index() int                      { return 5 }
func (x PrimitivePString) Interface() interface{}          { return x.Value }
func (x PrimitivePString) Name() string                    { return "PString" }
func (x PrimitivePString) __VDLReflect(__PrimitiveReflect) {}

func (m PrimitivePString) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	fieldsTarget1, err := t.StartFields(__VDLType_dump_v_io_v23_vom_Primitive)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("PString")
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

func (m PrimitivePString) MakeVDLTarget() vdl.Target {
	return nil
}

func (m PrimitivePString) IsZero() bool {

	unionField2 := false
	return unionField2
}

func (x PrimitivePControl) Index() int                      { return 6 }
func (x PrimitivePControl) Interface() interface{}          { return x.Value }
func (x PrimitivePControl) Name() string                    { return "PControl" }
func (x PrimitivePControl) __VDLReflect(__PrimitiveReflect) {}

func (m PrimitivePControl) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	fieldsTarget1, err := t.StartFields(__VDLType_dump_v_io_v23_vom_Primitive)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("PControl")
	if err != nil {
		return err
	}

	if err := m.Value.FillVDLTarget(fieldTarget3, __VDLType_dump_v_io_v23_vom_ControlKind); err != nil {
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

func (m PrimitivePControl) MakeVDLTarget() vdl.Target {
	return nil
}

func (m PrimitivePControl) IsZero() bool {

	unionField2 := false
	return unionField2
}

// DumpAtom describes a single indivisible piece of the vom encoding.  The vom
// encoding is composed of a stream of these atoms.
type DumpAtom struct {
	Kind  DumpKind  // The kind of this atom.
	Bytes []byte    // Raw bytes in the vom encoding representing this atom.
	Data  Primitive // Primitive data corresponding to the raw bytes.
	Debug string    // Free-form debug string with more information.
}

func (DumpAtom) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vom.DumpAtom"`
}) {
}

func (m *DumpAtom) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	if __VDLType_dump_v_io_v23_vom_DumpAtom == nil || __VDLTypedump0 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	var2 := m.Kind.IsZero()
	if !var2 {
		keyTarget3, fieldTarget4, err := fieldsTarget1.StartField("Kind")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {

			if err := m.Kind.FillVDLTarget(fieldTarget4, __VDLType_dump_v_io_v23_vom_DumpKind); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget3, fieldTarget4); err != nil {
				return err
			}
		}
	}
	var var5 bool
	if len(m.Bytes) == 0 {
		var5 = true
	}
	if !var5 {
		keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("Bytes")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {

			if err := fieldTarget7.FromBytes([]byte(m.Bytes), __VDLTypedump1); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
				return err
			}
		}
	}
	var8 := m.Data.IsZero()
	if !var8 {
		keyTarget9, fieldTarget10, err := fieldsTarget1.StartField("Data")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {

			unionValue11 := m.Data
			if unionValue11 == nil {
				unionValue11 = PrimitivePBool{}
			}
			if err := unionValue11.FillVDLTarget(fieldTarget10, __VDLType_dump_v_io_v23_vom_Primitive); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget9, fieldTarget10); err != nil {
				return err
			}
		}
	}
	var12 := (m.Debug == "")
	if !var12 {
		keyTarget13, fieldTarget14, err := fieldsTarget1.StartField("Debug")
		if err != vdl.ErrFieldNoExist && err != nil {
			return err
		}
		if err != vdl.ErrFieldNoExist {
			if err := fieldTarget14.FromString(string(m.Debug), vdl.StringType); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget13, fieldTarget14); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *DumpAtom) MakeVDLTarget() vdl.Target {
	return nil
}

func (m *DumpAtom) IsZero() bool {

	var1 := true
	var2 := m.Kind.IsZero()
	var1 = var1 && var2
	var var3 bool
	if len(m.Bytes) == 0 {
		var3 = true
	}
	var1 = var1 && var3
	var4 := m.Data.IsZero()
	var1 = var1 && var4
	var5 := (m.Debug == "")
	var1 = var1 && var5
	return var1
}

// DumpKind enumerates the different kinds of dump atoms.
type DumpKind int

const (
	DumpKindVersion DumpKind = iota
	DumpKindControl
	DumpKindMsgId
	DumpKindTypeMsg
	DumpKindValueMsg
	DumpKindMsgLen
	DumpKindAnyMsgLen
	DumpKindAnyLensLen
	DumpKindTypeIdsLen
	DumpKindTypeId
	DumpKindPrimValue
	DumpKindByteLen
	DumpKindValueLen
	DumpKindIndex
	DumpKindWireTypeIndex
)

// DumpKindAll holds all labels for DumpKind.
var DumpKindAll = [...]DumpKind{DumpKindVersion, DumpKindControl, DumpKindMsgId, DumpKindTypeMsg, DumpKindValueMsg, DumpKindMsgLen, DumpKindAnyMsgLen, DumpKindAnyLensLen, DumpKindTypeIdsLen, DumpKindTypeId, DumpKindPrimValue, DumpKindByteLen, DumpKindValueLen, DumpKindIndex, DumpKindWireTypeIndex}

// DumpKindFromString creates a DumpKind from a string label.
func DumpKindFromString(label string) (x DumpKind, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *DumpKind) Set(label string) error {
	switch label {
	case "Version", "version":
		*x = DumpKindVersion
		return nil
	case "Control", "control":
		*x = DumpKindControl
		return nil
	case "MsgId", "msgid":
		*x = DumpKindMsgId
		return nil
	case "TypeMsg", "typemsg":
		*x = DumpKindTypeMsg
		return nil
	case "ValueMsg", "valuemsg":
		*x = DumpKindValueMsg
		return nil
	case "MsgLen", "msglen":
		*x = DumpKindMsgLen
		return nil
	case "AnyMsgLen", "anymsglen":
		*x = DumpKindAnyMsgLen
		return nil
	case "AnyLensLen", "anylenslen":
		*x = DumpKindAnyLensLen
		return nil
	case "TypeIdsLen", "typeidslen":
		*x = DumpKindTypeIdsLen
		return nil
	case "TypeId", "typeid":
		*x = DumpKindTypeId
		return nil
	case "PrimValue", "primvalue":
		*x = DumpKindPrimValue
		return nil
	case "ByteLen", "bytelen":
		*x = DumpKindByteLen
		return nil
	case "ValueLen", "valuelen":
		*x = DumpKindValueLen
		return nil
	case "Index", "index":
		*x = DumpKindIndex
		return nil
	case "WireTypeIndex", "wiretypeindex":
		*x = DumpKindWireTypeIndex
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in vom.DumpKind", label)
}

// String returns the string label of x.
func (x DumpKind) String() string {
	switch x {
	case DumpKindVersion:
		return "Version"
	case DumpKindControl:
		return "Control"
	case DumpKindMsgId:
		return "MsgId"
	case DumpKindTypeMsg:
		return "TypeMsg"
	case DumpKindValueMsg:
		return "ValueMsg"
	case DumpKindMsgLen:
		return "MsgLen"
	case DumpKindAnyMsgLen:
		return "AnyMsgLen"
	case DumpKindAnyLensLen:
		return "AnyLensLen"
	case DumpKindTypeIdsLen:
		return "TypeIdsLen"
	case DumpKindTypeId:
		return "TypeId"
	case DumpKindPrimValue:
		return "PrimValue"
	case DumpKindByteLen:
		return "ByteLen"
	case DumpKindValueLen:
		return "ValueLen"
	case DumpKindIndex:
		return "Index"
	case DumpKindWireTypeIndex:
		return "WireTypeIndex"
	}
	return ""
}

func (DumpKind) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vom.DumpKind"`
	Enum struct{ Version, Control, MsgId, TypeMsg, ValueMsg, MsgLen, AnyMsgLen, AnyLensLen, TypeIdsLen, TypeId, PrimValue, ByteLen, ValueLen, Index, WireTypeIndex string }
}) {
}

func (m DumpKind) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromEnumLabel(m.String(), __VDLType_dump_v_io_v23_vom_DumpKind); err != nil {
		return err
	}
	return nil
}

func (m DumpKind) MakeVDLTarget() vdl.Target {
	return nil
}

func (m DumpKind) IsZero() bool {

	var1 := (m == DumpKindVersion)
	return var1
}

// ControlKind enumerates the different kinds of control bytes.
type ControlKind int

const (
	ControlKindNil ControlKind = iota
	ControlKindEnd
	ControlKindIncompleteType
)

// ControlKindAll holds all labels for ControlKind.
var ControlKindAll = [...]ControlKind{ControlKindNil, ControlKindEnd, ControlKindIncompleteType}

// ControlKindFromString creates a ControlKind from a string label.
func ControlKindFromString(label string) (x ControlKind, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *ControlKind) Set(label string) error {
	switch label {
	case "Nil", "nil":
		*x = ControlKindNil
		return nil
	case "End", "end":
		*x = ControlKindEnd
		return nil
	case "IncompleteType", "incompletetype":
		*x = ControlKindIncompleteType
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in vom.ControlKind", label)
}

// String returns the string label of x.
func (x ControlKind) String() string {
	switch x {
	case ControlKindNil:
		return "Nil"
	case ControlKindEnd:
		return "End"
	case ControlKindIncompleteType:
		return "IncompleteType"
	}
	return ""
}

func (ControlKind) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vom.ControlKind"`
	Enum struct{ Nil, End, IncompleteType string }
}) {
}

func (m ControlKind) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromEnumLabel(m.String(), __VDLType_dump_v_io_v23_vom_ControlKind); err != nil {
		return err
	}
	return nil
}

func (m ControlKind) MakeVDLTarget() vdl.Target {
	return nil
}

func (m ControlKind) IsZero() bool {

	var1 := (m == ControlKindNil)
	return var1
}

func init() {
	vdl.Register((*Primitive)(nil))
	vdl.Register((*DumpAtom)(nil))
	vdl.Register((*DumpKind)(nil))
	vdl.Register((*ControlKind)(nil))
}

var __VDLTypedump0 *vdl.Type = vdl.TypeOf((*DumpAtom)(nil))
var __VDLTypedump1 *vdl.Type = vdl.TypeOf([]byte(nil))
var __VDLType_dump_v_io_v23_vom_ControlKind *vdl.Type = vdl.TypeOf(ControlKindNil)
var __VDLType_dump_v_io_v23_vom_DumpAtom *vdl.Type = vdl.TypeOf(DumpAtom{
	Data: PrimitivePBool{false},
})
var __VDLType_dump_v_io_v23_vom_DumpKind *vdl.Type = vdl.TypeOf(DumpKindVersion)
var __VDLType_dump_v_io_v23_vom_Primitive *vdl.Type = vdl.TypeOf(Primitive(PrimitivePBool{false}))

func __VDLEnsureNativeBuilt_dump() {
}
