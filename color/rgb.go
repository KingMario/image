// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

import (
	"image/color"
)

type RGB struct {
	R, G, B uint8
}

func (c RGB) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R)<<8 | uint16(c.R))
	g = uint32(uint16(c.G)<<8 | uint16(c.G))
	b = uint32(uint16(c.B)<<8 | uint16(c.B))
	a = 0xFFFF
	return
}

func rgbModel(c color.Color) color.Color {
	if c, ok := c.(RGB); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
	}
}

type RGB48 struct {
	R, G, B uint16
}

func (c RGB48) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = 0xFFFF
	return
}

func rgb48Model(c color.Color) color.Color {
	if c, ok := c.(RGB48); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB48{
		R: uint16(r),
		G: uint16(g),
		B: uint16(b),
	}
}

type RGB96i struct {
	R, G, B int32
}

func (c RGB96i) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = 0xFFFF
	return
}

func rgb96iModel(c color.Color) color.Color {
	if c, ok := c.(RGB96i); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB96i{
		R: int32(r),
		G: int32(g),
		B: int32(b),
	}
}

type RGB96f struct {
	R, G, B float32
}

func (c RGB96f) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = 0xFFFF
	return
}

func rgb96fModel(c color.Color) color.Color {
	if c, ok := c.(RGB96f); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB96f{
		R: float32(r),
		G: float32(g),
		B: float32(b),
	}
}

type RGB192i struct {
	R, G, B int64
}

func (c RGB192i) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = 0xFFFF
	return
}

func rgb192iModel(c color.Color) color.Color {
	if c, ok := c.(RGB192i); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB192i{
		R: int64(r),
		G: int64(g),
		B: int64(b),
	}
}

type RGB192f struct {
	R, G, B float64
}

func (c RGB192f) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = 0xFFFF
	return
}

func rgb192fModel(c color.Color) color.Color {
	if c, ok := c.(RGB192f); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB192f{
		R: float64(r),
		G: float64(g),
		B: float64(b),
	}
}
