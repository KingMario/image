// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rawp

import (
	"image"
	"reflect"

	image_ext "github.com/chai2010/gopkg/image"
)

func defaultDepthKind(depth int) reflect.Kind {
	switch depth {
	case 8:
		return reflect.Uint8
	case 16:
		return reflect.Uint16
	case 32:
		return reflect.Float32
	}
	return reflect.Uint16
}

func newBytes(size int, buf []byte) []byte {
	if len(buf) >= size {
		return buf[:size]
	}
	return make([]byte, size)
}

func newGray(r image.Rectangle, buf image_ext.Buffer) *image.Gray {
	if buf != nil && r.In(buf.Bounds()) {
		if m, ok := buf.SubImage(r).(*image.Gray); ok {
			return m
		}
	}
	return image.NewGray(r)
}

func newGray16(r image.Rectangle, buf image_ext.Buffer) *image.Gray16 {
	if buf != nil && r.In(buf.Bounds()) {
		if m, ok := buf.SubImage(r).(*image.Gray16); ok {
			return m
		}
	}
	return image.NewGray16(r)
}

func newGray32f(r image.Rectangle, buf image_ext.Buffer) *image_ext.Gray32f {
	if buf != nil && r.In(buf.Bounds()) {
		if m, ok := buf.SubImage(r).(*image_ext.Gray32f); ok {
			return m
		}
	}
	return image_ext.NewGray32f(r)
}

func newRGB(r image.Rectangle, buf image_ext.Buffer) *image_ext.RGB {
	if buf != nil && r.In(buf.Bounds()) {
		if m, ok := buf.SubImage(r).(*image_ext.RGB); ok {
			return m
		}
	}
	return image_ext.NewRGB(r)
}

func newRGB48(r image.Rectangle, buf image_ext.Buffer) *image_ext.RGB48 {
	if buf != nil && r.In(buf.Bounds()) {
		if m, ok := buf.SubImage(r).(*image_ext.RGB48); ok {
			return m
		}
	}
	return image_ext.NewRGB48(r)
}

func newRGB96f(r image.Rectangle, buf image_ext.Buffer) *image_ext.RGB96f {
	if buf != nil && r.In(buf.Bounds()) {
		if m, ok := buf.SubImage(r).(*image_ext.RGB96f); ok {
			return m
		}
	}
	return image_ext.NewRGB96f(r)
}

func newRGBA(r image.Rectangle, buf image_ext.Buffer) *image.RGBA {
	if buf != nil && r.In(buf.Bounds()) {
		if m, ok := buf.SubImage(r).(*image.RGBA); ok {
			return m
		}
	}
	return image.NewRGBA(r)
}

func newRGBA64(r image.Rectangle, buf image_ext.Buffer) *image.RGBA64 {
	if buf != nil && r.In(buf.Bounds()) {
		if m, ok := buf.SubImage(r).(*image.RGBA64); ok {
			return m
		}
	}
	return image.NewRGBA64(r)
}

func newRGBA128f(r image.Rectangle, buf image_ext.Buffer) *image_ext.RGBA128f {
	if buf != nil && r.In(buf.Bounds()) {
		if m, ok := buf.SubImage(r).(*image_ext.RGBA128f); ok {
			return m
		}
	}
	return image_ext.NewRGBA128f(r)
}
