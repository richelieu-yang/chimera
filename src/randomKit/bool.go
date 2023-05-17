package randomKit

func Bool() bool {
	// i: [0, 2)
	i := r.Intn(2)
	return i == 1
}
