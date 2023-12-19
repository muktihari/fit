package kit

// Ptr returns new pointer of v
func Ptr[T any](v T) *T { return &v }
