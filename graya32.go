// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
	"image/color"
	"reflect"
)

type GrayA32 struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// NewGrayA32 returns a new GrayA32 with the given bounds.
func NewGrayA32(r image.Rectangle) *GrayA32 {
	return new(GrayA32).Init(make([]uint8, 4*r.Dx()*r.Dy()), 4*r.Dx(), r)
}

func (p *GrayA32) Init(pix []uint8, stride int, rect image.Rectangle) *GrayA32 {
	*p = GrayA32{
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

func (p *GrayA32) BaseType() image.Image { return p }
func (p *GrayA32) Pix() []byte           { return p.M.Pix }
func (p *GrayA32) Stride() int           { return p.M.Stride }
func (p *GrayA32) Rect() image.Rectangle { return p.M.Rect }
func (p *GrayA32) Channels() int         { return 2 }
func (p *GrayA32) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *GrayA32) ColorModel() color.Model { return color.RGBA64Model }

func (p *GrayA32) Bounds() image.Rectangle { return p.M.Rect }

func (p *GrayA32) At(x, y int) color.Color {
	c := p.GrayA32At(x, y)
	return color.RGBA64{
		R: c[0],
		G: c[0],
		B: c[0],
		A: c[1],
	}
}

func (p *GrayA32) GrayA32At(x, y int) [2]uint16 {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return [2]uint16{}
	}
	i := p.PixOffset(x, y)
	return [2]uint16{
		uint16(p.M.Pix[i+0])<<8 | uint16(p.M.Pix[i+1]),
		0xFFFF,
	}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *GrayA32) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*4
}

func (p *GrayA32) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	yy := color.Gray16Model.Convert(c).(color.Gray16).Y
	_, _, _, aa := c.RGBA()
	p.M.Pix[i+0] = uint8(yy >> 8)
	p.M.Pix[i+1] = uint8(yy)
	p.M.Pix[i+2] = uint8(aa >> 8)
	p.M.Pix[i+3] = uint8(aa)
	return
}

func (p *GrayA32) SetGrayA32(x, y int, c [2]uint16) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.M.Pix[i+0] = uint8(c[0] >> 8)
	p.M.Pix[i+1] = uint8(c[0])
	p.M.Pix[i+2] = uint8(c[1] >> 8)
	p.M.Pix[i+3] = uint8(c[1])
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *GrayA32) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &GrayA32{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(GrayA32).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *GrayA32) Opaque() bool {
	if p.M.Rect.Empty() {
		return true
	}
	i0, i1 := 2, p.M.Rect.Dx()*4
	for y := p.M.Rect.Min.Y; y < p.M.Rect.Max.Y; y++ {
		for i := i0; i < i1; i += 4 {
			if p.M.Pix[i+0] != 0xff || p.M.Pix[i+1] != 0xff {
				return false
			}
		}
		i0 += p.M.Stride
		i1 += p.M.Stride
	}
	return true
}

func (p *GrayA32) CopyFrom(m image.Image) *GrayA32 {
	panic("TODO")
}
