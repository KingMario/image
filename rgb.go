// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
)

type RGB struct {
	*image.RGBA
}

type RGB48 struct {
	*image.RGBA64
}

type RGB96f struct {
	*image.RGBA
}
