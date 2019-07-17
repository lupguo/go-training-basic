package main

import (
	"fmt"
	"sort"
)

func main()  {
	nums := []int{1,3,4,5,0,7,2,6,8,9}
	//sort.Ints(nums)

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	fmt.Println(nums)
}
