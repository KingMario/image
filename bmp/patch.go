// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bmp implements a BMP image decoder and encoder.
//
// The BMP specification is at http://www.digicamsoft.com/bmp/bmp.html.
package bmp

import (
	"image"
	"io"

	imageExt "github.com/chai2010/image"
)

func imageExtEncode(w io.Writer, m image.Image, opt imageExt.Options) error {
	return Encode(w, m)
}

func init() {
	imageExt.RegisterFormat(imageExt.Format{
		Name:         "bmp",
		Extensions:   []string{".bmp"},
		Magics:       []string{"BM????\x00\x00\x00\x00"},
		DecodeConfig: DecodeConfig,
		Decode:       Decode,
		Encode:       imageExtEncode,
	})
}
