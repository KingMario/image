// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Auto Generated By 'go generate', DONOT EDIT!!!

package image

import (
	"image"
	"image/color"
	"reflect"

	colorExt "github.com/chai2010/image/color"
)

type RGB48 struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// NewRGB48 returns a new RGB48 with the given bounds.
func NewRGB48(r image.Rectangle) *RGB48 {
	return new(RGB48).Init(make([]uint8, 6*r.Dx()*r.Dy()), 6*r.Dx(), r)
}

func (p *RGB48) Init(pix []uint8, stride int, rect image.Rectangle) *RGB48 {
	*p = RGB48{
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

func (p *RGB48) BaseType() image.Image { return asBaseType(p) }
func (p *RGB48) Pix() []byte           { return p.M.Pix }
func (p *RGB48) Stride() int           { return p.M.Stride }
func (p *RGB48) Rect() image.Rectangle { return p.M.Rect }
func (p *RGB48) Channels() int         { return 3 }
func (p *RGB48) Depth() reflect.Kind   { return reflect.Uint16 }

func (p *RGB48) ColorModel() color.Model { return colorExt.RGB48Model }

func (p *RGB48) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGB48) At(x, y int) color.Color {
	return p.RGB48At(x, y)
}

func (p *RGB48) RGB48At(x, y int) colorExt.RGB48 {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorExt.RGB48{}
	}
	i := p.PixOffset(x, y)
	return pRGB48At(p.M.Pix[i:])
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGB48) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*6
}

func (p *RGB48) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := colorExt.RGB48Model.Convert(c).(colorExt.RGB48)
	pSetRGB48(p.M.Pix[i:], c1)
	return
}

func (p *RGB48) SetRGB48(x, y int, c colorExt.RGB48) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	pSetRGB48(p.M.Pix[i:], c)
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGB48) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGB48{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(RGB48).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGB48) Opaque() bool {
	return true
}
