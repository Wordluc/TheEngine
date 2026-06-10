package base

type Vec[t Number] struct {
	UVec[t]
	length t
}

func NewLine[t Number](vec UVec[t], length t) Vec[t] {
	return Vec[t]{
		vec,
		length,
	}
}

func (l *Vec[t]) SetRotation(r float32) {
}

func (l *Vec[t]) GetRotation(r float32) {
}
