package setup

import (
	"github.com/robertgontarski/gokit"
	"os"
)

func Run() {
	if err := os.MkdirAll(gokit.Env["ENCODER_FILE_INPUT_FOLDER"], 0777); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(gokit.Env["ENCODER_FRAME_OUTPUT_FOLDER"], 0777); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(gokit.Env["ENCODER_MOVIE_OUTPUT_FOLDER"], 0777); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(gokit.Env["ENCODER_MOVIE_PREVIEW_OUTPUT_FOLDER"], 0777); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(gokit.Env["ENCODER_MOVIE_FRAMES_INPUT_FOLDER"], 0777); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(gokit.Env["DECODER_MOVIE_OUTPUT_FOLDER"], 0777); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(gokit.Env["DECODER_FILE_OUTPUT_FOLDER"], 0777); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(gokit.Env["DECODER_MOVIE_INPUT_FOLDER"], 0777); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(gokit.Env["DECODER_FRAME_INPUT_FOLDER"], 0777); err != nil {
		panic(err)
	}
}
