// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rawp

import (
	"image"
	"reflect"
	"testing"
)

func BenchmarkEncode_gray_buffer(b *testing.B) {
	m := image.NewGray(image.Rect(0, 0, 256, 256))
	encoder := pixEncoder{1, reflect.Uint8}
	dataBuf := make([]byte, 1<<20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(m, dataBuf)
	}
}

func BenchmarkEncode_gray_no_buffer(b *testing.B) {
	m := image.NewGray(image.Rect(0, 0, 256, 256))
	encoder := pixEncoder{1, reflect.Uint8}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(m, nil)
	}
}

func BenchmarkDecode_gray_buffer(b *testing.B) {
	m := image.NewGray(image.Rect(0, 0, 256, 256))
	encoder := pixEncoder{1, reflect.Uint8}
	decoder := pixDecoder{1, reflect.Uint8, m.Bounds().Dx(), m.Bounds().Dy()}
	data, _ := encoder.Encode(m, nil)
	imgBuf := image.NewGray(image.Rect(0, 0, 512, 512))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(data, imgBuf)
	}
}

func BenchmarkDecode_gray_no_buffer(b *testing.B) {
	m := image.NewGray(image.Rect(0, 0, 256, 256))
	encoder := pixEncoder{1, reflect.Uint8}
	decoder := pixDecoder{1, reflect.Uint8, m.Bounds().Dx(), m.Bounds().Dy()}
	data, _ := encoder.Encode(m, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(data, nil)
	}
}
