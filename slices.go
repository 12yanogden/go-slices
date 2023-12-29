package slices

// Remove an element from the slice by its index
func Remove[T any](slice []T, i int) []T {
	return append(slice[:i], slice[(i+1):]...)
}
