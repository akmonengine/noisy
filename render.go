package noisy

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"sort"
)

type Gradient map[float64]color.RGBA

func clampValue(value, lowerBound, upperBound int) int {
	if value < lowerBound {
		return lowerBound
	} else if value > upperBound {
		return upperBound
	} else {
		return value
	}
}

func (gradient Gradient) GetColor(position float64) (color.RGBA, error) {
	if len(gradient) < 2 {
		return color.RGBA{}, fmt.Errorf("a Gradient must have at least 2 values")
	}

	keys := []float64{}
	for k := range gradient {
		keys = append(keys, k)
	}

	sort.Float64s(keys)

	indexPos := 0

	for _, k := range keys {
		if position < k {
			break
		}
		indexPos++
	}

	index0 := clampValue(indexPos-1, 0, len(gradient)-1)
	index1 := clampValue(indexPos, 0, len(gradient)-1)

	if index0 == index1 {
		return gradient[keys[index1]], nil
	}

	input0 := keys[index0]
	input1 := keys[index1]
	alpha := (position - input0) / (input1 - input0)

	color0 := gradient[keys[index0]]
	color1 := gradient[keys[index1]]

	return linearInterpColor(color0, color1, alpha), nil
}

func blendChannel(channel0, channel1 uint8, alpha float64) uint8 {
	c0 := float64(channel0) / 255.0
	c1 := float64(channel1) / 255.0
	return uint8(((c1 * alpha) + (c0 * (1.0 - alpha))) * 255.0)
}

func linearInterpColor(color0, color1 color.RGBA, alpha float64) color.RGBA {
	a := blendChannel(color0.A, color1.A, alpha)
	r := blendChannel(color0.R, color1.R, alpha)
	g := blendChannel(color0.G, color1.G, alpha)
	b := blendChannel(color0.B, color1.B, alpha)

	return color.RGBA{r, g, b, a}
}

func RenderImg(source SourceInterface, gradient Gradient, filename string, width int, height int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{width, height},
	})

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			value := source.GetValue(float64(x)/float64(width), float64(y)/float64(height), 0.0)

			pixelColor, err := gradient.GetColor(value)
			if err != nil {
				return err
			}

			img.Set(x, height-y-1, pixelColor)
		}
	}

	return png.Encode(file, img.SubImage(img.Bounds()))
}
