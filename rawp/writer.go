// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rawp

import (
	"hash/crc32"
	"image"
	"image/color"
	"io"
	"unsafe"

	"code.google.com/p/snappy-go/snappy"
	image_ext "github.com/chai2010/gopkg/image"
	"github.com/chai2010/gopkg/image/convert"
)

// Encode writes the image m to w in RawP format.
func Encode(w io.Writer, m image.Image, opt *Options) (err error) {
	if opt != nil && opt.RawPColorModel != nil {
		m = convert.ColorModel(m, opt.RawPColorModel)
	}
	m = adjustImage(m)

	var useSnappy bool
	if opt != nil {
		useSnappy = opt.UseSnappy
	}

	hdr, err := rawpMakeHeader(m.Bounds().Dx(), m.Bounds().Dy(), m.ColorModel(), useSnappy)
	if err != nil {
		return
	}

	// encode raw pix
	encoder, err := rawpPixEncoder(hdr)
	if err != nil {
		return
	}
	pix, err := encoder.Encode(m, nil)
	if err != nil {
		return
	}
	if useSnappy {
		pix, err = snappy.Encode(nil, pix)
		if err != nil {
			return
		}
	}

	hdr.DataSize = uint32(len(pix))
	hdr.DataCheckSum = crc32.ChecksumIEEE(pix)
	hdr.Data = pix

	if _, err = w.Write(((*[1 << 30]byte)(unsafe.Pointer(hdr)))[:rawpHeaderSize]); err != nil {
		return
	}
	if _, err = w.Write(hdr.Data); err != nil {
		return
	}
	return
}

func adjustImage(m image.Image) image.Image {
	switch m := m.(type) {
	case *image.Gray, *image.Gray16, *image_ext.Gray32f:
		return m
	case *image_ext.RGB, *image_ext.RGB48, *image_ext.RGB96f:
		return m
	case *image.RGBA, *image.RGBA64, *image_ext.RGBA128f:
		return m
	default:
		b := m.Bounds()
		rgba := image.NewRGBA(b)
		dstColorRGBA64 := &color.RGBA64{}
		dstColor := color.Color(dstColorRGBA64)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				pr, pg, pb, pa := m.At(x, y).RGBA()
				dstColorRGBA64.R = uint16(pr)
				dstColorRGBA64.G = uint16(pg)
				dstColorRGBA64.B = uint16(pb)
				dstColorRGBA64.A = uint16(pa)
				rgba.Set(x, y, dstColor)
			}
		}
		return rgba
	}
}
