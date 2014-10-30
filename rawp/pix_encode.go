// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rawp

import (
	"fmt"
	"image"
	"image/color"
	"reflect"

	imageExt "github.com/chai2010/image"
)

type pixEncoder struct {
	Channels int          // 1/2/3/4
	DataType reflect.Kind // Uint8/Uint16/Int32/Int64/Float32/Float64
}

func (p *pixEncoder) Encode(m image.Image, buf []byte) (data []byte, err error) {
	// Gray/Gray16/Gray32f
	if p.Channels == 1 && p.DataType == reflect.Uint8 {
		return p.encodeGray(m, buf)
	}
	if p.Channels == 1 && p.DataType == reflect.Uint16 {
		return p.encodeGray16(m, buf)
	}
	if p.Channels == 1 && p.DataType == reflect.Float32 {
		return p.encodeGray32f(m, buf)
	}

	// RGB/RGB48/RGB96f
	if p.Channels == 3 && p.DataType == reflect.Uint8 {
		return p.encodeRGB(m, buf)
	}
	if p.Channels == 3 && p.DataType == reflect.Uint16 {
		return p.encodeRGB48(m, buf)
	}
	if p.Channels == 3 && p.DataType == reflect.Float32 {
		return p.encodeRGB96f(m, buf)
	}

	// RGBA/RGBA64/RGBA128f
	if p.Channels == 4 && p.DataType == reflect.Uint8 {
		return p.encodeRGBA(m, buf)
	}
	if p.Channels == 4 && p.DataType == reflect.Uint16 {
		return p.encodeRGBA64(m, buf)
	}
	if p.Channels == 4 && p.DataType == reflect.Float32 {
		return p.encodeRGBA128f(m, buf)
	}

	// Unknown
	err = fmt.Errorf("image/rawp: Encode, unknown image format, channels = %v, dataType = %v", p.Channels, p.DataType)
	return
}

func (p *pixEncoder) encodeGray(m image.Image, buf []byte) (data []byte, err error) {
	b := m.Bounds()
	d := newBytes(b.Dx()*b.Dy(), buf)
	switch m := m.(type) {
	case *image.Gray:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			copy(d[off:][:b.Dx()], m.Pix[y*m.Stride:])
			off += b.Dx()
		}
	case *image.Gray16:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				d[off] = uint8(m.Gray16At(x, y).Y >> 8)
				off++
			}
		}
	case *image.YCbCr:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			copy(d[off:][:b.Dx()], m.Y[y*m.YStride:])
			off += b.Dx()
		}
	default:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				d[off] = color.GrayModel.Convert(m.At(x, y)).(color.Gray).Y
				off++
			}
		}
	}
	data = d
	return
}

func (p *pixEncoder) encodeGray16(m image.Image, buf []byte) (data []byte, err error) {
	b := m.Bounds()
	d := newBytes(b.Dx()*b.Dy()*2, buf)
	switch m := m.(type) {
	case *image.Gray:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				builtin.PutUint16(d[off:], uint16(m.GrayAt(x, y).Y)<<8)
				off += 2
			}
		}
	case *image.Gray16:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Gray16At(x, y)
				builtin.PutUint16(d[off:], v.Y)
				off += 2
			}
		}
	case *image.YCbCr:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.YCbCrAt(x, y)
				builtin.PutUint16(d[off:], uint16(v.Y)<<8)
				off += 2
			}
		}
	default:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := color.Gray16Model.Convert(m.At(x, y)).(color.Gray16)
				builtin.PutUint16(d[off:], v.Y)
				off += 2
			}
		}
	}
	data = d
	return
}

