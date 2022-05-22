package wid

import "io"

type Reader interface {
	io.Closer
	Next() (string, error)
}
