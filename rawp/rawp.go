// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rawp implements a decoder and encoder for RawP images.
//
// RawP Image Structs (Little Endian):
//	type RawPImage struct {
//		Sig          [4]byte // 4Bytes, RAWP
//		Magic        uint32  // 4Bytes, 0x1BF2380A
//		Width        uint16  // 2Bytes, image Width
//		Height       uint16  // 2Bytes, image Height
//		Channels     byte    // 1Bytes, 1=Gray, 2=GrayA, 3=RGB, 4=RGBA
//		Depth        byte    // 1Bytes, 8/16/32/64 bits
//		DataType     byte    // 1Bytes, 1=Uint, 2=Int, 3=Float
//		UseSnappy    byte    // 1Bytes, 0=disabled, 1=enabled (RawPImage.Data)
//		DataSize     uint32  // 4Bytes, image data size (RawPImage.Data)
//		DataCheckSum uint32  // 4Bytes, CRC32(RawPImage.Data[RawPImage.DataSize])
//		Data         []byte  // ?Bytes, image data (RawPImage.DataSize)
//	}
//
// Note: RawP.DataType only support Uin8/Uint16/Int32/Int64/Float32/Float64 formats!!!
//
// Please report bugs to chaishushan{AT}gmail.com.
//
// Thanks!
package rawp

import (
	"fmt"
	"hash/crc32"
	"image/color"
	"math"
	"reflect"
	"unsafe"

	colorExt "github.com/chai2010/image/color"
	"github.com/chai2010/image/rawp/internal/snappy"
)

const (
	rawpHeaderSize = 24
	rawpSig        = "RAWP"
	rawpMagic      = 0x1BF2380A
)

// data type
const (
	rawpDataType_UInt  = 1
	rawpDataType_Int   = 2
	rawpDataType_Float = 3
)

// RawP Image Spec (Little Endian), 24Bytes.
type rawpHeader struct {
	Sig          [4]byte // 4Bytes, WEWP
	Magic        uint32  // 4Bytes, 0x1BF2380A
	Width        uint16  // 2Bytes, image Width
	Height       uint16  // 2Bytes, image Height
	Channels     byte    // 1Bytes, 1=Gray, 3=RGB, 4=RGBA
	Depth        byte    // 1Bytes, 8/16/32/64 bits
	DataType     byte    // 1Bytes, 1=Uint, 2=Int, 3=Float
	UseSnappy    byte    // 1Bytes, 0=disabled, 1=enabled (Header.Data)
	DataSize     uint32  // 4Bytes, image data size (Header.Data)
	DataCheckSum uint32  // 4Bytes, CRC32(RawPHeader.Data[RawPHeader.DataSize])
	Data         []byte  // ?Bytes, image data (RawPHeader.DataSize)
}

func (p *rawpHeader) String() string {
	return fmt.Sprintf(`
image/rawp.rawpHeader{
	Sig:          %q
	Magic:        0x%x
	Width:        %d
	Height:       %d
	Channels:     %d
	Depth:        %d
	DataType:     %d
	UseSnappy:    %d
	DataSize:     %d
	DataCheckSum: 0x%x
	Data:         ?
}
`[1:],
		p.Sig,
		p.Magic,
		p.Width,
		p.Height,
		p.Channels,
		p.Depth,
		p.DataType,
		p.UseSnappy,
		p.DataSize,
		p.DataCheckSum,
	)
}

func rawpIsValidChannels(channels byte) bool {
	return channels >= 1 && channels <= 4
}

func rawpIsValidDepth(depth byte) bool {
	return depth == 8 || depth == 16 || depth == 32 || depth == 64
}

func rawpIsValidDataType(depth, dataType byte) bool {
	switch depth {
	case 8:
		return dataType == rawpDataType_UInt
	case 16:
		return dataType == rawpDataType_UInt
	case 32:
		return dataType == rawpDataType_Int || dataType == rawpDataType_Float
	case 64:
		return dataType == rawpDataType_Int || dataType == rawpDataType_Float
	}
	return false
}

