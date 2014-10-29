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

type Gray64f struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// NewGray64f returns a new Gray64f with the given bounds.
func NewGray64f(r image.Rectangle) *Gray64f {
	return new(Gray64f).Init(make([]uint8, 4*r.Dx()*r.Dy()), 4*r.Dx(), r)
}

func (p *Gray64f) Init(pix []uint8, stride int, rect image.Rectangle) *Gray64f {
	*p = Gray64f{
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

func (p *Gray64f) BaseType() image.Image { return p }
func (p *Gray64f) Pix() []byte           { return p.M.Pix }
func (p *Gray64f) Stride() int           { return p.M.Stride }
func (p *Gray64f) Rect() image.Rectangle { return p.M.Rect }
func (p *Gray64f) Channels() int         { return 1 }
func (p *Gray64f) Depth() reflect.Kind   { return reflect.Float32 }

func (p *Gray64f) ColorModel() color.Model { return color.Gray16Model }

func (p *Gray64f) Bounds() image.Rectangle { return p.M.Rect }

func (p *Gray64f) At(x, y int) color.Color {
	return color.Gray16{
		Y: uint16(p.Gray32fAt(x, y)),
	}
}

func (p *Gray64f) Gray32fAt(x, y int) float32 {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return 0
	}
	i := p.PixOffset(x, y)
	v := math.Float32frombits(binary.BigEndian.Uint32(p.M.Pix[i:]))
	return v
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *Gray64f) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*4
}

func (p *Gray64f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	v := float32(color.Gray16Model.Convert(c).(color.Gray16).Y)
	binary.BigEndian.PutUint32(p.M.Pix[i:], math.Float32bits(v))
	return
}

func (p *Gray64f) SetGray32f(x, y int, c float32) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	binary.BigEndian.PutUint32(p.M.Pix[i:], math.Float32bits(c))
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *Gray64f) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &Gray64f{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(Gray64f).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *Gray64f) Opaque() bool {
	return true
}

func (p *Gray64f) CopyFrom(m image.Image) *Gray64f {
	panic("TODO")
}

func (p *Gray64f) Draw(r image.Rectangle, src Image, sp image.Point) Image {
	panic("TODO")
}
