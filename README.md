image
=====

PkgDoc: [http://godoc.org/github.com/chai2010/image](http://godoc.org/github.com/chai2010/image)

Install
=======

Install `GCC` or `MinGW` ([download here](http://tdm-gcc.tdragon.net/download)) at first,
and then run these commands:

1. `go get github.com/chai2010/image`
2. `go run hello.go`

Example
=======

```Go
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
```

BUGS
====

Report bugs to <chaishushan@gmail.com>.

Thanks!
