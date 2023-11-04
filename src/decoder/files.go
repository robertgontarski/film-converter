package decoder

import (
	"image/png"
	"os"
)

type ConfigDecoderFile struct {
	Input  FilePath
	Output FilePath
	Frames ConfigDecoderFrames
}

func GenerateFile(config ConfigDecoderFile) error {
	if err := GenerateFrames(config.Frames); err != nil {
		return err
	}

	files, err := os.ReadDir(config.Input.Path)
	if err != nil {
		return err
	}

	allBytes, err := decodeAllFramesIntoBytes(config, files)
	if err != nil {
		return err
	}

	return writeBytesToFile(config, allBytes)
}

func decodeAllFramesIntoBytes(config ConfigDecoderFile, files []os.DirEntry) ([]byte, error) {
	var allBytes []byte

	for frameNumber := 1; frameNumber <= len(files); frameNumber++ {
		file, err := os.Open(config.Input.GetOutputImagePath(frameNumber))
		if err != nil {
			return nil, err
		}

		img, err := png.Decode(file)
		if err != nil {
			return nil, err
		}

		width := img.Bounds().Dx()
		height := img.Bounds().Dy()

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

		if err := file.Close(); err != nil {
			return nil, err
		}

		allBytes = append(allBytes, bytes...)
	}

	return allBytes, nil
}

func writeBytesToFile(config ConfigDecoderFile, allBytes []byte) error {
	file, err := os.Create(config.Output.GetPath())
	if err != nil {
		return err
	}

	_, err = file.Write(allBytes)
	if err != nil {
		return err
	}

	return file.Close()
}
