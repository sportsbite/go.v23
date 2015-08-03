// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package conversions

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"v.io/syncbase/v23/syncbase/nosql/internal/query/query_parser"
)

func ConvertValueToString(o *query_parser.Operand) (*query_parser.Operand, error) {
	var c query_parser.Operand
	c.Type = query_parser.TypStr
	c.Off = o.Off
	switch o.Type {
	case query_parser.TypBigInt:
		c.Str = o.BigInt.String()
	case query_parser.TypBigRat:
		c.Str = o.BigRat.String()
	case query_parser.TypBool:
		c.Str = strconv.FormatBool(o.Bool)
	case query_parser.TypComplex:
		c.Str = fmt.Sprintf("%g", o.Complex)
	case query_parser.TypFloat:
		c.Str = strconv.FormatFloat(o.Float, 'f', -1, 64)
	case query_parser.TypInt:
		c.Str = strconv.FormatInt(o.Int, 10)
	case query_parser.TypStr:
		c.Str = o.Str
		c.Regex = o.Regex         // non-empty for rhs of like expressions
		c.CompRegex = o.CompRegex // non-nil for rhs of like expressions
	case query_parser.TypUint:
		c.Str = strconv.FormatUint(o.Uint, 10)
	case query_parser.TypObject:
		return nil, errors.New("Cannot convert object to string.")
	default:
		// TODO(jkline): Log this logic error and all other similar cases.
		return nil, errors.New("Cannot convert operand to string.")
	}
	return &c, nil
}

func ConvertValueToTime(o *query_parser.Operand) (*query_parser.Operand, error) {
	switch o.Type {
	case query_parser.TypTime:
		return o, nil
	default:
		return nil, errors.New("Cannot convert operand to time.")
	}
}

func ConvertValueToComplex(o *query_parser.Operand) (*query_parser.Operand, error) {
	var c query_parser.Operand
	c.Type = query_parser.TypComplex
	switch o.Type {
	case query_parser.TypComplex:
		return o, nil
	case query_parser.TypFloat:
		c.Complex = complex(o.Float, 0.0i)
	case query_parser.TypInt:
		c.Complex = complex(float64(o.Int), 0.0i)
	case query_parser.TypUint:
		c.Complex = complex(float64(o.Uint), 0.0i)
	default:
		return nil, errors.New("Cannot convert operand to Complex.")
	}
	return &c, nil
}

func ConvertValueToBigRat(o *query_parser.Operand) (*query_parser.Operand, error) {
	// operand cannot be string literal.
	var c query_parser.Operand
	c.Type = query_parser.TypBigRat
	switch o.Type {
	case query_parser.TypBigInt:
		var b big.Rat
		c.BigRat = b.SetInt(o.BigInt)
	case query_parser.TypBigRat:
		c.BigRat = o.BigRat
	case query_parser.TypBool:
		return nil, errors.New("Cannot convert bool to big.Rat.")
	case query_parser.TypFloat:
		var b big.Rat
		c.BigRat = b.SetFloat64(o.Float)
	case query_parser.TypInt:
		c.BigRat = big.NewRat(o.Int, 1)
	case query_parser.TypUint:
		var bi big.Int
		bi.SetUint64(o.Uint)
		var br big.Rat
		c.BigRat = br.SetInt(&bi)
	case query_parser.TypObject:
		return nil, errors.New("Cannot convert object to big.Rat.")
	default:
		// TODO(jkline): Log this logic error and all other similar cases.
		return nil, errors.New("Cannot convert operand to big.Rat.")
	}
	return &c, nil
}

func ConvertValueToFloat(o *query_parser.Operand) (*query_parser.Operand, error) {
	// Operand cannot be literal, big.Rat or big.Int
	var c query_parser.Operand
	c.Type = query_parser.TypFloat
	switch o.Type {
	case query_parser.TypBool:
		return nil, errors.New("Cannot convert bool to float64.")
	case query_parser.TypFloat:
		c.Float = o.Float
	case query_parser.TypInt:
		c.Float = float64(o.Int)
	case query_parser.TypUint:
		c.Float = float64(o.Uint)
	case query_parser.TypObject:
		return nil, errors.New("Cannot convert object to float64.")
	default:
		// TODO(jkline): Log this logic error and all other similar cases.
		return nil, errors.New("Cannot convert operand to float64.")
	}
	return &c, nil
}

func ConvertValueToBigInt(o *query_parser.Operand) (*query_parser.Operand, error) {
	// Operand cannot be literal, big.Rat or float.
	var c query_parser.Operand
	c.Type = query_parser.TypBigInt
	switch o.Type {
	case query_parser.TypBigInt:
		c.BigInt = o.BigInt
	case query_parser.TypBool:
		return nil, errors.New("Cannot convert bool to big.Int.")
	case query_parser.TypInt:
		c.BigInt = big.NewInt(o.Int)
	case query_parser.TypUint:
		var b big.Int
		b.SetUint64(o.Uint)
		c.BigInt = &b
	case query_parser.TypObject:
		return nil, errors.New("Cannot convert object to big.Int.")
	default:
		// TODO(jkline): Log this logic error and all other similar cases.
		return nil, errors.New("Cannot convert operand to big.Int.")
	}
	return &c, nil
}

func ConvertValueToInt(o *query_parser.Operand) (*query_parser.Operand, error) {
	// Operand cannot be literal, big.Rat or float or uint64.
	var c query_parser.Operand
	c.Type = query_parser.TypInt
	switch o.Type {
	case query_parser.TypBool:
		return nil, errors.New("Cannot convert bool to int64.")
	case query_parser.TypInt:
		c.Int = o.Int
	case query_parser.TypObject:
		return nil, errors.New("Cannot convert object to int64.")
	default:
		// TODO(jkline): Log this logic error and all other similar cases.
		return nil, errors.New("Cannot convert operand to int64.")
	}
	return &c, nil
}

func ConvertValueToUint(o *query_parser.Operand) (*query_parser.Operand, error) {
	// Operand cannot be literal, big.Rat or float or int64.
	var c query_parser.Operand
	c.Type = query_parser.TypUint
	switch o.Type {
	case query_parser.TypBool:
		return nil, errors.New("Cannot convert bool to int64.")
	case query_parser.TypUint:
		c.Uint = o.Uint
	case query_parser.TypObject:
		return nil, errors.New("Cannot convert object to int64.")
	default:
		// TODO(jkline): Log this logic error and all other similar cases.
		return nil, errors.New("Cannot convert operand to int64.")
	}
	return &c, nil
}
