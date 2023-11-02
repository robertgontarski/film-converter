package writeFilm

import (
	"fmt"
	"github.com/robertgontarski/gokit"
	"os"
	"os/exec"
)

func Init() {
	cmdArgs := []string{
		"-y",
		"-framerate", gokit.Env["ENCODER_MOVIE_FPS"],
		"-i", fmt.Sprintf("%v/%v_%%d.%v",
			gokit.Env["ENCODER_MOVIE_FRAMES_INPUT_FOLDER"],
			gokit.Env["ENCODER_MOVIE_FRAMES_INPUT_FILE_NAME"],
			gokit.Env["ENCODER_MOVIE_FRAMES_INPUT_FILE_EXTENSION"],
		),
		"-c:v", gokit.Env["ENCODER_MOVIE_CODEC"],
		"-crf", gokit.Env["ENCODER_MOVIE_CRF"],
		"-preset", gokit.Env["ENCODER_MOVIE_PRESET"],
		"-r", gokit.Env["ENCODER_MOVIE_FPS"],
		fmt.Sprintf(
			"%v/%v.%v",
			gokit.Env["ENCODER_MOVIE_OUTPUT_FOLDER"],
			gokit.Env["ENCODER_MOVIE_OUTPUT_FILE_NAME"],
			gokit.Env["ENCODER_MOVIE_OUTPUT_FILE_EXTENSION"],
		),
	}

	fmt.Println(cmdArgs)

	cmd := exec.Command("ffmpeg", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	creatShowFIle()
}

func creatShowFIle() {
	cmdArgs := []string{
		"-y",
		"-framerate", gokit.Env["ENCODER_MOVIE_FPS"],
		"-i", fmt.Sprintf("%v/%v_%%d.%v",
			gokit.Env["ENCODER_MOVIE_FRAMES_INPUT_FOLDER"],
			gokit.Env["ENCODER_MOVIE_FRAMES_INPUT_FILE_NAME"],
			gokit.Env["ENCODER_MOVIE_FRAMES_INPUT_FILE_EXTENSION"],
		),
		"-pix_fmt", gokit.Env["ENCODER_MOVIE_PREVIEW_PICTURE_FORMAT"],
		"-c:v", gokit.Env["ENCODER_MOVIE_PREVIEW_CODEC"],
		"-crf", gokit.Env["ENCODER_MOVIE_PREVIEW_CRF"],
		"-preset", gokit.Env["ENCODER_MOVIE_PREVIEW_PRESET"],
		"-r", gokit.Env["ENCODER_MOVIE_PREVIEW_FPS"],
		fmt.Sprintf(
			"%v/%v.%v",
			gokit.Env["ENCODER_MOVIE_PREVIEW_OUTPUT_FOLDER"],
			gokit.Env["ENCODER_MOVIE_PREVIEW_OUTPUT_FILE_NAME"],
			gokit.Env["ENCODER_MOVIE_PREVIEW_OUTPUT_FILE_EXTENSION"],
		),
	}

	cmd := exec.Command("ffmpeg", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
