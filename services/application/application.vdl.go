// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: application

// Package application defines types for describing applications.
package application

import (
	"time"
	"v.io/v23/security"
	"v.io/v23/vdl"
	vdltime "v.io/v23/vdlroot/time"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// SignedFile represents a file accompanied by a signature of its contents.
type SignedFile struct {
	//  File is the object name of the file.
	File string
	// Signature represents a signature on the sha256 hash of the file
	// contents by the publisher principal.
	Signature security.Signature
}

func (SignedFile) VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/application.SignedFile"`
}) {
}

func (x SignedFile) VDLIsZero() bool {
	if x.File != "" {
		return false
	}
	if !x.Signature.VDLIsZero() {
		return false
	}
	return true
}

func (x SignedFile) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_1); err != nil {
		return err
	}
	if x.File != "" {
		if err := enc.NextFieldValueString(0, vdl.StringType, x.File); err != nil {
			return err
		}
	}
	if !x.Signature.VDLIsZero() {
		if err := enc.NextField(1); err != nil {
			return err
		}
		if err := x.Signature.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *SignedFile) VDLRead(dec vdl.Decoder) error {
	*x = SignedFile{}
	if err := dec.StartValue(__VDLType_struct_1); err != nil {
		return err
	}
	decType := dec.Type()
	for {
		index, err := dec.NextField()
		switch {
		case err != nil:
			return err
		case index == -1:
			return dec.FinishValue()
		}
		if decType != __VDLType_struct_1 {
			index = __VDLType_struct_1.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.File = value
			}
		case 1:
			if err := x.Signature.VDLRead(dec); err != nil {
				return err
			}
		}
	}
}

// Packages represents a set of packages. The map key is the local
// file/directory name, relative to the instance's packages directory, where the
// package should be installed. For archives, this name represents a directory
// into which the archive is to be extracted, and for regular files it
// represents the name for the file.  The map value is the package
// specification.
//
// Each object's media type determines how to install it.
//
// For example, with key=pkg1,value=SignedFile{File:binaryrepo/configfiles} (an
// archive), the "configfiles" package will be installed under the "pkg1"
// directory. With key=pkg2,value=SignedFile{File:binaryrepo/binfile} (a
// binary), the "binfile" file will be installed as the "pkg2" file.
//
// The keys must be valid file/directory names, without path separators.
//
// Any number of packages may be specified.
type Packages map[string]SignedFile

func (Packages) VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/application.Packages"`
}) {
}

func (x Packages) VDLIsZero() bool {
	return len(x) == 0
}

func (x Packages) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_map_3); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for key, elem := range x {
		if err := enc.NextEntryValueString(vdl.StringType, key); err != nil {
			return err
		}
		if err := elem.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Packages) VDLRead(dec vdl.Decoder) error {
	if err := dec.StartValue(__VDLType_map_3); err != nil {
		return err
	}
	var tmpMap Packages
	if len := dec.LenHint(); len > 0 {
		tmpMap = make(Packages, len)
	}
	for {
		switch done, key, err := dec.NextEntryValueString(); {
		case err != nil:
			return err
		case done:
			*x = tmpMap
			return dec.FinishValue()
		default:
			var elem SignedFile
			if err := elem.VDLRead(dec); err != nil {
				return err
			}
			if tmpMap == nil {
				tmpMap = make(Packages)
			}
			tmpMap[key] = elem
		}
	}
}

// Envelope is a collection of metadata that describes an application.
type Envelope struct {
	// Title is the publisher-assigned application title.  Application
	// installations with the same title are considered as belonging to the
	// same application by the application management system.
	//
	// A change in the title signals a new application.
	Title string
	// Args is an array of command-line arguments to be used when executing
	// the binary.
	Args []string
	// Binary identifies the application binary.
	Binary SignedFile
	// Publisher represents the set of blessings that have been bound to
	// the principal who published this binary.
	Publisher security.Blessings
	// Env is an array that stores the environment variable values to be
	// used when executing the binary.
	Env []string
	// Packages is the set of packages to install on the local filesystem
	// before executing the binary
	Packages Packages
	// Restarts specifies how many times the device manager will attempt
	// to automatically restart an application that has crashed before
	// giving up and marking the application as NotRunning.
	Restarts int32
	// RestartTimeWindow is the time window within which an
	// application exit is considered a crash that counts against the
	// Restarts budget. If the application crashes after less than
	// RestartTimeWindow time for Restarts consecutive times, the
	// application is marked NotRunning and no more restart attempts
	// are made. If the application has run continuously for more
	// than RestartTimeWindow, subsequent crashes will again benefit
	// from up to Restarts restarts (that is, the Restarts budget is
	// reset by a successful run of at least RestartTimeWindow
	// duration).
	RestartTimeWindow time.Duration
}

func (Envelope) VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/application.Envelope"`
}) {
}

func (x Envelope) VDLIsZero() bool {
	if x.Title != "" {
		return false
	}
	if len(x.Args) != 0 {
		return false
	}
	if !x.Binary.VDLIsZero() {
		return false
	}
	if !x.Publisher.IsZero() {
		return false
	}
	if len(x.Env) != 0 {
		return false
	}
	if len(x.Packages) != 0 {
		return false
	}
	if x.Restarts != 0 {
		return false
	}
	if x.RestartTimeWindow != 0 {
		return false
	}
	return true
}

