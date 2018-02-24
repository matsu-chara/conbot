package libs

// FilterByKey filters by key
func FilterByKey(m map[string][]string, f func(string) bool) map[string][]string {
	result := map[string][]string{}
	for k := range m {
		if f(k) {
			result[k] = m[k]
		}
	}
	return result
}

// FilterByValues filter by values
// value contains only matched. eg. FilterByValues(["a"]{"1", "2", "3"}, "2") -> ["a"]{"2"}
func FilterByValues(m map[string][]string, f func([]string) ([]string, bool)) map[string][]string {
	result := map[string][]string{}
	for k := range m {
		values, isMatched := f(m[k])
		if isMatched {
			result[k] = values
		}
	}
	return result
}

// FromSlice convert string to hashmap
func FromSlice(strs []string) map[string]struct{} {
	result := map[string]struct{}{}
	for _, s := range strs {
		result[s] = struct{}{}
	}
	return result
}
