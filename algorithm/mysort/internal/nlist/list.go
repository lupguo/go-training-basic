package nlist

func GetList(n int, sorted bool) []int {
	if n < 1 {
		panic("error n")
	}
	list := make([]int, n)
	if sorted {
		for i := 0; i < n; i++ {
			list[i] = i
		}
	} else {
		for i := 0; i < n; i++ {
			list[i] = n - i
		}
	}
	return list
}

func GetConsList(n int) []int {
	if n < 1 {
		panic("error n")
	}
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = 6
	}
	return list
}
