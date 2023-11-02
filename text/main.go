package text

import (
	"fmt"
	"github.com/robertgontarski/gokit"
	"io/ioutil"
	"os"
)

func GetString() []byte {
	file, err := os.Open(fmt.Sprintf(
		"%s/%s.%s",
		gokit.Env["ENCODER_FILE_INPUT_FOLDER"],
		gokit.Env["ENCODER_FILE_INPUT_FILE_NAME"],
		gokit.Env["ENCODER_FILE_INPUT_FILE_EXTENSION"],
	))
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
