// #9
package main

import (
	"fmt"
	"strconv"
)

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//
// 示例 1:
//
// 输入: 123
//输出: 321
//
//
// 示例 2:
//
// 输入: -123
//输出: -321
//
//
// 示例 3:
//
// 输入: 120
//输出: 21
//
//
// 注意:
//
// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231, 231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
// Related Topics 数学

//leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) int {
	var neg, ds string
	xs := strconv.Itoa(x)
	if x < 0 {
		xs, neg = xs[1:], "-"
	}
	for i := len(xs) - 1; i >= 0; i-- {
		ds += string(xs[i])
	}
	ds = neg + ds
	n, err := strconv.Atoi(ds)
	if err != nil || n < -1<<31-1 || n > 1<<31-1 {
		return 0
	}
	return n
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(reverse(1534236469), reverse(-1563847412), reverse(123), reverse(-123), reverse(120), reverse(2<<32-1), reverse(-31010))
}
