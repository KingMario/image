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
	_ Image = (*GrayA)(nil)
	_ Image = (*GrayA16)(nil)
	_ Image = (*GrayA32f)(nil)
)

type GrayA struct {
	m *image.Gray
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
	return nil
}

func (p *GrayA) GrayAAt(x, y int) [2]uint8 {
	return [2]uint8{}
}

// PixOffset returns the index of the first element of _Pix that corresponds to
// the pixel at (x, y).
func (p *GrayA) PixOffset(x, y int) int {
	return 0
}

func (p *GrayA) Set(x, y int, c color.Color) {
	return
}

func (p *GrayA) SetGrayA(x, y int, c [2]uint8) {
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *GrayA) SubImage(r image.Rectangle) image.Image {
	return nil
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *GrayA) Opaque() bool {
	return true
}

func NewGrayA(r image.Rectangle) *GrayA {
	return nil
}

type GrayA16 struct {
	m *image.Gray16
}

func (p *GrayA16) BaseType() image.Image { return p }
func (p *GrayA16) Pix() []byte           { return p.m.Pix }
func (p *GrayA16) Stride() int           { return p.m.Stride }
func (p *GrayA16) Rect() image.Rectangle { return p.m.Rect }
func (p *GrayA16) Channels() int         { return 2 }
func (p *GrayA16) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *GrayA16) ColorModel() color.Model { return color.RGBAModel }

func (p *GrayA16) Bounds() image.Rectangle { return p.Rect() }

func (p *GrayA16) At(x, y int) color.Color {
	return nil
}

func (p *GrayA16) GrayA16At(x, y int) [2]uint16 {
	return [2]uint16{}
}

// PixOffset returns the index of the first element of _Pix that corresponds to
// the pixel at (x, y).
func (p *GrayA16) PixOffset(x, y int) int {
	return 0
}

func (p *GrayA16) Set(x, y int, c color.Color) {
	return
}

func (p *GrayA16) SetGrayA16(x, y int, c [2]uint16) {
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *GrayA16) SubImage(r image.Rectangle) image.Image {
	return nil
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *GrayA16) Opaque() bool {
	return true
}

func NewGrayA16(r image.Rectangle) *GrayA16 {
	return nil
}

type GrayA32f struct {
	m *image.Gray
}

func (p *GrayA32f) BaseType() image.Image { return p }
func (p *GrayA32f) Pix() []byte           { return p.m.Pix }
func (p *GrayA32f) Stride() int           { return p.m.Stride }
func (p *GrayA32f) Rect() image.Rectangle { return p.m.Rect }
func (p *GrayA32f) Channels() int         { return 2 }
func (p *GrayA32f) Depth() reflect.Kind   { return reflect.Float32 }

func (p *GrayA32f) ColorModel() color.Model { return color.RGBAModel }

func (p *GrayA32f) Bounds() image.Rectangle { return p.Rect() }

func (p *GrayA32f) At(x, y int) color.Color {
	return nil
}

func (p *GrayA32f) GrayA32fAt(x, y int) [2]float32 {
	return [2]float32{}
}

// PixOffset returns the index of the first element of _Pix that corresponds to
// the pixel at (x, y).
func (p *GrayA32f) PixOffset(x, y int) int {
	return 0
}

func (p *GrayA32f) Set(x, y int, c color.Color) {
	return
}

func (p *GrayA32f) SetGrayA32f(x, y int, c [2]float32) {
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *GrayA32f) SubImage(r image.Rectangle) image.Image {
	return nil
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *GrayA32f) Opaque() bool {
	return true
}

func NewGrayA32f(r image.Rectangle) *GrayA32f {
	return nil
}
