package fastimage

func (f *decoder) getBMPImageSize() (*ImageSize, error) {
	slice, err := f.reader.(*xbuffer).Slice(18, 8)
	if err != nil {
		return nil, err
	}

	imageSize := ImageSize{}

	imageSize.Width = uint32(readUint32(slice[0:4]))
	imageSize.Height = uint32(readUint32(slice[4:8]))

	return &imageSize, nil
}
