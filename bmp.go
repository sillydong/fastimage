package fastimage

func (f *FastImage) getBMPImageSize() (*ImageSize, error) {
	slice, err := f.getBytes(18, 8)
	if err != nil {
		return nil, err
	}

	imageSize := ImageSize{}

	imageSize.Width = uint32(readUint32(slice[0:4]))
	imageSize.Height = uint32(readUint32(slice[4:8]))

	return &imageSize, nil
}
