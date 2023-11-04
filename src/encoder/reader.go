package encoder

import (
	"io"
	"os"
)

func readFile(path FilePath) ([]byte, error) {
	file, err := os.Open(path.GetPath())
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := file.Close(); err != nil {
		return nil, err
	}

	return data, nil
}
