package noisy

// SourceInterface defines a SourceA/Operator.
//
// Its method GetValue returns a value between [-1;1], for a given 3 dimension position.
type SourceInterface interface {
	GetValue(x, y, z float64) float64
}

// Constant implements SourceInterface.
//
// It stores a static Value.
type Constant struct {
	Value float64
}

// GetValue returns the static Constant.Value.
func (constant Constant) GetValue(x, y, z float64) float64 {
	return constant.Value
}
