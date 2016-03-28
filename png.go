package fastimage

import (
	"encoding/binary"
)

func (f *FastImage) getPNGImageSize() (*ImageSize, error) {
	slice := make([]byte, 8)
	if _, err := f.reader.ReadAt(slice, 16); err != nil {
		return nil, err
	}

	imageSize := ImageSize{}

	imageSize.Width = binary.BigEndian.Uint32(slice[0:4])
	imageSize.Height = binary.BigEndian.Uint32(slice[4:8])

	return &imageSize, nil
}
