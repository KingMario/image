// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image_test

import (
	"image"
	"os"
	"testing"

	imageExt "github.com/chai2010/image"
	_ "github.com/chai2010/image/bmp"
	_ "github.com/chai2010/image/gif"
	_ "github.com/chai2010/image/jpeg"
	_ "github.com/chai2010/image/png"
	_ "github.com/chai2010/image/tiff"
	_ "github.com/chai2010/image/webp"
)

type tFormatTester struct {
	FileName      string
	Format        string
	DecodeEnabled bool
	EncodeEnabled bool
}

var tFormatTesterList = []tFormatTester{
	tFormatTester{
		FileName:      "video-001.bmp",
		Format:        "bmp",
		DecodeEnabled: true,
		EncodeEnabled: true,
	},
	tFormatTester{
		FileName:      "video-001.gif",
		Format:        "gif",
		DecodeEnabled: true,
		EncodeEnabled: true,
	},
	tFormatTester{
		FileName:      "video-001.jpeg",
		Format:        "jpeg",
		DecodeEnabled: true,
		EncodeEnabled: true,
	},
	tFormatTester{
		FileName:      "video-001.png",
		Format:        "png",
		DecodeEnabled: true,
		EncodeEnabled: true,
	},
	tFormatTester{
		FileName:      "video-001.tiff",
		Format:        "tiff",
		DecodeEnabled: true,
		EncodeEnabled: true,
	},
	tFormatTester{
		FileName:      "video-001.webp",
		Format:        "webp",
		DecodeEnabled: true,
		EncodeEnabled: true,
	},
}

func TestFormats(t *testing.T) {
	os.MkdirAll("tempdir", 0666)
	defer os.RemoveAll("tempdir")

	golden, _, err := imageExt.Load("testdata/video-001.png")
	if err != nil {
		t.Fatalf("Load golden fialed: %v", err)
	}

	for i, v := range tFormatTesterList {
		if v.DecodeEnabled {
			m, format, err := imageExt.Load("testdata/" + v.FileName)
			if err != nil {
				t.Fatalf("%d, Load(%q) fail: %v", i, v.FileName, err)
			}
			if format != v.Format {
				t.Fatalf(
					"%d: %s, bad format; got %v, want <= %v",
					i, v.FileName, format, v.Format,
				)
			}

			// Compare the average delta to the tolerance level.
			want := int64(12 << 8)
			if got := averageDelta(golden, m); got > want {
				t.Fatalf(
					"%d, %s, average delta too high; got %d, want <= %d",
					i, v.FileName, got, want,
				)
			}
		}
		if v.EncodeEnabled && v.DecodeEnabled {
			err := imageExt.Save("tempdir/"+v.FileName, golden, nil)
			if err != nil {
				t.Fatalf("%d, Save(%q) fail: %v", i, v.FileName, err)
			}

			// load again
			m, format, err := imageExt.Load("tempdir/" + v.FileName)
			if err != nil {
				t.Fatalf("%d, Load(%q) fail: %v", i, v.FileName, err)
			}
			if format != v.Format {
				t.Fatalf(
					"%d: %s, bad format; got %v, want <= %v",
					i, v.FileName, format, v.Format,
				)
			}

			// Compare the average delta to the tolerance level.
			want := int64(12 << 8)
			if got := averageDelta(golden, m); got > want {
				t.Fatalf(
					"%d, %s, average delta too high; got %d, want <= %d",
					i, v.FileName, got, want,
				)
			}
		}
	}
}

// averageDelta returns the average delta in RGB space. The two images must
// have the same bounds.
func averageDelta(m0, m1 image.Image) int64 {
	b := m0.Bounds()
	var sum, n int64
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c0 := m0.At(x, y)
			c1 := m1.At(x, y)
			r0, g0, b0, _ := c0.RGBA()
			r1, g1, b1, _ := c1.RGBA()
			sum += delta(r0, r1)
			sum += delta(g0, g1)
			sum += delta(b0, b1)
			n += 3
		}
	}
	return sum / n
}

func delta(u0, u1 uint32) int64 {
	d := int64(u0) - int64(u1)
	if d < 0 {
		return -d
	}
	return d
}
