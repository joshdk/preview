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

If run on a Mac, this code will open up the [Preview utility](https://en.wikipedia.org/wiki/Preview_(macOS)), displaying the downloaded image.

## License

This library is distributed under the [MIT License](https://opensource.org/licenses/MIT), see LICENSE.txt for more information.