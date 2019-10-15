// #21
package main

import "fmt"

//将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
// 示例：
//
// 输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//
// Related Topics 链表

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var p = &ListNode{}
	var dp = p
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 == nil {
		p.Next = l2
	} else {
		p.Next = l1
	}
	return dp.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}

	// print
	ptlink := func(l *ListNode) {
		for k := l; k != nil; k = k.Next {
			fmt.Printf("%d ", k.Val)
		}
	}

	//ptlink(l1)
	//ptlink(l2)
	ptlink(mergeTwoLists(l1, l2))
}
