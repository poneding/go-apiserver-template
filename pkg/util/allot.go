package util

// Ptr returns a pointer to v.
func Ptr[T any](v T) *T {
	return &v
}

// ValueIf returns v if cond is true, otherwise def.
func ValueIf[T any](cond bool, v, def T) T {
	if cond {
		return v
	}
	return def
}