func (p *pixEncoder) encodeGray32f(m image.Image, buf []byte) (data []byte, err error) {
	b := m.Bounds()
	d := newBytes(b.Dx()*b.Dy()*4, buf)
	switch m := m.(type) {
	case *image.Gray:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				builtin.PutFloat32(d[off:], float32(uint16(m.GrayAt(x, y).Y)<<8))
				off += 4
			}
		}
	case *image.Gray16:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Gray16At(x, y)
				builtin.PutFloat32(d[off:], float32(v.Y))
				off += 4
			}
		}
	case *imageExt.Gray32f:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Gray32fAt(x, y)
				builtin.PutFloat32(d[off:], v.Y)
				off += 4
			}
		}
	case *imageExt.RGB96f:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGB96fAt(x, y)
				builtin.PutFloat32(d[off:], 0.2990*v.R+0.5870*v.G+0.1140*v.B)
				off += 4
			}
		}
	case *imageExt.RGBA128f:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBA128fAt(x, y)
				builtin.PutFloat32(d[off:], 0.2990*v.R+0.5870*v.G+0.1140*v.B)
				off += 4
			}
		}
	case *image.YCbCr:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.YCbCrAt(x, y)
				builtin.PutFloat32(d[off:], float32(uint16(v.Y)<<8))
				off += 4
			}
		}
	default:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := color.Gray16Model.Convert(m.At(x, y)).(color.Gray16)
				builtin.PutFloat32(d[off:], float32(v.Y))
				off += 4
			}
		}
	}
	data = d
	return
}

func (p *pixEncoder) encodeRGB(m image.Image, buf []byte) (data []byte, err error) {
	b := m.Bounds()
	d := newBytes(b.Dx()*b.Dy()*3, buf)
	switch m := m.(type) {
	case *image.Gray:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.GrayAt(x, y)
				d[off+0] = v.Y
				d[off+1] = v.Y
				d[off+2] = v.Y
				off += 3
			}
		}
	case *image.Gray16:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Gray16At(x, y)
				d[off+0] = uint8(v.Y >> 8)
				d[off+1] = uint8(v.Y >> 8)
				d[off+2] = uint8(v.Y >> 8)
				off += 3
			}
		}
	case *imageExt.RGB:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			copy(d[off:][:b.Dx()*3], m.Pix[y*m.Stride:])
			off += b.Dx() * 3
		}
	case *image.RGBA:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBAAt(x, y)
				d[off+0] = v.R
				d[off+1] = v.G
				d[off+2] = v.B
				off += 3
			}
		}
	case *image.YCbCr:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.YCbCrAt(x, y)
				R, G, B := color.YCbCrToRGB(v.Y, v.Cb, v.Cr)
				d[off+0] = R
				d[off+1] = G
				d[off+2] = B
				off += 3
			}
		}
	default:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := color.RGBAModel.Convert(m.At(x, y)).(color.RGBA)
				d[off+0] = v.R
				d[off+1] = v.G
				d[off+2] = v.B
				off += 3
			}
		}
	}
	data = d
	return
}

