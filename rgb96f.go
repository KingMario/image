// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"encoding/binary"
	"image"
	"image/color"
	"math"
	"reflect"
)

type RGB96f struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

func (p *RGB96f) BaseType() image.Image { return p }
func (p *RGB96f) Pix() []byte           { return p.M.Pix }
func (p *RGB96f) Stride() int           { return p.M.Stride }
func (p *RGB96f) Rect() image.Rectangle { return p.M.Rect }
func (p *RGB96f) Channels() int         { return 3 }
func (p *RGB96f) Depth() reflect.Kind   { return reflect.Float32 }

func (p *RGB96f) ColorModel() color.Model { return color.RGBA64Model }

func (p *RGB96f) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGB96f) At(x, y int) color.Color {
	c := p.RGB96fAt(x, y)
	return color.RGBA64{
		R: uint16(c[0]),
		G: uint16(c[1]),
		B: uint16(c[2]),
		A: 0xFFFF,
	}
}

func (p *RGB96f) RGB96fAt(x, y int) [3]float32 {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return [3]float32{}
	}
	i := p.PixOffset(x, y)
	bitsR := binary.BigEndian.Uint32(p.M.Pix[i+4*0:])
	bitsG := binary.BigEndian.Uint32(p.M.Pix[i+4*1:])
	bitsB := binary.BigEndian.Uint32(p.M.Pix[i+4*2:])
	r := math.Float32frombits(bitsR)
	g := math.Float32frombits(bitsG)
	b := math.Float32frombits(bitsB)
	return [3]float32{r, g, b}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGB96f) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*12
}

func (p *RGB96f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	v := color.RGBA64Model.Convert(c).(color.RGBA64)
	bitsR := math.Float32bits(float32(v.R))
	bitsG := math.Float32bits(float32(v.G))
	bitsB := math.Float32bits(float32(v.B))
	binary.BigEndian.PutUint32(p.M.Pix[i+4*0:], bitsR)
	binary.BigEndian.PutUint32(p.M.Pix[i+4*1:], bitsG)
	binary.BigEndian.PutUint32(p.M.Pix[i+4*2:], bitsB)
	return
}

func (p *RGB96f) SetRGB96f(x, y int, c [3]float32) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	bitsR := math.Float32bits(c[0])
	bitsG := math.Float32bits(c[1])
	bitsB := math.Float32bits(c[2])
	binary.BigEndian.PutUint32(p.M.Pix[i+4*0:], bitsR)
	binary.BigEndian.PutUint32(p.M.Pix[i+4*1:], bitsG)
	binary.BigEndian.PutUint32(p.M.Pix[i+4*2:], bitsB)
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGB96f) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGB96f{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(RGB96f).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGB96f) Opaque() bool {
	return true
}

// NewRGB96f returns a new RGB96f with the given bounds.
func NewRGB96f(r image.Rectangle) *RGB96f {
	return new(RGB96f).Init(make([]uint8, 12*r.Dx()*r.Dy()), 12*r.Dx(), r)
}

func (p *RGB96f) Init(pix []uint8, stride int, rect image.Rectangle) *RGB96f {
	*p = RGB96f{
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

func (p *RGB96f) CopyFrom(m image.Image) *RGB96f {
	panic("TODO")
}
