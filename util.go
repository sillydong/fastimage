package fastimage

import "encoding/binary"

func readUint16(buffer []byte) uint16 {
	return binary.BigEndian.Uint16(buffer)
}

func readULint16(buffer []byte) uint16 {
	return binary.LittleEndian.Uint16(buffer)
}

func readUint32(b []byte) uint32 {
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}
