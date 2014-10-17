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
	_ Image = (*RGB)(nil)
	_ Image = (*RGB48)(nil)
	_ Image = (*RGB96f)(nil)
)

type RGB struct {
	m *image.RGBA
}

func (p *RGB) BaseType() image.Image { return p }
func (p *RGB) Pix() []byte           { return p.m.Pix }
func (p *RGB) Stride() int           { return p.m.Stride }
func (p *RGB) Rect() image.Rectangle { return p.m.Rect }
func (p *RGB) Channels() int         { return 1 }
func (p *RGB) Depth() reflect.Kind   { return reflect.Float32 }

func (p *RGB) ColorModel() color.Model { return color.RGBAModel }

func (p *RGB) Bounds() image.Rectangle { return p.Rect() }

func (p *RGB) At(x, y int) color.Color {
	return nil
}

func (p *RGB) ColorAt(x, y int) float32 {
	return 0
}

// PixOffset returns the index of the first element of _Pix that corresponds to
// the pixel at (x, y).
func (p *RGB) PixOffset(x, y int) int {
	return 0
}

func (p *RGB) Set(x, y int, c color.Color) {
	return
}

func (p *RGB) SetColor(x, y int, c float32) {
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGB) SubImage(r image.Rectangle) image.Image {
	return nil
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGB) Opaque() bool {
	return true
}

func NewRGB(r image.Rectangle) *RGB {
	return nil
}

type RGB48 struct {
	m *image.RGBA64
}

func (p *RGB48) BaseType() image.Image { return p }
func (p *RGB48) Pix() []byte           { return p.m.Pix }
func (p *RGB48) Stride() int           { return p.m.Stride }
func (p *RGB48) Rect() image.Rectangle { return p.m.Rect }
func (p *RGB48) Channels() int         { return 1 }
func (p *RGB48) Depth() reflect.Kind   { return reflect.Float32 }

func (p *RGB48) ColorModel() color.Model { return color.RGBAModel }

func (p *RGB48) Bounds() image.Rectangle { return p.Rect() }

func (p *RGB48) At(x, y int) color.Color {
	return nil
}

func (p *RGB48) ColorAt(x, y int) float32 {
	return 0
}

// PixOffset returns the index of the first element of _Pix that corresponds to
// the pixel at (x, y).
func (p *RGB48) PixOffset(x, y int) int {
	return 0
}

func (p *RGB48) Set(x, y int, c color.Color) {
	return
}

func (p *RGB48) SetColor(x, y int, c float32) {
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGB48) SubImage(r image.Rectangle) image.Image {
	return nil
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGB48) Opaque() bool {
	return true
}

func NewRGB48(r image.Rectangle) *RGB48 {
	return nil
}

type RGB96f struct {
	m *image.RGBA
}

func (p *RGB96f) BaseType() image.Image { return p }
func (p *RGB96f) Pix() []byte           { return p.m.Pix }
func (p *RGB96f) Stride() int           { return p.m.Stride }
func (p *RGB96f) Rect() image.Rectangle { return p.m.Rect }
func (p *RGB96f) Channels() int         { return 1 }
func (p *RGB96f) Depth() reflect.Kind   { return reflect.Float32 }

func (p *RGB96f) ColorModel() color.Model { return color.RGBAModel }

func (p *RGB96f) Bounds() image.Rectangle { return p.Rect() }

func (p *RGB96f) At(x, y int) color.Color {
	return nil
}

func (p *RGB96f) ColorAt(x, y int) float32 {
	return 0
}

// PixOffset returns the index of the first element of _Pix that corresponds to
// the pixel at (x, y).
func (p *RGB96f) PixOffset(x, y int) int {
	return 0
}

func (p *RGB96f) Set(x, y int, c color.Color) {
	return
}

func (p *RGB96f) SetColor(x, y int, c float32) {
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGB96f) SubImage(r image.Rectangle) image.Image {
	return nil
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGB96f) Opaque() bool {
	return true
}

func NewRGB96f(r image.Rectangle) *RGB96f {
	return nil
}
