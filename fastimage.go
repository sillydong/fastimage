package fastimage

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	Header                http.Header
	InsecureSkipVerify    bool
	AllowedContentTypes   []string
	AllowedFileExtensions []string
	DialTimeout           time.Duration
	ReadTimeout           time.Duration
}

// FastImage instance needs to be initialized before use
type FastImage struct {
	config *Config
	client *http.Client
	header http.Header
}

const (
	DefaultDialTimeout time.Duration = time.Second
	DefaultReadTimeout               = time.Second
)

//NewFastImage returns a FastImage client
func NewFastImage(cfg *Config) *FastImage {
	combinedHeaders := http.Header{}
	if cfg != nil && cfg.Header != nil {
		combinedHeaders = cfg.Header
	}
	if _, exists := combinedHeaders["User-Agent"]; !exists {
		combinedHeaders.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.87 Safari/537.36")
	}

	dialTimeout, readTimeout := DefaultDialTimeout, DefaultReadTimeout
	if cfg != nil {
		if cfg.DialTimeout != 0 {
			dialTimeout = cfg.DialTimeout
		}
		if cfg.ReadTimeout != 0 {
			readTimeout = cfg.ReadTimeout
		}
	}

	return &FastImage{
		config: cfg,
		client: &http.Client{
			Transport: &http.Transport{
				Dial:            (&net.Dialer{Timeout: dialTimeout}).Dial,
				TLSClientConfig: &tls.Config{InsecureSkipVerify: cfg.InsecureSkipVerify},
			},
			Timeout: readTimeout,
		},
		header: combinedHeaders,
	}
}

type decoder struct {
	reader io.ReaderAt
}

func (f *FastImage) newRequest(url *url.URL, fakeHost string) *http.Request {
	header := f.header.Clone()
	header.Set("Referer", url.Scheme+"://"+url.Host)

	req := &http.Request{
		Method:     "GET",
		URL:        url,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     header,
	}
	if _, exists := header["Host"]; exists {
		req.Host = header.Get("Host")
	}

	if fakeHost != "" {
		req.Host = fakeHost
	}

	return req
}

//Detect image type and size
func (f *FastImage) Detect(uri string, fakeHosts ...string) (ImageType, *ImageSize, error) {
	//start := time.Now().UnixNano()
	fakeHost := ""
	if len(fakeHosts) > 0 {
		fakeHost = fakeHosts[0]
	}
	u, err := url.Parse(uri)
	if err != nil {
		return Unknown, nil, err
	}

	if len(f.config.AllowedFileExtensions) > 0 {
		ext := filepath.Ext(u.Path)
		allowed := false
		for _, allowedExt := range f.config.AllowedFileExtensions {
			if strings.HasSuffix(ext, allowedExt) {
				allowed = true
				break
			}
		}
		if !allowed {
			return Unknown, nil, fmt.Errorf("%v has file extension %s, allowed extensions are %v", uri, ext, f.config.AllowedFileExtensions)
		}
	}

	req := f.newRequest(u, fakeHost)

	resp, err2 := f.client.Do(req)
	if err2 != nil {
		return Unknown, nil, err2
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Unknown, nil, fmt.Errorf(resp.Status)
	}
	if len(f.config.AllowedContentTypes) > 0 {
		contentType := resp.Header.Get("Content-Type")
		allowed := false
		for _, allowedContentType := range f.config.AllowedContentTypes {
			if strings.Contains(contentType, allowedContentType) {
				allowed = true
			}
		}
		if !allowed {
			return Unknown, nil, fmt.Errorf("%v has Content-Type %s, allowed Content-Types are %v", uri, contentType, f.config.AllowedContentTypes)
		}
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
	return NewFastImage(nil).Detect(url)
}
