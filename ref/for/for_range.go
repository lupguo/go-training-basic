package main

import "fmt"

func main() {
	// 1. range array or slice
	fmt.Println("range array or slice:")
	var list []int
	list = append(list, 70, 80)
	for idx, val := range list {
		fmt.Println(idx, "=>", val)
	}

	// 2. range string
	fmt.Println("range string:")
	for idx, val := range "Hey,浪潮!" {
		fmt.Printf("%d => %c\n",idx,val)
	}

	// 3. range map
	fmt.Println("range map:")
	user := map[string]interface{}{
		"name": "Clark Smith",
		"age":  31,
		"like": []string{
			"pingpong",
			"swimming",
		},
	}
	for key, val := range user {
		fmt.Println(key, "=>", val)
	}

	// 4. range channel
	fmt.Println("range channel:")
	ch := make(chan int, 2)
	ch <- 1988
	ch <- 2019
	close(ch)
	for chVal := range ch {
		fmt.Println("chan val:", chVal)
	}
}
