// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"log"
	"text/template"
)

type TypeInfo struct {
	FileName   string // graya64f.go
	TypeName   string // GrayA64f
	PixCommnet string // []struct{ Y, A float32 }
	Channels   int    // 2
	DepthType  string // reflect.Float32
	PixelSize  int    // 8
	HasAlpha   bool
}

func main() {
	for i := 0; i < len(types); i++ {
		out := bytes.NewBuffer([]byte{})
		if err := tmpl.Execute(out, types[i]); err != nil {
			log.Fatalf("%d, err = %v", i, err)
		}

		data, err := format.Source(out.Bytes())
		if err != nil {
			log.Fatalf("%d, err = %v", i, err)
		}
		if err := ioutil.WriteFile(types[i].FileName, data, 0666); err != nil {
			log.Fatalf("%d, err = %v", i, err)
		}
	}
}

var types = []TypeInfo{
	// Gray*
	TypeInfo{
		FileName:   `gray.go`,
		TypeName:   `Gray`,
		PixCommnet: `[]struct{ Y uint8 }`,
		Channels:   1,
		DepthType:  `reflect.Uint8`,
		PixelSize:  1 * 1,
	},
	TypeInfo{
		FileName:   `gray16.go`,
		TypeName:   `Gray16`,
		PixCommnet: `[]struct{ Y uint16 }`,
		Channels:   1,
		DepthType:  `reflect.Uint16`,
		PixelSize:  1 * 2,
	},
	TypeInfo{
		FileName:   `gray32i.go`,
		TypeName:   `Gray32i`,
		PixCommnet: `[]struct{ Y int32 }`,
		Channels:   1,
		DepthType:  `reflect.Int32`,
		PixelSize:  1 * 4,
	},
	TypeInfo{
		FileName:   `gray32f.go`,
		TypeName:   `Gray32f`,
		PixCommnet: `[]struct{ Y float32 }`,
		Channels:   1,
		DepthType:  `reflect.Float32`,
		PixelSize:  1 * 4,
	},
	TypeInfo{
		FileName:   `gray64i.go`,
		TypeName:   `Gray64i`,
		PixCommnet: `[]struct{ Y int64 }`,
		Channels:   1,
		DepthType:  `reflect.Int64`,
		PixelSize:  1 * 8,
	},
	TypeInfo{
		FileName:   `gray64f.go`,
		TypeName:   `Gray64f`,
		PixCommnet: `[]struct{ Y float64 }`,
		Channels:   1,
		DepthType:  `reflect.Float64`,
		PixelSize:  1 * 8,
	},

	// GrayA*
	TypeInfo{
		FileName:   `graya.go`,
		TypeName:   `GrayA`,
		PixCommnet: `[]struct{ Y, A uint8 }`,
		Channels:   2,
		DepthType:  `reflect.Uint8`,
		PixelSize:  2 * 1,
		HasAlpha:   true,
	},
	TypeInfo{
		FileName:   `graya32.go`,
		TypeName:   `GrayA32`,
		PixCommnet: `[]struct{ Y, A uint16 }`,
		Channels:   2,
		DepthType:  `reflect.Uint16`,
		PixelSize:  2 * 2,
		HasAlpha:   true,
	},
	TypeInfo{
		FileName:   `graya64i.go`,
		TypeName:   `GrayA64i`,
		PixCommnet: `[]struct{ Y, A int32 }`,
		Channels:   2,
		DepthType:  `reflect.Int32`,
		PixelSize:  2 * 4,
		HasAlpha:   true,
	},
	TypeInfo{
		FileName:   `graya64f.go`,
		TypeName:   `GrayA64f`,
		PixCommnet: `[]struct{ Y, A float32 }`,
		Channels:   2,
		DepthType:  `reflect.Float32`,
		PixelSize:  2 * 4,
		HasAlpha:   true,
	},
	TypeInfo{
		FileName:   `graya128i.go`,
		TypeName:   `GrayA128i`,
		PixCommnet: `[]struct{ Y, A int64 }`,
		Channels:   2,
		DepthType:  `reflect.Int64`,
		PixelSize:  2 * 8,
		HasAlpha:   true,
	},
	TypeInfo{
		FileName:   `graya128f.go`,
		TypeName:   `GrayA128f`,
		PixCommnet: `[]struct{ Y, A float64 }`,
		Channels:   2,
		DepthType:  `reflect.Float64`,
		PixelSize:  2 * 8,
		HasAlpha:   true,
	},

	// RGB*
	TypeInfo{
		FileName:   `rgb.go`,
		TypeName:   `RGB`,
		PixCommnet: `[]struct{ R, G, B uint8 }`,
		Channels:   3,
		DepthType:  `reflect.Uint8`,
		PixelSize:  3 * 1,
	},
	TypeInfo{
		FileName:   `rgb48.go`,
		TypeName:   `RGB48`,
		PixCommnet: `[]struct{ R, G, B uint16 }`,
		Channels:   3,
		DepthType:  `reflect.Uint16`,
		PixelSize:  3 * 2,
	},
	TypeInfo{
		FileName:   `rgb96i.go`,
		TypeName:   `RGB96i`,
		PixCommnet: `[]struct{ R, G, B int32 }`,
		Channels:   3,
		DepthType:  `reflect.Int32`,
		PixelSize:  3 * 4,
	},
	TypeInfo{
		FileName:   `rgb96f.go`,
		TypeName:   `RGB96f`,
		PixCommnet: `[]struct{ R, G, B float32 }`,
		Channels:   3,
		DepthType:  `reflect.Float32`,
		PixelSize:  3 * 4,
	},
	TypeInfo{
		FileName:   `rgb192i.go`,
		TypeName:   `RGB192i`,
		PixCommnet: `[]struct{ R, G, B int64 }`,
		Channels:   3,
		DepthType:  `reflect.Int64`,
		PixelSize:  3 * 8,
	},
	TypeInfo{
		FileName:   `rgb192f.go`,
		TypeName:   `RGB192f`,
		PixCommnet: `[]struct{ R, G, B float64 }`,
		Channels:   3,
		DepthType:  `reflect.Float64`,
		PixelSize:  3 * 8,
	},

	// RGBA*
	TypeInfo{
		FileName:   `rgba.go`,
		TypeName:   `RGBA`,
		PixCommnet: `[]struct{ R, G, B, A uint8 }`,
		Channels:   4,
		DepthType:  `reflect.Uint8`,
		PixelSize:  4 * 1,
		HasAlpha:   true,
	},
	TypeInfo{
		FileName:   `rgba64.go`,
		TypeName:   `RGBA64`,
		PixCommnet: `[]struct{ R, G, B, A uint16 }`,
		Channels:   4,
		DepthType:  `reflect.Uint16`,
		PixelSize:  4 * 2,
		HasAlpha:   true,
	},
	TypeInfo{
		FileName:   `rgba128i.go`,
		TypeName:   `RGBA128i`,
		PixCommnet: `[]struct{ R, G, B, A int32 }`,
		Channels:   4,
		DepthType:  `reflect.Int32`,
		PixelSize:  4 * 4,
		HasAlpha:   true,
	},
	TypeInfo{
		FileName:   `rgba128f.go`,
		TypeName:   `RGBA128f`,
		PixCommnet: `[]struct{ R, G, B, A float32 }`,
		Channels:   4,
		DepthType:  `reflect.Float32`,
		PixelSize:  4 * 4,
		HasAlpha:   true,
	},
	TypeInfo{
		FileName:   `rgba256i.go`,
		TypeName:   `RGBA256i`,
		PixCommnet: `[]struct{ R, G, B, A int64 }`,
		Channels:   4,
		DepthType:  `reflect.Int64`,
		PixelSize:  4 * 8,
		HasAlpha:   true,
	},
	TypeInfo{
		FileName:   `rgba256f.go`,
		TypeName:   `RGBA256f`,
		PixCommnet: `[]struct{ R, G, B, A float64 }`,
		Channels:   4,
		DepthType:  `reflect.Float64`,
		PixelSize:  4 * 8,
		HasAlpha:   true,
	},
}