func (p *pixEncoder) encodeRGB48(m image.Image, buf []byte) (data []byte, err error) {
	b := m.Bounds()
	d := newBytes(b.Dx()*b.Dy()*6, buf)
	switch m := m.(type) {
	case *image.Gray:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.GrayAt(x, y)
				builtin.PutUint16(d[off+0:], uint16(v.Y)<<8)
				builtin.PutUint16(d[off+2:], uint16(v.Y)<<8)
				builtin.PutUint16(d[off+4:], uint16(v.Y)<<8)
				off += 6
			}
		}
	case *image.Gray16:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Gray16At(x, y)
				builtin.PutUint16(d[off+0:], v.Y)
				builtin.PutUint16(d[off+2:], v.Y)
				builtin.PutUint16(d[off+4:], v.Y)
				off += 6
			}
		}
	case *imageExt.RGB:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBAt(x, y)
				builtin.PutUint16(d[off+0:], uint16(v.R)<<8)
				builtin.PutUint16(d[off+2:], uint16(v.G)<<8)
				builtin.PutUint16(d[off+4:], uint16(v.B)<<8)
				off += 6
			}
		}
	case *image.RGBA:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBAAt(x, y)
				builtin.PutUint16(d[off+0:], uint16(v.R)<<8)
				builtin.PutUint16(d[off+2:], uint16(v.G)<<8)
				builtin.PutUint16(d[off+4:], uint16(v.B)<<8)
				off += 6
			}
		}
	case *imageExt.RGB48:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGB48At(x, y)
				builtin.PutUint16(d[off+0:], v.R)
				builtin.PutUint16(d[off+2:], v.G)
				builtin.PutUint16(d[off+4:], v.B)
				off += 6
			}
		}
	case *image.RGBA64:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBA64At(x, y)
				builtin.PutUint16(d[off+0:], v.R)
				builtin.PutUint16(d[off+2:], v.G)
				builtin.PutUint16(d[off+4:], v.B)
				off += 6
			}
		}
	case *image.YCbCr:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.YCbCrAt(x, y)
				R, G, B := color.YCbCrToRGB(v.Y, v.Cb, v.Cr)
				builtin.PutUint16(d[off+0:], uint16(R)<<8)
				builtin.PutUint16(d[off+2:], uint16(G)<<8)
				builtin.PutUint16(d[off+4:], uint16(B)<<8)
				off += 6
			}
		}
	default:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := color.RGBA64Model.Convert(m.At(x, y)).(color.RGBA64)
				builtin.PutUint16(d[off+0:], v.R)
				builtin.PutUint16(d[off+2:], v.G)
				builtin.PutUint16(d[off+4:], v.B)
				off += 6
			}
		}
	}
	data = d
	return
}

func (p *pixEncoder) encodeRGB96f(m image.Image, buf []byte) (data []byte, err error) {
	b := m.Bounds()
	d := newBytes(b.Dx()*b.Dy()*12, buf)
	switch m := m.(type) {
	case *image.Gray:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.GrayAt(x, y)
				builtin.PutFloat32(d[off+0:], float32(uint16(v.Y)<<8))
				builtin.PutFloat32(d[off+4:], float32(uint16(v.Y)<<8))
				builtin.PutFloat32(d[off+8:], float32(uint16(v.Y)<<8))
				off += 12
			}
		}
	case *image.Gray16:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Gray16At(x, y)
				builtin.PutFloat32(d[off+0:], float32(v.Y))
				builtin.PutFloat32(d[off+4:], float32(v.Y))
				builtin.PutFloat32(d[off+8:], float32(v.Y))
				off += 12
			}
		}
	case *imageExt.Gray32f:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Gray32fAt(x, y)
				builtin.PutFloat32(d[off+0:], v.Y)
				builtin.PutFloat32(d[off+4:], v.Y)
				builtin.PutFloat32(d[off+8:], v.Y)
				off += 12
			}
		}
	case *imageExt.RGB:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBAt(x, y)
				builtin.PutFloat32(d[off+0:], float32(uint16(v.R)<<8))
				builtin.PutFloat32(d[off+4:], float32(uint16(v.G)<<8))
				builtin.PutFloat32(d[off+8:], float32(uint16(v.B)<<8))
				off += 12
			}
		}
	case *image.RGBA:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBAAt(x, y)
				builtin.PutFloat32(d[off+0:], float32(uint16(v.R)<<8))
				builtin.PutFloat32(d[off+4:], float32(uint16(v.G)<<8))
				builtin.PutFloat32(d[off+8:], float32(uint16(v.B)<<8))
				off += 12
			}
		}
	case *image.RGBA64:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBA64At(x, y)
				builtin.PutFloat32(d[off+0:], float32(v.R))
				builtin.PutFloat32(d[off+4:], float32(v.G))
				builtin.PutFloat32(d[off+8:], float32(v.B))
				off += 12
			}
		}
	case *imageExt.RGB96f:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGB96fAt(x, y)
				builtin.PutFloat32(d[off+0:], v.R)
				builtin.PutFloat32(d[off+4:], v.G)
				builtin.PutFloat32(d[off+8:], v.B)
				off += 12
			}
		}
	case *imageExt.RGBA128f:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBA128fAt(x, y)
				builtin.PutFloat32(d[off+0:], v.R)
				builtin.PutFloat32(d[off+4:], v.G)
				builtin.PutFloat32(d[off+8:], v.B)
				off += 12
			}
		}
	case *image.YCbCr:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.YCbCrAt(x, y)
				R, G, B := color.YCbCrToRGB(v.Y, v.Cb, v.Cr)
				builtin.PutFloat32(d[off+0:], float32(uint16(R)<<8))
				builtin.PutFloat32(d[off+4:], float32(uint16(G)<<8))
				builtin.PutFloat32(d[off+8:], float32(uint16(B)<<8))
				off += 12
			}
		}
	default:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := color.RGBA64Model.Convert(m.At(x, y)).(color.RGBA64)
				builtin.PutFloat32(d[off+0:], float32(v.R))
				builtin.PutFloat32(d[off+4:], float32(v.G))
				builtin.PutFloat32(d[off+8:], float32(v.B))
				off += 12
			}
		}
	}
	data = d
	return
}

