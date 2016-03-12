// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: time

// Package time defines standard representations of absolute and relative times.
//
// The representations described below are required to provide wire
// compatibility between different programming environments.  Generated code for
// different environments typically provide automatic conversions into native
// representations, for simpler idiomatic usage.
package time

import (
	"fmt"
	"time"
	"v.io/v23/vdl"
)

// Time represents an absolute point in time with up to nanosecond precision.
//
// Time is represented as the duration before or after a fixed epoch.  The zero
// Time represents the epoch 0001-01-01T00:00:00.000000000Z.  This uses the
// proleptic Gregorian calendar; the calendar runs on an exact 400 year cycle.
// Leap seconds are "smeared", ensuring that no leap second table is necessary
// for interpretation.
//
// This is similar to Go time.Time, but always in the UTC location.
// http://golang.org/pkg/time/#Time
//
// This is similar to conventional "unix time", but with the epoch defined at
// year 1 rather than year 1970.  This allows the zero Time to be used as a
// natural sentry, since it isn't a valid time for many practical applications.
// http://en.wikipedia.org/wiki/Unix_time
type Time struct {
	Seconds int64
	Nanos   int32
}

func (Time) __VDLReflect(struct {
	Name string `vdl:"time.Time"`
}) {
}

func (m *Time) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	__VDLEnsureNativeBuilt()
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Seconds")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromInt(int64(m.Seconds), vdl.Int64Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Nanos")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromInt(int64(m.Nanos), vdl.Int32Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Time) MakeVDLTarget() vdl.Target {
	return nil
}

type TimeTarget struct {
	Value     *time.Time
	wireValue Time
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *TimeTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	if !vdl.Compatible(tt, __VDLType_time_Time) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_time_Time)
	}
	return t, nil
}
func (t *TimeTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Seconds":
		val, err := &vdl.Int64Target{Value: &t.wireValue.Seconds}, error(nil)
		return nil, val, err
	case "Nanos":
		val, err := &vdl.Int32Target{Value: &t.wireValue.Nanos}, error(nil)
		return nil, val, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_time_Time)
	}
}
func (t *TimeTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *TimeTarget) FinishFields(_ vdl.FieldsTarget) error {

	if err := TimeToNative(t.wireValue, t.Value); err != nil {
		return err
	}
	return nil
}

// Duration represents the elapsed duration between two points in time, with
// up to nanosecond precision.
type Duration struct {
	// Seconds represents the seconds in the duration.  The range is roughly
	// +/-290 billion years, larger than the estimated age of the universe.
	Seconds int64
	// Nanos represents the fractions of a second at nanosecond resolution.  Must
	// be in the inclusive range between +/-999,999,999.
	//
	// In normalized form, durations less than one second are represented with 0
	// Seconds and +/-Nanos.  For durations one second or more, the sign of Nanos
	// must match Seconds, or be 0.
	Nanos int32
}

func (Duration) __VDLReflect(struct {
	Name string `vdl:"time.Duration"`
}) {
}

func (m *Duration) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	__VDLEnsureNativeBuilt()
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Seconds")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromInt(int64(m.Seconds), vdl.Int64Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Nanos")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromInt(int64(m.Nanos), vdl.Int32Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Duration) MakeVDLTarget() vdl.Target {
	return nil
}

type DurationTarget struct {
	Value     *time.Duration
	wireValue Duration
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *DurationTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	if !vdl.Compatible(tt, __VDLType_time_Duration) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_time_Duration)
	}
	return t, nil
}
func (t *DurationTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Seconds":
		val, err := &vdl.Int64Target{Value: &t.wireValue.Seconds}, error(nil)
		return nil, val, err
	case "Nanos":
		val, err := &vdl.Int32Target{Value: &t.wireValue.Nanos}, error(nil)
		return nil, val, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_time_Duration)
	}
}
func (t *DurationTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *DurationTarget) FinishFields(_ vdl.FieldsTarget) error {

	if err := DurationToNative(t.wireValue, t.Value); err != nil {
		return err
	}
	return nil
}

// WireDeadline represents the deadline for an operation, where the operation is
// expected to finish before the deadline.  The intended usage is for a client
// to set a deadline on an operation, say one minute from "now", and send the
// deadline to a server.  The server is expected to finish the operation before
// the deadline.
//
// On a single device, it is simplest to represent a deadline as an absolute
// time; when the time now reaches the deadline, the deadline has expired.
// However when sending a deadline between devices with potential clock skew, it
// is often more robust to represent the deadline as a duration from "now".  The
// sender computes the duration from its notion of "now", while the receiver
// computes the absolute deadline from its own notion of "now".
//
// This representation doesn't account for propagation delay, but does ensure
// that the deadline used by the receiver is no earlier than the deadline
// intended by the client.  In many common scenarios the propagation delay is
// small compared to the potential clock skew, making this a simple but
// effective approach.
//
// WireDeadline typically has a native representation called Deadline that is an
// absolute Time, which automatically performs the sender and receiver
// conversions from "now".
type WireDeadline struct {
	// FromNow represents the deadline as a duration from "now".
	FromNow time.Duration
	// NoDeadline indicates there is no deadline; the analogous sentry for the
	// native Deadline is the zero Time.
	NoDeadline bool
}

