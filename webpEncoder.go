package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	_ "image/jpeg"
	"io/ioutil"
	"log"

	"github.com/chai2010/webp"
)

func encode() {
	var buf bytes.Buffer
	var data []byte
	var err error

	// Load file data
	if data, err = ioutil.ReadFile("./testdata/1.png"); err != nil {
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

	if err = ioutil.WriteFile("./testdata/1.webp", buf.Bytes(), 0666); err != nil {
		log.Println(err)
	}

	fmt.Println("Saved ./testdata/1.webp ok")
}
