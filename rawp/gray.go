// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rawp

import (
	"github.com/chai2010/image"
)

func decodeGray(data []byte, size image.Point) *image.Gray {
	gray := newGray(image.Rect(0, 0, p.Width, p.Height), buf)
	var off = 0
	for y := 0; y < p.Height; y++ {
		copy(gray.Pix[y*gray.Stride:][:p.Width], data[off:])
		off += p.Width
	}
	m = gray
	return
}