func rawpIsValidHeader(hdr *rawpHeader) error {
	if string(hdr.Sig[:]) != rawpSig {
		return fmt.Errorf("image/rawp: bad Sig, %v", hdr.Sig)
	}
	if hdr.Magic != rawpMagic {
		return fmt.Errorf("image/rawp: bad Magic, %x", hdr.Magic)
	}

	if hdr.Width <= 0 || hdr.Height <= 0 {
		return fmt.Errorf("image/rawp: bad size, width = %v, height = %v", hdr.Width, hdr.Height)
	}
	if !rawpIsValidChannels(hdr.Channels) {
		return fmt.Errorf("image/rawp: bad Channels, %v", hdr.Channels)
	}
	if !rawpIsValidDepth(hdr.Depth) {
		return fmt.Errorf("image/rawp: bad Depth, %v", hdr.Depth)
	}
	if !rawpIsValidDataType(hdr.Depth, hdr.DataType) {
		return fmt.Errorf("image/rawp: bad format, Depth = %v, DataType = %v", hdr.Depth, hdr.DataType)
	}

	if hdr.UseSnappy != 0 && hdr.UseSnappy != 1 {
		return fmt.Errorf("image/rawp: bad UseSnappy, %v", hdr.UseSnappy)
	}
	if hdr.DataSize <= 0 {
		return fmt.Errorf("image/rawp: bad DataSize, %v", hdr.DataSize)
	}

	// check type more ...
	if hdr.Depth == 8 || hdr.Depth == 16 {
		if hdr.DataType == rawpDataType_Float {
			return fmt.Errorf("image/rawp: bad Depth, %v", hdr.Depth)
		}
	}

	// check data size more ...
	if hdr.UseSnappy != 0 {
		n, err := snappy.DecodedLen(hdr.Data)
		if err != nil {
			return fmt.Errorf("image/rawp: snappy.DecodedLen, err = %v", err)
		}
		if x := int(hdr.Width) * int(hdr.Height) * int(hdr.Channels) * int(hdr.Depth) / 8; n != x {
			return fmt.Errorf("image/rawp: snappy.DecodedLen, n = %v", n)
		}
	} else {
		n := int(hdr.DataSize)
		if x := int(hdr.Width) * int(hdr.Height) * int(hdr.Channels) * int(hdr.Depth) / 8; n != x {
			return fmt.Errorf("image/rawp: bad DataSize, %v", hdr.DataSize)
		}
	}

	// Check CRC32
	if v := crc32.ChecksumIEEE(hdr.Data); v != hdr.DataCheckSum {
		return fmt.Errorf("image/rawp: bad DataCheckSum, expect = %x, got = %x", hdr.DataCheckSum, v)
	}

	return nil
}

func rawpColorModel(hdr *rawpHeader) (color.Model, error) {
	switch {
	case hdr.Channels == 1:
		switch {
		case hdr.Depth == 8 && hdr.DataType == rawpDataType_UInt:
			return colorExt.GrayModel, nil
		case hdr.Depth == 16 && hdr.DataType == rawpDataType_UInt:
			return colorExt.Gray16Model, nil
		case hdr.Depth == 32 && hdr.DataType == rawpDataType_Int:
			return colorExt.Gray32iModel, nil
		case hdr.Depth == 32 && hdr.DataType == rawpDataType_Float:
			return colorExt.Gray32fModel, nil
		case hdr.Depth == 64 && hdr.DataType == rawpDataType_Int:
			return colorExt.Gray64iModel, nil
		case hdr.Depth == 64 && hdr.DataType == rawpDataType_Float:
			return colorExt.Gray64fModel, nil
		}
	case hdr.Channels == 2:
		switch {
		case hdr.Depth == 8 && hdr.DataType == rawpDataType_UInt:
			return colorExt.GrayAModel, nil
		case hdr.Depth == 16 && hdr.DataType == rawpDataType_UInt:
			return colorExt.GrayA32Model, nil
		case hdr.Depth == 32 && hdr.DataType == rawpDataType_Int:
			return colorExt.GrayA64iModel, nil
		case hdr.Depth == 32 && hdr.DataType == rawpDataType_Float:
			return colorExt.GrayA64fModel, nil
		case hdr.Depth == 64 && hdr.DataType == rawpDataType_Int:
			return colorExt.GrayA128iModel, nil
		case hdr.Depth == 64 && hdr.DataType == rawpDataType_Float:
			return colorExt.GrayA128fModel, nil
		}
	case hdr.Channels == 3:
		switch {
		case hdr.Depth == 8 && hdr.DataType == rawpDataType_UInt:
			return colorExt.RGBModel, nil
		case hdr.Depth == 16 && hdr.DataType == rawpDataType_UInt:
			return colorExt.RGB48Model, nil
		case hdr.Depth == 32 && hdr.DataType == rawpDataType_Int:
			return colorExt.RGB96iModel, nil
		case hdr.Depth == 32 && hdr.DataType == rawpDataType_Float:
			return colorExt.RGB96fModel, nil
		case hdr.Depth == 64 && hdr.DataType == rawpDataType_Int:
			return colorExt.RGB192iModel, nil
		case hdr.Depth == 64 && hdr.DataType == rawpDataType_Float:
			return colorExt.RGB192fModel, nil
		}
	case hdr.Channels == 4:
		switch {
		case hdr.Depth == 8 && hdr.DataType == rawpDataType_UInt:
			return color.RGBAModel, nil
		case hdr.Depth == 16 && hdr.DataType == rawpDataType_UInt:
			return color.RGBA64Model, nil
		case hdr.Depth == 32 && hdr.DataType == rawpDataType_Int:
			return colorExt.RGBA128iModel, nil
		case hdr.Depth == 32 && hdr.DataType == rawpDataType_Float:
			return colorExt.RGBA128fModel, nil
		case hdr.Depth == 64 && hdr.DataType == rawpDataType_Int:
			return colorExt.RGBA256iModel, nil
		case hdr.Depth == 64 && hdr.DataType == rawpDataType_Float:
			return colorExt.RGBA256fModel, nil
		}
	}
	return nil, fmt.Errorf("image/rawp: unsupport color model, hdr = %v", hdr)
}

