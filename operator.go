package noisy

import "math"

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
	return divide.SourceA.GetValue(x, y, z) / divide.SourceB.GetValue(x, y, z)
}

// Invert takes 1 source from which to inverse the values.
type Invert struct {
	SourceA SourceInterface
}

// GetValue inverse (e.g. 1 => -1) the value from the SourceA.
func (invert Invert) GetValue(x, y, z float64) float64 {
	return -invert.SourceA.GetValue(x, y, z)
}

// Max takes 1 source and a maximum value.
type Max struct {
	SourceA SourceInterface
	Max     float64
}

// GetValue returns the value from SourceA, or the Max value if it is higher.
func (max Max) GetValue(x, y, z float64) float64 {
	value := max.SourceA.GetValue(x, y, z)

	value = math.Max(value, max.Max)

	return value
}

// Min takes 1 source and a minimum value.
type Min struct {
	SourceA SourceInterface
	Min     float64
}

// GetValue returns the value from SourceA, or the Min value if it is lower.
func (min Min) GetValue(x, y, z float64) float64 {
	value := min.SourceA.GetValue(x, y, z)

	value = math.Min(value, min.Min)

	return value
}

// Clamp takes 1 source and min/max values.
type Clamp struct {
	SourceA  SourceInterface
	Min, Max float64
}

// GetValue returns the value from SourceA, clamped between Min/Max.
func (clamp Clamp) GetValue(x, y, z float64) float64 {
	value := clamp.SourceA.GetValue(x, y, z)

	if clamp.Min > clamp.Max {
		return math.NaN()
	}
	value = math.Min(value, clamp.Max)
	value = math.Max(value, clamp.Min)

	return value
}

// Abs takes 1 source on which to fetch the absolute values.
type Abs struct {
	SourceA SourceInterface
}

// GetValue returns the absolute value from SourceA.
func (abs Abs) GetValue(x, y, z float64) float64 {
	return math.Abs(abs.SourceA.GetValue(x, y, z))
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

// Exponent takes 1 source, and an exponent.
type Exponent struct {
	SourceA  SourceInterface
	Exponent float64
}

// GetValue returns the value from SourceA, powered by Exponent.
func (exponent Exponent) GetValue(x, y, z float64) float64 {
	value := exponent.SourceA.GetValue(z, y, z)

	return math.Pow(math.Abs((value+1.0)/2.0), exponent.Exponent)*2.0 - 1.0
}
