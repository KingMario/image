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

type RGBA256f struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// NewRGBA256f returns a new RGBA256f with the given bounds.
func NewRGBA256f(r image.Rectangle) *RGBA256f {
	return new(RGBA256f).Init(make([]uint8, 32*r.Dx()*r.Dy()), 32*r.Dx(), r)
}

func (p *RGBA256f) Init(pix []uint8, stride int, rect image.Rectangle) *RGBA256f {
	*p = RGBA256f{
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

func (p *RGBA256f) BaseType() image.Image { return asBaseType(p) }
func (p *RGBA256f) Pix() []byte           { return p.M.Pix }
func (p *RGBA256f) Stride() int           { return p.M.Stride }
func (p *RGBA256f) Rect() image.Rectangle { return p.M.Rect }
func (p *RGBA256f) Channels() int         { return 4 }
func (p *RGBA256f) Depth() reflect.Kind   { return reflect.Float64 }

func (p *RGBA256f) ColorModel() color.Model { return colorExt.RGBA256fModel }

func (p *RGBA256f) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGBA256f) At(x, y int) color.Color {
	return p.RGBA256fAt(x, y)
}

func (p *RGBA256f) RGBA256fAt(x, y int) colorExt.RGBA256f {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorExt.RGBA256f{}
	}
	i := p.PixOffset(x, y)
	return pRGBA256fAt(p.M.Pix[i:])
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGBA256f) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*32
}

func (p *RGBA256f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := colorExt.RGBA256fModel.Convert(c).(colorExt.RGBA256f)
	pSetRGBA256f(p.M.Pix[i:], c1)
	return
}

func (p *RGBA256f) SetRGBA256f(x, y int, c colorExt.RGBA256f) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	pSetRGBA256f(p.M.Pix[i:], c)
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGBA256f) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGBA256f{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(RGBA256f).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGBA256f) Opaque() bool {
	if p.M.Rect.Empty() {
		return true
	}
	for y := p.M.Rect.Min.Y; y < p.M.Rect.Max.Y; y++ {
		for x := p.M.Rect.Min.X; x < p.M.Rect.Max.X; x++ {
			if _, _, _, a := p.At(x, y).RGBA(); a == 0xFFFF {
				return false
			}
		}
	}
	return true
}
