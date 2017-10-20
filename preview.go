// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package preview

import (
	"bytes"
	"errors"
	"image"
	"image/png"
	"io/ioutil"
	"os/exec"
	"path"
)

func File(filename string) error {
	return show(filename)
}

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

func render(img image.Image, filename string) error {

	var buf bytes.Buffer

	if err := png.Encode(&buf, img); err != nil {
		return err
	}

	return ioutil.WriteFile(filename, buf.Bytes(), 0644)
}

func show(filename string) error {

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
