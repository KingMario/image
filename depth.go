// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"reflect"
)

type depthType reflect.Kind

func (d depthType) IsValid() bool {
	_, ok := depthTypeTable[reflect.Kind(d)]
	return ok
}

func (d depthType) IsIntType() bool {
	switch reflect.Kind(d) {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}

func (d depthType) IsUintType() bool {
	switch reflect.Kind(d) {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	}
	return false
}

func (d depthType) IsFloatType() bool {
	switch reflect.Kind(d) {
	case reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

func (d depthType) IsByteType() bool {
	return reflect.Kind(d) == reflect.Invalid
}

func (d depthType) Depth() int {
	v, _ := depthTypeTable[reflect.Kind(d)]
	return (v >> 8) // bits
}

func (d depthType) ByteSize() int {
	v, _ := depthTypeTable[reflect.Kind(d)]
	return v
}

var depthTypeTable = map[reflect.Kind]int{
	reflect.Invalid: 1, // Byte
	reflect.Int8:    1,
	reflect.Int16:   2,
	reflect.Int32:   4,
	reflect.Int64:   8,
	reflect.Uint8:   1,
	reflect.Uint16:  2,
	reflect.Uint32:  4,
	reflect.Uint64:  8,
	reflect.Float32: 4,
	reflect.Float64: 8,
}
