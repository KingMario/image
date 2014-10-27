// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
	"reflect"
)

type Gray struct {
	*image.Gray
}

// NewGray returns a new Gray with the given bounds.
func NewGray(r image.Rectangle) *Gray {
	return new(Gray).Init(make([]uint8, 1*r.Dx()*r.Dy()), 1*r.Dx(), r)
}

func (p *Gray) Init(pix []uint8, stride int, rect image.Rectangle) *Gray {
	*p = Gray{
		Gray: &image.Gray{
			Pix:    pix,
			Stride: stride,
			Rect:   rect,
		},
	}
	return p
}

func (p *Gray) BaseType() image.Image { return p.Gray }
func (p *Gray) Pix() []byte           { return p.Gray.Pix }
func (p *Gray) Stride() int           { return p.Gray.Stride }
func (p *Gray) Rect() image.Rectangle { return p.Gray.Rect }
func (p *Gray) Channels() int         { return 1 }
func (p *Gray) Depth() reflect.Kind   { return reflect.Uint8 }

func (p *Gray) CopyFrom(m image.Image) *Gray {
	panic("TODO")
}
