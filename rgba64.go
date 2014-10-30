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

type RGBA64 struct {
	M struct {
		Pix    []uint8 // []struct{ R, G, B, A uint16 }
		Stride int
		Rect   image.Rectangle
	}
}

// NewRGBA64 returns a new RGBA64 with the given bounds.
func NewRGBA64(r image.Rectangle) *RGBA64 {
	return new(RGBA64).Init(make([]uint8, 1*r.Dx()*r.Dy()), 1*r.Dx(), r)
}

func (p *RGBA64) Init(pix []uint8, stride int, rect image.Rectangle) *RGBA64 {
	*p = RGBA64{
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

func (p *RGBA64) BaseType() image.Image { return asBaseType(p) }
func (p *RGBA64) Pix() []byte           { return p.M.Pix }
func (p *RGBA64) Stride() int           { return p.M.Stride }
func (p *RGBA64) Rect() image.Rectangle { return p.M.Rect }
func (p *RGBA64) Channels() int         { return 1 }
func (p *RGBA64) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *RGBA64) ColorModel() color.Model { return colorExt.RGBA64Model }

func (p *RGBA64) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGBA64) At(x, y int) color.Color {
	return p.RGBA64At(x, y)
}

func (p *RGBA64) RGBA64At(x, y int) colorExt.RGBA64 {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorExt.RGBA64{}
	}
	i := p.PixOffset(x, y)
	return *(*colorExt.RGBA64)(unsafe.Pointer(&p.M.Pix[i]))
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGBA64) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*4
}

func (p *RGBA64) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := p.ColorModel().Convert(c).(colorExt.RGBA64)
	*(*colorExt.RGBA64)(unsafe.Pointer(&p.M.Pix[i])) = c1
	return
}

func (p *RGBA64) SetRGBA64(x, y int, c colorExt.RGBA64) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	*(*colorExt.RGBA64)(unsafe.Pointer(&p.M.Pix[i])) = c
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGBA64) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGBA64{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(RGBA64).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGBA64) Opaque() bool {
	return true
}

func (p *RGBA64) Draw(r image.Rectangle, src Image, sp image.Point) Image {
	panic("TODO")
}
