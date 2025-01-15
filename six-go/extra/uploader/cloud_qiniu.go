package uploader

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"path"

	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"github.com/spf13/cast"
	"six-go/config"
)

type Qiniu struct {
	accessKey string
	secretKey string
	bucket    string
	region    string
	url       string
	mac       *credentials.Credentials
	err       error
}

func newQiniu() Uploader {
	var qiniu = new(Qiniu)
	if qiniu.err = qiniu.config(); qiniu.err != nil {
		return qiniu
	}
	qiniu.mac = credentials.NewCredentials(qiniu.accessKey, qiniu.secretKey)
	return qiniu
}

func (c *Qiniu) Upload(f *multipart.FileHeader) (res Result) {
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

	client := uploader.NewUploadManager(&uploader.UploadManagerOptions{
		Options: http_client.Options{
			Credentials: c.mac,
		},
	})

	res.Error = client.UploadReader(context.Background(), bytes.NewReader(all), &uploader.ObjectOptions{
		BucketName: c.bucket,
		ObjectName: &filename,
	}, nil)
	res.Url = fmt.Sprintf("%s/%s", c.url, filename)
	return
}

func (c *Qiniu) Remove(path string) error {
	region, ok := storage.GetRegionByID(storage.RegionID(c.region))
	if !ok {
		return fmt.Errorf("qiniu region is not exists: %s", c.region)
	}
	return storage.NewBucketManager(c.mac, &storage.Config{Region: &region}).Delete(c.bucket, path)
}

func (c *Qiniu) config() error {
	cosConf := config.Get("qiniu")
	if cosConf == nil {
		return errors.New("qiniu 配置 未初始化")
	}
	cosConfM := cast.ToStringMapString(cosConf)
	if v, ok := cosConfM["access_key"]; ok {
		c.accessKey = v
	} else {
		return errors.New("qiniu access_key not found. ")
	}

	if v, ok := cosConfM["secret_key"]; ok {
		c.secretKey = v
	} else {
		return errors.New("qiniu secret_key not found. ")
	}

	if v, ok := cosConfM["bucket"]; ok {
		c.bucket = v
	} else {
		return errors.New("qiniu bucket not found. ")
	}

	if v, ok := cosConfM["url"]; ok {
		c.url = v
	} else {
		return errors.New("qiniu region not found. ")
	}

	if v, ok := cosConfM["region"]; ok {
		c.region = v
	}

	return nil
}
