package simple

func Max(list []int) int {
	if len(list) == 0 {
		return 0
	}
	max := list[0]
	for _, v := range list {
		if v > max {
			max = v
		}
	}
	return max
}


