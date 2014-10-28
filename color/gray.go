// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

import (
	"image/color"
)

type Gray struct {
	Y uint8
}

func (c Gray) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y)<<8 | uint16(c.Y))
	return y, y, y, 0xFFFF
}

func grayModel(c color.Color) color.Color {
	if c, ok := c.(Gray); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return Gray{Y: uint8(y >> 8)}
}

type Gray16 struct {
	Y uint16
}

func (c Gray16) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	return y, y, y, 0xFFFF
}

func gray16Model(c color.Color) color.Color {
	if c, ok := c.(Gray16); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return Gray16{Y: uint16(y)}
}

type Gray32i struct {
	Y int32
}

func (c Gray32i) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	return y, y, y, 0xFFFF
}

func gray32iModel(c color.Color) color.Color {
	if c, ok := c.(Gray32i); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return Gray32i{Y: int32(y)}
}

type Gray32f struct {
	Y float32
}

func (c Gray32f) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	return y, y, y, 0xFFFF
}

func gray32fModel(c color.Color) color.Color {
	if c, ok := c.(Gray32f); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return Gray32f{Y: float32(y)}
}

type Gray64i struct {
	Y int64
}

func (c Gray64i) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	return y, y, y, 0xFFFF
}

func gray64iModel(c color.Color) color.Color {
	if c, ok := c.(Gray64i); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return Gray64i{Y: int64(y)}
}

type Gray64f struct {
	Y float64
}

func (c Gray64f) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	return y, y, y, 0xFFFF
}

func gray64fModel(c color.Color) color.Color {
	if c, ok := c.(Gray64f); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return Gray64f{Y: float64(y)}
}
