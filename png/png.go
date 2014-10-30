// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package png implements a PNG image decoder and encoder.
//
// The PNG specification is at http://www.w3.org/TR/PNG/.
package png

import (
	"image"
	"image/png"
	"io"

	imageExt "github.com/chai2010/image"
)

const pngHeader = "\x89PNG\r\n\x1a\n"

// DecodeConfig returns the color model and dimensions of a PNG image
// without decoding the entire image.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	return png.DecodeConfig(r)
}

// Decode reads a PNG image from r and returns it as an image.Image.
// The type of Image returned depends on the PNG contents.
func Decode(r io.Reader) (m image.Image, err error) {
	return png.Decode(r)
}

// Encode writes the Image m to w in PNG format.
// Any Image may be encoded, but images that are not image.NRGBA
// might be encoded lossily.
func Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}

func imageExtEncode(w io.Writer, m image.Image, opt imageExt.Options) error {
	return Encode(w, m)
}

func init() {
	imageExt.RegisterFormat(imageExt.Format{
		Name:         "png",
		Extensions:   []string{".png"},
		Magics:       []string{pngHeader},
		DecodeConfig: DecodeConfig,
		Decode:       Decode,
		Encode:       imageExtEncode,
	})
}
