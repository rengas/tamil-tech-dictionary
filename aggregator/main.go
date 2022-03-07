package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir := "../letters"
	fs, err := os.ReadDir("../letters")
	if err != nil {
		panic(err)
	}

	path, _ := filepath.Abs(dir)

	for _, f := range fs {
		in, err := f.Info()
		if err != nil {
			panic(err)
		}

		f, err := os.OpenFile(path+"/"+in.Name(), os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Fatalf("open file error: %v", err)
			return
		}
		defer f.Close()

	}
}
