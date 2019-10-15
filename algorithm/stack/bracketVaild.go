// #20
package main

import "fmt"

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
// 注意空字符串可被认为是有效字符串。
//
// 示例 1:
//
// 输入: "()"
//输出: true
//
//
// 示例 2:
//
// 输入: "()[]{}"
//输出: true
//
//
// 示例 3:
//
// 输入: "(]"
//输出: false
//
//
// 示例 4:
//
// 输入: "([)]"
//输出: false
//
//
// 示例 5:
//
// 输入: "{[]}"
//输出: true
// Related Topics 栈 字符串

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	var bs []byte
	isMatch := func(c1, c2 byte) bool {
		return c1 == '(' && c2 == ')' ||
			c1 == '[' && c2 == ']' ||
			c1 == '{' && c2 == '}'
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			bs = append(bs, s[i])
		case ')', ']', '}':
			if len(bs) == 0 || !isMatch(bs[len(bs)-1], s[i]) {
				return false
			}
			bs = bs[:len(bs)-1]
		default:
			// unexpect character
			return false
		}
	}
	return len(bs) == 0
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	for _, v := range []string{"()]", "()", "()[]{}", "(]", "([)]", "{[]}", "{{}}()", "", "xO{}", "]"} {
		fmt.Println(v, isValid(v))
	}
}
