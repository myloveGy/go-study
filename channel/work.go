package channel

func DigitsSum(num int64) int64 {
	var sum int64 = 0
	for num > 0 {
		sum += num % 10
		num = num / 10
	}

	return sum
}
