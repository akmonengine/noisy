package noisy

import (
	"math/rand"
)

// WhiteNoise implements SourceInterface.
//
// It returns a completely random value.
type WhiteNoise struct {
}

// GetValue returns a random value between [-1;1] for a given 3D position.
func (whiteNoise WhiteNoise) GetValue(x, y, z float64) float64 {
	return rand.Float64()
}
