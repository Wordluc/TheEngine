package base

type Number interface {
	int32 | float32
}

type Vec[t Number] struct {
	X t
	Y t
}

func NewVec[t Number](x, y t) Vec[t] {
	return Vec[t]{
		X: x,
		Y: y,
	}
}

func (v *Vec[t]) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v Vec[t]) Get() (t, t) {
	return v.X, v.Y
}

func (v *Vec[t]) Add(a Vec[t]) {
	v.X += a.X
	v.Y += a.Y
}

func (v *Vec[t]) Sub(a Vec[t]) {
	v.X -= a.X
	v.Y -= a.Y
}

func (v *Vec[t]) AddScalars(x, y t) {
	v.X += x
	v.Y += y
}

func (v *Vec[t]) SubScalars(x, y t) {
	v.X -= x
	v.Y -= y
}

func CastVec[from, to Number](a Vec[from]) Vec[to] {
	return Vec[to]{
		X: to(a.X),
		Y: to(a.Y),
	}
}
