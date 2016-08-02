package fastimage

import (
	"encoding/binary"
)

func (f *decoder) getPNGImageSize() (*ImageSize, error) {
	slice, err := f.reader.(*xbuffer).Slice(16, 8)
	if err != nil {
		return nil, err
	}

	imageSize := ImageSize{}

	imageSize.Width = binary.BigEndian.Uint32(slice[0:4])
	imageSize.Height = binary.BigEndian.Uint32(slice[4:8])

	return &imageSize, nil
}
