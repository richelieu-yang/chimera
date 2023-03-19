package sliceKit

func EmptyToNil[T any](s []T) []T {
	if IsEmpty(s) {
		return nil
	}
	return s
}

// IsEmpty
/*
@param s 可以为nil
*/
func IsEmpty[T any](s []T) bool {
	return len(s) == 0
}

// IsNotEmpty
/*
@param s 可以为nil
*/
func IsNotEmpty[T any](s []T) bool {
	return len(s) > 0
}
