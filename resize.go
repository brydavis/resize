package resize

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func ResizePixels(reader io.Reader, width, height int) (*image.RGBA, error) {
	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Println(err)
	}

	dx := float64(width) / float64(img.Bounds().Max.X)
	dy := float64(height) / float64(img.Bounds().Max.Y)

	rect := image.Rect(0, 0, width, height)
	i := image.NewRGBA(rect)

	for iy := 0; iy <= rect.Max.Y; iy += 1 {
		for ix := 0; ix <= rect.Max.X; ix += 1 {
			i.Set(ix, iy, img.At(int(float64(ix)/dx), int(float64(iy)/dy)))
		}
	}

	return i, nil
}

func ResizeFolderPercent(dir string, percent float64, prefix string) {
	d, _ := ioutil.ReadDir(dir)
	for _, item := range d {
		filename := item.Name()
		file, _ := os.Open(filename)
		ext := path.Ext(filename)

		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			fmt.Printf("File not an image: %s\n", filename)
		}

		img, err := ResizePercent(file, percent)
		if err != nil {
			fmt.Println(err)
		} else {
			new_filename := fmt.Sprintf("%s_%s", prefix, filename)
			if err := WriteToFile(new_filename, img); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func ResizeFolderPixels(dir string, width, height int, prefix string) {
	d, _ := ioutil.ReadDir(dir)
	for _, item := range d {
		filename := item.Name()
		file, _ := os.Open(filename)
		ext := path.Ext(filename)

		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			fmt.Printf("File not an image: %s\n", filename)
		} else {
			img, err := ResizePixels(file, width, height)
			if err != nil {
				fmt.Println(err)
			} else {
				new_filename := fmt.Sprintf("%s_%s", prefix, filename)
				if err := WriteToFile(new_filename, img); err != nil {
					fmt.Println(err)
				}
			}
		}

	}

}

func WriteToFile(filename string, img *image.RGBA) error {
	out, _ := os.Create(filename)

	switch ext := path.Ext(filename); ext {
	case ".jpg", ".jpeg":
		var opt jpeg.Options
		opt.Quality = 1000
		jpeg.Encode(out, img, &opt)
	case ".png":
		png.Encode(out, img)
	default:
		return errors.New("Unable to encode image")
	}
	return nil
}

func ResizePercent(reader io.Reader, percent float64) (*image.RGBA, error) {
	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Println(err)
	}

	rect := image.Rect(0, 0, int(float64(img.Bounds().Max.X)*percent), int(float64(img.Bounds().Max.Y)*percent))
	out := image.NewRGBA(rect)

	for iy := 0; iy <= rect.Max.Y; iy += 1 {
		for ix := 0; ix <= rect.Max.X; ix += 1 {
			out.Set(ix, iy, img.At(int(float64(ix)/percent), int(float64(iy)/percent)))
		}
	}

	return out, nil
}
