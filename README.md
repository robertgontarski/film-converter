[Film converter](https://github.com/robertgontarski/film-converter) is the concept of unlimited cloud storage based on youtube private videos

Script is currently in alpha stage and still needs a lot of development

## Usage

To start using app, clone project

```shell
git clone git@github.com:robertgontarski/film-converter.git
```

uncomment or comment out the elements of the application that are not supposed to work in the main.go file

```go
package main

func main() {
	write.Init()
	read.Init()
	writeFilm.Init()
	readFilm.Init()
	readFrames.Init()
}
```

run docker config to run main.go file

```shell
docker compose up
```

## API Reference

### write

This part is responsible for converting the content.txt text file into images for the images folder. You can set the size and width of the generated images in the file

### read 

This part is responsible for processing the files generated and set in the images folder. The end result will be the generation of a file in the files folder based on the images

### writeFilm

One will be with the title of the show which will go bad for previewing, while the other will be a video for further processing without encoding and compression. You can find the generated movies in the movie folder. The standard fps is set to 1s

### readFilm

The process processes the video into individual frames (photos) saves them to the frames folder. The standard fps is set to 1s

### readFrames

Process identical to read only it returns the result to the console 

## What's next - goals

### more formats 

Adding support for more file formats. Adding recognition of automatic file format by the application. 

### plug&paly 

Adding environment configuration so that you don't need to have it on your device 

### interface

Adding a console interface for simple configuration. The next step will be to add gui for users 

### more configs

Adding the ability to configure the operation of the entire application with a single end / yml file

### standards and code quality 

Rewriting the code to comply with the writing methodology and improving the readability of the code

### and much more 

The list is constantly being updated with new ideas if you too have an idea boldly feel free to [write to me ](https://www.linkedin.com/in/robert-gontarski/)

## Patch notes

- [02-11-2023](.doc/patchNotes/02-11-2023.md)

## License

[MIT](LICENSE)