// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package preview

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
