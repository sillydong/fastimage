package fastimage

func (f *FastImage) getGIFImageSize() (*ImageSize, error) {
	slice := make([]byte, 4)
	if _, err := f.reader.ReadAt(slice, 6); err != nil {
		return nil, err
	}

	imageSize := ImageSize{}

	imageSize.Width = uint32(readULint16(slice[0:2]))
	imageSize.Height = uint32(readULint16(slice[2:4]))

	return &imageSize, nil
}