func rawpPixDecoder(hdr *rawpHeader) (decoder *pixDecoder, err error) {
	switch {
	case hdr.Channels == 1:
		switch {
		case hdr.Depth == 8 && hdr.DataType == rawpDataType_UInt:
			decoder = &pixDecoder{
				Channels: int(hdr.Channels),
				DataType: reflect.Uint8,
				Width:    int(hdr.Width),
				Height:   int(hdr.Height),
			}
			return
		case hdr.Depth == 16 && hdr.DataType == rawpDataType_UInt:
			decoder = &pixDecoder{
				Channels: int(hdr.Channels),
				DataType: reflect.Uint16,
				Width:    int(hdr.Width),
				Height:   int(hdr.Height),
			}
			return
		case hdr.Depth == 32 && hdr.DataType == rawpDataType_Float:
			decoder = &pixDecoder{
				Channels: int(hdr.Channels),
				DataType: reflect.Float32,
				Width:    int(hdr.Width),
				Height:   int(hdr.Height),
			}
			return
		}
	case hdr.Channels == 3:
		switch {
		case hdr.Depth == 8 && hdr.DataType == rawpDataType_UInt:
			decoder = &pixDecoder{
				Channels: int(hdr.Channels),
				DataType: reflect.Uint8,
				Width:    int(hdr.Width),
				Height:   int(hdr.Height),
			}
			return
		case hdr.Depth == 16 && hdr.DataType == rawpDataType_UInt:
			decoder = &pixDecoder{
				Channels: int(hdr.Channels),
				DataType: reflect.Uint16,
				Width:    int(hdr.Width),
				Height:   int(hdr.Height),
			}
			return
		case hdr.Depth == 32 && hdr.DataType == rawpDataType_Float:
			decoder = &pixDecoder{
				Channels: int(hdr.Channels),
				DataType: reflect.Float32,
				Width:    int(hdr.Width),
				Height:   int(hdr.Height),
			}
			return
		}
	case hdr.Channels == 4:
		switch {
		case hdr.Depth == 8 && hdr.DataType == rawpDataType_UInt:
			decoder = &pixDecoder{
				Channels: int(hdr.Channels),
				DataType: reflect.Uint8,
				Width:    int(hdr.Width),
				Height:   int(hdr.Height),
			}
			return
		case hdr.Depth == 16 && hdr.DataType == rawpDataType_UInt:
			decoder = &pixDecoder{
				Channels: int(hdr.Channels),
				DataType: reflect.Uint16,
				Width:    int(hdr.Width),
				Height:   int(hdr.Height),
			}
			return
		case hdr.Depth == 32 && hdr.DataType == rawpDataType_Float:
			decoder = &pixDecoder{
				Channels: int(hdr.Channels),
				DataType: reflect.Float32,
				Width:    int(hdr.Width),
				Height:   int(hdr.Height),
			}
			return
		}
	}
	return nil, fmt.Errorf("image/rawp: unsupport color model, hdr = %v", hdr)
}

