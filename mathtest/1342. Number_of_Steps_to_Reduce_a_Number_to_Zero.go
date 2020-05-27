package mathtest

func numberOfSteps(num int) int {
	step := 0
	for num != 0 {
		if num%2 == 0 {
			num = num / 2
			step += 1
		} else {
			num = num - 1
			step += 1
		}
	}
	return step
}

func numberOfSteps0(num int) int {
	if num == 0 {
		return 0
	}
	if num%2 == 0 {
		return 1 + numberOfSteps0(num/2)
	} else {
		return 1 + numberOfSteps0(num-1)
	}
}
