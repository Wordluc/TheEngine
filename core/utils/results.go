package utils

type Result[t any] struct {
	value t
	err   error
}

func NewResult[t any](r t, err error) Result[t] {
	return Result[t]{
		value: r,
		err:   err,
	}
}

func (r Result[t]) Open() (t, error) {
	return r.value, r.err
}

func ResultErr[t any](err error) Result[t] {
	return Result[t]{
		err: err,
	}
}

func ResultOk[t any](r t) Result[t] {
	return Result[t]{
		value: r,
	}
}

func CheckError[t any](r Result[t]) t {
	if r.err != nil {
		panic(r.err)
	}
	return r.value
}

func ResultsArray[t any](results ...Result[t]) (res []t, err error) {
	for _, r := range results {
		if r.err != nil {
			return nil, err
		}
		res = append(res, r.value)
	}
	return res, nil
}

func ResultsMap[key comparable, t any](results map[key]Result[t]) (res map[key]t, err error) {
	res = map[key]t{}
	for key, r := range results {
		if r.err != nil {
			return nil, err
		}
		res[key] = r.value
	}
	return res, nil
}
