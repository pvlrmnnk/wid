package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pvlrmnnk/wid"
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

	r, err := wid.NewTarGzReader(fr)
	if err != nil {
		fail(err)
	}
	defer r.Close()

	for {
		file, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fail(err)
		}

		fmt.Println(file)
	}
}

func fail(err error) {
	log.Fatalln(err)
}
