[![License](https://img.shields.io/github/license/joshdk/preview.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/joshdk/preview?status.svg)](https://godoc.org/github.com/joshdk/preview)
[![Go Report Card](https://goreportcard.com/badge/github.com/joshdk/preview)](https://goreportcard.com/report/github.com/joshdk/preview)
[![CircleCI](https://circleci.com/gh/joshdk/preview.svg?&style=shield)](https://circleci.com/gh/joshdk/preview/tree/master)

# Preview

🎨 Simple cross platform image viewing for developers

## Installing

You can fetch this library by running the following

    go get -u github.com/joshdk/preview

## Usage

```go
import (
	"image/jpeg"
	"net/http"
	"github.com/joshdk/preview"
)

resp, err := http.Get("https://i.imgur.com/X9GB4Pu.jpg")
if err != nil {
	panic(err.Error())
}

img, err := jpeg.Decode(resp.Body)
if err != nil {
	panic(err.Error())
}

preview.Image(img)
```

If run on a Mac, this code will open up the [Preview utility](https://en.wikipedia.org/wiki/Preview_(macOS)), displaying the downloaded image. On Linux, the [display](https://linux.die.net/man/1/display) or [xv](https://en.wikipedia.org/wiki/Xv_(software)) utility is used.

## License

This library is distributed under the [MIT License](https://opensource.org/licenses/MIT), see LICENSE.txt for more information.