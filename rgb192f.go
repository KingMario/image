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

type RGB192f struct {
	M struct {
		Pix    []uint8 // []struct{ R, G, B float64 }
		Stride int
		Rect   image.Rectangle
	}
}

// NewRGB192f returns a new RGB192f with the given bounds.
func NewRGB192f(r image.Rectangle) *RGB192f {
	return new(RGB192f).Init(make([]uint8, 1*r.Dx()*r.Dy()), 1*r.Dx(), r)
}

func (p *RGB192f) Init(pix []uint8, stride int, rect image.Rectangle) *RGB192f {
	*p = RGB192f{
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

func (p *RGB192f) BaseType() image.Image { return asBaseType(p) }
func (p *RGB192f) Pix() []byte           { return p.M.Pix }
func (p *RGB192f) Stride() int           { return p.M.Stride }
func (p *RGB192f) Rect() image.Rectangle { return p.M.Rect }
func (p *RGB192f) Channels() int         { return 1 }
func (p *RGB192f) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *RGB192f) ColorModel() color.Model { return colorExt.RGB192fModel }

func (p *RGB192f) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGB192f) At(x, y int) color.Color {
	return p.RGB192fAt(x, y)
}

func (p *RGB192f) RGB192fAt(x, y int) colorExt.RGB192f {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorExt.RGB192f{}
	}
	i := p.PixOffset(x, y)
	return *(*colorExt.RGB192f)(unsafe.Pointer(&p.M.Pix[i]))
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGB192f) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*4
}

func (p *RGB192f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := p.ColorModel().Convert(c).(colorExt.RGB192f)
	*(*colorExt.RGB192f)(unsafe.Pointer(&p.M.Pix[i])) = c1
	return
}

func (p *RGB192f) SetRGB192f(x, y int, c colorExt.RGB192f) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	*(*colorExt.RGB192f)(unsafe.Pointer(&p.M.Pix[i])) = c
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGB192f) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGB192f{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(RGB192f).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGB192f) Opaque() bool {
	return true
}

func (p *RGB192f) Draw(r image.Rectangle, src Image, sp image.Point) Image {
	panic("TODO")
}
