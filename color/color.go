// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package color implements a basic color library.
package color

import (
	"image/color"
)

var (
	_ color.Color = (*Gray)(nil)
	_ color.Color = (*Gray16)(nil)
	_ color.Color = (*Gray32i)(nil)
	_ color.Color = (*Gray32f)(nil)
	_ color.Color = (*Gray64i)(nil)
	_ color.Color = (*Gray64f)(nil)
	_ color.Color = (*GrayA)(nil)
	_ color.Color = (*GrayA32)(nil)
	_ color.Color = (*GrayA64i)(nil)
	_ color.Color = (*GrayA64f)(nil)
	_ color.Color = (*GrayA128i)(nil)
	_ color.Color = (*GrayA128f)(nil)
	_ color.Color = (*RGB)(nil)
	_ color.Color = (*RGB48)(nil)
	_ color.Color = (*RGB96i)(nil)
	_ color.Color = (*RGB96f)(nil)
	_ color.Color = (*RGB192i)(nil)
	_ color.Color = (*RGB192f)(nil)
	_ color.Color = (*RGBA)(nil)
	_ color.Color = (*RGBA64)(nil)
	_ color.Color = (*RGBA128i)(nil)
	_ color.Color = (*RGBA128f)(nil)
	_ color.Color = (*RGBA256i)(nil)
	_ color.Color = (*RGBA256f)(nil)
)

// Models for the standard color types.
var (
	GrayModel      color.Model = color.ModelFunc(grayModel)
	Gray16Model    color.Model = color.ModelFunc(gray16Model)
	Gray32iModel   color.Model = color.ModelFunc(gray32iModel)
	Gray32fModel   color.Model = color.ModelFunc(gray32fModel)
	Gray64iModel   color.Model = color.ModelFunc(gray64iModel)
	Gray64fModel   color.Model = color.ModelFunc(gray64fModel)
	GrayAModel     color.Model = color.ModelFunc(grayAModel)
	GrayA32Model   color.Model = color.ModelFunc(grayA32Model)
	GrayA64iModel  color.Model = color.ModelFunc(grayA64iModel)
	GrayA64fModel  color.Model = color.ModelFunc(grayA64fModel)
	GrayA128iModel color.Model = color.ModelFunc(grayA128iModel)
	GrayA128fModel color.Model = color.ModelFunc(grayA128fModel)
	RGBModel       color.Model = color.ModelFunc(rgbModel)
	RGB48Model     color.Model = color.ModelFunc(rgb48Model)
	RGB96iModel    color.Model = color.ModelFunc(rgb96iModel)
	RGB96fModel    color.Model = color.ModelFunc(rgb96fModel)
	RGB192iModel   color.Model = color.ModelFunc(rgb192iModel)
	RGB192fModel   color.Model = color.ModelFunc(rgb192fModel)
	RGBAModel      color.Model = color.ModelFunc(rgbaModel)
	RGBA64Model    color.Model = color.ModelFunc(rgba64Model)
	RGBA128iModel  color.Model = color.ModelFunc(rgba128iModel)
	RGBA128fModel  color.Model = color.ModelFunc(rgba128fModel)
	RGBA256iModel  color.Model = color.ModelFunc(rgba256iModel)
	RGBA256fModel  color.Model = color.ModelFunc(rgba256fModel)
)

func colorRgbToGray(r, g, b uint32) uint32 {
	y := (299*r + 587*g + 114*b + 500) / 1000
	return y
}

func colorRgbToGrayI32(r, g, b int32) int32 {
	y := (299*r + 587*g + 114*b + 500) / 1000
	return y
}

func colorRgbToGrayF32(r, g, b float32) float32 {
	y := (299*r + 587*g + 114*b + 500) / 1000
	return y
}

func colorRgbToGrayI64(r, g, b int64) int64 {
	y := (299*r + 587*g + 114*b + 500) / 1000
	return y
}

func colorRgbToGrayF64(r, g, b float64) float64 {
	y := (299*r + 587*g + 114*b + 500) / 1000
	return y
}
