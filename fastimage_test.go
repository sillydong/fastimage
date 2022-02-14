package fastimage

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

var fastimage *FastImage

func init() {
	cfg := Config{
		Header: http.Header{
			// "Host": []string{"www.baidu.com"}, // for common fake host
		},
		ReadTimeout:        5 * time.Second,
		DialTimeout:        2 * time.Second,
		InsecureSkipVerify: true,
	}
	fastimage = NewFastImage(&cfg)
}

func TestBuffer(b *testing.T) {
	url := "http://host.domain.com/path.jpg"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		defer resp.Body.Close()

		reader := newReaderAt(resp.Body)

		buffers, err := reader.(*xbuffer).Slice(0, 4)
		if err != nil {
			fmt.Printf("%+v\n", err)
		} else {
			fmt.Printf("%+v\n", buffers)
		}

		tmp := make([]byte, 3)
		size, err := reader.(*xbuffer).ReadFull(tmp)
		if err != nil {
			fmt.Printf("%+v\n", err)
		} else {
			fmt.Printf("%+v\n", tmp)
			fmt.Printf("%+v\n", size)
		}

		buffer, err := reader.(*xbuffer).ReadByte()
		if err != nil {
			fmt.Printf("%+v\n", err)
		} else {
			fmt.Printf("%+v\n", buffer)
		}

		buffers2, err := reader.(*xbuffer).ReadBytes(2)
		if err != nil {
			fmt.Printf("%+v\n", err)
		} else {
			fmt.Printf("%+v\n", buffers2)
		}
	}
}

func TestBytes(b *testing.T) {
	bytes := []byte("abcdefghijklmnopqrstuvwxyz")
	fmt.Printf("%+v\n", bytes)
	fmt.Printf("%+v\n", bytes[2])
	fmt.Printf("%+v\n", bytes[2:4])
	fmt.Printf("%+v\n", bytes[22:25])
	fmt.Println(len(bytes))
}

func TestImage(t *testing.T) {
	url := "http://host.domain.com/path.jpg"
	imagetype, size, err := fastimage.Detect(url)
	fmt.Println(imagetype)
	fmt.Printf("%v\n", size)
	fmt.Printf("%+v\n", err)
}

func TestPNGImageA(b *testing.T) {
	url := "http://fc08.deviantart.net/fs71/f/2012/214/7/c/futurama__bender_by_suzura-d59kq1p.png"

	imagetype, size, err := fastimage.Detect(url)
	fmt.Println(imagetype)
	fmt.Printf("%v\n", size)
	fmt.Printf("%+v\n", err)
}

func BenchmarkPNGImageA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://fc08.deviantart.net/fs71/f/2012/214/7/c/futurama__bender_by_suzura-d59kq1p.png"

		imagetype, size, err := fastimage.Detect(url)
		fmt.Println(imagetype)
		fmt.Printf("%v\n", size)
		fmt.Printf("%+v\n", err)
	}
}

func TestJPEGImageA(b *testing.T) {
	url := "http://upload.wikimedia.org/wikipedia/commons/9/9a/SKA_dishes_big.jpg"

	imagetype, size, err := fastimage.Detect(url)
	fmt.Println(imagetype)
	fmt.Printf("%v\n", size)
	fmt.Printf("%+v\n", err)
}

func BenchmarkJPEGImageA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://upload.wikimedia.org/wikipedia/commons/9/9a/SKA_dishes_big.jpg"

		imagetype, size, err := fastimage.Detect(url)
		fmt.Println(imagetype)
		fmt.Printf("%v\n", size)
		fmt.Printf("%+v\n", err)
	}
}

func TestGIFImageA(b *testing.T) {
	url := "http://media.giphy.com/media/gXcIuJBbRi2Va/giphy.gif"

	imagetype, size, err := fastimage.Detect(url)
	fmt.Println(imagetype)
	fmt.Printf("%v\n", size)
	fmt.Printf("%+v\n", err)
}

func BenchmarkGIFImageA(b *testing.B) {
	for i := 0; i < b.N; i++ {

		url := "http://media.giphy.com/media/gXcIuJBbRi2Va/giphy.gif"

		imagetype, size, err := fastimage.Detect(url)
		fmt.Println(imagetype)
		fmt.Printf("%v\n", size)
		fmt.Printf("%+v\n", err)
	}
}

func TestBMPImageA(b *testing.T) {
	url := "http://www.fileformat.info/format/bmp/sample/1d71eff930af4773a836a32229fde106/download"
	imagetype, size, err := fastimage.Detect(url)
	fmt.Println(imagetype)
	fmt.Printf("%v\n", size)
	fmt.Printf("%+v\n", err)
}

func BenchmarkBMPImageA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://www.fileformat.info/format/bmp/sample/1d71eff930af4773a836a32229fde106/download"
		imagetype, size, err := fastimage.Detect(url)
		fmt.Println(imagetype)
		fmt.Printf("%v\n", size)
		fmt.Printf("%+v\n", err)
	}
}

func TestTIFFImageA(b *testing.T) {
	a := time.Now().UnixNano()
	url := "http://www.fileformat.info/format/tiff/sample/928c96cd555b40e19fad31d5c06374c5/download"
	imagetype, size, err := fastimage.Detect(url)
	fmt.Println(time.Now().UnixNano() - a)
	fmt.Println(imagetype)
	fmt.Printf("%v\n", size)
	fmt.Printf("%+v\n", err)
}

func BenchmarkTIFFImageA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://www.fileformat.info/format/tiff/sample/928c96cd555b40e19fad31d5c06374c5/download"
		imagetype, size, err := fastimage.Detect(url)
		fmt.Println(imagetype)
		fmt.Printf("%v\n", size)
		fmt.Printf("%+v\n", err)
	}
}

func TestWEBPImageA(b *testing.T) {
	url := "http://www.etherdream.com/WebP/Test.webp"
	imagetype, size, err := fastimage.Detect(url)
	fmt.Println(imagetype)
	fmt.Printf("%v\n", size)
	fmt.Printf("%+v\n", err)
}

func BenchmarkWEBPImageA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://www.etherdream.com/WebP/Test.webp"
		imagetype, size, err := fastimage.Detect(url)
		fmt.Println(imagetype)
		fmt.Printf("%v\n", size)
		fmt.Printf("%+v\n", err)
	}
}
