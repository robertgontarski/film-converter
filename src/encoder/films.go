package encoder

import (
	"os"
	"os/exec"
)

type ConfigEncoderFilms struct {
	Film      ConfigFilm
	IsPreview bool
	Preview   ConfigFilm
	Images    ConfigEncoderImages
}

type ConfigFilm struct {
	Output FilePath
	Input  FilePath
	Fps    string
	Format string
	Codec  string
	Crf    string
	Preset string
}

func GenerateFilm(config ConfigEncoderFilms) error {
	if err := GenerateImages(config.Images); err != nil {
		return err
	}

	if err := executeCmd(prepareFilmCmdArguments(config.Film)); err != nil {
		return err
	}

	if config.IsPreview == false {
		return nil
	}

	if err := executeCmd(prepareFilmPreviewCmdArguments(config.Preview)); err != nil {
		return err
	}

	return nil
}

func prepareFilmCmdArguments(config ConfigFilm) []string {
	return []string{
		"-y",
		"-framerate", config.Fps,
		"-i", config.Input.GetInputImagePath(),
		"-c:v", config.Codec,
		"-crf", config.Crf,
		"-preset", config.Preset,
		"-r", config.Fps,
		config.Output.GetPath(),
	}
}

func prepareFilmPreviewCmdArguments(config ConfigFilm) []string {
	return []string{
		"-y",
		"-framerate", config.Fps,
		"-i", config.Input.GetInputImagePath(),
		"-pix_fmt", config.Format,
		"-c:v", config.Codec,
		"-crf", config.Crf,
		"-preset", config.Preset,
		"-r", config.Fps,
		config.Output.GetPath(),
	}

}

func executeCmd(arguments []string) error {
	cmd := exec.Command("ffmpeg", arguments...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
