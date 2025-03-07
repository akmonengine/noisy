package noisy

import "math"

type RidgedMulti struct {
	Frequency, Lacunarity, Persistence, Offset, Gain float64
	OctaveCount, Seed                                int
}

func ridge(value float64) float64 {
	return 2 * (0.5 - math.Abs(0.5-value))
}

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
