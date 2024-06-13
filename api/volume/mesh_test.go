package volume

import (
	"github.com/chazu/protometry/internal/utils"
	"math/rand"
	"testing"
)

func TestMesh_Mutate(t *testing.T) {
	rand.Seed(1337)
	m := NewMeshSquareCuboid(1, true)
	mutated := m.Mutate(10)
	utils.Equals(t, false, mutated == m)
}

func TestMesh_Clone(t *testing.T) {
	m := NewMeshSquareCuboid(1, true)
	nm := m.Clone()
	utils.Equals(t, false, m == nm)
}
