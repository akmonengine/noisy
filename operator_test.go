package noisy

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-9

func checkFloatEquality(a, b float64) bool {
	if math.IsNaN(a) && math.IsNaN(b) {
		return true
	}

	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestAdd_GetValue(t *testing.T) {
	tests := []struct {
		name      string
		operation Add
		want      float64
	}{
		{"addition1", Add{
			SourceA: Constant{0.2},
			SourceB: Constant{0.3},
		}, 0.5},
		{"addition2", Add{
			SourceA: Constant{0.2},
			SourceB: Constant{-0.3},
		}, -0.1},
		{"addition3", Add{
			SourceA: Constant{1.5},
			SourceB: Constant{0.0},
		}, 1.5},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value := test.operation.GetValue(0.5, 0.5, 0.5)
			if !checkFloatEquality(value, test.want) {
				t.Errorf("the operation %s returned %f, should be equal to %f", test.name, value, test.want)
			}
		})
	}
}

func TestMultiply_GetValue(t *testing.T) {
	tests := []struct {
		name      string
		operation Multiply
		want      float64
	}{
		{"multiplication1", Multiply{
			SourceA: Constant{0.2},
			SourceB: Constant{0.3},
		}, 0.06},
		{"multiplication2", Multiply{
			SourceA: Constant{0.2},
			SourceB: Constant{-0.3},
		}, -0.06},
		{"multiplication3", Multiply{
			SourceA: Constant{1.5},
			SourceB: Constant{0.0},
		}, 0.0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value := test.operation.GetValue(0.5, 0.5, 0.5)
			if !checkFloatEquality(value, test.want) {
				t.Errorf("the operation %s returned %f, should be equal to %f", test.name, value, test.want)
			}
		})
	}
}

func TestDivide_GetValue(t *testing.T) {
	tests := []struct {
		name      string
		operation Divide
		want      float64
	}{
		{"divide1", Divide{
			SourceA: Constant{0.2},
			SourceB: Constant{0.3},
		}, 0.666666667},
		{"divide2", Divide{
			SourceA: Constant{0.2},
			SourceB: Constant{-0.3},
		}, -0.666666667},
		{"divide3", Divide{
			SourceA: Constant{1.5},
			SourceB: Constant{0.0},
		}, math.NaN()},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value := test.operation.GetValue(0.5, 0.5, 0.5)
			if !checkFloatEquality(value, test.want) {
				t.Errorf("the operation %s returned %f, should be equal to %f", test.name, value, test.want)
			}
		})
	}
}

func TestInvert_GetValue(t *testing.T) {
	tests := []struct {
		name      string
		operation Invert
		want      float64
	}{
		{"invert1", Invert{
			Source: Constant{0.2},
		}, -0.2},
		{"invert2", Invert{
			Source: Constant{-0.2},
		}, 0.2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value := test.operation.GetValue(0.5, 0.5, 0.5)
			if !checkFloatEquality(value, test.want) {
				t.Errorf("the operation %s returned %f, should be equal to %f", test.name, value, test.want)
			}
		})
	}
}

func TestMax_GetValue(t *testing.T) {
	tests := []struct {
		name      string
		operation Max
		want      float64
	}{
		{"max1", Max{
			SourceA: Constant{0.2},
			SourceB: Constant{0.3},
		}, 0.3},
		{"max2", Max{
			SourceA: Constant{0.3},
			SourceB: Constant{0.2},
		}, 0.3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value := test.operation.GetValue(0.5, 0.5, 0.5)
			if !checkFloatEquality(value, test.want) {
				t.Errorf("the operation %s returned %f, should be equal to %f", test.name, value, test.want)
			}
		})
	}
}

func TestMin_GetValue(t *testing.T) {
	tests := []struct {
		name      string
		operation Min
		want      float64
	}{
		{"min1", Min{
			SourceA: Constant{0.2},
			SourceB: Constant{0.1},
		}, 0.1},
		{"min2", Min{
			SourceA: Constant{0.1},
			SourceB: Constant{0.2},
		}, 0.1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value := test.operation.GetValue(0.5, 0.5, 0.5)
			if !checkFloatEquality(value, test.want) {
				t.Errorf("the operation %s returned %f, should be equal to %f", test.name, value, test.want)
			}
		})
	}
}

func TestClamp_GetValue(t *testing.T) {
	tests := []struct {
		name      string
		operation Clamp
		want      float64
	}{
		{"clamp1", Clamp{
			Source:    Constant{0.2},
			SourceMin: Constant{0.1},
			SourceMax: Constant{0.3},
		}, 0.2},
		{"clamp2", Clamp{
			Source:    Constant{-0.1},
			SourceMin: Constant{0.1},
			SourceMax: Constant{0.3},
		}, 0.1},
		{"clamp3", Clamp{
			Source:    Constant{0.4},
			SourceMin: Constant{0.1},
			SourceMax: Constant{0.3},
		}, 0.3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value := test.operation.GetValue(0.5, 0.5, 0.5)
			if !checkFloatEquality(value, test.want) {
				t.Errorf("the operation %s returned %f, should be equal to %f", test.name, value, test.want)
			}
		})
	}
}

func TestAbs_GetValue(t *testing.T) {
	tests := []struct {
		name      string
		operation Abs
		want      float64
	}{
		{"abs1", Abs{
			Source: Constant{-0.2},
		}, 0.2},
		{"abs2", Abs{
			Source: Constant{0.2},
		}, 0.2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value := test.operation.GetValue(0.5, 0.5, 0.5)
			if !checkFloatEquality(value, test.want) {
				t.Errorf("the operation %s returned %f, should be equal to %f", test.name, value, test.want)
			}
		})
	}
}

func TestPower_GetValue(t *testing.T) {
	tests := []struct {
		name      string
		operation Power
		want      float64
	}{
		{"power1", Power{
			SourceA: Constant{0.5},
			SourceB: Constant{0.5},
		}, 0.707106781},
		{"power2", Power{
			SourceA: Constant{-2},
			SourceB: Constant{1.8},
		}, math.NaN()},
		{"power3", Power{
			SourceA: Constant{1.5},
			SourceB: Constant{0.0},
		}, 1.0},
		{"power4", Power{
			SourceA: Constant{-2},
			SourceB: Constant{0.0},
		}, 1.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value := test.operation.GetValue(0.5, 0.5, 0.5)
			if !checkFloatEquality(value, test.want) {
				t.Errorf("the operation %s returned %f, should be equal to %f", test.name, value, test.want)
			}
		})
	}
}
