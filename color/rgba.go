// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

import (
	"image/color"
)

type RGBA color.RGBA

func (c RGBA) RGBA() (r, g, b, a uint32) {
	return
}

type RGBA64 color.RGBA64

func (c RGBA64) RGBA() (r, g, b, a uint32) {
	return
}

type RGBA128i struct {
	Y int32
}

func (c RGBA128i) RGBA() (r, g, b, a uint32) {
	return
}

type RGBA128f struct {
	Y float32
}

func (c RGBA128f) RGBA() (r, g, b, a uint32) {
	return
}

type RGBA256i struct {
	Y int64
}

func (c RGBA256i) RGBA() (r, g, b, a uint32) {
	return
}

type RGBA256f struct {
	Y float64
}

func (c RGBA256f) RGBA() (r, g, b, a uint32) {
	return
}
