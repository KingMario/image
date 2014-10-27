// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
	"reflect"
)

type RGBA struct {
	*image.RGBA
}

func (p *RGBA) BaseType() image.Image { return p.RGBA }
func (p *RGBA) Pix() []byte           { return p.RGBA.Pix }
func (p *RGBA) Stride() int           { return p.RGBA.Stride }
func (p *RGBA) Rect() image.Rectangle { return p.RGBA.Rect }
func (p *RGBA) Channels() int         { return 4 }
func (p *RGBA) Depth() reflect.Kind   { return reflect.Uint8 }

// NewRGBA returns a new RGBA with the given bounds.
func NewRGBA(r image.Rectangle) *RGBA {
	return new(RGBA).Init(make([]uint8, 4*r.Dx()*r.Dy()), 4*r.Dx(), r)
}

func (p *RGBA) Init(pix []uint8, stride int, rect image.Rectangle) *RGBA {
	*p = RGBA{
		RGBA: &image.RGBA{
			Pix:    pix,
			Stride: stride,
			Rect:   rect,
		},
	}
	return p
}

func (p *RGBA) CopyFrom(m image.Image) *RGBA {
	panic("TODO")
}
