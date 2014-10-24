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
	_ Image = (*GrayA)(nil)
	_ Image = (*GrayA32)(nil)
	_ Image = (*GrayA64f)(nil)
)

type GrayA struct {
	m image.Gray
}

func (p *GrayA) BaseType() image.Image { return p }
func (p *GrayA) Pix() []byte           { return p.m.Pix }
func (p *GrayA) Stride() int           { return p.m.Stride }
func (p *GrayA) Rect() image.Rectangle { return p.m.Rect }
func (p *GrayA) Channels() int         { return 2 }
func (p *GrayA) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *GrayA) ColorModel() color.Model { return color.RGBAModel }

func (p *GrayA) Bounds() image.Rectangle { return p.Rect() }

func (p *GrayA) At(x, y int) color.Color {
	c := p.GrayAAt(x, y)
	return color.RGBA{
		R: c[0],
		G: c[0],
		B: c[0],
		A: c[1],
	}
}

func (p *GrayA) GrayAAt(x, y int) [2]uint8 {
	if !(image.Point{x, y}.In(p.m.Rect)) {
		return [2]uint8{}
	}
	i := p.PixOffset(x, y)
	return [2]uint8{
		p.m.Pix[i+0],
		p.m.Pix[i+1],
	}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *GrayA) PixOffset(x, y int) int {
	return (y-p.m.Rect.Min.Y)*p.m.Stride + (x-p.m.Rect.Min.X)*2
}

func (p *GrayA) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.m.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	rr, gg, bb, aa := c.RGBA()
	yy := colorRgbToGray(rr, gg, bb)
	p.m.Pix[i+0] = uint8(yy >> 8)
	p.m.Pix[i+1] = uint8(aa >> 8)
	return
}

func (p *GrayA) SetGrayA(x, y int, c [2]uint8) {
	if !(image.Point{x, y}.In(p.m.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.m.Pix[i+0] = c[0]
	p.m.Pix[i+1] = c[1]
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *GrayA) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.m.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &GrayA{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &GrayA{
		m: image.Gray{
			Pix:    p.m.Pix[i:],
			Stride: p.m.Stride,
			Rect:   r,
		},
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *GrayA) Opaque() bool {
	if p.m.Rect.Empty() {
		return true
	}
	i0, i1 := 1, p.m.Rect.Dx()*2
	for y := p.m.Rect.Min.Y; y < p.m.Rect.Max.Y; y++ {
		for i := i0; i < i1; i += 2 {
			if p.m.Pix[i] != 0xFF {
				return false
			}
		}
		i0 += p.m.Stride
		i1 += p.m.Stride
	}
	return true
}

// NewGrayA returns a new GrayA with the given bounds.
func NewGrayA(r image.Rectangle) *GrayA {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 2*w*h)
	return &GrayA{
		m: image.Gray{
			Pix:    pix,
			Stride: 2 * w,
			Rect:   r,
		},
	}
}

func (p *GrayA) Init(pix []uint8, stride int, rect image.Rectangle) Image {
	*p = GrayA{
		m: image.Gray{
			Pix:    pix,
			Stride: stride,
			Rect:   rect,
		},
	}
	return p
}

type GrayA32 struct {
	m image.Gray16
}

func (p *GrayA32) BaseType() image.Image { return p }
func (p *GrayA32) Pix() []byte           { return p.m.Pix }
func (p *GrayA32) Stride() int           { return p.m.Stride }
func (p *GrayA32) Rect() image.Rectangle { return p.m.Rect }
func (p *GrayA32) Channels() int         { return 2 }
func (p *GrayA32) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *GrayA32) ColorModel() color.Model { return color.RGBA64Model }

func (p *GrayA32) Bounds() image.Rectangle { return p.Rect() }

func (p *GrayA32) At(x, y int) color.Color {
	c := p.GrayA32At(x, y)
	rr, gg, bb, aa := colorGrayA32(c).RGBA()
	return color.RGBA64{
		R: uint16(rr),
		G: uint16(gg),
		B: uint16(bb),
		A: uint16(aa),
	}
}

func (p *GrayA32) GrayA32At(x, y int) [2]uint16 {
	if !(image.Point{x, y}.In(p.m.Rect)) {
		return [2]uint16{}
	}
	i := p.PixOffset(x, y)
	return [2]uint16{
		uint16(p.m.Pix[i+0])<<8 | uint16(p.m.Pix[i+1]),
		0xFFFF,
	}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *GrayA32) PixOffset(x, y int) int {
	return (y-p.m.Rect.Min.Y)*p.m.Stride + (x-p.m.Rect.Min.X)*4
}

func (p *GrayA32) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.m.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	yy := color.Gray16Model.Convert(c).(color.Gray16).Y
	_, _, _, aa := c.RGBA()
	p.m.Pix[i+0] = uint8(yy >> 8)
	p.m.Pix[i+1] = uint8(yy)
	p.m.Pix[i+2] = uint8(aa >> 8)
	p.m.Pix[i+3] = uint8(aa)
	return
}

