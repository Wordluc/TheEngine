package base

import "math"

type Number interface {
	int32 | float32
}

type UVec[t Number] struct {
	X t
	Y t
}

func NewVec[t Number](x, y t) UVec[t] {
	return UVec[t]{
		X: x,
		Y: y,
	}
}

func (v UVec[t]) Get() (t, t) {
	return v.X, v.Y
}

func (v *UVec[t]) Clone() *UVec[t] {
	return &UVec[t]{
		v.X,
		v.Y,
	}

}

func (v *UVec[t]) CapAt(a UVec[t]) *UVec[t] {
	if v.X > 0 && v.X > a.X {
		v.X = a.X
	} else if v.X < 0 && v.X < -a.X {
		v.X = -a.X
	}
	if v.Y > 0 && v.Y > a.Y {
		v.Y = a.Y
	} else if v.Y < 0 && v.Y < -a.Y {
		v.Y = -a.Y
	}
	return v
}

func (v *UVec[t]) Magnitude() t {
	return t(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v *UVec[t]) Normalize() *UVec[t] {
	mag := v.Magnitude()
	if mag == 0 {
		return &UVec[t]{}
	}
	return &UVec[t]{
		X: v.X / mag,
		Y: v.Y / mag,
	}
}

func (v *UVec[t]) MultScalar(a t) *UVec[t] {
	v.X *= a
	v.Y *= a
	return v
}

func (v *UVec[t]) Add(a UVec[t]) *UVec[t] {
	v.X += a.X
	v.Y += a.Y
	return v
}

func (v *UVec[t]) Sub(a UVec[t]) *UVec[t] {
	v.X -= a.X
	v.Y -= a.Y
	return v
}

func (v *UVec[t]) AddScalars(x, y t) *UVec[t] {
	v.X += x
	v.Y += y
	return v
}

func (v *UVec[t]) SubScalars(x, y t) *UVec[t] {
	v.X -= x
	v.Y -= y
	return v
}

func FromAtoBVec[t Number](a, b UVec[t]) UVec[t] {
	return UVec[t]{X: b.X - a.X, Y: b.Y - a.Y}
}

func (v *UVec[t]) IsNull() bool {
	return v.X == 0 && v.Y == 0
}

func SubVecs[t Number](a, b UVec[t]) UVec[t] {
	return UVec[t]{
		a.X - b.X,
		a.Y - b.Y,
	}
}
func AddVecs[t Number](a, b UVec[t]) UVec[t] {
	return UVec[t]{
		a.X + b.X,
		a.Y + b.Y,
	}
}

func CastVec[from, to Number](a UVec[from]) UVec[to] {
	return UVec[to]{
		X: to(a.X),
		Y: to(a.Y),
	}
}
