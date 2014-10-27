// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

import (
	"reflect"
)

type Pixel struct {
	Channels int          // 1:Gray, 2:GrayA, 3:RGB, 4:RGBA
	Depth    reflect.Kind // Uint8/Uint16/Int32/Int64/Float32/Float64
	Value    []byte       // Value is big-endian format
}

func (c Pixel) RGBA() (r, g, b, a uint32) {
	return
}
