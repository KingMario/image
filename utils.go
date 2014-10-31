// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"encoding/binary"
	"math"

	colorExt "github.com/chai2010/image/color"
)

func pGrayAt(pix []byte) colorExt.Gray {
	return colorExt.Gray{
		Y: pix[1*0],
	}
}
func pSetGray(pix []byte, c colorExt.Gray) {
	pix[1*0] = c.Y
}

func pGray16At(pix []byte) colorExt.Gray16 {
	return colorExt.Gray16{
		Y: binary.BigEndian.Uint16(pix[2*0:]),
	}
}
func pSetGray16(pix []byte, c colorExt.Gray16) {
	binary.BigEndian.PutUint16(pix[2*0:], c.Y)
}

func pGray32iAt(pix []byte) colorExt.Gray32i {
	return colorExt.Gray32i{
		Y: int32(binary.BigEndian.Uint32(pix[4*0:])),
	}
}
func pSetGray32i(pix []byte, c colorExt.Gray32i) {
	binary.BigEndian.PutUint32(pix[4*0:], uint32(c.Y))
}

func pGray32fAt(pix []byte) colorExt.Gray32f {
	return colorExt.Gray32f{
		Y: math.Float32frombits(binary.BigEndian.Uint32(pix[4*0:])),
	}
}
func pSetGray32f(pix []byte, c colorExt.Gray32f) {
	binary.BigEndian.PutUint32(pix[4*0:], math.Float32bits(c.Y))
}

func pGray64iAt(pix []byte) colorExt.Gray64i {
	return colorExt.Gray64i{
		Y: int64(binary.BigEndian.Uint64(pix[8*0:])),
	}
}
func pSetGray64i(pix []byte, c colorExt.Gray64i) {
	binary.BigEndian.PutUint64(pix[8*0:], uint64(c.Y))
}

func pGray64fAt(pix []byte) colorExt.Gray64f {
	return colorExt.Gray64f{
		Y: math.Float64frombits(binary.BigEndian.Uint64(pix[8*0:])),
	}
}
func pSetGray64f(pix []byte, c colorExt.Gray64f) {
	binary.BigEndian.PutUint64(pix[8*0:], math.Float64bits(c.Y))
}

func pGrayAAt(pix []byte) colorExt.GrayA {
	return colorExt.GrayA{
		Y: pix[1*0],
		A: pix[1*1],
	}
}
func pSetGrayA(pix []byte, c colorExt.GrayA) {
	pix[1*0] = c.Y
	pix[1*1] = c.A
}

func pGrayA32At(pix []byte) colorExt.GrayA32 {
	return colorExt.GrayA32{
		Y: binary.BigEndian.Uint16(pix[2*0:]),
		A: binary.BigEndian.Uint16(pix[2*1:]),
	}
}
func pSetGrayA32(pix []byte, c colorExt.GrayA32) {
	binary.BigEndian.PutUint16(pix[2*0:], c.Y)
	binary.BigEndian.PutUint16(pix[2*1:], c.A)
}

func pGrayA64iAt(pix []byte) colorExt.GrayA64i {
	return colorExt.GrayA64i{
		Y: int32(binary.BigEndian.Uint32(pix[4*0:])),
		A: int32(binary.BigEndian.Uint32(pix[4*1:])),
	}
}
func pSetGrayA64i(pix []byte, c colorExt.GrayA64i) {
	binary.BigEndian.PutUint32(pix[4*0:], uint32(c.Y))
	binary.BigEndian.PutUint32(pix[4*1:], uint32(c.A))
}

func pGrayA64fAt(pix []byte) colorExt.GrayA64f {
	return colorExt.GrayA64f{
		Y: math.Float32frombits(binary.BigEndian.Uint32(pix[4*0:])),
		A: math.Float32frombits(binary.BigEndian.Uint32(pix[4*1:])),
	}
}
func pSetGrayA64f(pix []byte, c colorExt.GrayA64f) {
	binary.BigEndian.PutUint32(pix[4*0:], math.Float32bits(c.Y))
	binary.BigEndian.PutUint32(pix[4*1:], math.Float32bits(c.A))
}

func pGrayA128iAt(pix []byte) colorExt.GrayA128i {
	return colorExt.GrayA128i{
		Y: int64(binary.BigEndian.Uint64(pix[8*0:])),
		A: int64(binary.BigEndian.Uint64(pix[8*1:])),
	}
}
func pSetGrayA128i(pix []byte, c colorExt.GrayA128i) {
	binary.BigEndian.PutUint64(pix[8*0:], uint64(c.Y))
	binary.BigEndian.PutUint64(pix[8*1:], uint64(c.A))
}

