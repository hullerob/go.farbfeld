package main

import (
	"github.com/hullerob/go.farbfeld"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		usage()
		os.Exit(1)
	}
	m, _, err := image.Decode(os.Stdin)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
	err = imagefile.Encode(os.Stdout, m)
	os.Stdout.Sync()
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func usage() {
	os.Stderr.WriteString("usage: img2ff\n")
}
