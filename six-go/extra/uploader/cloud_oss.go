package uploader

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"mime/multipart"
	"path"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/cast"
	"six-go/config"
)

type OSS struct {
	// domain          string
	accessKeyId     string
	accessKeySecret string
	bucket          string
	endpoint        string
	client          *oss.Client
	bucketObj       *oss.Bucket
	err             error
}

func newOss() Uploader {
	var osss = new(OSS)
	if osss.err = osss.config(); osss.err != nil {
		return osss
	}

	osss.client, osss.err = oss.New(osss.endpoint, osss.accessKeyId, osss.accessKeySecret)
	if osss.err != nil {
		return osss
	}
	osss.bucketObj, osss.err = osss.client.Bucket(osss.bucket)
	if osss.err != nil {
		return osss
	}
	return osss
}

func (c *OSS) Upload(f *multipart.FileHeader) (res Result) {
	if c.err != nil {
		res.Error = fmt.Errorf("upload err:%v", c.err)
		return
	}

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

	var filename string
	if conf().name == "md5" {
		filename = filenameByMd5(all, ext)
	} else if conf().name == "times" {
		filename = filenameByTimes(ext)
	}
	if res.Error = c.bucketObj.PutObject(filename, bytes.NewReader(all)); res.Error != nil {
		return
	}
	res.Url, res.Error = c.bucketObj.SignURL(filename, oss.HTTPGet, math.MaxInt)
	return
}

func (c *OSS) Remove(path string) error {
	return c.bucketObj.DeleteObject(path)
}

func (c *OSS) config() error {
	cosConf := config.Get("oss")
	if cosConf == nil {
		return errors.New("oss 配置 未初始化")
	}
	cosConfM := cast.ToStringMapString(cosConf)
	if v, ok := cosConfM["access_key_id"]; ok {
		c.accessKeyId = v
	} else {
		return errors.New("oss secret_id not found. ")
	}

	if v, ok := cosConfM["access_key_secret"]; ok {
		c.accessKeySecret = v
	} else {
		return errors.New("oss secret_key not found. ")
	}

	if v, ok := cosConfM["bucket"]; ok {
		c.bucket = v
	} else {
		return errors.New("oss bucket not found. ")
	}

	if v, ok := cosConfM["endpoint"]; ok {
		c.endpoint = v
	} else {
		return errors.New("oss endpoint not found. ")
	}

	return nil
}
