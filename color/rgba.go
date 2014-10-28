// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

import (
	"image/color"
)

type RGBA struct {
	R, G, B, A uint8
}

func (c RGBA) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R)<<8 | uint16(c.R))
	g = uint32(uint16(c.G)<<8 | uint16(c.G))
	b = uint32(uint16(c.B)<<8 | uint16(c.B))
	a = uint32(uint16(c.A)<<8 | uint16(c.A))
	return
}

func rgbaModel(c color.Color) color.Color {
	if c, ok := c.(RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return RGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: uint8(a >> 8),
	}
}

type RGBA64 struct {
	R, G, B, A uint16
}

func (c RGBA64) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = uint32(uint16(c.A))
	return
}

func rgba64Model(c color.Color) color.Color {
	if c, ok := c.(RGBA64); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return RGBA64{
		R: uint16(r),
		G: uint16(g),
		B: uint16(b),
		A: uint16(a),
	}
}

type RGBA128i struct {
	R, G, B, A int32
}

func (c RGBA128i) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = uint32(uint16(c.A))
	return
}

func rgba128iModel(c color.Color) color.Color {
	if c, ok := c.(RGBA128i); ok {
		return c
	}
	switch c := c.(type) {
	case Gray32i:
		return RGBA128i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
			A: 0xFFFF,
		}
	case Gray32f:
		return RGBA128i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
			A: 0xFFFF,
		}
	case Gray64i:
		return RGBA128i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
			A: 0xFFFF,
		}
	case Gray64f:
		return RGBA128i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
			A: 0xFFFF,
		}
	case GrayA64i:
		return RGBA128i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
			A: int32(c.A),
		}
	case GrayA64f:
		return RGBA128i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
			A: int32(c.A),
		}
	case GrayA128i:
		return RGBA128i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
			A: int32(c.A),
		}
	case GrayA128f:
		return RGBA128i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
			A: int32(c.A),
		}
	case RGB96i:
		return RGBA128i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
			A: 0xFFFF,
		}
	case RGB96f:
		return RGBA128i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
			A: 0xFFFF,
		}
	case RGB192i:
		return RGBA128i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
			A: 0xFFFF,
		}
	case RGB192f:
		return RGBA128i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
			A: 0xFFFF,
		}
	case RGBA128i:
		return RGBA128i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
			A: int32(c.A),
		}
	case RGBA128f:
		return RGBA128i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
			A: int32(c.A),
		}
	case RGBA256i:
		return RGBA128i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
			A: int32(c.A),
		}
	case RGBA256f:
		return RGBA128i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
			A: int32(c.A),
		}
	}
	r, g, b, a := c.RGBA()
	return RGBA128i{
		R: int32(r),
		G: int32(g),
		B: int32(b),
		A: int32(a),
	}
}

type RGBA128f struct {
	R, G, B, A float32
}

func (c RGBA128f) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = uint32(uint16(c.A))
	return
}

func rgba128fModel(c color.Color) color.Color {
	if c, ok := c.(RGBA128f); ok {
		return c
	}
	switch c := c.(type) {
	case Gray32i:
		return RGBA128f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
			A: 0xFFFF,
		}
	case Gray32f:
		return RGBA128f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
			A: 0xFFFF,
		}
	case Gray64i:
		return RGBA128f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
			A: 0xFFFF,
		}
	case Gray64f:
		return RGBA128f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
			A: 0xFFFF,
		}
	case GrayA64i:
		return RGBA128f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
			A: float32(c.A),
		}
	case GrayA64f:
		return RGBA128f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
			A: float32(c.A),
		}
	case GrayA128i:
		return RGBA128f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
			A: float32(c.A),
		}
	case GrayA128f:
		return RGBA128f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
			A: float32(c.A),
		}
	case RGB96i:
		return RGBA128f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
			A: 0xFFFF,
		}
	case RGB96f:
		return RGBA128f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
			A: 0xFFFF,
		}
	case RGB192i:
		return RGBA128f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
			A: 0xFFFF,
		}
	case RGB192f:
		return RGBA128f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
			A: 0xFFFF,
		}
	case RGBA128i:
		return RGBA128f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
			A: float32(c.A),
		}
	case RGBA128f:
		return RGBA128f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
			A: float32(c.A),
		}
	case RGBA256i:
		return RGBA128f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
			A: float32(c.A),
		}
	case RGBA256f:
		return RGBA128f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
			A: float32(c.A),
		}
	}
	r, g, b, a := c.RGBA()
	return RGBA128f{
		R: float32(r),
		G: float32(g),
		B: float32(b),
		A: float32(a),
	}
}

