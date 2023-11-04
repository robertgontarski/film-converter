package decoder

import "fmt"

type FilePath struct {
	Path      string
	Name      string
	Extension string
}

func (f *FilePath) GetPath() string {
	return fmt.Sprintf("%s/%s.%s",
		f.Path,
		f.Name,
		f.Extension,
	)
}

func (f *FilePath) GetInputImagePath() string {
	return fmt.Sprintf("%v/%v_%%d.%v",
		f.Path,
		f.Name,
		f.Extension,
	)
}

func (f *FilePath) GetOutputImagePath(number int) string {
	return fmt.Sprintf("%v/%v_%d.%v",
		f.Path,
		f.Name,
		number,
		f.Extension,
	)
}
