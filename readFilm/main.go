package readFilm

import (
	"log"
	"os/exec"
)

func Init() {
	videoPath := "films/film.mp4"
	outputPath := "frames/frame_%d.png"

	// Ustawianie argumentów dla ffmpeg
	cmdArgs := []string{
		"-i", videoPath,
		"-vf", "fps=1",
		outputPath,
	}

	cmd := exec.Command("ffmpeg", cmdArgs...)
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error while generating frames: %v", err)
	}
}
