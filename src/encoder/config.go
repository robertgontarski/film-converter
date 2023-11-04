package encoder

import "github.com/robertgontarski/gokit"

func GetImagesConfig() ConfigEncoderImages {
	return ConfigEncoderImages{
		Width:  gokit.Env["ENCODER_FRAME_WIDTH"],
		Height: gokit.Env["ENCODER_FRAME_HEIGHT"],
		Input: FilePath{
			Path:      gokit.Env["ENCODER_FILE_INPUT_FOLDER"],
			Name:      gokit.Env["ENCODER_FILE_INPUT_FILE_NAME"],
			Extension: gokit.Env["ENCODER_FILE_INPUT_FILE_EXTENSION"],
		},
		Output: FilePath{
			Path:      gokit.Env["ENCODER_FRAME_OUTPUT_FOLDER"],
			Name:      gokit.Env["ENCODER_FRAME_OUTPUT_FILE_NAME"],
			Extension: gokit.Env["ENCODER_FRAME_OUTPUT_FILE_EXTENSION"],
		},
	}
}

func GetFilmConfig() ConfigEncoderFilms {
	return ConfigEncoderFilms{
		Film: ConfigFilm{
			Output: FilePath{
				Path:      gokit.Env["ENCODER_MOVIE_OUTPUT_FOLDER"],
				Name:      gokit.Env["ENCODER_MOVIE_OUTPUT_FILE_NAME"],
				Extension: gokit.Env["ENCODER_MOVIE_OUTPUT_FILE_EXTENSION"],
			},
			Input: FilePath{
				Path:      gokit.Env["ENCODER_MOVIE_FRAMES_INPUT_FOLDER"],
				Name:      gokit.Env["ENCODER_MOVIE_FRAMES_INPUT_FILE_NAME"],
				Extension: gokit.Env["ENCODER_MOVIE_FRAMES_INPUT_FILE_EXTENSION"],
			},
			Fps:    gokit.Env["ENCODER_MOVIE_FPS"],
			Codec:  gokit.Env["ENCODER_MOVIE_CODEC"],
			Crf:    gokit.Env["ENCODER_MOVIE_CRF"],
			Preset: gokit.Env["ENCODER_MOVIE_PRESET"],
		},
		IsPreview: true,
		Preview: ConfigFilm{
			Output: FilePath{
				Path:      gokit.Env["ENCODER_MOVIE_PREVIEW_OUTPUT_FOLDER"],
				Name:      gokit.Env["ENCODER_MOVIE_PREVIEW_OUTPUT_FILE_NAME"],
				Extension: gokit.Env["ENCODER_MOVIE_PREVIEW_OUTPUT_FILE_EXTENSION"],
			},
			Input: FilePath{
				Path:      gokit.Env["ENCODER_MOVIE_FRAMES_INPUT_FOLDER"],
				Name:      gokit.Env["ENCODER_MOVIE_FRAMES_INPUT_FILE_NAME"],
				Extension: gokit.Env["ENCODER_MOVIE_FRAMES_INPUT_FILE_EXTENSION"],
			},
			Fps:    gokit.Env["ENCODER_MOVIE_PREVIEW_FPS"],
			Format: gokit.Env["ENCODER_MOVIE_PREVIEW_PICTURE_FORMAT"],
			Codec:  gokit.Env["ENCODER_MOVIE_PREVIEW_CODEC"],
			Crf:    gokit.Env["ENCODER_MOVIE_PREVIEW_CRF"],
			Preset: gokit.Env["ENCODER_MOVIE_PREVIEW_PRESET"],
		},
		Images: GetImagesConfig(),
	}
}
