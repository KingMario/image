// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
	"reflect"
)

type RGBA64 struct {
	*image.RGBA64
}

func (p *RGBA64) BaseType() image.Image { return p.RGBA64 }
func (p *RGBA64) Pix() []byte           { return p.RGBA64.Pix }
func (p *RGBA64) Stride() int           { return p.RGBA64.Stride }
func (p *RGBA64) Rect() image.Rectangle { return p.RGBA64.Rect }
func (p *RGBA64) Channels() int         { return 4 }
func (p *RGBA64) Depth() reflect.Kind   { return reflect.Uint16 }

// NewRGBA64 returns a new RGBA64 with the given bounds.
func NewRGBA64(r image.Rectangle) *RGBA64 {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 8*w*h)
	return &RGBA64{
		RGBA64: &image.RGBA64{
			Pix:    pix,
			Stride: 8 * w,
			Rect:   r,
		},
	}
}

func (p *RGBA64) Init(pix []uint8, stride int, rect image.Rectangle) *RGBA64 {
	*p = RGBA64{
		RGBA64: &image.RGBA64{
			Pix:    pix,
			Stride: stride,
			Rect:   rect,
		},
	}
	return p
}

func (p *RGBA64) CopyFrom(m image.Image) *RGBA64 {
	panic("TODO")
}
