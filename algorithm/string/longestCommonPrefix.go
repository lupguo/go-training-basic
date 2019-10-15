// 14
package main

import "fmt"

//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
// 示例 1:
//
// 输入: ["flower","flow","flight"]
//输出: "fl"
//
//
// 示例 2:
//
// 输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//
//
// 说明:
//
// 所有输入只包含小写字母 a-z 。
// Related Topics 字符串



//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	switch {
	case len(strs) == 0:
		return ""
	case len(strs) == 1:
		return strs[0]
	default:
		base := strs[0]
		for i, _ := range base {
			for _, s := range strs[1:]{
				if len(s) < i+1 || base[:i+1] != s[:i+1] {
					return base[:i]
				}
			}
		}
		return base
	}
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
	fmt.Println(longestCommonPrefix([]string{"dog","dig"}))
	fmt.Println(longestCommonPrefix([]string{"dog","","car"}))
	fmt.Println(longestCommonPrefix([]string{"dog","","car"}))
	fmt.Println(longestCommonPrefix([]string{}))
	fmt.Println(longestCommonPrefix([]string{"a"}))
	fmt.Println(longestCommonPrefix([]string{"a", "a"}))
}
