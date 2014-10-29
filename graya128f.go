// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
	"image/color"
	"reflect"
)

type GrayA128f struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// NewGrayA128f returns a new GrayA128f with the given bounds.
func NewGrayA128f(r image.Rectangle) *GrayA128f {
	return new(GrayA128f).Init(make([]uint8, 2*r.Dx()*r.Dy()), 2*r.Dx(), r)
}

func (p *GrayA128f) Init(pix []uint8, stride int, rect image.Rectangle) *GrayA128f {
	*p = GrayA128f{
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

func (p *GrayA128f) BaseType() image.Image { return p }
func (p *GrayA128f) Pix() []byte           { return p.M.Pix }
func (p *GrayA128f) Stride() int           { return p.M.Stride }
func (p *GrayA128f) Rect() image.Rectangle { return p.M.Rect }
func (p *GrayA128f) Channels() int         { return 2 }
func (p *GrayA128f) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *GrayA128f) ColorModel() color.Model { return color.RGBAModel }

func (p *GrayA128f) Bounds() image.Rectangle { return p.M.Rect }

func (p *GrayA128f) At(x, y int) color.Color {
	c := p.GrayAAt(x, y)
	return color.RGBA{
		R: c[0],
		G: c[0],
		B: c[0],
		A: c[1],
	}
}

func (p *GrayA128f) GrayAAt(x, y int) [2]uint8 {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return [2]uint8{}
	}
	i := p.PixOffset(x, y)
	return [2]uint8{
		p.M.Pix[i+0],
		p.M.Pix[i+1],
	}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *GrayA128f) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*2
}

func (p *GrayA128f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	rr, gg, bb, aa := c.RGBA()
	yy := colorRgbToGray(rr, gg, bb)
	p.M.Pix[i+0] = uint8(yy >> 8)
	p.M.Pix[i+1] = uint8(aa >> 8)
	return
}

func (p *GrayA128f) SetGrayA(x, y int, c [2]uint8) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.M.Pix[i+0] = c[0]
	p.M.Pix[i+1] = c[1]
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *GrayA128f) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &GrayA128f{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(GrayA128f).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *GrayA128f) Opaque() bool {
	if p.M.Rect.Empty() {
		return true
	}
	i0, i1 := 1, p.M.Rect.Dx()*2
	for y := p.M.Rect.Min.Y; y < p.M.Rect.Max.Y; y++ {
		for i := i0; i < i1; i += 2 {
			if p.M.Pix[i] != 0xFF {
				return false
			}
		}
		i0 += p.M.Stride
		i1 += p.M.Stride
	}
	return true
}

func (p *GrayA128f) CopyFrom(m image.Image) *GrayA128f {
	panic("TODO")
}

func (p *GrayA128f) Draw(r image.Rectangle, src Image, sp image.Point) Image {
	panic("TODO")
}
