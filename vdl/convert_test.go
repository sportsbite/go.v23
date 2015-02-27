package vdl_test

// This test is in the vdl_test package to avoid a cyclic dependency with the
// "v.io/v23/verror" package.  We need to import the verror package in
// order to test error conversions.
//
// TODO(toddw): test values of recursive types.

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"v.io/v23/vdl"
	"v.io/v23/verror"
)

func errorValue(e verror.E) *vdl.Value {
	verr := vdl.NonNilZeroValue(vdl.ErrorType)
	vv := verr.Elem()
	vv.StructField(0).AssignString(string(e.ID))
	vv.StructField(1).AssignEnumLabel(retryFromAction(e.Action))
	vv.StructField(2).AssignString(e.Msg)
	vv.StructField(3).AssignLen(len(e.ParamList))
	for ix, p := range e.ParamList {
		vv.StructField(3).Index(ix).Assign(vdl.ValueOf(p))
	}
	return verr
}

func retryFromAction(action verror.ActionCode) string {
	switch action.RetryAction() {
	case verror.NoRetry:
		return vdl.WireRetryCodeNoRetry.String()
	case verror.RetryConnection:
		return vdl.WireRetryCodeRetryConnection.String()
	case verror.RetryRefetch:
		return vdl.WireRetryCodeRetryRefetch.String()
	case verror.RetryBackoff:
		return vdl.WireRetryCodeRetryBackoff.String()
	}
	// Default to NoRetry.
	return vdl.WireRetryCodeNoRetry.String()
}

