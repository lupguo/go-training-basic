package main

import "fmt"

// 每头牛活6年，每第3年、第5年，生产一头小牛，牛到第6年就会死亡
// 假定当前农场仅有1头牛，N年后农场的牛数目是多少

func getCowNum(year int, cowNum map[int]int) (num int) {
	if _, ok := cowNum[year]; ok {
		return cowNum[year]
	}

	switch {
	case year < 1:
		num = 0
	case year == 1:
		num = 1
	default:
		// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 11 12
		// a  a  a	a  a  0  0  0  0  0	 0  0
		// 		 b+	b  b  b  b  0  0  0	 0  0
		//			   c+ c  c  c  c  0	 0  0
		// 			   d+ d  d  d  d  0  0  0
		// 					 e+ e  e  e	 e  0
		//					 f+ f  f  f	 f  0
		//					 g+ g  g  g  g  0
		//					 	   h+ h  h  h
		//						   i+ i  i  i
		//						   j+ j  j  j
		//						   k+ k  k  k
		//						   l+ l  l  l
		//						   		 m+ m
		//						   		 n+ n
		//						   		 o+ o
		//						   		 p+ p
		//						   		 q+ q
		//						   		 r+ r
		//						   		 s+ s
		//						   		 t+ t
		// 每3,5,7,9,11，基数年活牛数量翻倍；每偶数年，活牛数量会减少5年前数量的一半；
		if year%2 == 1 {
			num = getCowNum(year-1, cowNum) * 2
		} else {
			num = getCowNum(year-1, cowNum) - (getCowNum(year-5, cowNum)+1)/2
		}
	}
	cowNum[year] = num
	return
}

func main() {
	cowNum := make(map[int]int)
	for n := 1; n < 100; n++ {
		fmt.Println("第", n, "年，牛的数量=", getCowNum(n, cowNum))
	}
}
