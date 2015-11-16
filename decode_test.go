package imagefile

import (
	"bytes"
	"image"
	"io"
	"testing"
)

var decodeTests = []struct {
	descr string
	input []byte
	img   []byte
	w, h  int
}{
	{
		"empty image",
		[]byte("farbfeld\000\000\000\000\000\000\000\000"),
		[]byte{},
		0, 0,
	},
	{
		"image 1x1",
		[]byte("farbfeld\000\000\000\001\000\000\000\001rRgGbBaA"),
		[]byte("rRgGbBaA"),
		1, 1,
	},
}

func TestDecodeImage(t *testing.T) {
	for _, test := range decodeTests {
		r := bytes.NewBuffer(test.input)
		img, err := Decode(r)
		if err != nil {
			t.Errorf("test %s: err not nil: %v", test.descr, err)
			continue
		}
		nrgba64, ok := img.(*image.NRGBA64)
		if !ok {
			t.Errorf("%s: decoded image has wrong type", test.descr)
			continue
		}
		if dx, dy := img.Bounds().Dx(), img.Bounds().Dy(); dx != test.w || dy != test.h {
			t.Errorf("%s: decoded image has wrong size: %d x %d | expected: %d x %d",
				test.descr, dx, dy, test.w, test.h)
		}
		if bytes.Compare(test.img, nrgba64.Pix) != 0 {
			t.Errorf("%s: image data differs")
		}
	}
}

func TestDecodeImageBadHeader(t *testing.T) {
	r := bytes.NewBuffer([]byte("farbfeld\000\000\000\000\000\000\000"))
	img, err := Decode(r)
	if err != io.ErrUnexpectedEOF {
		t.Errorf("returned error is wrong, expected: %#v | got: %#v", io.ErrUnexpectedEOF, err)
	}
	if img != nil {
		t.Errorf("returned image is not nil")
	}
}

func TestDecodeImageBad(t *testing.T) {
	r := bytes.NewBuffer([]byte("farbfeld\000\000\000\001\000\000\000\001rgb"))
	_, err := Decode(r)
	if err != io.ErrUnexpectedEOF {
		t.Errorf("returned error is wrong, expected: %#v | got: %#v", io.ErrUnexpectedEOF, err)
	}
}
