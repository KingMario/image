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

type RGB96f struct {
	M struct {
		Pix    []uint8 // []struct{ R, G, B float32 }
		Stride int
		Rect   image.Rectangle
	}
}

// NewRGB96f returns a new RGB96f with the given bounds.
func NewRGB96f(r image.Rectangle) *RGB96f {
	return new(RGB96f).Init(make([]uint8, 1*r.Dx()*r.Dy()), 1*r.Dx(), r)
}

func (p *RGB96f) Init(pix []uint8, stride int, rect image.Rectangle) *RGB96f {
	*p = RGB96f{
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

func (p *RGB96f) BaseType() image.Image { return asBaseType(p) }
func (p *RGB96f) Pix() []byte           { return p.M.Pix }
func (p *RGB96f) Stride() int           { return p.M.Stride }
func (p *RGB96f) Rect() image.Rectangle { return p.M.Rect }
func (p *RGB96f) Channels() int         { return 1 }
func (p *RGB96f) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *RGB96f) ColorModel() color.Model { return colorExt.RGB96fModel }

func (p *RGB96f) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGB96f) At(x, y int) color.Color {
	return p.RGB96fAt(x, y)
}

func (p *RGB96f) RGB96fAt(x, y int) colorExt.RGB96f {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorExt.RGB96f{}
	}
	i := p.PixOffset(x, y)
	return *(*colorExt.RGB96f)(unsafe.Pointer(&p.M.Pix[i]))
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGB96f) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*4
}

func (p *RGB96f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := p.ColorModel().Convert(c).(colorExt.RGB96f)
	*(*colorExt.RGB96f)(unsafe.Pointer(&p.M.Pix[i])) = c1
	return
}

func (p *RGB96f) SetRGB96f(x, y int, c colorExt.RGB96f) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	*(*colorExt.RGB96f)(unsafe.Pointer(&p.M.Pix[i])) = c
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

func (p *RGB96f) Draw(r image.Rectangle, src Image, sp image.Point) Image {
	panic("TODO")
}
