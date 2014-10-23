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

func (p *Unknown) BaseType() image.Image { return p }
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

func (p *Unknown) PixelAt(x, y int) interface{} {
	return nil
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *Unknown) PixOffset(x, y int) int {
	return 0
}

func (p *Unknown) Set(x, y int, c color.Color) {
	//
}

func (p *Unknown) SetPixel(x, y int, c interface{}) {
	//
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *Unknown) SubImage(r image.Rectangle) image.Image {
	return nil
}

// NewUnknown returns a new Unknown with the given bounds.
func NewUnknown(r image.Rectangle) *Unknown {
	return nil
}
