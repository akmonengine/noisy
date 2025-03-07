package noisy

// Perlin implements SourceInterface.
//
// - Frequency sets the first (biggest) Octave.
// - Lacunarity sets the multiplier to the frequency for each successive Octave.
// - Persistence sets the amplitude for each successive Octave.
// - OctaveCount sets the number of Octaves to generate and blend.
// - Seed sets the random seed, useful to regenerate the same image if required.
type Perlin struct {
	Frequency, Lacunarity, Persistence float64
	OctaveCount, Seed                  int
}

// GetValue returns the value between [-1;1] for a given 3D position.
func (perlin Perlin) GetValue(x, y, z float64) float64 {
	value := 0.0
	persistence := 1.0
	frequency := perlin.Frequency
	seed := perlin.Seed

	for range perlin.OctaveCount {
		value += persistence * getNoise(seed, x*frequency, y*frequency, z*frequency)
		// prepare the persistency & frequency for the next octave
		persistence *= perlin.Persistence
		frequency *= perlin.Lacunarity
		// offset the seed to avoid seeing patterns
		seed++
	}

	return value
}
