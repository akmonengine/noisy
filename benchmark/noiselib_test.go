package benchmark

import (
	"github.com/ragoune/noiselib"
	"testing"
)

func BenchmarkPerlinNoiselib(b *testing.B) {
	perlin := noiselib.DefaultPerlin()
	perlin.Seed = seed
	perlin.Frequency = 3.0
	perlin.Lacunarity = 2.0
	perlin.Persistence = 0.5
	perlin.OctaveCount = 6

	for b.Loop() {
		perlin.GetValue(0.5, 0.5, 0.5)
	}

	b.ReportAllocs()
}

func BenchmarkAddPerlinNoiselib(b *testing.B) {
	perlin1 := noiselib.DefaultPerlin()
	perlin1.Seed = seed
	perlin1.Frequency = 3.0
	perlin1.Lacunarity = 2.0
	perlin1.Persistence = 0.5
	perlin1.OctaveCount = 6

	perlin2 := noiselib.DefaultPerlin()
	perlin2.Seed = seed
	perlin2.Frequency = 3.0
	perlin2.Lacunarity = 2.0
	perlin2.Persistence = 0.5
	perlin2.OctaveCount = 6

	invert := noiselib.Invert{SourceModule: make([]noiselib.Module, noiselib.InvertModuleCount)}
	invert.SetSourceModule(0, perlin2)
	addModule := noiselib.Add{SourceModule: make([]noiselib.Module, noiselib.AddModuleCount)}
	addModule.SetSourceModule(0, perlin1)
	addModule.SetSourceModule(1, invert)

	for b.Loop() {
		addModule.GetValue(0.5, 0.5, 0.5)
	}

	b.ReportAllocs()
}
