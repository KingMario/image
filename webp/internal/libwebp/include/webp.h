// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef WEBP_H_
#define WEBP_H_

#include <stddef.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

int webpGetInfo(
	const uint8_t* data, size_t data_size,
	int* width, int* height,
	int* has_alpha
);

uint8_t* webpDecodeGray(
	const uint8_t* data, size_t data_size,
	int* width, int* height
);
uint8_t* webpDecodeRGB(
	const uint8_t* data, size_t data_size,
	int* width, int* height
);
uint8_t* webpDecodeRGBA(
	const uint8_t* data, size_t data_size,
	int* width, int* height
);

size_t webpEncodeGray(
	const uint8_t* gray, int width, int height, int stride, float quality_factor,
	uint8_t** output
);
size_t webpEncodeRGB(
	const uint8_t* rgb, int width, int height, int stride, float quality_factor,
	uint8_t** output
);
size_t webpEncodeRGBA(
	const uint8_t* rgba, int width, int height, int stride, float quality_factor,
	uint8_t** output
);

size_t webpEncodeLosslessGray(
	const uint8_t* gray, int width, int height, int stride,
	uint8_t** output
);
size_t webpEncodeLosslessRGB(
	const uint8_t* rgb, int width, int height, int stride,
	uint8_t** output
);
size_t webpEncodeLosslessRGBA(
	const uint8_t* rgba, int width, int height, int stride,
	uint8_t** output
);

void webpFree(void* p);

#ifdef __cplusplus
}
#endif
#endif // WEBP_H_
