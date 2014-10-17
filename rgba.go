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

type RGBA64 struct {
	*image.RGBA64
}

func (p *RGBA64) BaseType() image.Image { return p.RGBA64 }
func (p *RGBA64) Pix() []byte           { return p.RGBA64.Pix }
func (p *RGBA64) Stride() int           { return p.RGBA64.Stride }
func (p *RGBA64) Rect() image.Rectangle { return p.RGBA64.Rect }
func (p *RGBA64) Channels() int         { return 4 }
func (p *RGBA64) Depth() reflect.Kind   { return reflect.Uint16 }

type RGBA128f struct {
	m *image.RGBA
}

func (p *RGBA128f) BaseType() image.Image { return p }
func (p *RGBA128f) Pix() []byte           { return p.m.Pix }
func (p *RGBA128f) Stride() int           { return p.m.Stride }
func (p *RGBA128f) Rect() image.Rectangle { return p.m.Rect }
func (p *RGBA128f) Channels() int         { return 4 }
func (p *RGBA128f) Depth() reflect.Kind   { return reflect.Float32 }

func (p *RGBA128f) ColorModel() color.Model { return color.RGBA64Model }

func (p *RGBA128f) Bounds() image.Rectangle { return p.Rect() }

func (p *RGBA128f) At(x, y int) color.Color {
	return nil
}

func (p *RGBA128f) RGBA128fAt(x, y int) [4]float32 {
	return [4]float32{}
}

// PixOffset returns the index of the first element of _Pix that corresponds to
// the pixel at (x, y).
func (p *RGBA128f) PixOffset(x, y int) int {
	return 0
}

func (p *RGBA128f) Set(x, y int, c color.Color) {
	return
}

func (p *RGBA128f) SetRGBA128f(x, y int, c [4]float32) {
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGBA128f) SubImage(r image.Rectangle) image.Image {
	return nil
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGBA128f) Opaque() bool {
	return true
}

func NewRGBA128f(r image.Rectangle) *RGBA128f {
	return nil
}
