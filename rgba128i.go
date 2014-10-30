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

type RGBA128i struct {
	M struct {
		Pix    []uint8 // []struct{ R, G, B, A int32 }
		Stride int
		Rect   image.Rectangle
	}
}

// NewRGBA128i returns a new RGBA128i with the given bounds.
func NewRGBA128i(r image.Rectangle) *RGBA128i {
	return new(RGBA128i).Init(make([]uint8, 1*r.Dx()*r.Dy()), 1*r.Dx(), r)
}

func (p *RGBA128i) Init(pix []uint8, stride int, rect image.Rectangle) *RGBA128i {
	*p = RGBA128i{
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

func (p *RGBA128i) BaseType() image.Image { return asBaseType(p) }
func (p *RGBA128i) Pix() []byte           { return p.M.Pix }
func (p *RGBA128i) Stride() int           { return p.M.Stride }
func (p *RGBA128i) Rect() image.Rectangle { return p.M.Rect }
func (p *RGBA128i) Channels() int         { return 1 }
func (p *RGBA128i) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *RGBA128i) ColorModel() color.Model { return colorExt.RGBA128iModel }

func (p *RGBA128i) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGBA128i) At(x, y int) color.Color {
	return p.RGBA128iAt(x, y)
}

func (p *RGBA128i) RGBA128iAt(x, y int) colorExt.RGBA128i {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorExt.RGBA128i{}
	}
	i := p.PixOffset(x, y)
	return *(*colorExt.RGBA128i)(unsafe.Pointer(&p.M.Pix[i]))
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGBA128i) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*4
}

func (p *RGBA128i) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := p.ColorModel().Convert(c).(colorExt.RGBA128i)
	*(*colorExt.RGBA128i)(unsafe.Pointer(&p.M.Pix[i])) = c1
	return
}

func (p *RGBA128i) SetRGBA128i(x, y int, c colorExt.RGBA128i) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	*(*colorExt.RGBA128i)(unsafe.Pointer(&p.M.Pix[i])) = c
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGBA128i) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGBA128i{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(RGBA128i).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGBA128i) Opaque() bool {
	return true
}

func (p *RGBA128i) Draw(r image.Rectangle, src Image, sp image.Point) Image {
	panic("TODO")
}
