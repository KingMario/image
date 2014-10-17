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
	m *image.Gray
}

func (p *Gray32f) BaseType() image.Image { return p }
func (p *Gray32f) Pix() []byte           { return p.m.Pix }
func (p *Gray32f) Stride() int           { return p.m.Stride }
func (p *Gray32f) Rect() image.Rectangle { return p.m.Rect }
func (p *Gray32f) Channels() int         { return 1 }
func (p *Gray32f) Depth() reflect.Kind   { return reflect.Float32 }

func (p *Gray32f) ColorModel() color.Model { return color.RGBAModel }

func (p *Gray32f) Bounds() image.Rectangle { return p.Rect() }

func (p *Gray32f) At(x, y int) color.Color {
	return nil
}

func (p *Gray32f) ColorAt(x, y int) float32 {
	return 0
}

// PixOffset returns the index of the first element of _Pix that corresponds to
// the pixel at (x, y).
func (p *Gray32f) PixOffset(x, y int) int {
	return 0
}

func (p *Gray32f) Set(x, y int, c color.Color) {
	return
}

func (p *Gray32f) SetColor(x, y int, c float32) {
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *Gray32f) SubImage(r image.Rectangle) image.Image {
	return nil
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *Gray32f) Opaque() bool {
	return true
}

func NewGray32f(r image.Rectangle) *Gray32f {
	return nil
}
