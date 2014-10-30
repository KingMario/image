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

type GrayA128i struct {
	M struct {
		Pix    []uint8 // []struct{ Y, A int64 }
		Stride int
		Rect   image.Rectangle
	}
}

// NewGrayA128i returns a new GrayA128i with the given bounds.
func NewGrayA128i(r image.Rectangle) *GrayA128i {
	return new(GrayA128i).Init(make([]uint8, 1*r.Dx()*r.Dy()), 1*r.Dx(), r)
}

func (p *GrayA128i) Init(pix []uint8, stride int, rect image.Rectangle) *GrayA128i {
	*p = GrayA128i{
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

func (p *GrayA128i) BaseType() image.Image { return asBaseType(p) }
func (p *GrayA128i) Pix() []byte           { return p.M.Pix }
func (p *GrayA128i) Stride() int           { return p.M.Stride }
func (p *GrayA128i) Rect() image.Rectangle { return p.M.Rect }
func (p *GrayA128i) Channels() int         { return 1 }
func (p *GrayA128i) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *GrayA128i) ColorModel() color.Model { return colorExt.GrayA128iModel }

func (p *GrayA128i) Bounds() image.Rectangle { return p.M.Rect }

func (p *GrayA128i) At(x, y int) color.Color {
	return p.GrayA128iAt(x, y)
}

func (p *GrayA128i) GrayA128iAt(x, y int) colorExt.GrayA128i {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorExt.GrayA128i{}
	}
	i := p.PixOffset(x, y)
	return *(*colorExt.GrayA128i)(unsafe.Pointer(&p.M.Pix[i]))
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *GrayA128i) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*4
}

func (p *GrayA128i) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := p.ColorModel().Convert(c).(colorExt.GrayA128i)
	*(*colorExt.GrayA128i)(unsafe.Pointer(&p.M.Pix[i])) = c1
	return
}

func (p *GrayA128i) SetGrayA128i(x, y int, c colorExt.GrayA128i) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	*(*colorExt.GrayA128i)(unsafe.Pointer(&p.M.Pix[i])) = c
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *GrayA128i) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &GrayA128i{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(GrayA128i).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *GrayA128i) Opaque() bool {
	return true
}

func (p *GrayA128i) Draw(r image.Rectangle, src Image, sp image.Point) Image {
	panic("TODO")
}
