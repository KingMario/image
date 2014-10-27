// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"image"
)

var (
	_ Drawer = (*Gray)(nil)
	_ Drawer = (*Gray16)(nil)
	_ Drawer = (*Gray32f)(nil)
	_ Drawer = (*GrayA)(nil)
	_ Drawer = (*GrayA32)(nil)
	_ Drawer = (*GrayA64f)(nil)
	_ Drawer = (*RGB)(nil)
	_ Drawer = (*RGB48)(nil)
	_ Drawer = (*RGB96f)(nil)
	_ Drawer = (*RGBA)(nil)
	_ Drawer = (*RGBA64)(nil)
	_ Drawer = (*RGBA128f)(nil)
	_ Drawer = (*Unknown)(nil)
)

type Drawer interface {
	// Draw aligns r.Min in dst with sp in src and then replaces the
	// rectangle r in dst with the result of drawing src on dst.
	Draw(dst Image, r image.Rectangle, src Image, sp image.Point) Image
}
