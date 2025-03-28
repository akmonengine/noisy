package noisy

import "math"

// Sphere is a SourceInterface.
//
// Frequency is the quantity of inner spheres, and the Offsets allow a padding to center/move away the sphere(s).
type Sphere struct {
	Frequency                 float64
	OffsetX, OffsetY, OffsetZ float64
}

// GetValue returns the value between [-1;1] for a given 3D position.
func (sphere Sphere) GetValue(x, y, z float64) float64 {
	x *= sphere.Frequency
	y *= sphere.Frequency
	z *= sphere.Frequency
	x += sphere.OffsetX
	y += sphere.OffsetY
	z += sphere.OffsetZ

	centerDistance := math.Sqrt(x*x + y*y + z*z)
	sphereDistance := centerDistance - math.Floor(centerDistance)
	nearestDist := math.Min(sphereDistance, 1.0-sphereDistance)

	return 1.0 - (nearestDist * 4.0)
}
