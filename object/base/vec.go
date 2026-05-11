package base

type Vec struct {
	x int32
	y int32
}

func NewVec(x, y int32) Vec {
	return Vec{
		x: x,
		y: y,
	}
}

func (v *Vec) Get() (int32, int32) {
	return v.x, v.y
}

func (v *Vec) Add(a Vec) {
	v.x += a.x
	v.y += a.y
}

func (v *Vec) Sub(a Vec) {
	v.x -= a.x
	v.y -= a.y
}

func (v *Vec) AddScalars(x, y int32) {
	v.x += x
	v.y += y
}

func (v *Vec) SubScalars(x, y int32) {
	v.x -= x
	v.y -= y
}
