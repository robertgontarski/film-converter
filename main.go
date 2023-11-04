package main

import (
	"github.com/robertgontarski/film-converter/src/setup"
)

func init() {
	setup.Run()
}

func main() {
	//if err := encoder.GenerateImages(encoder.GetImagesConfig()); err != nil {
	//	gokit.ErrorLogger.Println(err)
	//	fmt.Println("App failed to generate images. Please check the logs for more information.")
	//	return
	//}
	//
	//if err := encoder.GenerateFilm(encoder.GetFilmConfig()); err != nil {
	//	gokit.ErrorLogger.Println(err)
	//	fmt.Println("App failed to generate film. Please check the logs for more information.")
	//	return
	//}
	//
	//if err := decoder.GenerateFrames(decoder.GetFrameConfig()); err != nil {
	//	gokit.ErrorLogger.Println(err)
	//	fmt.Println("App failed to generate frames. Please check the logs for more information.")
	//	return
	//}
	//
	//if err := decoder.GenerateFile(decoder.GetFileConfig()); err != nil {
	//	gokit.ErrorLogger.Println(err)
	//	fmt.Println("App failed to generate file. Please check the logs for more information.")
	//	return
	//}
}
