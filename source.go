package noisy

type SourceInterface interface {
	GetValue(x, y, z float64) float64
}

type Constant struct {
	Value float64
}

func (constant Constant) GetValue(x, y, z float64) float64 {
	return constant.Value
}

func Generate() float64 {
	//	generator := Divide[Perlin, Invert[Constant]]{
	//		Perlin{Frequency: 0.2},
	//		Invert[Constant]{Constant{0.3}},
	//	}

	generator := Perlin{
		OctaveCount: 8,
		Persistence: 0.7,
		Lacunarity:  2.0,
		Frequency:   1.0,
		Seed:        18272,
	}

	RenderImg(generator, "test.png", 1024, 1024)

	return generator.GetValue(0.5, 0.5, 0.0)
}
