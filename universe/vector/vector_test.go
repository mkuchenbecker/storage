package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Parallel()
	vec := New(1, 2, 3)

	x, y, z := vec.Get()
	assert.Equal(t, float64(1), x)
	assert.Equal(t, float64(2), y)
	assert.Equal(t, float64(3), z)
}

func TestAdd(t *testing.T) {
	t.Parallel()
	a := New(1, 2, 3)
	b := New(1, 2, 3)

	sum := a.Add(b)
	assert.Equal(t, New(2, 4, 6), sum)
}

func TestEquals(t *testing.T) {
	t.Parallel()
	a := New(1, 2, 3)

	assert.True(t, a.Equals(New(1, 2, 3)))
}

func TestSquareMagnitude(t *testing.T) {
	t.Parallel()
	a := New(1, 2, 3)

	assert.Equal(t, float64(14), a.SquaredMagnitude())
}

func TestNegative(t *testing.T) {
	t.Parallel()
	a := New(1, 2, 3)

	assert.Equal(t, New(-1, -2, -3), a.Negative())
}

func TestNormalize(t *testing.T) {
	t.Parallel()
	x := New(1000, 0, 0)
	y := New(0, 1000, 0)
	z := New(0, 0, 1000)

	assert.Equal(t, New(1, 0, 0), x.Normalize())
	assert.Equal(t, New(0, 1, 0), y.Normalize())
	assert.Equal(t, New(0, 0, 1), z.Normalize())
}

func TestProto(t *testing.T) {
	t.Parallel()
	vec := New(1, 2, 3)
	protoVec := vec.ToProto()
	back := FromProto(protoVec)

	assert.Equal(t, vec, back)
}
