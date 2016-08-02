package fastimage

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

func (f *decoder) getTIFFImageSize() (*ImageSize, error) {
	p := make([]byte, 8)
	if _, err := f.reader.ReadAt(p, 0); err != nil {
		return nil, err
	}

	var bo binary.ByteOrder
	switch string(p[0:4]) {
	case leHeader:
		bo = binary.LittleEndian
	case beHeader:
		bo = binary.BigEndian
	default:
		return nil, fmt.Errorf("malformed header")
	}
	ifdOffset := int64(bo.Uint32(p[4:8]))

	// The first two bytes contain the number of entries (12 bytes each).
	if _, err := f.reader.ReadAt(p[0:2], ifdOffset); err != nil {
		return nil, err
	}
	numItems := int(bo.Uint16(p[0:2]))
	// All IFD entries are read in one chunk.
	p = make([]byte, ifdLen*numItems)
	if _, err := f.reader.ReadAt(p, ifdOffset+2); err != nil {
		return nil, err
	}

	imageSize := ImageSize{}
	for i := 0; i < len(p); i += ifdLen {
		t := p[i : i+ifdLen]
		tag := bo.Uint16(t[0:2])
		switch tag {
		case tImageWidth:
			val, _ := ifdUint(f.reader, bo, t)
			imageSize.Width = uint32(val[0])
		case tImageLength:
			val, _ := ifdUint(f.reader, bo, t)
			imageSize.Height = uint32(val[0])
		default:
			break
		}
	}

	return &imageSize, nil
}

var lengths = [...]uint32{0, 1, 1, 2, 4, 8}

func ifdUint(reader io.ReaderAt, bo binary.ByteOrder, p []byte) (u []uint, err error) {
	var raw []byte
	if len(p) < ifdLen {
		return nil, fmt.Errorf("bad IFD entry")
	}

	datatype := bo.Uint16(p[2:4])
	if dt := int(datatype); dt <= 0 || dt >= len(lengths) {
		return nil, fmt.Errorf("IFD entry datatype")
	}

	count := bo.Uint32(p[4:8])
	if count > math.MaxInt32/lengths[datatype] {
		return nil, fmt.Errorf("IFD data too large")
	}
	if datalen := lengths[datatype] * count; datalen > 4 {
		// The IFD contains a pointer to the real value.
		raw = make([]byte, datalen)
		_, err = reader.ReadAt(raw, int64(bo.Uint32(p[8:12])))
	} else {
		raw = p[8 : 8+datalen]
	}
	if err != nil {
		return nil, err
	}

	u = make([]uint, count)
	switch datatype {
	case dtByte:
		for i := uint32(0); i < count; i++ {
			u[i] = uint(raw[i])
		}
	case dtShort:
		for i := uint32(0); i < count; i++ {
			u[i] = uint(bo.Uint16(raw[2*i : 2*(i+1)]))
		}
	case dtLong:
		for i := uint32(0); i < count; i++ {
			u[i] = uint(bo.Uint32(raw[4*i : 4*(i+1)]))
		}
	default:
		return nil, fmt.Errorf("data type")
	}
	return u, nil
}

const (
	tImageWidth  = 256
	tImageLength = 257

	leHeader = "II\x2A\x00" // Header for little-endian files.
	beHeader = "MM\x00\x2A" // Header for big-endian files.

	ifdLen     = 12 // Length of an IFD entry in bytes.
	dtByte     = 1
	dtASCII    = 2
	dtShort    = 3
	dtLong     = 4
	dtRational = 5
)
