# Noisy
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/akmonengine/noisy)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Reference](https://img.shields.io/badge/reference-%23007D9C?logo=go&logoColor=white&labelColor=gray)](https://pkg.go.dev/github.com/akmonengine/noisy)
[![Go Report Card](https://goreportcard.com/badge/github.com/akmonengine/noisy)](https://goreportcard.com/report/github.com/akmonengine/noisy)
![Tests](https://img.shields.io/github/actions/workflow/status/akmonengine/noisy/code_coverage.yml?label=tests)
![Codecov](https://img.shields.io/codecov/c/github/akmonengine/noisy)
![GitHub Issues or Pull Requests](https://img.shields.io/github/issues/akmonengine/noisy)
![GitHub Issues or Pull Requests](https://img.shields.io/github/issues-pr/akmonengine/noisy)

Noisy is a Go tool to generate textures and heightmaps with noises algorithm (up to 3 dimensions). This is useful
for procedural generation, for example in game development.
Using Perlin Noise (https://mrl.cs.nyu.edu/~perlin/noise/), it allows to combine different Sources (perlin noise, ridged multifractal) using Operators (e.g. add, multiply, clamps).

## Basic Usage
Note that the returned values are always contained between [-1;1].

### Sources
Multiple Sources are available, and you can create your own if you implement the SourceInterface.

#### Constant
Constant has a defined Value, and returns this value whatever is the position x,y,z.
```go
source := noisy.Constant{1.0}
```

#### Sphere
Sphere has a Frequency setting the number of inner spheres, and an Offset to translate it.
```go
source := noisy.Sphere{
	Frequency: 1.0,
	OffsetX: -0.5,
	OffsetY: -0.5,
	OffsetZ: -0.5,
}
```

#### Perlin Noise
Perlin Noise is a standard for noise generation.
```go
source := noisy.Perlin{
    Frequency:   5.0,
    Lacunarity:  2.0,
    Persistence: 0.5,
    OctaveCount: 6,
    Seed:        42,
}
```
The parameters allow to custom the generated result:
- Frequency sets the first (biggest) Octave.
- Lacunarity sets the multiplier to the frequency for each successive Octave.
- Persistence sets the amplitude for each successive Octave.
- OctaveCount sets the number of Octaves to generate and blend.
- Seed sets the random seed, useful to regenerate the same image if required.

#### Billow
Billow Noise is similar to Perlin Noise, except it returns only absolute values.
So the resulted content would be between [0;1].
```go
source := noisy.Billow{
    Frequency:   5.0,
    Lacunarity:  2.0,
    Persistence: 0.5,
    OctaveCount: 6,
    Seed:        42,
}
```

#### Ridged Multifractal
Ridged Multifractal is based on the Perlin Noise, but allows visual effects similar to mountains.
```go
source := noisy.RidgedMulti{
    Frequency:   5.0,
    Lacunarity:  2.0,
    Persistence: 0.5, 
    Offset:      0.4, 
    Gain:        0.6,
    OctaveCount: 6,
    Seed:        42,
}
```
The new parameters define the ridges:
- Offset is added for each successive Octave, enabling more ridges, a rougher result.
- Gain sets the factor for the high-frequency ridges, enabling more detailed noises in high Octaves.

### Operators
#### Add
The Add operator sums two Sources.

#### Multiply
The Multiply operator multiply the values from two Sources.
#### Divide
The Divide operator divides the value from a Source A by the value from a Source B.

#### Invert
The Invert operator inverts the value. i.e. 1 => -1 and -1 => 1.

#### Max
The Max operator returns the highest value between a Source and the its parameter Max.

#### Min
The Min operator returns the lowest value between a Source and the its parameter Min.

#### Clamp
The Clamp operator returns the value from a Source, clamped between [Min;Max].

#### Abs
The Abs operator returns the absolute value from a Source, limiting it to [0;1].

#### Power
The Power operator returns the value from a Source A, powered by the value from a Source B.

#### Exponent
The Exponent operator returns the value from a Source, powered by Exponent.

### Result
Once your generator is built, you can either fetch the value for one position:
```go
value := generator.GetValue(0.0, 0.0, 0.0)
```
This value is contained between [-1;1].

You can also generate an image, stored on the filesystem:
```go
err := RenderImg(generator, map[float64]color.RGBA{
-1.0: {0, 0, 0, 255},
1.0: {255, 255, 255, 255},
}, "noise.png", 1024, 1024)

if err != nil {
    fmt.Println(err.Error())
}
```

### Complex Example
Let's create a volcanic island for a game. For the base of our island, we can use a circle (a Sphere source).
```go
seed := 876310720398733142
generator := Multiply{
    Sphere{
        Frequency: 2.0,
        OffsetX:   -1.0,
        OffsetY:   -1.0,
        OffsetZ:   -1.0,
    },
    Clamp{
        SourceA: Add{
            Perlin{
                Frequency:   1.1,
                Lacunarity:  2.0,
                Persistence: 0.5,
                OctaveCount: 6,
                Seed:        seed,
            },
            RidgedMulti{
                Frequency:   3.0,
                Lacunarity:  2.0,
                Persistence: 0.5,
                Offset:      0.5,
                Gain:        0.8,
                OctaveCount: 5,
                Seed:        seed,
            },
        },
        Min: 0.0,
        Max: 1.0,
    },
}
```

## What's next ?
### Sources
- Checkboard source
- Cylinder source
- Sine source
- Voronoi source
- White noise source

### Operators
- Displace
- Terrace

### Processes
- Gaussian Blur
- Erosion

## Sources
- https://mrl.cs.nyu.edu/~perlin/noise/

## Contributing Guidelines

See [how to contribute](CONTRIBUTING.md).

## Licence
This project is distributed under the [Apache 2.0 licence](LICENCE.md).