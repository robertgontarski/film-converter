package writeFilm

import (
	"os"
	"os/exec"
	"strconv"
)

func Init() {
	outputFileName := "films/film.mp4"
	fps := 1

	cmdArgs := []string{
		"-y",
		"-framerate", strconv.Itoa(fps),
		"-i", "images/obraz_%d.png",
		"-c:v", "libx264rgb",
		//"-pix_fmt", "yuv420p",
		//"-c:v", "libx264",
		"-crf", "0",
		"-preset", "veryslow",
		"-r", strconv.Itoa(fps),
		outputFileName,
	}

	cmd := exec.Command("ffmpeg", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Uruchomienie komendy

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	creatShowFIle()
}

func creatShowFIle() {
	outputFileName := "films/film_show.mp4"
	fps := 1

	cmdArgs := []string{
		"-y",
		"-framerate", strconv.Itoa(fps),
		"-i", "images/obraz_%d.png",
		"-pix_fmt", "yuv420p",
		"-c:v", "libx264",
		"-crf", "1",
		"-preset", "veryslow",
		"-r", strconv.Itoa(fps),
		outputFileName,
	}

	cmd := exec.Command("ffmpeg", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Uruchomienie komendy

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
