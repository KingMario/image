// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"bytes"
	"io/ioutil"
	"log"
)

func main() {
	// gen Gray*
	genFile("gray.go", "gray16.go", "Gray", "Gray16", `// []struct{ Y uint8 }`, `// []struct{ Y uint16 }`)
	genFile("gray.go", "gray32i.go", "Gray", "Gray32i", `// []struct{ Y uint8 }`, `// []struct{ Y int32 }`)
	genFile("gray.go", "gray32f.go", "Gray", "Gray32f", `// []struct{ Y uint8 }`, `// []struct{ Y float32 }`)
	genFile("gray.go", "gray64i.go", "Gray", "Gray64i", `// []struct{ Y uint8 }`, `// []struct{ Y int64 }`)
	genFile("gray.go", "gray64f.go", "Gray", "Gray64f", `// []struct{ Y uint8 }`, `// []struct{ Y float64 }`)

	// gen GrayA*
	genFile("gray.go", "graya.go", "Gray", "GrayA", `// []struct{ Y uint8 }`, `// []struct{ Y, A uint8 }`)
	genFile("gray.go", "graya32.go", "Gray", "GrayA32", `// []struct{ Y uint8 }`, `// []struct{ Y, A uint16 }`)
	genFile("gray.go", "graya64i.go", "Gray", "GrayA64i", `// []struct{ Y uint8 }`, `// []struct{ Y, A int32 }`)
	genFile("gray.go", "graya64f.go", "Gray", "GrayA64f", `// []struct{ Y uint8 }`, `// []struct{ Y, A float32 }`)
	genFile("gray.go", "graya128i.go", "Gray", "GrayA128i", `// []struct{ Y uint8 }`, `// []struct{ Y, A int64 }`)
	genFile("gray.go", "graya128f.go", "Gray", "GrayA128f", `// []struct{ Y uint8 }`, `// []struct{ Y, A float64 }`)

	// gen RGB*
	genFile("gray.go", "rgb.go", "Gray", "RGB", `// []struct{ Y uint8 }`, `// []struct{ R, G, B uint8 }`)
	genFile("gray.go", "rgb48.go", "Gray", "RGB48", `// []struct{ Y uint8 }`, `// []struct{ R, G, B uint16 }`)
	genFile("gray.go", "rgb96i.go", "Gray", "RGB96i", `// []struct{ Y uint8 }`, `// []struct{ R, G, B int32 }`)
	genFile("gray.go", "rgb96f.go", "Gray", "RGB96f", `// []struct{ Y uint8 }`, `// []struct{ R, G, B float32 }`)
	genFile("gray.go", "rgb192i.go", "Gray", "RGB192i", `// []struct{ Y uint8 }`, `// []struct{ R, G, B int64 }`)
	genFile("gray.go", "rgb192f.go", "Gray", "RGB192f", `// []struct{ Y uint8 }`, `// []struct{ R, G, B float64 }`)

	// gen RGBA*
	genFile("gray.go", "rgba.go", "Gray", "RGBA", `// []struct{ Y uint8 }`, `// []struct{ R, G, B, A uint8 }`)
	genFile("gray.go", "rgba64.go", "Gray", "RGBA64", `// []struct{ Y uint8 }`, `// []struct{ R, G, B, A uint16 }`)
	genFile("gray.go", "rgba128i.go", "Gray", "RGBA128i", `// []struct{ Y uint8 }`, `// []struct{ R, G, B, A int32 }`)
	genFile("gray.go", "rgba128f.go", "Gray", "RGBA128f", `// []struct{ Y uint8 }`, `// []struct{ R, G, B, A float32 }`)
	genFile("gray.go", "rgba256i.go", "Gray", "RGBA256i", `// []struct{ Y uint8 }`, `// []struct{ R, G, B, A int64 }`)
	genFile("gray.go", "rgba256f.go", "Gray", "RGBA256f", `// []struct{ Y uint8 }`, `// []struct{ R, G, B, A float64 }`)
}

func genFile(srcName, dstName, oldType, newType, oldPixComment, newPixComment string) {
	data, err := ioutil.ReadFile(srcName)
	if err != nil {
		log.Fatal(err)
	}
	data = bytes.Replace(data, []byte(oldType), []byte(newType), -1)
	data = bytes.Replace(data, []byte(oldPixComment), []byte(newPixComment), -1)
	err = ioutil.WriteFile(dstName, data, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
