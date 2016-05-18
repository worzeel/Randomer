package randomer

import "strconv"

type Randomer struct {
	calls int
}

func (r *Randomer) Random() string {
	var result string

	r.calls++
	result = strconv.Itoa(r.calls)

	return result
}
