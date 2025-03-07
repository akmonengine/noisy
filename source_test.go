package noisy

import (
	"fmt"
	"image/color"
	"math/rand/v2"
	"testing"
)

func TestGenerate(t *testing.T) {
	seed := rand.Int()
	generator := Multiply{
		Sphere{
			Frequency: 2.0,
			OffsetX:   -1.0,
			OffsetY:   -1.0,
			OffsetZ:   -1.0,
		},
		Clamp{
			SourceA: Add{
				Perlin{
					Frequency:   1.1,
					Lacunarity:  2.0,
					Persistence: 0.5,
					OctaveCount: 6,
					Seed:        seed,
				},
				RidgedMulti{
					Frequency:   3.0,
					Lacunarity:  2.0,
					Persistence: 0.5,
					Offset:      0.5,
					Gain:        0.8,
					OctaveCount: 5,
					Seed:        seed,
				},
			},
			Min: 0.0,
			Max: 1.0,
		},
	}

	err := RenderImg(generator, map[float64]color.RGBA{
		-1.0: {0, 0, 0, 255},
		0.0:  {0, 0, 0, 255},
		0.5:  {255, 255, 255, 255},
	}, "test.png", 1024, 1024)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(seed, generator.GetValue(0.5, 0.5, 0.0))
}
