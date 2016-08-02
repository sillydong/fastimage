package fastimage

func (f *decoder) getGIFImageSize() (*ImageSize, error) {
	slice, err := f.reader.(*xbuffer).Slice(6, 4)
	if err != nil {
		return nil, err
	}

	imageSize := ImageSize{}

	imageSize.Width = uint32(readULint16(slice[0:2]))
	imageSize.Height = uint32(readULint16(slice[2:4]))

	return &imageSize, nil
}
