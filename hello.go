// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"fmt"
	"log"

	imageExt "github.com/chai2010/image"
	_ "github.com/chai2010/image/jpeg"
	_ "github.com/chai2010/image/webp"
)

func main() {
	lena, _, err := imageExt.Load("testdata/lena.jpg")
	if err != nil {
		log.Fatalf("Load fialed: %v", err)
	}
	err = imageExt.Save("lena.webp", lena, imageExt.NewOptions(true, 0))
	if err != nil {
		log.Fatalf("Save fialed: %v", err)
	}
	fmt.Println("Save as lossless lena.webp !")
}
