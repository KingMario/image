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
	GrayModel      color.Model = color.GrayModel
	Gray16Model    color.Model = color.Gray16Model
	Gray32iModel   color.Model = color.ModelFunc(xModel)
	Gray32fModel   color.Model = color.ModelFunc(xModel)
	Gray64iModel   color.Model = color.ModelFunc(xModel)
	Gray64fModel   color.Model = color.ModelFunc(xModel)
	GrayAModel     color.Model = color.GrayModel
	GrayA32Model   color.Model = color.Gray16Model
	GrayA64iModel  color.Model = color.ModelFunc(xModel)
	GrayA64fModel  color.Model = color.ModelFunc(xModel)
	GrayA128iModel color.Model = color.ModelFunc(xModel)
	GrayA128fModel color.Model = color.ModelFunc(xModel)
	RGBModel       color.Model = color.GrayModel
	RGB48Model     color.Model = color.Gray16Model
	RGB96iModel    color.Model = color.ModelFunc(xModel)
	RGB96fModel    color.Model = color.ModelFunc(xModel)
	RGB192iModel   color.Model = color.ModelFunc(xModel)
	RGB192fModel   color.Model = color.ModelFunc(xModel)
	RGBAModel      color.Model = color.RGBAModel
	RGBA64Model    color.Model = color.RGBA64Model
	RGB128iModel   color.Model = color.ModelFunc(xModel)
	RGB128fModel   color.Model = color.ModelFunc(xModel)
	RGB256iModel   color.Model = color.ModelFunc(xModel)
	RGB256fModel   color.Model = color.ModelFunc(xModel)
)

func xModel(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	return c
}
