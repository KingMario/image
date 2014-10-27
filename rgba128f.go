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

type RGBA128f struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

func (p *RGBA128f) BaseType() image.Image { return p }
func (p *RGBA128f) Pix() []byte           { return p.M.Pix }
func (p *RGBA128f) Stride() int           { return p.M.Stride }
func (p *RGBA128f) Rect() image.Rectangle { return p.M.Rect }
func (p *RGBA128f) Channels() int         { return 4 }
func (p *RGBA128f) Depth() reflect.Kind   { return reflect.Float32 }

func (p *RGBA128f) ColorModel() color.Model { return color.RGBA64Model }

func (p *RGBA128f) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGBA128f) At(x, y int) color.Color {
	c := p.RGBA128fAt(x, y)
	return color.RGBA64{
		R: uint16(c[0]),
		G: uint16(c[1]),
		B: uint16(c[2]),
		A: uint16(c[3]),
	}
}

func (p *RGBA128f) RGBA128fAt(x, y int) [4]float32 {
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
func (p *RGBA128f) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*16
}

func (p *RGBA128f) Set(x, y int, c color.Color) {
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

func (p *RGBA128f) SetRGBA128f(x, y int, c [4]float32) {
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
func (p *RGBA128f) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGBA128f{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGBA128f{
		M: struct {
			Pix    []uint8
			Stride int
			Rect   image.Rectangle
		}{
			Pix:    p.M.Pix[i:],
			Stride: p.M.Stride,
			Rect:   r,
		},
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGBA128f) Opaque() bool {
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

// NewRGBA128f returns a new RGBA128f with the given bounds.
func NewRGBA128f(r image.Rectangle) *RGBA128f {
	return new(RGBA128f).Init(make([]uint8, 4*r.Dx()*r.Dy()), 4*r.Dx(), r)
}

func (p *RGBA128f) Init(pix []uint8, stride int, rect image.Rectangle) *RGBA128f {
	*p = RGBA128f{
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

func (p *RGBA128f) CopyFrom(m image.Image) *RGBA128f {
	panic("TODO")
}
