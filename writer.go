package seriesrw

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"
)

type BinaryReader interface {
	Read(...interface{}) error
}

type BinaryWriter interface {
	Write(...interface{}) error
}

func NewFileWriter(path string, size int) (*FileWriter, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewWriterSize(f, size)
	return &FileWriter{BinaryWriter: &BinaryReadWriter{w: buf}, f: f, buf: buf}, nil
}

type FileWriter struct {
	BinaryWriter
	f   *os.File
	buf *bufio.Writer
}

func (f *FileWriter) Close() error {
	if err := f.buf.Flush(); err != nil {
		return err
	}
	if err := f.f.Sync(); err != nil {
		return err
	}
	// should close in any case
	return f.f.Close()
}

type BinaryReadWriter struct {
	w io.Writer
	r io.Reader
}

func (w *BinaryReadWriter) Write(items ...interface{}) error {
	for _, data := range items {
		if err := binary.Write(w.w, binary.BigEndian, data); err != nil {
			return err
		}
	}
	return nil
}

func (w *BinaryReadWriter) Read(items ...interface{}) error {
	for _, data := range items {
		if err := binary.Read(w.r, binary.BigEndian, data); err != nil {
			return err
		}
	}
	return nil
}
