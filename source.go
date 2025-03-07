package noisy

type SourceInterface interface {
	GetValue(x, y, z float64) float64
}

type Constant struct {
	Value float64
}

func (constant Constant) GetValue(x, y, z float64) float64 {
	return constant.Value
}
