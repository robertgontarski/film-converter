package readFilm

import (
	"fmt"
	"github.com/robertgontarski/gokit"
	"os/exec"
)

func Init() {
	cmdArgs := []string{
		"-i", fmt.Sprintf(
			"%s/%s.%s",
			gokit.Env["DECODER_MOVIE_INPUT_FOLDER"],
			gokit.Env["DECODER_MOVIE_INPUT_FILE_NAME"],
			gokit.Env["DECODER_MOVIE_INPUT_FILE_EXTENSION"],
		),
		"-vf",
		fmt.Sprintf("fps=%s", gokit.Env["DECODER_MOVIE_INPUT_FILE_FPS"]),
		fmt.Sprintf("%s/%s_%%d.%s",
			gokit.Env["DECODER_MOVIE_OUTPUT_FOLDER"],
			gokit.Env["DECODER_MOVIE_OUTPUT_FILE_NAME"],
			gokit.Env["DECODER_MOVIE_OUTPUT_FILE_EXTENSION"],
		),
	}

	cmd := exec.Command("ffmpeg", cmdArgs...)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
