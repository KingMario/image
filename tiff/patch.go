// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

// BUG(chai2010): support Gray32f/RGB/RGB48/RGB96f/RGBA128f.

import (
	"image"
	"io"

	imageExt "github.com/chai2010/image"
)

type internalOptions struct {
	Options
}

func (opt *internalOptions) Lossless() bool {
	return true
}

func (opt *internalOptions) Quality() float32 {
	return 0
}

func (opt *Options) Interface() imageExt.Options {
	return &internalOptions{
		Options: *opt,
	}
}

func toOptions(opt imageExt.Options) *Options {
	if opt, ok := opt.(*internalOptions); ok {
		return &opt.Options
	}
	return nil
}

func imageExtEncode(w io.Writer, m image.Image, opt imageExt.Options) error {
	return Encode(w, m, toOptions(opt))
}

func init() {
	imageExt.RegisterFormat(imageExt.Format{
		Name:         "tiff",
		Extensions:   []string{".tiff", ".tif"},
		Magics:       []string{leHeader, beHeader},
		DecodeConfig: DecodeConfig,
		Decode:       Decode,
		Encode:       imageExtEncode,
	})
}
