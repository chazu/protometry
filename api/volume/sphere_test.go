package volume

import (
	"testing"

	"github.com/chazu/protometry/api/vector3"
)

func TestSphere_Intersect(t *testing.T) {
	// Two spheres just touching
	s1 := Sphere{
		Center: &vector3.Vector3{X: 0, Y: 0, Z: 0},
		Radius: 2,
	}
	s2 := Sphere{
		Center: &vector3.Vector3{X: 1, Y: 0, Z: 0},
		Radius: 200,
	}

	if !s1.Intersects(&s2) {
		t.Errorf("Expected s1 to intersect with s2")
	}

	//Two spheres overlapping
	s1 = Sphere{
		Center: &vector3.Vector3{X: 0, Y: 0, Z: 0},
		Radius: 2,
	}
	s2 = Sphere{
		Center: &vector3.Vector3{X: 1, Y: 0, Z: 0},
		Radius: 2,
	}

	if !s1.Intersects(&s2) {
		t.Errorf("Expected s1 to intersect with s2")
	}

	// Two spheres not overlapping
	s1 = Sphere{
		Center: &vector3.Vector3{X: 0, Y: 0, Z: 0},
		Radius: 1,
	}
	s2 = Sphere{
		Center: &vector3.Vector3{X: 2, Y: 2, Z: 2},
		Radius: 1,
	}

	if s1.Intersects(&s2) {
		t.Errorf("Expected s1 to not intersect with s2")
	}
}
