package noisy

import "math"

type Billow struct {
	Frequency, Lacunarity, Persistence float64
	OctaveCount, Seed                  int
}

func (billow Billow) GetValue(x, y, z float64) float64 {
	value := 0.0
	persistence := 1.0
	frequency := billow.Frequency
	seed := billow.Seed

	for range billow.OctaveCount {
		noise := persistence * getNoise(seed, x*frequency, y*frequency, z*frequency)
		value += math.Abs(noise)
		// prepare the persistency & frequency for the next octave
		persistence *= billow.Persistence
		frequency *= billow.Lacunarity
		// offset the seed to avoid seeing patterns
		seed++
	}

	return value
}
