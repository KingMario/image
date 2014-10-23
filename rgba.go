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

var (
	_ Image = (*RGBA)(nil)
	_ Image = (*RGBA64)(nil)
	_ Image = (*RGBA128f)(nil)
)

type RGBA struct {
	*image.RGBA
}

func (p *RGBA) BaseType() image.Image { return p.RGBA }
func (p *RGBA) Pix() []byte           { return p.RGBA.Pix }
func (p *RGBA) Stride() int           { return p.RGBA.Stride }
func (p *RGBA) Rect() image.Rectangle { return p.RGBA.Rect }
func (p *RGBA) Channels() int         { return 4 }
func (p *RGBA) Depth() reflect.Kind   { return reflect.Uint8 }

// NewRGBA returns a new RGBA with the given bounds.
func NewRGBA(r image.Rectangle) *RGBA {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 4*w*h)
	return &RGBA{
		RGBA: &image.RGBA{
			Pix:    pix,
			Stride: 4 * w,
			Rect:   r,
		},
	}
}

type RGBA64 struct {
	*image.RGBA64
}

func (p *RGBA64) BaseType() image.Image { return p.RGBA64 }
func (p *RGBA64) Pix() []byte           { return p.RGBA64.Pix }
func (p *RGBA64) Stride() int           { return p.RGBA64.Stride }
func (p *RGBA64) Rect() image.Rectangle { return p.RGBA64.Rect }
func (p *RGBA64) Channels() int         { return 4 }
func (p *RGBA64) Depth() reflect.Kind   { return reflect.Uint16 }

// NewRGBA64 returns a new RGBA64 with the given bounds.
func NewRGBA64(r image.Rectangle) *RGBA64 {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 8*w*h)
	return &RGBA64{
		RGBA64: &image.RGBA64{
			Pix:    pix,
			Stride: 8 * w,
			Rect:   r,
		},
	}
}

type RGBA128f struct {
	m image.RGBA
}

func (p *RGBA128f) BaseType() image.Image { return p }
func (p *RGBA128f) Pix() []byte           { return p.m.Pix }
func (p *RGBA128f) Stride() int           { return p.m.Stride }
func (p *RGBA128f) Rect() image.Rectangle { return p.m.Rect }
func (p *RGBA128f) Channels() int         { return 4 }
func (p *RGBA128f) Depth() reflect.Kind   { return reflect.Float32 }

func (p *RGBA128f) ColorModel() color.Model { return color.RGBA64Model }

func (p *RGBA128f) Bounds() image.Rectangle { return p.m.Rect }

func (p *RGBA128f) At(x, y int) color.Color {
	c := p.RGBA128fAt(x, y)
	rr, gg, bb, aa := colorRGBA128f(c).RGBA()
	return color.RGBA64{
		R: uint16(rr),
		G: uint16(gg),
		B: uint16(bb),
		A: uint16(aa),
	}
}

func (p *RGBA128f) RGBA128fAt(x, y int) [4]float32 {
	if !(image.Point{x, y}.In(p.m.Rect)) {
		return [4]float32{}
	}
	i := p.PixOffset(x, y)
	bitsR := binary.BigEndian.Uint32(p.m.Pix[i+4*0:])
	bitsG := binary.BigEndian.Uint32(p.m.Pix[i+4*1:])
	bitsB := binary.BigEndian.Uint32(p.m.Pix[i+4*2:])
	bitsA := binary.BigEndian.Uint32(p.m.Pix[i+4*3:])
	r := math.Float32frombits(bitsR)
	g := math.Float32frombits(bitsG)
	b := math.Float32frombits(bitsB)
	a := math.Float32frombits(bitsA)
	return [4]float32{r, g, b, a}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGBA128f) PixOffset(x, y int) int {
	return (y-p.m.Rect.Min.Y)*p.m.Stride + (x-p.m.Rect.Min.X)*16
}

func (p *RGBA128f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.m.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	v := color.RGBA64Model.Convert(c).(color.RGBA64)
	bitsR := math.Float32bits(float32(v.R))
	bitsG := math.Float32bits(float32(v.G))
	bitsB := math.Float32bits(float32(v.B))
	bitsA := math.Float32bits(float32(v.A))
	binary.BigEndian.PutUint32(p.m.Pix[i+4*0:], bitsR)
	binary.BigEndian.PutUint32(p.m.Pix[i+4*1:], bitsG)
	binary.BigEndian.PutUint32(p.m.Pix[i+4*2:], bitsB)
	binary.BigEndian.PutUint32(p.m.Pix[i+4*3:], bitsA)
	return
}

func (p *RGBA128f) SetRGBA128f(x, y int, c [4]float32) {
	if !(image.Point{x, y}.In(p.m.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	bitsR := math.Float32bits(c[0])
	bitsG := math.Float32bits(c[1])
	bitsB := math.Float32bits(c[2])
	bitsA := math.Float32bits(c[3])
	binary.BigEndian.PutUint32(p.m.Pix[i+4*0:], bitsR)
	binary.BigEndian.PutUint32(p.m.Pix[i+4*1:], bitsG)
	binary.BigEndian.PutUint32(p.m.Pix[i+4*2:], bitsB)
	binary.BigEndian.PutUint32(p.m.Pix[i+4*3:], bitsA)
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGBA128f) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.m.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGBA128f{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGBA128f{
		m: image.RGBA{
			Pix:    p.m.Pix[i:],
			Stride: p.m.Stride,
			Rect:   r,
		},
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGBA128f) Opaque() bool {
	if p.m.Rect.Empty() {
		return true
	}
	i0, i1 := 12, p.m.Rect.Dx()*16
	for y := p.m.Rect.Min.Y; y < p.m.Rect.Max.Y; y++ {
		for i := i0; i < i1; i += 16 {
			if math.Float32frombits(binary.BigEndian.Uint32(p.m.Pix[i:])) < 0xFFFF {
				return false
			}
		}
		i0 += p.m.Stride
		i1 += p.m.Stride
	}
	return true
}

// NewRGBA128f returns a new RGBA128f with the given bounds.
func NewRGBA128f(r image.Rectangle) *RGBA128f {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 16*w*h)
	return &RGBA128f{
		m: image.RGBA{
			Pix:    pix,
			Stride: 16 * w,
			Rect:   r,
		},
	}
}
