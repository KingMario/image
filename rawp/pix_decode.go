// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rawp

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"reflect"

	imageExt "github.com/chai2010/image"
	colorExt "github.com/chai2010/image/color"
)

type pixDecoder struct {
	Channels int          // 1/2/3/4
	DataType reflect.Kind // Uint8/Uint16/Int32/Int64/Float32/Float64
	Width    int          // need for Decode
	Height   int          // need for Decode
}

func (p *pixDecoder) Decode(data []byte, buf imageExt.Buffer) (m draw.Image, err error) {
	// Gray/Gray16/Gray32i/Gray32f/Gray64i/Gray64f
	if p.Channels == 1 && p.DataType == reflect.Uint8 {
		return p.decodeGray(data, buf)
	}
	if p.Channels == 1 && p.DataType == reflect.Uint16 {
		return p.decodeGray16(data, buf)
	}
	if p.Channels == 1 && p.DataType == reflect.Int32 {
		return p.decodeGray32f(data, buf)
	}
	if p.Channels == 1 && p.DataType == reflect.Float32 {
		return p.decodeGray32f(data, buf)
	}
	if p.Channels == 1 && p.DataType == reflect.Int64 {
		return p.decodeGray64f(data, buf)
	}
	if p.Channels == 1 && p.DataType == reflect.Float64 {
		return p.decodeGray64f(data, buf)
	}

	// GrayA/GrayA32/GrayA64i/GrayA64f/GrayA128i/GrayA128f
	if p.Channels == 1 && p.DataType == reflect.Uint8 {
		return p.decodeGrayA(data, buf)
	}
	if p.Channels == 1 && p.DataType == reflect.Uint16 {
		return p.decodeGrayA32(data, buf)
	}
	if p.Channels == 1 && p.DataType == reflect.Int32 {
		return p.decodeGrayA64i(data, buf)
	}
	if p.Channels == 1 && p.DataType == reflect.Float32 {
		return p.decodeGrayA64f(data, buf)
	}
	if p.Channels == 1 && p.DataType == reflect.Int64 {
		return p.decodeGrayA128i(data, buf)
	}
	if p.Channels == 1 && p.DataType == reflect.Float64 {
		return p.decodeGrayA128f(data, buf)
	}

	// RGB/RGB48/RGB96i/RGB96f/RGB192i/RGB192f
	if p.Channels == 3 && p.DataType == reflect.Uint8 {
		return p.decodeRGB(data, buf)
	}
	if p.Channels == 3 && p.DataType == reflect.Uint16 {
		return p.decodeRGB48(data, buf)
	}
	if p.Channels == 3 && p.DataType == reflect.Int32 {
		return p.decodeRGB96i(data, buf)
	}
	if p.Channels == 3 && p.DataType == reflect.Float32 {
		return p.decodeRGB96f(data, buf)
	}
	if p.Channels == 3 && p.DataType == reflect.Int64 {
		return p.decodeRGB192i(data, buf)
	}
	if p.Channels == 3 && p.DataType == reflect.Float64 {
		return p.decodeRGB192f(data, buf)
	}

	// RGBA/RGBA64/RGBA128f
	if p.Channels == 4 && p.DataType == reflect.Uint8 {
		return p.decodeRGBA(data, buf)
	}
	if p.Channels == 4 && p.DataType == reflect.Uint16 {
		return p.decodeRGBA64(data, buf)
	}
	if p.Channels == 4 && p.DataType == reflect.Int32 {
		return p.decodeRGBA128i(data, buf)
	}
	if p.Channels == 4 && p.DataType == reflect.Float32 {
		return p.decodeRGBA128f(data, buf)
	}
	if p.Channels == 4 && p.DataType == reflect.Int64 {
		return p.decodeRGBA256i(data, buf)
	}
	if p.Channels == 4 && p.DataType == reflect.Float64 {
		return p.decodeRGBA256f(data, buf)
	}

	// Unknown
	err = fmt.Errorf(
		"image/rawp: Decode, unknown image format, channels = %v, dataType = %v",
		p.Channels, p.DataType,
	)
	return
}

func (p *pixDecoder) getPixelSize() int {
	switch p.DataType {
	case reflect.Uint8:
		return p.Channels * 1
	case reflect.Uint16:
		return p.Channels * 2
	case reflect.Int32:
		return p.Channels * 4
	case reflect.Float32:
		return p.Channels * 4
	case reflect.Int64:
		return p.Channels * 8
	case reflect.Float64:
		return p.Channels * 8
	}
	panic("image/rawp: getPixelSize, unreachable")
}

func (p *pixDecoder) getImageDataSize() int {
	return p.getPixelSize() * p.Width * p.Height
}

func (p *pixDecoder) decodeGray(data []byte, buf imageExt.Buffer) (m draw.Image, err error) {
	if size := p.getImageDataSize(); len(data) != size {
		err = fmt.Errorf("image/rawp: decodeGray, bad data size, expect = %d, got = %d", size, len(data))
		return
	}
	gray := newGray(image.Rect(0, 0, p.Width, p.Height), buf)
	var off = 0
	for y := 0; y < p.Height; y++ {
		copy(gray.Pix[y*gray.Stride:][:p.Width], data[off:])
		off += p.Width
	}
	m = gray
	return
}

func (p *pixDecoder) decodeGray16(data []byte, buf imageExt.Buffer) (m draw.Image, err error) {
	if size := p.getImageDataSize(); len(data) != size {
		err = fmt.Errorf("image/rawp: decodeGray16, bad data size, expect = %d, got = %d", size, len(data))
		return
	}
	gray16 := newGray16(image.Rect(0, 0, p.Width, p.Height), buf)
	var off = 0
	for y := 0; y < p.Height; y++ {
		u16Pix := builtin.Slice(data[off:], reflect.TypeOf([]uint16(nil))).([]uint16)
		for x := 0; x < p.Width; x++ {
			gray16.SetGray16(x, y, color.Gray16{u16Pix[x]})
		}
		off += p.Width * 2
	}
	m = gray16
	return
}