func rawpPixEncoder(hdr *rawpHeader) (encoder *pixEncoder, err error) {
	switch {
	case hdr.Channels == 1:
		switch {
		case hdr.Depth == 8 && hdr.DataType == rawpDataType_UInt:
			encoder = &pixEncoder{
				Channels: int(hdr.Channels),
				DataType: reflect.Uint8,
			}
			return
		case hdr.Depth == 16 && hdr.DataType == rawpDataType_UInt:
			encoder = &pixEncoder{
				Channels: int(hdr.Channels),
				DataType: reflect.Uint16,
			}
			return
		case hdr.Depth == 32 && hdr.DataType == rawpDataType_Float:
			encoder = &pixEncoder{
				Channels: int(hdr.Channels),
				DataType: reflect.Float32,
			}
			return
		}
	case hdr.Channels == 3:
		switch {
		case hdr.Depth == 8 && hdr.DataType == rawpDataType_UInt:
			encoder = &pixEncoder{
				Channels: int(hdr.Channels),
				DataType: reflect.Uint8,
			}
			return
		case hdr.Depth == 16 && hdr.DataType == rawpDataType_UInt:
			encoder = &pixEncoder{
				Channels: int(hdr.Channels),
				DataType: reflect.Uint16,
			}
			return
		case hdr.Depth == 32 && hdr.DataType == rawpDataType_Float:
			encoder = &pixEncoder{
				Channels: int(hdr.Channels),
				DataType: reflect.Float32,
			}
			return
		}
	case hdr.Channels == 4:
		switch {
		case hdr.Depth == 8 && hdr.DataType == rawpDataType_UInt:
			encoder = &pixEncoder{
				Channels: int(hdr.Channels),
				DataType: reflect.Uint8,
			}
			return
		case hdr.Depth == 16 && hdr.DataType == rawpDataType_UInt:
			encoder = &pixEncoder{
				Channels: int(hdr.Channels),
				DataType: reflect.Uint16,
			}
			return
		case hdr.Depth == 32 && hdr.DataType == rawpDataType_Float:
			encoder = &pixEncoder{
				Channels: int(hdr.Channels),
				DataType: reflect.Float32,
			}
			return
		}
	}
	return nil, fmt.Errorf("image/rawp: unsupport color model, hdr = %v", hdr)
}

func rawpMakeHeader(width, height int, model color.Model, useSnappy bool) (hdr *rawpHeader, err error) {
	if width <= 0 || width > math.MaxUint16 {
		err = fmt.Errorf("image/rawp: image size overflow: width = %v, height = %v", width, height)
		return
	}
	if height <= 0 || height > math.MaxUint16 {
		err = fmt.Errorf("image/rawp: image size overflow: width = %v, height = %v", width, height)
		return
	}

	hdr = &rawpHeader{
		Sig:    [4]byte{'R', 'A', 'W', 'P'},
		Magic:  rawpMagic,
		Width:  uint16(width),
		Height: uint16(height),
	}
	if useSnappy {
		hdr.UseSnappy = 1
	}

	switch model {
	case color.GrayModel:
		hdr.Channels = 1
		hdr.Depth = 8
		hdr.DataType = rawpDataType_UInt
		return
	case color.Gray16Model:
		hdr.Channels = 1
		hdr.Depth = 16
		hdr.DataType = rawpDataType_UInt
		return
	case color_ext.Gray32fModel:
		hdr.Channels = 1
		hdr.Depth = 32
		hdr.DataType = rawpDataType_Float
		return
	case color_ext.RGBModel:
		hdr.Channels = 3
		hdr.Depth = 8
		hdr.DataType = rawpDataType_UInt
		return
	case color_ext.RGB48Model:
		hdr.Channels = 3
		hdr.Depth = 16
		hdr.DataType = rawpDataType_UInt
		return
	case color_ext.RGB96fModel:
		hdr.Channels = 3
		hdr.Depth = 32
		hdr.DataType = rawpDataType_Float
		return
	case color.RGBAModel:
		hdr.Channels = 4
		hdr.Depth = 8
		hdr.DataType = rawpDataType_UInt
		return
	case color.RGBA64Model:
		hdr.Channels = 4
		hdr.Depth = 16
		hdr.DataType = rawpDataType_UInt
		return
	case color_ext.RGBA128fModel:
		hdr.Channels = 4
		hdr.Depth = 32
		hdr.DataType = rawpDataType_Float
		return
	}
	return nil, fmt.Errorf("image/rawp: unsupport color model, %T", model)
}

func rawpDecodeHeader(data []byte) (hdr *rawpHeader, err error) {
	if len(data) < rawpHeaderSize {
		err = fmt.Errorf("image/rawp: bad header.")
		return
	}

	// reader header
	hdr = new(rawpHeader)
	copy(((*[1 << 30]byte)(unsafe.Pointer(hdr)))[:rawpHeaderSize], data)
	hdr.Data = data[rawpHeaderSize:]

	// check header
	if err = rawpIsValidHeader(hdr); err != nil {
		return
	}
	return
}
