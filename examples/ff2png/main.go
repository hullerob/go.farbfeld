package main

import (
	_ "github.com/hullerob/go.farbfeld"
	"image"
	"image/png"
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
	err = png.Encode(os.Stdout, m)
	os.Stdout.Sync()
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func usage() {
	os.Stderr.WriteString("usage: ff2png\n")
}
