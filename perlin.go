package noisy

type Perlin struct {
	Frequency, Lacunarity, Persistence float64
	OctaveCount, Seed                  int
}

func (perlin Perlin) GetValue(x, y, z float64) float64 {
	value := 0.0
	persistence := 1.0
	frequency := perlin.Frequency
	seed := perlin.Seed

	for range perlin.OctaveCount {
		value += persistence * getNoise(seed, x*frequency, y*frequency, z*frequency)
		// prepapre the persistency & frequency for the next octave
		persistence *= perlin.Persistence
		frequency *= perlin.Lacunarity
		// offset the seed to avoid seeing patterns
		seed++
	}

	return value
}
