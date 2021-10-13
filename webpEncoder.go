package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
)

//Maybe use https://github.com/h2non/bimg
//instead of github.com/chai2010/webp
func encode(path string) {
	var buf bytes.Buffer
	var data []byte
	var err error

	// Load file data
	if data, err = ioutil.ReadFile(path); err != nil {
		log.Println(err)
	}

	m, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Println(err)
	}
	// Encode lossless webp

	if err = webp.Encode(&buf, m, &webp.Options{Lossless: true}); err != nil {
		log.Println(err)
	}


	path = newFilePath(path)
	if err = ioutil.WriteFile(path, buf.Bytes(), 0666); err != nil {
		log.Println(err)
	}

	fmt.Println("Encoded to webp" + path)
}

func newFilePath(filePath string) string {
	dir, file := filepath.Split(filePath)
	if pos := strings.LastIndexByte(file, '.'); pos != -1 {
		file = file[:pos]
	}
	return dir + file + ".webp"
}
