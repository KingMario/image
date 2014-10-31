// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The byte order fallacy. By Rob Pike
// http://commandcenter.blogspot.de/2012/04/byte-order-fallacy.html

package image

import (
	"fmt"
	"image"
	"image/draw"
	"reflect"

	colorExt "github.com/chai2010/image/color"
)

var (
	_ Image = (*Gray)(nil)
	_ Image = (*Gray16)(nil)
	_ Image = (*Gray32i)(nil)
	_ Image = (*Gray32f)(nil)
	_ Image = (*Gray64i)(nil)
	_ Image = (*Gray64f)(nil)
	_ Image = (*GrayA)(nil)
	_ Image = (*GrayA32)(nil)
	_ Image = (*GrayA64i)(nil)
	_ Image = (*GrayA64f)(nil)
	_ Image = (*GrayA128i)(nil)
	_ Image = (*GrayA128f)(nil)
	_ Image = (*RGB)(nil)
	_ Image = (*RGB48)(nil)
	_ Image = (*RGB96i)(nil)
	_ Image = (*RGB96f)(nil)
	_ Image = (*RGB192i)(nil)
	_ Image = (*RGB192f)(nil)
	_ Image = (*RGBA)(nil)
	_ Image = (*RGBA64)(nil)
	_ Image = (*RGBA128i)(nil)
	_ Image = (*RGBA128f)(nil)
	_ Image = (*RGBA256i)(nil)
	_ Image = (*RGBA256f)(nil)
)

type Image interface {
	// Get original type, such as *image.Gray, *image.RGBA, etc.
	BaseType() image.Image

	// Pix holds the image's pixels, as pixel values in big-endian order format. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*PixelSize].
	Pix() []byte
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride() int
	// Rect is the image's bounds.
	Rect() image.Rectangle

	// 1:Gray, 2:GrayA, 3:RGB, 4:RGBA
	Channels() int
	// Uint8/Uint16/Int32/Int64/Float32/Float64
	Depth() reflect.Kind

	draw.Image
}

func newRGBAFromImage(m image.Image) *RGBA {
	b := m.Bounds()
	rgba := NewRGBA(b)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			pr, pg, pb, pa := m.At(x, y).RGBA()
			rgba.SetRGBA(x, y, colorExt.RGBA{
				R: uint8(pr >> 8),
				G: uint8(pg >> 8),
				B: uint8(pb >> 8),
				A: uint8(pa >> 8),
			})
		}
	}
	return rgba
}

func newRGBA64FromImage(m image.Image) *RGBA64 {
	b := m.Bounds()
	rgba64 := NewRGBA64(b)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			pr, pg, pb, pa := m.At(x, y).RGBA()
			rgba64.SetRGBA64(x, y, colorExt.RGBA64{
				R: uint16(pr),
				G: uint16(pg),
				B: uint16(pb),
				A: uint16(pa),
			})
		}
	}
	return rgba64
}

func asBaseType(m Image) image.Image {
	switch channels, depth := m.Channels(), m.Depth(); {
	case channels == 1 && depth == reflect.Uint8:
		return &image.Gray{
			Pix:    m.Pix(),
			Stride: m.Stride(),
			Rect:   m.Rect(),
		}
	case channels == 1 && depth == reflect.Uint16:
		return &image.Gray16{
			Pix:    m.Pix(),
			Stride: m.Stride(),
			Rect:   m.Rect(),
		}
	case channels == 4 && depth == reflect.Uint8:
		return &image.RGBA{
			Pix:    m.Pix(),
			Stride: m.Stride(),
			Rect:   m.Rect(),
		}
	case channels == 4 && depth == reflect.Uint16:
		return &image.RGBA64{
			Pix:    m.Pix(),
			Stride: m.Stride(),
			Rect:   m.Rect(),
		}
	}
	return m
}

func AsImage(m image.Image) Image {
	if p, ok := m.(Image); ok {
		return p
	}

	switch m := m.(type) {
	case *image.Gray:
		return new(Gray).Init(m.Pix, m.Stride, m.Rect)
	case *image.Gray16:
		return new(Gray16).Init(m.Pix, m.Stride, m.Rect)
	case *image.RGBA:
		return new(RGBA).Init(m.Pix, m.Stride, m.Rect)
	case *image.RGBA64:
		return new(RGBA64).Init(m.Pix, m.Stride, m.Rect)
	case *image.YCbCr:
		return newRGBAFromImage(m)
	}

	return newRGBA64FromImage(m)
}

