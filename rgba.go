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

type RGBA struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// NewRGBA returns a new RGBA with the given bounds.
func NewRGBA(r image.Rectangle) *RGBA {
	return new(RGBA).Init(make([]uint8, 4*r.Dx()*r.Dy()), 4*r.Dx(), r)
}

func (p *RGBA) Init(pix []uint8, stride int, rect image.Rectangle) *RGBA {
	*p = RGBA{
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

func (p *RGBA) BaseType() image.Image { return asBaseType(p) }
func (p *RGBA) Pix() []byte           { return p.M.Pix }
func (p *RGBA) Stride() int           { return p.M.Stride }
func (p *RGBA) Rect() image.Rectangle { return p.M.Rect }
func (p *RGBA) Channels() int         { return 4 }
func (p *RGBA) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *RGBA) ColorModel() color.Model { return colorExt.RGBAModel }

func (p *RGBA) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGBA) At(x, y int) color.Color {
	return p.RGBAAt(x, y)
}

func (p *RGBA) RGBAAt(x, y int) colorExt.RGBA {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorExt.RGBA{}
	}
	i := p.PixOffset(x, y)
	return pRGBAAt(p.M.Pix[i:])
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGBA) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*4
}

func (p *RGBA) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := colorExt.RGBAModel.Convert(c).(colorExt.RGBA)
	pSetRGBA(p.M.Pix[i:], c1)
	return
}

func (p *RGBA) SetRGBA(x, y int, c colorExt.RGBA) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	pSetRGBA(p.M.Pix[i:], c)
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGBA) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGBA{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(RGBA).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGBA) Opaque() bool {
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
