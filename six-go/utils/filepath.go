package utils

import "os"

type FilePath struct {
	err  error
	file os.FileInfo
}

func NewFilepath(path string) *FilePath {
	stat, err := os.Stat(path)
	return &FilePath{err: err, file: stat}
}

func (f *FilePath) Exists() bool {
	return f.err == nil
}

func (f *FilePath) IsDir() bool {
	return f.Exists() && f.file.IsDir()
}

func (f *FilePath) IsFile() bool {
	return f.Exists() && !f.file.IsDir()
}