func (WireDeadline) __VDLReflect(struct {
	Name string `vdl:"time.WireDeadline"`
}) {
}

func (m *WireDeadline) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	__VDLEnsureNativeBuilt()
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	var wireValue2 Duration
	if err := DurationFromNative(&wireValue2, m.FromNow); err != nil {
		return err
	}

	keyTarget3, fieldTarget4, err := fieldsTarget1.StartField("FromNow")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := wireValue2.FillVDLTarget(fieldTarget4, __VDLType_time_Duration); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget3, fieldTarget4); err != nil {
			return err
		}
	}
	keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("NoDeadline")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget6.FromBool(bool(m.NoDeadline), vdl.BoolType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *WireDeadline) MakeVDLTarget() vdl.Target {
	return nil
}

type WireDeadlineTarget struct {
	Value     *Deadline
	wireValue WireDeadline
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *WireDeadlineTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	if !vdl.Compatible(tt, __VDLType_time_WireDeadline) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_time_WireDeadline)
	}
	return t, nil
}
func (t *WireDeadlineTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "FromNow":
		val, err := &DurationTarget{Value: &t.wireValue.FromNow}, error(nil)
		return nil, val, err
	case "NoDeadline":
		val, err := &vdl.BoolTarget{Value: &t.wireValue.NoDeadline}, error(nil)
		return nil, val, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_time_WireDeadline)
	}
}
func (t *WireDeadlineTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *WireDeadlineTarget) FinishFields(_ vdl.FieldsTarget) error {

	if err := WireDeadlineToNative(t.wireValue, t.Value); err != nil {
		return err
	}
	return nil
}

func init() {
	vdl.RegisterNative(DurationToNative, DurationFromNative)
	vdl.RegisterNative(TimeToNative, TimeFromNative)
	vdl.RegisterNative(WireDeadlineToNative, WireDeadlineFromNative)
	vdl.Register((*Time)(nil))
	vdl.Register((*Duration)(nil))
	vdl.Register((*WireDeadline)(nil))
}

// Type-check Duration conversion functions.
var _ func(Duration, *time.Duration) error = DurationToNative
var _ func(*Duration, time.Duration) error = DurationFromNative

// Type-check Time conversion functions.
var _ func(Time, *time.Time) error = TimeToNative
var _ func(*Time, time.Time) error = TimeFromNative

// Type-check WireDeadline conversion functions.
var _ func(WireDeadline, *Deadline) error = WireDeadlineToNative
var _ func(*WireDeadline, Deadline) error = WireDeadlineFromNative

var __VDLType1 *vdl.Type

func __VDLType1_gen() *vdl.Type {
	__VDLType1Builder := vdl.TypeBuilder{}

	__VDLType11 := __VDLType1Builder.Optional()
	__VDLType12 := __VDLType1Builder.Struct()
	__VDLType13 := __VDLType1Builder.Named("time.Duration").AssignBase(__VDLType12)
	__VDLType14 := vdl.Int64Type
	__VDLType12.AppendField("Seconds", __VDLType14)
	__VDLType15 := vdl.Int32Type
	__VDLType12.AppendField("Nanos", __VDLType15)
	__VDLType11.AssignElem(__VDLType13)
	__VDLType1Builder.Build()
	__VDLType1v, err := __VDLType11.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType1v
}
func init() {
	__VDLType1 = __VDLType1_gen()
}

var __VDLType0 *vdl.Type

func __VDLType0_gen() *vdl.Type {
	__VDLType0Builder := vdl.TypeBuilder{}

	__VDLType01 := __VDLType0Builder.Optional()
	__VDLType02 := __VDLType0Builder.Struct()
	__VDLType03 := __VDLType0Builder.Named("time.Time").AssignBase(__VDLType02)
	__VDLType04 := vdl.Int64Type
	__VDLType02.AppendField("Seconds", __VDLType04)
	__VDLType05 := vdl.Int32Type
	__VDLType02.AppendField("Nanos", __VDLType05)
	__VDLType01.AssignElem(__VDLType03)
	__VDLType0Builder.Build()
	__VDLType0v, err := __VDLType01.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType0v
}
func init() {
	__VDLType0 = __VDLType0_gen()
}

var __VDLType2 *vdl.Type

