package fastimage

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

var fastimage *FastImage

func init() {
	fastimage = NewFastImage(2, nil)
}

func TestBuffer(b *testing.T) {
	url := "http://pic.hualongxiang.com/app/image/2016/0405/09-54-25-1459821265.s.293x355.jpg"
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
	//url := "http://img03.store.sogou.com/net/a/04/link?appid=100520031&w=710&url=http%3A%2F%2Fmmbiz.qpic.cn%2Fmmbiz%2FQUZRHutbdrGlNSQbzcvHInkz4jRWMYjl0tYssEgtHR8qS5rEzMMCickFPulIcPj5xwy6pIriczRrRu0YAibAEJ2xA%2F0%3Fwx_fmt%3Dgif"
	//url :="http://pic.hualongxiang.com/app/image/2016/0405/09-54-25-1459821265.s.293x355.jpg"
	//url := "https://mmbiz.qlogo.cn/mmbiz/5gKn2ibOCyceiccOz6knZXUkOpom3HVXia6yToaDAAWQdc8uRL5VFViakV7Fa2O5J38oZOC2ib1Cyuaib0nIgTTdCiaHw/0?wx_fmt=jpeg"
	//url := "http://p.bydonline.com/img/27.jpg"
	//url := "http://pic.bbs.zszhili.com/data/attachment/forum/201604/24/110256gugqu9tzgtnauawe.jpg"
	url := "http://pics.18qiang.com/attachment/photo/Mon_1604/15253_736f1461223031839eece92bc6254.jpg"
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
