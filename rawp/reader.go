// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rawp

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"io/ioutil"

	imageExt "github.com/chai2010/image"
	"github.com/chai2010/image/rawp/internal/snappy"
)

// Options are the encoding and decoding parameters.
type Options struct {
	RawPColorModel color.Model
	UseSnappy      bool
}

func (opt *Options) ColorModel() color.Model {
	if opt != nil {
		return opt.RawPColorModel
	}
	return nil
}

func (opt *Options) Lossless() bool {
	return false
}

func (opt *Options) Quality() float32 {
	return 0
}

// DecodeConfig returns the color model and dimensions of a RawP image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (config image.Config, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	hdr, err := rawpDecodeHeader(data)
	if err != nil {
		return
	}

	model, err := rawpColorModel(hdr)
	if err != nil {
		return
	}

	config = image.Config{
		ColorModel: model,
		Width:      int(hdr.Width),
		Height:     int(hdr.Height),
	}
	return
}

// Decode reads a RawP image from r and returns it as an image.Image.
// The type of Image returned depends on the contents of the RawP.
func Decode(r io.Reader, opt *Options) (m image.Image, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	hdr, err := rawpDecodeHeader(data)
	if err != nil {
		return
	}

	// new decoder
	decoder, err := rawpPixDecoder(hdr)
	if err != nil {
		return
	}

	// decode snappy
	pix := hdr.Data
	if hdr.UseSnappy != 0 {
		if pix, err = snappy.Decode(nil, hdr.Data); err != nil {
			err = fmt.Errorf("image/rawp: Decode, snappy err: %v", err)
			return
		}
	}

	// decode raw pix
	m, err = decoder.Decode(pix, nil)
	if err != nil {
		return
	}

	// convert color model
	if opt != nil && opt.RawPColorModel != nil {
		m = convert.ColorModel(m, opt.RawPColorModel)
	}

	return
}

func toOptions(opt imageExt.Options) *Options {
	if opt, ok := opt.(*Options); ok {
		return opt
	}
	if opt != nil {
		return &Options{
			RawPColorModel: opt.ColorModel(),
		}
	}
	return nil
}

func imageDecode(r io.Reader) (image.Image, error) {
	return Decode(r, nil)
}

func imageExtDecode(r io.Reader, opt imageExt.Options) (image.Image, error) {
	return Decode(r, toOptions(opt))
}

func imageExtEncode(w io.Writer, m image.Image, opt imageExt.Options) error {
	return Encode(w, m, toOptions(opt))
}

func init() {
	image.RegisterFormat("rawp", "RAWP\x1B\xF2\x38\x0A", imageDecode, DecodeConfig)

	imageExt.RegisterFormat(imageExt.Format{
		Name:         "rawp",
		Extensions:   []string{".rawp"},
		Magics:       []string{"RAWP\x0A\x38\xF2\x1B"}, // rawSig + rawpMagic(Little Endian)
		DecodeConfig: DecodeConfig,
		Decode:       imageExtDecode,
		Encode:       imageExtEncode,
	})
}
