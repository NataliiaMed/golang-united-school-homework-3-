package homework

func average(input [15]float32) (result float32) {
	var summ float32 = 0
	for i := 0; i < len(input); i++ {
		summ += input[i]
	}
	result = summ / float32(len(input))

	return (result)
}