var tmpl = template.Must(template.New("").Parse(`
// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Auto Generated By 'go generate', DONOT EDIT!!!

package image

import (
	"image"
	"image/color"
	"reflect"
	"unsafe"

	colorExt "github.com/chai2010/image/color"
)

type {{.TypeName}} struct {
	M struct {
		Pix    []uint8 // {{.PixCommnet}}
		Stride int
		Rect   image.Rectangle
	}
}

// New{{.TypeName}} returns a new {{.TypeName}} with the given bounds.
func New{{.TypeName}}(r image.Rectangle) *{{.TypeName}} {
	return new({{.TypeName}}).Init(make([]uint8, {{.PixelSize}}*r.Dx()*r.Dy()), {{.PixelSize}}*r.Dx(), r)
}

func (p *{{.TypeName}}) Init(pix []uint8, stride int, rect image.Rectangle) *{{.TypeName}} {
	*p = {{.TypeName}}{
		M: struct {
			Pix    []uint8
			Stride int
			Rect   image.Rectangle
		}{
			Pix:    pix,
			Stride: stride,
			Rect:   rect,
		},
	}
	return p
}

func (p *{{.TypeName}}) BaseType() image.Image { return asBaseType(p) }
func (p *{{.TypeName}}) Pix() []byte           { return p.M.Pix }
func (p *{{.TypeName}}) Stride() int           { return p.M.Stride }
func (p *{{.TypeName}}) Rect() image.Rectangle { return p.M.Rect }
func (p *{{.TypeName}}) Channels() int         { return {{.Channels}} }
func (p *{{.TypeName}}) Depth() reflect.Kind   { return {{.DepthType}} }

func (p *{{.TypeName}}) ColorModel() color.Model { return colorExt.{{.TypeName}}Model }

func (p *{{.TypeName}}) Bounds() image.Rectangle { return p.M.Rect }

func (p *{{.TypeName}}) At(x, y int) color.Color {
	return p.{{.TypeName}}At(x, y)
}

func (p *{{.TypeName}}) {{.TypeName}}At(x, y int) colorExt.{{.TypeName}} {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorExt.{{.TypeName}}{}
	}
	i := p.PixOffset(x, y)
	return *(*colorExt.{{.TypeName}})(unsafe.Pointer(&p.M.Pix[i]))
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *{{.TypeName}}) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*{{.PixelSize}}
}

func (p *{{.TypeName}}) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := p.ColorModel().Convert(c).(colorExt.{{.TypeName}})
	*(*colorExt.{{.TypeName}})(unsafe.Pointer(&p.M.Pix[i])) = c1
	return
}

func (p *{{.TypeName}}) Set{{.TypeName}}(x, y int, c colorExt.{{.TypeName}}) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	*(*colorExt.{{.TypeName}})(unsafe.Pointer(&p.M.Pix[i])) = c
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *{{.TypeName}}) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &{{.TypeName}}{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new({{.TypeName}}).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *{{.TypeName}}) Opaque() bool {
{{if .HasAlpha}}if p.M.Rect.Empty() {
		return true
	}
	for y := p.M.Rect.Min.Y; y < p.M.Rect.Max.Y; y++ {
		for x := p.M.Rect.Min.X; x < p.M.Rect.Max.X; x++ {
			if _, _, _, a := p.At(x, y).RGBA(); a == 0xFFFF {
				return false
			}
		}
	}
{{end}}return true
}

func (p *{{.TypeName}}) Draw(r image.Rectangle, src Image, sp image.Point) Image {
	panic("TODO")
}
`[1:]))
