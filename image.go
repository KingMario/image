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

var (
	_ Image = (*Gray)(nil)
	_ Image = (*Gray16)(nil)
	_ Image = (*Gray32f)(nil)
	_ Image = (*GrayA)(nil)
	_ Image = (*GrayA32)(nil)
	_ Image = (*GrayA64f)(nil)
	_ Image = (*RGB)(nil)
	_ Image = (*RGB48)(nil)
	_ Image = (*RGB96f)(nil)
	_ Image = (*RGBA)(nil)
	_ Image = (*RGBA64)(nil)
	_ Image = (*RGBA128f)(nil)
	_ Image = (*Unknown)(nil)
)

type Image interface {
	// Get original type, such as *image.Gray, *image.RGBA, etc.
	BaseType() image.Image

	// Pix holds the image's pixels, as pixel values in machine order format. The pixel at
	// (x, y) starts at []Type(Pix[(y-Rect.Min.Y)*Stride:])[(x-Rect.Min.X)*Channels:].
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

func newRGBA64FromImage(m image.Image) *RGBA64 {
	b := m.Bounds()
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

func AsImage(m image.Image) Image {
	if p, ok := m.(Image); ok {
		return p
	}

	switch m := m.(type) {
	case *image.Gray:
		return &Gray{m}
	case *image.Gray16:
		return &Gray16{m}
	case *image.RGBA:
		return &RGBA{m}
	case *image.RGBA64:
		return &RGBA64{m}
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
		case channels == 1 && depth == reflect.Float32:
			return new(Gray32f).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 2 && depth == reflect.Uint8:
			return new(GrayA).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 2 && depth == reflect.Uint16:
			return new(GrayA32).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 2 && depth == reflect.Float32:
			return new(GrayA64f).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 3 && depth == reflect.Uint8:
			return new(RGB).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 3 && depth == reflect.Uint16:
			return new(RGB48).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 3 && depth == reflect.Float32:
			return new(RGB96f).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 4 && depth == reflect.Uint8:
			return new(RGBA).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 4 && depth == reflect.Uint16:
			return new(RGBA64).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		case channels == 4 && depth == reflect.Float32:
			return new(RGBA128f).Init(append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect())
		}

		return new(Unknown).Init(
			append([]uint8(nil), m.Pix()...), m.Stride(), m.Rect(),
			m.Channels(), m.Depth(),
		)
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
	}

	return newRGBA64FromImage(m)
}

func CopyImage(buf, src image.Image) (dst Image) {
	if buf == nil {
		return CloneImage(src)
	}

	switch buf := buf.(type) {
	case *image.Gray:
		return (&Gray{buf}).CopyFrom(src)
	case *image.Gray16:
		return (&Gray16{buf}).CopyFrom(src)
	case *image.RGBA:
		return (&RGBA{buf}).CopyFrom(src)
	case *image.RGBA64:
		return (&RGBA64{buf}).CopyFrom(src)

	case *Gray:
		return buf.CopyFrom(src)
	case *Gray16:
		return buf.CopyFrom(src)
	case *Gray32f:
		return buf.CopyFrom(src)

	case *GrayA:
		return buf.CopyFrom(src)
	case *GrayA32:
		return buf.CopyFrom(src)
	case *GrayA64f:
		return buf.CopyFrom(src)

	case *RGB:
		return buf.CopyFrom(src)
	case *RGB48:
		return buf.CopyFrom(src)
	case *RGB96f:
		return buf.CopyFrom(src)

	case *RGBA:
		return buf.CopyFrom(src)
	case *RGBA64:
		return buf.CopyFrom(src)
	case *RGBA128f:
		return buf.CopyFrom(src)

	case *Unknown:
		return buf.CopyFrom(src)
	}

	return CloneImage(src)
}

func ConvertCopyImage(m image.Image, channels int, depth reflect.Kind) (Image, error) {
	buf, err := NewImage(m.Bounds(), channels, depth)
	if err != nil {
		return nil, err
	}
	p := CopyImage(buf, m)
	return p, nil
}

func ConvertImage(m image.Image, channels int, depth reflect.Kind) (Image, error) {
	if p, ok := m.(Image); ok {
		if channels == p.Channels() && depth == p.Depth() {
			return p, nil
		}
	}
	buf, err := NewImage(m.Bounds(), channels, depth)
	if err != nil {
		return nil, err
	}
	p := CopyImage(buf, m)
	return p, nil
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
