// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package vom

import (
	"fmt"
	"io"
	"math"
	"reflect"
	"strings"
	"testing"
)

func hex2Bin(t *testing.T, hex string) []byte {
	binStr, err := binFromHexPat(hex)
	if err != nil {
		t.Fatalf("error converting %q to binary: %v", hex, err)
	}
	return []byte(binStr)
}

func TestBinaryEncodeDecode(t *testing.T) {
	tests := []struct {
		v   interface{}
		hex string
	}{
		{false, "00"},
		{true, "01"},
		{byte(0x80), "80"},
		{byte(0xbf), "bf"},
		{byte(0xc0), "c0"},
		{byte(0xdf), "df"},
		{byte(0xe0), "e0"},
		{byte(0xef), "ef"},
		{uint64(0), "00"},
		{uint64(1), "01"},
		{uint64(2), "02"},
		{uint64(127), "7f"},
		{uint64(128), "ff80"},
		{uint64(255), "ffff"},
		{uint64(256), "fe0100"},
		{uint64(257), "fe0101"},
		{uint64(0xffff), "feffff"},
		{uint64(0xffffff), "fdffffff"},
		{uint64(0xffffffff), "fcffffffff"},
		{uint64(0xffffffffff), "fbffffffffff"},
		{uint64(0xffffffffffff), "faffffffffffff"},
		{uint64(0xffffffffffffff), "f9ffffffffffffff"},
		{uint64(0xffffffffffffffff), "f8ffffffffffffffff"},
		{int64(0), "00"},
		{int64(1), "02"},
		{int64(2), "04"},
		{int64(63), "7e"},
		{int64(64), "ff80"},
		{int64(65), "ff82"},
		{int64(127), "fffe"},
		{int64(128), "fe0100"},
		{int64(129), "fe0102"},
		{int64(math.MaxInt16), "fefffe"},
		{int64(math.MaxInt32), "fcfffffffe"},
		{int64(math.MaxInt64), "f8fffffffffffffffe"},
		{int64(-1), "01"},
		{int64(-2), "03"},
		{int64(-64), "7f"},
		{int64(-65), "ff81"},
		{int64(-66), "ff83"},
		{int64(-128), "ffff"},
		{int64(-129), "fe0101"},
		{int64(-130), "fe0103"},
		{int64(math.MinInt16), "feffff"},
		{int64(math.MinInt32), "fcffffffff"},
		{int64(math.MinInt64), "f8ffffffffffffffff"},
		{float64(0), "00"},
		{float64(1), "fef03f"},
		{float64(17), "fe3140"},
		{float64(18), "fe3240"},
		{"", "00"},
		{"abc", "03616263"},
		{"defghi", "06646566676869"},
	}
	for _, test := range tests {
		// Test encode
		encbuf := newEncbuf()
		var buf []byte
		switch val := test.v.(type) {
		case byte:
			binaryEncodeControl(encbuf, val)
		case bool:
			binaryEncodeBool(encbuf, val)
		case uint64:
			binaryEncodeUint(encbuf, val)
			buf = make([]byte, maxEncodedUintBytes)
			buf = buf[binaryEncodeUintEnd(buf, val):]
		case int64:
			binaryEncodeInt(encbuf, val)
			buf = make([]byte, maxEncodedUintBytes)
			buf = buf[binaryEncodeIntEnd(buf, val):]
		case float64:
			binaryEncodeFloat(encbuf, val)
		case string:
			binaryEncodeString(encbuf, val)
		}
		if got, want := fmt.Sprintf("%x", encbuf.Bytes()), test.hex; got != want {
			t.Errorf("binary encode %T(%v): GOT 0x%v WANT 0x%v", test.v, test.v, got, want)
		}
		if buf != nil {
			if got, want := fmt.Sprintf("%x", buf), test.hex; got != want {
				t.Errorf("binary encode end %T(%v): GOT 0x%v WANT 0x%v", test.v, test.v, got, want)
			}
		}
		// Test decode
		var bin string
		if _, err := fmt.Sscanf(test.hex, "%x", &bin); err != nil {
			t.Errorf("couldn't scan 0x%v as hex: %v", test.hex, err)
			continue
		}
		decbuf := newDecbuf(strings.NewReader(bin))
		decbuf2 := newDecbuf(strings.NewReader(bin))
		decbuf3 := newDecbuf(strings.NewReader(bin))
		decbuf4 := newDecbuf(strings.NewReader(bin))
		decbuf5 := newDecbuf(strings.NewReader(bin))
		var v, v2 interface{}
		var err, err2, err3 error
		var vmr, vmr2 interface{}
		var errmr, errmr2, errmr3 error
		switch test.v.(type) {
		case byte:
			_, v, err = binaryDecodeUintWithControl(decbuf)
			decbuf2.Skip(1)
			vmr, err = binaryPeekControl(decbuf3)
			decbuf3.Skip(1)
			_, vmr2, errmr2 = binaryDecodeUintWithControl(decbuf4)
			errmr3 = binaryIgnoreUint(decbuf5)
		case bool:
			decbuf.Skip(1)
			decbuf2.Skip(1)
			vmr, errmr = binaryDecodeBool(decbuf3)
			errmr2 = binaryIgnoreUint(decbuf4)
		case uint64:
			v, err = binaryDecodeUint(decbuf)
			v2, _, err2 = binaryDecodeUintWithControl(decbuf2)
			vmr, errmr = binaryDecodeUint(decbuf3)
			vmr2, _, errmr2 = binaryDecodeUintWithControl(decbuf4)
			errmr3 = binaryIgnoreUint(decbuf5)
		case int64:
			v, err = binaryDecodeInt(decbuf)
			v2, err2 = binaryDecodeInt(decbuf2)
			vmr, errmr = binaryDecodeInt(decbuf3)
			errmr2 = binaryIgnoreUint(decbuf4)
		case float64:
			binaryDecodeUint(decbuf)  // skip
			binaryDecodeUint(decbuf2) // skip
			vmr, errmr = binaryDecodeFloat(decbuf3)
			errmr2 = binaryIgnoreUint(decbuf4)
		case string:
			l, _ := binaryDecodeUint(decbuf)
			decbuf.Skip(int(l))
			l, _ = binaryDecodeUint(decbuf2)
			decbuf2.Skip(int(l))
			vmr, errmr = binaryDecodeString(decbuf3)
			errmr2 = binaryIgnoreString(decbuf4)
		}
		if v2 != nil && v != v2 {
			t.Errorf("binary decode %T(0x%v) differently: %v %v", test.v, test.hex, v, v2)
			continue
		}
		if vmr2 != nil && vmr != vmr2 {
			t.Errorf("message reader binary decode %T(0x%v) differently: %v %v", test.v, test.hex, v, v2)
			continue
		}
		if err != nil || err2 != nil || err3 != nil {
			t.Errorf("binary decode %T(0x%v): %v %v %v", test.v, test.hex, err, err2, err3)
			continue
		}
		if errmr != nil || errmr2 != nil || errmr3 != nil {
			t.Errorf("message reader binary decode %T(0x%v): %v %v %v", test.v, test.hex, errmr, errmr2, errmr3)
			continue
		}
		if b, err := decbuf.ReadByte(); err != io.EOF {
			t.Errorf("binary decode %T(0x%v) leftover byte r=%x", test.v, test.hex, b)
			continue
		}
		if b, err := decbuf2.ReadByte(); err != io.EOF {
			t.Errorf("binary decode %T(0x%v) leftover byte r2=%x", test.v, test.hex, b)
			continue
		}
		if v != nil {
			if !reflect.DeepEqual(v, test.v) {
				t.Errorf("binary decode %T(0x%v): GOT %v WANT %v", test.v, test.hex, v, test.v)
			}
		}
		if vmr != nil {
			if !reflect.DeepEqual(vmr, test.v) {
				t.Errorf("binary decode %T(0x%v): GOT %v WANT %v", test.v, test.hex, vmr, test.v)
			}
		}
	}
}
