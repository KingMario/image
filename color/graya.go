// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

type GrayA struct {
	Y, A uint8
}

func (c GrayA) RGBA() (r, g, b, a uint32) {
	return
}

type GrayA32 struct {
	Y, A uint16
}

func (c GrayA32) RGBA() (r, g, b, a uint32) {
	return
}

type GrayA64i struct {
	Y, A int32
}

func (c GrayA64i) RGBA() (r, g, b, a uint32) {
	return
}

type GrayA64f struct {
	Y, A float32
}

func (c GrayA64f) RGBA() (r, g, b, a uint32) {
	return
}

type GrayA128i struct {
	Y, A int64
}

func (c GrayA128i) RGBA() (r, g, b, a uint32) {
	return
}

type GrayA128f struct {
	Y, A float64
}

func (c GrayA128f) RGBA() (r, g, b, a uint32) {
	return
}
