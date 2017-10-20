// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package preview

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"path"

	"github.com/Arafatk/glot"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

// Color will open a viewer for the color.Color specified by img.
// The previewed image will be 256x256 pixels in size, completely filled with the given color.
func Color(clr color.Color) error {
	img := image.NewRGBA(image.Rect(0, 0, 256, 256))
	draw.Draw(img, img.Bounds(), &image.Uniform{clr}, image.ZP, draw.Src)

	return Image(img)
}

// File will open a viewer for the path specified by filename.
func File(filename string) error {
	return view(filename)
}

// Glot will open a viewer for the glot.Plot specified by plt.
func Glot(plt *glot.Plot) error {

	filename, err := tempFile()
	if err != nil {
		return err
	}

	if err := plt.SavePlot(filename); err != nil {
		return err
	}

	return File(filename)
}

// Gonum will open a viewer for the plot.Plot specified by plt.
// The previewed image will be 4x4 inches in size.
func Gonum(plt *plot.Plot) error {

	filename, err := tempFile()
	if err != nil {
		return err
	}

	if err := plt.Save(4*vg.Inch, 4*vg.Inch, filename); err != nil {
		return err
	}

	return File(filename)
}

// Image will open a viewer for the image.Image specified by img.
func Image(img image.Image) error {

	filename, err := tempFile()
	if err != nil {
		return err
	}

	if err := render(img, filename); err != nil {
		return err
	}

	return File(filename)
}

// Show will open a viewer for the type specified by any.
func Show(any interface{}) error {
	switch data := any.(type) {
	case color.Color:
		return Color(data)
	case string:
		return File(data)
	case *glot.Plot:
		return Glot(data)
	case *plot.Plot:
		return Gonum(data)
	case image.Image:
		return Image(data)
	default:
		return errors.New("unsupported type")
	}
}

func render(img image.Image, filename string) error {

	var buf bytes.Buffer

	if err := png.Encode(&buf, img); err != nil {
		return err
	}

	return ioutil.WriteFile(filename, buf.Bytes(), 0644)
}

func tempFile() (string, error) {
	directory, err := ioutil.TempDir("", "preview")
	if err != nil {
		return "", err
	}

	return path.Join(directory, "image.png"), nil
}
