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
	_ Image = (*Gray)(nil)
	_ Image = (*Gray16)(nil)
	_ Image = (*Gray32f)(nil)
)

type Gray struct {
	*image.Gray
}

func (p *Gray) BaseType() image.Image { return p.Gray }
func (p *Gray) Pix() []byte           { return p.Gray.Pix }
func (p *Gray) Stride() int           { return p.Gray.Stride }
func (p *Gray) Rect() image.Rectangle { return p.Gray.Rect }
func (p *Gray) Channels() int         { return 1 }
func (p *Gray) Depth() reflect.Kind   { return reflect.Uint8 }

type Gray16 struct {
	*image.Gray16
}

func (p *Gray16) BaseType() image.Image { return p.Gray16 }
func (p *Gray16) Pix() []byte           { return p.Gray16.Pix }
func (p *Gray16) Stride() int           { return p.Gray16.Stride }
func (p *Gray16) Rect() image.Rectangle { return p.Gray16.Rect }
func (p *Gray16) Channels() int         { return 1 }
func (p *Gray16) Depth() reflect.Kind   { return reflect.Uint16 }

type Gray32f struct {
	m image.Gray
}

func (p *Gray32f) BaseType() image.Image { return p }
func (p *Gray32f) Pix() []byte           { return p.m.Pix }
func (p *Gray32f) Stride() int           { return p.m.Stride }
func (p *Gray32f) Rect() image.Rectangle { return p.m.Rect }
func (p *Gray32f) Channels() int         { return 1 }
func (p *Gray32f) Depth() reflect.Kind   { return reflect.Float32 }

func (p *Gray32f) ColorModel() color.Model { return color.Gray16Model }

func (p *Gray32f) Bounds() image.Rectangle { return p.m.Rect }

func (p *Gray32f) At(x, y int) color.Color {
	return color.Gray16{
		Y: uint16(p.Gray32fAt(x, y)),
	}
}

func (p *Gray32f) Gray32fAt(x, y int) float32 {
	if !(image.Point{x, y}.In(p.m.Rect)) {
		return 0
	}
	i := p.PixOffset(x, y)
	v := math.Float32frombits(binary.BigEndian.Uint32(p.m.Pix[i:]))
	return v
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *Gray32f) PixOffset(x, y int) int {
	return (y-p.m.Rect.Min.Y)*p.m.Stride + (x-p.m.Rect.Min.X)*4
}

func (p *Gray32f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.m.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	v := float32(color.Gray16Model.Convert(c).(color.Gray16).Y)
	binary.BigEndian.PutUint32(p.m.Pix[i:], math.Float32bits(v))
	return
}

func (p *Gray32f) SetGray32f(x, y int, c float32) {
	if !(image.Point{x, y}.In(p.m.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	binary.BigEndian.PutUint32(p.m.Pix[i:], math.Float32bits(c))
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *Gray32f) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.m.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &Gray32f{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &Gray32f{
		m: image.Gray{
			Pix:    p.m.Pix[i:],
			Stride: p.m.Stride,
			Rect:   r,
		},
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *Gray32f) Opaque() bool {
	return true
}

// NewGray32f returns a new Gray32f with the given bounds.
func NewGray32f(r image.Rectangle) *Gray32f {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 4*w*h)
	return &Gray32f{
		m: image.Gray{
			Pix:    pix,
			Stride: 4 * w,
			Rect:   r,
		},
	}
}