// Each group of values in vvNAME and rvNAME are all mutually convertible.
var (
	rvError1 = verror.E{
		ID:        verror.ID("id1"),
		Action:    verror.NoRetry,
		Msg:       "msg1",
		ParamList: nil,
	}
	rvError2 = verror.E{
		ID:        verror.ID("id2"),
		Action:    verror.RetryConnection,
		Msg:       "msg2",
		ParamList: []interface{}{"abc", int32(123)},
	}
	rvError3 = verror.E{
		ID:        verror.ID("id3"),
		Action:    verror.RetryBackoff,
		Msg:       "msg3",
		ParamList: []interface{}{rvError1, &rvError2},
	}
	vvError1 = errorValue(rvError1)
	vvError2 = errorValue(rvError2)
	vvError3 = errorValue(rvError3)

	vvBoolTrue = []*vdl.Value{
		vdl.BoolValue(true), vdl.ZeroValue(vdl.BoolTypeN).AssignBool(true),
	}
	vvStrABC = []*vdl.Value{
		vdl.StringValue("ABC"), vdl.ZeroValue(vdl.StringTypeN).AssignString("ABC"),
		vdl.ZeroValue(vdl.BytesType).AssignBytes([]byte("ABC")), vdl.ZeroValue(vdl.BytesTypeN).AssignBytes([]byte("ABC")),
		vdl.ZeroValue(vdl.Bytes3Type).AssignBytes([]byte("ABC")), vdl.ZeroValue(vdl.Bytes3TypeN).AssignBytes([]byte("ABC")),
		vdl.ZeroValue(vdl.EnumTypeN).AssignEnumLabel("ABC"),
	}
	vvTypeObjectBool = []*vdl.Value{
		vdl.TypeObjectValue(vdl.BoolType),
	}
	vvSeq123 = []*vdl.Value{
		vdl.SeqNumValue(vdl.Array3Uint64Type, 1, 2, 3),
		vdl.SeqNumValue(vdl.Array3Uint64TypeN, 1, 2, 3),
		vdl.SeqNumValue(vdl.Array3Int64Type, 1, 2, 3),
		vdl.SeqNumValue(vdl.Array3Int64TypeN, 1, 2, 3),
		vdl.SeqNumValue(vdl.Array3Float64Type, 1, 2, 3),
		vdl.SeqNumValue(vdl.Array3Float64TypeN, 1, 2, 3),
		vdl.SeqNumValue(vdl.Array3Complex64Type, 1, 2, 3),
		vdl.SeqNumValue(vdl.Array3Complex64TypeN, 1, 2, 3),
		vdl.SeqNumValue(vdl.ListUint64Type, 1, 2, 3),
		vdl.SeqNumValue(vdl.ListUint64TypeN, 1, 2, 3),
		vdl.SeqNumValue(vdl.ListInt64Type, 1, 2, 3),
		vdl.SeqNumValue(vdl.ListInt64TypeN, 1, 2, 3),
		vdl.SeqNumValue(vdl.ListFloat64Type, 1, 2, 3),
		vdl.SeqNumValue(vdl.ListFloat64TypeN, 1, 2, 3),
		vdl.SeqNumValue(vdl.ListComplex64Type, 1, 2, 3),
		vdl.SeqNumValue(vdl.ListComplex64TypeN, 1, 2, 3),
	}
	vvSet123 = []*vdl.Value{
		vdl.SetNumValue(vdl.SetUint64Type, 1, 2, 3),
		vdl.SetNumValue(vdl.SetUint64TypeN, 1, 2, 3),
		vdl.SetNumValue(vdl.SetInt64Type, 1, 2, 3),
		vdl.SetNumValue(vdl.SetInt64TypeN, 1, 2, 3),
		vdl.SetNumValue(vdl.SetFloat64Type, 1, 2, 3),
		vdl.SetNumValue(vdl.SetFloat64TypeN, 1, 2, 3),
		vdl.SetNumValue(vdl.SetComplex64Type, 1, 2, 3),
		vdl.SetNumValue(vdl.SetComplex64TypeN, 1, 2, 3),
	}
	vvMap123True = []*vdl.Value{
		vdl.MapNumBoolValue(vdl.MapUint64BoolType, vdl.NB{1, true}, vdl.NB{2, true}, vdl.NB{3, true}),
		vdl.MapNumBoolValue(vdl.MapUint64BoolTypeN, vdl.NB{1, true}, vdl.NB{2, true}, vdl.NB{3, true}),
		vdl.MapNumBoolValue(vdl.MapInt64BoolType, vdl.NB{1, true}, vdl.NB{2, true}, vdl.NB{3, true}),
		vdl.MapNumBoolValue(vdl.MapInt64BoolTypeN, vdl.NB{1, true}, vdl.NB{2, true}, vdl.NB{3, true}),
		vdl.MapNumBoolValue(vdl.MapFloat64BoolType, vdl.NB{1, true}, vdl.NB{2, true}, vdl.NB{3, true}),
		vdl.MapNumBoolValue(vdl.MapFloat64BoolTypeN, vdl.NB{1, true}, vdl.NB{2, true}, vdl.NB{3, true}),
		vdl.MapNumBoolValue(vdl.MapComplex64BoolType, vdl.NB{1, true}, vdl.NB{2, true}, vdl.NB{3, true}),
		vdl.MapNumBoolValue(vdl.MapComplex64BoolTypeN, vdl.NB{1, true}, vdl.NB{2, true}, vdl.NB{3, true}),
	}
	vvSetMap123       = append(vvSet123, vvMap123True...)
	vvMap123FalseTrue = []*vdl.Value{
		vdl.MapNumBoolValue(vdl.MapUint64BoolType, vdl.NB{1, false}, vdl.NB{2, true}, vdl.NB{3, false}),
		vdl.MapNumBoolValue(vdl.MapUint64BoolTypeN, vdl.NB{1, false}, vdl.NB{2, true}, vdl.NB{3, false}),
		vdl.MapNumBoolValue(vdl.MapInt64BoolType, vdl.NB{1, false}, vdl.NB{2, true}, vdl.NB{3, false}),
		vdl.MapNumBoolValue(vdl.MapInt64BoolTypeN, vdl.NB{1, false}, vdl.NB{2, true}, vdl.NB{3, false}),
		vdl.MapNumBoolValue(vdl.MapFloat64BoolType, vdl.NB{1, false}, vdl.NB{2, true}, vdl.NB{3, false}),
		vdl.MapNumBoolValue(vdl.MapFloat64BoolTypeN, vdl.NB{1, false}, vdl.NB{2, true}, vdl.NB{3, false}),
		vdl.MapNumBoolValue(vdl.MapComplex64BoolType, vdl.NB{1, false}, vdl.NB{2, true}, vdl.NB{3, false}),
		vdl.MapNumBoolValue(vdl.MapComplex64BoolTypeN, vdl.NB{1, false}, vdl.NB{2, true}, vdl.NB{3, false}),
	}
	vvSetXYZ = []*vdl.Value{
		vdl.SetStringValue(vdl.SetStringType, "X", "Y", "Z"),
		vdl.SetStringValue(vdl.SetStringTypeN, "X", "Y", "Z"),
	}
	vvMapXYZTrue = []*vdl.Value{
		vdl.MapStringBoolValue(vdl.MapStringBoolType, vdl.SB{"X", true}, vdl.SB{"Y", true}, vdl.SB{"Z", true}),
		vdl.MapStringBoolValue(vdl.MapStringBoolTypeN, vdl.SB{"X", true}, vdl.SB{"Y", true}, vdl.SB{"Z", true}),
	}
	vvStructXYZTrue = []*vdl.Value{
		vdl.StructBoolValue(vdl.StructXYZBoolType, vdl.SB{"X", true}, vdl.SB{"Y", true}, vdl.SB{"Z", true}),
		vdl.StructBoolValue(vdl.StructXYZBoolTypeN, vdl.SB{"X", true}, vdl.SB{"Y", true}, vdl.SB{"Z", true}),
	}
	vvSetMapStructXYZ       = append(append(vvSetXYZ, vvMapXYZTrue...), vvStructXYZTrue...)
	vvMapStructXYZFalseTrue = []*vdl.Value{
		vdl.MapStringBoolValue(vdl.MapStringBoolType, vdl.SB{"X", false}, vdl.SB{"Y", true}, vdl.SB{"Z", false}),
		vdl.MapStringBoolValue(vdl.MapStringBoolTypeN, vdl.SB{"X", false}, vdl.SB{"Y", true}, vdl.SB{"Z", false}),
		vdl.StructBoolValue(vdl.StructXYZBoolType, vdl.SB{"X", false}, vdl.SB{"Y", true}, vdl.SB{"Z", false}),
		vdl.StructBoolValue(vdl.StructXYZBoolTypeN, vdl.SB{"X", false}, vdl.SB{"Y", true}, vdl.SB{"Z", false}),
	}
	vvMapStructXYZEmpty = []*vdl.Value{
		vdl.MapStringEmptyValue(vdl.MapStringEmptyType, "X", "Y", "Z"),
		vdl.MapStringEmptyValue(vdl.MapStringEmptyTypeN, "X", "Y", "Z"),
		vdl.ZeroValue(vdl.StructXYZEmptyType), vdl.ZeroValue(vdl.StructXYZEmptyTypeN),
	}
	vvStructWXFalseTrue = []*vdl.Value{
		vdl.StructBoolValue(vdl.StructWXBoolType, vdl.SB{"W", false}, vdl.SB{"X", true}),
		vdl.StructBoolValue(vdl.StructWXBoolTypeN, vdl.SB{"W", false}, vdl.SB{"X", true}),
	}
	vvMapVWX123 = []*vdl.Value{
		vdl.MapStringNumValue(vdl.MapStringUint64Type, vdl.SN{"V", 1}, vdl.SN{"W", 2}, vdl.SN{"X", 3}),
		vdl.MapStringNumValue(vdl.MapStringUint64TypeN, vdl.SN{"V", 1}, vdl.SN{"W", 2}, vdl.SN{"X", 3}),
		vdl.MapStringNumValue(vdl.MapStringInt64Type, vdl.SN{"V", 1}, vdl.SN{"W", 2}, vdl.SN{"X", 3}),
		vdl.MapStringNumValue(vdl.MapStringInt64TypeN, vdl.SN{"V", 1}, vdl.SN{"W", 2}, vdl.SN{"X", 3}),
		vdl.MapStringNumValue(vdl.MapStringFloat64Type, vdl.SN{"V", 1}, vdl.SN{"W", 2}, vdl.SN{"X", 3}),
		vdl.MapStringNumValue(vdl.MapStringFloat64TypeN, vdl.SN{"V", 1}, vdl.SN{"W", 2}, vdl.SN{"X", 3}),
		vdl.MapStringNumValue(vdl.MapStringComplex64Type, vdl.SN{"V", 1}, vdl.SN{"W", 2}, vdl.SN{"X", 3}),
		vdl.MapStringNumValue(vdl.MapStringComplex64TypeN, vdl.SN{"V", 1}, vdl.SN{"W", 2}, vdl.SN{"X", 3}),
	}
	vvStructVWX123 = []*vdl.Value{
		vdl.StructNumValue(vdl.StructVWXUint64Type, vdl.SN{"V", 1}, vdl.SN{"W", 2}, vdl.SN{"X", 3}),
		vdl.StructNumValue(vdl.StructVWXUint64TypeN, vdl.SN{"V", 1}, vdl.SN{"W", 2}, vdl.SN{"X", 3}),
		vdl.StructNumValue(vdl.StructVWXInt64Type, vdl.SN{"V", 1}, vdl.SN{"W", 2}, vdl.SN{"X", 3}),
		vdl.StructNumValue(vdl.StructVWXInt64TypeN, vdl.SN{"V", 1}, vdl.SN{"W", 2}, vdl.SN{"X", 3}),
		vdl.StructNumValue(vdl.StructVWXFloat64Type, vdl.SN{"V", 1}, vdl.SN{"W", 2}, vdl.SN{"X", 3}),
		vdl.StructNumValue(vdl.StructVWXFloat64TypeN, vdl.SN{"V", 1}, vdl.SN{"W", 2}, vdl.SN{"X", 3}),
		vdl.StructNumValue(vdl.StructVWXComplex64Type, vdl.SN{"V", 1}, vdl.SN{"W", 2}, vdl.SN{"X", 3}),
		vdl.StructNumValue(vdl.StructVWXComplex64TypeN, vdl.SN{"V", 1}, vdl.SN{"W", 2}, vdl.SN{"X", 3}),
	}
	vvMapStructVWX123 = append(vvMapVWX123, vvStructVWX123...)
	vvStructUV01      = []*vdl.Value{
		vdl.StructNumValue(vdl.StructUVUint64Type, vdl.SN{"U", 0}, vdl.SN{"V", 1}),
		vdl.StructNumValue(vdl.StructUVUint64TypeN, vdl.SN{"U", 0}, vdl.SN{"V", 1}),
		vdl.StructNumValue(vdl.StructUVInt64Type, vdl.SN{"U", 0}, vdl.SN{"V", 1}),
		vdl.StructNumValue(vdl.StructUVInt64TypeN, vdl.SN{"U", 0}, vdl.SN{"V", 1}),
		vdl.StructNumValue(vdl.StructUVFloat64Type, vdl.SN{"U", 0}, vdl.SN{"V", 1}),
		vdl.StructNumValue(vdl.StructUVFloat64TypeN, vdl.SN{"U", 0}, vdl.SN{"V", 1}),
		vdl.StructNumValue(vdl.StructUVComplex64Type, vdl.SN{"U", 0}, vdl.SN{"V", 1}),
		vdl.StructNumValue(vdl.StructUVComplex64TypeN, vdl.SN{"U", 0}, vdl.SN{"V", 1}),
	}
	vvEmptyStruct = []*vdl.Value{vdl.ZeroValue(vdl.EmptyType), vdl.ZeroValue(vdl.EmptyTypeN)}

	rvBoolTrue = []interface{}{
		bool(true), vdl.NBool(true),
	}
	rvStrABC = []interface{}{
		string("ABC"), []byte("ABC"), [3]byte{'A', 'B', 'C'},
		vdl.NString("ABC"), vdl.NSliceUint8("ABC"), vdl.NArray3Uint8{'A', 'B', 'C'},
		vdl.NEnumABC,
	}
	rvTypeObjectBool = []interface{}{
		vdl.BoolType, vdl.NType(vdl.BoolType),
	}
	rvSeq123 = []interface{}{
		[3]uint64{1, 2, 3}, []uint64{1, 2, 3}, vdl.NArray3Uint64{1, 2, 3}, vdl.NSliceUint64{1, 2, 3},
		[3]int64{1, 2, 3}, []int64{1, 2, 3}, vdl.NArray3Int64{1, 2, 3}, vdl.NSliceInt64{1, 2, 3},
		[3]float64{1, 2, 3}, []float64{1, 2, 3}, vdl.NArray3Float64{1, 2, 3}, vdl.NSliceFloat64{1, 2, 3},
		[3]complex64{1, 2, 3}, []complex64{1, 2, 3}, vdl.NArray3Complex64{1, 2, 3}, vdl.NSliceComplex64{1, 2, 3},
	}
	rvSet123 = []interface{}{
		map[uint64]struct{}{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		map[int64]struct{}{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		map[float64]struct{}{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		map[complex64]struct{}{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		vdl.NMapUint64Empty{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		vdl.NMapInt64Empty{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		vdl.NMapFloat64Empty{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		vdl.NMapComplex64Empty{1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
	}
	rvMap123True = []interface{}{
		map[uint64]bool{1: true, 2: true, 3: true},
		map[int64]bool{1: true, 2: true, 3: true},
		map[float64]bool{1: true, 2: true, 3: true},
		map[complex64]bool{1: true, 2: true, 3: true},
		vdl.NMapUint64Bool{1: true, 2: true, 3: true},
		vdl.NMapInt64Bool{1: true, 2: true, 3: true},
		vdl.NMapFloat64Bool{1: true, 2: true, 3: true},
		vdl.NMapComplex64Bool{1: true, 2: true, 3: true},
	}
	rvSetMap123       = append(rvSet123, rvMap123True...)
	rvMap123FalseTrue = []interface{}{
		map[uint64]bool{1: false, 2: true, 3: false},
		map[int64]bool{1: false, 2: true, 3: false},
		map[float64]bool{1: false, 2: true, 3: false},
		map[complex64]bool{1: false, 2: true, 3: false},
		vdl.NMapUint64Bool{1: false, 2: true, 3: false},
		vdl.NMapInt64Bool{1: false, 2: true, 3: false},
		vdl.NMapFloat64Bool{1: false, 2: true, 3: false},
		vdl.NMapComplex64Bool{1: false, 2: true, 3: false},
	}
	rvSetXYZ = []interface{}{
		map[string]struct{}{"X": struct{}{}, "Y": struct{}{}, "Z": struct{}{}},
		vdl.NMapStringEmpty{"X": struct{}{}, "Y": struct{}{}, "Z": struct{}{}},
	}
	rvMapXYZTrue = []interface{}{
		map[string]bool{"X": true, "Y": true, "Z": true},
		vdl.NMapStringBool{"X": true, "Y": true, "Z": true},
	}
	rvStructXYZTrue = []interface{}{
		struct{ X, Y, Z bool }{X: true, Y: true, Z: true},
		struct{ a, X, b, Y, c, Z, d bool }{X: true, Y: true, Z: true},
		vdl.NStructXYZBool{X: true, Y: true, Z: true},
		vdl.NStructXYZBoolUnexported{X: true, Y: true, Z: true},
	}
	rvSetMapStructXYZ       = append(append(rvSetXYZ, rvMapXYZTrue...), rvStructXYZTrue...)
	rvMapStructXYZFalseTrue = []interface{}{
		map[string]bool{"X": false, "Y": true, "Z": false},
		vdl.NMapStringBool{"X": false, "Y": true, "Z": false},
		struct{ X, Y, Z bool }{X: false, Y: true, Z: false},
		struct{ a, X, b, Y, c, Z, d bool }{X: false, Y: true, Z: false},
		vdl.NStructXYZBool{X: false, Y: true, Z: false},
		vdl.NStructXYZBoolUnexported{X: false, Y: true, Z: false},
	}
	rvMapStructXYZEmpty = []interface{}{
		map[string]vdl.NEmpty{"X": vdl.NEmpty{}, "Y": vdl.NEmpty{}, "Z": vdl.NEmpty{}},
		vdl.NMapStringNEmpty{"X": vdl.NEmpty{}, "Y": vdl.NEmpty{}, "Z": vdl.NEmpty{}},
		vdl.NStructXYZEmpty{}, vdl.NStructXYZNEmpty{},
	}
	rvStructWXFalseTrue = []interface{}{
		struct{ W, X bool }{W: false, X: true},
		vdl.NStructWXBool{W: false, X: true},
	}
	rvMapVWX123 = []interface{}{
		map[string]uint64{"V": 1, "W": 2, "X": 3},
		map[string]int64{"V": 1, "W": 2, "X": 3},
		map[string]float64{"V": 1, "W": 2, "X": 3},
		map[string]complex64{"V": 1, "W": 2, "X": 3},
		vdl.NMapStringUint64{"V": 1, "W": 2, "X": 3},
		vdl.NMapStringInt64{"V": 1, "W": 2, "X": 3},
		vdl.NMapStringFloat64{"V": 1, "W": 2, "X": 3},
		vdl.NMapStringComplex64{"V": 1, "W": 2, "X": 3},
	}
	rvStructVWX123 = []interface{}{
		struct{ V, W, X uint64 }{V: 1, W: 2, X: 3},
		struct{ V, W, X int64 }{V: 1, W: 2, X: 3},
		struct{ V, W, X float64 }{V: 1, W: 2, X: 3},
		struct{ V, W, X complex64 }{V: 1, W: 2, X: 3},
		struct {
			// Interleave unexported fields, which are ignored.
			a bool
			V int64
			b string
			W float64
			c []byte
			X complex64
			d interface{}
		}{V: 1, W: 2, X: 3},
		vdl.NStructVWXUint64{V: 1, W: 2, X: 3},
		vdl.NStructVWXInt64{V: 1, W: 2, X: 3},
		vdl.NStructVWXFloat64{V: 1, W: 2, X: 3},
		vdl.NStructVWXComplex64{V: 1, W: 2, X: 3},
		vdl.NStructVWXMixed{V: 1, W: 2, X: 3},
	}
	rvMapStructVWX123 = append(rvMapVWX123, rvStructVWX123...)
	rvStructUV01      = []interface{}{
		struct{ U, V uint64 }{U: 0, V: 1},
		struct{ U, V int64 }{U: 0, V: 1},
		struct{ U, V float64 }{U: 0, V: 1},
		struct{ U, V complex64 }{U: 0, V: 1},
		struct {
			// Interleave unexported fields, which are ignored.
			a bool
			U int64
			b string
			V float64
			c []byte
		}{U: 0, V: 1},
		vdl.NStructUVUint64{U: 0, V: 1},
		vdl.NStructUVInt64{U: 0, V: 1},
		vdl.NStructUVFloat64{U: 0, V: 1},
		vdl.NStructUVComplex64{U: 0, V: 1},
		vdl.NStructUVMixed{U: 0, V: 1},
	}
	rvEmptyStruct = []interface{}{struct{}{}, vdl.NEmpty{}}

	ttBools           = ttTypes(vvBoolTrue)
	ttStrs            = ttTypes(vvStrABC)
	ttTypeObjects     = ttTypes(vvTypeObjectBool)
	ttSeq123          = ttTypes(vvSeq123)
	ttSet123          = ttTypes(vvSet123)
	ttSetMap123       = ttTypes(vvSetMap123)
	ttSetXYZ          = ttTypes(vvSetXYZ)
	ttMapXYZBool      = ttTypes(vvMapXYZTrue)
	ttStructXYZBool   = ttTypes(vvStructXYZTrue)
	ttSetMapStructXYZ = ttTypes(vvSetMapStructXYZ)
	ttStructWXBool    = ttTypes(vvStructWXFalseTrue)
	ttMapVWXNum       = ttTypes(vvMapVWX123)
	ttStructVWXNum    = ttTypes(vvStructVWX123)
	ttMapStructVWXNum = ttTypes(vvMapStructVWX123)
	ttStructUVNum     = ttTypes(vvStructUV01)
	ttUints           = []*vdl.Type{
		vdl.ByteType, vdl.ByteTypeN,
		vdl.Uint16Type, vdl.Uint16TypeN,
		vdl.Uint32Type, vdl.Uint32TypeN,
		vdl.Uint64Type, vdl.Uint64TypeN,
	}
	ttInts = []*vdl.Type{
		vdl.Int16Type, vdl.Int16TypeN,
		vdl.Int32Type, vdl.Int32TypeN,
		vdl.Int64Type, vdl.Int64TypeN,
	}
	ttFloat32s    = []*vdl.Type{vdl.Float32Type, vdl.Float32TypeN}
	ttFloat64s    = []*vdl.Type{vdl.Float64Type, vdl.Float64TypeN}
	ttFloats      = ttJoin(ttFloat32s, ttFloat64s)
	ttComplex64s  = []*vdl.Type{vdl.Complex64Type, vdl.Complex64TypeN}
	ttComplex128s = []*vdl.Type{vdl.Complex128Type, vdl.Complex128TypeN}
	ttComplexes   = ttJoin(ttComplex64s, ttComplex128s)
	ttIntegers    = ttJoin(ttUints, ttInts)
	ttNumbers     = ttJoin(ttIntegers, ttFloats, ttComplexes)
	ttAllTypes    = ttJoin(ttBools, ttStrs, ttTypeObjects, ttNumbers, ttSeq123, ttSetMap123, ttSetMapStructXYZ, ttMapStructVWXNum)

	rtBools           = rtTypes(rvBoolTrue)
	rtStrs            = rtTypes(rvStrABC)
	rtTypeObjects     = rtTypes(rvTypeObjectBool)
	rtSeq123          = rtTypes(rvSeq123)
	rtSet123          = rtTypes(rvSet123)
	rtSetMap123       = rtTypes(rvSetMap123)
	rtSetXYZ          = rtTypes(rvSetXYZ)
	rtMapXYZBool      = rtTypes(rvMapXYZTrue)
	rtStructXYZBool   = rtTypes(rvStructXYZTrue)
	rtSetMapStructXYZ = rtTypes(rvSetMapStructXYZ)
	rtStructWXBool    = rtTypes(rvStructWXFalseTrue)
	rtMapVWXNum       = rtTypes(rvMapVWX123)
	rtStructVWXNum    = rtTypes(rvStructVWX123)
	rtMapStructVWXNum = rtTypes(rvMapStructVWX123)
	rtStructUVNum     = rtTypes(rvStructUV01)
	rtUints           = []reflect.Type{
		reflect.TypeOf(uint8(0)), reflect.TypeOf(vdl.NUint8(0)),
		reflect.TypeOf(uint16(0)), reflect.TypeOf(vdl.NUint16(0)),
		reflect.TypeOf(uint32(0)), reflect.TypeOf(vdl.NUint32(0)),
		reflect.TypeOf(uint64(0)), reflect.TypeOf(vdl.NUint64(0)),
	}
	rtInts = []reflect.Type{
		reflect.TypeOf(int8(0)), reflect.TypeOf(vdl.NInt8(0)),
		reflect.TypeOf(int16(0)), reflect.TypeOf(vdl.NInt16(0)),
		reflect.TypeOf(int32(0)), reflect.TypeOf(vdl.NInt32(0)),
		reflect.TypeOf(int64(0)), reflect.TypeOf(vdl.NInt64(0)),
	}
	rtFloat32s = []reflect.Type{
		reflect.TypeOf(float32(0)), reflect.TypeOf(vdl.NFloat32(0)),
	}
	rtFloat64s = []reflect.Type{
		reflect.TypeOf(float64(0)), reflect.TypeOf(vdl.NFloat64(0)),
	}
	rtFloats     = rtJoin(rtFloat32s, rtFloat64s)
	rtComplex64s = []reflect.Type{
		reflect.TypeOf(complex64(0)), reflect.TypeOf(vdl.NComplex64(0)),
	}
	rtComplex128s = []reflect.Type{
		reflect.TypeOf(complex128(0)), reflect.TypeOf(vdl.NComplex128(0)),
	}
	rtComplexes = rtJoin(rtComplex64s, rtComplex128s)
	rtIntegers  = rtJoin(rtUints, rtInts)
	rtNumbers   = rtJoin(rtIntegers, rtFloats, rtComplexes)
	rtAllTypes  = rtJoin(rtBools, rtStrs, rtTypeObjects, rtNumbers, rtSeq123, rtSetMap123, rtSetMapStructXYZ, rtMapStructVWXNum)

	rtInterface = reflect.TypeOf((*interface{})(nil)).Elem()
	rtPtrToType = reflect.TypeOf((*vdl.Type)(nil))
)

// Helpers to manipulate slices of *Type
func ttSetToSlice(set map[*vdl.Type]bool) (result []*vdl.Type) {
	for tt, _ := range set {
		result = append(result, tt)
	}
	return
}
func ttTypes(values []*vdl.Value) []*vdl.Type {
	uniq := make(map[*vdl.Type]bool)
	for _, v := range values {
		uniq[v.Type()] = true
	}
	return ttSetToSlice(uniq)
}
func ttJoin(types ...[]*vdl.Type) []*vdl.Type {
	uniq := make(map[*vdl.Type]bool)
	for _, ttSlice := range types {
		for _, tt := range ttSlice {
			uniq[tt] = true
		}
	}
	return ttSetToSlice(uniq)
}

// ttOtherThan returns all types from types that aren't represented in other.
func ttOtherThan(types []*vdl.Type, other ...[]*vdl.Type) (result []*vdl.Type) {
	otherMap := make(map[*vdl.Type]bool)
	for _, oo := range other {
		for _, o := range oo {
			otherMap[o] = true
		}
	}
	for _, t := range types {
		if !otherMap[t] {
			result = append(result, t)
		}
	}
	return
}

// Helpers to manipulate slices of reflect.Type
func rtSetToSlice(set map[reflect.Type]bool) (result []reflect.Type) {
	for rt, _ := range set {
		result = append(result, rt)
	}
	return
}
func rtTypes(values []interface{}) []reflect.Type {
	uniq := make(map[reflect.Type]bool)
	for _, v := range values {
		uniq[reflect.TypeOf(v)] = true
	}
	return rtSetToSlice(uniq)
}
func rtJoin(types ...[]reflect.Type) (result []reflect.Type) {
	uniq := make(map[reflect.Type]bool)
	for _, rtSlice := range types {
		for _, rt := range rtSlice {
			uniq[rt] = true
		}
	}
	return rtSetToSlice(uniq)
}

// rtOtherThan returns all types from types that aren't represented in other.
func rtOtherThan(types []reflect.Type, other ...[]reflect.Type) (result []reflect.Type) {
	otherMap := make(map[reflect.Type]bool)
	for _, oo := range other {
		for _, o := range oo {
			otherMap[o] = true
		}
	}
	for _, t := range types {
		if !otherMap[t] {
			result = append(result, t)
		}
	}
	return
}

// vvOnlyFrom returns all values from values that are represented in from.
func vvOnlyFrom(values []*vdl.Value, from []*vdl.Type) (result []*vdl.Value) {
	fromMap := make(map[*vdl.Type]bool)
	for _, f := range from {
		fromMap[f] = true
	}
	for _, v := range values {
		if fromMap[v.Type()] {
			result = append(result, v)
		}
	}
	return
}

// rvOnlyFrom returns all values from values that are represented in from.
func rvOnlyFrom(values []interface{}, from []reflect.Type) (result []interface{}) {
	fromMap := make(map[reflect.Type]bool)
	for _, f := range from {
		fromMap[f] = true
	}
	for _, v := range values {
		if fromMap[reflect.TypeOf(v)] {
			result = append(result, v)
		}
	}
	return
}

// vvFromUint returns all *Values that can represent u without loss of precision.
func vvFromUint(u uint64) (result []*vdl.Value) {
	b, i, f, c := byte(u), int64(u), float64(u), complex(float64(u), 0)
	switch {
	case u <= math.MaxInt8:
		fallthrough
	case u <= math.MaxUint8:
		result = append(result,
			vdl.ZeroValue(vdl.ByteType).AssignByte(b),
			vdl.ZeroValue(vdl.ByteTypeN).AssignByte(b))
		fallthrough
	case u <= math.MaxInt16:
		result = append(result,
			vdl.ZeroValue(vdl.Int16Type).AssignInt(i),
			vdl.ZeroValue(vdl.Int16TypeN).AssignInt(i))
		fallthrough
	case u <= math.MaxUint16:
		result = append(result,
			vdl.ZeroValue(vdl.Uint16Type).AssignUint(u),
			vdl.ZeroValue(vdl.Uint16TypeN).AssignUint(u))
		fallthrough
	case u <= 1<<24:
		result = append(result,
			vdl.ZeroValue(vdl.Float32Type).AssignFloat(f),
			vdl.ZeroValue(vdl.Float32TypeN).AssignFloat(f))
		result = append(result,
			vdl.ZeroValue(vdl.Complex64Type).AssignComplex(c),
			vdl.ZeroValue(vdl.Complex64TypeN).AssignComplex(c))
		fallthrough
	case u <= math.MaxInt32:
		result = append(result,
			vdl.ZeroValue(vdl.Int32Type).AssignInt(i),
			vdl.ZeroValue(vdl.Int32TypeN).AssignInt(i))
		fallthrough
	case u <= math.MaxUint32:
		result = append(result,
			vdl.ZeroValue(vdl.Uint32Type).AssignUint(u),
			vdl.ZeroValue(vdl.Uint32TypeN).AssignUint(u))
		fallthrough
	case u <= 1<<53:
		result = append(result,
			vdl.ZeroValue(vdl.Float64Type).AssignFloat(f),
			vdl.ZeroValue(vdl.Float64TypeN).AssignFloat(f))
		result = append(result,
			vdl.ZeroValue(vdl.Complex128Type).AssignComplex(c),
			vdl.ZeroValue(vdl.Complex128TypeN).AssignComplex(c))
		fallthrough
	case u <= math.MaxInt64:
		result = append(result,
			vdl.ZeroValue(vdl.Int64Type).AssignInt(i),
			vdl.ZeroValue(vdl.Int64TypeN).AssignInt(i))
		fallthrough
	default:
		result = append(result,
			vdl.ZeroValue(vdl.Uint64Type).AssignUint(u),
			vdl.ZeroValue(vdl.Uint64TypeN).AssignUint(u))
	}
	return result
}

// rvFromUint returns all values that can represent u without loss of precision.
func rvFromUint(u uint64) (result []interface{}) {
	c64, c128 := complex(float32(u), 0), complex(float64(u), 0)
	switch {
	case u <= math.MaxInt8:
		result = append(result, int8(u), vdl.NInt8(u))
		fallthrough
	case u <= math.MaxUint8:
		result = append(result, uint8(u), vdl.NUint8(u))
		fallthrough
	case u <= math.MaxInt16:
		result = append(result, int16(u), vdl.NInt16(u))
		fallthrough
	case u <= math.MaxUint16:
		result = append(result, uint16(u), vdl.NUint16(u))
		fallthrough
	case u <= 1<<24:
		result = append(result, float32(u), vdl.NFloat32(u))
		result = append(result, c64, vdl.NComplex64(c64))
		fallthrough
	case u <= math.MaxInt32:
		result = append(result, int32(u), vdl.NInt32(u))
		fallthrough
	case u <= math.MaxUint32:
		result = append(result, uint32(u), vdl.NUint32(u))
		fallthrough
	case u <= 1<<53:
		result = append(result, float64(u), vdl.NFloat64(u))
		result = append(result, c128, vdl.NComplex128(c128))
		fallthrough
	case u <= math.MaxInt64:
		result = append(result, int64(u), vdl.NInt64(u))
		fallthrough
	default:
		result = append(result, uint64(u), vdl.NUint64(u))
	}
	return result
}

// vvFromInt returns all *Values that can represent i without loss of precision.
func vvFromInt(i int64) (result []*vdl.Value) {
	b, u, f, c := byte(i), uint64(i), float64(i), complex(float64(i), 0)
	switch {
	case math.MinInt8 <= i && i <= math.MaxInt8:
		fallthrough
	case math.MinInt16 <= i && i <= math.MaxInt16:
		result = append(result,
			vdl.ZeroValue(vdl.Int16Type).AssignInt(i),
			vdl.ZeroValue(vdl.Int16TypeN).AssignInt(i))
		fallthrough
	case -1<<24 <= i && i <= 1<<24:
		result = append(result,
			vdl.ZeroValue(vdl.Float32Type).AssignFloat(f),
			vdl.ZeroValue(vdl.Float32TypeN).AssignFloat(f))
		result = append(result,
			vdl.ZeroValue(vdl.Complex64Type).AssignComplex(c),
			vdl.ZeroValue(vdl.Complex64TypeN).AssignComplex(c))
		fallthrough
	case math.MinInt32 <= i && i <= math.MaxInt32:
		result = append(result,
			vdl.ZeroValue(vdl.Int32Type).AssignInt(i),
			vdl.ZeroValue(vdl.Int32TypeN).AssignInt(i))
		fallthrough
	case -1<<53 <= i && i <= 1<<53:
		result = append(result,
			vdl.ZeroValue(vdl.Float64Type).AssignFloat(f),
			vdl.ZeroValue(vdl.Float64TypeN).AssignFloat(f))
		result = append(result,
			vdl.ZeroValue(vdl.Complex128Type).AssignComplex(c),
			vdl.ZeroValue(vdl.Complex128TypeN).AssignComplex(c))
		fallthrough
	default:
		result = append(result,
			vdl.ZeroValue(vdl.Int64Type).AssignInt(i),
			vdl.ZeroValue(vdl.Int64TypeN).AssignInt(i))
	}
	if i < 0 {
		return
	}
	switch {
	case i <= math.MaxUint8:
		result = append(result,
			vdl.ZeroValue(vdl.ByteType).AssignByte(b),
			vdl.ZeroValue(vdl.ByteTypeN).AssignByte(b))
		fallthrough
	case i <= math.MaxUint16:
		result = append(result,
			vdl.ZeroValue(vdl.Uint16Type).AssignUint(u),
			vdl.ZeroValue(vdl.Uint16TypeN).AssignUint(u))
		fallthrough
	case i <= math.MaxUint32:
		result = append(result,
			vdl.ZeroValue(vdl.Uint32Type).AssignUint(u),
			vdl.ZeroValue(vdl.Uint32TypeN).AssignUint(u))
		fallthrough
	default:
		result = append(result,
			vdl.ZeroValue(vdl.Uint64Type).AssignUint(u),
			vdl.ZeroValue(vdl.Uint64TypeN).AssignUint(u))
	}
	return
}

// rvFromInt returns all values that can represent i without loss of precision.
func rvFromInt(i int64) (result []interface{}) {
	c64, c128 := complex(float32(i), 0), complex(float64(i), 0)
	switch {
	case math.MinInt8 <= i && i <= math.MaxInt8:
		result = append(result, int8(i), vdl.NInt8(i))
		fallthrough
	case math.MinInt16 <= i && i <= math.MaxInt16:
		result = append(result, int16(i), vdl.NInt16(i))
		fallthrough
	case -1<<24 <= i && i <= 1<<24:
		result = append(result, float32(i), vdl.NFloat32(i))
		result = append(result, c64, vdl.NComplex64(c64))
		fallthrough
	case math.MinInt32 <= i && i <= math.MaxInt32:
		result = append(result, int32(i), vdl.NInt32(i))
		fallthrough
	case -1<<53 <= i && i <= 1<<53:
		result = append(result, float64(i), vdl.NFloat64(i))
		result = append(result, c128, vdl.NComplex128(c128))
		fallthrough
	default:
		result = append(result, int64(i), vdl.NInt64(i))
	}
	if i < 0 {
		return
	}
	switch {
	case i <= math.MaxUint8:
		result = append(result, uint8(i), vdl.NUint8(i))
		fallthrough
	case i <= math.MaxUint16:
		result = append(result, uint16(i), vdl.NUint16(i))
		fallthrough
	case i <= math.MaxUint32:
		result = append(result, uint32(i), vdl.NUint32(i))
		fallthrough
	default:
		result = append(result, uint64(i), vdl.NUint64(i))
	}
	return
}

func vvFloat(f float64) []*vdl.Value {
	c := complex(f, 0)
	return []*vdl.Value{
		vdl.ZeroValue(vdl.Float32Type).AssignFloat(f),
		vdl.ZeroValue(vdl.Float32TypeN).AssignFloat(f),
		vdl.ZeroValue(vdl.Float64Type).AssignFloat(f),
		vdl.ZeroValue(vdl.Float64TypeN).AssignFloat(f),
		vdl.ZeroValue(vdl.Complex64Type).AssignComplex(c),
		vdl.ZeroValue(vdl.Complex64TypeN).AssignComplex(c),
		vdl.ZeroValue(vdl.Complex128Type).AssignComplex(c),
		vdl.ZeroValue(vdl.Complex128TypeN).AssignComplex(c),
	}
}

func vvComplex(c complex128) []*vdl.Value {
	return []*vdl.Value{
		vdl.ZeroValue(vdl.Complex64Type).AssignComplex(c),
		vdl.ZeroValue(vdl.Complex64TypeN).AssignComplex(c),
		vdl.ZeroValue(vdl.Complex128Type).AssignComplex(c),
		vdl.ZeroValue(vdl.Complex128TypeN).AssignComplex(c),
	}
}

func rvFloat(f float64) []interface{} {
	c64, c128 := complex(float32(f), 0), complex(f, 0)
	return []interface{}{
		float32(f), vdl.NFloat32(f),
		float64(f), vdl.NFloat64(f),
		c64, vdl.NComplex64(c64),
		c128, vdl.NComplex128(c128),
	}
}

func rvComplex(c128 complex128) []interface{} {
	c64 := complex64(c128)
	return []interface{}{
		c64, vdl.NComplex64(c64),
		c128, vdl.NComplex128(c128),
	}
}

// Test successful conversions.  Each test contains a set of values and
// interfaces that are all equivalent and convertible to each other.
func TestConverter(t *testing.T) {
	tests := []struct {
		vv []*vdl.Value
		rv []interface{}
	}{
		{[]*vdl.Value{vvError1}, []interface{}{rvError1}},
		{[]*vdl.Value{vvError2}, []interface{}{rvError2}},
		{[]*vdl.Value{vvError3}, []interface{}{rvError3}},
		{vvBoolTrue, rvBoolTrue},
		{vvStrABC, rvStrABC},
		{vvFromUint(math.MaxUint8), rvFromUint(math.MaxUint8)},
		{vvFromUint(math.MaxUint16), rvFromUint(math.MaxUint16)},
		{vvFromUint(math.MaxUint32), rvFromUint(math.MaxUint32)},
		{vvFromUint(math.MaxUint64), rvFromUint(math.MaxUint64)},
		{vvFromInt(math.MaxInt8), rvFromInt(math.MaxInt8)},
		{vvFromInt(math.MaxInt16), rvFromInt(math.MaxInt16)},
		{vvFromInt(math.MaxInt32), rvFromInt(math.MaxInt32)},
		{vvFromInt(math.MaxInt64), rvFromInt(math.MaxInt64)},
		{vvFromInt(math.MinInt8), rvFromInt(math.MinInt8)},
		{vvFromInt(math.MinInt16), rvFromInt(math.MinInt16)},
		{vvFromInt(math.MinInt32), rvFromInt(math.MinInt32)},
		{vvFromInt(math.MinInt64), rvFromInt(math.MinInt64)},
		{vvFromInt(vdl.Float32MaxInt), rvFromInt(vdl.Float32MaxInt)},
		{vvFromInt(vdl.Float64MaxInt), rvFromInt(vdl.Float64MaxInt)},
		{vvFromInt(vdl.Float32MinInt), rvFromInt(vdl.Float32MinInt)},
		{vvFromInt(vdl.Float64MinInt), rvFromInt(vdl.Float64MinInt)},
		{vvTypeObjectBool, rvTypeObjectBool},
		{vvSeq123, rvSeq123},
		{vvSetMap123, rvSetMap123},
		{vvMap123FalseTrue, rvMap123FalseTrue},
		{vvSetMapStructXYZ, rvSetMapStructXYZ},
		{vvMapStructXYZFalseTrue, rvMapStructXYZFalseTrue},
		{vvMapStructXYZEmpty, rvMapStructXYZEmpty},
		{vvStructWXFalseTrue, rvStructWXFalseTrue},
		{vvMapStructVWX123, rvMapStructVWX123},
		{vvStructUV01, rvStructUV01},
		{nil, []interface{}{vdl.NNative(0)}},
		{nil, []interface{}{vdl.NNative(1)}},
		{nil, []interface{}{vdl.NNative(-(1 << 63))}},
		{nil, []interface{}{vdl.NNative((1 << 63) - 1)}},
		{nil, []interface{}{vdl.NUnionNative("A=false")}},
		{nil, []interface{}{vdl.NUnionNative("A=true")}},
		{nil, []interface{}{vdl.NUnionNative("B=123")}},
		{nil, []interface{}{vdl.NUnionNative("B=-123")}},
	}
	for _, test := range tests {
		testConverterWantSrc(t, vvrv{test.vv, test.rv}, vvrv{test.vv, test.rv})
	}
}

// Test successful conversions that drop and ignore fields in the dst struct.
func TestConverterStructDropIgnore(t *testing.T) {
	tests := []struct {
		vvWant []*vdl.Value
		rvWant []interface{}
		vvSrc  []*vdl.Value
		rvSrc  []interface{}
	}{
		{vvStructWXFalseTrue, rvStructWXFalseTrue, vvSetMapStructXYZ, rvSetMapStructXYZ},
		{vvStructUV01, rvStructUV01, vvMapStructVWX123, rvMapStructVWX123},
	}
	for _, test := range tests {
		testConverterWantSrc(t, vvrv{test.vvWant, test.rvWant}, vvrv{test.vvSrc, test.rvSrc})
	}
}

// Test successful conversions to and from union values.
func TestConverterUnion(t *testing.T) {
	// values for union component types
	vvTrue := vdl.BoolValue(true)
	vv123 := vdl.Int64Value(123)
	vvAbc := vdl.StringValue("Abc")
	vvStruct123 := vdl.ZeroValue(vdl.StructInt64TypeN)
	vvStruct123.StructField(0).Assign(vv123)
	rvTrue := bool(true)
	rv123 := int64(123)
	rvAbc := string("Abc")
	rvStruct123 := vdl.NStructInt64{123}
	// values for union{A bool;B string;C struct}
	vvTrueABC := vdl.ZeroValue(vdl.UnionABCTypeN).AssignUnionField(0, vvTrue)
	vvAbcABC := vdl.ZeroValue(vdl.UnionABCTypeN).AssignUnionField(1, vvAbc)
	vvStruct123ABC := vdl.ZeroValue(vdl.UnionABCTypeN).AssignUnionField(2, vvStruct123)
	rvTrueABC := vdl.NUnionABCA{rvTrue}
	rvAbcABC := vdl.NUnionABCB{rvAbc}
	rvStruct123ABC := vdl.NUnionABCC{rvStruct123}
	rvTrueABCi := vdl.NUnionABC(rvTrueABC)
	rvAbcABCi := vdl.NUnionABC(rvAbcABC)
	rvStruct123ABCi := vdl.NUnionABC(rvStruct123ABC)
	// values for union{B string;C struct;D int64}
	vvAbcBCD := vdl.ZeroValue(vdl.UnionBCDTypeN).AssignUnionField(0, vvAbc)
	vvStruct123BCD := vdl.ZeroValue(vdl.UnionBCDTypeN).AssignUnionField(1, vvStruct123)
	vv123BCD := vdl.ZeroValue(vdl.UnionBCDTypeN).AssignUnionField(2, vv123)
	rvAbcBCD := vdl.NUnionBCDB{rvAbc}
	rvStruct123BCD := vdl.NUnionBCDC{rvStruct123}
	rv123BCD := vdl.NUnionBCDD{rv123}
	rvAbcBCDi := vdl.NUnionBCD(rvAbcBCD)
	rvStruct123BCDi := vdl.NUnionBCD(rvStruct123BCD)
	rv123BCDi := vdl.NUnionBCD(rv123BCD)
	// values for union{X string;Y struct}, which has no Go equivalent.
	vvAbcXY := vdl.ZeroValue(vdl.UnionXYTypeN).AssignUnionField(0, vvAbc)
	vvStruct123XY := vdl.ZeroValue(vdl.UnionXYTypeN).AssignUnionField(1, vvStruct123)

	tests := []struct {
		vvWant *vdl.Value
		rvWant interface{}
		vvSrc  *vdl.Value
		rvSrc  interface{}
	}{
		// Convert source and target same union.
		{vvTrueABC, rvTrueABC, vvTrueABC, rvTrueABC},
		{vv123BCD, rv123BCD, vv123BCD, rv123BCD},
		{vvAbcABC, rvAbcABC, vvAbcABC, rvAbcABC},
		{vvAbcBCD, rvAbcBCD, vvAbcBCD, rvAbcBCD},
		{vvStruct123ABC, rvStruct123ABC, vvStruct123ABC, rvStruct123ABC},
		{vvStruct123BCD, rvStruct123BCD, vvStruct123BCD, rvStruct123BCD},
		// Same thing, but with pointers to the interface type.
		{vvTrueABC, &rvTrueABCi, vvTrueABC, &rvTrueABCi},
		{vv123BCD, &rv123BCDi, vv123BCD, &rv123BCD},
		{vvAbcABC, &rvAbcABCi, vvAbcABC, &rvAbcABC},
		{vvAbcBCD, &rvAbcBCDi, vvAbcBCD, &rvAbcBCD},
		{vvStruct123ABC, &rvStruct123ABCi, vvStruct123ABC, &rvStruct123ABCi},
		{vvStruct123BCD, &rvStruct123BCDi, vvStruct123BCD, &rvStruct123BCDi},

		// Convert source and target different union.
		{vvAbcABC, rvAbcABC, vvAbcBCD, rvAbcBCD},
		{vvAbcBCD, rvAbcBCD, vvAbcABC, rvAbcABC},
		{vvStruct123ABC, rvStruct123ABC, vvStruct123BCD, rvStruct123BCD},
		{vvStruct123BCD, rvStruct123BCD, vvStruct123ABC, rvStruct123ABC},
		// Same thing, but with pointers to the interface type.
		{vvAbcABC, &rvAbcABCi, vvAbcBCD, &rvAbcBCDi},
		{vvAbcBCD, &rvAbcBCDi, vvAbcABC, &rvAbcABCi},
		{vvStruct123ABC, &rvStruct123ABCi, vvStruct123BCD, &rvStruct123BCDi},
		{vvStruct123BCD, &rvStruct123BCDi, vvStruct123ABC, &rvStruct123ABCi},

		// Test unions that have no Go equivalent.
		{vvAbcXY, nil, vvAbcXY, nil},
		{vvStruct123XY, nil, vvStruct123XY, nil},
	}
	for _, test := range tests {
		testConverterWantSrc(t,
			vvrv{vvSlice(test.vvWant), rvSlice(test.rvWant)},
			vvrv{vvSlice(test.vvSrc), rvSlice(test.rvSrc)})
	}
}

func vvSlice(v *vdl.Value) []*vdl.Value {
	if v != nil {
		return []*vdl.Value{v}
	}
	return nil
}

func rvSlice(v interface{}) []interface{} {
	if v != nil {
		return []interface{}{v}
	}
	return nil
}

// Test successful conversions to and from nil values.
func TestConverterNil(t *testing.T) {
	vvNil := vdl.ZeroValue(vdl.AnyType)
	rvNil := new(interface{})
	vvNilError := vdl.ZeroValue(vdl.ErrorType)
	rvNilError := new(error)
	vvNilPtrStruct := vdl.ZeroValue(vdl.TypeOf((*vdl.NStructInt)(nil)))
	rvNilPtrStruct := (*vdl.NStructInt)(nil)
	vvStructNilStructField := vdl.ZeroValue(vdl.TypeOf(vdl.NStructOptionalStruct{}))
	rvStructNilStructField := vdl.NStructOptionalStruct{X: nil}
	vvStructNilAnyField := vdl.ZeroValue(vdl.TypeOf(vdl.NStructOptionalAny{}))
	rvStructNilAnyField := vdl.NStructOptionalAny{X: nil}
	tests := []struct {
		vvWant *vdl.Value
		rvWant interface{}
		vvSrc  *vdl.Value
		rvSrc  interface{}
	}{
		// Conversion source and target are the same.
		{vvNil, rvNil, vvNil, rvNil},
		{vvNilError, rvNilError, vvNilError, rvNilError},
		{vvNilPtrStruct, rvNilPtrStruct, vvNilPtrStruct, rvNilPtrStruct},
		{vvStructNilStructField, rvStructNilStructField, vvStructNilStructField, rvStructNilStructField},
		{vvStructNilAnyField, rvStructNilAnyField, vvStructNilAnyField, rvStructNilAnyField},
		// All typed nil targets may be converted from any(nil).
		{vvNilError, rvNilError, vvNil, rvNil},
		{vvNilPtrStruct, rvNilPtrStruct, vvNil, rvNil},
		{vvStructNilStructField, rvStructNilStructField, vvStructNilAnyField, rvStructNilAnyField},
	}
	for _, test := range tests {
		testConverterWantSrc(t,
			vvrv{[]*vdl.Value{test.vvWant}, []interface{}{test.rvWant}},
			vvrv{[]*vdl.Value{test.vvSrc}, []interface{}{test.rvSrc}})
	}
}

type vvrv struct {
	vv []*vdl.Value
	rv []interface{}
}

func testConverterWantSrc(t *testing.T, vvrvWant, vvrvSrc vvrv) {
	// We run each testConvert helper twice; the first call tests filling in a
	// zero dst, and the second call tests filling in a non-zero dst.

	// Tests of filling from *Value
	for _, vvSrc := range vvrvSrc.vv {
		for _, vvWant := range vvrvWant.vv {
			// Test filling *Value from *Value
			vvDst := vdl.ZeroValue(vvWant.Type())
			testConvert(t, "vv1", vvDst, vvSrc, vvWant, 0, false)
			testConvert(t, "vv2", vvDst, vvSrc, vvWant, 0, false)
		}
		for _, want := range vvrvWant.rv {
			// Test filling reflect.Value from *Value
			dst := reflect.New(reflect.TypeOf(want)).Interface()
			testConvert(t, "vv3", dst, vvSrc, want, 1, false)
			testConvert(t, "vv4", dst, vvSrc, want, 1, false)
		}
		// Test filling Any from *Value
		vvDst := vdl.ZeroValue(vdl.AnyType)
		testConvert(t, "vv5", vvDst, vvSrc, vdl.ZeroValue(vdl.AnyType).Assign(vvSrc), 0, false)
		testConvert(t, "vv6", vvDst, vvSrc, vdl.ZeroValue(vdl.AnyType).Assign(vvSrc), 0, false)
		// Test filling Optional from *Value
		if vvSrc.Type().CanBeOptional() {
			ttNil := vdl.OptionalType(vvSrc.Type())
			vvNil := vdl.ZeroValue(ttNil)
			vvOptWant := vdl.OptionalValue(vvSrc)
			testConvert(t, "vv7", vvNil, vvSrc, vvOptWant, 0, false)
			testConvert(t, "vv8", vvNil, vvSrc, vvOptWant, 0, false)
		}
		// Test filling **Value(nil) from *Value
		var vvValue *vdl.Value
		testConvert(t, "vv9", &vvValue, vvSrc, vvSrc, 1, false)
		testConvert(t, "vv10", &vvValue, vvSrc, vvSrc, 1, false)
		// Test filling interface{} from *Value
		var dst interface{}
		var want interface{} = vvSrc
		if !canCreateGoObject(vvSrc.Type()) {
			// We only run these tests if we *can't* create an actual Go object for
			// this type, so we'll end up with *vdl.Value.
			//
			// The problem is that we don't know what value to expect.  It seems
			// pointless to run a conversion to the rv src, since it's a
			// self-fulfilling test.  So we just skip this, and let the rv tests below
			// check Go object generation.
			//
			// TODO(toddw): Test with parallel rv and vv, to get better coverage.
			if vvSrc.Type().CanBeNil() && vvSrc.IsNil() {
				want = nil // filling interface{} from any(nil) yields nil.
			}
			testConvert(t, "vv11", &dst, vvSrc, want, 1, false)
			testConvert(t, "vv12", &dst, vvSrc, want, 1, false)
		}
		if vvSrc.Kind() == vdl.Struct {
			// Every struct may be converted to the empty struct
			testConvert(t, "vv13", vdl.ZeroValue(vdl.EmptyType), vvSrc, vdl.ZeroValue(vdl.EmptyType), 0, false)
			testConvert(t, "vv14", vdl.ZeroValue(vdl.EmptyTypeN), vvSrc, vdl.ZeroValue(vdl.EmptyTypeN), 0, false)
			var empty struct{}
			var emptyN vdl.NEmpty
			testConvert(t, "vv15", &empty, vvSrc, struct{}{}, 1, false)
			testConvert(t, "vv16", &emptyN, vvSrc, vdl.NEmpty{}, 1, false)
			// The empty struct may be converted to the zero value of any struct
			vvZeroSrc := vdl.ZeroValue(vvSrc.Type())
			testConvert(t, "vv17", vdl.ZeroValue(vvSrc.Type()), vdl.ZeroValue(vdl.EmptyType), vvZeroSrc, 0, false)
			testConvert(t, "vv18", vdl.ZeroValue(vvSrc.Type()), vdl.ZeroValue(vdl.EmptyTypeN), vvZeroSrc, 0, false)
			testConvert(t, "vv19", vdl.ZeroValue(vvSrc.Type()), struct{}{}, vvZeroSrc, 0, false)
			testConvert(t, "vv20", vdl.ZeroValue(vvSrc.Type()), vdl.NEmpty{}, vvZeroSrc, 0, false)
		}
	}

	// Tests of filling from reflect.Value
	for _, src := range vvrvSrc.rv {
		rtSrc := reflect.TypeOf(src)
		for _, vvWant := range vvrvWant.vv {
			// Test filling *Value from reflect.Value
			vvDst := vdl.ZeroValue(vvWant.Type())
			testConvert(t, "rv1", vvDst, src, vvWant, 0, false)
			testConvert(t, "rv2", vvDst, src, vvWant, 0, false)
		}
		for _, want := range vvrvWant.rv {
			// Test filling reflect.Value from reflect.Value
			dst := reflect.New(reflect.TypeOf(want)).Interface()
			testConvert(t, "rv3", dst, src, want, 1, false)
			testConvert(t, "rv4", dst, src, want, 1, false)
		}
		vvWant := vdl.ValueOf(src)
		// Test filling Any from reflect.Value
		vvDst := vdl.ZeroValue(vdl.AnyType)
		testConvert(t, "rv5", vvDst, src, vdl.ZeroValue(vdl.AnyType).Assign(vvWant), 0, true)
		testConvert(t, "rv6", vvDst, src, vdl.ZeroValue(vdl.AnyType).Assign(vvWant), 0, true)
		// Test filling **Value(nil) from reflect.Value
		var vvValue *vdl.Value
		testConvert(t, "rv7", &vvValue, src, vvWant, 1, false)
		testConvert(t, "rv8", &vvValue, src, vvWant, 1, false)
		// Test filling interface{} from reflect.Value
		var dst interface{}
		var want interface{} = src
		// Handle special-cases for want values.
		srcInt8, isInt8 := src.(int8)
		switch {
		case isInt8:
			// VDL represents int8 as int16, so set our expectations accordingly.
			want = int16(srcInt8)
		case rtSrc != rtPtrToType && rtSrc.ConvertibleTo(rtPtrToType):
			// VDL doesn't represent named TypeObject, so our converter automatically
			// creates *vdl.Type.
			want = reflect.ValueOf(src).Convert(rtPtrToType).Interface().(*vdl.Type)
		case vvWant.Kind() == vdl.Any && vvWant.IsNil():
			want = nil // filling interface{} from any(nil) yields nil.
		case !canCreateGoObject(vvWant.Type()):
			// We can't create an actual Go object for this type; e.g. perhaps it's
			// named, and isn't registered.  We should get a *vdl.Value back.
			want = vvWant
		}
		testConvert(t, "rv9", &dst, src, want, 1, true)
		testConvert(t, "rv10", &dst, src, want, 1, true)
		ttSrc, err := vdl.TypeFromReflect(rtSrc)
		if err != nil {
			t.Error(err)
			continue
		}
		if rtSrc.Kind() == reflect.Struct && ttSrc.Kind() != vdl.Union {
			// Every struct may be converted to the empty struct
			testConvert(t, "rv11", vdl.ZeroValue(vdl.EmptyType), src, vdl.ZeroValue(vdl.EmptyType), 0, false)
			testConvert(t, "rv12", vdl.ZeroValue(vdl.EmptyTypeN), src, vdl.ZeroValue(vdl.EmptyTypeN), 0, false)
			var empty struct{}
			var emptyN vdl.NEmpty
			testConvert(t, "rv13", &empty, src, struct{}{}, 1, false)
			testConvert(t, "rv14", &emptyN, src, vdl.NEmpty{}, 1, false)
			// The empty struct may be converted to the zero value of any struct
			rvZeroSrc := reflect.Zero(rtSrc).Interface()
			testConvert(t, "rv15", reflect.New(rtSrc).Interface(), vdl.ZeroValue(vdl.EmptyType), rvZeroSrc, 1, false)
			testConvert(t, "rv16", reflect.New(rtSrc).Interface(), vdl.ZeroValue(vdl.EmptyTypeN), rvZeroSrc, 1, false)
			testConvert(t, "rv17", reflect.New(rtSrc).Interface(), struct{}{}, rvZeroSrc, 1, false)
			testConvert(t, "rv18", reflect.New(rtSrc).Interface(), vdl.NEmpty{}, rvZeroSrc, 1, false)
		}
	}
}

// canCreateGoObject returns true iff we can create a regular Go object from a
// value of type tt.  We can create a real Go object if the Go type for tt has
// been registered, or if tt is the special-cased error type.
func canCreateGoObject(tt *vdl.Type) bool {
	return vdl.TypeToReflect(tt) != nil || tt == vdl.ErrorType || tt == vdl.ErrorType.Elem()
}

func testConvert(t *testing.T, prefix string, dst, src, want interface{}, deref int, optWant bool) {
	const ptrDepth = 3
	rvDst := reflect.ValueOf(dst)
	for dstptrs := 0; dstptrs < ptrDepth; dstptrs++ {
		rvSrc := reflect.ValueOf(src)
		for srcptrs := 0; srcptrs < ptrDepth; srcptrs++ {
			tname := fmt.Sprintf("%s ReflectTarget(%v).From(%v)", prefix, rvDst.Type(), rvSrc.Type())
			// This is tricky - if optWant is set, we might need to change the want
			// value to become optional or non-optional.
			eWant, rvWant, ttWant := want, reflect.ValueOf(want), vdl.TypeOf(want)
			if optWant {
				vvWant, wantIsVV := want.(*vdl.Value)
				if srcptrs > 0 {
					if wantIsVV {
						switch {
						case vvWant.Kind() == vdl.Any && !vvWant.IsNil() && vvWant.Elem().Type().CanBeOptional():
							// Turn any(struct{...}) into any(?struct{...})
							eWant = vdl.ZeroValue(vdl.AnyType).Assign(vdl.OptionalValue(vvWant.Elem()))
						case vvWant.Type().CanBeOptional():
							// Turn struct{...} into ?struct{...}
							eWant = vdl.OptionalValue(vvWant)
						}
					} else if (ttWant.Kind() == vdl.Optional || ttWant.CanBeOptional()) && rvWant.Kind() != reflect.Ptr {
						// Add a pointer to non-pointers that can be optional.
						rvPtrWant := reflect.New(rvWant.Type())
						rvPtrWant.Elem().Set(rvWant)
						eWant = rvPtrWant.Interface()
					}
				}
				if !wantIsVV && ttWant.Kind() != vdl.TypeObject && !ttWant.CanBeOptional() && rvWant.Kind() == reflect.Ptr && !rvWant.IsNil() {
					// Remove a pointer from  anything that can't be optional.
					eWant = rvWant.Elem().Interface()
				}
			}
			target, err := vdl.ReflectTarget(rvDst)
			vdl.ExpectErr(t, err, "", tname)
			err = vdl.FromReflect(target, rvSrc)
			vdl.ExpectErr(t, err, "", tname)
			expectConvert(t, tname, dst, eWant, deref)
			// Next iteration adds a pointer to src.
			rvNewSrc := reflect.New(rvSrc.Type())
			rvNewSrc.Elem().Set(rvSrc)
			rvSrc = rvNewSrc
		}
		// Next iteration adds a pointer to dst.
		rvNewDst := reflect.New(rvDst.Type())
		rvNewDst.Elem().Set(rvDst)
		rvDst = rvNewDst
	}
}

func expectConvert(t *testing.T, tname string, got, want interface{}, deref int) {
	rvGot := reflect.ValueOf(got)
	for d := 0; d < deref; d++ {
		if rvGot.Kind() != reflect.Ptr || rvGot.IsNil() {
			t.Errorf("%s can't deref %d %T %v", deref, got, got)
			return
		}
		rvGot = rvGot.Elem()
	}
	got = rvGot.Interface()
	vvGot, ok1 := got.(*vdl.Value)
	vvWant, ok2 := want.(*vdl.Value)
	if ok1 && ok2 {
		if !vdl.EqualValue(vvGot, vvWant) {
			t.Errorf("%s\nGOT  %v\nWANT %v", tname, vvGot, vvWant)
		}
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s\nGOT  %#v\nWANT %#v", tname, got, want)
	}
}

// Test failed conversions.
func TestConverterError(t *testing.T) {
	tests := []struct {
		ttDst []*vdl.Type
		rtDst []reflect.Type
		vvSrc []*vdl.Value
		rvSrc []interface{}
	}{
		{ttOtherThan(ttAllTypes, ttBools), rtOtherThan(rtAllTypes, rtBools),
			vvBoolTrue, rvBoolTrue},
		{ttOtherThan(ttAllTypes, ttStrs), rtOtherThan(rtAllTypes, rtStrs),
			vvStrABC, rvStrABC},
		{ttOtherThan(ttAllTypes, ttTypeObjects), rtOtherThan(rtAllTypes, rtTypeObjects),
			vvTypeObjectBool, rvTypeObjectBool},
		{ttOtherThan(ttAllTypes, ttSeq123), rtOtherThan(rtAllTypes, rtSeq123),
			vvSeq123, rvSeq123},
		{ttOtherThan(ttAllTypes, ttSetMap123), rtOtherThan(rtAllTypes, rtSetMap123),
			vvSetMap123, rvSetMap123},
		{ttOtherThan(ttAllTypes, ttSetMapStructXYZ), rtOtherThan(rtAllTypes, rtSetMapStructXYZ),
			vvSetMapStructXYZ, rvSetMapStructXYZ},
		{ttOtherThan(ttAllTypes, ttMapStructVWXNum), rtOtherThan(rtAllTypes, rtMapStructVWXNum),
			vvMapStructVWX123, rvMapStructVWX123},
		// Test invalid conversions to set types
		{ttSet123, rtSet123, vvMap123FalseTrue, rvMap123FalseTrue},
		{ttSetXYZ, rtSetXYZ, vvMapStructXYZFalseTrue, rvMapStructXYZFalseTrue},
		{ttSetMapStructXYZ, rtSetMapStructXYZ, vvMapStructXYZEmpty, rvMapStructXYZEmpty},
		// Test invalid conversions to struct types: mismatched field types
		{ttStructXYZBool, rtStructXYZBool, vvMapStructVWX123, rvMapStructVWX123},
		{ttStructVWXNum, rtStructVWXNum, vvSetMapStructXYZ, rvSetMapStructXYZ},
		// Test invalid conversions to struct types: no fields in common
		{ttStructWXBool, rtStructWXBool, vvStructUV01, rvStructUV01},
		{ttStructUVNum, rtStructUVNum, vvStructWXFalseTrue, rvStructWXFalseTrue},
		// Test uint values one past the max bound.
		{ttOtherThan(ttIntegers, ttTypes(vvFromUint(math.MaxUint8))),
			rtOtherThan(rtIntegers, rtTypes(rvFromUint(math.MaxUint8))),
			vvFromUint(math.MaxUint8 + 1),
			rvFromUint(math.MaxUint8 + 1)},
		{ttOtherThan(ttIntegers, ttTypes(vvFromUint(math.MaxUint16))),
			rtOtherThan(rtIntegers, rtTypes(rvFromUint(math.MaxUint16))),
			vvFromUint(math.MaxUint16 + 1),
			rvFromUint(math.MaxUint16 + 1)},
		{ttOtherThan(ttIntegers, ttTypes(vvFromUint(math.MaxUint32))),
			rtOtherThan(rtIntegers, rtTypes(rvFromUint(math.MaxUint32))),
			vvFromUint(math.MaxUint32 + 1),
			rvFromUint(math.MaxUint32 + 1)},
		// Test int values one past the max bound.
		{ttOtherThan(ttIntegers, ttTypes(vvFromUint(math.MaxInt8))),
			rtOtherThan(rtIntegers, rtTypes(rvFromUint(math.MaxInt8))),
			vvFromUint(math.MaxInt8 + 1),
			rvFromUint(math.MaxInt8 + 1)},
		{ttOtherThan(ttIntegers, ttTypes(vvFromUint(math.MaxInt16))),
			rtOtherThan(rtIntegers, rtTypes(rvFromUint(math.MaxInt16))),
			vvFromUint(math.MaxInt16 + 1),
			rvFromUint(math.MaxInt16 + 1)},
		{ttOtherThan(ttIntegers, ttTypes(vvFromUint(math.MaxInt32))),
			rtOtherThan(rtIntegers, rtTypes(rvFromUint(math.MaxInt32))),
			vvFromUint(math.MaxInt32 + 1),
			rvFromUint(math.MaxInt32 + 1)},
		{ttOtherThan(ttIntegers, ttTypes(vvFromUint(math.MaxInt64))),
			rtOtherThan(rtIntegers, rtTypes(rvFromUint(math.MaxInt64))),
			vvFromUint(math.MaxInt64 + 1),
			rvFromUint(math.MaxInt64 + 1)},
		// Test int values one past the min bound.
		{ttOtherThan(ttIntegers, ttTypes(vvFromInt(math.MinInt8))),
			rtOtherThan(rtIntegers, rtTypes(rvFromInt(math.MinInt8))),
			vvFromInt(math.MinInt8 - 1),
			rvFromInt(math.MinInt8 - 1)},
		{ttOtherThan(ttIntegers, ttTypes(vvFromInt(math.MinInt16))),
			rtOtherThan(rtIntegers, rtTypes(rvFromInt(math.MinInt16))),
			vvFromInt(math.MinInt16 - 1),
			rvFromInt(math.MinInt16 - 1)},
		{ttOtherThan(ttIntegers, ttTypes(vvFromInt(math.MinInt32))),
			rtOtherThan(rtIntegers, rtTypes(rvFromInt(math.MinInt32))),
			vvFromInt(math.MinInt32 - 1),
			rvFromInt(math.MinInt32 - 1)},
		// Test int to float max bound.
		{ttJoin(ttFloat32s, ttComplex64s), rtJoin(rtFloat32s, rtComplex64s),
			vvOnlyFrom(vvFromInt(vdl.Float32MaxInt+1), ttIntegers),
			rvOnlyFrom(rvFromInt(vdl.Float32MaxInt+1), rtIntegers)},
		{ttJoin(ttFloat64s, ttComplex128s), rtJoin(rtFloat64s, rtComplex128s),
			vvOnlyFrom(vvFromInt(vdl.Float64MaxInt+1), ttIntegers),
			rvOnlyFrom(rvFromInt(vdl.Float64MaxInt+1), rtIntegers)},
		// Test int to float min bound.
		{ttJoin(ttFloat32s, ttComplex64s), rtJoin(rtFloat32s, rtComplex64s),
			vvOnlyFrom(vvFromInt(vdl.Float32MinInt-1), ttIntegers),
			rvOnlyFrom(rvFromInt(vdl.Float32MinInt-1), rtIntegers)},
		{ttJoin(ttFloat64s, ttComplex128s), rtJoin(rtFloat64s, rtComplex128s),
			vvOnlyFrom(vvFromInt(vdl.Float64MinInt-1), ttIntegers),
			rvOnlyFrom(rvFromInt(vdl.Float64MinInt-1), rtIntegers)},
		// Test negative uints, fractional integers, imaginary non-complex numbers.
		{ttUints, rtUints, vvFromInt(-1), rvFromInt(-1)},
		{ttIntegers, rtIntegers, vvFloat(1.5), rvFloat(1.5)},
		{ttOtherThan(ttNumbers, ttComplexes), rtOtherThan(rtNumbers, rtComplexes),
			vvComplex(1 + 2i), rvComplex(1 + 2i)},
	}
	for _, test := range tests {
		for _, ttDst := range test.ttDst {
			tname := fmt.Sprintf("ValueTarget(%v)", ttDst)
			vvDst := vdl.ZeroValue(ttDst)
			target, err := vdl.ValueTarget(vvDst)
			if !vdl.ExpectErr(t, err, "", tname) {
				continue
			}
			for _, vvSrc := range test.vvSrc {
				if err := vdl.FromValue(target, vvSrc); err == nil {
					t.Errorf("%s FromValue(%v) got %v, want error", tname, vvSrc.Type(), vvDst)
				}
			}
			for _, src := range test.rvSrc {
				rvSrc := reflect.ValueOf(src)
				if err := vdl.FromReflect(target, rvSrc); err == nil {
					t.Errorf("%s FromReflect(%v) got %v, want error", tname, rvSrc.Type(), vvDst)
				}
			}
		}
		for _, rtDst := range test.rtDst {
			tname := fmt.Sprintf("ReflectTarget(%v)", rtDst)
			rvDst := reflect.New(rtDst)
			target, err := vdl.ReflectTarget(rvDst)
			if !vdl.ExpectErr(t, err, "", tname) {
				continue
			}
			got := rvDst.Elem().Interface()
			for _, vvSrc := range test.vvSrc {
				if err := vdl.FromValue(target, vvSrc); err == nil {
					t.Errorf("%s FromValue(%v) got %v, want error", tname, vvSrc.Type(), got)
				}
			}
			for _, src := range test.rvSrc {
				rvSrc := reflect.ValueOf(src)
				if err := vdl.FromReflect(target, rvSrc); err == nil {
					t.Errorf("%s FromReflect(%v) got %v, want error", tname, rvSrc.Type(), got)
				}
			}
		}
	}
}
