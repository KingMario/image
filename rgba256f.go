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

type RGBA256f struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// NewRGBA256f returns a new RGBA256f with the given bounds.
func NewRGBA256f(r image.Rectangle) *RGBA256f {
	return new(RGBA256f).Init(make([]uint8, 4*r.Dx()*r.Dy()), 4*r.Dx(), r)
}

func (p *RGBA256f) Init(pix []uint8, stride int, rect image.Rectangle) *RGBA256f {
	*p = RGBA256f{
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

func (p *RGBA256f) BaseType() image.Image { return p }
func (p *RGBA256f) Pix() []byte           { return p.M.Pix }
func (p *RGBA256f) Stride() int           { return p.M.Stride }
func (p *RGBA256f) Rect() image.Rectangle { return p.M.Rect }
func (p *RGBA256f) Channels() int         { return 4 }
func (p *RGBA256f) Depth() reflect.Kind   { return reflect.Float32 }

func (p *RGBA256f) ColorModel() color.Model { return color.RGBA64Model }

func (p *RGBA256f) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGBA256f) At(x, y int) color.Color {
	c := p.RGBA256fAt(x, y)
	return color.RGBA64{
		R: uint16(c[0]),
		G: uint16(c[1]),
		B: uint16(c[2]),
		A: uint16(c[3]),
	}
}

func (p *RGBA256f) RGBA256fAt(x, y int) [4]float32 {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return [4]float32{}
	}
	i := p.PixOffset(x, y)
	bitsR := binary.BigEndian.Uint32(p.M.Pix[i+4*0:])
	bitsG := binary.BigEndian.Uint32(p.M.Pix[i+4*1:])
	bitsB := binary.BigEndian.Uint32(p.M.Pix[i+4*2:])
	bitsA := binary.BigEndian.Uint32(p.M.Pix[i+4*3:])
	r := math.Float32frombits(bitsR)
	g := math.Float32frombits(bitsG)
	b := math.Float32frombits(bitsB)
	a := math.Float32frombits(bitsA)
	return [4]float32{r, g, b, a}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGBA256f) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*16
}

func (p *RGBA256f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	v := color.RGBA64Model.Convert(c).(color.RGBA64)
	bitsR := math.Float32bits(float32(v.R))
	bitsG := math.Float32bits(float32(v.G))
	bitsB := math.Float32bits(float32(v.B))
	bitsA := math.Float32bits(float32(v.A))
	binary.BigEndian.PutUint32(p.M.Pix[i+4*0:], bitsR)
	binary.BigEndian.PutUint32(p.M.Pix[i+4*1:], bitsG)
	binary.BigEndian.PutUint32(p.M.Pix[i+4*2:], bitsB)
	binary.BigEndian.PutUint32(p.M.Pix[i+4*3:], bitsA)
	return
}

func (p *RGBA256f) SetRGBA256f(x, y int, c [4]float32) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	bitsR := math.Float32bits(c[0])
	bitsG := math.Float32bits(c[1])
	bitsB := math.Float32bits(c[2])
	bitsA := math.Float32bits(c[3])
	binary.BigEndian.PutUint32(p.M.Pix[i+4*0:], bitsR)
	binary.BigEndian.PutUint32(p.M.Pix[i+4*1:], bitsG)
	binary.BigEndian.PutUint32(p.M.Pix[i+4*2:], bitsB)
	binary.BigEndian.PutUint32(p.M.Pix[i+4*3:], bitsA)
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
	i0, i1 := 12, p.M.Rect.Dx()*16
	for y := p.M.Rect.Min.Y; y < p.M.Rect.Max.Y; y++ {
		for i := i0; i < i1; i += 16 {
			if math.Float32frombits(binary.BigEndian.Uint32(p.M.Pix[i:])) < 0xFFFF {
				return false
			}
		}
		i0 += p.M.Stride
		i1 += p.M.Stride
	}
	return true
}

func (p *RGBA256f) CopyFrom(m image.Image) *RGBA256f {
	panic("TODO")
}

func (p *RGBA256f) Draw(r image.Rectangle, src Image, sp image.Point) Image {
	panic("TODO")
}