type RGBA256i struct {
	R, G, B, A int64
}

func (c RGBA256i) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = uint32(uint16(c.A))
	return
}

func rgba256iModel(c color.Color) color.Color {
	if c, ok := c.(RGBA256i); ok {
		return c
	}
	switch c := c.(type) {
	case Gray32i:
		return RGBA256i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
			A: 0xFFFF,
		}
	case Gray32f:
		return RGBA256i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
			A: 0xFFFF,
		}
	case Gray64i:
		return RGBA256i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
			A: 0xFFFF,
		}
	case Gray64f:
		return RGBA256i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
			A: 0xFFFF,
		}
	case GrayA64i:
		return RGBA256i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
			A: int64(c.A),
		}
	case GrayA64f:
		return RGBA256i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
			A: int64(c.A),
		}
	case GrayA128i:
		return RGBA256i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
			A: int64(c.A),
		}
	case GrayA128f:
		return RGBA256i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
			A: int64(c.A),
		}
	case RGB96i:
		return RGBA256i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
			A: 0xFFFF,
		}
	case RGB96f:
		return RGBA256i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
			A: 0xFFFF,
		}
	case RGB192i:
		return RGBA256i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
			A: 0xFFFF,
		}
	case RGB192f:
		return RGBA256i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
			A: 0xFFFF,
		}
	case RGBA128i:
		return RGBA256i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
			A: int64(c.A),
		}
	case RGBA128f:
		return RGBA256i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
			A: int64(c.A),
		}
	case RGBA256i:
		return RGBA256i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
			A: int64(c.A),
		}
	case RGBA256f:
		return RGBA256i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
			A: int64(c.A),
		}
	}
	r, g, b, a := c.RGBA()
	return RGBA256i{
		R: int64(r),
		G: int64(g),
		B: int64(b),
		A: int64(a),
	}
}

type RGBA256f struct {
	R, G, B, A float64
}

func (c RGBA256f) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = uint32(uint16(c.A))
	return
}

func rgba256fModel(c color.Color) color.Color {
	if c, ok := c.(RGBA256f); ok {
		return c
	}
	switch c := c.(type) {
	case Gray32i:
		return RGBA256f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
			A: 0xFFFF,
		}
	case Gray32f:
		return RGBA256f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
			A: 0xFFFF,
		}
	case Gray64i:
		return RGBA256f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
			A: 0xFFFF,
		}
	case Gray64f:
		return RGBA256f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
			A: 0xFFFF,
		}
	case GrayA64i:
		return RGBA256f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
			A: float64(c.A),
		}
	case GrayA64f:
		return RGBA256f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
			A: float64(c.A),
		}
	case GrayA128i:
		return RGBA256f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
			A: float64(c.A),
		}
	case GrayA128f:
		return RGBA256f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
			A: float64(c.A),
		}
	case RGB96i:
		return RGBA256f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
			A: 0xFFFF,
		}
	case RGB96f:
		return RGBA256f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
			A: 0xFFFF,
		}
	case RGB192i:
		return RGBA256f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
			A: 0xFFFF,
		}
	case RGB192f:
		return RGBA256f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
			A: 0xFFFF,
		}
	case RGBA128i:
		return RGBA256f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
			A: float64(c.A),
		}
	case RGBA128f:
		return RGBA256f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
			A: float64(c.A),
		}
	case RGBA256i:
		return RGBA256f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
			A: float64(c.A),
		}
	case RGBA256f:
		return RGBA256f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
			A: float64(c.A),
		}
	}
	r, g, b, a := c.RGBA()
	return RGBA256f{
		R: float64(r),
		G: float64(g),
		B: float64(b),
		A: float64(a),
	}
}
