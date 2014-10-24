// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"reflect"
)

type Pixel struct {
	Channels int
	Depth    reflect.Kind
	Value    []byte
}

func (c Pixel) RGBA() (r, g, b, a uint32) {
	panic("TODO")
}
func (c Pixel) PutRGBA(r, g, b, a uint32) {
	panic("TODO")
}

func (c Pixel) Int8N() []int8 {
	panic("TODO")
}
func (c Pixel) PutInt8N(v []int8) {
	panic("TODO")
}

func (c Pixel) Int16N() []int16 {
	panic("TODO")
}
func (c Pixel) PutInt16N(v []int16) {
	panic("TODO")
}

func (c Pixel) Int32N() []int32 {
	panic("TODO")
}
func (c Pixel) PutInt32N(v []int32) {
	panic("TODO")
}

func (c Pixel) Int64N() []int64 {
	panic("TODO")
}
func (c Pixel) PutInt64N(v []int64) {
	panic("TODO")
}

func (c Pixel) Uint8N() []uint8 {
	panic("TODO")
}
func (c Pixel) PutUint8N(v []uint8) {
	panic("TODO")
}

func (c Pixel) Uint16N() []uint16 {
	panic("TODO")
}
func (c Pixel) PutUint16N(v []uint16) {
	panic("TODO")
}

func (c Pixel) Uint32N() []uint32 {
	panic("TODO")
}
func (c Pixel) PutUint32N(v []uint32) {
	panic("TODO")
}

func (c Pixel) Uint64N() []uint64 {
	panic("TODO")
}
func (c Pixel) PutUint64N(v []uint64) {
	panic("TODO")
}

func (c Pixel) Float32N() []float32 {
	panic("TODO")
}
func (c Pixel) PutFloat32N(v []float32) {
	panic("TODO")
}

func (c Pixel) Float64N() []float64 {
	panic("TODO")
}
func (c Pixel) PutFloat64N(v []float64) {
	panic("TODO")
}
