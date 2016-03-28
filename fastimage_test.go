package fastimage

import (
	"fmt"
	"testing"
"time"
)

func TestBytes(b *testing.T) {
	bytes := []byte("abcdefghijklmnopqrstuvwxyz")
	fmt.Printf("%+v\n", bytes)
	fmt.Printf("%+v\n", bytes[2])
	fmt.Printf("%+v\n", bytes[2:4])
	fmt.Printf("%+v\n", bytes[22:25])
	fmt.Println(len(bytes))
}

func TestPNGImageA(b *testing.T) {
	url := "http://fc08.deviantart.net/fs71/f/2012/214/7/c/futurama__bender_by_suzura-d59kq1p.png"

	imagetype, size, err := GetImageSize(url)
	fmt.Println(imagetype)
	fmt.Printf("%v\n", size)
	fmt.Printf("%+v\n", err)
}

func BenchmarkPNGImageA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://fc08.deviantart.net/fs71/f/2012/214/7/c/futurama__bender_by_suzura-d59kq1p.png"

		imagetype, size, err := GetImageSize(url)
		fmt.Println(imagetype)
		fmt.Printf("%v\n", size)
		fmt.Printf("%+v\n", err)
	}
}

func TestJPEGImageA(b *testing.T) {
	url := "http://upload.wikimedia.org/wikipedia/commons/9/9a/SKA_dishes_big.jpg"

	imagetype, size, err := GetImageSize(url)
	fmt.Println(imagetype)
	fmt.Printf("%v\n", size)
	fmt.Printf("%+v\n", err)
}

func BenchmarkJPEGImageA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://upload.wikimedia.org/wikipedia/commons/9/9a/SKA_dishes_big.jpg"

		imagetype, size, err := GetImageSize(url)
		fmt.Println(imagetype)
		fmt.Printf("%v\n", size)
		fmt.Printf("%+v\n", err)
	}
}

func TestGIFImageA(b *testing.T) {
	url := "http://media.giphy.com/media/gXcIuJBbRi2Va/giphy.gif"

	imagetype, size, err := GetImageSize(url)
	fmt.Println(imagetype)
	fmt.Printf("%v\n", size)
	fmt.Printf("%+v\n", err)
}

func BenchmarkGIFImageA(b *testing.B) {
	for i := 0; i < b.N; i++ {

		url := "http://media.giphy.com/media/gXcIuJBbRi2Va/giphy.gif"

		imagetype, size, err := GetImageSize(url)
		fmt.Println(imagetype)
		fmt.Printf("%v\n", size)
		fmt.Printf("%+v\n", err)
	}
}

func TestBMPImageA(b *testing.T) {
	url := "http://www.fileformat.info/format/bmp/sample/1d71eff930af4773a836a32229fde106/download"
	imagetype, size, err := GetImageSize(url)
	fmt.Println(imagetype)
	fmt.Printf("%v\n", size)
	fmt.Printf("%+v\n", err)
}

func BenchmarkBMPImageA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://www.fileformat.info/format/bmp/sample/1d71eff930af4773a836a32229fde106/download"
		imagetype, size, err := GetImageSize(url)
		fmt.Println(imagetype)
		fmt.Printf("%v\n", size)
		fmt.Printf("%+v\n", err)
	}
}

func TestTIFFImageA(b *testing.T) {
	a := time.Now().UnixNano()
	url := "http://www.fileformat.info/format/tiff/sample/928c96cd555b40e19fad31d5c06374c5/download"
	imagetype, size, err := GetImageSize(url)
	fmt.Println(time.Now().UnixNano() - a)
	fmt.Println(imagetype)
	fmt.Printf("%v\n", size)
	fmt.Printf("%+v\n", err)
}

func BenchmarkTIFFImageA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://www.fileformat.info/format/tiff/sample/928c96cd555b40e19fad31d5c06374c5/download"
		imagetype, size, err := GetImageSize(url)
		fmt.Println(imagetype)
		fmt.Printf("%v\n", size)
		fmt.Printf("%+v\n", err)
	}
}

func TestWEBPImageA(b *testing.T) {
	url := "http://mindprod.com/image/jgloss/lossy.webp"
	imagetype, size, err := GetImageSize(url)
	fmt.Println(imagetype)
	fmt.Printf("%v\n", size)
	fmt.Printf("%+v\n", err)
}

func BenchmarkWEBPImageA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://mindprod.com/image/jgloss/lossy.webp"
		imagetype, size, err := GetImageSize(url)
		fmt.Println(imagetype)
		fmt.Printf("%v\n", size)
		fmt.Printf("%+v\n", err)
	}
}
