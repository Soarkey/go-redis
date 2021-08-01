package utils

// BytesEquals 判断给定的两个bytes是否相等
func BytesEquals(a []byte, b []byte) bool {
	if (a == nil && b != nil) || (a != nil && b == nil) || len(a) != len(b) {
		return false
	}
	size := len(a)
	for i := 0; i < size; i++ {
		av := a[i]
		bv := b[i]
		if av != bv {
			return false
		}
	}
	return true
}
