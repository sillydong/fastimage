package fastimage

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

type FastImage struct {
	Url     string
	Timeout time.Duration

	resp   *http.Response
	reader io.ReaderAt
}

func (f *FastImage) Detect() (ImageType, *ImageSize, error) {
	start := time.Now().UnixNano()
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
				Timeout: f.Timeout,
			}).Dial,
		},
	}

	var err2 error
	f.resp, err2 = client.Do(req)
	if err2 != nil {
		return Unknown, nil, err2
	}
	defer f.resp.Body.Close()

	if f.resp.StatusCode != 200 {
		return Unknown, nil, fmt.Errorf(f.resp.Status)
	}

	f.reader = newReaderAt(f.resp.Body)

	var t ImageType
	var s *ImageSize
	var e error

	typebuf := make([]byte, 2)
	if _, err := f.reader.ReadAt(typebuf, 0); err != nil {
		return Unknown, nil, err
	}

	switch {
	case string(typebuf) == "BM":
		t = BMP
		s, e = f.getBMPImageSize()
	case bytes.Equal(typebuf, []byte{0x47, 0x49}):
		t = GIF
		s, e = f.getGIFImageSize()
	case bytes.Equal(typebuf, []byte{0xFF, 0xD8}):
		t = JPEG
		s, e = f.getJPEGImageSize()
	case bytes.Equal(typebuf, []byte{0x89, 0x50}):
		t = PNG
		s, e = f.getPNGImageSize()
	case string(typebuf) == "II" || string(typebuf) == "MM":
		t = TIFF
		s, e = f.getTIFFImageSize()
	case string(typebuf) == "RI":
		t = WEBP
		s, e = f.getWEBPImageSize()
	default:
		t = Unknown
		e = fmt.Errorf("Unkown image type[%v]", typebuf)
	}
	stop := time.Now().UnixNano()
	if stop-start > 500000000 {
		fmt.Printf("[%v]%v\n", stop-start, f.Url)
	}
	return t, s, e
}

func GetImageSize(url string) (ImageType, *ImageSize, error) {
	return (&FastImage{Url: url, Timeout: 3 * time.Second}).Detect()
}
