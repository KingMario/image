// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
	"image/color"
	"reflect"
)

type RGB48 struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// NewRGB48 returns a new RGB48 with the given bounds.
func NewRGB48(r image.Rectangle) *RGB48 {
	return new(RGB48).Init(make([]uint8, 6*r.Dx()*r.Dy()), 6*r.Dx(), r)
}

func (p *RGB48) Init(pix []uint8, stride int, rect image.Rectangle) *RGB48 {
	*p = RGB48{
		M: struct {
			Pix    []uint8
			Stride int
			Rect   image.Rectangle
		}{
			Pix:    p.M.Pix,
			Stride: p.M.Stride,
			Rect:   p.M.Rect,
		},
	}
	return p
}

func (p *RGB48) BaseType() image.Image { return p }
func (p *RGB48) Pix() []byte           { return p.M.Pix }
func (p *RGB48) Stride() int           { return p.M.Stride }
func (p *RGB48) Rect() image.Rectangle { return p.M.Rect }
func (p *RGB48) Channels() int         { return 3 }
func (p *RGB48) Depth() reflect.Kind   { return reflect.Uint16 }

func (p *RGB48) ColorModel() color.Model { return color.RGBA64Model }

func (p *RGB48) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGB48) At(x, y int) color.Color {
	c := p.RGB48At(x, y)
	return color.RGBA64{
		R: c[0],
		G: c[1],
		B: c[2],
		A: 0xFFFF,
	}
}

func (p *RGB48) RGB48At(x, y int) [3]uint16 {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return [3]uint16{}
	}
	i := p.PixOffset(x, y)
	return [3]uint16{
		uint16(p.M.Pix[i+0])<<8 | uint16(p.M.Pix[i+1]),
		uint16(p.M.Pix[i+2])<<8 | uint16(p.M.Pix[i+3]),
		uint16(p.M.Pix[i+4])<<8 | uint16(p.M.Pix[i+5]),
	}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGB48) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*6
}

func (p *RGB48) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := color.RGBA64Model.Convert(c).(color.RGBA64)
	p.M.Pix[i+0] = uint8(c1.R >> 8)
	p.M.Pix[i+1] = uint8(c1.R)
	p.M.Pix[i+2] = uint8(c1.G >> 8)
	p.M.Pix[i+3] = uint8(c1.G)
	p.M.Pix[i+4] = uint8(c1.B >> 8)
	p.M.Pix[i+5] = uint8(c1.B)
	return
}

func (p *RGB48) SetRGB48(x, y int, c [3]uint16) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.M.Pix[i+0] = uint8(c[0] >> 8)
	p.M.Pix[i+1] = uint8(c[0])
	p.M.Pix[i+2] = uint8(c[1] >> 8)
	p.M.Pix[i+3] = uint8(c[1])
	p.M.Pix[i+4] = uint8(c[2] >> 8)
	p.M.Pix[i+5] = uint8(c[2])
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGB48) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGB48{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(RGB48).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGB48) Opaque() bool {
	return true
}

func (p *RGB48) CopyFrom(m image.Image) *RGB48 {
	panic("TODO")
}

func (p *RGB48) Draw(r image.Rectangle, src Image, sp image.Point) Image {
	panic("TODO")
}
