package noisy

import "math"

type Add[A, B SourceInterface] struct {
	SourceA A
	SourceB B
}

func (add Add[A, B]) GetValue(x, y, z float64) float64 {
	return add.SourceA.GetValue(x, y, z) + add.SourceB.GetValue(x, y, z)
}

type Multiply[A, B SourceInterface] struct {
	SourceA A
	SourceB B
}

func (multiply Multiply[A, B]) GetValue(x, y, z float64) float64 {
	return multiply.SourceA.GetValue(x, y, z) * multiply.SourceB.GetValue(x, y, z)
}

type Divide[A, B SourceInterface] struct {
	SourceA A
	SourceB B
}

func (divide Divide[A, B]) GetValue(x, y, z float64) float64 {
	return divide.SourceA.GetValue(x, y, z) / divide.SourceB.GetValue(x, y, z)
}

type Invert[A SourceInterface] struct {
	SourceA A
}

func (invert Invert[A]) GetValue(x, y, z float64) float64 {
	return -invert.SourceA.GetValue(x, y, z)
}

type Max[A SourceInterface] struct {
	SourceA A
	Max     float64
}

func (max Max[A]) GetValue(x, y, z float64) float64 {
	value := max.SourceA.GetValue(x, y, z)

	value = math.Max(value, max.Max)

	return value
}

type Min[A SourceInterface] struct {
	SourceA A
	Min     float64
}

func (min Min[A]) GetValue(x, y, z float64) float64 {
	value := min.SourceA.GetValue(x, y, z)

	value = math.Min(value, min.Min)

	return value
}

type Clamp[A SourceInterface] struct {
	SourceA  A
	Min, Max float64
}

func (clamp Clamp[A]) GetValue(x, y, z float64) float64 {
	value := clamp.SourceA.GetValue(x, y, z)

	if clamp.Min > clamp.Max {
		return math.NaN()
	}
	value = math.Min(value, clamp.Max)
	value = math.Max(value, clamp.Min)

	return value
}

type Power[A, B SourceInterface] struct {
	SourceA A
	SourceB B
}

func (power Power[A, B]) GetValue(x, y, z float64) float64 {
	return math.Pow(power.SourceA.GetValue(x, y, z), power.SourceB.GetValue(x, y, z))
}

type Exponent[A SourceInterface] struct {
	SourceA  A
	Exponent float64
}

func (exponent Exponent[A]) GetValue(x, y, z float64) float64 {
	value := exponent.SourceA.GetValue(z, y, z)

	return math.Pow(math.Abs((value+1.0)/2.0), exponent.Exponent)*2.0 - 1.0
}
