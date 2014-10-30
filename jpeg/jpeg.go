// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package jpeg implements a JPEG image decoder and encoder.
//
// JPEG is defined in ITU-T T.81: http://www.w3.org/Graphics/JPEG/itu-t81.pdf.
package jpeg

import (
	"image"
	"image/jpeg"
	"io"

	imageExt "github.com/chai2010/image"
)

// Options are the encoding and decoding parameters.
type Options struct {
	jpeg.Options
}

func (opt *Options) Lossless() bool {
	return false
}

func (opt *Options) Quality() float32 {
	if opt != nil {
		return float32(opt.Options.Quality)
	}
	return 0
}

// DecodeConfig returns the color model and dimensions of a JPEG image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	return jpeg.DecodeConfig(r)
}

// Decode reads a JPEG image from r and returns it as an image.Image.
func Decode(r io.Reader) (m image.Image, err error) {
	return jpeg.Decode(r)
}

// Encode writes the Image m to w in JPEG 4:2:0 baseline format with the given
// options. Default parameters are used if a nil *Options is passed.
func Encode(w io.Writer, m image.Image, opt *Options) error {
	if opt != nil {
		return jpeg.Encode(w, m, &opt.Options)
	} else {
		return jpeg.Encode(w, m, nil)
	}
}

func toOptions(opt imageExt.Options) *Options {
	if opt, ok := opt.(*Options); ok {
		return opt
	}
	if opt != nil {
		return &Options{
			Options: jpeg.Options{
				Quality: int(opt.Quality()),
			},
		}
	}
	return nil
}

func imageExtDecode(r io.Reader) (image.Image, error) {
	return Decode(r)
}

func imageExtEncode(w io.Writer, m image.Image, opt imageExt.Options) error {
	return Encode(w, m, toOptions(opt))
}

func init() {
	imageExt.RegisterFormat(imageExt.Format{
		Name:         "jpeg",
		Extensions:   []string{".jpeg", ".jpg"},
		Magics:       []string{"\xff\xd8"},
		DecodeConfig: DecodeConfig,
		Decode:       imageExtDecode,
		Encode:       imageExtEncode,
	})
}
