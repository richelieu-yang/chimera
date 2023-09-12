package mapKit

func IsEmpty[K comparable, V interface{}](m map[K]V) bool {
	return len(m) == 0
}

func IsNotEmpty[K comparable, V interface{}](m map[K]V) bool {
	return len(m) > 0
}
