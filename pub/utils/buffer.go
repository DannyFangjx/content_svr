package utils

import (
	"bytes"
	"content_svr/pub/errors"
	"io"
	"mime/multipart"
	"sync"
)

var bytesBufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func GetBytesBuffer() *bytes.Buffer {
	return bytesBufPool.Get().(*bytes.Buffer)
}

func PutBytesBuffer(buf *bytes.Buffer) {
	bytesBufPool.Put(buf)
}

func ReadFile2Buf(file multipart.File, buf *bytes.Buffer) error {
	buf.Reset()
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return errors.Wrap(err)
	}
	_, err = io.Copy(buf, file)
	return err
}
