package read

import (
	"fmt"
	"github.com/robertgontarski/gokit"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

var width int
var height int

func Init() {
	files, err := ioutil.ReadDir(gokit.Env["DECODER_FRAME_INPUT_FOLDER"])
	if err != nil {
		log.Fatal(err)
	}

	var allBytes []byte

	for i := 0; i < len(files); i++ {
		file, err := os.Open(fmt.Sprintf(
			"%v/%v_%d.%v",
			gokit.Env["DECODER_FRAME_INPUT_FOLDER"],
			gokit.Env["DECODER_FRAME_INPUT_FILE_NAME"],
			i,
			gokit.Env["DECODER_FRAME_INPUT_FILE_EXTENSION"],
		))
		if err != nil {
			panic(err)
		}

		img, err := png.Decode(file)
		if err != nil {
			panic(err)
		}

		width = img.Bounds().Dx()
		height = img.Bounds().Dy()

		var bytes []byte
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				r, g, b, _ := img.At(x, y).RGBA()
				r = r >> 8
				g = g >> 8
				b = b >> 8

				if r == 0 && g == 0 && b == 0 {
					break
				}

				bytes = append(bytes, byte(r), byte(g), byte(b))
			}
		}

		file.Close()
		allBytes = append(allBytes, bytes...)
	}

	file, err := os.Create(fmt.Sprintf(
		"%v/%v.%v",
		gokit.Env["DECODER_FILE_OUTPUT_FOLDER"],
		gokit.Env["DECODER_FILE_OUTPUT_FILE_NAME"],
		gokit.Env["DECODER_FILE_OUTPUT_FILE_EXTENSION"],
	))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(allBytes)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
