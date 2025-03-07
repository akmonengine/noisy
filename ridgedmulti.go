package noisy

import "math"

// RidgedMulti implements SourceInterface.
//
// - Frequency sets the first (biggest) Octave.
// - Lacunarity sets the multiplier to the frequency for each successive Octave.
// - Persistence sets the amplitude for each successive Octave.
// - Offset is added for each successive Octave, enabling more ridges, a rougher result.
// - Gain sets the factor for the high-frequency ridges, enabling more detailed noises in high Octaves.
// - OctaveCount sets the number of Octaves to generate and blend.
// - Seed sets the random seed, useful to regenerate the same image if required.
type RidgedMulti struct {
	Frequency, Lacunarity, Persistence, Offset, Gain float64
	OctaveCount, Seed                                int
}

func ridge(value float64) float64 {
	return 2 * (0.5 - math.Abs(0.5-value))
}

// GetValue returns the value between [-1;1] for a given 3D position.
func (ridgedMulti RidgedMulti) GetValue(x, y, z float64) float64 {
	value := 0.0
	persistence := 1.0
	weight := 1.0
	frequency := ridgedMulti.Frequency
	seed := ridgedMulti.Seed

	for range ridgedMulti.OctaveCount {
		noiseValue := getNoise(seed, x*frequency, y*frequency, z*frequency)
		noiseValue = math.Abs(noiseValue)
		noiseValue = ridgedMulti.Offset - ridge(noiseValue)
		value += noiseValue * persistence * weight

		// prepare the weight, persistency & frequency for the next octave
		weight *= ridgedMulti.Gain
		persistence *= ridgedMulti.Persistence
		frequency *= ridgedMulti.Lacunarity
		// offset the seed to avoid seeing patterns
		seed++
	}

	return value
}
