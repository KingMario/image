// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
	"image/color"
	"image/draw"
	"reflect"
)

type Image interface {
	// Get original type, such as *image.Gray, *image.RGBA, etc.
	BaseType() image.Image

	// Pix holds the image's pixels, as pixel values in big-endian format. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*Channels*sizeof(DataType)].
	Pix() []byte
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride() int
	// Rect is the image's bounds.
	Rect() image.Rectangle

	// 1=Gray, 2=GrayA, 3=RGB, 4=RGBA
	Channels() int
	// Uint8/Uint16/Float32/...
	Depth() reflect.Kind

	draw.Image
}

func AsImage(x image.Image) (m Image) {
	if p, ok := x.(Image); ok {
		return p
	}

	switch x := x.(type) {
	case *image.Gray:
		return &Gray{x}
	case *image.Gray16:
		return &Gray16{x}
	case *image.RGBA:
		return &RGBA{x}
	case *image.RGBA64:
		return &RGBA64{x}
	}

	b := x.Bounds()
	rgba64 := NewRGBA64(b)
	dstColorRGBA64 := &color.RGBA64{}
	dstColor := color.Color(dstColorRGBA64)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			pr, pg, pb, pa := m.At(x, y).RGBA()
			dstColorRGBA64.R = uint16(pr)
			dstColorRGBA64.G = uint16(pg)
			dstColorRGBA64.B = uint16(pb)
			dstColorRGBA64.A = uint16(pa)
			rgba64.Set(x, y, dstColor)
		}
	}

	return rgba64
}

func CopyImage(x image.Image) (m Image) {
	panic("TODO")
}

func ConvertImage(x image.Image, channels int, dataType reflect.Kind) (m Image) {
	panic("TODO")
}

func CopyConvertImage(x image.Image, channels int, dataType reflect.Kind) (m Image) {
	panic("TODO")
}

func NewImage(r image.Rectangle, channels int, dataType reflect.Kind) (m Image, err error) {
	switch {
	case channels == 1 && dataType == reflect.Uint8:
		m = NewGray(r)
		return
	case channels == 1 && dataType == reflect.Uint16:
		m = NewGray16(r)
		return
	case channels == 1 && dataType == reflect.Float32:
		m = NewGray32f(r)
		return

	case channels == 2 && dataType == reflect.Uint8:
		m = NewGrayA(r)
		return
	case channels == 2 && dataType == reflect.Uint16:
		m = NewGrayA32(r)
		return
	case channels == 2 && dataType == reflect.Float32:
		m = NewGrayA64f(r)
		return

	case channels == 3 && dataType == reflect.Uint8:
		m = NewRGB(r)
		return
	case channels == 3 && dataType == reflect.Uint16:
		m = NewRGB48(r)
		return
	case channels == 3 && dataType == reflect.Float32:
		m = NewRGB96f(r)
		return

	case channels == 4 && dataType == reflect.Uint8:
		m = NewRGBA(r)
		return
	case channels == 4 && dataType == reflect.Uint16:
		m = NewRGBA64(r)
		return
	case channels == 4 && dataType == reflect.Float32:
		m = NewRGBA128f(r)
		return

	default:
		m, err = NewUnknown(r, channels, dataType)
		return
	}
}
