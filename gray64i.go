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

type Gray64i struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// NewGray64i returns a new Gray64i with the given bounds.
func NewGray64i(r image.Rectangle) *Gray64i {
	return new(Gray64i).Init(make([]uint8, 4*r.Dx()*r.Dy()), 4*r.Dx(), r)
}

func (p *Gray64i) Init(pix []uint8, stride int, rect image.Rectangle) *Gray64i {
	*p = Gray64i{
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

func (p *Gray64i) BaseType() image.Image { return p }
func (p *Gray64i) Pix() []byte           { return p.M.Pix }
func (p *Gray64i) Stride() int           { return p.M.Stride }
func (p *Gray64i) Rect() image.Rectangle { return p.M.Rect }
func (p *Gray64i) Channels() int         { return 1 }
func (p *Gray64i) Depth() reflect.Kind   { return reflect.Float32 }

func (p *Gray64i) ColorModel() color.Model { return color.Gray16Model }

func (p *Gray64i) Bounds() image.Rectangle { return p.M.Rect }

func (p *Gray64i) At(x, y int) color.Color {
	return color.Gray16{
		Y: uint16(p.Gray32fAt(x, y)),
	}
}

func (p *Gray64i) Gray32fAt(x, y int) float32 {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return 0
	}
	i := p.PixOffset(x, y)
	v := math.Float32frombits(binary.BigEndian.Uint32(p.M.Pix[i:]))
	return v
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *Gray64i) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*4
}

func (p *Gray64i) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	v := float32(color.Gray16Model.Convert(c).(color.Gray16).Y)
	binary.BigEndian.PutUint32(p.M.Pix[i:], math.Float32bits(v))
	return
}

func (p *Gray64i) SetGray32f(x, y int, c float32) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	binary.BigEndian.PutUint32(p.M.Pix[i:], math.Float32bits(c))
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *Gray64i) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &Gray64i{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(Gray64i).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *Gray64i) Opaque() bool {
	return true
}

func (p *Gray64i) CopyFrom(m image.Image) *Gray64i {
	panic("TODO")
}

func (p *Gray64i) Draw(r image.Rectangle, src Image, sp image.Point) Image {
	panic("TODO")
}
