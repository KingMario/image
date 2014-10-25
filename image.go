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

	// if Depth() != Invalid { 1:Gray, 2:GrayA, 3:RGB, 4:RGBA, N:[N]Type }
	// if Depth() == Invalid { N:[N]byte }
	Channels() int
	// Invalid/Uint8/Uint16/Uint32/Uint64/Int8/Int16/Int32/Int64/Float32/Float64
	// Invalid is equal Byte type.
	Depth() reflect.Kind

	draw.Image
}

func newRGBA64FromImage(x image.Image) (m *RGBA64) {
	b := x.Bounds()
	rgba64 := NewRGBA64(b)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			pr, pg, pb, pa := m.At(x, y).RGBA()
			rgba64.SetRGBA64(x, y, color.RGBA64{
				R: uint16(pr),
				G: uint16(pg),
				B: uint16(pb),
				A: uint16(pa),
			})
		}
	}
	return rgba64
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

	return newRGBA64FromImage(m)
}

func CopyImage(x image.Image) (m Image) {
	if x, ok := x.(Image); ok {
		switch channels, depth := x.Channels(), x.Depth(); {
		case channels == 1 && depth == reflect.Uint8:
			return new(Gray).Init(append([]uint8(nil), x.Pix()...), x.Stride(), x.Rect())
		case channels == 1 && depth == reflect.Uint16:
			return new(Gray16).Init(append([]uint8(nil), x.Pix()...), x.Stride(), x.Rect())
		case channels == 1 && depth == reflect.Float32:
			return new(Gray32f).Init(append([]uint8(nil), x.Pix()...), x.Stride(), x.Rect())
		case channels == 2 && depth == reflect.Uint8:
			return new(GrayA).Init(append([]uint8(nil), x.Pix()...), x.Stride(), x.Rect())
		case channels == 2 && depth == reflect.Uint16:
			return new(GrayA32).Init(append([]uint8(nil), x.Pix()...), x.Stride(), x.Rect())
		case channels == 2 && depth == reflect.Float32:
			return new(GrayA64f).Init(append([]uint8(nil), x.Pix()...), x.Stride(), x.Rect())
		case channels == 3 && depth == reflect.Uint8:
			return new(RGB).Init(append([]uint8(nil), x.Pix()...), x.Stride(), x.Rect())
		case channels == 3 && depth == reflect.Uint16:
			return new(RGB48).Init(append([]uint8(nil), x.Pix()...), x.Stride(), x.Rect())
		case channels == 3 && depth == reflect.Float32:
			return new(RGB96f).Init(append([]uint8(nil), x.Pix()...), x.Stride(), x.Rect())
		case channels == 4 && depth == reflect.Uint8:
			return new(RGBA).Init(append([]uint8(nil), x.Pix()...), x.Stride(), x.Rect())
		case channels == 4 && depth == reflect.Uint16:
			return new(RGBA64).Init(append([]uint8(nil), x.Pix()...), x.Stride(), x.Rect())
		case channels == 4 && depth == reflect.Float32:
			return new(RGBA128f).Init(append([]uint8(nil), x.Pix()...), x.Stride(), x.Rect())
		}

		return new(Unknown).Init(
			append([]uint8(nil), x.Pix()...), x.Stride(), x.Rect(),
			x.Channels(), x.Depth(),
		)
	}

	switch x := x.(type) {
	case *image.Gray:
		return new(Gray).Init(append([]uint8(nil), x.Pix...), x.Stride, x.Rect)
	case *image.Gray16:
		return new(Gray16).Init(append([]uint8(nil), x.Pix...), x.Stride, x.Rect)
	case *image.RGBA:
		return new(RGBA).Init(append([]uint8(nil), x.Pix...), x.Stride, x.Rect)
	case *image.RGBA64:
		return new(RGBA64).Init(append([]uint8(nil), x.Pix...), x.Stride, x.Rect)
	}

	return newRGBA64FromImage(m)
}

func ConvertImage(x image.Image, channels int, depth reflect.Kind) (m Image) {
	if x, ok := x.(Image); ok {
		if x.Channels() == channels && x.Depth() == depth {
			// speed up
		}
	}

	switch x.(type) {
	case *image.Gray:
		if channels == 1 && depth == reflect.Uint8 {
			// speed up
		}
	case *image.Gray16:
		if channels == 1 && depth == reflect.Uint16 {
			// speed up
		}
	case *image.RGBA:
		if channels == 4 && depth == reflect.Uint8 {
			// speed up
		}
	case *image.RGBA64:
		if channels == 4 && depth == reflect.Uint16 {
			// speed up
		}
	}

	panic("TODO")
}

func CopyConvertImage(x image.Image, channels int, depth reflect.Kind) (m Image) {
	if x, ok := x.(Image); ok {
		if x.Channels() == channels && x.Depth() == depth {
			// speed up
		}
	}

	switch x.(type) {
	case *image.Gray:
		if channels == 1 && depth == reflect.Uint8 {
			// speed up
		}
	case *image.Gray16:
		if channels == 1 && depth == reflect.Uint16 {
			// speed up
		}
	case *image.RGBA:
		if channels == 4 && depth == reflect.Uint8 {
			// speed up
		}
	case *image.RGBA64:
		if channels == 4 && depth == reflect.Uint16 {
			// speed up
		}
	}

	panic("TODO")
}

func NewImage(r image.Rectangle, channels int, depth reflect.Kind) (m Image, err error) {
	switch {
	case channels == 1 && depth == reflect.Uint8:
		m = NewGray(r)
		return
	case channels == 1 && depth == reflect.Uint16:
		m = NewGray16(r)
		return
	case channels == 1 && depth == reflect.Float32:
		m = NewGray32f(r)
		return

	case channels == 2 && depth == reflect.Uint8:
		m = NewGrayA(r)
		return
	case channels == 2 && depth == reflect.Uint16:
		m = NewGrayA32(r)
		return
	case channels == 2 && depth == reflect.Float32:
		m = NewGrayA64f(r)
		return

	case channels == 3 && depth == reflect.Uint8:
		m = NewRGB(r)
		return
	case channels == 3 && depth == reflect.Uint16:
		m = NewRGB48(r)
		return
	case channels == 3 && depth == reflect.Float32:
		m = NewRGB96f(r)
		return

	case channels == 4 && depth == reflect.Uint8:
		m = NewRGBA(r)
		return
	case channels == 4 && depth == reflect.Uint16:
		m = NewRGBA64(r)
		return
	case channels == 4 && depth == reflect.Float32:
		m = NewRGBA128f(r)
		return

	default:
		m, err = NewUnknown(r, channels, depth)
		return
	}
}
