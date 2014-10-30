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
	"testing"

	imageExt "github.com/chai2010/image"
	colorExt "github.com/chai2010/image/color"
)

type tTester struct {
	Image    draw.Image
	Model    color.Model
	DataType reflect.Kind
	Channels int
}

var tTesterList = []tTester{
	// Gray/Gray16/Gray32f
	tTester{
		Image:    image.NewGray(image.Rect(0, 0, 10, 10)),
		Model:    color.GrayModel,
		DataType: reflect.Uint8,
		Channels: 1,
	},
	tTester{
		Image:    image.NewGray16(image.Rect(0, 0, 10, 10)),
		Model:    color.Gray16Model,
		DataType: reflect.Uint16,
		Channels: 1,
	},
	tTester{
		Image:    imageExt.NewGray32f(image.Rect(0, 0, 10, 10)),
		Model:    colorExt.Gray32fModel,
		DataType: reflect.Float32,
		Channels: 1,
	},
	// RGB/RGB48/RGB96f
	tTester{
		Image:    imageExt.NewRGB(image.Rect(0, 0, 10, 10)),
		Model:    colorExt.RGBModel,
		DataType: reflect.Uint8,
		Channels: 3,
	},
	tTester{
		Image:    imageExt.NewRGB48(image.Rect(0, 0, 10, 10)),
		Model:    colorExt.RGB48Model,
		DataType: reflect.Uint16,
		Channels: 3,
	},
	tTester{
		Image:    imageExt.NewRGB96f(image.Rect(0, 0, 10, 10)),
		Model:    colorExt.RGB96fModel,
		DataType: reflect.Float32,
		Channels: 3,
	},
	// RGBA/RGBA48/RGBA128f
	tTester{
		Image:    image.NewRGBA(image.Rect(0, 0, 10, 10)),
		Model:    color.RGBAModel,
		DataType: reflect.Uint8,
		Channels: 4,
	},
	tTester{
		Image:    image.NewRGBA64(image.Rect(0, 0, 10, 10)),
		Model:    color.RGBA64Model,
		DataType: reflect.Uint16,
		Channels: 4,
	},
	tTester{
		Image:    imageExt.NewRGBA128f(image.Rect(0, 0, 10, 10)),
		Model:    colorExt.RGBA128fModel,
		DataType: reflect.Float32,
		Channels: 4,
	},
}

func TestEncodeAndDecode(t *testing.T) {
	for _, v := range tTesterList {
		v.Image.Set(6, 3, color.RGBA{0xAA, 0xBB, 0xCC, 0xDD})
	}
	for i, v := range tTesterList {
		encoder := pixEncoder{v.Channels, v.DataType}
		decoder := pixDecoder{v.Channels, v.DataType, v.Image.Bounds().Dx(), v.Image.Bounds().Dy()}

		data, err := encoder.Encode(v.Image, nil)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}

		m, err := decoder.Decode(data, nil)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}

		err = tCompareImage(v.Image, v.Channels, v.Model, m)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}
	}
}

func tCompareImage(img0 image.Image, channels int, model color.Model, img1 image.Image) error {
	if img1.ColorModel() != model {
		return fmt.Errorf("img1 wrong image model: want %v, got %v", model, img1.ColorModel())
	}
	if !img1.Bounds().Eq(img0.Bounds()) {
		return fmt.Errorf("wrong image size: want %v, got %v", img0.Bounds(), img1.Bounds())
	}

	switch img0.ColorModel() {
	case colorExt.Gray32fModel:
		b := img1.Bounds()
		img0 := img0.(*imageExt.Gray32f)
		img1 := img1.(*imageExt.Gray32f)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				c0 := img0.Gray32fAt(x, y)
				c1 := img1.Gray32fAt(x, y)
				if c0 != c1 {
					return fmt.Errorf("pixel at (%d, %d) has wrong color: want %v, got %v", x, y, c0, c1)
				}
			}
		}
	case colorExt.RGB96fModel:
		b := img1.Bounds()
		img0 := img0.(*imageExt.RGB96f)
		img1 := img1.(*imageExt.RGB96f)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				c0 := img0.RGB96fAt(x, y)
				c1 := img1.RGB96fAt(x, y)
				if c0 != c1 {
					return fmt.Errorf("pixel at (%d, %d) has wrong color: want %v, got %v", x, y, c0, c1)
				}
			}
		}
	case colorExt.RGBA128fModel:
		b := img1.Bounds()
		img0 := img0.(*imageExt.RGBA128f)
		img1 := img1.(*imageExt.RGBA128f)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				c0 := img0.RGBA128fAt(x, y)
				c1 := img1.RGBA128fAt(x, y)
				if channels == 3 {
					c0.A, c1.A = 0, 0
				}
				if c0 != c1 {
					return fmt.Errorf("pixel at (%d, %d) has wrong color: want %v, got %v", x, y, c0, c1)
				}
			}
		}
	default:
		b := img1.Bounds()
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				switch channels {
				case 4:
					c0 := img0.At(x, y)
					c1 := img1.At(x, y)
					r0, g0, b0, a0 := c0.RGBA()
					r1, g1, b1, a1 := c1.RGBA()
					if r0 != r1 || g0 != g1 || b0 != b1 || a0 != a1 {
						return fmt.Errorf("pixel at (%d, %d) has wrong color: want %v, got %v", x, y, c0, c1)
					}
				case 3:
					c0 := img0.At(x, y)
					c1 := img1.At(x, y)
					r0, g0, b0, _ := c0.RGBA()
					r1, g1, b1, _ := c1.RGBA()
					if r0 != r1 || g0 != g1 || b0 != b1 {
						return fmt.Errorf("pixel at (%d, %d) has wrong color: want %v, got %v", x, y, c0, c1)
					}
				case 1:
					c0 := color.GrayModel.Convert(img0.At(x, y)).(color.Gray)
					c1 := color.GrayModel.Convert(img1.At(x, y)).(color.Gray)
					if c0 != c1 {
						return fmt.Errorf("pixel at (%d, %d) has wrong color: want %v, got %v", x, y, c0, c1)
					}
				}
			}
		}
	}
	return nil
}

// new a zero YCbCr
func tNewYCbCr(r image.Rectangle, subsampleRatio image.YCbCrSubsampleRatio) *image.YCbCr {
	m := image.NewYCbCr(r, subsampleRatio)
	for i := 0; i < len(m.Cb); i++ {
		m.Cb[i] = tZeroYCbCr.Cb
	}
	for i := 0; i < len(m.Cr); i++ {
		m.Cr[i] = tZeroYCbCr.Cr
	}
	return m
}

// YCbCr.Set
func tSetYCbCr(p *image.YCbCr, x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	yi := p.YOffset(x, y)
	ci := p.COffset(x, y)
	c1 := color.YCbCrModel.Convert(c).(color.YCbCr)
	p.Y[yi] = c1.Y
	p.Cb[ci] = c1.Cb
	p.Cr[ci] = c1.Cr
}

var tZeroYCbCr = func() (c color.YCbCr) {
	c.Y, c.Cb, c.Cr = color.RGBToYCbCr(0, 0, 0)
	return
}()
