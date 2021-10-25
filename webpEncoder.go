package main

import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/chai2010/webp"
)

//Takes filepath as argument. Converts .jpeg or .png to .webp
//And saves it.
func encode(path string) {
	// Load file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
	}

	m, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Println(err)
	}

	// Encode lossless webp
	var buf bytes.Buffer
	if err := webp.Encode(&buf, m, &webp.Options{Lossless: true}); err != nil {
		log.Println(err)
	}

	// Change filename extensions to .webp
	if pos := strings.LastIndexByte(path, '.'); pos != -1 {
		path = path[:pos] + ".webp"
	}

	// Save to file
	if err := ioutil.WriteFile(path, buf.Bytes(), 0666); err != nil {
		log.Println(err)
	}

	log.Println("Encoded to webp" + path)
}

//Delete file
func remove(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Println(err)
	}
}
