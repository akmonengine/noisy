package noisy

import "math"

// Billow implements SourceInterface.
//
// - Frequency sets the first (biggest) Octave.
// - Lacunarity sets the multiplier to the frequency for each successive Octave.
// - Persistence sets the amplitude for each successive Octave.
// - OctaveCount sets the number of Octaves to generate and blend.
// - Seed sets the random seed, useful to regenerate the same image if required.
type Billow struct {
	Frequency, Lacunarity, Persistence float64
	OctaveCount, Seed                  int
}

// GetValue returns the value between [0;1] for a given 3D position.
//
// It uses the Perlin Noise as a base, but returns the absolute values.
func (billow Billow) GetValue(x, y, z float64) float64 {
	value := 0.0
	persistence := 1.0
	frequency := billow.Frequency
	seed := billow.Seed

	for range billow.OctaveCount {
		noise := persistence * getNoise(seed, x*frequency, y*frequency, z*frequency)
		// instead of Perlin Noise, we store the absolute value
		value += math.Abs(noise)
		// prepare the persistency & frequency for the next octave
		persistence *= billow.Persistence
		frequency *= billow.Lacunarity
		// offset the seed to avoid seeing patterns
		seed++
	}

	return value
}
