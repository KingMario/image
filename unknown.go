// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"fmt"
	"image"
	"image/color"
	"reflect"
)

var (
	_ Image = (*Unknown)(nil)
)

type Unknown struct {
	M struct {
		Pix      []uint8
		Stride   int
		Rect     image.Rectangle
		Channels int
		Depth    reflect.Kind
	}
}

func (p *Unknown) BaseType() image.Image {
	switch channels, depth := p.M.Channels, p.M.Depth; {
	case channels == 1 && depth == reflect.Uint8:
		return &image.Gray{
			Pix:    p.M.Pix,
			Stride: p.M.Stride,
			Rect:   p.M.Rect,
		}
	case channels == 1 && depth == reflect.Uint16:
		return &image.Gray16{
			Pix:    p.M.Pix,
			Stride: p.M.Stride,
			Rect:   p.M.Rect,
		}
	case channels == 1 && depth == reflect.Float32:
		return new(Gray32f).Init(p.M.Pix, p.M.Stride, p.M.Rect)

	case channels == 2 && depth == reflect.Uint8:
		return new(GrayA).Init(p.M.Pix, p.M.Stride, p.M.Rect)
	case channels == 2 && depth == reflect.Uint16:
		return new(GrayA32).Init(p.M.Pix, p.M.Stride, p.M.Rect)
	case channels == 2 && depth == reflect.Float32:
		return new(GrayA64f).Init(p.M.Pix, p.M.Stride, p.M.Rect)

	case channels == 3 && depth == reflect.Uint8:
		return new(RGB).Init(p.M.Pix, p.M.Stride, p.M.Rect)
	case channels == 3 && depth == reflect.Uint16:
		return new(RGB48).Init(p.M.Pix, p.M.Stride, p.M.Rect)
	case channels == 3 && depth == reflect.Float32:
		return new(RGB96f).Init(p.M.Pix, p.M.Stride, p.M.Rect)

	case channels == 4 && depth == reflect.Uint8:
		return &image.RGBA{
			Pix:    p.M.Pix,
			Stride: p.M.Stride,
			Rect:   p.M.Rect,
		}
	case channels == 4 && depth == reflect.Uint16:
		return &image.RGBA64{
			Pix:    p.M.Pix,
			Stride: p.M.Stride,
			Rect:   p.M.Rect,
		}
	case channels == 4 && depth == reflect.Float32:
		return new(RGBA128f).Init(p.M.Pix, p.M.Stride, p.M.Rect)
	}

	return p
}

func (p *Unknown) Pix() []byte           { return p.M.Pix }
func (p *Unknown) Stride() int           { return p.M.Stride }
func (p *Unknown) Rect() image.Rectangle { return p.M.Rect }
func (p *Unknown) Channels() int         { return p.M.Channels }
func (p *Unknown) Depth() reflect.Kind   { return p.M.Depth }

func (p *Unknown) ColorModel() color.Model { return color.RGBA64Model }

func (p *Unknown) Bounds() image.Rectangle { return p.M.Rect }

func (p *Unknown) At(x, y int) color.Color {
	r, g, b, a := p.PixelAt(x, y).RGBA()
	return color.RGBA64{
		R: uint16(r),
		G: uint16(g),
		B: uint16(b),
		A: uint16(a),
	}
}

func (p *Unknown) PixelAt(x, y int) Pixel {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return Pixel{}
	}
	i := p.PixOffset(x, y)
	return Pixel{
		Channels: p.M.Channels,
		Depth:    p.M.Depth,
		Value:    p.M.Pix[i:][:depthType(p.M.Depth).ByteSize()*p.M.Channels],
	}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *Unknown) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*depthType(p.M.Depth).ByteSize()*p.M.Channels
}

func (p *Unknown) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	var c1 Pixel
	c1.PutRGBA(c.RGBA())
	copy(p.M.Pix[i:], c1.Value)
}

func (p *Unknown) SetPixel(x, y int, c Pixel) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	switch {
	case c.Channels == p.M.Channels && c.Depth == p.M.Depth:
		copy(p.M.Pix[i:], c.Value)
	case c.Channels == p.M.Channels:
		var c1 Pixel
		c1.PutValueN(c.ValueN())
		copy(p.M.Pix[i:], c1.Value)
	default:
		var c1 Pixel
		c1.PutRGBA(c.RGBA())
		copy(p.M.Pix[i:], c1.Value)
	}
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *Unknown) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &Unknown{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &Unknown{
		M: struct {
			Pix      []uint8
			Stride   int
			Rect     image.Rectangle
			Channels int
			Depth    reflect.Kind
		}{
			Pix:      p.M.Pix[i:],
			Stride:   p.M.Stride,
			Rect:     r,
			Channels: p.M.Channels,
			Depth:    p.M.Depth,
		},
	}
}

// NewUnknown returns a new Unknown with the given bounds.
func NewUnknown(r image.Rectangle, channels int, depth reflect.Kind) (m *Unknown, err error) {
	if channels <= 0 || !depthType(depth).IsValid() {
		err = fmt.Errorf("image: NewUnknown, invalid format: channels = %d, depth = %v", channels, depth)
		return
	}
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, depthType(depth).ByteSize()*channels*w*h)
	m = &Unknown{
		M: struct {
			Pix      []uint8
			Stride   int
			Rect     image.Rectangle
			Channels int
			Depth    reflect.Kind
		}{
			Pix:      pix,
			Stride:   depthType(depth).ByteSize() * channels * w,
			Rect:     r,
			Channels: channels,
			Depth:    depth,
		},
	}
	return
}

func (p *Unknown) Init(pix []uint8, stride int, rect image.Rectangle, channels int, depth reflect.Kind) Image {
	p.M.Pix = pix
	p.M.Stride = stride
	p.M.Rect = rect
	p.M.Channels = channels
	p.M.Depth = depth
	return p
}
