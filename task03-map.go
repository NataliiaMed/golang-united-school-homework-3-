package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	value := make([]string, 0, len(input))

	for k := 0; k < len(keys); k++ {
		num := keys[k]
		value = append(value, input[num])
	}

	return (value)

}
