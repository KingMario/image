// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

import (
	"image/color"
)

type Gray color.Gray

func (c Gray) RGBA() (r, g, b, a uint32) {
	return
}

type Gray16 color.Gray16

func (c Gray16) RGBA() (r, g, b, a uint32) {
	return
}

type Gray32i struct {
	Y int32
}

func (c Gray32i) RGBA() (r, g, b, a uint32) {
	return
}

type Gray32f struct {
	Y float32
}

func (c Gray32f) RGBA() (r, g, b, a uint32) {
	return
}

type Gray64i struct {
	Y int64
}

func (c Gray64i) RGBA() (r, g, b, a uint32) {
	return
}

type Gray64f struct {
	Y float64
}

func (c Gray64f) RGBA() (r, g, b, a uint32) {
	return
}
