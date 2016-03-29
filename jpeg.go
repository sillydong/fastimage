package fastimage

import "fmt"

func (f *FastImage) getJPEGImageSize() (*ImageSize, error) {
	return f.parseJPEGData(2, nextSegment)
}

const (
	nextSegment = iota
	sofSegment
	skipSegment
	parseSegment
	eioSegment
)

func (f *FastImage) parseJPEGData(offset int, segment int) (*ImageSize, error) {
	switch segment {
	case nextSegment:
		newOffset := offset + 1
		bytes, err := f.reader.(*xbuffer).Slice(newOffset, 1)
		if err != nil {
			return nil, err
		}
		b := bytes[0]
		if b == 0xFF {
			return f.parseJPEGData(newOffset, sofSegment)
		}
		return f.parseJPEGData(newOffset, nextSegment)
	case sofSegment:
		newOffset := offset + 1
		bytes, err := f.reader.(*xbuffer).Slice(newOffset, 1)
		if err != nil {
			return nil, err
		}
		b := bytes[0]
		if b >= 0xE0 && b <= 0xEF {
			return f.parseJPEGData(newOffset, skipSegment)
		}
		if (b >= 0xC0 && b <= 0xC3) || (b >= 0xC5 && b <= 0xC7) || (b >= 0xC9 && b <= 0xCB) || b >= 0xCD && b <= 0xCF {
			return f.parseJPEGData(newOffset, parseSegment)
		}
		if b == 0xFF {
			return f.parseJPEGData(newOffset, sofSegment)
		}
		if b == 0xD9 {
			return f.parseJPEGData(newOffset, eioSegment)
		}
		return f.parseJPEGData(newOffset, skipSegment)
	case skipSegment:
		bytes, err := f.reader.(*xbuffer).Slice(offset+1, 2)
		if err != nil {
			return nil, err
		}
		length := readUint16(bytes)

		newOffset := offset + int(length)
		return f.parseJPEGData(newOffset, nextSegment)
	case parseSegment:
		bytes, err := f.reader.(*xbuffer).Slice(offset+4, 4)
		if err != nil {
			return nil, err
		}

		imageSize := ImageSize{}

		imageSize.Width = uint32(readUint16(bytes[2:4]))
		imageSize.Height = uint32(readUint16(bytes[0:2]))

		return &imageSize, nil
	default:
		return nil, fmt.Errorf("error parse image data")
	}
}
