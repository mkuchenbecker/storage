package vector

import "math"

type Vector interface {
	Get() (float64, float64, float64)
	Add(Vector) Vector
	SquaredMagnitude() float64
	Negative() Vector
	Equals(other Vector) bool
	ToProto() *Vector3
	Normalize() Vector
}

type vector struct {
	x float64
	y float64
	z float64
}

func New(x, y, z float64) Vector {
	return vector{
		x: x,
		y: y,
		z: z,
	}
}

func (v vector) Get() (float64, float64, float64) {
	return v.x, v.y, v.z
}

func (v vector) Add(other Vector) Vector {
	x, y, z := v.Get()
	ox, oy, oz := other.Get()

	return New(x+ox, y+oy, z+oz)
}

func (v vector) Equals(other Vector) bool {
	x, y, z := v.Get()

	ox, oy, oz := other.Get()
	return x == ox && y == oy && oz == z
}

func (v vector) SquaredMagnitude() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v vector) Negative() Vector {
	return New(-v.x, -v.y, -v.z)
}

func (v vector) ToProto() *Vector3 {
	return &Vector3{X: v.x, Y: v.y, Z: v.z}
}

func FromProto(v *Vector3) Vector {
	return New(v.X, v.Y, v.Z)
}

func (v vector) Normalize() Vector {
	squareMagnitude := v.SquaredMagnitude()
	inverseMagnitude := 1 / math.Sqrt(squareMagnitude)
	return New(v.x*inverseMagnitude, v.y*inverseMagnitude, v.z*inverseMagnitude)
}
