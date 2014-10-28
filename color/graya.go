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
	switch c := c.(type) {
	case Gray32i:
		return GrayA64i{
			Y: int32(c.Y),
			A: 0xFFFF,
		}
	case Gray32f:
		return GrayA64i{
			Y: int32(c.Y),
			A: 0xFFFF,
		}
	case Gray64i:
		return GrayA64i{
			Y: int32(c.Y),
			A: 0xFFFF,
		}
	case Gray64f:
		return GrayA64i{
			Y: int32(c.Y),
			A: 0xFFFF,
		}
	case GrayA64i:
		return GrayA64i{
			Y: int32(c.Y),
			A: int32(c.A),
		}
	case GrayA64f:
		return GrayA64i{
			Y: int32(c.Y),
			A: int32(c.A),
		}
	case GrayA128i:
		return GrayA64i{
			Y: int32(c.Y),
			A: int32(c.A),
		}
	case GrayA128f:
		return GrayA64i{
			Y: int32(c.Y),
			A: int32(c.A),
		}
	case RGB96i:
		return GrayA64i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
			A: 0xFFFF,
		}
	case RGB96f:
		return GrayA64i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
			A: 0xFFFF,
		}
	case RGB192i:
		return GrayA64i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
			A: 0xFFFF,
		}
	case RGB192f:
		return GrayA64i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
			A: 0xFFFF,
		}
	case RGBA128i:
		return GrayA64i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
			A: int32(c.A),
		}
	case RGBA128f:
		return GrayA64i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
			A: int32(c.A),
		}
	case RGBA256i:
		return GrayA64i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
			A: int32(c.A),
		}
	case RGBA256f:
		return GrayA64i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
			A: int32(c.A),
		}
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
	switch c := c.(type) {
	case Gray32i:
		return GrayA64f{
			Y: float32(c.Y),
			A: 0xFFFF,
		}
	case Gray32f:
		return GrayA64f{
			Y: float32(c.Y),
			A: 0xFFFF,
		}
	case Gray64i:
		return GrayA64f{
			Y: float32(c.Y),
			A: 0xFFFF,
		}
	case Gray64f:
		return GrayA64f{
			Y: float32(c.Y),
			A: 0xFFFF,
		}
	case GrayA64i:
		return GrayA64f{
			Y: float32(c.Y),
			A: float32(c.A),
		}
	case GrayA64f:
		return GrayA64f{
			Y: float32(c.Y),
			A: float32(c.A),
		}
	case GrayA128i:
		return GrayA64f{
			Y: float32(c.Y),
			A: float32(c.A),
		}
	case GrayA128f:
		return GrayA64f{
			Y: float32(c.Y),
			A: float32(c.A),
		}
	case RGB96i:
		return GrayA64f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
			A: 0xFFFF,
		}
	case RGB96f:
		return GrayA64f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
			A: 0xFFFF,
		}
	case RGB192i:
		return GrayA64f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
			A: 0xFFFF,
		}
	case RGB192f:
		return GrayA64f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
			A: 0xFFFF,
		}
	case RGBA128i:
		return GrayA64f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
			A: float32(c.A),
		}
	case RGBA128f:
		return GrayA64f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
			A: float32(c.A),
		}
	case RGBA256i:
		return GrayA64f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
			A: float32(c.A),
		}
	case RGBA256f:
		return GrayA64f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
			A: float32(c.A),
		}
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
	switch c := c.(type) {
	case Gray32i:
		return GrayA128i{
			Y: int64(c.Y),
			A: 0xFFFF,
		}
	case Gray32f:
		return GrayA128i{
			Y: int64(c.Y),
			A: 0xFFFF,
		}
	case Gray64i:
		return GrayA128i{
			Y: int64(c.Y),
			A: 0xFFFF,
		}
	case Gray64f:
		return GrayA128i{
			Y: int64(c.Y),
			A: 0xFFFF,
		}
	case GrayA64i:
		return GrayA128i{
			Y: int64(c.Y),
			A: int64(c.A),
		}
	case GrayA64f:
		return GrayA128i{
			Y: int64(c.Y),
			A: int64(c.A),
		}
	case GrayA128i:
		return GrayA128i{
			Y: int64(c.Y),
			A: int64(c.A),
		}
	case GrayA128f:
		return GrayA128i{
			Y: int64(c.Y),
			A: int64(c.A),
		}
	case RGB96i:
		return GrayA128i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
			A: 0xFFFF,
		}
	case RGB96f:
		return GrayA128i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
			A: 0xFFFF,
		}
	case RGB192i:
		return GrayA128i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
			A: 0xFFFF,
		}
	case RGB192f:
		return GrayA128i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
			A: 0xFFFF,
		}
	case RGBA128i:
		return GrayA128i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
			A: int64(c.A),
		}
	case RGBA128f:
		return GrayA128i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
			A: int64(c.A),
		}
	case RGBA256i:
		return GrayA128i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
			A: int64(c.A),
		}
	case RGBA256f:
		return GrayA128i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
			A: int64(c.A),
		}
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
	switch c := c.(type) {
	case Gray32i:
		return GrayA128f{
			Y: float64(c.Y),
			A: 0xFFFF,
		}
	case Gray32f:
		return GrayA128f{
			Y: float64(c.Y),
			A: 0xFFFF,
		}
	case Gray64i:
		return GrayA128f{
			Y: float64(c.Y),
			A: 0xFFFF,
		}
	case Gray64f:
		return GrayA128f{
			Y: float64(c.Y),
			A: 0xFFFF,
		}
	case GrayA64i:
		return GrayA128f{
			Y: float64(c.Y),
			A: float64(c.A),
		}
	case GrayA64f:
		return GrayA128f{
			Y: float64(c.Y),
			A: float64(c.A),
		}
	case GrayA128i:
		return GrayA128f{
			Y: float64(c.Y),
			A: float64(c.A),
		}
	case GrayA128f:
		return GrayA128f{
			Y: float64(c.Y),
			A: float64(c.A),
		}
	case RGB96i:
		return GrayA128f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
			A: 0xFFFF,
		}
	case RGB96f:
		return GrayA128f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
			A: 0xFFFF,
		}
	case RGB192i:
		return GrayA128f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
			A: 0xFFFF,
		}
	case RGB192f:
		return GrayA128f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
			A: 0xFFFF,
		}
	case RGBA128i:
		return GrayA128f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
			A: float64(c.A),
		}
	case RGBA128f:
		return GrayA128f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
			A: float64(c.A),
		}
	case RGBA256i:
		return GrayA128f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
			A: float64(c.A),
		}
	case RGBA256f:
		return GrayA128f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
			A: float64(c.A),
		}
	}
	r, g, b, a := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return GrayA128f{
		Y: float64(y),
		A: float64(a),
	}
}