func (p *pixDecoder) decodeGray32f(data []byte, buf imageExt.Buffer) (m draw.Image, err error) {
	if size := p.getImageDataSize(); len(data) != size {
		err = fmt.Errorf("image/rawp: decodeGray32f, bad data size, expect = %d, got = %d", size, len(data))
		return
	}
	gray32f := newGray32f(image.Rect(0, 0, p.Width, p.Height), buf)
	var off = 0
	for y := 0; y < p.Height; y++ {
		copy(gray32f.Pix[y*gray32f.Stride:][:p.Width*4], data[off:])
		off += p.Width * 4
	}
	m = gray32f
	return
}

func (p *pixDecoder) decodeRGB(data []byte, buf imageExt.Buffer) (m draw.Image, err error) {
	if size := p.getImageDataSize(); len(data) != size {
		err = fmt.Errorf("image/rawp: decodeRGB, bad data size, expect = %d, got = %d", size, len(data))
		return
	}
	rgb := newRGB(image.Rect(0, 0, p.Width, p.Height), buf)
	var off = 0
	for y := 0; y < p.Height; y++ {
		for x := 0; x < p.Width; x++ {
			rgb.SetRGB(x, y, colorExt.RGB{
				R: data[off+0],
				G: data[off+1],
				B: data[off+2],
			})
			off += 3
		}
	}
	m = rgb
	return
}

func (p *pixDecoder) decodeRGB48(data []byte, buf imageExt.Buffer) (m draw.Image, err error) {
	if size := p.getImageDataSize(); len(data) != size {
		err = fmt.Errorf("image/rawp: decodeRGB48, bad data size, expect = %d, got = %d", size, len(data))
		return
	}
	rgb48 := newRGB48(image.Rect(0, 0, p.Width, p.Height), buf)
	var off = 0
	for y := 0; y < p.Height; y++ {
		u16Pix := builtin.Slice(data[off:], reflect.TypeOf([]uint16(nil))).([]uint16)
		for x := 0; x < p.Width; x++ {
			rgb48.SetRGB48(x, y, colorExt.RGB48{
				R: u16Pix[x*3+0],
				G: u16Pix[x*3+1],
				B: u16Pix[x*3+2],
			})
		}
		off += p.Width * 6
	}
	m = rgb48
	return
}

func (p *pixDecoder) decodeRGB96f(data []byte, buf imageExt.Buffer) (m draw.Image, err error) {
	if size := p.getImageDataSize(); len(data) != size {
		err = fmt.Errorf("image/rawp: decodeRGB96f, bad data size, expect = %d, got = %d", size, len(data))
		return
	}
	rgb96f := newRGB96f(image.Rect(0, 0, p.Width, p.Height), buf)
	var off = 0
	for y := 0; y < p.Height; y++ {
		for x := 0; x < p.Width; x++ {
			rgb96f.SetRGB96f(x, y, colorExt.RGB96f{
				R: builtin.Float32(data[off+0:]),
				G: builtin.Float32(data[off+4:]),
				B: builtin.Float32(data[off+8:]),
			})
			off += 12
		}
	}
	m = rgb96f
	return
}

func (p *pixDecoder) decodeRGBA(data []byte, buf imageExt.Buffer) (m draw.Image, err error) {
	if size := p.getImageDataSize(); len(data) != size {
		err = fmt.Errorf("image/rawp: decodeRGBA, bad data size, expect = %d, got = %d", size, len(data))
		return
	}
	rgba := newRGBA(image.Rect(0, 0, p.Width, p.Height), buf)
	var off = 0
	for y := 0; y < p.Height; y++ {
		copy(rgba.Pix[y*rgba.Stride:][:p.Width*4], data[off:])
		off += p.Width * 4
	}
	m = rgba
	return
}

func (p *pixDecoder) decodeRGBA64(data []byte, buf imageExt.Buffer) (m draw.Image, err error) {
	if size := p.getImageDataSize(); len(data) != size {
		err = fmt.Errorf("image/rawp: decodeRGBA64, bad data size, expect = %d, got = %d", size, len(data))
		return
	}
	rgba64 := newRGBA64(image.Rect(0, 0, p.Width, p.Height), buf)
	var off = 0
	for y := 0; y < p.Height; y++ {
		u16Pix := builtin.Slice(data[off:], reflect.TypeOf([]uint16(nil))).([]uint16)
		for x := 0; x < p.Width; x++ {
			rgba64.SetRGBA64(x, y, color.RGBA64{
				R: u16Pix[x*4+0],
				G: u16Pix[x*4+1],
				B: u16Pix[x*4+2],
				A: u16Pix[x*4+3],
			})
		}
		off += p.Width * 8
	}
	m = rgba64
	return
}

func (p *pixDecoder) decodeRGBA128f(data []byte, buf imageExt.Buffer) (m draw.Image, err error) {
	if size := p.getImageDataSize(); len(data) != size {
		err = fmt.Errorf("image/rawp: decodeRGBA128f, bad data size, expect = %d, got = %d", size, len(data))
		return
	}
	rgba128f := newRGBA128f(image.Rect(0, 0, p.Width, p.Height), buf)
	var off = 0
	for y := 0; y < p.Height; y++ {
		copy(rgba128f.Pix[y*rgba128f.Stride:][:p.Width*16], data[off:])
		off += p.Width * 16
	}
	m = rgba128f
	return
}
