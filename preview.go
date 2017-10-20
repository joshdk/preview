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
	"os/exec"
	"path"
)

// File will open a viewer for the path specified by filename.
func File(filename string) error {
	return view(filename)
}

// Image will open a viewer for the image.Image specified by img.
func Image(img image.Image) error {

	directory, err := ioutil.TempDir("", "preview")
	if err != nil {
		return err
	}

	filename := path.Join(directory, "image.png")

	if err := render(img, filename); err != nil {
		return err
	}

	return File(filename)
}

// Color will open a viewer for the color.Color specified by img.
// The previewed image will be 256x256 pixels in size, completely filled with the given color.
func Color(clr color.Color) error {
	img := image.NewRGBA(image.Rect(0, 0, 256, 256))
	draw.Draw(img, img.Bounds(), &image.Uniform{clr}, image.ZP, draw.Src)

	return Image(img)
}

// Show will open a viewer for the type specified by any.
func Show(any interface{}) error {
	switch data := any.(type) {
	case color.Color:
		return Color(data)
	case image.Image:
		return Image(data)
	case string:
		return File(data)
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

func view(filename string) error {

	for _, viewer := range viewers() {
		cmd := exec.Command(viewer.Name, append(viewer.Args, filename)...)

		if err := cmd.Start(); err != nil {

			if execErr, ok := err.(*exec.Error); ok {
				// This view does not exist, try the next one
				if execErr.Err == exec.ErrNotFound {
					continue
				}
			}

			return err
		}

		return nil
	}

	return errors.New("no viewers available")
}
