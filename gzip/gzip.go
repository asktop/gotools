package gzip

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"io"
	"io/ioutil"
)

//GZIP压缩
func Gzip(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w, err := gzip.NewWriterLevel(&b, 9)
	if err != nil {
		return nil, err
	}
	defer w.Close()
	_, err = w.Write(data)
	if err != nil {
		return nil, err
	}
	err = w.Flush()
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

//GZIP解压缩
func UnGzip(data []byte) ([]byte, error) {
	b := new(bytes.Buffer)
	err := binary.Write(b, binary.LittleEndian, data)
	if err != nil {
		return nil, err
	}
	r, err := gzip.NewReader(b)
	if err != nil && err != io.EOF {
		return nil, err
	}
	defer r.Close()
	undata, err := ioutil.ReadAll(r)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return undata, nil
}
