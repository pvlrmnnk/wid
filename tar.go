package wid

import (
	"archive/tar"
	"compress/gzip"
	"io"
)

type tarGzReader struct {
	gzr *gzip.Reader
	tr  *tar.Reader
}

func (tgzr *tarGzReader) Close() error {
	return tgzr.gzr.Close()
}

func (tgzr *tarGzReader) Next() (string, error) {
	th, err := tgzr.tr.Next()
	if err != nil {
		return "", err
	}
	if th == nil {
		return "", io.EOF
	}

	return th.Name, nil
}

func NewTarGzReader(r io.Reader) (Reader, error) {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	return &tarGzReader{gzr, tar.NewReader(gzr)}, nil
}
