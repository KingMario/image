// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webp

/*
#cgo CFLAGS: -I./internal/libwebp/include  -I./internal/libwebp/src
#cgo !windows LDFLAGS: -lm

#include "webp.h"

#include <stdlib.h>
#include <string.h>

struct cgoWebpGetInfoReturn {
	int ok;
	int width;
	int height;
	int has_alpha;
} cgoWebpGetInfo(const uint8_t* data, size_t data_size) {
	struct cgoWebpGetInfoReturn t;
	t.ok = webpGetInfo(data, data_size, &t.width, &t.height, &t.has_alpha);
	return t;
}

struct cgoWebpDecodeGrayReturn {
	int ok;
	int width;
	int height;
	uint8_t* ptr;
} cgoWebpDecodeGray(const uint8_t* data, size_t data_size) {
	struct cgoWebpDecodeGrayReturn t;
	t.ptr = webpDecodeGray(data, data_size, &t.width, &t.height);
	t.ok = (t.ptr != NULL)? 1: 0;
	return t;
}

struct cgoWebpDecodeRGBReturn {
	int ok;
	int width;
	int height;
	uint8_t* ptr;
} cgoWebpDecodeRGB(const uint8_t* data, size_t data_size) {
	struct cgoWebpDecodeRGBReturn t;
	t.ptr = webpDecodeRGB(data, data_size, &t.width, &t.height);
	t.ok = (t.ptr != NULL)? 1: 0;
	return t;
}

struct cgoWebpDecodeRGBAReturn {
	int ok;
	int width;
	int height;
	uint8_t* ptr;
} cgoWebpDecodeRGBA(const uint8_t* data, size_t data_size) {
	struct cgoWebpDecodeRGBAReturn t;
	t.ptr = webpDecodeRGBA(data, data_size, &t.width, &t.height);
	t.ok = (t.ptr != NULL)? 1: 0;
	return t;
}

struct cgoWebpEncodeGrayReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpEncodeGray(const uint8_t* data, int width, int height, int stride, float quality_factor) {
	struct cgoWebpEncodeGrayReturn t;
	t.size = webpEncodeGray(data, width, height, stride, quality_factor, &t.ptr);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}

struct cgoWebpEncodeRGBReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpEncodeRGB(const uint8_t* data, int width, int height, int stride, float quality_factor) {
	struct cgoWebpEncodeRGBReturn t;
	t.size = webpEncodeRGB(data, width, height, stride, quality_factor, &t.ptr);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}

struct cgoWebpEncodeRGBAReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpEncodeRGBA(const uint8_t* data, int width, int height, int stride, float quality_factor) {
	struct cgoWebpEncodeRGBAReturn t;
	t.size = webpEncodeRGBA(data, width, height, stride, quality_factor, &t.ptr);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}

struct cgoWebpEncodeLosslessGrayReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpEncodeLosslessGray(const uint8_t* data, int width, int height, int stride) {
	struct cgoWebpEncodeLosslessGrayReturn t;
	t.size = webpEncodeLosslessGray(data, width, height, stride, &t.ptr);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}

struct cgoWebpEncodeLosslessRGBReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpEncodeLosslessRGB(const uint8_t* data, int width, int height, int stride) {
	struct cgoWebpEncodeLosslessRGBReturn t;
	t.size = webpEncodeLosslessRGB(data, width, height, stride, &t.ptr);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}

struct cgoWebpEncodeLosslessRGBAReturn {
	int ok;
	size_t size;
	uint8_t* ptr;
} cgoWebpEncodeLosslessRGBA(const uint8_t* data, int width, int height, int stride) {
	struct cgoWebpEncodeLosslessRGBAReturn t;
	t.size = webpEncodeLosslessRGBA(data, width, height, stride, &t.ptr);
	t.ok = (t.size != 0)? 1: 0;
	return t;
}

*/
import "C"
import (
	"errors"
	"unsafe"
)

const maxWebpHeaderSize = 32

func webpGetInfo(data []byte) (width, height int, has_alpha bool, err error) {
	if len(data) == 0 {
		err = errors.New("webpGetInfo: bad arguments")
		return
	}
	if len(data) > maxWebpHeaderSize {
		data = data[:maxWebpHeaderSize]
	}
	cData := cgoSafePtr(data)
	defer cgoFreePtr(cData)

	rv := C.cgoWebpGetInfo((*C.uint8_t)(cData), C.size_t(len(data)))
	if rv.ok != 1 {
		err = errors.New("webpGetInfo: failed")
		return
	}
	width, height = int(rv.width), int(rv.height)
	has_alpha = (rv.has_alpha != 0)
	return
}

