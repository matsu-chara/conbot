package libs

// Filter filters slice
func Filter(vs []string, f func(string) bool) []string {
	result := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map maps slice
func Map(vs []string, f func(string) string) []string {
	result := make([]string, len(vs))
	for i, v := range vs {
		result[i] = f(v)
	}
	return result
}

// Contains check contains
func Contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

// RemoveDuplicates remove duplicates
func RemoveDuplicates(xs *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *xs {
		if !found[x] {
			found[x] = true
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}

// Grouped return grouped slice
// Grouped([a,b,c,d], 2) // [a,b], [c,d]
func Grouped(xs []string, n int) [][]string {
	grouped := [][]string{}
	for i := 0; i < len(xs); i += n {
		to := i + n
		if to > len(xs) {
			to = len(xs)
		}
		grouped = append(grouped, xs[i:to])
	}
	return grouped
}

// SameLength checks xss has same elements
func SameLength(xss [][]string) bool {
	if len(xss) == 0 {
		return true
	}
	for i := 1; i < len(xss); i++ {
		if len(xss[i]) != len(xss[0]) {
			return false
		}
	}
	return true
}

// Flatten flatten
func Flatten(xss [][]string) []string {
	result := []string{}
	for _, xs := range xss {
		result = append(result, xs...)
	}
	return result
}
