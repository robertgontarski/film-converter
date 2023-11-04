package decoder

import (
	"fmt"
	"os/exec"
)

type ConfigDecoderFrames struct {
	Input  FilePath
	Output FilePath
	Fps    string
}

func GenerateFrames(config ConfigDecoderFrames) error {
	return executeCmd(prepareCmdArguments(config))
}

func prepareCmdArguments(config ConfigDecoderFrames) []string {
	return []string{
		"-i", config.Input.GetPath(),
		"-vf",
		fmt.Sprintf("fps=%s", config.Fps),
		config.Output.GetInputImagePath(),
	}
}

func executeCmd(arguments []string) error {
	cmd := exec.Command("ffmpeg", arguments...)
	return cmd.Run()
}
