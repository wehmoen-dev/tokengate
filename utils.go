package tokengate

import "strings"

func lowercaseKeys[T ~string](input map[T]int) map[T]int {
	output := make(map[T]int, len(input))
	for k, v := range input {
		output[T(strings.ToLower(string(k)))] = v
	}
	return output
}
