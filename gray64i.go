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

type Gray64i struct {
	M struct {
		Pix    []uint8 // []struct{ Y int64 }
		Stride int
		Rect   image.Rectangle
	}
}

// NewGray64i returns a new Gray64i with the given bounds.
func NewGray64i(r image.Rectangle) *Gray64i {
	return new(Gray64i).Init(make([]uint8, 1*r.Dx()*r.Dy()), 1*r.Dx(), r)
}

func (p *Gray64i) Init(pix []uint8, stride int, rect image.Rectangle) *Gray64i {
	*p = Gray64i{
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

func (p *Gray64i) BaseType() image.Image { return asBaseType(p) }
func (p *Gray64i) Pix() []byte           { return p.M.Pix }
func (p *Gray64i) Stride() int           { return p.M.Stride }
func (p *Gray64i) Rect() image.Rectangle { return p.M.Rect }
func (p *Gray64i) Channels() int         { return 1 }
func (p *Gray64i) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *Gray64i) ColorModel() color.Model { return colorExt.Gray64iModel }

func (p *Gray64i) Bounds() image.Rectangle { return p.M.Rect }

func (p *Gray64i) At(x, y int) color.Color {
	return p.Gray64iAt(x, y)
}

func (p *Gray64i) Gray64iAt(x, y int) colorExt.Gray64i {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorExt.Gray64i{}
	}
	i := p.PixOffset(x, y)
	return *(*colorExt.Gray64i)(unsafe.Pointer(&p.M.Pix[i]))
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
	c1 := p.ColorModel().Convert(c).(colorExt.Gray64i)
	*(*colorExt.Gray64i)(unsafe.Pointer(&p.M.Pix[i])) = c1
	return
}

func (p *Gray64i) SetGray64i(x, y int, c colorExt.Gray64i) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	*(*colorExt.Gray64i)(unsafe.Pointer(&p.M.Pix[i])) = c
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

func (p *Gray64i) Draw(r image.Rectangle, src Image, sp image.Point) Image {
	panic("TODO")
}