func pGrayA128fAt(pix []byte) colorExt.GrayA128f {
	return colorExt.GrayA128f{
		Y: math.Float64frombits(binary.BigEndian.Uint64(pix[8*0:])),
		A: math.Float64frombits(binary.BigEndian.Uint64(pix[8*1:])),
	}
}
func pSetGrayA128f(pix []byte, c colorExt.GrayA128f) {
	binary.BigEndian.PutUint64(pix[8*0:], math.Float64bits(c.Y))
	binary.BigEndian.PutUint64(pix[8*1:], math.Float64bits(c.A))
}

func pRGBAt(pix []byte) colorExt.RGB {
	return colorExt.RGB{
		R: pix[1*0],
		G: pix[1*1],
		B: pix[1*2],
	}
}
func pSetRGB(pix []byte, c colorExt.RGB) {
	pix[1*0] = c.R
	pix[1*1] = c.G
	pix[1*2] = c.B
}

func pRGB48At(pix []byte) colorExt.RGB48 {
	return colorExt.RGB48{
		R: binary.BigEndian.Uint16(pix[2*0:]),
		G: binary.BigEndian.Uint16(pix[2*1:]),
		B: binary.BigEndian.Uint16(pix[2*2:]),
	}
}
func pSetRGB48(pix []byte, c colorExt.RGB48) {
	binary.BigEndian.PutUint16(pix[2*0:], c.R)
	binary.BigEndian.PutUint16(pix[2*1:], c.G)
	binary.BigEndian.PutUint16(pix[2*2:], c.B)
}

func pRGB96iAt(pix []byte) colorExt.RGB96i {
	return colorExt.RGB96i{
		R: int32(binary.BigEndian.Uint32(pix[4*0:])),
		G: int32(binary.BigEndian.Uint32(pix[4*1:])),
		B: int32(binary.BigEndian.Uint32(pix[4*2:])),
	}
}
func pSetRGB96i(pix []byte, c colorExt.RGB96i) {
	binary.BigEndian.PutUint32(pix[4*0:], uint32(c.R))
	binary.BigEndian.PutUint32(pix[4*1:], uint32(c.G))
	binary.BigEndian.PutUint32(pix[4*2:], uint32(c.B))
}

func pRGB96fAt(pix []byte) colorExt.RGB96f {
	return colorExt.RGB96f{
		R: math.Float32frombits(binary.BigEndian.Uint32(pix[4*0:])),
		G: math.Float32frombits(binary.BigEndian.Uint32(pix[4*1:])),
		B: math.Float32frombits(binary.BigEndian.Uint32(pix[4*2:])),
	}
}
func pSetRGB96f(pix []byte, c colorExt.RGB96f) {
	binary.BigEndian.PutUint32(pix[4*0:], math.Float32bits(c.R))
	binary.BigEndian.PutUint32(pix[4*1:], math.Float32bits(c.G))
	binary.BigEndian.PutUint32(pix[4*2:], math.Float32bits(c.B))
}

func pRGB192iAt(pix []byte) colorExt.RGB192i {
	return colorExt.RGB192i{
		R: int64(binary.BigEndian.Uint64(pix[8*0:])),
		G: int64(binary.BigEndian.Uint64(pix[8*1:])),
		B: int64(binary.BigEndian.Uint64(pix[8*2:])),
	}
}
func pSetRGB192i(pix []byte, c colorExt.RGB192i) {
	binary.BigEndian.PutUint64(pix[8*0:], uint64(c.R))
	binary.BigEndian.PutUint64(pix[8*1:], uint64(c.G))
	binary.BigEndian.PutUint64(pix[8*2:], uint64(c.B))
}

func pRGB192fAt(pix []byte) colorExt.RGB192f {
	return colorExt.RGB192f{
		R: math.Float64frombits(binary.BigEndian.Uint64(pix[8*0:])),
		G: math.Float64frombits(binary.BigEndian.Uint64(pix[8*1:])),
		B: math.Float64frombits(binary.BigEndian.Uint64(pix[8*2:])),
	}
}
func pSetRGB192f(pix []byte, c colorExt.RGB192f) {
	binary.BigEndian.PutUint64(pix[8*0:], math.Float64bits(c.R))
	binary.BigEndian.PutUint64(pix[8*1:], math.Float64bits(c.G))
	binary.BigEndian.PutUint64(pix[8*2:], math.Float64bits(c.B))
}

