package fastimage

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

type FastImage struct {
	Url string

	resp   *http.Response
	buffer bytes.Buffer
}

func (f *FastImage) Detect() (ImageType, *ImageSize, error) {
	u, err := url.Parse(f.Url)
	if err != nil {
		return Unknown, nil, err
	}

	header := make(http.Header)
	header.Set("Referer", u.Scheme+"://"+u.Host)
	header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.87 Safari/537.36")

	req := &http.Request{
		Method:     "GET",
		URL:        u,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     header,
		Host:       u.Host,
	}

	client := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 2 * time.Second,
			}).Dial,
		},
	}

	var err2 error
	f.resp, err2 = client.Do(req)
	defer f.resp.Body.Close()

	if err2 != nil {
		return Unknown, nil, err2
	}

	var t ImageType
	var s *ImageSize
	var e error

	typebuf, _ := f.getBytes(0, 2)
	switch {
	case string(typebuf) == "BM":
		t = BMP
		s, e = f.getBMPImageSize()
	case bytes.Equal(typebuf, []byte{0x47, 0x49}):
		t = GIF
		s, e = f.getGIFImageSize()
	//case bytes.Equal(typebuf, []byte{0xFF, 0xD8}):
	//	t = JPEG
	//	s = f.getJPEGImageSize()
	case bytes.Equal(typebuf, []byte{0x89, 0x50}):
		t = PNG
		s, e = f.getPNGImageSize()
	//case string(typebuf) == "II" || string(typebuf) == "MM":
	//	t = TIFF
	//	s = f.getTIFFImageSize()
	case string(typebuf) == "RI":
		t = WEBP
		s, e = f.getWEBPImageSize()
	default:
		t = Unknown
		e = fmt.Errorf("Unkown image type")
	}

	fmt.Println(f.buffer.Len())

	return t, s, e
}

func (f *FastImage) getBytes(start, size int) ([]byte, error) {
	if f.buffer.Len() < start+size {
		err := f.readToBuffer(start + size - f.buffer.Len())
		if err != nil {
			return []byte{}, err
		}
	}

	bytes := f.buffer.Bytes()
	return bytes[start : start+size], nil
}

func (f *FastImage) readToBuffer(size int) error {
	chunk := make([]byte, size)
	_, err := f.resp.Body.Read(chunk)
	if err != nil {
		return err
	}
	f.buffer.Write(chunk)
	return nil
}

func GetFastImage(url string) (ImageType, *ImageSize, error) {
	return (&FastImage{Url: url}).Detect()
}
