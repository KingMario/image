// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
	"image/color"
	"reflect"
)

type RGB192i struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// NewRGB192i returns a new RGB192i with the given bounds.
func NewRGB192i(r image.Rectangle) *RGB192i {
	return new(RGB192i).Init(make([]uint8, 3*r.Dx()*r.Dy()), 3*r.Dx(), r)
}

func (p *RGB192i) Init(pix []uint8, stride int, rect image.Rectangle) *RGB192i {
	*p = RGB192i{
		M: struct {
			Pix    []uint8
			Stride int
			Rect   image.Rectangle
		}{
			Pix:    p.M.Pix,
			Stride: p.M.Stride,
			Rect:   p.M.Rect,
		},
	}
	return p
}

func (p *RGB192i) BaseType() image.Image { return p }
func (p *RGB192i) Pix() []byte           { return p.M.Pix }
func (p *RGB192i) Stride() int           { return p.M.Stride }
func (p *RGB192i) Rect() image.Rectangle { return p.M.Rect }
func (p *RGB192i) Channels() int         { return 3 }
func (p *RGB192i) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *RGB192i) ColorModel() color.Model { return color.RGBAModel }

func (p *RGB192i) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGB192i) At(x, y int) color.Color {
	c := p.RGBAt(x, y)
	return color.RGBA{
		R: c[0],
		G: c[1],
		B: c[2],
		A: 0xFF,
	}
}

func (p *RGB192i) RGBAt(x, y int) [3]uint8 {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return [3]uint8{}
	}
	i := p.PixOffset(x, y)
	return [3]uint8{
		p.M.Pix[i+0],
		p.M.Pix[i+1],
		p.M.Pix[i+2],
	}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGB192i) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*3
}

func (p *RGB192i) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := color.RGBAModel.Convert(c).(color.RGBA)
	p.M.Pix[i+0] = c1.R
	p.M.Pix[i+1] = c1.G
	p.M.Pix[i+2] = c1.B
	return
}

func (p *RGB192i) SetRGB(x, y int, c [3]uint8) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.M.Pix[i+0] = c[0]
	p.M.Pix[i+1] = c[1]
	p.M.Pix[i+2] = c[2]
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGB192i) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGB192i{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(RGB192i).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGB192i) Opaque() bool {
	return true
}

func (p *RGB192i) CopyFrom(m image.Image) *RGB192i {
	panic("TODO")
}

func (p *RGB192i) Draw(r image.Rectangle, src Image, sp image.Point) Image {
	panic("TODO")
}