func pRGBAAt(pix []byte) colorExt.RGBA {
	return colorExt.RGBA{
		R: pix[1*0],
		G: pix[1*1],
		B: pix[1*2],
		A: pix[1*3],
	}
}
func pSetRGBA(pix []byte, c colorExt.RGBA) {
	pix[1*0] = c.R
	pix[1*1] = c.G
	pix[1*2] = c.B
	pix[1*3] = c.A
}

func pRGBA64At(pix []byte) colorExt.RGBA64 {
	return colorExt.RGBA64{
		R: binary.BigEndian.Uint16(pix[2*0:]),
		G: binary.BigEndian.Uint16(pix[2*1:]),
		B: binary.BigEndian.Uint16(pix[2*2:]),
		A: binary.BigEndian.Uint16(pix[2*3:]),
	}
}
func pSetRGBA64(pix []byte, c colorExt.RGBA64) {
	binary.BigEndian.PutUint16(pix[2*0:], c.R)
	binary.BigEndian.PutUint16(pix[2*1:], c.G)
	binary.BigEndian.PutUint16(pix[2*2:], c.B)
	binary.BigEndian.PutUint16(pix[2*3:], c.A)
}

func pRGBA128iAt(pix []byte) colorExt.RGBA128i {
	return colorExt.RGBA128i{
		R: int32(binary.BigEndian.Uint32(pix[4*0:])),
		G: int32(binary.BigEndian.Uint32(pix[4*1:])),
		B: int32(binary.BigEndian.Uint32(pix[4*2:])),
		A: int32(binary.BigEndian.Uint32(pix[4*3:])),
	}
}
func pSetRGBA128i(pix []byte, c colorExt.RGBA128i) {
	binary.BigEndian.PutUint32(pix[4*0:], uint32(c.R))
	binary.BigEndian.PutUint32(pix[4*1:], uint32(c.G))
	binary.BigEndian.PutUint32(pix[4*2:], uint32(c.B))
	binary.BigEndian.PutUint32(pix[4*3:], uint32(c.A))
}

func pRGBA128fAt(pix []byte) colorExt.RGBA128f {
	return colorExt.RGBA128f{
		R: math.Float32frombits(binary.BigEndian.Uint32(pix[4*0:])),
		G: math.Float32frombits(binary.BigEndian.Uint32(pix[4*1:])),
		B: math.Float32frombits(binary.BigEndian.Uint32(pix[4*2:])),
		A: math.Float32frombits(binary.BigEndian.Uint32(pix[4*3:])),
	}
}
func pSetRGBA128f(pix []byte, c colorExt.RGBA128f) {
	binary.BigEndian.PutUint32(pix[4*0:], math.Float32bits(c.R))
	binary.BigEndian.PutUint32(pix[4*1:], math.Float32bits(c.G))
	binary.BigEndian.PutUint32(pix[4*2:], math.Float32bits(c.B))
	binary.BigEndian.PutUint32(pix[4*3:], math.Float32bits(c.A))
}

func pRGBA256iAt(pix []byte) colorExt.RGBA256i {
	return colorExt.RGBA256i{
		R: int64(binary.BigEndian.Uint64(pix[8*0:])),
		G: int64(binary.BigEndian.Uint64(pix[8*1:])),
		B: int64(binary.BigEndian.Uint64(pix[8*2:])),
		A: int64(binary.BigEndian.Uint64(pix[8*3:])),
	}
}
func pSetRGBA256i(pix []byte, c colorExt.RGBA256i) {
	binary.BigEndian.PutUint64(pix[8*0:], uint64(c.R))
	binary.BigEndian.PutUint64(pix[8*1:], uint64(c.G))
	binary.BigEndian.PutUint64(pix[8*2:], uint64(c.B))
	binary.BigEndian.PutUint64(pix[8*3:], uint64(c.A))
}

func pRGBA256fAt(pix []byte) colorExt.RGBA256f {
	return colorExt.RGBA256f{
		R: math.Float64frombits(binary.BigEndian.Uint64(pix[8*0:])),
		G: math.Float64frombits(binary.BigEndian.Uint64(pix[8*1:])),
		B: math.Float64frombits(binary.BigEndian.Uint64(pix[8*2:])),
		A: math.Float64frombits(binary.BigEndian.Uint64(pix[8*3:])),
	}
}
func pSetRGBA256f(pix []byte, c colorExt.RGBA256f) {
	binary.BigEndian.PutUint64(pix[8*0:], math.Float64bits(c.R))
	binary.BigEndian.PutUint64(pix[8*1:], math.Float64bits(c.G))
	binary.BigEndian.PutUint64(pix[8*2:], math.Float64bits(c.B))
	binary.BigEndian.PutUint64(pix[8*3:], math.Float64bits(c.A))
}
