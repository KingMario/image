// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
	"reflect"
)

type Gray16 struct {
	*image.Gray16
}

// NewGray16 returns a new Gray16 with the given bounds.
func NewGray16(r image.Rectangle) *Gray16 {
	return new(Gray16).Init(make([]uint8, 2*r.Dx()*r.Dy()), 2*r.Dx(), r)
}

func (p *Gray16) Init(pix []uint8, stride int, rect image.Rectangle) *Gray16 {
	*p = Gray16{
		Gray16: &image.Gray16{
			Pix:    pix,
			Stride: stride,
			Rect:   rect,
		},
	}
	return p
}

func (p *Gray16) BaseType() image.Image { return p.Gray16 }
func (p *Gray16) Pix() []byte           { return p.Gray16.Pix }
func (p *Gray16) Stride() int           { return p.Gray16.Stride }
func (p *Gray16) Rect() image.Rectangle { return p.Gray16.Rect }
func (p *Gray16) Channels() int         { return 1 }
func (p *Gray16) Depth() reflect.Kind   { return reflect.Uint16 }

func (p *Gray16) CopyFrom(m image.Image) *Gray16 {
	panic("TODO")
}
