// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

import (
	"image/color"
)

type GrayA struct {
	Y, A uint8
}

func (c GrayA) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y)<<8 | uint16(c.Y))
	a = uint32(uint16(c.A)<<8 | uint16(c.A))
	return y, y, y, a
}

func grayAModel(c color.Color) color.Color {
	if c, ok := c.(GrayA); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return GrayA{
		Y: uint8(y >> 8),
		A: 0xFF,
	}
}

type GrayA32 struct {
	Y, A uint16
}

func (c GrayA32) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	a = uint32(uint16(c.A))
	return y, y, y, a
}

func grayA32Model(c color.Color) color.Color {
	if c, ok := c.(GrayA32); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return GrayA32{
		Y: uint16(y),
		A: uint16(a),
	}
}

type GrayA64i struct {
	Y, A int32
}

func (c GrayA64i) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	a = uint32(uint16(c.A))
	return y, y, y, a
}

func grayA64iModel(c color.Color) color.Color {
	if c, ok := c.(GrayA64i); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return GrayA64i{
		Y: int32(y),
		A: int32(a),
	}
}

type GrayA64f struct {
	Y, A float32
}

func (c GrayA64f) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	a = uint32(uint16(c.A))
	return y, y, y, a
}

func grayA64fModel(c color.Color) color.Color {
	if c, ok := c.(GrayA64f); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return GrayA64f{
		Y: float32(y),
		A: float32(a),
	}
}

type GrayA128i struct {
	Y, A int64
}

func (c GrayA128i) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	a = uint32(uint16(c.A))
	return y, y, y, a
}

func grayA128iModel(c color.Color) color.Color {
	if c, ok := c.(GrayA128i); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return GrayA128i{
		Y: int64(y),
		A: int64(a),
	}
}

type GrayA128f struct {
	Y, A float64
}

func (c GrayA128f) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	a = uint32(uint16(c.A))
	return y, y, y, a
}

func grayA128fModel(c color.Color) color.Color {
	if c, ok := c.(GrayA128f); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return GrayA128f{
		Y: float64(y),
		A: float64(a),
	}
}
