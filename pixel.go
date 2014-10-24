// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"encoding/binary"
	"math"
	"reflect"
	"unsafe"
)

type Pixel struct {
	Channels int
	Depth    reflect.Kind
	Value    []byte
}

func (c Pixel) IsValid() bool {
	if c.Channels <= 0 || !depthType(c.Depth).IsValid() {
		return false
	}
	if len(c.Value) < depthType(c.Depth).ByteSize()*c.Channels {
		return false
	}
	return true
}

func (c Pixel) IsIntType() bool {
	return depthType(c.Depth).IsIntType()
}

func (c Pixel) IsUintType() bool {
	return depthType(c.Depth).IsUintType()
}

func (c Pixel) IsFloatType() bool {
	return depthType(c.Depth).IsFloatType()
}

func (c Pixel) IsByteType() bool {
	return depthType(c.Depth).IsByteType()
}

func (c Pixel) ByteN() []byte {
	if c.Depth != reflect.Invalid {
		panic("image: Pixel.ByteN, invalid type")
	}
	if c.Value == nil {
		return nil
	}
	n := depthType(c.Depth).ByteSize() * c.Channels
	d := ((*[1 << 30]byte)(unsafe.Pointer(&c.Value[0])))[0:n:n]
	return append([]byte(nil), d...)
}
func (c Pixel) PutByteN(v []byte) {
	if c.Depth != reflect.Invalid {
		panic("image: Pixel.PutByteN, invalid type")
	}
	if c.Value == nil {
		c.Value = make([]byte, depthType(c.Depth).ByteSize()*c.Channels)
	}
	n := depthType(c.Depth).ByteSize() * c.Channels
	d := ((*[1 << 30]byte)(unsafe.Pointer(&c.Value[0])))[0:n:n]
	copy(d, v)
}

func (c Pixel) Int8N() []int8 {
	if c.Depth != reflect.Int8 {
		panic("image: Pixel.Int8N, invalid type")
	}
	if c.Value == nil {
		return nil
	}
	d := ((*[1 << 30]int8)(unsafe.Pointer(&c.Value[0])))[0:c.Channels:c.Channels]
	return append([]int8(nil), d...)
}
func (c Pixel) PutInt8N(v []int8) {
	if c.Depth != reflect.Int8 {
		panic("image: Pixel.PutInt8N, invalid type")
	}
	if c.Value == nil {
		c.Value = make([]byte, depthType(c.Depth).ByteSize()*c.Channels)
	}
	d := ((*[1 << 30]int8)(unsafe.Pointer(&c.Value[0])))[0:c.Channels:c.Channels]
	copy(d, v)
}

func (c Pixel) Int16N() []int16 {
	if c.Depth != reflect.Int16 {
		panic("image: Pixel.Int16N, invalid type")
	}
	if c.Value == nil {
		return nil
	}
	t := make([]uint16, c.Channels)
	v := ((*[1 << 30]int16)(unsafe.Pointer(&t)))[0:c.Channels:c.Channels]
	for i := 0; i < c.Channels; i++ {
		t[i] = binary.BigEndian.Uint16(c.Value[i*2:])
	}
	return v
}
func (c Pixel) PutInt16N(v []int16) {
	if c.Depth != reflect.Int16 {
		panic("image: Pixel.PutInt16N, invalid type")
	}
	if c.Value == nil {
		c.Value = make([]byte, depthType(c.Depth).ByteSize()*c.Channels)
	}
	for i := 0; i < c.Channels && i < len(v); i++ {
		binary.BigEndian.PutUint16(c.Value[i*2:], *(*uint16)(unsafe.Pointer(&v[i])))
	}
}

func (c Pixel) Int32N() []int32 {
	if c.Depth != reflect.Int32 {
		panic("image: Pixel.Int32N, invalid type")
	}
	if c.Value == nil {
		return nil
	}
	t := make([]uint32, c.Channels)
	v := ((*[1 << 30]int32)(unsafe.Pointer(&t)))[0:c.Channels:c.Channels]
	for i := 0; i < c.Channels; i++ {
		t[i] = binary.BigEndian.Uint32(c.Value[i*4:])
	}
	return v
}
func (c Pixel) PutInt32N(v []int32) {
	if c.Depth != reflect.Int32 {
		panic("image: Pixel.PutInt32N, invalid type")
	}
	if c.Value == nil {
		c.Value = make([]byte, depthType(c.Depth).ByteSize()*c.Channels)
	}
	for i := 0; i < c.Channels && i < len(v); i++ {
		binary.BigEndian.PutUint32(c.Value[i*4:], *(*uint32)(unsafe.Pointer(&v[i])))
	}
}

func (c Pixel) Int64N() []int64 {
	if c.Depth != reflect.Int64 {
		panic("image: Pixel.Int64N, invalid type")
	}
	if c.Value == nil {
		return nil
	}
	t := make([]uint64, c.Channels)
	v := ((*[1 << 30]int64)(unsafe.Pointer(&t)))[0:c.Channels:c.Channels]
	for i := 0; i < c.Channels; i++ {
		t[i] = binary.BigEndian.Uint64(c.Value[i*8:])
	}
	return v
}
func (c Pixel) PutInt64N(v []int64) {
	if c.Depth != reflect.Int64 {
		panic("image: Pixel.PutInt64N, invalid type")
	}
	if c.Value == nil {
		c.Value = make([]byte, depthType(c.Depth).ByteSize()*c.Channels)
	}
	for i := 0; i < c.Channels && i < len(v); i++ {
		binary.BigEndian.PutUint64(c.Value[i*8:], *(*uint64)(unsafe.Pointer(&v[i])))
	}
}