func webpDecodeGray(data []byte) (pix []byte, width, height int, err error) {
	if len(data) == 0 {
		err = errors.New("webpDecodeGray: bad arguments")
		return
	}
	cData := cgoSafePtr(data)
	defer cgoFreePtr(cData)

	rv := C.cgoWebpDecodeGray((*C.uint8_t)(cData), C.size_t(len(data)))
	if rv.ok != 1 {
		err = errors.New("webpDecodeGray: failed")
		return
	}

	width, height = int(rv.width), int(rv.height)
	pix = make([]byte, width*height*1)
	copy(pix, ((*[1 << 30]byte)(unsafe.Pointer(rv.ptr)))[0:len(pix):len(pix)])
	C.webpFree(unsafe.Pointer(rv.ptr))
	return
}

func webpDecodeRGB(data []byte) (pix []byte, width, height int, err error) {
	if len(data) == 0 {
		err = errors.New("webpDecodeRGB: bad arguments")
		return
	}
	cData := cgoSafePtr(data)
	defer cgoFreePtr(cData)

	rv := C.cgoWebpDecodeRGB((*C.uint8_t)(cData), C.size_t(len(data)))
	if rv.ok != 1 {
		err = errors.New("webpDecodeRGB: failed")
		return
	}

	width, height = int(rv.width), int(rv.height)
	pix = make([]byte, width*height*3)
	copy(pix, ((*[1 << 30]byte)(unsafe.Pointer(rv.ptr)))[0:len(pix):len(pix)])
	C.webpFree(unsafe.Pointer(rv.ptr))
	return
}

func webpDecodeRGBA(data []byte) (pix []byte, width, height int, err error) {
	if len(data) == 0 {
		err = errors.New("webpDecodeRGBA: bad arguments")
		return
	}
	cData := cgoSafePtr(data)
	defer cgoFreePtr(cData)

	rv := C.cgoWebpDecodeRGBA((*C.uint8_t)(cData), C.size_t(len(data)))
	if rv.ok != 1 {
		err = errors.New("webpDecodeRGBA: failed")
		return
	}

	width, height = int(rv.width), int(rv.height)
	pix = make([]byte, width*height*4)
	copy(pix, ((*[1 << 30]byte)(unsafe.Pointer(rv.ptr)))[0:len(pix):len(pix)])
	C.webpFree(unsafe.Pointer(rv.ptr))
	return
}

func webpEncodeGray(
	pix []byte, width, height, stride int,
	quality_factor float32,
) (output []byte, err error) {
	if len(pix) == 0 || width <= 0 || height <= 0 || stride <= 0 || quality_factor < 0.0 {
		err = errors.New("webpEncodeGray: bad arguments")
		return
	}
	if stride < width*1 && len(pix) < height*stride {
		err = errors.New("webpEncodeGray: bad arguments")
		return
	}
	cPix := cgoSafePtr(pix)
	defer cgoFreePtr(cPix)

	rv := C.cgoWebpEncodeGray(
		(*C.uint8_t)(cPix), C.int(width), C.int(height),
		C.int(stride), C.float(quality_factor),
	)
	if rv.ok != 1 {
		err = errors.New("webpEncodeGray: failed")
		return
	}

	output = make([]byte, int(rv.size))
	copy(output, ((*[1 << 30]byte)(unsafe.Pointer(rv.ptr)))[0:len(output):len(output)])
	C.webpFree(unsafe.Pointer(rv.ptr))
	return
}

func webpEncodeRGB(
	pix []byte, width, height, stride int,
	quality_factor float32,
) (output []byte, err error) {
	if len(pix) == 0 || width <= 0 || height <= 0 || stride <= 0 || quality_factor < 0.0 {
		err = errors.New("webpEncodeRGB: bad arguments")
		return
	}
	if stride < width*3 && len(pix) < height*stride {
		err = errors.New("webpEncodeRGB: bad arguments")
		return
	}
	cPix := cgoSafePtr(pix)
	defer cgoFreePtr(cPix)

	rv := C.cgoWebpEncodeRGB(
		(*C.uint8_t)(cPix), C.int(width), C.int(height),
		C.int(stride), C.float(quality_factor),
	)
	if rv.ok != 1 {
		err = errors.New("webpEncodeRGB: failed")
		return
	}

	output = make([]byte, int(rv.size))
	copy(output, ((*[1 << 30]byte)(unsafe.Pointer(rv.ptr)))[0:len(output):len(output)])
	C.webpFree(unsafe.Pointer(rv.ptr))
	return
}

