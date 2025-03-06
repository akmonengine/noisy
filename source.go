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

	generator := Sphere{
		Frequency: 1.0,
		OffsetX:   -0.5,
		OffsetY:   -0.5,
		OffsetZ:   -0.5,
	}

	RenderImg(generator, "test.png", 1024, 1024)

	return generator.GetValue(0.5, 0.5, 0.0)
}
