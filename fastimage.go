package fastimage

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// FastImage instance needs to be initialized before use
type FastImage struct {
	Client *http.Client
	header *http.Header
}

//NewFastImage returns a FastImage client
func NewFastImage(timeout int, headers map[string]string) *FastImage {

	header := &http.Header{}
	if headers != nil {
		for headerKey, headerValue := range headers {
			header.Add(headerKey, headerValue)
		}
	}

	return &FastImage{
		Client: &http.Client{
			Transport: &http.Transport{
				Dial:            (&net.Dialer{Timeout: time.Duration(timeout) * time.Second}).Dial,
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
		header: header,
	}
}

type decoder struct {
	reader io.ReaderAt
}

func (f *FastImage) newRequest(url *url.URL) *http.Request {
	header := make(http.Header)
	header.Set("Referer", url.Scheme+"://"+url.Host)
	header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.87 Safari/537.36")

	//override default header settings
	for chKey := range *f.header {
		header.Set(chKey, f.header.Get(chKey))
	}

	req := &http.Request{
		Method:     "GET",
		URL:        url,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     header,
	}
	//Host in header affects nothing in golang, consider a not-fixed-yet bug.
	if _, present := (*f.header)["Host"]; present {
		req.Host = f.header.Get("Host")
	}

	return req
}

//Detect image type and size
func (f *FastImage) Detect(uri string) (ImageType, *ImageSize, error) {
	//start := time.Now().UnixNano()
	u, err := url.Parse(uri)
	if err != nil {
		return Unknown, nil, err
	}

	req := f.newRequest(u)

	resp, err2 := f.Client.Do(req)
	if err2 != nil {
		return Unknown, nil, err2
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Unknown, nil, fmt.Errorf(resp.Status)
	}
	if !strings.Contains(resp.Header.Get("Content-Type"), "image") {
		return Unknown, nil, fmt.Errorf("%v is not image", uri)
	}

	d := &decoder{
		reader: newReaderAt(resp.Body),
	}

	var t ImageType
	var s *ImageSize
	var e error

	typebuf := make([]byte, 2)
	if _, err := d.reader.ReadAt(typebuf, 0); err != nil {
		return Unknown, nil, err
	}

	switch {
	case string(typebuf) == "BM":
		t = BMP
		s, e = d.getBMPImageSize()
	case bytes.Equal(typebuf, []byte{0x47, 0x49}):
		t = GIF
		s, e = d.getGIFImageSize()
	case bytes.Equal(typebuf, []byte{0xFF, 0xD8}):
		t = JPEG
		s, e = d.getJPEGImageSize()
	case bytes.Equal(typebuf, []byte{0x89, 0x50}):
		t = PNG
		s, e = d.getPNGImageSize()
	case string(typebuf) == "II" || string(typebuf) == "MM":
		t = TIFF
		s, e = d.getTIFFImageSize()
	case string(typebuf) == "RI":
		t = WEBP
		s, e = d.getWEBPImageSize()
	default:
		t = Unknown
		e = fmt.Errorf("Unkown image type[%v]", typebuf)
	}
	//stop := time.Now().UnixNano()
	//if stop-start > 500000000 {
	//	fmt.Printf("[%v]%v\n", stop-start, f.Url)
	//}
	return t, s, e
}

//GetImageSize create a default fastimage instance to detect image type and size
func GetImageSize(url string) (ImageType, *ImageSize, error) {
	return NewFastImage(2, nil).Detect(url)
}