func webpEncodeRGBA(
	pix []byte, width, height, stride int,
	quality_factor float32,
) (output []byte, err error) {
	if len(pix) == 0 || width <= 0 || height <= 0 || stride <= 0 || quality_factor < 0.0 {
		err = errors.New("webpEncodeRGBA: bad arguments")
		return
	}
	if stride < width*4 && len(pix) < height*stride {
		err = errors.New("webpEncodeRGBA: bad arguments")
		return
	}
	cPix := cgoSafePtr(pix)
	defer cgoFreePtr(cPix)

	rv := C.cgoWebpEncodeRGBA(
		(*C.uint8_t)(cPix), C.int(width), C.int(height),
		C.int(stride), C.float(quality_factor),
	)
	if rv.ok != 1 {
		err = errors.New("webpEncodeRGBA: failed")
		return
	}

	output = make([]byte, int(rv.size))
	copy(output, ((*[1 << 30]byte)(unsafe.Pointer(rv.ptr)))[0:len(output):len(output)])
	C.webpFree(unsafe.Pointer(rv.ptr))
	return
}

func webpEncodeLosslessGray(
	pix []byte, width, height, stride int,
) (output []byte, err error) {
	if len(pix) == 0 || width <= 0 || height <= 0 || stride <= 0 {
		err = errors.New("webpEncodeLosslessGray: bad arguments")
		return
	}
	if stride < width*1 && len(pix) < height*stride {
		err = errors.New("webpEncodeLosslessGray: bad arguments")
		return
	}
	cPix := cgoSafePtr(pix)
	defer cgoFreePtr(cPix)

	rv := C.cgoWebpEncodeLosslessGray(
		(*C.uint8_t)(cPix), C.int(width), C.int(height),
		C.int(stride),
	)
	if rv.ok != 1 {
		err = errors.New("webpEncodeLosslessGray: failed")
		return
	}

	output = make([]byte, int(rv.size))
	copy(output, ((*[1 << 30]byte)(unsafe.Pointer(rv.ptr)))[0:len(output):len(output)])
	C.webpFree(unsafe.Pointer(rv.ptr))
	return
}

func webpEncodeLosslessRGB(
	pix []byte, width, height, stride int,
) (output []byte, err error) {
	if len(pix) == 0 || width <= 0 || height <= 0 || stride <= 0 {
		err = errors.New("webpEncodeLosslessRGB: bad arguments")
		return
	}
	if stride < width*3 && len(pix) < height*stride {
		err = errors.New("webpEncodeLosslessRGB: bad arguments")
		return
	}
	cPix := cgoSafePtr(pix)
	defer cgoFreePtr(cPix)

	rv := C.cgoWebpEncodeLosslessRGB(
		(*C.uint8_t)(cPix), C.int(width), C.int(height),
		C.int(stride),
	)
	if rv.ok != 1 {
		err = errors.New("webpEncodeLosslessRGB: failed")
		return
	}

	output = make([]byte, int(rv.size))
	copy(output, ((*[1 << 30]byte)(unsafe.Pointer(rv.ptr)))[0:len(output):len(output)])
	C.webpFree(unsafe.Pointer(rv.ptr))
	return
}

func webpEncodeLosslessRGBA(
	pix []byte, width, height, stride int,
) (output []byte, err error) {
	if len(pix) == 0 || width <= 0 || height <= 0 || stride <= 0 {
		err = errors.New("webpEncodeLosslessRGBA: bad arguments")
		return
	}
	if stride < width*4 && len(pix) < height*stride {
		err = errors.New("webpEncodeLosslessRGBA: bad arguments")
		return
	}
	cPix := cgoSafePtr(pix)
	defer cgoFreePtr(cPix)

	rv := C.cgoWebpEncodeLosslessRGBA(
		(*C.uint8_t)(cPix), C.int(width), C.int(height),
		C.int(stride),
	)
	if rv.ok != 1 {
		err = errors.New("webpEncodeLosslessRGBA: failed")
		return
	}

	output = make([]byte, int(rv.size))
	copy(output, ((*[1 << 30]byte)(unsafe.Pointer(rv.ptr)))[0:len(output):len(output)])
	C.webpFree(unsafe.Pointer(rv.ptr))
	return
}
