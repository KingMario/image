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

type Gray32f struct {
	M struct {
		Pix    []uint8 // []struct{ Y float32 }
		Stride int
		Rect   image.Rectangle
	}
}

// NewGray32f returns a new Gray32f with the given bounds.
func NewGray32f(r image.Rectangle) *Gray32f {
	return new(Gray32f).Init(make([]uint8, 1*r.Dx()*r.Dy()), 1*r.Dx(), r)
}

func (p *Gray32f) Init(pix []uint8, stride int, rect image.Rectangle) *Gray32f {
	*p = Gray32f{
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

func (p *Gray32f) BaseType() image.Image { return asBaseType(p) }
func (p *Gray32f) Pix() []byte           { return p.M.Pix }
func (p *Gray32f) Stride() int           { return p.M.Stride }
func (p *Gray32f) Rect() image.Rectangle { return p.M.Rect }
func (p *Gray32f) Channels() int         { return 1 }
func (p *Gray32f) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *Gray32f) ColorModel() color.Model { return colorExt.Gray32fModel }

func (p *Gray32f) Bounds() image.Rectangle { return p.M.Rect }

func (p *Gray32f) At(x, y int) color.Color {
	return p.Gray32fAt(x, y)
}

func (p *Gray32f) Gray32fAt(x, y int) colorExt.Gray32f {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorExt.Gray32f{}
	}
	i := p.PixOffset(x, y)
	return *(*colorExt.Gray32f)(unsafe.Pointer(&p.M.Pix[i]))
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *Gray32f) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*4
}

func (p *Gray32f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := p.ColorModel().Convert(c).(colorExt.Gray32f)
	*(*colorExt.Gray32f)(unsafe.Pointer(&p.M.Pix[i])) = c1
	return
}

func (p *Gray32f) SetGray32f(x, y int, c colorExt.Gray32f) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	*(*colorExt.Gray32f)(unsafe.Pointer(&p.M.Pix[i])) = c
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *Gray32f) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &Gray32f{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(Gray32f).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *Gray32f) Opaque() bool {
	return true
}

func (p *Gray32f) Draw(r image.Rectangle, src Image, sp image.Point) Image {
	panic("TODO")
}
