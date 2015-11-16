package imagefile

import (
	"bytes"
	"image"
	"image/color"
	"testing"
)

func TestEncodeEmptyImage(t *testing.T) {
	img := image.NewRGBA64(image.Rect(0, 0, 0, 0))
	w := new(bytes.Buffer)
	err := Encode(w, img)
	if err != nil {
		t.Errorf("err is not nil: %v", err)
	}
	if 0 != bytes.Compare(w.Bytes(), []byte("farbfeld\000\000\000\000\000\000\000\000")) {
		t.Errorf("encoded image differs")
	}
}

func TestEncodeSmallImage(t *testing.T) {
	img := image.NewNRGBA64(image.Rect(0, 0, 1, 1))
	img.Pix = []byte("aAbBcCdD")
	w := new(bytes.Buffer)
	err := Encode(w, img)
	if err != nil {
		t.Errorf("err is not nil: %v", err)
	}
	if 0 != bytes.Compare(w.Bytes(), []byte("farbfeld\000\000\000\001\000\000\000\001aAbBcCdD")) {
		t.Errorf("encoded image differs")
	}
}

func TestEncodeSubImage(t *testing.T) {
	img := image.NewNRGBA64(image.Rect(0, 0, 4, 4))
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			c := uint16(y*4*4 + x*4)
			img.SetNRGBA64(x, y, color.NRGBA64{c, c + 1, c + 2, c + 3})
		}
	}
	w := new(bytes.Buffer)
	err := Encode(w, img.SubImage(image.Rect(1, 1, 3, 3)))
	if err != nil {
		t.Errorf("err is not nil: %v", err)
	}
	if 0 != bytes.Compare(w.Bytes(), []byte("farbfeld\000\000\000\002\000\000\000\002"+
		"\x00\x14\x00\x15\x00\x16\x00\x17\x00\x18\x00\x19\x00\x1a\x00\x1b"+
		"\x00\x24\x00\x25\x00\x26\x00\x27\x00\x28\x00\x29\x00\x2a\x00\x2b")) {
		t.Errorf("encoded image differs")
	}
}

func TestEncodeNRGBA(t *testing.T) {
	img := image.NewNRGBA(image.Rect(0, 0, 1, 1))
	img.Pix = []byte{0x00, 0x55, 0xa0, 0xff}
	w := new(bytes.Buffer)
	err := Encode(w, img)
	if err != nil {
		t.Errorf("err is not nil: %v", err)
	}
	if 0 != bytes.Compare(w.Bytes(), []byte("farbfeld\000\000\000\001\000\000\000\001"+
		"\x00\x00\x55\x55\xa0\xa0\xff\xff")) {
		t.Errorf("encoded image differs")
	}
}
