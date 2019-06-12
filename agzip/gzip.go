package agzip

import (
	"bytes"
	"compress/gzip"
	"io"
)

//做gzip压缩
func Gzip(data []byte) []byte {
	var buf bytes.Buffer
	zip := gzip.NewWriter(&buf)
	_, err := zip.Write(data)
	if err != nil {
		return nil
	}
	zip.Close()

	return buf.Bytes()
}

//做gzip解压缩
func UnGzip(data []byte) []byte {
	var buf bytes.Buffer
	content := bytes.NewReader(data)
	zipdata, err := gzip.NewReader(content)
	if err != nil {
		return nil
	}
	io.Copy(&buf, zipdata)
	zipdata.Close()
	return buf.Bytes()
}
