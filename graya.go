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

type GrayA struct {
	M struct {
		Pix    []uint8 // []struct{ Y, A uint8 }
		Stride int
		Rect   image.Rectangle
	}
}

// NewGrayA returns a new GrayA with the given bounds.
func NewGrayA(r image.Rectangle) *GrayA {
	return new(GrayA).Init(make([]uint8, 1*r.Dx()*r.Dy()), 1*r.Dx(), r)
}

func (p *GrayA) Init(pix []uint8, stride int, rect image.Rectangle) *GrayA {
	*p = GrayA{
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

func (p *GrayA) BaseType() image.Image { return asBaseType(p) }
func (p *GrayA) Pix() []byte           { return p.M.Pix }
func (p *GrayA) Stride() int           { return p.M.Stride }
func (p *GrayA) Rect() image.Rectangle { return p.M.Rect }
func (p *GrayA) Channels() int         { return 1 }
func (p *GrayA) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *GrayA) ColorModel() color.Model { return colorExt.GrayAModel }

func (p *GrayA) Bounds() image.Rectangle { return p.M.Rect }

func (p *GrayA) At(x, y int) color.Color {
	return p.GrayAAt(x, y)
}

func (p *GrayA) GrayAAt(x, y int) colorExt.GrayA {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorExt.GrayA{}
	}
	i := p.PixOffset(x, y)
	return *(*colorExt.GrayA)(unsafe.Pointer(&p.M.Pix[i]))
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *GrayA) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*4
}

func (p *GrayA) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := p.ColorModel().Convert(c).(colorExt.GrayA)
	*(*colorExt.GrayA)(unsafe.Pointer(&p.M.Pix[i])) = c1
	return
}

func (p *GrayA) SetGrayA(x, y int, c colorExt.GrayA) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	*(*colorExt.GrayA)(unsafe.Pointer(&p.M.Pix[i])) = c
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *GrayA) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &GrayA{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(GrayA).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *GrayA) Opaque() bool {
	return true
}

func (p *GrayA) Draw(r image.Rectangle, src Image, sp image.Point) Image {
	panic("TODO")
}