func (p *pixEncoder) encodeRGBA(m image.Image, buf []byte) (data []byte, err error) {
	b := m.Bounds()
	d := newBytes(b.Dx()*b.Dy()*4, buf)
	switch m := m.(type) {
	case *image.Gray:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.GrayAt(x, y)
				d[off+0] = v.Y
				d[off+1] = v.Y
				d[off+2] = v.Y
				d[off+3] = 0xFF
				off += 4
			}
		}
	case *image.Gray16:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Gray16At(x, y)
				d[off+0] = uint8(v.Y >> 8)
				d[off+1] = uint8(v.Y >> 8)
				d[off+2] = uint8(v.Y >> 8)
				d[off+3] = 0xFF
				off += 4
			}
		}
	case *imageExt.RGB:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBAt(x, y)
				d[off+0] = uint8(v.R >> 8)
				d[off+1] = uint8(v.G >> 8)
				d[off+2] = uint8(v.B >> 8)
				d[off+3] = 0xFF
				off += 4
			}
		}
	case *image.RGBA:
		var off = 0
		for y := 0; y < b.Max.Y-b.Min.Y; y++ {
			copy(d[off:][:b.Dx()*4], m.Pix[y*m.Stride:])
			off += b.Dx() * 4
		}
	case *imageExt.RGB48:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGB48At(x, y)
				d[off+0] = uint8(v.R >> 8)
				d[off+1] = uint8(v.G >> 8)
				d[off+2] = uint8(v.B >> 8)
				d[off+3] = 0xFF
				off += 4
			}
		}
	case *image.RGBA64:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBA64At(x, y)
				d[off+0] = uint8(v.R >> 8)
				d[off+1] = uint8(v.G >> 8)
				d[off+2] = uint8(v.B >> 8)
				d[off+3] = uint8(v.A >> 8)
				off += 4
			}
		}
	case *image.YCbCr:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.YCbCrAt(x, y)
				R, G, B := color.YCbCrToRGB(v.Y, v.Cb, v.Cr)
				d[off+0] = R
				d[off+1] = G
				d[off+2] = B
				d[off+3] = 0xFF
				off += 4
			}
		}
	default:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := color.RGBAModel.Convert(m.At(x, y)).(color.RGBA)
				d[off+0] = v.R
				d[off+1] = v.G
				d[off+2] = v.B
				d[off+3] = v.A
				off += 4
			}
		}
	}
	data = d
	return
}

