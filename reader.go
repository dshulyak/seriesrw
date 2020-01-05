package seriesrw

import (
	"bufio"
	"os"
)

func NewFileReader(path string, size int) (*FileReader, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewReaderSize(f, size)
	return &FileReader{BinaryReader: &BinaryReadWriter{r: buf}, buf: buf, f: f}, nil
}

type FileReader struct {
	BinaryReader
	buf *bufio.Reader
	f   *os.File
}

func (f *FileReader) Close() error {
	f.buf.Reset(nil)
	return f.f.Close()
}
