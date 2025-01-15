package uploader

import (
	"crypto/md5"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/spf13/cast"
	"six-go/config"
)

type Uploader interface {
	Upload(f *multipart.FileHeader) Result
	Remove(path string) error
}

type Result struct {
	Error   error  `json:"error"`
	Url     string `json:"url"`
	IsLocal bool   `json:"is_local"`
}

type _config struct {
	typ  string
	name string
	path string
	err  error
}

const (
	defaultUploadType = "local"
	defaultUploadPath = "uploads"
	defaultUploadName = "md5"
)

func conf() *_config {
	var cfg = &_config{}

	m := cast.ToStringMapString(config.Get("upload"))

	typ, ok := m["type"]
	if !ok {
		cfg.typ = defaultUploadType
	} else {
		cfg.typ = typ
	}

	name, nok := m["name"]
	if !nok {
		cfg.name = defaultUploadName
	} else {
		cfg.name = name
	}

	if cfg.typ == "local" {
		if path, pok := m["path"]; !pok || path == "" {
			cfg.path = defaultUploadPath
		} else {
			cfg.path = path
		}
	}
	return cfg
}

var mode = map[string]Uploader{
	"local": newLocal(),
	"cos":   newCos(),
	"oss":   newOss(),
	"qiniu": newQiniu(),
}

func New() Uploader {
	return mode[conf().typ]
}

func filenameByMd5(f []byte, ext string) string {
	return fmt.Sprintf("%s/%x", conf().path, md5.Sum(f)) + ext
}

func filenameByTimes(ext string) string {
	timeFormat := "20060102150405.000000000"
	return conf().path + "/" + time.Now().Format(timeFormat) + ext
}
