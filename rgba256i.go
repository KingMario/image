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

type RGBA256i struct {
	M struct {
		Pix    []uint8 // []struct{ R, G, B, A int64 }
		Stride int
		Rect   image.Rectangle
	}
}

// NewRGBA256i returns a new RGBA256i with the given bounds.
func NewRGBA256i(r image.Rectangle) *RGBA256i {
	return new(RGBA256i).Init(make([]uint8, 1*r.Dx()*r.Dy()), 1*r.Dx(), r)
}

func (p *RGBA256i) Init(pix []uint8, stride int, rect image.Rectangle) *RGBA256i {
	*p = RGBA256i{
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

func (p *RGBA256i) BaseType() image.Image { return asBaseType(p) }
func (p *RGBA256i) Pix() []byte           { return p.M.Pix }
func (p *RGBA256i) Stride() int           { return p.M.Stride }
func (p *RGBA256i) Rect() image.Rectangle { return p.M.Rect }
func (p *RGBA256i) Channels() int         { return 1 }
func (p *RGBA256i) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *RGBA256i) ColorModel() color.Model { return colorExt.RGBA256iModel }

func (p *RGBA256i) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGBA256i) At(x, y int) color.Color {
	return p.RGBA256iAt(x, y)
}

func (p *RGBA256i) RGBA256iAt(x, y int) colorExt.RGBA256i {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorExt.RGBA256i{}
	}
	i := p.PixOffset(x, y)
	return *(*colorExt.RGBA256i)(unsafe.Pointer(&p.M.Pix[i]))
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGBA256i) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*4
}

func (p *RGBA256i) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := p.ColorModel().Convert(c).(colorExt.RGBA256i)
	*(*colorExt.RGBA256i)(unsafe.Pointer(&p.M.Pix[i])) = c1
	return
}

func (p *RGBA256i) SetRGBA256i(x, y int, c colorExt.RGBA256i) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	*(*colorExt.RGBA256i)(unsafe.Pointer(&p.M.Pix[i])) = c
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGBA256i) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGBA256i{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(RGBA256i).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGBA256i) Opaque() bool {
	return true
}

func (p *RGBA256i) Draw(r image.Rectangle, src Image, sp image.Point) Image {
	panic("TODO")
}
