package uploader

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"

	"github.com/spf13/cast"
	"github.com/tencentyun/cos-go-sdk-v5"
	"six-go/config"
)

type COS struct {
	domain    string
	secretId  string
	secretKey string
	bucket    string
	region    string
	client    *cos.Client
	err       error
}

func newCos() Uploader {
	var coss = new(COS)
	if coss.err = coss.config(); coss.err != nil {
		return coss
	}

	coss.domain = fmt.Sprintf("https://%s.cos.%s.myqcloud.com", coss.bucket, coss.region)

	parse, err := url.Parse(coss.domain)
	if err != nil {
		coss.err = err
		return coss
	}

	coss.client = cos.NewClient(&cos.BaseURL{BucketURL: parse}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  coss.secretId,
			SecretKey: coss.secretKey,
		},
	})

	return coss
}

func (c *COS) Upload(f *multipart.FileHeader) (res Result) {
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

	_, res.Error = c.client.Object.Put(context.Background(), filename, bytes.NewReader(all), nil)
	if res.Error == nil {
		res.Url = fmt.Sprintf("%s/%s", c.domain, filename)
	}
	return
}

func (c *COS) Remove(path string) error {
	_, err := c.client.Object.Delete(context.Background(), path)
	return err
}

func (c *COS) config() error {
	cosConf := config.Get("cos")
	if cosConf == nil {
		return errors.New("cos 配置 未初始化")
	}

	cosConfM := cast.ToStringMapString(cosConf)
	if v, ok := cosConfM["secret_id"]; ok {
		c.secretId = v
	} else {
		return errors.New("cos secret_id not found. ")
	}

	if v, ok := cosConfM["secret_key"]; ok {
		c.secretKey = v
	} else {
		return errors.New("cos secret_key not found. ")
	}

	if v, ok := cosConfM["bucket"]; ok {
		c.bucket = v
	} else {
		return errors.New("cos bucket not found. ")
	}

	if v, ok := cosConfM["region"]; ok {
		c.region = v
	} else {
		return errors.New("cos region not found. ")
	}

	return nil
}
