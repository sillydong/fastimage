package fastimage

import (
	"fmt"
	"testing"
)

func TestPNGImageA(b *testing.T) {
	url := "http://fc08.deviantart.net/fs71/f/2012/214/7/c/futurama__bender_by_suzura-d59kq1p.png"

	imagetype, size, err := GetFastImage(url)
	fmt.Println(imagetype)
	fmt.Printf("%v\n", size)
	fmt.Printf("%+v\n", err)
}

func BenchmarkPNGImageA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://fc08.deviantart.net/fs71/f/2012/214/7/c/futurama__bender_by_suzura-d59kq1p.png"

		imagetype, size, err := GetFastImage(url)
		fmt.Println(imagetype)
		fmt.Printf("%v\n", size)
		fmt.Printf("%+v\n", err)
	}
}

//
//func BenchmarkJPEGImageA(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		url := "http://7xiq7k.com1.z0.glb.clouddn.com/forum/201601/02/162328qeqeok71quzkx6z7.JPG"
//
//		imagetype, size, err := Detect(url)
//		fmt.Println(imagetype)
//		fmt.Printf("%v\n", size)
//		fmt.Printf("%+v\n", err)
//	}
//}
//
func TestGIFImageA(b *testing.T) {
	url := "http://media.giphy.com/media/gXcIuJBbRi2Va/giphy.gif"

	imagetype, size, err := GetFastImage(url)
	fmt.Println(imagetype)
	fmt.Printf("%v\n", size)
	fmt.Printf("%+v\n", err)
}

func BenchmarkGIFImageA(b *testing.B) {
	for i := 0; i < b.N; i++ {

		url := "http://media.giphy.com/media/gXcIuJBbRi2Va/giphy.gif"

		imagetype, size, err := GetFastImage(url)
		fmt.Println(imagetype)
		fmt.Printf("%v\n", size)
		fmt.Printf("%+v\n", err)
	}
}

func TestBMPImageA(b *testing.T) {
	url := "http://www.ac-grenoble.fr/ien.vienne1-2/spip/IMG/bmp_Image004.bmp"
	imagetype, size, err := GetFastImage(url)
	fmt.Println(imagetype)
	fmt.Printf("%v\n", size)
	fmt.Printf("%+v\n", err)
}

func BenchmarkBMPImageA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://www.ac-grenoble.fr/ien.vienne1-2/spip/IMG/bmp_Image004.bmp"
		imagetype, size, err := GetFastImage(url)
		fmt.Println(imagetype)
		fmt.Printf("%v\n", size)
		fmt.Printf("%+v\n", err)
	}
}

//func BenchmarkTIFFImageA(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		url := "http://www.fileformat.info/format/tiff/sample/928c96cd555b40e19fad31d5c06374c5/download"
//		imagetype, size, err := Detect(url)
//		fmt.Println(imagetype)
//		fmt.Printf("%v\n", size)
//		fmt.Printf("%+v\n", err)
//	}
//}

func TestWEBPImageA(b *testing.T) {
	url := "http://mindprod.com/image/jgloss/lossy.webp"
	imagetype, size, err := GetFastImage(url)
	fmt.Println(imagetype)
	fmt.Printf("%v\n", size)
	fmt.Printf("%+v\n", err)
}

func BenchmarkWEBPImageA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://mindprod.com/image/jgloss/lossy.webp"
		imagetype, size, err := GetFastImage(url)
		fmt.Println(imagetype)
		fmt.Printf("%v\n", size)
		fmt.Printf("%+v\n", err)
	}
}
