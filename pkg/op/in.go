package op

// In 包含.
func In[T comparable](elem T, elems []T) bool {
	for _, e := range elems {
		if elem == e {
			return true
		}
	}
	return false
}
