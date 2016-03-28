package fastimage
import (
"io"
"fmt"
)

type xbuffer struct {
	r   io.Reader
	buf []byte
}

func (b *xbuffer) fill(end int) error {
	fmt.Printf("fill %+v\n",end)
	m := len(b.buf)
	if end > m {
		if end > cap(b.buf) {
			newbuf := make([]byte, end, end)
			copy(newbuf, b.buf)
			b.buf = newbuf
		} else {
			b.buf = b.buf[:end]
		}
		if n, err := io.ReadFull(b.r, b.buf[m:end]); err != nil {
			end = m + n
			b.buf = b.buf[:end]
			return err
		}
	}
	return nil
}

func (b *xbuffer) ReadAt(p []byte, off int64) (int, error) {
	o := int(off)
	end := o + len(p)
	if int64(end) != off + int64(len(p)) {
		return 0, io.ErrUnexpectedEOF
	}

	err := b.fill(end)
	return copy(p, b.buf[o:end]), err
}

func (b *xbuffer) Slice(off, n int) ([]byte, error) {
	end := off + n
	if err := b.fill(end); err != nil {
		return nil, err
	}
	return b.buf[off:end], nil
}

func newReaderAt(r io.Reader) io.ReaderAt {
	if ra, ok := r.(io.ReaderAt); ok {
		return ra
	}
	return &xbuffer{
		r:   r,
		buf: make([]byte, 0, 2),
	}
}
