package decoder

import "github.com/robertgontarski/gokit"

func GetFrameConfig() ConfigDecoderFrames {
	return ConfigDecoderFrames{
		Input: FilePath{
			Path:      gokit.Env["DECODER_MOVIE_INPUT_FOLDER"],
			Name:      gokit.Env["DECODER_MOVIE_INPUT_FILE_NAME"],
			Extension: gokit.Env["DECODER_MOVIE_INPUT_FILE_EXTENSION"],
		},
		Output: FilePath{
			Path:      gokit.Env["DECODER_MOVIE_OUTPUT_FOLDER"],
			Name:      gokit.Env["DECODER_MOVIE_OUTPUT_FILE_NAME"],
			Extension: gokit.Env["DECODER_MOVIE_OUTPUT_FILE_EXTENSION"],
		},
		Fps: gokit.Env["DECODER_MOVIE_INPUT_FILE_FPS"],
	}
}

func GetFileConfig() ConfigDecoderFile {
	return ConfigDecoderFile{
		Input: FilePath{
			Path:      gokit.Env["DECODER_FRAME_INPUT_FOLDER"],
			Name:      gokit.Env["DECODER_FRAME_INPUT_FILE_NAME"],
			Extension: gokit.Env["DECODER_FRAME_INPUT_FILE_EXTENSION"],
		},
		Output: FilePath{
			Path:      gokit.Env["DECODER_FILE_OUTPUT_FOLDER"],
			Name:      gokit.Env["DECODER_FILE_OUTPUT_FILE_NAME"],
			Extension: gokit.Env["DECODER_FILE_OUTPUT_FILE_EXTENSION"],
		},
		Frames: GetFrameConfig(),
	}
}