func (p *pixEncoder) encodeRGBA64(m image.Image, buf []byte) (data []byte, err error) {
	b := m.Bounds()
	d := newBytes(b.Dx()*b.Dy()*8, buf)
	switch m := m.(type) {
	case *image.Gray:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.GrayAt(x, y)
				builtin.PutUint16(d[off+0:], uint16(v.Y)<<8)
				builtin.PutUint16(d[off+2:], uint16(v.Y)<<8)
				builtin.PutUint16(d[off+4:], uint16(v.Y)<<8)
				builtin.PutUint16(d[off+6:], 0xFFFF)
				off += 8
			}
		}
	case *image.Gray16:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Gray16At(x, y)
				builtin.PutUint16(d[off+0:], v.Y)
				builtin.PutUint16(d[off+2:], v.Y)
				builtin.PutUint16(d[off+4:], v.Y)
				builtin.PutUint16(d[off+6:], 0xFFFF)
				off += 8
			}
		}
	case *imageExt.RGB:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBAt(x, y)
				builtin.PutUint16(d[off+0:], uint16(v.R)<<8)
				builtin.PutUint16(d[off+2:], uint16(v.G)<<8)
				builtin.PutUint16(d[off+4:], uint16(v.B)<<8)
				builtin.PutUint16(d[off+6:], 0xFFFF)
				off += 8
			}
		}
	case *imageExt.RGB48:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGB48At(x, y)
				builtin.PutUint16(d[off+0:], v.R)
				builtin.PutUint16(d[off+2:], v.G)
				builtin.PutUint16(d[off+4:], v.B)
				builtin.PutUint16(d[off+6:], 0xFFFF)
				off += 8
			}
		}
	case *image.RGBA:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBAAt(x, y)
				builtin.PutUint16(d[off+0:], uint16(v.R)<<8)
				builtin.PutUint16(d[off+2:], uint16(v.G)<<8)
				builtin.PutUint16(d[off+4:], uint16(v.B)<<8)
				builtin.PutUint16(d[off+6:], uint16(v.A)<<8)
				off += 8
			}
		}
	case *image.RGBA64:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBA64At(x, y)
				builtin.PutUint16(d[off+0:], v.R)
				builtin.PutUint16(d[off+2:], v.G)
				builtin.PutUint16(d[off+4:], v.B)
				builtin.PutUint16(d[off+6:], v.A)
				off += 8
			}
		}
	case *image.YCbCr:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.YCbCrAt(x, y)
				R, G, B := color.YCbCrToRGB(v.Y, v.Cb, v.Cr)
				builtin.PutUint16(d[off+0:], uint16(R)<<8)
				builtin.PutUint16(d[off+2:], uint16(G)<<8)
				builtin.PutUint16(d[off+4:], uint16(B)<<8)
				builtin.PutUint16(d[off+6:], 0xFFFF)
				off += 8
			}
		}
	default:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := color.RGBA64Model.Convert(m.At(x, y)).(color.RGBA64)
				builtin.PutUint16(d[off+0:], v.R)
				builtin.PutUint16(d[off+2:], v.G)
				builtin.PutUint16(d[off+4:], v.B)
				builtin.PutUint16(d[off+6:], v.A)
				off += 8
			}
		}
	}
	data = d
	return
}

