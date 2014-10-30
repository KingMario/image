// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package image

import (
	"bufio"
	"image"
	"io"
	"os"
	"strings"
)

// Options are the encoding and decoding parameters.
type Options interface {
	Lossless() bool
	Quality() float32
}

// A Format holds an image format's name, magic header and how to decode it.
// Name is the name of the format, like "jpeg" or "png".
// Extensions is the name extensions, like ".jpg" or ".jpeg".
// Magics is the magic prefix that identifies the format's encoding. The magic
// string can contain "?" wildcards that each match any one byte.
// Decode is the function that decodes the encoded image.
// DecodeConfig is the function that decodes just its configuration.
// Encode is the function that encodes just its configuration.
type Format struct {
	Name         string
	Extensions   []string
	Magics       []string
	DecodeConfig func(r io.Reader) (image.Config, error)
	Decode       func(r io.Reader) (image.Image, error)
	Encode       func(w io.Writer, m image.Image, opt Options) error
}

// Formats is the list of registered formats.
var formats []Format

// RegisterFormat registers an image format for use by Encode and Decode.
func RegisterFormat(fmt Format) {
	formats = append(formats, Format{
		Name:         fmt.Name,
		Extensions:   append([]string(nil), fmt.Extensions...),
		Magics:       append([]string(nil), fmt.Magics...),
		DecodeConfig: fmt.DecodeConfig,
		Decode:       fmt.Decode,
		Encode:       fmt.Encode,
	})
}

// A reader is an io.Reader that can also peek ahead.
type reader interface {
	io.Reader
	Peek(int) ([]byte, error)
}

// asReader converts an io.Reader to a reader.
func asReader(r io.Reader) reader {
	if rr, ok := r.(reader); ok {
		return rr
	}
	return bufio.NewReader(r)
}

// Match reports whether magic matches b. Magic may contain "?" wildcards.
func match(magic string, b []byte) bool {
	if len(magic) != len(b) {
		return false
	}
	for i, c := range b {
		if magic[i] != c && magic[i] != '?' {
			return false
		}
	}
	return true
}

// Sniff determines the format by filename extension.
func sniffByName(filename string) Format {
	if idx := strings.LastIndex(filename, "."); idx >= 0 {
		ext := strings.ToLower(filename[idx:])
		for _, f := range formats {
			for _, extensions := range f.Extensions {
				if ext == extensions {
					return f
				}
			}
		}
	}
	return Format{}
}

// Sniff determines the format of r's data.
func sniffByMagic(r reader) Format {
	for _, f := range formats {
		for _, magic := range f.Magics {
			b, err := r.Peek(len(magic))
			if err == nil && match(magic, b) {
				return f
			}
		}
	}
	return Format{}
}

// Decode decodes an image that has been encoded in a registered format.
// The string returned is the format name used during format registration.
// Format registration is typically done by an init function in the codec-
// specific package.
func Decode(r io.Reader) (image.Image, string, error) {
	rr := asReader(r)
	f := sniffByMagic(rr)
	if f.Decode == nil {
		return nil, "", image.ErrFormat
	}
	m, err := f.Decode(rr)
	return m, f.Name, err
}

// DecodeConfig decodes the color model and dimensions of an image that has
// been encoded in a registered format. The string returned is the format name
// used during format registration. Format registration is typically done by
// an init function in the codec-specific package.
func DecodeConfig(r io.Reader) (image.Config, string, error) {
	rr := asReader(r)
	f := sniffByMagic(rr)
	if f.DecodeConfig == nil {
		return image.Config{}, "", image.ErrFormat
	}
	c, err := f.DecodeConfig(rr)
	return c, f.Name, err
}

// Encode encodes an image as a registered format.
// The format is the format name used during format registration.
// Format registration is typically done by an init function in the codec-
// specific package.
func Encode(format string, w io.Writer, m image.Image, opt Options) error {
	for _, f := range formats {
		if f.Name == format {
			return f.Encode(w, m, opt)
		}
	}
	return image.ErrFormat
}

func Load(filename string) (m image.Image, format string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	m, format, err = Decode(f)
	if err != nil {
		return
	}
	return
}

func Save(filename string, m image.Image, opt Options) (err error) {
	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()

	format := sniffByName(filename)
	if format.Encode == nil {
		return image.ErrFormat
	}
	if err = format.Encode(f, m, opt); err != nil {
		return
	}
	return
}
