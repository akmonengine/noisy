package noisy

import (
	"fmt"
	"image/color"
	"testing"
)

func TestGenerate(t *testing.T) {
	generator := RidgedMulti{
		Frequency:   5.0,
		Lacunarity:  2.0,
		Persistence: 0.5,
		Offset:      0.3,
		Gain:        0.9,
		OctaveCount: 6,
		Seed:        48498,
	}

	RenderImg(generator, map[float64]color.RGBA{
		-1.0: {0, 0, 0, 255},
		0.0:  {100, 120, 0, 255},
		0.5:  {120, 200, 0, 255},
	}, "test.png", 1024, 1024)

	fmt.Println(generator.GetValue(0.5, 0.5, 0.0))
}
