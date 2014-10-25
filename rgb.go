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
	_ Image = (*RGB)(nil)
	_ Image = (*RGB48)(nil)
	_ Image = (*RGB96f)(nil)
)

type RGB struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

func (p *RGB) BaseType() image.Image { return p }
func (p *RGB) Pix() []byte           { return p.M.Pix }
func (p *RGB) Stride() int           { return p.M.Stride }
func (p *RGB) Rect() image.Rectangle { return p.M.Rect }
func (p *RGB) Channels() int         { return 3 }
func (p *RGB) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *RGB) ColorModel() color.Model { return color.RGBAModel }

func (p *RGB) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGB) At(x, y int) color.Color {
	c := p.RGBAt(x, y)
	return color.RGBA{
		R: c[0],
		G: c[1],
		B: c[2],
		A: 0xFF,
	}
}

func (p *RGB) RGBAt(x, y int) [3]uint8 {
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
func (p *RGB) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*3
}

func (p *RGB) Set(x, y int, c color.Color) {
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

func (p *RGB) SetRGB(x, y int, c [3]uint8) {
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
func (p *RGB) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGB{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGB{
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
func (p *RGB) Opaque() bool {
	return true
}

// NewRGB returns a new RGB with the given bounds.
func NewRGB(r image.Rectangle) *RGB {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 3*w*h)
	return &RGB{
		M: struct {
			Pix    []uint8
			Stride int
			Rect   image.Rectangle
		}{
			Pix:    pix,
			Stride: 3 * w,
			Rect:   r,
		},
	}
}

func (p *RGB) Init(pix []uint8, stride int, rect image.Rectangle) Image {
	*p = RGB{
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

type RGB48 struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

func (p *RGB48) BaseType() image.Image { return p }
func (p *RGB48) Pix() []byte           { return p.M.Pix }
func (p *RGB48) Stride() int           { return p.M.Stride }
func (p *RGB48) Rect() image.Rectangle { return p.M.Rect }
func (p *RGB48) Channels() int         { return 3 }
func (p *RGB48) Depth() reflect.Kind   { return reflect.Uint16 }

func (p *RGB48) ColorModel() color.Model { return color.RGBA64Model }

func (p *RGB48) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGB48) At(x, y int) color.Color {
	c := p.RGB48At(x, y)
	return color.RGBA64{
		R: c[0],
		G: c[1],
		B: c[2],
		A: 0xFFFF,
	}
}

func (p *RGB48) RGB48At(x, y int) [3]uint16 {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return [3]uint16{}
	}
	i := p.PixOffset(x, y)
	return [3]uint16{
		uint16(p.M.Pix[i+0])<<8 | uint16(p.M.Pix[i+1]),
		uint16(p.M.Pix[i+2])<<8 | uint16(p.M.Pix[i+3]),
		uint16(p.M.Pix[i+4])<<8 | uint16(p.M.Pix[i+5]),
	}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGB48) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*6
}

func (p *RGB48) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := color.RGBA64Model.Convert(c).(color.RGBA64)
	p.M.Pix[i+0] = uint8(c1.R >> 8)
	p.M.Pix[i+1] = uint8(c1.R)
	p.M.Pix[i+2] = uint8(c1.G >> 8)
	p.M.Pix[i+3] = uint8(c1.G)
	p.M.Pix[i+4] = uint8(c1.B >> 8)
	p.M.Pix[i+5] = uint8(c1.B)
	return
}

func (p *RGB48) SetRGB48(x, y int, c [3]uint16) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.M.Pix[i+0] = uint8(c[0] >> 8)
	p.M.Pix[i+1] = uint8(c[0])
	p.M.Pix[i+2] = uint8(c[1] >> 8)
	p.M.Pix[i+3] = uint8(c[1])
	p.M.Pix[i+4] = uint8(c[2] >> 8)
	p.M.Pix[i+5] = uint8(c[2])
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGB48) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGB48{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGB48{
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
func (p *RGB48) Opaque() bool {
	return true
}

// NewRGB48 returns a new RGB48 with the given bounds.
func NewRGB48(r image.Rectangle) *RGB48 {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 6*w*h)
	return &RGB48{
		M: struct {
			Pix    []uint8
			Stride int
			Rect   image.Rectangle
		}{
			Pix:    pix,
			Stride: 6 * w,
			Rect:   r,
		},
	}
}

func (p *RGB48) Init(pix []uint8, stride int, rect image.Rectangle) Image {
	*p = RGB48{
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

type RGB96f struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

func (p *RGB96f) BaseType() image.Image { return p }
func (p *RGB96f) Pix() []byte           { return p.M.Pix }
func (p *RGB96f) Stride() int           { return p.M.Stride }
func (p *RGB96f) Rect() image.Rectangle { return p.M.Rect }
func (p *RGB96f) Channels() int         { return 3 }
func (p *RGB96f) Depth() reflect.Kind   { return reflect.Float32 }

func (p *RGB96f) ColorModel() color.Model { return color.RGBA64Model }

func (p *RGB96f) Bounds() image.Rectangle { return p.M.Rect }

func (p *RGB96f) At(x, y int) color.Color {
	c := p.RGB96fAt(x, y)
	return color.RGBA64{
		R: uint16(c[0]),
		G: uint16(c[1]),
		B: uint16(c[2]),
		A: 0xFFFF,
	}
}

func (p *RGB96f) RGB96fAt(x, y int) [3]float32 {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return [3]float32{}
	}
	i := p.PixOffset(x, y)
	bitsR := binary.BigEndian.Uint32(p.M.Pix[i+4*0:])
	bitsG := binary.BigEndian.Uint32(p.M.Pix[i+4*1:])
	bitsB := binary.BigEndian.Uint32(p.M.Pix[i+4*2:])
	r := math.Float32frombits(bitsR)
	g := math.Float32frombits(bitsG)
	b := math.Float32frombits(bitsB)
	return [3]float32{r, g, b}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGB96f) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*12
}

func (p *RGB96f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	v := color.RGBA64Model.Convert(c).(color.RGBA64)
	bitsR := math.Float32bits(float32(v.R))
	bitsG := math.Float32bits(float32(v.G))
	bitsB := math.Float32bits(float32(v.B))
	binary.BigEndian.PutUint32(p.M.Pix[i+4*0:], bitsR)
	binary.BigEndian.PutUint32(p.M.Pix[i+4*1:], bitsG)
	binary.BigEndian.PutUint32(p.M.Pix[i+4*2:], bitsB)
	return
}

func (p *RGB96f) SetRGB96f(x, y int, c [3]float32) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	bitsR := math.Float32bits(c[0])
	bitsG := math.Float32bits(c[1])
	bitsB := math.Float32bits(c[2])
	binary.BigEndian.PutUint32(p.M.Pix[i+4*0:], bitsR)
	binary.BigEndian.PutUint32(p.M.Pix[i+4*1:], bitsG)
	binary.BigEndian.PutUint32(p.M.Pix[i+4*2:], bitsB)
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGB96f) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGB96f{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGB96f{
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
func (p *RGB96f) Opaque() bool {
	return true
}

// NewRGB96f returns a new RGB96f with the given bounds.
func NewRGB96f(r image.Rectangle) *RGB96f {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 12*w*h)
	return &RGB96f{
		M: struct {
			Pix    []uint8
			Stride int
			Rect   image.Rectangle
		}{
			Pix:    pix,
			Stride: 12 * w,
			Rect:   r,
		},
	}
}

func (p *RGB96f) Init(pix []uint8, stride int, rect image.Rectangle) Image {
	*p = RGB96f{
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
