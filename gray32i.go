// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
	"image/color"
	"reflect"
	"unsafe"

	colorExt "github.com/chai2010/image/color"
)

type Gray32i struct {
	M struct {
		Pix    []uint8 // []struct{ Y int32 }
		Stride int
		Rect   image.Rectangle
	}
}

// NewGray32i returns a new Gray32i with the given bounds.
func NewGray32i(r image.Rectangle) *Gray32i {
	return new(Gray32i).Init(make([]uint8, 1*r.Dx()*r.Dy()), 1*r.Dx(), r)
}

func (p *Gray32i) Init(pix []uint8, stride int, rect image.Rectangle) *Gray32i {
	*p = Gray32i{
		M: struct {
			Pix    []uint8
			Stride int
			Rect   image.Rectangle
		}{
			Pix:    pix,
			Stride: stride,
			Rect:   rect,
		},
	}
	return p
}

func (p *Gray32i) BaseType() image.Image { return asBaseType(p) }
func (p *Gray32i) Pix() []byte           { return p.M.Pix }
func (p *Gray32i) Stride() int           { return p.M.Stride }
func (p *Gray32i) Rect() image.Rectangle { return p.M.Rect }
func (p *Gray32i) Channels() int         { return 1 }
func (p *Gray32i) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *Gray32i) ColorModel() color.Model { return colorExt.Gray32iModel }

func (p *Gray32i) Bounds() image.Rectangle { return p.M.Rect }

func (p *Gray32i) At(x, y int) color.Color {
	return p.Gray32iAt(x, y)
}

func (p *Gray32i) Gray32iAt(x, y int) colorExt.Gray32i {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorExt.Gray32i{}
	}
	i := p.PixOffset(x, y)
	return *(*colorExt.Gray32i)(unsafe.Pointer(&p.M.Pix[i]))
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *Gray32i) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*4
}

func (p *Gray32i) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := p.ColorModel().Convert(c).(colorExt.Gray32i)
	*(*colorExt.Gray32i)(unsafe.Pointer(&p.M.Pix[i])) = c1
	return
}

func (p *Gray32i) SetGray32i(x, y int, c colorExt.Gray32i) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	*(*colorExt.Gray32i)(unsafe.Pointer(&p.M.Pix[i])) = c
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *Gray32i) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &Gray32i{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(Gray32i).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *Gray32i) Opaque() bool {
	return true
}

func (p *Gray32i) Draw(r image.Rectangle, src Image, sp image.Point) Image {
	panic("TODO")
}