func (p *pixEncoder) encodeRGBA128f(m image.Image, buf []byte) (data []byte, err error) {
	b := m.Bounds()
	d := newBytes(b.Dx()*b.Dy()*16, buf)
	switch m := m.(type) {
	case *image.Gray:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.GrayAt(x, y)
				builtin.PutFloat32(d[off+0:], float32(uint16(v.Y)<<8))
				builtin.PutFloat32(d[off+4:], float32(uint16(v.Y)<<8))
				builtin.PutFloat32(d[off+8:], float32(uint16(v.Y)<<8))
				builtin.PutFloat32(d[off+12:], 0xFFFF)
				off += 16
			}
		}
	case *image.Gray16:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Gray16At(x, y)
				builtin.PutFloat32(d[off+0:], float32(v.Y))
				builtin.PutFloat32(d[off+4:], float32(v.Y))
				builtin.PutFloat32(d[off+8:], float32(v.Y))
				builtin.PutFloat32(d[off+12:], 0xFFFF)
				off += 16
			}
		}
	case *imageExt.Gray32f:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.Gray32fAt(x, y)
				builtin.PutFloat32(d[off+0:], v.Y)
				builtin.PutFloat32(d[off+4:], v.Y)
				builtin.PutFloat32(d[off+8:], v.Y)
				builtin.PutFloat32(d[off+12:], 0xFFFF)
				off += 16
			}
		}
	case *imageExt.RGB:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBAt(x, y)
				builtin.PutFloat32(d[off+0:], float32(uint16(v.R)<<8))
				builtin.PutFloat32(d[off+4:], float32(uint16(v.G)<<8))
				builtin.PutFloat32(d[off+8:], float32(uint16(v.B)<<8))
				builtin.PutFloat32(d[off+12:], 0xFFFF)
				off += 16
			}
		}
	case *imageExt.RGB48:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGB48At(x, y)
				builtin.PutFloat32(d[off+0:], float32(v.R))
				builtin.PutFloat32(d[off+4:], float32(v.G))
				builtin.PutFloat32(d[off+8:], float32(v.B))
				builtin.PutFloat32(d[off+12:], 0xFFFF)
				off += 16
			}
		}
	case *imageExt.RGB96f:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGB96fAt(x, y)
				builtin.PutFloat32(d[off+0:], v.R)
				builtin.PutFloat32(d[off+4:], v.G)
				builtin.PutFloat32(d[off+8:], v.B)
				builtin.PutFloat32(d[off+12:], 0xFFFF)
				off += 16
			}
		}
	case *image.RGBA:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBAAt(x, y)
				builtin.PutFloat32(d[off+0:], float32(uint16(v.R)<<8))
				builtin.PutFloat32(d[off+4:], float32(uint16(v.G)<<8))
				builtin.PutFloat32(d[off+8:], float32(uint16(v.B)<<8))
				builtin.PutFloat32(d[off+12:], float32(uint16(v.A)<<8))
				off += 16
			}
		}
	case *image.RGBA64:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBA64At(x, y)
				builtin.PutFloat32(d[off+0:], float32(v.R))
				builtin.PutFloat32(d[off+4:], float32(v.G))
				builtin.PutFloat32(d[off+8:], float32(v.B))
				builtin.PutFloat32(d[off+12:], float32(v.A))
				off += 16
			}
		}
	case *imageExt.RGBA128f:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.RGBA128fAt(x, y)
				builtin.PutFloat32(d[off+0:], v.R)
				builtin.PutFloat32(d[off+4:], v.G)
				builtin.PutFloat32(d[off+8:], v.B)
				builtin.PutFloat32(d[off+12:], v.A)
				off += 16
			}
		}
	case *image.YCbCr:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := m.YCbCrAt(x, y)
				R, G, B := color.YCbCrToRGB(v.Y, v.Cb, v.Cr)
				builtin.PutFloat32(d[off+0:], float32(uint16(R)<<8))
				builtin.PutFloat32(d[off+4:], float32(uint16(G)<<8))
				builtin.PutFloat32(d[off+8:], float32(uint16(B)<<8))
				builtin.PutFloat32(d[off+12:], 0xFFFF)
				off += 16
			}
		}
	default:
		var off = 0
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				v := color.RGBA64Model.Convert(m.At(x, y)).(color.RGBA64)
				builtin.PutFloat32(d[off+0:], float32(v.R))
				builtin.PutFloat32(d[off+4:], float32(v.G))
				builtin.PutFloat32(d[off+8:], float32(v.B))
				builtin.PutFloat32(d[off+12:], float32(v.A))
				off += 16
			}
		}
	}
	data = d
	return
}
