// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gif implements a GIF image decoder and encoder.
//
// The GIF specification is at http://www.w3.org/Graphics/GIF/spec-gif89a.txt.
package gif

import (
	"image"
	"image/color"
	"image/gif"
	"io"

	image_ext "github.com/chai2010/gopkg/image"
	"github.com/chai2010/gopkg/image/convert"
)

// Options are the encoding and decoding parameters.
type Options struct {
	*gif.Options
	GifColorModel color.Model
}

func (opt *Options) ColorModel() color.Model {
	if opt != nil {
		return opt.GifColorModel
	}
	return nil
}

func (opt *Options) Lossless() bool {
	return false
}

func (opt *Options) Quality() float32 {
	return 0
}

// DecodeConfig returns the global color model and dimensions of a GIF image
// without decoding the entire image.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	return gif.DecodeConfig(r)
}

// Decode reads a GIF image from r and returns the first embedded
// image as an image.Image.
func Decode(r io.Reader, opt *Options) (m image.Image, err error) {
	if m, err = gif.Decode(r); err != nil {
		return
	}
	if opt != nil && opt.GifColorModel != nil {
		m = convert.ColorModel(m, opt.GifColorModel)
	}
	return
}

// DecodeAll reads a GIF image from r and returns the sequential frames
// and timing information.
func DecodeAll(r io.Reader) (*gif.GIF, error) {
	return gif.DecodeAll(r)
}

// EncodeAll writes the images in g to w in GIF format with the
// given loop count and delay between frames.
func EncodeAll(w io.Writer, g *gif.GIF) error {
	return gif.EncodeAll(w, g)
}

// Encode writes the Image m to w in GIF format.
func Encode(w io.Writer, m image.Image, opt *Options) error {
	if opt != nil && opt.GifColorModel != nil {
		m = convert.ColorModel(m, opt.GifColorModel)
	}
	if opt != nil && opt.Options != nil {
		return gif.Encode(w, m, opt.Options)
	} else {
		return gif.Encode(w, m, nil)
	}
}

func toOptions(opt image_ext.Options) *Options {
	if opt, ok := opt.(*Options); ok {
		return opt
	}
	if opt != nil {
		return &Options{
			GifColorModel: opt.ColorModel(),
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
		Name:         "gif",
		Extensions:   []string{".gif"},
		Magics:       []string{"GIF8?a"},
		DecodeConfig: DecodeConfig,
		Decode:       imageExtDecode,
		Encode:       imageExtEncode,
	})
}
