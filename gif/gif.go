// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gif implements a GIF image decoder and encoder.
//
// The GIF specification is at http://www.w3.org/Graphics/GIF/spec-gif89a.txt.
package gif

import (
	"image"
	"image/gif"
	"io"

	imageExt "github.com/chai2010/image"
)

// Options are the encoding and decoding parameters.
type Options struct {
	gif.Options
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
func Decode(r io.Reader) (m image.Image, err error) {
	return gif.Decode(r)
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
	if opt != nil {
		return gif.Encode(w, m, &opt.Options)
	} else {
		return gif.Encode(w, m, nil)
	}
}

func toOptions(opt imageExt.Options) *Options {
	if opt, ok := opt.(*Options); ok {
		return opt
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
		Name:         "gif",
		Extensions:   []string{".gif"},
		Magics:       []string{"GIF8?a"},
		DecodeConfig: DecodeConfig,
		Decode:       imageExtDecode,
		Encode:       imageExtEncode,
	})
}
