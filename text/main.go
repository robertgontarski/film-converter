package text

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetString() []byte {
	// Open the file
	file, err := os.Open("content.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return data
}
