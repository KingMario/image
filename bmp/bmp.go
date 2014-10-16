// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bmp implements a BMP image decoder and encoder.
//
// The BMP specification is at http://www.digicamsoft.com/bmp/bmp.html.
package bmp

import (
	"image"
	"image/color"
	"io"

	"code.google.com/p/go.image/bmp"
	image_ext "github.com/chai2010/gopkg/image"
	"github.com/chai2010/gopkg/image/convert"
)

// Options are the encoding and decoding parameters.
type Options struct {
	BmpColorModel color.Model
}

func (opt *Options) ColorModel() color.Model {
	if opt != nil {
		return opt.BmpColorModel
	}
	return nil
}

func (opt *Options) Lossless() bool {
	return false
}

func (opt *Options) Quality() float32 {
	return 0
}

// DecodeConfig returns the color model and dimensions of a BMP image without
// decoding the entire image.
// Limitation: The file must be 8 or 24 bits per pixel.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	return bmp.DecodeConfig(r)
}

// Decode reads a BMP image from r and returns it as an image.Image.
// Limitation: The file must be 8 or 24 bits per pixel.
func Decode(r io.Reader, opt *Options) (m image.Image, err error) {
	if m, err = bmp.Decode(r); err != nil {
		return
	}
	if opt != nil && opt.BmpColorModel != nil {
		m = convert.ColorModel(m, opt.BmpColorModel)
	}
	return
}

// Encode writes the image m to w in BMP format.
func Encode(w io.Writer, m image.Image, opt *Options) error {
	if opt != nil && opt.BmpColorModel != nil {
		m = convert.ColorModel(m, opt.BmpColorModel)
	}
	return bmp.Encode(w, m)
}

func toOptions(opt image_ext.Options) *Options {
	if opt, ok := opt.(*Options); ok {
		return opt
	}
	if opt != nil {
		return &Options{
			BmpColorModel: opt.ColorModel(),
		}
	}
	return nil
}

func imageExtDecode(r io.Reader, opt image_ext.Options) (image.Image, error) {
	return Decode(r, toOptions(opt))
}

func imageExtEncode(w io.Writer, m image.Image, opt image_ext.Options) error {
	return Encode(w, m, toOptions(opt))
}

func init() {
	image_ext.RegisterFormat(image_ext.Format{
		Name:         "bmp",
		Extensions:   []string{".bmp"},
		Magics:       []string{"BM????\x00\x00\x00\x00"},
		DecodeConfig: DecodeConfig,
		Decode:       imageExtDecode,
		Encode:       imageExtEncode,
	})
}
