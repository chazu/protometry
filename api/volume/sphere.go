package volume

import "math"

// Fit check if the given volume is entirely contained in the other one
func (s *Sphere) Fit(other Volume) bool {
	return false
}

// Intersects check if a volume intersects with another one
func (s *Sphere) Intersects(other Volume) bool {
	// Calculate the distance between the centers of the spheres
	o := other.(*Sphere)
	distance := math.Sqrt(
		math.Pow(s.Center.X-o.Center.X, 2) +
			math.Pow(s.Center.Y-o.Center.Y, 2) +
			math.Pow(s.Center.Z-o.Center.Z, 2))

	// Check if the distance is less than or equal to the sum of the radii
	return distance <= (s.Radius + o.Radius)
}

// Average create a new volume averaged on 2 volumes
func (s *Sphere) Average(other Volume) Volume {
	return nil
}

// Mutate create a new volume with random mutations
func (s *Sphere) Mutate(rate float64) Volume {
	return nil
}
