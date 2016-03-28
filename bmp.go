package fastimage

func (f *FastImage) getBMPImageSize() (*ImageSize, error) {
	slice := make([]byte,8)
	if _,err:= f.reader.ReadAt(slice,18);err!=nil{
		return nil,err
	}

	imageSize := ImageSize{}

	imageSize.Width = uint32(readUint32(slice[0:4]))
	imageSize.Height = uint32(readUint32(slice[4:8]))

	return &imageSize, nil
}
