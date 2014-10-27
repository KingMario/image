// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"encoding/binary"
	"image"
	"image/color"
	"math"
	"reflect"
)

type GrayA64f struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// NewGrayA64f returns a new GrayA64f with the given bounds.
func NewGrayA64f(r image.Rectangle) *GrayA64f {
	return new(GrayA64f).Init(make([]uint8, 8*r.Dx()*r.Dy()), 8*r.Dx(), r)
}

func (p *GrayA64f) Init(pix []uint8, stride int, rect image.Rectangle) *GrayA64f {
	*p = GrayA64f{
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

func (p *GrayA64f) BaseType() image.Image { return p }
func (p *GrayA64f) Pix() []byte           { return p.M.Pix }
func (p *GrayA64f) Stride() int           { return p.M.Stride }
func (p *GrayA64f) Rect() image.Rectangle { return p.M.Rect }
func (p *GrayA64f) Channels() int         { return 2 }
func (p *GrayA64f) Depth() reflect.Kind   { return reflect.Float32 }

func (p *GrayA64f) ColorModel() color.Model { return color.RGBA64Model }

func (p *GrayA64f) Bounds() image.Rectangle { return p.M.Rect }

func (p *GrayA64f) At(x, y int) color.Color {
	c := p.GrayA64fAt(x, y)
	return color.RGBA64{
		R: uint16(c[0]),
		G: uint16(c[0]),
		B: uint16(c[0]),
		A: uint16(c[1]),
	}
}

func (p *GrayA64f) GrayA64fAt(x, y int) [2]float32 {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return [2]float32{}
	}
	i := p.PixOffset(x, y)
	v0 := math.Float32frombits(binary.BigEndian.Uint32(p.M.Pix[i+4*0:]))
	v1 := math.Float32frombits(binary.BigEndian.Uint32(p.M.Pix[i+4*1:]))
	return [2]float32{v0, v1}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *GrayA64f) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*8
}

func (p *GrayA64f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	yy := color.Gray16Model.Convert(c).(color.Gray16).Y
	_, _, _, aa := c.RGBA()
	binary.BigEndian.PutUint32(p.M.Pix[i+4*0:], math.Float32bits(float32(yy)))
	binary.BigEndian.PutUint32(p.M.Pix[i+4*1:], math.Float32bits(float32(aa)))
	return
}

func (p *GrayA64f) SetGrayA64f(x, y int, c [2]float32) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	binary.BigEndian.PutUint32(p.M.Pix[i+4*0:], math.Float32bits(c[0]))
	binary.BigEndian.PutUint32(p.M.Pix[i+4*1:], math.Float32bits(c[1]))
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *GrayA64f) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &GrayA64f{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(GrayA64f).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *GrayA64f) Opaque() bool {
	if p.M.Rect.Empty() {
		return true
	}
	i0, i1 := 4, p.M.Rect.Dx()*8
	for y := p.M.Rect.Min.Y; y < p.M.Rect.Max.Y; y++ {
		for i := i0; i < i1; i += 8 {
			if math.Float32frombits(binary.BigEndian.Uint32(p.M.Pix[i:])) < 0xFFFF {
				return false
			}
		}
		i0 += p.M.Stride
		i1 += p.M.Stride
	}
	return true
}

func (p *GrayA64f) CopyFrom(m image.Image) *GrayA64f {
	panic("TODO")
}

func (p *GrayA64f) Draw(dst Image, r image.Rectangle, src Image, sp image.Point) Image {
	panic("TODO")
}
