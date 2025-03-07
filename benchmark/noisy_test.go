package benchmark

import (
	"github.com/akmonengine/noisy"
	"testing"
)

func BenchmarkPerlinNoisy(b *testing.B) {
	generator := noisy.Perlin{
		Frequency:   3.0,
		Lacunarity:  2.0,
		Persistence: 0.5,
		OctaveCount: 6,
		Seed:        seed,
	}

	for b.Loop() {
		generator.GetValue(0.5, 0.5, 0.5)
	}

	b.ReportAllocs()
}

func BenchmarkAddPerlinNoisy(b *testing.B) {
	generator := noisy.Add{
		noisy.Perlin{
			Frequency:   3.0,
			Lacunarity:  2.0,
			Persistence: 0.5,
			OctaveCount: 6,
			Seed:        seed,
		},
		noisy.Invert{noisy.Perlin{
			Frequency:   3.0,
			Lacunarity:  2.0,
			Persistence: 0.5,
			OctaveCount: 6,
			Seed:        seed * 100,
		}},
	}

	for b.Loop() {
		generator.GetValue(0.5, 0.5, 0.5)
	}

	b.ReportAllocs()
}
