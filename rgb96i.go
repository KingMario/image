// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Auto Generated By 'go generate', DONOT EDIT!!!

package image

import (
	"image"
	"image/color"
	"reflect"
	"unsafe"

	colorExt "github.com/chai2010/image/color"
)

type RGB96i struct {
	M struct {
		Pix    []uint8 // []struct{ R, G, B int32 }
		Stride int
		Rect   image.Rectangle
	}
}

// NewRGB96i returns a new RGB96i with the given bounds.
func NewRGB96i(r image.Rectangle) *RGB96i {
	return new(RGB96i).Init(make([]uint8, 12*r.Dx()*r.Dy()), 12*r.Dx(), r)
}

func (p *RGB96i) Init(pix []uint8, stride int, rect image.Rectangle) *RGB96i {
	*p = RGB96i{
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

func (p *RGB96i) BaseType() image.Image { return asBaseType(p) }
func (p *RGB96i) Pix() []byte           { return p.M.Pix }
func (p *RGB96i) Stride() int           { return p.M.Stride }
func (p *RGB96i) Rect() image.Rectangle { return p.M.Rect }
func (p *RGB96i) Channels() int         { return 3 }
func (p *RGB96i) Depth() reflect.Kind   { return reflect.Int32 }

func (p *RGB96i) ColorModel() color.Model { return colorExt.RGB96iModel }

func (p *RGB96i) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGB96i) At(x, y int) color.Color {
	return p.RGB96iAt(x, y)
}

func (p *RGB96i) RGB96iAt(x, y int) colorExt.RGB96i {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorExt.RGB96i{}
	}
	i := p.PixOffset(x, y)
	return *(*colorExt.RGB96i)(unsafe.Pointer(&p.M.Pix[i]))
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGB96i) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*12
}

func (p *RGB96i) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := p.ColorModel().Convert(c).(colorExt.RGB96i)
	*(*colorExt.RGB96i)(unsafe.Pointer(&p.M.Pix[i])) = c1
	return
}

func (p *RGB96i) SetRGB96i(x, y int, c colorExt.RGB96i) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	*(*colorExt.RGB96i)(unsafe.Pointer(&p.M.Pix[i])) = c
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGB96i) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGB96i{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(RGB96i).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGB96i) Opaque() bool {
	return true
}

func (p *RGB96i) Draw(r image.Rectangle, src Image, sp image.Point) Image {
	panic("TODO")
}
