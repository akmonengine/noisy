package noisy

import (
	"math"
)

// Add takes 2 sources to do the sum.
type Add struct {
	SourceA SourceInterface
	SourceB SourceInterface
}

// GetValue returns the sum of the values from the 2 sources.
func (add Add) GetValue(x, y, z float64) float64 {
	return add.SourceA.GetValue(x, y, z) + add.SourceB.GetValue(x, y, z)
}

// Multiply takes 2 sources to be multiplied.
type Multiply struct {
	SourceA SourceInterface
	SourceB SourceInterface
}

// GetValue returns the multiplication of the values from the 2 sources.
func (multiply Multiply) GetValue(x, y, z float64) float64 {
	return multiply.SourceA.GetValue(x, y, z) * multiply.SourceB.GetValue(x, y, z)
}

// Divide takes 2 sources to divide the value from the SourceA by the value of the SourceB.
type Divide struct {
	SourceA SourceInterface
	SourceB SourceInterface
}

// GetValue returns the division of the value from the SourceA by the value of the SourceB.
func (divide Divide) GetValue(x, y, z float64) float64 {
	if divide.SourceB.GetValue(x, y, z) == 0.0 {
		return math.NaN()
	}

	return divide.SourceA.GetValue(x, y, z) / divide.SourceB.GetValue(x, y, z)
}

// Invert takes 1 source from which to inverse the values.
type Invert struct {
	Source SourceInterface
}

// GetValue inverse (e.g. 1 => -1) the value from the Source.
func (invert Invert) GetValue(x, y, z float64) float64 {
	return -invert.Source.GetValue(x, y, z)
}

// Max takes 1 source and a maximum value.
type Max struct {
	SourceA SourceInterface
	SourceB SourceInterface
}

// GetValue returns the value from SourceA, or the SourceB value if it is higher.
func (max Max) GetValue(x, y, z float64) float64 {
	value := math.Max(max.SourceA.GetValue(x, y, z), max.SourceB.GetValue(x, y, z))

	return value
}

// Min takes 1 source and a minimum value.
type Min struct {
	SourceA SourceInterface
	SourceB SourceInterface
}

// GetValue returns the value from SourceA, or the SourceB value if it is lower.
func (min Min) GetValue(x, y, z float64) float64 {
	value := math.Min(min.SourceA.GetValue(x, y, z), min.SourceB.GetValue(x, y, z))

	return value
}

// Clamp takes 1 source and min/max values.
type Clamp struct {
	Source               SourceInterface
	SourceMin, SourceMax SourceInterface
}

// GetValue returns the value from Source, clamped between the values of SourceMin/SourceMax.
func (clamp Clamp) GetValue(x, y, z float64) float64 {
	value := clamp.Source.GetValue(x, y, z)
	minValue := clamp.SourceMin.GetValue(x, y, z)
	maxValue := clamp.SourceMax.GetValue(x, y, z)

	if minValue > maxValue {
		return math.NaN()
	}
	value = math.Min(value, maxValue)
	value = math.Max(value, minValue)

	return value
}

// Abs takes 1 source on which to fetch the absolute values.
type Abs struct {
	Source SourceInterface
}

// GetValue returns the absolute value from Source.
func (abs Abs) GetValue(x, y, z float64) float64 {
	return math.Abs(abs.Source.GetValue(x, y, z))
}

// Power takes 2 sources.
type Power struct {
	SourceA SourceInterface
	SourceB SourceInterface
}

// GetValue returns the value from SourceA, powered by the value from SourceB.
func (power Power) GetValue(x, y, z float64) float64 {
	return math.Pow(power.SourceA.GetValue(x, y, z), power.SourceB.GetValue(x, y, z))
}