func CloneImage(m image.Image) Image {
	if m, ok := m.(Image); ok {
		switch channels, depth := m.Channels(), m.Depth(); {
		case channels == 1 && depth == reflect.Uint8:
			return new(Gray).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 1 && depth == reflect.Uint16:
			return new(Gray16).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 1 && depth == reflect.Int32:
			return new(Gray32i).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 1 && depth == reflect.Float32:
			return new(Gray32f).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 1 && depth == reflect.Int64:
			return new(Gray64i).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 1 && depth == reflect.Float64:
			return new(Gray64f).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())

		case channels == 2 && depth == reflect.Uint8:
			return new(GrayA).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 2 && depth == reflect.Uint16:
			return new(GrayA32).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 2 && depth == reflect.Int32:
			return new(GrayA64i).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 2 && depth == reflect.Float32:
			return new(GrayA64f).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 2 && depth == reflect.Int64:
			return new(GrayA128i).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 2 && depth == reflect.Float64:
			return new(GrayA128f).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())

		case channels == 3 && depth == reflect.Uint8:
			return new(RGB).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 3 && depth == reflect.Uint16:
			return new(RGB48).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 3 && depth == reflect.Int32:
			return new(RGB96i).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 3 && depth == reflect.Float32:
			return new(RGB96f).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 3 && depth == reflect.Int64:
			return new(RGB192i).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 3 && depth == reflect.Float64:
			return new(RGB192f).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())

		case channels == 4 && depth == reflect.Uint8:
			return new(RGBA).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 4 && depth == reflect.Uint16:
			return new(RGBA64).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 4 && depth == reflect.Int32:
			return new(RGBA128i).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 4 && depth == reflect.Float32:
			return new(RGBA128f).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 4 && depth == reflect.Int64:
			return new(RGBA256i).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 4 && depth == reflect.Float64:
			return new(RGBA256f).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())

		default:
			panic(fmt.Errorf("image: CloneImage, invalid format: channels = %v, depth = %v", channels, depth))
		}
	}

	switch m := m.(type) {
	case *image.Gray:
		return new(Gray).Init(append([]uint8(nil), m.Pix...), m.Stride, m.Rect)
	case *image.Gray16:
		return new(Gray16).Init(append([]uint8(nil), m.Pix...), m.Stride, m.Rect)
	case *image.RGBA:
		return new(RGBA).Init(append([]uint8(nil), m.Pix...), m.Stride, m.Rect)
	case *image.RGBA64:
		return new(RGBA64).Init(append([]uint8(nil), m.Pix...), m.Stride, m.Rect)
	case *image.YCbCr:
		return newRGBAFromImage(m)
	}

	return newRGBA64FromImage(m)
}

func NewImage(r image.Rectangle, channels int, depth reflect.Kind) (m Image, err error) {
	switch {
	case channels == 1 && depth == reflect.Uint8:
		m = NewGray(r)
		return
	case channels == 1 && depth == reflect.Uint16:
		m = NewGray16(r)
		return
	case channels == 1 && depth == reflect.Int32:
		m = NewGray32i(r)
		return
	case channels == 1 && depth == reflect.Float32:
		m = NewGray32f(r)
		return
	case channels == 1 && depth == reflect.Int64:
		m = NewGray64i(r)
		return
	case channels == 1 && depth == reflect.Float64:
		m = NewGray64f(r)
		return

	case channels == 2 && depth == reflect.Uint8:
		m = NewGrayA(r)
		return
	case channels == 2 && depth == reflect.Uint16:
		m = NewGrayA32(r)
		return
	case channels == 2 && depth == reflect.Int32:
		m = NewGrayA64f(r)
		return
	case channels == 2 && depth == reflect.Float32:
		m = NewGrayA64f(r)
		return
	case channels == 2 && depth == reflect.Int64:
		m = NewGrayA128f(r)
		return
	case channels == 2 && depth == reflect.Float64:
		m = NewGrayA128f(r)
		return

	case channels == 3 && depth == reflect.Uint8:
		m = NewRGB(r)
		return
	case channels == 3 && depth == reflect.Uint16:
		m = NewRGB48(r)
		return
	case channels == 3 && depth == reflect.Int32:
		m = NewRGB96i(r)
		return
	case channels == 3 && depth == reflect.Float32:
		m = NewRGB96f(r)
		return
	case channels == 3 && depth == reflect.Int64:
		m = NewRGB192i(r)
		return
	case channels == 3 && depth == reflect.Float64:
		m = NewRGB192f(r)
		return

	case channels == 4 && depth == reflect.Uint8:
		m = NewRGBA(r)
		return
	case channels == 4 && depth == reflect.Uint16:
		m = NewRGBA64(r)
		return
	case channels == 4 && depth == reflect.Int32:
		m = NewRGBA128i(r)
		return
	case channels == 4 && depth == reflect.Float32:
		m = NewRGBA128f(r)
		return
	case channels == 4 && depth == reflect.Int64:
		m = NewRGBA256i(r)
		return
	case channels == 4 && depth == reflect.Float64:
		m = NewRGBA256f(r)
		return

	default:
		err = fmt.Errorf("image: NewImage, invalid format: channels = %v, depth = %v", channels, depth)
		return
	}
}
