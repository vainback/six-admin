package uploader

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
)

type Local struct {
	path string
}

func newLocal() Uploader {
	return &Local{path: conf().path}
}

func (l *Local) Upload(f *multipart.FileHeader) (res Result) {
	res.IsLocal = true

	ext := path.Ext(f.Filename)
	open, err := f.Open()
	if err != nil {
		res.Error = fmt.Errorf("upload err:%v", err)
		return
	}

	all, err2 := io.ReadAll(open)
	if err2 != nil {
		res.Error = fmt.Errorf("upload err:%v", err)
		return
	}

	if conf().name == "md5" {
		res.Url = filenameByMd5(all, ext)
	} else if conf().name == "times" {
		res.Url = filenameByTimes(ext)
	}
	return
}

func (l *Local) Remove(filepath string) error {
	return os.Remove(filepath)
}
