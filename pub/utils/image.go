package utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/disintegration/imaging"
	"io/ioutil"
	"os"
)

func ResizeImage(ctx context.Context, fileName string, buf *bytes.Buffer) (oriFile, file *os.File, err error) {
	imageFile, err := imaging.Decode(buf)
	if err != nil {
		return
	}
	//生成缩略图，128*128
	resizeFile := imaging.Resize(imageFile, 128, 128, imaging.Lanczos)
	cwd, err := os.Getwd()
	if err != nil {
		return
	}
	file, err = ioutil.TempFile(cwd, fmt.Sprintf("tem-*-%v", fileName))
	if err != nil {
		return
	}
	err = imaging.Save(resizeFile, file.Name())
	if err != nil {
		return
	}
	oriFile, err = ioutil.TempFile(cwd, fmt.Sprintf("tem-*-%v", fileName))
	if err != nil {
		return
	}
	err = imaging.Save(imageFile, oriFile.Name())
	if err != nil {
		return
	}
	return
}
