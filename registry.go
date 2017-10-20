// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package preview

import (
	"errors"
	"os/exec"
)

type viewer struct {
	Name string
	Args []string
}

var registry []viewer

func register(name string, args ...string) {
	registry = append(registry, viewer{
		Name: name,
		Args: args,
	})
}

func viewers() []viewer {
	return registry
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
