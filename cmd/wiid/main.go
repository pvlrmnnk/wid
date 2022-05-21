package main

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fail(errors.New("wrong usage"))
	}

	fn := os.Args[1]
	log.Println("File: ", fn)

	fr, err := os.Open(fn)
	if err != nil {
		fail(err)
	}
	defer fr.Close()

	gzr, err := gzip.NewReader(fr)
	if err != nil {
		fail(err)
	}
	defer gzr.Close()

	list := make([]string, 0, 1000)

	tr := tar.NewReader(gzr)
	for {
		th, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fail(err)
		}
		log.Println(th)
		list = append(list, th.Name)
	}

	log.Println(list)
}

func fail(err error) {
	log.Fatalln(err)
}
