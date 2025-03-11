package noisy

import (
	"testing"
)

func TestWhiteNoise_GetValue(t *testing.T) {
	generator := WhiteNoise{}
	value := generator.GetValue(0.5, 0.5, 0.5)

	if value < -1.0 {
		t.Errorf("generated value %f should not be lower than -1", value)
	}

	if value > 1.0 {
		t.Errorf("generated value %f should not be higher than 1", value)
	}
}
