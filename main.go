package main

import (
	"image/read"
	"image/readFilm"
	"image/readFrames"
	"image/write"
	"image/writeFilm"
	"os"
)

func init() {
	if err := os.RemoveAll("images"); err != nil {
		panic(err)
	}
	if err := os.Mkdir("images", 0777); err != nil {
		panic(err)
	}
	if err := os.RemoveAll("films"); err != nil {
		panic(err)
	}
	if err := os.Mkdir("films", 0777); err != nil {
		panic(err)
	}
	if err := os.RemoveAll("frames"); err != nil {
		panic(err)
	}
	if err := os.Mkdir("frames", 0777); err != nil {
		panic(err)
	}
	if err := os.RemoveAll("files"); err != nil {
		panic(err)
	}
	if err := os.Mkdir("files", 0777); err != nil {
		panic(err)
	}
}

func main() {
	write.Init()
	read.Init()
	writeFilm.Init()
	readFilm.Init()
	readFrames.Init()
}