func __VDLType2_gen() *vdl.Type {
	__VDLType2Builder := vdl.TypeBuilder{}

	__VDLType21 := __VDLType2Builder.Optional()
	__VDLType22 := __VDLType2Builder.Struct()
	__VDLType23 := __VDLType2Builder.Named("time.WireDeadline").AssignBase(__VDLType22)
	__VDLType24 := __VDLType2Builder.Struct()
	__VDLType25 := __VDLType2Builder.Named("time.Duration").AssignBase(__VDLType24)
	__VDLType26 := vdl.Int64Type
	__VDLType24.AppendField("Seconds", __VDLType26)
	__VDLType27 := vdl.Int32Type
	__VDLType24.AppendField("Nanos", __VDLType27)
	__VDLType22.AppendField("FromNow", __VDLType25)
	__VDLType28 := vdl.BoolType
	__VDLType22.AppendField("NoDeadline", __VDLType28)
	__VDLType21.AssignElem(__VDLType23)
	__VDLType2Builder.Build()
	__VDLType2v, err := __VDLType21.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType2v
}
func init() {
	__VDLType2 = __VDLType2_gen()
}

var __VDLType_time_Duration *vdl.Type

func __VDLType_time_Duration_gen() *vdl.Type {
	__VDLType_time_DurationBuilder := vdl.TypeBuilder{}

	__VDLType_time_Duration1 := __VDLType_time_DurationBuilder.Struct()
	__VDLType_time_Duration2 := __VDLType_time_DurationBuilder.Named("time.Duration").AssignBase(__VDLType_time_Duration1)
	__VDLType_time_Duration3 := vdl.Int64Type
	__VDLType_time_Duration1.AppendField("Seconds", __VDLType_time_Duration3)
	__VDLType_time_Duration4 := vdl.Int32Type
	__VDLType_time_Duration1.AppendField("Nanos", __VDLType_time_Duration4)
	__VDLType_time_DurationBuilder.Build()
	__VDLType_time_Durationv, err := __VDLType_time_Duration2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_time_Durationv
}
func init() {
	__VDLType_time_Duration = __VDLType_time_Duration_gen()
}

var __VDLType_time_Time *vdl.Type

func __VDLType_time_Time_gen() *vdl.Type {
	__VDLType_time_TimeBuilder := vdl.TypeBuilder{}

	__VDLType_time_Time1 := __VDLType_time_TimeBuilder.Struct()
	__VDLType_time_Time2 := __VDLType_time_TimeBuilder.Named("time.Time").AssignBase(__VDLType_time_Time1)
	__VDLType_time_Time3 := vdl.Int64Type
	__VDLType_time_Time1.AppendField("Seconds", __VDLType_time_Time3)
	__VDLType_time_Time4 := vdl.Int32Type
	__VDLType_time_Time1.AppendField("Nanos", __VDLType_time_Time4)
	__VDLType_time_TimeBuilder.Build()
	__VDLType_time_Timev, err := __VDLType_time_Time2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_time_Timev
}
func init() {
	__VDLType_time_Time = __VDLType_time_Time_gen()
}

var __VDLType_time_WireDeadline *vdl.Type

func __VDLType_time_WireDeadline_gen() *vdl.Type {
	__VDLType_time_WireDeadlineBuilder := vdl.TypeBuilder{}

	__VDLType_time_WireDeadline1 := __VDLType_time_WireDeadlineBuilder.Struct()
	__VDLType_time_WireDeadline2 := __VDLType_time_WireDeadlineBuilder.Named("time.WireDeadline").AssignBase(__VDLType_time_WireDeadline1)
	__VDLType_time_WireDeadline3 := __VDLType_time_WireDeadlineBuilder.Struct()
	__VDLType_time_WireDeadline4 := __VDLType_time_WireDeadlineBuilder.Named("time.Duration").AssignBase(__VDLType_time_WireDeadline3)
	__VDLType_time_WireDeadline5 := vdl.Int64Type
	__VDLType_time_WireDeadline3.AppendField("Seconds", __VDLType_time_WireDeadline5)
	__VDLType_time_WireDeadline6 := vdl.Int32Type
	__VDLType_time_WireDeadline3.AppendField("Nanos", __VDLType_time_WireDeadline6)
	__VDLType_time_WireDeadline1.AppendField("FromNow", __VDLType_time_WireDeadline4)
	__VDLType_time_WireDeadline7 := vdl.BoolType
	__VDLType_time_WireDeadline1.AppendField("NoDeadline", __VDLType_time_WireDeadline7)
	__VDLType_time_WireDeadlineBuilder.Build()
	__VDLType_time_WireDeadlinev, err := __VDLType_time_WireDeadline2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_time_WireDeadlinev
}
func init() {
	__VDLType_time_WireDeadline = __VDLType_time_WireDeadline_gen()
}
func __VDLEnsureNativeBuilt() {
	if __VDLType1 == nil {
		__VDLType1 = __VDLType1_gen()
	}
	if __VDLType0 == nil {
		__VDLType0 = __VDLType0_gen()
	}
	if __VDLType2 == nil {
		__VDLType2 = __VDLType2_gen()
	}
	if __VDLType_time_Duration == nil {
		__VDLType_time_Duration = __VDLType_time_Duration_gen()
	}
	if __VDLType_time_Time == nil {
		__VDLType_time_Time = __VDLType_time_Time_gen()
	}
	if __VDLType_time_WireDeadline == nil {
		__VDLType_time_WireDeadline = __VDLType_time_WireDeadline_gen()
	}
}
