// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

type (
	colorGrayA    [2]uint8
	colorGrayA16  [2]uint16
	colorGrayA32f [2]float32
	colorRGB      [3]uint8
	colorRGB48    [3]uint16
	colorRGB96f   [3]float32
	colorRGBA128f [4]float32
)

func colorRgbToGray(r, g, b uint32) uint32 {
	y := (299*r + 587*g + 114*b + 500) / 1000
	return y
}

func (c colorGrayA) RGBA() (r, g, b, a uint32) {
	r = uint32(c[0]) << 8
	g = uint32(c[0]) << 8
	b = uint32(c[0]) << 8
	a = uint32(c[1]) << 8
	return
}

func (c colorGrayA16) RGBA() (r, g, b, a uint32) {
	r = uint32(c[0])
	g = uint32(c[0])
	b = uint32(c[0])
	a = uint32(c[1])
	return
}

func (c colorGrayA32f) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c[0]))
	g = uint32(uint16(c[0]))
	b = uint32(uint16(c[0]))
	a = uint32(uint16(c[1]))
	return
}

func (c colorRGB) RGBA() (r, g, b, a uint32) {
	r = uint32(c[0]) << 8
	g = uint32(c[1]) << 8
	b = uint32(c[2]) << 8
	a = 0xFFFF
	return
}

func (c colorRGB48) RGBA() (r, g, b, a uint32) {
	r = uint32(c[0])
	g = uint32(c[1])
	b = uint32(c[2])
	a = 0xFFFF
	return
}

func (c colorRGB96f) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c[0]))
	g = uint32(uint16(c[1]))
	b = uint32(uint16(c[2]))
	a = 0xFFFF
	return
}

func (c colorRGBA128f) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c[0]))
	g = uint32(uint16(c[1]))
	b = uint32(uint16(c[2]))
	a = uint32(uint16(c[3]))
	return
}
