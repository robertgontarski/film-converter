package write

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"image/text"
	"os"
)

const (
	width  = 1280
	height = 720
)

func Init() {
	bytes := text.GetString()

	countImg := len(bytes) / (width * height)
	if len(bytes)%(width*height) != 0 {
		countImg++
	}

	var slices []color.RGBA

	for i := 0; i < len(bytes); i += 3 {
		end := i + 3
		if end > len(bytes) {
			end = len(bytes)
		}
		slice := bytes[i:end]

		if len(slice) < 2 {
			slice = append(slice, []byte{0, 0}...)
		} else if len(slice) < 3 {
			slice = append(slice, []byte{0}...)
		}

		slices = append(slices, color.RGBA{
			R: slice[0],
			G: slice[1],
			B: slice[2],
			A: 255,
		})
	}

	j := 0

	for i := 0; i < countImg; i++ {
		if j >= len(slices) {
			break
		}

		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if j >= len(slices) {
					break
				}

				img.Set(x, y, slices[j])
				j++
			}
		}

		file, err := os.Create(fmt.Sprintf("images/obraz_%d.png", i))
		if err != nil {
			panic(err)
		}

		if err := png.Encode(file, img); err != nil {
			panic(err)
		}

		file.Close()
	}
}
