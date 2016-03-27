package fastimage

import "fmt"

// ImageType represents the type of the image detected, or `Unknown`.
type ImageType uint

//go:generate stringer -type=ImageType -output=image_type_string.go
const (
	// GIF represents a GIF image
	GIF ImageType = iota
	// PNG represents a PNG image
	PNG
	// JPEG represents a JPEG image
	JPEG
	// BMP represents a BMP image
	BMP
	// TIFF represents a TIFF image
	TIFF
	// WEBP represendts a WEBP image
	WEBP
	// Unknown represents an unknown image type
	Unknown
)

const _ImageType_name = "GIFPNGJPEGBMPTIFFWEBPUnknown"

var _ImageType_index = [...]uint8{0, 3, 6, 10, 13, 17, 21, 28}

func (i ImageType) String() string {
	if i+1 >= ImageType(len(_ImageType_index)) {
		return fmt.Sprintf("ImageType(%d)", i)
	}
	return _ImageType_name[_ImageType_index[i]:_ImageType_index[i+1]]
}
