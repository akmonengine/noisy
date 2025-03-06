package noisy

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func RenderImg(source SourceInterface, filename string, width int, height int) error {
	file, err := os.Create(filename)
	if err != nil {
		panic("Something went wrong with opening a file.")
	}
	defer file.Close()

	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{width, height},
	})

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			value := source.GetValue(float64(x)/float64(width), float64(y)/float64(height), 0.0)
			value = (value + 1) / 2
			if value > 1.0 {
				value = 1.0
			} else if value < 0.0 {
				value = 0.0
			}
			valueInt := uint8(255 * value)
			pixelColor := color.RGBA{valueInt, valueInt, valueInt, 255}

			img.Set(x, height-y-1, pixelColor)
		}
	}

	png.Encode(file, img.SubImage(img.Bounds()))

	return nil
}
