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

func (v Vec[t]) Get() (t, t) {
	return v.X, v.Y
}

func (v *Vec[t]) Clone() *Vec[t] {
	return &Vec[t]{
		v.X,
		v.Y,
	}

}

func (v *Vec[t]) CapAt(a Vec[t]) *Vec[t] {
	if v.X > a.X {
		v.X = a.X
	}
	if v.Y > a.Y {
		v.Y = a.Y
	}
	return v
}

func (v *Vec[t]) MultScalar(a t) *Vec[t] {
	v.X *= a
	v.Y *= a
	return v
}

func (v *Vec[t]) Add(a Vec[t]) *Vec[t] {
	v.X += a.X
	v.Y += a.Y
	return v
}

func (v *Vec[t]) Sub(a Vec[t]) *Vec[t] {
	v.X -= a.X
	v.Y -= a.Y
	return v
}

func (v *Vec[t]) AddScalars(x, y t) *Vec[t] {
	v.X += x
	v.Y += y
	return v
}

func (v *Vec[t]) SubScalars(x, y t) *Vec[t] {
	v.X -= x
	v.Y -= y
	return v
}

func FromAtoBVec[t Number](a, b Vec[t]) Vec[t] {
	return Vec[t]{X: b.X - a.X, Y: b.Y - a.Y}
}

func (v *Vec[t]) IsNull() bool {
	return v.X == 0 && v.Y == 0
}

func SubVecs[t Number](a, b Vec[t]) Vec[t] {
	return Vec[t]{
		a.X - b.X,
		a.Y - b.Y,
	}
}
func AddVecs[t Number](a, b Vec[t]) Vec[t] {
	return Vec[t]{
		a.X + b.X,
		a.Y + b.Y,
	}
}

func CastVec[from, to Number](a Vec[from]) Vec[to] {
	return Vec[to]{
		X: to(a.X),
		Y: to(a.Y),
	}
}
