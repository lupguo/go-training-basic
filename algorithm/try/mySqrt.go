package main

import "fmt"

// 69
//实现 int sqrt(int x) 函数。
//
// 计算并返回 x 的平方根，其中 x 是非负整数。
//
// 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
//
// 示例 1:
//
// 输入: 4
//输出: 2
//
//
// 示例 2:
//
// 输入: 8
//输出: 2
//说明: 8 的平方根是 2.82842...,
//     由于返回类型是整数，小数部分将被舍去。
//
// Related Topics 数学 二分查找

//leetcode submit region begin(Prohibit modification and deletion)
func mySqrt(x int) int {
	p := x / 2
	for {
		switch {
		case p*p > x:
			p = p / 2
		case p*p < x:
			if (p+1)*(p+1) > x {
				return p
			} else {
				p = p + 1
			}
		default:
			return p
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println("sqrt(", i, ")=", mySqrt(i))
	}
}
