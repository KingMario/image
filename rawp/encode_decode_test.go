// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rawp

import (
	"bytes"
	"fmt"
	"image"
	"testing"

	image_ext "github.com/chai2010/gopkg/image"
)

func diff(m0, m1 image.Image) error {
	b0, b1 := m0.Bounds(), m1.Bounds()
	if !b0.Size().Eq(b1.Size()) {
		return fmt.Errorf("dimensions differ: %v vs %v", b0, b1)
	}
	if m0.ColorModel() != m1.ColorModel() {
		return fmt.Errorf("differ ColorModel: %T vs %T", m0.ColorModel(), m1.ColorModel())
	}
	dx := b1.Min.X - b0.Min.X
	dy := b1.Min.Y - b0.Min.Y
	for y := b0.Min.Y; y < b0.Max.Y; y++ {
		for x := b0.Min.X; x < b0.Max.X; x++ {
			c0 := m0.At(x, y)
			c1 := m1.At(x+dx, y+dy)
			r0, g0, b0, a0 := c0.RGBA()
			r1, g1, b1, a1 := c1.RGBA()
			if r0 != r1 || g0 != g1 || b0 != b1 || a0 != a1 {
				return fmt.Errorf("colors differ at (%d, %d): %v vs %v", x, y, c0, c1)
			}
		}
	}
	return nil
}

func encodeDecode(m image.Image) (image.Image, error) {
	var b bytes.Buffer
	err := Encode(&b, m, nil)
	if err != nil {
		return nil, err
	}
	m, err = Decode(&b, nil)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func TestEncodeDecode(t *testing.T) {
	imgs := []image.Image{
		image.NewGray(image.Rect(0, 0, 10, 10)),
		image.NewGray16(image.Rect(0, 0, 20, 20)),
		image_ext.NewGray32f(image.Rect(0, 0, 30, 30)),
		image_ext.NewRGB(image.Rect(0, 0, 40, 40)),
		image_ext.NewRGB48(image.Rect(0, 0, 50, 50)),
		image_ext.NewRGB96f(image.Rect(0, 0, 60, 60)),
		image.NewRGBA(image.Rect(0, 0, 70, 70)),
		image.NewRGBA64(image.Rect(0, 0, 80, 80)),
		image_ext.NewRGBA128f(image.Rect(0, 0, 90, 90)),
	}
	for i, m0 := range imgs {
		m1, err := encodeDecode(m0)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}
		err = diff(m0, m1)
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}
	}
}
