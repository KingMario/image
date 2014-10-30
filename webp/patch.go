// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

import (
	"image"
	"io"

	imageExt "github.com/chai2010/image"
)

type internalOptions struct {
	Options
}

func (opt *internalOptions) Lossless() bool {
	return opt.Options.Lossless
}

func (opt *internalOptions) Quality() float32 {
	return opt.Options.Quality
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
	if opt != nil {
		return &Options{
			Lossless: opt.Lossless(),
			Quality:  opt.Quality(),
		}
	}
	return nil
}

func imageExtEncode(w io.Writer, m image.Image, opt imageExt.Options) error {
	return Encode(w, m, toOptions(opt))
}

func init() {
	imageExt.RegisterFormat(imageExt.Format{
		Name:         "webp",
		Extensions:   []string{".webp"},
		Magics:       []string{"RIFF????WEBPVP8"},
		DecodeConfig: DecodeConfig,
		Decode:       Decode,
		Encode:       imageExtEncode,
	})
}
