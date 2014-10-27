// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

import (
	"image/color"
)

type RGB struct {
	color.Gray
}

func (c RGB) RGBA() (r, g, b, a uint32) {
	return
}

type RGB48 struct {
	color.Gray16
}

func (c RGB48) RGBA() (r, g, b, a uint32) {
	return
}

type RGB96i struct {
	Y int32
}

func (c RGB96i) RGBA() (r, g, b, a uint32) {
	return
}

type RGB96f struct {
	Y float32
}

func (c RGB96f) RGBA() (r, g, b, a uint32) {
	return
}

type RGB192i struct {
	Y int64
}

func (c RGB192i) RGBA() (r, g, b, a uint32) {
	return
}

type RGB192f struct {
	Y float64
}

func (c RGB192f) RGBA() (r, g, b, a uint32) {
	return
}
