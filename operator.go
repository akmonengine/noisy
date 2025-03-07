package noisy

import "math"

type Add struct {
	SourceA SourceInterface
	SourceB SourceInterface
}

func (add Add) GetValue(x, y, z float64) float64 {
	return add.SourceA.GetValue(x, y, z) + add.SourceB.GetValue(x, y, z)
}

type Multiply struct {
	SourceA SourceInterface
	SourceB SourceInterface
}

func (multiply Multiply) GetValue(x, y, z float64) float64 {
	return multiply.SourceA.GetValue(x, y, z) * multiply.SourceB.GetValue(x, y, z)
}

type Divide struct {
	SourceA SourceInterface
	SourceB SourceInterface
}

func (divide Divide) GetValue(x, y, z float64) float64 {
	return divide.SourceA.GetValue(x, y, z) / divide.SourceB.GetValue(x, y, z)
}

type Invert struct {
	SourceA SourceInterface
}

func (invert Invert) GetValue(x, y, z float64) float64 {
	return -invert.SourceA.GetValue(x, y, z)
}

type Max struct {
	SourceA SourceInterface
	Max     float64
}

func (max Max) GetValue(x, y, z float64) float64 {
	value := max.SourceA.GetValue(x, y, z)

	value = math.Max(value, max.Max)

	return value
}

type Min struct {
	SourceA SourceInterface
	Min     float64
}

func (min Min) GetValue(x, y, z float64) float64 {
	value := min.SourceA.GetValue(x, y, z)

	value = math.Min(value, min.Min)

	return value
}

type Clamp struct {
	SourceA  SourceInterface
	Min, Max float64
}

func (clamp Clamp) GetValue(x, y, z float64) float64 {
	value := clamp.SourceA.GetValue(x, y, z)

	if clamp.Min > clamp.Max {
		return math.NaN()
	}
	value = math.Min(value, clamp.Max)
	value = math.Max(value, clamp.Min)

	return value
}

type Abs struct {
	SourceA SourceInterface
}

func (abs Abs) GetValue(x, y, z float64) float64 {
	return math.Abs(abs.SourceA.GetValue(x, y, z))
}

type Power struct {
	SourceA SourceInterface
	SourceB SourceInterface
}

func (power Power) GetValue(x, y, z float64) float64 {
	return math.Pow(power.SourceA.GetValue(x, y, z), power.SourceB.GetValue(x, y, z))
}

type Exponent struct {
	SourceA  SourceInterface
	Exponent float64
}

func (exponent Exponent) GetValue(x, y, z float64) float64 {
	value := exponent.SourceA.GetValue(z, y, z)

	return math.Pow(math.Abs((value+1.0)/2.0), exponent.Exponent)*2.0 - 1.0
}
