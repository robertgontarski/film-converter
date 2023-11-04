package encoder

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

type ConfigEncoderImages struct {
	Width  string
	Height string
	Input  FilePath
	Output FilePath
}

func (c *ConfigEncoderImages) GetWidth() int {
	width, _ := strconv.Atoi(c.Width)
	return width
}

func (c *ConfigEncoderImages) GetHeight() int {
	height, _ := strconv.Atoi(c.Height)
	return height
}

func GenerateImages(config ConfigEncoderImages) error {

	fileBytes, err := readFile(config.Input)
	if err != nil {
		return err
	}

	lenFileBytes := len(fileBytes)

	countImg := lenFileBytes / (config.GetWidth() * config.GetHeight())
	if lenFileBytes%(config.GetWidth()*config.GetHeight()) != 0 {
		countImg++
	}

	return compileArrayColorsIntoImage(config, prepareColorMaps(fileBytes), countImg)
}

func prepareColorMaps(fileBytes []byte) []color.RGBA {
	lenFileBytes := len(fileBytes)
	var arrayColors []color.RGBA

	for i := 0; i < lenFileBytes; i += 3 {
		end := i + 3
		if end > lenFileBytes {
			end = lenFileBytes
		}
		slice := fileBytes[i:end]

		if len(slice) < 2 {
			slice = append(slice, []byte{0, 0}...)
		} else if len(slice) < 3 {
			slice = append(slice, []byte{0}...)
		}

		arrayColors = append(arrayColors, color.RGBA{
			R: slice[0],
			G: slice[1],
			B: slice[2],
			A: 255,
		})
	}

	return arrayColors
}

func compileArrayColorsIntoImage(config ConfigEncoderImages, arrayColors []color.RGBA, totalNumberImages int) error {
	lenArrayColors := len(arrayColors)
	colorIndex := 0

	for imageNumber := 0; imageNumber < totalNumberImages; imageNumber++ {
		if colorIndex >= lenArrayColors {
			break
		}

		img := image.NewRGBA(image.Rect(0, 0, config.GetWidth(), config.GetHeight()))

		for positionY := 0; positionY < config.GetHeight(); positionY++ {
			for positionX := 0; positionX < config.GetWidth(); positionX++ {
				if colorIndex >= lenArrayColors {
					break
				}

				img.Set(positionX, positionY, arrayColors[colorIndex])
				colorIndex++
			}
		}

		file, err := os.Create(config.Output.GetOutputImagePath(imageNumber))
		if err != nil {
			return err
		}

		if err := png.Encode(file, img); err != nil {
			return err
		}

		if err := file.Close(); err != nil {
			return err
		}
	}

	return nil
}