func (c Pixel) Uint8N() []uint8 {
	if c.Depth != reflect.Uint8 {
		panic("image: Pixel.Uint8N, invalid type")
	}
	if c.Value == nil {
		return nil
	}
	return append([]uint8(nil), c.Value...)
}
func (c Pixel) PutUint8N(v []uint8) {
	if c.Depth != reflect.Uint8 {
		panic("image: Pixel.PutUint8N, invalid type")
	}
	if c.Value == nil {
		c.Value = make([]byte, depthType(c.Depth).ByteSize()*c.Channels)
	}
	copy(c.Value, v)
}

func (c Pixel) Uint16N() []uint16 {
	if c.Depth != reflect.Uint16 {
		panic("image: Pixel.Uint16N, invalid type")
	}
	if c.Value == nil {
		return nil
	}
	v := make([]uint16, c.Channels)
	for i := 0; i < c.Channels; i++ {
		v[i] = binary.BigEndian.Uint16(c.Value[i*2:])
	}
	return v
}
func (c Pixel) PutUint16N(v []uint16) {
	if c.Depth != reflect.Uint16 {
		panic("image: Pixel.PutUint16N, invalid type")
	}
	if c.Value == nil {
		c.Value = make([]byte, depthType(c.Depth).ByteSize()*c.Channels)
	}
	for i := 0; i < c.Channels && i < len(v); i++ {
		binary.BigEndian.PutUint16(c.Value[i*2:], v[i])
	}
}

func (c Pixel) Uint32N() []uint32 {
	if c.Depth != reflect.Uint32 {
		panic("image: Pixel.Uint32N, invalid type")
	}
	if c.Value == nil {
		return nil
	}
	v := make([]uint32, c.Channels)
	for i := 0; i < c.Channels; i++ {
		v[i] = binary.BigEndian.Uint32(c.Value[i*4:])
	}
	return v
}
func (c Pixel) PutUint32N(v []uint32) {
	if c.Depth != reflect.Uint32 {
		panic("image: Pixel.PutUint32N, invalid type")
	}
	if c.Value == nil {
		c.Value = make([]byte, depthType(c.Depth).ByteSize()*c.Channels)
	}
	for i := 0; i < c.Channels && i < len(v); i++ {
		binary.BigEndian.PutUint32(c.Value[i*4:], v[i])
	}
}

func (c Pixel) Uint64N() []uint64 {
	if c.Depth != reflect.Uint64 {
		panic("image: Pixel.Uint64N, invalid type")
	}
	if c.Value == nil {
		return nil
	}
	v := make([]uint64, c.Channels)
	for i := 0; i < c.Channels; i++ {
		v[i] = binary.BigEndian.Uint64(c.Value[i*8:])
	}
	return v
}
func (c Pixel) PutUint64N(v []uint64) {
	if c.Depth != reflect.Uint64 {
		panic("image: Pixel.PutUint64N, invalid type")
	}
	if c.Value == nil {
		c.Value = make([]byte, depthType(c.Depth).ByteSize()*c.Channels)
	}
	for i := 0; i < c.Channels && i < len(v); i++ {
		binary.BigEndian.PutUint64(c.Value[i*8:], v[i])
	}
}

func (c Pixel) Float32N() []float32 {
	if c.Depth != reflect.Float32 {
		panic("image: Pixel.Float32N, invalid type")
	}
	if c.Value == nil {
		return nil
	}
	v := make([]float32, c.Channels)
	for i := 0; i < c.Channels; i++ {
		v[i] = math.Float32frombits(binary.BigEndian.Uint32(c.Value[i*4:]))
	}
	return v
}
func (c Pixel) PutFloat32N(v []float32) {
	if c.Depth != reflect.Float32 {
		panic("image: Pixel.PutFloat32N, invalid type")
	}
	if c.Value == nil {
		c.Value = make([]byte, depthType(c.Depth).ByteSize()*c.Channels)
	}
	for i := 0; i < c.Channels && i < len(v); i++ {
		binary.BigEndian.PutUint32(c.Value[i*4:], math.Float32bits(v[i]))
	}
}

func (c Pixel) Float64N() []float64 {
	if c.Depth != reflect.Float64 {
		panic("image: Pixel.Float64N, invalid type")
	}
	if c.Value == nil {
		return nil
	}
	v := make([]float64, c.Channels)
	for i := 0; i < c.Channels; i++ {
		v[i] = math.Float64frombits(binary.BigEndian.Uint64(c.Value[i*8:]))
	}
	return v
}
func (c Pixel) PutFloat64N(v []float64) {
	if c.Depth != reflect.Float64 {
		panic("image: Pixel.PutFloat64N, invalid type")
	}
	if c.Value == nil {
		c.Value = make([]byte, depthType(c.Depth).ByteSize()*c.Channels)
	}
	for i := 0; i < c.Channels && i < len(v); i++ {
		binary.BigEndian.PutUint64(c.Value[i*8:], math.Float64bits(v[i]))
	}
}

func (c Pixel) RGBA() (r, g, b, a uint32) {
	if c.Value == nil {
		return
	}
	panic("TODO")
}
func (c Pixel) PutRGBA(r, g, b, a uint32) {
	if c.Value == nil {
		c.Value = make([]byte, depthType(c.Depth).ByteSize()*c.Channels)
	}
	panic("TODO")
}