func (p *GrayA32) SetGrayA32(x, y int, c [2]uint16) {
	if !(image.Point{x, y}.In(p.m.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.m.Pix[i+0] = uint8(c[0] >> 8)
	p.m.Pix[i+1] = uint8(c[0])
	p.m.Pix[i+2] = uint8(c[1] >> 8)
	p.m.Pix[i+3] = uint8(c[1])
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *GrayA32) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.m.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &GrayA32{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &GrayA32{
		m: image.Gray16{
			Pix:    p.m.Pix[i:],
			Stride: p.m.Stride,
			Rect:   r,
		},
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *GrayA32) Opaque() bool {
	if p.m.Rect.Empty() {
		return true
	}
	i0, i1 := 2, p.m.Rect.Dx()*4
	for y := p.m.Rect.Min.Y; y < p.m.Rect.Max.Y; y++ {
		for i := i0; i < i1; i += 4 {
			if p.m.Pix[i+0] != 0xff || p.m.Pix[i+1] != 0xff {
				return false
			}
		}
		i0 += p.m.Stride
		i1 += p.m.Stride
	}
	return true
}

// NewGrayA32 returns a new GrayA32 with the given bounds.
func NewGrayA32(r image.Rectangle) *GrayA32 {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 4*w*h)
	return &GrayA32{
		m: image.Gray16{
			Pix:    pix,
			Stride: 4 * w,
			Rect:   r,
		},
	}
}

func (p *GrayA32) Init(pix []uint8, stride int, rect image.Rectangle) Image {
	*p = GrayA32{
		m: image.Gray16{
			Pix:    pix,
			Stride: stride,
			Rect:   rect,
		},
	}
	return p
}

type GrayA64f struct {
	m image.Gray
}

func (p *GrayA64f) BaseType() image.Image { return p }
func (p *GrayA64f) Pix() []byte           { return p.m.Pix }
func (p *GrayA64f) Stride() int           { return p.m.Stride }
func (p *GrayA64f) Rect() image.Rectangle { return p.m.Rect }
func (p *GrayA64f) Channels() int         { return 2 }
func (p *GrayA64f) Depth() reflect.Kind   { return reflect.Float32 }

func (p *GrayA64f) ColorModel() color.Model { return color.RGBA64Model }

func (p *GrayA64f) Bounds() image.Rectangle { return p.Rect() }

func (p *GrayA64f) At(x, y int) color.Color {
	c := p.GrayA64fAt(x, y)
	rr, gg, bb, aa := colorGrayA64f(c).RGBA()
	return color.RGBA64{
		R: uint16(rr),
		G: uint16(gg),
		B: uint16(bb),
		A: uint16(aa),
	}
}

func (p *GrayA64f) GrayA64fAt(x, y int) [2]float32 {
	if !(image.Point{x, y}.In(p.m.Rect)) {
		return [2]float32{}
	}
	i := p.PixOffset(x, y)
	v0 := math.Float32frombits(binary.BigEndian.Uint32(p.m.Pix[i+4*0:]))
	v1 := math.Float32frombits(binary.BigEndian.Uint32(p.m.Pix[i+4*1:]))
	return [2]float32{v0, v1}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *GrayA64f) PixOffset(x, y int) int {
	return (y-p.m.Rect.Min.Y)*p.m.Stride + (x-p.m.Rect.Min.X)*8
}

func (p *GrayA64f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.m.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	yy := color.Gray16Model.Convert(c).(color.Gray16).Y
	_, _, _, aa := c.RGBA()
	binary.BigEndian.PutUint32(p.m.Pix[i+4*0:], math.Float32bits(float32(yy)))
	binary.BigEndian.PutUint32(p.m.Pix[i+4*1:], math.Float32bits(float32(aa)))
	return
}

func (p *GrayA64f) SetGrayA64f(x, y int, c [2]float32) {
	if !(image.Point{x, y}.In(p.m.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	binary.BigEndian.PutUint32(p.m.Pix[i+4*0:], math.Float32bits(c[0]))
	binary.BigEndian.PutUint32(p.m.Pix[i+4*1:], math.Float32bits(c[1]))
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *GrayA64f) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.m.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &GrayA64f{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &GrayA64f{
		m: image.Gray{
			Pix:    p.m.Pix[i:],
			Stride: p.m.Stride,
			Rect:   r,
		},
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *GrayA64f) Opaque() bool {
	if p.m.Rect.Empty() {
		return true
	}
	i0, i1 := 4, p.m.Rect.Dx()*8
	for y := p.m.Rect.Min.Y; y < p.m.Rect.Max.Y; y++ {
		for i := i0; i < i1; i += 8 {
			if math.Float32frombits(binary.BigEndian.Uint32(p.m.Pix[i:])) < 0xFFFF {
				return false
			}
		}
		i0 += p.m.Stride
		i1 += p.m.Stride
	}
	return true
}

// NewGrayA64f returns a new GrayA64f with the given bounds.
func NewGrayA64f(r image.Rectangle) *GrayA64f {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 8*w*h)
	return &GrayA64f{
		m: image.Gray{
			Pix:    pix,
			Stride: 8 * w,
			Rect:   r,
		},
	}
}

func (p *GrayA64f) Init(pix []uint8, stride int, rect image.Rectangle) Image {
	*p = GrayA64f{
		m: image.Gray{
			Pix:    pix,
			Stride: stride,
			Rect:   rect,
		},
	}
	return p
}
