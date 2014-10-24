// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
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

func (p *Unknown) ColorModel() color.Model {
	return nil
}

func (p *Unknown) Bounds() image.Rectangle {
	return image.Rectangle{}
}

func (p *Unknown) At(x, y int) color.Color {
	return nil
}

func (p *Unknown) PixelAt(x, y int) Pixel {
	return Pixel{}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *Unknown) PixOffset(x, y int) int {
	return 0
}

func (p *Unknown) Set(x, y int, c color.Color) {
	//
}

func (p *Unknown) SetPixel(x, y int, c Pixel) {
	//
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *Unknown) SubImage(r image.Rectangle) image.Image {
	return nil
}

// NewUnknown returns a new Unknown with the given bounds.
func NewUnknown(r image.Rectangle, channels int, depth reflect.Kind) (m *Unknown, err error) {
	return
}

func (p *Unknown) Init(pix []uint8, stride int, rect image.Rectangle) Image {
	p.M.Pix = pix
	p.M.Stride = stride
	p.M.Rect = rect
	return p
}