func (x Envelope) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_4); err != nil {
		return err
	}
	if x.Title != "" {
		if err := enc.NextFieldValueString(0, vdl.StringType, x.Title); err != nil {
			return err
		}
	}
	if len(x.Args) != 0 {
		if err := enc.NextField(1); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_1(enc, x.Args); err != nil {
			return err
		}
	}
	if !x.Binary.VDLIsZero() {
		if err := enc.NextField(2); err != nil {
			return err
		}
		if err := x.Binary.VDLWrite(enc); err != nil {
			return err
		}
	}
	if !x.Publisher.IsZero() {
		if err := enc.NextField(3); err != nil {
			return err
		}
		var wire security.WireBlessings
		if err := security.WireBlessingsFromNative(&wire, x.Publisher); err != nil {
			return err
		}
		if err := wire.VDLWrite(enc); err != nil {
			return err
		}
	}
	if len(x.Env) != 0 {
		if err := enc.NextField(4); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_1(enc, x.Env); err != nil {
			return err
		}
	}
	if len(x.Packages) != 0 {
		if err := enc.NextField(5); err != nil {
			return err
		}
		if err := x.Packages.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.Restarts != 0 {
		if err := enc.NextFieldValueInt(6, vdl.Int32Type, int64(x.Restarts)); err != nil {
			return err
		}
	}
	if x.RestartTimeWindow != 0 {
		if err := enc.NextField(7); err != nil {
			return err
		}
		var wire vdltime.Duration
		if err := vdltime.DurationFromNative(&wire, x.RestartTimeWindow); err != nil {
			return err
		}
		if err := wire.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_list_1(enc vdl.Encoder, x []string) error {
	if err := enc.StartValue(__VDLType_list_5); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for _, elem := range x {
		if err := enc.NextEntryValueString(vdl.StringType, elem); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Envelope) VDLRead(dec vdl.Decoder) error {
	*x = Envelope{}
	if err := dec.StartValue(__VDLType_struct_4); err != nil {
		return err
	}
	decType := dec.Type()
	for {
		index, err := dec.NextField()
		switch {
		case err != nil:
			return err
		case index == -1:
			return dec.FinishValue()
		}
		if decType != __VDLType_struct_4 {
			index = __VDLType_struct_4.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Title = value
			}
		case 1:
			if err := __VDLReadAnon_list_1(dec, &x.Args); err != nil {
				return err
			}
		case 2:
			if err := x.Binary.VDLRead(dec); err != nil {
				return err
			}
		case 3:
			var wire security.WireBlessings
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := security.WireBlessingsToNative(wire, &x.Publisher); err != nil {
				return err
			}
		case 4:
			if err := __VDLReadAnon_list_1(dec, &x.Env); err != nil {
				return err
			}
		case 5:
			if err := x.Packages.VDLRead(dec); err != nil {
				return err
			}
		case 6:
			switch value, err := dec.ReadValueInt(32); {
			case err != nil:
				return err
			default:
				x.Restarts = int32(value)
			}
		case 7:
			var wire vdltime.Duration
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := vdltime.DurationToNative(wire, &x.RestartTimeWindow); err != nil {
				return err
			}
		}
	}
}

func __VDLReadAnon_list_1(dec vdl.Decoder, x *[]string) error {
	if err := dec.StartValue(__VDLType_list_5); err != nil {
		return err
	}
	if len := dec.LenHint(); len > 0 {
		*x = make([]string, 0, len)
	} else {
		*x = nil
	}
	for {
		switch done, elem, err := dec.NextEntryValueString(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		default:
			*x = append(*x, elem)
		}
	}
}

//////////////////////////////////////////////////
// Const definitions

// Device manager application envelopes must present this title.
const DeviceManagerTitle = "device manager"

// Hold type definitions in package-level variables, for better performance.
var (
	__VDLType_struct_1 *vdl.Type
	__VDLType_struct_2 *vdl.Type
	__VDLType_map_3    *vdl.Type
	__VDLType_struct_4 *vdl.Type
	__VDLType_list_5   *vdl.Type
	__VDLType_struct_6 *vdl.Type
	__VDLType_struct_7 *vdl.Type
)

var __VDLInitCalled bool

// __VDLInit performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = __VDLInit()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func __VDLInit() struct{} {
	if __VDLInitCalled {
		return struct{}{}
	}
	__VDLInitCalled = true

	// Register types.
	vdl.Register((*SignedFile)(nil))
	vdl.Register((*Packages)(nil))
	vdl.Register((*Envelope)(nil))

	// Initialize type definitions.
	__VDLType_struct_1 = vdl.TypeOf((*SignedFile)(nil)).Elem()
	__VDLType_struct_2 = vdl.TypeOf((*security.Signature)(nil)).Elem()
	__VDLType_map_3 = vdl.TypeOf((*Packages)(nil))
	__VDLType_struct_4 = vdl.TypeOf((*Envelope)(nil)).Elem()
	__VDLType_list_5 = vdl.TypeOf((*[]string)(nil))
	__VDLType_struct_6 = vdl.TypeOf((*security.WireBlessings)(nil)).Elem()
	__VDLType_struct_7 = vdl.TypeOf((*vdltime.Duration)(nil)).Elem()

	return struct{}{}
}
